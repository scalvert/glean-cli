# FeedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentIds** | Pointer to **[]int64** | List of experiment ids for the corresponding request. | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular feed response. | [optional] 
**ServerTimestamp** | **int32** | Server unix timestamp (in seconds since epoch UTC). | 
**Results** | Pointer to [**[]FeedResult**](FeedResult.md) |  | [optional] 
**BackendTimeMillis** | Pointer to **int64** | Time in milliseconds the backend took to respond to the request. | [optional] 
**DatasourceAffinity** | Pointer to **map[string]float32** | A mapping from datasources to affinity of the user to each with scores. | [optional] 
**CompanyResourcesCollectionId** | Pointer to **int32** | The unique ID of the collection for company resources. | [optional] 
**FacetResults** | Pointer to [**map[string][]FacetResult**](array.md) | Map from category to the list of facets that can be used to filter the entry&#39;s content. | [optional] 
**MentionsTimeWindowInHours** | Pointer to **int32** | The time window (in hours) used for generating user mentions. | [optional] 

## Methods

### NewFeedResponse

`func NewFeedResponse(serverTimestamp int32, ) *FeedResponse`

NewFeedResponse instantiates a new FeedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedResponseWithDefaults

`func NewFeedResponseWithDefaults() *FeedResponse`

NewFeedResponseWithDefaults instantiates a new FeedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentIds

`func (o *FeedResponse) GetExperimentIds() []int64`

GetExperimentIds returns the ExperimentIds field if non-nil, zero value otherwise.

### GetExperimentIdsOk

`func (o *FeedResponse) GetExperimentIdsOk() (*[]int64, bool)`

GetExperimentIdsOk returns a tuple with the ExperimentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentIds

`func (o *FeedResponse) SetExperimentIds(v []int64)`

SetExperimentIds sets ExperimentIds field to given value.

### HasExperimentIds

`func (o *FeedResponse) HasExperimentIds() bool`

HasExperimentIds returns a boolean if a field has been set.

### GetTrackingToken

`func (o *FeedResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *FeedResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *FeedResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *FeedResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetServerTimestamp

`func (o *FeedResponse) GetServerTimestamp() int32`

GetServerTimestamp returns the ServerTimestamp field if non-nil, zero value otherwise.

### GetServerTimestampOk

`func (o *FeedResponse) GetServerTimestampOk() (*int32, bool)`

GetServerTimestampOk returns a tuple with the ServerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimestamp

`func (o *FeedResponse) SetServerTimestamp(v int32)`

SetServerTimestamp sets ServerTimestamp field to given value.


### GetResults

`func (o *FeedResponse) GetResults() []FeedResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FeedResponse) GetResultsOk() (*[]FeedResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FeedResponse) SetResults(v []FeedResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *FeedResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetBackendTimeMillis

`func (o *FeedResponse) GetBackendTimeMillis() int64`

GetBackendTimeMillis returns the BackendTimeMillis field if non-nil, zero value otherwise.

### GetBackendTimeMillisOk

`func (o *FeedResponse) GetBackendTimeMillisOk() (*int64, bool)`

GetBackendTimeMillisOk returns a tuple with the BackendTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTimeMillis

`func (o *FeedResponse) SetBackendTimeMillis(v int64)`

SetBackendTimeMillis sets BackendTimeMillis field to given value.

### HasBackendTimeMillis

`func (o *FeedResponse) HasBackendTimeMillis() bool`

HasBackendTimeMillis returns a boolean if a field has been set.

### GetDatasourceAffinity

`func (o *FeedResponse) GetDatasourceAffinity() map[string]float32`

GetDatasourceAffinity returns the DatasourceAffinity field if non-nil, zero value otherwise.

### GetDatasourceAffinityOk

`func (o *FeedResponse) GetDatasourceAffinityOk() (*map[string]float32, bool)`

GetDatasourceAffinityOk returns a tuple with the DatasourceAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceAffinity

`func (o *FeedResponse) SetDatasourceAffinity(v map[string]float32)`

SetDatasourceAffinity sets DatasourceAffinity field to given value.

### HasDatasourceAffinity

`func (o *FeedResponse) HasDatasourceAffinity() bool`

HasDatasourceAffinity returns a boolean if a field has been set.

### GetCompanyResourcesCollectionId

`func (o *FeedResponse) GetCompanyResourcesCollectionId() int32`

GetCompanyResourcesCollectionId returns the CompanyResourcesCollectionId field if non-nil, zero value otherwise.

### GetCompanyResourcesCollectionIdOk

`func (o *FeedResponse) GetCompanyResourcesCollectionIdOk() (*int32, bool)`

GetCompanyResourcesCollectionIdOk returns a tuple with the CompanyResourcesCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyResourcesCollectionId

`func (o *FeedResponse) SetCompanyResourcesCollectionId(v int32)`

SetCompanyResourcesCollectionId sets CompanyResourcesCollectionId field to given value.

### HasCompanyResourcesCollectionId

`func (o *FeedResponse) HasCompanyResourcesCollectionId() bool`

HasCompanyResourcesCollectionId returns a boolean if a field has been set.

### GetFacetResults

`func (o *FeedResponse) GetFacetResults() map[string][]FacetResult`

GetFacetResults returns the FacetResults field if non-nil, zero value otherwise.

### GetFacetResultsOk

`func (o *FeedResponse) GetFacetResultsOk() (*map[string][]FacetResult, bool)`

GetFacetResultsOk returns a tuple with the FacetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetResults

`func (o *FeedResponse) SetFacetResults(v map[string][]FacetResult)`

SetFacetResults sets FacetResults field to given value.

### HasFacetResults

`func (o *FeedResponse) HasFacetResults() bool`

HasFacetResults returns a boolean if a field has been set.

### GetMentionsTimeWindowInHours

`func (o *FeedResponse) GetMentionsTimeWindowInHours() int32`

GetMentionsTimeWindowInHours returns the MentionsTimeWindowInHours field if non-nil, zero value otherwise.

### GetMentionsTimeWindowInHoursOk

`func (o *FeedResponse) GetMentionsTimeWindowInHoursOk() (*int32, bool)`

GetMentionsTimeWindowInHoursOk returns a tuple with the MentionsTimeWindowInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionsTimeWindowInHours

`func (o *FeedResponse) SetMentionsTimeWindowInHours(v int32)`

SetMentionsTimeWindowInHours sets MentionsTimeWindowInHours field to given value.

### HasMentionsTimeWindowInHours

`func (o *FeedResponse) HasMentionsTimeWindowInHours() bool`

HasMentionsTimeWindowInHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


