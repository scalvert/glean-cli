# ActivityEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Universally unique identifier of the event. To allow for reliable retransmission, only the earliest received event of a given UUID is considered valid by the server and subsequent are ignored. | [optional] 
**Action** | **string** | The type of activity this represents. | 
**Params** | Pointer to [**ActivityEventParams**](ActivityEventParams.md) |  | [optional] 
**Timestamp** | **time.Time** | The ISO 8601 timestamp when the activity began. | 
**Url** | **string** | The URL of the activity. | 

## Methods

### NewActivityEvent

`func NewActivityEvent(action string, timestamp time.Time, url string, ) *ActivityEvent`

NewActivityEvent instantiates a new ActivityEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityEventWithDefaults

`func NewActivityEventWithDefaults() *ActivityEvent`

NewActivityEventWithDefaults instantiates a new ActivityEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAction

`func (o *ActivityEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActivityEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActivityEvent) SetAction(v string)`

SetAction sets Action field to given value.


### GetParams

`func (o *ActivityEvent) GetParams() ActivityEventParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ActivityEvent) GetParamsOk() (*ActivityEventParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ActivityEvent) SetParams(v ActivityEventParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *ActivityEvent) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetTimestamp

`func (o *ActivityEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUrl

`func (o *ActivityEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActivityEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActivityEvent) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


