# RecommendationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingToken** | Pointer to **string** | A token that should be passed for additional requests related to this request (such as more results requests). | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**Results** | Pointer to [**[]SearchResult**](SearchResult.md) |  | [optional] 
**StructuredResults** | Pointer to [**[]StructuredResult**](StructuredResult.md) |  | [optional] 
**GeneratedQnaResult** | Pointer to [**GeneratedQna**](GeneratedQna.md) |  | [optional] 
**ErrorInfo** | Pointer to [**ErrorInfo**](ErrorInfo.md) |  | [optional] 
**RequestID** | Pointer to **string** | A platform-generated request ID to correlate backend logs. | [optional] 
**BackendTimeMillis** | Pointer to **int64** | Time in milliseconds the backend took to respond to the request. | [optional] 

## Methods

### NewRecommendationsResponse

`func NewRecommendationsResponse() *RecommendationsResponse`

NewRecommendationsResponse instantiates a new RecommendationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationsResponseWithDefaults

`func NewRecommendationsResponseWithDefaults() *RecommendationsResponse`

NewRecommendationsResponseWithDefaults instantiates a new RecommendationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingToken

`func (o *RecommendationsResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *RecommendationsResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *RecommendationsResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *RecommendationsResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetSessionInfo

`func (o *RecommendationsResponse) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *RecommendationsResponse) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *RecommendationsResponse) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *RecommendationsResponse) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetResults

`func (o *RecommendationsResponse) GetResults() []SearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RecommendationsResponse) GetResultsOk() (*[]SearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RecommendationsResponse) SetResults(v []SearchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *RecommendationsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStructuredResults

`func (o *RecommendationsResponse) GetStructuredResults() []StructuredResult`

GetStructuredResults returns the StructuredResults field if non-nil, zero value otherwise.

### GetStructuredResultsOk

`func (o *RecommendationsResponse) GetStructuredResultsOk() (*[]StructuredResult, bool)`

GetStructuredResultsOk returns a tuple with the StructuredResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredResults

`func (o *RecommendationsResponse) SetStructuredResults(v []StructuredResult)`

SetStructuredResults sets StructuredResults field to given value.

### HasStructuredResults

`func (o *RecommendationsResponse) HasStructuredResults() bool`

HasStructuredResults returns a boolean if a field has been set.

### GetGeneratedQnaResult

`func (o *RecommendationsResponse) GetGeneratedQnaResult() GeneratedQna`

GetGeneratedQnaResult returns the GeneratedQnaResult field if non-nil, zero value otherwise.

### GetGeneratedQnaResultOk

`func (o *RecommendationsResponse) GetGeneratedQnaResultOk() (*GeneratedQna, bool)`

GetGeneratedQnaResultOk returns a tuple with the GeneratedQnaResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedQnaResult

`func (o *RecommendationsResponse) SetGeneratedQnaResult(v GeneratedQna)`

SetGeneratedQnaResult sets GeneratedQnaResult field to given value.

### HasGeneratedQnaResult

`func (o *RecommendationsResponse) HasGeneratedQnaResult() bool`

HasGeneratedQnaResult returns a boolean if a field has been set.

### GetErrorInfo

`func (o *RecommendationsResponse) GetErrorInfo() ErrorInfo`

GetErrorInfo returns the ErrorInfo field if non-nil, zero value otherwise.

### GetErrorInfoOk

`func (o *RecommendationsResponse) GetErrorInfoOk() (*ErrorInfo, bool)`

GetErrorInfoOk returns a tuple with the ErrorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorInfo

`func (o *RecommendationsResponse) SetErrorInfo(v ErrorInfo)`

SetErrorInfo sets ErrorInfo field to given value.

### HasErrorInfo

`func (o *RecommendationsResponse) HasErrorInfo() bool`

HasErrorInfo returns a boolean if a field has been set.

### GetRequestID

`func (o *RecommendationsResponse) GetRequestID() string`

GetRequestID returns the RequestID field if non-nil, zero value otherwise.

### GetRequestIDOk

`func (o *RecommendationsResponse) GetRequestIDOk() (*string, bool)`

GetRequestIDOk returns a tuple with the RequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestID

`func (o *RecommendationsResponse) SetRequestID(v string)`

SetRequestID sets RequestID field to given value.

### HasRequestID

`func (o *RecommendationsResponse) HasRequestID() bool`

HasRequestID returns a boolean if a field has been set.

### GetBackendTimeMillis

`func (o *RecommendationsResponse) GetBackendTimeMillis() int64`

GetBackendTimeMillis returns the BackendTimeMillis field if non-nil, zero value otherwise.

### GetBackendTimeMillisOk

`func (o *RecommendationsResponse) GetBackendTimeMillisOk() (*int64, bool)`

GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTimeMillis

`func (o *RecommendationsResponse) SetBackendTimeMillis(v int64)`

SetBackendTimeMillis sets BackendTimeMillis field to given value.

### HasBackendTimeMillis

`func (o *RecommendationsResponse) HasBackendTimeMillis() bool`

HasBackendTimeMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


