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

func TestAnnouncementsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
	assert.Contains(t, b.String(), "create")
	// No "list" subcommand — the Glean API doesn't expose list announcements.
	assert.NotContains(t, b.String(), "list        ")
}

func TestAnnouncementsNoListSubcommand(t *testing.T) {
	// The Glean API has no list-announcements endpoint, so no list subcommand.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.Error(t, err, "list must not exist — no list announcements endpoint in the API")
}

// create

func TestAnnouncementsCreateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{"title":"Test"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "Test", req["title"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "endTime": "0001-01-01T00:00:00Z",
  "startTime": "0001-01-01T00:00:00Z",
  "title": "Test"
}
`))
}

func TestAnnouncementsCreateMissingJSON(t *testing.T) {
	cmd := NewCmdAnnouncements()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAnnouncementsCreateInvalidJSON(t *testing.T) {
	cmd := NewCmdAnnouncements()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"create", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAnnouncementsCreateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--json", `{"title":"Test"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// update

func TestAnnouncementsUpdateDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--dry-run", "--json", `{"id":42}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "id")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "endTime": "0001-01-01T00:00:00Z",
  "id": 42,
  "startTime": "0001-01-01T00:00:00Z",
  "title": ""
}
`))
}

func TestAnnouncementsUpdateMissingJSON(t *testing.T) {
	cmd := NewCmdAnnouncements()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAnnouncementsUpdateInvalidJSON(t *testing.T) {
	cmd := NewCmdAnnouncements()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"update", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAnnouncementsUpdateLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"update", "--json", `{"id":42}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// delete

func TestAnnouncementsDeleteDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
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

func TestAnnouncementsDeleteMissingJSON(t *testing.T) {
	cmd := NewCmdAnnouncements()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestAnnouncementsDeleteInvalidJSON(t *testing.T) {
	cmd := NewCmdAnnouncements()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestAnnouncementsDeleteLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdAnnouncements()
	cmd.SetOut(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"delete", "--json", `{"id":42}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
