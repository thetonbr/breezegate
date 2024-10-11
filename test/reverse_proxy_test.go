package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thetonbr/breezegate/internal/proxy"
)

func TestReverseProxy_ServeHTTP(t *testing.T) {
	// Create a backend server to proxy to
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from backend"))
	}))
	defer backend.Close()

	// Create a ReverseProxy instance pointing to the backend server
	rp, err := proxy.NewReverseProxy(backend.URL) // Ensure correct package usage
	if err != nil {
		t.Fatalf("Error creating reverse proxy: %s", err)
	}

	// Create a new request to the reverse proxy
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Serve the request through the reverse proxy
	rp.ServeHTTP(w, req)

	// Check the response status and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	if body != "Hello from backend" {
		t.Errorf("Expected 'Hello from backend', got '%s'", body)
	}
}
