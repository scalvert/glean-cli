package cmd

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/gleanwork/glean-cli/internal/testutils"
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

// list

func TestToolsListDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdTools()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run"})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{}
`))
}

func TestToolsListInvalidJSON(t *testing.T) {
	cmd := NewCmdTools()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"list", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
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

// run

func TestToolsRunDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdTools()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"run", "--dry-run", "--json", `{"name":"my-tool","parameters":{}}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "my-tool", req["name"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "name": "my-tool",
  "parameters": {}
}
`))
}

func TestToolsRunMissingJSON(t *testing.T) {
	cmd := NewCmdTools()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"run"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestToolsRunInvalidJSON(t *testing.T) {
	cmd := NewCmdTools()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"run", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestToolsRunLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdTools()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"run", "--json", `{"name":"my-tool","parameters":{}}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
