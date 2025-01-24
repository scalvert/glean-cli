# FacetResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceName** | Pointer to **string** | The source of this facet (e.g. container_name, type, last_updated_at). | [optional] 
**OperatorName** | Pointer to **string** | How to display this facet. Currently supportes &#39;SelectSingle&#39; and &#39;SelectMultiple&#39;. | [optional] 
**Buckets** | Pointer to [**[]FacetBucket**](FacetBucket.md) | A list of unique buckets that exist within this result set. | [optional] 
**HasMoreBuckets** | Pointer to **bool** | Returns true if more buckets exist than those returned. Additional buckets can be retrieve by requesting again with a higher facetBucketSize. | [optional] 
**GroupName** | Pointer to **string** | For most facets this will be the empty string, meaning the facet is high-level and applies to all documents for the datasource. When non-empty, this is used to group facets together (i.e. group facets for each doctype for a certain datasource) | [optional] 

## Methods

### NewFacetResult

`func NewFacetResult() *FacetResult`

NewFacetResult instantiates a new FacetResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetResultWithDefaults

`func NewFacetResultWithDefaults() *FacetResult`

NewFacetResultWithDefaults instantiates a new FacetResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceName

`func (o *FacetResult) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *FacetResult) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *FacetResult) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *FacetResult) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetOperatorName

`func (o *FacetResult) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *FacetResult) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *FacetResult) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.

### HasOperatorName

`func (o *FacetResult) HasOperatorName() bool`

HasOperatorName returns a boolean if a field has been set.

### GetBuckets

`func (o *FacetResult) GetBuckets() []FacetBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *FacetResult) GetBucketsOk() (*[]FacetBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *FacetResult) SetBuckets(v []FacetBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *FacetResult) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetHasMoreBuckets

`func (o *FacetResult) GetHasMoreBuckets() bool`

GetHasMoreBuckets returns the HasMoreBuckets field if non-nil, zero value otherwise.

### GetHasMoreBucketsOk

`func (o *FacetResult) GetHasMoreBucketsOk() (*bool, bool)`

GetHasMoreBucketsOk returns a tuple with the HasMoreBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreBuckets

`func (o *FacetResult) SetHasMoreBuckets(v bool)`

SetHasMoreBuckets sets HasMoreBuckets field to given value.

### HasHasMoreBuckets

`func (o *FacetResult) HasHasMoreBuckets() bool`

HasHasMoreBuckets returns a boolean if a field has been set.

### GetGroupName

`func (o *FacetResult) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *FacetResult) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *FacetResult) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *FacetResult) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


