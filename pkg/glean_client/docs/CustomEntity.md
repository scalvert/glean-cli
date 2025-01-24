# CustomEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**ObjectPermissions**](ObjectPermissions.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier. | [optional] 
**Title** | Pointer to **string** | Title or name of the custom entity. | [optional] 
**Datasource** | Pointer to **string** | The datasource the custom entity is from. | [optional] 
**ObjectType** | Pointer to **string** | The type of the entity. Interpretation is specific to each datasource | [optional] 
**Metadata** | Pointer to [**CustomEntityMetadata**](CustomEntityMetadata.md) |  | [optional] 
**Roles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the custom entity explicitly granted by the owner. | [optional] 

## Methods

### NewCustomEntity

`func NewCustomEntity() *CustomEntity`

NewCustomEntity instantiates a new CustomEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntityWithDefaults

`func NewCustomEntityWithDefaults() *CustomEntity`

NewCustomEntityWithDefaults instantiates a new CustomEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *CustomEntity) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CustomEntity) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CustomEntity) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CustomEntity) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *CustomEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *CustomEntity) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomEntity) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomEntity) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomEntity) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDatasource

`func (o *CustomEntity) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *CustomEntity) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *CustomEntity) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *CustomEntity) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetObjectType

`func (o *CustomEntity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CustomEntity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CustomEntity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CustomEntity) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomEntity) GetMetadata() CustomEntityMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomEntity) GetMetadataOk() (*CustomEntityMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomEntity) SetMetadata(v CustomEntityMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomEntity) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRoles

`func (o *CustomEntity) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CustomEntity) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CustomEntity) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CustomEntity) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


