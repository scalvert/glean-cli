package search

// Options holds all CLI-level search options parsed from flags.
// These are mapped to the SDK's components.SearchRequest inside performSearch.
type Options struct {
	InputDetails   *SearchInputDetails `json:"inputDetails,omitempty"`
	SessionInfo    *SessionInfo        `json:"sessionInfo,omitempty"`
	RequestOptions *RequestOptions     `json:"requestOptions,omitempty"`
	Query          string              `json:"query"`
	Cursor         string              `json:"cursor,omitempty"`
	Timestamp      string              `json:"timestamp,omitempty"`
	TrackingToken  string              `json:"trackingToken,omitempty"`
	ResultTabIds   []string            `json:"resultTabIds,omitempty"`
	PageSize       int                 `json:"pageSize,omitempty"`
	MaxSnippetSize int                 `json:"maxSnippetSize,omitempty"`
	TimeoutMillis  int                 `json:"timeoutMillis,omitempty"`
	DisableSpellcheck bool             `json:"disableSpellcheck,omitempty"`
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

type RestrictionFilters struct{}
type FacetBucketFilter struct{}

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
