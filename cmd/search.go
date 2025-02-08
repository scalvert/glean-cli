package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mattn/go-tty"
	"github.com/scalvert/glean-cli/pkg/api"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/scalvert/glean-cli/pkg/output"
	"github.com/scalvert/glean-cli/pkg/theme"
	"github.com/spf13/cobra"
)

// Test mode configuration for automated testing
var (
	testMode  bool   // Whether we're running in test mode
	testInput string // Simulated user input for testing
)

var defaultTemplate = `{{- range $i, $result := .Results -}}
{{if $i}}

{{end}}{{gleanBlue (add $i 1)}} {{gleanBlue (formatDatasource $result.Document.Datasource)}} | {{bold $result.Document.Title}}
{{gleanYellow $result.Document.URL}}
{{- range $result.Snippets}}
{{.Text}}{{end}}
{{- end}}{{if .SuggestedSpellCorrectedQuery}}

Did you mean: {{.SuggestedSpellCorrectedQuery}}?{{end}}{{if .RewrittenQuery}}

Showing results for: {{.RewrittenQuery}}{{end}}`

type Document = api.Document
type DocumentMetadata = api.DocumentMetadata
type Person = api.Person
type PersonMetadata = api.PersonMetadata
type RelatedDocument = api.RelatedDocument
type Shortcut = api.Shortcut
type StructuredResult = api.StructuredResult

type SearchOptions struct {
	InputDetails      *SearchInputDetails `json:"inputDetails,omitempty"`
	SessionInfo       *SessionInfo        `json:"sessionInfo,omitempty"`
	SourceDocument    *Document           `json:"sourceDocument,omitempty"`
	RequestOptions    *RequestOptions     `json:"requestOptions,omitempty"`
	Template          string
	OutputFormat      string
	Query             string   `json:"query"`
	Cursor            string   `json:"cursor,omitempty"`
	Timestamp         string   `json:"timestamp,omitempty"`
	TrackingToken     string   `json:"trackingToken,omitempty"`
	People            []Person `json:"people,omitempty"`
	ResultTabIds      []string `json:"resultTabIds,omitempty"`
	PageSize          int      `json:"pageSize,omitempty"`
	MaxSnippetSize    int      `json:"maxSnippetSize,omitempty"`
	TimeoutMillis     int      `json:"timeoutMillis,omitempty"`
	DisableSpellcheck bool     `json:"disableSpellcheck,omitempty"`
	NoColor           bool
}

type SearchInputDetails struct {
	HasCopyPaste bool `json:"hasCopyPaste,omitempty"`
}

type RequestOptions struct {
	Exclusions                   *RestrictionFilters `json:"exclusions,omitempty"`
	FacetBucketFilter            *FacetBucketFilter  `json:"facetBucketFilter,omitempty"`
	Inclusions                   *RestrictionFilters `json:"inclusions,omitempty"`
	DatasourceFilter             string              `json:"datasourceFilter,omitempty"`
	FacetFilters                 []FacetFilter       `json:"facetFilters,omitempty"`
	FacetFilterSets              []FacetFilterSet    `json:"facetFilterSets,omitempty"`
	AuthTokens                   []AuthToken         `json:"authTokens,omitempty"`
	ResponseHints                []string            `json:"responseHints,omitempty"`
	DefaultFacets                []string            `json:"defaultFacets,omitempty"`
	DatasourcesFilter            []string            `json:"datasourcesFilter,omitempty"`
	FacetBucketSize              int                 `json:"facetBucketSize"`
	TimezoneOffset               int                 `json:"timezoneOffset,omitempty"`
	DisableQueryAutocorrect      bool                `json:"disableQueryAutocorrect,omitempty"`
	DisableSpellcheck            bool                `json:"disableSpellcheck,omitempty"`
	FetchAllDatasourceCounts     bool                `json:"fetchAllDatasourceCounts,omitempty"`
	QueryOverridesFacetFilters   bool                `json:"queryOverridesFacetFilters,omitempty"`
	ReturnLlmContentOverSnippets bool                `json:"returnLlmContentOverSnippets,omitempty"`
}

type AuthToken struct {
	// Add fields as needed
}

