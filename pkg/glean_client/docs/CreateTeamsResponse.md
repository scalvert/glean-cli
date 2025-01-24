# CreateTeamsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumErrors** | Pointer to **int32** | Number of teams that failed to be created | [optional] 

## Methods

### NewCreateTeamsResponse

`func NewCreateTeamsResponse() *CreateTeamsResponse`

NewCreateTeamsResponse instantiates a new CreateTeamsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeamsResponseWithDefaults

`func NewCreateTeamsResponseWithDefaults() *CreateTeamsResponse`

NewCreateTeamsResponseWithDefaults instantiates a new CreateTeamsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumErrors

`func (o *CreateTeamsResponse) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *CreateTeamsResponse) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *CreateTeamsResponse) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *CreateTeamsResponse) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


