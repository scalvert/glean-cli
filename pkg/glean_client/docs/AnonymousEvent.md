# AnonymousEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to [**TimeInterval**](TimeInterval.md) |  | [optional] 
**EventType** | Pointer to **string** | The nature of the event, for example \&quot;out of office\&quot;. | [optional] 

## Methods

### NewAnonymousEvent

`func NewAnonymousEvent() *AnonymousEvent`

NewAnonymousEvent instantiates a new AnonymousEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnonymousEventWithDefaults

`func NewAnonymousEventWithDefaults() *AnonymousEvent`

NewAnonymousEventWithDefaults instantiates a new AnonymousEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *AnonymousEvent) GetTime() TimeInterval`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AnonymousEvent) GetTimeOk() (*TimeInterval, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AnonymousEvent) SetTime(v TimeInterval)`

SetTime sets Time field to given value.

### HasTime

`func (o *AnonymousEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetEventType

`func (o *AnonymousEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *AnonymousEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *AnonymousEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *AnonymousEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


