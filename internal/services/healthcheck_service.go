/*
Package services provides various services like health checks and ACME client integration.
*/
package services

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/thetonbr/breezegate/internal/domain"
)

const (
	healthCheckTimeout = 5 * time.Second
)

// HealthCheck performs periodic health checks on a backend server at specified intervals.
func HealthCheck(server *domain.Server, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		// Perform a HEAD request to check if the server is responding
		ctx, cancel := context.WithTimeout(context.Background(), healthCheckTimeout)
		req, err := http.NewRequestWithContext(ctx, http.MethodHead, server.URL.String(), http.NoBody)
		if err != nil {
			cancel()
			log.Printf("Health check failed for %s. Marking server as unhealthy\n", server.URL.String())
			server.SetHealthStatus(false)
			continue
		}

		resp, err := http.DefaultClient.Do(req)
		cancel()
		if err != nil {
			// Mark the server as unhealthy
			log.Printf("Health check failed for %s. Marking server as unhealthy\n", server.URL.String())
			server.SetHealthStatus(false)
			continue
		}
		if resp.Body != nil {
			if closeErr := resp.Body.Close(); closeErr != nil {
				log.Printf("Error closing response body: %v", closeErr)
			}
		}
		if resp.StatusCode != http.StatusOK {
			// Mark the server as unhealthy
			log.Printf("Health check failed for %s. Marking server as unhealthy\n", server.URL.String())
			server.SetHealthStatus(false)
		} else {
			// Mark the server as healthy
			server.SetHealthStatus(true)
		}
	}
}
