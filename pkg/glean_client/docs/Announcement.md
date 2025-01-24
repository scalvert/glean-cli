# Announcement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | The date and time at which the announcement becomes active. | [optional] 
**EndTime** | Pointer to **time.Time** | The date and time at which the announcement expires. | [optional] 
**Title** | Pointer to **string** | The headline of the announcement. | [optional] 
**Body** | Pointer to [**StructuredText**](StructuredText.md) |  | [optional] 
**Emoji** | Pointer to **string** | An emoji used to indicate the nature of the announcement. | [optional] 
**Thumbnail** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**Banner** | Pointer to [**Thumbnail**](Thumbnail.md) |  | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see the announcement. Values are taken from the corresponding filters in people search. | [optional] 
**SourceDocumentId** | Pointer to **string** | The Glean Document ID of the source document this Announcement was created from (e.g. Slack thread). | [optional] 
**HideAttribution** | Pointer to **bool** | Whether or not to hide an author attribution. | [optional] 
**Channel** | Pointer to **string** | This determines whether this is a Social Feed post or a regular announcement. | [optional] 
**PostType** | Pointer to **string** | This determines whether this is an external-link post or a regular announcement post. TEXT - Regular announcement that can contain rich text. LINK - Announcement that is linked to an external site. | [optional] 
**IsPrioritized** | Pointer to **bool** | Used by the Social Feed to pin posts to the front of the feed. | [optional] 
**ViewUrl** | Pointer to **string** | URL for viewing the announcement. It will be set to document URL for announcements from other datasources e.g. simpplr. Can only be written when channel&#x3D;\&quot;SOCIAL_FEED\&quot;. | [optional] 
**DraftId** | Pointer to **int32** | The opaque id of the associated draft. | [optional] 
**Permissions** | Pointer to [**ObjectPermissions**](ObjectPermissions.md) |  | [optional] 
**Id** | Pointer to **int32** | The opaque id of the announcement. | [optional] 
**Author** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreateTimestamp** | Pointer to **int32** | Server Unix timestamp of the creation time (in seconds since epoch UTC). | [optional] 
**LastUpdateTimestamp** | Pointer to **int32** | Server Unix timestamp of the last update time (in seconds since epoch UTC). | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**ViewerInfo** | Pointer to [**AnnouncementAllOfViewerInfo**](AnnouncementAllOfViewerInfo.md) |  | [optional] 
**SourceDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**IsPublished** | Pointer to **bool** | Whether or not the announcement is published. | [optional] 

## Methods

### NewAnnouncement

`func NewAnnouncement() *Announcement`

NewAnnouncement instantiates a new Announcement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementWithDefaults

`func NewAnnouncementWithDefaults() *Announcement`

NewAnnouncementWithDefaults instantiates a new Announcement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *Announcement) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Announcement) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Announcement) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Announcement) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *Announcement) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Announcement) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Announcement) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Announcement) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetTitle

`func (o *Announcement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Announcement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Announcement) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Announcement) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *Announcement) GetBody() StructuredText`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Announcement) GetBodyOk() (*StructuredText, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Announcement) SetBody(v StructuredText)`

SetBody sets Body field to given value.

### HasBody

`func (o *Announcement) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEmoji

`func (o *Announcement) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *Announcement) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *Announcement) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *Announcement) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetThumbnail

`func (o *Announcement) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Announcement) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Announcement) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Announcement) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetBanner

`func (o *Announcement) GetBanner() Thumbnail`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *Announcement) GetBannerOk() (*Thumbnail, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *Announcement) SetBanner(v Thumbnail)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *Announcement) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *Announcement) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *Announcement) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *Announcement) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *Announcement) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetSourceDocumentId

`func (o *Announcement) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *Announcement) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *Announcement) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *Announcement) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### GetHideAttribution

`func (o *Announcement) GetHideAttribution() bool`

GetHideAttribution returns the HideAttribution field if non-nil, zero value otherwise.

### GetHideAttributionOk

`func (o *Announcement) GetHideAttributionOk() (*bool, bool)`

GetHideAttributionOk returns a tuple with the HideAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAttribution

`func (o *Announcement) SetHideAttribution(v bool)`

SetHideAttribution sets HideAttribution field to given value.

### HasHideAttribution

`func (o *Announcement) HasHideAttribution() bool`

HasHideAttribution returns a boolean if a field has been set.

### GetChannel

`func (o *Announcement) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Announcement) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Announcement) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Announcement) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetPostType

