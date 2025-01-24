# QueryInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogTimestamp** | Pointer to **int32** | Unix timestamp of the last activity processed to make the response (in seconds since epoch UTC). | [optional] 
**QueryInsights** | Pointer to [**[]QueryInsight**](QueryInsight.md) | Insights for queries. | [optional] 
**LowPerformingQueryInsights** | Pointer to [**[]QueryInsight**](QueryInsight.md) | Insights for low performing queries without good results. | [optional] 
**Departments** | Pointer to **[]string** | list of departments applicable for queries tab. | [optional] 
**MinVisitorThreshold** | Pointer to **int32** | Min threshold in number of visitors while populating results, otherwise 0. | [optional] 

## Methods

### NewQueryInsightsResponse

`func NewQueryInsightsResponse() *QueryInsightsResponse`

NewQueryInsightsResponse instantiates a new QueryInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInsightsResponseWithDefaults

`func NewQueryInsightsResponseWithDefaults() *QueryInsightsResponse`

NewQueryInsightsResponseWithDefaults instantiates a new QueryInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogTimestamp

`func (o *QueryInsightsResponse) GetLastLogTimestamp() int32`

GetLastLogTimestamp returns the LastLogTimestamp field if non-nil, zero value otherwise.

### GetLastLogTimestampOk

`func (o *QueryInsightsResponse) GetLastLogTimestampOk() (*int32, bool)`

GetLastLogTimestampOk returns a tuple with the LastLogTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogTimestamp

`func (o *QueryInsightsResponse) SetLastLogTimestamp(v int32)`

SetLastLogTimestamp sets LastLogTimestamp field to given value.

### HasLastLogTimestamp

`func (o *QueryInsightsResponse) HasLastLogTimestamp() bool`

HasLastLogTimestamp returns a boolean if a field has been set.

### GetQueryInsights

`func (o *QueryInsightsResponse) GetQueryInsights() []QueryInsight`

GetQueryInsights returns the QueryInsights field if non-nil, zero value otherwise.

### GetQueryInsightsOk

`func (o *QueryInsightsResponse) GetQueryInsightsOk() (*[]QueryInsight, bool)`

GetQueryInsightsOk returns a tuple with the QueryInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryInsights

`func (o *QueryInsightsResponse) SetQueryInsights(v []QueryInsight)`

SetQueryInsights sets QueryInsights field to given value.

### HasQueryInsights

`func (o *QueryInsightsResponse) HasQueryInsights() bool`

HasQueryInsights returns a boolean if a field has been set.

### GetLowPerformingQueryInsights

`func (o *QueryInsightsResponse) GetLowPerformingQueryInsights() []QueryInsight`

GetLowPerformingQueryInsights returns the LowPerformingQueryInsights field if non-nil, zero value otherwise.

### GetLowPerformingQueryInsightsOk

`func (o *QueryInsightsResponse) GetLowPerformingQueryInsightsOk() (*[]QueryInsight, bool)`

GetLowPerformingQueryInsightsOk returns a tuple with the LowPerformingQueryInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPerformingQueryInsights

`func (o *QueryInsightsResponse) SetLowPerformingQueryInsights(v []QueryInsight)`

SetLowPerformingQueryInsights sets LowPerformingQueryInsights field to given value.

### HasLowPerformingQueryInsights

`func (o *QueryInsightsResponse) HasLowPerformingQueryInsights() bool`

HasLowPerformingQueryInsights returns a boolean if a field has been set.

### GetDepartments

`func (o *QueryInsightsResponse) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *QueryInsightsResponse) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *QueryInsightsResponse) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *QueryInsightsResponse) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.

### GetMinVisitorThreshold

`func (o *QueryInsightsResponse) GetMinVisitorThreshold() int32`

GetMinVisitorThreshold returns the MinVisitorThreshold field if non-nil, zero value otherwise.

### GetMinVisitorThresholdOk

`func (o *QueryInsightsResponse) GetMinVisitorThresholdOk() (*int32, bool)`

GetMinVisitorThresholdOk returns a tuple with the MinVisitorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVisitorThreshold

`func (o *QueryInsightsResponse) SetMinVisitorThreshold(v int32)`

SetMinVisitorThreshold sets MinVisitorThreshold field to given value.

### HasMinVisitorThreshold

`func (o *QueryInsightsResponse) HasMinVisitorThreshold() bool`

HasMinVisitorThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


