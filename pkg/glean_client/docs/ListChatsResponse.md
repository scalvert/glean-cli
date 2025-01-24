# ListChatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatResults** | Pointer to [**[]ChatMetadataResult**](ChatMetadataResult.md) |  | [optional] 

## Methods

### NewListChatsResponse

`func NewListChatsResponse() *ListChatsResponse`

NewListChatsResponse instantiates a new ListChatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChatsResponseWithDefaults

`func NewListChatsResponseWithDefaults() *ListChatsResponse`

NewListChatsResponseWithDefaults instantiates a new ListChatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatResults

`func (o *ListChatsResponse) GetChatResults() []ChatMetadataResult`

GetChatResults returns the ChatResults field if non-nil, zero value otherwise.

### GetChatResultsOk

`func (o *ListChatsResponse) GetChatResultsOk() (*[]ChatMetadataResult, bool)`

GetChatResultsOk returns a tuple with the ChatResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatResults

`func (o *ListChatsResponse) SetChatResults(v []ChatMetadataResult)`

SetChatResults sets ChatResults field to given value.

### HasChatResults

`func (o *ListChatsResponse) HasChatResults() bool`

HasChatResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


