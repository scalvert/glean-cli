# DocumentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | Pointer to **string** |  | [optional] 
**DatasourceInstance** | Pointer to **string** | The datasource instance from which the document was extracted. | [optional] 
**ObjectType** | Pointer to **string** | The type of the result. Interpretation is specific to each datasource. (e.g. for Jira issues, this is the issue type such as Bug or Feature Request). | [optional] 
**Container** | Pointer to **string** | The name of the container (higher level parent, not direct parent) of the result. Interpretation is specific to each datasource (e.g. Channels for Slack, Project for Jira). cf. parentId | [optional] 
**ContainerId** | Pointer to **string** | The Glean Document ID of the container. Uniquely identifies the container. | [optional] 
**SuperContainerId** | Pointer to **string** | The Glean Document ID of the super container. Super container represents a broader abstraction that contains many containers. For example, whereas container might refer to a folder, super container would refer to a drive. | [optional] 
**ParentId** | Pointer to **string** | The id of the direct parent of the result. Interpretation is specific to each datasource (e.g. parent issue for Jira). cf. container | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** | The index-wide unique identifier. | [optional] 
**DocumentIdHash** | Pointer to **string** | Hash of the Glean Document ID. | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 
**Author** | Pointer to [**Person**](Person.md) |  | [optional] 
**Owner** | Pointer to [**Person**](Person.md) |  | [optional] 
**Visibility** | Pointer to [**DocumentVisibility**](DocumentVisibility.md) |  | [optional] 
**Components** | Pointer to **[]string** | A list of components this result is associated with. Interpretation is specific to each datasource. (e.g. for Jira issues, these are [components](https://confluence.atlassian.com/jirasoftwarecloud/organizing-work-with-components-764478279.html).) | [optional] 
**Status** | Pointer to **string** | The status or disposition of the result. Interpretation is specific to each datasource. (e.g. for Jira issues, this is the issue status such as Done, In Progress or Will Not Fix). | [optional] 
**StatusCategory** | Pointer to **string** | The status category of the result. Meant to be more general than status. Interpretation is specific to each datasource. | [optional] 
**Pins** | Pointer to [**[]PinDocument**](PinDocument.md) | A list of stars associated with this result.  \&quot;Pin\&quot; is an older name. | [optional] 
**Priority** | Pointer to **string** | The document priority. Interpretation is datasource specific. | [optional] 
**AssignedTo** | Pointer to [**Person**](Person.md) |  | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**Labels** | Pointer to **[]string** | A list of tags for the document. Interpretation is datasource specific. | [optional] 
**Collections** | Pointer to [**[]Collection**](Collection.md) | A list of collections that the document belongs to. | [optional] 
**DatasourceId** | Pointer to **string** | The user-visible datasource specific id (e.g. Salesforce case number for example, GitHub PR number). | [optional] 
**Interactions** | Pointer to [**DocumentInteractions**](DocumentInteractions.md) |  | [optional] 
**Verification** | Pointer to [**Verification**](Verification.md) |  | [optional] 
**ViewerInfo** | Pointer to [**ViewerInfo**](ViewerInfo.md) |  | [optional] 
**Permissions** | Pointer to [**ObjectPermissions**](ObjectPermissions.md) |  | [optional] 
**VisitCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 
**Shortcuts** | Pointer to [**[]Shortcut**](Shortcut.md) | A list of shortcuts of which destination URL is for the document. | [optional] 
**Path** | Pointer to **string** | For file datasources like onedrive/github etc this has the path to the file | [optional] 
**CustomData** | Pointer to [**map[string]CustomDataValue**](CustomDataValue.md) | Custom fields specific to individual datasources | [optional] 
**DocumentCategory** | Pointer to **string** | The document&#39;s document_category(.proto). | [optional] 
**ContactPerson** | Pointer to [**Person**](Person.md) |  | [optional] 
**Thumbnail** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**IndexStatus** | Pointer to [**IndexStatus**](IndexStatus.md) |  | [optional] 
**Ancestors** | Pointer to [**[]Document**](Document.md) | A list of documents that are ancestors of this document in the hierarchy of the document&#39;s datasource, for example parent folders or containers. Ancestors can be of different types and some may not be indexed. Higher level ancestors appear earlier in the list. | [optional] 

## Methods

### NewDocumentMetadata

`func NewDocumentMetadata() *DocumentMetadata`

NewDocumentMetadata instantiates a new DocumentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentMetadataWithDefaults

`func NewDocumentMetadataWithDefaults() *DocumentMetadata`

NewDocumentMetadataWithDefaults instantiates a new DocumentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *DocumentMetadata) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *DocumentMetadata) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *DocumentMetadata) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *DocumentMetadata) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetDatasourceInstance

