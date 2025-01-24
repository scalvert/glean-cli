# DeleteCollectionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **float32** | The ID of the Collection to remove an item in. | 
**ItemId** | **string** | The item ID of the CollectionItem to remove from this Collection. | 
**DocumentId** | Pointer to **string** | The (optional) Glean Document ID of the CollectionItem to remove from this Collection if this is an indexed document. | [optional] 

## Methods

### NewDeleteCollectionItemRequest

`func NewDeleteCollectionItemRequest(collectionId float32, itemId string, ) *DeleteCollectionItemRequest`

NewDeleteCollectionItemRequest instantiates a new DeleteCollectionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCollectionItemRequestWithDefaults

`func NewDeleteCollectionItemRequestWithDefaults() *DeleteCollectionItemRequest`

NewDeleteCollectionItemRequestWithDefaults instantiates a new DeleteCollectionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *DeleteCollectionItemRequest) GetCollectionId() float32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *DeleteCollectionItemRequest) GetCollectionIdOk() (*float32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *DeleteCollectionItemRequest) SetCollectionId(v float32)`

SetCollectionId sets CollectionId field to given value.


### GetItemId

`func (o *DeleteCollectionItemRequest) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DeleteCollectionItemRequest) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DeleteCollectionItemRequest) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetDocumentId

`func (o *DeleteCollectionItemRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DeleteCollectionItemRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DeleteCollectionItemRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *DeleteCollectionItemRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


