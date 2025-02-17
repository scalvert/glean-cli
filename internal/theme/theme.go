// Package theme provides Glean's brand colors and theme utilities for consistent CLI styling
package theme

import (
	"fmt"
	"text/template"

	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
)

// Brand colors as hex strings
const (
	// Primary brand colors
	GleanBlueHex   = "#4339F2" // Primary brand color
	GleanYellowHex = "#DFFC6A" // Secondary brand color
	GleanPurpleHex = "#7C4DFF" // Accent color

	// Neutral colors for text and backgrounds
	TextPrimary   = "#24292E" // Primary text color (dark)
	TextSecondary = "#586069" // Secondary text color (medium)
	Background    = "#FFFFFF" // Background color
)

// Color represents a themed color with both terminal and TUI support
type Color struct {
	*color.Color
}

// Brand colors as Color objects
var (
	GleanBlue   = Color{color.New(color.FgHiBlue)}
	GleanYellow = Color{color.New(color.FgHiYellow)}
)

// ToLipgloss converts the theme color to a Lipgloss color
func (c Color) ToLipgloss() lipgloss.Color {
	// Map our color values to lipgloss colors
	switch c.Color {
	case color.New(color.FgHiBlue):
		return lipgloss.Color("39") // Bright blue
	case color.New(color.FgHiYellow):
		return lipgloss.Color("220") // Bright yellow
	default:
		return lipgloss.Color("") // No color
	}
}

// SprintFunc returns a function that colorizes text
func (c Color) SprintFunc() func(a ...interface{}) string {
	return c.Color.SprintFunc()
}

// ColorFunc is a function that applies color to text
type ColorFunc = func(a ...interface{}) string

// StyleFunc returns a template function that applies styling based on noColor setting
func StyleFunc(noColor bool, style ColorFunc) func(any) string {
	return func(s any) string {
		if noColor {
			return fmt.Sprint(s)
		}
		return style(s)
	}
}

// Color functions for consistent styling
var (
	// Blue returns text styled with Glean's brand blue
	Blue = GleanBlue.SprintFunc()

	// Yellow returns text styled with a muted version of Glean's yellow
	Yellow = GleanYellow.SprintFunc()

	// Secondary returns text styled with a muted color for less emphasis
	Secondary = color.New(color.FgHiBlack).SprintFunc()

	// Bold returns text in bold
	Bold = color.New(color.Bold).SprintFunc()
)

// TemplateFuncs returns a map of template functions that respect color settings
func TemplateFuncs(noColor bool) template.FuncMap {
	return template.FuncMap{
		"gleanBlue":   StyleFunc(noColor, Blue),
		"gleanYellow": StyleFunc(noColor, Yellow),
		"secondary":   StyleFunc(noColor, Secondary),
		"bold":        StyleFunc(noColor, Bold),
	}
}

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
