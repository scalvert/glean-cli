# GetAnnouncementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Announcement** | Pointer to [**Announcement**](Announcement.md) |  | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular announcement. To be used for /feedback reporting. | [optional] 
**Error** | Pointer to [**AnnouncementError**](AnnouncementError.md) |  | [optional] 

## Methods

### NewGetAnnouncementResponse

`func NewGetAnnouncementResponse() *GetAnnouncementResponse`

NewGetAnnouncementResponse instantiates a new GetAnnouncementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnnouncementResponseWithDefaults

`func NewGetAnnouncementResponseWithDefaults() *GetAnnouncementResponse`

NewGetAnnouncementResponseWithDefaults instantiates a new GetAnnouncementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnouncement

`func (o *GetAnnouncementResponse) GetAnnouncement() Announcement`

GetAnnouncement returns the Announcement field if non-nil, zero value otherwise.

### GetAnnouncementOk

`func (o *GetAnnouncementResponse) GetAnnouncementOk() (*Announcement, bool)`

GetAnnouncementOk returns a tuple with the Announcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncement

`func (o *GetAnnouncementResponse) SetAnnouncement(v Announcement)`

SetAnnouncement sets Announcement field to given value.

### HasAnnouncement

`func (o *GetAnnouncementResponse) HasAnnouncement() bool`

HasAnnouncement returns a boolean if a field has been set.

### GetTrackingToken

`func (o *GetAnnouncementResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *GetAnnouncementResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *GetAnnouncementResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *GetAnnouncementResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetError

`func (o *GetAnnouncementResponse) GetError() AnnouncementError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAnnouncementResponse) GetErrorOk() (*AnnouncementError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAnnouncementResponse) SetError(v AnnouncementError)`

SetError sets Error field to given value.

### HasError

`func (o *GetAnnouncementResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


