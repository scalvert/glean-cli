# UserInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastLogTimestamp** | Pointer to **int32** | Unix timestamp of the last activity processed to make the response (in seconds since epoch UTC). | [optional] 
**ActivityInsights** | Pointer to [**[]UserActivityInsight**](UserActivityInsight.md) | Insights for all active users with respect to set of actions. | [optional] 
**InactiveInsights** | Pointer to [**[]UserActivityInsight**](UserActivityInsight.md) | Insights for all in inactive users with respect to set of actions and time period. Activity count will be set to 0. | [optional] 
**TotalTeammates** | Pointer to **int32** | Total number of teammates that have logged in to the product, that are still valid teammates. | [optional] 
**TotalActiveUsers** | Pointer to **int32** | Total number of active users in the requested period. | [optional] 
**Departments** | Pointer to **[]string** | list of departments applicable for users tab. | [optional] 

## Methods

### NewUserInsightsResponse

`func NewUserInsightsResponse() *UserInsightsResponse`

NewUserInsightsResponse instantiates a new UserInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInsightsResponseWithDefaults

`func NewUserInsightsResponseWithDefaults() *UserInsightsResponse`

NewUserInsightsResponseWithDefaults instantiates a new UserInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastLogTimestamp

`func (o *UserInsightsResponse) GetLastLogTimestamp() int32`

GetLastLogTimestamp returns the LastLogTimestamp field if non-nil, zero value otherwise.

### GetLastLogTimestampOk

`func (o *UserInsightsResponse) GetLastLogTimestampOk() (*int32, bool)`

GetLastLogTimestampOk returns a tuple with the LastLogTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogTimestamp

`func (o *UserInsightsResponse) SetLastLogTimestamp(v int32)`

SetLastLogTimestamp sets LastLogTimestamp field to given value.

### HasLastLogTimestamp

`func (o *UserInsightsResponse) HasLastLogTimestamp() bool`

HasLastLogTimestamp returns a boolean if a field has been set.

### GetActivityInsights

`func (o *UserInsightsResponse) GetActivityInsights() []UserActivityInsight`

GetActivityInsights returns the ActivityInsights field if non-nil, zero value otherwise.

### GetActivityInsightsOk

`func (o *UserInsightsResponse) GetActivityInsightsOk() (*[]UserActivityInsight, bool)`

GetActivityInsightsOk returns a tuple with the ActivityInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityInsights

`func (o *UserInsightsResponse) SetActivityInsights(v []UserActivityInsight)`

SetActivityInsights sets ActivityInsights field to given value.

### HasActivityInsights

`func (o *UserInsightsResponse) HasActivityInsights() bool`

HasActivityInsights returns a boolean if a field has been set.

### GetInactiveInsights

`func (o *UserInsightsResponse) GetInactiveInsights() []UserActivityInsight`

GetInactiveInsights returns the InactiveInsights field if non-nil, zero value otherwise.

### GetInactiveInsightsOk

`func (o *UserInsightsResponse) GetInactiveInsightsOk() (*[]UserActivityInsight, bool)`

GetInactiveInsightsOk returns a tuple with the InactiveInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveInsights

`func (o *UserInsightsResponse) SetInactiveInsights(v []UserActivityInsight)`

SetInactiveInsights sets InactiveInsights field to given value.

### HasInactiveInsights

`func (o *UserInsightsResponse) HasInactiveInsights() bool`

HasInactiveInsights returns a boolean if a field has been set.

### GetTotalTeammates

`func (o *UserInsightsResponse) GetTotalTeammates() int32`

GetTotalTeammates returns the TotalTeammates field if non-nil, zero value otherwise.

### GetTotalTeammatesOk

`func (o *UserInsightsResponse) GetTotalTeammatesOk() (*int32, bool)`

GetTotalTeammatesOk returns a tuple with the TotalTeammates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTeammates

`func (o *UserInsightsResponse) SetTotalTeammates(v int32)`

SetTotalTeammates sets TotalTeammates field to given value.

### HasTotalTeammates

`func (o *UserInsightsResponse) HasTotalTeammates() bool`

HasTotalTeammates returns a boolean if a field has been set.

### GetTotalActiveUsers

`func (o *UserInsightsResponse) GetTotalActiveUsers() int32`

GetTotalActiveUsers returns the TotalActiveUsers field if non-nil, zero value otherwise.

### GetTotalActiveUsersOk

`func (o *UserInsightsResponse) GetTotalActiveUsersOk() (*int32, bool)`

GetTotalActiveUsersOk returns a tuple with the TotalActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActiveUsers

`func (o *UserInsightsResponse) SetTotalActiveUsers(v int32)`

SetTotalActiveUsers sets TotalActiveUsers field to given value.

### HasTotalActiveUsers

`func (o *UserInsightsResponse) HasTotalActiveUsers() bool`

HasTotalActiveUsers returns a boolean if a field has been set.

### GetDepartments

`func (o *UserInsightsResponse) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *UserInsightsResponse) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *UserInsightsResponse) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *UserInsightsResponse) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


