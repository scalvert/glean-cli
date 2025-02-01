package testutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// SetupTestConfig creates a test configuration and returns a cleanup function
func SetupTestConfig(t *testing.T) func() {
	t.Helper()

	configDir := t.TempDir()
	configPath := filepath.Join(configDir, "config.json")

	configData := `{
		"host": "https://test.glean.com",
		"token": "test-token"
	}`

	err := os.WriteFile(configPath, []byte(configData), 0600)
	require.NoError(t, err)

	oldConfigPath := os.Getenv("GLEAN_CONFIG_PATH")
	os.Setenv("GLEAN_CONFIG_PATH", configPath)

	return func() {
		if oldConfigPath != "" {
			os.Setenv("GLEAN_CONFIG_PATH", oldConfigPath)
		} else {
			os.Unsetenv("GLEAN_CONFIG_PATH")
		}
	}
}
