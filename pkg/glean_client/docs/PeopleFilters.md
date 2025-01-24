# PeopleFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Facets used for filtering people search | [optional] 
**Query** | Pointer to **string** | A query string to search for people that each person in the response must conform to. An empty query does not filter any people. | [optional] 

## Methods

### NewPeopleFilters

`func NewPeopleFilters() *PeopleFilters`

NewPeopleFilters instantiates a new PeopleFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleFiltersWithDefaults

`func NewPeopleFiltersWithDefaults() *PeopleFilters`

NewPeopleFiltersWithDefaults instantiates a new PeopleFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *PeopleFilters) GetFilter() []FacetFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PeopleFilters) GetFilterOk() (*[]FacetFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PeopleFilters) SetFilter(v []FacetFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PeopleFilters) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetQuery

`func (o *PeopleFilters) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PeopleFilters) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PeopleFilters) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PeopleFilters) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


