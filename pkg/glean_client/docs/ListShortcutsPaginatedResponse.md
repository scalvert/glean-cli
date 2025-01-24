# ListShortcutsPaginatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shortcuts** | [**[]Shortcut**](Shortcut.md) | List of all shortcuts accessible to the user | 
**FacetResults** | Pointer to [**[]FacetResult**](FacetResult.md) |  | [optional] 
**Meta** | [**ShortcutsPaginationMetadata**](ShortcutsPaginationMetadata.md) |  | 

## Methods

### NewListShortcutsPaginatedResponse

`func NewListShortcutsPaginatedResponse(shortcuts []Shortcut, meta ShortcutsPaginationMetadata, ) *ListShortcutsPaginatedResponse`

NewListShortcutsPaginatedResponse instantiates a new ListShortcutsPaginatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListShortcutsPaginatedResponseWithDefaults

`func NewListShortcutsPaginatedResponseWithDefaults() *ListShortcutsPaginatedResponse`

NewListShortcutsPaginatedResponseWithDefaults instantiates a new ListShortcutsPaginatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortcuts

`func (o *ListShortcutsPaginatedResponse) GetShortcuts() []Shortcut`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *ListShortcutsPaginatedResponse) GetShortcutsOk() (*[]Shortcut, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *ListShortcutsPaginatedResponse) SetShortcuts(v []Shortcut)`

SetShortcuts sets Shortcuts field to given value.


### GetFacetResults

`func (o *ListShortcutsPaginatedResponse) GetFacetResults() []FacetResult`

GetFacetResults returns the FacetResults field if non-nil, zero value otherwise.

### GetFacetResultsOk

`func (o *ListShortcutsPaginatedResponse) GetFacetResultsOk() (*[]FacetResult, bool)`

GetFacetResultsOk returns a tuple with the FacetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetResults

`func (o *ListShortcutsPaginatedResponse) SetFacetResults(v []FacetResult)`

SetFacetResults sets FacetResults field to given value.

### HasFacetResults

`func (o *ListShortcutsPaginatedResponse) HasFacetResults() bool`

HasFacetResults returns a boolean if a field has been set.

### GetMeta

`func (o *ListShortcutsPaginatedResponse) GetMeta() ShortcutsPaginationMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListShortcutsPaginatedResponse) GetMetaOk() (*ShortcutsPaginationMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListShortcutsPaginatedResponse) SetMeta(v ShortcutsPaginationMetadata)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


