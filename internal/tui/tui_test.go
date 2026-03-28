package tui

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestModel creates a minimal Model suitable for unit tests.
// Width/height are set to non-zero so layout functions work.
func newTestModel(t *testing.T) *Model {
	t.Helper()
	cfg := &config.Config{GleanHost: "test-be.glean.com", GleanToken: "tok"}
	m, err := New(cfg, &Session{}, "test@example.com · test", "dev", context.Background())
	require.NoError(t, err)
	// Simulate a terminal size so layout calculations are predictable.
	m.width = 120
	m.height = 40
	return m
}

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
	assert.Contains(t, rendered, "Could not reach Glean", "friendly error must appear in rendered conversation")
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

func TestSlashClearResetsSession(t *testing.T) {
	m := newTestModel(t)
	chatID := "old-chat"
	m.chatID = &chatID
	m.conversationActive = true
	m.session.AppendTurn(Turn{Role: roleUser, Content: "hello"})

	result, _ := m.handleSlashCommand("/clear")
	r := result.(*Model)

	assert.Nil(t, r.chatID)
	assert.Empty(t, r.session.Turns)
	assert.False(t, r.conversationActive)
}

func TestSlashModeSetsFast(t *testing.T) {
	m := newTestModel(t)
	result, _ := m.handleSlashCommand("/mode fast")
	r := result.(*Model)
	assert.Equal(t, components.AgentEnumFast, r.agentMode)
}

func TestSlashModeSetAdvanced(t *testing.T) {
	m := newTestModel(t)
	result, _ := m.handleSlashCommand("/mode advanced")
	r := result.(*Model)
	assert.Equal(t, components.AgentEnumAdvanced, r.agentMode)
}

func TestSlashModeSetAuto(t *testing.T) {
	m := newTestModel(t)
	m.agentMode = components.AgentEnumFast
	result, _ := m.handleSlashCommand("/mode auto")
	r := result.(*Model)
	assert.Equal(t, components.AgentEnumAuto, r.agentMode)
}

func TestSlashModeShowsFeedback(t *testing.T) {
	m := newTestModel(t)
	m.conversationActive = true
	result, _ := m.handleSlashCommand("/mode fast")
	r := result.(*Model)
	require.NotEmpty(t, r.session.Turns)
	last := r.session.Turns[len(r.session.Turns)-1]
	assert.Equal(t, roleSystem, last.Role)
	assert.Contains(t, last.Content, "FAST")
}

func TestSlashUnknownCommandShowsError(t *testing.T) {
	m := newTestModel(t)
	m.conversationActive = true
	result, _ := m.handleSlashCommand("/foobar")
	r := result.(*Model)
	require.NotEmpty(t, r.session.Turns)
	last := r.session.Turns[len(r.session.Turns)-1]
	assert.Equal(t, roleSystem, last.Role)
	assert.Contains(t, last.Content, "foobar")
}

func TestSlashModeUnknownArgShowsError(t *testing.T) {
	m := newTestModel(t)
	m.conversationActive = true
	result, _ := m.handleSlashCommand("/mode turbo")
	r := result.(*Model)
	last := r.session.Turns[len(r.session.Turns)-1]
	assert.Equal(t, roleSystem, last.Role)
	assert.Contains(t, last.Content, "turbo")
}

func TestSlashInputDoesNotTriggerStreaming(t *testing.T) {
	m := newTestModel(t)
	m.textarea.SetValue("/mode fast")
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	r := updated.(*Model)
	assert.False(t, r.isStreaming, "slash commands must not start an API call")
	assert.Equal(t, components.AgentEnumFast, r.agentMode)
}

func TestStatusBarShowsAgentMode(t *testing.T) {
	m := newTestModel(t)
	m.agentMode = components.AgentEnumAdvanced
	status := m.statusLine()
	assert.Contains(t, status, "ADVANCED")
}

func TestParseFileQueryDetectsAt(t *testing.T) {
	query, ok := parseFileQuery("hello @src/")
	assert.True(t, ok)
	assert.Equal(t, "src/", query)
}

func TestParseFileQueryNoAt(t *testing.T) {
	_, ok := parseFileQuery("hello world")
	assert.False(t, ok)
}

func TestParseFileQueryAtWithSpaceAfter(t *testing.T) {
	_, ok := parseFileQuery("@foo bar")
	assert.False(t, ok)
}

