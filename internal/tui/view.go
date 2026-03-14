package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// gleanLogo is a braille Unicode rendering of the Glean wordmark, generated
// from the official logo image via chafa (--symbols braille --size 60x6).
// Braille's 2×4 dot grid reproduces the circular "g" and curved letterforms
// far more faithfully than ASCII art. Each line is 30 terminal columns wide.
const gleanLogo = "⠀⠀⠀⠀⠀⠀⠀⢸⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀\n" +
	"⠀⢀⣀⣀⣀⣴⡄⢸⣿⠀⠀⠀⣀⣀⣀⠀⠀⠀⢀⣀⣀⡀⠀⠀⠀⢀⣀⣀⡀⠀\n" +
	"⣰⡿⠛⠛⠻⣿⡀⢸⣿⠀⣰⡿⠛⢛⣻⣷⡀⣰⡿⠛⠛⠻⣷⡀⣴⡿⠛⠛⢿⣦\n" +
	"⢿⣇⠀⠀⢀⣿⠇⢸⣿⠀⢿⣷⠿⠟⢋⣭⠄⣿⣇⠀⠀⢀⣿⡇⣿⡇⠀⠀⢸⣿\n" +
	"⠈⠻⠿⠾⠿⠋⣠⡈⠻⠿⠈⠻⠿⠾⠿⠋⠀⠈⠻⠿⠾⠿⠿⠇⠿⠇⠀⠀⠸⠿\n" +
	"⠀⠀⣶⣶⣶⠿⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀"

const gleanTagline = "AI-powered search for your company's knowledge"

// logoHeaderLines is the number of rows the header occupies, used by
// recalculateLayout to size the viewport correctly.
const logoHeaderLines = 10 // 1 blank + 6 braille + 1 blank + 1 identity + 1 blank

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

	// Body: welcome hints (empty session) or conversation viewport.
	// Viewport is full terminal width — same outer edge as the input box.
	var body string
	if len(m.session.Turns) == 0 && m.lastErr == nil && m.currentResponse.Len() == 0 {
		body = m.welcomeBody()
	} else {
		body = m.viewport.View()
	}

	// Input box — rounded border, full width.
	inputBox := styleInputFocused.
		Width(m.width - 2).
		Render(m.textarea.View())

	// After first ctrl+c, replace status bar with exit hint.
	bottom := m.statusLine()
	if m.showExitHint {
		bottom = styleExitHint.Render("  Press ctrl+c again to exit  ·  esc to cancel")
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		body,
		"",
		inputBox,
		bottom,
	)
}

// headerView renders the logo and identity line — shown on every screen.
func (m *Model) headerView() string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(centerBlock(gleanLogo, styleLogo, m.width))
	sb.WriteString("\n\n")
	if m.identity != "" {
		sb.WriteString(center(styleTagline.Render(m.identity), m.width))
	} else {
		sb.WriteString(center(styleTagline.Render(gleanTagline), m.width))
	}
	sb.WriteString("\n")
	return sb.String()
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
func (m *Model) statusLine() string {
	// Status bar left: streaming phase indicator OR turn count.
	// Identity lives in the header — no duplication here.
	var left string
	switch {
	case m.isStreaming && !m.streamHasContent:
		left = m.spinner.View() + " " + styleStatusBar.Render("Thinking…")
	case m.isStreaming:
		left = m.spinner.View() + " " + styleStatusBar.Render("Responding…")
	default:
		turns := len(m.session.Turns)
		if turns > 0 {
			left = styleStatusAccent.Render(fmt.Sprintf("%d", turns)) +
				styleStatusBar.Render(" turn"+pluralS(turns))
		}
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
	}

	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(center(styleLogo.Render("Keyboard shortcuts"), m.width))
	sb.WriteString("\n\n")
	for _, s := range shortcuts {
		line := "  " +
			styleHelpKey.Render(fmt.Sprintf("%-26s", s.key)) +
			"  " +
			styleHelpDesc.Render(s.desc)
		sb.WriteString(line + "\n")
	}
	sb.WriteString("\n")
	sb.WriteString(center(styleStatusBar.Render("press ctrl+h to close"), m.width))
	sb.WriteString("\n")
	return sb.String()
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
