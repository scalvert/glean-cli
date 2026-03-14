package cmd

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// searchResponse builds a minimal Glean SearchResponse JSON body with the given document titles.
func searchResponse(titles ...string) []byte {
	type doc struct {
		Title string `json:"title"`
	}
	type result struct {
		Document doc `json:"document"`
	}
	var rs []result
	for _, title := range titles {
		rs = append(rs, result{Document: doc{Title: title}})
	}
	b, _ := json.Marshal(map[string]any{
		"results": rs,
	})
	return b
}

func TestSearchCommand_BasicQuery(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, searchResponse("Vacation Policy", "Holiday Guide"))
	defer cleanup()

	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetArgs([]string{"search", "vacation policy"})
	err := root.Execute()
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "Vacation Policy")
}

func TestSearchCommand_MissingQuery(t *testing.T) {
	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetErr(buf)
	root.SetArgs([]string{"search"})
	err := root.Execute()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "requires a query argument")
}

func TestSearchCommand_DryRun(t *testing.T) {
	// Dry-run still calls NewFromConfig before checking --dry-run, so credentials are needed.
	_, cleanup := testutils.SetupTestWithResponse(t, searchResponse())
	defer cleanup()

	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetArgs([]string{"search", "--dry-run", "test query"})
	err := root.Execute()
	require.NoError(t, err)
	var req map[string]any
	require.NoError(t, json.Unmarshal(buf.Bytes(), &req), "dry-run output must be valid JSON")
	assert.Equal(t, "test query", req["query"])
}

func TestSearchCommand_JSONPayload(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, searchResponse("Engineering Docs"))
	defer cleanup()

	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetArgs([]string{"search", "--json", `{"query":"engineering","pageSize":5}`})
	err := root.Execute()
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "Engineering Docs")
}

func TestSearchCommand_OutputNDJSON(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, searchResponse("Doc A", "Doc B"))
	defer cleanup()

	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetArgs([]string{"search", "--output", "ndjson", "test"})
	err := root.Execute()
	require.NoError(t, err)
	lines := bytes.Split(bytes.TrimSpace(buf.Bytes()), []byte("\n"))
	assert.Greater(t, len(lines), 0)
	for _, line := range lines {
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		var obj map[string]any
		assert.NoError(t, json.Unmarshal(line, &obj))
	}
}
