# CreateCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the Collection. | 
**Description** | **string** | A brief summary of the Collection&#39;s contents. | 
**AddedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of added user roles for the Collection. | [optional] 
**RemovedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of removed user roles for the Collection. | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see this Collection. Values are taken from the corresponding filters in people search. | [optional] 
**Icon** | Pointer to **string** | The emoji icon of this Collection. | [optional] 
**AdminLocked** | Pointer to **bool** | Indicates whether edits are allowed for everyone or only admins. | [optional] 
**ParentId** | Pointer to **int32** | The parent of this Collection, or 0 if it&#39;s a top-level Collection. | [optional] 
**Thumbnail** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**AllowedDatasource** | Pointer to **string** | The datasource type this Collection can hold. | [optional] 
**Permissions** | Pointer to [**ObjectPermissions**](ObjectPermissions.md) |  | [optional] 
**Id** | **int32** | The unique ID of the Collection. | 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 
**Creator** | Pointer to [**Person**](Person.md) |  | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**ItemCount** | Pointer to **int32** | The number of items currently in the Collection. Separated from the actual items so we can grab the count without items. | [optional] 
**ChildCount** | Pointer to **int32** | The number of children Collections. Separated from the actual children so we can grab the count without children. | [optional] 
**Items** | Pointer to [**[]CollectionItem**](CollectionItem.md) | The items in this Collection. | [optional] 
**PinMetadata** | Pointer to [**CollectionPinnedMetadata**](CollectionPinnedMetadata.md) |  | [optional] 
**Shortcuts** | Pointer to **[]string** | The names of the shortcuts (Go Links) that point to this Collection. | [optional] 
**Children** | Pointer to [**[]Collection**](Collection.md) | The children Collections of this Collection. | [optional] 
**Roles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the Collection. | [optional] 
**ErrorCode** | **string** |  | 
**Collection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**Error** | Pointer to [**CollectionError**](CollectionError.md) |  | [optional] 

## Methods

### NewCreateCollectionResponse

`func NewCreateCollectionResponse(name string, description string, id int32, errorCode string, ) *CreateCollectionResponse`

NewCreateCollectionResponse instantiates a new CreateCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionResponseWithDefaults

`func NewCreateCollectionResponseWithDefaults() *CreateCollectionResponse`

NewCreateCollectionResponseWithDefaults instantiates a new CreateCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCollectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCollectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCollectionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCollectionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCollectionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCollectionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAddedRoles

`func (o *CreateCollectionResponse) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *CreateCollectionResponse) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *CreateCollectionResponse) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *CreateCollectionResponse) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *CreateCollectionResponse) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *CreateCollectionResponse) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *CreateCollectionResponse) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *CreateCollectionResponse) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *CreateCollectionResponse) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *CreateCollectionResponse) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *CreateCollectionResponse) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *CreateCollectionResponse) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetIcon

`func (o *CreateCollectionResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateCollectionResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateCollectionResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateCollectionResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetAdminLocked

`func (o *CreateCollectionResponse) GetAdminLocked() bool`

GetAdminLocked returns the AdminLocked field if non-nil, zero value otherwise.

### GetAdminLockedOk

`func (o *CreateCollectionResponse) GetAdminLockedOk() (*bool, bool)`

GetAdminLockedOk returns a tuple with the AdminLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocked

`func (o *CreateCollectionResponse) SetAdminLocked(v bool)`

SetAdminLocked sets AdminLocked field to given value.

### HasAdminLocked

`func (o *CreateCollectionResponse) HasAdminLocked() bool`

HasAdminLocked returns a boolean if a field has been set.

### GetParentId

`func (o *CreateCollectionResponse) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateCollectionResponse) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateCollectionResponse) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateCollectionResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetThumbnail

`func (o *CreateCollectionResponse) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *CreateCollectionResponse) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *CreateCollectionResponse) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *CreateCollectionResponse) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetAllowedDatasource

`func (o *CreateCollectionResponse) GetAllowedDatasource() string`

