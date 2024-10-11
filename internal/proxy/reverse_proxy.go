/*
Package proxy contains utilities for handling reverse proxy operations.
*/

package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ReverseProxy encapsulates the reverse proxy that forwards requests to the backend server.
type ReverseProxy struct {
	proxy *httputil.ReverseProxy
}

// NewReverseProxy creates a new reverse proxy for the specified backend server.
func NewReverseProxy(targetURL string) (*ReverseProxy, error) {
	// Parse the backend server URL
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	// Create a reverse proxy to the backend server
	return &ReverseProxy{
		proxy: httputil.NewSingleHostReverseProxy(parsedURL),
	}, nil
}

// ServeHTTP handles HTTP requests and forwards them to the backend server via reverse proxy.
func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Optionally add custom headers
	w.Header().Add("X-Proxy", "Reverse Proxy")

	// Forward the request to the backend server using the proxy
	rp.proxy.ServeHTTP(w, r)
}