`func (o *Announcement) GetPostType() string`

GetPostType returns the PostType field if non-nil, zero value otherwise.

### GetPostTypeOk

`func (o *Announcement) GetPostTypeOk() (*string, bool)`

GetPostTypeOk returns a tuple with the PostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostType

`func (o *Announcement) SetPostType(v string)`

SetPostType sets PostType field to given value.

### HasPostType

`func (o *Announcement) HasPostType() bool`

HasPostType returns a boolean if a field has been set.

### GetIsPrioritized

`func (o *Announcement) GetIsPrioritized() bool`

GetIsPrioritized returns the IsPrioritized field if non-nil, zero value otherwise.

### GetIsPrioritizedOk

`func (o *Announcement) GetIsPrioritizedOk() (*bool, bool)`

GetIsPrioritizedOk returns a tuple with the IsPrioritized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrioritized

`func (o *Announcement) SetIsPrioritized(v bool)`

SetIsPrioritized sets IsPrioritized field to given value.

### HasIsPrioritized

`func (o *Announcement) HasIsPrioritized() bool`

HasIsPrioritized returns a boolean if a field has been set.

### GetViewUrl

`func (o *Announcement) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *Announcement) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *Announcement) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *Announcement) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.

### GetDraftId

`func (o *Announcement) GetDraftId() int32`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *Announcement) GetDraftIdOk() (*int32, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *Announcement) SetDraftId(v int32)`

SetDraftId sets DraftId field to given value.

### HasDraftId

`func (o *Announcement) HasDraftId() bool`

HasDraftId returns a boolean if a field has been set.

### GetPermissions

`func (o *Announcement) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Announcement) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Announcement) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Announcement) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetId

`func (o *Announcement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Announcement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Announcement) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Announcement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthor

`func (o *Announcement) GetAuthor() Person`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Announcement) GetAuthorOk() (*Person, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Announcement) SetAuthor(v Person)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Announcement) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *Announcement) GetCreateTimestamp() int32`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *Announcement) GetCreateTimestampOk() (*int32, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *Announcement) SetCreateTimestamp(v int32)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *Announcement) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *Announcement) GetLastUpdateTimestamp() int32`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *Announcement) GetLastUpdateTimestampOk() (*int32, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *Announcement) SetLastUpdateTimestamp(v int32)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *Announcement) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Announcement) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Announcement) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Announcement) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Announcement) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetViewerInfo

`func (o *Announcement) GetViewerInfo() AnnouncementAllOfViewerInfo`

GetViewerInfo returns the ViewerInfo field if non-nil, zero value otherwise.

### GetViewerInfoOk

`func (o *Announcement) GetViewerInfoOk() (*AnnouncementAllOfViewerInfo, bool)`

GetViewerInfoOk returns a tuple with the ViewerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerInfo

`func (o *Announcement) SetViewerInfo(v AnnouncementAllOfViewerInfo)`

SetViewerInfo sets ViewerInfo field to given value.

### HasViewerInfo

`func (o *Announcement) HasViewerInfo() bool`

HasViewerInfo returns a boolean if a field has been set.

### GetSourceDocument

`func (o *Announcement) GetSourceDocument() Document`

GetSourceDocument returns the SourceDocument field if non-nil, zero value otherwise.

### GetSourceDocumentOk

`func (o *Announcement) GetSourceDocumentOk() (*Document, bool)`

GetSourceDocumentOk returns a tuple with the SourceDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocument

`func (o *Announcement) SetSourceDocument(v Document)`

SetSourceDocument sets SourceDocument field to given value.

### HasSourceDocument

`func (o *Announcement) HasSourceDocument() bool`

HasSourceDocument returns a boolean if a field has been set.

### GetIsPublished

`func (o *Announcement) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *Announcement) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *Announcement) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *Announcement) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


