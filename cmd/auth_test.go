package cmd

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

func TestAuthCmd_Help(t *testing.T) {
	cmd := NewCmdAuth()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--help"})

	err := cmd.Execute()
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "auth")
}

func TestAuthCmd_HasSubcommands(t *testing.T) {
	cmd := NewCmdAuth()
	subNames := make([]string, 0, len(cmd.Commands()))
	for _, sub := range cmd.Commands() {
		subNames = append(subNames, sub.Name())
	}
	assert.Contains(t, subNames, "login")
	assert.Contains(t, subNames, "logout")
	assert.Contains(t, subNames, "status")
}

func isolateAuthState(t *testing.T) {
	t.Helper()

	home := t.TempDir()
	t.Setenv("HOME", home)

	oldPath := config.ConfigPath
	config.ConfigPath = filepath.Join(home, ".glean", "config.json")
	t.Cleanup(func() { config.ConfigPath = oldPath })

	oldService := config.ServiceName
	config.ServiceName = "glean-cli-test-cmd-auth"
	t.Cleanup(func() { config.ServiceName = oldService })

	keyring.MockInit()
}

func TestAuthStatusCmd_NoConfig(t *testing.T) {
	isolateAuthState(t)

	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"auth", "status"})

	_ = root.Execute()
}

func TestAuthLogoutCmd_NoPanic(t *testing.T) {
	isolateAuthState(t)

	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"auth", "logout"})

	_ = root.Execute()
}
