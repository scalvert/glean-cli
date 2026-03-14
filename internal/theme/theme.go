// Package theme provides Glean's brand colors and theme utilities for consistent CLI styling.
// Colors sourced directly from www.glean.com (#343CED confirmed as dominant primary, 60 occurrences).
package theme

import (
	"fmt"
	"text/template"

	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
)

// Glean brand palette — sourced from www.glean.com.
const (
	GleanBlueHex   = "#343CED" // Primary brand blue/indigo (was incorrectly #4339F2)
	GleanGreenHex  = "#D8FD49" // Yellow-green accent
	GleanPurpleHex = "#E16BFF" // Purple accent
	GleanOrangeHex = "#FF7E4C" // Orange accent
	GleanErrorHex  = "#FF492C" // Coral red (error / destructive)
	GleanMutedHex  = "#777867" // Warm gray (muted text, on-brand)
	GleanCreamHex  = "#E1DFD7" // Warm off-white (surface)

	// Legacy aliases kept for backward compatibility.
	GleanYellowHex = GleanGreenHex
)

// Color wraps fatih/color with the hex value for lipgloss compatibility.
type Color struct {
	*color.Color
	hex string
}

// ToLipgloss returns the equivalent lipgloss.Color.
func (c Color) ToLipgloss() lipgloss.Color {
	return lipgloss.Color(c.hex)
}

// SprintFunc returns a function that colorizes text.
func (c Color) SprintFunc() func(a ...interface{}) string {
	return c.Color.SprintFunc()
}

// ColorFunc is a function that applies color to text.
type ColorFunc = func(a ...interface{}) string

// StyleFunc returns a template function that respects noColor.
func StyleFunc(noColor bool, style ColorFunc) func(any) string {
	return func(s any) string {
		if noColor {
			return fmt.Sprint(s)
		}
		return style(s)
	}
}

// Branded color functions for consistent CLI stdout output.
var (
	// Blue applies Glean's primary brand color (#343CED).
	Blue = Color{color.New(color.FgHiBlue), GleanBlueHex}.SprintFunc()

	// Muted applies Glean's warm gray for secondary / dimmed text.
	Muted = Color{color.New(color.FgHiBlack), GleanMutedHex}.SprintFunc()

	// Bold makes text bold.
	Bold = color.New(color.Bold).SprintFunc()

	// Success applies a bright accent for positive outcomes.
	Success = Color{color.New(color.FgHiGreen), GleanGreenHex}.SprintFunc()

	// Err applies Glean's coral red for errors.
	Err = Color{color.New(color.FgHiRed), GleanErrorHex}.SprintFunc()

	// Secondary is an alias for Muted (backward compat).
	Secondary = Muted

	// Yellow is kept for backward compat (maps to green accent now).
	Yellow = Color{color.New(color.FgHiYellow), GleanGreenHex}.SprintFunc()
)

// TemplateFuncs returns template helpers that respect color settings.
func TemplateFuncs(noColor bool) template.FuncMap {
	return template.FuncMap{
		"gleanBlue":   StyleFunc(noColor, Blue),
		"gleanYellow": StyleFunc(noColor, Yellow),
		"secondary":   StyleFunc(noColor, Muted),
		"bold":        StyleFunc(noColor, Bold),
		"success":     StyleFunc(noColor, Success),
		"gleanError":  StyleFunc(noColor, Err),
	}
}

// NoColor returns true if color output is disabled.
func NoColor() bool {
	return color.NoColor
}

// DisableColors disables color output.
func DisableColors() {
	color.NoColor = true
}

// EnableColors enables color output.
func EnableColors() {
	color.NoColor = false
}
