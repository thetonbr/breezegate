package test

import (
	"net/url"
	"testing"

	"github.com/thetonbr/breezegate/internal/domain"
)

func newTestServer(host string, healthy bool) *domain.Server {
	u, _ := url.Parse(host)
	return &domain.Server{
		URL:       u,
		IsHealthy: healthy,
	}
}

func TestLoadBalancer_AddRoute(t *testing.T) {
	tests := []struct {
		name     string
		route    string
		servers  []*domain.Server
		expected int
	}{
		{
			name:     "Single Route",
			route:    "/api",
			servers:  []*domain.Server{newTestServer("http://localhost:8080", true)},
			expected: 1,
		},
		{
			name:     "Multiple Routes",
			route:    "/api/v1",
			servers:  []*domain.Server{newTestServer("http://localhost:8080", true), newTestServer("http://localhost:8081", true)},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := domain.NewLoadBalancer()
			lb.AddRoute(tt.route, tt.servers)
			if len(lb.Routes) != 1 {
				t.Errorf("Expected 1 route, got %d", len(lb.Routes))
			}
			if len(lb.Routes[tt.route].Backends) != tt.expected {
				t.Errorf("Expected %d backends, got %d", tt.expected, len(lb.Routes[tt.route].Backends))
			}
		})
	}
}

func TestLoadBalancer_GetBackendForPath(t *testing.T) {
	tests := []struct {
		name            string
		route           string
		servers         []*domain.Server
		expectedHealthy bool
		expectedURL     string
	}{
		{
			name:            "Single Healthy Server",
			route:           "/api",
			servers:         []*domain.Server{newTestServer("http://localhost:8080", true)},
			expectedHealthy: true,
			expectedURL:     "http://localhost:8080",
		},
		{
			name:            "One Healthy, One Unhealthy",
			route:           "/api",
			servers:         []*domain.Server{newTestServer("http://localhost:8080", false), newTestServer("http://localhost:8081", true)},
			expectedHealthy: true,
			expectedURL:     "http://localhost:8081",
		},
		{
			name:            "All Unhealthy",
			route:           "/api",
			servers:         []*domain.Server{newTestServer("http://localhost:8080", false), newTestServer("http://localhost:8081", false)},
			expectedHealthy: false,
			expectedURL:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := domain.NewLoadBalancer()
			lb.AddRoute(tt.route, tt.servers)

			server := lb.GetBackendForPath(tt.route)
			if tt.expectedHealthy && server == nil {
				t.Error("Expected a healthy server, but got nil")
			}
			if !tt.expectedHealthy && server != nil {
				t.Error("Expected no healthy server, but got one")
			}
			if server != nil && server.URL.String() != tt.expectedURL {
				t.Errorf("Expected server URL %s, but got %s", tt.expectedURL, server.URL.String())
			}
		})
	}
}
