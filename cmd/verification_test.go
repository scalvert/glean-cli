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

func TestVerificationHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

// list

func TestVerificationListDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run"})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{}
`))
}

func TestVerificationListInvalidJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"list", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestVerificationListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

// verify

func TestVerificationVerifyDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"verify", "--dry-run", "--json", `{"documentId":"doc-123"}`})
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

func TestVerificationVerifyMissingJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"verify"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestVerificationVerifyInvalidJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"verify", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestVerificationVerifyLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"verify", "--json", `{"documentId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// remind

func TestVerificationRemindDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"remind", "--dry-run", "--json", `{"documentId":"doc-123"}`})
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

func TestVerificationRemindMissingJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"remind"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestVerificationRemindInvalidJSON(t *testing.T) {
	cmd := NewCmdVerification()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"remind", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestVerificationRemindLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"remind", "--json", `{"documentId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
