package tui

import "github.com/charmbracelet/lipgloss"

// Glean brand colors
const (
	colorBlue   = "#4339F2"
	colorYellow = "#DFFC6A"
	colorGray   = "#586069"
	colorWhite  = "#FFFFFF"
)

var (
	styleStatusBar = lipgloss.NewStyle().
			Background(lipgloss.Color(colorBlue)).
			Foreground(lipgloss.Color(colorWhite)).
			Padding(0, 1)

	styleUserPrompt = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBlue)).
			Bold(true)

	styleAIResponse = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite))

	styleSource = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorGray)).
			Italic(true)

	styleHelpKey = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorYellow)).
			Bold(true)

	styleHelpDesc = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorGray))

	styleBorder = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(colorBlue))
)
