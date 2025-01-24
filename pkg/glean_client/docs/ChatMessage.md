# ChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentConfig** | Pointer to [**AgentConfig**](AgentConfig.md) |  | [optional] 
**Author** | Pointer to **string** |  | [optional] [default to "USER"]
**Citations** | Pointer to [**[]ChatMessageCitation**](ChatMessageCitation.md) | A list of Citations used to generate the message. | [optional] 
**Fragments** | Pointer to [**[]ChatMessageFragment**](ChatMessageFragment.md) | A list of chat results. | [optional] 
**Metadata** | Pointer to **string** | Metadata associated with the message (not displayed to the user but stored in the app). | [optional] 
**Ts** | Pointer to **string** | Timestamp of the message. | [optional] 
**MessageId** | Pointer to **string** | Unique ID of the message. | [optional] 
**MessageTrackingToken** | Pointer to **string** | Opaque tracking token generated server-side. | [optional] 
**MessageType** | Pointer to **string** | Used to determine the type of UI treatment to apply to this message. UPDATE - intermediate state message for progress updates before content responses. CONTENT - contains content relevant to the user query. CONTEXT - contains additional context relevant to the user query. DEBUG - contains debug information of ChatBot behavior. ERROR - an error happened on server side. | [optional] 
**HasMoreFragments** | Pointer to **bool** | Signals there are more fragments incoming. | [optional] 

## Methods

### NewChatMessage

`func NewChatMessage() *ChatMessage`

NewChatMessage instantiates a new ChatMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageWithDefaults

`func NewChatMessageWithDefaults() *ChatMessage`

NewChatMessageWithDefaults instantiates a new ChatMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentConfig

`func (o *ChatMessage) GetAgentConfig() AgentConfig`

GetAgentConfig returns the AgentConfig field if non-nil, zero value otherwise.

### GetAgentConfigOk

`func (o *ChatMessage) GetAgentConfigOk() (*AgentConfig, bool)`

GetAgentConfigOk returns a tuple with the AgentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentConfig

`func (o *ChatMessage) SetAgentConfig(v AgentConfig)`

SetAgentConfig sets AgentConfig field to given value.

### HasAgentConfig

`func (o *ChatMessage) HasAgentConfig() bool`

HasAgentConfig returns a boolean if a field has been set.

### GetAuthor

`func (o *ChatMessage) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ChatMessage) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ChatMessage) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ChatMessage) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCitations

`func (o *ChatMessage) GetCitations() []ChatMessageCitation`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *ChatMessage) GetCitationsOk() (*[]ChatMessageCitation, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *ChatMessage) SetCitations(v []ChatMessageCitation)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *ChatMessage) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetFragments

`func (o *ChatMessage) GetFragments() []ChatMessageFragment`

GetFragments returns the Fragments field if non-nil, zero value otherwise.

### GetFragmentsOk

`func (o *ChatMessage) GetFragmentsOk() (*[]ChatMessageFragment, bool)`

GetFragmentsOk returns a tuple with the Fragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragments

`func (o *ChatMessage) SetFragments(v []ChatMessageFragment)`

SetFragments sets Fragments field to given value.

### HasFragments

`func (o *ChatMessage) HasFragments() bool`

HasFragments returns a boolean if a field has been set.

### GetMetadata

`func (o *ChatMessage) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChatMessage) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChatMessage) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ChatMessage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTs

`func (o *ChatMessage) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ChatMessage) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ChatMessage) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ChatMessage) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetMessageId

`func (o *ChatMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ChatMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ChatMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ChatMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageTrackingToken

`func (o *ChatMessage) GetMessageTrackingToken() string`

GetMessageTrackingToken returns the MessageTrackingToken field if non-nil, zero value otherwise.

### GetMessageTrackingTokenOk

`func (o *ChatMessage) GetMessageTrackingTokenOk() (*string, bool)`

GetMessageTrackingTokenOk returns a tuple with the MessageTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTrackingToken

`func (o *ChatMessage) SetMessageTrackingToken(v string)`

SetMessageTrackingToken sets MessageTrackingToken field to given value.

### HasMessageTrackingToken

`func (o *ChatMessage) HasMessageTrackingToken() bool`

HasMessageTrackingToken returns a boolean if a field has been set.

### GetMessageType

`func (o *ChatMessage) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ChatMessage) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ChatMessage) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ChatMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetHasMoreFragments

`func (o *ChatMessage) GetHasMoreFragments() bool`

GetHasMoreFragments returns the HasMoreFragments field if non-nil, zero value otherwise.

### GetHasMoreFragmentsOk

`func (o *ChatMessage) GetHasMoreFragmentsOk() (*bool, bool)`

GetHasMoreFragmentsOk returns a tuple with the HasMoreFragments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreFragments

`func (o *ChatMessage) SetHasMoreFragments(v bool)`

SetHasMoreFragments sets HasMoreFragments field to given value.

### HasHasMoreFragments

`func (o *ChatMessage) HasHasMoreFragments() bool`

HasHasMoreFragments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


