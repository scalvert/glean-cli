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

func TestDocumentsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

// get

func TestDocumentsGetDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"documentSpecs":[{"url":"https://glean.com"}]}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "documentSpecs")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "documentSpecs": [
    {
      "url": "https://glean.com"
    }
  ]
}
`))
}

func TestDocumentsGetInvalidJSON(t *testing.T) {
	cmd := NewCmdDocuments()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestDocumentsGetLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--json", `{"documentSpecs":[{"url":"https://glean.com"}]}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// get-by-facets

func TestDocumentsGetByFacetsDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	// GetDocumentsByFacetsRequest uses "filterSets" (required), not "datasource".
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get-by-facets", "--dry-run", "--json", `{"filterSets":[]}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Contains(t, req, "filterSets")
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "filterSets": []
}
`))
}

func TestDocumentsGetByFacetsInvalidJSON(t *testing.T) {
	cmd := NewCmdDocuments()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get-by-facets", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestDocumentsGetByFacetsLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get-by-facets", "--json", `{"filterSets":[]}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// get-permissions

func TestDocumentsGetPermissionsDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	// GetDocPermissionsRequest uses "documentId", not "id".
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get-permissions", "--dry-run", "--json", `{"documentId":"doc-123"}`})
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

func TestDocumentsGetPermissionsMissingJSON(t *testing.T) {
	cmd := NewCmdDocuments()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get-permissions"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestDocumentsGetPermissionsInvalidJSON(t *testing.T) {
	cmd := NewCmdDocuments()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get-permissions", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestDocumentsGetPermissionsLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get-permissions", "--json", `{"documentId":"doc-123"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

// summarize

func TestDocumentsSummarizeDryRun(t *testing.T) {
	// Dry-run must not require auth — SDK init is deferred until after the dry-run check.
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"summarize", "--dry-run", "--json", `{"query":"What is this about?"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(b.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "What is this about?", req["query"])
	snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`{
  "documentSpecs": null,
  "query": "What is this about?"
}
`))
}

func TestDocumentsSummarizeMissingJSON(t *testing.T) {
	cmd := NewCmdDocuments()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"summarize"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}

func TestDocumentsSummarizeInvalidJSON(t *testing.T) {
	cmd := NewCmdDocuments()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"summarize", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestDocumentsSummarizeLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"summarize", "--json", `{"query":"What is this about?"}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
