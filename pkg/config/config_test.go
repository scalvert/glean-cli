package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

// mockKeyring implements a file-based mock of the keyring for testing
type mockKeyring struct {
	dir string
}

func (m *mockKeyring) Get(service, key string) (string, error) {
	data, err := os.ReadFile(filepath.Join(m.dir, service+"_"+key))
	if err != nil {
		if os.IsNotExist(err) {
			return "", keyring.ErrNotFound
		}
		return "", err
	}
	return string(data), nil
}

func (m *mockKeyring) Set(service, key, value string) error {
	err := os.MkdirAll(m.dir, 0700)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(m.dir, service+"_"+key), []byte(value), 0600)
}

func (m *mockKeyring) Delete(service, key string) error {
	err := os.Remove(filepath.Join(m.dir, service+"_"+key))
	if err != nil {
		if os.IsNotExist(err) {
			return keyring.ErrNotFound
		}
		return err
	}
	return nil
}

func setupTestKeyring(t *testing.T) func() {
	t.Helper()

	// Create a temporary directory for the mock keyring
	tmpDir := t.TempDir()

	// Store the original keyring implementation
	originalKeyring := keyringImpl

	// Set up the mock keyring
	keyringImpl = &mockKeyring{dir: tmpDir}

	// Return a cleanup function
	return func() {
		keyringImpl = originalKeyring
	}
}

func TestValidateAndTransformHost(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        string
		wantErr     bool
		errContains string
	}{
		{
			name:  "simple instance name",
			input: "linkedin",
			want:  "linkedin-be.glean.com",
		},
		{
			name:  "full valid hostname",
			input: "linkedin-be.glean.com",
			want:  "linkedin-be.glean.com",
		},
		{
			name:        "invalid domain",
			input:       "linkedin.example.com",
			wantErr:     true,
			errContains: "invalid host format",
		},
		{
			name:        "missing -be suffix",
			input:       "linkedin.glean.com",
			wantErr:     true,
			errContains: "invalid host format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateAndTransformHost(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestConfigPath(t *testing.T) {
	t.Run("default config path", func(t *testing.T) {
		homeDir, err := os.UserHomeDir()
		require.NoError(t, err)
		expected := filepath.Join(homeDir, ".glean", "config.json")
		assert.Equal(t, expected, ConfigPath)
	})
}

func TestConfigOperations(t *testing.T) {
	// Set up mock keyring and get cleanup function
	cleanup := setupTestKeyring(t)
	defer cleanup()

	t.Run("save and load config", func(t *testing.T) {
		// Save config
		err := SaveConfig("linkedin", "test-token", "test@example.com")
		require.NoError(t, err)

		// Load config
		cfg, err := LoadConfig()
		require.NoError(t, err)

		assert.Equal(t, "linkedin-be.glean.com", cfg.GleanHost)
		assert.Equal(t, "test-token", cfg.GleanToken)
		assert.Equal(t, "test@example.com", cfg.GleanEmail)
	})

	t.Run("clear config", func(t *testing.T) {
		// First save some config
		err := SaveConfig("linkedin", "test-token", "test@example.com")
		require.NoError(t, err)

		// Clear config
		err = ClearConfig()
		require.NoError(t, err)

		// Load config should return empty values
		cfg, err := LoadConfig()
		require.NoError(t, err)
		assert.Empty(t, cfg.GleanHost)
		assert.Empty(t, cfg.GleanToken)
		assert.Empty(t, cfg.GleanEmail)
	})

	t.Run("save invalid host", func(t *testing.T) {
		err := SaveConfig("invalid.example.com", "test-token", "test@example.com")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid host format")
	})
}
