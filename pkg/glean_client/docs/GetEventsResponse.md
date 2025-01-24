# GetEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]FeedResult**](FeedResult.md) | The requested events. Uses a FeedResult for now as a Calendar Event doesn&#39;t have all event info such as attachments. | [optional] 

## Methods

### NewGetEventsResponse

`func NewGetEventsResponse() *GetEventsResponse`

NewGetEventsResponse instantiates a new GetEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventsResponseWithDefaults

`func NewGetEventsResponseWithDefaults() *GetEventsResponse`

NewGetEventsResponseWithDefaults instantiates a new GetEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *GetEventsResponse) GetEvents() []FeedResult`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetEventsResponse) GetEventsOk() (*[]FeedResult, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetEventsResponse) SetEvents(v []FeedResult)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


