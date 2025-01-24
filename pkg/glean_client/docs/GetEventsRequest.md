# GetEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | The ids of the calendar events to be retrieved. | 
**AuthTokens** | Pointer to [**[]AuthToken**](AuthToken.md) | Auth tokens if client-side authentication. | [optional] 
**Datasource** | Pointer to **string** | The app or other repository type from which the event was extracted | [optional] 
**Annotate** | Pointer to **bool** | Whether relevant content and documents, via GeneratedAttachments, should be attached to the events. | [optional] 

## Methods

### NewGetEventsRequest

`func NewGetEventsRequest(ids []string, ) *GetEventsRequest`

NewGetEventsRequest instantiates a new GetEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventsRequestWithDefaults

`func NewGetEventsRequestWithDefaults() *GetEventsRequest`

NewGetEventsRequestWithDefaults instantiates a new GetEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *GetEventsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *GetEventsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *GetEventsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetAuthTokens

`func (o *GetEventsRequest) GetAuthTokens() []AuthToken`

GetAuthTokens returns the AuthTokens field if non-nil, zero value otherwise.

### GetAuthTokensOk

`func (o *GetEventsRequest) GetAuthTokensOk() (*[]AuthToken, bool)`

GetAuthTokensOk returns a tuple with the AuthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokens

`func (o *GetEventsRequest) SetAuthTokens(v []AuthToken)`

SetAuthTokens sets AuthTokens field to given value.

### HasAuthTokens

`func (o *GetEventsRequest) HasAuthTokens() bool`

HasAuthTokens returns a boolean if a field has been set.

### GetDatasource

`func (o *GetEventsRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *GetEventsRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *GetEventsRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *GetEventsRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetAnnotate

`func (o *GetEventsRequest) GetAnnotate() bool`

GetAnnotate returns the Annotate field if non-nil, zero value otherwise.

### GetAnnotateOk

`func (o *GetEventsRequest) GetAnnotateOk() (*bool, bool)`

GetAnnotateOk returns a tuple with the Annotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotate

`func (o *GetEventsRequest) SetAnnotate(v bool)`

SetAnnotate sets Annotate field to given value.

### HasAnnotate

`func (o *GetEventsRequest) HasAnnotate() bool`

HasAnnotate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


