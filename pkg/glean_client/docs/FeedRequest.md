# FeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshType** | **string** | Type of refresh requested. Intended for logging and future optimizations. | 
**Categories** | Pointer to **[]string** | Categories of content requested. An allowlist gives flexibility to request content separately or together. | [optional] 
**RequestOptions** | Pointer to [**FeedRequestOptions**](FeedRequestOptions.md) |  | [optional] 
**ClientData** | Pointer to [**ClientData**](ClientData.md) |  | [optional] 
**TimeoutMillis** | Pointer to **int32** | Timeout in milliseconds for the request. A &#x60;408&#x60; error will be returned if handling the request takes longer. | [optional] 
**SessionInfo** | Pointer to [**SessionInfo**](SessionInfo.md) |  | [optional] 

## Methods

### NewFeedRequest

`func NewFeedRequest(refreshType string, ) *FeedRequest`

NewFeedRequest instantiates a new FeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedRequestWithDefaults

`func NewFeedRequestWithDefaults() *FeedRequest`

NewFeedRequestWithDefaults instantiates a new FeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshType

`func (o *FeedRequest) GetRefreshType() string`

GetRefreshType returns the RefreshType field if non-nil, zero value otherwise.

### GetRefreshTypeOk

`func (o *FeedRequest) GetRefreshTypeOk() (*string, bool)`

GetRefreshTypeOk returns a tuple with the RefreshType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshType

`func (o *FeedRequest) SetRefreshType(v string)`

SetRefreshType sets RefreshType field to given value.


### GetCategories

`func (o *FeedRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *FeedRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *FeedRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *FeedRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetRequestOptions

`func (o *FeedRequest) GetRequestOptions() FeedRequestOptions`

GetRequestOptions returns the RequestOptions field if non-nil, zero value otherwise.

### GetRequestOptionsOk

`func (o *FeedRequest) GetRequestOptionsOk() (*FeedRequestOptions, bool)`

GetRequestOptionsOk returns a tuple with the RequestOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestOptions

`func (o *FeedRequest) SetRequestOptions(v FeedRequestOptions)`

SetRequestOptions sets RequestOptions field to given value.

### HasRequestOptions

`func (o *FeedRequest) HasRequestOptions() bool`

HasRequestOptions returns a boolean if a field has been set.

### GetClientData

`func (o *FeedRequest) GetClientData() ClientData`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *FeedRequest) GetClientDataOk() (*ClientData, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *FeedRequest) SetClientData(v ClientData)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *FeedRequest) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *FeedRequest) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *FeedRequest) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *FeedRequest) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *FeedRequest) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetSessionInfo

`func (o *FeedRequest) GetSessionInfo() SessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *FeedRequest) GetSessionInfoOk() (*SessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *FeedRequest) SetSessionInfo(v SessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *FeedRequest) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


