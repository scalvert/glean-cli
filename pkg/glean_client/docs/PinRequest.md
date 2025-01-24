# PinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | Pointer to **[]string** | The query strings for which the pinned result will show. | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see the pinned document. Values are taken from the corresponding filters in people search. | [optional] 
**DocumentId** | Pointer to **string** | The document to be pinned. | [optional] 

## Methods

### NewPinRequest

`func NewPinRequest() *PinRequest`

NewPinRequest instantiates a new PinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinRequestWithDefaults

`func NewPinRequestWithDefaults() *PinRequest`

NewPinRequestWithDefaults instantiates a new PinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *PinRequest) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *PinRequest) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *PinRequest) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *PinRequest) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *PinRequest) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *PinRequest) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *PinRequest) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *PinRequest) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetDocumentId

`func (o *PinRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *PinRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *PinRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *PinRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


