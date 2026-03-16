package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnnouncementsNoListSubcommand(t *testing.T) {
	// The Glean API has no list-announcements endpoint, so no list subcommand.
	// Attempting to run it should fail with "unknown command".
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.Error(t, err, "list must not exist — no list announcements endpoint in the API")
}

func TestAnnouncementsCreateDryRun(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{"title":"Test","startTime":"2026-01-01T00:00:00Z","endTime":"2026-01-02T00:00:00Z"}`})
	err := cmd.Execute()
	require.NoError(t, err, "announcements create --dry-run must not crash")
	assert.Contains(t, b.String(), "title")
}

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