`func (o *DocumentMetadata) GetDatasourceInstance() string`

GetDatasourceInstance returns the DatasourceInstance field if non-nil, zero value otherwise.

### GetDatasourceInstanceOk

`func (o *DocumentMetadata) GetDatasourceInstanceOk() (*string, bool)`

GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstance

`func (o *DocumentMetadata) SetDatasourceInstance(v string)`

SetDatasourceInstance sets DatasourceInstance field to given value.

### HasDatasourceInstance

`func (o *DocumentMetadata) HasDatasourceInstance() bool`

HasDatasourceInstance returns a boolean if a field has been set.

### GetObjectType

`func (o *DocumentMetadata) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DocumentMetadata) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DocumentMetadata) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *DocumentMetadata) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetContainer

`func (o *DocumentMetadata) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *DocumentMetadata) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *DocumentMetadata) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *DocumentMetadata) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerId

`func (o *DocumentMetadata) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *DocumentMetadata) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *DocumentMetadata) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *DocumentMetadata) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetSuperContainerId

`func (o *DocumentMetadata) GetSuperContainerId() string`

GetSuperContainerId returns the SuperContainerId field if non-nil, zero value otherwise.

### GetSuperContainerIdOk

`func (o *DocumentMetadata) GetSuperContainerIdOk() (*string, bool)`

GetSuperContainerIdOk returns a tuple with the SuperContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperContainerId

`func (o *DocumentMetadata) SetSuperContainerId(v string)`

SetSuperContainerId sets SuperContainerId field to given value.

### HasSuperContainerId

`func (o *DocumentMetadata) HasSuperContainerId() bool`

HasSuperContainerId returns a boolean if a field has been set.

### GetParentId

`func (o *DocumentMetadata) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DocumentMetadata) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DocumentMetadata) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DocumentMetadata) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetMimeType

`func (o *DocumentMetadata) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *DocumentMetadata) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *DocumentMetadata) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *DocumentMetadata) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetDocumentId

`func (o *DocumentMetadata) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DocumentMetadata) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DocumentMetadata) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *DocumentMetadata) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentIdHash

`func (o *DocumentMetadata) GetDocumentIdHash() string`

GetDocumentIdHash returns the DocumentIdHash field if non-nil, zero value otherwise.

### GetDocumentIdHashOk

`func (o *DocumentMetadata) GetDocumentIdHashOk() (*string, bool)`

GetDocumentIdHashOk returns a tuple with the DocumentIdHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIdHash

`func (o *DocumentMetadata) SetDocumentIdHash(v string)`

SetDocumentIdHash sets DocumentIdHash field to given value.

### HasDocumentIdHash

`func (o *DocumentMetadata) HasDocumentIdHash() bool`

HasDocumentIdHash returns a boolean if a field has been set.

### GetCreateTime

