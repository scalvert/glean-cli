# ResultTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the tab. Can be passed in a search request to get results for that tab. | [optional] 
**Count** | Pointer to **int32** | The number of results in this tab for the current query. | [optional] 
**Datasource** | Pointer to **string** | The datasource associated with the tab, if any. | [optional] 
**DatasourceInstance** | Pointer to **string** | The datasource instance associated with the tab, if any. | [optional] 

## Methods

### NewResultTab

`func NewResultTab() *ResultTab`

NewResultTab instantiates a new ResultTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultTabWithDefaults

`func NewResultTabWithDefaults() *ResultTab`

NewResultTabWithDefaults instantiates a new ResultTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultTab) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultTab) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultTab) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResultTab) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCount

`func (o *ResultTab) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResultTab) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResultTab) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResultTab) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDatasource

`func (o *ResultTab) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *ResultTab) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *ResultTab) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *ResultTab) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetDatasourceInstance

`func (o *ResultTab) GetDatasourceInstance() string`

GetDatasourceInstance returns the DatasourceInstance field if non-nil, zero value otherwise.

### GetDatasourceInstanceOk

`func (o *ResultTab) GetDatasourceInstanceOk() (*string, bool)`

GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstance

`func (o *ResultTab) SetDatasourceInstance(v string)`

SetDatasourceInstance sets DatasourceInstance field to given value.

### HasDatasourceInstance

`func (o *ResultTab) HasDatasourceInstance() bool`

HasDatasourceInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


