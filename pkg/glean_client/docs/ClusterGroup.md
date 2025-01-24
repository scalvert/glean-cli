# ClusterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusteredResults** | Pointer to [**[]SearchResult**](SearchResult.md) | A list of results that should be displayed as associated with this result. | [optional] 
**ClusterType** | Pointer to [**ClusterTypeEnum**](ClusterTypeEnum.md) |  | [optional] 
**VisibleCountHint** | **int32** | The default number of results to display before truncating and showing a \&quot;see more\&quot; link | 

## Methods

### NewClusterGroup

`func NewClusterGroup(visibleCountHint int32, ) *ClusterGroup`

NewClusterGroup instantiates a new ClusterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterGroupWithDefaults

`func NewClusterGroupWithDefaults() *ClusterGroup`

NewClusterGroupWithDefaults instantiates a new ClusterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusteredResults

`func (o *ClusterGroup) GetClusteredResults() []SearchResult`

GetClusteredResults returns the ClusteredResults field if non-nil, zero value otherwise.

### GetClusteredResultsOk

`func (o *ClusterGroup) GetClusteredResultsOk() (*[]SearchResult, bool)`

GetClusteredResultsOk returns a tuple with the ClusteredResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusteredResults

`func (o *ClusterGroup) SetClusteredResults(v []SearchResult)`

SetClusteredResults sets ClusteredResults field to given value.

### HasClusteredResults

`func (o *ClusterGroup) HasClusteredResults() bool`

HasClusteredResults returns a boolean if a field has been set.

### GetClusterType

`func (o *ClusterGroup) GetClusterType() ClusterTypeEnum`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *ClusterGroup) GetClusterTypeOk() (*ClusterTypeEnum, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *ClusterGroup) SetClusterType(v ClusterTypeEnum)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *ClusterGroup) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetVisibleCountHint

`func (o *ClusterGroup) GetVisibleCountHint() int32`

GetVisibleCountHint returns the VisibleCountHint field if non-nil, zero value otherwise.

### GetVisibleCountHintOk

`func (o *ClusterGroup) GetVisibleCountHintOk() (*int32, bool)`

GetVisibleCountHintOk returns a tuple with the VisibleCountHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleCountHint

`func (o *ClusterGroup) SetVisibleCountHint(v int32)`

SetVisibleCountHint sets VisibleCountHint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


