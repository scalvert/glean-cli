# TimeInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **string** | The RFC3339 timestamp formatted start time of this event. | 
**End** | **string** | The RFC3339 timestamp formatted end time of this event. | 

## Methods

### NewTimeInterval

`func NewTimeInterval(start string, end string, ) *TimeInterval`

NewTimeInterval instantiates a new TimeInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeIntervalWithDefaults

`func NewTimeIntervalWithDefaults() *TimeInterval`

NewTimeIntervalWithDefaults instantiates a new TimeInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TimeInterval) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TimeInterval) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TimeInterval) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *TimeInterval) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TimeInterval) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TimeInterval) SetEnd(v string)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


