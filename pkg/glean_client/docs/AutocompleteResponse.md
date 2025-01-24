# AutocompleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentIds** | Pointer to **[]int64** | List of experiment ids for the corresponding request. | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular set of autocomplete results. To be used for /feedback reporting. | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**Results** | Pointer to [**[]AutocompleteResult**](AutocompleteResult.md) |  | [optional] 
**Groups** | Pointer to [**[]AutocompleteResultGroup**](AutocompleteResultGroup.md) | Subsections of the results list from which distinct sections should be created. | [optional] 
**ErrorInfo** | Pointer to [**ErrorInfo**](ErrorInfo.md) |  | [optional] 
**BackendTimeMillis** | Pointer to **int64** | Time in milliseconds the backend took to respond to the request. | [optional] 

## Methods

### NewAutocompleteResponse

`func NewAutocompleteResponse() *AutocompleteResponse`

NewAutocompleteResponse instantiates a new AutocompleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteResponseWithDefaults

`func NewAutocompleteResponseWithDefaults() *AutocompleteResponse`

NewAutocompleteResponseWithDefaults instantiates a new AutocompleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentIds

`func (o *AutocompleteResponse) GetExperimentIds() []int64`

GetExperimentIds returns the ExperimentIds field if non-nil, zero value otherwise.

### GetExperimentIdsOk

`func (o *AutocompleteResponse) GetExperimentIdsOk() (*[]int64, bool)`

GetExperimentIdsOk returns a tuple with the ExperimentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIds

`func (o *AutocompleteResponse) SetExperimentIds(v []int64)`

SetExperimentIds sets ExperimentIds field to given value.

### HasExperimentIds

`func (o *AutocompleteResponse) HasExperimentIds() bool`

HasExperimentIds returns a boolean if a field has been set.

### GetTrackingToken

`func (o *AutocompleteResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *AutocompleteResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *AutocompleteResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *AutocompleteResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetSessionInfo

`func (o *AutocompleteResponse) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *AutocompleteResponse) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *AutocompleteResponse) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *AutocompleteResponse) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetResults

`func (o *AutocompleteResponse) GetResults() []AutocompleteResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AutocompleteResponse) GetResultsOk() (*[]AutocompleteResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AutocompleteResponse) SetResults(v []AutocompleteResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *AutocompleteResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetGroups

`func (o *AutocompleteResponse) GetGroups() []AutocompleteResultGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AutocompleteResponse) GetGroupsOk() (*[]AutocompleteResultGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AutocompleteResponse) SetGroups(v []AutocompleteResultGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AutocompleteResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetErrorInfo

`func (o *AutocompleteResponse) GetErrorInfo() ErrorInfo`

GetErrorInfo returns the ErrorInfo field if non-nil, zero value otherwise.

### GetErrorInfoOk

`func (o *AutocompleteResponse) GetErrorInfoOk() (*ErrorInfo, bool)`

GetErrorInfoOk returns a tuple with the ErrorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorInfo

`func (o *AutocompleteResponse) SetErrorInfo(v ErrorInfo)`

SetErrorInfo sets ErrorInfo field to given value.

### HasErrorInfo

`func (o *AutocompleteResponse) HasErrorInfo() bool`

HasErrorInfo returns a boolean if a field has been set.

### GetBackendTimeMillis

`func (o *AutocompleteResponse) GetBackendTimeMillis() int64`

GetBackendTimeMillis returns the BackendTimeMillis field if non-nil, zero value otherwise.

### GetBackendTimeMillisOk

`func (o *AutocompleteResponse) GetBackendTimeMillisOk() (*int64, bool)`

GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTimeMillis

`func (o *AutocompleteResponse) SetBackendTimeMillis(v int64)`

SetBackendTimeMillis sets BackendTimeMillis field to given value.

### HasBackendTimeMillis

`func (o *AutocompleteResponse) HasBackendTimeMillis() bool`

HasBackendTimeMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


