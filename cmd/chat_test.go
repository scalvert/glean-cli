package cmd

import (
	"bytes"
	"testing"

	"github.com/scalvert/glean-cli/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatCommand(t *testing.T) {
	t.Run("basic chat response", func(t *testing.T) {
		// Create a response with multiple streaming messages
		response := []byte(`{"messages":[{"author":"GLEAN_AI","fragments":[{"text":"Hello"}],"hasMoreFragments":false}],"chatSessionTrackingToken":"token1"}
{"messages":[{"author":"GLEAN_AI","fragments":[{"text":"How"}],"hasMoreFragments":true}],"chatSessionTrackingToken":"token2"}
{"messages":[{"author":"GLEAN_AI","fragments":[{"text":" can"}],"hasMoreFragments":true}],"chatSessionTrackingToken":"token3"}
{"messages":[{"author":"GLEAN_AI","fragments":[{"text":" I"}],"hasMoreFragments":true}],"chatSessionTrackingToken":"token4"}
{"messages":[{"author":"GLEAN_AI","fragments":[{"text":" help?"}],"hasMoreFragments":false}],"chatSessionTrackingToken":"token5"}`)

		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"What can you do?"})

		err := cmd.Execute()
		require.NoError(t, err)

		output := b.String()
		assert.Contains(t, output, "Hello")
		assert.Contains(t, output, "How can I help?")
	})

	t.Run("chat with error response", func(t *testing.T) {
		// Create an error response that's not in the expected ChatResponse format
		response := []byte(`{"error": "Something went wrong"}\n`)
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"Test error"})

		err := cmd.Execute()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "error parsing response line")
	})

	t.Run("chat with invalid JSON response", func(t *testing.T) {
		response := []byte(`invalid json`)
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"Test invalid"})

		err := cmd.Execute()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "error parsing response line")
	})

	t.Run("chat with empty response", func(t *testing.T) {
		response := []byte(``)
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"Test empty"})

		err := cmd.Execute()
		require.NoError(t, err)
		assert.Empty(t, b.String())
	})

	t.Run("chat with timeout flag", func(t *testing.T) {
		response := []byte(`{"messages":[{"author":"GLEAN_AI","fragments":[{"text":"Quick response"}],"hasMoreFragments":false}],"chatSessionTrackingToken":"token1"}`)
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--timeout", "60000", "Test timeout"})

		err := cmd.Execute()
		require.NoError(t, err)
		assert.Contains(t, b.String(), "Quick response")
	})

	t.Run("chat with save flag disabled", func(t *testing.T) {
		response := []byte(`{"messages":[{"author":"GLEAN_AI","fragments":[{"text":"Not saved"}],"hasMoreFragments":false}],"chatSessionTrackingToken":"token1"}`)
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--save=false", "Test no save"})

		err := cmd.Execute()
		require.NoError(t, err)
		assert.Contains(t, b.String(), "Not saved")
	})
}
