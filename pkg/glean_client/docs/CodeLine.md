# CodeLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineNumber** | Pointer to **int32** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Ranges** | Pointer to [**[]TextRange**](TextRange.md) | Index ranges depicting matched sections of the line | [optional] 

## Methods

### NewCodeLine

`func NewCodeLine() *CodeLine`

NewCodeLine instantiates a new CodeLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeLineWithDefaults

`func NewCodeLineWithDefaults() *CodeLine`

NewCodeLineWithDefaults instantiates a new CodeLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineNumber

`func (o *CodeLine) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *CodeLine) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *CodeLine) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *CodeLine) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetContent

`func (o *CodeLine) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CodeLine) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CodeLine) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CodeLine) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRanges

`func (o *CodeLine) GetRanges() []TextRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *CodeLine) GetRangesOk() (*[]TextRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *CodeLine) SetRanges(v []TextRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *CodeLine) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


