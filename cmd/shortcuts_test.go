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

func TestShortcutsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

// list

func TestShortcutsListDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run"})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "pageSize": 0
}
`))
}

func TestShortcutsListInvalidJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"list", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestShortcutsListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{"shortcuts":[]}`))

	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

// get

func TestShortcutsGetDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"alias":"go/test"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "go/test", req["alias"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "alias": "go/test"
}
`))
}

func TestShortcutsGetMissingJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestShortcutsGetInvalidJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestShortcutsGetLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--json", `{"alias":"go/test"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// create

func TestShortcutsCreateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{"data":{"inputAlias":"go/test","destinationUrl":"https://example.com"}}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "data")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "data": {
    "inputAlias": "go/test",
    "destinationUrl": "https://example.com"
  }
}
`))
}

func TestShortcutsCreateMissingJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestShortcutsCreateInvalidJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestShortcutsCreateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--json", `{"data":{"inputAlias":"go/test","destinationUrl":"https://example.com"}}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// update

func TestShortcutsUpdateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--dry-run", "--json", `{"id":42}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "id")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "id": 42
}
`))
}

func TestShortcutsUpdateMissingJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestShortcutsUpdateInvalidJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestShortcutsUpdateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--json", `{"id":42}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// delete

func TestShortcutsDeleteDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"delete", "--dry-run", "--json", `{"id":42}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "id")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "id": 42
}
`))
}

func TestShortcutsDeleteMissingJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestShortcutsDeleteInvalidJSON(t *testing.T) {
	cmd := NewCmdShortcuts()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestShortcutsDeleteLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdShortcuts()
	cmd.SetOut(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete", "--json", `{"id":42}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