type RestrictionFilters struct {
	// Add fields as needed
}

type FacetBucketFilter struct {
	// Add fields as needed
}

type FacetFilter struct {
	FieldName string        `json:"fieldName"`
	Values    []FilterValue `json:"values"`
}

type FacetFilterSet struct {
	Filters []FacetFilter `json:"filters"`
}

type FilterValue struct {
	Value        string `json:"value"`
	RelationType string `json:"relationType"`
}

type SessionInfo struct {
	LastQuery            string `json:"lastQuery,omitempty"`
	LastSeen             string `json:"lastSeen,omitempty"`
	SessionTrackingToken string `json:"sessionTrackingToken,omitempty"`
	TabId                string `json:"tabId,omitempty"`
}

// Additional response types
type ErrorInfo struct {
	ErrorMessages []ErrorMessage `json:"errorMessages,omitempty"`
}

type ErrorMessage struct {
	Source       string `json:"source"`
	ErrorMessage string `json:"errorMessage"`
}

type FacetResult struct {
	FieldName string        `json:"fieldName"`
	Buckets   []FacetBucket `json:"buckets"`
}

type FacetBucket struct {
	Value       interface{} `json:"value"`
	DisplayName string      `json:"displayName,omitempty"`
	Count       int         `json:"count"`
}

type GeneratedQna struct {
	Answer     string  `json:"answer"`
	Confidence float64 `json:"confidence"`
}

type SearchResponseMetadata struct {
	TotalResults int `json:"totalResults"`
}

type ResultsDescription struct {
	Description string `json:"description"`
}

type ResultTab struct {
	TabId       string `json:"tabId"`
	TabName     string `json:"tabName"`
	ResultCount int    `json:"resultCount"`
}

type SearchResponse struct {
	ErrorInfo                    *ErrorInfo              `json:"errorInfo,omitempty"`
	SessionInfo                  *SessionInfo            `json:"sessionInfo"`
	ResultsDescription           *ResultsDescription     `json:"resultsDescription,omitempty"`
	Metadata                     *SearchResponseMetadata `json:"metadata,omitempty"`
	GeneratedQna                 *GeneratedQna           `json:"generatedQnaResult,omitempty"`
	TrackingToken                string                  `json:"trackingToken"`
	ResponseTrackingToken        string                  `json:"responseTrackingToken"`
	SuggestedSpellCorrectedQuery string                  `json:"suggestedSpellCorrectedQuery,omitempty"`
	RewrittenQuery               string                  `json:"rewrittenQuery,omitempty"`
	Cursor                       string                  `json:"cursor"`
	RequestID                    string                  `json:"requestID,omitempty"`
	FacetResults                 []FacetResult           `json:"facetResults,omitempty"`
	Results                      []SearchResult          `json:"results"`
	ExperimentIds                []int64                 `json:"experimentIds,omitempty"`
	StructuredResults            []StructuredResult      `json:"structuredResults,omitempty"`
	ResultTabIds                 []string                `json:"resultTabIds,omitempty"`
	ResultTabs                   []ResultTab             `json:"resultTabs,omitempty"`
	BackendTimeMillis            int                     `json:"backendTimeMillis,omitempty"`
	HasMoreResults               bool                    `json:"hasMoreResults"`
}

type RewrittenFacetFilter struct {
	FieldName string   `json:"fieldName"`
	Values    []string `json:"values"`
}

type SearchResult struct {
	Document               *Document              `json:"document"`
	MustIncludeSuggestions map[string]interface{} `json:"mustIncludeSuggestions"`
	DebugInfo              map[string]interface{} `json:"debugInfo"`
	TrackingToken          string                 `json:"trackingToken"`
	Title                  string                 `json:"title"`
	URL                    string                 `json:"url"`
	Snippets               []SearchSnippet        `json:"snippets"`
}

type SearchSnippet struct {
	Snippet             string         `json:"snippet"`
	MimeType            string         `json:"mimeType"`
	Text                string         `json:"text"`
	URL                 string         `json:"url,omitempty"`
	Ranges              []SnippetRange `json:"ranges,omitempty"`
	SnippetTextOrdering int            `json:"snippetTextOrdering,omitempty"`
}

