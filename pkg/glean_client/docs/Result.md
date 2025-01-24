# Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructuredResults** | Pointer to [**[]StructuredResult**](StructuredResult.md) | An array of entities in the work graph retrieved via a data request. | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular result in this particular query. To be used for /feedback reporting. | [optional] 

## Methods

### NewResult

`func NewResult() *Result`

NewResult instantiates a new Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultWithDefaults

`func NewResultWithDefaults() *Result`

NewResultWithDefaults instantiates a new Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructuredResults

`func (o *Result) GetStructuredResults() []StructuredResult`

GetStructuredResults returns the StructuredResults field if non-nil, zero value otherwise.

### GetStructuredResultsOk

`func (o *Result) GetStructuredResultsOk() (*[]StructuredResult, bool)`

GetStructuredResultsOk returns a tuple with the StructuredResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredResults

`func (o *Result) SetStructuredResults(v []StructuredResult)`

SetStructuredResults sets StructuredResults field to given value.

### HasStructuredResults

`func (o *Result) HasStructuredResults() bool`

HasStructuredResults returns a boolean if a field has been set.

### GetTrackingToken

`func (o *Result) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *Result) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *Result) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *Result) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


