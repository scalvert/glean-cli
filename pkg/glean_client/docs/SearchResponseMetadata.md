# SearchResponseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RewrittenQuery** | Pointer to **string** | A cleaned up or updated version of the query to be displayed in the query box. Useful for mapping visual facets to search operators. | [optional] 
**SearchedQuery** | Pointer to **string** | The actual query used to perform search and return results. | [optional] 
**SearchedQueryRanges** | Pointer to [**[]TextRange**](TextRange.md) | The bolded ranges within the searched query. | [optional] 
**OriginalQuery** | Pointer to **string** | The query text sent by the client in the request. | [optional] 
**QuerySuggestion** | Pointer to [**QuerySuggestion**](QuerySuggestion.md) |  | [optional] 
**AdditionalQuerySuggestions** | Pointer to [**QuerySuggestionList**](QuerySuggestionList.md) |  | [optional] 
**NegatedTerms** | Pointer to **[]string** | A list of terms that were negated when processing the query. | [optional] 
**ModifiedQueryWasUsed** | Pointer to **bool** | A different query was performed than the one requested. | [optional] 
**OriginalQueryHadNoResults** | Pointer to **bool** | No results were found for the original query. The usage of this bit in conjunction with modifiedQueryWasUsed will dictate whether the full page replacement is 0-result or few-result based. | [optional] 
**SearchWarning** | Pointer to [**SearchWarning**](SearchWarning.md) |  | [optional] 
**TriggeredExpertDetection** | Pointer to **bool** | Whether the query triggered expert detection results in the People tab. | [optional] 

## Methods

### NewSearchResponseMetadata

`func NewSearchResponseMetadata() *SearchResponseMetadata`

NewSearchResponseMetadata instantiates a new SearchResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseMetadataWithDefaults

`func NewSearchResponseMetadataWithDefaults() *SearchResponseMetadata`

NewSearchResponseMetadataWithDefaults instantiates a new SearchResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRewrittenQuery

`func (o *SearchResponseMetadata) GetRewrittenQuery() string`

GetRewrittenQuery returns the RewrittenQuery field if non-nil, zero value otherwise.

### GetRewrittenQueryOk

`func (o *SearchResponseMetadata) GetRewrittenQueryOk() (*string, bool)`

GetRewrittenQueryOk returns a tuple with the RewrittenQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewrittenQuery

`func (o *SearchResponseMetadata) SetRewrittenQuery(v string)`

SetRewrittenQuery sets RewrittenQuery field to given value.

### HasRewrittenQuery

`func (o *SearchResponseMetadata) HasRewrittenQuery() bool`

HasRewrittenQuery returns a boolean if a field has been set.

### GetSearchedQuery

`func (o *SearchResponseMetadata) GetSearchedQuery() string`

GetSearchedQuery returns the SearchedQuery field if non-nil, zero value otherwise.

### GetSearchedQueryOk

`func (o *SearchResponseMetadata) GetSearchedQueryOk() (*string, bool)`

GetSearchedQueryOk returns a tuple with the SearchedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchedQuery

`func (o *SearchResponseMetadata) SetSearchedQuery(v string)`

SetSearchedQuery sets SearchedQuery field to given value.

### HasSearchedQuery

`func (o *SearchResponseMetadata) HasSearchedQuery() bool`

HasSearchedQuery returns a boolean if a field has been set.

### GetSearchedQueryRanges

`func (o *SearchResponseMetadata) GetSearchedQueryRanges() []TextRange`

GetSearchedQueryRanges returns the SearchedQueryRanges field if non-nil, zero value otherwise.

### GetSearchedQueryRangesOk

`func (o *SearchResponseMetadata) GetSearchedQueryRangesOk() (*[]TextRange, bool)`

GetSearchedQueryRangesOk returns a tuple with the SearchedQueryRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchedQueryRanges

`func (o *SearchResponseMetadata) SetSearchedQueryRanges(v []TextRange)`

SetSearchedQueryRanges sets SearchedQueryRanges field to given value.

### HasSearchedQueryRanges

`func (o *SearchResponseMetadata) HasSearchedQueryRanges() bool`

HasSearchedQueryRanges returns a boolean if a field has been set.

### GetOriginalQuery

`func (o *SearchResponseMetadata) GetOriginalQuery() string`

GetOriginalQuery returns the OriginalQuery field if non-nil, zero value otherwise.

### GetOriginalQueryOk

`func (o *SearchResponseMetadata) GetOriginalQueryOk() (*string, bool)`

GetOriginalQueryOk returns a tuple with the OriginalQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalQuery

`func (o *SearchResponseMetadata) SetOriginalQuery(v string)`

SetOriginalQuery sets OriginalQuery field to given value.

### HasOriginalQuery

`func (o *SearchResponseMetadata) HasOriginalQuery() bool`

HasOriginalQuery returns a boolean if a field has been set.

### GetQuerySuggestion

`func (o *SearchResponseMetadata) GetQuerySuggestion() QuerySuggestion`

GetQuerySuggestion returns the QuerySuggestion field if non-nil, zero value otherwise.

### GetQuerySuggestionOk

`func (o *SearchResponseMetadata) GetQuerySuggestionOk() (*QuerySuggestion, bool)`

GetQuerySuggestionOk returns a tuple with the QuerySuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySuggestion

