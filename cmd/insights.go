package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdInsights() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insights",
		Short: "Retrieve Glean usage insights",
		Long: `Retrieve Glean usage insights.

Insights provide analytics on how your organization uses Glean — search trends, popular content, and more.

Example:
  glean insights get --json '{"overviewRequest":{"disablePerUserInsights":false}}'`,
	}
	cmd.AddCommand(newInsightsGetCmd())
	return cmd
}

func newInsightsGetCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.InsightsRequest]{
		Use:          "get",
		Short:        "Get usage insights",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.InsightsRequest) (any, error) {
			resp, err := sdk.Client.Insights.Retrieve(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.InsightsResponse, nil
		},
	})
}
