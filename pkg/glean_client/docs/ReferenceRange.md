# ReferenceRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TextRange** | [**TextRange**](TextRange.md) |  | 
**Snippets** | Pointer to [**[]SearchResultSnippet**](SearchResultSnippet.md) |  | [optional] 

## Methods

### NewReferenceRange

`func NewReferenceRange(textRange TextRange, ) *ReferenceRange`

NewReferenceRange instantiates a new ReferenceRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceRangeWithDefaults

`func NewReferenceRangeWithDefaults() *ReferenceRange`

NewReferenceRangeWithDefaults instantiates a new ReferenceRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTextRange

`func (o *ReferenceRange) GetTextRange() TextRange`

GetTextRange returns the TextRange field if non-nil, zero value otherwise.

### GetTextRangeOk

`func (o *ReferenceRange) GetTextRangeOk() (*TextRange, bool)`

GetTextRangeOk returns a tuple with the TextRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextRange

`func (o *ReferenceRange) SetTextRange(v TextRange)`

SetTextRange sets TextRange field to given value.


### GetSnippets

`func (o *ReferenceRange) GetSnippets() []SearchResultSnippet`

GetSnippets returns the Snippets field if non-nil, zero value otherwise.

### GetSnippetsOk

`func (o *ReferenceRange) GetSnippetsOk() (*[]SearchResultSnippet, bool)`

GetSnippetsOk returns a tuple with the Snippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippets

`func (o *ReferenceRange) SetSnippets(v []SearchResultSnippet)`

SetSnippets sets Snippets field to given value.

### HasSnippets

`func (o *ReferenceRange) HasSnippets() bool`

HasSnippets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


