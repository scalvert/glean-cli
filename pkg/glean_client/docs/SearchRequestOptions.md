# SearchRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasourceFilter** | Pointer to **string** | Filter results to a single datasource name (e.g. gmail, slack). All results are returned if missing. | [optional] 
**DatasourcesFilter** | Pointer to **[]string** | Filter results to one or more datasources (e.g. gmail, slack). All results are returned if missing. | [optional] 
**QueryOverridesFacetFilters** | Pointer to **bool** | If true, the operators in the query are taken to override any operators in facetFilters in the case of conflict. This is used to correctly set rewrittenFacetFilters and rewrittenQuery. | [optional] 
**FacetFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | A list of filters for the query. An AND is assumed between different facetFilters. For example, owner Sumeet and type Spreadsheet shows documents that are by Sumeet AND are Spreadsheets. | [optional] 
**FacetFilterSets** | Pointer to [**[]FacetFilterSet**](FacetFilterSet.md) | A list of facet filter sets that will be OR&#39;ed together. SearchRequestOptions where both facetFilterSets and facetFilters set are considered as bad request. Callers should set only one of these fields. | [optional] 
**FacetBucketFilter** | Pointer to [**FacetBucketFilter**](FacetBucketFilter.md) |  | [optional] 
**FacetBucketSize** | **int32** | The maximum number of FacetBuckets to return in each FacetResult. | 
**DefaultFacets** | Pointer to **[]string** | Facets for which FacetResults should be fetched and that don&#39;t apply to a particular datasource. If specified, these values will replace the standard default facets (last_updated_at, from, etc.). The requested facets will be returned alongside datasource-specific facets if searching a single datasource. | [optional] 
**AuthTokens** | Pointer to [**[]AuthToken**](AuthToken.md) | Auth tokens which may be used for non-indexed, federated results (e.g. Gmail). | [optional] 
**FetchAllDatasourceCounts** | Pointer to **bool** | Hints that the QE should return result counts (via the datasource facet result) for all supported datasources, rather than just those specified in the datasource[s]Filter | [optional] 
**ResponseHints** | Pointer to **[]string** | Array of hints containing which fields should be populated in the response. | [optional] 
**TimezoneOffset** | Pointer to **int32** | The offset of the client&#39;s timezone in minutes from UTC. e.g. PDT is -420 because it&#39;s 7 hours behind UTC. | [optional] 
**ForceNegation** | Pointer to **bool** | Whether or not to force not ignoring of negation, i.e. force negated terms to be negated. | [optional] 
**DisableSpellcheck** | Pointer to **bool** | Whether or not to disable spellcheck. | [optional] 
**DisableQueryAutocorrect** | Pointer to **bool** | Disables automatic adjustment of the input query for spelling corrections or other reasons. | [optional] 
**ReturnLlmContentOverSnippets** | Pointer to **bool** | [beta] Enables expanded content to be returned for LLM usage. The size of content per result returned should be modified using maxSnippetSize. Server may return less or more than what is specified in maxSnippetSize. For more details, https://docs.google.com/document/d/1CTOLSxWWT9WDEnHVLoCUaxbGYyXYP8kctPRF-RluSQY/edit. Requires sufficient permissions. | [optional] 

## Methods

### NewSearchRequestOptions

`func NewSearchRequestOptions(facetBucketSize int32, ) *SearchRequestOptions`

NewSearchRequestOptions instantiates a new SearchRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestOptionsWithDefaults

`func NewSearchRequestOptionsWithDefaults() *SearchRequestOptions`

NewSearchRequestOptionsWithDefaults instantiates a new SearchRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasourceFilter

`func (o *SearchRequestOptions) GetDatasourceFilter() string`

GetDatasourceFilter returns the DatasourceFilter field if non-nil, zero value otherwise.

### GetDatasourceFilterOk

`func (o *SearchRequestOptions) GetDatasourceFilterOk() (*string, bool)`

