# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelatedObjects** | Pointer to [**map[string]RelatedObjectEdge**](RelatedObjectEdge.md) | A list of objects related to a source object. | [optional] 
**Permissions** | Pointer to [**ObjectPermissions**](ObjectPermissions.md) |  | [optional] 
**Id** | **string** | Unique identifier | 
**Name** | **string** | Team name | 
**Description** | Pointer to **string** | A description of the team | [optional] 
**BusinessUnit** | Pointer to **string** | Typically the highest level organizational unit; generally applies to bigger companies with multiple distinct businesses. | [optional] 
**Department** | Pointer to **string** | An organizational unit where everyone has a similar task, e.g. &#x60;Engineering&#x60;. | [optional] 
**PhotoUrl** | Pointer to **string** | A link to the team&#39;s photo. | [optional] 
**BannerUrl** | Pointer to **string** | A link to the team&#39;s banner photo. | [optional] 
**ExternalLink** | Pointer to **string** | Link to a team page on the internet or your company&#39;s intranet | [optional] 
**Members** | Pointer to [**[]PersonToTeamRelationship**](PersonToTeamRelationship.md) | The members on this team | [optional] 
**MemberCount** | Pointer to **int32** | Number of members on this team (recursive; includes all individuals that belong to this team, and all individuals that belong to a subteam within this team) | [optional] 
**Emails** | Pointer to [**[]TeamEmail**](TeamEmail.md) | The emails for this team | [optional] 
**DatasourceProfiles** | Pointer to [**[]DatasourceProfile**](DatasourceProfile.md) | The datasource profiles of the team | [optional] 
**Datasource** | Pointer to **string** | the data source of the team, e.g. GDRIVE | [optional] 
**CreatedFrom** | Pointer to **string** | For teams created from docs, the doc title. Otherwise empty. | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | when this team was last updated. | [optional] 
**Status** | Pointer to **string** | whether this team is fully processed or there are still unprocessed operations that&#39;ll affect it | [optional] [default to "PROCESSED"]
**CanBeDeleted** | Pointer to **bool** | can this team be deleted. Some manually ingested teams like GCS_CSV or PUSH_API cannot | [optional] [default to true]
**LoggingId** | Pointer to **string** | The logging id of the team used in scrubbed logs, client analytics, and metrics. | [optional] 

## Methods

### NewTeam

`func NewTeam(id string, name string, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelatedObjects

`func (o *Team) GetRelatedObjects() map[string]RelatedObjectEdge`

GetRelatedObjects returns the RelatedObjects field if non-nil, zero value otherwise.

### GetRelatedObjectsOk

`func (o *Team) GetRelatedObjectsOk() (*map[string]RelatedObjectEdge, bool)`

GetRelatedObjectsOk returns a tuple with the RelatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjects

`func (o *Team) SetRelatedObjects(v map[string]RelatedObjectEdge)`

SetRelatedObjects sets RelatedObjects field to given value.

### HasRelatedObjects

`func (o *Team) HasRelatedObjects() bool`

HasRelatedObjects returns a boolean if a field has been set.

### GetPermissions

`func (o *Team) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Team) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Team) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Team) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *Team) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Team) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Team) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *Team) GetBusinessUnit() string`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *Team) GetBusinessUnitOk() (*string, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *Team) SetBusinessUnit(v string)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *Team) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetDepartment

`func (o *Team) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Team) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Team) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Team) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *Team) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *Team) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *Team) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *Team) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetBannerUrl

`func (o *Team) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *Team) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *Team) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *Team) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### GetExternalLink

`func (o *Team) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *Team) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *Team) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *Team) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetMembers

`func (o *Team) GetMembers() []PersonToTeamRelationship`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Team) GetMembersOk() (*[]PersonToTeamRelationship, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Team) SetMembers(v []PersonToTeamRelationship)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Team) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMemberCount

`func (o *Team) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *Team) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *Team) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *Team) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetEmails

`func (o *Team) GetEmails() []TeamEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *Team) GetEmailsOk() (*[]TeamEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *Team) SetEmails(v []TeamEmail)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *Team) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetDatasourceProfiles

`func (o *Team) GetDatasourceProfiles() []DatasourceProfile`

GetDatasourceProfiles returns the DatasourceProfiles field if non-nil, zero value otherwise.

### GetDatasourceProfilesOk

`func (o *Team) GetDatasourceProfilesOk() (*[]DatasourceProfile, bool)`

GetDatasourceProfilesOk returns a tuple with the DatasourceProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceProfiles

`func (o *Team) SetDatasourceProfiles(v []DatasourceProfile)`

SetDatasourceProfiles sets DatasourceProfiles field to given value.

### HasDatasourceProfiles

`func (o *Team) HasDatasourceProfiles() bool`

HasDatasourceProfiles returns a boolean if a field has been set.

### GetDatasource

`func (o *Team) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *Team) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *Team) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *Team) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetCreatedFrom

`func (o *Team) GetCreatedFrom() string`

GetCreatedFrom returns the CreatedFrom field if non-nil, zero value otherwise.

### GetCreatedFromOk

`func (o *Team) GetCreatedFromOk() (*string, bool)`

GetCreatedFromOk returns a tuple with the CreatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFrom

`func (o *Team) SetCreatedFrom(v string)`

SetCreatedFrom sets CreatedFrom field to given value.

### HasCreatedFrom

`func (o *Team) HasCreatedFrom() bool`

HasCreatedFrom returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *Team) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *Team) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *Team) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *Team) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Team) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Team) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Team) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Team) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCanBeDeleted

`func (o *Team) GetCanBeDeleted() bool`

GetCanBeDeleted returns the CanBeDeleted field if non-nil, zero value otherwise.

### GetCanBeDeletedOk

`func (o *Team) GetCanBeDeletedOk() (*bool, bool)`

GetCanBeDeletedOk returns a tuple with the CanBeDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeDeleted

`func (o *Team) SetCanBeDeleted(v bool)`

SetCanBeDeleted sets CanBeDeleted field to given value.

### HasCanBeDeleted

`func (o *Team) HasCanBeDeleted() bool`

HasCanBeDeleted returns a boolean if a field has been set.

### GetLoggingId

`func (o *Team) GetLoggingId() string`

GetLoggingId returns the LoggingId field if non-nil, zero value otherwise.

### GetLoggingIdOk

`func (o *Team) GetLoggingIdOk() (*string, bool)`

GetLoggingIdOk returns a tuple with the LoggingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingId

`func (o *Team) SetLoggingId(v string)`

SetLoggingId sets LoggingId field to given value.

### HasLoggingId

`func (o *Team) HasLoggingId() bool`

HasLoggingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


