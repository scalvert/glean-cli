# GetDocumentAnalyticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentSpecs** | [**[]DocumentSpec**](DocumentSpec.md) | The specification for the documents for which analytics will be retrieved. | 
**DayRange** | [**Period**](Period.md) |  | 
**WithClickerCounts** | Pointer to **bool** | Whether response should include click information or not. Default is to not include click information. | [optional] 
**WithImpressionCounts** | Pointer to **bool** | Whether response should include impression information or not. Default is to not include impression information. | [optional] 
**WithFacetAggregations** | Pointer to **bool** | Whether the results will include aggregate counts/info for facets like location, department, etc. | [optional] 
**WithVisitCounts** | Pointer to **bool** | Whether response should include visit counts or not. Default is to return only visitor counts. | [optional] 

## Methods

### NewGetDocumentAnalyticsRequest

`func NewGetDocumentAnalyticsRequest(documentSpecs []DocumentSpec, dayRange Period, ) *GetDocumentAnalyticsRequest`

NewGetDocumentAnalyticsRequest instantiates a new GetDocumentAnalyticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentAnalyticsRequestWithDefaults

`func NewGetDocumentAnalyticsRequestWithDefaults() *GetDocumentAnalyticsRequest`

NewGetDocumentAnalyticsRequestWithDefaults instantiates a new GetDocumentAnalyticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentSpecs

`func (o *GetDocumentAnalyticsRequest) GetDocumentSpecs() []DocumentSpec`

GetDocumentSpecs returns the DocumentSpecs field if non-nil, zero value otherwise.

### GetDocumentSpecsOk

`func (o *GetDocumentAnalyticsRequest) GetDocumentSpecsOk() (*[]DocumentSpec, bool)`

GetDocumentSpecsOk returns a tuple with the DocumentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSpecs

`func (o *GetDocumentAnalyticsRequest) SetDocumentSpecs(v []DocumentSpec)`

SetDocumentSpecs sets DocumentSpecs field to given value.


### GetDayRange

`func (o *GetDocumentAnalyticsRequest) GetDayRange() Period`

GetDayRange returns the DayRange field if non-nil, zero value otherwise.

### GetDayRangeOk

`func (o *GetDocumentAnalyticsRequest) GetDayRangeOk() (*Period, bool)`

GetDayRangeOk returns a tuple with the DayRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayRange

`func (o *GetDocumentAnalyticsRequest) SetDayRange(v Period)`

SetDayRange sets DayRange field to given value.


### GetWithClickerCounts

`func (o *GetDocumentAnalyticsRequest) GetWithClickerCounts() bool`

GetWithClickerCounts returns the WithClickerCounts field if non-nil, zero value otherwise.

### GetWithClickerCountsOk

`func (o *GetDocumentAnalyticsRequest) GetWithClickerCountsOk() (*bool, bool)`

GetWithClickerCountsOk returns a tuple with the WithClickerCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithClickerCounts

`func (o *GetDocumentAnalyticsRequest) SetWithClickerCounts(v bool)`

SetWithClickerCounts sets WithClickerCounts field to given value.

### HasWithClickerCounts

`func (o *GetDocumentAnalyticsRequest) HasWithClickerCounts() bool`

HasWithClickerCounts returns a boolean if a field has been set.

### GetWithImpressionCounts

`func (o *GetDocumentAnalyticsRequest) GetWithImpressionCounts() bool`

GetWithImpressionCounts returns the WithImpressionCounts field if non-nil, zero value otherwise.

### GetWithImpressionCountsOk

`func (o *GetDocumentAnalyticsRequest) GetWithImpressionCountsOk() (*bool, bool)`

GetWithImpressionCountsOk returns a tuple with the WithImpressionCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithImpressionCounts

`func (o *GetDocumentAnalyticsRequest) SetWithImpressionCounts(v bool)`

SetWithImpressionCounts sets WithImpressionCounts field to given value.

### HasWithImpressionCounts

`func (o *GetDocumentAnalyticsRequest) HasWithImpressionCounts() bool`

HasWithImpressionCounts returns a boolean if a field has been set.

### GetWithFacetAggregations

`func (o *GetDocumentAnalyticsRequest) GetWithFacetAggregations() bool`

GetWithFacetAggregations returns the WithFacetAggregations field if non-nil, zero value otherwise.

### GetWithFacetAggregationsOk

`func (o *GetDocumentAnalyticsRequest) GetWithFacetAggregationsOk() (*bool, bool)`

GetWithFacetAggregationsOk returns a tuple with the WithFacetAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithFacetAggregations

`func (o *GetDocumentAnalyticsRequest) SetWithFacetAggregations(v bool)`

SetWithFacetAggregations sets WithFacetAggregations field to given value.

### HasWithFacetAggregations

`func (o *GetDocumentAnalyticsRequest) HasWithFacetAggregations() bool`

HasWithFacetAggregations returns a boolean if a field has been set.

### GetWithVisitCounts

`func (o *GetDocumentAnalyticsRequest) GetWithVisitCounts() bool`

GetWithVisitCounts returns the WithVisitCounts field if non-nil, zero value otherwise.

### GetWithVisitCountsOk

`func (o *GetDocumentAnalyticsRequest) GetWithVisitCountsOk() (*bool, bool)`

GetWithVisitCountsOk returns a tuple with the WithVisitCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithVisitCounts

`func (o *GetDocumentAnalyticsRequest) SetWithVisitCounts(v bool)`

SetWithVisitCounts sets WithVisitCounts field to given value.

### HasWithVisitCounts

`func (o *GetDocumentAnalyticsRequest) HasWithVisitCounts() bool`

HasWithVisitCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


