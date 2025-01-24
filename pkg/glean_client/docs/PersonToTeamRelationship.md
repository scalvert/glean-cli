# PersonToTeamRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Person** | [**Person**](Person.md) |  | 
**Relationship** | Pointer to **string** | The team member&#39;s relationship to the team. This defaults to MEMBER if not set. | [optional] [default to "MEMBER"]
**CustomRelationshipStr** | Pointer to **string** | Displayed name for the relationship if relationship is set to &#x60;OTHER&#x60;. | [optional] 
**JoinDate** | Pointer to **time.Time** | The team member&#39;s start date | [optional] 

## Methods

### NewPersonToTeamRelationship

`func NewPersonToTeamRelationship(person Person, ) *PersonToTeamRelationship`

NewPersonToTeamRelationship instantiates a new PersonToTeamRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonToTeamRelationshipWithDefaults

`func NewPersonToTeamRelationshipWithDefaults() *PersonToTeamRelationship`

NewPersonToTeamRelationshipWithDefaults instantiates a new PersonToTeamRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerson

`func (o *PersonToTeamRelationship) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *PersonToTeamRelationship) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *PersonToTeamRelationship) SetPerson(v Person)`

SetPerson sets Person field to given value.


### GetRelationship

`func (o *PersonToTeamRelationship) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *PersonToTeamRelationship) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *PersonToTeamRelationship) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *PersonToTeamRelationship) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetCustomRelationshipStr

`func (o *PersonToTeamRelationship) GetCustomRelationshipStr() string`

GetCustomRelationshipStr returns the CustomRelationshipStr field if non-nil, zero value otherwise.

### GetCustomRelationshipStrOk

`func (o *PersonToTeamRelationship) GetCustomRelationshipStrOk() (*string, bool)`

GetCustomRelationshipStrOk returns a tuple with the CustomRelationshipStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRelationshipStr

`func (o *PersonToTeamRelationship) SetCustomRelationshipStr(v string)`

SetCustomRelationshipStr sets CustomRelationshipStr field to given value.

### HasCustomRelationshipStr

`func (o *PersonToTeamRelationship) HasCustomRelationshipStr() bool`

HasCustomRelationshipStr returns a boolean if a field has been set.

### GetJoinDate

`func (o *PersonToTeamRelationship) GetJoinDate() time.Time`

GetJoinDate returns the JoinDate field if non-nil, zero value otherwise.

### GetJoinDateOk

`func (o *PersonToTeamRelationship) GetJoinDateOk() (*time.Time, bool)`

GetJoinDateOk returns a tuple with the JoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDate

`func (o *PersonToTeamRelationship) SetJoinDate(v time.Time)`

SetJoinDate sets JoinDate field to given value.

### HasJoinDate

`func (o *PersonToTeamRelationship) HasJoinDate() bool`

HasJoinDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


