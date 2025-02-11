package cmd

import (
	"fmt"

	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/search"
	"github.com/spf13/cobra"
)

// NewCmdSearch creates and returns the search command.
// The search command allows users to search across their Glean instance,
// with support for pagination, custom output formats, and filtering options.
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

You can search for any content that is indexed in your Glean instance.
The results will be displayed in a formatted list by default.

Example:
  glean search "vacation policy"
  glean search -i                    # Start interactive mode
  glean search -i "vacation policy"  # Start interactive mode with initial query
  glean search --page-size 20 "engineering docs"`,
		Args: func(cmd *cobra.Command, args []string) error {
			interactive, _ := cmd.Flags().GetBool("interactive")
			if interactive && len(args) == 0 {
				return nil // Allow no args in interactive mode
			}
			if len(args) != 1 {
				return fmt.Errorf("requires exactly one query argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadConfig()
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			client, err := http.NewClient(cfg)
			if err != nil {
				return err
			}

			// Handle all the flag options first
			if datasources, flagErr := cmd.Flags().GetStringSlice("datasource"); flagErr == nil && len(datasources) > 0 {
				search.AddFacetFilter(opts, "datasource", datasources)
			}
			if types, flagErr := cmd.Flags().GetStringSlice("type"); flagErr == nil && len(types) > 0 {
				search.AddFacetFilter(opts, "type", types)
			}

			if people, flagErr := cmd.Flags().GetStringSlice("person"); flagErr == nil && len(people) > 0 {
				opts.People = make([]search.Person, len(people))
				for i, email := range people {
					opts.People[i] = search.Person{
						ObfuscatedId: email,
					}
				}
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

			if opts.Interactive {
				query := ""
				if len(args) > 0 {
					query = args[0]
				}
				opts.Query = query
				return search.RunInteractiveSearch(opts, client)
			}

			opts.Query = args[0]
			return search.RunSearch(opts, client)
		},
	}

	cmd.Flags().StringVar(&opts.OutputFormat, "output", "", "Output format (json)")
	cmd.Flags().IntVar(&opts.PageSize, "page-size", 10, "Number of results per page")
	cmd.Flags().IntVar(&opts.MaxSnippetSize, "max-snippet-size", 0, "Maximum size of snippets")
	cmd.Flags().IntVar(&opts.TimeoutMillis, "timeout", 30000, "Request timeout in milliseconds")
	cmd.Flags().BoolVar(&opts.DisableSpellcheck, "disable-spellcheck", false, "Disable spellcheck")
	cmd.Flags().BoolVar(&opts.NoColor, "no-color", false, "Disable color output")
	cmd.Flags().BoolVarP(&opts.Interactive, "interactive", "i", false, "Run in interactive mode")

	cmd.Flags().StringSliceP("datasource", "d", nil, "Filter by datasource (can be specified multiple times)")
	cmd.Flags().StringSliceP("type", "y", nil, "Filter by document type (can be specified multiple times)")
	cmd.Flags().StringSliceP("person", "p", nil, "Filter by person email (can be specified multiple times)")
	cmd.Flags().StringSlice("tab", nil, "Filter by result tab IDs (can be specified multiple times)")

	cmd.Flags().Bool("disable-query-autocorrect", false, "Disable automatic query corrections")
	cmd.Flags().Bool("fetch-all-datasource-counts", false, "Return result counts for all supported datasources")
	cmd.Flags().Bool("query-overrides-facet-filters", false, "Let query operators override facet filters")
	cmd.Flags().Bool("return-llm-content", false, "Return expanded content for LLM usage")
	cmd.Flags().StringSlice("response-hints", []string{"RESULTS", "QUERY_METADATA"}, "Response hints (RESULTS, QUERY_METADATA, etc)")
	cmd.Flags().Int("facet-bucket-size", 10, "Maximum number of facet buckets to return")

	return cmd
}
