package tui

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/config"
)

const (
	roleUser          = "user"
	roleAssistant     = "assistant"
	defaultDatasource = "glean"
)

// streamLineMsg carries one parsed NDJSON line from the chat stream,
// along with the scanner so the pump can schedule the next read.
type streamLineMsg struct {
	line    string
	scanner *bufio.Scanner
	body    io.ReadCloser
}

// streamDoneMsg signals the stream has ended (normally or with error).
type streamDoneMsg struct {
	err error
}

// maxViewportHeight caps the conversation viewport height.
const maxViewportHeight = 20

// Model is the Bubbletea model for the glean chat TUI.
type Model struct {
	// UI components
	viewport viewport.Model
	textarea textarea.Model
	spinner  spinner.Model
	renderer *glamour.TermRenderer

	// State
	cfg              *config.Config
	session          *Session
	ctx              context.Context
	identity         string                   // "email · host" shown in header + status
	conversationMsgs []components.ChatMessage // full history sent on each turn (multi-turn context)
	history          strings.Builder          // rendered conversation for the viewport
	currentResponse  strings.Builder          // accumulates the in-progress assistant response
	currentSources   []Source                 // accumulates sources for the in-progress response
	width            int
	height           int
	isStreaming      bool
	showHelp         bool
}

// New creates a fully-initialized TUI model.
func New(cfg *config.Config, session *Session, identity string, ctx context.Context) (*Model, error) {
	ta := textarea.New()
	ta.Placeholder = "Message Glean…  (shift+enter for a new line)"
	ta.Focus()
	ta.ShowLineNumbers = false
	ta.CharLimit = 4096
	ta.KeyMap.InsertNewline.SetKeys("shift+enter")
	ta.SetHeight(3)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	vp := viewport.New(0, 0)

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = styleStatusAccent

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		renderer = nil
	}

	m := &Model{
		viewport: vp,
		textarea: ta,
		spinner:  sp,
		renderer: renderer,
		cfg:      cfg,
		session:  session,
		identity: identity,
		ctx:      ctx,
	}

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
		taCmd tea.Cmd
		vpCmd tea.Cmd
		spCmd tea.Cmd
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

		case "ctrl+h":
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
			m.currentResponse.Reset()
			m.currentSources = nil

			turn := Turn{Role: roleUser, Content: question}
			m.addTurnToHistory(turn)
			m.addTurnToConversation(turn)
			m.session.AddTurn(roleUser, question, nil)
			m.viewport.SetContent(m.history.String())
			m.viewport.GotoBottom()

			return m, tea.Batch(m.spinner.Tick, m.callAPI())
		}

	case spinner.TickMsg:
		if m.isStreaming {
			m.spinner, spCmd = m.spinner.Update(msg)
			return m, spCmd
		}

	// Each line from the stream: parse it, update the display, pump next line.
	case streamLineMsg:
		m.processStreamLine(msg.line)
		next := msg.scanner
		body := msg.body
		return m, func() tea.Msg {
			return pumpNextLine(next, body)
		}

	case streamDoneMsg:
		m.isStreaming = false

		if msg.err != nil {
			m.appendError(msg.err)
		} else if m.currentResponse.Len() > 0 {
			// Finalize the assistant turn: render full response + sources.
			text := m.currentResponse.String()
			turn := Turn{Role: roleAssistant, Content: text, Sources: m.currentSources}
			// The content has been rendered incrementally; finalize sources only.
			if len(m.currentSources) > 0 {
				m.history.WriteString(styleSourceHeader.Render("  Sources\n"))
				for i, s := range m.currentSources {
					title := s.Title
					if title == "" {
						title = s.URL
					}
					ds := s.Datasource
					if ds == "" {
						ds = defaultDatasource
					}
					m.history.WriteString(styleSourceItem.Render(
						"  ── [" + itoa(i+1) + "] " + ds + ": " + title + "\n",
					))
				}
				m.history.WriteString("\n")
			}
			m.session.AddTurn(roleAssistant, text, m.currentSources)
			m.addTurnToConversation(turn)
		}

		m.viewport.SetContent(m.history.String())
		m.viewport.GotoBottom()
		m.currentResponse.Reset()
		m.currentSources = nil
		return m, nil
	}

	m.textarea, taCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	return m, tea.Batch(taCmd, vpCmd)
}

