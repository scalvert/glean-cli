package config

import (
	"fmt"
	"strings"

	"github.com/zalando/go-keyring"
)

const (
	serviceName = "glean-cli"
	hostKey     = "host"
	tokenKey    = "token"
	emailKey    = "email"
)

type Config struct {
	GleanHost  string
	GleanToken string
	GleanEmail string
}

func ValidateAndTransformHost(host string) (string, error) {
	if !strings.Contains(host, ".") {
		return fmt.Sprintf("%s-be.glean.com", host), nil
	}

	if !strings.HasSuffix(host, ".glean.com") {
		return "", fmt.Errorf("invalid host format. Must be either 'instance' or 'instance-be.glean.com'")
	}

	if !strings.HasSuffix(strings.TrimSuffix(host, ".glean.com"), "-be") {
		return "", fmt.Errorf("invalid host format. Must end with '-be.glean.com'")
	}

	return host, nil
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	if host, err := keyring.Get(serviceName, hostKey); err == nil {
		cfg.GleanHost = host
	}

	if token, err := keyring.Get(serviceName, tokenKey); err == nil {
		cfg.GleanToken = token
	}

	if email, err := keyring.Get(serviceName, emailKey); err == nil {
		cfg.GleanEmail = email
	}

	return cfg, nil
}

func SaveConfig(host, token, email string) error {
	if host != "" {
		validHost, err := ValidateAndTransformHost(host)
		if err != nil {
			return err
		}
		if err := keyring.Set(serviceName, hostKey, validHost); err != nil {
			return err
		}
	}

	if token != "" {
		if err := keyring.Set(serviceName, tokenKey, token); err != nil {
			return err
		}
	}

	if email != "" {
		if err := keyring.Set(serviceName, emailKey, email); err != nil {
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
	if err := keyring.Delete(serviceName, emailKey); err != nil && err != keyring.ErrNotFound {
		return err
	}
	return nil
}
