# ShortcutInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogTimestamp** | Pointer to **int32** | Unix timestamp of the last activity processed to make the response (in seconds since epoch UTC). | [optional] 
**ShortcutInsights** | Pointer to [**[]ShortcutInsight**](ShortcutInsight.md) | Insights for shortcuts. | [optional] 
**Departments** | Pointer to **[]string** | list of departments applicable for shortcuts tab. | [optional] 
**MinVisitorThreshold** | Pointer to **int32** | Min threshold in number of visitors while populating results, otherwise 0. | [optional] 

## Methods

### NewShortcutInsightsResponse

`func NewShortcutInsightsResponse() *ShortcutInsightsResponse`

NewShortcutInsightsResponse instantiates a new ShortcutInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortcutInsightsResponseWithDefaults

`func NewShortcutInsightsResponseWithDefaults() *ShortcutInsightsResponse`

NewShortcutInsightsResponseWithDefaults instantiates a new ShortcutInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogTimestamp

`func (o *ShortcutInsightsResponse) GetLastLogTimestamp() int32`

GetLastLogTimestamp returns the LastLogTimestamp field if non-nil, zero value otherwise.

### GetLastLogTimestampOk

`func (o *ShortcutInsightsResponse) GetLastLogTimestampOk() (*int32, bool)`

GetLastLogTimestampOk returns a tuple with the LastLogTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogTimestamp

`func (o *ShortcutInsightsResponse) SetLastLogTimestamp(v int32)`

SetLastLogTimestamp sets LastLogTimestamp field to given value.

### HasLastLogTimestamp

`func (o *ShortcutInsightsResponse) HasLastLogTimestamp() bool`

HasLastLogTimestamp returns a boolean if a field has been set.

### GetShortcutInsights

`func (o *ShortcutInsightsResponse) GetShortcutInsights() []ShortcutInsight`

GetShortcutInsights returns the ShortcutInsights field if non-nil, zero value otherwise.

### GetShortcutInsightsOk

`func (o *ShortcutInsightsResponse) GetShortcutInsightsOk() (*[]ShortcutInsight, bool)`

GetShortcutInsightsOk returns a tuple with the ShortcutInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutInsights

`func (o *ShortcutInsightsResponse) SetShortcutInsights(v []ShortcutInsight)`

SetShortcutInsights sets ShortcutInsights field to given value.

### HasShortcutInsights

`func (o *ShortcutInsightsResponse) HasShortcutInsights() bool`

HasShortcutInsights returns a boolean if a field has been set.

### GetDepartments

`func (o *ShortcutInsightsResponse) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *ShortcutInsightsResponse) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *ShortcutInsightsResponse) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *ShortcutInsightsResponse) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.

### GetMinVisitorThreshold

`func (o *ShortcutInsightsResponse) GetMinVisitorThreshold() int32`

GetMinVisitorThreshold returns the MinVisitorThreshold field if non-nil, zero value otherwise.

### GetMinVisitorThresholdOk

`func (o *ShortcutInsightsResponse) GetMinVisitorThresholdOk() (*int32, bool)`

GetMinVisitorThresholdOk returns a tuple with the MinVisitorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVisitorThreshold

`func (o *ShortcutInsightsResponse) SetMinVisitorThreshold(v int32)`

SetMinVisitorThreshold sets MinVisitorThreshold field to given value.

### HasMinVisitorThreshold

`func (o *ShortcutInsightsResponse) HasMinVisitorThreshold() bool`

HasMinVisitorThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


