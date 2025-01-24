# WriteAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolName** | Pointer to **string** | The name of the tool. | [optional] 
**ToolConfig** | Pointer to [**ToolConfig**](ToolConfig.md) |  | [optional] 
**RedirectUrl** | Pointer to **string** | If a &#x60;REDIRECT&#x60; action, the URL to visit to execute the action. | [optional] 
**Parameters** | Pointer to [**map[string]WriteActionParameter**](WriteActionParameter.md) | The parameters to be passed to the redirect URL for actions. | [optional] 

## Methods

### NewWriteAction

`func NewWriteAction() *WriteAction`

NewWriteAction instantiates a new WriteAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteActionWithDefaults

`func NewWriteActionWithDefaults() *WriteAction`

NewWriteActionWithDefaults instantiates a new WriteAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolName

`func (o *WriteAction) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *WriteAction) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *WriteAction) SetToolName(v string)`

SetToolName sets ToolName field to given value.

### HasToolName

`func (o *WriteAction) HasToolName() bool`

HasToolName returns a boolean if a field has been set.

### GetToolConfig

`func (o *WriteAction) GetToolConfig() ToolConfig`

GetToolConfig returns the ToolConfig field if non-nil, zero value otherwise.

### GetToolConfigOk

`func (o *WriteAction) GetToolConfigOk() (*ToolConfig, bool)`

GetToolConfigOk returns a tuple with the ToolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfig

`func (o *WriteAction) SetToolConfig(v ToolConfig)`

SetToolConfig sets ToolConfig field to given value.

### HasToolConfig

`func (o *WriteAction) HasToolConfig() bool`

HasToolConfig returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *WriteAction) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *WriteAction) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *WriteAction) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *WriteAction) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetParameters

`func (o *WriteAction) GetParameters() map[string]WriteActionParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WriteAction) GetParametersOk() (*map[string]WriteActionParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WriteAction) SetParameters(v map[string]WriteActionParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WriteAction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


