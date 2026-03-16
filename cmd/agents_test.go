package cmd

import (
	"bytes"
	"testing"

	"github.com/gleanwork/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAgentsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

func TestAgentsListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

func TestAgentsRunDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"run", "--dry-run", "--json", `{}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

func TestAgentsRunInvalidJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"run", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAgentsRunMissingJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"run"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAgentsListInvalidJSON(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"list", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}
