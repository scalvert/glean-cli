# ListEntitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Person**](Person.md) |  | [optional] 
**TeamResults** | Pointer to [**[]Team**](Team.md) |  | [optional] 
**CustomEntityResults** | Pointer to [**[]CustomEntity**](CustomEntity.md) |  | [optional] 
**FacetResults** | Pointer to [**[]FacetResult**](FacetResult.md) |  | [optional] 
**Cursor** | Pointer to **string** | Pagination cursor. A previously received opaque token representing the position in the overall results at which to start. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of entities available | [optional] 
**HasMoreResults** | Pointer to **bool** | Whether or not more entities can be fetched. | [optional] 
**SortOptions** | Pointer to [**[]EntitiesSortOrder**](EntitiesSortOrder.md) | Sort options from EntitiesSortOrder supported for this response. Default is empty list. | [optional] 
**CustomFacetNames** | Pointer to **[]string** | list of Person attributes that are custom setup by deployment | [optional] 

## Methods

### NewListEntitiesResponse

`func NewListEntitiesResponse() *ListEntitiesResponse`

NewListEntitiesResponse instantiates a new ListEntitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntitiesResponseWithDefaults

`func NewListEntitiesResponseWithDefaults() *ListEntitiesResponse`

NewListEntitiesResponseWithDefaults instantiates a new ListEntitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListEntitiesResponse) GetResults() []Person`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListEntitiesResponse) GetResultsOk() (*[]Person, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListEntitiesResponse) SetResults(v []Person)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListEntitiesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTeamResults

`func (o *ListEntitiesResponse) GetTeamResults() []Team`

GetTeamResults returns the TeamResults field if non-nil, zero value otherwise.

### GetTeamResultsOk

`func (o *ListEntitiesResponse) GetTeamResultsOk() (*[]Team, bool)`

GetTeamResultsOk returns a tuple with the TeamResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamResults

`func (o *ListEntitiesResponse) SetTeamResults(v []Team)`

SetTeamResults sets TeamResults field to given value.

### HasTeamResults

`func (o *ListEntitiesResponse) HasTeamResults() bool`

HasTeamResults returns a boolean if a field has been set.

### GetCustomEntityResults

`func (o *ListEntitiesResponse) GetCustomEntityResults() []CustomEntity`

GetCustomEntityResults returns the CustomEntityResults field if non-nil, zero value otherwise.

### GetCustomEntityResultsOk

`func (o *ListEntitiesResponse) GetCustomEntityResultsOk() (*[]CustomEntity, bool)`

GetCustomEntityResultsOk returns a tuple with the CustomEntityResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEntityResults

`func (o *ListEntitiesResponse) SetCustomEntityResults(v []CustomEntity)`

SetCustomEntityResults sets CustomEntityResults field to given value.

### HasCustomEntityResults

`func (o *ListEntitiesResponse) HasCustomEntityResults() bool`

HasCustomEntityResults returns a boolean if a field has been set.

### GetFacetResults

`func (o *ListEntitiesResponse) GetFacetResults() []FacetResult`

GetFacetResults returns the FacetResults field if non-nil, zero value otherwise.

### GetFacetResultsOk

`func (o *ListEntitiesResponse) GetFacetResultsOk() (*[]FacetResult, bool)`

GetFacetResultsOk returns a tuple with the FacetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetResults

`func (o *ListEntitiesResponse) SetFacetResults(v []FacetResult)`

SetFacetResults sets FacetResults field to given value.

### HasFacetResults

`func (o *ListEntitiesResponse) HasFacetResults() bool`

HasFacetResults returns a boolean if a field has been set.

### GetCursor

`func (o *ListEntitiesResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListEntitiesResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListEntitiesResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ListEntitiesResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListEntitiesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListEntitiesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListEntitiesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListEntitiesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetHasMoreResults

`func (o *ListEntitiesResponse) GetHasMoreResults() bool`

GetHasMoreResults returns the HasMoreResults field if non-nil, zero value otherwise.

### GetHasMoreResultsOk

`func (o *ListEntitiesResponse) GetHasMoreResultsOk() (*bool, bool)`

GetHasMoreResultsOk returns a tuple with the HasMoreResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoreResults

`func (o *ListEntitiesResponse) SetHasMoreResults(v bool)`

SetHasMoreResults sets HasMoreResults field to given value.

### HasHasMoreResults

`func (o *ListEntitiesResponse) HasHasMoreResults() bool`

HasHasMoreResults returns a boolean if a field has been set.

### GetSortOptions

`func (o *ListEntitiesResponse) GetSortOptions() []EntitiesSortOrder`

GetSortOptions returns the SortOptions field if non-nil, zero value otherwise.

### GetSortOptionsOk

`func (o *ListEntitiesResponse) GetSortOptionsOk() (*[]EntitiesSortOrder, bool)`

GetSortOptionsOk returns a tuple with the SortOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOptions

`func (o *ListEntitiesResponse) SetSortOptions(v []EntitiesSortOrder)`

SetSortOptions sets SortOptions field to given value.

### HasSortOptions

`func (o *ListEntitiesResponse) HasSortOptions() bool`

HasSortOptions returns a boolean if a field has been set.

### GetCustomFacetNames

`func (o *ListEntitiesResponse) GetCustomFacetNames() []string`

GetCustomFacetNames returns the CustomFacetNames field if non-nil, zero value otherwise.

### GetCustomFacetNamesOk

`func (o *ListEntitiesResponse) GetCustomFacetNamesOk() (*[]string, bool)`

GetCustomFacetNamesOk returns a tuple with the CustomFacetNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFacetNames

`func (o *ListEntitiesResponse) SetCustomFacetNames(v []string)`

SetCustomFacetNames sets CustomFacetNames field to given value.

### HasCustomFacetNames

`func (o *ListEntitiesResponse) HasCustomFacetNames() bool`

HasCustomFacetNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