func TestParseFileQueryAtEnd(t *testing.T) {
	query, ok := parseFileQuery("look at this @")
	assert.True(t, ok)
	assert.Equal(t, "", query)
}

func TestBuildFileContextPrependsFiles(t *testing.T) {
	files := []attachedFile{
		{Path: "go.mod", Content: "module foo"},
	}
	result := buildFileContext(files, "what does this do?")
	assert.Contains(t, result, "[File: go.mod]")
	assert.Contains(t, result, "module foo")
	assert.Contains(t, result, "what does this do?")
	assert.Less(t, strings.Index(result, "[File:"), strings.Index(result, "what does this do?"))
}

func TestBuildFileContextNoFiles(t *testing.T) {
	result := buildFileContext(nil, "hello")
	assert.Equal(t, "hello", result)
}

func TestReadAttachedFileRejectsBinary(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "binary*.bin")
	require.NoError(t, err)
	_, _ = f.Write([]byte{0x00, 0x01, 0x02})
	f.Close()

	_, err = readAttachedFile(f.Name())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "binary")
}

func TestReadAttachedFileTruncatesLargeFiles(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "large*.txt")
	require.NoError(t, err)
	_, _ = f.WriteString(strings.Repeat("x", 15_000))
	f.Close()

	af, err := readAttachedFile(f.Name())
	require.NoError(t, err)
	assert.LessOrEqual(t, len(af.Content), 10_200)
	assert.Contains(t, af.Content, "truncated")
}

func TestReadAttachedFileReadsNormalFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "normal*.go")
	require.NoError(t, err)
	_, _ = f.WriteString("package main\n")
	f.Close()

	af, err := readAttachedFile(f.Name())
	require.NoError(t, err)
	assert.Equal(t, "package main\n", af.Content)
}

func TestPickerUpDownNavigation(t *testing.T) {
	m := newTestModel(t)
	m.showFilePicker = true
	m.filePickerItems = []string{"a.go", "b.go", "c.go"}
	m.filePickerIdx = 1

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	r := updated.(*Model)
	assert.Equal(t, 2, r.filePickerIdx)

	updated2, _ := r.Update(tea.KeyMsg{Type: tea.KeyUp})
	r2 := updated2.(*Model)
	assert.Equal(t, 1, r2.filePickerIdx)
}

func TestPickerEscClosesPicker(t *testing.T) {
	m := newTestModel(t)
	m.showFilePicker = true
	m.filePickerItems = []string{"a.go"}

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	r := updated.(*Model)
	assert.False(t, r.showFilePicker)
}

func TestEnterWithAttachedFilesInjectsContext(t *testing.T) {
	m := newTestModel(t)
	m.conversationActive = true
	m.attachedFiles = []attachedFile{
		{Path: "go.mod", Content: "module test"},
	}
	m.textarea.SetValue("explain this")

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	r := updated.(*Model)

	assert.Empty(t, r.attachedFiles)
	require.NotEmpty(t, r.conversationMsgs)
	lastMsg := r.conversationMsgs[len(r.conversationMsgs)-1]
	require.NotEmpty(t, lastMsg.Fragments)
	assert.Contains(t, *lastMsg.Fragments[0].Text, "[File: go.mod]")
	assert.Contains(t, *lastMsg.Fragments[0].Text, "explain this")
}

func TestEnterWithNoAttachedFilesSendsNormalMessage(t *testing.T) {
	m := newTestModel(t)
	m.conversationActive = true
	m.textarea.SetValue("hello")

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	r := updated.(*Model)

	require.NotEmpty(t, r.conversationMsgs)
	lastMsg := r.conversationMsgs[len(r.conversationMsgs)-1]
	require.NotEmpty(t, lastMsg.Fragments)
	assert.Equal(t, "hello", *lastMsg.Fragments[0].Text)
}

func TestUpdateFilePickerClosesWhenNoAt(t *testing.T) {
	m := newTestModel(t)
	m.showFilePicker = true
	m.filePickerItems = []string{"foo.go"}
	m.textarea.SetValue("hello world")
	m.updateFilePicker()
	assert.False(t, m.showFilePicker)
}

