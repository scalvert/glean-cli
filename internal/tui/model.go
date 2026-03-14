package tui

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
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
	streamRenderLen  int                      // chars since last glamour render (throttle)
	width            int
	height           int
	isStreaming      bool
	showHelp         bool
}

// New creates a fully-initialized TUI model.
func New(cfg *config.Config, session *Session, identity string, ctx context.Context) (*Model, error) {
	ta := textarea.New()
	ta.Placeholder = "Message Glean…"
	ta.Focus()
	ta.ShowLineNumbers = false
	ta.CharLimit = 4096
	// shift+enter is terminal-dependent and unreliable; disable the claim.
	ta.KeyMap.InsertNewline.SetEnabled(false)
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
			m.resizeViewportToContent()
			return m, nil

		case "ctrl+r":
			m.session = &Session{}
			m.conversationMsgs = nil
			m.history.Reset()
			m.viewport.SetContent("")
			m.resizeViewportToContent()
			return m, nil

			// Scroll keys: route to viewport when conversation exists.
		case "up", "down", "pgup", "pgdown":
			if m.history.Len() > 0 {
				m.viewport, vpCmd = m.viewport.Update(msg)
				return m, vpCmd
			}

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
			m.streamRenderLen = 0

			turn := Turn{Role: roleUser, Content: question}
			m.addTurnToHistory(turn)
			m.addTurnToConversation(turn)
			m.session.AddTurn(roleUser, question, nil)
			m.viewport.SetContent(m.history.String())
			m.resizeViewportToContent()
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
			text := m.currentResponse.String()
			turn := Turn{Role: roleAssistant, Content: text, Sources: m.currentSources}
			// addTurnToHistory renders markdown and appends sources to m.history.
			// This is what was missing — the streaming text was never committed
			// to m.history, so it vanished when currentResponse was reset.
			m.addTurnToHistory(turn)
			m.addTurnToConversation(turn)
			m.session.AddTurn(roleAssistant, text, m.currentSources)
		}

		m.viewport.SetContent(m.history.String())
		m.resizeViewportToContent()
		m.viewport.GotoBottom()
		m.currentResponse.Reset()
		m.currentSources = nil
		return m, nil

	// Mouse scroll events go to the viewport when there is content.
	case tea.MouseMsg:
		if m.history.Len() > 0 {
			m.viewport, vpCmd = m.viewport.Update(msg)
			return m, vpCmd
		}
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

	chunk := stripStagePreamble(newText.String())
	if strings.TrimSpace(chunk) == "" {
		return
	}

	// Append chunk to the accumulated response.
	m.currentResponse.WriteString(chunk)
	m.streamRenderLen += len(chunk)

	// Throttle glamour renders: only re-render every ~80 chars or when a
	// newline arrives (sentence/paragraph boundary). This prevents the glamour
	// renderer from becoming a bottleneck on rapid single-word token streams.
	if m.streamRenderLen >= 80 || strings.ContainsAny(chunk, "\n.!?") {
		m.streamRenderLen = 0
		m.rebuildViewport()
	}
}

// stripStagePreamble removes Glean chat stage preamble lines from a chunk.
// Stage markers arrive interleaved with real content in the same fragment —
// dropping the whole chunk (as isStageMarker did) removes actual content too.
// This strips individual marker lines and keeps everything else.
func stripStagePreamble(text string) string {
	lines := strings.Split(text, "\n")
	kept := lines[:0]
	for _, line := range lines {
		if !isStageLine(strings.TrimSpace(line)) {
			kept = append(kept, line)
		}
	}
	return strings.Join(kept, "\n")
}

// isStageLine returns true if a single line is a Glean preamble stage marker.
func isStageLine(line string) bool {
	if line == "" {
		return false
	}
	// Explicit stage prefixes: **Searching:** query, **Reading:** docs, **Writing:**
	for _, prefix := range []string{"**Searching:**", "**Reading:**", "**Writing:**"} {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	// Search topic lines: **Searching some query** (bold wrap, no colon)
	if strings.HasPrefix(line, "**Searching ") && strings.HasSuffix(line, "**") {
		return true
	}
	// Reading topic lines: **Reading some doc**
	if strings.HasPrefix(line, "**Reading ") && strings.HasSuffix(line, "**") {
		return true
	}
	return false
}

// rebuildViewport regenerates the viewport content from saved history +
// the current in-progress streaming response, rendered through glamour.
func (m *Model) rebuildViewport() {
	var buf strings.Builder
	buf.WriteString(m.history.String())

	// Render the in-progress response through glamour so markdown is
	// readable during streaming (not raw **bold** text).
	if m.currentResponse.Len() > 0 {
		rendered := m.renderMarkdown(m.currentResponse.String())
		buf.WriteString(rendered)
	}

	m.viewport.SetContent(buf.String())
	m.resizeViewportToContent()
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

// recalculateLayout updates widget widths and the maximum viewport height on
// terminal resize, then sizes the viewport to its current content.
func (m *Model) recalculateLayout() {
	if m.width == 0 || m.height == 0 {
		return
	}
	// Viewport width matches the input box content width so text and input
	// share the same horizontal margins (left border + padding = 2 chars each side).
	m.viewport.Width = m.width
	m.textarea.SetWidth(m.width - 4)
	m.resizeViewportToContent()

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

// resizeViewportToContent sets the viewport height to exactly the number of
// content lines when they fit in the available space, so the input box appears
// right below the content (Claude Code style). Once content fills the screen
// the viewport caps at maxVpH and becomes scrollable.
func (m *Model) resizeViewportToContent() {
	if m.width == 0 || m.height == 0 {
		return
	}
	inputH := 5
	statusH := 1
	spacerH := 1
	maxVpH := m.height - logoHeaderLines - spacerH - inputH - statusH
	if maxVpH < 4 {
		maxVpH = 4
	}

	// Count rendered content lines.
	contentLines := strings.Count(m.history.String(), "\n") + 1
	if m.currentResponse.Len() > 0 {
		contentLines += strings.Count(m.currentResponse.String(), "\n") + 1
	}

	vpH := contentLines
	if vpH > maxVpH {
		vpH = maxVpH
	}
	if vpH < 1 {
		vpH = 1
	}
	m.viewport.Height = vpH
}

// addTurnToHistory appends a rendered turn to the viewport history buffer.
func (m *Model) addTurnToHistory(turn Turn) {
	switch turn.Role {
	case roleUser:
		// User message: blue left-border block with "you" label.
		// Left-border style requires no width calculation — adapts naturally.
		inner := styleUserLabel.Render("you") + "  " + styleUserText.Render(turn.Content)
		m.history.WriteString("\n")
		m.history.WriteString(styleUserMsg.Render(inner))
		m.history.WriteString("\n\n")

	case roleAssistant:
		rendered := m.renderMarkdown(turn.Content)
		m.history.WriteString(rendered)
		if len(turn.Sources) > 0 {
			m.history.WriteString(styleSourceHeader.Render("Sources") + "\n")
			for i, s := range turn.Sources {
				title := s.Title
				if title == "" {
					title = s.URL
				}
				// Truncate long titles so they fit cleanly in the viewport.
				const maxTitle = 60
				if len([]rune(title)) > maxTitle {
					title = string([]rune(title)[:maxTitle-1]) + "…"
				}
				ds := s.Datasource
				if ds == "" {
					ds = defaultDatasource
				}
				line := fmt.Sprintf("  [%d] %s — %s", i+1, ds, title)
				m.history.WriteString(styleSourceItem.Render(line) + "\n")
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
