# AnswerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | [**Answer**](Answer.md) |  | 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular Answer. To be used for &#x60;/feedback&#x60; reporting. | [optional] 

## Methods

### NewAnswerResult

`func NewAnswerResult(answer Answer, ) *AnswerResult`

NewAnswerResult instantiates a new AnswerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerResultWithDefaults

`func NewAnswerResultWithDefaults() *AnswerResult`

NewAnswerResultWithDefaults instantiates a new AnswerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *AnswerResult) GetAnswer() Answer`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *AnswerResult) GetAnswerOk() (*Answer, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *AnswerResult) SetAnswer(v Answer)`

SetAnswer sets Answer field to given value.


### GetTrackingToken

`func (o *AnswerResult) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *AnswerResult) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *AnswerResult) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *AnswerResult) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