type SnippetRange struct {
	Type       string `json:"type"`
	StartIndex int    `json:"startIndex"`
	EndIndex   int    `json:"endIndex"`
}

type SearchMetadata struct {
	Container  string  `json:"container,omitempty"`
	CreateTime string  `json:"createTime,omitempty"`
	Datasource string  `json:"datasource,omitempty"`
	Author     *Person `json:"author,omitempty"`
	DocumentId string  `json:"documentId,omitempty"`
	UpdateTime string  `json:"updateTime,omitempty"`
	MimeType   string  `json:"mimeType,omitempty"`
	ObjectType string  `json:"objectType,omitempty"`
}

// NewCmdSearch creates and returns the search command.
// The search command allows users to search across their Glean instance,
// with support for pagination, custom output formats, and filtering options.
func NewCmdSearch() *cobra.Command {
	opts := &SearchOptions{
		RequestOptions: &RequestOptions{
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
  glean search --page-size 20 "engineering docs"
  glean search --template "{{range .Results}}{{.Title}}\n{{end}}" "meeting notes"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Query = args[0]
			return runSearch(cmd, opts)
		},
	}

	cmd.Flags().IntVarP(&opts.PageSize, "page-size", "n", 10, "Number of results per page")
	cmd.Flags().BoolVar(&opts.DisableSpellcheck, "disable-spellcheck", false, "Disable spellcheck suggestions")
	cmd.Flags().IntVar(&opts.MaxSnippetSize, "max-snippet-size", 200, "Maximum length of result snippets")
	cmd.Flags().IntVar(&opts.TimeoutMillis, "timeout", 30000, "Request timeout in milliseconds")

	cmd.Flags().StringVarP(&opts.Template, "template", "t", "", "Go template for formatting results")
	cmd.Flags().StringVarP(&opts.OutputFormat, "output", "o", "text", "Output format: text, json")
	cmd.Flags().BoolVar(&opts.NoColor, "no-color", false, "Disable colorized output")

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

func runSearch(cmd *cobra.Command, opts *SearchOptions) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client, err := http.NewClient(cfg)
	if err != nil {
		return err
	}

	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spin.Prefix = "Searching Company Knowledge "
	spin.Start()
	defer spin.Stop()

	if datasources, flagErr := cmd.Flags().GetStringSlice("datasource"); flagErr == nil && len(datasources) > 0 {
		addFacetFilter(opts, "datasource", datasources)
	}
	if types, flagErr := cmd.Flags().GetStringSlice("type"); flagErr == nil && len(types) > 0 {
		addFacetFilter(opts, "type", types)
	}

	if people, flagErr := cmd.Flags().GetStringSlice("person"); flagErr == nil && len(people) > 0 {
		opts.People = make([]Person, len(people))
		for i, email := range people {
			opts.People[i] = Person{
				ObfuscatedId: email,
			}
		}
	}

	if tabs, flagErr := cmd.Flags().GetStringSlice("tab"); flagErr == nil && len(tabs) > 0 {
		opts.ResultTabIds = tabs
	}

	if opts.RequestOptions == nil {
		opts.RequestOptions = &RequestOptions{}
	}

	opts.RequestOptions.TimezoneOffset = getTimezoneOffset()

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

	resp, err := performSearch(client, opts, "", "")
	if err != nil {
		return err
	}

	spin.Stop()

	if opts.OutputFormat == "json" {
		jsonBytes, marshalErr := json.Marshal(resp)
		if marshalErr != nil {
			return fmt.Errorf("error marshaling JSON: %w", marshalErr)
		}
		return output.Write(cmd.OutOrStdout(), jsonBytes, output.Options{
			NoColor: opts.NoColor,
			Format:  "json",
		})
	}

	tmpl := defaultTemplate
	if opts.Template != "" {
		tmpl = opts.Template
	}

	funcs := template.FuncMap{
		"add":              func(a, b int) int { return a + b },
		"formatDatasource": formatDatasource,
	}

	for k, v := range theme.TemplateFuncs(opts.NoColor) {
		funcs[k] = v
	}

	t, err := template.New("search").Funcs(funcs).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	err = t.Execute(cmd.OutOrStdout(), resp)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	for resp.HasMoreResults {
		fmt.Fprint(cmd.OutOrStdout(), "\n\nPress 'q' to quit, any other key to load more results...")

		if testMode {
			if testInput == "q" || testInput == "Q" {
				return nil
			}
			break
		}

		var readErr error
		ttyInput, readErr := tty.Open()
		if readErr != nil {
			// Fall back to standard input if TTY is not available
			var input string
			if _, readErr := fmt.Scanln(&input); readErr != nil {
				if readErr == io.EOF {
					return nil // User pressed Ctrl+D
				}
				return fmt.Errorf("error reading input: %w", readErr)
			}
			if input == "q" || input == "Q" {
				return nil
			}
		} else {
			defer ttyInput.Close()
			r, readErr := ttyInput.ReadRune()
			if readErr != nil {
				return fmt.Errorf("error reading input: %w", readErr)
			}
			if r == 'q' || r == 'Q' {
				return nil
			}
		}

		resp, err = performSearch(client, opts, resp.Cursor, resp.TrackingToken)
		if err != nil {
			return err
		}

		err = t.Execute(cmd.OutOrStdout(), resp)
		if err != nil {
			return fmt.Errorf("error executing template: %w", err)
		}
	}

	return nil
}

