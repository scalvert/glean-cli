package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/gleanwork/glean-cli/internal/client"
	"github.com/gleanwork/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdTools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tools",
		Short: "List and run Glean tools",
		Long: `List and run Glean tools.

Tools are callable functions exposed by the Glean platform for agent workflows.

Example:
  glean tools list
  glean tools run --json '{"toolName":"search","parameters":{"query":"Q1 results"}}'`,
	}
	cmd.AddCommand(
		newToolsListCmd(),
		newToolsRunCmd(),
	)
	return cmd
}

func newToolsListCmd() *cobra.Command {
	var outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available tools",
		RunE: func(cmd *cobra.Command, args []string) error {
			if dryRun {
				return output.WriteJSON(cmd.OutOrStdout(), struct{}{})
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			resp, err := sdk.Client.Tools.List(cmd.Context(), nil)
			if err != nil {
				return err
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
		},
	}
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print request without sending")
	return cmd
}

func newToolsRunCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a tool",
		Long: `Run a tool by name with typed parameters.

The --json payload must match the ToolsCallRequest schema:

  {
    "name": "toolName",
    "parameters": {
      "paramName": {"name": "paramName", "value": "paramValue"},
      "nested": {"name": "nested", "properties": {"key": {"name": "key", "value": "val"}}}
    }
  }

Each parameter entry is a ToolsCallParameter with fields:
  name        (string, required) - parameter name
  value       (string)           - value for primitive types
  items       (array)            - value for array types (each element is a ToolsCallParameter)
  properties  (object)           - value for object types (map of string to ToolsCallParameter)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required\n\nRun '%s --help' for the expected payload format", cmd.CommandPath())
			}
			var req components.ToolsCallRequest
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
			resp, err := sdk.Client.Tools.Run(cmd.Context(), req)
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
