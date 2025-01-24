# UpdateShortcutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The opaque id of the user generated content. | 
**InputAlias** | Pointer to **string** | Link text following go/ prefix as entered by the user. | [optional] 
**DestinationUrl** | Pointer to **string** | Destination URL for the shortcut. | [optional] 
**DestinationDocumentId** | Pointer to **string** | Glean Document ID for the URL, if known. | [optional] 
**Description** | Pointer to **string** | A short, plain text blurb to help people understand the intent of the shortcut. | [optional] 
**Unlisted** | Pointer to **bool** | Whether this shortcut is unlisted or not. Unlisted shortcuts are visible to author + admins only. | [optional] 
**UrlTemplate** | Pointer to **string** | For variable shortcuts, contains the URL template; note, &#x60;destinationUrl&#x60; contains default URL. | [optional] 
**AddedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles added for the Shortcut. | [optional] 
**RemovedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles removed for the Shortcut. | [optional] 

## Methods

### NewUpdateShortcutRequest

`func NewUpdateShortcutRequest(id int32, ) *UpdateShortcutRequest`

NewUpdateShortcutRequest instantiates a new UpdateShortcutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShortcutRequestWithDefaults

`func NewUpdateShortcutRequestWithDefaults() *UpdateShortcutRequest`

NewUpdateShortcutRequestWithDefaults instantiates a new UpdateShortcutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateShortcutRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateShortcutRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateShortcutRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetInputAlias

`func (o *UpdateShortcutRequest) GetInputAlias() string`

GetInputAlias returns the InputAlias field if non-nil, zero value otherwise.

### GetInputAliasOk

`func (o *UpdateShortcutRequest) GetInputAliasOk() (*string, bool)`

GetInputAliasOk returns a tuple with the InputAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAlias

`func (o *UpdateShortcutRequest) SetInputAlias(v string)`

SetInputAlias sets InputAlias field to given value.

### HasInputAlias

`func (o *UpdateShortcutRequest) HasInputAlias() bool`

HasInputAlias returns a boolean if a field has been set.

### GetDestinationUrl

`func (o *UpdateShortcutRequest) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *UpdateShortcutRequest) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *UpdateShortcutRequest) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *UpdateShortcutRequest) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### GetDestinationDocumentId

`func (o *UpdateShortcutRequest) GetDestinationDocumentId() string`

GetDestinationDocumentId returns the DestinationDocumentId field if non-nil, zero value otherwise.

### GetDestinationDocumentIdOk

`func (o *UpdateShortcutRequest) GetDestinationDocumentIdOk() (*string, bool)`

GetDestinationDocumentIdOk returns a tuple with the DestinationDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDocumentId

`func (o *UpdateShortcutRequest) SetDestinationDocumentId(v string)`

SetDestinationDocumentId sets DestinationDocumentId field to given value.

### HasDestinationDocumentId

`func (o *UpdateShortcutRequest) HasDestinationDocumentId() bool`

HasDestinationDocumentId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateShortcutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateShortcutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateShortcutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateShortcutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnlisted

`func (o *UpdateShortcutRequest) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *UpdateShortcutRequest) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *UpdateShortcutRequest) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *UpdateShortcutRequest) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *UpdateShortcutRequest) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *UpdateShortcutRequest) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *UpdateShortcutRequest) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *UpdateShortcutRequest) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.

### GetAddedRoles

`func (o *UpdateShortcutRequest) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *UpdateShortcutRequest) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *UpdateShortcutRequest) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *UpdateShortcutRequest) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *UpdateShortcutRequest) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *UpdateShortcutRequest) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *UpdateShortcutRequest) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *UpdateShortcutRequest) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


