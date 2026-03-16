package llm

import (
	"encoding/json"
	"testing"

	"github.com/gleanwork/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// sdkChatResponse mirrors the components.ChatResponse shape the SDK will parse.
type sdkChatResponse struct {
	Messages []struct {
		Author    string `json:"author"`
		Fragments []struct {
			Text string `json:"text"`
		} `json:"fragments"`
	} `json:"messages"`
}

func makeSuccessResp(text string) []byte {
	resp := sdkChatResponse{
		Messages: []struct {
			Author    string `json:"author"`
			Fragments []struct {
				Text string `json:"text"`
			} `json:"fragments"`
		}{
			{
				Author: "GLEAN_AI",
				Fragments: []struct {
					Text string `json:"text"`
				}{
					{Text: text},
				},
			},
		},
	}
	b, _ := json.Marshal(resp)
	return b
}

//nolint:govet
type testCase struct {
	name        string
	input       string
	prompt      string
	model       string
	errContains string
	mockResp    []byte
	wantErr     bool
}

func TestGenerateOpenAPISpec(t *testing.T) {
	yamlSpec := "openapi: 3.0.0\ninfo:\n  title: Test API\n  version: 1.0.0\npaths:\n  /test:\n    get:\n      description: Test endpoint"

	tests := []testCase{
		{
			name:        "empty response returns error",
			input:       "GET /api/users",
			mockResp:    []byte(`{"messages": []}`),
			wantErr:     true,
			errContains: "no response from LLM",
		},
		{
			name:     "with prompt",
			input:    "GET /api/users - Get a list of users",
			prompt:   "Include authentication details",
			model:    "gpt-4",
			mockResp: makeSuccessResp(yamlSpec),
		},
		{
			name:     "curl command",
			input:    `curl -X POST https://api.example.com/users -d '{"name":"test"}'`,
			model:    "gpt-4",
			mockResp: makeSuccessResp(yamlSpec),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, cleanup := testutils.SetupTestWithResponse(t, tt.mockResp)
			defer cleanup()

			spec, err := GenerateOpenAPISpec(tt.input, tt.prompt, tt.model)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, spec)
				assert.Contains(t, spec, "openapi:")
				assert.Contains(t, spec, "paths:")
			}
		})
	}
}
