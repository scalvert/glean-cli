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

func TestMessagesHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdMessages()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

// get

func TestMessagesGetDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdMessages()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"idType":"THREAD_ID","id":"test-id","datasource":"SLACK"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "THREAD_ID", req["idType"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "idType": "THREAD_ID",
  "id": "test-id",
  "datasource": "SLACK"
}
`))
}

func TestMessagesGetMissingJSON(t *testing.T) {
	cmd := NewCmdMessages()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestMessagesGetInvalidJSON(t *testing.T) {
	cmd := NewCmdMessages()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestMessagesGetLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdMessages()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--json", `{"idType":"THREAD_ID","id":"test-id","datasource":"SLACK"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
