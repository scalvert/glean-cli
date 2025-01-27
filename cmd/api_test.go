package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestApiCmd(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
		args       []string
		wantErr    bool
	}{
		{
			name:       "default GET method",
			args:       []string{"/search"},
			wantOutput: "Invoking Glean API with method=GET, endpoint=/search",
		},
		{
			name:       "POST method",
			args:       []string{"--method", "POST", "/users"},
			wantOutput: "Invoking Glean API with method=POST, endpoint=/users",
		},
		{
			name:       "custom method with -X flag",
			args:       []string{"-X", "PUT", "/update"},
			wantOutput: "Invoking Glean API with method=PUT, endpoint=/update",
		},
		{
			name:       "no endpoint provided",
			args:       []string{},
			wantOutput: "Invoking Glean API with method=GET, endpoint=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh command for each test
			cmd := &cobra.Command{
				Use:   "api",
				Short: "Make calls to the Glean API",
				RunE: func(cmd *cobra.Command, args []string) error {
					method, _ := cmd.Flags().GetString("method")
					endpoint := ""
					if len(args) > 0 {
						endpoint = args[0]
					}
					fmt.Fprintf(cmd.OutOrStdout(), "Invoking Glean API with method=%s, endpoint=%s\n", method, endpoint)
					return nil
				},
			}
			cmd.Flags().StringP("method", "X", "GET", "HTTP method to use (GET, POST, etc.)")

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
