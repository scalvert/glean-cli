package cmd

import (
	"bytes"
	"testing"

	"github.com/gleanwork/glean-cli/internal/auth/authtest"
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
	authtest.IsolateAuthState(t)

	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"auth", "status"})

	_ = root.Execute()
}

func TestAuthLogoutCmd_NoPanic(t *testing.T) {
	authtest.IsolateAuthState(t)

	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"auth", "logout"})

	_ = root.Execute()
}
