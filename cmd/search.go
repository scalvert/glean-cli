package cmd

import (
	gleanClient "github.com/scalvert/glean-cli/internal/client"
	"github.com/scalvert/glean-cli/internal/search"
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

	cmd := &cobra.Command{
		Use:   "search [query]",
		Short: "Search for content in your Glean instance",
		Long: `Search for content in your Glean instance.

Results are written as JSON to stdout, making the output easy to pipe
to jq or other tools.

Example:
  glean search "vacation policy"
  glean search "vacation policy" | jq '.results[].document.title'
  glean search --page-size 20 "engineering docs"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			sdk, err := gleanClient.NewFromConfig()
			if err != nil {
				return err
			}

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
			return search.RunSearch(cmd.Context(), opts, sdk, cmd.OutOrStdout())
		},
	}

	cmd.Flags().IntVar(&opts.PageSize, "page-size", 10, "Number of results per page")
	cmd.Flags().IntVar(&opts.MaxSnippetSize, "max-snippet-size", 0, "Maximum size of snippets")
	cmd.Flags().IntVar(&opts.TimeoutMillis, "timeout", 30000, "Request timeout in milliseconds")
	cmd.Flags().BoolVar(&opts.DisableSpellcheck, "disable-spellcheck", false, "Disable spellcheck")

	cmd.Flags().StringSliceP("datasource", "d", nil, "Filter by datasource (can be specified multiple times)")
	cmd.Flags().StringSliceP("type", "y", nil, "Filter by document type (can be specified multiple times)")
	cmd.Flags().StringSlice("tab", nil, "Filter by result tab IDs (can be specified multiple times)")

	cmd.Flags().Bool("disable-query-autocorrect", false, "Disable automatic query corrections")
	cmd.Flags().Bool("fetch-all-datasource-counts", false, "Return result counts for all supported datasources")
	cmd.Flags().Bool("query-overrides-facet-filters", false, "Let query operators override facet filters")
	cmd.Flags().Bool("return-llm-content", false, "Return expanded content for LLM usage")
	cmd.Flags().StringSlice("response-hints", []string{"RESULTS", "QUERY_METADATA"}, "Response hints (RESULTS, QUERY_METADATA, etc)")
	cmd.Flags().Int("facet-bucket-size", 10, "Maximum number of facet buckets to return")

	return cmd
}
