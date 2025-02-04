package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zalando/go-keyring"
)

// keyringProvider defines the interface for keyring operations
type keyringProvider interface {
	Get(service, key string) (string, error)
	Set(service, key, value string) error
	Delete(service, key string) error
}

// systemKeyring implements keyringProvider using the system keyring
type systemKeyring struct{}

func (s *systemKeyring) Get(service, key string) (string, error) {
	return keyring.Get(service, key)
}

func (s *systemKeyring) Set(service, key, value string) error {
	return keyring.Set(service, key, value)
}

func (s *systemKeyring) Delete(service, key string) error {
	return keyring.Delete(service, key)
}

// keyringImpl is the current keyring implementation, can be swapped for testing
var keyringImpl keyringProvider = &systemKeyring{}

const (
	serviceName = "glean-cli"
	hostKey     = "host"
	tokenKey    = "token"
	emailKey    = "email"
)

// ConfigPath is the path to the config file. This can be overridden for testing.
var ConfigPath string

func init() {
	homeDir, err := os.UserHomeDir()
	if err == nil {
		ConfigPath = filepath.Join(homeDir, ".glean", "config.json")
	}
}

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

	if host, err := keyringImpl.Get(serviceName, hostKey); err == nil {
		cfg.GleanHost = host
	}

	if token, err := keyringImpl.Get(serviceName, tokenKey); err == nil {
		cfg.GleanToken = token
	}

	if email, err := keyringImpl.Get(serviceName, emailKey); err == nil {
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
		if err := keyringImpl.Set(serviceName, hostKey, validHost); err != nil {
			return err
		}
	}

	if token != "" {
		if err := keyringImpl.Set(serviceName, tokenKey, token); err != nil {
			return err
		}
	}

	if email != "" {
		if err := keyringImpl.Set(serviceName, emailKey, email); err != nil {
			return err
		}
	}

	return nil
}

func ClearConfig() error {
	if err := keyringImpl.Delete(serviceName, hostKey); err != nil && err != keyring.ErrNotFound {
		return err
	}
	if err := keyringImpl.Delete(serviceName, tokenKey); err != nil && err != keyring.ErrNotFound {
		return err
	}
	if err := keyringImpl.Delete(serviceName, emailKey); err != nil && err != keyring.ErrNotFound {
		return err
	}
	return nil
}
