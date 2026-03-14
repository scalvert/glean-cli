package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdActivity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "activity",
		Short: "Report user activity and feedback to Glean",
		Long: `Report user activity and feedback to Glean.

Send activity events (views, clicks, feedback) to improve Glean's search personalization.

Example:
  glean activity report --json '{"events":[{"action":"VIEW","docId":{"datasource":"confluence","objectId":"12345"}}]}'`,
	}
	cmd.AddCommand(
		newActivityReportCmd(),
		newActivityFeedbackCmd(),
	)
	return cmd
}

func newActivityReportCmd() *cobra.Command {
	var jsonPayload string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Report user activity events",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.Activity
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
			_, err = sdk.Client.Activity.Report(cmd.Context(), req)
			return err
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body (required)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print request without sending")
	return cmd
}

func newActivityFeedbackCmd() *cobra.Command {
	var jsonPayload string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "feedback",
		Short: "Submit feedback",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.Feedback
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
			_, err = sdk.Client.Activity.Feedback(cmd.Context(), nil, &req)
			return err
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body (required)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print request without sending")
	return cmd
}
