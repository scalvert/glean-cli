# PreviewShortcutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shortcut** | Pointer to [**Shortcut**](Shortcut.md) |  | [optional] 
**ExistingUrlShortcuts** | Pointer to [**[]Shortcut**](Shortcut.md) | Exising shortcuts that have a similar destination URL. | [optional] 
**Error** | Pointer to [**ShortcutError**](ShortcutError.md) |  | [optional] 

## Methods

### NewPreviewShortcutResponse

`func NewPreviewShortcutResponse() *PreviewShortcutResponse`

NewPreviewShortcutResponse instantiates a new PreviewShortcutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewShortcutResponseWithDefaults

`func NewPreviewShortcutResponseWithDefaults() *PreviewShortcutResponse`

NewPreviewShortcutResponseWithDefaults instantiates a new PreviewShortcutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortcut

`func (o *PreviewShortcutResponse) GetShortcut() Shortcut`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *PreviewShortcutResponse) GetShortcutOk() (*Shortcut, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *PreviewShortcutResponse) SetShortcut(v Shortcut)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *PreviewShortcutResponse) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetExistingUrlShortcuts

`func (o *PreviewShortcutResponse) GetExistingUrlShortcuts() []Shortcut`

GetExistingUrlShortcuts returns the ExistingUrlShortcuts field if non-nil, zero value otherwise.

### GetExistingUrlShortcutsOk

`func (o *PreviewShortcutResponse) GetExistingUrlShortcutsOk() (*[]Shortcut, bool)`

GetExistingUrlShortcutsOk returns a tuple with the ExistingUrlShortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingUrlShortcuts

`func (o *PreviewShortcutResponse) SetExistingUrlShortcuts(v []Shortcut)`

SetExistingUrlShortcuts sets ExistingUrlShortcuts field to given value.

### HasExistingUrlShortcuts

`func (o *PreviewShortcutResponse) HasExistingUrlShortcuts() bool`

HasExistingUrlShortcuts returns a boolean if a field has been set.

### GetError

`func (o *PreviewShortcutResponse) GetError() ShortcutError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PreviewShortcutResponse) GetErrorOk() (*ShortcutError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PreviewShortcutResponse) SetError(v ShortcutError)`

SetError sets Error field to given value.

### HasError

`func (o *PreviewShortcutResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


