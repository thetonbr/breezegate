/*
Package handlers contains HTTP handlers for the load balancer, including routing and TLS setup.
*/

package handlers

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/thetonbr/breezegate/internal/services"
)

// SetupACMEAutoTLS configures and starts an HTTPS server with automatic TLS certificates from Let's Encrypt.
func SetupACMEAutoTLS(acmeService *services.ACMEClient, domain string, handler http.Handler) {
	cert, err := acmeService.ObtainCertificate(domain)
	if err != nil {
		log.Fatalf("Failed to obtain certificate: %s\n", err.Error())
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: handler,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*cert},
			MinVersion:   tls.VersionTLS12,
		},
	}

	log.Println("Starting HTTPS server with Let's Encrypt for domain:", domain)
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalf("Error starting HTTPS server: %s\n", err.Error())
	}
}
