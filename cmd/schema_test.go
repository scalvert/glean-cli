package cmd

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaCmd_ListAll(t *testing.T) {
	cmd := NewCmdSchema()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	require.NoError(t, err)

	var result map[string][]string
	require.NoError(t, json.Unmarshal(buf.Bytes(), &result))

	commands, ok := result["commands"]
	assert.True(t, ok, "response should have a 'commands' key")
	assert.Contains(t, commands, "search")
	assert.Contains(t, commands, "chat")
}

func TestSchemaCmd_SpecificCommand(t *testing.T) {
	cmd := NewCmdSchema()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"search"})

	err := cmd.Execute()
	require.NoError(t, err)

	var result map[string]any
	require.NoError(t, json.Unmarshal(buf.Bytes(), &result))

	assert.Equal(t, "search", result["command"])
	flags, ok := result["flags"].(map[string]any)
	assert.True(t, ok, "response should have a 'flags' map")
	assert.Contains(t, flags, "--query")
}

func TestSchemaCmd_Nonexistent(t *testing.T) {
	cmd := NewCmdSchema()
	buf := &bytes.Buffer{}
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"nonexistent"})

	err := cmd.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no schema registered")
}
