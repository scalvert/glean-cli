# ContentInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogTimestamp** | Pointer to **int32** | Unix timestamp of the last activity processed to make the response (in seconds since epoch UTC). | [optional] 
**DocumentInsights** | Pointer to [**[]DocumentInsight**](DocumentInsight.md) | Insights for documents. | [optional] 
**Departments** | Pointer to **[]string** | list of departments applicable for contents tab. | [optional] 
**MinDepartmentSizeThreshold** | Pointer to **int32** | Min threshold in size of departments while populating results, otherwise 0. | [optional] 
**MinVisitorThreshold** | Pointer to **int32** | Minimum number of visitors to a document required to be included in insights. | [optional] 

## Methods

### NewContentInsightsResponse

`func NewContentInsightsResponse() *ContentInsightsResponse`

NewContentInsightsResponse instantiates a new ContentInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentInsightsResponseWithDefaults

`func NewContentInsightsResponseWithDefaults() *ContentInsightsResponse`

NewContentInsightsResponseWithDefaults instantiates a new ContentInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogTimestamp

`func (o *ContentInsightsResponse) GetLastLogTimestamp() int32`

GetLastLogTimestamp returns the LastLogTimestamp field if non-nil, zero value otherwise.

### GetLastLogTimestampOk

`func (o *ContentInsightsResponse) GetLastLogTimestampOk() (*int32, bool)`

GetLastLogTimestampOk returns a tuple with the LastLogTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogTimestamp

`func (o *ContentInsightsResponse) SetLastLogTimestamp(v int32)`

SetLastLogTimestamp sets LastLogTimestamp field to given value.

### HasLastLogTimestamp

`func (o *ContentInsightsResponse) HasLastLogTimestamp() bool`

HasLastLogTimestamp returns a boolean if a field has been set.

### GetDocumentInsights

`func (o *ContentInsightsResponse) GetDocumentInsights() []DocumentInsight`

GetDocumentInsights returns the DocumentInsights field if non-nil, zero value otherwise.

### GetDocumentInsightsOk

`func (o *ContentInsightsResponse) GetDocumentInsightsOk() (*[]DocumentInsight, bool)`

GetDocumentInsightsOk returns a tuple with the DocumentInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentInsights

`func (o *ContentInsightsResponse) SetDocumentInsights(v []DocumentInsight)`

SetDocumentInsights sets DocumentInsights field to given value.

### HasDocumentInsights

`func (o *ContentInsightsResponse) HasDocumentInsights() bool`

HasDocumentInsights returns a boolean if a field has been set.

### GetDepartments

`func (o *ContentInsightsResponse) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *ContentInsightsResponse) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *ContentInsightsResponse) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *ContentInsightsResponse) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.

### GetMinDepartmentSizeThreshold

`func (o *ContentInsightsResponse) GetMinDepartmentSizeThreshold() int32`

GetMinDepartmentSizeThreshold returns the MinDepartmentSizeThreshold field if non-nil, zero value otherwise.

### GetMinDepartmentSizeThresholdOk

`func (o *ContentInsightsResponse) GetMinDepartmentSizeThresholdOk() (*int32, bool)`

GetMinDepartmentSizeThresholdOk returns a tuple with the MinDepartmentSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDepartmentSizeThreshold

`func (o *ContentInsightsResponse) SetMinDepartmentSizeThreshold(v int32)`

SetMinDepartmentSizeThreshold sets MinDepartmentSizeThreshold field to given value.

### HasMinDepartmentSizeThreshold

`func (o *ContentInsightsResponse) HasMinDepartmentSizeThreshold() bool`

HasMinDepartmentSizeThreshold returns a boolean if a field has been set.

### GetMinVisitorThreshold

`func (o *ContentInsightsResponse) GetMinVisitorThreshold() int32`

GetMinVisitorThreshold returns the MinVisitorThreshold field if non-nil, zero value otherwise.

### GetMinVisitorThresholdOk

`func (o *ContentInsightsResponse) GetMinVisitorThresholdOk() (*int32, bool)`

GetMinVisitorThresholdOk returns a tuple with the MinVisitorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVisitorThreshold

`func (o *ContentInsightsResponse) SetMinVisitorThreshold(v int32)`

SetMinVisitorThreshold sets MinVisitorThreshold field to given value.

### HasMinVisitorThreshold

`func (o *ContentInsightsResponse) HasMinVisitorThreshold() bool`

HasMinVisitorThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


