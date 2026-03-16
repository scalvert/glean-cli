package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gleanwork/api-client-go/models/components"
	gleanClient "github.com/gleanwork/glean-cli/internal/client"
	"github.com/gleanwork/glean-cli/internal/output"
	"github.com/gleanwork/glean-cli/internal/search"
	"github.com/spf13/cobra"
)

// NewCmdSearch creates and returns the search command.
func NewCmdSearch() *cobra.Command {
	opts := &search.Options{
		RequestOptions: &search.RequestOptions{
			FacetBucketSize: 10,
			ResponseHints:   []string{"RESULTS", "QUERY_METADATA"},
		},
	}
	var jsonPayload string
	var outputFormat string
	var dryRun bool
	var fields string

	cmd := &cobra.Command{
		Use:   "search [query]",
		Short: "Search for content in your Glean instance",
		Long: `Search for content in your Glean instance.

Results are written as JSON to stdout by default, making the output easy to
pipe to jq or other tools.

Example:
  glean search "vacation policy"
  glean search "vacation policy" | jq '.results[].document.title'
  glean search --json '{"query":"Q1 reports","pageSize":5}' | jq .
  glean search --output ndjson "engineering docs" | head -3 | jq .document.title
  glean search --dry-run "test"`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if jsonPayload == "" && len(args) == 0 {
				return fmt.Errorf("requires a query argument or --json payload")
			}

			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}

			// --json path: parse directly into SearchRequest
			if jsonPayload != "" {
				var req components.SearchRequest
				if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
					return fmt.Errorf("invalid --json payload: %w", err)
				}
				if dryRun {
					return output.WriteJSON(cmd.OutOrStdout(), req)
				}
				resp, err := sdk.Client.Search.Query(cmd.Context(), req, nil)
				if err != nil {
					return fmt.Errorf("search request failed: %w", err)
				}
				return output.WriteFormatted(cmd.OutOrStdout(), resp.SearchResponse, outputFormat, nil)
			}

			// flag-based path
			if datasources, flagErr := cmd.Flags().GetStringSlice("datasource"); flagErr == nil && len(datasources) > 0 {
				search.AddFacetFilter(opts, "datasource", datasources)
			}
			if types, flagErr := cmd.Flags().GetStringSlice("type"); flagErr == nil && len(types) > 0 {
				search.AddFacetFilter(opts, "type", types)
			}
			if tabs, flagErr := cmd.Flags().GetStringSlice("tab"); flagErr == nil && len(tabs) > 0 {
				opts.ResultTabIds = tabs
			}
			if opts.RequestOptions == nil {
				opts.RequestOptions = &search.RequestOptions{}
			}
			opts.RequestOptions.TimezoneOffset = search.GetTimezoneOffset()
			if size, flagErr := cmd.Flags().GetInt("facet-bucket-size"); flagErr == nil {
				opts.RequestOptions.FacetBucketSize = size
			}
			if hints, flagErr := cmd.Flags().GetStringSlice("response-hints"); flagErr == nil {
				opts.RequestOptions.ResponseHints = hints
			}
			if disable, flagErr := cmd.Flags().GetBool("disable-query-autocorrect"); flagErr == nil {
				opts.RequestOptions.DisableQueryAutocorrect = disable
			}
			if fetch, flagErr := cmd.Flags().GetBool("fetch-all-datasource-counts"); flagErr == nil {
				opts.RequestOptions.FetchAllDatasourceCounts = fetch
			}
			if override, flagErr := cmd.Flags().GetBool("query-overrides-facet-filters"); flagErr == nil {
				opts.RequestOptions.QueryOverridesFacetFilters = override
			}
			if llm, flagErr := cmd.Flags().GetBool("return-llm-content"); flagErr == nil {
				opts.RequestOptions.ReturnLlmContentOverSnippets = llm
			}

			opts.Query = args[0]

			if dryRun {
				return output.WriteJSON(cmd.OutOrStdout(), search.BuildSearchRequest(opts))
			}

			resp, err := search.RunSearchSDK(cmd.Context(), opts, sdk)
			if err != nil {
				return err
			}
			if fields != "" {
				return output.ProjectFields(cmd.OutOrStdout(), resp, fields)
			}
			return output.WriteFormatted(cmd.OutOrStdout(), resp, outputFormat, nil)
		},
	}

	cmd.Flags().StringVar(&jsonPayload, "json", "", "Complete JSON request body (overrides all other flags)")
	cmd.Flags().StringVar(&outputFormat, "output", "json", "Output format: json, ndjson, or text")
	cmd.Flags().StringVar(&fields, "fields", "", "Comma-separated dot-path fields to include (e.g. results.document.title,results.document.url). Results where all projected fields are missing appear as {}")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print the request body without sending it")
	cmd.Flags().IntVar(&opts.PageSize, "page-size", 10, "Number of results per page")
	cmd.Flags().IntVar(&opts.MaxSnippetSize, "max-snippet-size", 0, "Maximum size of snippets")
	cmd.Flags().IntVar(&opts.TimeoutMillis, "timeout", 30000, "Request timeout in milliseconds")
	cmd.Flags().BoolVar(&opts.DisableSpellcheck, "disable-spellcheck", false, "Disable spellcheck")

	cmd.Flags().StringSliceP("datasource", "d", nil, "Filter by datasource (can be specified multiple times)")
	cmd.Flags().StringSliceP("type", "t", nil, "Filter by document type (can be specified multiple times)")
	cmd.Flags().StringSlice("tab", nil, "Filter by result tab IDs (can be specified multiple times)")

	cmd.Flags().Bool("disable-query-autocorrect", false, "Disable automatic query corrections")
	cmd.Flags().Bool("fetch-all-datasource-counts", false, "Return result counts for all supported datasources")
	cmd.Flags().Bool("query-overrides-facet-filters", false, "Let query operators override facet filters")
	cmd.Flags().Bool("return-llm-content", false, "Return expanded content for LLM usage")
	cmd.Flags().StringSlice("response-hints", []string{"RESULTS", "QUERY_METADATA"}, "Response hints (RESULTS, QUERY_METADATA, etc)")
	cmd.Flags().Int("facet-bucket-size", 10, "Maximum number of facet buckets to return")

	return cmd
}
