# CollectionItemMutableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The optional name of the Collection item. | [optional] 
**Description** | Pointer to **string** | A helpful description of why this CollectionItem is in the Collection that it&#39;s in. | [optional] 
**Icon** | Pointer to **string** | The emoji icon for this CollectionItem. Only used for Text type items. | [optional] 

## Methods

### NewCollectionItemMutableProperties

`func NewCollectionItemMutableProperties() *CollectionItemMutableProperties`

NewCollectionItemMutableProperties instantiates a new CollectionItemMutableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionItemMutablePropertiesWithDefaults

`func NewCollectionItemMutablePropertiesWithDefaults() *CollectionItemMutableProperties`

NewCollectionItemMutablePropertiesWithDefaults instantiates a new CollectionItemMutableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CollectionItemMutableProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionItemMutableProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionItemMutableProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionItemMutableProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CollectionItemMutableProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionItemMutableProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionItemMutableProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionItemMutableProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *CollectionItemMutableProperties) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CollectionItemMutableProperties) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CollectionItemMutableProperties) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CollectionItemMutableProperties) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


