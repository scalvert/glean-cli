# QuerySuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MissingTerm** | Pointer to **string** | A query term missing from the original query on which this suggestion is based | [optional] 
**Query** | **string** | The query being suggested (e.g. enforcing the missing term from the original query) | 
**Label** | Pointer to **string** | A user-facing description to display for the suggestion | [optional] 
**Datasource** | Pointer to **string** | The datasource associated with the suggestion | [optional] 
**RequestOptions** | Pointer to [**SearchRequestOptions**](SearchRequestOptions.md) |  | [optional] 
**Ranges** | Pointer to [**[]TextRange**](TextRange.md) | The bolded ranges within the query of the QuerySuggestion. | [optional] 

## Methods

### NewQuerySuggestion

`func NewQuerySuggestion(query string, ) *QuerySuggestion`

NewQuerySuggestion instantiates a new QuerySuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySuggestionWithDefaults

`func NewQuerySuggestionWithDefaults() *QuerySuggestion`

NewQuerySuggestionWithDefaults instantiates a new QuerySuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMissingTerm

`func (o *QuerySuggestion) GetMissingTerm() string`

GetMissingTerm returns the MissingTerm field if non-nil, zero value otherwise.

### GetMissingTermOk

`func (o *QuerySuggestion) GetMissingTermOk() (*string, bool)`

GetMissingTermOk returns a tuple with the MissingTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingTerm

`func (o *QuerySuggestion) SetMissingTerm(v string)`

SetMissingTerm sets MissingTerm field to given value.

### HasMissingTerm

`func (o *QuerySuggestion) HasMissingTerm() bool`

HasMissingTerm returns a boolean if a field has been set.

### GetQuery

`func (o *QuerySuggestion) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QuerySuggestion) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QuerySuggestion) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetLabel

`func (o *QuerySuggestion) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *QuerySuggestion) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *QuerySuggestion) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *QuerySuggestion) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDatasource

`func (o *QuerySuggestion) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *QuerySuggestion) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *QuerySuggestion) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *QuerySuggestion) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetRequestOptions

`func (o *QuerySuggestion) GetRequestOptions() SearchRequestOptions`

GetRequestOptions returns the RequestOptions field if non-nil, zero value otherwise.

### GetRequestOptionsOk

`func (o *QuerySuggestion) GetRequestOptionsOk() (*SearchRequestOptions, bool)`

GetRequestOptionsOk returns a tuple with the RequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOptions

`func (o *QuerySuggestion) SetRequestOptions(v SearchRequestOptions)`

SetRequestOptions sets RequestOptions field to given value.

### HasRequestOptions

`func (o *QuerySuggestion) HasRequestOptions() bool`

HasRequestOptions returns a boolean if a field has been set.

### GetRanges

`func (o *QuerySuggestion) GetRanges() []TextRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *QuerySuggestion) GetRangesOk() (*[]TextRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *QuerySuggestion) SetRanges(v []TextRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *QuerySuggestion) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


