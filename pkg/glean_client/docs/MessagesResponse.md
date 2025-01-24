# MessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Whether there are more results for client to continue requesting. | 
**SearchResponse** | Pointer to [**SearchResponse**](SearchResponse.md) |  | [optional] 
**RootMessage** | Pointer to [**SearchResult**](SearchResult.md) |  | [optional] 

## Methods

### NewMessagesResponse

`func NewMessagesResponse(hasMore bool, ) *MessagesResponse`

NewMessagesResponse instantiates a new MessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesResponseWithDefaults

`func NewMessagesResponseWithDefaults() *MessagesResponse`

NewMessagesResponseWithDefaults instantiates a new MessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *MessagesResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *MessagesResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *MessagesResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetSearchResponse

`func (o *MessagesResponse) GetSearchResponse() SearchResponse`

GetSearchResponse returns the SearchResponse field if non-nil, zero value otherwise.

### GetSearchResponseOk

`func (o *MessagesResponse) GetSearchResponseOk() (*SearchResponse, bool)`

GetSearchResponseOk returns a tuple with the SearchResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResponse

`func (o *MessagesResponse) SetSearchResponse(v SearchResponse)`

SetSearchResponse sets SearchResponse field to given value.

### HasSearchResponse

`func (o *MessagesResponse) HasSearchResponse() bool`

HasSearchResponse returns a boolean if a field has been set.

### GetRootMessage

`func (o *MessagesResponse) GetRootMessage() SearchResult`

GetRootMessage returns the RootMessage field if non-nil, zero value otherwise.

### GetRootMessageOk

`func (o *MessagesResponse) GetRootMessageOk() (*SearchResult, bool)`

GetRootMessageOk returns a tuple with the RootMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootMessage

`func (o *MessagesResponse) SetRootMessage(v SearchResult)`

SetRootMessage sets RootMessage field to given value.

### HasRootMessage

`func (o *MessagesResponse) HasRootMessage() bool`

HasRootMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


