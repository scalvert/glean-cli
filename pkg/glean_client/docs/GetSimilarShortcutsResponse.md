# GetSimilarShortcutsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shortcuts** | [**[]Shortcut**](Shortcut.md) | Shortcuts with similar aliases to the given. Includes shortcut with the exact alias if it exists. | 

## Methods

### NewGetSimilarShortcutsResponse

`func NewGetSimilarShortcutsResponse(shortcuts []Shortcut, ) *GetSimilarShortcutsResponse`

NewGetSimilarShortcutsResponse instantiates a new GetSimilarShortcutsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSimilarShortcutsResponseWithDefaults

`func NewGetSimilarShortcutsResponseWithDefaults() *GetSimilarShortcutsResponse`

NewGetSimilarShortcutsResponseWithDefaults instantiates a new GetSimilarShortcutsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortcuts

`func (o *GetSimilarShortcutsResponse) GetShortcuts() []Shortcut`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *GetSimilarShortcutsResponse) GetShortcutsOk() (*[]Shortcut, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *GetSimilarShortcutsResponse) SetShortcuts(v []Shortcut)`

SetShortcuts sets Shortcuts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


