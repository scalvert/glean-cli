package llm

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/scalvert/glean-cli/internal/config"
	"github.com/scalvert/glean-cli/internal/http"
)

type chatRequest struct {
	AgentConfig struct {
		Agent string `json:"agent"`
		Mode  string `json:"mode"`
	} `json:"agentConfig"`
	Messages []message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type message struct {
	Author      string     `json:"author"`
	MessageType string     `json:"messageType"`
	Fragments   []fragment `json:"fragments"`
}

type fragment struct {
	Text string `json:"text"`
}

type chatResponse struct {
	Messages []struct {
		Fragments []fragment `json:"fragments"`
	} `json:"messages"`
}

// GenerateOpenAPISpec generates an OpenAPI specification from the given input
func GenerateOpenAPISpec(input, prompt, model string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	client, err := http.NewClient(cfg)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP client: %w", err)
	}

	systemPrompt := `You are an expert in OpenAPI specifications. Generate an OpenAPI 3.0 specification in YAML format based on the provided API description or curl command.

Return ONLY the YAML content with no additional text, markdown formatting, or code blocks. Follow these guidelines:

1. Include detailed descriptions for paths, parameters, and responses
2. Use appropriate data types and formats
3. Include example values where applicable
4. Document error responses
5. Follow OpenAPI 3.0 best practices`

	if prompt != "" {
		systemPrompt += "\n\nAdditional instructions:\n" + prompt
	}

	messages := []message{
		{
			Author:      "SYSTEM",
			MessageType: "CONTENT",
			Fragments:   []fragment{{Text: systemPrompt}},
		},
		{
			Author:      "USER",
			MessageType: "CONTENT",
			Fragments:   []fragment{{Text: input}},
		},
	}

	req := &http.Request{
		Method: "POST",
		Path:   "/rest/api/v1/chat",
		Body: chatRequest{
			Messages: messages,
			Stream:   false,
			AgentConfig: struct {
				Agent string `json:"agent"`
				Mode  string `json:"mode"`
			}{
				Agent: "GPT",
				Mode:  "DEFAULT",
			},
		},
	}

	resp, err := client.SendRequest(req)
	if err != nil {
		return "", fmt.Errorf("failed to generate OpenAPI spec: %w", err)
	}

	var chatResp chatResponse
	if err := json.Unmarshal(resp, &chatResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(chatResp.Messages) == 0 {
		return "", fmt.Errorf("no response from LLM")
	}

	var spec string
	for _, fragment := range chatResp.Messages[len(chatResp.Messages)-1].Fragments {
		spec += fragment.Text
	}

	return strings.TrimSpace(spec), nil
}
