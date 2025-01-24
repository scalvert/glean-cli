# UpdateAnnouncementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **time.Time** | The date and time at which the announcement becomes active. | 
**EndTime** | **time.Time** | The date and time at which the announcement expires. | 
**Title** | **string** | The headline of the announcement. | 
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
**Id** | **int32** | The opaque id of the announcement. | 

## Methods

### NewUpdateAnnouncementRequest

`func NewUpdateAnnouncementRequest(startTime time.Time, endTime time.Time, title string, id int32, ) *UpdateAnnouncementRequest`

NewUpdateAnnouncementRequest instantiates a new UpdateAnnouncementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAnnouncementRequestWithDefaults

`func NewUpdateAnnouncementRequestWithDefaults() *UpdateAnnouncementRequest`

NewUpdateAnnouncementRequestWithDefaults instantiates a new UpdateAnnouncementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *UpdateAnnouncementRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdateAnnouncementRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpdateAnnouncementRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *UpdateAnnouncementRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UpdateAnnouncementRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UpdateAnnouncementRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetTitle

`func (o *UpdateAnnouncementRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateAnnouncementRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateAnnouncementRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *UpdateAnnouncementRequest) GetBody() StructuredText`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateAnnouncementRequest) GetBodyOk() (*StructuredText, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateAnnouncementRequest) SetBody(v StructuredText)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateAnnouncementRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEmoji

`func (o *UpdateAnnouncementRequest) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *UpdateAnnouncementRequest) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *UpdateAnnouncementRequest) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *UpdateAnnouncementRequest) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetThumbnail

`func (o *UpdateAnnouncementRequest) GetThumbnail() Thumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *UpdateAnnouncementRequest) GetThumbnailOk() (*Thumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *UpdateAnnouncementRequest) SetThumbnail(v Thumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *UpdateAnnouncementRequest) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetBanner

`func (o *UpdateAnnouncementRequest) GetBanner() Thumbnail`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *UpdateAnnouncementRequest) GetBannerOk() (*Thumbnail, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *UpdateAnnouncementRequest) SetBanner(v Thumbnail)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *UpdateAnnouncementRequest) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *UpdateAnnouncementRequest) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *UpdateAnnouncementRequest) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *UpdateAnnouncementRequest) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *UpdateAnnouncementRequest) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetSourceDocumentId

`func (o *UpdateAnnouncementRequest) GetSourceDocumentId() string`

GetSourceDocumentId returns the SourceDocumentId field if non-nil, zero value otherwise.

### GetSourceDocumentIdOk

`func (o *UpdateAnnouncementRequest) GetSourceDocumentIdOk() (*string, bool)`

GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentId

`func (o *UpdateAnnouncementRequest) SetSourceDocumentId(v string)`

SetSourceDocumentId sets SourceDocumentId field to given value.

### HasSourceDocumentId

`func (o *UpdateAnnouncementRequest) HasSourceDocumentId() bool`

HasSourceDocumentId returns a boolean if a field has been set.

### GetHideAttribution

`func (o *UpdateAnnouncementRequest) GetHideAttribution() bool`

GetHideAttribution returns the HideAttribution field if non-nil, zero value otherwise.

### GetHideAttributionOk

`func (o *UpdateAnnouncementRequest) GetHideAttributionOk() (*bool, bool)`

GetHideAttributionOk returns a tuple with the HideAttribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAttribution

`func (o *UpdateAnnouncementRequest) SetHideAttribution(v bool)`

SetHideAttribution sets HideAttribution field to given value.

### HasHideAttribution

`func (o *UpdateAnnouncementRequest) HasHideAttribution() bool`

HasHideAttribution returns a boolean if a field has been set.

### GetChannel

`func (o *UpdateAnnouncementRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UpdateAnnouncementRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UpdateAnnouncementRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UpdateAnnouncementRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetPostType

`func (o *UpdateAnnouncementRequest) GetPostType() string`

GetPostType returns the PostType field if non-nil, zero value otherwise.

### GetPostTypeOk

`func (o *UpdateAnnouncementRequest) GetPostTypeOk() (*string, bool)`

GetPostTypeOk returns a tuple with the PostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostType

`func (o *UpdateAnnouncementRequest) SetPostType(v string)`

SetPostType sets PostType field to given value.

### HasPostType

`func (o *UpdateAnnouncementRequest) HasPostType() bool`

HasPostType returns a boolean if a field has been set.

### GetIsPrioritized

`func (o *UpdateAnnouncementRequest) GetIsPrioritized() bool`

GetIsPrioritized returns the IsPrioritized field if non-nil, zero value otherwise.

### GetIsPrioritizedOk

`func (o *UpdateAnnouncementRequest) GetIsPrioritizedOk() (*bool, bool)`

GetIsPrioritizedOk returns a tuple with the IsPrioritized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrioritized

`func (o *UpdateAnnouncementRequest) SetIsPrioritized(v bool)`

SetIsPrioritized sets IsPrioritized field to given value.

### HasIsPrioritized

`func (o *UpdateAnnouncementRequest) HasIsPrioritized() bool`

HasIsPrioritized returns a boolean if a field has been set.

### GetViewUrl

`func (o *UpdateAnnouncementRequest) GetViewUrl() string`

GetViewUrl returns the ViewUrl field if non-nil, zero value otherwise.

### GetViewUrlOk

`func (o *UpdateAnnouncementRequest) GetViewUrlOk() (*string, bool)`

GetViewUrlOk returns a tuple with the ViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUrl

`func (o *UpdateAnnouncementRequest) SetViewUrl(v string)`

SetViewUrl sets ViewUrl field to given value.

### HasViewUrl

`func (o *UpdateAnnouncementRequest) HasViewUrl() bool`

HasViewUrl returns a boolean if a field has been set.

### GetId

`func (o *UpdateAnnouncementRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAnnouncementRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAnnouncementRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


