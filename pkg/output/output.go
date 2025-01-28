// Package output provides standardized colorized output functionality
package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"golang.org/x/term"
)

// Options represents configuration for output formatting
type Options struct {
	NoColor bool
	Format  string // json, yaml, etc.
}

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
	return !opts.NoColor && term.IsTerminal(int(os.Stdout.Fd()))
}

func writeRaw(w io.Writer, content []byte, format string) error {
	switch format {
	case "json":
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, content, "", "  "); err != nil {
			return fmt.Errorf("failed to format JSON: %w", err)
		}
		_, err := fmt.Fprintln(w, prettyJSON.String())
		return err
	default:
		_, err := w.Write(content)
		return err
	}
}

func writeColorized(w io.Writer, content []byte, format string) error {
	var lexer chroma.Lexer
	switch format {
	case "json":
		lexer = lexers.Get("json")
	case "yaml":
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

	return formatter.Format(w, style, iterator)
}
