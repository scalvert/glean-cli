package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

// contentOnly extracts the user-visible text from a chat NDJSON stream.
// It skips UPDATE/CONTROL/DEBUG messages (internal reasoning/progress)
// and only collects CONTENT message fragments, stripping stage preamble lines.
func contentOnly(ndjson string) string {
	var sb strings.Builder
	for _, line := range strings.Split(ndjson, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var resp components.ChatResponse
		if err := json.Unmarshal([]byte(line), &resp); err != nil {
			continue
		}
		for _, msg := range resp.Messages {
			// Only CONTENT messages carry user-facing text.
			if msg.MessageType != nil && *msg.MessageType != components.MessageTypeContent {
				continue
			}
			for _, frag := range msg.Fragments {
				if frag.Text != nil && *frag.Text != "" {
					sb.WriteString(*frag.Text)
				}
			}
		}
	}
	return sb.String()
}

// NewCmdChat creates and returns the chat command.
func NewCmdChat() *cobra.Command {
	var timeoutMillis int
	var saveChat bool
	var jsonPayload string
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "chat [message]",
		Short: "Have a conversation with Glean's chat API",
		Long: `Have a conversation with Glean's chat API.

The chat API allows you to have natural language conversations with Glean's AI.

Example:
  glean chat "What are the company holidays?"
  glean chat --json '{"messages":[{"author":"USER","messageType":"CONTENT","fragments":[{"text":"What is Glean?"}]}]}'
  glean chat --dry-run "test question"`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" && len(args) == 0 {
				return fmt.Errorf("requires a message argument or --json payload")
			}
			if jsonPayload != "" {
				var chatReq components.ChatRequest
				if err := json.Unmarshal([]byte(jsonPayload), &chatReq); err != nil {
					return fmt.Errorf("invalid --json payload: %w", err)
				}
				stream := true
				chatReq.Stream = &stream
				if dryRun {
					return output.WriteJSON(cmd.OutOrStdout(), chatReq)
				}
				return executeChat(cmd, chatReq, false)
			}
			if dryRun {
				// Show what would be sent
				q := args[0]
				timeout := int64(timeoutMillis)
				save := saveChat
				stream := true
				agentDefault := components.AgentEnumDefault
				modeDefault := components.ModeDefault
				authorUser := components.AuthorUser
				chatReq := components.ChatRequest{
					Messages: []components.ChatMessage{{
						Author:      authorUser.ToPointer(),
						MessageType: components.MessageTypeContent.ToPointer(),
						Fragments:   []components.ChatMessageFragment{{Text: &q}},
					}},
					AgentConfig:   &components.AgentConfig{Agent: agentDefault.ToPointer(), Mode: modeDefault.ToPointer()},
					SaveChat:      &save,
					TimeoutMillis: &timeout,
					Stream:        &stream,
				}
				return output.WriteJSON(cmd.OutOrStdout(), chatReq)
			}
			return executeChatMessage(cmd, args[0], timeoutMillis, saveChat)
		},
	}

	cmd.Flags().StringVar(&jsonPayload, "json", "", "Complete JSON chat request body (overrides all other flags)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print the request body without sending it")
	cmd.Flags().IntVar(&timeoutMillis, "timeout", 60000, "Request timeout in milliseconds (default 60000 — 60 seconds)")
	cmd.Flags().BoolVar(&saveChat, "save", true, "Save the chat for later continuation")

	return cmd
}

// executeChatMessage builds a ChatRequest from a plain message string and calls executeChat.
func executeChatMessage(cmd *cobra.Command, question string, timeoutMillis int, saveChat bool) error {
	agentDefault := components.AgentEnumDefault
	modeDefault := components.ModeDefault
	authorUser := components.AuthorUser
	timeout := int64(timeoutMillis)
	save := saveChat
	stream := true

	chatReq := components.ChatRequest{
		Messages: []components.ChatMessage{
			{
				Author:      authorUser.ToPointer(),
				MessageType: components.MessageTypeContent.ToPointer(),
				Fragments:   []components.ChatMessageFragment{{Text: &question}},
			},
		},
		AgentConfig: &components.AgentConfig{
			Agent: agentDefault.ToPointer(),
			Mode:  modeDefault.ToPointer(),
		},
		SaveChat:      &save,
		TimeoutMillis: &timeout,
		Stream:        &stream,
	}
	return executeChat(cmd, chatReq, true)
}

// executeChat sends a ChatRequest and writes the response to cmd.OutOrStdout().
// Only CONTENT messages are shown — UPDATE/CONTROL/DEBUG (internal reasoning)
// are filtered. The full response is collected before printing so agents get
// a clean, complete string without per-token newlines.
func executeChat(cmd *cobra.Command, chatReq components.ChatRequest, showSpinner bool) error {
	ctx := cmd.Context()

	sdk, err := gleanClient.NewFromConfig()
	if err != nil {
		return err
	}

	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spin.Prefix = "  "
	if showSpinner {
		spin.Start()
		defer spin.Stop()
	}

	resp, err := sdk.Client.Chat.CreateStream(ctx, chatReq, nil)
	if err != nil {
		return fmt.Errorf("chat request failed: %w", err)
	}
	spin.Stop()

	if resp.ChatRequestStream == nil {
		return nil
	}

	text := contentOnly(*resp.ChatRequestStream)
	if text == "" {
		return nil
	}
	fmt.Fprintln(cmd.OutOrStdout(), strings.TrimSpace(text))
	return nil
}

