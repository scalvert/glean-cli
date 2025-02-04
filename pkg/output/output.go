// Package output provides standardized colorized output functionality
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

const (
	FormatJSON = "json"
	FormatYAML = "yaml"
)

// Options represents configuration for output formatting
type Options struct {
	Format  string
	NoColor bool
}

// For testing
var isTerminalCheck = term.IsTerminal

// Write formats and writes the content to the writer with optional syntax highlighting
func Write(w io.Writer, content []byte, opts Options) error {
	if !shouldColorize(opts) {
		return writeRaw(w, content, opts.Format)
	}

	return writeColorized(w, content, opts.Format)
}

// WriteString is a convenience wrapper for Write that takes a string
func WriteString(w io.Writer, content string, opts Options) error {
	return Write(w, []byte(content), opts)
}

func shouldColorize(opts Options) bool {
	return !opts.NoColor && isTerminalCheck(int(os.Stdout.Fd()))
}

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
