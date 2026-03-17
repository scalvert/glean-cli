package cmd

import (
	gleanClient "github.com/gleanwork/glean-cli/internal/client"
	gleamcp "github.com/gleanwork/glean-cli/internal/mcp"
	"github.com/spf13/cobra"
)

// NewCmdMCP creates and returns the mcp command.
func NewCmdMCP() *cobra.Command {
	return &cobra.Command{
		Use:   "mcp",
		Short: "Start a stdio MCP server exposing Glean tools to AI agents",
		Long: `Start a Model Context Protocol (MCP) stdio server.

AI agents (such as Claude Code) can connect to this server to invoke
Glean operations as structured tools without needing the CLI in their context.

Available MCP tools:
  glean_search  — Search company knowledge
  glean_chat    — Ask Glean AI a question
  glean_schema  — Introspect CLI command schemas
  glean_people  — Search for employees

Usage with Claude Code (add to .claude/settings.json):
  {
    "mcpServers": {
      "glean": {
        "command": "glean",
        "args": ["mcp"]
      }
    }
  }`,
		RunE: func(cmd *cobra.Command, args []string) error {
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			return gleamcp.Serve(sdk)
		},
	}
}
