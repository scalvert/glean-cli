package cmd

import (
	"context"
	"fmt"
	"io"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/gleanwork/glean-cli/internal/output"
	"github.com/spf13/cobra"
)

func NewCmdCollections() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collections",
		Short: "Manage Glean collections",
		Long: `Manage Glean collections.

Collections are curated groups of documents and links organized around a topic or project.

Example:
  glean collections list
  glean collections create --json '{"name":"Onboarding","description":"New hire resources"}'`,
	}
	cmd.AddCommand(
		newCollectionsListCmd(),
		newCollectionsCreateCmd(),
		newCollectionsDeleteCmd(),
		newCollectionsUpdateCmd(),
		newCollectionsAddItemsCmd(),
		newCollectionsDeleteItemCmd(),
	)
	return cmd
}

func newCollectionsListCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.ListCollectionsRequest]{
		Use:   "list",
		Short: "List collections",
		TextFn: func(w io.Writer, v any) error {
			resp, ok := v.(*components.ListCollectionsResponse)
			if !ok {
				return output.WriteJSON(w, v)
			}
			rows := make([][]string, len(resp.Collections))
			for i, c := range resp.Collections {
				rows[i] = []string{
					fmt.Sprintf("%d", c.ID),
					c.Name,
					output.Truncate(c.Description, 60),
				}
			}
			return output.WriteTable(w, []string{"ID", "NAME", "DESCRIPTION"}, rows)
		},
		Run: func(ctx context.Context, sdk *glean.Glean, req components.ListCollectionsRequest) (any, error) {
			resp, err := sdk.Client.Collections.List(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.ListCollectionsResponse, nil
		},
	})
}

func newCollectionsCreateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.CreateCollectionRequest]{
		Use:          "create",
		Short:        "Create a collection",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.CreateCollectionRequest) (any, error) {
			resp, err := sdk.Client.Collections.Create(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.OneOf, nil
		},
	})
}

func newCollectionsDeleteCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.DeleteCollectionRequest]{
		Use:          "delete",
		Short:        "Delete a collection",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.DeleteCollectionRequest) (any, error) {
			_, err := sdk.Client.Collections.Delete(ctx, req, nil)
			return nil, err
		},
	})
}

func newCollectionsUpdateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.EditCollectionRequest]{
		Use:          "update",
		Short:        "Update a collection",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.EditCollectionRequest) (any, error) {
			resp, err := sdk.Client.Collections.Update(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.EditCollectionResponse, nil
		},
	})
}

func newCollectionsAddItemsCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.AddCollectionItemsRequest]{
		Use:          "add-items",
		Short:        "Add items to a collection",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.AddCollectionItemsRequest) (any, error) {
			resp, err := sdk.Client.Collections.AddItems(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.AddCollectionItemsResponse, nil
		},
	})
}

func newCollectionsDeleteItemCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.DeleteCollectionItemRequest]{
		Use:          "delete-item",
		Short:        "Delete an item from a collection",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.DeleteCollectionItemRequest) (any, error) {
			_, err := sdk.Client.Collections.DeleteItem(ctx, req, nil)
			return nil, err
		},
	})
}
