# IndexStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCrawledTime** | Pointer to **time.Time** | When the document was last crawled | [optional] 
**LastIndexedTime** | Pointer to **time.Time** | When the document was last indexed | [optional] 

## Methods

### NewIndexStatus

`func NewIndexStatus() *IndexStatus`

NewIndexStatus instantiates a new IndexStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexStatusWithDefaults

`func NewIndexStatusWithDefaults() *IndexStatus`

NewIndexStatusWithDefaults instantiates a new IndexStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCrawledTime

`func (o *IndexStatus) GetLastCrawledTime() time.Time`

GetLastCrawledTime returns the LastCrawledTime field if non-nil, zero value otherwise.

### GetLastCrawledTimeOk

`func (o *IndexStatus) GetLastCrawledTimeOk() (*time.Time, bool)`

GetLastCrawledTimeOk returns a tuple with the LastCrawledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCrawledTime

`func (o *IndexStatus) SetLastCrawledTime(v time.Time)`

SetLastCrawledTime sets LastCrawledTime field to given value.

### HasLastCrawledTime

`func (o *IndexStatus) HasLastCrawledTime() bool`

HasLastCrawledTime returns a boolean if a field has been set.

### GetLastIndexedTime

`func (o *IndexStatus) GetLastIndexedTime() time.Time`

GetLastIndexedTime returns the LastIndexedTime field if non-nil, zero value otherwise.

### GetLastIndexedTimeOk

`func (o *IndexStatus) GetLastIndexedTimeOk() (*time.Time, bool)`

GetLastIndexedTimeOk returns a tuple with the LastIndexedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexedTime

`func (o *IndexStatus) SetLastIndexedTime(v time.Time)`

SetLastIndexedTime sets LastIndexedTime field to given value.

### HasLastIndexedTime

`func (o *IndexStatus) HasLastIndexedTime() bool`

HasLastIndexedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


