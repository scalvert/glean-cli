package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
		args       []string
		wantErr    bool
	}{
		{
			name:       "shows help with no args",
			args:       []string{},
			wantOutput: "Work seamlessly with Glean from your command line",
		},
		{
			name:       "shows help with --help flag",
			args:       []string{"--help"},
			wantOutput: "Work seamlessly with Glean from your command line",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewCmdRoot()
			cmd.SetOut(b)
			cmd.SetErr(b)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				output := b.String()
				assert.Contains(t, output, tt.wantOutput)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	// Redirect stdout to capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Reset stdout after test
	defer func() {
		os.Stdout = old
	}()

	// Execute with help flag
	os.Args = []string{"glean", "--help"}
	if err := Execute(); err != nil {
		t.Errorf("Execute() error = %v", err)
	}

	// Close writer and read output
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check output contains expected text
	if !strings.Contains(output, "Work seamlessly with Glean from your command line") {
		t.Errorf("Expected help text not found in output")
	}
}
