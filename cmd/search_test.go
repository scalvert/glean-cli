package cmd

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/scalvert/glean-cli/pkg/config"
	gleanhttp "github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	basicSearchResponse = `{
		"results": [{
			"document": {
				"datasource": "confluence",
				"title": "Test Document",
				"url": "https://test.com/doc"
			}
		}]
	}`
)

func TestSearchCommand(t *testing.T) {
	t.Run("basic search", func(t *testing.T) {
		_, cleanup := testutils.SetupTestWithResponse(t, []byte(basicSearchResponse))
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdSearch()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--no-color", "test query"})

		err := cmd.Execute()
		require.NoError(t, err)

		output := b.String()
		assert.Contains(t, output, "Confluence | Test Document")
		assert.Contains(t, output, "https://test.com/doc")
	})

	t.Run("search with spell correction", func(t *testing.T) {
		response := `{
			"results": [{
				"document": {
					"datasource": "confluence",
					"title": "Test Document",
					"url": "https://test.com/doc"
				}
			}],
			"suggestedSpellCorrectedQuery": "correct query"
		}`

		_, cleanup := testutils.SetupTestWithResponse(t, []byte(response))
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdSearch()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--no-color", "test qurey"})

		err := cmd.Execute()
		require.NoError(t, err)

		output := b.String()
		assert.Contains(t, output, "Did you mean: correct query?")
	})

	t.Run("search with rewritten query", func(t *testing.T) {
		response := `{
			"results": [{
				"document": {
					"datasource": "confluence",
					"title": "Test Document",
					"url": "https://test.com/doc"
				}
			}],
			"rewrittenQuery": "rewritten query"
		}`

		_, cleanup := testutils.SetupTestWithResponse(t, []byte(response))
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdSearch()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--no-color", "test query"})

		err := cmd.Execute()
		require.NoError(t, err)

		output := b.String()
		assert.Contains(t, output, "Showing results for: rewritten query")
	})

	t.Run("search with json output", func(t *testing.T) {
		_, cleanup := testutils.SetupTestWithResponse(t, []byte(basicSearchResponse))
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdSearch()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--output", "json", "test query"})

		err := cmd.Execute()
		require.NoError(t, err)

		var result map[string]interface{}
		err = json.Unmarshal(b.Bytes(), &result)
		require.NoError(t, err)
		assert.NotNil(t, result["results"])
	})

	t.Run("search with custom template", func(t *testing.T) {
		_, cleanup := testutils.SetupTestWithResponse(t, []byte(basicSearchResponse))
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdSearch()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--no-color", "--template", "{{range .Results}}{{.Document.Title}}{{end}}", "test query"})

		err := cmd.Execute()
		require.NoError(t, err)

		assert.Equal(t, "Test Document", strings.TrimSpace(b.String()))
	})

	t.Run("search with datasource filter", func(t *testing.T) {
		response := `{
			"results": [{
				"document": {
					"datasource": "confluence",
					"title": "Test Document",
					"url": "https://test.com/doc"
				}
			}]
		}`

		_, cleanup := testutils.SetupTestWithResponse(t, []byte(response))
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdSearch()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--no-color", "--datasource", "confluence", "test query"})

		err := cmd.Execute()
		require.NoError(t, err)

		output := b.String()
		assert.Contains(t, output, "Confluence | Test Document")
	})

	t.Run("search with golink result", func(t *testing.T) {
		response := `{
			"results": [{
				"document": {
					"datasource": "nonindexedshortcut",
					"title": "Test GoLink",
					"url": "https://test.com/go/link"
				}
			}]
		}`

		_, cleanup := testutils.SetupTestWithResponse(t, []byte(response))
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdSearch()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--no-color", "test query"})

		err := cmd.Execute()
		require.NoError(t, err)

		output := b.String()
		assert.Contains(t, output, "GoLink | Test GoLink")
	})

	t.Run("search with pagination", func(t *testing.T) {
		firstResponse := `{
			"results": [{
				"document": {
					"datasource": "confluence",
					"title": "First Document",
					"url": "https://test.com/first"
				}
			}],
			"hasMoreResults": true,
			"cursor": "next-page"
		}`

		secondResponse := `{
			"results": [{
				"document": {
					"datasource": "confluence",
					"title": "Second Document",
					"url": "https://test.com/second"
				}
			}],
			"hasMoreResults": false
		}`

		cleanupConfig := testutils.SetupTestConfig(t)
		defer cleanupConfig()

		// Setup mock client with multiple responses
		mock := &testutils.MockClient{
			Responses: [][]byte{[]byte(firstResponse), []byte(secondResponse)},
		}
		origFunc := gleanhttp.NewClientFunc
		gleanhttp.NewClientFunc = func(cfg *config.Config) (gleanhttp.Client, error) {
			return mock, nil
		}
		defer func() { gleanhttp.NewClientFunc = origFunc }()

		b := bytes.NewBufferString("")
		cmd := NewCmdSearch()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--no-color", "--page-size", "1", "test query"})

		err := cmd.Execute()
		require.NoError(t, err)

		output := b.String()
		assert.Contains(t, output, "First Document")
		assert.Contains(t, output, "Press 'q' to quit")
	})
}

func TestSearch(t *testing.T) {
	//nolint:govet // Ignoring fieldalignment issues in test code
	type testCase struct {
		args     []string
		contains []string
		name     string
		wantErr  bool
	}

	tests := []testCase{
		{
			args: []string{"test query"},
			contains: []string{
				"Confluence | Test Document",
				"Did you mean: correct query?",
			},
			name:    "basic search",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, cleanup := testutils.SetupTestWithResponse(t, []byte(basicSearchResponse))
			defer cleanup()

			cmd := NewCmdSearch()
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			out := b.String()
			for _, s := range tt.contains {
				if !strings.Contains(out, s) {
					t.Errorf("Output should contain %q but got: %v", s, out)
				}
			}
		})
	}
}

func TestSearchWithPagination(t *testing.T) {
	cmd := NewCmdSearch()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"test query"})

	err := cmd.Execute()
	if err != nil {
		t.Errorf("Execute() error = %v", err)
		return
	}

	out := b.String()
	expected := []string{
		"Confluence | Test Document",
		"Press 'q' to quit",
	}

	for _, s := range expected {
		if !strings.Contains(out, s) {
			t.Errorf("Output should contain %q but got: %v", s, out)
		}
	}
}
