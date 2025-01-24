# AutocompleteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** |  | 
**Keywords** | Pointer to **[]string** | A list of all possible keywords for given result. | [optional] 
**ResultType** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** | Higher indicates a more confident match. | [optional] 
**OperatorMetadata** | Pointer to [**OperatorMetadata**](OperatorMetadata.md) |  | [optional] 
**Quicklink** | Pointer to [**Quicklink**](Quicklink.md) |  | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**StructuredResult** | Pointer to [**StructuredResult**](StructuredResult.md) |  | [optional] 
**TrackingToken** | Pointer to **string** | A token to be passed in /feedback events associated with this autocomplete result. | [optional] 
**Ranges** | Pointer to [**[]TextRange**](TextRange.md) | Subsections of the result string to which some special formatting should be applied (eg. bold) | [optional] 

## Methods

### NewAutocompleteResult

`func NewAutocompleteResult(result string, ) *AutocompleteResult`

NewAutocompleteResult instantiates a new AutocompleteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteResultWithDefaults

`func NewAutocompleteResultWithDefaults() *AutocompleteResult`

NewAutocompleteResultWithDefaults instantiates a new AutocompleteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AutocompleteResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AutocompleteResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AutocompleteResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetKeywords

`func (o *AutocompleteResult) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *AutocompleteResult) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *AutocompleteResult) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *AutocompleteResult) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetResultType

`func (o *AutocompleteResult) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *AutocompleteResult) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *AutocompleteResult) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *AutocompleteResult) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetScore

`func (o *AutocompleteResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AutocompleteResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AutocompleteResult) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *AutocompleteResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetOperatorMetadata

`func (o *AutocompleteResult) GetOperatorMetadata() OperatorMetadata`

GetOperatorMetadata returns the OperatorMetadata field if non-nil, zero value otherwise.

### GetOperatorMetadataOk

`func (o *AutocompleteResult) GetOperatorMetadataOk() (*OperatorMetadata, bool)`

GetOperatorMetadataOk returns a tuple with the OperatorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorMetadata

`func (o *AutocompleteResult) SetOperatorMetadata(v OperatorMetadata)`

SetOperatorMetadata sets OperatorMetadata field to given value.

### HasOperatorMetadata

`func (o *AutocompleteResult) HasOperatorMetadata() bool`

HasOperatorMetadata returns a boolean if a field has been set.

### GetQuicklink

`func (o *AutocompleteResult) GetQuicklink() Quicklink`

GetQuicklink returns the Quicklink field if non-nil, zero value otherwise.

### GetQuicklinkOk

`func (o *AutocompleteResult) GetQuicklinkOk() (*Quicklink, bool)`

GetQuicklinkOk returns a tuple with the Quicklink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklink

`func (o *AutocompleteResult) SetQuicklink(v Quicklink)`

SetQuicklink sets Quicklink field to given value.

### HasQuicklink

`func (o *AutocompleteResult) HasQuicklink() bool`

HasQuicklink returns a boolean if a field has been set.

### GetDocument

`func (o *AutocompleteResult) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *AutocompleteResult) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *AutocompleteResult) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *AutocompleteResult) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetUrl

`func (o *AutocompleteResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutocompleteResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutocompleteResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutocompleteResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStructuredResult

`func (o *AutocompleteResult) GetStructuredResult() StructuredResult`

GetStructuredResult returns the StructuredResult field if non-nil, zero value otherwise.

### GetStructuredResultOk

`func (o *AutocompleteResult) GetStructuredResultOk() (*StructuredResult, bool)`

GetStructuredResultOk returns a tuple with the StructuredResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredResult

`func (o *AutocompleteResult) SetStructuredResult(v StructuredResult)`

SetStructuredResult sets StructuredResult field to given value.

### HasStructuredResult

`func (o *AutocompleteResult) HasStructuredResult() bool`

HasStructuredResult returns a boolean if a field has been set.

### GetTrackingToken

`func (o *AutocompleteResult) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *AutocompleteResult) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *AutocompleteResult) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *AutocompleteResult) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetRanges

`func (o *AutocompleteResult) GetRanges() []TextRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AutocompleteResult) GetRangesOk() (*[]TextRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AutocompleteResult) SetRanges(v []TextRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *AutocompleteResult) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


