# Collection

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

## Methods

### NewCollection

`func NewCollection(name string, description string, id int32, ) *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Collection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collection) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Collection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Collection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Collection) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAddedRoles

`func (o *Collection) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *Collection) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *Collection) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *Collection) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *Collection) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *Collection) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *Collection) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *Collection) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *Collection) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *Collection) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *Collection) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *Collection) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetIcon

`func (o *Collection) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Collection) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Collection) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Collection) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetAdminLocked

`func (o *Collection) GetAdminLocked() bool`

GetAdminLocked returns the AdminLocked field if non-nil, zero value otherwise.

### GetAdminLockedOk

`func (o *Collection) GetAdminLockedOk() (*bool, bool)`

GetAdminLockedOk returns a tuple with the AdminLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLocked

`func (o *Collection) SetAdminLocked(v bool)`

SetAdminLocked sets AdminLocked field to given value.

### HasAdminLocked

`func (o *Collection) HasAdminLocked() bool`

HasAdminLocked returns a boolean if a field has been set.

### GetParentId

`func (o *Collection) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Collection) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Collection) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Collection) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetThumbnail

`func (o *Collection) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Collection) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Collection) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Collection) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetAllowedDatasource

`func (o *Collection) GetAllowedDatasource() string`

GetAllowedDatasource returns the AllowedDatasource field if non-nil, zero value otherwise.

### GetAllowedDatasourceOk

`func (o *Collection) GetAllowedDatasourceOk() (*string, bool)`

GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDatasource

`func (o *Collection) SetAllowedDatasource(v string)`

SetAllowedDatasource sets AllowedDatasource field to given value.

### HasAllowedDatasource

`func (o *Collection) HasAllowedDatasource() bool`

HasAllowedDatasource returns a boolean if a field has been set.

### GetPermissions

`func (o *Collection) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Collection) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Collection) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Collection) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *Collection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Collection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Collection) SetId(v int32)`

SetId sets Id field to given value.


### GetCreateTime

`func (o *Collection) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Collection) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Collection) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Collection) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Collection) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Collection) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Collection) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Collection) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetCreator

`func (o *Collection) GetCreator() Person`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Collection) GetCreatorOk() (*Person, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Collection) SetCreator(v Person)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Collection) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Collection) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Collection) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Collection) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Collection) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetItemCount

`func (o *Collection) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *Collection) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *Collection) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *Collection) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetChildCount

`func (o *Collection) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *Collection) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *Collection) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *Collection) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### GetItems

`func (o *Collection) GetItems() []CollectionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Collection) GetItemsOk() (*[]CollectionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Collection) SetItems(v []CollectionItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Collection) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPinMetadata

`func (o *Collection) GetPinMetadata() CollectionPinnedMetadata`

GetPinMetadata returns the PinMetadata field if non-nil, zero value otherwise.

### GetPinMetadataOk

`func (o *Collection) GetPinMetadataOk() (*CollectionPinnedMetadata, bool)`

GetPinMetadataOk returns a tuple with the PinMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinMetadata

`func (o *Collection) SetPinMetadata(v CollectionPinnedMetadata)`

SetPinMetadata sets PinMetadata field to given value.

### HasPinMetadata

`func (o *Collection) HasPinMetadata() bool`

HasPinMetadata returns a boolean if a field has been set.

### GetShortcuts

`func (o *Collection) GetShortcuts() []string`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *Collection) GetShortcutsOk() (*[]string, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *Collection) SetShortcuts(v []string)`

SetShortcuts sets Shortcuts field to given value.

### HasShortcuts

`func (o *Collection) HasShortcuts() bool`

HasShortcuts returns a boolean if a field has been set.

### GetChildren

`func (o *Collection) GetChildren() []Collection`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Collection) GetChildrenOk() (*[]Collection, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Collection) SetChildren(v []Collection)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Collection) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetRoles

`func (o *Collection) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Collection) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Collection) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Collection) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


