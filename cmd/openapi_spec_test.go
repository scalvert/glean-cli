package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/scalvert/glean-cli/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenAPISpec(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	err := os.WriteFile(testFile, []byte("GET /api/users - Get a list of users"), 0600)
	require.NoError(t, err)

	// Mock LLM response
	response := `{
		"messages": [{
			"fragments": [{
				"text": "openapi: 3.0.0\ninfo:\n  title: Glean API\n  version: 1.0.0"
			}]
		}]
	}`

	_, cleanup := testutils.SetupTestWithResponse(t, []byte(response))
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdOpenAPISpec()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"-f", testFile})

	err = cmd.Execute()
	require.NoError(t, err)

	assert.Contains(t, b.String(), "Glean API")
}
