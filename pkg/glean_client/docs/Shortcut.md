# Shortcut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The opaque id of the user generated content. | [optional] 
**InputAlias** | **string** | Link text following go/ prefix as entered by the user. | 
**DestinationUrl** | Pointer to **string** | Destination URL for the shortcut. | [optional] 
**DestinationDocumentId** | Pointer to **string** | Glean Document ID for the URL, if known. | [optional] 
**Description** | Pointer to **string** | A short, plain text blurb to help people understand the intent of the shortcut. | [optional] 
**Unlisted** | Pointer to **bool** | Whether this shortcut is unlisted or not. Unlisted shortcuts are visible to author + admins only. | [optional] 
**UrlTemplate** | Pointer to **string** | For variable shortcuts, contains the URL template; note, &#x60;destinationUrl&#x60; contains default URL. | [optional] 
**AddedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles added for the Shortcut. | [optional] 
**RemovedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles removed for the Shortcut. | [optional] 
**Permissions** | Pointer to [**ObjectPermissions**](ObjectPermissions.md) |  | [optional] 
**CreatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** | The time the shortcut was created in ISO format (ISO 8601). | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time the shortcut was updated in ISO format (ISO 8601). | [optional] 
**DestinationDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**IntermediateUrl** | Pointer to **string** | The URL from which the user is then redirected to the destination URL. Full replacement for https://go/&lt;inputAlias&gt;. | [optional] 
**ViewPrefix** | Pointer to **string** | The part of the shortcut preceding the input alias when used for showing shortcuts to users. Should end with \&quot;/\&quot;. e.g. \&quot;go/\&quot; for native shortcuts. | [optional] 
**IsExternal** | Pointer to **bool** | Indicates whether a shortcut is native or external. | [optional] 
**EditUrl** | Pointer to **string** | The URL using which the user can access the edit page of the shortcut. | [optional] 
**Alias** | Pointer to **string** | canonical link text following go/ prefix where hyphen/underscore is removed. | [optional] 
**Title** | Pointer to **string** | Title for the Go Link | [optional] 
**Roles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the Go Link. | [optional] 

## Methods

### NewShortcut

`func NewShortcut(inputAlias string, ) *Shortcut`

NewShortcut instantiates a new Shortcut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortcutWithDefaults

`func NewShortcutWithDefaults() *Shortcut`

NewShortcutWithDefaults instantiates a new Shortcut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Shortcut) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Shortcut) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Shortcut) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Shortcut) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputAlias

`func (o *Shortcut) GetInputAlias() string`

GetInputAlias returns the InputAlias field if non-nil, zero value otherwise.

### GetInputAliasOk

`func (o *Shortcut) GetInputAliasOk() (*string, bool)`

GetInputAliasOk returns a tuple with the InputAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAlias

`func (o *Shortcut) SetInputAlias(v string)`

SetInputAlias sets InputAlias field to given value.


### GetDestinationUrl

`func (o *Shortcut) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *Shortcut) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *Shortcut) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *Shortcut) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### GetDestinationDocumentId

`func (o *Shortcut) GetDestinationDocumentId() string`

GetDestinationDocumentId returns the DestinationDocumentId field if non-nil, zero value otherwise.

### GetDestinationDocumentIdOk

`func (o *Shortcut) GetDestinationDocumentIdOk() (*string, bool)`

GetDestinationDocumentIdOk returns a tuple with the DestinationDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDocumentId

`func (o *Shortcut) SetDestinationDocumentId(v string)`

SetDestinationDocumentId sets DestinationDocumentId field to given value.

### HasDestinationDocumentId

`func (o *Shortcut) HasDestinationDocumentId() bool`

HasDestinationDocumentId returns a boolean if a field has been set.

### GetDescription

`func (o *Shortcut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Shortcut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Shortcut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Shortcut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnlisted

`func (o *Shortcut) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *Shortcut) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *Shortcut) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *Shortcut) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *Shortcut) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *Shortcut) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *Shortcut) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *Shortcut) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.

### GetAddedRoles

`func (o *Shortcut) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *Shortcut) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *Shortcut) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *Shortcut) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *Shortcut) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *Shortcut) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *Shortcut) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *Shortcut) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetPermissions

`func (o *Shortcut) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Shortcut) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Shortcut) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Shortcut) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Shortcut) GetCreatedBy() Person`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Shortcut) GetCreatedByOk() (*Person, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Shortcut) SetCreatedBy(v Person)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Shortcut) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *Shortcut) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Shortcut) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Shortcut) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Shortcut) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Shortcut) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Shortcut) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Shortcut) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Shortcut) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Shortcut) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Shortcut) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Shortcut) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Shortcut) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetDestinationDocument

`func (o *Shortcut) GetDestinationDocument() Document`

GetDestinationDocument returns the DestinationDocument field if non-nil, zero value otherwise.

### GetDestinationDocumentOk

`func (o *Shortcut) GetDestinationDocumentOk() (*Document, bool)`

GetDestinationDocumentOk returns a tuple with the DestinationDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDocument

`func (o *Shortcut) SetDestinationDocument(v Document)`

SetDestinationDocument sets DestinationDocument field to given value.

### HasDestinationDocument

`func (o *Shortcut) HasDestinationDocument() bool`

HasDestinationDocument returns a boolean if a field has been set.

### GetIntermediateUrl

`func (o *Shortcut) GetIntermediateUrl() string`

GetIntermediateUrl returns the IntermediateUrl field if non-nil, zero value otherwise.

### GetIntermediateUrlOk

`func (o *Shortcut) GetIntermediateUrlOk() (*string, bool)`

GetIntermediateUrlOk returns a tuple with the IntermediateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateUrl

`func (o *Shortcut) SetIntermediateUrl(v string)`

SetIntermediateUrl sets IntermediateUrl field to given value.

### HasIntermediateUrl

`func (o *Shortcut) HasIntermediateUrl() bool`

HasIntermediateUrl returns a boolean if a field has been set.

### GetViewPrefix

`func (o *Shortcut) GetViewPrefix() string`

GetViewPrefix returns the ViewPrefix field if non-nil, zero value otherwise.

### GetViewPrefixOk

`func (o *Shortcut) GetViewPrefixOk() (*string, bool)`

GetViewPrefixOk returns a tuple with the ViewPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPrefix

`func (o *Shortcut) SetViewPrefix(v string)`

SetViewPrefix sets ViewPrefix field to given value.

### HasViewPrefix

`func (o *Shortcut) HasViewPrefix() bool`

HasViewPrefix returns a boolean if a field has been set.

### GetIsExternal

`func (o *Shortcut) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *Shortcut) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *Shortcut) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *Shortcut) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetEditUrl

`func (o *Shortcut) GetEditUrl() string`

GetEditUrl returns the EditUrl field if non-nil, zero value otherwise.

### GetEditUrlOk

`func (o *Shortcut) GetEditUrlOk() (*string, bool)`

GetEditUrlOk returns a tuple with the EditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditUrl

`func (o *Shortcut) SetEditUrl(v string)`

SetEditUrl sets EditUrl field to given value.

### HasEditUrl

`func (o *Shortcut) HasEditUrl() bool`

HasEditUrl returns a boolean if a field has been set.

### GetAlias

`func (o *Shortcut) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Shortcut) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Shortcut) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Shortcut) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetTitle

`func (o *Shortcut) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Shortcut) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Shortcut) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Shortcut) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRoles

`func (o *Shortcut) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Shortcut) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Shortcut) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Shortcut) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


