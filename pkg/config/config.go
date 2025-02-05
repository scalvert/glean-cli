package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zalando/go-keyring"
)

// MaskToken masks a token by showing only the first and last 4 characters
// and replacing the rest with asterisks.
func MaskToken(token string) string {
	if len(token) <= 8 {
		return strings.Repeat("*", len(token))
	}
	return token[:4] + strings.Repeat("*", len(token)-8) + token[len(token)-4:]
}

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

// ServiceName is the service name used for keyring operations
var ServiceName = "glean-cli"

const (
	hostKey  = "host"
	tokenKey = "token"
	emailKey = "email"
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
	GleanHost  string `json:"host"`
	GleanToken string `json:"token"`
	GleanEmail string `json:"email"`
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

// loadFromKeyring attempts to load config from the system keyring
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

// loadFromFile attempts to load config from the config file
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

// LoadConfig loads the configuration from keyring first, falling back to config file
func LoadConfig() (*Config, error) {
	// Try keyring first
	cfg := loadFromKeyring()

	// If no values were loaded from keyring, try config file
	if cfg.GleanHost == "" && cfg.GleanToken == "" && cfg.GleanEmail == "" {
		var err error
		cfg, err = loadFromFile()
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

// saveToFile saves the config to the local config file
func saveToFile(cfg *Config) error {
	if ConfigPath == "" {
		return fmt.Errorf("config path not set")
	}

	// Ensure the directory exists
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

// SaveConfig saves the configuration to both keyring and config file
func SaveConfig(host, token, email string) error {
	if host != "" {
		validHost, err := ValidateAndTransformHost(host)
		if err != nil {
			return err
		}
		host = validHost
	}

	// Try to save to keyring first
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

	// Always try to save to file as well
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
		// Only return error if both storage methods failed
		return fmt.Errorf("failed to save config: keyring error: %v, file error: %v", keyringErr, err)
	}

	return nil
}

// ClearConfig removes the configuration from both keyring and config file
func ClearConfig() error {
	var keyringErr error

	// Clear keyring
	if err := keyringImpl.Delete(ServiceName, hostKey); err != nil && err != keyring.ErrNotFound {
		keyringErr = err
	}
	if err := keyringImpl.Delete(ServiceName, tokenKey); err != nil && err != keyring.ErrNotFound {
		keyringErr = err
	}
	if err := keyringImpl.Delete(ServiceName, emailKey); err != nil && err != keyring.ErrNotFound {
		keyringErr = err
	}

	// Clear config file
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
