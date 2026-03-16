package cmd

import (
	"bytes"
	"testing"

	"github.com/gleanwork/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPinsHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

func TestPinsListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{"pins":[]}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}

func TestPinsCreateDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdPins()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"create", "--dry-run", "--json", `{}`})
	err := cmd.Execute()
	require.NoError(t, err)
}
