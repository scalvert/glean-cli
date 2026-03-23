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

func TestActivityHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdActivity()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

// report

func TestActivityReportDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdActivity()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"report", "--dry-run", "--json", `{"events":[]}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "events")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "events": []
}
`))
}

func TestActivityReportMissingJSON(t *testing.T) {
	cmd := NewCmdActivity()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"report"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestActivityReportInvalidJSON(t *testing.T) {
	cmd := NewCmdActivity()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"report", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestActivityReportLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdActivity()
	cmd.SetOut(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"report", "--json", `{"events":[]}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// feedback

func TestActivityFeedbackDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdActivity()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"feedback", "--dry-run", "--json", `{}`})
	err := cmd.Execute()
	require.NoError(t, err)
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "event": "",
  "trackingTokens": null
}
`))
}

func TestActivityFeedbackMissingJSON(t *testing.T) {
	cmd := NewCmdActivity()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"feedback"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestActivityFeedbackInvalidJSON(t *testing.T) {
	cmd := NewCmdActivity()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"feedback", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestActivityFeedbackLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	cmd := NewCmdActivity()
	cmd.SetOut(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"feedback", "--json", `{}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
