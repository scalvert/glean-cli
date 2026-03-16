package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerificationHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})
	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

func TestVerificationListLive(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()
	b := bytes.NewBufferString("")
	cmd := NewCmdVerification()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list"})
	err := cmd.Execute()
	require.NoError(t, err)
}
