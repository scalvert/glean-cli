package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCmd(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
		args       []string
		wantErr    bool
	}{
		{
			name:       "shows help with no args",
			args:       []string{},
			wantOutput: "Use this command to generate various resources",
		},
		{
			name:       "shows help with --help flag",
			args:       []string{"--help"},
			wantOutput: "Use this command to generate various resources",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh command for each test
			cmd := &cobra.Command{
				Use:   "generate",
				Short: "Generate resources or code for Glean",
				Long: `Use this command to generate various resources,
such as OpenAPI specs, configurations, or other Glean-related assets.`,
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
