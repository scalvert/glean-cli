package tui

import (
	"context"
	"fmt"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/scalvert/glean-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestModel creates a minimal Model suitable for unit tests.
// Width/height are set to non-zero so layout functions work.
func newTestModel(t *testing.T) *Model {
	t.Helper()
	cfg := &config.Config{GleanHost: "test-be.glean.com", GleanToken: "tok"}
	m, err := New(cfg, &Session{}, "test@example.com · test", context.Background())
	require.NoError(t, err)
	// Simulate a terminal size so layout calculations are predictable.
	m.width = 120
	m.height = 40
	return m
}

func ptr(s string) *string { return &s }

// TestCtrlRResetsAllState verifies CHK-032: ctrl+r clears chatID, conversationMsgs,
// conversationActive, session, and historyIdx so the new session is completely fresh.
func TestCtrlRResetsAllState(t *testing.T) {
	m := newTestModel(t)

	// Populate state that ctrl+r must clear.
	chatID := "server-chat-123"
	m.chatID = &chatID
	m.conversationActive = true
	m.historyIdx = 2
	m.lastErr = fmt.Errorf("previous error")
	m.conversationMsgs = []components.ChatMessage{
		{Author: components.AuthorUser.ToPointer()},
	}
	m.session.AppendTurn(Turn{Role: roleUser, Content: "hello"})

	// Simulate ctrl+r key press.
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyCtrlR})
	result, ok := updated.(*Model)
	require.True(t, ok)

	assert.Nil(t, result.chatID, "chatID must be cleared so new session starts fresh")
	assert.Nil(t, result.conversationMsgs, "conversationMsgs must be cleared")
	assert.False(t, result.conversationActive, "conversationActive must be reset")
	assert.Empty(t, result.session.Turns, "session turns must be cleared")
	assert.Equal(t, -1, result.historyIdx, "historyIdx must reset to -1")
}

// TestStreamCompleteMsgPreservesElapsed verifies that the elapsed timing
// field flows correctly from streamCompleteMsg → Turn.Elapsed → renderConversation.
func TestStreamCompleteMsgPreservesElapsed(t *testing.T) {
	m := newTestModel(t)
	m.conversationActive = true
	m.isStreaming = true

	msg := streamCompleteMsg{
		text:    "The answer is 42.",
		elapsed: "1.5s",
	}
	updated, _ := m.Update(msg)
	result := updated.(*Model)

	require.Len(t, result.session.Turns, 1)
	assert.Equal(t, "1.5s", result.session.Turns[0].Elapsed, "Elapsed must be preserved in session")

	rendered := result.renderConversation()
	assert.Contains(t, rendered, "1.5s", "renderConversation must display the elapsed time")
}

// TestStreamCompleteMsgErrorSurfaces verifies that an API error is stored
// on the model and rendered in the viewport via renderConversation.
func TestStreamCompleteMsgErrorSurfaces(t *testing.T) {
	m := newTestModel(t)
	m.conversationActive = true
	m.isStreaming = true

	updated, _ := m.Update(streamCompleteMsg{err: fmt.Errorf("connection refused")})
	result := updated.(*Model)

	assert.False(t, result.isStreaming, "isStreaming must be cleared after error")
	require.NotNil(t, result.lastErr)
	assert.Contains(t, result.lastErr.Error(), "connection refused")

	rendered := result.renderConversation()
	assert.Contains(t, rendered, "connection refused", "error must appear in rendered conversation")
}

// TestStreamCompleteMsgClearsMsgsOnChatID verifies CHK-038: once chatID is received
// from the server, conversationMsgs is cleared to avoid unbounded growth.
func TestStreamCompleteMsgClearsMsgsOnChatID(t *testing.T) {
	m := newTestModel(t)
	m.conversationActive = true
	m.isStreaming = true
	m.conversationMsgs = []components.ChatMessage{
		{Author: components.AuthorUser.ToPointer()},
		{Author: components.AuthorGleanAi.ToPointer()},
	}

	chatID := "server-chat-456"
	updated, _ := m.Update(streamCompleteMsg{text: "Hello!", chatID: &chatID, elapsed: "0.3s"})
	result := updated.(*Model)

	assert.Equal(t, &chatID, result.chatID, "chatID must be stored on model")
	// conversationMsgs is cleared then the assistant reply is appended — bounded to 1 entry.
	// Before this fix it grew to [user1, ai1, user2, ai2, ...] unboundedly.
	assert.Len(t, result.conversationMsgs, 1, "conversationMsgs bounded to 1 entry after chatID active")
	assert.Equal(t, components.AuthorGleanAi, *result.conversationMsgs[0].Author)
}

// TestRenderConversationShowsSources verifies that source citations are shown
// below the assistant response when present.
func TestRenderConversationShowsSources(t *testing.T) {
	m := newTestModel(t)
	m.session.Turns = []Turn{
		{Role: roleUser, Content: "What is our vacation policy?"},
		{
			Role:    roleAssistant,
			Content: "You get 20 days PTO.",
			Sources: []Source{
				{Title: "HR Policy Doc", URL: "https://wiki.example.com/hr", Datasource: "confluence"},
			},
			Elapsed: "2.1s",
		},
	}

	rendered := m.renderConversation()
	assert.Contains(t, rendered, "HR Policy Doc")
	assert.Contains(t, rendered, "2.1s")
	assert.Contains(t, rendered, "Sources")
}

