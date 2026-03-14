package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "messages",
		Short: "Retrieve Glean messages",
		Long: `Retrieve Glean messages.

Messages are threaded communications (Slack, Teams, email, etc.) indexed in Glean.

Example:
  glean messages get --json '{"messageIds":["<id>"]}'`,
	}
	cmd.AddCommand(newMessagesGetCmd())
	return cmd
}

func newMessagesGetCmd() *cobra.Command {
	var jsonPayload, outputFormat string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" {
				return fmt.Errorf("--json is required")
			}
			var req components.MessagesRequest
			if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
				return fmt.Errorf("invalid --json: %w", err)
			}
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}
			resp, err := sdk.Client.Messages.Retrieve(cmd.Context(), req, nil)
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
