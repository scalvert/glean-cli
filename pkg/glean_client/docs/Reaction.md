# Reaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** | The count of the reaction type on the document. | [optional] 
**Reactors** | Pointer to [**[]Person**](Person.md) |  | [optional] 
**ReactedByViewer** | Pointer to **bool** | Whether the user in context reacted with this type to the document. | [optional] 

## Methods

### NewReaction

`func NewReaction() *Reaction`

NewReaction instantiates a new Reaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionWithDefaults

`func NewReactionWithDefaults() *Reaction`

NewReactionWithDefaults instantiates a new Reaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Reaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Reaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Reaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Reaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCount

`func (o *Reaction) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Reaction) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Reaction) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Reaction) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetReactors

`func (o *Reaction) GetReactors() []Person`

GetReactors returns the Reactors field if non-nil, zero value otherwise.

### GetReactorsOk

`func (o *Reaction) GetReactorsOk() (*[]Person, bool)`

GetReactorsOk returns a tuple with the Reactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactors

`func (o *Reaction) SetReactors(v []Person)`

SetReactors sets Reactors field to given value.

### HasReactors

`func (o *Reaction) HasReactors() bool`

HasReactors returns a boolean if a field has been set.

### GetReactedByViewer

`func (o *Reaction) GetReactedByViewer() bool`

GetReactedByViewer returns the ReactedByViewer field if non-nil, zero value otherwise.

### GetReactedByViewerOk

`func (o *Reaction) GetReactedByViewerOk() (*bool, bool)`

GetReactedByViewerOk returns a tuple with the ReactedByViewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactedByViewer

`func (o *Reaction) SetReactedByViewer(v bool)`

SetReactedByViewer sets ReactedByViewer field to given value.

### HasReactedByViewer

`func (o *Reaction) HasReactedByViewer() bool`

HasReactedByViewer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


