# SearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructuredResults** | Pointer to [**[]StructuredResult**](StructuredResult.md) | An array of entities in the work graph retrieved via a data request. | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular result in this particular query. To be used for /feedback reporting. | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**NativeAppUrl** | Pointer to **string** | A deep link, if available, into the datasource&#39;s native application for the user&#39;s platform (e.g. slack://...). | [optional] 
**Snippets** | Pointer to [**[]SearchResultSnippet**](SearchResultSnippet.md) | Text content from the result document which contains search query terms, if available. | [optional] 
**ExpandedSnippets** | Pointer to **[]string** | The expanded snippets for this result. This is only populated if the query has the expand_snippets parameter set to true. | [optional] 
**FullText** | Pointer to **string** | The full body text of the result if not already contained in the snippets | [optional] 
**FullTextList** | Pointer to **[]string** | The full body text of the result if not already contained in the snippets; each item in the array represents a separate line in the original text | [optional] 
**RelatedResults** | Pointer to [**[]RelatedDocuments**](RelatedDocuments.md) | A list of results related to this search result. Eg. for conversation results it contains individual messages from the conversation document which will be shown on SERP. | [optional] 
**ClusteredResults** | Pointer to [**[]SearchResult**](SearchResult.md) | A list of results that should be displayed as associated with this result. | [optional] 
**AllClusteredResults** | Pointer to [**[]ClusterGroup**](ClusterGroup.md) | A list of results that should be displayed as associated with this result. | [optional] 
**AttachmentCount** | Pointer to **int32** | The total number of attachments. | [optional] 
**Attachments** | Pointer to [**[]SearchResult**](SearchResult.md) | A (potentially partial) list of results representing documents attached to the main result document. | [optional] 
**BacklinkResults** | Pointer to [**[]SearchResult**](SearchResult.md) | A list of results that should be displayed as backlinks of this result in reverse chronological order. | [optional] 
**ClusterType** | Pointer to [**ClusterTypeEnum**](ClusterTypeEnum.md) |  | [optional] 
**MustIncludeSuggestions** | Pointer to [**QuerySuggestionList**](QuerySuggestionList.md) |  | [optional] 
**QuerySuggestion** | Pointer to [**QuerySuggestion**](QuerySuggestion.md) |  | [optional] 
**Prominence** | Pointer to [**SearchResultProminenceEnum**](SearchResultProminenceEnum.md) |  | [optional] 
**AttachmentContext** | Pointer to **string** | Additional context for the relationship between the result and the document it&#39;s attached to. | [optional] 
**Pins** | Pointer to [**[]PinDocument**](PinDocument.md) | A list of pins associated with this search result. | [optional] 

## Methods

### NewSearchResult

`func NewSearchResult(url string, ) *SearchResult`

NewSearchResult instantiates a new SearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWithDefaults

`func NewSearchResultWithDefaults() *SearchResult`

NewSearchResultWithDefaults instantiates a new SearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructuredResults

`func (o *SearchResult) GetStructuredResults() []StructuredResult`

GetStructuredResults returns the StructuredResults field if non-nil, zero value otherwise.

### GetStructuredResultsOk

`func (o *SearchResult) GetStructuredResultsOk() (*[]StructuredResult, bool)`

GetStructuredResultsOk returns a tuple with the StructuredResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredResults

`func (o *SearchResult) SetStructuredResults(v []StructuredResult)`

SetStructuredResults sets StructuredResults field to given value.

### HasStructuredResults

`func (o *SearchResult) HasStructuredResults() bool`

HasStructuredResults returns a boolean if a field has been set.

### GetTrackingToken

`func (o *SearchResult) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *SearchResult) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *SearchResult) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *SearchResult) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetDocument

`func (o *SearchResult) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *SearchResult) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *SearchResult) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *SearchResult) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetTitle

