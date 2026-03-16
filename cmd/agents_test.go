package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAgentsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

func TestAgentsListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

func TestAgentsRunDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdAgents()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"run", "--dry-run", "--json", `{}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
