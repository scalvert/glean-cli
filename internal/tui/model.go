package tui

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
)

// streamDoneMsg signals that the streaming response has completed.
type streamDoneMsg struct {
	ndjson string
	err    error
}

// Model is the Bubbletea model for the glean chat TUI.
type Model struct {
	viewport    viewport.Model
	textarea    textarea.Model
	sdk         *glean.Glean
	session     *Session
	renderer    *glamour.TermRenderer
	err         error
	history     strings.Builder // rendered conversation for viewport
	streaming   strings.Builder // current in-progress response text
	width       int
	height      int
	showHelp    bool
	isStreaming  bool
}

// New creates a new TUI model.
func New(sdk *glean.Glean, session *Session) (*Model, error) {
	ta := textarea.New()
	ta.Placeholder = "Ask Glean anything… (Enter to send, Shift+Enter for newline)"
	ta.Focus()
	ta.SetWidth(80)
	ta.SetHeight(3)
	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetKeys("shift+enter")

	vp := viewport.New(80, 20)
	vp.SetContent("")

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
	)
	if err != nil {
		return nil, fmt.Errorf("glamour renderer: %w", err)
	}

	m := &Model{
		viewport:  vp,
		textarea:  ta,
		sdk:       sdk,
		session:   session,
		renderer:  renderer,
	}

	// Render saved history
	for _, turn := range session.Turns {
		m.appendTurnToHistory(turn)
	}
	m.viewport.SetContent(m.history.String())
	m.viewport.GotoBottom()

	return m, nil
}

// Init implements tea.Model.
func (m *Model) Init() tea.Cmd {
	return textarea.Blink
}

// Update implements tea.Model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		taCmd  tea.Cmd
		vpCmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		headerH := 1  // status bar
		footerH := m.textarea.Height() + 2 // input + border
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - headerH - footerH
		m.textarea.SetWidth(msg.Width - 2)
		if m.renderer != nil {
			m.renderer, _ = glamour.NewTermRenderer(
				glamour.WithAutoStyle(),
				glamour.WithWordWrap(msg.Width-4),
			)
		}

	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c" || msg.String() == "esc":
			return m, tea.Quit

		case msg.String() == "?":
			m.showHelp = !m.showHelp
			return m, nil

		case msg.String() == "ctrl+l":
			m.history.Reset()
			m.viewport.SetContent("")
			return m, nil

		case msg.String() == "ctrl+r":
			m.session = &Session{}
			m.history.Reset()
			m.viewport.SetContent("")
			return m, nil

		case msg.String() == "enter" && !m.isStreaming:
			question := strings.TrimSpace(m.textarea.Value())
			if question == "" {
				return m, nil
			}
			m.textarea.Reset()
			m.isStreaming = true

			// Append user turn to history immediately
			userTurn := Turn{Role: "user", Content: question}
			m.appendTurnToHistory(userTurn)
			m.viewport.SetContent(m.history.String())
			m.viewport.GotoBottom()

			// Start the API call in a goroutine
			return m, m.sendMessage(question)
		}

	case streamDoneMsg:
		m.isStreaming = false
		if msg.err != nil {
			m.err = msg.err
			m.appendError(msg.err)
			m.viewport.SetContent(m.history.String())
			m.viewport.GotoBottom()
			return m, nil
		}

		// Process NDJSON response
		text, sources := m.parseNDJSON(msg.ndjson)
		if text != "" {
			rendered, err := m.renderer.Render(text)
			if err != nil {
				rendered = text
			}
			aiTurn := Turn{Role: "assistant", Content: text, Sources: sources}
			m.session.AddTurn("user", m.lastUserMessage(), nil)
			m.session.AddTurn("assistant", text, sources)

			m.appendRenderedResponse(rendered, sources)
			m.viewport.SetContent(m.history.String())
			m.viewport.GotoBottom()
			_ = aiTurn // already saved via session
		}
		return m, nil
	}

	m.textarea, taCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	return m, tea.Batch(taCmd, vpCmd)
}

// View implements tea.Model.
func (m *Model) View() string {
	if m.width == 0 {
		return "Loading…"
	}

	statusText := fmt.Sprintf(" glean chat | %d turns | ? for help", len(m.session.Turns))
	if m.isStreaming {
		statusText = " glean chat | thinking… | ? for help"
	}
	status := styleStatusBar.Width(m.width).Render(statusText)

	var body string
	if m.showHelp {
		body = m.helpView()
	} else {
		body = m.viewport.View()
	}

	inputBorder := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(colorBlue)).
		Width(m.width - 2)

	return lipgloss.JoinVertical(lipgloss.Left,
		status,
		body,
		inputBorder.Render(m.textarea.View()),
	)
}

