package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdVerification() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verification",
		Short: "Manage document verification",
		Long: `Manage document verification.

Document verification lets teams mark documents as reviewed and accurate for a given time window.

Example:
  glean verification list
  glean verification verify --json '{"documentId":"doc123","action":"VERIFY"}'
  glean verification remind --json '{"documentId":"doc123","remindInDays":30}'`,
	}
	cmd.AddCommand(
		newVerificationListCmd(),
		newVerificationVerifyCmd(),
		newVerificationRemindCmd(),
	)
	return cmd
}

func newVerificationListCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[struct{}]{
		Use:   "list",
		Short: "List verifications pending review",
		Run: func(ctx context.Context, sdk *glean.Glean, req struct{}) (any, error) {
			resp, err := sdk.Client.Verification.List(ctx, nil, nil)
			if err != nil {
				return nil, err
			}
			return resp.VerificationFeed, nil
		},
	})
}

func newVerificationVerifyCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.VerifyRequest]{
		Use:          "verify",
		Short:        "Mark a document as verified",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.VerifyRequest) (any, error) {
			resp, err := sdk.Client.Verification.Verify(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.Verification, nil
		},
	})
}

func newVerificationRemindCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.ReminderRequest]{
		Use:          "remind",
		Short:        "Send a verification reminder",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.ReminderRequest) (any, error) {
			resp, err := sdk.Client.Verification.AddReminder(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.Verification, nil
		},
	})
}
