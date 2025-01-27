package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
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
			// Create a fresh command for each test
			cmd := &cobra.Command{
				Use:   "glean",
				Short: "Glean CLI - A command-line interface for Glean operations.",
				Long: `Work seamlessly with Glean from your command line.

To get started, run 'glean --help'.`,
				Run: func(cmd *cobra.Command, args []string) {
					cmd.Help()
				},
			}

			b := bytes.NewBufferString("")
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
	// Save original stdout
	oldOut := rootCmd.OutOrStdout()
	defer func() {
		rootCmd.SetOut(oldOut)
	}()

	// Create a buffer to capture output
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)

	// Execute with no args should show help
	err := Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Work seamlessly with Glean from your command line")
}
