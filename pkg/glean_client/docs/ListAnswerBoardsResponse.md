# ListAnswerBoardsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardResults** | [**[]AnswerBoardResult**](AnswerBoardResult.md) | List of all Answer Boards, no Answers are included. | 

## Methods

### NewListAnswerBoardsResponse

`func NewListAnswerBoardsResponse(boardResults []AnswerBoardResult, ) *ListAnswerBoardsResponse`

NewListAnswerBoardsResponse instantiates a new ListAnswerBoardsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnswerBoardsResponseWithDefaults

`func NewListAnswerBoardsResponseWithDefaults() *ListAnswerBoardsResponse`

NewListAnswerBoardsResponseWithDefaults instantiates a new ListAnswerBoardsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardResults

`func (o *ListAnswerBoardsResponse) GetBoardResults() []AnswerBoardResult`

GetBoardResults returns the BoardResults field if non-nil, zero value otherwise.

### GetBoardResultsOk

`func (o *ListAnswerBoardsResponse) GetBoardResultsOk() (*[]AnswerBoardResult, bool)`

GetBoardResultsOk returns a tuple with the BoardResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardResults

`func (o *ListAnswerBoardsResponse) SetBoardResults(v []AnswerBoardResult)`

SetBoardResults sets BoardResults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


