package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/thetonbr/breezegate/internal/domain"
	"github.com/thetonbr/breezegate/internal/services"
)

func TestHealthCheck(t *testing.T) {
	tests := []struct {
		name            string
		serverFunc      func() *httptest.Server
		expectedHealthy bool
	}{
		{
			name: "Healthy Server",
			serverFunc: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
			},
			expectedHealthy: true,
		},
		{
			name: "Unhealthy Server",
			serverFunc: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}))
			},
			expectedHealthy: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := tt.serverFunc()
			defer ts.Close()

			server := &domain.Server{
				URL:       parseURL(ts.URL),
				IsHealthy: false, // initially unhealthy
			}

			// Ensure that services.HealthCheck is being called correctly
			go services.HealthCheck(server, 1*time.Second)
			time.Sleep(2 * time.Second) // Give time for health check

			if server.GetHealthStatus() != tt.expectedHealthy {
				t.Errorf("Expected server health to be %v, but got %v", tt.expectedHealthy, server.GetHealthStatus())
			}
		})
	}
}

func parseURL(rawURL string) *url.URL {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsed
}
