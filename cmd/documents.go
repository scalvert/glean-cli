package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdDocuments() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "documents",
		Short: "Retrieve and summarize Glean documents",
		Long: `Retrieve and summarize Glean documents.

Fetch full document content by URL or Glean document ID, or get an AI-generated summary.

Example:
  glean documents get --json '{"documentSpecs":[{"url":"https://app.glean.com/doc/abc"}]}'
  glean documents get --json '{"documentSpecs":[{"id":"<glean-doc-id>"}]}'
  glean documents summarize --json '{"documentSpecs":[{"url":"https://app.glean.com/doc/abc"}]}'`,
	}
	cmd.AddCommand(
		newDocumentsGetCmd(),
		newDocumentsGetByFacetsCmd(),
		newDocumentsGetPermissionsCmd(),
		newDocumentsSummarizeCmd(),
	)
	return cmd
}

func newDocumentsGetCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.GetDocumentsRequest]{
		Use:   "get",
		Short: "Get documents by ID",
		Run: func(ctx context.Context, sdk *glean.Glean, req components.GetDocumentsRequest) (any, error) {
			resp, err := sdk.Client.Documents.Retrieve(ctx, nil, &req)
			if err != nil {
				return nil, err
			}
			return resp.GetDocumentsResponse, nil
		},
	})
}

func newDocumentsGetByFacetsCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.GetDocumentsByFacetsRequest]{
		Use:   "get-by-facets",
		Short: "Get documents by facets",
		Run: func(ctx context.Context, sdk *glean.Glean, req components.GetDocumentsByFacetsRequest) (any, error) {
			resp, err := sdk.Client.Documents.RetrieveByFacets(ctx, nil, &req)
			if err != nil {
				return nil, err
			}
			return resp.GetDocumentsByFacetsResponse, nil
		},
	})
}

func newDocumentsGetPermissionsCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.GetDocPermissionsRequest]{
		Use:          "get-permissions",
		Short:        "Get document permissions",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.GetDocPermissionsRequest) (any, error) {
			resp, err := sdk.Client.Documents.RetrievePermissions(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.GetDocPermissionsResponse, nil
		},
	})
}

func newDocumentsSummarizeCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.SummarizeRequest]{
		Use:          "summarize",
		Short:        "Summarize a document",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.SummarizeRequest) (any, error) {
			resp, err := sdk.Client.Documents.Summarize(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.SummarizeResponse, nil
		},
	})
}
