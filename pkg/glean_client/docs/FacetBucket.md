# FacetBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Estimated number of results in this facet. | [optional] 
**Datasource** | Pointer to **string** | The datasource the value belongs to. This will be used by the all tab to show types across all datasources. | [optional] 
**Percentage** | Pointer to **int32** | Estimated percentage of results in this facet. | [optional] 
**Value** | Pointer to [**FacetValue**](FacetValue.md) |  | [optional] 

## Methods

### NewFacetBucket

`func NewFacetBucket() *FacetBucket`

NewFacetBucket instantiates a new FacetBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetBucketWithDefaults

`func NewFacetBucketWithDefaults() *FacetBucket`

NewFacetBucketWithDefaults instantiates a new FacetBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FacetBucket) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FacetBucket) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FacetBucket) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FacetBucket) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDatasource

`func (o *FacetBucket) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *FacetBucket) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *FacetBucket) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *FacetBucket) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetPercentage

`func (o *FacetBucket) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *FacetBucket) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *FacetBucket) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *FacetBucket) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetValue

`func (o *FacetBucket) GetValue() FacetValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FacetBucket) GetValueOk() (*FacetValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FacetBucket) SetValue(v FacetValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *FacetBucket) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


