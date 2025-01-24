# PeopleSuggestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | [**[]PeopleSuggestionCategory**](PeopleSuggestionCategory.md) | Categories of data requested. Request can include single or multiple categories. | 
**Departments** | Pointer to **[]string** | Departments that the data is requested for. If empty, corresponds to whole company. | [optional] 

## Methods

### NewPeopleSuggestRequest

`func NewPeopleSuggestRequest(categories []PeopleSuggestionCategory, ) *PeopleSuggestRequest`

NewPeopleSuggestRequest instantiates a new PeopleSuggestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleSuggestRequestWithDefaults

`func NewPeopleSuggestRequestWithDefaults() *PeopleSuggestRequest`

NewPeopleSuggestRequestWithDefaults instantiates a new PeopleSuggestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *PeopleSuggestRequest) GetCategories() []PeopleSuggestionCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PeopleSuggestRequest) GetCategoriesOk() (*[]PeopleSuggestionCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PeopleSuggestRequest) SetCategories(v []PeopleSuggestionCategory)`

SetCategories sets Categories field to given value.


### GetDepartments

`func (o *PeopleSuggestRequest) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *PeopleSuggestRequest) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *PeopleSuggestRequest) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *PeopleSuggestRequest) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


