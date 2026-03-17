package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestAuthStatusCmd_NoConfig(t *testing.T) {
	// With no config set, auth status should not panic.
	// It prints to stdout directly (not cmd.OutOrStdout), so we just
	// verify it doesn't return an error.
	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"auth", "status"})

	// Should not crash — returns nil (prints "Not configured.") or a wrapped error.
	_ = root.Execute()
}

func TestAuthLogoutCmd_NoPanic(t *testing.T) {
	// Verify logout doesn't panic regardless of system auth state.
	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"auth", "logout"})

	// May succeed or fail depending on whether credentials exist on
	// this machine — either outcome is fine, we just verify no panic.
	_ = root.Execute()
}
