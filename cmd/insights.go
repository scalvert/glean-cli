package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/gleanwork/glean-cli/internal/client"
	"github.com/gleanwork/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdInsights() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insights",
		Short: "Retrieve Glean usage insights",
		Long: `Retrieve Glean usage insights.

Insights provide analytics on how your organization uses Glean — search trends, popular content, and more.

Example:
  glean insights get --json '{"insightType":"SEARCH","timeRange":"LAST_30_DAYS"}'`,
	}
	cmd.AddCommand(newInsightsGetCmd())
	return cmd
}

func newInsightsGetCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get usage insights",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.InsightsRequest
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
			resp, err := sdk.Client.Insights.Retrieve(cmd.Context(), req, nil)
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
