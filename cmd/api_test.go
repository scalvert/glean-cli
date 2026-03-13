package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPICommand_StdinNotTTY_ReturnsQuickly(t *testing.T) {
	// Simulate a non-TTY stdin with no data (like piping from /dev/null)
	pr, pw, err := os.Pipe()
	require.NoError(t, err)
	pw.Close() // EOF immediately — no data
	oldStdin := os.Stdin
	os.Stdin = pr
	defer func() { os.Stdin = oldStdin; pr.Close() }()

	// No mock needed — command should error before hitting API
	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"api", "users/me"})
	// This must complete without hanging
	err = root.Execute()
	// We expect an error because no body was provided and stdin was empty
	assert.Error(t, err)
}

func TestAPICommand_Preview_WritesToCmdOut(t *testing.T) {
	root := NewCmdRoot()
	buf := &bytes.Buffer{}
	root.SetOut(buf)
	root.SetErr(&bytes.Buffer{})
	root.SetArgs([]string{"api", "search", "--method", "POST", "--raw-field", `{"query":"test"}`, "--preview"})
	_ = root.Execute()
	// Preview output must appear in buf (cmd.OutOrStdout())
	// It should contain the HTTP method and endpoint info
	assert.NotEmpty(t, buf.String(), "preview output should be written to cmd.OutOrStdout()")
}

func TestApiCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no endpoint provided",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "help flag",
			args:    []string{"--help"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			cmd := NewCmdAPI()
			cmd.SetOut(b)
			cmd.SetErr(b)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
