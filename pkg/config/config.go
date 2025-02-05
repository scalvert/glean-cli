// Package config manages the Glean CLI's configuration, providing secure storage
// of credentials using the system keyring with fallback to file-based storage.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zalando/go-keyring"
)

// keyringProvider defines operations for secure credential storage.
type keyringProvider interface {
	Get(service, key string) (string, error)
	Set(service, key, value string) error
	Delete(service, key string) error
}

// ServiceName is the service identifier used for keyring operations.
var ServiceName = "glean-cli"

// ConfigPath is the path to the fallback config file.
var ConfigPath string

const (
	hostKey  = "host"
	tokenKey = "token"
	emailKey = "email"
)

// Config holds the Glean API credentials and connection settings.
type Config struct {
	GleanHost  string `json:"host"`
	GleanToken string `json:"token"`
	GleanEmail string `json:"email"`
}

// MaskToken masks a token by showing only the first and last 4 characters
// and replacing the rest with asterisks.
func MaskToken(token string) string {
	if len(token) <= 8 {
		return strings.Repeat("*", len(token))
	}
	return token[:4] + strings.Repeat("*", len(token)-8) + token[len(token)-4:]
}

// ValidateAndTransformHost ensures the Glean host is in the correct format,
// transforming short names (e.g., "linkedin") to full hostnames (e.g., "linkedin-be.glean.com").
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

// LoadConfig retrieves configuration from the system keyring, falling back to
// file-based storage if keyring access fails.
func LoadConfig() (*Config, error) {
	cfg := loadFromKeyring()

	if cfg.GleanHost == "" && cfg.GleanToken == "" && cfg.GleanEmail == "" {
		var err error
		cfg, err = loadFromFile()
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

// SaveConfig stores configuration in both the system keyring and file storage.
// It returns an error only if both storage methods fail.
func SaveConfig(host, token, email string) error {
	if host != "" {
		validHost, err := ValidateAndTransformHost(host)
		if err != nil {
			return err
		}
		host = validHost
	}

	var keyringErr error
	if host != "" {
		if err := keyringImpl.Set(ServiceName, hostKey, host); err != nil {
			keyringErr = err
		}
	}
	if token != "" {
		if err := keyringImpl.Set(ServiceName, tokenKey, token); err != nil {
			keyringErr = err
		}
	}
	if email != "" {
		if err := keyringImpl.Set(ServiceName, emailKey, email); err != nil {
			keyringErr = err
		}
	}

	cfg := &Config{}
	existingCfg, err := loadFromFile()
	if err == nil {
		cfg = existingCfg
	}

	if host != "" {
		cfg.GleanHost = host
	}
	if token != "" {
		cfg.GleanToken = token
	}
	if email != "" {
		cfg.GleanEmail = email
	}

	if err := saveToFile(cfg); err != nil && keyringErr != nil {
		return fmt.Errorf("failed to save config: keyring error: %v, file error: %v", keyringErr, err)
	}

	return nil
}

// ClearConfig removes all stored configuration from both keyring and file storage.
func ClearConfig() error {
	var keyringErr error

	if err := keyringImpl.Delete(ServiceName, hostKey); err != nil && err != keyring.ErrNotFound {
		keyringErr = err
	}
	if err := keyringImpl.Delete(ServiceName, tokenKey); err != nil && err != keyring.ErrNotFound {
		keyringErr = err
	}
	if err := keyringImpl.Delete(ServiceName, emailKey); err != nil && err != keyring.ErrNotFound {
		keyringErr = err
	}

	if ConfigPath != "" {
		if err := os.Remove(ConfigPath); err != nil && !os.IsNotExist(err) {
			if keyringErr != nil {
				return fmt.Errorf("failed to clear config: keyring error: %v, file error: %v", keyringErr, err)
			}
			return fmt.Errorf("error removing config file: %w", err)
		}
	}

	if keyringErr != nil {
		return fmt.Errorf("error clearing keyring: %w", keyringErr)
	}

	return nil
}

// systemKeyring provides the default keyring implementation.
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

var keyringImpl keyringProvider = &systemKeyring{}

func init() {
	homeDir, err := os.UserHomeDir()
	if err == nil {
		ConfigPath = filepath.Join(homeDir, ".glean", "config.json")
	}
}

func loadFromKeyring() *Config {
	cfg := &Config{}

	if host, err := keyringImpl.Get(ServiceName, hostKey); err == nil {
		cfg.GleanHost = host
	}

	if token, err := keyringImpl.Get(ServiceName, tokenKey); err == nil {
		cfg.GleanToken = token
	}

	if email, err := keyringImpl.Get(ServiceName, emailKey); err == nil {
		cfg.GleanEmail = email
	}

	return cfg
}

func loadFromFile() (*Config, error) {
	if ConfigPath == "" {
		return nil, fmt.Errorf("config path not set")
	}

	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return &cfg, nil
}

func saveToFile(cfg *Config) error {
	if ConfigPath == "" {
		return fmt.Errorf("config path not set")
	}

	if err := os.MkdirAll(filepath.Dir(ConfigPath), 0700); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	if err := os.WriteFile(ConfigPath, data, 0600); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}
