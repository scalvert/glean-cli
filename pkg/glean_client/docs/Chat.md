# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]ChatMessage**](ChatMessage.md) | The chat messages within a Chat. | [optional] 
**Id** | Pointer to **string** | The opaque id of the Chat. | [optional] 
**CreateTime** | Pointer to **int32** | Server Unix timestamp of the creation time (in seconds since epoch UTC). | [optional] 
**CreatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**UpdateTime** | Pointer to **int32** | Server Unix timestamp of the update time (in seconds since epoch UTC). | [optional] 
**Name** | Pointer to **string** | The name of the Chat. | [optional] 

## Methods

### NewChat

`func NewChat() *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *Chat) GetMessages() []ChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *Chat) GetMessagesOk() (*[]ChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *Chat) SetMessages(v []ChatMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *Chat) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetId

`func (o *Chat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Chat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Chat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Chat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTime

`func (o *Chat) GetCreateTime() int32`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Chat) GetCreateTimeOk() (*int32, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Chat) SetCreateTime(v int32)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Chat) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Chat) GetCreatedBy() Person`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Chat) GetCreatedByOk() (*Person, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Chat) SetCreatedBy(v Person)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Chat) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Chat) GetUpdateTime() int32`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Chat) GetUpdateTimeOk() (*int32, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Chat) SetUpdateTime(v int32)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Chat) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetName

`func (o *Chat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Chat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Chat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Chat) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