GetAllowedDatasource returns the AllowedDatasource field if non-nil, zero value otherwise.

### GetAllowedDatasourceOk

`func (o *CreateCollectionResponse) GetAllowedDatasourceOk() (*string, bool)`

GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDatasource

`func (o *CreateCollectionResponse) SetAllowedDatasource(v string)`

SetAllowedDatasource sets AllowedDatasource field to given value.

### HasAllowedDatasource

`func (o *CreateCollectionResponse) HasAllowedDatasource() bool`

HasAllowedDatasource returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateCollectionResponse) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateCollectionResponse) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateCollectionResponse) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateCollectionResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *CreateCollectionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCollectionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCollectionResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCreateTime

`func (o *CreateCollectionResponse) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CreateCollectionResponse) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CreateCollectionResponse) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CreateCollectionResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CreateCollectionResponse) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CreateCollectionResponse) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CreateCollectionResponse) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CreateCollectionResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetCreator

`func (o *CreateCollectionResponse) GetCreator() Person`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CreateCollectionResponse) GetCreatorOk() (*Person, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CreateCollectionResponse) SetCreator(v Person)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CreateCollectionResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CreateCollectionResponse) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CreateCollectionResponse) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CreateCollectionResponse) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CreateCollectionResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetItemCount

`func (o *CreateCollectionResponse) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *CreateCollectionResponse) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *CreateCollectionResponse) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *CreateCollectionResponse) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetChildCount

`func (o *CreateCollectionResponse) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *CreateCollectionResponse) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *CreateCollectionResponse) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *CreateCollectionResponse) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### GetItems

`func (o *CreateCollectionResponse) GetItems() []CollectionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateCollectionResponse) GetItemsOk() (*[]CollectionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateCollectionResponse) SetItems(v []CollectionItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *CreateCollectionResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPinMetadata

`func (o *CreateCollectionResponse) GetPinMetadata() CollectionPinnedMetadata`

GetPinMetadata returns the PinMetadata field if non-nil, zero value otherwise.

### GetPinMetadataOk

`func (o *CreateCollectionResponse) GetPinMetadataOk() (*CollectionPinnedMetadata, bool)`

GetPinMetadataOk returns a tuple with the PinMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinMetadata

`func (o *CreateCollectionResponse) SetPinMetadata(v CollectionPinnedMetadata)`

SetPinMetadata sets PinMetadata field to given value.

### HasPinMetadata

`func (o *CreateCollectionResponse) HasPinMetadata() bool`

HasPinMetadata returns a boolean if a field has been set.

### GetShortcuts

`func (o *CreateCollectionResponse) GetShortcuts() []string`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *CreateCollectionResponse) GetShortcutsOk() (*[]string, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *CreateCollectionResponse) SetShortcuts(v []string)`

SetShortcuts sets Shortcuts field to given value.

### HasShortcuts

`func (o *CreateCollectionResponse) HasShortcuts() bool`

HasShortcuts returns a boolean if a field has been set.

### GetChildren

`func (o *CreateCollectionResponse) GetChildren() []Collection`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CreateCollectionResponse) GetChildrenOk() (*[]Collection, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CreateCollectionResponse) SetChildren(v []Collection)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *CreateCollectionResponse) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRoles

`func (o *CreateCollectionResponse) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateCollectionResponse) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateCollectionResponse) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateCollectionResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetErrorCode

`func (o *CreateCollectionResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateCollectionResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateCollectionResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetCollection

`func (o *CreateCollectionResponse) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CreateCollectionResponse) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CreateCollectionResponse) SetCollection(v Collection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *CreateCollectionResponse) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetError

`func (o *CreateCollectionResponse) GetError() CollectionError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateCollectionResponse) GetErrorOk() (*CollectionError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateCollectionResponse) SetError(v CollectionError)`

SetError sets Error field to given value.

### HasError

`func (o *CreateCollectionResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


