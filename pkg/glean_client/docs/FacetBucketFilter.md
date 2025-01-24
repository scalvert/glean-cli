# FacetBucketFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facet** | Pointer to **string** | The facet whose buckets should be filtered. | [optional] 
**Prefix** | Pointer to **string** | The per-term prefix that facet buckets should be filtered on. | [optional] 

## Methods

### NewFacetBucketFilter

`func NewFacetBucketFilter() *FacetBucketFilter`

NewFacetBucketFilter instantiates a new FacetBucketFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetBucketFilterWithDefaults

`func NewFacetBucketFilterWithDefaults() *FacetBucketFilter`

NewFacetBucketFilterWithDefaults instantiates a new FacetBucketFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacet

`func (o *FacetBucketFilter) GetFacet() string`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *FacetBucketFilter) GetFacetOk() (*string, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *FacetBucketFilter) SetFacet(v string)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *FacetBucketFilter) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetPrefix

`func (o *FacetBucketFilter) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *FacetBucketFilter) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *FacetBucketFilter) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *FacetBucketFilter) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


