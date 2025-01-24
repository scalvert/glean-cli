# SummarizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The ISO 8601 timestamp associated with the client request. | [optional] 
**Query** | Pointer to **string** | Optional query that the summary should be about | [optional] 
**PreferredSummaryLength** | Pointer to **int32** | Optional length of summary output. If not given, defaults to 500 chars. | [optional] 
**DocumentSpecs** | [**[]DocumentSpec**](DocumentSpec.md) | Specifications of documents to summarize | 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular result. To be used for /feedback reporting. | [optional] 

## Methods

### NewSummarizeRequest

`func NewSummarizeRequest(documentSpecs []DocumentSpec, ) *SummarizeRequest`

NewSummarizeRequest instantiates a new SummarizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummarizeRequestWithDefaults

`func NewSummarizeRequestWithDefaults() *SummarizeRequest`

NewSummarizeRequestWithDefaults instantiates a new SummarizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SummarizeRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SummarizeRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SummarizeRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SummarizeRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetQuery

`func (o *SummarizeRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SummarizeRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SummarizeRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SummarizeRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetPreferredSummaryLength

`func (o *SummarizeRequest) GetPreferredSummaryLength() int32`

GetPreferredSummaryLength returns the PreferredSummaryLength field if non-nil, zero value otherwise.

### GetPreferredSummaryLengthOk

`func (o *SummarizeRequest) GetPreferredSummaryLengthOk() (*int32, bool)`

GetPreferredSummaryLengthOk returns a tuple with the PreferredSummaryLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSummaryLength

`func (o *SummarizeRequest) SetPreferredSummaryLength(v int32)`

SetPreferredSummaryLength sets PreferredSummaryLength field to given value.

### HasPreferredSummaryLength

`func (o *SummarizeRequest) HasPreferredSummaryLength() bool`

HasPreferredSummaryLength returns a boolean if a field has been set.

### GetDocumentSpecs

`func (o *SummarizeRequest) GetDocumentSpecs() []DocumentSpec`

GetDocumentSpecs returns the DocumentSpecs field if non-nil, zero value otherwise.

### GetDocumentSpecsOk

`func (o *SummarizeRequest) GetDocumentSpecsOk() (*[]DocumentSpec, bool)`

GetDocumentSpecsOk returns a tuple with the DocumentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSpecs

`func (o *SummarizeRequest) SetDocumentSpecs(v []DocumentSpec)`

SetDocumentSpecs sets DocumentSpecs field to given value.


### GetTrackingToken

`func (o *SummarizeRequest) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *SummarizeRequest) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *SummarizeRequest) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *SummarizeRequest) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


