package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gleanwork/api-client-go/models/components"
)

// handleSlashCommand parses and executes a slash command entered in the TUI input.
// Returns without making any API call. Feedback is written to the viewport as a
// roleSystem turn.
func (m *Model) handleSlashCommand(input string) (tea.Model, tea.Cmd) {
	parts := strings.Fields(strings.TrimPrefix(strings.TrimSpace(input), "/"))
	if len(parts) == 0 {
		return m, nil
	}
	cmd := strings.ToLower(parts[0])
	args := parts[1:]

	switch cmd {
	case "clear":
		m.session = &Session{}
		m.conversationMsgs = nil
		m.chatID = nil
		m.lastErr = nil
		m.historyIdx = -1
		m.conversationActive = false
		m.viewport.Height = 1
		m.viewport.SetContent(m.renderConversation())
		m.resizeViewportToContent()

	case "mode":
		if len(args) == 0 {
			m.addSystemMessage("Usage: /mode fast | advanced | auto")
			break
		}
		switch strings.ToLower(args[0]) {
		case "fast":
			m.agentMode = components.AgentEnumFast
			m.addSystemMessage("Mode set to FAST — quicker responses, lighter reasoning")
		case "advanced":
			m.agentMode = components.AgentEnumAdvanced
			m.addSystemMessage("Mode set to ADVANCED — deeper reasoning, more thorough answers")
		case "auto":
			m.agentMode = components.AgentEnumAuto
			m.addSystemMessage("Mode set to AUTO — adapts reasoning depth to each question")
		default:
			m.addSystemMessage(fmt.Sprintf("Unknown mode %q — try: fast, advanced, auto", args[0]))
		}

	case "help":
		m.showHelp = true

	default:
		m.addSystemMessage(fmt.Sprintf("Unknown command /%s — try: /clear, /mode, /help", cmd))
	}

	return m, nil
}

// addSystemMessage appends a system-role turn to the session and refreshes the viewport.
// System turns are rendered in the viewport but never sent to the Glean API.
func (m *Model) addSystemMessage(text string) {
	turn := Turn{Role: roleSystem, Content: text}
	m.session.AppendTurn(turn)
	if !m.conversationActive {
		m.conversationActive = true
		m.viewport.Height = m.maxViewportHeight()
	}
	m.viewport.SetContent(m.renderConversation())
	m.viewport.GotoBottom()
}
