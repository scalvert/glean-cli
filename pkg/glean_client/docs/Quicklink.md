# Quicklink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Full action name. Used in autocomplete. | [optional] 
**ShortName** | Pointer to **string** | Shortened name. Used in app cards. | [optional] 
**Url** | Pointer to **string** | The URL of the action. | [optional] 
**IconConfig** | Pointer to [**IconConfig**](IconConfig.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier of this quicklink | [optional] 
**Scopes** | Pointer to **[]string** | The scopes for which this quicklink is applicable | [optional] 

## Methods

### NewQuicklink

`func NewQuicklink() *Quicklink`

NewQuicklink instantiates a new Quicklink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuicklinkWithDefaults

`func NewQuicklinkWithDefaults() *Quicklink`

NewQuicklinkWithDefaults instantiates a new Quicklink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Quicklink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Quicklink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Quicklink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Quicklink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortName

`func (o *Quicklink) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *Quicklink) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *Quicklink) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *Quicklink) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetUrl

`func (o *Quicklink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Quicklink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Quicklink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Quicklink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIconConfig

`func (o *Quicklink) GetIconConfig() IconConfig`

GetIconConfig returns the IconConfig field if non-nil, zero value otherwise.

### GetIconConfigOk

`func (o *Quicklink) GetIconConfigOk() (*IconConfig, bool)`

GetIconConfigOk returns a tuple with the IconConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconConfig

`func (o *Quicklink) SetIconConfig(v IconConfig)`

SetIconConfig sets IconConfig field to given value.

### HasIconConfig

`func (o *Quicklink) HasIconConfig() bool`

HasIconConfig returns a boolean if a field has been set.

### GetId

`func (o *Quicklink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Quicklink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Quicklink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Quicklink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScopes

`func (o *Quicklink) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Quicklink) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Quicklink) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Quicklink) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


