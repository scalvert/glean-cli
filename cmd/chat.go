package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/scalvert/glean-cli/pkg/api"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/theme"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

const (
	StageSearching ChatStageType = "Searching"
	StageReading   ChatStageType = "Reading"
	StageWriting   ChatStageType = "Writing"
	StageSummary   ChatStageType = "" // Empty string since we use the exact text
)

var (
	// summarizePattern matches variations of summarize/summarizing at start of text
	summarizePattern = regexp.MustCompile(`^(?i)summariz(e|ing)\b`)
)

// ChatMessage represents a message in a chat conversation with fragments of text.
type ChatMessage struct {
	Author      string `json:"author"`      // USER or GLEAN_AI
	MessageType string `json:"messageType"` // e.g., CONTENT
	Fragments   []struct {
		Text string `json:"text"`
	} `json:"fragments"`
}

// ChatRequest represents a request to the Glean chat API.
type ChatRequest struct {
	AgentConfig   AgentConfig   `json:"agentConfig"`
	ApplicationID string        `json:"applicationId,omitempty"`
	ChatID        string        `json:"chatId,omitempty"`
	Messages      []ChatMessage `json:"messages"`
	TimeoutMillis int           `json:"timeoutMillis"`
	Stream        bool          `json:"stream"`
	SaveChat      bool          `json:"saveChat"`
}

// AgentConfig configures the behavior of the chat agent.
type AgentConfig struct {
	Agent string `json:"agent"` // e.g., GPT
	Mode  string `json:"mode"`  // e.g., DEFAULT
}

// ChatResponse represents a response from the Glean chat API.
type ChatResponse struct {
	ChatSessionTrackingToken string `json:"chatSessionTrackingToken"`
	Messages                 []struct {
		Author    string `json:"author"`
		Fragments []struct {
			Text              string                 `json:"text"`
			StructuredResults []api.StructuredResult `json:"structuredResults,omitempty"`
		} `json:"fragments"`
		HasMoreFragments bool `json:"hasMoreFragments,omitempty"`
	} `json:"messages"`
}

// stageInfo represents a parsed chat stage
type stageInfo struct {
	stage  ChatStageType
	detail string
}

// ChatStageType represents different stages of chat output
type ChatStageType string

// ChatState holds the state for processing chat responses
type ChatState struct {
	cmd           *cobra.Command
	searchStage   *stageInfo
	readingStage  *stageInfo
	isStageOutput bool
	firstLine     bool
}

// NewCmdChat creates and returns the chat command.
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

// sendChatRequest creates and sends a chat request to the Glean API
func sendChatRequest(client http.Client, question string, timeoutMillis int, saveChat bool) (io.ReadCloser, error) {
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

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	httpReq := &http.Request{
		Method: "POST",
		Path:   "chat",
		Body:   json.RawMessage(jsonBytes),
		Stream: true,
	}

	return client.SendStreamingRequest(httpReq)
}

// processFragment handles a single chat message fragment
func (s *ChatState) processFragment(fragment struct {
	Text              string                 `json:"text"`
	StructuredResults []api.StructuredResult `json:"structuredResults,omitempty"`
}, hasMoreFragments bool) {
	if len(fragment.StructuredResults) > 0 {
		if s.readingStage == nil {
			fmt.Fprintln(s.cmd.OutOrStdout(), formatChatStage(StageReading, formatReadingStage(fragment.StructuredResults)))
			s.isStageOutput = true
		}
		s.readingStage = nil
		return
	}

	if fragment.Text == "" {
		return
	}

	if stage := isStage(fragment.Text); stage != nil {
		if stage.stage == StageReading {
			s.readingStage = stage
			return
		}
		if stage.stage == StageSearching && stage.detail == "" {
			s.searchStage = stage
			return
		}
		fmt.Fprintln(s.cmd.OutOrStdout(), formatChatStage(stage.stage, stage.detail))
		if stage.stage == StageSummary {
			fmt.Fprintln(s.cmd.OutOrStdout())
		}
		if stage.stage != StageSummary {
			s.isStageOutput = true
		}
	} else if s.searchStage != nil {
		fmt.Fprintln(s.cmd.OutOrStdout(), formatChatStage(s.searchStage.stage, fragment.Text))
		s.searchStage = nil
		s.isStageOutput = true
	} else {
		if s.isStageOutput {
			fmt.Fprint(s.cmd.OutOrStdout(), formatChatResponse(fragment.Text))
			s.isStageOutput = false
		} else {
			fmt.Fprint(s.cmd.OutOrStdout(), fragment.Text)
			if !hasMoreFragments {
				fmt.Println()
				if s.firstLine {
					fmt.Println()
					s.firstLine = false
				}
			}
		}
	}
}

