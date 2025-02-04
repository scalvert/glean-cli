package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	// Save original check function and restore after test
	origCheck := isTerminalCheck
	defer func() {
		isTerminalCheck = origCheck
	}()

	// Mock terminal check to always return true
	isTerminalCheck = func(fd int) bool {
		return true
	}

	tests := []struct {
		name        string
		input       []byte
		opts        Options
		wantErr     bool
		checkOutput func(t *testing.T, output string)
	}{
		{
			name:  "writes JSON with color",
			input: []byte(`{"key": "value"}`),
			opts: Options{
				Format:  FormatJSON,
				NoColor: false,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				// Check that output contains color codes
				assert.Contains(t, output, "\x1b[")

				// Remove color codes and check JSON structure
				cleanOutput := stripAnsiCodes(output)
				var result map[string]interface{}
				err := json.Unmarshal([]byte(cleanOutput), &result)
				require.NoError(t, err)
				assert.Equal(t, "value", result["key"])
			},
		},
		{
			name:  "writes JSON without color",
			input: []byte(`{"key": "value"}`),
			opts: Options{
				Format:  FormatJSON,
				NoColor: true,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.NotContains(t, output, "\x1b[")
				assert.Equal(t, "{\n  \"key\": \"value\"\n}\n", output)
			},
		},
		{
			name:  "handles invalid JSON",
			input: []byte("{invalid json}"),
			opts: Options{
				Format: FormatJSON,
			},
			wantErr: true,
			checkOutput: func(t *testing.T, output string) {
				assert.Empty(t, output)
			},
		},
		{
			name:  "writes YAML with color",
			input: []byte("key: value"),
			opts: Options{
				Format:  FormatYAML,
				NoColor: false,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.Contains(t, output, "\x1b[")
				cleanOutput := stripAnsiCodes(output)
				assert.Equal(t, "key: value\n", cleanOutput)
			},
		},
		{
			name:  "writes YAML without color",
			input: []byte("key: value"),
			opts: Options{
				Format:  FormatYAML,
				NoColor: true,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.NotContains(t, output, "\x1b[")
				assert.Equal(t, "key: value\n", output)
			},
		},
		{
			name:  "handles invalid YAML",
			input: []byte("key: : invalid"),
			opts: Options{
				Format: FormatYAML,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.Equal(t, "key: : invalid\n", stripAnsiCodes(output))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := Write(&buf, tt.input, tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			tt.checkOutput(t, buf.String())
		})
	}
}

func TestWriteString(t *testing.T) {
	// Save original check function and restore after test
	origCheck := isTerminalCheck
	defer func() {
		isTerminalCheck = origCheck
	}()

	// Mock terminal check to always return true
	isTerminalCheck = func(fd int) bool {
		return true
	}

	tests := []struct {
		name        string
		input       string
		opts        Options
		wantErr     bool
		checkOutput func(t *testing.T, output string)
	}{
		{
			name:  "writes valid JSON string with color",
			input: `{"key": "value"}`,
			opts: Options{
				Format:  FormatJSON,
				NoColor: false,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.Contains(t, output, "\x1b[")
				cleanOutput := stripAnsiCodes(output)
				var result map[string]interface{}
				err := json.Unmarshal([]byte(cleanOutput), &result)
				require.NoError(t, err)
				assert.Equal(t, "value", result["key"])
			},
		},
		{
			name:  "writes valid JSON string without color",
			input: `{"key": "value"}`,
			opts: Options{
				Format:  FormatJSON,
				NoColor: true,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.NotContains(t, output, "\x1b[")
				assert.Equal(t, "{\n  \"key\": \"value\"\n}\n", output)
			},
		},
		{
			name:  "handles empty string",
			input: "",
			opts: Options{
				Format: FormatJSON,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.Equal(t, "\n", output)
			},
		},
		{
			name:  "handles invalid JSON string",
			input: "{invalid json}",
			opts: Options{
				Format: FormatJSON,
			},
			wantErr: true,
			checkOutput: func(t *testing.T, output string) {
				assert.Empty(t, output)
			},
		},
		{
			name:  "writes valid YAML string with color",
			input: "key: value",
			opts: Options{
				Format:  FormatYAML,
				NoColor: false,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.Contains(t, output, "\x1b[")
				cleanOutput := stripAnsiCodes(output)
				assert.Equal(t, "key: value\n", cleanOutput)
			},
		},
		{
			name:  "writes valid YAML string without color",
			input: "key: value",
			opts: Options{
				Format:  FormatYAML,
				NoColor: true,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.NotContains(t, output, "\x1b[")
				assert.Equal(t, "key: value\n", output)
			},
		},
		{
			name:  "handles invalid YAML string",
			input: "key: : invalid",
			opts: Options{
				Format: FormatYAML,
			},
			wantErr: false,
			checkOutput: func(t *testing.T, output string) {
				assert.Equal(t, "key: : invalid\n", stripAnsiCodes(output))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := WriteString(&buf, tt.input, tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			tt.checkOutput(t, buf.String())
		})
	}
}

func TestShouldColorize(t *testing.T) {
	// Save original check function and restore after test
	origCheck := isTerminalCheck
	defer func() {
		isTerminalCheck = origCheck
	}()

	tests := []struct {
		name       string
		opts       Options
		isTerminal bool
		expected   bool
	}{
		{
			name: "returns false when NoColor is true",
			opts: Options{
				NoColor: true,
			},
			isTerminal: true,
			expected:   false,
		},
		{
			name: "returns false when not a terminal",
			opts: Options{
				NoColor: false,
			},
			isTerminal: false,
			expected:   false,
		},
		{
			name: "returns true when terminal and NoColor is false",
			opts: Options{
				NoColor: false,
			},
			isTerminal: true,
			expected:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isTerminalCheck = func(fd int) bool {
				return tt.isTerminal
			}

			result := shouldColorize(tt.opts)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// stripAnsiCodes removes ANSI escape codes from a string
func stripAnsiCodes(s string) string {
	var result strings.Builder
	inEscape := false
	for _, c := range s {
		if c == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
				inEscape = false
			}
			continue
		}
		result.WriteRune(c)
	}
	return result.String()
}
