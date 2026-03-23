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

func TestAnswersHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

// list

func TestAnswersListDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run"})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{}
`))
}

func TestAnswersListInvalidJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"list", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAnswersListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

// get

func TestAnswersGetDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"docId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "doc-123", req["docId"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "docId": "doc-123"
}
`))
}

func TestAnswersGetMissingJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAnswersGetInvalidJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAnswersGetLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--json", `{"docId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// create

func TestAnswersCreateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	// CreateAnswerRequest has a top-level "data" field wrapping the answer content.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{"data":{}}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "data")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "data": {}
}
`))
}

func TestAnswersCreateMissingJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAnswersCreateInvalidJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAnswersCreateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--json", `{"data":{}}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// update

func TestAnswersUpdateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--dry-run", "--json", `{"docId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "doc-123", req["docId"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "id": 0,
  "docId": "doc-123"
}
`))
}

func TestAnswersUpdateMissingJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAnswersUpdateInvalidJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAnswersUpdateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--json", `{"docId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// delete

func TestAnswersDeleteDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"delete", "--dry-run", "--json", `{"docId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "doc-123", req["docId"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "id": 0,
  "docId": "doc-123"
}
`))
}

func TestAnswersDeleteMissingJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAnswersDeleteInvalidJSON(t *testing.T) {
	cmd := NewCmdAnswers()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAnswersDeleteLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdAnswers()
	cmd.SetOut(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete", "--json", `{"docId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