// processChatResponse processes a single line of chat response
func (s *ChatState) processChatResponse(line string) error {
	var chatResp ChatResponse
	if err := json.Unmarshal([]byte(line), &chatResp); err != nil {
		return fmt.Errorf("error parsing response line: %w", err)
	}

	for _, msg := range chatResp.Messages {
		for _, fragment := range msg.Fragments {
			s.processFragment(fragment, msg.HasMoreFragments)
		}
	}

	return nil
}

// executeChat handles the chat interaction with Glean's API, streaming the response
// and formatting the output for the terminal.
func executeChat(cmd *cobra.Command, question string, timeoutMillis int, saveChat bool) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client, err := http.NewClient(cfg)
	if err != nil {
		return err
	}

	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spin.Prefix = "Waiting for response "
	spin.Start()
	defer spin.Stop()

	responseBody, err := sendChatRequest(client, question, timeoutMillis, saveChat)
	if err != nil {
		return err
	}
	defer responseBody.Close()

	state := &ChatState{
		cmd:       cmd,
		firstLine: true,
	}

	reader := bufio.NewReader(responseBody)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading response: %w", err)
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if state.firstLine {
			spin.Stop()
		}

		if err := state.processChatResponse(line); err != nil {
			return err
		}
	}

	return nil
}

// formatChatStage formats a chat stage output with a colored checkmark and appropriate spacing.
func formatChatStage(stage ChatStageType, detail string) string {
	const check = "✓"
	if stage == StageSummary {
		return fmt.Sprintf("%s %s", theme.Blue(check), detail)
	}
	return fmt.Sprintf("%s %s: %s", theme.Blue(check), stage, detail)
}

// formatChatResponse formats the final chat response with proper spacing and divider.
func formatChatResponse(response string) string {
	const divider = "\n----------------------------------------\n\n"
	return fmt.Sprintf("%s%s", divider, response)
}

// isStage checks if a text fragment represents a chat stage and returns the stage info
func isStage(text string) *stageInfo {
	stagePatterns := map[string]ChatStageType{
		"**Searching:**": StageSearching,
		"**Reading:**":   StageReading,
		"**Writing:**":   StageWriting,
	}

	for pattern, stageType := range stagePatterns {
		if strings.HasPrefix(text, pattern) {
			detail := strings.TrimPrefix(text, pattern)
			// Special case for searching stage which may have an empty detail
			// as it comes in a separate fragment
			if stageType == StageSearching && strings.TrimSpace(detail) == "" {
				return &stageInfo{stage: stageType}
			}
			detail = strings.TrimSpace(detail)
			return &stageInfo{stage: stageType, detail: detail}
		}
	}

	if summarizePattern.MatchString(text) {
		return &stageInfo{stage: StageSummary, detail: text}
	}

	return nil
}

// formatReadingStage formats the reading stage with document details
func formatReadingStage(sources []api.StructuredResult) string {
	if len(sources) == 0 {
		return "Reading: no sources found"
	}

	// Group sources by datasource for better organization
	sourcesByType := make(map[string][]api.StructuredResult)
	for _, source := range sources {
		if source.Document != nil {
			ds := source.Document.Metadata.Datasource
			sourcesByType[ds] = append(sourcesByType[ds], source)
		}
	}

	datasources := maps.Keys(sourcesByType)
	sort.Strings(datasources)

	var details []string
	for _, ds := range datasources {
		sources := sourcesByType[ds]
		for _, source := range sources {
			doc := source.Document
			if doc.Title != "" {
				details = append(details, fmt.Sprintf("%s: %s (%s)",
					formatDatasource(ds),
					doc.Title,
					doc.URL))
			} else {
				details = append(details, fmt.Sprintf("%s: %s",
					formatDatasource(ds),
					doc.URL))
			}
		}
	}

	return strings.Join(details, "\n         ")
}
