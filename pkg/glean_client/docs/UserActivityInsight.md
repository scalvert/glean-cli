# UserActivityInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**Person**](Person.md) |  | 
**Activity** | **string** | Activity e.g. search, home page visit or all. | 
**LastActivityTimestamp** | Pointer to **int32** | Unix timestamp of the last activity (in seconds since epoch UTC). | [optional] 
**ActivityCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 
**ActiveDayCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 

## Methods

### NewUserActivityInsight

`func NewUserActivityInsight(user Person, activity string, ) *UserActivityInsight`

NewUserActivityInsight instantiates a new UserActivityInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityInsightWithDefaults

`func NewUserActivityInsightWithDefaults() *UserActivityInsight`

NewUserActivityInsightWithDefaults instantiates a new UserActivityInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserActivityInsight) GetUser() Person`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserActivityInsight) GetUserOk() (*Person, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserActivityInsight) SetUser(v Person)`

SetUser sets User field to given value.


### GetActivity

`func (o *UserActivityInsight) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *UserActivityInsight) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *UserActivityInsight) SetActivity(v string)`

SetActivity sets Activity field to given value.


### GetLastActivityTimestamp

`func (o *UserActivityInsight) GetLastActivityTimestamp() int32`

GetLastActivityTimestamp returns the LastActivityTimestamp field if non-nil, zero value otherwise.

### GetLastActivityTimestampOk

`func (o *UserActivityInsight) GetLastActivityTimestampOk() (*int32, bool)`

GetLastActivityTimestampOk returns a tuple with the LastActivityTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityTimestamp

`func (o *UserActivityInsight) SetLastActivityTimestamp(v int32)`

SetLastActivityTimestamp sets LastActivityTimestamp field to given value.

### HasLastActivityTimestamp

`func (o *UserActivityInsight) HasLastActivityTimestamp() bool`

HasLastActivityTimestamp returns a boolean if a field has been set.

### GetActivityCount

`func (o *UserActivityInsight) GetActivityCount() CountInfo`

GetActivityCount returns the ActivityCount field if non-nil, zero value otherwise.

### GetActivityCountOk

`func (o *UserActivityInsight) GetActivityCountOk() (*CountInfo, bool)`

GetActivityCountOk returns a tuple with the ActivityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCount

`func (o *UserActivityInsight) SetActivityCount(v CountInfo)`

SetActivityCount sets ActivityCount field to given value.

### HasActivityCount

`func (o *UserActivityInsight) HasActivityCount() bool`

HasActivityCount returns a boolean if a field has been set.

### GetActiveDayCount

`func (o *UserActivityInsight) GetActiveDayCount() CountInfo`

GetActiveDayCount returns the ActiveDayCount field if non-nil, zero value otherwise.

### GetActiveDayCountOk

`func (o *UserActivityInsight) GetActiveDayCountOk() (*CountInfo, bool)`

GetActiveDayCountOk returns a tuple with the ActiveDayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDayCount

`func (o *UserActivityInsight) SetActiveDayCount(v CountInfo)`

SetActiveDayCount sets ActiveDayCount field to given value.

### HasActiveDayCount

`func (o *UserActivityInsight) HasActiveDayCount() bool`

HasActiveDayCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


