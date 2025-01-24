# DisplayableList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | The type of data that backs this displayable list | [optional] 
**Id** | Pointer to **int32** | Unique identifier of this list. Unique amongst only DisplayableLists, not unique amongst other types of UGC. | [optional] 
**SourceId** | Pointer to **string** | Unstructured identifier for the source to render (ID, URL, query). | [optional] 
**Config** | Pointer to [**DisplayableListConfig**](DisplayableListConfig.md) |  | [optional] 

## Methods

### NewDisplayableList

`func NewDisplayableList() *DisplayableList`

NewDisplayableList instantiates a new DisplayableList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayableListWithDefaults

`func NewDisplayableListWithDefaults() *DisplayableList`

NewDisplayableListWithDefaults instantiates a new DisplayableList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *DisplayableList) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DisplayableList) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DisplayableList) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DisplayableList) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetId

`func (o *DisplayableList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisplayableList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisplayableList) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DisplayableList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceId

`func (o *DisplayableList) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DisplayableList) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DisplayableList) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *DisplayableList) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetConfig

`func (o *DisplayableList) GetConfig() DisplayableListConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DisplayableList) GetConfigOk() (*DisplayableListConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DisplayableList) SetConfig(v DisplayableListConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DisplayableList) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


