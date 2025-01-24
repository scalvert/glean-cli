package config

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

const (
	serviceName = "glean-cli"
	hostKey     = "host"
	tokenKey    = "token"
)

type Config struct {
	GleanHost  string
	GleanToken string
}

func LoadConfig() (*Config, error) {
	host, err := keyring.Get(serviceName, hostKey)
	if err != nil {
		return nil, fmt.Errorf("GLEAN_HOST not configured. Run 'glean config --host <host>'")
	}

	token, err := keyring.Get(serviceName, tokenKey)
	if err != nil {
		return nil, fmt.Errorf("GLEAN_TOKEN not configured. Run 'glean config --token <token>'")
	}

	return &Config{
		GleanHost:  host,
		GleanToken: token,
	}, nil
}

func SaveConfig(host, token string) error {
	if host != "" {
		if err := keyring.Set(serviceName, hostKey, host); err != nil {
			return fmt.Errorf("failed to save host: %w", err)
		}
	}

	if token != "" {
		if err := keyring.Set(serviceName, tokenKey, token); err != nil {
			return fmt.Errorf("failed to save token: %w", err)
		}
	}

	return nil
}

func ClearConfig() error {
	if err := keyring.Delete(serviceName, hostKey); err != nil && err != keyring.ErrNotFound {
		return fmt.Errorf("failed to clear host: %w", err)
	}
	if err := keyring.Delete(serviceName, tokenKey); err != nil && err != keyring.ErrNotFound {
		return fmt.Errorf("failed to clear token: %w", err)
	}
	return nil
}
