# ExtractedQnA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Heading** | Pointer to **string** | Heading text that was matched to produce this result. | [optional] 
**Question** | Pointer to **string** | Question text that was matched to produce this result. | [optional] 
**QuestionResult** | Pointer to [**SearchResult**](SearchResult.md) |  | [optional] 

## Methods

### NewExtractedQnA

`func NewExtractedQnA() *ExtractedQnA`

NewExtractedQnA instantiates a new ExtractedQnA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractedQnAWithDefaults

`func NewExtractedQnAWithDefaults() *ExtractedQnA`

NewExtractedQnAWithDefaults instantiates a new ExtractedQnA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeading

`func (o *ExtractedQnA) GetHeading() string`

GetHeading returns the Heading field if non-nil, zero value otherwise.

### GetHeadingOk

`func (o *ExtractedQnA) GetHeadingOk() (*string, bool)`

GetHeadingOk returns a tuple with the Heading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeading

`func (o *ExtractedQnA) SetHeading(v string)`

SetHeading sets Heading field to given value.

### HasHeading

`func (o *ExtractedQnA) HasHeading() bool`

HasHeading returns a boolean if a field has been set.

### GetQuestion

`func (o *ExtractedQnA) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *ExtractedQnA) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *ExtractedQnA) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *ExtractedQnA) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetQuestionResult

`func (o *ExtractedQnA) GetQuestionResult() SearchResult`

GetQuestionResult returns the QuestionResult field if non-nil, zero value otherwise.

### GetQuestionResultOk

`func (o *ExtractedQnA) GetQuestionResultOk() (*SearchResult, bool)`

GetQuestionResultOk returns a tuple with the QuestionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionResult

`func (o *ExtractedQnA) SetQuestionResult(v SearchResult)`

SetQuestionResult sets QuestionResult field to given value.

### HasQuestionResult

`func (o *ExtractedQnA) HasQuestionResult() bool`

HasQuestionResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


