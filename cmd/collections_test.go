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

func TestCollectionsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "list")
}

// list

func TestCollectionsListDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run"})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{}
`))
}

func TestCollectionsListInvalidJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"list", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestCollectionsListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
	assert.NotContains(t, b.String(), "Usage:", "list should not show parent help")
}

// create

func TestCollectionsCreateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{"name":"Test Collection"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "Test Collection", req["name"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "name": "Test Collection"
}
`))
}

func TestCollectionsCreateMissingJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestCollectionsCreateInvalidJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestCollectionsCreateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--json", `{"name":"Test Collection"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// update

func TestCollectionsUpdateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	// EditCollectionRequest requires "name" (required field).
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--dry-run", "--json", `{"name":"Updated Collection"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "Updated Collection", req["name"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "name": "Updated Collection",
  "id": 0
}
`))
}

func TestCollectionsUpdateMissingJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestCollectionsUpdateInvalidJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestCollectionsUpdateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--json", `{"name":"Updated Collection"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// add-items

func TestCollectionsAddItemsDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"add-items", "--dry-run", "--json", `{"collectionId":1,"items":[]}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "collectionId")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "collectionId": 1
}
`))
}

func TestCollectionsAddItemsMissingJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"add-items"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestCollectionsAddItemsInvalidJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"add-items", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestCollectionsAddItemsLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"add-items", "--json", `{"collectionId":1,"items":[]}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// delete

func TestCollectionsDeleteDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	// DeleteCollectionRequest uses "ids" ([]int64), not "collectionId".
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"delete", "--dry-run", "--json", `{"ids":[1]}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "ids")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "ids": [
    1
  ]
}
`))
}

func TestCollectionsDeleteMissingJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestCollectionsDeleteInvalidJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestCollectionsDeleteLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdCollections()
	cmd.SetOut(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete", "--json", `{"ids":[1]}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// delete-item

func TestCollectionsDeleteItemDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"delete-item", "--dry-run", "--json", `{"collectionId":1,"collectionItemId":2}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "collectionId")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "collectionId": 1,
  "itemId": ""
}
`))
}

func TestCollectionsDeleteItemMissingJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete-item"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestCollectionsDeleteItemInvalidJSON(t *testing.T) {
	cmd := NewCmdCollections()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete-item", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestCollectionsDeleteItemLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdCollections()
	cmd.SetOut(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete-item", "--json", `{"collectionId":1,"collectionItemId":2}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
