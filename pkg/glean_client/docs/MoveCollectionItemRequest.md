# MoveCollectionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **int32** | The ID of the Collection to move items in. | 
**ItemId** | **string** | The item ID of the item being moved. | 
**NewNextItemId** | Pointer to **string** | The (optional) item ID of the item that is the new next of itemId, or empty if this is now the last item. This item does not move, it&#39;s used as a reference position to put the itemId in the right position. | [optional] 

## Methods

### NewMoveCollectionItemRequest

`func NewMoveCollectionItemRequest(collectionId int32, itemId string, ) *MoveCollectionItemRequest`

NewMoveCollectionItemRequest instantiates a new MoveCollectionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveCollectionItemRequestWithDefaults

`func NewMoveCollectionItemRequestWithDefaults() *MoveCollectionItemRequest`

NewMoveCollectionItemRequestWithDefaults instantiates a new MoveCollectionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *MoveCollectionItemRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *MoveCollectionItemRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *MoveCollectionItemRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.


### GetItemId

`func (o *MoveCollectionItemRequest) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *MoveCollectionItemRequest) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *MoveCollectionItemRequest) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetNewNextItemId

`func (o *MoveCollectionItemRequest) GetNewNextItemId() string`

GetNewNextItemId returns the NewNextItemId field if non-nil, zero value otherwise.

### GetNewNextItemIdOk

`func (o *MoveCollectionItemRequest) GetNewNextItemIdOk() (*string, bool)`

GetNewNextItemIdOk returns a tuple with the NewNextItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNextItemId

`func (o *MoveCollectionItemRequest) SetNewNextItemId(v string)`

SetNewNextItemId sets NewNextItemId field to given value.

### HasNewNextItemId

`func (o *MoveCollectionItemRequest) HasNewNextItemId() bool`

HasNewNextItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


