package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

// mockKeyring implements a file-based mock of the keyring for testing
type mockKeyring struct {
	dir         string
	err         error  // Used to simulate keyring errors
	serviceName string // Test-specific service name
}

func (m *mockKeyring) Get(service, key string) (string, error) {
	if service != m.serviceName {
		return "", keyring.ErrNotFound
	}
	if m.err != nil {
		return "", m.err
	}
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
	if service != m.serviceName {
		return fmt.Errorf("attempted to write to non-test service: %s", service)
	}
	if m.err != nil {
		return m.err
	}
	err := os.MkdirAll(m.dir, 0700)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(m.dir, service+"_"+key), []byte(value), 0600)
}

func (m *mockKeyring) Delete(service, key string) error {
	if service != m.serviceName {
		return keyring.ErrNotFound
	}
	if m.err != nil {
		return m.err
	}
	err := os.Remove(filepath.Join(m.dir, service+"_"+key))
	if err != nil {
		if os.IsNotExist(err) {
			return keyring.ErrNotFound
		}
		return err
	}
	return nil
}

func setupTestKeyring(t *testing.T) (*mockKeyring, func()) {
	t.Helper()

	// Create a temporary directory for the mock keyring
	tmpDir := t.TempDir()

	// Store the original keyring implementation and service name
	originalKeyring := keyringImpl
	originalService := ServiceName

	// Use a test-specific service name
	testService := "glean-cli-test"
	ServiceName = testService

	// Create and set up the mock keyring
	mock := &mockKeyring{
		dir:         tmpDir,
		serviceName: testService,
	}
	keyringImpl = mock

	// Return the mock and cleanup function
	return mock, func() {
		keyringImpl = originalKeyring
		ServiceName = originalService
	}
}

func setupTestConfig(t *testing.T) (string, func()) {
	t.Helper()

	// Create a temporary directory for the config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	// Store the original config path
	originalPath := ConfigPath
	ConfigPath = configPath

	// Return the config path and cleanup function
	return configPath, func() {
		ConfigPath = originalPath
	}
}

// isolateAuthState redirects config path, service name, and keyring to
// temporary locations so tests never touch real credentials.
// This mirrors authtest.IsolateAuthState but lives in the config package
// to avoid a circular import (authtest imports config).
func isolateAuthState(t *testing.T) {
	t.Helper()

	home := t.TempDir()
	t.Setenv("HOME", home)

	oldConfigPath := ConfigPath
	ConfigPath = filepath.Join(home, ".glean", "config.json")
	t.Cleanup(func() { ConfigPath = oldConfigPath })

	oldServiceName := ServiceName
	ServiceName = "glean-cli-test-isolated"
	t.Cleanup(func() { ServiceName = oldServiceName })

	keyring.MockInit()
}

func TestValidateAndTransformHost(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        string
		errContains string
		wantErr     bool
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
	// Set up both keyring and config file
	mock, cleanupKeyring := setupTestKeyring(t)
	configPath, cleanupConfig := setupTestConfig(t)
	defer cleanupKeyring()
	defer cleanupConfig()

	t.Run("save and load config with working keyring", func(t *testing.T) {
		// Save config
		err := SaveConfig("linkedin", "test-token")
		require.NoError(t, err)

		// Load config
		cfg, err := LoadConfig()
		require.NoError(t, err)

		assert.Equal(t, "linkedin-be.glean.com", cfg.GleanHost)
		assert.Equal(t, "test-token", cfg.GleanToken)

		// Verify config file was also created
		assert.FileExists(t, configPath)
	})

	t.Run("fallback to config file when keyring fails", func(t *testing.T) {
		// First save config successfully
		err := SaveConfig("linkedin", "test-token")
		require.NoError(t, err)

		// Now simulate keyring failure
		mock.err = assert.AnError

		// Load config should still work using file
		cfg, err := LoadConfig()
		require.NoError(t, err)

		assert.Equal(t, "linkedin-be.glean.com", cfg.GleanHost)
		assert.Equal(t, "test-token", cfg.GleanToken)
	})

	t.Run("clear config removes from both storages", func(t *testing.T) {
		// First save some config
		err := SaveConfig("linkedin", "test-token")
		require.NoError(t, err)

		// Reset mock error
		mock.err = nil

		// Clear config
		err = ClearConfig()
		require.NoError(t, err)

		// Verify keyring is cleared
		_, err = keyringImpl.Get(ServiceName, hostKey)
		assert.Equal(t, keyring.ErrNotFound, err)

		// Verify config file is removed
		_, err = os.Stat(configPath)
		assert.True(t, os.IsNotExist(err))

		// Load config should return empty values
		cfg, err := LoadConfig()
		require.NoError(t, err)
		assert.Empty(t, cfg.GleanHost)
		assert.Empty(t, cfg.GleanToken)
	})

	t.Run("save with both storages failing", func(t *testing.T) {
		// Simulate keyring failure
		mock.err = assert.AnError

		// Create config directory
		configDir := filepath.Dir(configPath)
		err := os.MkdirAll(configDir, 0700)
		require.NoError(t, err)

		// Remove any existing config file
		os.Remove(configPath)

		// Create a file instead of the config directory to make writes fail
		err = os.Remove(configDir)
		require.NoError(t, err)
		err = os.WriteFile(configDir, []byte("not a directory"), 0600)
		require.NoError(t, err)
		defer func() {
			// Clean up for other tests
			os.Remove(configDir)
			os.MkdirAll(configDir, 0700)
		}()

		err = SaveConfig("linkedin", "test-token")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to save config")
		assert.Contains(t, err.Error(), "keyring error")
		assert.Contains(t, err.Error(), "file error")

		// Reset mock error for other tests
		mock.err = nil
	})
}

