# UserActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to [**Person**](Person.md) |  | [optional] 
**Timestamp** | Pointer to **int32** | Unix timestamp of the activity (in seconds since epoch UTC). | [optional] 
**Action** | Pointer to **string** | The action for the activity | [optional] 
**AggregateVisitCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 

## Methods

### NewUserActivity

`func NewUserActivity() *UserActivity`

NewUserActivity instantiates a new UserActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityWithDefaults

`func NewUserActivityWithDefaults() *UserActivity`

NewUserActivityWithDefaults instantiates a new UserActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *UserActivity) GetActor() Person`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *UserActivity) GetActorOk() (*Person, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *UserActivity) SetActor(v Person)`

SetActor sets Actor field to given value.

### HasActor

`func (o *UserActivity) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetTimestamp

`func (o *UserActivity) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserActivity) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserActivity) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UserActivity) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAction

`func (o *UserActivity) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserActivity) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserActivity) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserActivity) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAggregateVisitCount

`func (o *UserActivity) GetAggregateVisitCount() CountInfo`

GetAggregateVisitCount returns the AggregateVisitCount field if non-nil, zero value otherwise.

### GetAggregateVisitCountOk

`func (o *UserActivity) GetAggregateVisitCountOk() (*CountInfo, bool)`

GetAggregateVisitCountOk returns a tuple with the AggregateVisitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateVisitCount

`func (o *UserActivity) SetAggregateVisitCount(v CountInfo)`

SetAggregateVisitCount sets AggregateVisitCount field to given value.

### HasAggregateVisitCount

`func (o *UserActivity) HasAggregateVisitCount() bool`

HasAggregateVisitCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


