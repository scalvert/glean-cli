# Badge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | An auto generated unique identifier. | [optional] 
**DisplayName** | Pointer to **string** | The badge name displayed to users | [optional] 
**IconConfig** | Pointer to [**IconConfig**](IconConfig.md) |  | [optional] 
**Pinned** | Pointer to **bool** | The badge should be shown on the PersonAttribution | [optional] 

## Methods

### NewBadge

`func NewBadge() *Badge`

NewBadge instantiates a new Badge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadgeWithDefaults

`func NewBadgeWithDefaults() *Badge`

NewBadgeWithDefaults instantiates a new Badge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Badge) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Badge) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Badge) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Badge) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDisplayName

`func (o *Badge) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Badge) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Badge) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Badge) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIconConfig

`func (o *Badge) GetIconConfig() IconConfig`

GetIconConfig returns the IconConfig field if non-nil, zero value otherwise.

### GetIconConfigOk

`func (o *Badge) GetIconConfigOk() (*IconConfig, bool)`

GetIconConfigOk returns a tuple with the IconConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconConfig

`func (o *Badge) SetIconConfig(v IconConfig)`

SetIconConfig sets IconConfig field to given value.

### HasIconConfig

`func (o *Badge) HasIconConfig() bool`

HasIconConfig returns a boolean if a field has been set.

### GetPinned

`func (o *Badge) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *Badge) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *Badge) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *Badge) HasPinned() bool`

HasPinned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


