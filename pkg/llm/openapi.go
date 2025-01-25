package llm

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/glean_client"
)

func GenerateOpenAPISpec(input, prompt, model string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("configuration error: %w", err)
	}

	if cfg.GleanEmail == "" {
		return "", fmt.Errorf("email not configured. Run 'glean config --email <your-email>'")
	}

	clientConfig := glean_client.NewConfiguration()
	clientConfig.Host = cfg.GleanHost
	clientConfig.Scheme = "https"
	clientConfig.DefaultHeader = map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", cfg.GleanToken),
	}
	client := glean_client.NewAPIClient(clientConfig)

	systemPrompt := heredoc.Doc(`
		You are an expert at converting API descriptions and curl commands into OpenAPI specifications.
		Generate a valid OpenAPI 3.0 specification in YAML format based on the input provided.
		Include reasonable defaults for schema types and ensure the spec is complete and valid.
		Focus on accuracy and practicality.

		When returning the response, only return the OpenAPI spec, no other text, comments, and don't wrap the response in a code block.
	`)

	strPtr := func(s string) *string {
		return &s
	}

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

	req := client.ChatAPI.Chat(context.Background())
	stream := false
	req = req.XScioActas(cfg.GleanEmail)
	req = req.Payload(glean_client.ChatRequest{
		Messages: messages,
		Stream:   &stream,
		AgentConfig: &glean_client.AgentConfig{
			Agent: strPtr("GPT"),
			Mode:  strPtr("DEFAULT"),
		},
	})

	resp, _, err := req.Execute()
	if err != nil {
		return "", err
	}

	if len(resp.Messages) == 0 {
		return "", fmt.Errorf("no response received from LLM")
	}

	lastMessage := resp.Messages[len(resp.Messages)-1]
	var spec string
	for _, fragment := range lastMessage.Fragments {
		if fragment.Text != nil {
			spec += *fragment.Text
		}
	}

	return spec, nil
}
