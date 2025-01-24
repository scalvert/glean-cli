# ListShortcutsPaginatedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeFields** | Pointer to **[]string** | Array of fields/data to be included in response that are not included by default | [optional] 
**PageSize** | **int32** |  | 
**Cursor** | Pointer to **string** | A token specifying the position in the overall results to start at. Received from the endpoint and iterated back. Currently being used as page no (as we implement offset pagination) | [optional] 
**Filters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | A list of filters for the query. An AND is assumed between different filters. We support filters on Go Link name, author, department and type. | [optional] 
**Sort** | Pointer to [**SortOptions**](SortOptions.md) |  | [optional] 
**Query** | Pointer to **string** | Search query that should be a substring in atleast one of the fields (alias , inputAlias, destinationUrl, description). Empty query does not filter shortcuts. | [optional] 

## Methods

### NewListShortcutsPaginatedRequest

`func NewListShortcutsPaginatedRequest(pageSize int32, ) *ListShortcutsPaginatedRequest`

NewListShortcutsPaginatedRequest instantiates a new ListShortcutsPaginatedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListShortcutsPaginatedRequestWithDefaults

`func NewListShortcutsPaginatedRequestWithDefaults() *ListShortcutsPaginatedRequest`

NewListShortcutsPaginatedRequestWithDefaults instantiates a new ListShortcutsPaginatedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeFields

`func (o *ListShortcutsPaginatedRequest) GetIncludeFields() []string`

GetIncludeFields returns the IncludeFields field if non-nil, zero value otherwise.

### GetIncludeFieldsOk

`func (o *ListShortcutsPaginatedRequest) GetIncludeFieldsOk() (*[]string, bool)`

GetIncludeFieldsOk returns a tuple with the IncludeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFields

`func (o *ListShortcutsPaginatedRequest) SetIncludeFields(v []string)`

SetIncludeFields sets IncludeFields field to given value.

### HasIncludeFields

`func (o *ListShortcutsPaginatedRequest) HasIncludeFields() bool`

HasIncludeFields returns a boolean if a field has been set.

### GetPageSize

`func (o *ListShortcutsPaginatedRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListShortcutsPaginatedRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListShortcutsPaginatedRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetCursor

`func (o *ListShortcutsPaginatedRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListShortcutsPaginatedRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListShortcutsPaginatedRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListShortcutsPaginatedRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetFilters

`func (o *ListShortcutsPaginatedRequest) GetFilters() []FacetFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListShortcutsPaginatedRequest) GetFiltersOk() (*[]FacetFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListShortcutsPaginatedRequest) SetFilters(v []FacetFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListShortcutsPaginatedRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSort

`func (o *ListShortcutsPaginatedRequest) GetSort() SortOptions`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListShortcutsPaginatedRequest) GetSortOk() (*SortOptions, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListShortcutsPaginatedRequest) SetSort(v SortOptions)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListShortcutsPaginatedRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetQuery

`func (o *ListShortcutsPaginatedRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListShortcutsPaginatedRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListShortcutsPaginatedRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ListShortcutsPaginatedRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


