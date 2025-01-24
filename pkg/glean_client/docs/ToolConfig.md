# ToolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Human understandable name of the tool. Max 50 characters. | [optional] 
**ObjectName** | Pointer to **string** | Name of the generated object. This will be used to indicate to the end user what the generated object contains. | [optional] 
**LogoUrl** | Pointer to **string** | URL used to fetch the logo. | [optional] 
**Type** | Pointer to **string** | Valid only for ACTION tools. Represents the type of action tool REDIRECT - The client renders the URL which contains information for carrying out the action. EXECUTION - Send a request to an external server and execute the action. | [optional] 

## Methods

### NewToolConfig

`func NewToolConfig() *ToolConfig`

NewToolConfig instantiates a new ToolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolConfigWithDefaults

`func NewToolConfigWithDefaults() *ToolConfig`

NewToolConfigWithDefaults instantiates a new ToolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ToolConfig) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ToolConfig) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ToolConfig) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ToolConfig) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetObjectName

`func (o *ToolConfig) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ToolConfig) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ToolConfig) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ToolConfig) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *ToolConfig) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ToolConfig) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ToolConfig) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *ToolConfig) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetType

`func (o *ToolConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


