# ChannelInviteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**CommunicationChannel**](CommunicationChannel.md) |  | [optional] 
**IsAutoInvite** | Pointer to **bool** | Bit that tracks if this invite was automatically sent or user-sent | [optional] 
**Inviter** | Pointer to [**Person**](Person.md) |  | [optional] 
**InviteTime** | Pointer to **time.Time** | The time this person was invited in ISO format (ISO 8601). | [optional] 
**ReminderTime** | Pointer to **time.Time** | The time this person was reminded in ISO format (ISO 8601) if a reminder was sent. | [optional] 

## Methods

### NewChannelInviteInfo

`func NewChannelInviteInfo() *ChannelInviteInfo`

NewChannelInviteInfo instantiates a new ChannelInviteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelInviteInfoWithDefaults

`func NewChannelInviteInfoWithDefaults() *ChannelInviteInfo`

NewChannelInviteInfoWithDefaults instantiates a new ChannelInviteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelInviteInfo) GetChannel() CommunicationChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelInviteInfo) GetChannelOk() (*CommunicationChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelInviteInfo) SetChannel(v CommunicationChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChannelInviteInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetIsAutoInvite

`func (o *ChannelInviteInfo) GetIsAutoInvite() bool`

GetIsAutoInvite returns the IsAutoInvite field if non-nil, zero value otherwise.

### GetIsAutoInviteOk

`func (o *ChannelInviteInfo) GetIsAutoInviteOk() (*bool, bool)`

GetIsAutoInviteOk returns a tuple with the IsAutoInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoInvite

`func (o *ChannelInviteInfo) SetIsAutoInvite(v bool)`

SetIsAutoInvite sets IsAutoInvite field to given value.

### HasIsAutoInvite

`func (o *ChannelInviteInfo) HasIsAutoInvite() bool`

HasIsAutoInvite returns a boolean if a field has been set.

### GetInviter

`func (o *ChannelInviteInfo) GetInviter() Person`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *ChannelInviteInfo) GetInviterOk() (*Person, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *ChannelInviteInfo) SetInviter(v Person)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *ChannelInviteInfo) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetInviteTime

`func (o *ChannelInviteInfo) GetInviteTime() time.Time`

GetInviteTime returns the InviteTime field if non-nil, zero value otherwise.

### GetInviteTimeOk

`func (o *ChannelInviteInfo) GetInviteTimeOk() (*time.Time, bool)`

GetInviteTimeOk returns a tuple with the InviteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteTime

`func (o *ChannelInviteInfo) SetInviteTime(v time.Time)`

SetInviteTime sets InviteTime field to given value.

### HasInviteTime

`func (o *ChannelInviteInfo) HasInviteTime() bool`

HasInviteTime returns a boolean if a field has been set.

### GetReminderTime

`func (o *ChannelInviteInfo) GetReminderTime() time.Time`

GetReminderTime returns the ReminderTime field if non-nil, zero value otherwise.

### GetReminderTimeOk

`func (o *ChannelInviteInfo) GetReminderTimeOk() (*time.Time, bool)`

GetReminderTimeOk returns a tuple with the ReminderTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderTime

`func (o *ChannelInviteInfo) SetReminderTime(v time.Time)`

SetReminderTime sets ReminderTime field to given value.

### HasReminderTime

`func (o *ChannelInviteInfo) HasReminderTime() bool`

HasReminderTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


