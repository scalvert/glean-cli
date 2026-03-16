package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentsGetDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"docIds":[]}`})

	err := cmd.Execute()
	require.NoError(t, err)
	assert.NotEmpty(t, b.String())
}

func TestDocumentsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdDocuments()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})

	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}
