# PublicConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThemeKeys** | Pointer to **[]string** | A list of theme keys to include in the response. | [optional] 
**BoolKeys** | Pointer to **[]string** | A list of boolean flag keys to include in the response. | [optional] 

## Methods

### NewPublicConfigRequest

`func NewPublicConfigRequest() *PublicConfigRequest`

NewPublicConfigRequest instantiates a new PublicConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicConfigRequestWithDefaults

`func NewPublicConfigRequestWithDefaults() *PublicConfigRequest`

NewPublicConfigRequestWithDefaults instantiates a new PublicConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThemeKeys

`func (o *PublicConfigRequest) GetThemeKeys() []string`

GetThemeKeys returns the ThemeKeys field if non-nil, zero value otherwise.

### GetThemeKeysOk

`func (o *PublicConfigRequest) GetThemeKeysOk() (*[]string, bool)`

GetThemeKeysOk returns a tuple with the ThemeKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeKeys

`func (o *PublicConfigRequest) SetThemeKeys(v []string)`

SetThemeKeys sets ThemeKeys field to given value.

### HasThemeKeys

`func (o *PublicConfigRequest) HasThemeKeys() bool`

HasThemeKeys returns a boolean if a field has been set.

### GetBoolKeys

`func (o *PublicConfigRequest) GetBoolKeys() []string`

GetBoolKeys returns the BoolKeys field if non-nil, zero value otherwise.

### GetBoolKeysOk

`func (o *PublicConfigRequest) GetBoolKeysOk() (*[]string, bool)`

GetBoolKeysOk returns a tuple with the BoolKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolKeys

`func (o *PublicConfigRequest) SetBoolKeys(v []string)`

SetBoolKeys sets BoolKeys field to given value.

### HasBoolKeys

`func (o *PublicConfigRequest) HasBoolKeys() bool`

HasBoolKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


