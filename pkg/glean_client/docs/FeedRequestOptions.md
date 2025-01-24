# FeedRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultSize** | **int32** | Number of results asked in response. If a result is a collection, counts as one. | 
**TimezoneOffset** | Pointer to **int32** | The offset of the client&#39;s timezone in minutes from UTC. e.g. PDT is -420 because it&#39;s 7 hours behind UTC. | [optional] 
**CategoryToResultSize** | Pointer to [**map[string]FeedRequestOptionsCategoryToResultSizeValue**](FeedRequestOptionsCategoryToResultSizeValue.md) | Mapping from category to number of results asked for the category. | [optional] 
**DatasourceFilter** | Pointer to **[]string** | Datasources for which content should be included. Empty is for all. | [optional] 
**AuthTokens** | Pointer to [**[]AuthToken**](AuthToken.md) | Auth tokens which may be used for federated retrieval of Feed entries. | [optional] 

## Methods

### NewFeedRequestOptions

`func NewFeedRequestOptions(resultSize int32, ) *FeedRequestOptions`

NewFeedRequestOptions instantiates a new FeedRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedRequestOptionsWithDefaults

`func NewFeedRequestOptionsWithDefaults() *FeedRequestOptions`

NewFeedRequestOptionsWithDefaults instantiates a new FeedRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultSize

`func (o *FeedRequestOptions) GetResultSize() int32`

GetResultSize returns the ResultSize field if non-nil, zero value otherwise.

### GetResultSizeOk

`func (o *FeedRequestOptions) GetResultSizeOk() (*int32, bool)`

GetResultSizeOk returns a tuple with the ResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSize

`func (o *FeedRequestOptions) SetResultSize(v int32)`

SetResultSize sets ResultSize field to given value.


### GetTimezoneOffset

`func (o *FeedRequestOptions) GetTimezoneOffset() int32`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *FeedRequestOptions) GetTimezoneOffsetOk() (*int32, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *FeedRequestOptions) SetTimezoneOffset(v int32)`

SetTimezoneOffset sets TimezoneOffset field to given value.

### HasTimezoneOffset

`func (o *FeedRequestOptions) HasTimezoneOffset() bool`

HasTimezoneOffset returns a boolean if a field has been set.

### GetCategoryToResultSize

`func (o *FeedRequestOptions) GetCategoryToResultSize() map[string]FeedRequestOptionsCategoryToResultSizeValue`

GetCategoryToResultSize returns the CategoryToResultSize field if non-nil, zero value otherwise.

### GetCategoryToResultSizeOk

`func (o *FeedRequestOptions) GetCategoryToResultSizeOk() (*map[string]FeedRequestOptionsCategoryToResultSizeValue, bool)`

GetCategoryToResultSizeOk returns a tuple with the CategoryToResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryToResultSize

`func (o *FeedRequestOptions) SetCategoryToResultSize(v map[string]FeedRequestOptionsCategoryToResultSizeValue)`

SetCategoryToResultSize sets CategoryToResultSize field to given value.

### HasCategoryToResultSize

`func (o *FeedRequestOptions) HasCategoryToResultSize() bool`

HasCategoryToResultSize returns a boolean if a field has been set.

### GetDatasourceFilter

`func (o *FeedRequestOptions) GetDatasourceFilter() []string`

GetDatasourceFilter returns the DatasourceFilter field if non-nil, zero value otherwise.

### GetDatasourceFilterOk

`func (o *FeedRequestOptions) GetDatasourceFilterOk() (*[]string, bool)`

GetDatasourceFilterOk returns a tuple with the DatasourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceFilter

`func (o *FeedRequestOptions) SetDatasourceFilter(v []string)`

SetDatasourceFilter sets DatasourceFilter field to given value.

### HasDatasourceFilter

`func (o *FeedRequestOptions) HasDatasourceFilter() bool`

HasDatasourceFilter returns a boolean if a field has been set.

### GetAuthTokens

`func (o *FeedRequestOptions) GetAuthTokens() []AuthToken`

GetAuthTokens returns the AuthTokens field if non-nil, zero value otherwise.

### GetAuthTokensOk

`func (o *FeedRequestOptions) GetAuthTokensOk() (*[]AuthToken, bool)`

GetAuthTokensOk returns a tuple with the AuthTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokens

`func (o *FeedRequestOptions) SetAuthTokens(v []AuthToken)`

SetAuthTokens sets AuthTokens field to given value.

### HasAuthTokens

`func (o *FeedRequestOptions) HasAuthTokens() bool`

HasAuthTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


