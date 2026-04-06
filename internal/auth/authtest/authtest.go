// Package authtest provides shared test helpers for auth-related tests.
package authtest

import (
	"path/filepath"
	"testing"

	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/zalando/go-keyring"
)

// IsolateAuthState redirects config, HOME, and keyring to temporary
// locations so tests never touch real credentials. All mutations are
// reverted via t.Cleanup.
func IsolateAuthState(t *testing.T) {
	t.Helper()

	home := t.TempDir()
	t.Setenv("HOME", home)

	oldConfigPath := config.ConfigPath
	config.ConfigPath = filepath.Join(home, ".glean", "config.json")
	t.Cleanup(func() { config.ConfigPath = oldConfigPath })

	oldServiceName := config.ServiceName
	config.ServiceName = "glean-cli-test-isolated"
	t.Cleanup(func() { config.ServiceName = oldServiceName })

	keyring.MockInit()
}
