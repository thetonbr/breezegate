/*
Package services provides various services like health checks and ACME client integration.
*/

package services

import (
	"log"
	"net/http"
	"time"

	"github.com/thetonbr/breezegate/internal/domain"
)

// HealthCheck performs periodic health checks on a backend server at specified intervals.
func HealthCheck(server *domain.Server, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		// Perform a HEAD request to check if the server is responding
		resp, err := http.Head(server.URL.String())
		if err != nil || resp.StatusCode != http.StatusOK {
			// Mark the server as unhealthy
			log.Printf("Health check failed for %s. Marking server as unhealthy\n", server.URL.String())
			server.SetHealthStatus(false)
		} else {
			// Mark the server as healthy
			server.SetHealthStatus(true)
		}
	}
}