`func (o *DocumentMetadata) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DocumentMetadata) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DocumentMetadata) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DocumentMetadata) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *DocumentMetadata) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *DocumentMetadata) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *DocumentMetadata) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *DocumentMetadata) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetAuthor

`func (o *DocumentMetadata) GetAuthor() Person`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *DocumentMetadata) GetAuthorOk() (*Person, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *DocumentMetadata) SetAuthor(v Person)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *DocumentMetadata) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetOwner

`func (o *DocumentMetadata) GetOwner() Person`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DocumentMetadata) GetOwnerOk() (*Person, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DocumentMetadata) SetOwner(v Person)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DocumentMetadata) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetVisibility

`func (o *DocumentMetadata) GetVisibility() DocumentVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *DocumentMetadata) GetVisibilityOk() (*DocumentVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *DocumentMetadata) SetVisibility(v DocumentVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *DocumentMetadata) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetComponents

`func (o *DocumentMetadata) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DocumentMetadata) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DocumentMetadata) SetComponents(v []string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *DocumentMetadata) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *DocumentMetadata) GetStatusCategory() string`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *DocumentMetadata) GetStatusCategoryOk() (*string, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *DocumentMetadata) SetStatusCategory(v string)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *DocumentMetadata) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetPins

`func (o *DocumentMetadata) GetPins() []PinDocument`

GetPins returns the Pins field if non-nil, zero value otherwise.

### GetPinsOk

`func (o *DocumentMetadata) GetPinsOk() (*[]PinDocument, bool)`

GetPinsOk returns a tuple with the Pins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPins

`func (o *DocumentMetadata) SetPins(v []PinDocument)`

SetPins sets Pins field to given value.

### HasPins

`func (o *DocumentMetadata) HasPins() bool`

HasPins returns a boolean if a field has been set.

### GetPriority

`func (o *DocumentMetadata) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DocumentMetadata) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DocumentMetadata) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DocumentMetadata) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetAssignedTo

`func (o *DocumentMetadata) GetAssignedTo() Person`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *DocumentMetadata) GetAssignedToOk() (*Person, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *DocumentMetadata) SetAssignedTo(v Person)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *DocumentMetadata) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DocumentMetadata) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DocumentMetadata) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DocumentMetadata) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DocumentMetadata) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetLabels

`func (o *DocumentMetadata) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *DocumentMetadata) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *DocumentMetadata) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *DocumentMetadata) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCollections

`func (o *DocumentMetadata) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *DocumentMetadata) GetCollectionsOk() (*[]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *DocumentMetadata) SetCollections(v []Collection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *DocumentMetadata) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetDatasourceId

`func (o *DocumentMetadata) GetDatasourceId() string`

GetDatasourceId returns the DatasourceId field if non-nil, zero value otherwise.

### GetDatasourceIdOk

`func (o *DocumentMetadata) GetDatasourceIdOk() (*string, bool)`

GetDatasourceIdOk returns a tuple with the DatasourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceId

`func (o *DocumentMetadata) SetDatasourceId(v string)`

SetDatasourceId sets DatasourceId field to given value.

### HasDatasourceId

`func (o *DocumentMetadata) HasDatasourceId() bool`

HasDatasourceId returns a boolean if a field has been set.

### GetInteractions

`func (o *DocumentMetadata) GetInteractions() DocumentInteractions`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *DocumentMetadata) GetInteractionsOk() (*DocumentInteractions, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *DocumentMetadata) SetInteractions(v DocumentInteractions)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *DocumentMetadata) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.

### GetVerification

`func (o *DocumentMetadata) GetVerification() Verification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *DocumentMetadata) GetVerificationOk() (*Verification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *DocumentMetadata) SetVerification(v Verification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *DocumentMetadata) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetViewerInfo

`func (o *DocumentMetadata) GetViewerInfo() ViewerInfo`

GetViewerInfo returns the ViewerInfo field if non-nil, zero value otherwise.

### GetViewerInfoOk

`func (o *DocumentMetadata) GetViewerInfoOk() (*ViewerInfo, bool)`

GetViewerInfoOk returns a tuple with the ViewerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerInfo

