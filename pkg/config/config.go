package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	GleanHost  string
	GleanToken string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	host := os.Getenv("GLEAN_HOST")
	if host == "" {
		return nil, fmt.Errorf("GLEAN_HOST environment variable is required")
	}

	token := os.Getenv("GLEAN_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GLEAN_TOKEN environment variable is required")
	}

	return &Config{
		GleanHost:  host,
		GleanToken: token,
	}, nil
}
