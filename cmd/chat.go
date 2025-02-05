package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/spf13/cobra"
)

type ChatMessage struct {
	Author      string `json:"author"`
	MessageType string `json:"messageType"`
	Fragments   []struct {
		Text string `json:"text"`
	} `json:"fragments"`
}

type ChatRequest struct {
	AgentConfig   AgentConfig   `json:"agentConfig"`
	ApplicationID string        `json:"applicationId,omitempty"`
	ChatID        string        `json:"chatId,omitempty"`
	Messages      []ChatMessage `json:"messages"`
	TimeoutMillis int           `json:"timeoutMillis"`
	Stream        bool          `json:"stream"`
	SaveChat      bool          `json:"saveChat"`
}

type AgentConfig struct {
	Agent string `json:"agent"`
	Mode  string `json:"mode"`
}

type ChatResponse struct {
	ChatSessionTrackingToken string `json:"chatSessionTrackingToken"`
	Messages                 []struct {
		Author    string `json:"author"`
		Fragments []struct {
			Text string `json:"text"`
		} `json:"fragments"`
		HasMoreFragments bool `json:"hasMoreFragments,omitempty"`
	} `json:"messages"`
}

// cleanMarkdown removes markdown formatting from text
func cleanMarkdown(text string) string {
	// Remove bold/italic markers
	text = regexp.MustCompile(`\*\*`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`\*`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`__`).ReplaceAllString(text, "")
	text = regexp.MustCompile(`_`).ReplaceAllString(text, "")

	// Convert markdown links to plain text
	text = regexp.MustCompile(`\[([^\]]+)\]\([^)]+\)`).ReplaceAllString(text, "$1")

	// Remove backticks
	text = regexp.MustCompile("`").ReplaceAllString(text, "")

	// Remove multiple blank lines
	text = regexp.MustCompile(`\n\s*\n\s*\n`).ReplaceAllString(text, "\n\n")

	return text
}

func NewCmdChat() *cobra.Command {
	var timeoutMillis int
	var saveChat bool

	cmd := &cobra.Command{
		Use:   "chat [message]",
		Short: "Have a conversation with Glean's chat API",
		Long: `Have a conversation with Glean's chat API.

The chat API allows you to have natural language conversations with Glean's AI.
The response will be streamed as it becomes available.

Example:
  glean chat "What are the company holidays?"
  glean chat --timeout 60000 "Tell me about the engineering team"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeChat(cmd, args[0], timeoutMillis, saveChat)
		},
	}

	cmd.Flags().IntVar(&timeoutMillis, "timeout", 30000, "Request timeout in milliseconds")
	cmd.Flags().BoolVar(&saveChat, "save", true, "Save the chat for later continuation")

	return cmd
}

func executeChat(cmd *cobra.Command, question string, timeoutMillis int, saveChat bool) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client, err := http.NewClient(cfg)
	if err != nil {
		return err
	}

	// Create chat request
	req := ChatRequest{
		Messages: []ChatMessage{
			{
				Author:      "USER",
				MessageType: "CONTENT",
				Fragments: []struct {
					Text string `json:"text"`
				}{
					{Text: question},
				},
			},
		},
		Stream: true,
		AgentConfig: AgentConfig{
			Agent: "DEFAULT",
			Mode:  "DEFAULT",
		},
		SaveChat:      saveChat,
		TimeoutMillis: timeoutMillis,
	}

	// Convert request to JSON
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	// Create HTTP request
	httpReq := &http.Request{
		Method: "POST",
		Path:   "chat",
		Body:   json.RawMessage(jsonBytes),
		Stream: true,
	}

	// Start spinner
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spin.Prefix = "Waiting for response "
	spin.Start()
	defer spin.Stop()

	// Send request and get streaming response
	responseBody, err := client.SendStreamingRequest(httpReq)
	if err != nil {
		return err
	}
	defer responseBody.Close()

	// Create a reader for the streaming response
	reader := bufio.NewReader(responseBody)
	firstLine := true

	// Read response line by line
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading response: %w", err)
		}

		// Skip empty lines
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Stop spinner after first line
		if firstLine {
			spin.Stop()
			firstLine = false
		}

		// Parse and print the response
		var chatResp ChatResponse
		if err := json.Unmarshal([]byte(line), &chatResp); err != nil {
			return fmt.Errorf("error parsing response line: %w", err)
		}

		// Print each message
		for _, msg := range chatResp.Messages {
			for _, fragment := range msg.Fragments {
				cleanedText := cleanMarkdown(fragment.Text)
				if cleanedText != "" {
					fmt.Fprint(cmd.OutOrStdout(), cleanedText)
					if !msg.HasMoreFragments {
						fmt.Fprintln(cmd.OutOrStdout())
					}
				}
			}
		}
	}

	return nil
}
