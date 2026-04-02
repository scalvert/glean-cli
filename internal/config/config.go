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
)

// Config holds the Glean API credentials and connection settings.
type Config struct {
	GleanHost         string `json:"host"`
	GleanToken        string `json:"token"`
	OAuthClientID     string `json:"oauth_client_id,omitempty"`
	OAuthClientSecret string `json:"oauth_client_secret,omitempty"`
}

// MaskToken masks a token by showing only the first and last 4 characters
// and replacing the rest with asterisks.
func MaskToken(token string) string {
	if len(token) <= 8 {
		return strings.Repeat("*", len(token))
	}
	return token[:4] + strings.Repeat("*", len(token)-8) + token[len(token)-4:]
}

// NormalizeHost ensures the Glean host is in the correct format,
// transforming short names (e.g., "linkedin") to full hostnames (e.g., "linkedin-be.glean.com").
// Full hostnames (containing a ".") are returned unchanged.
func NormalizeHost(host string) string {
	if !strings.Contains(host, ".") {
		return host + "-be.glean.com"
	}
	return host
}

// ValidateAndTransformHost is a compatibility wrapper around NormalizeHost.
func ValidateAndTransformHost(host string) (string, error) {
	return NormalizeHost(host), nil
}

// LoadConfig retrieves configuration using the following priority order:
//  1. Environment variables (GLEAN_API_TOKEN, GLEAN_HOST)
//  2. System keyring
//  3. ~/.glean/config.json
func LoadConfig() (*Config, error) {
	cfg := loadFromEnv()

	if cfg.GleanHost == "" || cfg.GleanToken == "" {
		keyringCfg := loadFromKeyring()
		if cfg.GleanHost == "" {
			cfg.GleanHost = keyringCfg.GleanHost
		}
		if cfg.GleanToken == "" {
			cfg.GleanToken = keyringCfg.GleanToken
		}
	}

	if cfg.GleanHost == "" || cfg.GleanToken == "" {
		fileCfg, err := loadFromFile()
		if err != nil {
			return nil, err
		}
		if cfg.GleanHost == "" {
			cfg.GleanHost = fileCfg.GleanHost
		}
		if cfg.GleanToken == "" {
			cfg.GleanToken = fileCfg.GleanToken
		}
		if cfg.OAuthClientID == "" {
			cfg.OAuthClientID = fileCfg.OAuthClientID
		}
		if cfg.OAuthClientSecret == "" {
			cfg.OAuthClientSecret = fileCfg.OAuthClientSecret
		}
	}

	return cfg, nil
}

// loadFromEnv reads config values from environment variables.
// GLEAN_API_TOKEN takes precedence over all other credential sources.
func loadFromEnv() *Config {
	cfg := &Config{}
	if v := os.Getenv("GLEAN_API_TOKEN"); v != "" {
		cfg.GleanToken = v
	}
	if v := os.Getenv("GLEAN_HOST"); v != "" {
		cfg.GleanHost = v
	}
	return cfg
}

// SaveConfig stores host and token in both the system keyring and file storage.
// It returns an error only if both storage methods fail.
func SaveConfig(host, token string) error {
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

	if err := saveToFile(cfg); err != nil && keyringErr != nil {
		return fmt.Errorf("failed to save config: keyring error: %v, file error: %v", keyringErr, err)
	}

	return nil
}

// SaveHostToFile persists only the host in ~/.glean/config.json without touching
// the system keyring. This is intended for OAuth flows where the host is not
// secret and persisting it should not trigger OS keychain prompts.
func SaveHostToFile(host string) error {
	if host != "" {
		validHost, err := ValidateAndTransformHost(host)
		if err != nil {
			return err
		}
		host = validHost
	}

	cfg := &Config{}
	existingCfg, err := loadFromFile()
	if err == nil {
		cfg = existingCfg
	}
	cfg.GleanHost = host

	return saveToFile(cfg)
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
