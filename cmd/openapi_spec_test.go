package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fragment struct {
	Text string `json:"text"`
}

type testCase struct {
	setupFiles  map[string]string
	name        string
	wantOutput  string
	errContains string
	args        []string
	wantErr     bool
}

func TestOpenapiSpecCmd(t *testing.T) {
	tempDir := t.TempDir()

	type testCase struct {
		name        string
		args        []string
		setupFiles  map[string]string
		wantErr     bool
		errContains string
		wantOutput  string
	}

	tests := []testCase{
		{
			name:        "no input provided",
			args:        []string{},
			wantErr:     true,
			errContains: "no input provided",
		},
		{
			name: "input file provided",
			args: []string{"-f", filepath.Join(tempDir, "api.txt")},
			setupFiles: map[string]string{
				"api.txt": "GET /api/users - Get a list of users",
			},
			wantOutput: "openapi:",
		},
		{
			name: "with custom model",
			args: []string{"-f", filepath.Join(tempDir, "api.txt"), "--model", "gpt-3.5-turbo"},
			setupFiles: map[string]string{
				"api.txt": "GET /api/users - Get a list of users",
			},
			wantOutput: "openapi:",
		},
		{
			name: "with output file",
			args: []string{"-f", filepath.Join(tempDir, "api.txt"), "-o", filepath.Join(tempDir, "spec.yaml")},
			setupFiles: map[string]string{
				"api.txt": "GET /api/users - Get a list of users",
			},
		},
		{
			name: "with prompt",
			args: []string{"-f", filepath.Join(tempDir, "api.txt"), "--prompt", "Include authentication details"},
			setupFiles: map[string]string{
				"api.txt": "GET /api/users - Get a list of users",
			},
			wantOutput: "openapi:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up test files
			for path, content := range tt.setupFiles {
				err := os.WriteFile(filepath.Join(tempDir, path), []byte(content), 0644)
				require.NoError(t, err)
			}

			b := bytes.NewBufferString("")
			cmd := NewCmdOpenAPISpec()
			cmd.SetOut(b)
			cmd.SetErr(b)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				assert.NoError(t, err)
				if tt.wantOutput != "" {
					output := b.String()
					assert.Contains(t, output, tt.wantOutput)
				}
			}
		})
	}
}