GetDatasourceFilterOk returns a tuple with the DatasourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceFilter

`func (o *SearchRequestOptions) SetDatasourceFilter(v string)`

SetDatasourceFilter sets DatasourceFilter field to given value.

### HasDatasourceFilter

`func (o *SearchRequestOptions) HasDatasourceFilter() bool`

HasDatasourceFilter returns a boolean if a field has been set.

### GetDatasourcesFilter

`func (o *SearchRequestOptions) GetDatasourcesFilter() []string`

GetDatasourcesFilter returns the DatasourcesFilter field if non-nil, zero value otherwise.

### GetDatasourcesFilterOk

`func (o *SearchRequestOptions) GetDatasourcesFilterOk() (*[]string, bool)`

GetDatasourcesFilterOk returns a tuple with the DatasourcesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourcesFilter

`func (o *SearchRequestOptions) SetDatasourcesFilter(v []string)`

SetDatasourcesFilter sets DatasourcesFilter field to given value.

### HasDatasourcesFilter

`func (o *SearchRequestOptions) HasDatasourcesFilter() bool`

HasDatasourcesFilter returns a boolean if a field has been set.

### GetQueryOverridesFacetFilters

`func (o *SearchRequestOptions) GetQueryOverridesFacetFilters() bool`

GetQueryOverridesFacetFilters returns the QueryOverridesFacetFilters field if non-nil, zero value otherwise.

### GetQueryOverridesFacetFiltersOk

`func (o *SearchRequestOptions) GetQueryOverridesFacetFiltersOk() (*bool, bool)`

GetQueryOverridesFacetFiltersOk returns a tuple with the QueryOverridesFacetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryOverridesFacetFilters

`func (o *SearchRequestOptions) SetQueryOverridesFacetFilters(v bool)`

SetQueryOverridesFacetFilters sets QueryOverridesFacetFilters field to given value.

### HasQueryOverridesFacetFilters

`func (o *SearchRequestOptions) HasQueryOverridesFacetFilters() bool`

HasQueryOverridesFacetFilters returns a boolean if a field has been set.

### GetFacetFilters

`func (o *SearchRequestOptions) GetFacetFilters() []FacetFilter`

GetFacetFilters returns the FacetFilters field if non-nil, zero value otherwise.

### GetFacetFiltersOk

`func (o *SearchRequestOptions) GetFacetFiltersOk() (*[]FacetFilter, bool)`

GetFacetFiltersOk returns a tuple with the FacetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetFilters

`func (o *SearchRequestOptions) SetFacetFilters(v []FacetFilter)`

SetFacetFilters sets FacetFilters field to given value.

### HasFacetFilters

`func (o *SearchRequestOptions) HasFacetFilters() bool`

HasFacetFilters returns a boolean if a field has been set.

### GetFacetFilterSets

`func (o *SearchRequestOptions) GetFacetFilterSets() []FacetFilterSet`

GetFacetFilterSets returns the FacetFilterSets field if non-nil, zero value otherwise.

### GetFacetFilterSetsOk

`func (o *SearchRequestOptions) GetFacetFilterSetsOk() (*[]FacetFilterSet, bool)`

GetFacetFilterSetsOk returns a tuple with the FacetFilterSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetFilterSets

`func (o *SearchRequestOptions) SetFacetFilterSets(v []FacetFilterSet)`

SetFacetFilterSets sets FacetFilterSets field to given value.

### HasFacetFilterSets

`func (o *SearchRequestOptions) HasFacetFilterSets() bool`

HasFacetFilterSets returns a boolean if a field has been set.

### GetFacetBucketFilter

`func (o *SearchRequestOptions) GetFacetBucketFilter() FacetBucketFilter`

GetFacetBucketFilter returns the FacetBucketFilter field if non-nil, zero value otherwise.

### GetFacetBucketFilterOk

