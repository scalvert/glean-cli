# TeamsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Team**](Team.md) | A Team and a deep copy of all its members for each ID in the request | [optional] 
**Errors** | Pointer to **[]string** | A list of IDs that could not be found. | [optional] 

## Methods

### NewTeamsResponse

`func NewTeamsResponse() *TeamsResponse`

NewTeamsResponse instantiates a new TeamsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsResponseWithDefaults

`func NewTeamsResponseWithDefaults() *TeamsResponse`

NewTeamsResponseWithDefaults instantiates a new TeamsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *TeamsResponse) GetResults() []Team`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TeamsResponse) GetResultsOk() (*[]Team, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TeamsResponse) SetResults(v []Team)`

SetResults sets Results field to given value.

### HasResults

`func (o *TeamsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetErrors

`func (o *TeamsResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *TeamsResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *TeamsResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *TeamsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


