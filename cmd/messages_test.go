package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessagesGetDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdMessages()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"get", "--dry-run", "--json", `{"idType":"THREAD_ID","id":"test-id","datasource":"SLACK"}`})

	err := cmd.Execute()
	require.NoError(t, err)
}

func TestMessagesHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdMessages()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})

	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}
