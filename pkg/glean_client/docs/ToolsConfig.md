# ToolsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableTools** | Pointer to [**[]ToolMetadata**](ToolMetadata.md) | List of tools available to the user. | [optional] 

## Methods

### NewToolsConfig

`func NewToolsConfig() *ToolsConfig`

NewToolsConfig instantiates a new ToolsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolsConfigWithDefaults

`func NewToolsConfigWithDefaults() *ToolsConfig`

NewToolsConfigWithDefaults instantiates a new ToolsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableTools

`func (o *ToolsConfig) GetAvailableTools() []ToolMetadata`

GetAvailableTools returns the AvailableTools field if non-nil, zero value otherwise.

### GetAvailableToolsOk

`func (o *ToolsConfig) GetAvailableToolsOk() (*[]ToolMetadata, bool)`

GetAvailableToolsOk returns a tuple with the AvailableTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTools

`func (o *ToolsConfig) SetAvailableTools(v []ToolMetadata)`

SetAvailableTools sets AvailableTools field to given value.

### HasAvailableTools

`func (o *ToolsConfig) HasAvailableTools() bool`

HasAvailableTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


