/*
Package domain contains the core logic for the load balancer, including routing and server management.
*/
package domain

import (
	"sync"
)

// Route represents a URL path and its associated backend servers.
type Route struct {
	Path     string
	Backends []*Server
	current  int
	mu       sync.Mutex
}

// LoadBalancer manages the routing of requests to backend servers based on defined routes.
type LoadBalancer struct {
	Routes map[string]*Route
	mu     sync.RWMutex
}

// NewLoadBalancer creates a new instance of LoadBalancer.
func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		Routes: make(map[string]*Route),
	}
}

// AddRoute adds a new route and its associated backends to the load balancer.
func (lb *LoadBalancer) AddRoute(path string, backends []*Server) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.Routes[path] = &Route{
		Path:     path,
		Backends: backends,
	}
}

// GetBackendForPath retrieves a healthy backend server for the given path using Round Robin algorithm.
func (lb *LoadBalancer) GetBackendForPath(path string) *Server {
	lb.mu.RLock()
	route, exists := lb.Routes[path]
	lb.mu.RUnlock()
	if !exists {
		return nil
	}

	route.mu.Lock()
	defer route.mu.Unlock()

	numBackends := len(route.Backends)
	if numBackends == 0 {
		return nil
	}

	for i := 0; i < numBackends; i++ {
		idx := route.current % numBackends
		server := route.Backends[idx]
		route.current++

		if server.GetHealthStatus() {
			return server
		}
	}
	return nil
}