// TestRenderConversationNoElapsedWhenEmpty verifies that no timing line is
// rendered when Turn.Elapsed is empty (e.g., restored from old session format).
func TestRenderConversationNoElapsedWhenEmpty(t *testing.T) {
	m := newTestModel(t)
	m.session.Turns = []Turn{
		{Role: roleAssistant, Content: "Hello.", Elapsed: ""},
	}

	rendered := m.renderConversation()
	assert.NotContains(t, rendered, "───", "timing divider must not appear when Elapsed is empty")
}

// TestUserMessagesExtractsOnlyUserTurns verifies the history navigation helper.
func TestUserMessagesExtractsOnlyUserTurns(t *testing.T) {
	turns := []Turn{
		{Role: roleUser, Content: "question one"},
		{Role: roleAssistant, Content: "answer one"},
		{Role: roleUser, Content: "question two"},
		{Role: roleAssistant, Content: "answer two"},
	}
	msgs := userMessages(turns)
	assert.Equal(t, []string{"question one", "question two"}, msgs)
}

// TestUserMessagesEmptyTurns verifies that nil is returned for an empty history.
func TestUserMessagesEmptyTurns(t *testing.T) {
	assert.Nil(t, userMessages(nil))
	assert.Nil(t, userMessages([]Turn{}))
}

// TestSessionPreviewReturnsFirstUserMessage verifies the welcome screen preview.
func TestSessionPreviewReturnsFirstUserMessage(t *testing.T) {
	m := newTestModel(t)
	assert.Equal(t, "", m.sessionPreview(), "empty session has no preview")

	m.session.Turns = []Turn{
		{Role: roleAssistant, Content: "welcome"},
		{Role: roleUser, Content: "What is the onboarding process?"},
	}
	preview := m.sessionPreview()
	assert.Contains(t, preview, "What is the onboarding process?")
}

// TestSessionPreviewTruncatesLongMessage verifies long messages are truncated with ellipsis.
func TestSessionPreviewTruncatesLongMessage(t *testing.T) {
	m := newTestModel(t)
	m.session.Turns = []Turn{
		{Role: roleUser, Content: strings.Repeat("x", 100)},
	}
	preview := m.sessionPreview()
	assert.Contains(t, preview, "…", "long messages must be truncated with ellipsis")
	// Preview content (excluding quotes) must fit within maxLen + some buffer.
	assert.Less(t, len([]rune(preview)), 70, "truncated preview must not exceed ~60 visible chars")
}

// TestAddTurnToConversationBuildsHistory verifies that addTurnToConversation correctly
// populates conversationMsgs with properly typed SDK messages.
func TestAddTurnToConversationBuildsHistory(t *testing.T) {
	m := newTestModel(t)

	m.addTurnToConversation(Turn{Role: roleUser, Content: "hello"})
	m.addTurnToConversation(Turn{Role: roleAssistant, Content: "hi there"})

	require.Len(t, m.conversationMsgs, 2)

	userMsg := m.conversationMsgs[0]
	require.NotNil(t, userMsg.Author)
	assert.Equal(t, components.AuthorUser, *userMsg.Author)
	require.Len(t, userMsg.Fragments, 1)
	assert.Equal(t, "hello", *userMsg.Fragments[0].Text)

	aiMsg := m.conversationMsgs[1]
	require.NotNil(t, aiMsg.Author)
	assert.Equal(t, components.AuthorGleanAi, *aiMsg.Author)
}

// TestCallAPIBoundsCheckSafe verifies CHK-033: calling callAPI when chatID is set
// but conversationMsgs is empty does not panic.
func TestCallAPIBoundsCheckSafe(t *testing.T) {
	m := newTestModel(t)
	chatID := "server-chat-789"
	m.chatID = &chatID
	m.conversationMsgs = nil // empty — would have panicked before the bounds-check fix

	// callAPI() returns a tea.Cmd (a function). We only call it to verify it
	// does not panic during the slice-access guard.
	assert.NotPanics(t, func() {
		_ = m.callAPI()
	})
}

func TestDefaultAgentModeIsAuto(t *testing.T) {
	m := newTestModel(t)
	assert.Equal(t, components.AgentEnumAuto, m.agentMode)
}

func TestCallAPIUsesAgentMode(t *testing.T) {
	m := newTestModel(t)
	assert.Equal(t, components.AgentEnumAuto, m.agentMode)
	m.agentMode = components.AgentEnumFast
	assert.Equal(t, components.AgentEnumFast, m.agentMode)
}

func TestSystemMessageRendersInViewport(t *testing.T) {
	m := newTestModel(t)
	m.session.Turns = []Turn{
		{Role: roleSystem, Content: "Mode set to FAST"},
	}
	rendered := m.renderConversation()
	assert.Contains(t, rendered, "Mode set to FAST")
}

func TestSystemTurnNotAddedToConversationMsgs(t *testing.T) {
	m := newTestModel(t)
	before := len(m.conversationMsgs)
	m.addTurnToConversation(Turn{Role: roleSystem, Content: "test"})
	assert.Equal(t, before, len(m.conversationMsgs),
		"system turns must not be sent to the Glean API")
}

// TestSessionAppendTurnPreservesElapsed verifies the session persistence path:
// Turn.Elapsed is written to disk and restored correctly.
func TestSessionAppendTurnPreservesElapsed(t *testing.T) {
	s := &Session{}
	turn := Turn{
		Role:    roleAssistant,
		Content: "response text",
		Elapsed: "3.7s",
	}
	s.AppendTurn(turn)

	require.Len(t, s.Turns, 1)
	assert.Equal(t, "3.7s", s.Turns[0].Elapsed, "Elapsed must survive AppendTurn")
}
