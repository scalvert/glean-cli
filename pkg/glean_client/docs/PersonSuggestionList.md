# PersonSuggestionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**PeopleSuggestionCategory**](PeopleSuggestionCategory.md) |  | 
**People** | Pointer to [**[]Person**](Person.md) | Information about suggested users. | [optional] 

## Methods

### NewPersonSuggestionList

`func NewPersonSuggestionList(category PeopleSuggestionCategory, ) *PersonSuggestionList`

NewPersonSuggestionList instantiates a new PersonSuggestionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSuggestionListWithDefaults

`func NewPersonSuggestionListWithDefaults() *PersonSuggestionList`

NewPersonSuggestionListWithDefaults instantiates a new PersonSuggestionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PersonSuggestionList) GetCategory() PeopleSuggestionCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PersonSuggestionList) GetCategoryOk() (*PeopleSuggestionCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PersonSuggestionList) SetCategory(v PeopleSuggestionCategory)`

SetCategory sets Category field to given value.


### GetPeople

`func (o *PersonSuggestionList) GetPeople() []Person`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PersonSuggestionList) GetPeopleOk() (*[]Person, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PersonSuggestionList) SetPeople(v []Person)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PersonSuggestionList) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


