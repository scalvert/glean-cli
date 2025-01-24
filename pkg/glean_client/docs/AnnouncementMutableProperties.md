# AnnouncementMutableProperties

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

## Methods

### NewAnnouncementMutableProperties

`func NewAnnouncementMutableProperties() *AnnouncementMutableProperties`

NewAnnouncementMutableProperties instantiates a new AnnouncementMutableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementMutablePropertiesWithDefaults

`func NewAnnouncementMutablePropertiesWithDefaults() *AnnouncementMutableProperties`

NewAnnouncementMutablePropertiesWithDefaults instantiates a new AnnouncementMutableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *AnnouncementMutableProperties) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AnnouncementMutableProperties) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AnnouncementMutableProperties) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AnnouncementMutableProperties) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *AnnouncementMutableProperties) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AnnouncementMutableProperties) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AnnouncementMutableProperties) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AnnouncementMutableProperties) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetTitle

`func (o *AnnouncementMutableProperties) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AnnouncementMutableProperties) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AnnouncementMutableProperties) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AnnouncementMutableProperties) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *AnnouncementMutableProperties) GetBody() StructuredText`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AnnouncementMutableProperties) GetBodyOk() (*StructuredText, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AnnouncementMutableProperties) SetBody(v StructuredText)`

SetBody sets Body field to given value.

### HasBody

`func (o *AnnouncementMutableProperties) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEmoji

`func (o *AnnouncementMutableProperties) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *AnnouncementMutableProperties) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *AnnouncementMutableProperties) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *AnnouncementMutableProperties) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetThumbnail

`func (o *AnnouncementMutableProperties) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *AnnouncementMutableProperties) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *AnnouncementMutableProperties) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *AnnouncementMutableProperties) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetBanner

`func (o *AnnouncementMutableProperties) GetBanner() Thumbnail`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *AnnouncementMutableProperties) GetBannerOk() (*Thumbnail, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *AnnouncementMutableProperties) SetBanner(v Thumbnail)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *AnnouncementMutableProperties) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *AnnouncementMutableProperties) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *AnnouncementMutableProperties) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *AnnouncementMutableProperties) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *AnnouncementMutableProperties) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetSourceDocumentId

`func (o *AnnouncementMutableProperties) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *AnnouncementMutableProperties) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *AnnouncementMutableProperties) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *AnnouncementMutableProperties) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### GetHideAttribution

`func (o *AnnouncementMutableProperties) GetHideAttribution() bool`

GetHideAttribution returns the HideAttribution field if non-nil, zero value otherwise.

### GetHideAttributionOk

`func (o *AnnouncementMutableProperties) GetHideAttributionOk() (*bool, bool)`

GetHideAttributionOk returns a tuple with the HideAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAttribution

`func (o *AnnouncementMutableProperties) SetHideAttribution(v bool)`

SetHideAttribution sets HideAttribution field to given value.

### HasHideAttribution

`func (o *AnnouncementMutableProperties) HasHideAttribution() bool`

HasHideAttribution returns a boolean if a field has been set.

### GetChannel

`func (o *AnnouncementMutableProperties) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AnnouncementMutableProperties) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AnnouncementMutableProperties) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AnnouncementMutableProperties) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetPostType

`func (o *AnnouncementMutableProperties) GetPostType() string`

GetPostType returns the PostType field if non-nil, zero value otherwise.

### GetPostTypeOk

`func (o *AnnouncementMutableProperties) GetPostTypeOk() (*string, bool)`

GetPostTypeOk returns a tuple with the PostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostType

`func (o *AnnouncementMutableProperties) SetPostType(v string)`

SetPostType sets PostType field to given value.

### HasPostType

`func (o *AnnouncementMutableProperties) HasPostType() bool`

HasPostType returns a boolean if a field has been set.

### GetIsPrioritized

`func (o *AnnouncementMutableProperties) GetIsPrioritized() bool`

GetIsPrioritized returns the IsPrioritized field if non-nil, zero value otherwise.

### GetIsPrioritizedOk

`func (o *AnnouncementMutableProperties) GetIsPrioritizedOk() (*bool, bool)`

GetIsPrioritizedOk returns a tuple with the IsPrioritized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrioritized

`func (o *AnnouncementMutableProperties) SetIsPrioritized(v bool)`

SetIsPrioritized sets IsPrioritized field to given value.

### HasIsPrioritized

`func (o *AnnouncementMutableProperties) HasIsPrioritized() bool`

HasIsPrioritized returns a boolean if a field has been set.

### GetViewUrl

`func (o *AnnouncementMutableProperties) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *AnnouncementMutableProperties) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *AnnouncementMutableProperties) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *AnnouncementMutableProperties) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


