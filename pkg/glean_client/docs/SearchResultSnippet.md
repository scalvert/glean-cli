# SearchResultSnippet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snippet** | **string** | A matching snippet from the document. Query term matches are marked by the unicode characters uE006 and uE007. | 
**MimeType** | Pointer to **string** | The mime type of the snippets, currently either text/plain or text/html. | [optional] 
**Text** | Pointer to **string** | A matching snippet from the document with no highlights. | [optional] 
**SnippetTextOrdering** | Pointer to **int32** | Used for sorting based off the snippet&#39;s location within all_snippetable_text | [optional] 
**Ranges** | Pointer to [**[]TextRange**](TextRange.md) | The bolded ranges within text. | [optional] 
**Url** | Pointer to **string** | A URL, generated based on availability, that links to the position of the snippet text or to the nearest header above the snippet text. | [optional] 

## Methods

### NewSearchResultSnippet

`func NewSearchResultSnippet(snippet string, ) *SearchResultSnippet`

NewSearchResultSnippet instantiates a new SearchResultSnippet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultSnippetWithDefaults

`func NewSearchResultSnippetWithDefaults() *SearchResultSnippet`

NewSearchResultSnippetWithDefaults instantiates a new SearchResultSnippet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnippet

`func (o *SearchResultSnippet) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *SearchResultSnippet) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *SearchResultSnippet) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.


### GetMimeType

`func (o *SearchResultSnippet) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *SearchResultSnippet) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *SearchResultSnippet) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *SearchResultSnippet) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetText

`func (o *SearchResultSnippet) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SearchResultSnippet) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SearchResultSnippet) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SearchResultSnippet) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSnippetTextOrdering

`func (o *SearchResultSnippet) GetSnippetTextOrdering() int32`

GetSnippetTextOrdering returns the SnippetTextOrdering field if non-nil, zero value otherwise.

### GetSnippetTextOrderingOk

`func (o *SearchResultSnippet) GetSnippetTextOrderingOk() (*int32, bool)`

GetSnippetTextOrderingOk returns a tuple with the SnippetTextOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetTextOrdering

`func (o *SearchResultSnippet) SetSnippetTextOrdering(v int32)`

SetSnippetTextOrdering sets SnippetTextOrdering field to given value.

### HasSnippetTextOrdering

`func (o *SearchResultSnippet) HasSnippetTextOrdering() bool`

HasSnippetTextOrdering returns a boolean if a field has been set.

### GetRanges

`func (o *SearchResultSnippet) GetRanges() []TextRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *SearchResultSnippet) GetRangesOk() (*[]TextRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *SearchResultSnippet) SetRanges(v []TextRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *SearchResultSnippet) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetUrl

`func (o *SearchResultSnippet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SearchResultSnippet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SearchResultSnippet) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SearchResultSnippet) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


