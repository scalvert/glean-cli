# EditAnswerBoardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardResult** | Pointer to [**AnswerBoardResult**](AnswerBoardResult.md) |  | [optional] 
**Error** | Pointer to [**AnswerBoardError**](AnswerBoardError.md) |  | [optional] 

## Methods

### NewEditAnswerBoardResponse

`func NewEditAnswerBoardResponse() *EditAnswerBoardResponse`

NewEditAnswerBoardResponse instantiates a new EditAnswerBoardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAnswerBoardResponseWithDefaults

`func NewEditAnswerBoardResponseWithDefaults() *EditAnswerBoardResponse`

NewEditAnswerBoardResponseWithDefaults instantiates a new EditAnswerBoardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardResult

`func (o *EditAnswerBoardResponse) GetBoardResult() AnswerBoardResult`

GetBoardResult returns the BoardResult field if non-nil, zero value otherwise.

### GetBoardResultOk

`func (o *EditAnswerBoardResponse) GetBoardResultOk() (*AnswerBoardResult, bool)`

GetBoardResultOk returns a tuple with the BoardResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardResult

`func (o *EditAnswerBoardResponse) SetBoardResult(v AnswerBoardResult)`

SetBoardResult sets BoardResult field to given value.

### HasBoardResult

`func (o *EditAnswerBoardResponse) HasBoardResult() bool`

HasBoardResult returns a boolean if a field has been set.

### GetError

`func (o *EditAnswerBoardResponse) GetError() AnswerBoardError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EditAnswerBoardResponse) GetErrorOk() (*AnswerBoardError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EditAnswerBoardResponse) SetError(v AnswerBoardError)`

SetError sets Error field to given value.

### HasError

`func (o *EditAnswerBoardResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


