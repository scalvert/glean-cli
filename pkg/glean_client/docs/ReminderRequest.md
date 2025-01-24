# ReminderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** | The document which the verification is for new reminders and/or update. | 
**Assignee** | Pointer to **string** | The obfuscated id of the person this verification is assigned to. | [optional] 
**RemindInDays** | Pointer to **int32** | Reminder for the next verifications in terms of days. For deletion, this will be omitted. | [optional] 
**Reason** | Pointer to **string** | An optional free-text reason for the reminder. This is particularly useful when a reminder is used to ask for verification from another user (for example, \&quot;Duplicate\&quot;, \&quot;Incomplete\&quot;, \&quot;Incorrect\&quot;). | [optional] 

## Methods

### NewReminderRequest

`func NewReminderRequest(documentId string, ) *ReminderRequest`

NewReminderRequest instantiates a new ReminderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReminderRequestWithDefaults

`func NewReminderRequestWithDefaults() *ReminderRequest`

NewReminderRequestWithDefaults instantiates a new ReminderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *ReminderRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *ReminderRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *ReminderRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetAssignee

`func (o *ReminderRequest) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ReminderRequest) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ReminderRequest) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ReminderRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetRemindInDays

`func (o *ReminderRequest) GetRemindInDays() int32`

GetRemindInDays returns the RemindInDays field if non-nil, zero value otherwise.

### GetRemindInDaysOk

`func (o *ReminderRequest) GetRemindInDaysOk() (*int32, bool)`

GetRemindInDaysOk returns a tuple with the RemindInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemindInDays

`func (o *ReminderRequest) SetRemindInDays(v int32)`

SetRemindInDays sets RemindInDays field to given value.

### HasRemindInDays

`func (o *ReminderRequest) HasRemindInDays() bool`

HasRemindInDays returns a boolean if a field has been set.

### GetReason

`func (o *ReminderRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ReminderRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ReminderRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ReminderRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


