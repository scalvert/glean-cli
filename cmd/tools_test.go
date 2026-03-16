package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToolsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdTools()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "list")
	assert.Contains(t, b.String(), "run")
}

func TestToolsRunHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdTools()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"run", "--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "ToolsCallParameter")
}

func TestToolsListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{"tools":[]}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdTools()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

func TestToolsRunInvalidJSON(t *testing.T) {
	cmd := NewCmdTools()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"run", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestToolsRunMissingJSON(t *testing.T) {
	cmd := NewCmdTools()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"run"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}
