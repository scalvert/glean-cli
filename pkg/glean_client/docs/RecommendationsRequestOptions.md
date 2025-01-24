# RecommendationsRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasourceFilter** | Pointer to **string** | Filter results to a single datasource name (e.g. gmail, slack). All results are returned if missing. | [optional] 
**DatasourcesFilter** | Pointer to **[]string** | Filter results to only those relevant to one or more datasources (e.g. jira, gdrive). All results are returned if missing. | [optional] 
**Context** | Pointer to [**Document**](Document.md) |  | [optional] 
**ResultProminence** | Pointer to [**[]SearchResultProminenceEnum**](SearchResultProminenceEnum.md) | The types of prominence wanted in results returned. Default is any type. | [optional] 

## Methods

### NewRecommendationsRequestOptions

`func NewRecommendationsRequestOptions() *RecommendationsRequestOptions`

NewRecommendationsRequestOptions instantiates a new RecommendationsRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationsRequestOptionsWithDefaults

`func NewRecommendationsRequestOptionsWithDefaults() *RecommendationsRequestOptions`

NewRecommendationsRequestOptionsWithDefaults instantiates a new RecommendationsRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasourceFilter

`func (o *RecommendationsRequestOptions) GetDatasourceFilter() string`

GetDatasourceFilter returns the DatasourceFilter field if non-nil, zero value otherwise.

### GetDatasourceFilterOk

`func (o *RecommendationsRequestOptions) GetDatasourceFilterOk() (*string, bool)`

GetDatasourceFilterOk returns a tuple with the DatasourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceFilter

`func (o *RecommendationsRequestOptions) SetDatasourceFilter(v string)`

SetDatasourceFilter sets DatasourceFilter field to given value.

### HasDatasourceFilter

`func (o *RecommendationsRequestOptions) HasDatasourceFilter() bool`

HasDatasourceFilter returns a boolean if a field has been set.

### GetDatasourcesFilter

`func (o *RecommendationsRequestOptions) GetDatasourcesFilter() []string`

GetDatasourcesFilter returns the DatasourcesFilter field if non-nil, zero value otherwise.

### GetDatasourcesFilterOk

`func (o *RecommendationsRequestOptions) GetDatasourcesFilterOk() (*[]string, bool)`

GetDatasourcesFilterOk returns a tuple with the DatasourcesFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourcesFilter

`func (o *RecommendationsRequestOptions) SetDatasourcesFilter(v []string)`

SetDatasourcesFilter sets DatasourcesFilter field to given value.

### HasDatasourcesFilter

`func (o *RecommendationsRequestOptions) HasDatasourcesFilter() bool`

HasDatasourcesFilter returns a boolean if a field has been set.

### GetContext

`func (o *RecommendationsRequestOptions) GetContext() Document`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RecommendationsRequestOptions) GetContextOk() (*Document, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RecommendationsRequestOptions) SetContext(v Document)`

SetContext sets Context field to given value.

### HasContext

`func (o *RecommendationsRequestOptions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetResultProminence

`func (o *RecommendationsRequestOptions) GetResultProminence() []SearchResultProminenceEnum`

GetResultProminence returns the ResultProminence field if non-nil, zero value otherwise.

### GetResultProminenceOk

`func (o *RecommendationsRequestOptions) GetResultProminenceOk() (*[]SearchResultProminenceEnum, bool)`

GetResultProminenceOk returns a tuple with the ResultProminence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultProminence

`func (o *RecommendationsRequestOptions) SetResultProminence(v []SearchResultProminenceEnum)`

SetResultProminence sets ResultProminence field to given value.

### HasResultProminence

`func (o *RecommendationsRequestOptions) HasResultProminence() bool`

HasResultProminence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


