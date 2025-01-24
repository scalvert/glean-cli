# EditCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the Collection. | 
**Description** | Pointer to **string** | A brief summary of the Collection&#39;s contents. | [optional] 
**AddedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of added user roles for the Collection. | [optional] 
**RemovedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of removed user roles for the Collection. | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see this Collection. Values are taken from the corresponding filters in people search. | [optional] 
**Icon** | Pointer to **string** | The emoji icon of this Collection. | [optional] 
**AdminLocked** | Pointer to **bool** | Indicates whether edits are allowed for everyone or only admins. | [optional] 
**ParentId** | Pointer to **int32** | The parent of this Collection, or 0 if it&#39;s a top-level Collection. | [optional] 
**Thumbnail** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**AllowedDatasource** | Pointer to **string** | The datasource type this Collection can hold. | [optional] 
**Id** | **int32** | The ID of the Collection to modify. | 

## Methods

### NewEditCollectionRequest

`func NewEditCollectionRequest(name string, id int32, ) *EditCollectionRequest`

NewEditCollectionRequest instantiates a new EditCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditCollectionRequestWithDefaults

`func NewEditCollectionRequestWithDefaults() *EditCollectionRequest`

NewEditCollectionRequestWithDefaults instantiates a new EditCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditCollectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditCollectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditCollectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EditCollectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditCollectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditCollectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditCollectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddedRoles

`func (o *EditCollectionRequest) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *EditCollectionRequest) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *EditCollectionRequest) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *EditCollectionRequest) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *EditCollectionRequest) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *EditCollectionRequest) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *EditCollectionRequest) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *EditCollectionRequest) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *EditCollectionRequest) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *EditCollectionRequest) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *EditCollectionRequest) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *EditCollectionRequest) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetIcon

`func (o *EditCollectionRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EditCollectionRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EditCollectionRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EditCollectionRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetAdminLocked

`func (o *EditCollectionRequest) GetAdminLocked() bool`

GetAdminLocked returns the AdminLocked field if non-nil, zero value otherwise.

### GetAdminLockedOk

`func (o *EditCollectionRequest) GetAdminLockedOk() (*bool, bool)`

GetAdminLockedOk returns a tuple with the AdminLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocked

`func (o *EditCollectionRequest) SetAdminLocked(v bool)`

SetAdminLocked sets AdminLocked field to given value.

### HasAdminLocked

`func (o *EditCollectionRequest) HasAdminLocked() bool`

HasAdminLocked returns a boolean if a field has been set.

### GetParentId

`func (o *EditCollectionRequest) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *EditCollectionRequest) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *EditCollectionRequest) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *EditCollectionRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetThumbnail

`func (o *EditCollectionRequest) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *EditCollectionRequest) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *EditCollectionRequest) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *EditCollectionRequest) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetAllowedDatasource

`func (o *EditCollectionRequest) GetAllowedDatasource() string`

GetAllowedDatasource returns the AllowedDatasource field if non-nil, zero value otherwise.

### GetAllowedDatasourceOk

`func (o *EditCollectionRequest) GetAllowedDatasourceOk() (*string, bool)`

GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDatasource

`func (o *EditCollectionRequest) SetAllowedDatasource(v string)`

SetAllowedDatasource sets AllowedDatasource field to given value.

### HasAllowedDatasource

`func (o *EditCollectionRequest) HasAllowedDatasource() bool`

HasAllowedDatasource returns a boolean if a field has been set.

### GetId

`func (o *EditCollectionRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditCollectionRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditCollectionRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


