package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/models/operations"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdPins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pins",
		Short: "Manage Glean pins",
		Long: `Manage Glean pins.

Pins are manually promoted search results that appear at the top for specific queries.

Example:
  glean pins list
  glean pins create --json '{"query":"onboarding","url":"https://wiki.example.com/onboarding"}'`,
	}
	cmd.AddCommand(
		newPinsListCmd(),
		newPinsGetCmd(),
		newPinsCreateCmd(),
		newPinsUpdateCmd(),
		newPinsRemoveCmd(),
	)
	return cmd
}

func newPinsListCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List pins",
		RunE: func(cmd *cobra.Command, args []string) error {
			var req operations.ListpinsRequestBody
			if jsonPayload != "" {
				if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
					return fmt.Errorf("invalid --json: %w", err)
				}
			}
			if dryRun {
				return output.WriteJSON(cmd.OutOrStdout(), req)
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			resp, err := sdk.Client.Pins.List(cmd.Context(), req, nil)
			if err != nil {
				return err
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body")
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print request without sending")
	return cmd
}

func newPinsGetCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a pin",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.GetPinRequest
			if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
				return fmt.Errorf("invalid --json: %w", err)
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			resp, err := sdk.Client.Pins.Retrieve(cmd.Context(), req, nil)
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

func newPinsCreateCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a pin",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.PinRequest
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
			resp, err := sdk.Client.Pins.Create(cmd.Context(), req, nil)
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

func newPinsUpdateCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a pin",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.EditPinRequest
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
			resp, err := sdk.Client.Pins.Update(cmd.Context(), req, nil)
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

func newPinsRemoveCmd() *cobra.Command {
	var jsonPayload string
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a pin",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.Unpin
			if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
				return fmt.Errorf("invalid --json: %w", err)
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			_, err = sdk.Client.Pins.Remove(cmd.Context(), req, nil)
			return err
		},
	}
	cmd.Flags().StringVar(&jsonPayload, "json", "", "JSON request body (required)")
	return cmd
}
