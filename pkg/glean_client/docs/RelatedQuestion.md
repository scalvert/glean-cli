# RelatedQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | Pointer to **string** | The text of the related question | [optional] 
**Answer** | Pointer to **string** | The answer for the related question | [optional] 
**Ranges** | Pointer to [**[]TextRange**](TextRange.md) | Subsections of the answer string to which some special formatting should be applied (eg. bold) | [optional] 

## Methods

### NewRelatedQuestion

`func NewRelatedQuestion() *RelatedQuestion`

NewRelatedQuestion instantiates a new RelatedQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedQuestionWithDefaults

`func NewRelatedQuestionWithDefaults() *RelatedQuestion`

NewRelatedQuestionWithDefaults instantiates a new RelatedQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *RelatedQuestion) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *RelatedQuestion) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *RelatedQuestion) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *RelatedQuestion) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswer

`func (o *RelatedQuestion) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *RelatedQuestion) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *RelatedQuestion) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *RelatedQuestion) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetRanges

`func (o *RelatedQuestion) GetRanges() []TextRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *RelatedQuestion) GetRangesOk() (*[]TextRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *RelatedQuestion) SetRanges(v []TextRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *RelatedQuestion) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


