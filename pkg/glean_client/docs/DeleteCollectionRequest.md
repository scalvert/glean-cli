# DeleteCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]int32** | The IDs of the Collections to delete. | 
**AllowedDatasource** | Pointer to **string** | The datasource allowed in the Collection to be deleted. | [optional] 

## Methods

### NewDeleteCollectionRequest

`func NewDeleteCollectionRequest(ids []int32, ) *DeleteCollectionRequest`

NewDeleteCollectionRequest instantiates a new DeleteCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCollectionRequestWithDefaults

`func NewDeleteCollectionRequestWithDefaults() *DeleteCollectionRequest`

NewDeleteCollectionRequestWithDefaults instantiates a new DeleteCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *DeleteCollectionRequest) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DeleteCollectionRequest) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DeleteCollectionRequest) SetIds(v []int32)`

SetIds sets Ids field to given value.


### GetAllowedDatasource

`func (o *DeleteCollectionRequest) GetAllowedDatasource() string`

GetAllowedDatasource returns the AllowedDatasource field if non-nil, zero value otherwise.

### GetAllowedDatasourceOk

`func (o *DeleteCollectionRequest) GetAllowedDatasourceOk() (*string, bool)`

GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDatasource

`func (o *DeleteCollectionRequest) SetAllowedDatasource(v string)`

SetAllowedDatasource sets AllowedDatasource field to given value.

### HasAllowedDatasource

`func (o *DeleteCollectionRequest) HasAllowedDatasource() bool`

HasAllowedDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


