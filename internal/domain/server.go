package domain

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

// Server represents a backend server that receives traffic from the load balancer.
type Server struct {
	URL       *url.URL
	IsHealthy bool
	mu        sync.Mutex
}

// NewServer creates a new Server instance with the provided URL.
func NewServer(serverURL string) (*Server, error) {
	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}

	return &Server{
		URL:       parsedURL,
		IsHealthy: true, // Assume the server is healthy initially
	}, nil
}

// SetHealthStatus sets the health status of the server (true = healthy, false = unhealthy).
func (s *Server) SetHealthStatus(isHealthy bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.IsHealthy = isHealthy
}

// GetHealthStatus returns the current health status of the server (true = healthy, false = unhealthy).
func (s *Server) GetHealthStatus() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.IsHealthy
}

// ReverseProxy returns a reverse proxy that forwards the requests to the backend server.
func (s *Server) ReverseProxy() *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(s.URL)
}
