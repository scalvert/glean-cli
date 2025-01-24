# CalendarEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**EventType** | Pointer to **string** | The nature of the event, for example \&quot;out of office\&quot;. | [optional] 
**Id** | **string** | The calendar event id | 
**Url** | **string** | A permalink for this calendar event | 
**Attendees** | Pointer to [**CalendarAttendees**](CalendarAttendees.md) |  | [optional] 
**Location** | Pointer to **string** | The location that this event is taking place at. | [optional] 
**ConferenceData** | Pointer to [**ConferenceData**](ConferenceData.md) |  | [optional] 
**Description** | Pointer to **string** | The HTML description of the event. | [optional] 
**Datasource** | Pointer to **string** | The app or other repository type from which the event was extracted | [optional] 
**Classifications** | Pointer to [**[]EventClassification**](EventClassification.md) |  | [optional] 
**GeneratedAttachments** | Pointer to [**[]GeneratedAttachment**](GeneratedAttachment.md) |  | [optional] 

## Methods

### NewCalendarEvent

`func NewCalendarEvent(id string, url string, ) *CalendarEvent`

NewCalendarEvent instantiates a new CalendarEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarEventWithDefaults

`func NewCalendarEventWithDefaults() *CalendarEvent`

NewCalendarEventWithDefaults instantiates a new CalendarEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *CalendarEvent) GetTime() TimeInterval`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CalendarEvent) GetTimeOk() (*TimeInterval, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CalendarEvent) SetTime(v TimeInterval)`

SetTime sets Time field to given value.

### HasTime

`func (o *CalendarEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetEventType

`func (o *CalendarEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CalendarEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CalendarEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CalendarEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *CalendarEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalendarEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalendarEvent) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *CalendarEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CalendarEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CalendarEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAttendees

`func (o *CalendarEvent) GetAttendees() CalendarAttendees`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *CalendarEvent) GetAttendeesOk() (*CalendarAttendees, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendees

`func (o *CalendarEvent) SetAttendees(v CalendarAttendees)`

SetAttendees sets Attendees field to given value.

### HasAttendees

`func (o *CalendarEvent) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### GetLocation

`func (o *CalendarEvent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CalendarEvent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CalendarEvent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CalendarEvent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetConferenceData

`func (o *CalendarEvent) GetConferenceData() ConferenceData`

GetConferenceData returns the ConferenceData field if non-nil, zero value otherwise.

### GetConferenceDataOk

`func (o *CalendarEvent) GetConferenceDataOk() (*ConferenceData, bool)`

GetConferenceDataOk returns a tuple with the ConferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceData

`func (o *CalendarEvent) SetConferenceData(v ConferenceData)`

SetConferenceData sets ConferenceData field to given value.

### HasConferenceData

`func (o *CalendarEvent) HasConferenceData() bool`

HasConferenceData returns a boolean if a field has been set.

### GetDescription

`func (o *CalendarEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CalendarEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CalendarEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CalendarEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDatasource

`func (o *CalendarEvent) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *CalendarEvent) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *CalendarEvent) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *CalendarEvent) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetClassifications

`func (o *CalendarEvent) GetClassifications() []EventClassification`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *CalendarEvent) GetClassificationsOk() (*[]EventClassification, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *CalendarEvent) SetClassifications(v []EventClassification)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *CalendarEvent) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetGeneratedAttachments

`func (o *CalendarEvent) GetGeneratedAttachments() []GeneratedAttachment`

GetGeneratedAttachments returns the GeneratedAttachments field if non-nil, zero value otherwise.

### GetGeneratedAttachmentsOk

`func (o *CalendarEvent) GetGeneratedAttachmentsOk() (*[]GeneratedAttachment, bool)`

GetGeneratedAttachmentsOk returns a tuple with the GeneratedAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAttachments

`func (o *CalendarEvent) SetGeneratedAttachments(v []GeneratedAttachment)`

SetGeneratedAttachments sets GeneratedAttachments field to given value.

### HasGeneratedAttachments

`func (o *CalendarEvent) HasGeneratedAttachments() bool`

HasGeneratedAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


