# DeleteQueryHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | Pointer to **[]string** | Queries to delete. | [optional] 

## Methods

### NewDeleteQueryHistoryRequest

`func NewDeleteQueryHistoryRequest() *DeleteQueryHistoryRequest`

NewDeleteQueryHistoryRequest instantiates a new DeleteQueryHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteQueryHistoryRequestWithDefaults

`func NewDeleteQueryHistoryRequestWithDefaults() *DeleteQueryHistoryRequest`

NewDeleteQueryHistoryRequestWithDefaults instantiates a new DeleteQueryHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *DeleteQueryHistoryRequest) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *DeleteQueryHistoryRequest) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *DeleteQueryHistoryRequest) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *DeleteQueryHistoryRequest) HasQueries() bool`

HasQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


