# DeleteQueryHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**DeleteQueryHistoryError**](DeleteQueryHistoryError.md) |  | [optional] 

## Methods

### NewDeleteQueryHistoryResponse

`func NewDeleteQueryHistoryResponse() *DeleteQueryHistoryResponse`

NewDeleteQueryHistoryResponse instantiates a new DeleteQueryHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteQueryHistoryResponseWithDefaults

`func NewDeleteQueryHistoryResponseWithDefaults() *DeleteQueryHistoryResponse`

NewDeleteQueryHistoryResponseWithDefaults instantiates a new DeleteQueryHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DeleteQueryHistoryResponse) GetError() DeleteQueryHistoryError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeleteQueryHistoryResponse) GetErrorOk() (*DeleteQueryHistoryError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeleteQueryHistoryResponse) SetError(v DeleteQueryHistoryError)`

SetError sets Error field to given value.

### HasError

`func (o *DeleteQueryHistoryResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


