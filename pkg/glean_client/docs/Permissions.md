# Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanAdminSearch** | Pointer to **bool** | TODO--deprecate in favor of the read and write properties. True if the user has access to /adminsearch | [optional] 
**CanAdminClientApiGlobalTokens** | Pointer to **bool** | TODO--deprecate in favor of the read and write properties. True if the user can administrate client API tokens with global scope | [optional] 
**CanDlp** | Pointer to **bool** | TODO--deprecate in favor of the read and write properties. True if the user has access to data loss prevention (DLP) features | [optional] 
**Read** | Pointer to [**map[string][]ReadPermission**](array.md) | Describes the read permission levels that a user has for permissioned features. Key must be PermissionedFeatureOrObject | [optional] 
**Write** | Pointer to [**map[string][]WritePermission**](array.md) | Describes the write permissions levels that a user has for permissioned features. Key must be PermissionedFeatureOrObject | [optional] 
**Grant** | Pointer to [**map[string][]GrantPermission**](array.md) | Describes the grant permission levels that a user has for permissioned features. Key must be PermissionedFeatureOrObject | [optional] 
**Role** | Pointer to **string** | The roleId of the canonical role a user has. The displayName is equal to the roleId. | [optional] 
**Roles** | Pointer to **[]string** | The roleIds of the roles a user has. | [optional] 

## Methods

### NewPermissions

`func NewPermissions() *Permissions`

NewPermissions instantiates a new Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsWithDefaults

`func NewPermissionsWithDefaults() *Permissions`

NewPermissionsWithDefaults instantiates a new Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanAdminSearch

`func (o *Permissions) GetCanAdminSearch() bool`

GetCanAdminSearch returns the CanAdminSearch field if non-nil, zero value otherwise.

### GetCanAdminSearchOk

`func (o *Permissions) GetCanAdminSearchOk() (*bool, bool)`

GetCanAdminSearchOk returns a tuple with the CanAdminSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdminSearch

`func (o *Permissions) SetCanAdminSearch(v bool)`

SetCanAdminSearch sets CanAdminSearch field to given value.

### HasCanAdminSearch

`func (o *Permissions) HasCanAdminSearch() bool`

HasCanAdminSearch returns a boolean if a field has been set.

### GetCanAdminClientApiGlobalTokens

`func (o *Permissions) GetCanAdminClientApiGlobalTokens() bool`

GetCanAdminClientApiGlobalTokens returns the CanAdminClientApiGlobalTokens field if non-nil, zero value otherwise.

### GetCanAdminClientApiGlobalTokensOk

`func (o *Permissions) GetCanAdminClientApiGlobalTokensOk() (*bool, bool)`

GetCanAdminClientApiGlobalTokensOk returns a tuple with the CanAdminClientApiGlobalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdminClientApiGlobalTokens

`func (o *Permissions) SetCanAdminClientApiGlobalTokens(v bool)`

SetCanAdminClientApiGlobalTokens sets CanAdminClientApiGlobalTokens field to given value.

### HasCanAdminClientApiGlobalTokens

`func (o *Permissions) HasCanAdminClientApiGlobalTokens() bool`

HasCanAdminClientApiGlobalTokens returns a boolean if a field has been set.

### GetCanDlp

`func (o *Permissions) GetCanDlp() bool`

GetCanDlp returns the CanDlp field if non-nil, zero value otherwise.

### GetCanDlpOk

`func (o *Permissions) GetCanDlpOk() (*bool, bool)`

GetCanDlpOk returns a tuple with the CanDlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDlp

`func (o *Permissions) SetCanDlp(v bool)`

SetCanDlp sets CanDlp field to given value.

### HasCanDlp

`func (o *Permissions) HasCanDlp() bool`

HasCanDlp returns a boolean if a field has been set.

### GetRead

`func (o *Permissions) GetRead() map[string][]ReadPermission`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *Permissions) GetReadOk() (*map[string][]ReadPermission, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *Permissions) SetRead(v map[string][]ReadPermission)`

SetRead sets Read field to given value.

### HasRead

`func (o *Permissions) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetWrite

`func (o *Permissions) GetWrite() map[string][]WritePermission`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *Permissions) GetWriteOk() (*map[string][]WritePermission, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *Permissions) SetWrite(v map[string][]WritePermission)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *Permissions) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetGrant

`func (o *Permissions) GetGrant() map[string][]GrantPermission`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *Permissions) GetGrantOk() (*map[string][]GrantPermission, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *Permissions) SetGrant(v map[string][]GrantPermission)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *Permissions) HasGrant() bool`

HasGrant returns a boolean if a field has been set.

### GetRole

`func (o *Permissions) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Permissions) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Permissions) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Permissions) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoles

`func (o *Permissions) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Permissions) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Permissions) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Permissions) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


