# EmailRequestChatFeedbackPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | Pointer to **string** | Rating given to the conversation (currently either \&quot;upvoted\&quot; or \&quot;downvoted\&quot;). | [optional] 
**Comments** | Pointer to **string** | Additional freeform comments provided by the reporter. | [optional] 
**PreviousMessages** | Pointer to **[]string** | Previous messages in this conversation. | [optional] 

## Methods

### NewEmailRequestChatFeedbackPayload

`func NewEmailRequestChatFeedbackPayload() *EmailRequestChatFeedbackPayload`

NewEmailRequestChatFeedbackPayload instantiates a new EmailRequestChatFeedbackPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRequestChatFeedbackPayloadWithDefaults

`func NewEmailRequestChatFeedbackPayloadWithDefaults() *EmailRequestChatFeedbackPayload`

NewEmailRequestChatFeedbackPayloadWithDefaults instantiates a new EmailRequestChatFeedbackPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *EmailRequestChatFeedbackPayload) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *EmailRequestChatFeedbackPayload) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *EmailRequestChatFeedbackPayload) SetRating(v string)`

SetRating sets Rating field to given value.

### HasRating

`func (o *EmailRequestChatFeedbackPayload) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetComments

`func (o *EmailRequestChatFeedbackPayload) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *EmailRequestChatFeedbackPayload) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *EmailRequestChatFeedbackPayload) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *EmailRequestChatFeedbackPayload) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPreviousMessages

`func (o *EmailRequestChatFeedbackPayload) GetPreviousMessages() []string`

GetPreviousMessages returns the PreviousMessages field if non-nil, zero value otherwise.

### GetPreviousMessagesOk

`func (o *EmailRequestChatFeedbackPayload) GetPreviousMessagesOk() (*[]string, bool)`

GetPreviousMessagesOk returns a tuple with the PreviousMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMessages

`func (o *EmailRequestChatFeedbackPayload) SetPreviousMessages(v []string)`

SetPreviousMessages sets PreviousMessages field to given value.

### HasPreviousMessages

`func (o *EmailRequestChatFeedbackPayload) HasPreviousMessages() bool`

HasPreviousMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


