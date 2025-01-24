# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The ISO 8601 timestamp associated with the client request. | [optional] 
**TrackingToken** | Pointer to **string** | A previously received trackingToken for a search associated with the same query. Useful for more requests and requests for other tabs. | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 
**SourceDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**PageSize** | Pointer to **int32** | Hint to the server about how many results to send back. Server may return less or more. Structured results and clustered results don&#39;t count towards pageSize. | [optional] 
**MaxSnippetSize** | Pointer to **int32** | Hint to the server about how many characters long a snippet may be. Server may return less or more. | [optional] 
**Query** | **string** | The search terms. | 
**Cursor** | Pointer to **string** | Pagination cursor. A previously received opaque token representing the position in the overall results at which to start. | [optional] 
**ResultTabIds** | Pointer to **[]string** | The unique IDs of the result tabs for which to fetch results. This will have precedence over datasource filters if both are specified and in conflict. | [optional] 
**InputDetails** | Pointer to [**SearchRequestInputDetails**](SearchRequestInputDetails.md) |  | [optional] 
**RequestOptions** | Pointer to [**SearchRequestOptions**](SearchRequestOptions.md) |  | [optional] 
**TimeoutMillis** | Pointer to **int32** | Timeout in milliseconds for the request. A &#x60;408&#x60; error will be returned if handling the request takes longer. | [optional] 
**People** | Pointer to [**[]Person**](Person.md) | People associated with the search request. Hints to the server to fetch additional information for these people. Note that in this request, an email may be used as a person&#39;s obfuscatedId value. | [optional] 
**DisableSpellcheck** | Pointer to **bool** | Whether or not to disable spellcheck. | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest(query string, ) *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SearchRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SearchRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SearchRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SearchRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTrackingToken

`func (o *SearchRequest) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *SearchRequest) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *SearchRequest) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *SearchRequest) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetSessionInfo

`func (o *SearchRequest) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *SearchRequest) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *SearchRequest) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *SearchRequest) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetSourceDocument

`func (o *SearchRequest) GetSourceDocument() Document`

GetSourceDocument returns the SourceDocument field if non-nil, zero value otherwise.

### GetSourceDocumentOk

`func (o *SearchRequest) GetSourceDocumentOk() (*Document, bool)`

GetSourceDocumentOk returns a tuple with the SourceDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocument

`func (o *SearchRequest) SetSourceDocument(v Document)`

SetSourceDocument sets SourceDocument field to given value.

### HasSourceDocument

`func (o *SearchRequest) HasSourceDocument() bool`

HasSourceDocument returns a boolean if a field has been set.

### GetPageSize

`func (o *SearchRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SearchRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SearchRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SearchRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetMaxSnippetSize

`func (o *SearchRequest) GetMaxSnippetSize() int32`

GetMaxSnippetSize returns the MaxSnippetSize field if non-nil, zero value otherwise.

### GetMaxSnippetSizeOk

`func (o *SearchRequest) GetMaxSnippetSizeOk() (*int32, bool)`

GetMaxSnippetSizeOk returns a tuple with the MaxSnippetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnippetSize

`func (o *SearchRequest) SetMaxSnippetSize(v int32)`

SetMaxSnippetSize sets MaxSnippetSize field to given value.

### HasMaxSnippetSize

`func (o *SearchRequest) HasMaxSnippetSize() bool`

HasMaxSnippetSize returns a boolean if a field has been set.

### GetQuery

`func (o *SearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetCursor

`func (o *SearchRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *SearchRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *SearchRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *SearchRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetResultTabIds

`func (o *SearchRequest) GetResultTabIds() []string`

GetResultTabIds returns the ResultTabIds field if non-nil, zero value otherwise.

### GetResultTabIdsOk

`func (o *SearchRequest) GetResultTabIdsOk() (*[]string, bool)`

GetResultTabIdsOk returns a tuple with the ResultTabIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTabIds

`func (o *SearchRequest) SetResultTabIds(v []string)`

SetResultTabIds sets ResultTabIds field to given value.

### HasResultTabIds

`func (o *SearchRequest) HasResultTabIds() bool`

HasResultTabIds returns a boolean if a field has been set.

### GetInputDetails

`func (o *SearchRequest) GetInputDetails() SearchRequestInputDetails`

GetInputDetails returns the InputDetails field if non-nil, zero value otherwise.

### GetInputDetailsOk

`func (o *SearchRequest) GetInputDetailsOk() (*SearchRequestInputDetails, bool)`

GetInputDetailsOk returns a tuple with the InputDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDetails

`func (o *SearchRequest) SetInputDetails(v SearchRequestInputDetails)`

SetInputDetails sets InputDetails field to given value.

### HasInputDetails

`func (o *SearchRequest) HasInputDetails() bool`

HasInputDetails returns a boolean if a field has been set.

### GetRequestOptions

`func (o *SearchRequest) GetRequestOptions() SearchRequestOptions`

GetRequestOptions returns the RequestOptions field if non-nil, zero value otherwise.

### GetRequestOptionsOk

`func (o *SearchRequest) GetRequestOptionsOk() (*SearchRequestOptions, bool)`

GetRequestOptionsOk returns a tuple with the RequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOptions

`func (o *SearchRequest) SetRequestOptions(v SearchRequestOptions)`

SetRequestOptions sets RequestOptions field to given value.

### HasRequestOptions

`func (o *SearchRequest) HasRequestOptions() bool`

HasRequestOptions returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *SearchRequest) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *SearchRequest) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *SearchRequest) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *SearchRequest) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetPeople

`func (o *SearchRequest) GetPeople() []Person`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *SearchRequest) GetPeopleOk() (*[]Person, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *SearchRequest) SetPeople(v []Person)`

SetPeople sets People field to given value.

### HasPeople

`func (o *SearchRequest) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetDisableSpellcheck

`func (o *SearchRequest) GetDisableSpellcheck() bool`

GetDisableSpellcheck returns the DisableSpellcheck field if non-nil, zero value otherwise.

### GetDisableSpellcheckOk

`func (o *SearchRequest) GetDisableSpellcheckOk() (*bool, bool)`

GetDisableSpellcheckOk returns a tuple with the DisableSpellcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSpellcheck

`func (o *SearchRequest) SetDisableSpellcheck(v bool)`

SetDisableSpellcheck sets DisableSpellcheck field to given value.

### HasDisableSpellcheck

`func (o *SearchRequest) HasDisableSpellcheck() bool`

HasDisableSpellcheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


