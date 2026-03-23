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

func TestPinsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

// list

func TestPinsListDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run"})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{}
`))
}

func TestPinsListInvalidJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"list", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestPinsListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{"pins":[]}`))

	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

// get

func TestPinsGetDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"id":"pin-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "id")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "id": "pin-123"
}
`))
}

func TestPinsGetMissingJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestPinsGetInvalidJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestPinsGetLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--json", `{"id":"pin-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// create

func TestPinsCreateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{"documentId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "doc-123", req["documentId"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "documentId": "doc-123"
}
`))
}

func TestPinsCreateMissingJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestPinsCreateInvalidJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestPinsCreateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--json", `{"documentId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// update

func TestPinsUpdateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--dry-run", "--json", `{"id":"pin-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "id")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "id": "pin-123"
}
`))
}

func TestPinsUpdateMissingJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestPinsUpdateInvalidJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestPinsUpdateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--json", `{"id":"pin-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// remove

func TestPinsRemoveDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"remove", "--dry-run", "--json", `{"id":"pin-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "id")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "id": "pin-123"
}
`))
}

func TestPinsRemoveMissingJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"remove"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestPinsRemoveInvalidJSON(t *testing.T) {
	cmd := NewCmdPins()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"remove", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestPinsRemoveLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdPins()
	cmd.SetOut(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"remove", "--json", `{"id":"pin-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
