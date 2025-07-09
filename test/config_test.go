package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/thetonbr/breezegate/internal/config"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	testConfig := config.Config{
		Port:                ":8080",
		HealthCheckInterval: "10s",
		Domains: []config.Domain{
			{
				DomainName: "example.com",
				Email:      "admin@example.com",
				UseTLS:     true,
				Routes: []config.Route{
					{
						Path: "/api",
						Backends: []config.Backend{
							{URL: "http://localhost:8081", Healthy: true},
							{URL: "http://localhost:8082", Healthy: true},
						},
					},
				},
			},
		},
	}

	configData, err := json.MarshalIndent(testConfig, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	err = os.WriteFile(configPath, configData, 0o644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test loading the config
	loadedConfig, err := config.LoadConfig(configPath)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify the loaded config
	if loadedConfig.Port != testConfig.Port {
		t.Errorf("Expected port %s, got %s", testConfig.Port, loadedConfig.Port)
	}

	if loadedConfig.HealthCheckInterval != testConfig.HealthCheckInterval {
		t.Errorf("Expected health check interval %s, got %s", testConfig.HealthCheckInterval, loadedConfig.HealthCheckInterval)
	}

	if len(loadedConfig.Domains) != len(testConfig.Domains) {
		t.Errorf("Expected %d domains, got %d", len(testConfig.Domains), len(loadedConfig.Domains))
	}

	if len(loadedConfig.Domains) > 0 {
		domain := loadedConfig.Domains[0]
		expectedDomain := testConfig.Domains[0]

		if domain.DomainName != expectedDomain.DomainName {
			t.Errorf("Expected domain name %s, got %s", expectedDomain.DomainName, domain.DomainName)
		}

		if domain.Email != expectedDomain.Email {
			t.Errorf("Expected email %s, got %s", expectedDomain.Email, domain.Email)
		}

		if domain.UseTLS != expectedDomain.UseTLS {
			t.Errorf("Expected UseTLS %v, got %v", expectedDomain.UseTLS, domain.UseTLS)
		}

		if len(domain.Routes) != len(expectedDomain.Routes) {
			t.Errorf("Expected %d routes, got %d", len(expectedDomain.Routes), len(domain.Routes))
		}
	}
}

func TestLoadConfigInvalidFile(t *testing.T) {
	// Test loading a non-existent config file
	_, err := config.LoadConfig("nonexistent.json")
	if err == nil {
		t.Error("Expected error when loading non-existent config file")
	}
}

func TestLoadConfigInvalidJSON(t *testing.T) {
	// Create a temporary file with invalid JSON
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid.json")

	err := os.WriteFile(configPath, []byte("invalid json"), 0o644)
	if err != nil {
		t.Fatalf("Failed to write invalid config file: %v", err)
	}

	_, err = config.LoadConfig(configPath)
	if err == nil {
		t.Error("Expected error when loading invalid JSON config file")
	}
}
