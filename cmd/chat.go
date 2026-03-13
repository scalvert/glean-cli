package cmd

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/theme"
	"github.com/scalvert/glean-cli/internal/utils"
	"github.com/spf13/cobra"
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

// ChatStageType represents different stages of chat output
type ChatStageType string

// stageInfo represents a parsed chat stage
type stageInfo struct {
	stage  ChatStageType
	detail string
}

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

// executeChat handles the chat interaction with Glean's API.
func executeChat(cmd *cobra.Command, question string, timeoutMillis int, saveChat bool) error {
	ctx := cmd.Context()

	sdk, err := gleanClient.NewFromConfig()
	if err != nil {
		return err
	}

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
				Fragments: []components.ChatMessageFragment{
					{Text: &question},
				},
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

	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spin.Prefix = "Waiting for response "
	spin.Start()
	defer spin.Stop()

	resp, err := sdk.Client.Chat.CreateStream(ctx, chatReq, nil)
	if err != nil {
		return fmt.Errorf("chat request failed: %w", err)
	}

	if resp.ChatRequestStream == nil {
		return nil
	}

	state := &ChatState{
		cmd:       cmd,
		firstLine: true,
	}

	lines := strings.Split(*resp.ChatRequestStream, "\n")
	for _, line := range lines {
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

// processChatResponse processes a single line of chat response JSON.
func (s *ChatState) processChatResponse(line string) error {
	var chatResp components.ChatResponse
	if err := json.Unmarshal([]byte(line), &chatResp); err != nil {
		return fmt.Errorf("error parsing response line: %w", err)
	}

	for _, msg := range chatResp.Messages {
		hasMore := msg.HasMoreFragments != nil && *msg.HasMoreFragments
		for _, fragment := range msg.Fragments {
			s.processFragment(fragment, hasMore)
		}
	}

	return nil
}

// processFragment handles a single chat message fragment.
func (s *ChatState) processFragment(fragment components.ChatMessageFragment, hasMoreFragments bool) {
	if len(fragment.StructuredResults) > 0 {
		if s.readingStage == nil {
			fmt.Fprintln(s.cmd.OutOrStdout(), formatChatStage(StageReading, formatReadingStage(fragment.StructuredResults)))
			s.isStageOutput = true
		}
		s.readingStage = nil
		return
	}

	text := ""
	if fragment.Text != nil {
		text = *fragment.Text
	}

	if text == "" {
		return
	}

	if stage := isStage(text); stage != nil {
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
		fmt.Fprintln(s.cmd.OutOrStdout(), formatChatStage(s.searchStage.stage, text))
		s.searchStage = nil
		s.isStageOutput = true
	} else {
		if s.isStageOutput {
			fmt.Fprint(s.cmd.OutOrStdout(), formatChatResponse(text))
			s.isStageOutput = false
		} else {
			fmt.Fprint(s.cmd.OutOrStdout(), text)
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

// formatChatStage formats a chat stage output with a colored checkmark.
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

// isStage checks if a text fragment represents a chat stage.
func isStage(text string) *stageInfo {
	stagePatterns := map[string]ChatStageType{
		"**Searching:**": StageSearching,
		"**Reading:**":   StageReading,
		"**Writing:**":   StageWriting,
	}

	for pattern, stageType := range stagePatterns {
		if strings.HasPrefix(text, pattern) {
			detail := strings.TrimPrefix(text, pattern)
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

// formatReadingStage formats the reading stage with document details.
func formatReadingStage(sources []components.StructuredResult) string {
	if len(sources) == 0 {
		return "Reading: no sources found"
	}

	sourcesByType := make(map[string][]components.StructuredResult)
	for _, source := range sources {
		if source.Document != nil {
			ds := ""
			if source.Document.Datasource != nil {
				ds = *source.Document.Datasource
			} else if source.Document.Metadata != nil && source.Document.Metadata.Datasource != nil {
				ds = *source.Document.Metadata.Datasource
			}
			sourcesByType[ds] = append(sourcesByType[ds], source)
		}
	}

	datasources := make([]string, 0, len(sourcesByType))
	for ds := range sourcesByType {
		datasources = append(datasources, ds)
	}
	sort.Strings(datasources)

	var details []string
	for _, ds := range datasources {
		for _, source := range sourcesByType[ds] {
			doc := source.Document
			title := ""
			if doc.Title != nil {
				title = *doc.Title
			}
			docURL := ""
			if doc.URL != nil {
				docURL = *doc.URL
			}
			if title != "" {
				details = append(details, fmt.Sprintf("%s: %s (%s)",
					utils.FormatDatasource(ds),
					title,
					utils.MaybeAnonymizeURL(docURL)))
			} else {
				details = append(details, fmt.Sprintf("%s: %s",
					utils.FormatDatasource(ds),
					utils.MaybeAnonymizeURL(docURL)))
			}
		}
	}

	return strings.Join(details, "\n         ")
}
