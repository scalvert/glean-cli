# AiAppActionCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSlackbotResponses** | Pointer to **int32** | Total number of Slackbot responses, both proactive and reactive. | [optional] 
**TotalSlackbotResponsesShared** | Pointer to **int32** | Total number of Slackbot responses shared publicly (upvoted). | [optional] 
**TotalSlackbotResponsesNotHelpful** | Pointer to **int32** | Total number of Slackbot responses rejected as not helpful (downvoted). | [optional] 
**TotalChatMessages** | Pointer to **int32** | Total number of Chat messages sent in requested period. | [optional] 
**TotalUpvotes** | Pointer to **int32** | Total number of Chat messages which received upvotes by the user. | [optional] 
**TotalDownvotes** | Pointer to **int32** | Total number of Chat messages which received downvotes by the user. | [optional] 

## Methods

### NewAiAppActionCounts

`func NewAiAppActionCounts() *AiAppActionCounts`

NewAiAppActionCounts instantiates a new AiAppActionCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAppActionCountsWithDefaults

`func NewAiAppActionCountsWithDefaults() *AiAppActionCounts`

NewAiAppActionCountsWithDefaults instantiates a new AiAppActionCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSlackbotResponses

`func (o *AiAppActionCounts) GetTotalSlackbotResponses() int32`

GetTotalSlackbotResponses returns the TotalSlackbotResponses field if non-nil, zero value otherwise.

### GetTotalSlackbotResponsesOk

`func (o *AiAppActionCounts) GetTotalSlackbotResponsesOk() (*int32, bool)`

GetTotalSlackbotResponsesOk returns a tuple with the TotalSlackbotResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSlackbotResponses

`func (o *AiAppActionCounts) SetTotalSlackbotResponses(v int32)`

SetTotalSlackbotResponses sets TotalSlackbotResponses field to given value.

### HasTotalSlackbotResponses

`func (o *AiAppActionCounts) HasTotalSlackbotResponses() bool`

HasTotalSlackbotResponses returns a boolean if a field has been set.

### GetTotalSlackbotResponsesShared

`func (o *AiAppActionCounts) GetTotalSlackbotResponsesShared() int32`

GetTotalSlackbotResponsesShared returns the TotalSlackbotResponsesShared field if non-nil, zero value otherwise.

### GetTotalSlackbotResponsesSharedOk

`func (o *AiAppActionCounts) GetTotalSlackbotResponsesSharedOk() (*int32, bool)`

GetTotalSlackbotResponsesSharedOk returns a tuple with the TotalSlackbotResponsesShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSlackbotResponsesShared

`func (o *AiAppActionCounts) SetTotalSlackbotResponsesShared(v int32)`

SetTotalSlackbotResponsesShared sets TotalSlackbotResponsesShared field to given value.

### HasTotalSlackbotResponsesShared

`func (o *AiAppActionCounts) HasTotalSlackbotResponsesShared() bool`

HasTotalSlackbotResponsesShared returns a boolean if a field has been set.

### GetTotalSlackbotResponsesNotHelpful

`func (o *AiAppActionCounts) GetTotalSlackbotResponsesNotHelpful() int32`

GetTotalSlackbotResponsesNotHelpful returns the TotalSlackbotResponsesNotHelpful field if non-nil, zero value otherwise.

### GetTotalSlackbotResponsesNotHelpfulOk

`func (o *AiAppActionCounts) GetTotalSlackbotResponsesNotHelpfulOk() (*int32, bool)`

GetTotalSlackbotResponsesNotHelpfulOk returns a tuple with the TotalSlackbotResponsesNotHelpful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSlackbotResponsesNotHelpful

`func (o *AiAppActionCounts) SetTotalSlackbotResponsesNotHelpful(v int32)`

SetTotalSlackbotResponsesNotHelpful sets TotalSlackbotResponsesNotHelpful field to given value.

### HasTotalSlackbotResponsesNotHelpful

`func (o *AiAppActionCounts) HasTotalSlackbotResponsesNotHelpful() bool`

HasTotalSlackbotResponsesNotHelpful returns a boolean if a field has been set.

### GetTotalChatMessages

`func (o *AiAppActionCounts) GetTotalChatMessages() int32`

GetTotalChatMessages returns the TotalChatMessages field if non-nil, zero value otherwise.

### GetTotalChatMessagesOk

`func (o *AiAppActionCounts) GetTotalChatMessagesOk() (*int32, bool)`

GetTotalChatMessagesOk returns a tuple with the TotalChatMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChatMessages

`func (o *AiAppActionCounts) SetTotalChatMessages(v int32)`

SetTotalChatMessages sets TotalChatMessages field to given value.

### HasTotalChatMessages

`func (o *AiAppActionCounts) HasTotalChatMessages() bool`

HasTotalChatMessages returns a boolean if a field has been set.

### GetTotalUpvotes

`func (o *AiAppActionCounts) GetTotalUpvotes() int32`

GetTotalUpvotes returns the TotalUpvotes field if non-nil, zero value otherwise.

### GetTotalUpvotesOk

`func (o *AiAppActionCounts) GetTotalUpvotesOk() (*int32, bool)`

GetTotalUpvotesOk returns a tuple with the TotalUpvotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUpvotes

`func (o *AiAppActionCounts) SetTotalUpvotes(v int32)`

SetTotalUpvotes sets TotalUpvotes field to given value.

### HasTotalUpvotes

`func (o *AiAppActionCounts) HasTotalUpvotes() bool`

HasTotalUpvotes returns a boolean if a field has been set.

### GetTotalDownvotes

`func (o *AiAppActionCounts) GetTotalDownvotes() int32`

GetTotalDownvotes returns the TotalDownvotes field if non-nil, zero value otherwise.

### GetTotalDownvotesOk

`func (o *AiAppActionCounts) GetTotalDownvotesOk() (*int32, bool)`

GetTotalDownvotesOk returns a tuple with the TotalDownvotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDownvotes

`func (o *AiAppActionCounts) SetTotalDownvotes(v int32)`

SetTotalDownvotes sets TotalDownvotes field to given value.

### HasTotalDownvotes

`func (o *AiAppActionCounts) HasTotalDownvotes() bool`

HasTotalDownvotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


