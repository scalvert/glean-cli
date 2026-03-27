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

func TestAgentsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

// list

func TestAgentsListDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run"})
	err := cmd.Execute()
	require.NoError(t, err)
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{}
`))
}

func TestAgentsListInvalidJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"list", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
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

func TestAgentsListFields(t *testing.T) {
	body, _ := json.Marshal(map[string]any{
		"agents": []map[string]any{
			{"agent_id": "agent-1", "name": "Research Agent", "capabilities": map[string]any{}},
			{"agent_id": "agent-2", "name": "Data Analyst", "capabilities": map[string]any{}},
		},
	})
	_, cleanup := testutils.SetupTestWithResponse(t, body)
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--fields", "agents.agent_id,agents.name"})
	err := cmd.Execute()
	require.NoError(t, err)

	out := b.String()
	assert.Contains(t, out, "agent-1")
	assert.Contains(t, out, "Research Agent")
	// capabilities should be filtered out
	assert.NotContains(t, out, "capabilities")
}

func TestAgentsListOutputText(t *testing.T) {
	body, _ := json.Marshal(map[string]any{
		"agents": []map[string]any{
			{"agent_id": "agent-1", "name": "Research Agent", "description": "Finds things", "capabilities": map[string]any{}},
		},
	})
	_, cleanup := testutils.SetupTestWithResponse(t, body)
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--output", "text"})
	err := cmd.Execute()
	require.NoError(t, err)

	out := b.String()
	assert.Contains(t, out, "agent-1")
	assert.Contains(t, out, "Research Agent")
	assert.Contains(t, out, "Finds things")
}

// get

func TestAgentsGetDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"agentId":"test-agent"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "test-agent", req["agentId"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "agentId": "test-agent"
}
`))
}

func TestAgentsGetMissingJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAgentsGetInvalidJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAgentsGetLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--json", `{"agentId":"test-agent"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// schemas

func TestAgentsSchemasDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"schemas", "--dry-run", "--json", `{"agentId":"test-agent"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "agentId": "test-agent"
}
`))
}

func TestAgentsSchemasMissingJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"schemas"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAgentsSchemasInvalidJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"schemas", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAgentsSchemasLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"schemas", "--json", `{"agentId":"test-agent"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// run

func TestAgentsRunDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"run", "--dry-run", "--json", `{"agent_id":"test-agent","messages":[]}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "agent_id": "test-agent"
}
`))
}

func TestAgentsRunMissingJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"run"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAgentsRunInvalidJSON(t *testing.T) {
	cmd := NewCmdAgents()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"run", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAgentsRunLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"run", "--json", `{"agent_id":"test-agent","messages":[]}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