`func (o *SearchRequestOptions) GetFacetBucketFilterOk() (*FacetBucketFilter, bool)`

GetFacetBucketFilterOk returns a tuple with the FacetBucketFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetBucketFilter

`func (o *SearchRequestOptions) SetFacetBucketFilter(v FacetBucketFilter)`

SetFacetBucketFilter sets FacetBucketFilter field to given value.

### HasFacetBucketFilter

`func (o *SearchRequestOptions) HasFacetBucketFilter() bool`

HasFacetBucketFilter returns a boolean if a field has been set.

### GetFacetBucketSize

`func (o *SearchRequestOptions) GetFacetBucketSize() int32`

GetFacetBucketSize returns the FacetBucketSize field if non-nil, zero value otherwise.

### GetFacetBucketSizeOk

`func (o *SearchRequestOptions) GetFacetBucketSizeOk() (*int32, bool)`

GetFacetBucketSizeOk returns a tuple with the FacetBucketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetBucketSize

`func (o *SearchRequestOptions) SetFacetBucketSize(v int32)`

SetFacetBucketSize sets FacetBucketSize field to given value.


### GetDefaultFacets

`func (o *SearchRequestOptions) GetDefaultFacets() []string`

GetDefaultFacets returns the DefaultFacets field if non-nil, zero value otherwise.

### GetDefaultFacetsOk

`func (o *SearchRequestOptions) GetDefaultFacetsOk() (*[]string, bool)`

GetDefaultFacetsOk returns a tuple with the DefaultFacets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFacets

`func (o *SearchRequestOptions) SetDefaultFacets(v []string)`

SetDefaultFacets sets DefaultFacets field to given value.

### HasDefaultFacets

`func (o *SearchRequestOptions) HasDefaultFacets() bool`

HasDefaultFacets returns a boolean if a field has been set.

### GetAuthTokens

`func (o *SearchRequestOptions) GetAuthTokens() []AuthToken`

GetAuthTokens returns the AuthTokens field if non-nil, zero value otherwise.

### GetAuthTokensOk

`func (o *SearchRequestOptions) GetAuthTokensOk() (*[]AuthToken, bool)`

GetAuthTokensOk returns a tuple with the AuthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokens

`func (o *SearchRequestOptions) SetAuthTokens(v []AuthToken)`

SetAuthTokens sets AuthTokens field to given value.

### HasAuthTokens

`func (o *SearchRequestOptions) HasAuthTokens() bool`

HasAuthTokens returns a boolean if a field has been set.

### GetFetchAllDatasourceCounts

`func (o *SearchRequestOptions) GetFetchAllDatasourceCounts() bool`

GetFetchAllDatasourceCounts returns the FetchAllDatasourceCounts field if non-nil, zero value otherwise.

### GetFetchAllDatasourceCountsOk

`func (o *SearchRequestOptions) GetFetchAllDatasourceCountsOk() (*bool, bool)`

GetFetchAllDatasourceCountsOk returns a tuple with the FetchAllDatasourceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchAllDatasourceCounts

`func (o *SearchRequestOptions) SetFetchAllDatasourceCounts(v bool)`

SetFetchAllDatasourceCounts sets FetchAllDatasourceCounts field to given value.

### HasFetchAllDatasourceCounts

`func (o *SearchRequestOptions) HasFetchAllDatasourceCounts() bool`

HasFetchAllDatasourceCounts returns a boolean if a field has been set.

### GetResponseHints

`func (o *SearchRequestOptions) GetResponseHints() []string`

GetResponseHints returns the ResponseHints field if non-nil, zero value otherwise.

### GetResponseHintsOk

`func (o *SearchRequestOptions) GetResponseHintsOk() (*[]string, bool)`

GetResponseHintsOk returns a tuple with the ResponseHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHints

`func (o *SearchRequestOptions) SetResponseHints(v []string)`

SetResponseHints sets ResponseHints field to given value.

