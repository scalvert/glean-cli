# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingToken** | Pointer to **string** | A token that should be passed for additional requests related to this request (such as more results requests). | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**Results** | Pointer to [**[]SearchResult**](SearchResult.md) |  | [optional] 
**StructuredResults** | Pointer to [**[]StructuredResult**](StructuredResult.md) |  | [optional] 
**GeneratedQnaResult** | Pointer to [**GeneratedQna**](GeneratedQna.md) |  | [optional] 
**ErrorInfo** | Pointer to [**ErrorInfo**](ErrorInfo.md) |  | [optional] 
**RequestID** | Pointer to **string** | A platform-generated request ID to correlate backend logs. | [optional] 
**BackendTimeMillis** | Pointer to **int64** | Time in milliseconds the backend took to respond to the request. | [optional] 
**ExperimentIds** | Pointer to **[]int64** | List of experiment ids for the corresponding request. | [optional] 
**Metadata** | Pointer to [**SearchResponseMetadata**](SearchResponseMetadata.md) |  | [optional] 
**FacetResults** | Pointer to [**[]FacetResult**](FacetResult.md) |  | [optional] 
**ResultTabs** | Pointer to [**[]ResultTab**](ResultTab.md) | All result tabs available for the current query. Populated if QUERY_METADATA is specified in the request. | [optional] 
**ResultTabIds** | Pointer to **[]string** | The unique IDs of the result tabs to which this response belongs. | [optional] 
**ResultsDescription** | Pointer to [**ResultsDescription**](ResultsDescription.md) |  | [optional] 
**RewrittenFacetFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | The actual applied facet filters based on the operators and facetFilters in the query. Useful for mapping typed operators to visual facets. | [optional] 
**Cursor** | Pointer to **string** | Cursor that indicates the start of the next page of results. To be passed in \&quot;more\&quot; requests for this query. | [optional] 
**HasMoreResults** | Pointer to **bool** | Whether more results are available. Use cursor to retrieve them. | [optional] 

## Methods

### NewSearchResponse

