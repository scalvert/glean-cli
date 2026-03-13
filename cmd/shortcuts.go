package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdShortcuts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shortcuts",
		Short: "Manage Glean shortcuts (go-links)",
	}
	cmd.AddCommand(
		newShortcutsListCmd(),
		newShortcutsGetCmd(),
		newShortcutsCreateCmd(),
		newShortcutsUpdateCmd(),
		newShortcutsDeleteCmd(),
	)
	return cmd
}

func newShortcutsListCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List shortcuts",
		RunE: func(cmd *cobra.Command, args []string) error {
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			var req components.ListShortcutsPaginatedRequest
			if jsonPayload != "" {
				if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
					return fmt.Errorf("invalid --json: %w", err)
				}
			}
			resp, err := sdk.Client.Shortcuts.List(cmd.Context(), req, nil)
			if err != nil {
				return err
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp.ListShortcutsPaginatedResponse, outputFormat, nil)
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body")
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format: json, ndjson, text")
	return cmd
}

func newShortcutsGetCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a shortcut by alias or ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			var req components.GetShortcutRequestUnion
			if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
				return fmt.Errorf("invalid --json: %w", err)
			}
			resp, err := sdk.Client.Shortcuts.Retrieve(cmd.Context(), req, nil)
			if err != nil {
				return err
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body (required)")
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format")
	return cmd
}

func newShortcutsCreateCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a shortcut",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.CreateShortcutRequest
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
			resp, err := sdk.Client.Shortcuts.Create(cmd.Context(), req, nil)
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

func newShortcutsUpdateCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a shortcut",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.UpdateShortcutRequest
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
			resp, err := sdk.Client.Shortcuts.Update(cmd.Context(), req, nil)
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

func newShortcutsDeleteCmd() *cobra.Command {
	var jsonPayload string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a shortcut",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.DeleteShortcutRequest
			if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
				return fmt.Errorf("invalid --json: %w", err)
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			_, err = sdk.Client.Shortcuts.Delete(cmd.Context(), req, nil)
			return err
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body (required)")
	return cmd
}
