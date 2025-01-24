# CreateTeamsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | Pointer to [**[]TeamCreationData**](TeamCreationData.md) | List of all teams from data sources to create | [optional] 

## Methods

### NewCreateTeamsRequest

`func NewCreateTeamsRequest() *CreateTeamsRequest`

NewCreateTeamsRequest instantiates a new CreateTeamsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeamsRequestWithDefaults

`func NewCreateTeamsRequestWithDefaults() *CreateTeamsRequest`

NewCreateTeamsRequestWithDefaults instantiates a new CreateTeamsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *CreateTeamsRequest) GetTeams() []TeamCreationData`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *CreateTeamsRequest) GetTeamsOk() (*[]TeamCreationData, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *CreateTeamsRequest) SetTeams(v []TeamCreationData)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *CreateTeamsRequest) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


