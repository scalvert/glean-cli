package cmd

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/scalvert/glean-cli/internal/api"
	"github.com/scalvert/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		cupaloy.SnapshotT(t, b.String())
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
		cupaloy.SnapshotT(t, b.String())
	})

	t.Run("chat with error response", func(t *testing.T) {
		response := fixtures.LoadAsStream("error_response")
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
		response := fixtures.LoadAsStream("invalid_json_response")
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

	t.Run("chat with timeout flag", func(t *testing.T) {
		response := fixtures.LoadAsStream("timeout_response")
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

func TestStageDetection(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		expectedStage ChatStageType
		expectedInfo  bool
	}{
		{
			name:          "searching stage",
			text:          "**Searching:** for relevant documents",
			expectedStage: StageSearching,
			expectedInfo:  true,
		},
		{
			name:          "reading stage",
			text:          "**Reading:** found documents",
			expectedStage: StageReading,
			expectedInfo:  true,
		},
		{
			name:          "writing stage",
			text:          "**Writing:** the response",
			expectedStage: StageWriting,
			expectedInfo:  true,
		},
		{
			name:          "summarizing stage",
			text:          "Summarizing the information",
			expectedStage: StageSummary,
			expectedInfo:  true,
		},
		{
			name:          "alternate summarize stage",
			text:          "Summarize the gathered information",
			expectedStage: StageSummary,
			expectedInfo:  true,
		},
		{
			name:         "not a stage",
			text:         "This is a regular message",
			expectedInfo: false,
		},
		{
			name:         "empty text",
			text:         "",
			expectedInfo: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := isStage(tt.text)
			if tt.expectedInfo {
				assert.NotNil(t, info)
				assert.Equal(t, tt.expectedStage, info.stage)
				expectedDetail := strings.TrimPrefix(strings.TrimPrefix(tt.text, "**"+string(tt.expectedStage)+":**"), "**Summarize:**")
				expectedDetail = strings.TrimSpace(expectedDetail)
				assert.Equal(t, expectedDetail, info.detail)
			} else {
				assert.Nil(t, info)
			}
		})
	}
}

func TestReadingStageFormatting(t *testing.T) {
	fixtures := testutils.NewFixtures(t, "reading_stage_sources.json")
	data := fixtures.Load("reading_stage_sources")

	var fixtureData struct {
		Sources []api.StructuredResult `json:"sources"`
	}
	err := json.Unmarshal(data, &fixtureData)
	require.NoError(t, err)

	tests := []struct {
		name    string
		sources []api.StructuredResult
	}{
		{
			name:    "no sources",
			sources: nil,
		},
		{
			name:    "single source with title",
			sources: fixtureData.Sources[:1],
		},
		{
			name:    "single source without title",
			sources: fixtureData.Sources[1:2],
		},
		{
			name:    "multiple sources from same datasource",
			sources: fixtureData.Sources[2:4],
		},
		{
			name:    "multiple sources from different datasources",
			sources: []api.StructuredResult{fixtureData.Sources[0], fixtureData.Sources[1]},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatReadingStage(tt.sources)
			cupaloy.SnapshotT(t, result)
		})
	}
}

func TestStageFormatting(t *testing.T) {
	tests := []struct {
		name   string
		stage  ChatStageType
		detail string
	}{
		{
			name:   "searching stage",
			stage:  StageSearching,
			detail: "for relevant documents",
		},
		{
			name:   "reading stage",
			stage:  StageReading,
			detail: "found 5 documents",
		},
		{
			name:   "writing stage",
			stage:  StageWriting,
			detail: "the response",
		},
		{
			name:   "summarizing stage",
			stage:  StageSummary,
			detail: "the gathered information",
		},
		{
			name:   "empty detail",
			stage:  StageSearching,
			detail: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatChatStage(tt.stage, tt.detail)
			cupaloy.SnapshotT(t, result)
		})
	}
}

func TestFormatChatResponse(t *testing.T) {
	testCases := []struct {
		name     string
		response string
	}{
		{
			name:     "simple response",
			response: "Hello, how can I help?",
		},
		{
			name:     "multiline response",
			response: "Here are some points:\n1. First point\n2. Second point",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := formatChatResponse(tc.response)
			cupaloy.SnapshotT(t, result)
		})
	}
}
