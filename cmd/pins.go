package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/api-client-go/models/operations"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
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
  glean pins create --json '{"queries":["onboarding","new hire"],"documentId":"https://wiki.example.com/onboarding"}'`,
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
	return cmdutil.Build(cmdutil.Spec[operations.ListpinsRequestBody]{
		Use:   "list",
		Short: "List pins",
		Run: func(ctx context.Context, sdk *glean.Glean, req operations.ListpinsRequestBody) (any, error) {
			resp, err := sdk.Client.Pins.List(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.ListPinsResponse, nil
		},
	})
}

func newPinsGetCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.GetPinRequest]{
		Use:          "get",
		Short:        "Get a pin",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.GetPinRequest) (any, error) {
			resp, err := sdk.Client.Pins.Retrieve(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.GetPinResponse, nil
		},
	})
}

func newPinsCreateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.PinRequest]{
		Use:          "create",
		Short:        "Create a pin",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.PinRequest) (any, error) {
			resp, err := sdk.Client.Pins.Create(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.PinDocument, nil
		},
	})
}

func newPinsUpdateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.EditPinRequest]{
		Use:          "update",
		Short:        "Update a pin",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.EditPinRequest) (any, error) {
			resp, err := sdk.Client.Pins.Update(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.PinDocument, nil
		},
	})
}

func newPinsRemoveCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.Unpin]{
		Use:          "remove",
		Short:        "Remove a pin",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.Unpin) (any, error) {
			_, err := sdk.Client.Pins.Remove(ctx, req, nil)
			return nil, err
		},
	})
}
