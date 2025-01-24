# ChatMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The opaque id of the Chat. | [optional] 
**CreateTime** | Pointer to **int32** | Server Unix timestamp of the creation time (in seconds since epoch UTC). | [optional] 
**CreatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**UpdateTime** | Pointer to **int32** | Server Unix timestamp of the update time (in seconds since epoch UTC). | [optional] 
**Name** | Pointer to **string** | The name of the Chat. | [optional] 

## Methods

### NewChatMetadata

`func NewChatMetadata() *ChatMetadata`

NewChatMetadata instantiates a new ChatMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMetadataWithDefaults

`func NewChatMetadataWithDefaults() *ChatMetadata`

NewChatMetadataWithDefaults instantiates a new ChatMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTime

`func (o *ChatMetadata) GetCreateTime() int32`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ChatMetadata) GetCreateTimeOk() (*int32, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ChatMetadata) SetCreateTime(v int32)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ChatMetadata) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ChatMetadata) GetCreatedBy() Person`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChatMetadata) GetCreatedByOk() (*Person, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChatMetadata) SetCreatedBy(v Person)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ChatMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdateTime

`func (o *ChatMetadata) GetUpdateTime() int32`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ChatMetadata) GetUpdateTimeOk() (*int32, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ChatMetadata) SetUpdateTime(v int32)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *ChatMetadata) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetName

`func (o *ChatMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChatMetadata) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