// helpView renders the keyboard shortcut overlay.
func (m *Model) helpView() string {
	shortcuts := []struct{ key, desc string }{
		{"Enter", "Send message"},
		{"Shift+Enter", "New line in input"},
		{"Ctrl+C / Esc", "Quit"},
		{"Ctrl+L", "Clear conversation"},
		{"Ctrl+R", "Start new session"},
		{"PgUp / PgDn", "Scroll history"},
		{"?", "Toggle this help"},
	}

	var sb strings.Builder
	sb.WriteString("\n  Keyboard Shortcuts\n\n")
	for _, s := range shortcuts {
		sb.WriteString(fmt.Sprintf("  %s  %s\n",
			styleHelpKey.Render(fmt.Sprintf("%-20s", s.key)),
			styleHelpDesc.Render(s.desc),
		))
	}
	return sb.String()
}

// sendMessage fires the Glean chat API and returns the NDJSON as a streamDoneMsg.
func (m *Model) sendMessage(question string) tea.Cmd {
	return func() tea.Msg {
		agentDefault := components.AgentEnumDefault
		modeDefault := components.ModeDefault
		authorUser := components.AuthorUser
		stream := true

		messages := []components.ChatMessage{
			{
				Author:      authorUser.ToPointer(),
				MessageType: components.MessageTypeContent.ToPointer(),
				Fragments:   []components.ChatMessageFragment{{Text: &question}},
			},
		}

		chatReq := components.ChatRequest{
			Messages:    messages,
			AgentConfig: &components.AgentConfig{Agent: agentDefault.ToPointer(), Mode: modeDefault.ToPointer()},
			Stream:      &stream,
		}

		resp, err := m.sdk.Client.Chat.CreateStream(context.Background(), chatReq, nil)
		if err != nil {
			return streamDoneMsg{err: err}
		}

		ndjson := ""
		if resp.ChatRequestStream != nil {
			ndjson = *resp.ChatRequestStream
		}
		return streamDoneMsg{ndjson: ndjson}
	}
}

// parseNDJSON processes the NDJSON stream string and extracts the full text + sources.
func (m *Model) parseNDJSON(ndjson string) (string, []Source) {
	var textParts []string
	var sources []Source

	for _, line := range strings.Split(ndjson, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var resp components.ChatResponse
		if err := json.Unmarshal([]byte(line), &resp); err != nil {
			continue
		}
		for _, msg := range resp.Messages {
			for _, frag := range msg.Fragments {
				if frag.Text != nil && *frag.Text != "" {
					textParts = append(textParts, *frag.Text)
				}
				for _, sr := range frag.StructuredResults {
					if sr.Document != nil {
						src := Source{}
						if sr.Document.Title != nil {
							src.Title = *sr.Document.Title
						}
						if sr.Document.URL != nil {
							src.URL = *sr.Document.URL
						}
						if sr.Document.Datasource != nil {
							src.Datasource = *sr.Document.Datasource
						} else if sr.Document.Metadata != nil && sr.Document.Metadata.Datasource != nil {
							src.Datasource = *sr.Document.Metadata.Datasource
						}
						sources = append(sources, src)
					}
				}
			}
		}
	}

	return strings.Join(textParts, ""), sources
}

// appendTurnToHistory renders a session turn into the history buffer.
func (m *Model) appendTurnToHistory(turn Turn) {
	switch turn.Role {
	case "user":
		m.history.WriteString(styleUserPrompt.Render("You: "))
		m.history.WriteString(turn.Content)
		m.history.WriteString("\n\n")
	case "assistant":
		rendered, err := m.renderer.Render(turn.Content)
		if err != nil {
			rendered = turn.Content
		}
		m.appendRenderedResponse(rendered, turn.Sources)
	}
}

// appendRenderedResponse adds a rendered AI response + source list to history.
func (m *Model) appendRenderedResponse(rendered string, sources []Source) {
	m.history.WriteString(rendered)
	if len(sources) > 0 {
		m.history.WriteString(styleSource.Render("\nSources:\n"))
		for i, s := range sources {
			title := s.Title
			if title == "" {
				title = s.URL
			}
			m.history.WriteString(styleSource.Render(fmt.Sprintf("  [%d] %s — %s\n", i+1, s.Datasource, title)))
		}
		m.history.WriteString("\n")
	}
}

func (m *Model) appendError(err error) {
	m.history.WriteString(fmt.Sprintf("\nError: %v\n\n", err))
}

func (m *Model) lastUserMessage() string {
	for i := len(m.session.Turns) - 1; i >= 0; i-- {
		if m.session.Turns[i].Role == "user" {
			return m.session.Turns[i].Content
		}
	}
	return ""
}
