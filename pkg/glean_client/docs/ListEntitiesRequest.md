# ListEntitiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**[]FacetFilter**](FacetFilter.md) |  | [optional] 
**Sort** | Pointer to [**[]SortOptions**](SortOptions.md) | Use EntitiesSortOrder enum for SortOptions.sortBy | [optional] 
**EntityType** | Pointer to **string** |  | [optional] [default to "PEOPLE"]
**Datasource** | Pointer to **string** | The datasource associated with the entity type, most commonly used with CUSTOM_ENTITIES | [optional] 
**Query** | Pointer to **string** | A query string to search for entities that each entity in the response must conform to. An empty query does not filter any entities. | [optional] 
**IncludeFields** | Pointer to **[]string** | List of entity fields to return (that aren&#39;t returned by default) | [optional] 
**PageSize** | Pointer to **int32** | Hint to the server about how many results to send back. Server may return less. | [optional] 
**Cursor** | Pointer to **string** | Pagination cursor. A previously received opaque token representing the position in the overall results at which to start. | [optional] 
**Source** | Pointer to **string** | A string denoting the search surface from which the endpoint is called. | [optional] 

## Methods

### NewListEntitiesRequest

`func NewListEntitiesRequest() *ListEntitiesRequest`

NewListEntitiesRequest instantiates a new ListEntitiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitiesRequestWithDefaults

`func NewListEntitiesRequestWithDefaults() *ListEntitiesRequest`

NewListEntitiesRequestWithDefaults instantiates a new ListEntitiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ListEntitiesRequest) GetFilter() []FacetFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListEntitiesRequest) GetFilterOk() (*[]FacetFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListEntitiesRequest) SetFilter(v []FacetFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListEntitiesRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *ListEntitiesRequest) GetSort() []SortOptions`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListEntitiesRequest) GetSortOk() (*[]SortOptions, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListEntitiesRequest) SetSort(v []SortOptions)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListEntitiesRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetEntityType

`func (o *ListEntitiesRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ListEntitiesRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ListEntitiesRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ListEntitiesRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetDatasource

`func (o *ListEntitiesRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *ListEntitiesRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *ListEntitiesRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *ListEntitiesRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetQuery

`func (o *ListEntitiesRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListEntitiesRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListEntitiesRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListEntitiesRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetIncludeFields

`func (o *ListEntitiesRequest) GetIncludeFields() []string`

GetIncludeFields returns the IncludeFields field if non-nil, zero value otherwise.

### GetIncludeFieldsOk

`func (o *ListEntitiesRequest) GetIncludeFieldsOk() (*[]string, bool)`

GetIncludeFieldsOk returns a tuple with the IncludeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFields

`func (o *ListEntitiesRequest) SetIncludeFields(v []string)`

SetIncludeFields sets IncludeFields field to given value.

### HasIncludeFields

`func (o *ListEntitiesRequest) HasIncludeFields() bool`

HasIncludeFields returns a boolean if a field has been set.

### GetPageSize

`func (o *ListEntitiesRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListEntitiesRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListEntitiesRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListEntitiesRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetCursor

`func (o *ListEntitiesRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListEntitiesRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListEntitiesRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListEntitiesRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetSource

`func (o *ListEntitiesRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListEntitiesRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListEntitiesRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListEntitiesRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


