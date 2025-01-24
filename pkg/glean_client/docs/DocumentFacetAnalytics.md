# DocumentFacetAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facet** | Pointer to [**FacetFilter**](FacetFilter.md) |  | [optional] 
**Analytics** | Pointer to [**DocumentAnalytics**](DocumentAnalytics.md) |  | [optional] 

## Methods

### NewDocumentFacetAnalytics

`func NewDocumentFacetAnalytics() *DocumentFacetAnalytics`

NewDocumentFacetAnalytics instantiates a new DocumentFacetAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentFacetAnalyticsWithDefaults

`func NewDocumentFacetAnalyticsWithDefaults() *DocumentFacetAnalytics`

NewDocumentFacetAnalyticsWithDefaults instantiates a new DocumentFacetAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacet

`func (o *DocumentFacetAnalytics) GetFacet() FacetFilter`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *DocumentFacetAnalytics) GetFacetOk() (*FacetFilter, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *DocumentFacetAnalytics) SetFacet(v FacetFilter)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *DocumentFacetAnalytics) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetAnalytics

`func (o *DocumentFacetAnalytics) GetAnalytics() DocumentAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *DocumentFacetAnalytics) GetAnalyticsOk() (*DocumentAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *DocumentFacetAnalytics) SetAnalytics(v DocumentAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *DocumentFacetAnalytics) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


