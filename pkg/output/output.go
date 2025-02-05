// Package output provides standardized output formatting with syntax highlighting
// for various data formats like JSON and YAML. It handles terminal detection
// and colorization based on the environment and user preferences.
package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"golang.org/x/term"
)

// For testing
var isTerminalCheck = term.IsTerminal

// Format constants for supported output formats.
const (
	FormatJSON = "json"
	FormatYAML = "yaml"
)

// Options configures the output formatting behavior.
type Options struct {
	// Format specifies the output format (json or yaml)
	Format string
	// NoColor disables syntax highlighting even in terminal environments
	NoColor bool
}

// Write formats and writes content to the writer with optional syntax highlighting.
// It supports JSON and YAML formats, automatically indenting and colorizing the output
// when appropriate for the terminal environment.
func Write(w io.Writer, content []byte, opts Options) error {
	if !shouldColorize(opts) {
		return writeRaw(w, content, opts.Format)
	}

	return writeColorized(w, content, opts.Format)
}

// WriteString is a convenience wrapper for Write that accepts a string input.
// It converts the string to bytes and delegates to Write for formatting and output.
func WriteString(w io.Writer, content string, opts Options) error {
	return Write(w, []byte(content), opts)
}

// shouldColorize determines if syntax highlighting should be applied based on
// the terminal environment and user preferences.
func shouldColorize(opts Options) bool {
	return !opts.NoColor && isTerminalCheck(int(os.Stdout.Fd()))
}

// writeRaw writes the content without syntax highlighting, handling indentation
// and formatting based on the specified format.
func writeRaw(w io.Writer, content []byte, format string) error {
	switch format {
	case FormatJSON:
		// Handle empty content
		if len(content) == 0 {
			_, err := fmt.Fprintln(w)
			return err
		}
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, content, "", "  "); err != nil {
			return fmt.Errorf("failed to format JSON: %w", err)
		}
		_, err := fmt.Fprintln(w, prettyJSON.String())
		return err
	case FormatYAML:
		output := string(content)
		if !strings.HasSuffix(output, "\n") {
			output += "\n"
		}
		_, err := fmt.Fprint(w, output)
		return err
	default:
		output := string(content)
		if !strings.HasSuffix(output, "\n") {
			output += "\n"
		}
		_, err := fmt.Fprint(w, output)
		return err
	}
}

// writeColorized writes the content with syntax highlighting using chroma.
// It handles lexer selection, formatting, and style application based on
// the content format.
func writeColorized(w io.Writer, content []byte, format string) error {
	// Handle empty content
	if len(content) == 0 {
		_, err := fmt.Fprintln(w)
		return err
	}

	var lexer chroma.Lexer
	switch format {
	case FormatJSON:
		lexer = lexers.Get("json")
		// Try to validate JSON before colorizing
		if !json.Valid(content) {
			return fmt.Errorf("failed to format JSON: invalid JSON content")
		}
	case FormatYAML:
		lexer = lexers.Get("yaml")
	default:
		lexer = lexers.Fallback
	}

	formatter := formatters.Get("terminal")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}

	iterator, err := lexer.Tokenise(nil, string(content))
	if err != nil {
		return fmt.Errorf("failed to tokenize content: %w", err)
	}

	if formatErr := formatter.Format(w, style, iterator); formatErr != nil {
		return formatErr
	}

	// Add newline if not present
	if format == FormatYAML || format == FormatJSON {
		output := string(content)
		if !strings.HasSuffix(output, "\n") {
			_, err = fmt.Fprint(w, "\n")
		}
	}
	return err
}
