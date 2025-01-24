# CountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | The counter value | 
**Period** | Pointer to [**Period**](Period.md) |  | [optional] 
**Org** | Pointer to **string** | The unit of organization over which we did the count aggregation, e.g. org (department) or company | [optional] 

## Methods

### NewCountInfo

`func NewCountInfo(count int32, ) *CountInfo`

NewCountInfo instantiates a new CountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountInfoWithDefaults

`func NewCountInfoWithDefaults() *CountInfo`

NewCountInfoWithDefaults instantiates a new CountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CountInfo) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CountInfo) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CountInfo) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPeriod

`func (o *CountInfo) GetPeriod() Period`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CountInfo) GetPeriodOk() (*Period, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CountInfo) SetPeriod(v Period)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CountInfo) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetOrg

`func (o *CountInfo) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CountInfo) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CountInfo) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CountInfo) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


