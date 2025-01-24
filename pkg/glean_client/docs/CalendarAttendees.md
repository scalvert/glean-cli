# CalendarAttendees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**People** | Pointer to [**[]CalendarAttendee**](CalendarAttendee.md) | Full details of some of the attendees of this event | [optional] 
**IsLimit** | Pointer to **bool** | Whether the total count of the people returned is at the retrieval limit. | [optional] 
**Total** | Pointer to **int32** | Total number of attendees in this event. | [optional] 
**NumAccepted** | Pointer to **int32** | Total number of attendees who have accepted this event. | [optional] 
**NumDeclined** | Pointer to **int32** | Total number of attendees who have declined this event. | [optional] 
**NumNoResponse** | Pointer to **int32** | Total number of attendees who have not responded to this event. | [optional] 
**NumTentative** | Pointer to **int32** | Total number of attendees who have responded tentatively (i.e. responded maybe) to this event. | [optional] 

## Methods

### NewCalendarAttendees

`func NewCalendarAttendees() *CalendarAttendees`

NewCalendarAttendees instantiates a new CalendarAttendees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarAttendeesWithDefaults

`func NewCalendarAttendeesWithDefaults() *CalendarAttendees`

NewCalendarAttendeesWithDefaults instantiates a new CalendarAttendees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeople

`func (o *CalendarAttendees) GetPeople() []CalendarAttendee`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *CalendarAttendees) GetPeopleOk() (*[]CalendarAttendee, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *CalendarAttendees) SetPeople(v []CalendarAttendee)`

SetPeople sets People field to given value.

### HasPeople

`func (o *CalendarAttendees) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetIsLimit

`func (o *CalendarAttendees) GetIsLimit() bool`

GetIsLimit returns the IsLimit field if non-nil, zero value otherwise.

### GetIsLimitOk

`func (o *CalendarAttendees) GetIsLimitOk() (*bool, bool)`

GetIsLimitOk returns a tuple with the IsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLimit

`func (o *CalendarAttendees) SetIsLimit(v bool)`

SetIsLimit sets IsLimit field to given value.

### HasIsLimit

`func (o *CalendarAttendees) HasIsLimit() bool`

HasIsLimit returns a boolean if a field has been set.

### GetTotal

`func (o *CalendarAttendees) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CalendarAttendees) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CalendarAttendees) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CalendarAttendees) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetNumAccepted

`func (o *CalendarAttendees) GetNumAccepted() int32`

GetNumAccepted returns the NumAccepted field if non-nil, zero value otherwise.

### GetNumAcceptedOk

`func (o *CalendarAttendees) GetNumAcceptedOk() (*int32, bool)`

GetNumAcceptedOk returns a tuple with the NumAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAccepted

`func (o *CalendarAttendees) SetNumAccepted(v int32)`

SetNumAccepted sets NumAccepted field to given value.

### HasNumAccepted

`func (o *CalendarAttendees) HasNumAccepted() bool`

HasNumAccepted returns a boolean if a field has been set.

### GetNumDeclined

`func (o *CalendarAttendees) GetNumDeclined() int32`

GetNumDeclined returns the NumDeclined field if non-nil, zero value otherwise.

### GetNumDeclinedOk

`func (o *CalendarAttendees) GetNumDeclinedOk() (*int32, bool)`

GetNumDeclinedOk returns a tuple with the NumDeclined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeclined

`func (o *CalendarAttendees) SetNumDeclined(v int32)`

SetNumDeclined sets NumDeclined field to given value.

### HasNumDeclined

`func (o *CalendarAttendees) HasNumDeclined() bool`

HasNumDeclined returns a boolean if a field has been set.

### GetNumNoResponse

`func (o *CalendarAttendees) GetNumNoResponse() int32`

GetNumNoResponse returns the NumNoResponse field if non-nil, zero value otherwise.

### GetNumNoResponseOk

`func (o *CalendarAttendees) GetNumNoResponseOk() (*int32, bool)`

GetNumNoResponseOk returns a tuple with the NumNoResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNoResponse

`func (o *CalendarAttendees) SetNumNoResponse(v int32)`

SetNumNoResponse sets NumNoResponse field to given value.

### HasNumNoResponse

`func (o *CalendarAttendees) HasNumNoResponse() bool`

HasNumNoResponse returns a boolean if a field has been set.

### GetNumTentative

`func (o *CalendarAttendees) GetNumTentative() int32`

GetNumTentative returns the NumTentative field if non-nil, zero value otherwise.

### GetNumTentativeOk

`func (o *CalendarAttendees) GetNumTentativeOk() (*int32, bool)`

GetNumTentativeOk returns a tuple with the NumTentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTentative

`func (o *CalendarAttendees) SetNumTentative(v int32)`

SetNumTentative sets NumTentative field to given value.

### HasNumTentative

`func (o *CalendarAttendees) HasNumTentative() bool`

HasNumTentative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


