# Reminder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | [**Person**](Person.md) |  | 
**Requestor** | Pointer to [**Person**](Person.md) |  | [optional] 
**RemindAt** | **int32** | Unix timestamp for when the reminder should trigger (in seconds since epoch UTC). | 
**CreatedAt** | Pointer to **int32** | Unix timestamp for when the reminder was first created (in seconds since epoch UTC). | [optional] 
**Reason** | Pointer to **string** | An optional free-text reason for the reminder. This is particularly useful when a reminder is used to ask for verification from another user (for example, \&quot;Duplicate\&quot;, \&quot;Incomplete\&quot;, \&quot;Incorrect\&quot;). | [optional] 

## Methods

### NewReminder

`func NewReminder(assignee Person, remindAt int32, ) *Reminder`

NewReminder instantiates a new Reminder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReminderWithDefaults

`func NewReminderWithDefaults() *Reminder`

NewReminderWithDefaults instantiates a new Reminder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *Reminder) GetAssignee() Person`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Reminder) GetAssigneeOk() (*Person, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Reminder) SetAssignee(v Person)`

SetAssignee sets Assignee field to given value.


### GetRequestor

`func (o *Reminder) GetRequestor() Person`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *Reminder) GetRequestorOk() (*Person, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *Reminder) SetRequestor(v Person)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *Reminder) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetRemindAt

`func (o *Reminder) GetRemindAt() int32`

GetRemindAt returns the RemindAt field if non-nil, zero value otherwise.

### GetRemindAtOk

`func (o *Reminder) GetRemindAtOk() (*int32, bool)`

GetRemindAtOk returns a tuple with the RemindAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemindAt

`func (o *Reminder) SetRemindAt(v int32)`

SetRemindAt sets RemindAt field to given value.


### GetCreatedAt

`func (o *Reminder) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Reminder) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Reminder) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Reminder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetReason

`func (o *Reminder) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Reminder) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Reminder) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Reminder) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