// processStreamLine parses one NDJSON line and appends text/sources to the
// in-progress response, updating the viewport after each chunk.
func (m *Model) processStreamLine(line string) {
	// Strip SSE "data: " prefix if present.
	line = strings.TrimPrefix(line, "data: ")
	if line == "" || line == "[DONE]" {
		return
	}

	var resp components.ChatResponse
	if err := json.Unmarshal([]byte(line), &resp); err != nil {
		return
	}

	var newText strings.Builder
	for _, msg := range resp.Messages {
		for _, frag := range msg.Fragments {
			if frag.Text != nil && *frag.Text != "" {
				newText.WriteString(*frag.Text)
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
				if !containsSource(m.currentSources, src.URL) {
					m.currentSources = append(m.currentSources, src)
				}
			}
		}
	}

	chunk := newText.String()
	if chunk == "" {
		return
	}

	// Append chunk to the accumulated response.
	m.currentResponse.WriteString(chunk)

	// Re-render the full in-progress response incrementally.
	// Rebuild history from the saved turns + the current partial response.
	m.rebuildViewport()
}

// rebuildViewport regenerates the viewport content from saved history +
// the current in-progress streaming response.
func (m *Model) rebuildViewport() {
	var buf strings.Builder
	buf.WriteString(m.history.String())

	// Append the in-progress response as plain text (no markdown yet —
	// glamour needs the full text to render correctly).
	if m.currentResponse.Len() > 0 {
		buf.WriteString(m.currentResponse.String())
	}

	m.viewport.SetContent(buf.String())
	m.viewport.GotoBottom()
}

// callAPI starts the streaming chat request and returns the first pump cmd.
func (m *Model) callAPI() tea.Cmd {
	msgs := make([]components.ChatMessage, len(m.conversationMsgs))
	copy(msgs, m.conversationMsgs)
	cfg := m.cfg
	ctx := m.ctx

	return func() tea.Msg {
		body, err := client.StreamChat(ctx, cfg, msgs)
		if err != nil {
			return streamDoneMsg{err: err}
		}
		scanner := bufio.NewScanner(body)
		return pumpNextLine(scanner, body)
	}
}

// pumpNextLine reads the next line from the scanner and returns the
// appropriate tea.Msg. This is the streaming pump — each call schedules itself
// as the next Cmd, giving bubbletea a chance to re-render between chunks.
func pumpNextLine(scanner *bufio.Scanner, body io.ReadCloser) tea.Msg {
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // skip blank lines between SSE events
		}
		return streamLineMsg{line: line, scanner: scanner, body: body}
	}
	body.Close()
	if err := scanner.Err(); err != nil {
		return streamDoneMsg{err: err}
	}
	return streamDoneMsg{}
}

// recalculateLayout updates widget sizes based on current terminal dimensions.
func (m *Model) recalculateLayout() {
	if m.width == 0 || m.height == 0 {
		return
	}

	inputH := 5  // 3 content rows + 2 border rows
	statusH := 1 // status bar
	spacerH := 1 // blank between body and input
	vpH := m.height - logoHeaderLines - spacerH - inputH - statusH
	if vpH > maxViewportHeight {
		vpH = maxViewportHeight
	}
	if vpH < 4 {
		vpH = 4
	}

	m.viewport.Width = m.width
	m.viewport.Height = vpH
	m.textarea.SetWidth(m.width - 4)

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

// addTurnToHistory appends a rendered turn to the viewport history buffer.
func (m *Model) addTurnToHistory(turn Turn) {
	switch turn.Role {
	case roleUser:
		m.history.WriteString("\n")
		m.history.WriteString(styleUserLabel.Render("  you  "))
		m.history.WriteString(styleUserText.Render(turn.Content))
		m.history.WriteString("\n\n")

	case roleAssistant:
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
					ds = defaultDatasource
				}
				m.history.WriteString(styleSourceItem.Render(
					"  ── [" + itoa(i+1) + "] " + ds + ": " + title + "\n",
				))
			}
			m.history.WriteString("\n")
		}
	}
}

// addTurnToConversation appends an SDK ChatMessage for multi-turn context.
func (m *Model) addTurnToConversation(turn Turn) {
	switch turn.Role {
	case roleUser:
		text := turn.Content
		m.conversationMsgs = append(m.conversationMsgs, components.ChatMessage{
			Author:      components.AuthorUser.ToPointer(),
			MessageType: components.MessageTypeContent.ToPointer(),
			Fragments:   []components.ChatMessageFragment{{Text: &text}},
		})
	case roleAssistant:
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
