# AgentClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentConfig** | Pointer to [**AgentConfig**](AgentConfig.md) |  | [optional] 
**InputCharLimit** | Pointer to **int32** | The character limit of an input to GleanChat under the specified AgentConfig. | [optional] 

## Methods

### NewAgentClientConfig

`func NewAgentClientConfig() *AgentClientConfig`

NewAgentClientConfig instantiates a new AgentClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentClientConfigWithDefaults

`func NewAgentClientConfigWithDefaults() *AgentClientConfig`

NewAgentClientConfigWithDefaults instantiates a new AgentClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentConfig

`func (o *AgentClientConfig) GetAgentConfig() AgentConfig`

GetAgentConfig returns the AgentConfig field if non-nil, zero value otherwise.

### GetAgentConfigOk

`func (o *AgentClientConfig) GetAgentConfigOk() (*AgentConfig, bool)`

GetAgentConfigOk returns a tuple with the AgentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfig

`func (o *AgentClientConfig) SetAgentConfig(v AgentConfig)`

SetAgentConfig sets AgentConfig field to given value.

### HasAgentConfig

`func (o *AgentClientConfig) HasAgentConfig() bool`

HasAgentConfig returns a boolean if a field has been set.

### GetInputCharLimit

`func (o *AgentClientConfig) GetInputCharLimit() int32`

GetInputCharLimit returns the InputCharLimit field if non-nil, zero value otherwise.

### GetInputCharLimitOk

`func (o *AgentClientConfig) GetInputCharLimitOk() (*int32, bool)`

GetInputCharLimitOk returns a tuple with the InputCharLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCharLimit

`func (o *AgentClientConfig) SetInputCharLimit(v int32)`

SetInputCharLimit sets InputCharLimit field to given value.

### HasInputCharLimit

`func (o *AgentClientConfig) HasInputCharLimit() bool`

HasInputCharLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


