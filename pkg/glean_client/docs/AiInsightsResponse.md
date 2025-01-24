# AiInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogTimestamp** | Pointer to **int32** | Unix timestamp of the last activity processed to make the response (in seconds since epoch UTC). | [optional] 
**AssistantInsights** | Pointer to [**[]UserActivityInsight**](UserActivityInsight.md) |  | [optional] 
**TotalActiveAssistantUsers** | Pointer to **int32** | Total number of Active Assistant users (chat, summary, AIA) in requested period. | [optional] 
**TotalChatMessages** | Pointer to **int32** | Total number of Chat messages sent in requested period. | [optional] 
**TotalAiSummarizations** | Pointer to **int32** | Total number of AI Document Summarizations invoked in the requested period. | [optional] 
**TotalAiAnswers** | Pointer to **int32** | Total number of AI Answers generated in the requested period. | [optional] 
**TotalUpvotes** | Pointer to **int32** | Total number of Chat messages which received upvotes by the user. | [optional] 
**TotalDownvotes** | Pointer to **int32** | Total number of Chat messages which received downvotes by the user. | [optional] 
**TotalGleanbotResponses** | Pointer to **int32** | Total number of Gleanbot responses, both proactive and reactive. | [optional] 
**TotalGleanbotResponsesShared** | Pointer to **int32** | Total number of Gleanbot responses shared publicly (upvoted). | [optional] 
**TotalGleanbotResponsesNotHelpful** | Pointer to **int32** | Total number of Glean responses rejected as not helpful (downvoted). | [optional] 
**Departments** | Pointer to **[]string** | list of departments applicable for users tab. | [optional] 

## Methods

### NewAiInsightsResponse

`func NewAiInsightsResponse() *AiInsightsResponse`

NewAiInsightsResponse instantiates a new AiInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiInsightsResponseWithDefaults

`func NewAiInsightsResponseWithDefaults() *AiInsightsResponse`

NewAiInsightsResponseWithDefaults instantiates a new AiInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogTimestamp

`func (o *AiInsightsResponse) GetLastLogTimestamp() int32`

GetLastLogTimestamp returns the LastLogTimestamp field if non-nil, zero value otherwise.

### GetLastLogTimestampOk

`func (o *AiInsightsResponse) GetLastLogTimestampOk() (*int32, bool)`

GetLastLogTimestampOk returns a tuple with the LastLogTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogTimestamp

`func (o *AiInsightsResponse) SetLastLogTimestamp(v int32)`

SetLastLogTimestamp sets LastLogTimestamp field to given value.

### HasLastLogTimestamp

`func (o *AiInsightsResponse) HasLastLogTimestamp() bool`

HasLastLogTimestamp returns a boolean if a field has been set.

### GetAssistantInsights

`func (o *AiInsightsResponse) GetAssistantInsights() []UserActivityInsight`

GetAssistantInsights returns the AssistantInsights field if non-nil, zero value otherwise.

### GetAssistantInsightsOk

`func (o *AiInsightsResponse) GetAssistantInsightsOk() (*[]UserActivityInsight, bool)`

GetAssistantInsightsOk returns a tuple with the AssistantInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantInsights

`func (o *AiInsightsResponse) SetAssistantInsights(v []UserActivityInsight)`

SetAssistantInsights sets AssistantInsights field to given value.

### HasAssistantInsights

`func (o *AiInsightsResponse) HasAssistantInsights() bool`

HasAssistantInsights returns a boolean if a field has been set.

### GetTotalActiveAssistantUsers

`func (o *AiInsightsResponse) GetTotalActiveAssistantUsers() int32`

GetTotalActiveAssistantUsers returns the TotalActiveAssistantUsers field if non-nil, zero value otherwise.

### GetTotalActiveAssistantUsersOk

`func (o *AiInsightsResponse) GetTotalActiveAssistantUsersOk() (*int32, bool)`

GetTotalActiveAssistantUsersOk returns a tuple with the TotalActiveAssistantUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActiveAssistantUsers

`func (o *AiInsightsResponse) SetTotalActiveAssistantUsers(v int32)`

SetTotalActiveAssistantUsers sets TotalActiveAssistantUsers field to given value.

### HasTotalActiveAssistantUsers

`func (o *AiInsightsResponse) HasTotalActiveAssistantUsers() bool`

HasTotalActiveAssistantUsers returns a boolean if a field has been set.

### GetTotalChatMessages

`func (o *AiInsightsResponse) GetTotalChatMessages() int32`

GetTotalChatMessages returns the TotalChatMessages field if non-nil, zero value otherwise.

### GetTotalChatMessagesOk

`func (o *AiInsightsResponse) GetTotalChatMessagesOk() (*int32, bool)`

GetTotalChatMessagesOk returns a tuple with the TotalChatMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChatMessages

`func (o *AiInsightsResponse) SetTotalChatMessages(v int32)`

SetTotalChatMessages sets TotalChatMessages field to given value.

### HasTotalChatMessages

`func (o *AiInsightsResponse) HasTotalChatMessages() bool`

HasTotalChatMessages returns a boolean if a field has been set.

### GetTotalAiSummarizations

`func (o *AiInsightsResponse) GetTotalAiSummarizations() int32`

GetTotalAiSummarizations returns the TotalAiSummarizations field if non-nil, zero value otherwise.

### GetTotalAiSummarizationsOk

`func (o *AiInsightsResponse) GetTotalAiSummarizationsOk() (*int32, bool)`

GetTotalAiSummarizationsOk returns a tuple with the TotalAiSummarizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAiSummarizations

