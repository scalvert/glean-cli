package tui

import (
	"fmt"
	"strings"

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

// View implements tea.Model.
func (m *Model) View() string {
	if m.width == 0 {
		return ""
	}
	if m.showHelp {
		return m.helpView()
	}

	// Main conversation area.
	var mainArea string
	if m.history.Len() == 0 {
		mainArea = m.welcomeView()
	} else {
		mainArea = m.viewport.View()
	}

	// Input box — rounded border.
	inputBox := styleInputFocused.
		Width(m.width - 4).
		PaddingLeft(1).
		PaddingRight(1).
		Render(m.textarea.View())

	// Status line.
	status := m.statusLine()

	return lipgloss.JoinVertical(lipgloss.Left,
		mainArea,
		"",
		inputBox,
		status,
	)
}

// welcomeView renders the compact welcome panel shown on an empty session.
// No vertical padding to fill the viewport — without alt-screen the terminal
// handles whitespace naturally.
func (m *Model) welcomeView() string {
	var sb strings.Builder

	// Logo.
	sb.WriteString("\n")
	sb.WriteString(centerBlock(gleanLogo, styleLogo, m.width))
	sb.WriteString("\n\n")

	// Identity line: "email · host" or just the tagline if no identity.
	if m.identity != "" {
		sb.WriteString(center(styleTagline.Render(m.identity), m.width))
		sb.WriteString("\n")
		sb.WriteString(center(styleSourceHeader.Render(gleanTagline), m.width))
	} else {
		sb.WriteString(center(styleTagline.Render(gleanTagline), m.width))
	}
	sb.WriteString("\n")

	// Session preview: first user message from the saved session.
	if preview := m.sessionPreview(); preview != "" {
		sb.WriteString("\n")
		sb.WriteString(center(styleSourceHeader.Render("Last session: "+preview), m.width))
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	sb.WriteString(center(styleSourceHeader.Render("Start typing to begin a conversation"), m.width))
	sb.WriteString("\n")

	return sb.String()
}

// sessionPreview returns a truncated first user message from the current session,
// or "" if the session is empty.
func (m *Model) sessionPreview() string {
	for _, t := range m.session.Turns {
		if t.Role == "user" && t.Content != "" {
			msg := t.Content
			const maxLen = 55
			if len([]rune(msg)) > maxLen {
				msg = string([]rune(msg)[:maxLen]) + "…"
			}
			return "\u201c" + msg + "\u201d" // "…"
		}
	}
	return ""
}

// statusLine renders the one-line hint bar at the bottom of the screen.
func (m *Model) statusLine() string {
	var left string
	if m.isStreaming {
		left = m.spinner.View() + " " + styleStatusBar.Render("Searching…")
	} else if m.identity != "" {
		left = styleStatusBar.Render(m.identity)
	} else {
		turns := len(m.session.Turns)
		if turns > 0 {
			left = styleStatusAccent.Render(fmt.Sprintf("%d", turns)) +
				styleStatusBar.Render(" turn"+pluralS(turns))
		}
	}

	right := styleStatusBar.Render("ctrl+r new  ctrl+l clear  ctrl+h help  ctrl+c quit")

	leftW := lipgloss.Width(left)
	rightW := lipgloss.Width(right)
	gap := m.width - leftW - rightW - 2
	if gap < 1 {
		gap = 1
	}

	return left + strings.Repeat(" ", gap) + right
}

// helpView renders the full-screen keyboard shortcut reference.
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
// and centers each line independently within termWidth columns.
func centerBlock(s string, style lipgloss.Style, termWidth int) string {
	lines := strings.Split(s, "\n")
	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = center(style.Render(line), termWidth)
	}
	return strings.Join(result, "\n")
}

// center horizontally centers a (possibly ANSI-styled) single-line string
// within termWidth columns.
func center(s string, termWidth int) string {
	visible := lipgloss.Width(s)
	pad := (termWidth - visible) / 2
	if pad < 0 {
		pad = 0
	}
	return strings.Repeat(" ", pad) + s
}

func pluralS(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
