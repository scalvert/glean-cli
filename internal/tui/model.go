package tui

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
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

// streamCompleteMsg carries the complete collected response from a single API call.
type streamCompleteMsg struct {
	text    string   // all CONTENT message text concatenated
	sources []Source // all source citations collected
	chatID  *string  // Glean chatId for conversation continuity
	elapsed string   // formatted elapsed time, e.g. "4.2s"
	err     error
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
	chatID             *string                  // Glean chatId — server manages conversation context
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
			return m, spCmd
		}

	case streamCompleteMsg:
		m.isStreaming = false
		if msg.chatID != nil {
			m.chatID = msg.chatID // use Glean's chatId for subsequent turns
		}
		if msg.err != nil {
			m.lastErr = msg.err
		} else if msg.text != "" {
			turn := Turn{
				Role:    roleAssistant,
				Content: msg.text,
				Sources: msg.sources,
				Elapsed: msg.elapsed,
			}
			m.addTurnToConversation(turn)
			m.session.AppendTurn(turn) // preserves Elapsed for renderConversation
		}
		m.viewport.SetContent(m.renderConversation())
		m.viewport.GotoBottom()
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

// callAPI runs the streaming chat request in a goroutine, collects the
// complete CONTENT response, and returns a single streamCompleteMsg.
// Uses Glean's chatId for server-side conversation context when available,
// which gives better multi-turn behavior than manually passing history.
func (m *Model) callAPI() tea.Cmd {
	// When we have a chatId from Glean, only send the current user message.
	// Glean manages conversation context server-side — more reliable than
	// manually building message history which can bias the model.
	var msgs []components.ChatMessage
	if m.chatID != nil {
		// Send only the most recent user message; chatId carries the context.
		if last := m.conversationMsgs[len(m.conversationMsgs)-1]; true {
			msgs = []components.ChatMessage{last}
		}
	} else {
		msgs = make([]components.ChatMessage, len(m.conversationMsgs))
		copy(msgs, m.conversationMsgs)
	}
	chatID := m.chatID
	cfg := m.cfg
	ctx := m.ctx

	return func() tea.Msg {
		save := true
		req := components.ChatRequest{Messages: msgs, SaveChat: &save}
		if chatID != nil {
			req.ChatID = chatID
		}
		body, err := client.StreamChat(ctx, cfg, req)
		if err != nil {
			return streamCompleteMsg{err: err}
		}
		defer body.Close()

		start := time.Now()

		var textBuf strings.Builder
		var sources []Source
		var returnedChatID *string
		seen := map[string]bool{}

		scanner := bufio.NewScanner(body)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || line == "[DONE]" {
				continue
			}
			var resp components.ChatResponse
			if err := json.Unmarshal([]byte(line), &resp); err != nil {
				continue
			}
			// Capture chatId for server-side conversation continuity.
			if resp.ChatID != nil && returnedChatID == nil {
				returnedChatID = resp.ChatID
			}
			for _, msg := range resp.Messages {
				if msg.MessageType != nil && *msg.MessageType != components.MessageTypeContent {
					continue
				}
				for _, frag := range msg.Fragments {
					if frag.Text != nil && *frag.Text != "" {
						textBuf.WriteString(*frag.Text)
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
						if !seen[src.URL] {
							seen[src.URL] = true
							sources = append(sources, src)
						}
					}
				}
			}
		}
		if err := scanner.Err(); err != nil {
			return streamCompleteMsg{err: err}
		}

		elapsed := time.Since(start).Round(10 * time.Millisecond)
		elapsedStr := fmt.Sprintf("%.1fs", elapsed.Seconds())

		return streamCompleteMsg{
			text:    textBuf.String(),
			sources: sources,
			chatID:  returnedChatID,
			elapsed: elapsedStr,
		}
	}
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
			// Response timing indicator — muted, below the response.
			if turn.Elapsed != "" {
				sb.WriteString(styleSourceHeader.Render("  ─── " + turn.Elapsed + " ───"))
				sb.WriteString("\n\n")
			}
		}
	}
	if m.lastErr != nil {
		sb.WriteString("\n")
		sb.WriteString(styleError.Render("  Error: " + m.lastErr.Error()))
		sb.WriteString("\n\n")
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
