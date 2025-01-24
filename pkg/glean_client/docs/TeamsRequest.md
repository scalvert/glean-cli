# TeamsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | The IDs of the teams to retrieve. | [optional] 
**IncludeFields** | Pointer to **[]string** | List of teams fields to return that aren&#39;t returned by default | [optional] 

## Methods

### NewTeamsRequest

`func NewTeamsRequest() *TeamsRequest`

NewTeamsRequest instantiates a new TeamsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsRequestWithDefaults

`func NewTeamsRequestWithDefaults() *TeamsRequest`

NewTeamsRequestWithDefaults instantiates a new TeamsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *TeamsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *TeamsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *TeamsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *TeamsRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetIncludeFields

`func (o *TeamsRequest) GetIncludeFields() []string`

GetIncludeFields returns the IncludeFields field if non-nil, zero value otherwise.

### GetIncludeFieldsOk

`func (o *TeamsRequest) GetIncludeFieldsOk() (*[]string, bool)`

GetIncludeFieldsOk returns a tuple with the IncludeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFields

`func (o *TeamsRequest) SetIncludeFields(v []string)`

SetIncludeFields sets IncludeFields field to given value.

### HasIncludeFields

`func (o *TeamsRequest) HasIncludeFields() bool`

HasIncludeFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


