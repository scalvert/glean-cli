package cmd

import (
	"context"
	"io"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/gleanwork/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

// agentIDRequest is a CLI-only request struct for commands that take an agent
// ID as their only input. Using camelCase JSON tag for CLI consistency.
type agentIDRequest struct {
	AgentID string `json:"agentId"`
}

func NewCmdAgents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agents",
		Short: "Manage and run Glean agents",
		Long: `Manage and run Glean agents.

Agents are AI-powered workflows that can search, reason, and act on your company's knowledge.

Example:
  glean agents list
  glean agents get --json '{"agentId":"<id>"}'
  glean agents schemas --json '{"agentId":"<id>"}'
  glean agents run --json '{"agentId":"<id>","messages":[{"author":"USER","fragments":[{"text":"summarize Q1 results"}]}]}'`,
	}
	cmd.AddCommand(
		newAgentsListCmd(),
		newAgentsGetCmd(),
		newAgentsSchemasCmd(),
		newAgentsRunCmd(),
	)
	return cmd
}

func newAgentsListCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.SearchAgentsRequest]{
		Use:   "list",
		Short: "List available agents",
		TextFn: func(w io.Writer, v any) error {
			resp, ok := v.(*components.SearchAgentsResponse)
			if !ok {
				return output.WriteJSON(w, v)
			}
			rows := make([][]string, len(resp.Agents))
			for i, a := range resp.Agents {
				desc := ""
				if a.Description != nil {
					desc = output.Truncate(*a.Description, 60)
				}
				rows[i] = []string{a.AgentID, a.Name, desc}
			}
			return output.WriteTable(w, []string{"ID", "NAME", "DESCRIPTION"}, rows)
		},
		Run: func(ctx context.Context, sdk *glean.Glean, req components.SearchAgentsRequest) (any, error) {
			resp, err := sdk.Client.Agents.List(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.SearchAgentsResponse, nil
		},
	})
}

func newAgentsGetCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[agentIDRequest]{
		Use:          "get",
		Short:        "Get an agent by ID",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req agentIDRequest) (any, error) {
			resp, err := sdk.Client.Agents.Retrieve(ctx, req.AgentID, nil, nil)
			if err != nil {
				return nil, err
			}
			return resp.Agent, nil
		},
	})
}

func newAgentsSchemasCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[agentIDRequest]{
		Use:          "schemas",
		Short:        "Get the schemas for an agent",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req agentIDRequest) (any, error) {
			resp, err := sdk.Client.Agents.RetrieveSchemas(ctx, req.AgentID, nil, nil)
			if err != nil {
				return nil, err
			}
			return resp.AgentSchemas, nil
		},
	})
}

func newAgentsRunCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.AgentRunCreate]{
		Use:          "run",
		Short:        "Run an agent (synchronous)",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.AgentRunCreate) (any, error) {
			resp, err := sdk.Client.Agents.Run(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.AgentRunWaitResponse, nil
		},
	})
}
