# Themes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Light** | Pointer to **map[string]string** | A map of {string, string} pairs where the key represents a known theme option and the value is the CSS color to apply. Supported options are documented in https://dev.glean.com/meta/browser_api/interfaces/Theme.html. | [optional] 
**Dark** | Pointer to **map[string]string** | A map of {string, string} pairs where the key represents a known theme option and the value is the CSS color to apply. Supported options are documented in https://dev.glean.com/meta/browser_api/interfaces/Theme.html. | [optional] 

## Methods

### NewThemes

`func NewThemes() *Themes`

NewThemes instantiates a new Themes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThemesWithDefaults

`func NewThemesWithDefaults() *Themes`

NewThemesWithDefaults instantiates a new Themes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLight

`func (o *Themes) GetLight() map[string]string`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *Themes) GetLightOk() (*map[string]string, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *Themes) SetLight(v map[string]string)`

SetLight sets Light field to given value.

### HasLight

`func (o *Themes) HasLight() bool`

HasLight returns a boolean if a field has been set.

### GetDark

`func (o *Themes) GetDark() map[string]string`

GetDark returns the Dark field if non-nil, zero value otherwise.

### GetDarkOk

`func (o *Themes) GetDarkOk() (*map[string]string, bool)`

GetDarkOk returns a tuple with the Dark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDark

`func (o *Themes) SetDark(v map[string]string)`

SetDark sets Dark field to given value.

### HasDark

`func (o *Themes) HasDark() bool`

HasDark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


