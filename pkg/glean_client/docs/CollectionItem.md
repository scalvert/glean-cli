# CollectionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The optional name of the Collection item. | [optional] 
**Description** | Pointer to **string** | A helpful description of why this CollectionItem is in the Collection that it&#39;s in. | [optional] 
**Icon** | Pointer to **string** | The emoji icon for this CollectionItem. Only used for Text type items. | [optional] 
**CollectionId** | **int32** | The Collection ID of the Collection that this CollectionItem belongs in. | 
**DocumentId** | Pointer to **string** | If this CollectionItem is indexed, the Glean Document ID of that document. | [optional] 
**Url** | Pointer to **string** | The URL of this CollectionItem. | [optional] 
**ItemId** | Pointer to **string** | Unique identifier for the item within the Collection it belongs to. | [optional] 
**CreatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Unix timestamp for when the item was first added (in seconds since epoch UTC). | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Shortcut** | Pointer to [**Shortcut**](Shortcut.md) |  | [optional] 
**Collection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**ItemType** | **string** |  | 

## Methods

### NewCollectionItem

`func NewCollectionItem(collectionId int32, itemType string, ) *CollectionItem`

NewCollectionItem instantiates a new CollectionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionItemWithDefaults

`func NewCollectionItemWithDefaults() *CollectionItem`

NewCollectionItemWithDefaults instantiates a new CollectionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CollectionItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CollectionItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CollectionItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CollectionItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CollectionItem) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CollectionItem) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetCollectionId

`func (o *CollectionItem) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CollectionItem) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CollectionItem) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.


### GetDocumentId

`func (o *CollectionItem) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *CollectionItem) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *CollectionItem) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *CollectionItem) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetUrl

`func (o *CollectionItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CollectionItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CollectionItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CollectionItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetItemId

`func (o *CollectionItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CollectionItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CollectionItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *CollectionItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CollectionItem) GetCreatedBy() Person`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CollectionItem) GetCreatedByOk() (*Person, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CollectionItem) SetCreatedBy(v Person)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CollectionItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CollectionItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CollectionItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDocument

`func (o *CollectionItem) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *CollectionItem) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *CollectionItem) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *CollectionItem) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetShortcut

`func (o *CollectionItem) GetShortcut() Shortcut`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *CollectionItem) GetShortcutOk() (*Shortcut, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *CollectionItem) SetShortcut(v Shortcut)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *CollectionItem) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetCollection

`func (o *CollectionItem) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CollectionItem) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CollectionItem) SetCollection(v Collection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *CollectionItem) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetItemType

`func (o *CollectionItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CollectionItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CollectionItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