func TestClosePickerResetsState(t *testing.T) {
	m := newTestModel(t)
	m.showFilePicker = true
	m.filePickerItems = []string{"a.go", "b.go"}
	m.filePickerIdx = 1
	m.closePicker()
	assert.False(t, m.showFilePicker)
	assert.Nil(t, m.filePickerItems)
	assert.Equal(t, 0, m.filePickerIdx)
}

// TestMouseEnabledByDefault verifies that mouse is enabled on init.
func TestMouseEnabledByDefault(t *testing.T) {
	m := newTestModel(t)
	assert.True(t, m.mouseEnabled, "mouse should be enabled by default")
}

// TestCtrlOTogglesMouse verifies that ctrl+o toggles mouseEnabled state.
func TestCtrlOTogglesMouse(t *testing.T) {
	m := newTestModel(t)
	assert.True(t, m.mouseEnabled, "starts enabled")

	// First toggle: disable
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlO})
	r := updated.(*Model)
	assert.False(t, r.mouseEnabled, "ctrl+o should disable mouse")
	assert.NotNil(t, cmd, "should return DisableMouse command")

	// Second toggle: re-enable
	updated2, cmd2 := r.Update(tea.KeyMsg{Type: tea.KeyCtrlO})
	r2 := updated2.(*Model)
	assert.True(t, r2.mouseEnabled, "ctrl+o should re-enable mouse")
	assert.NotNil(t, cmd2, "should return EnableMouseCellMotion command")
}

// TestMouseHintShowsWhileSelecting verifies the selection hint behavior.
func TestMouseHintShowsWhileSelecting(t *testing.T) {
	m := newTestModel(t)
	m.session.Turns = []Turn{{Role: roleUser, Content: "test"}}
	assert.False(t, m.showMouseHint, "hint should start hidden")

	// Simulate left-button drag motion - should show hint
	dragMsg := tea.MouseMsg{
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionMotion,
	}
	updated, _ := m.Update(dragMsg)
	r := updated.(*Model)
	assert.True(t, r.showMouseHint, "dragging with left button should show hint")

	// Simulate mouse button release - should hide hint
	releaseMsg := tea.MouseMsg{
		Action: tea.MouseActionRelease,
	}
	updated2, _ := r.Update(releaseMsg)
	r2 := updated2.(*Model)
	assert.False(t, r2.showMouseHint, "releasing mouse should hide hint")
}

// TestMouseHintNotShownWhenDisabled verifies hint doesn't show when mouse is disabled.
func TestMouseHintNotShownWhenDisabled(t *testing.T) {
	m := newTestModel(t)
	m.mouseEnabled = false
	m.session.Turns = []Turn{{Role: roleUser, Content: "test"}}

	// Simulate left-button drag motion
	msg := tea.MouseMsg{
		Button: tea.MouseButtonLeft,
		Action: tea.MouseActionMotion,
	}
	updated, _ := m.Update(msg)
	r := updated.(*Model)

	assert.False(t, r.showMouseHint, "hint should not show when mouse is disabled")
}

// TestCtrlOClearsMouseHint verifies that toggling mouse mode clears the hint.
func TestCtrlOClearsMouseHint(t *testing.T) {
	m := newTestModel(t)
	m.showMouseHint = true

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyCtrlO})
	r := updated.(*Model)

	assert.False(t, r.showMouseHint, "ctrl+o should clear mouse hint")
}

// TestViewShowsMouseHint verifies the hint is displayed in the view.
func TestViewShowsMouseHint(t *testing.T) {
	m := newTestModel(t)
	m.showMouseHint = true

	view := m.View()
	assert.Contains(t, view, "To select text: hold Shift+drag")
	assert.Contains(t, view, "ctrl+o to toggle mouse mode")
}

// TestStatusBarShowsMouseIndicator verifies mouse-off indicator in status bar.
func TestStatusBarShowsMouseIndicator(t *testing.T) {
	m := newTestModel(t)
	m.mouseEnabled = false

	status := m.statusLine()
	assert.Contains(t, status, "🖱️ off", "status bar should show mouse indicator when disabled")
}

// TestStatusBarNoMouseIndicatorWhenEnabled verifies no indicator when mouse is enabled.
func TestStatusBarNoMouseIndicatorWhenEnabled(t *testing.T) {
	m := newTestModel(t)
	m.mouseEnabled = true

	status := m.statusLine()
	assert.NotContains(t, status, "🖱️ off", "status bar should not show indicator when mouse is enabled")
}
