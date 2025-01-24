package llm

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/glean_client"
)

// GenerateOpenAPISpec takes an API description or curl command and generates an OpenAPI spec
func GenerateOpenAPISpec(input, prompt, model string) (string, error) {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	// Create Glean client
	clientConfig := glean_client.NewConfiguration()
	clientConfig.Host = cfg.GleanHost
	clientConfig.Scheme = "https"
	clientConfig.DefaultHeader = map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", cfg.GleanToken),
	}
	client := glean_client.NewAPIClient(clientConfig)

	// Base prompt for OpenAPI generation
	systemPrompt := heredoc.Doc(`
		You are an expert at converting API descriptions and curl commands into OpenAPI specifications.
		Generate a valid OpenAPI 3.0 specification in YAML format based on the input provided.
		Include reasonable defaults for schema types and ensure the spec is complete and valid.
		Focus on accuracy and practicality.
	`)

	// Helper function to create string pointer
	strPtr := func(s string) *string {
		return &s
	}

	// Create chat messages
	messages := []glean_client.ChatMessage{
		{
			Author:      strPtr("SYSTEM"),
			MessageType: strPtr("CONTENT"),
			Fragments: []glean_client.ChatMessageFragment{
				{Text: strPtr(systemPrompt)},
			},
		},
		{
			Author:      strPtr("USER"),
			MessageType: strPtr("CONTENT"),
			Fragments: []glean_client.ChatMessageFragment{
				{Text: strPtr(input)},
			},
		},
	}

	// Add custom prompt if provided
	if prompt != "" {
		messages = append(messages, glean_client.ChatMessage{
			Author:      strPtr("USER"),
			MessageType: strPtr("CONTENT"),
			Fragments: []glean_client.ChatMessageFragment{
				{Text: strPtr(prompt)},
			},
		})
	}

	// Create chat request with GPT agent config
	req := client.ChatAPI.Chat(context.Background())
	stream := false
	req = req.Payload(glean_client.ChatRequest{
		Messages: messages,
		Stream:   &stream,
		AgentConfig: &glean_client.AgentConfig{
			Agent: strPtr("GPT"),
			Mode:  strPtr("DEFAULT"),
		},
	})

	// Execute request
	resp, _, err := req.Execute()
	if err != nil {
		return "", fmt.Errorf("failed to generate OpenAPI spec: %w", err)
	}

	// Get the generated spec from the response
	if len(resp.Messages) == 0 {
		return "", fmt.Errorf("no response received from LLM")
	}

	// Combine all fragments from the last message
	lastMessage := resp.Messages[len(resp.Messages)-1]
	var spec string
	for _, fragment := range lastMessage.Fragments {
		if fragment.Text != nil {
			spec += *fragment.Text
		}
	}

	return spec, nil
}
