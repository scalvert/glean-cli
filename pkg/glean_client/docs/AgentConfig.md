# AgentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Name of the agent. DEFAULT - Integrates with your company&#39;s knowledge. GPT - Communicates directly with the LLM. | [optional] 
**Mode** | Pointer to **string** | Top level modes to run GleanChat in. DEFAULT - Used if no mode supplied. QUICK - Trades accuracy and precision for speed. | [optional] 

## Methods

### NewAgentConfig

`func NewAgentConfig() *AgentConfig`

NewAgentConfig instantiates a new AgentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentConfigWithDefaults

`func NewAgentConfigWithDefaults() *AgentConfig`

NewAgentConfigWithDefaults instantiates a new AgentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AgentConfig) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentConfig) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentConfig) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentConfig) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetMode

`func (o *AgentConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AgentConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AgentConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AgentConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


