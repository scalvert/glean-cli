# CreateCollectionRequest

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
**NewNextItemId** | Pointer to **string** | The (optional) ItemId of the next CollectionItem in sequence. If omitted, will be added to the end of the Collection. Only used if parentId is specified. | [optional] 

## Methods

### NewCreateCollectionRequest

`func NewCreateCollectionRequest(name string, ) *CreateCollectionRequest`

NewCreateCollectionRequest instantiates a new CreateCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionRequestWithDefaults

`func NewCreateCollectionRequestWithDefaults() *CreateCollectionRequest`

NewCreateCollectionRequestWithDefaults instantiates a new CreateCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCollectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCollectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCollectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCollectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCollectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCollectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCollectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddedRoles

`func (o *CreateCollectionRequest) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *CreateCollectionRequest) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *CreateCollectionRequest) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *CreateCollectionRequest) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *CreateCollectionRequest) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *CreateCollectionRequest) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *CreateCollectionRequest) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *CreateCollectionRequest) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *CreateCollectionRequest) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *CreateCollectionRequest) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *CreateCollectionRequest) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *CreateCollectionRequest) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetIcon

`func (o *CreateCollectionRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateCollectionRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateCollectionRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateCollectionRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetAdminLocked

`func (o *CreateCollectionRequest) GetAdminLocked() bool`

GetAdminLocked returns the AdminLocked field if non-nil, zero value otherwise.

### GetAdminLockedOk

`func (o *CreateCollectionRequest) GetAdminLockedOk() (*bool, bool)`

GetAdminLockedOk returns a tuple with the AdminLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocked

`func (o *CreateCollectionRequest) SetAdminLocked(v bool)`

SetAdminLocked sets AdminLocked field to given value.

### HasAdminLocked

`func (o *CreateCollectionRequest) HasAdminLocked() bool`

HasAdminLocked returns a boolean if a field has been set.

### GetParentId

`func (o *CreateCollectionRequest) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateCollectionRequest) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateCollectionRequest) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateCollectionRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetThumbnail

`func (o *CreateCollectionRequest) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *CreateCollectionRequest) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *CreateCollectionRequest) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *CreateCollectionRequest) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetAllowedDatasource

`func (o *CreateCollectionRequest) GetAllowedDatasource() string`

GetAllowedDatasource returns the AllowedDatasource field if non-nil, zero value otherwise.

### GetAllowedDatasourceOk

`func (o *CreateCollectionRequest) GetAllowedDatasourceOk() (*string, bool)`

GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDatasource

`func (o *CreateCollectionRequest) SetAllowedDatasource(v string)`

SetAllowedDatasource sets AllowedDatasource field to given value.

### HasAllowedDatasource

`func (o *CreateCollectionRequest) HasAllowedDatasource() bool`

HasAllowedDatasource returns a boolean if a field has been set.

### GetNewNextItemId

`func (o *CreateCollectionRequest) GetNewNextItemId() string`

GetNewNextItemId returns the NewNextItemId field if non-nil, zero value otherwise.

### GetNewNextItemIdOk

`func (o *CreateCollectionRequest) GetNewNextItemIdOk() (*string, bool)`

GetNewNextItemIdOk returns a tuple with the NewNextItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNextItemId

`func (o *CreateCollectionRequest) SetNewNextItemId(v string)`

SetNewNextItemId sets NewNextItemId field to given value.

### HasNewNextItemId

`func (o *CreateCollectionRequest) HasNewNextItemId() bool`

HasNewNextItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


