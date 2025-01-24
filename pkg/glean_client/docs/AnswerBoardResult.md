# AnswerBoardResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Board** | [**AnswerBoard**](AnswerBoard.md) |  | 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular Answer Board. To be used for /feedback reporting. | [optional] 

## Methods

### NewAnswerBoardResult

`func NewAnswerBoardResult(board AnswerBoard, ) *AnswerBoardResult`

NewAnswerBoardResult instantiates a new AnswerBoardResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerBoardResultWithDefaults

`func NewAnswerBoardResultWithDefaults() *AnswerBoardResult`

NewAnswerBoardResultWithDefaults instantiates a new AnswerBoardResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoard

`func (o *AnswerBoardResult) GetBoard() AnswerBoard`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *AnswerBoardResult) GetBoardOk() (*AnswerBoard, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *AnswerBoardResult) SetBoard(v AnswerBoard)`

SetBoard sets Board field to given value.


### GetTrackingToken

`func (o *AnswerBoardResult) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *AnswerBoardResult) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *AnswerBoardResult) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *AnswerBoardResult) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


