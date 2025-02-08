// Package theme provides Glean's brand colors and theme utilities for consistent CLI styling
package theme

import (
	"fmt"
	"text/template"

	"github.com/fatih/color"
)

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

// ColorFunc is a function that applies color to text
type ColorFunc func(string) string

// StyleFunc returns a template function that applies styling based on noColor setting
func StyleFunc(noColor bool, style ColorFunc) func(any) string {
	return func(s any) string {
		if noColor {
			return fmt.Sprint(s)
		}
		return style(fmt.Sprint(s))
	}
}

// Color functions for consistent styling
var (
	blueColor      = color.New(color.FgHiBlue).SprintFunc()
	yellowColor    = color.New(color.FgHiYellow).SprintFunc()
	secondaryColor = color.New(color.FgHiBlack).SprintFunc()

	// Blue returns text styled with Glean's brand blue
	Blue = func(s string) string { return blueColor(s) }

	// Yellow returns text styled with a muted version of Glean's yellow
	Yellow = func(s string) string { return yellowColor(s) }

	// Secondary returns text styled with a muted color for less emphasis
	Secondary = func(s string) string { return secondaryColor(s) }

	// Bold returns text in bold
	Bold = func(s string) string {
		return fmt.Sprintf("\033[1m%s\033[0m", s)
	}
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
