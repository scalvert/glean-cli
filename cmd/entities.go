package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdEntities() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "entities",
		Short: "List and read Glean entities and people",
		Long: `List and read Glean entities and people.

Entities represent structured objects in your company's knowledge graph — people, teams, projects, and more.

Example:
  glean entities list --json '{"entityType":"PERSON","query":"engineering"}'
  glean entities get <entity-id>`,
	}
	cmd.AddCommand(
		newEntitiesListCmd(),
		newEntitiesReadPeopleCmd(),
	)
	return cmd
}

func newEntitiesListCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List entities",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.ListEntitiesRequest
			if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
				if strings.Contains(err.Error(), "ListEntitiesRequestEntityType") {
					return fmt.Errorf("invalid entityType: valid values are PEOPLE, TEAMS, CUSTOM_ENTITIES")
				}
				return fmt.Errorf("invalid --json: %w", err)
			}
			if dryRun {
				return output.WriteJSON(cmd.OutOrStdout(), req)
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			resp, err := sdk.Client.Entities.List(cmd.Context(), req, nil)
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

func newEntitiesReadPeopleCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "read-people",
		Short: "Read people records",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.PeopleRequest
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
			resp, err := sdk.Client.Entities.ReadPeople(cmd.Context(), req, nil)
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
