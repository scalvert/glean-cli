# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**FollowUpPrompts** | Pointer to **[]string** | Follow-up prompts based on the summarized doc | [optional] 

## Methods

### NewSummary

`func NewSummary() *Summary`

NewSummary instantiates a new Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryWithDefaults

`func NewSummaryWithDefaults() *Summary`

NewSummaryWithDefaults instantiates a new Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *Summary) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Summary) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Summary) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Summary) HasText() bool`

HasText returns a boolean if a field has been set.

### GetFollowUpPrompts

`func (o *Summary) GetFollowUpPrompts() []string`

GetFollowUpPrompts returns the FollowUpPrompts field if non-nil, zero value otherwise.

### GetFollowUpPromptsOk

`func (o *Summary) GetFollowUpPromptsOk() (*[]string, bool)`

GetFollowUpPromptsOk returns a tuple with the FollowUpPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpPrompts

`func (o *Summary) SetFollowUpPrompts(v []string)`

SetFollowUpPrompts sets FollowUpPrompts field to given value.

### HasFollowUpPrompts

`func (o *Summary) HasFollowUpPrompts() bool`

HasFollowUpPrompts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


