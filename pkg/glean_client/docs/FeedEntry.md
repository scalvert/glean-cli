# FeedEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | Pointer to **string** | optional ID associated with a single feed entry (displayable_list_id) | [optional] 
**Title** | **string** | Title for the result. Can be document title, event title and so on. | 
**Thumbnail** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**CreatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**UiConfig** | Pointer to [**FeedEntryUiConfig**](FeedEntryUiConfig.md) |  | [optional] 
**Snippet** | Pointer to **string** | A textual snippet representing this entry, dependent on type. For example, for USER_MENTION, it may contain the sentence in which the mention occurred. | [optional] 
**JustificationType** | Pointer to **string** | Type of the justification. | [optional] 
**Justification** | Pointer to **string** | Server side generated justification string if server provides one. | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular feed entry in this particular response. To be used for /feedback reporting. | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Event** | Pointer to [**CalendarEvent**](CalendarEvent.md) |  | [optional] 
**Announcement** | Pointer to [**Announcement**](Announcement.md) |  | [optional] 
**Collection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**CollectionItem** | Pointer to [**CollectionItem**](CollectionItem.md) |  | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**App** | Pointer to [**AppResult**](AppResult.md) |  | [optional] 
**Activities** | Pointer to [**[]UserActivity**](UserActivity.md) | List of activity where each activity has user, action, timestamp. | [optional] 
**DocumentVisitorCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 
**ViewUrl** | Pointer to **string** | View URL for the entry if based on links that are not documents in Glean. | [optional] 
**AdditionalClientActions** | Pointer to [**[]ClientAction**](ClientAction.md) | List of client actions suggested by the backend to be included for entry. | [optional] 

## Methods

### NewFeedEntry

`func NewFeedEntry(title string, ) *FeedEntry`

NewFeedEntry instantiates a new FeedEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedEntryWithDefaults

`func NewFeedEntryWithDefaults() *FeedEntry`

NewFeedEntryWithDefaults instantiates a new FeedEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *FeedEntry) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *FeedEntry) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *FeedEntry) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *FeedEntry) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetTitle

`func (o *FeedEntry) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FeedEntry) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FeedEntry) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetThumbnail

`func (o *FeedEntry) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *FeedEntry) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *FeedEntry) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *FeedEntry) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FeedEntry) GetCreatedBy() Person`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FeedEntry) GetCreatedByOk() (*Person, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FeedEntry) SetCreatedBy(v Person)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FeedEntry) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUiConfig

`func (o *FeedEntry) GetUiConfig() FeedEntryUiConfig`

GetUiConfig returns the UiConfig field if non-nil, zero value otherwise.

### GetUiConfigOk

`func (o *FeedEntry) GetUiConfigOk() (*FeedEntryUiConfig, bool)`

GetUiConfigOk returns a tuple with the UiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiConfig

`func (o *FeedEntry) SetUiConfig(v FeedEntryUiConfig)`

SetUiConfig sets UiConfig field to given value.

### HasUiConfig

`func (o *FeedEntry) HasUiConfig() bool`

HasUiConfig returns a boolean if a field has been set.

### GetSnippet

`func (o *FeedEntry) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *FeedEntry) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *FeedEntry) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *FeedEntry) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetJustificationType

`func (o *FeedEntry) GetJustificationType() string`

GetJustificationType returns the JustificationType field if non-nil, zero value otherwise.

### GetJustificationTypeOk

`func (o *FeedEntry) GetJustificationTypeOk() (*string, bool)`

GetJustificationTypeOk returns a tuple with the JustificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustificationType

`func (o *FeedEntry) SetJustificationType(v string)`

SetJustificationType sets JustificationType field to given value.

### HasJustificationType

`func (o *FeedEntry) HasJustificationType() bool`

HasJustificationType returns a boolean if a field has been set.

### GetJustification

