package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
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
	return cmdutil.Build(cmdutil.Spec[components.Activity]{
		Use:          "report",
		Short:        "Report user activity events",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.Activity) (any, error) {
			_, err := sdk.Client.Activity.Report(ctx, req)
			return nil, err
		},
	})
}

func newActivityFeedbackCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.Feedback]{
		Use:          "feedback",
		Short:        "Submit feedback",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.Feedback) (any, error) {
			_, err := sdk.Client.Activity.Feedback(ctx, nil, &req)
			return nil, err
		},
	})
}
