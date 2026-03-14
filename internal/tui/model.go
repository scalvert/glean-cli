package tui

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/atotto/clipboard"
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
	cfg                *config.Config
	session            *Session
	ctx                context.Context
	identity           string                   // "email · host" shown in header + status
	conversationMsgs   []components.ChatMessage // full history sent on each turn (multi-turn context)
	currentResponse    strings.Builder          // accumulates the in-progress assistant response
	currentSources     []Source                 // accumulates sources for the in-progress response
	streamRenderLen    int                      // chars since last glamour render (throttle)
	streamHasContent   bool                     // true once first CONTENT message received
	startTime          time.Time                // session start, for stats on quit
	lastCtrlC          time.Time                // for double ctrl+c detection
	showExitHint       bool                     // show "press ctrl+c again to exit" hint
	lastErr            error
	width              int
	height             int
	isStreaming        bool
	showHelp           bool
	historyIdx         int
	conversationActive bool
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
	ta.SetHeight(1)
	ta.Prompt = styleStatusAccent.Render("❯") + " "
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	vp := viewport.New(0, 0)

	// Custom braille circular spinner — looks like a spinning ball.
	// Braille characters create a smooth circular motion that feels more
	// premium than the default dot spinner.
	sp := spinner.New()
	sp.Spinner = spinner.Spinner{
		Frames: []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "},
		FPS:    time.Second / 10,
	}
	sp.Style = styleStatusAccent

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(100),
	)
	if err != nil {
		renderer = nil
	}

	m := &Model{
		viewport:   vp,
		textarea:   ta,
		spinner:    sp,
		renderer:   renderer,
		cfg:        cfg,
		session:    session,
		identity:   identity,
		ctx:        ctx,
		startTime:  time.Now(),
		historyIdx: -1,
	}

	for _, turn := range session.Turns {
		m.addTurnToConversation(turn)
	}
	m.viewport.SetContent(m.renderConversation())
	m.viewport.GotoBottom()

	return m, nil
}

