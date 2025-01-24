# WritePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeType** | Pointer to [**ScopeType**](ScopeType.md) |  | [optional] 
**Create** | Pointer to **bool** | True if user has create permission for this feature and scope | [optional] 
**Update** | Pointer to **bool** | True if user has update permission for this feature and scope | [optional] 
**Delete** | Pointer to **bool** | True if user has delete permission for this feature and scope | [optional] 

## Methods

### NewWritePermission

`func NewWritePermission() *WritePermission`

NewWritePermission instantiates a new WritePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritePermissionWithDefaults

`func NewWritePermissionWithDefaults() *WritePermission`

NewWritePermissionWithDefaults instantiates a new WritePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeType

`func (o *WritePermission) GetScopeType() ScopeType`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *WritePermission) GetScopeTypeOk() (*ScopeType, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *WritePermission) SetScopeType(v ScopeType)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *WritePermission) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetCreate

`func (o *WritePermission) GetCreate() bool`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *WritePermission) GetCreateOk() (*bool, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *WritePermission) SetCreate(v bool)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *WritePermission) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *WritePermission) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *WritePermission) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *WritePermission) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *WritePermission) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDelete

`func (o *WritePermission) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *WritePermission) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *WritePermission) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *WritePermission) HasDelete() bool`

HasDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


