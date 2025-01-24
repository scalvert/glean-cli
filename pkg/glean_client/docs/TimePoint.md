# TimePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EpochSeconds** | Pointer to **int32** | Epoch seconds. Has precedence over daysFromNow. | [optional] 
**DaysFromNow** | Pointer to **int32** | The number of days from now. Specification relative to current time. Can be negative. | [optional] 

## Methods

### NewTimePoint

`func NewTimePoint() *TimePoint`

NewTimePoint instantiates a new TimePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimePointWithDefaults

`func NewTimePointWithDefaults() *TimePoint`

NewTimePointWithDefaults instantiates a new TimePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpochSeconds

`func (o *TimePoint) GetEpochSeconds() int32`

GetEpochSeconds returns the EpochSeconds field if non-nil, zero value otherwise.

### GetEpochSecondsOk

`func (o *TimePoint) GetEpochSecondsOk() (*int32, bool)`

GetEpochSecondsOk returns a tuple with the EpochSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochSeconds

`func (o *TimePoint) SetEpochSeconds(v int32)`

SetEpochSeconds sets EpochSeconds field to given value.

### HasEpochSeconds

`func (o *TimePoint) HasEpochSeconds() bool`

HasEpochSeconds returns a boolean if a field has been set.

### GetDaysFromNow

`func (o *TimePoint) GetDaysFromNow() int32`

GetDaysFromNow returns the DaysFromNow field if non-nil, zero value otherwise.

### GetDaysFromNowOk

`func (o *TimePoint) GetDaysFromNowOk() (*int32, bool)`

GetDaysFromNowOk returns a tuple with the DaysFromNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysFromNow

`func (o *TimePoint) SetDaysFromNow(v int32)`

SetDaysFromNow sets DaysFromNow field to given value.

### HasDaysFromNow

`func (o *TimePoint) HasDaysFromNow() bool`

HasDaysFromNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


