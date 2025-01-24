# TextRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartIndex** | **int32** | The inclusive start index of the range. | 
**EndIndex** | Pointer to **int32** | The exclusive end index of the range. | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | The URL associated with the range, if applicable. For example, the linked URL for a LINK range. | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 

## Methods

### NewTextRange

`func NewTextRange(startIndex int32, ) *TextRange`

NewTextRange instantiates a new TextRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextRangeWithDefaults

`func NewTextRangeWithDefaults() *TextRange`

NewTextRangeWithDefaults instantiates a new TextRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartIndex

`func (o *TextRange) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *TextRange) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *TextRange) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


### GetEndIndex

`func (o *TextRange) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *TextRange) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *TextRange) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *TextRange) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetType

`func (o *TextRange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TextRange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TextRange) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TextRange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *TextRange) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TextRange) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TextRange) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TextRange) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDocument

`func (o *TextRange) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *TextRange) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *TextRange) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *TextRange) HasDocument() bool`

HasDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


