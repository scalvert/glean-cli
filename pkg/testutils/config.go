package testutils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// SetupTestConfig creates a test configuration and returns a cleanup function
func SetupTestConfig(t *testing.T) func() {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")
	configData := `{
		"glean_host": "test-company",
		"glean_token": "test-token",
		"glean_email": "test@example.com"
	}`

	err := os.WriteFile(configPath, []byte(configData), 0644)
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
