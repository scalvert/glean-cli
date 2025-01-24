# AskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsQuestion** | **bool** | Whether or not the query was a question. | 
**Question** | Pointer to **string** | The part of the query which was used as a question for search | [optional] 
**SearchResponse** | Pointer to [**SearchResponse**](SearchResponse.md) |  | [optional] 

## Methods

### NewAskResponse

`func NewAskResponse(isQuestion bool, ) *AskResponse`

NewAskResponse instantiates a new AskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskResponseWithDefaults

`func NewAskResponseWithDefaults() *AskResponse`

NewAskResponseWithDefaults instantiates a new AskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsQuestion

`func (o *AskResponse) GetIsQuestion() bool`

GetIsQuestion returns the IsQuestion field if non-nil, zero value otherwise.

### GetIsQuestionOk

`func (o *AskResponse) GetIsQuestionOk() (*bool, bool)`

GetIsQuestionOk returns a tuple with the IsQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQuestion

`func (o *AskResponse) SetIsQuestion(v bool)`

SetIsQuestion sets IsQuestion field to given value.


### GetQuestion

`func (o *AskResponse) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AskResponse) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AskResponse) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *AskResponse) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetSearchResponse

`func (o *AskResponse) GetSearchResponse() SearchResponse`

GetSearchResponse returns the SearchResponse field if non-nil, zero value otherwise.

### GetSearchResponseOk

`func (o *AskResponse) GetSearchResponseOk() (*SearchResponse, bool)`

GetSearchResponseOk returns a tuple with the SearchResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResponse

`func (o *AskResponse) SetSearchResponse(v SearchResponse)`

SetSearchResponse sets SearchResponse field to given value.

### HasSearchResponse

`func (o *AskResponse) HasSearchResponse() bool`

HasSearchResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


