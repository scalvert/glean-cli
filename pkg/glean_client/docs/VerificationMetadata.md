# VerificationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastVerifier** | Pointer to [**Person**](Person.md) |  | [optional] 
**LastVerificationTs** | Pointer to **int32** | The unix timestamp of the verification (in seconds since epoch UTC). | [optional] 
**ExpirationTs** | Pointer to **int32** | The unix timestamp of the verification expiration if applicable (in seconds since epoch UTC). | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Reminders** | Pointer to [**[]Reminder**](Reminder.md) | Info about all outstanding verification reminders for the document if exists. | [optional] 
**LastReminder** | Pointer to [**Reminder**](Reminder.md) |  | [optional] 
**VisitorCount** | Pointer to [**[]CountInfo**](CountInfo.md) | Number of visitors to the document during included time periods. | [optional] 
**CandidateVerifiers** | Pointer to [**[]Person**](Person.md) | List of potential verifiers for the document e.g. old verifiers and/or users with view/edit permissions. | [optional] 

## Methods

### NewVerificationMetadata

`func NewVerificationMetadata() *VerificationMetadata`

NewVerificationMetadata instantiates a new VerificationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationMetadataWithDefaults

`func NewVerificationMetadataWithDefaults() *VerificationMetadata`

NewVerificationMetadataWithDefaults instantiates a new VerificationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastVerifier

`func (o *VerificationMetadata) GetLastVerifier() Person`

GetLastVerifier returns the LastVerifier field if non-nil, zero value otherwise.

### GetLastVerifierOk

`func (o *VerificationMetadata) GetLastVerifierOk() (*Person, bool)`

GetLastVerifierOk returns a tuple with the LastVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifier

`func (o *VerificationMetadata) SetLastVerifier(v Person)`

SetLastVerifier sets LastVerifier field to given value.

### HasLastVerifier

`func (o *VerificationMetadata) HasLastVerifier() bool`

HasLastVerifier returns a boolean if a field has been set.

### GetLastVerificationTs

`func (o *VerificationMetadata) GetLastVerificationTs() int32`

GetLastVerificationTs returns the LastVerificationTs field if non-nil, zero value otherwise.

### GetLastVerificationTsOk

`func (o *VerificationMetadata) GetLastVerificationTsOk() (*int32, bool)`

GetLastVerificationTsOk returns a tuple with the LastVerificationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerificationTs

`func (o *VerificationMetadata) SetLastVerificationTs(v int32)`

SetLastVerificationTs sets LastVerificationTs field to given value.

### HasLastVerificationTs

`func (o *VerificationMetadata) HasLastVerificationTs() bool`

HasLastVerificationTs returns a boolean if a field has been set.

### GetExpirationTs

`func (o *VerificationMetadata) GetExpirationTs() int32`

GetExpirationTs returns the ExpirationTs field if non-nil, zero value otherwise.

### GetExpirationTsOk

`func (o *VerificationMetadata) GetExpirationTsOk() (*int32, bool)`

GetExpirationTsOk returns a tuple with the ExpirationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTs

`func (o *VerificationMetadata) SetExpirationTs(v int32)`

SetExpirationTs sets ExpirationTs field to given value.

### HasExpirationTs

`func (o *VerificationMetadata) HasExpirationTs() bool`

HasExpirationTs returns a boolean if a field has been set.

### GetDocument

`func (o *VerificationMetadata) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *VerificationMetadata) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *VerificationMetadata) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *VerificationMetadata) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetReminders

`func (o *VerificationMetadata) GetReminders() []Reminder`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *VerificationMetadata) GetRemindersOk() (*[]Reminder, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *VerificationMetadata) SetReminders(v []Reminder)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *VerificationMetadata) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetLastReminder

`func (o *VerificationMetadata) GetLastReminder() Reminder`

GetLastReminder returns the LastReminder field if non-nil, zero value otherwise.

### GetLastReminderOk

`func (o *VerificationMetadata) GetLastReminderOk() (*Reminder, bool)`

GetLastReminderOk returns a tuple with the LastReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReminder

`func (o *VerificationMetadata) SetLastReminder(v Reminder)`

SetLastReminder sets LastReminder field to given value.

### HasLastReminder

`func (o *VerificationMetadata) HasLastReminder() bool`

HasLastReminder returns a boolean if a field has been set.

### GetVisitorCount

`func (o *VerificationMetadata) GetVisitorCount() []CountInfo`

GetVisitorCount returns the VisitorCount field if non-nil, zero value otherwise.

### GetVisitorCountOk

`func (o *VerificationMetadata) GetVisitorCountOk() (*[]CountInfo, bool)`

GetVisitorCountOk returns a tuple with the VisitorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorCount

`func (o *VerificationMetadata) SetVisitorCount(v []CountInfo)`

SetVisitorCount sets VisitorCount field to given value.

### HasVisitorCount

`func (o *VerificationMetadata) HasVisitorCount() bool`

HasVisitorCount returns a boolean if a field has been set.

### GetCandidateVerifiers

`func (o *VerificationMetadata) GetCandidateVerifiers() []Person`

GetCandidateVerifiers returns the CandidateVerifiers field if non-nil, zero value otherwise.

### GetCandidateVerifiersOk

`func (o *VerificationMetadata) GetCandidateVerifiersOk() (*[]Person, bool)`

GetCandidateVerifiersOk returns a tuple with the CandidateVerifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateVerifiers

`func (o *VerificationMetadata) SetCandidateVerifiers(v []Person)`

SetCandidateVerifiers sets CandidateVerifiers field to given value.

### HasCandidateVerifiers

`func (o *VerificationMetadata) HasCandidateVerifiers() bool`

HasCandidateVerifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


