# CreateAnswerBoardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardResult** | Pointer to [**AnswerBoardResult**](AnswerBoardResult.md) |  | [optional] 
**Error** | Pointer to [**AnswerBoardError**](AnswerBoardError.md) |  | [optional] 

## Methods

### NewCreateAnswerBoardResponse

`func NewCreateAnswerBoardResponse() *CreateAnswerBoardResponse`

NewCreateAnswerBoardResponse instantiates a new CreateAnswerBoardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAnswerBoardResponseWithDefaults

`func NewCreateAnswerBoardResponseWithDefaults() *CreateAnswerBoardResponse`

NewCreateAnswerBoardResponseWithDefaults instantiates a new CreateAnswerBoardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardResult

`func (o *CreateAnswerBoardResponse) GetBoardResult() AnswerBoardResult`

GetBoardResult returns the BoardResult field if non-nil, zero value otherwise.

### GetBoardResultOk

`func (o *CreateAnswerBoardResponse) GetBoardResultOk() (*AnswerBoardResult, bool)`

GetBoardResultOk returns a tuple with the BoardResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardResult

`func (o *CreateAnswerBoardResponse) SetBoardResult(v AnswerBoardResult)`

SetBoardResult sets BoardResult field to given value.

### HasBoardResult

`func (o *CreateAnswerBoardResponse) HasBoardResult() bool`

HasBoardResult returns a boolean if a field has been set.

### GetError

`func (o *CreateAnswerBoardResponse) GetError() AnswerBoardError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateAnswerBoardResponse) GetErrorOk() (*AnswerBoardError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateAnswerBoardResponse) SetError(v AnswerBoardError)`

SetError sets Error field to given value.

### HasError

`func (o *CreateAnswerBoardResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


