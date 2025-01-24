# Period

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinDaysFromNow** | Pointer to **int32** | DEPRECATED - The number of days from now in the past to define upper boundary of time period. | [optional] 
**MaxDaysFromNow** | Pointer to **int32** | DEPRECATED - The number of days from now in the past to define lower boundary of time period. | [optional] 
**Start** | Pointer to [**TimePoint**](TimePoint.md) |  | [optional] 
**End** | Pointer to [**TimePoint**](TimePoint.md) |  | [optional] 

## Methods

### NewPeriod

`func NewPeriod() *Period`

NewPeriod instantiates a new Period object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeriodWithDefaults

`func NewPeriodWithDefaults() *Period`

NewPeriodWithDefaults instantiates a new Period object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinDaysFromNow

`func (o *Period) GetMinDaysFromNow() int32`

GetMinDaysFromNow returns the MinDaysFromNow field if non-nil, zero value otherwise.

### GetMinDaysFromNowOk

`func (o *Period) GetMinDaysFromNowOk() (*int32, bool)`

GetMinDaysFromNowOk returns a tuple with the MinDaysFromNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDaysFromNow

`func (o *Period) SetMinDaysFromNow(v int32)`

SetMinDaysFromNow sets MinDaysFromNow field to given value.

### HasMinDaysFromNow

`func (o *Period) HasMinDaysFromNow() bool`

HasMinDaysFromNow returns a boolean if a field has been set.

### GetMaxDaysFromNow

`func (o *Period) GetMaxDaysFromNow() int32`

GetMaxDaysFromNow returns the MaxDaysFromNow field if non-nil, zero value otherwise.

### GetMaxDaysFromNowOk

`func (o *Period) GetMaxDaysFromNowOk() (*int32, bool)`

GetMaxDaysFromNowOk returns a tuple with the MaxDaysFromNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDaysFromNow

`func (o *Period) SetMaxDaysFromNow(v int32)`

SetMaxDaysFromNow sets MaxDaysFromNow field to given value.

### HasMaxDaysFromNow

`func (o *Period) HasMaxDaysFromNow() bool`

HasMaxDaysFromNow returns a boolean if a field has been set.

### GetStart

`func (o *Period) GetStart() TimePoint`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Period) GetStartOk() (*TimePoint, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Period) SetStart(v TimePoint)`

SetStart sets Start field to given value.

### HasStart

`func (o *Period) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *Period) GetEnd() TimePoint`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Period) GetEndOk() (*TimePoint, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Period) SetEnd(v TimePoint)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Period) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


