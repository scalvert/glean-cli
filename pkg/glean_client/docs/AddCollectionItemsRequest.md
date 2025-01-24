# AddCollectionItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **float32** | The ID of the Collection to add items to. | 
**AddedCollectionItemDescriptors** | Pointer to [**[]CollectionItemDescriptor**](CollectionItemDescriptor.md) | The CollectionItemDescriptors of the items being added. | [optional] 

## Methods

### NewAddCollectionItemsRequest

`func NewAddCollectionItemsRequest(collectionId float32, ) *AddCollectionItemsRequest`

NewAddCollectionItemsRequest instantiates a new AddCollectionItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCollectionItemsRequestWithDefaults

`func NewAddCollectionItemsRequestWithDefaults() *AddCollectionItemsRequest`

NewAddCollectionItemsRequestWithDefaults instantiates a new AddCollectionItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *AddCollectionItemsRequest) GetCollectionId() float32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *AddCollectionItemsRequest) GetCollectionIdOk() (*float32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *AddCollectionItemsRequest) SetCollectionId(v float32)`

SetCollectionId sets CollectionId field to given value.


### GetAddedCollectionItemDescriptors

`func (o *AddCollectionItemsRequest) GetAddedCollectionItemDescriptors() []CollectionItemDescriptor`

GetAddedCollectionItemDescriptors returns the AddedCollectionItemDescriptors field if non-nil, zero value otherwise.

### GetAddedCollectionItemDescriptorsOk

`func (o *AddCollectionItemsRequest) GetAddedCollectionItemDescriptorsOk() (*[]CollectionItemDescriptor, bool)`

GetAddedCollectionItemDescriptorsOk returns a tuple with the AddedCollectionItemDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCollectionItemDescriptors

`func (o *AddCollectionItemsRequest) SetAddedCollectionItemDescriptors(v []CollectionItemDescriptor)`

SetAddedCollectionItemDescriptors sets AddedCollectionItemDescriptors field to given value.

### HasAddedCollectionItemDescriptors

`func (o *AddCollectionItemsRequest) HasAddedCollectionItemDescriptors() bool`

HasAddedCollectionItemDescriptors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


