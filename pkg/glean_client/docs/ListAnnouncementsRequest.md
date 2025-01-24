# ListAnnouncementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**AnnouncementChannel**](AnnouncementChannel.md) |  | [optional] 

## Methods

### NewListAnnouncementsRequest

`func NewListAnnouncementsRequest() *ListAnnouncementsRequest`

NewListAnnouncementsRequest instantiates a new ListAnnouncementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnnouncementsRequestWithDefaults

`func NewListAnnouncementsRequestWithDefaults() *ListAnnouncementsRequest`

NewListAnnouncementsRequestWithDefaults instantiates a new ListAnnouncementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ListAnnouncementsRequest) GetChannel() AnnouncementChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ListAnnouncementsRequest) GetChannelOk() (*AnnouncementChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ListAnnouncementsRequest) SetChannel(v AnnouncementChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ListAnnouncementsRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


