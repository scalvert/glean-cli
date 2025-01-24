# GetShortcutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The opaque id of the user generated content. | [optional] 
**Alias** | **string** | The alias for the shortcut, including any arguments for variable shortcuts. | 

## Methods

### NewGetShortcutRequest

`func NewGetShortcutRequest(alias string, ) *GetShortcutRequest`

NewGetShortcutRequest instantiates a new GetShortcutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShortcutRequestWithDefaults

`func NewGetShortcutRequestWithDefaults() *GetShortcutRequest`

NewGetShortcutRequestWithDefaults instantiates a new GetShortcutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetShortcutRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetShortcutRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetShortcutRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetShortcutRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *GetShortcutRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetShortcutRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetShortcutRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


