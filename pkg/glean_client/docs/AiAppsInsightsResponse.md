# AiAppsInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogTimestamp** | Pointer to **int32** | Unix timestamp of the last activity processed to make the response (in seconds since epoch UTC). | [optional] 
**AiAppInsights** | Pointer to [**[]UserActivityInsight**](UserActivityInsight.md) |  | [optional] 
**TotalActiveUsers** | Pointer to **int32** | Total number of active users on the Ai App in the requested period. | [optional] 
**ActionCounts** | Pointer to [**AiAppActionCounts**](AiAppActionCounts.md) |  | [optional] 
**Departments** | Pointer to **[]string** | list of departments applicable for users tab. | [optional] 

## Methods

### NewAiAppsInsightsResponse

`func NewAiAppsInsightsResponse() *AiAppsInsightsResponse`

NewAiAppsInsightsResponse instantiates a new AiAppsInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAppsInsightsResponseWithDefaults

`func NewAiAppsInsightsResponseWithDefaults() *AiAppsInsightsResponse`

NewAiAppsInsightsResponseWithDefaults instantiates a new AiAppsInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogTimestamp

`func (o *AiAppsInsightsResponse) GetLastLogTimestamp() int32`

GetLastLogTimestamp returns the LastLogTimestamp field if non-nil, zero value otherwise.

### GetLastLogTimestampOk

`func (o *AiAppsInsightsResponse) GetLastLogTimestampOk() (*int32, bool)`

GetLastLogTimestampOk returns a tuple with the LastLogTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogTimestamp

`func (o *AiAppsInsightsResponse) SetLastLogTimestamp(v int32)`

SetLastLogTimestamp sets LastLogTimestamp field to given value.

### HasLastLogTimestamp

`func (o *AiAppsInsightsResponse) HasLastLogTimestamp() bool`

HasLastLogTimestamp returns a boolean if a field has been set.

### GetAiAppInsights

`func (o *AiAppsInsightsResponse) GetAiAppInsights() []UserActivityInsight`

GetAiAppInsights returns the AiAppInsights field if non-nil, zero value otherwise.

### GetAiAppInsightsOk

`func (o *AiAppsInsightsResponse) GetAiAppInsightsOk() (*[]UserActivityInsight, bool)`

GetAiAppInsightsOk returns a tuple with the AiAppInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAppInsights

`func (o *AiAppsInsightsResponse) SetAiAppInsights(v []UserActivityInsight)`

SetAiAppInsights sets AiAppInsights field to given value.

### HasAiAppInsights

`func (o *AiAppsInsightsResponse) HasAiAppInsights() bool`

HasAiAppInsights returns a boolean if a field has been set.

### GetTotalActiveUsers

`func (o *AiAppsInsightsResponse) GetTotalActiveUsers() int32`

GetTotalActiveUsers returns the TotalActiveUsers field if non-nil, zero value otherwise.

### GetTotalActiveUsersOk

`func (o *AiAppsInsightsResponse) GetTotalActiveUsersOk() (*int32, bool)`

GetTotalActiveUsersOk returns a tuple with the TotalActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActiveUsers

`func (o *AiAppsInsightsResponse) SetTotalActiveUsers(v int32)`

SetTotalActiveUsers sets TotalActiveUsers field to given value.

### HasTotalActiveUsers

`func (o *AiAppsInsightsResponse) HasTotalActiveUsers() bool`

HasTotalActiveUsers returns a boolean if a field has been set.

### GetActionCounts

`func (o *AiAppsInsightsResponse) GetActionCounts() AiAppActionCounts`

GetActionCounts returns the ActionCounts field if non-nil, zero value otherwise.

### GetActionCountsOk

`func (o *AiAppsInsightsResponse) GetActionCountsOk() (*AiAppActionCounts, bool)`

GetActionCountsOk returns a tuple with the ActionCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCounts

`func (o *AiAppsInsightsResponse) SetActionCounts(v AiAppActionCounts)`

SetActionCounts sets ActionCounts field to given value.

### HasActionCounts

`func (o *AiAppsInsightsResponse) HasActionCounts() bool`

HasActionCounts returns a boolean if a field has been set.

### GetDepartments

`func (o *AiAppsInsightsResponse) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *AiAppsInsightsResponse) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *AiAppsInsightsResponse) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *AiAppsInsightsResponse) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


