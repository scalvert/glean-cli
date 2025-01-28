package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockClient struct {
	response []byte
	err      error
}

func (m *mockClient) SendRequest(req *http.Request) ([]byte, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.response, nil
}

func (m *mockClient) GetFullURL(path string) string {
	return "https://test-company-be.glean.com" + path
}

func setupMockClient(response []byte, err error) func() {
	origFunc := http.NewClientFunc
	http.NewClientFunc = func(cfg *config.Config) (http.Client, error) {
		return &mockClient{response: response, err: err}, nil
	}
	return func() {
		http.NewClientFunc = origFunc
	}
}

func setupTestConfig(t *testing.T) func() {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")
	configData := `{
		"glean_host": "test-company",
		"glean_token": "test-token",
		"glean_email": "test@example.com"
	}`

	err := os.WriteFile(configPath, []byte(configData), 0644)
	require.NoError(t, err)

	oldConfigPath := os.Getenv("GLEAN_CONFIG_PATH")
	os.Setenv("GLEAN_CONFIG_PATH", configPath)

	return func() {
		if oldConfigPath != "" {
			os.Setenv("GLEAN_CONFIG_PATH", oldConfigPath)
		} else {
			os.Unsetenv("GLEAN_CONFIG_PATH")
		}
	}
}

type fragment struct {
	Text string `json:"text"`
}

func TestOpenapiSpecCmd(t *testing.T) {
	cleanup := setupTestConfig(t)
	defer cleanup()

	// Mock successful response
	successResp := struct {
		Messages []struct {
			Fragments []fragment `json:"fragments"`
		} `json:"messages"`
	}{
		Messages: []struct {
			Fragments []fragment `json:"fragments"`
		}{
			{
				Fragments: []fragment{
					{Text: "openapi: 3.0.0\ninfo:\n  title: Test API\n  version: 1.0.0\npaths:\n  /test:\n    get:\n      description: Test endpoint"},
				},
			},
		},
	}
	successRespBytes, err := json.Marshal(successResp)
	require.NoError(t, err)

	tests := []struct {
		name        string
		args        []string
		input       string
		setupFiles  map[string]string
		wantErr     bool
		errContains string
		wantOutput  string
	}{
		{
			name:        "no input provided",
			args:        []string{},
			wantErr:     true,
			errContains: "no input provided",
		},
		{
			name: "input file provided",
			args: []string{"-f", "testdata/api.txt"},
			setupFiles: map[string]string{
				"testdata/api.txt": "GET /api/users - Get a list of users",
			},
			wantOutput: "openapi:",
		},
		{
			name: "with custom model",
			args: []string{"-f", "testdata/api.txt", "--model", "gpt-3.5-turbo"},
			setupFiles: map[string]string{
				"testdata/api.txt": "GET /api/users - Get a list of users",
			},
			wantOutput: "openapi:",
		},
		{
			name: "with output file",
			args: []string{"-f", "testdata/api.txt", "-o", "testdata/spec.yaml"},
			setupFiles: map[string]string{
				"testdata/api.txt": "GET /api/users - Get a list of users",
			},
			wantOutput: "OpenAPI spec written to testdata/spec.yaml",
		},
		{
			name: "with prompt",
			args: []string{"-f", "testdata/api.txt", "--prompt", "Include authentication details"},
			setupFiles: map[string]string{
				"testdata/api.txt": "GET /api/users - Get a list of users",
			},
			wantOutput: "openapi:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock client for each test
			cleanupMock := setupMockClient(successRespBytes, nil)
			defer cleanupMock()

			// Setup test files
			if len(tt.setupFiles) > 0 {
				err := os.MkdirAll("testdata", 0755)
				require.NoError(t, err)
				defer os.RemoveAll("testdata")

				for path, content := range tt.setupFiles {
					err := os.WriteFile(path, []byte(content), 0644)
					require.NoError(t, err)
				}
			}

			// Create a fresh command for each test
			cmd := newOpenapiSpecCmd()
			out := &bytes.Buffer{}
			errOut := &bytes.Buffer{}
			cmd.SetOut(out)
			cmd.SetErr(errOut)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				assert.NoError(t, err)
				output := out.String() + errOut.String()
				t.Logf("Command output: %s", output)
				assert.Contains(t, output, tt.wantOutput)

				// If output file was specified, verify it exists and contains OpenAPI content
				if outputFile := cmd.Flag("output").Value.String(); outputFile != "" {
					content, err := os.ReadFile(outputFile)
					require.NoError(t, err)
					assert.Contains(t, string(content), "openapi:")
					// Cleanup output file
					os.Remove(outputFile)
				}
			}
		})
	}
}
