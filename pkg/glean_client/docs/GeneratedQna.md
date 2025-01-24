# GeneratedQna

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | Pointer to **string** | Search query rephrased into a question. | [optional] 
**Answer** | Pointer to **string** | Answer generated for the given query or the generated question. | [optional] 
**FollowUpPrompts** | Pointer to **[]string** | List of all follow-up prompts generated for the given query or the generated question. | [optional] 
**Ranges** | Pointer to [**[]TextRange**](TextRange.md) | Answer subsections to mark with special formatting (citations, bolding etc) | [optional] 
**Status** | Pointer to **string** | Status of backend generating the answer | [optional] 
**Cursor** | Pointer to **string** | An opaque cursor representing the search request | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular result in this particular query. To be used for /feedback reporting. | [optional] 

## Methods

### NewGeneratedQna

`func NewGeneratedQna() *GeneratedQna`

NewGeneratedQna instantiates a new GeneratedQna object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratedQnaWithDefaults

`func NewGeneratedQnaWithDefaults() *GeneratedQna`

NewGeneratedQnaWithDefaults instantiates a new GeneratedQna object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *GeneratedQna) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *GeneratedQna) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *GeneratedQna) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *GeneratedQna) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswer

`func (o *GeneratedQna) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *GeneratedQna) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *GeneratedQna) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *GeneratedQna) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetFollowUpPrompts

`func (o *GeneratedQna) GetFollowUpPrompts() []string`

GetFollowUpPrompts returns the FollowUpPrompts field if non-nil, zero value otherwise.

### GetFollowUpPromptsOk

`func (o *GeneratedQna) GetFollowUpPromptsOk() (*[]string, bool)`

GetFollowUpPromptsOk returns a tuple with the FollowUpPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpPrompts

`func (o *GeneratedQna) SetFollowUpPrompts(v []string)`

SetFollowUpPrompts sets FollowUpPrompts field to given value.

### HasFollowUpPrompts

`func (o *GeneratedQna) HasFollowUpPrompts() bool`

HasFollowUpPrompts returns a boolean if a field has been set.

### GetRanges

`func (o *GeneratedQna) GetRanges() []TextRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *GeneratedQna) GetRangesOk() (*[]TextRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *GeneratedQna) SetRanges(v []TextRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *GeneratedQna) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetStatus

`func (o *GeneratedQna) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GeneratedQna) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GeneratedQna) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GeneratedQna) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCursor

`func (o *GeneratedQna) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GeneratedQna) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GeneratedQna) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GeneratedQna) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetTrackingToken

`func (o *GeneratedQna) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *GeneratedQna) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *GeneratedQna) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *GeneratedQna) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


