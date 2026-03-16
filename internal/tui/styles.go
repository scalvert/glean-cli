package tui

import "github.com/charmbracelet/lipgloss"

// Glean brand palette — sourced from www.glean.com.
// Primary blue confirmed as #343CED (60 occurrences, dominant brand color).
const (
	colorBrand  = "#343CED" // Primary brand blue/indigo
	colorGreen  = "#D8FD49" // Yellow-green accent (success, highlights)
	colorPurple = "#E16BFF" // Purple accent
	colorOrange = "#FF7E4C" // Orange accent
	colorError  = "#FF492C" // Coral red (on-brand error color)
	colorMuted  = "#777867" // Warm gray (muted text — more on-brand than generic #6B7280)
	colorCream  = "#E1DFD7" // Warm off-white (surface hint)
	colorDim    = "#484848" // Dark gray (very subtle elements)
)

var (
	// Logo — rendered in brand blue, centered.
	styleLogo = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBrand)).
			Bold(true)

	styleTagline = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted)).
			Italic(true)

	// User message — brand blue left-border block.
	styleUserLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBrand)).
			Bold(true)

	styleUserText = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#111111", Dark: colorCream})

	styleUserMsg = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder(), false, false, false, true).
			BorderForeground(lipgloss.Color(colorBrand)).
			PaddingLeft(1)

	// Cited sources block.
	styleSourceHeader = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorMuted)).
				Italic(true)

	styleSourceItem = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted))

	// System messages — command feedback, mode changes.
	styleSystem = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBrand)).
			Italic(true)

	// Error text.
	styleError = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorError)).
			Bold(true)

	// Status / hint bar — one line at the very bottom.
	styleStatusBar = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted))

	styleStatusAccent = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorBrand)).
				Bold(true)

	// Stats line shown on quit.
	styleStatLabel = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted))

	styleStatValue = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBrand)).
			Bold(true)

	// Input border — rounded corners, brand blue when focused.
	styleInputFocused = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(colorBrand))

	// Help overlay.
	styleHelpKey  = lipgloss.NewStyle().Foreground(lipgloss.Color(colorBrand)).Bold(true)
	styleHelpDesc = lipgloss.NewStyle().Foreground(lipgloss.Color(colorMuted))

	// Viewport delimiter — thin rule above and below the scrollable content area.
	styleDelimiter = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorDim))

	// Exit hint shown after first ctrl+c.
	styleExitHint = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorOrange)).
			Bold(true)

	// File picker overlay styles.
	stylePickerHeader = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorMuted)).
				Italic(true)

	stylePickerItem = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorMuted))

	stylePickerSelected = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorBrand)).
				Bold(true)

	styleAttached = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBrand))
)
