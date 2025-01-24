# ChatMetadataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | Pointer to [**ChatMetadata**](ChatMetadata.md) |  | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular Chat. To be used for &#x60;/feedback&#x60; reporting. | [optional] 

## Methods

### NewChatMetadataResult

`func NewChatMetadataResult() *ChatMetadataResult`

NewChatMetadataResult instantiates a new ChatMetadataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMetadataResultWithDefaults

`func NewChatMetadataResultWithDefaults() *ChatMetadataResult`

NewChatMetadataResultWithDefaults instantiates a new ChatMetadataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *ChatMetadataResult) GetChat() ChatMetadata`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ChatMetadataResult) GetChatOk() (*ChatMetadata, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ChatMetadataResult) SetChat(v ChatMetadata)`

SetChat sets Chat field to given value.

### HasChat

`func (o *ChatMetadataResult) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetTrackingToken

`func (o *ChatMetadataResult) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *ChatMetadataResult) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *ChatMetadataResult) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *ChatMetadataResult) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


