package config

import (
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

// LoadConfig loads the configuration and returns error only if keyring access fails
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// Load host
	if host, err := keyring.Get(serviceName, hostKey); err == nil {
		cfg.GleanHost = host
	}

	// Load token
	if token, err := keyring.Get(serviceName, tokenKey); err == nil {
		cfg.GleanToken = token
	}

	return cfg, nil
}

func SaveConfig(host, token string) error {
	if host != "" {
		if err := keyring.Set(serviceName, hostKey, host); err != nil {
			return err
		}
	}

	if token != "" {
		if err := keyring.Set(serviceName, tokenKey, token); err != nil {
			return err
		}
	}

	return nil
}

func ClearConfig() error {
	if err := keyring.Delete(serviceName, hostKey); err != nil && err != keyring.ErrNotFound {
		return err
	}
	if err := keyring.Delete(serviceName, tokenKey); err != nil && err != keyring.ErrNotFound {
		return err
	}
	return nil
}
