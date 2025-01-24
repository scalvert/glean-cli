# UserRoleSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDocumentSpec** | Pointer to [**DocumentSpec**](DocumentSpec.md) |  | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Group** | Pointer to [**Group**](Group.md) |  | [optional] 
**Role** | [**UserRole**](UserRole.md) |  | 

## Methods

### NewUserRoleSpecification

`func NewUserRoleSpecification(role UserRole, ) *UserRoleSpecification`

NewUserRoleSpecification instantiates a new UserRoleSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleSpecificationWithDefaults

`func NewUserRoleSpecificationWithDefaults() *UserRoleSpecification`

NewUserRoleSpecificationWithDefaults instantiates a new UserRoleSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDocumentSpec

`func (o *UserRoleSpecification) GetSourceDocumentSpec() DocumentSpec`

GetSourceDocumentSpec returns the SourceDocumentSpec field if non-nil, zero value otherwise.

### GetSourceDocumentSpecOk

`func (o *UserRoleSpecification) GetSourceDocumentSpecOk() (*DocumentSpec, bool)`

GetSourceDocumentSpecOk returns a tuple with the SourceDocumentSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentSpec

`func (o *UserRoleSpecification) SetSourceDocumentSpec(v DocumentSpec)`

SetSourceDocumentSpec sets SourceDocumentSpec field to given value.

### HasSourceDocumentSpec

`func (o *UserRoleSpecification) HasSourceDocumentSpec() bool`

HasSourceDocumentSpec returns a boolean if a field has been set.

### GetPerson

`func (o *UserRoleSpecification) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *UserRoleSpecification) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *UserRoleSpecification) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *UserRoleSpecification) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetGroup

`func (o *UserRoleSpecification) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UserRoleSpecification) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UserRoleSpecification) SetGroup(v Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UserRoleSpecification) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetRole

`func (o *UserRoleSpecification) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserRoleSpecification) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserRoleSpecification) SetRole(v UserRole)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


