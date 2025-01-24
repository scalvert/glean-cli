# EditCollectionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **int32** | The ID of the Collection to edit CollectionItems in. | 
**ItemId** | **string** | The ID of the CollectionItem to edit. | 
**Name** | Pointer to **string** | The optional name of the Collection item. | [optional] 
**Description** | Pointer to **string** | A helpful description of why this CollectionItem is in the Collection that it&#39;s in. | [optional] 
**Icon** | Pointer to **string** | The emoji icon for this CollectionItem. Only used for Text type items. | [optional] 

## Methods

### NewEditCollectionItemRequest

`func NewEditCollectionItemRequest(collectionId int32, itemId string, ) *EditCollectionItemRequest`

NewEditCollectionItemRequest instantiates a new EditCollectionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditCollectionItemRequestWithDefaults

`func NewEditCollectionItemRequestWithDefaults() *EditCollectionItemRequest`

NewEditCollectionItemRequestWithDefaults instantiates a new EditCollectionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *EditCollectionItemRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *EditCollectionItemRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *EditCollectionItemRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.


### GetItemId

`func (o *EditCollectionItemRequest) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *EditCollectionItemRequest) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *EditCollectionItemRequest) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetName

`func (o *EditCollectionItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditCollectionItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditCollectionItemRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditCollectionItemRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EditCollectionItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditCollectionItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditCollectionItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditCollectionItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *EditCollectionItemRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EditCollectionItemRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EditCollectionItemRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EditCollectionItemRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


