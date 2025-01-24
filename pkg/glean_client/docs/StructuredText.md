# StructuredText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**StructuredList** | Pointer to [**[]StructuredTextItem**](StructuredTextItem.md) | An array of objects each of which contains either a string or a link which optionally corresponds to a document. | [optional] 

## Methods

### NewStructuredText

`func NewStructuredText(text string, ) *StructuredText`

NewStructuredText instantiates a new StructuredText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuredTextWithDefaults

`func NewStructuredTextWithDefaults() *StructuredText`

NewStructuredTextWithDefaults instantiates a new StructuredText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *StructuredText) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StructuredText) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StructuredText) SetText(v string)`

SetText sets Text field to given value.


### GetStructuredList

`func (o *StructuredText) GetStructuredList() []StructuredTextItem`

GetStructuredList returns the StructuredList field if non-nil, zero value otherwise.

### GetStructuredListOk

`func (o *StructuredText) GetStructuredListOk() (*[]StructuredTextItem, bool)`

GetStructuredListOk returns a tuple with the StructuredList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredList

`func (o *StructuredText) SetStructuredList(v []StructuredTextItem)`

SetStructuredList sets StructuredList field to given value.

### HasStructuredList

`func (o *StructuredText) HasStructuredList() bool`

HasStructuredList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


