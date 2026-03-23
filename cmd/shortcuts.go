package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdShortcuts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shortcuts",
		Short: "Manage Glean shortcuts (go-links)",
		Long: `Manage Glean shortcuts (go-links).

Shortcuts are memorable short URLs that redirect to longer internal resources.

Example:
  glean shortcuts list
  glean shortcuts create --json '{"data":{"inputAlias":"onboarding","destinationUrl":"https://wiki.example.com/onboarding"}}'
  glean shortcuts get --json '{"alias":"onboarding"}'`,
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
	return cmdutil.Build(cmdutil.Spec[components.ListShortcutsPaginatedRequest]{
		Use:   "list",
		Short: "List shortcuts",
		Run: func(ctx context.Context, sdk *glean.Glean, req components.ListShortcutsPaginatedRequest) (any, error) {
			resp, err := sdk.Client.Shortcuts.List(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.ListShortcutsPaginatedResponse, nil
		},
	})
}

func newShortcutsGetCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.GetShortcutRequestUnion]{
		Use:          "get",
		Short:        "Get a shortcut by alias or ID",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.GetShortcutRequestUnion) (any, error) {
			resp, err := sdk.Client.Shortcuts.Retrieve(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.GetShortcutResponse, nil
		},
	})
}

func newShortcutsCreateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.CreateShortcutRequest]{
		Use:          "create",
		Short:        "Create a shortcut",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.CreateShortcutRequest) (any, error) {
			resp, err := sdk.Client.Shortcuts.Create(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.CreateShortcutResponse, nil
		},
	})
}

func newShortcutsUpdateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.UpdateShortcutRequest]{
		Use:          "update",
		Short:        "Update a shortcut",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.UpdateShortcutRequest) (any, error) {
			resp, err := sdk.Client.Shortcuts.Update(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.UpdateShortcutResponse, nil
		},
	})
}

func newShortcutsDeleteCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.DeleteShortcutRequest]{
		Use:          "delete",
		Short:        "Delete a shortcut",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.DeleteShortcutRequest) (any, error) {
			_, err := sdk.Client.Shortcuts.Delete(ctx, req, nil)
			return nil, err
		},
	})
}
