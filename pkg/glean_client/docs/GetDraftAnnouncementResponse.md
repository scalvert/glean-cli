# GetDraftAnnouncementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Announcement** | Pointer to [**Announcement**](Announcement.md) |  | [optional] 
**Error** | Pointer to [**AnnouncementError**](AnnouncementError.md) |  | [optional] 

## Methods

### NewGetDraftAnnouncementResponse

`func NewGetDraftAnnouncementResponse() *GetDraftAnnouncementResponse`

NewGetDraftAnnouncementResponse instantiates a new GetDraftAnnouncementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDraftAnnouncementResponseWithDefaults

`func NewGetDraftAnnouncementResponseWithDefaults() *GetDraftAnnouncementResponse`

NewGetDraftAnnouncementResponseWithDefaults instantiates a new GetDraftAnnouncementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncement

`func (o *GetDraftAnnouncementResponse) GetAnnouncement() Announcement`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *GetDraftAnnouncementResponse) GetAnnouncementOk() (*Announcement, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *GetDraftAnnouncementResponse) SetAnnouncement(v Announcement)`

SetAnnouncement sets Announcement field to given value.

### HasAnnouncement

`func (o *GetDraftAnnouncementResponse) HasAnnouncement() bool`

HasAnnouncement returns a boolean if a field has been set.

### GetError

`func (o *GetDraftAnnouncementResponse) GetError() AnnouncementError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetDraftAnnouncementResponse) GetErrorOk() (*AnnouncementError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetDraftAnnouncementResponse) SetError(v AnnouncementError)`

SetError sets Error field to given value.

### HasError

`func (o *GetDraftAnnouncementResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


