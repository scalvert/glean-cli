# SearchRequestInputDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasCopyPaste** | Pointer to **bool** | Whether the associated query was at least partially copy-pasted.  If subsequent requests are issued after a copy-pasted query is constructed (e.g. with facet modifications), this bit should continue to be set for those requests. | [optional] 

## Methods

### NewSearchRequestInputDetails

`func NewSearchRequestInputDetails() *SearchRequestInputDetails`

NewSearchRequestInputDetails instantiates a new SearchRequestInputDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestInputDetailsWithDefaults

`func NewSearchRequestInputDetailsWithDefaults() *SearchRequestInputDetails`

NewSearchRequestInputDetailsWithDefaults instantiates a new SearchRequestInputDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasCopyPaste

`func (o *SearchRequestInputDetails) GetHasCopyPaste() bool`

GetHasCopyPaste returns the HasCopyPaste field if non-nil, zero value otherwise.

### GetHasCopyPasteOk

`func (o *SearchRequestInputDetails) GetHasCopyPasteOk() (*bool, bool)`

GetHasCopyPasteOk returns a tuple with the HasCopyPaste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCopyPaste

`func (o *SearchRequestInputDetails) SetHasCopyPaste(v bool)`

SetHasCopyPaste sets HasCopyPaste field to given value.

### HasHasCopyPaste

`func (o *SearchRequestInputDetails) HasHasCopyPaste() bool`

HasHasCopyPaste returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


