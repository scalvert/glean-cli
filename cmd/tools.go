package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
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
	return cmdutil.Build(cmdutil.Spec[struct{}]{
		Use:   "list",
		Short: "List available tools",
		Run: func(ctx context.Context, sdk *glean.Glean, req struct{}) (any, error) {
			resp, err := sdk.Client.Tools.List(ctx, nil)
			if err != nil {
				return nil, err
			}
			return resp.ToolsListResponse, nil
		},
	})
}

func newToolsRunCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.ToolsCallRequest]{
		Use:          "run",
		Short:        "Run a tool",
		JSONRequired: true,
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
		Run: func(ctx context.Context, sdk *glean.Glean, req components.ToolsCallRequest) (any, error) {
			resp, err := sdk.Client.Tools.Run(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.ToolsCallResponse, nil
		},
	})
}
