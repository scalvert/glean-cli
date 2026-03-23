package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
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
	return cmdutil.Build(cmdutil.Spec[components.CreateAnnouncementRequest]{
		Use:          "create",
		Short:        "Create an announcement",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.CreateAnnouncementRequest) (any, error) {
			resp, err := sdk.Client.Announcements.Create(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.Announcement, nil
		},
	})
}

func newAnnouncementsUpdateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.UpdateAnnouncementRequest]{
		Use:          "update",
		Short:        "Update an announcement",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.UpdateAnnouncementRequest) (any, error) {
			resp, err := sdk.Client.Announcements.Update(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.Announcement, nil
		},
	})
}

func newAnnouncementsDeleteCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.DeleteAnnouncementRequest]{
		Use:          "delete",
		Short:        "Delete an announcement",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.DeleteAnnouncementRequest) (any, error) {
			_, err := sdk.Client.Announcements.Delete(ctx, req, nil)
			return nil, err
		},
	})
}
