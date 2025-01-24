# FeedResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | Category of the result, one of the requested categories in incoming request. | 
**PrimaryEntry** | [**FeedEntry**](FeedEntry.md) |  | 
**SecondaryEntries** | Pointer to [**[]FeedEntry**](FeedEntry.md) | Secondary entries for the result e.g. suggested docs for the calendar, carousel. | [optional] 
**Rank** | Pointer to **int32** | Rank of the result. Rank is suggested by server. Client side rank may differ. | [optional] 
**FacetResults** | Pointer to [**[]FacetResult**](FacetResult.md) | DEPRECATED - List of facets that can be used to filter the entry&#39;s content. | [optional] 

## Methods

### NewFeedResult

`func NewFeedResult(category string, primaryEntry FeedEntry, ) *FeedResult`

NewFeedResult instantiates a new FeedResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedResultWithDefaults

`func NewFeedResultWithDefaults() *FeedResult`

NewFeedResultWithDefaults instantiates a new FeedResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *FeedResult) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FeedResult) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FeedResult) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetPrimaryEntry

`func (o *FeedResult) GetPrimaryEntry() FeedEntry`

GetPrimaryEntry returns the PrimaryEntry field if non-nil, zero value otherwise.

### GetPrimaryEntryOk

`func (o *FeedResult) GetPrimaryEntryOk() (*FeedEntry, bool)`

GetPrimaryEntryOk returns a tuple with the PrimaryEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEntry

`func (o *FeedResult) SetPrimaryEntry(v FeedEntry)`

SetPrimaryEntry sets PrimaryEntry field to given value.


### GetSecondaryEntries

`func (o *FeedResult) GetSecondaryEntries() []FeedEntry`

GetSecondaryEntries returns the SecondaryEntries field if non-nil, zero value otherwise.

### GetSecondaryEntriesOk

`func (o *FeedResult) GetSecondaryEntriesOk() (*[]FeedEntry, bool)`

GetSecondaryEntriesOk returns a tuple with the SecondaryEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEntries

`func (o *FeedResult) SetSecondaryEntries(v []FeedEntry)`

SetSecondaryEntries sets SecondaryEntries field to given value.

### HasSecondaryEntries

`func (o *FeedResult) HasSecondaryEntries() bool`

HasSecondaryEntries returns a boolean if a field has been set.

### GetRank

`func (o *FeedResult) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *FeedResult) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *FeedResult) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *FeedResult) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetFacetResults

`func (o *FeedResult) GetFacetResults() []FacetResult`

GetFacetResults returns the FacetResults field if non-nil, zero value otherwise.

### GetFacetResultsOk

`func (o *FeedResult) GetFacetResultsOk() (*[]FacetResult, bool)`

GetFacetResultsOk returns a tuple with the FacetResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetResults

`func (o *FeedResult) SetFacetResults(v []FacetResult)`

SetFacetResults sets FacetResults field to given value.

### HasFacetResults

`func (o *FeedResult) HasFacetResults() bool`

HasFacetResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


