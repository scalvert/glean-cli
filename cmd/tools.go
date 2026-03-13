package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdTools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tools",
		Short: "List and run Glean tools",
	}
	cmd.AddCommand(
		newToolsListCmd(),
		newToolsRunCmd(),
	)
	return cmd
}

func newToolsListCmd() *cobra.Command {
	var outputFormat string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available tools",
		RunE: func(cmd *cobra.Command, args []string) error {
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
	return cmd
}

func newToolsRunCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
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
