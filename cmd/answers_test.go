package cmd

import (
	"bytes"
	"testing"

	"github.com/gleanwork/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnswersListDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{"answers":[]}`))
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run"})

	err := cmd.Execute()
	require.NoError(t, err)
}

func TestAnswersHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdAnswers()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})

	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}
