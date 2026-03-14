package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdAgents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agents",
		Short: "Manage and run Glean agents",
		Long: `Manage and run Glean agents.

Agents are AI-powered workflows that can search, reason, and act on your company's knowledge.

Example:
  glean agents list
  glean agents get <agent-id>
  glean agents run --json '{"agentId":"<id>","query":"summarize Q1 results"}'`,
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
	var jsonPayload, outputFormat string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available agents",
		RunE: func(cmd *cobra.Command, args []string) error {
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			var req components.SearchAgentsRequest
			if jsonPayload != "" {
				if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
					return fmt.Errorf("invalid --json: %w", err)
				}
			}
			resp, err := sdk.Client.Agents.List(cmd.Context(), req)
			if err != nil {
				return err
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body")
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format: json, ndjson, text")
	return cmd
}

func newAgentsGetCmd() *cobra.Command {
	var agentID, outputFormat string
	cmd := &cobra.Command{
		Use:   "get <agent-id>",
		Short: "Get an agent by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			agentID = args[0]
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			resp, err := sdk.Client.Agents.Retrieve(cmd.Context(), agentID, nil, nil)
			if err != nil {
				return err
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
		},
	}
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format")
	return cmd
}

func newAgentsSchemasCmd() *cobra.Command {
	var outputFormat string
	cmd := &cobra.Command{
		Use:   "schemas <agent-id>",
		Short: "Get the schemas for an agent",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			resp, err := sdk.Client.Agents.RetrieveSchemas(cmd.Context(), args[0], nil, nil)
			if err != nil {
				return err
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
		},
	}
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format")
	return cmd
}

func newAgentsRunCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run an agent (synchronous)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.AgentRunCreate
			if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
				return fmt.Errorf("invalid --json: %w", err)
			}
			if dryRun {
				return output.WriteJSON(cmd.OutOrStdout(), req)
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			resp, err := sdk.Client.Agents.Run(cmd.Context(), req)
			if err != nil {
				return err
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body (required)")
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print request without sending")
	return cmd
}