`func (o *FeedEntry) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *FeedEntry) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *FeedEntry) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *FeedEntry) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### GetTrackingToken

`func (o *FeedEntry) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *FeedEntry) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *FeedEntry) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *FeedEntry) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetDocument

`func (o *FeedEntry) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *FeedEntry) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *FeedEntry) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *FeedEntry) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetEvent

`func (o *FeedEntry) GetEvent() CalendarEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *FeedEntry) GetEventOk() (*CalendarEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *FeedEntry) SetEvent(v CalendarEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *FeedEntry) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetAnnouncement

`func (o *FeedEntry) GetAnnouncement() Announcement`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *FeedEntry) GetAnnouncementOk() (*Announcement, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *FeedEntry) SetAnnouncement(v Announcement)`

SetAnnouncement sets Announcement field to given value.

### HasAnnouncement

`func (o *FeedEntry) HasAnnouncement() bool`

HasAnnouncement returns a boolean if a field has been set.

### GetCollection

`func (o *FeedEntry) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *FeedEntry) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *FeedEntry) SetCollection(v Collection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *FeedEntry) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetCollectionItem

`func (o *FeedEntry) GetCollectionItem() CollectionItem`

GetCollectionItem returns the CollectionItem field if non-nil, zero value otherwise.

### GetCollectionItemOk

`func (o *FeedEntry) GetCollectionItemOk() (*CollectionItem, bool)`

GetCollectionItemOk returns a tuple with the CollectionItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionItem

`func (o *FeedEntry) SetCollectionItem(v CollectionItem)`

SetCollectionItem sets CollectionItem field to given value.

### HasCollectionItem

`func (o *FeedEntry) HasCollectionItem() bool`

HasCollectionItem returns a boolean if a field has been set.

### GetPerson

`func (o *FeedEntry) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *FeedEntry) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *FeedEntry) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *FeedEntry) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetApp

`func (o *FeedEntry) GetApp() AppResult`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *FeedEntry) GetAppOk() (*AppResult, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *FeedEntry) SetApp(v AppResult)`

SetApp sets App field to given value.

### HasApp

`func (o *FeedEntry) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetActivities

`func (o *FeedEntry) GetActivities() []UserActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *FeedEntry) GetActivitiesOk() (*[]UserActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *FeedEntry) SetActivities(v []UserActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *FeedEntry) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetDocumentVisitorCount

`func (o *FeedEntry) GetDocumentVisitorCount() CountInfo`

GetDocumentVisitorCount returns the DocumentVisitorCount field if non-nil, zero value otherwise.

### GetDocumentVisitorCountOk

`func (o *FeedEntry) GetDocumentVisitorCountOk() (*CountInfo, bool)`

GetDocumentVisitorCountOk returns a tuple with the DocumentVisitorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVisitorCount

`func (o *FeedEntry) SetDocumentVisitorCount(v CountInfo)`

SetDocumentVisitorCount sets DocumentVisitorCount field to given value.

### HasDocumentVisitorCount

`func (o *FeedEntry) HasDocumentVisitorCount() bool`

HasDocumentVisitorCount returns a boolean if a field has been set.

### GetViewUrl

`func (o *FeedEntry) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *FeedEntry) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *FeedEntry) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *FeedEntry) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.

### GetAdditionalClientActions

`func (o *FeedEntry) GetAdditionalClientActions() []ClientAction`

GetAdditionalClientActions returns the AdditionalClientActions field if non-nil, zero value otherwise.

### GetAdditionalClientActionsOk

`func (o *FeedEntry) GetAdditionalClientActionsOk() (*[]ClientAction, bool)`

GetAdditionalClientActionsOk returns a tuple with the AdditionalClientActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalClientActions

`func (o *FeedEntry) SetAdditionalClientActions(v []ClientAction)`

SetAdditionalClientActions sets AdditionalClientActions field to given value.

### HasAdditionalClientActions

`func (o *FeedEntry) HasAdditionalClientActions() bool`

HasAdditionalClientActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


