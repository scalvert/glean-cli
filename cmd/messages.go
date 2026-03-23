package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "messages",
		Short: "Retrieve Glean messages",
		Long: `Retrieve Glean messages.

Messages are threaded communications (Slack, Teams, email, etc.) indexed in Glean.
Each request identifies a single channel, thread, or conversation by its type, ID, and datasource.

Example:
  glean messages get --json '{"idType":"THREAD_ID","id":"<thread-id>","datasource":"SLACK"}'
  glean messages get --json '{"idType":"CONVERSATION_ID","id":"<conv-id>","datasource":"GCHAT"}'`,
	}
	cmd.AddCommand(newMessagesGetCmd())
	return cmd
}

func newMessagesGetCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.MessagesRequest]{
		Use:          "get",
		Short:        "Get messages",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.MessagesRequest) (any, error) {
			resp, err := sdk.Client.Messages.Retrieve(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.MessagesResponse, nil
		},
	})
}
