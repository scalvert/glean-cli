package cmd

import (
	"bytes"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/gleanwork/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChatJSONPayloadSetsStreamTrue(t *testing.T) {
	fixtures := testutils.NewFixtures(t, "basic_chat_response.json")
	response := fixtures.LoadAsStream("basic_chat_response")
	_, cleanup := testutils.SetupTestWithResponse(t, response)
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdChat()
	cmd.SetOut(b)
	cmd.SetArgs([]string{
		"--json",
		`{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is 2+2?"}]}]}`,
	})
	err := cmd.Execute()
	require.NoError(t, err, "chat --json must succeed (not fail with content-type error)")
}

func TestChatCommand(t *testing.T) {
	fixtures := testutils.NewFixtures(t,
		"basic_chat_response.json",
		"chat_with_stages.json",
		"error_response.json",
		"invalid_json_response.json",
		"empty_response.json",
		"timeout_response.json",
		"save_disabled_response.json",
	)

	t.Run("basic chat response", func(t *testing.T) {
		response := fixtures.LoadAsStream("basic_chat_response")
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"What can you do?"})

		err := cmd.Execute()
		require.NoError(t, err)

		// Verify output matches snapshot
		snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`Hello

How can I help?
`))
	})

	t.Run("chat with stages", func(t *testing.T) {
		response := fixtures.LoadAsStream("chat_with_stages")
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"Test stages"})

		err := cmd.Execute()
		require.NoError(t, err)

		// Verify output matches snapshot
		snaps.MatchInlineSnapshot(t, b.String(), snaps.Inline(`The Glean Assistant is an AI-powered tool designed to enhance workplace productivity and information accessibility. It can provide a variety of information, including:

1. **Intelligent Search**: Glean understands everyday language and can search for specific documents or provide an overview of team activities.
`))
	})

	t.Run("chat with error response", func(t *testing.T) {
		// Fully invalid NDJSON is silently skipped — agent-first design means
		// we don't fail the command for malformed lines; we just produce no output.
		response := fixtures.LoadAsStream("error_response")
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"Test error"})

		err := cmd.Execute()
		require.NoError(t, err)
		assert.Empty(t, b.String())
	})

	t.Run("chat with invalid JSON response", func(t *testing.T) {
		// Invalid lines are silently skipped; valid CONTENT lines still produce output.
		response := fixtures.LoadAsStream("invalid_json_response")
		_, cleanup := testutils.SetupTestWithResponse(t, response)
		defer cleanup()

		b := bytes.NewBufferString("")
		cmd := NewCmdChat()
		cmd.SetOut(b)
		cmd.SetArgs([]string{"Test invalid"})

		err := cmd.Execute()
		require.NoError(t, err)
		// Valid messages should still produce output.
		assert.NotEmpty(t, b.String())
	})

	t.Run("chat with empty response", func(t *testing.T) {
		response := fixtures.LoadAsStream("empty_response")
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

	t.Run("chat with save flag disabled", func(t *testing.T) {
		response := fixtures.LoadAsStream("save_disabled_response")
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
