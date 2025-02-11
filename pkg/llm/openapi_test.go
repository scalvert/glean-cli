package llm

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/scalvert/glean-cli/pkg/config"
	gleanhttp "github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockClient(response []byte, err error) func() {
	origFunc := gleanhttp.NewClientFunc
	gleanhttp.NewClientFunc = func(cfg *config.Config) (gleanhttp.Client, error) {
		return &testutils.MockClient{Response: response, Err: err}, nil
	}
	return func() {
		gleanhttp.NewClientFunc = origFunc
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

//nolint:govet // Ignoring fieldalignment as it's just an optimization
type testCase struct {
	mockErr     error
	name        string
	input       string
	prompt      string
	model       string
	errContains string
	mockResp    []byte
	wantErr     bool
}

func TestGenerateOpenAPISpec(t *testing.T) {
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

	tests := []testCase{
		{
			name:        "empty input",
			input:       "",
			mockResp:    []byte(`{"messages": []}`),
			wantErr:     true,
			errContains: "no response from LLM",
		},
		{
			name:     "with prompt",
			input:    "GET /api/users - Get a list of users",
			prompt:   "Include authentication details",
			model:    "gpt-4",
			mockResp: successRespBytes,
		},
		{
			name:     "curl command",
			input:    `curl -X POST https://api.example.com/users -H "Content-Type: application/json" -d '{"name":"test"}'`,
			model:    "gpt-4",
			mockResp: successRespBytes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := setupMockClient(tt.mockResp, tt.mockErr)
			defer cleanup()

			spec, err := GenerateOpenAPISpec(tt.input, tt.prompt, tt.model)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, spec)
				// Basic validation that it looks like YAML
				assert.Contains(t, spec, "openapi:")
				assert.Contains(t, spec, "paths:")
			}
		})
	}
}
