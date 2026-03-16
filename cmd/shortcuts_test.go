package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortcutsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

func TestShortcutsListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{"shortcuts":[]}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

func TestShortcutsCreateDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdShortcuts()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