### HasResponseHints

`func (o *SearchRequestOptions) HasResponseHints() bool`

HasResponseHints returns a boolean if a field has been set.

### GetTimezoneOffset

`func (o *SearchRequestOptions) GetTimezoneOffset() int32`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *SearchRequestOptions) GetTimezoneOffsetOk() (*int32, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *SearchRequestOptions) SetTimezoneOffset(v int32)`

SetTimezoneOffset sets TimezoneOffset field to given value.

### HasTimezoneOffset

`func (o *SearchRequestOptions) HasTimezoneOffset() bool`

HasTimezoneOffset returns a boolean if a field has been set.

### GetForceNegation

`func (o *SearchRequestOptions) GetForceNegation() bool`

GetForceNegation returns the ForceNegation field if non-nil, zero value otherwise.

### GetForceNegationOk

`func (o *SearchRequestOptions) GetForceNegationOk() (*bool, bool)`

GetForceNegationOk returns a tuple with the ForceNegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNegation

`func (o *SearchRequestOptions) SetForceNegation(v bool)`

SetForceNegation sets ForceNegation field to given value.

### HasForceNegation

`func (o *SearchRequestOptions) HasForceNegation() bool`

HasForceNegation returns a boolean if a field has been set.

### GetDisableSpellcheck

`func (o *SearchRequestOptions) GetDisableSpellcheck() bool`

GetDisableSpellcheck returns the DisableSpellcheck field if non-nil, zero value otherwise.

### GetDisableSpellcheckOk

`func (o *SearchRequestOptions) GetDisableSpellcheckOk() (*bool, bool)`

GetDisableSpellcheckOk returns a tuple with the DisableSpellcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSpellcheck

`func (o *SearchRequestOptions) SetDisableSpellcheck(v bool)`

SetDisableSpellcheck sets DisableSpellcheck field to given value.

### HasDisableSpellcheck

`func (o *SearchRequestOptions) HasDisableSpellcheck() bool`

HasDisableSpellcheck returns a boolean if a field has been set.

### GetDisableQueryAutocorrect

`func (o *SearchRequestOptions) GetDisableQueryAutocorrect() bool`

GetDisableQueryAutocorrect returns the DisableQueryAutocorrect field if non-nil, zero value otherwise.

### GetDisableQueryAutocorrectOk

`func (o *SearchRequestOptions) GetDisableQueryAutocorrectOk() (*bool, bool)`

GetDisableQueryAutocorrectOk returns a tuple with the DisableQueryAutocorrect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableQueryAutocorrect

`func (o *SearchRequestOptions) SetDisableQueryAutocorrect(v bool)`

SetDisableQueryAutocorrect sets DisableQueryAutocorrect field to given value.

### HasDisableQueryAutocorrect

`func (o *SearchRequestOptions) HasDisableQueryAutocorrect() bool`

HasDisableQueryAutocorrect returns a boolean if a field has been set.

### GetReturnLlmContentOverSnippets

`func (o *SearchRequestOptions) GetReturnLlmContentOverSnippets() bool`

GetReturnLlmContentOverSnippets returns the ReturnLlmContentOverSnippets field if non-nil, zero value otherwise.

### GetReturnLlmContentOverSnippetsOk

`func (o *SearchRequestOptions) GetReturnLlmContentOverSnippetsOk() (*bool, bool)`

GetReturnLlmContentOverSnippetsOk returns a tuple with the ReturnLlmContentOverSnippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLlmContentOverSnippets

`func (o *SearchRequestOptions) SetReturnLlmContentOverSnippets(v bool)`

SetReturnLlmContentOverSnippets sets ReturnLlmContentOverSnippets field to given value.

### HasReturnLlmContentOverSnippets

`func (o *SearchRequestOptions) HasReturnLlmContentOverSnippets() bool`

HasReturnLlmContentOverSnippets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


