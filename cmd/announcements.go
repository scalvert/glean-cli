package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/gleanwork/glean-cli/internal/client"
	"github.com/gleanwork/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdAnnouncements() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "announcements",
		Short: "Manage Glean announcements",
		Long: `Manage Glean announcements.

Announcements are time-bounded notices that appear prominently in Glean search results.

Example:
  glean announcements create --json '{"title":"Q2 kickoff","startTime":"2026-04-01T00:00:00Z","endTime":"2026-04-30T00:00:00Z"}'`,
	}
	cmd.AddCommand(
		newAnnouncementsCreateCmd(),
		newAnnouncementsUpdateCmd(),
		newAnnouncementsDeleteCmd(),
	)
	return cmd
}

func newAnnouncementsCreateCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an announcement",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.CreateAnnouncementRequest
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
			resp, err := sdk.Client.Announcements.Create(cmd.Context(), req, nil)
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

func newAnnouncementsUpdateCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an announcement",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.UpdateAnnouncementRequest
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
			resp, err := sdk.Client.Announcements.Update(cmd.Context(), req, nil)
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

func newAnnouncementsDeleteCmd() *cobra.Command {
	var jsonPayload string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an announcement",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.DeleteAnnouncementRequest
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
			_, err = sdk.Client.Announcements.Delete(cmd.Context(), req, nil)
			return err
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body (required)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print request without sending")
	return cmd
}
