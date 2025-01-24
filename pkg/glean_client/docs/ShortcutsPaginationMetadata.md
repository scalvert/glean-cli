# ShortcutsPaginationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** | Cursor indicates the start of the next page of results | [optional] 
**HasNextPage** | Pointer to **bool** |  | [optional] 
**TotalItemCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewShortcutsPaginationMetadata

`func NewShortcutsPaginationMetadata() *ShortcutsPaginationMetadata`

NewShortcutsPaginationMetadata instantiates a new ShortcutsPaginationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortcutsPaginationMetadataWithDefaults

`func NewShortcutsPaginationMetadataWithDefaults() *ShortcutsPaginationMetadata`

NewShortcutsPaginationMetadataWithDefaults instantiates a new ShortcutsPaginationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ShortcutsPaginationMetadata) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ShortcutsPaginationMetadata) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ShortcutsPaginationMetadata) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ShortcutsPaginationMetadata) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetHasNextPage

`func (o *ShortcutsPaginationMetadata) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *ShortcutsPaginationMetadata) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *ShortcutsPaginationMetadata) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.

### HasHasNextPage

`func (o *ShortcutsPaginationMetadata) HasHasNextPage() bool`

HasHasNextPage returns a boolean if a field has been set.

### GetTotalItemCount

`func (o *ShortcutsPaginationMetadata) GetTotalItemCount() int32`

GetTotalItemCount returns the TotalItemCount field if non-nil, zero value otherwise.

### GetTotalItemCountOk

`func (o *ShortcutsPaginationMetadata) GetTotalItemCountOk() (*int32, bool)`

GetTotalItemCountOk returns a tuple with the TotalItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItemCount

`func (o *ShortcutsPaginationMetadata) SetTotalItemCount(v int32)`

SetTotalItemCount sets TotalItemCount field to given value.

### HasTotalItemCount

`func (o *ShortcutsPaginationMetadata) HasTotalItemCount() bool`

HasTotalItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


