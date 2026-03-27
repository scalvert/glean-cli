package output

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/charmbracelet/lipgloss"
	libtable "github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"
)

// WriteTable writes a styled TUI table when writing to a terminal, or a plain
// tab-aligned table when output is piped, so both use cases work cleanly.
func WriteTable(w io.Writer, headers []string, rows [][]string) error {
	if isTerminalWriter(w) {
		return writeStyledTable(w, headers, rows)
	}
	return writePlainTable(w, headers, rows)
}

func isTerminalWriter(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}

func writeStyledTable(w io.Writer, headers []string, rows [][]string) error {
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#777867"))
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#343CED")).
		Bold(true).
		Padding(0, 1)
	cellStyle := lipgloss.NewStyle().Padding(0, 1)
	dimCellStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#AAAAAA")).
		Padding(0, 1)

	t := libtable.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(borderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == libtable.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return dimCellStyle
			default:
				return cellStyle
			}
		}).
		Headers(headers...).
		Rows(rows...)

	_, err := fmt.Fprintln(w, t.Render())
	return err
}

func writePlainTable(w io.Writer, headers []string, rows [][]string) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, strings.Join(headers, "\t"))
	for _, row := range rows {
		fmt.Fprintln(tw, strings.Join(row, "\t"))
	}
	return tw.Flush()
}

// Truncate shortens s to at most n runes, appending "…" if truncated.
func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n-1]) + "…"
}