`func (o *SearchResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SearchResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SearchResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SearchResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *SearchResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SearchResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SearchResult) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNativeAppUrl

`func (o *SearchResult) GetNativeAppUrl() string`

GetNativeAppUrl returns the NativeAppUrl field if non-nil, zero value otherwise.

### GetNativeAppUrlOk

`func (o *SearchResult) GetNativeAppUrlOk() (*string, bool)`

GetNativeAppUrlOk returns a tuple with the NativeAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeAppUrl

`func (o *SearchResult) SetNativeAppUrl(v string)`

SetNativeAppUrl sets NativeAppUrl field to given value.

### HasNativeAppUrl

`func (o *SearchResult) HasNativeAppUrl() bool`

HasNativeAppUrl returns a boolean if a field has been set.

### GetSnippets

`func (o *SearchResult) GetSnippets() []SearchResultSnippet`

GetSnippets returns the Snippets field if non-nil, zero value otherwise.

### GetSnippetsOk

`func (o *SearchResult) GetSnippetsOk() (*[]SearchResultSnippet, bool)`

GetSnippetsOk returns a tuple with the Snippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippets

`func (o *SearchResult) SetSnippets(v []SearchResultSnippet)`

SetSnippets sets Snippets field to given value.

### HasSnippets

`func (o *SearchResult) HasSnippets() bool`

HasSnippets returns a boolean if a field has been set.

### GetExpandedSnippets

`func (o *SearchResult) GetExpandedSnippets() []string`

GetExpandedSnippets returns the ExpandedSnippets field if non-nil, zero value otherwise.

### GetExpandedSnippetsOk

`func (o *SearchResult) GetExpandedSnippetsOk() (*[]string, bool)`

GetExpandedSnippetsOk returns a tuple with the ExpandedSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandedSnippets

`func (o *SearchResult) SetExpandedSnippets(v []string)`

SetExpandedSnippets sets ExpandedSnippets field to given value.

### HasExpandedSnippets

`func (o *SearchResult) HasExpandedSnippets() bool`

HasExpandedSnippets returns a boolean if a field has been set.

### GetFullText

`func (o *SearchResult) GetFullText() string`

GetFullText returns the FullText field if non-nil, zero value otherwise.

### GetFullTextOk

`func (o *SearchResult) GetFullTextOk() (*string, bool)`

GetFullTextOk returns a tuple with the FullText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullText

`func (o *SearchResult) SetFullText(v string)`

SetFullText sets FullText field to given value.

### HasFullText

`func (o *SearchResult) HasFullText() bool`

HasFullText returns a boolean if a field has been set.

### GetFullTextList

`func (o *SearchResult) GetFullTextList() []string`

GetFullTextList returns the FullTextList field if non-nil, zero value otherwise.

### GetFullTextListOk

`func (o *SearchResult) GetFullTextListOk() (*[]string, bool)`

GetFullTextListOk returns a tuple with the FullTextList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullTextList

`func (o *SearchResult) SetFullTextList(v []string)`

SetFullTextList sets FullTextList field to given value.

### HasFullTextList

`func (o *SearchResult) HasFullTextList() bool`

HasFullTextList returns a boolean if a field has been set.

### GetRelatedResults

`func (o *SearchResult) GetRelatedResults() []RelatedDocuments`

GetRelatedResults returns the RelatedResults field if non-nil, zero value otherwise.

### GetRelatedResultsOk

`func (o *SearchResult) GetRelatedResultsOk() (*[]RelatedDocuments, bool)`

GetRelatedResultsOk returns a tuple with the RelatedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResults

`func (o *SearchResult) SetRelatedResults(v []RelatedDocuments)`

SetRelatedResults sets RelatedResults field to given value.

### HasRelatedResults

`func (o *SearchResult) HasRelatedResults() bool`

HasRelatedResults returns a boolean if a field has been set.

### GetClusteredResults

`func (o *SearchResult) GetClusteredResults() []SearchResult`

GetClusteredResults returns the ClusteredResults field if non-nil, zero value otherwise.

### GetClusteredResultsOk

`func (o *SearchResult) GetClusteredResultsOk() (*[]SearchResult, bool)`

GetClusteredResultsOk returns a tuple with the ClusteredResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusteredResults

`func (o *SearchResult) SetClusteredResults(v []SearchResult)`

SetClusteredResults sets ClusteredResults field to given value.

### HasClusteredResults

`func (o *SearchResult) HasClusteredResults() bool`

HasClusteredResults returns a boolean if a field has been set.

### GetAllClusteredResults

`func (o *SearchResult) GetAllClusteredResults() []ClusterGroup`

GetAllClusteredResults returns the AllClusteredResults field if non-nil, zero value otherwise.

### GetAllClusteredResultsOk

`func (o *SearchResult) GetAllClusteredResultsOk() (*[]ClusterGroup, bool)`

GetAllClusteredResultsOk returns a tuple with the AllClusteredResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllClusteredResults

`func (o *SearchResult) SetAllClusteredResults(v []ClusterGroup)`

SetAllClusteredResults sets AllClusteredResults field to given value.

### HasAllClusteredResults

`func (o *SearchResult) HasAllClusteredResults() bool`

HasAllClusteredResults returns a boolean if a field has been set.

### GetAttachmentCount

`func (o *SearchResult) GetAttachmentCount() int32`

GetAttachmentCount returns the AttachmentCount field if non-nil, zero value otherwise.

### GetAttachmentCountOk

`func (o *SearchResult) GetAttachmentCountOk() (*int32, bool)`

GetAttachmentCountOk returns a tuple with the AttachmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentCount

`func (o *SearchResult) SetAttachmentCount(v int32)`

SetAttachmentCount sets AttachmentCount field to given value.

### HasAttachmentCount

`func (o *SearchResult) HasAttachmentCount() bool`

HasAttachmentCount returns a boolean if a field has been set.

### GetAttachments

`func (o *SearchResult) GetAttachments() []SearchResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SearchResult) GetAttachmentsOk() (*[]SearchResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SearchResult) SetAttachments(v []SearchResult)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SearchResult) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBacklinkResults

