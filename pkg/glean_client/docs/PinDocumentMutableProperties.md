# PinDocumentMutableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | Pointer to **[]string** | The query strings for which the pinned result will show. | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see the pinned document. Values are taken from the corresponding filters in people search. | [optional] 

## Methods

### NewPinDocumentMutableProperties

`func NewPinDocumentMutableProperties() *PinDocumentMutableProperties`

NewPinDocumentMutableProperties instantiates a new PinDocumentMutableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinDocumentMutablePropertiesWithDefaults

`func NewPinDocumentMutablePropertiesWithDefaults() *PinDocumentMutableProperties`

NewPinDocumentMutablePropertiesWithDefaults instantiates a new PinDocumentMutableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *PinDocumentMutableProperties) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *PinDocumentMutableProperties) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *PinDocumentMutableProperties) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *PinDocumentMutableProperties) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *PinDocumentMutableProperties) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *PinDocumentMutableProperties) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *PinDocumentMutableProperties) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *PinDocumentMutableProperties) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


