/*
Package main is the entry point for the BreezeGate Load Balancer.
It loads configurations, initializes servers, and handles HTTP/HTTPS requests with Let's Encrypt TLS support.
*/

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/thetonbr/breezegate/internal/config"
	"github.com/thetonbr/breezegate/internal/domain"
	"github.com/thetonbr/breezegate/internal/handlers"
	"github.com/thetonbr/breezegate/internal/services"
)

// main initializes the load balancer, loads configurations, and starts the HTTP/HTTPS servers.
func main() {
	// Load configurations from the config.json file
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err.Error())
	}

	// Initialize load balancer
	lb := domain.NewLoadBalancer()

	// Add routes and backend servers
	for _, domainConfig := range cfg.Domains {
		for _, route := range domainConfig.Routes {
			var backends []*domain.Server
			for _, backend := range route.Backends {
				server, err := domain.NewServer(backend.URL)
				if err != nil {
					log.Fatalf("Error creating server: %s", err.Error())
				}
				backends = append(backends, server)

				// Start health checks for each backend server
				healthCheckInterval, _ := time.ParseDuration(cfg.HealthCheckInterval)
				go services.HealthCheck(server, healthCheckInterval)
			}
			lb.AddRoute(route.Path, backends)
		}
	}

	// Initialize load balancer handler
	lbHandler := handlers.NewLoadBalancerHandler(lb)

	// Initialize ACME client for Let's Encrypt TLS certificates
	for _, domainConfig := range cfg.Domains {
		if domainConfig.UseTLS {
			// Create ACME client for the domain
			acmeClient, err := services.NewACMEClient(domainConfig.Email, domainConfig.DomainName)
			if err != nil {
				log.Fatalf("Error initializing ACME client: %s", err.Error())
			}
			// Start HTTPS server with Let's Encrypt
			go handlers.SetupACMEAutoTLS(acmeClient, domainConfig.DomainName, lbHandler)
		} else {
			// Start HTTP server
			go func() {
				log.Printf("Starting HTTP server for domain %s on port %s", domainConfig.DomainName, cfg.Port)
				err := http.ListenAndServe(cfg.Port, lbHandler)
				if err != nil {
					log.Fatalf("Error starting HTTP server: %s\n", err.Error())
				}
			}()
		}
	}

	// Block to keep the server running
	select {}
}
