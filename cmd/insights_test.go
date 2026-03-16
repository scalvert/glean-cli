package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsightsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdInsights()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

func TestInsightsGetDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdInsights()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{}`})
	err := cmd.Execute()
	require.NoError(t, err)
}

func TestInsightsGetInvalidJSON(t *testing.T) {
	cmd := NewCmdInsights()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get", "--json", "not valid json"})
	err := cmd.Execute()
	assert.Error(t, err, "invalid JSON must return error")
}

func TestInsightsGetMissingJSON(t *testing.T) {
	cmd := NewCmdInsights()
	cmd.SetErr(bytes.NewBufferString(""))
	cmd.SetArgs([]string{"get"})
	err := cmd.Execute()
	assert.Error(t, err, "missing --json must return error")
	assert.Contains(t, err.Error(), "--json is required")
}
