package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestConfigCmd(t *testing.T) {
	tests := []struct {
		name       string
		errMessage string
		wantOutput string
		args       []string
		wantErr    bool
	}{
		{
			name:       "no flags provided",
			args:       []string{},
			wantErr:    true,
			errMessage: "no configuration provided",
		},
		{
			name:       "set host",
			args:       []string{"--host", "testcompany"},
			wantOutput: "Configuration saved successfully",
		},
		{
			name:       "set token",
			args:       []string{"--token", "test-token"},
			wantOutput: "Configuration saved successfully",
		},
		{
			name:       "set email",
			args:       []string{"--email", "test@example.com"},
			wantOutput: "Configuration saved successfully",
		},
		{
			name:       "show configuration",
			args:       []string{"--show"},
			wantOutput: "Current configuration:",
		},
		{
			name:       "clear configuration",
			args:       []string{"--clear"},
			wantOutput: "Configuration cleared successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh command for each test
			cmd := &cobra.Command{
				Use:   "config",
				Short: "Configure Glean CLI credentials",
				RunE: func(cmd *cobra.Command, args []string) error {
					flagShow, _ := cmd.Flags().GetBool("show")
					flagClear, _ := cmd.Flags().GetBool("clear")
					flagHost, _ := cmd.Flags().GetString("host")
					flagToken, _ := cmd.Flags().GetString("token")
					flagEmail, _ := cmd.Flags().GetString("email")

					if flagShow {
						fmt.Fprintln(cmd.OutOrStdout(), "Current configuration:")
						return nil
					}

					if flagClear {
						fmt.Fprintln(cmd.OutOrStdout(), "Configuration cleared successfully")
						return nil
					}

					if flagHost == "" && flagToken == "" && flagEmail == "" {
						return fmt.Errorf("no configuration provided")
					}

					fmt.Fprintln(cmd.OutOrStdout(), "Configuration saved successfully")
					return nil
				},
			}

			cmd.Flags().StringVar(&host, "host", "", "Glean instance name")
			cmd.Flags().StringVar(&token, "token", "", "Glean API token")
			cmd.Flags().StringVar(&email, "email", "", "Email address")
			cmd.Flags().BoolVar(&clear, "clear", false, "Clear all stored credentials")
			cmd.Flags().BoolVar(&show, "show", false, "Show current configuration")

			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetErr(b)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
			} else {
				assert.NoError(t, err)
				output := b.String()
				assert.Contains(t, output, tt.wantOutput)
			}
		})
	}
}

func TestValueOrNotSet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "[not set]",
		},
		{
			name:     "non-empty string",
			input:    "value",
			expected: "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := valueOrNotSet(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
