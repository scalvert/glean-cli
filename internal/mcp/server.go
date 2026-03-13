// Package mcp implements a Model Context Protocol (MCP) stdio server that
// exposes Glean operations as MCP tools. Agents can connect to this server
// to invoke glean_search, glean_chat, glean_schema, and glean_people.
package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/scalvert/glean-cli/internal/schema"
)

// NewServer creates a configured MCP server with all Glean tools registered.
func NewServer(sdk *glean.Glean) *server.MCPServer {
	s := server.NewMCPServer(
		"glean",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	s.AddTool(searchTool(), searchHandler(sdk))
	s.AddTool(chatTool(), chatHandler(sdk))
	s.AddTool(schemaTool(), schemaHandler())
	s.AddTool(peopleTool(), peopleHandler(sdk))

	return s
}

// searchTool defines the glean_search MCP tool.
func searchTool() mcp.Tool {
	return mcp.NewTool("glean_search",
		mcp.WithDescription("Search Glean company knowledge. Returns JSON search results."),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search query string"),
		),
		mcp.WithNumber("pageSize",
			mcp.Description("Number of results (default 10)"),
		),
		mcp.WithString("datasource",
			mcp.Description("Filter by datasource name (e.g. confluence, gdrive)"),
		),
	)
}

func searchHandler(sdk *glean.Glean) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := req.GetString("query", "")
		if query == "" {
			return mcp.NewToolResultError("query is required"), nil
		}

		pageSize := int64(req.GetInt("pageSize", 10))
		searchReq := components.SearchRequest{
			Query:    query,
			PageSize: &pageSize,
		}

		resp, err := sdk.Client.Search.Query(ctx, searchReq, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("search failed", err), nil
		}

		result, err := mcp.NewToolResultJSON(resp.SearchResponse)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("marshal failed", err), nil
		}
		return result, nil
	}
}

// chatTool defines the glean_chat MCP tool.
func chatTool() mcp.Tool {
	return mcp.NewTool("glean_chat",
		mcp.WithDescription("Ask Glean AI a question. Returns the full response text."),
		mcp.WithString("message",
			mcp.Required(),
			mcp.Description("The question or message to send to Glean AI"),
		),
	)
}

func chatHandler(sdk *glean.Glean) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		message := req.GetString("message", "")
		if message == "" {
			return mcp.NewToolResultError("message is required"), nil
		}

		agentDefault := components.AgentEnumDefault
		modeDefault := components.ModeDefault
		authorUser := components.AuthorUser
		stream := true

		chatReq := components.ChatRequest{
			Messages: []components.ChatMessage{
				{
					Author:      authorUser.ToPointer(),
					MessageType: components.MessageTypeContent.ToPointer(),
					Fragments:   []components.ChatMessageFragment{{Text: &message}},
				},
			},
			AgentConfig: &components.AgentConfig{
				Agent: agentDefault.ToPointer(),
				Mode:  modeDefault.ToPointer(),
			},
			Stream: &stream,
		}

		resp, err := sdk.Client.Chat.CreateStream(ctx, chatReq, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("chat request failed", err), nil
		}

		if resp.ChatRequestStream == nil {
			return mcp.NewToolResultText("(no response)"), nil
		}

		// Extract text from NDJSON stream
		var textParts []string
		for _, line := range splitLines(*resp.ChatRequestStream) {
			var chatResp components.ChatResponse
			if err := json.Unmarshal([]byte(line), &chatResp); err != nil {
				continue
			}
			for _, msg := range chatResp.Messages {
				for _, frag := range msg.Fragments {
					if frag.Text != nil && *frag.Text != "" {
						textParts = append(textParts, *frag.Text)
					}
				}
			}
		}

		text := ""
		for _, p := range textParts {
			text += p
		}
		return mcp.NewToolResultText(text), nil
	}
}

// schemaTool defines the glean_schema MCP tool.
func schemaTool() mcp.Tool {
	return mcp.NewTool("glean_schema",
		mcp.WithDescription("Get the JSON schema for a glean CLI command. Call with no command to list all commands."),
		mcp.WithString("command",
			mcp.Description("Command name (e.g. search, chat, shortcuts). Omit to list all."),
		),
	)
}

func schemaHandler() server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		command := req.GetString("command", "")

		if command == "" {
			names := schema.List()
			data, _ := json.Marshal(map[string][]string{"commands": names})
			return mcp.NewToolResultText(string(data)), nil
		}

		s, err := schema.Get(command)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("schema not found", err), nil
		}

		data, _ := json.Marshal(s)
		return mcp.NewToolResultText(string(data)), nil
	}
}

// peopleTool defines the glean_people MCP tool.
func peopleTool() mcp.Tool {
	return mcp.NewTool("glean_people",
		mcp.WithDescription("Search for people (employees) in Glean."),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Name, email, or role to search for"),
		),
	)
}

func peopleHandler(sdk *glean.Glean) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := req.GetString("query", "")
		if query == "" {
			return mcp.NewToolResultError("query is required"), nil
		}

		entityType := components.ListEntitiesRequestEntityTypePeople
		peopleReq := components.ListEntitiesRequest{
			Query:      &query,
			EntityType: entityType.ToPointer(),
		}

		resp, err := sdk.Client.Entities.List(ctx, peopleReq, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("people search failed", err), nil
		}

		result, err := mcp.NewToolResultJSON(resp)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("marshal failed", err), nil
		}
		return result, nil
	}
}

func splitLines(s string) []string {
	var lines []string
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

// Serve starts the MCP stdio server — blocks until the client disconnects.
func Serve(sdk *glean.Glean) error {
	s := NewServer(sdk)
	if err := server.ServeStdio(s); err != nil {
		return fmt.Errorf("MCP server error: %w", err)
	}
	return nil
}
