# GetCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the Collection to be retrieved. | 
**WithItems** | Pointer to **bool** | Whether or not to include the Collection Items in this Collection. Only request if absolutely required, as this is expensive. | [optional] 
**WithHierarchy** | Pointer to **bool** | Whether or not to include the top level Collection in this Collection&#39;s hierarchy. | [optional] 
**AllowedDatasource** | Pointer to **string** | The datasource allowed in the Collection returned. | [optional] 

## Methods

### NewGetCollectionRequest

`func NewGetCollectionRequest(id int32, ) *GetCollectionRequest`

NewGetCollectionRequest instantiates a new GetCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionRequestWithDefaults

`func NewGetCollectionRequestWithDefaults() *GetCollectionRequest`

NewGetCollectionRequestWithDefaults instantiates a new GetCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCollectionRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCollectionRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCollectionRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetWithItems

`func (o *GetCollectionRequest) GetWithItems() bool`

GetWithItems returns the WithItems field if non-nil, zero value otherwise.

### GetWithItemsOk

`func (o *GetCollectionRequest) GetWithItemsOk() (*bool, bool)`

GetWithItemsOk returns a tuple with the WithItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithItems

`func (o *GetCollectionRequest) SetWithItems(v bool)`

SetWithItems sets WithItems field to given value.

### HasWithItems

`func (o *GetCollectionRequest) HasWithItems() bool`

HasWithItems returns a boolean if a field has been set.

### GetWithHierarchy

`func (o *GetCollectionRequest) GetWithHierarchy() bool`

GetWithHierarchy returns the WithHierarchy field if non-nil, zero value otherwise.

### GetWithHierarchyOk

`func (o *GetCollectionRequest) GetWithHierarchyOk() (*bool, bool)`

GetWithHierarchyOk returns a tuple with the WithHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithHierarchy

`func (o *GetCollectionRequest) SetWithHierarchy(v bool)`

SetWithHierarchy sets WithHierarchy field to given value.

### HasWithHierarchy

`func (o *GetCollectionRequest) HasWithHierarchy() bool`

HasWithHierarchy returns a boolean if a field has been set.

### GetAllowedDatasource

`func (o *GetCollectionRequest) GetAllowedDatasource() string`

GetAllowedDatasource returns the AllowedDatasource field if non-nil, zero value otherwise.

### GetAllowedDatasourceOk

`func (o *GetCollectionRequest) GetAllowedDatasourceOk() (*string, bool)`

GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDatasource

`func (o *GetCollectionRequest) SetAllowedDatasource(v string)`

SetAllowedDatasource sets AllowedDatasource field to given value.

### HasAllowedDatasource

`func (o *GetCollectionRequest) HasAllowedDatasource() bool`

HasAllowedDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


