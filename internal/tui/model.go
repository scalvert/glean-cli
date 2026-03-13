package tui

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
)

// streamDoneMsg signals that the API response has arrived.
type streamDoneMsg struct {
	ndjson string
	err    error
}

// Model is the Bubbletea model for the glean chat TUI.
type Model struct {
	// UI components
	viewport viewport.Model
	textarea textarea.Model
	spinner  spinner.Model
	renderer *glamour.TermRenderer

	// State
	sdk              *glean.Glean
	session          *Session
	conversationMsgs []components.ChatMessage // full history sent to SDK on each turn
	history          strings.Builder          // rendered HTML for the viewport
	width            int
	height           int
	isStreaming      bool
	showHelp         bool
}

// New creates a fully-initialized TUI model.
func New(sdk *glean.Glean, session *Session) (*Model, error) {
	ta := textarea.New()
	ta.Placeholder = "Message Glean…  (shift+enter for a new line)"
	ta.Focus()
	ta.ShowLineNumbers = false
	ta.CharLimit = 4096
	ta.KeyMap.InsertNewline.SetKeys("shift+enter")
	ta.SetHeight(3)

	vp := viewport.New(0, 0)

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = styleStatusAccent

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		// Non-fatal: fall back to plain text rendering
		renderer = nil
	}

	m := &Model{
		viewport: vp,
		textarea: ta,
		spinner:  sp,
		renderer: renderer,
		sdk:      sdk,
		session:  session,
	}

	// Replay saved session into the history buffer and build SDK message list.
	for _, turn := range session.Turns {
		m.addTurnToHistory(turn)
		m.addTurnToConversation(turn)
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
		spCmd  tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.recalculateLayout()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "?":
			m.showHelp = !m.showHelp
			return m, nil

		case "ctrl+l":
			m.history.Reset()
			m.viewport.SetContent("")
			return m, nil

		case "ctrl+r":
			m.session = &Session{}
			m.conversationMsgs = nil
			m.history.Reset()
			m.viewport.SetContent("")
			return m, nil

		case "enter":
			if m.isStreaming {
				return m, nil
			}
			question := strings.TrimSpace(m.textarea.Value())
			if question == "" {
				return m, nil
			}
			m.textarea.Reset()
			m.isStreaming = true

			// Record the user turn immediately (both visually and in session).
			turn := Turn{Role: "user", Content: question}
			m.addTurnToHistory(turn)
			m.addTurnToConversation(turn)
			m.session.AddTurn("user", question, nil)
			m.viewport.SetContent(m.history.String())
			m.viewport.GotoBottom()

			return m, tea.Batch(m.spinner.Tick, m.callAPI())
		}

	case spinner.TickMsg:
		if m.isStreaming {
			m.spinner, spCmd = m.spinner.Update(msg)
			return m, spCmd
		}

	case streamDoneMsg:
		m.isStreaming = false

		if msg.err != nil {
			m.appendError(msg.err)
			m.viewport.SetContent(m.history.String())
			m.viewport.GotoBottom()
			return m, nil
		}

		text, sources := parseNDJSON(msg.ndjson)
		if text != "" {
			rendered := m.renderMarkdown(text)
			turn := Turn{Role: "assistant", Content: text, Sources: sources}
			m.addTurnToHistory(turn)
			m.addTurnToConversation(turn)
			m.session.AddTurn("assistant", text, sources)
			_ = rendered // history already updated inside addTurnToHistory
			m.viewport.SetContent(m.history.String())
			m.viewport.GotoBottom()
		}
		return m, nil
	}

	m.textarea, taCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	return m, tea.Batch(taCmd, vpCmd)
}

// recalculateLayout updates widget sizes based on current terminal dimensions.
func (m *Model) recalculateLayout() {
	if m.width == 0 || m.height == 0 {
		return
	}

	inputH := 5  // 3 content rows + 2 border rows (rounded border)
	statusH := 1 // bottom status line
	vpH := m.height - inputH - statusH
	if vpH < 1 {
		vpH = 1
	}

	m.viewport.Width = m.width
	m.viewport.Height = vpH
	m.textarea.SetWidth(m.width - 4) // 2 border + 2 padding

	if m.renderer != nil {
		wrapWidth := m.width - 4
		if wrapWidth < 40 {
			wrapWidth = 40
		}
		if r, err := glamour.NewTermRenderer(
			glamour.WithStandardStyle("dark"),
			glamour.WithWordWrap(wrapWidth),
		); err == nil {
			m.renderer = r
		}
	}
}