`func (o *AiInsightsResponse) SetTotalAiSummarizations(v int32)`

SetTotalAiSummarizations sets TotalAiSummarizations field to given value.

### HasTotalAiSummarizations

`func (o *AiInsightsResponse) HasTotalAiSummarizations() bool`

HasTotalAiSummarizations returns a boolean if a field has been set.

### GetTotalAiAnswers

`func (o *AiInsightsResponse) GetTotalAiAnswers() int32`

GetTotalAiAnswers returns the TotalAiAnswers field if non-nil, zero value otherwise.

### GetTotalAiAnswersOk

`func (o *AiInsightsResponse) GetTotalAiAnswersOk() (*int32, bool)`

GetTotalAiAnswersOk returns a tuple with the TotalAiAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAiAnswers

`func (o *AiInsightsResponse) SetTotalAiAnswers(v int32)`

SetTotalAiAnswers sets TotalAiAnswers field to given value.

### HasTotalAiAnswers

`func (o *AiInsightsResponse) HasTotalAiAnswers() bool`

HasTotalAiAnswers returns a boolean if a field has been set.

### GetTotalUpvotes

`func (o *AiInsightsResponse) GetTotalUpvotes() int32`

GetTotalUpvotes returns the TotalUpvotes field if non-nil, zero value otherwise.

### GetTotalUpvotesOk

`func (o *AiInsightsResponse) GetTotalUpvotesOk() (*int32, bool)`

GetTotalUpvotesOk returns a tuple with the TotalUpvotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUpvotes

`func (o *AiInsightsResponse) SetTotalUpvotes(v int32)`

SetTotalUpvotes sets TotalUpvotes field to given value.

### HasTotalUpvotes

`func (o *AiInsightsResponse) HasTotalUpvotes() bool`

HasTotalUpvotes returns a boolean if a field has been set.

### GetTotalDownvotes

`func (o *AiInsightsResponse) GetTotalDownvotes() int32`

GetTotalDownvotes returns the TotalDownvotes field if non-nil, zero value otherwise.

### GetTotalDownvotesOk

`func (o *AiInsightsResponse) GetTotalDownvotesOk() (*int32, bool)`

GetTotalDownvotesOk returns a tuple with the TotalDownvotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDownvotes

`func (o *AiInsightsResponse) SetTotalDownvotes(v int32)`

SetTotalDownvotes sets TotalDownvotes field to given value.

### HasTotalDownvotes

`func (o *AiInsightsResponse) HasTotalDownvotes() bool`

HasTotalDownvotes returns a boolean if a field has been set.

### GetTotalGleanbotResponses

`func (o *AiInsightsResponse) GetTotalGleanbotResponses() int32`

GetTotalGleanbotResponses returns the TotalGleanbotResponses field if non-nil, zero value otherwise.

### GetTotalGleanbotResponsesOk

`func (o *AiInsightsResponse) GetTotalGleanbotResponsesOk() (*int32, bool)`

GetTotalGleanbotResponsesOk returns a tuple with the TotalGleanbotResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGleanbotResponses

`func (o *AiInsightsResponse) SetTotalGleanbotResponses(v int32)`

SetTotalGleanbotResponses sets TotalGleanbotResponses field to given value.

### HasTotalGleanbotResponses

`func (o *AiInsightsResponse) HasTotalGleanbotResponses() bool`

HasTotalGleanbotResponses returns a boolean if a field has been set.

### GetTotalGleanbotResponsesShared

`func (o *AiInsightsResponse) GetTotalGleanbotResponsesShared() int32`

GetTotalGleanbotResponsesShared returns the TotalGleanbotResponsesShared field if non-nil, zero value otherwise.

### GetTotalGleanbotResponsesSharedOk

`func (o *AiInsightsResponse) GetTotalGleanbotResponsesSharedOk() (*int32, bool)`

GetTotalGleanbotResponsesSharedOk returns a tuple with the TotalGleanbotResponsesShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGleanbotResponsesShared

`func (o *AiInsightsResponse) SetTotalGleanbotResponsesShared(v int32)`

SetTotalGleanbotResponsesShared sets TotalGleanbotResponsesShared field to given value.

### HasTotalGleanbotResponsesShared

`func (o *AiInsightsResponse) HasTotalGleanbotResponsesShared() bool`

HasTotalGleanbotResponsesShared returns a boolean if a field has been set.

### GetTotalGleanbotResponsesNotHelpful

`func (o *AiInsightsResponse) GetTotalGleanbotResponsesNotHelpful() int32`

GetTotalGleanbotResponsesNotHelpful returns the TotalGleanbotResponsesNotHelpful field if non-nil, zero value otherwise.

### GetTotalGleanbotResponsesNotHelpfulOk

`func (o *AiInsightsResponse) GetTotalGleanbotResponsesNotHelpfulOk() (*int32, bool)`

GetTotalGleanbotResponsesNotHelpfulOk returns a tuple with the TotalGleanbotResponsesNotHelpful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGleanbotResponsesNotHelpful

`func (o *AiInsightsResponse) SetTotalGleanbotResponsesNotHelpful(v int32)`

SetTotalGleanbotResponsesNotHelpful sets TotalGleanbotResponsesNotHelpful field to given value.

### HasTotalGleanbotResponsesNotHelpful

`func (o *AiInsightsResponse) HasTotalGleanbotResponsesNotHelpful() bool`

HasTotalGleanbotResponsesNotHelpful returns a boolean if a field has been set.

### GetDepartments

`func (o *AiInsightsResponse) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *AiInsightsResponse) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *AiInsightsResponse) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *AiInsightsResponse) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


