# ChatResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chat** | Pointer to [**Chat**](Chat.md) |  | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular Chat. To be used for &#x60;/feedback&#x60; reporting. | [optional] 

## Methods

### NewChatResult

`func NewChatResult() *ChatResult`

NewChatResult instantiates a new ChatResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatResultWithDefaults

`func NewChatResultWithDefaults() *ChatResult`

NewChatResultWithDefaults instantiates a new ChatResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChat

`func (o *ChatResult) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *ChatResult) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *ChatResult) SetChat(v Chat)`

SetChat sets Chat field to given value.

### HasChat

`func (o *ChatResult) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetTrackingToken

`func (o *ChatResult) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *ChatResult) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *ChatResult) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *ChatResult) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


