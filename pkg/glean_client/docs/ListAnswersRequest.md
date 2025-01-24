# ListAnswersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardId** | Pointer to **int32** | The Answer Board Id to list answers on. | [optional] 

## Methods

### NewListAnswersRequest

`func NewListAnswersRequest() *ListAnswersRequest`

NewListAnswersRequest instantiates a new ListAnswersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnswersRequestWithDefaults

`func NewListAnswersRequestWithDefaults() *ListAnswersRequest`

NewListAnswersRequestWithDefaults instantiates a new ListAnswersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardId

`func (o *ListAnswersRequest) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *ListAnswersRequest) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *ListAnswersRequest) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *ListAnswersRequest) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


