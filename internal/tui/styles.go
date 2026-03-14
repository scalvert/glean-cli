package tui

import "github.com/charmbracelet/lipgloss"

// Brand colors — taken directly from the Glean wordmark SVG.
const (
	colorBlue  = "#343CED" // Glean primary blue (from glean.svg fill)
	colorMuted = "#6B7280" // dimmed text (status bar, hints)
	colorError = "#EF4444" // error red
)

var (
	// Logo — rendered in Glean blue, centered.
	styleLogo = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBlue)).
			Bold(true)

	styleTagline = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			Italic(true)

	// User message box — distinct background block like Claude Code.
	styleUserLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBlue)).
			Bold(true)

	styleUserText = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#111111", Dark: "#E5E7EB"})

	// Cited sources block.
	styleSourceHeader = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorMuted)).
				Italic(true)

	styleSourceItem = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted))

	// Error text.
	styleError = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorError)).
			Bold(true)

	// Status / hint bar — one line at the very bottom.
	styleStatusBar = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted))

	styleStatusAccent = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorBlue)).
				Bold(true)

	// Input border — rounded corners, brand blue when focused.
	styleInputFocused = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(colorBlue))

	// Help overlay.
	styleHelpKey  = lipgloss.NewStyle().Foreground(lipgloss.Color(colorBlue)).Bold(true)
	styleHelpDesc = lipgloss.NewStyle().Foreground(lipgloss.Color(colorMuted))
)
