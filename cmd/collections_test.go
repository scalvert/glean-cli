package cmd

import (
	"bytes"
	"testing"

	"github.com/gleanwork/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectionsListExists(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
	assert.NotContains(t, b.String(), "Usage:", "list should not show parent help")
}

func TestCollectionsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "list")
}

func TestCollectionsCreateDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdCollections()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{"name":"Test Collection"}`})
	err := cmd.Execute()
	require.NoError(t, err)
	assert.Contains(t, b.String(), "name")
}
