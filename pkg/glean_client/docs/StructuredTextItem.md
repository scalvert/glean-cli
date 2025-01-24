# StructuredTextItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** |  | [optional] 
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 

## Methods

### NewStructuredTextItem

`func NewStructuredTextItem() *StructuredTextItem`

NewStructuredTextItem instantiates a new StructuredTextItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuredTextItemWithDefaults

`func NewStructuredTextItemWithDefaults() *StructuredTextItem`

NewStructuredTextItemWithDefaults instantiates a new StructuredTextItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *StructuredTextItem) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *StructuredTextItem) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *StructuredTextItem) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *StructuredTextItem) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetDocument

`func (o *StructuredTextItem) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *StructuredTextItem) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *StructuredTextItem) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *StructuredTextItem) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetText

`func (o *StructuredTextItem) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StructuredTextItem) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StructuredTextItem) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StructuredTextItem) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


