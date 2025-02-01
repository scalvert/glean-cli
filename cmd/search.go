package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"

	"github.com/mattn/go-tty"
	"github.com/scalvert/glean-cli/pkg/config"
	"github.com/scalvert/glean-cli/pkg/http"
	"github.com/spf13/cobra"
)

var testMode bool
var testInput string

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

type Person struct {
	Metadata     *PersonMetadata `json:"metadata,omitempty"`
	Name         string          `json:"name"`
	ObfuscatedId string          `json:"obfuscatedId"`
}

type PersonMetadata struct {
	RelatedDocuments []RelatedDocument `json:"relatedDocuments,omitempty"`
}

type RelatedDocument struct {
	// Add fields as needed
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

type Document struct {
	ParentDocument *Document         `json:"parentDocument,omitempty"`
	Metadata       *DocumentMetadata `json:"metadata"`
	ID             string            `json:"id"`
	Datasource     string            `json:"datasource"`
	DocType        string            `json:"docType"`
	Title          string            `json:"title"`
	URL            string            `json:"url"`
}

type DocumentMetadata struct {
	Datasource         string                 `json:"datasource"`
	DatasourceInstance string                 `json:"datasourceInstance"`
	ObjectType         string                 `json:"objectType"`
	Container          string                 `json:"container,omitempty"`
	ContainerId        string                 `json:"containerId,omitempty"`
	MimeType           string                 `json:"mimeType"`
	DocumentId         string                 `json:"documentId"`
	LoggingId          string                 `json:"loggingId"`
	CreateTime         string                 `json:"createTime"`
	UpdateTime         string                 `json:"updateTime"`
	Author             *Person                `json:"author,omitempty"`
	Owner              *Person                `json:"owner,omitempty"`
	Visibility         string                 `json:"visibility"`
	Status             string                 `json:"status,omitempty"`
	AssignedTo         *Person                `json:"assignedTo,omitempty"`
	DatasourceId       string                 `json:"datasourceId"`
	Interactions       map[string]interface{} `json:"interactions"`
	DocumentCategory   string                 `json:"documentCategory"`
	CustomData         map[string]interface{} `json:"customData,omitempty"`
	Shortcuts          []Shortcut             `json:"shortcuts,omitempty"`
}

type Shortcut struct {
	InputAlias     string `json:"inputAlias"`
	DestinationUrl string `json:"destinationUrl"`
	Description    string `json:"description"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
	ViewPrefix     string `json:"viewPrefix"`
	Alias          string `json:"alias"`
	Title          string `json:"title"`
}

var defaultTemplate = `{{- range $i, $result := .Results -}}
{{if $i}}

{{end}}{{add $i 1}} {{formatDatasource $result.Document.Datasource}} | {{gleanBlue $result.Document.Title}}
{{gleanYellow $result.Document.URL}}
{{- range $result.Snippets}}
{{.Text}}{{end}}
{{- end}}{{if .SuggestedSpellCorrectedQuery}}

Did you mean: {{.SuggestedSpellCorrectedQuery}}?{{end}}{{if .RewrittenQuery}}

Showing results for: {{.RewrittenQuery}}{{end}}`

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

	// Basic search options
	cmd.Flags().IntVarP(&opts.PageSize, "page-size", "n", 10, "Number of results per page")
	cmd.Flags().BoolVar(&opts.DisableSpellcheck, "disable-spellcheck", false, "Disable spellcheck suggestions")
	cmd.Flags().IntVar(&opts.MaxSnippetSize, "max-snippet-size", 200, "Maximum length of result snippets")
	cmd.Flags().IntVar(&opts.TimeoutMillis, "timeout", 30000, "Request timeout in milliseconds")

	// Output formatting
	cmd.Flags().StringVarP(&opts.Template, "template", "t", "", "Go template for formatting results")
	cmd.Flags().StringVarP(&opts.OutputFormat, "output", "o", "text", "Output format: text, json")
	cmd.Flags().BoolVar(&opts.NoColor, "no-color", false, "Disable colorized output")

	// Filtering options
	cmd.Flags().StringSliceP("datasource", "d", nil, "Filter by datasource (can be specified multiple times)")
	cmd.Flags().StringSliceP("type", "y", nil, "Filter by document type (can be specified multiple times)")
	cmd.Flags().StringSliceP("person", "p", nil, "Filter by person email (can be specified multiple times)")
	cmd.Flags().StringSlice("tab", nil, "Filter by result tab IDs (can be specified multiple times)")

	// Advanced options
	cmd.Flags().Bool("disable-query-autocorrect", false, "Disable automatic query corrections")
	cmd.Flags().Bool("fetch-all-datasource-counts", false, "Return result counts for all supported datasources")
	cmd.Flags().Bool("query-overrides-facet-filters", false, "Let query operators override facet filters")
	cmd.Flags().Bool("return-llm-content", false, "Return expanded content for LLM usage")
	cmd.Flags().StringSlice("response-hints", []string{"RESULTS", "QUERY_METADATA"}, "Response hints (RESULTS, QUERY_METADATA, etc)")
	cmd.Flags().Int("facet-bucket-size", 10, "Maximum number of facet buckets to return")

	return cmd
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
	Value string `json:"value"`
	Count int    `json:"count"`
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

type StructuredResult struct {
	Type     string          `json:"type"`
	Metadata json.RawMessage `json:"metadata,omitempty"`
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

func runSearch(cmd *cobra.Command, opts *SearchOptions) error {
	// Handle facet filters
	if datasources, err := cmd.Flags().GetStringSlice("datasource"); err == nil && len(datasources) > 0 {
		addFacetFilter(opts, "datasource", datasources)
	}
	if types, err := cmd.Flags().GetStringSlice("type"); err == nil && len(types) > 0 {
		addFacetFilter(opts, "type", types)
	}

	// Handle person filters
	if people, err := cmd.Flags().GetStringSlice("person"); err == nil && len(people) > 0 {
		opts.People = make([]Person, len(people))
		for i, email := range people {
			opts.People[i] = Person{
				ObfuscatedId: email,
			}
		}
	}

	// Handle result tab filters
	if tabs, err := cmd.Flags().GetStringSlice("tab"); err == nil && len(tabs) > 0 {
		opts.ResultTabIds = tabs
	}

	// Handle advanced options
	if opts.RequestOptions == nil {
		opts.RequestOptions = &RequestOptions{}
	}

	// Set timezone offset
	opts.RequestOptions.TimezoneOffset = getTimezoneOffset()

	// Set facet bucket size
	if size, err := cmd.Flags().GetInt("facet-bucket-size"); err == nil {
		opts.RequestOptions.FacetBucketSize = size
	}

	// Set response hints
	if hints, err := cmd.Flags().GetStringSlice("response-hints"); err == nil {
		opts.RequestOptions.ResponseHints = hints
	}

	// Set boolean flags
	if disable, err := cmd.Flags().GetBool("disable-query-autocorrect"); err == nil {
		opts.RequestOptions.DisableQueryAutocorrect = disable
	}
	if fetch, err := cmd.Flags().GetBool("fetch-all-datasource-counts"); err == nil {
		opts.RequestOptions.FetchAllDatasourceCounts = fetch
	}
	if override, err := cmd.Flags().GetBool("query-overrides-facet-filters"); err == nil {
		opts.RequestOptions.QueryOverridesFacetFilters = override
	}
	if llm, err := cmd.Flags().GetBool("return-llm-content"); err == nil {
		opts.RequestOptions.ReturnLlmContentOverSnippets = llm
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client, err := http.NewClient(cfg)
	if err != nil {
		return err
	}

	// Initial search request
	response, err := performSearch(client, opts, "", "")
	if err != nil {
		return err
	}

	if opts.OutputFormat == "json" {
		return json.NewEncoder(cmd.OutOrStdout()).Encode(response)
	}

	// Use template for output
	tmpl := defaultTemplate
	if opts.Template != "" {
		tmpl = opts.Template
	}

	t, err := template.New("search").Funcs(template.FuncMap{
		"gleanBlue": func(s string) string {
			if opts.NoColor {
				return s
			}
			return fmt.Sprintf("\033[38;2;82;105;255m%s\033[0m", s)
		},
		"gleanYellow": func(s string) string {
			if opts.NoColor {
				return s
			}
			return fmt.Sprintf("\033[38;2;236;240;115m%s\033[0m", s)
		},
		"add":              func(a, b int) int { return a + b },
		"formatDatasource": formatDatasource,
	}).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Print initial results
	err = t.Execute(cmd.OutOrStdout(), response)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	// Handle pagination if needed
	for response.HasMoreResults {
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

		response, err = performSearch(client, opts, response.Cursor, response.TrackingToken)
		if err != nil {
			return err
		}

		err = t.Execute(cmd.OutOrStdout(), response)
		if err != nil {
			return fmt.Errorf("error executing template: %w", err)
		}
	}

	return nil
}

// getTimezoneOffset returns the local timezone offset in minutes from UTC
func getTimezoneOffset() int {
	_, offset := time.Now().Zone()
	return offset / 60 // Convert seconds to minutes
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
