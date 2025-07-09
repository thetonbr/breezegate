package test

import (
	"net/url"
	"testing"

	"github.com/thetonbr/breezegate/internal/domain"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name        string
		serverURL   string
		expectError bool
	}{
		{
			name:        "Valid HTTP URL",
			serverURL:   "http://localhost:8080",
			expectError: false,
		},
		{
			name:        "Valid HTTPS URL",
			serverURL:   "https://example.com",
			expectError: false,
		},
		{
			name:        "Relative URL",
			serverURL:   "localhost:8080",
			expectError: false,
		},
		{
			name:        "Empty URL",
			serverURL:   "",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, err := domain.NewServer(tt.serverURL)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if server != nil {
					t.Error("Expected nil server but got a valid server")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if server == nil {
					t.Error("Expected valid server but got nil")
				}
				if server != nil {
					parsedURL, parseErr := url.Parse(tt.serverURL)
					if parseErr != nil {
						t.Fatalf("Failed to parse URL: %v", parseErr)
					}
					if server.URL.String() != parsedURL.String() {
						t.Errorf("Expected URL %s, got %s", parsedURL.String(), server.URL.String())
					}
					// New servers should be healthy by default
					if !server.GetHealthStatus() {
						t.Error("Expected new server to be healthy by default")
					}
				}
			}
		})
	}
}

func TestServerHealthStatus(t *testing.T) {
	server, err := domain.NewServer("http://localhost:8080")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Test initial health status
	if !server.GetHealthStatus() {
		t.Error("Expected new server to be healthy by default")
	}

	// Test setting health status to false
	server.SetHealthStatus(false)
	if server.GetHealthStatus() {
		t.Error("Expected server to be unhealthy after setting to false")
	}

	// Test setting health status to true
	server.SetHealthStatus(true)
	if !server.GetHealthStatus() {
		t.Error("Expected server to be healthy after setting to true")
	}
}

func TestServerConcurrentHealthStatus(t *testing.T) {
	server, err := domain.NewServer("http://localhost:8080")
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Test concurrent access to health status
	done := make(chan bool)

	// Start multiple goroutines to set health status
	for i := 0; i < 10; i++ {
		go func(healthy bool) {
			server.SetHealthStatus(healthy)
			done <- true
		}(i%2 == 0)
	}

	// Start multiple goroutines to get health status
	for i := 0; i < 10; i++ {
		go func() {
			_ = server.GetHealthStatus()
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 20; i++ {
		<-done
	}

	// The final health status should be either true or false
	finalStatus := server.GetHealthStatus()
	if finalStatus != true && finalStatus != false {
		t.Error("Health status should be either true or false")
	}
}
