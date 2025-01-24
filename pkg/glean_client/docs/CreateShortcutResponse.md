# CreateShortcutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shortcut** | Pointer to [**Shortcut**](Shortcut.md) |  | [optional] 
**Error** | Pointer to [**ShortcutError**](ShortcutError.md) |  | [optional] 

## Methods

### NewCreateShortcutResponse

`func NewCreateShortcutResponse() *CreateShortcutResponse`

NewCreateShortcutResponse instantiates a new CreateShortcutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShortcutResponseWithDefaults

`func NewCreateShortcutResponseWithDefaults() *CreateShortcutResponse`

NewCreateShortcutResponseWithDefaults instantiates a new CreateShortcutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortcut

`func (o *CreateShortcutResponse) GetShortcut() Shortcut`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *CreateShortcutResponse) GetShortcutOk() (*Shortcut, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *CreateShortcutResponse) SetShortcut(v Shortcut)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *CreateShortcutResponse) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetError

`func (o *CreateShortcutResponse) GetError() ShortcutError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateShortcutResponse) GetErrorOk() (*ShortcutError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateShortcutResponse) SetError(v ShortcutError)`

SetError sets Error field to given value.

### HasError

`func (o *CreateShortcutResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


