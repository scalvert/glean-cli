# CalendarAttendee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOrganizer** | Pointer to **bool** | Whether or not this attendee is an organizer. | [optional] 
**IsInGroup** | Pointer to **bool** | Whether or not this attendee is in a group. Needed temporarily at least to support both flat attendees and tree for compatibility. | [optional] 
**Person** | [**Person**](Person.md) |  | 
**GroupAttendees** | Pointer to [**[]CalendarAttendee**](CalendarAttendee.md) | If this attendee is a group, represents the list of individual attendees in the group. | [optional] 
**ResponseStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewCalendarAttendee

`func NewCalendarAttendee(person Person, ) *CalendarAttendee`

NewCalendarAttendee instantiates a new CalendarAttendee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarAttendeeWithDefaults

`func NewCalendarAttendeeWithDefaults() *CalendarAttendee`

NewCalendarAttendeeWithDefaults instantiates a new CalendarAttendee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOrganizer

`func (o *CalendarAttendee) GetIsOrganizer() bool`

GetIsOrganizer returns the IsOrganizer field if non-nil, zero value otherwise.

### GetIsOrganizerOk

`func (o *CalendarAttendee) GetIsOrganizerOk() (*bool, bool)`

GetIsOrganizerOk returns a tuple with the IsOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganizer

`func (o *CalendarAttendee) SetIsOrganizer(v bool)`

SetIsOrganizer sets IsOrganizer field to given value.

### HasIsOrganizer

`func (o *CalendarAttendee) HasIsOrganizer() bool`

HasIsOrganizer returns a boolean if a field has been set.

### GetIsInGroup

`func (o *CalendarAttendee) GetIsInGroup() bool`

GetIsInGroup returns the IsInGroup field if non-nil, zero value otherwise.

### GetIsInGroupOk

`func (o *CalendarAttendee) GetIsInGroupOk() (*bool, bool)`

GetIsInGroupOk returns a tuple with the IsInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInGroup

`func (o *CalendarAttendee) SetIsInGroup(v bool)`

SetIsInGroup sets IsInGroup field to given value.

### HasIsInGroup

`func (o *CalendarAttendee) HasIsInGroup() bool`

HasIsInGroup returns a boolean if a field has been set.

### GetPerson

`func (o *CalendarAttendee) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *CalendarAttendee) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *CalendarAttendee) SetPerson(v Person)`

SetPerson sets Person field to given value.


### GetGroupAttendees

`func (o *CalendarAttendee) GetGroupAttendees() []CalendarAttendee`

GetGroupAttendees returns the GroupAttendees field if non-nil, zero value otherwise.

### GetGroupAttendeesOk

`func (o *CalendarAttendee) GetGroupAttendeesOk() (*[]CalendarAttendee, bool)`

GetGroupAttendeesOk returns a tuple with the GroupAttendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttendees

`func (o *CalendarAttendee) SetGroupAttendees(v []CalendarAttendee)`

SetGroupAttendees sets GroupAttendees field to given value.

### HasGroupAttendees

`func (o *CalendarAttendee) HasGroupAttendees() bool`

HasGroupAttendees returns a boolean if a field has been set.

### GetResponseStatus

`func (o *CalendarAttendee) GetResponseStatus() string`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *CalendarAttendee) GetResponseStatusOk() (*string, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *CalendarAttendee) SetResponseStatus(v string)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *CalendarAttendee) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