`func NewSearchResponse() *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingToken

`func (o *SearchResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *SearchResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *SearchResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *SearchResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetSessionInfo

`func (o *SearchResponse) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *SearchResponse) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *SearchResponse) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *SearchResponse) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetResults

`func (o *SearchResponse) GetResults() []SearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchResponse) GetResultsOk() (*[]SearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchResponse) SetResults(v []SearchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStructuredResults

`func (o *SearchResponse) GetStructuredResults() []StructuredResult`

GetStructuredResults returns the StructuredResults field if non-nil, zero value otherwise.

### GetStructuredResultsOk

`func (o *SearchResponse) GetStructuredResultsOk() (*[]StructuredResult, bool)`

GetStructuredResultsOk returns a tuple with the StructuredResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredResults

`func (o *SearchResponse) SetStructuredResults(v []StructuredResult)`

SetStructuredResults sets StructuredResults field to given value.

### HasStructuredResults

`func (o *SearchResponse) HasStructuredResults() bool`

HasStructuredResults returns a boolean if a field has been set.

### GetGeneratedQnaResult

`func (o *SearchResponse) GetGeneratedQnaResult() GeneratedQna`

GetGeneratedQnaResult returns the GeneratedQnaResult field if non-nil, zero value otherwise.

### GetGeneratedQnaResultOk

`func (o *SearchResponse) GetGeneratedQnaResultOk() (*GeneratedQna, bool)`

GetGeneratedQnaResultOk returns a tuple with the GeneratedQnaResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedQnaResult

`func (o *SearchResponse) SetGeneratedQnaResult(v GeneratedQna)`

SetGeneratedQnaResult sets GeneratedQnaResult field to given value.

### HasGeneratedQnaResult

`func (o *SearchResponse) HasGeneratedQnaResult() bool`

HasGeneratedQnaResult returns a boolean if a field has been set.

### GetErrorInfo

`func (o *SearchResponse) GetErrorInfo() ErrorInfo`

GetErrorInfo returns the ErrorInfo field if non-nil, zero value otherwise.

### GetErrorInfoOk

`func (o *SearchResponse) GetErrorInfoOk() (*ErrorInfo, bool)`

GetErrorInfoOk returns a tuple with the ErrorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorInfo

`func (o *SearchResponse) SetErrorInfo(v ErrorInfo)`

SetErrorInfo sets ErrorInfo field to given value.

### HasErrorInfo

`func (o *SearchResponse) HasErrorInfo() bool`

HasErrorInfo returns a boolean if a field has been set.

### GetRequestID

`func (o *SearchResponse) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *SearchResponse) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *SearchResponse) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *SearchResponse) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.

### GetBackendTimeMillis

`func (o *SearchResponse) GetBackendTimeMillis() int64`

GetBackendTimeMillis returns the BackendTimeMillis field if non-nil, zero value otherwise.

### GetBackendTimeMillisOk

`func (o *SearchResponse) GetBackendTimeMillisOk() (*int64, bool)`

GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTimeMillis

`func (o *SearchResponse) SetBackendTimeMillis(v int64)`

SetBackendTimeMillis sets BackendTimeMillis field to given value.

### HasBackendTimeMillis

`func (o *SearchResponse) HasBackendTimeMillis() bool`

HasBackendTimeMillis returns a boolean if a field has been set.

### GetExperimentIds

`func (o *SearchResponse) GetExperimentIds() []int64`

GetExperimentIds returns the ExperimentIds field if non-nil, zero value otherwise.

### GetExperimentIdsOk

`func (o *SearchResponse) GetExperimentIdsOk() (*[]int64, bool)`

GetExperimentIdsOk returns a tuple with the ExperimentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIds

`func (o *SearchResponse) SetExperimentIds(v []int64)`

SetExperimentIds sets ExperimentIds field to given value.

### HasExperimentIds

`func (o *SearchResponse) HasExperimentIds() bool`

HasExperimentIds returns a boolean if a field has been set.

### GetMetadata

`func (o *SearchResponse) GetMetadata() SearchResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SearchResponse) GetMetadataOk() (*SearchResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SearchResponse) SetMetadata(v SearchResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SearchResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFacetResults

`func (o *SearchResponse) GetFacetResults() []FacetResult`

GetFacetResults returns the FacetResults field if non-nil, zero value otherwise.

### GetFacetResultsOk

`func (o *SearchResponse) GetFacetResultsOk() (*[]FacetResult, bool)`

GetFacetResultsOk returns a tuple with the FacetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetResults

`func (o *SearchResponse) SetFacetResults(v []FacetResult)`

SetFacetResults sets FacetResults field to given value.

### HasFacetResults

`func (o *SearchResponse) HasFacetResults() bool`

HasFacetResults returns a boolean if a field has been set.

### GetResultTabs

`func (o *SearchResponse) GetResultTabs() []ResultTab`

GetResultTabs returns the ResultTabs field if non-nil, zero value otherwise.

### GetResultTabsOk

`func (o *SearchResponse) GetResultTabsOk() (*[]ResultTab, bool)`

GetResultTabsOk returns a tuple with the ResultTabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTabs

`func (o *SearchResponse) SetResultTabs(v []ResultTab)`

SetResultTabs sets ResultTabs field to given value.

### HasResultTabs

`func (o *SearchResponse) HasResultTabs() bool`

HasResultTabs returns a boolean if a field has been set.

### GetResultTabIds

`func (o *SearchResponse) GetResultTabIds() []string`

GetResultTabIds returns the ResultTabIds field if non-nil, zero value otherwise.

### GetResultTabIdsOk

`func (o *SearchResponse) GetResultTabIdsOk() (*[]string, bool)`

GetResultTabIdsOk returns a tuple with the ResultTabIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTabIds

`func (o *SearchResponse) SetResultTabIds(v []string)`

SetResultTabIds sets ResultTabIds field to given value.

### HasResultTabIds

`func (o *SearchResponse) HasResultTabIds() bool`

HasResultTabIds returns a boolean if a field has been set.

### GetResultsDescription

`func (o *SearchResponse) GetResultsDescription() ResultsDescription`

GetResultsDescription returns the ResultsDescription field if non-nil, zero value otherwise.

### GetResultsDescriptionOk

`func (o *SearchResponse) GetResultsDescriptionOk() (*ResultsDescription, bool)`

GetResultsDescriptionOk returns a tuple with the ResultsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsDescription

`func (o *SearchResponse) SetResultsDescription(v ResultsDescription)`

SetResultsDescription sets ResultsDescription field to given value.

### HasResultsDescription

`func (o *SearchResponse) HasResultsDescription() bool`

HasResultsDescription returns a boolean if a field has been set.

### GetRewrittenFacetFilters

`func (o *SearchResponse) GetRewrittenFacetFilters() []FacetFilter`

GetRewrittenFacetFilters returns the RewrittenFacetFilters field if non-nil, zero value otherwise.

### GetRewrittenFacetFiltersOk

`func (o *SearchResponse) GetRewrittenFacetFiltersOk() (*[]FacetFilter, bool)`

GetRewrittenFacetFiltersOk returns a tuple with the RewrittenFacetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewrittenFacetFilters

`func (o *SearchResponse) SetRewrittenFacetFilters(v []FacetFilter)`

SetRewrittenFacetFilters sets RewrittenFacetFilters field to given value.

### HasRewrittenFacetFilters

`func (o *SearchResponse) HasRewrittenFacetFilters() bool`

HasRewrittenFacetFilters returns a boolean if a field has been set.

### GetCursor

`func (o *SearchResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *SearchResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *SearchResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *SearchResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetHasMoreResults

`func (o *SearchResponse) GetHasMoreResults() bool`

GetHasMoreResults returns the HasMoreResults field if non-nil, zero value otherwise.

### GetHasMoreResultsOk

`func (o *SearchResponse) GetHasMoreResultsOk() (*bool, bool)`

GetHasMoreResultsOk returns a tuple with the HasMoreResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreResults

`func (o *SearchResponse) SetHasMoreResults(v bool)`

SetHasMoreResults sets HasMoreResults field to given value.

### HasHasMoreResults

`func (o *SearchResponse) HasHasMoreResults() bool`

HasHasMoreResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


