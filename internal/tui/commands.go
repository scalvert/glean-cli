package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gleanwork/api-client-go/models/components"
)

// slashCmd is a registered slash command shown in the autocomplete picker.
type slashCmd struct {
	name string // e.g. "/mode auto"
	desc string // one-line description
}

// allSlashCommands is the canonical list shown in the slash picker.
var allSlashCommands = []slashCmd{
	{"/clear", "Start a new session"},
	{"/mode auto", "Adaptive reasoning — default"},
	{"/mode fast", "Faster, lighter responses"},
	{"/mode advanced", "Deeper reasoning, more thorough"},
	{"/help", "Show keyboard shortcuts"},
}

// updateSlashPicker opens/refreshes/closes the slash command picker based on
// whether the textarea starts with / and has matching commands.
func (m *Model) updateSlashPicker() {
	val := m.textarea.Value()
	if !strings.HasPrefix(val, "/") || m.showFilePicker {
		if m.showSlashPicker {
			m.showSlashPicker = false
			m.slashCandidates = nil
			m.slashPickerIdx = 0
		}
		return
	}
	query := strings.ToLower(val)
	var matches []slashCmd
	for _, cmd := range allSlashCommands {
		if strings.HasPrefix(cmd.name, query) {
			matches = append(matches, cmd)
		}
	}
	m.slashCandidates = matches
	m.showSlashPicker = len(matches) > 0
	if m.slashPickerIdx >= len(matches) {
		m.slashPickerIdx = 0
	}
}

// selectSlashItem executes the highlighted slash command from the picker.
func (m *Model) selectSlashItem() (tea.Model, tea.Cmd) {
	if m.slashPickerIdx >= len(m.slashCandidates) {
		return m, nil
	}
	name := m.slashCandidates[m.slashPickerIdx].name
	m.showSlashPicker = false
	m.slashCandidates = nil
	m.slashPickerIdx = 0
	m.textarea.Reset()
	return m.handleSlashCommand(name)
}

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
	if err := m.session.AppendTurn(turn); err != nil {
		sessionLog.Log("save failed: %v", err)
	}
	if !m.conversationActive {
		m.conversationActive = true
		m.viewport.Height = m.maxViewportHeight()
	}
	m.viewport.SetContent(m.renderConversation())
	m.viewport.GotoBottom()
}
