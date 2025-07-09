/*
Package handlers contains HTTP handlers for the load balancer, including routing and TLS setup.
*/
package handlers

import (
	"net/http"

	"github.com/thetonbr/breezegate/internal/domain"
)

// LoadBalancerHandler represents the HTTP handler for the load balancer.
type LoadBalancerHandler struct {
	lb *domain.LoadBalancer
}

// NewLoadBalancerHandler creates a new instance of LoadBalancerHandler.
func NewLoadBalancerHandler(lb *domain.LoadBalancer) *LoadBalancerHandler {
	return &LoadBalancerHandler{lb: lb}
}

// ServeHTTP implements the HTTP handler interface. It routes the request to the appropriate backend
// based on the URL path.
func (h *LoadBalancerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Extract the URL path
	path := r.URL.Path

	// Find the appropriate backend for the path
	server := h.lb.GetBackendForPath(path)
	if server == nil {
		http.Error(w, "No healthy server available for this route", http.StatusServiceUnavailable)
		return
	}

	// Optionally add custom headers
	w.Header().Add("X-Forwarded-Path", path)
	w.Header().Add("X-Forwarded-Host", r.Host)

	// Route the request to the backend server using reverse proxy
	server.ReverseProxy().ServeHTTP(w, r)
}
