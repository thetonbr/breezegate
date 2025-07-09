/*
Package config handles loading and parsing configuration files for BreezeGate.
*/

// Package config provides configuration management for the BreezeGate Load Balancer.
package config

import (
	"encoding/json"
	"os"
)

// Backend defines the backend server's structure with its URL and health status.
type Backend struct {
	URL     string `json:"url"`
	Healthy bool   `json:"healthy"`
}

// Route defines a routing path and its associated backends.
type Route struct {
	Path     string    `json:"path"`
	Backends []Backend `json:"backends"`
}

// Domain defines the domain configurations, including its routes and TLS usage.
type Domain struct {
	DomainName string  `json:"domainName"`
	Email      string  `json:"email"`
	Routes     []Route `json:"routes"`
	UseTLS     bool    `json:"useTLS"`
}

// Config holds the global configuration settings for BreezeGate.
type Config struct {
	Port                string   `json:"port"`
	HealthCheckInterval string   `json:"healthCheckInterval"`
	Domains             []Domain `json:"domains"`
}

// LoadConfig reads the configuration file and parses it into a Config struct.
func LoadConfig(file string) (Config, error) {
	var config Config
	data, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}
