# ResultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The ISO 8601 timestamp associated with the client request. | [optional] 
**TrackingToken** | Pointer to **string** | A previously received trackingToken for a search associated with the same query. Useful for more requests and requests for other tabs. | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**SourceDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**PageSize** | Pointer to **int32** | Hint to the server about how many results to send back. Server may return less or more. Structured results and clustered results don&#39;t count towards pageSize. | [optional] 
**MaxSnippetSize** | Pointer to **int32** | Hint to the server about how many characters long a snippet may be. Server may return less or more. | [optional] 

## Methods

### NewResultsRequest

`func NewResultsRequest() *ResultsRequest`

NewResultsRequest instantiates a new ResultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsRequestWithDefaults

`func NewResultsRequestWithDefaults() *ResultsRequest`

NewResultsRequestWithDefaults instantiates a new ResultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ResultsRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ResultsRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ResultsRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ResultsRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTrackingToken

`func (o *ResultsRequest) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *ResultsRequest) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *ResultsRequest) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *ResultsRequest) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetSessionInfo

`func (o *ResultsRequest) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *ResultsRequest) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *ResultsRequest) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *ResultsRequest) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetSourceDocument

`func (o *ResultsRequest) GetSourceDocument() Document`

GetSourceDocument returns the SourceDocument field if non-nil, zero value otherwise.

### GetSourceDocumentOk

`func (o *ResultsRequest) GetSourceDocumentOk() (*Document, bool)`

GetSourceDocumentOk returns a tuple with the SourceDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocument

`func (o *ResultsRequest) SetSourceDocument(v Document)`

SetSourceDocument sets SourceDocument field to given value.

### HasSourceDocument

`func (o *ResultsRequest) HasSourceDocument() bool`

HasSourceDocument returns a boolean if a field has been set.

### GetPageSize

`func (o *ResultsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ResultsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ResultsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ResultsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetMaxSnippetSize

`func (o *ResultsRequest) GetMaxSnippetSize() int32`

GetMaxSnippetSize returns the MaxSnippetSize field if non-nil, zero value otherwise.

### GetMaxSnippetSizeOk

`func (o *ResultsRequest) GetMaxSnippetSizeOk() (*int32, bool)`

GetMaxSnippetSizeOk returns a tuple with the MaxSnippetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnippetSize

`func (o *ResultsRequest) SetMaxSnippetSize(v int32)`

SetMaxSnippetSize sets MaxSnippetSize field to given value.

### HasMaxSnippetSize

`func (o *ResultsRequest) HasMaxSnippetSize() bool`

HasMaxSnippetSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


