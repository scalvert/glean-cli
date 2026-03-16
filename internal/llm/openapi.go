package llm

import (
	"context"
	"fmt"
	"strings"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/gleanwork/glean-cli/internal/client"
)

// GenerateOpenAPISpec generates an OpenAPI specification from the given input
func GenerateOpenAPISpec(input, prompt, model string) (string, error) {
	sdk, err := gleanClient.NewFromConfig()
	if err != nil {
		return "", fmt.Errorf("failed to create client: %w", err)
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

	agentGPT := components.AgentEnumGpt
	modeDefault := components.ModeDefault

	chatReq := components.ChatRequest{
		Messages: []components.ChatMessage{
			{
				Author:      components.AuthorUser.ToPointer(),
				MessageType: components.MessageTypeContext.ToPointer(),
				Fragments:   []components.ChatMessageFragment{{Text: &systemPrompt}},
			},
			{
				Author:      components.AuthorUser.ToPointer(),
				MessageType: components.MessageTypeContent.ToPointer(),
				Fragments:   []components.ChatMessageFragment{{Text: &input}},
			},
		},
		AgentConfig: &components.AgentConfig{
			Agent: agentGPT.ToPointer(),
			Mode:  modeDefault.ToPointer(),
		},
	}

	// Use context.Background() since this function doesn't have a context parameter yet
	// (context propagation is handled at the cmd layer for streaming; this is non-streaming)
	resp, err := sdk.Client.Chat.Create(context.Background(), chatReq, nil, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate OpenAPI spec: %w", err)
	}

	if resp.ChatResponse == nil || len(resp.ChatResponse.Messages) == 0 {
		return "", fmt.Errorf("no response from LLM")
	}

	var spec strings.Builder
	lastMsg := resp.ChatResponse.Messages[len(resp.ChatResponse.Messages)-1]
	for _, fragment := range lastMsg.Fragments {
		if fragment.Text != nil {
			spec.WriteString(*fragment.Text)
		}
	}

	return strings.TrimSpace(spec.String()), nil
}