`func (o *SearchResponseMetadata) SetQuerySuggestion(v QuerySuggestion)`

SetQuerySuggestion sets QuerySuggestion field to given value.

### HasQuerySuggestion

`func (o *SearchResponseMetadata) HasQuerySuggestion() bool`

HasQuerySuggestion returns a boolean if a field has been set.

### GetAdditionalQuerySuggestions

`func (o *SearchResponseMetadata) GetAdditionalQuerySuggestions() QuerySuggestionList`

GetAdditionalQuerySuggestions returns the AdditionalQuerySuggestions field if non-nil, zero value otherwise.

### GetAdditionalQuerySuggestionsOk

`func (o *SearchResponseMetadata) GetAdditionalQuerySuggestionsOk() (*QuerySuggestionList, bool)`

GetAdditionalQuerySuggestionsOk returns a tuple with the AdditionalQuerySuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalQuerySuggestions

`func (o *SearchResponseMetadata) SetAdditionalQuerySuggestions(v QuerySuggestionList)`

SetAdditionalQuerySuggestions sets AdditionalQuerySuggestions field to given value.

### HasAdditionalQuerySuggestions

`func (o *SearchResponseMetadata) HasAdditionalQuerySuggestions() bool`

HasAdditionalQuerySuggestions returns a boolean if a field has been set.

### GetNegatedTerms

`func (o *SearchResponseMetadata) GetNegatedTerms() []string`

GetNegatedTerms returns the NegatedTerms field if non-nil, zero value otherwise.

### GetNegatedTermsOk

`func (o *SearchResponseMetadata) GetNegatedTermsOk() (*[]string, bool)`

GetNegatedTermsOk returns a tuple with the NegatedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegatedTerms

`func (o *SearchResponseMetadata) SetNegatedTerms(v []string)`

SetNegatedTerms sets NegatedTerms field to given value.

### HasNegatedTerms

`func (o *SearchResponseMetadata) HasNegatedTerms() bool`

HasNegatedTerms returns a boolean if a field has been set.

### GetModifiedQueryWasUsed

`func (o *SearchResponseMetadata) GetModifiedQueryWasUsed() bool`

GetModifiedQueryWasUsed returns the ModifiedQueryWasUsed field if non-nil, zero value otherwise.

### GetModifiedQueryWasUsedOk

`func (o *SearchResponseMetadata) GetModifiedQueryWasUsedOk() (*bool, bool)`

GetModifiedQueryWasUsedOk returns a tuple with the ModifiedQueryWasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedQueryWasUsed

`func (o *SearchResponseMetadata) SetModifiedQueryWasUsed(v bool)`

SetModifiedQueryWasUsed sets ModifiedQueryWasUsed field to given value.

### HasModifiedQueryWasUsed

`func (o *SearchResponseMetadata) HasModifiedQueryWasUsed() bool`

HasModifiedQueryWasUsed returns a boolean if a field has been set.

### GetOriginalQueryHadNoResults

`func (o *SearchResponseMetadata) GetOriginalQueryHadNoResults() bool`

GetOriginalQueryHadNoResults returns the OriginalQueryHadNoResults field if non-nil, zero value otherwise.

### GetOriginalQueryHadNoResultsOk

`func (o *SearchResponseMetadata) GetOriginalQueryHadNoResultsOk() (*bool, bool)`

GetOriginalQueryHadNoResultsOk returns a tuple with the OriginalQueryHadNoResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalQueryHadNoResults

`func (o *SearchResponseMetadata) SetOriginalQueryHadNoResults(v bool)`

SetOriginalQueryHadNoResults sets OriginalQueryHadNoResults field to given value.

### HasOriginalQueryHadNoResults

`func (o *SearchResponseMetadata) HasOriginalQueryHadNoResults() bool`

HasOriginalQueryHadNoResults returns a boolean if a field has been set.

### GetSearchWarning

`func (o *SearchResponseMetadata) GetSearchWarning() SearchWarning`

GetSearchWarning returns the SearchWarning field if non-nil, zero value otherwise.

### GetSearchWarningOk

`func (o *SearchResponseMetadata) GetSearchWarningOk() (*SearchWarning, bool)`

GetSearchWarningOk returns a tuple with the SearchWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWarning

`func (o *SearchResponseMetadata) SetSearchWarning(v SearchWarning)`

SetSearchWarning sets SearchWarning field to given value.

### HasSearchWarning

`func (o *SearchResponseMetadata) HasSearchWarning() bool`

HasSearchWarning returns a boolean if a field has been set.

### GetTriggeredExpertDetection

`func (o *SearchResponseMetadata) GetTriggeredExpertDetection() bool`

GetTriggeredExpertDetection returns the TriggeredExpertDetection field if non-nil, zero value otherwise.

### GetTriggeredExpertDetectionOk

`func (o *SearchResponseMetadata) GetTriggeredExpertDetectionOk() (*bool, bool)`

GetTriggeredExpertDetectionOk returns a tuple with the TriggeredExpertDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredExpertDetection

`func (o *SearchResponseMetadata) SetTriggeredExpertDetection(v bool)`

SetTriggeredExpertDetection sets TriggeredExpertDetection field to given value.

### HasTriggeredExpertDetection

`func (o *SearchResponseMetadata) HasTriggeredExpertDetection() bool`

HasTriggeredExpertDetection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