func getTimezoneOffset() int {
	_, offset := time.Now().Zone()
	return offset / 60
}

func addFacetFilter(opts *SearchOptions, fieldName string, values []string) {
	filter := FacetFilter{
		FieldName: fieldName,
		Values:    make([]FilterValue, len(values)),
	}
	for i, value := range values {
		filter.Values[i] = FilterValue{
			Value:        value,
			RelationType: "EQUALS",
		}
	}
	opts.RequestOptions.FacetFilters = append(opts.RequestOptions.FacetFilters, filter)
}

func formatDatasource(s string) string {
	if s == "nonindexedshortcut" {
		return "GoLink"
	}

	words := []rune(s)
	if len(words) > 0 {
		words[0] = []rune(strings.ToUpper(string(words[0])))[0]
	}
	for i := 1; i < len(words); i++ {
		if words[i-1] == ' ' {
			words[i] = []rune(strings.ToUpper(string(words[i])))[0]
		}
	}
	return string(words)
}

func performSearch(client http.Client, opts *SearchOptions, cursor, trackingToken string) (*SearchResponse, error) {
	requestBody := map[string]interface{}{
		"query":             opts.Query,
		"pageSize":          opts.PageSize,
		"disableSpellcheck": opts.DisableSpellcheck,
		"maxSnippetSize":    opts.MaxSnippetSize,
		"timeoutMillis":     opts.TimeoutMillis,
	}

	// Add optional parameters if they're set
	if opts.InputDetails != nil {
		requestBody["inputDetails"] = opts.InputDetails
	}
	if len(opts.People) > 0 {
		requestBody["people"] = opts.People
	}
	if opts.RequestOptions != nil && len(opts.RequestOptions.FacetFilters) > 0 {
		requestBody["requestOptions"] = opts.RequestOptions
	}
	if len(opts.ResultTabIds) > 0 {
		requestBody["resultTabIds"] = opts.ResultTabIds
	}
	if opts.SessionInfo != nil {
		requestBody["sessionInfo"] = opts.SessionInfo
	}
	if opts.SourceDocument != nil {
		requestBody["sourceDocument"] = opts.SourceDocument
	}
	if opts.Timestamp != "" {
		requestBody["timestamp"] = opts.Timestamp
	}
	if cursor != "" {
		requestBody["cursor"] = cursor
	}
	if trackingToken != "" {
		requestBody["trackingToken"] = trackingToken
	}

	req := &http.Request{
		Method: "POST",
		Path:   "search",
		Body:   requestBody,
	}

	resp, err := client.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making search request: %w", err)
	}

	var searchResp SearchResponse
	if err := json.Unmarshal(resp, &searchResp); err != nil {
		return nil, fmt.Errorf("error parsing search response: %w", err)
	}

	return &searchResp, nil
}