// callAPI sends the full conversation history to Glean and returns a streamDoneMsg.
func (m *Model) callAPI() tea.Cmd {
	msgs := make([]components.ChatMessage, len(m.conversationMsgs))
	copy(msgs, m.conversationMsgs)

	sdk := m.sdk
	return func() tea.Msg {
		agentDefault := components.AgentEnumDefault
		modeDefault := components.ModeDefault
		stream := true

		chatReq := components.ChatRequest{
			Messages:    msgs,
			AgentConfig: &components.AgentConfig{Agent: agentDefault.ToPointer(), Mode: modeDefault.ToPointer()},
			Stream:      &stream,
		}

		resp, err := sdk.Client.Chat.CreateStream(context.Background(), chatReq, nil)
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

// addTurnToHistory appends a rendered turn to the viewport history buffer.
func (m *Model) addTurnToHistory(turn Turn) {
	switch turn.Role {
	case "user":
		m.history.WriteString("\n")
		m.history.WriteString(styleUserLabel.Render("  > "))
		m.history.WriteString(styleUserText.Render(turn.Content))
		m.history.WriteString("\n\n")

	case "assistant":
		rendered := m.renderMarkdown(turn.Content)
		m.history.WriteString(rendered)
		if len(turn.Sources) > 0 {
			m.history.WriteString(styleSourceHeader.Render("  Sources\n"))
			for i, s := range turn.Sources {
				title := s.Title
				if title == "" {
					title = s.URL
				}
				ds := s.Datasource
				if ds == "" {
					ds = "glean"
				}
				m.history.WriteString(styleSourceItem.Render(
					"  " + strings.Repeat("─", 2) + " [" + itoa(i+1) + "] " + ds + ": " + title + "\n",
				))
			}
			m.history.WriteString("\n")
		}
	}
}

// addTurnToConversation appends an SDK ChatMessage to the conversation history
// that will be sent on the next API call (enables multi-turn context).
func (m *Model) addTurnToConversation(turn Turn) {
	switch turn.Role {
	case "user":
		text := turn.Content
		m.conversationMsgs = append(m.conversationMsgs, components.ChatMessage{
			Author:      components.AuthorUser.ToPointer(),
			MessageType: components.MessageTypeContent.ToPointer(),
			Fragments:   []components.ChatMessageFragment{{Text: &text}},
		})
	case "assistant":
		text := turn.Content
		m.conversationMsgs = append(m.conversationMsgs, components.ChatMessage{
			Author:      components.AuthorGleanAi.ToPointer(),
			MessageType: components.MessageTypeContent.ToPointer(),
			Fragments:   []components.ChatMessageFragment{{Text: &text}},
		})
	}
}

// renderMarkdown renders text using Glamour, falling back to plain text.
func (m *Model) renderMarkdown(text string) string {
	if m.renderer == nil {
		return text + "\n"
	}
	rendered, err := m.renderer.Render(text)
	if err != nil {
		return text + "\n"
	}
	return rendered
}

func (m *Model) appendError(err error) {
	m.history.WriteString("\n")
	m.history.WriteString(styleError.Render("  Error: " + err.Error()))
	m.history.WriteString("\n\n")
}

// parseNDJSON extracts the full text and source citations from a buffered NDJSON response.
func parseNDJSON(ndjson string) (string, []Source) {
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
					if sr.Document == nil {
						continue
					}
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
					// Deduplicate by URL
					if !containsSource(sources, src.URL) {
						sources = append(sources, src)
					}
				}
			}
		}
	}

	return strings.Join(textParts, ""), sources
}

func containsSource(sources []Source, url string) bool {
	for _, s := range sources {
		if s.URL == url {
			return true
		}
	}
	return false
}

// itoa is a tiny int-to-string helper to avoid importing strconv.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	b := make([]byte, 0, 4)
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
