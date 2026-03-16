package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerificationHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

func TestVerificationListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

func TestVerificationVerifyInvalidJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"verify", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestVerificationVerifyMissingJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"verify"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestVerificationRemindInvalidJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"remind", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestVerificationRemindMissingJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"remind"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}
