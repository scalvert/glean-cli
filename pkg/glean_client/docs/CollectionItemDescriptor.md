# CollectionItemDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the item being added. | [optional] 
**DocumentId** | Pointer to **string** | The Glean Document ID of the item being added if it&#39;s an indexed document. | [optional] 
**NewNextItemId** | Pointer to **string** | The (optional) ItemId of the next CollectionItem in sequence. If omitted, will be added to the end of the Collection | [optional] 
**ItemType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The optional name of the Collection item. | [optional] 
**Description** | Pointer to **string** | A helpful description of why this CollectionItem is in the Collection that it&#39;s in. | [optional] 
**Icon** | Pointer to **string** | The emoji icon for this CollectionItem. Only used for Text type items. | [optional] 

## Methods

### NewCollectionItemDescriptor

`func NewCollectionItemDescriptor() *CollectionItemDescriptor`

NewCollectionItemDescriptor instantiates a new CollectionItemDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionItemDescriptorWithDefaults

`func NewCollectionItemDescriptorWithDefaults() *CollectionItemDescriptor`

NewCollectionItemDescriptorWithDefaults instantiates a new CollectionItemDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CollectionItemDescriptor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CollectionItemDescriptor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CollectionItemDescriptor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CollectionItemDescriptor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDocumentId

`func (o *CollectionItemDescriptor) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *CollectionItemDescriptor) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *CollectionItemDescriptor) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *CollectionItemDescriptor) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetNewNextItemId

`func (o *CollectionItemDescriptor) GetNewNextItemId() string`

GetNewNextItemId returns the NewNextItemId field if non-nil, zero value otherwise.

### GetNewNextItemIdOk

`func (o *CollectionItemDescriptor) GetNewNextItemIdOk() (*string, bool)`

GetNewNextItemIdOk returns a tuple with the NewNextItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNextItemId

`func (o *CollectionItemDescriptor) SetNewNextItemId(v string)`

SetNewNextItemId sets NewNextItemId field to given value.

### HasNewNextItemId

`func (o *CollectionItemDescriptor) HasNewNextItemId() bool`

HasNewNextItemId returns a boolean if a field has been set.

### GetItemType

`func (o *CollectionItemDescriptor) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CollectionItemDescriptor) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CollectionItemDescriptor) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *CollectionItemDescriptor) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetName

`func (o *CollectionItemDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionItemDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionItemDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionItemDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CollectionItemDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionItemDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionItemDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionItemDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CollectionItemDescriptor) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CollectionItemDescriptor) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CollectionItemDescriptor) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CollectionItemDescriptor) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


