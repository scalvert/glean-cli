package cmd

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenapiSpecCmd(t *testing.T) {
	tests := []struct {
		setupFunc   func() error
		cleanupFunc func() error
		name        string
		input       string
		errMessage  string
		args        []string
		wantErr     bool
	}{
		{
			name:       "no input provided",
			args:       []string{},
			wantErr:    true,
			errMessage: "no input found",
		},
		{
			name: "input file provided",
			args: []string{"-f", "testdata/sample_api.txt"},
			setupFunc: func() error {
				err := os.MkdirAll("testdata", 0755)
				if err != nil {
					return err
				}
				return os.WriteFile("testdata/sample_api.txt", []byte("GET /api/v1/users"), 0644)
			},
			cleanupFunc: func() error {
				return os.RemoveAll("testdata")
			},
		},
		{
			name: "with custom model",
			args: []string{"-f", "testdata/sample_api.txt", "--model", "gpt-3.5-turbo"},
			setupFunc: func() error {
				err := os.MkdirAll("testdata", 0755)
				if err != nil {
					return err
				}
				return os.WriteFile("testdata/sample_api.txt", []byte("GET /api/v1/users"), 0644)
			},
			cleanupFunc: func() error {
				return os.RemoveAll("testdata")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupFunc != nil {
				err := tt.setupFunc()
				require.NoError(t, err)
			}

			if tt.cleanupFunc != nil {
				defer func() {
					err := tt.cleanupFunc()
					require.NoError(t, err)
				}()
			}

			// Create a fresh command for each test
			cmd := &cobra.Command{
				Use:   "openapi-spec",
				Short: "Generate an OpenAPI spec from an API definition or curl command",
				RunE: func(cmd *cobra.Command, args []string) error {
					if inputFile == "" {
						return fmt.Errorf("no input found; please provide a file or pipe input to stdin")
					}
					return nil
				},
			}
			cmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input file containing the API/curl command")
			cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file for the OpenAPI spec")
			cmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Additional instructions for the LLM")
			cmd.Flags().StringVar(&model, "model", "gpt-4", "LLM model to use")

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
			}
		})
	}
}
