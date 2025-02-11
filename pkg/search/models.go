package search

import (
	"fmt"

	"github.com/scalvert/glean-cli/pkg/api"
	"github.com/scalvert/glean-cli/pkg/theme"
	"github.com/scalvert/glean-cli/pkg/utils"
)

type Document = api.Document
type DocumentMetadata = api.DocumentMetadata
type Person = api.Person
type PersonMetadata = api.PersonMetadata
type RelatedDocument = api.RelatedDocument
type Shortcut = api.Shortcut
type StructuredResult = api.StructuredResult

type Options struct {
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
	Interactive       bool // Whether to run in interactive mode
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

type Response struct {
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
	Results                      []Result                `json:"results"`
	ExperimentIds                []int64                 `json:"experimentIds,omitempty"`
	StructuredResults            []StructuredResult      `json:"structuredResults,omitempty"`
	ResultTabIds                 []string                `json:"resultTabIds,omitempty"`
	ResultTabs                   []ResultTab             `json:"resultTabs,omitempty"`
	BackendTimeMillis            int                     `json:"backendTimeMillis,omitempty"`
	HasMoreResults               bool                    `json:"hasMoreResults"`
}

type Result struct {
	Document               *Document              `json:"document"`
	MustIncludeSuggestions map[string]interface{} `json:"mustIncludeSuggestions"`
	DebugInfo              map[string]interface{} `json:"debugInfo"`
	TrackingToken          string                 `json:"trackingToken"`
	Title                  string                 `json:"title"`
	URL                    string                 `json:"url"`
	Snippets               []Snippet              `json:"snippets"`
	StructuredResults      []interface{}          `json:"structuredResults,omitempty"`
}

type Snippet struct {
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

// resultItem represents a search result in the list
type resultItem struct {
	title  string
	url    string
	desc   string
	source string
	index  int
}

func (i resultItem) Title() string {
	return fmt.Sprintf("%s %s | %s",
		theme.Blue(fmt.Sprint(i.index+1)),
		theme.Blue(utils.FormatDatasource(i.source)),
		theme.Bold(i.title),
	)
}

func (i resultItem) Description() string {
	return theme.Yellow(utils.MaybeAnonymizeURL(i.url)) + "\n" + i.desc
}

func (i resultItem) FilterValue() string {
	return i.title
}