func TestClearTokenFromStorage(t *testing.T) {
	isolateAuthState(t)

	t.Run("clears token but preserves host", func(t *testing.T) {
		require.NoError(t, SaveConfig("linkedin", "stale-api-token"))

		cfg, err := LoadConfig()
		require.NoError(t, err)
		assert.Equal(t, "stale-api-token", cfg.GleanToken)

		require.NoError(t, ClearTokenFromStorage())

		cfg, err = LoadConfig()
		require.NoError(t, err)
		assert.Empty(t, cfg.GleanToken, "token should be cleared")
		assert.Equal(t, "linkedin-be.glean.com", cfg.GleanHost, "host should be preserved")
	})

	t.Run("no-op when no token exists", func(t *testing.T) {
		// Clear state from previous subtest.
		_ = ClearConfig()

		require.NoError(t, SaveHostToFile("acme"))

		require.NoError(t, ClearTokenFromStorage())

		cfg, err := LoadConfig()
		require.NoError(t, err)
		assert.Empty(t, cfg.GleanToken)
		assert.Equal(t, "acme-be.glean.com", cfg.GleanHost)
	})
}

func TestLoadConfigEnvPriority(t *testing.T) {
	isolateAuthState(t)

	t.Run("GLEAN_API_TOKEN overrides keyring", func(t *testing.T) {
		t.Setenv("GLEAN_API_TOKEN", "env-token")
		t.Setenv("GLEAN_HOST", "env-be.glean.com")

		cfg, err := LoadConfig()
		require.NoError(t, err)
		assert.Equal(t, "env-token", cfg.GleanToken)
		assert.Equal(t, "env-be.glean.com", cfg.GleanHost)
	})

	t.Run("falls through to keyring when env vars absent", func(t *testing.T) {
		require.NoError(t, SaveConfig("linkedin", "keyring-token"))

		cfg, err := LoadConfig()
		require.NoError(t, err)
		assert.Equal(t, "keyring-token", cfg.GleanToken)
	})
}

func TestLoadConfig_EnvTokenWithKeyringHost(t *testing.T) {
	isolateAuthState(t)

	err := SaveConfig("myhost.glean.com", "")
	require.NoError(t, err)
	t.Setenv("GLEAN_API_TOKEN", "env-token")

	result, err := LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "env-token", result.GleanToken)
	assert.Equal(t, "myhost.glean.com", result.GleanHost, "host from keyring must be used even when token comes from env")
}

func TestLoadConfig_EnvHostWithFileToken(t *testing.T) {
	isolateAuthState(t)

	err := saveToFile(&Config{GleanToken: "file-token"})
	require.NoError(t, err)
	t.Setenv("GLEAN_HOST", "envhost.glean.com")

	result, err := LoadConfig()
	require.NoError(t, err)
	assert.Equal(t, "envhost.glean.com", result.GleanHost)
	assert.Equal(t, "file-token", result.GleanToken, "token from file must be used even when host comes from env")
}

func TestSaveHostToFile_DoesNotTouchKeyring(t *testing.T) {
	mock, cleanupKeyring := setupTestKeyring(t)
	_, cleanupConfig := setupTestConfig(t)
	defer cleanupKeyring()
	defer cleanupConfig()

	mock.err = assert.AnError
	require.NoError(t, SaveHostToFile("linkedin"))

	cfg, err := loadFromFile()
	require.NoError(t, err)
	assert.Equal(t, "linkedin-be.glean.com", cfg.GleanHost)
	assert.Empty(t, cfg.GleanToken)
}

func TestLoadFromFile(t *testing.T) {
	_, cleanup := setupTestConfig(t)
	defer cleanup()

	t.Run("load non-existent file returns empty config", func(t *testing.T) {
		cfg, err := loadFromFile()
		require.NoError(t, err)
		assert.Empty(t, cfg.GleanHost)
		assert.Empty(t, cfg.GleanToken)
	})

	t.Run("load invalid JSON returns error", func(t *testing.T) {
		err := os.WriteFile(ConfigPath, []byte("invalid json"), 0600)
		require.NoError(t, err)

		_, err = loadFromFile()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error parsing config file")
	})

	t.Run("load valid config file", func(t *testing.T) {
		cfg := Config{
			GleanHost:  "test-be.glean.com",
			GleanToken: "test-token",
		}
		data, err := json.MarshalIndent(cfg, "", "  ")
		require.NoError(t, err)

		err = os.WriteFile(ConfigPath, data, 0600)
		require.NoError(t, err)

		loadedCfg, err := loadFromFile()
		require.NoError(t, err)
		assert.Equal(t, cfg, *loadedCfg)
	})
}
