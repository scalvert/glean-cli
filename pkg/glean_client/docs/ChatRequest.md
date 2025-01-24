# ChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaveChat** | Pointer to **bool** | Save the current interaction as a Chat for the user to access later. | [optional] 
**ChatId** | Pointer to **string** | The id of the Chat that this message should be added to. An empty id signifies creating a new Chat if saveChat is true. | [optional] 
**Messages** | [**[]ChatMessage**](ChatMessage.md) | A list of chat messages, from most recent to least recent. It can be assumed that the first chat message in the list is the user&#39;s most recent query. | 
**AgentConfig** | Pointer to [**AgentConfig**](AgentConfig.md) |  | [optional] 
**Inclusions** | Pointer to [**RestrictionFilters**](RestrictionFilters.md) |  | [optional] 
**Exclusions** | Pointer to [**RestrictionFilters**](RestrictionFilters.md) |  | [optional] 
**TimeoutMillis** | Pointer to **int32** | Timeout in milliseconds for the request. A &#x60;408&#x60; error will be returned if handling the request takes longer. | [optional] 
**ApplicationId** | Pointer to **string** | The ID of the application this request originates from, used to determine the configuration of underlying chat processes. This should correspond to the ID set during admin setup. If not specified, the default chat experience will be used. | [optional] 
**Stream** | Pointer to **bool** | Whether to stream responses as they become available. If false, the entire response will be returned at once. Note if true and the model being used does not support streaming, the model&#39;s response will not be streamed but other messages from the endpoint still will. | [optional] 

## Methods

### NewChatRequest

`func NewChatRequest(messages []ChatMessage, ) *ChatRequest`

NewChatRequest instantiates a new ChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatRequestWithDefaults

`func NewChatRequestWithDefaults() *ChatRequest`

NewChatRequestWithDefaults instantiates a new ChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaveChat

`func (o *ChatRequest) GetSaveChat() bool`

GetSaveChat returns the SaveChat field if non-nil, zero value otherwise.

### GetSaveChatOk

`func (o *ChatRequest) GetSaveChatOk() (*bool, bool)`

GetSaveChatOk returns a tuple with the SaveChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveChat

`func (o *ChatRequest) SetSaveChat(v bool)`

SetSaveChat sets SaveChat field to given value.

### HasSaveChat

`func (o *ChatRequest) HasSaveChat() bool`

HasSaveChat returns a boolean if a field has been set.

### GetChatId

`func (o *ChatRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *ChatRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *ChatRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *ChatRequest) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetMessages

`func (o *ChatRequest) GetMessages() []ChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatRequest) GetMessagesOk() (*[]ChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatRequest) SetMessages(v []ChatMessage)`

SetMessages sets Messages field to given value.


### GetAgentConfig

`func (o *ChatRequest) GetAgentConfig() AgentConfig`

GetAgentConfig returns the AgentConfig field if non-nil, zero value otherwise.

### GetAgentConfigOk

`func (o *ChatRequest) GetAgentConfigOk() (*AgentConfig, bool)`

GetAgentConfigOk returns a tuple with the AgentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfig

`func (o *ChatRequest) SetAgentConfig(v AgentConfig)`

SetAgentConfig sets AgentConfig field to given value.

### HasAgentConfig

`func (o *ChatRequest) HasAgentConfig() bool`

HasAgentConfig returns a boolean if a field has been set.

### GetInclusions

`func (o *ChatRequest) GetInclusions() RestrictionFilters`

GetInclusions returns the Inclusions field if non-nil, zero value otherwise.

### GetInclusionsOk

`func (o *ChatRequest) GetInclusionsOk() (*RestrictionFilters, bool)`

GetInclusionsOk returns a tuple with the Inclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusions

`func (o *ChatRequest) SetInclusions(v RestrictionFilters)`

SetInclusions sets Inclusions field to given value.

### HasInclusions

`func (o *ChatRequest) HasInclusions() bool`

HasInclusions returns a boolean if a field has been set.

### GetExclusions

`func (o *ChatRequest) GetExclusions() RestrictionFilters`

GetExclusions returns the Exclusions field if non-nil, zero value otherwise.

### GetExclusionsOk

`func (o *ChatRequest) GetExclusionsOk() (*RestrictionFilters, bool)`

GetExclusionsOk returns a tuple with the Exclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusions

`func (o *ChatRequest) SetExclusions(v RestrictionFilters)`

SetExclusions sets Exclusions field to given value.

### HasExclusions

`func (o *ChatRequest) HasExclusions() bool`

HasExclusions returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *ChatRequest) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *ChatRequest) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *ChatRequest) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *ChatRequest) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetApplicationId

`func (o *ChatRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ChatRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ChatRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ChatRequest) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetStream

`func (o *ChatRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


