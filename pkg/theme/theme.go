// Package theme provides Glean's brand colors and theme utilities for consistent CLI styling
package theme

import "github.com/fatih/color"

// Brand colors
const (
	// Primary brand colors
	GleanBlue   = "#4339F2" // Primary brand color
	GleanYellow = "#DFFC6A" // Secondary brand color
	GleanPurple = "#7C4DFF" // Accent color

	// Neutral colors for text and backgrounds
	TextPrimary   = "#24292E" // Primary text color (dark)
	TextSecondary = "#586069" // Secondary text color (medium)
	Background    = "#FFFFFF" // Background color
)

// Color functions for consistent styling
var (
	// Blue returns text styled with Glean's brand blue
	Blue = color.New(color.FgHiBlue).SprintFunc()

	// Yellow returns text styled with a muted version of Glean's yellow
	Yellow = color.New(color.FgHiYellow).SprintFunc()

	// Secondary returns text styled with a muted color for less emphasis
	Secondary = color.New(color.FgHiBlack).SprintFunc()
)

// NoColor returns true if color output is disabled
func NoColor() bool {
	return color.NoColor
}

// DisableColors disables color output
func DisableColors() {
	color.NoColor = true
}

// EnableColors enables color output
func EnableColors() {
	color.NoColor = false
}