`func (o *SearchResult) GetBacklinkResults() []SearchResult`

GetBacklinkResults returns the BacklinkResults field if non-nil, zero value otherwise.

### GetBacklinkResultsOk

`func (o *SearchResult) GetBacklinkResultsOk() (*[]SearchResult, bool)`

GetBacklinkResultsOk returns a tuple with the BacklinkResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacklinkResults

`func (o *SearchResult) SetBacklinkResults(v []SearchResult)`

SetBacklinkResults sets BacklinkResults field to given value.

### HasBacklinkResults

`func (o *SearchResult) HasBacklinkResults() bool`

HasBacklinkResults returns a boolean if a field has been set.

### GetClusterType

`func (o *SearchResult) GetClusterType() ClusterTypeEnum`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *SearchResult) GetClusterTypeOk() (*ClusterTypeEnum, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *SearchResult) SetClusterType(v ClusterTypeEnum)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *SearchResult) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetMustIncludeSuggestions

`func (o *SearchResult) GetMustIncludeSuggestions() QuerySuggestionList`

GetMustIncludeSuggestions returns the MustIncludeSuggestions field if non-nil, zero value otherwise.

### GetMustIncludeSuggestionsOk

`func (o *SearchResult) GetMustIncludeSuggestionsOk() (*QuerySuggestionList, bool)`

GetMustIncludeSuggestionsOk returns a tuple with the MustIncludeSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustIncludeSuggestions

`func (o *SearchResult) SetMustIncludeSuggestions(v QuerySuggestionList)`

SetMustIncludeSuggestions sets MustIncludeSuggestions field to given value.

### HasMustIncludeSuggestions

`func (o *SearchResult) HasMustIncludeSuggestions() bool`

HasMustIncludeSuggestions returns a boolean if a field has been set.

### GetQuerySuggestion

`func (o *SearchResult) GetQuerySuggestion() QuerySuggestion`

GetQuerySuggestion returns the QuerySuggestion field if non-nil, zero value otherwise.

### GetQuerySuggestionOk

`func (o *SearchResult) GetQuerySuggestionOk() (*QuerySuggestion, bool)`

GetQuerySuggestionOk returns a tuple with the QuerySuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySuggestion

`func (o *SearchResult) SetQuerySuggestion(v QuerySuggestion)`

SetQuerySuggestion sets QuerySuggestion field to given value.

### HasQuerySuggestion

`func (o *SearchResult) HasQuerySuggestion() bool`

HasQuerySuggestion returns a boolean if a field has been set.

### GetProminence

`func (o *SearchResult) GetProminence() SearchResultProminenceEnum`

GetProminence returns the Prominence field if non-nil, zero value otherwise.

### GetProminenceOk

`func (o *SearchResult) GetProminenceOk() (*SearchResultProminenceEnum, bool)`

GetProminenceOk returns a tuple with the Prominence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProminence

`func (o *SearchResult) SetProminence(v SearchResultProminenceEnum)`

SetProminence sets Prominence field to given value.

### HasProminence

`func (o *SearchResult) HasProminence() bool`

HasProminence returns a boolean if a field has been set.

### GetAttachmentContext

`func (o *SearchResult) GetAttachmentContext() string`

GetAttachmentContext returns the AttachmentContext field if non-nil, zero value otherwise.

### GetAttachmentContextOk

`func (o *SearchResult) GetAttachmentContextOk() (*string, bool)`

GetAttachmentContextOk returns a tuple with the AttachmentContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentContext

`func (o *SearchResult) SetAttachmentContext(v string)`

SetAttachmentContext sets AttachmentContext field to given value.

### HasAttachmentContext

`func (o *SearchResult) HasAttachmentContext() bool`

HasAttachmentContext returns a boolean if a field has been set.

### GetPins

`func (o *SearchResult) GetPins() []PinDocument`

GetPins returns the Pins field if non-nil, zero value otherwise.

### GetPinsOk

`func (o *SearchResult) GetPinsOk() (*[]PinDocument, bool)`

GetPinsOk returns a tuple with the Pins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPins

`func (o *SearchResult) SetPins(v []PinDocument)`

SetPins sets Pins field to given value.

### HasPins

`func (o *SearchResult) HasPins() bool`

HasPins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