`func (o *DocumentMetadata) SetViewerInfo(v ViewerInfo)`

SetViewerInfo sets ViewerInfo field to given value.

### HasViewerInfo

`func (o *DocumentMetadata) HasViewerInfo() bool`

HasViewerInfo returns a boolean if a field has been set.

### GetPermissions

`func (o *DocumentMetadata) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DocumentMetadata) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DocumentMetadata) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DocumentMetadata) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetVisitCount

`func (o *DocumentMetadata) GetVisitCount() CountInfo`

GetVisitCount returns the VisitCount field if non-nil, zero value otherwise.

### GetVisitCountOk

`func (o *DocumentMetadata) GetVisitCountOk() (*CountInfo, bool)`

GetVisitCountOk returns a tuple with the VisitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitCount

`func (o *DocumentMetadata) SetVisitCount(v CountInfo)`

SetVisitCount sets VisitCount field to given value.

### HasVisitCount

`func (o *DocumentMetadata) HasVisitCount() bool`

HasVisitCount returns a boolean if a field has been set.

### GetShortcuts

`func (o *DocumentMetadata) GetShortcuts() []Shortcut`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *DocumentMetadata) GetShortcutsOk() (*[]Shortcut, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *DocumentMetadata) SetShortcuts(v []Shortcut)`

SetShortcuts sets Shortcuts field to given value.

### HasShortcuts

`func (o *DocumentMetadata) HasShortcuts() bool`

HasShortcuts returns a boolean if a field has been set.

### GetPath

`func (o *DocumentMetadata) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DocumentMetadata) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DocumentMetadata) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DocumentMetadata) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCustomData

`func (o *DocumentMetadata) GetCustomData() map[string]CustomDataValue`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *DocumentMetadata) GetCustomDataOk() (*map[string]CustomDataValue, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *DocumentMetadata) SetCustomData(v map[string]CustomDataValue)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *DocumentMetadata) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetDocumentCategory

`func (o *DocumentMetadata) GetDocumentCategory() string`

GetDocumentCategory returns the DocumentCategory field if non-nil, zero value otherwise.

### GetDocumentCategoryOk

`func (o *DocumentMetadata) GetDocumentCategoryOk() (*string, bool)`

GetDocumentCategoryOk returns a tuple with the DocumentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCategory

`func (o *DocumentMetadata) SetDocumentCategory(v string)`

SetDocumentCategory sets DocumentCategory field to given value.

### HasDocumentCategory

`func (o *DocumentMetadata) HasDocumentCategory() bool`

HasDocumentCategory returns a boolean if a field has been set.

### GetContactPerson

`func (o *DocumentMetadata) GetContactPerson() Person`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *DocumentMetadata) GetContactPersonOk() (*Person, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *DocumentMetadata) SetContactPerson(v Person)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *DocumentMetadata) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetThumbnail

`func (o *DocumentMetadata) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *DocumentMetadata) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *DocumentMetadata) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *DocumentMetadata) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetIndexStatus

`func (o *DocumentMetadata) GetIndexStatus() IndexStatus`

GetIndexStatus returns the IndexStatus field if non-nil, zero value otherwise.

### GetIndexStatusOk

`func (o *DocumentMetadata) GetIndexStatusOk() (*IndexStatus, bool)`

GetIndexStatusOk returns a tuple with the IndexStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStatus

`func (o *DocumentMetadata) SetIndexStatus(v IndexStatus)`

SetIndexStatus sets IndexStatus field to given value.

### HasIndexStatus

`func (o *DocumentMetadata) HasIndexStatus() bool`

HasIndexStatus returns a boolean if a field has been set.

### GetAncestors

`func (o *DocumentMetadata) GetAncestors() []Document`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *DocumentMetadata) GetAncestorsOk() (*[]Document, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *DocumentMetadata) SetAncestors(v []Document)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *DocumentMetadata) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


