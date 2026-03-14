package tui

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

const gleanTagline = "AI-powered search for your company's knowledge"

// logoHeaderLines is the number of rows the header occupies, used by
// recalculateLayout to size the viewport correctly.
// Layout: 1 solid header bar + 1 blank = 2 rows.
const logoHeaderLines = 2

// View implements tea.Model.
func (m *Model) View() string {
	if m.width == 0 {
		return ""
	}
	if m.showHelp {
		return m.helpView()
	}

	// Logo + identity — always visible, regardless of conversation state.
	header := m.headerView()

	// Input box — rounded border, full width.
	inputBox := styleInputFocused.
		Width(m.width - 2).
		Render(m.textarea.View())

	// After first ctrl+c, replace status bar with exit hint.
	bottom := m.statusLine()
	if m.showExitHint {
		bottom = styleExitHint.Render("  Press ctrl+c again to exit  ·  esc to cancel")
	}

	// Picker and chip are shown in both welcome and active states — the user
	// can type @ before sending their first message.
	picker := m.filePickerView()
	chip := m.attachedFilesView()

	// Welcome state (no conversation yet): no viewport or delimiter rules.
	if !m.conversationActive {
		parts := []string{header, m.welcomeBody(), ""}
		if picker != "" {
			parts = append(parts, picker)
		}
		if chip != "" {
			parts = append(parts, chip)
		}
		parts = append(parts, inputBox, bottom)
		return lipgloss.JoinVertical(lipgloss.Left, parts...)
	}

	// Active state: conversation viewport bounded by delimiter rules.
	rule := styleDelimiter.Render(strings.Repeat("─", m.width))
	parts := []string{header, rule, m.viewport.View(), rule}
	if picker != "" {
		parts = append(parts, picker)
	}
	if chip != "" {
		parts = append(parts, chip)
	}
	parts = append(parts, inputBox, bottom)
	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

// headerView renders a full-width solid header bar — the entire row is filled
// with brand blue so it reads as a proper application title bar, not a badge.
// Layout: left-padded "glean" · separator · identity · right-fill to terminal edge
func (m *Model) headerView() string {
	if m.width == 0 {
		return "\n"
	}

	identity := m.identity
	if identity == "" {
		identity = gleanTagline
	}

	// Every piece shares the same background so the bar renders as a solid stripe.
	bg := lipgloss.Color(colorBrand)

	gleanText := lipgloss.NewStyle().
		Background(bg).Foreground(lipgloss.Color("#FFFFFF")).Bold(true).
		Render("  glean")

	sepText := lipgloss.NewStyle().
		Background(bg).Foreground(lipgloss.Color("#8890FF")).
		Render("  │  ")

	rightText := lipgloss.NewStyle().
		Background(bg).Foreground(lipgloss.Color("#C8CAFF")).Italic(true).
		Render(identity)

	used := lipgloss.Width(gleanText) + lipgloss.Width(sepText) + lipgloss.Width(rightText)
	fillW := m.width - used
	if fillW < 0 {
		fillW = 0
	}
	fill := lipgloss.NewStyle().Background(bg).Render(strings.Repeat(" ", fillW))

	return gleanText + sepText + rightText + fill + "\n"
}

// welcomeBody renders the hints shown below the logo when no conversation exists.
func (m *Model) welcomeBody() string {
	var sb strings.Builder
	sb.WriteString("\n")

	if preview := m.sessionPreview(); preview != "" {
		sb.WriteString(center(styleSourceHeader.Render("Last session: "+preview), m.width))
		sb.WriteString("\n\n")
	} else {
		sb.WriteString(center(styleTagline.Render(gleanTagline), m.width))
		sb.WriteString("\n\n")
	}

	sb.WriteString(center(styleSourceHeader.Render("Start typing to ask Glean anything"), m.width))
	sb.WriteString("\n")
	return sb.String()
}

// sessionPreview returns a truncated first user message, or "".
func (m *Model) sessionPreview() string {
	for _, t := range m.session.Turns {
		if t.Role == roleUser && t.Content != "" {
			msg := t.Content
			const maxLen = 55
			if len([]rune(msg)) > maxLen {
				msg = string([]rune(msg)[:maxLen]) + "…"
			}
			return "\u201c" + msg + "\u201d"
		}
	}
	return ""
}

// statusLine renders the one-line hint bar at the bottom of the screen.
// The spinner lives in the viewport content area, not here.
func (m *Model) statusLine() string {
	// Left side: mode badge + optional turn count.
	modeLabel := styleStatusAccent.Render(string(m.agentMode))
	var left string
	turns := len(m.session.Turns)
	if turns > 0 {
		left = modeLabel +
			styleStatusBar.Render("  ·  ") +
			styleStatusAccent.Render(fmt.Sprintf("%d", turns)) +
			styleStatusBar.Render(" turn"+pluralS(turns))
	} else {
		left = modeLabel
	}

	right := styleStatusBar.Render("ctrl+r new  ctrl+l clear  ctrl+y copy  ctrl+h help  ctrl+c quit")

	leftW := lipgloss.Width(left)
	rightW := lipgloss.Width(right)
	gap := m.width - leftW - rightW - 2
	if gap < 1 {
		gap = 1
	}
	return left + strings.Repeat(" ", gap) + right
}

// helpView renders the keyboard shortcut reference.
func (m *Model) helpView() string {
	shortcuts := []struct{ key, desc string }{
		{"enter", "Send message"},
		{"shift+enter", "New line in input"},
		{"↑ / ↓  or  pgup / pgdn", "Scroll history"},
		{"ctrl+r", "New session (clear history)"},
		{"ctrl+l", "Clear screen"},
		{"ctrl+c  /  esc", "Quit"},
		{"ctrl+h", "Toggle this help"},
		{"", ""},
		{"/clear", "Start a new session"},
		{"/mode fast|advanced|auto", "Switch agent reasoning depth"},
		{"/help", "Show this help"},
	}

	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(center(styleLogo.Render("Keyboard shortcuts"), m.width))
	sb.WriteString("\n\n")
	for _, s := range shortcuts {
		if s.key == "" {
			sb.WriteString("\n")
			continue
		}
		line := "  " +
			styleHelpKey.Render(fmt.Sprintf("%-30s", s.key)) +
			"  " +
			styleHelpDesc.Render(s.desc)
		sb.WriteString(line + "\n")
	}
	sb.WriteString("\n")
	sb.WriteString(center(styleStatusBar.Render("press ctrl+h to close"), m.width))
	sb.WriteString("\n")
	return sb.String()
}

// filePickerView renders the file picker overlay shown when the user types @.
// Shows at most 5 items with the selected item highlighted in brand blue.
func (m *Model) filePickerView() string {
	if !m.showFilePicker || len(m.filePickerItems) == 0 {
		return ""
	}
	maxItems := 5
	items := m.filePickerItems
	if len(items) > maxItems {
		items = items[:maxItems]
	}
	var sb strings.Builder
	sb.WriteString(stylePickerHeader.Render("  @ file") + "\n")
	for i, item := range items {
		if i == m.filePickerIdx {
			sb.WriteString(stylePickerSelected.Render("  ▸ " + item))
		} else {
			sb.WriteString(stylePickerItem.Render("    " + item))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// attachedFilesView renders a one-line row of file chips for files staged
// for the next message. Returns "" when no files are attached.
func (m *Model) attachedFilesView() string {
	if len(m.attachedFiles) == 0 {
		return ""
	}
	var parts []string
	for _, f := range m.attachedFiles {
		parts = append(parts, styleAttached.Render("📎 "+filepath.Base(f.Path)))
	}
	return "  " + strings.Join(parts, "   ")
}

// centerBlock renders each line of a multi-line string with the given style
// and centers each line independently.
func centerBlock(s string, style lipgloss.Style, termWidth int) string {
	lines := strings.Split(s, "\n")
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = center(style.Render(line), termWidth)
	}
	return strings.Join(result, "\n")
}

// center horizontally centers a styled single-line string within termWidth columns.
func center(s string, termWidth int) string {
	visible := lipgloss.Width(s)
	pad := (termWidth - visible) / 2
	if pad < 0 {
		pad = 0
	}
	return strings.Repeat(" ", pad) + s
}

// StatsLine builds a brief session summary printed to stdout after the TUI exits.
// Format: "N turns  ·  5m 30s  ·  Thanks for using Glean"
func (m *Model) StatsLine() string {
	turns := len(m.session.Turns) / 2 // each turn = user + assistant
	if turns == 0 && len(m.session.Turns) > 0 {
		turns = 1
	}

	elapsed := time.Since(m.startTime).Round(time.Second)
	mins := int(elapsed.Minutes())
	secs := int(elapsed.Seconds()) % 60

	var durationStr string
	if mins > 0 {
		durationStr = fmt.Sprintf("%dm %ds", mins, secs)
	} else {
		durationStr = fmt.Sprintf("%ds", secs)
	}

	turnStr := fmt.Sprintf("%d turn%s", turns, pluralS(turns))

	left := styleStatValue.Render(turnStr) + styleStatLabel.Render("  ·  "+durationStr)
	right := styleStatLabel.Render("Thanks for using Glean")

	leftW := lipgloss.Width(left)
	rightW := lipgloss.Width(right)
	gap := m.width - leftW - rightW - 2
	if gap < 1 {
		gap = 1
	}
	return left + strings.Repeat(" ", gap) + right
}

func pluralS(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
