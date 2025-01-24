# PersonTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier | [optional] 
**Name** | Pointer to **string** | Team name | [optional] 
**ExternalLink** | Pointer to **string** | Link to a team page on the internet or your company&#39;s intranet | [optional] 
**Relationship** | Pointer to **string** | The team member&#39;s relationship to the team. This defaults to MEMBER if not set. | [optional] [default to "MEMBER"]
**JoinDate** | Pointer to **time.Time** | The team member&#39;s start date | [optional] 

## Methods

### NewPersonTeam

`func NewPersonTeam() *PersonTeam`

NewPersonTeam instantiates a new PersonTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonTeamWithDefaults

`func NewPersonTeamWithDefaults() *PersonTeam`

NewPersonTeamWithDefaults instantiates a new PersonTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonTeam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonTeam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PersonTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PersonTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonTeam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PersonTeam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalLink

`func (o *PersonTeam) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *PersonTeam) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *PersonTeam) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *PersonTeam) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetRelationship

`func (o *PersonTeam) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *PersonTeam) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *PersonTeam) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *PersonTeam) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetJoinDate

`func (o *PersonTeam) GetJoinDate() time.Time`

GetJoinDate returns the JoinDate field if non-nil, zero value otherwise.

### GetJoinDateOk

`func (o *PersonTeam) GetJoinDateOk() (*time.Time, bool)`

GetJoinDateOk returns a tuple with the JoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDate

`func (o *PersonTeam) SetJoinDate(v time.Time)`

SetJoinDate sets JoinDate field to given value.

### HasJoinDate

`func (o *PersonTeam) HasJoinDate() bool`

HasJoinDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


