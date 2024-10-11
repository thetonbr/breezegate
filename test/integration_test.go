package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/thetonbr/breezegate/internal/domain"
	"github.com/thetonbr/breezegate/internal/handlers"
)

func TestLoadBalancer_EndToEnd(t *testing.T) {
	backend1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Response from Backend 1"))
	}))
	defer backend1.Close()

	backend2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Response from Backend 2"))
	}))
	defer backend2.Close()

	servers := []*domain.Server{
		{URL: mustParseURL(backend1.URL), IsHealthy: true},
		{URL: mustParseURL(backend2.URL), IsHealthy: true},
	}

	lb := domain.NewLoadBalancer()
	lb.AddRoute("/api", servers)

	lbHandler := handlers.NewLoadBalancerHandler(lb)

	tests := []struct {
		name        string
		path        string
		expectedRes string
	}{
		{
			name:        "Route to Backend 1",
			path:        "/api",
			expectedRes: "Response from Backend 1",
		},
		{
			name:        "Route to Backend 2",
			path:        "/api",
			expectedRes: "Response from Backend 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()

			lbHandler.ServeHTTP(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", resp.StatusCode)
			}

			body := w.Body.String()
			if body != tt.expectedRes {
				t.Errorf("Expected response %s, got %s", tt.expectedRes, body)
			}
		})
	}
}

func mustParseURL(rawURL string) *url.URL {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsed
}
