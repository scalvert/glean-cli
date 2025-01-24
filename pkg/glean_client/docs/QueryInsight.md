# QueryInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The query string the information is about. | 
**SearchCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 
**SearchorCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 
**SearchWithClickCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 
**ClickCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 
**SimilarQueries** | Pointer to [**[]QueryInsight**](QueryInsight.md) | list of similar queries to current one. | [optional] 

## Methods

### NewQueryInsight

`func NewQueryInsight(query string, ) *QueryInsight`

NewQueryInsight instantiates a new QueryInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryInsightWithDefaults

`func NewQueryInsightWithDefaults() *QueryInsight`

NewQueryInsightWithDefaults instantiates a new QueryInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *QueryInsight) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryInsight) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryInsight) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetSearchCount

`func (o *QueryInsight) GetSearchCount() CountInfo`

GetSearchCount returns the SearchCount field if non-nil, zero value otherwise.

### GetSearchCountOk

`func (o *QueryInsight) GetSearchCountOk() (*CountInfo, bool)`

GetSearchCountOk returns a tuple with the SearchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCount

`func (o *QueryInsight) SetSearchCount(v CountInfo)`

SetSearchCount sets SearchCount field to given value.

### HasSearchCount

`func (o *QueryInsight) HasSearchCount() bool`

HasSearchCount returns a boolean if a field has been set.

### GetSearchorCount

`func (o *QueryInsight) GetSearchorCount() CountInfo`

GetSearchorCount returns the SearchorCount field if non-nil, zero value otherwise.

### GetSearchorCountOk

`func (o *QueryInsight) GetSearchorCountOk() (*CountInfo, bool)`

GetSearchorCountOk returns a tuple with the SearchorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchorCount

`func (o *QueryInsight) SetSearchorCount(v CountInfo)`

SetSearchorCount sets SearchorCount field to given value.

### HasSearchorCount

`func (o *QueryInsight) HasSearchorCount() bool`

HasSearchorCount returns a boolean if a field has been set.

### GetSearchWithClickCount

`func (o *QueryInsight) GetSearchWithClickCount() CountInfo`

GetSearchWithClickCount returns the SearchWithClickCount field if non-nil, zero value otherwise.

### GetSearchWithClickCountOk

`func (o *QueryInsight) GetSearchWithClickCountOk() (*CountInfo, bool)`

GetSearchWithClickCountOk returns a tuple with the SearchWithClickCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchWithClickCount

`func (o *QueryInsight) SetSearchWithClickCount(v CountInfo)`

SetSearchWithClickCount sets SearchWithClickCount field to given value.

### HasSearchWithClickCount

`func (o *QueryInsight) HasSearchWithClickCount() bool`

HasSearchWithClickCount returns a boolean if a field has been set.

### GetClickCount

`func (o *QueryInsight) GetClickCount() CountInfo`

GetClickCount returns the ClickCount field if non-nil, zero value otherwise.

### GetClickCountOk

`func (o *QueryInsight) GetClickCountOk() (*CountInfo, bool)`

GetClickCountOk returns a tuple with the ClickCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickCount

`func (o *QueryInsight) SetClickCount(v CountInfo)`

SetClickCount sets ClickCount field to given value.

### HasClickCount

`func (o *QueryInsight) HasClickCount() bool`

HasClickCount returns a boolean if a field has been set.

### GetSimilarQueries

`func (o *QueryInsight) GetSimilarQueries() []QueryInsight`

GetSimilarQueries returns the SimilarQueries field if non-nil, zero value otherwise.

### GetSimilarQueriesOk

`func (o *QueryInsight) GetSimilarQueriesOk() (*[]QueryInsight, bool)`

GetSimilarQueriesOk returns a tuple with the SimilarQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarQueries

`func (o *QueryInsight) SetSimilarQueries(v []QueryInsight)`

SetSimilarQueries sets SimilarQueries field to given value.

### HasSimilarQueries

`func (o *QueryInsight) HasSimilarQueries() bool`

HasSimilarQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


