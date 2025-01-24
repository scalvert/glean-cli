# StructuredLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The display name for the link | [optional] 
**Url** | Pointer to **string** | The URL for the link. | [optional] 
**IconConfig** | Pointer to [**IconConfig**](IconConfig.md) |  | [optional] 

## Methods

### NewStructuredLink

`func NewStructuredLink() *StructuredLink`

NewStructuredLink instantiates a new StructuredLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuredLinkWithDefaults

`func NewStructuredLinkWithDefaults() *StructuredLink`

NewStructuredLinkWithDefaults instantiates a new StructuredLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StructuredLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StructuredLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StructuredLink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StructuredLink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *StructuredLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StructuredLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StructuredLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StructuredLink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIconConfig

`func (o *StructuredLink) GetIconConfig() IconConfig`

GetIconConfig returns the IconConfig field if non-nil, zero value otherwise.

### GetIconConfigOk

`func (o *StructuredLink) GetIconConfigOk() (*IconConfig, bool)`

GetIconConfigOk returns a tuple with the IconConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconConfig

`func (o *StructuredLink) SetIconConfig(v IconConfig)`

SetIconConfig sets IconConfig field to given value.

### HasIconConfig

`func (o *StructuredLink) HasIconConfig() bool`

HasIconConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


