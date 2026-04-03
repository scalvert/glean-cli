package output

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// WriteTable uses a bytes.Buffer (not an *os.File), so isTerminalWriter returns
// false and the plain tabwriter path is exercised in all tests below.

func TestWriteTable_PlainFallback(t *testing.T) {
	var buf bytes.Buffer
	err := WriteTable(&buf, []string{"ID", "NAME"}, [][]string{
		{"agent-1", "Research Agent"},
		{"agent-2", "Data Analyst"},
	})
	require.NoError(t, err)

	out := buf.String()
	assert.Contains(t, out, "ID")
	assert.Contains(t, out, "NAME")
	assert.Contains(t, out, "agent-1")
	assert.Contains(t, out, "Research Agent")
	assert.Contains(t, out, "agent-2")
	assert.Contains(t, out, "Data Analyst")
}

func TestWriteTable_EmptyRows(t *testing.T) {
	var buf bytes.Buffer
	err := WriteTable(&buf, []string{"ID", "NAME"}, [][]string{})
	require.NoError(t, err)

	out := buf.String()
	assert.Contains(t, out, "ID")
	assert.Contains(t, out, "NAME")
}

func TestWriteTable_ColumnAlignment(t *testing.T) {
	var buf bytes.Buffer
	err := WriteTable(&buf, []string{"A", "B"}, [][]string{
		{"short", "x"},
		{"much-longer-value", "y"},
	})
	require.NoError(t, err)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	// Each row should be on its own line (header + 2 rows = 3 lines)
	assert.Len(t, lines, 3)
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		n        int
		expected string
	}{
		{"hello", 10, "hello"},
		{"hello", 5, "hello"},
		{"hello world", 8, "hello w…"},
		{"", 5, ""},
		{"短い", 10, "短い"},
		{"これは長いテキストです", 6, "これは長い…"},
		{"Line one\nLine two\nLine three", 20, "Line one Line two L…"},
		{"has\ttabs\tand  spaces", 20, "has tabs and spaces"},
		{"\n\nnewlines only\n\n", 15, "newlines only"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, Truncate(tt.input, tt.n))
		})
	}
}
