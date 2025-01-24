# DeleteTeamsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumErrors** | Pointer to **int32** | Number of teams that failed to be deleted | [optional] 

## Methods

### NewDeleteTeamsResponse

`func NewDeleteTeamsResponse() *DeleteTeamsResponse`

NewDeleteTeamsResponse instantiates a new DeleteTeamsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTeamsResponseWithDefaults

`func NewDeleteTeamsResponseWithDefaults() *DeleteTeamsResponse`

NewDeleteTeamsResponseWithDefaults instantiates a new DeleteTeamsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumErrors

`func (o *DeleteTeamsResponse) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *DeleteTeamsResponse) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *DeleteTeamsResponse) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *DeleteTeamsResponse) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


