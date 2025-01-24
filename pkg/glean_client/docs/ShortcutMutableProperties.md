# ShortcutMutableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputAlias** | Pointer to **string** | Link text following go/ prefix as entered by the user. | [optional] 
**DestinationUrl** | Pointer to **string** | Destination URL for the shortcut. | [optional] 
**DestinationDocumentId** | Pointer to **string** | Glean Document ID for the URL, if known. | [optional] 
**Description** | Pointer to **string** | A short, plain text blurb to help people understand the intent of the shortcut. | [optional] 
**Unlisted** | Pointer to **bool** | Whether this shortcut is unlisted or not. Unlisted shortcuts are visible to author + admins only. | [optional] 
**UrlTemplate** | Pointer to **string** | For variable shortcuts, contains the URL template; note, &#x60;destinationUrl&#x60; contains default URL. | [optional] 
**AddedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles added for the Shortcut. | [optional] 
**RemovedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles removed for the Shortcut. | [optional] 

## Methods

### NewShortcutMutableProperties

`func NewShortcutMutableProperties() *ShortcutMutableProperties`

NewShortcutMutableProperties instantiates a new ShortcutMutableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortcutMutablePropertiesWithDefaults

`func NewShortcutMutablePropertiesWithDefaults() *ShortcutMutableProperties`

NewShortcutMutablePropertiesWithDefaults instantiates a new ShortcutMutableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputAlias

`func (o *ShortcutMutableProperties) GetInputAlias() string`

GetInputAlias returns the InputAlias field if non-nil, zero value otherwise.

### GetInputAliasOk

`func (o *ShortcutMutableProperties) GetInputAliasOk() (*string, bool)`

GetInputAliasOk returns a tuple with the InputAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAlias

`func (o *ShortcutMutableProperties) SetInputAlias(v string)`

SetInputAlias sets InputAlias field to given value.

### HasInputAlias

`func (o *ShortcutMutableProperties) HasInputAlias() bool`

HasInputAlias returns a boolean if a field has been set.

### GetDestinationUrl

`func (o *ShortcutMutableProperties) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *ShortcutMutableProperties) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *ShortcutMutableProperties) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *ShortcutMutableProperties) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.

### GetDestinationDocumentId

`func (o *ShortcutMutableProperties) GetDestinationDocumentId() string`

GetDestinationDocumentId returns the DestinationDocumentId field if non-nil, zero value otherwise.

### GetDestinationDocumentIdOk

`func (o *ShortcutMutableProperties) GetDestinationDocumentIdOk() (*string, bool)`

GetDestinationDocumentIdOk returns a tuple with the DestinationDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDocumentId

`func (o *ShortcutMutableProperties) SetDestinationDocumentId(v string)`

SetDestinationDocumentId sets DestinationDocumentId field to given value.

### HasDestinationDocumentId

`func (o *ShortcutMutableProperties) HasDestinationDocumentId() bool`

HasDestinationDocumentId returns a boolean if a field has been set.

### GetDescription

`func (o *ShortcutMutableProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShortcutMutableProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShortcutMutableProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShortcutMutableProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnlisted

`func (o *ShortcutMutableProperties) GetUnlisted() bool`

GetUnlisted returns the Unlisted field if non-nil, zero value otherwise.

### GetUnlistedOk

`func (o *ShortcutMutableProperties) GetUnlistedOk() (*bool, bool)`

GetUnlistedOk returns a tuple with the Unlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlisted

`func (o *ShortcutMutableProperties) SetUnlisted(v bool)`

SetUnlisted sets Unlisted field to given value.

### HasUnlisted

`func (o *ShortcutMutableProperties) HasUnlisted() bool`

HasUnlisted returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *ShortcutMutableProperties) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *ShortcutMutableProperties) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *ShortcutMutableProperties) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *ShortcutMutableProperties) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.

### GetAddedRoles

`func (o *ShortcutMutableProperties) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *ShortcutMutableProperties) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *ShortcutMutableProperties) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *ShortcutMutableProperties) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *ShortcutMutableProperties) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *ShortcutMutableProperties) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *ShortcutMutableProperties) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *ShortcutMutableProperties) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