// Session returns the current session (used by cmd/root.go for post-exit stats).
func (m *Model) Session() *Session {
	return m.session
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
		case "esc":
			if m.showExitHint {
				// Cancel the pending exit.
				m.showExitHint = false
				m.lastCtrlC = time.Time{}
				return m, nil
			}
			return m, tea.Quit

		case "ctrl+c":
			now := time.Now()
			if !m.lastCtrlC.IsZero() && now.Sub(m.lastCtrlC) < time.Second {
				// Second press within 1s — exit.
				return m, tea.Quit
			}
			m.lastCtrlC = now
			m.showExitHint = true
			return m, nil

		case "ctrl+h":
			m.showHelp = !m.showHelp
			return m, nil

		case "ctrl+l":
			m.lastErr = nil
			m.viewport.SetContent(m.renderConversation())
			// viewport stays at max height — already active
			return m, nil

		case "ctrl+r":
			m.session = &Session{}
			m.conversationMsgs = nil
			m.lastErr = nil
			m.historyIdx = -1
			m.conversationActive = false
			m.viewport.Height = 1 // will be resized by resizeViewportToContent on next render
			m.viewport.SetContent(m.renderConversation())
			m.resizeViewportToContent()
			return m, nil

		case "pgup", "pgdown":
			if m.session != nil && len(m.session.Turns) > 0 {
				m.viewport, vpCmd = m.viewport.Update(msg)
				return m, vpCmd
			}

		case "up":
			// Shell-style history nav when input is single-line.
			if !m.isStreaming && m.textarea.LineCount() <= 1 {
				msgs := userMessages(m.session.Turns)
				if len(msgs) > 0 {
					if m.historyIdx == -1 {
						m.historyIdx = len(msgs) - 1
					} else if m.historyIdx > 0 {
						m.historyIdx--
					}
					m.textarea.SetValue(msgs[m.historyIdx])
					m.textarea.CursorEnd()
					return m, nil
				}
			}
			// Fall through to viewport scroll.
			if m.session != nil && len(m.session.Turns) > 0 {
				m.viewport, vpCmd = m.viewport.Update(msg)
				return m, vpCmd
			}

		case "down":
			if !m.isStreaming && m.historyIdx >= 0 {
				msgs := userMessages(m.session.Turns)
				m.historyIdx++
				if m.historyIdx >= len(msgs) {
					m.historyIdx = -1
					m.textarea.SetValue("")
				} else {
					m.textarea.SetValue(msgs[m.historyIdx])
					m.textarea.CursorEnd()
				}
				return m, nil
			}
			if m.session != nil && len(m.session.Turns) > 0 {
				m.viewport, vpCmd = m.viewport.Update(msg)
				return m, vpCmd
			}

		case "ctrl+y":
			for i := len(m.session.Turns) - 1; i >= 0; i-- {
				if m.session.Turns[i].Role == roleAssistant {
					_ = clipboard.WriteAll(m.session.Turns[i].Content)
					break
				}
			}
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
			m.historyIdx = -1
			m.isStreaming = true
			m.currentResponse.Reset()
			m.currentSources = nil
			m.streamRenderLen = 0
			m.streamHasContent = false

			// Transition to active state: fix viewport at max height.
			// This only runs once per session — after this the viewport never resizes.
			if !m.conversationActive {
				m.conversationActive = true
				m.viewport.Height = m.maxViewportHeight()
			}

			turn := Turn{Role: roleUser, Content: question}
			m.addTurnToConversation(turn)
			m.session.AddTurn(roleUser, question, nil)
			m.viewport.SetContent(m.renderConversation())
			m.viewport.GotoBottom()

			return m, tea.Batch(m.spinner.Tick, m.callAPI())
		}

	case spinner.TickMsg:
		if m.isStreaming {
			m.spinner, spCmd = m.spinner.Update(msg)
			if !m.streamHasContent {
				// Animate the thinking indicator in the viewport.
				m.viewport.SetContent(m.renderConversation())
				m.viewport.GotoBottom()
			}
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
			m.addTurnToConversation(turn)
			m.session.AddTurn(roleAssistant, text, m.currentSources)
		}
		m.viewport.SetContent(m.renderConversation())
		m.viewport.GotoBottom()
		m.currentResponse.Reset()
		m.currentSources = nil
		return m, nil

	// Mouse scroll events go to the viewport when there is content.
	case tea.MouseMsg:
		if m.session != nil && len(m.session.Turns) > 0 {
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
		// Only process CONTENT messages. UPDATE, CONTROL, DEBUG, DEBUG_EXTERNAL,
		// HEADING, WARNING, SERVER_TOOL etc. are internal signals not meant for
		// display. DEBUG/DEBUG_EXTERNAL carry Glean's internal "I'm thinking…"
		// reasoning steps which must not be shown to the user.
		if msg.MessageType != nil && *msg.MessageType != components.MessageTypeContent {
			continue
		}
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
	m.streamHasContent = true

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

// rebuildViewport updates viewport content during streaming.
// Never resizes — the viewport height is fixed once conversationActive is true.
func (m *Model) rebuildViewport() {
	m.viewport.SetContent(m.renderConversation())
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
// terminal resize. When the conversation is active the viewport is pinned to
// maxViewportHeight; otherwise it auto-sizes to its content.
func (m *Model) recalculateLayout() {
	if m.width == 0 || m.height == 0 {
		return
	}
	m.viewport.Width = m.width
	m.textarea.SetWidth(m.width - 4)

	// Always recalculate max height on resize; if active, pin to max.
	if m.conversationActive {
		m.viewport.Height = m.maxViewportHeight()
	} else {
		m.resizeViewportToContent()
	}

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

// resizeViewportToContent sets the viewport height to content size when small,
// or caps at maxVpH when the conversation fills the screen. Only called in
// the empty state (conversationActive == false).
func (m *Model) resizeViewportToContent() {
	if m.width == 0 || m.height == 0 {
		return
	}
	inputH := 3
	statusH := 1
	spacerH := 1
	maxVpH := m.height - logoHeaderLines - spacerH - inputH - statusH
	if maxVpH < 4 {
		maxVpH = 4
	}

	// If viewport is already at max height, don't touch it — avoids oscillation.
	if m.viewport.Height >= maxVpH {
		return
	}

	content := m.renderConversation()
	contentLines := strings.Count(content, "\n") + 1
	vpH := contentLines
	if vpH > maxVpH {
		vpH = maxVpH
	}
	if vpH < 1 {
		vpH = 1
	}
	m.viewport.Height = vpH
}

// maxViewportHeight returns the maximum viewport height that fits the terminal.
func (m *Model) maxViewportHeight() int {
	if m.width == 0 || m.height == 0 {
		return 4
	}
	const (
		inputH  = 3 // 1-line textarea + 2 border rows
		statusH = 1
		spacerH = 1
	)
	h := m.height - logoHeaderLines - spacerH - inputH - statusH
	if h < 4 {
		return 4
	}
	return h
}

// renderConversation rebuilds the full viewport content from session turns.
// Called on every viewport update — simpler than incremental updates.
func (m *Model) renderConversation() string {
	var sb strings.Builder
	for _, turn := range m.session.Turns {
		switch turn.Role {
		case roleUser:
			inner := styleUserLabel.Render("you") + "  " + styleUserText.Render(turn.Content)
			sb.WriteString("\n")
			sb.WriteString(styleUserMsg.Render(inner))
			sb.WriteString("\n\n")
		case roleAssistant:
			sb.WriteString(m.renderMarkdown(turn.Content))
			if len(turn.Sources) > 0 {
				sb.WriteString(styleSourceHeader.Render("Sources") + "\n")
				for i, s := range turn.Sources {
					title := s.Title
					if title == "" {
						title = s.URL
					}
					const maxTitle = 60
					if len([]rune(title)) > maxTitle {
						title = string([]rune(title)[:maxTitle-1]) + "…"
					}
					ds := s.Datasource
					if ds == "" {
						ds = defaultDatasource
					}
					sb.WriteString(styleSourceItem.Render(fmt.Sprintf("  [%d] %s — %s", i+1, ds, title)) + "\n")
				}
				sb.WriteString("\n")
			}
		}
	}
	if m.lastErr != nil {
		sb.WriteString("\n")
		sb.WriteString(styleError.Render("  Error: " + m.lastErr.Error()))
		sb.WriteString("\n\n")
	}
	// In-progress streaming response.
	if m.currentResponse.Len() > 0 {
		sb.WriteString(m.renderMarkdown(m.currentResponse.String()))
	}
	// Inline thinking/responding indicator — lives in the conversation, not the status bar.
	if m.isStreaming && m.currentResponse.Len() == 0 {
		sb.WriteString("\n  ")
		sb.WriteString(m.spinner.View())
		sb.WriteString("  ")
		sb.WriteString(styleSourceHeader.Render("Thinking…"))
		sb.WriteString("\n")
	}
	return sb.String()
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
	m.lastErr = err
}

// userMessages returns content of all user turns for history navigation.
func userMessages(turns []Turn) []string {
	var msgs []string
	for _, t := range turns {
		if t.Role == roleUser {
			msgs = append(msgs, t.Content)
		}
	}
	return msgs
}

func containsSource(sources []Source, url string) bool {
	for _, s := range sources {
		if s.URL == url {
			return true
		}
	}
	return false
}
