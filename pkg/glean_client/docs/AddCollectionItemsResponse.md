# AddCollectionItemsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**Error** | Pointer to [**AddCollectionItemsError**](AddCollectionItemsError.md) |  | [optional] 

## Methods

### NewAddCollectionItemsResponse

`func NewAddCollectionItemsResponse() *AddCollectionItemsResponse`

NewAddCollectionItemsResponse instantiates a new AddCollectionItemsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCollectionItemsResponseWithDefaults

`func NewAddCollectionItemsResponseWithDefaults() *AddCollectionItemsResponse`

NewAddCollectionItemsResponseWithDefaults instantiates a new AddCollectionItemsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *AddCollectionItemsResponse) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *AddCollectionItemsResponse) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *AddCollectionItemsResponse) SetCollection(v Collection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *AddCollectionItemsResponse) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetError

`func (o *AddCollectionItemsResponse) GetError() AddCollectionItemsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AddCollectionItemsResponse) GetErrorOk() (*AddCollectionItemsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AddCollectionItemsResponse) SetError(v AddCollectionItemsError)`

SetError sets Error field to given value.

### HasError

`func (o *AddCollectionItemsResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


