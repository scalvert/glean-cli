# ChatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]ChatMessage**](ChatMessage.md) |  | [optional] 
**ChatId** | Pointer to **string** | The id of the associated Chat the messages belong to, if one exists. | [optional] 
**FollowUpPrompts** | Pointer to **[]string** | Follow-up prompts for the user to potentially use | [optional] 
**AgentConfig** | Pointer to [**AgentConfig**](AgentConfig.md) |  | [optional] 
**BackendTimeMillis** | Pointer to **int64** | Time in milliseconds the backend took to respond to the request. | [optional] 
**ChatSessionTrackingToken** | Pointer to **string** | A token that is used to track the session. | [optional] 

## Methods

### NewChatResponse

`func NewChatResponse() *ChatResponse`

NewChatResponse instantiates a new ChatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatResponseWithDefaults

`func NewChatResponseWithDefaults() *ChatResponse`

NewChatResponseWithDefaults instantiates a new ChatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatResponse) GetMessages() []ChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatResponse) GetMessagesOk() (*[]ChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatResponse) SetMessages(v []ChatMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ChatResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetChatId

`func (o *ChatResponse) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ChatResponse) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ChatResponse) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *ChatResponse) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetFollowUpPrompts

`func (o *ChatResponse) GetFollowUpPrompts() []string`

GetFollowUpPrompts returns the FollowUpPrompts field if non-nil, zero value otherwise.

### GetFollowUpPromptsOk

`func (o *ChatResponse) GetFollowUpPromptsOk() (*[]string, bool)`

GetFollowUpPromptsOk returns a tuple with the FollowUpPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpPrompts

`func (o *ChatResponse) SetFollowUpPrompts(v []string)`

SetFollowUpPrompts sets FollowUpPrompts field to given value.

### HasFollowUpPrompts

`func (o *ChatResponse) HasFollowUpPrompts() bool`

HasFollowUpPrompts returns a boolean if a field has been set.

### GetAgentConfig

`func (o *ChatResponse) GetAgentConfig() AgentConfig`

GetAgentConfig returns the AgentConfig field if non-nil, zero value otherwise.

### GetAgentConfigOk

`func (o *ChatResponse) GetAgentConfigOk() (*AgentConfig, bool)`

GetAgentConfigOk returns a tuple with the AgentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfig

`func (o *ChatResponse) SetAgentConfig(v AgentConfig)`

SetAgentConfig sets AgentConfig field to given value.

### HasAgentConfig

`func (o *ChatResponse) HasAgentConfig() bool`

HasAgentConfig returns a boolean if a field has been set.

### GetBackendTimeMillis

`func (o *ChatResponse) GetBackendTimeMillis() int64`

GetBackendTimeMillis returns the BackendTimeMillis field if non-nil, zero value otherwise.

### GetBackendTimeMillisOk

`func (o *ChatResponse) GetBackendTimeMillisOk() (*int64, bool)`

GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTimeMillis

`func (o *ChatResponse) SetBackendTimeMillis(v int64)`

SetBackendTimeMillis sets BackendTimeMillis field to given value.

### HasBackendTimeMillis

`func (o *ChatResponse) HasBackendTimeMillis() bool`

HasBackendTimeMillis returns a boolean if a field has been set.

### GetChatSessionTrackingToken

`func (o *ChatResponse) GetChatSessionTrackingToken() string`

GetChatSessionTrackingToken returns the ChatSessionTrackingToken field if non-nil, zero value otherwise.

### GetChatSessionTrackingTokenOk

`func (o *ChatResponse) GetChatSessionTrackingTokenOk() (*string, bool)`

GetChatSessionTrackingTokenOk returns a tuple with the ChatSessionTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSessionTrackingToken

`func (o *ChatResponse) SetChatSessionTrackingToken(v string)`

SetChatSessionTrackingToken sets ChatSessionTrackingToken field to given value.

### HasChatSessionTrackingToken

`func (o *ChatResponse) HasChatSessionTrackingToken() bool`

HasChatSessionTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


