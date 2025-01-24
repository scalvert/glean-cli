# InviteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignUpTime** | Pointer to **time.Time** | The time this person signed up in ISO format (ISO 8601). | [optional] 
**Invites** | Pointer to [**[]ChannelInviteInfo**](ChannelInviteInfo.md) | Latest invites received by the user for each channel | [optional] 
**Inviter** | Pointer to [**Person**](Person.md) |  | [optional] 
**InviteTime** | Pointer to **time.Time** | The time this person was invited in ISO format (ISO 8601). | [optional] 
**ReminderTime** | Pointer to **time.Time** | The time this person was reminded in ISO format (ISO 8601) if a reminder was sent. | [optional] 

## Methods

### NewInviteInfo

`func NewInviteInfo() *InviteInfo`

NewInviteInfo instantiates a new InviteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteInfoWithDefaults

`func NewInviteInfoWithDefaults() *InviteInfo`

NewInviteInfoWithDefaults instantiates a new InviteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignUpTime

`func (o *InviteInfo) GetSignUpTime() time.Time`

GetSignUpTime returns the SignUpTime field if non-nil, zero value otherwise.

### GetSignUpTimeOk

`func (o *InviteInfo) GetSignUpTimeOk() (*time.Time, bool)`

GetSignUpTimeOk returns a tuple with the SignUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignUpTime

`func (o *InviteInfo) SetSignUpTime(v time.Time)`

SetSignUpTime sets SignUpTime field to given value.

### HasSignUpTime

`func (o *InviteInfo) HasSignUpTime() bool`

HasSignUpTime returns a boolean if a field has been set.

### GetInvites

`func (o *InviteInfo) GetInvites() []ChannelInviteInfo`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *InviteInfo) GetInvitesOk() (*[]ChannelInviteInfo, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *InviteInfo) SetInvites(v []ChannelInviteInfo)`

SetInvites sets Invites field to given value.

### HasInvites

`func (o *InviteInfo) HasInvites() bool`

HasInvites returns a boolean if a field has been set.

### GetInviter

`func (o *InviteInfo) GetInviter() Person`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *InviteInfo) GetInviterOk() (*Person, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *InviteInfo) SetInviter(v Person)`

SetInviter sets Inviter field to given value.

### HasInviter

`func (o *InviteInfo) HasInviter() bool`

HasInviter returns a boolean if a field has been set.

### GetInviteTime

`func (o *InviteInfo) GetInviteTime() time.Time`

GetInviteTime returns the InviteTime field if non-nil, zero value otherwise.

### GetInviteTimeOk

`func (o *InviteInfo) GetInviteTimeOk() (*time.Time, bool)`

GetInviteTimeOk returns a tuple with the InviteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteTime

`func (o *InviteInfo) SetInviteTime(v time.Time)`

SetInviteTime sets InviteTime field to given value.

### HasInviteTime

`func (o *InviteInfo) HasInviteTime() bool`

HasInviteTime returns a boolean if a field has been set.

### GetReminderTime

`func (o *InviteInfo) GetReminderTime() time.Time`

GetReminderTime returns the ReminderTime field if non-nil, zero value otherwise.

### GetReminderTimeOk

`func (o *InviteInfo) GetReminderTimeOk() (*time.Time, bool)`

GetReminderTimeOk returns a tuple with the ReminderTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderTime

`func (o *InviteInfo) SetReminderTime(v time.Time)`

SetReminderTime sets ReminderTime field to given value.

### HasReminderTime

`func (o *InviteInfo) HasReminderTime() bool`

HasReminderTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


