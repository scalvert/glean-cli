package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnnouncementsListExists(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAnnouncements()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.Error(t, err, "list should return an error since the API doesn't support it")
	assert.Contains(t, err.Error(), "does not expose a list announcements endpoint")
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
	assert.Contains(t, b.String(), "list")
}
