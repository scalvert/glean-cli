# QuerySuggestionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestions** | [**[]QuerySuggestion**](QuerySuggestion.md) |  | 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 

## Methods

### NewQuerySuggestionList

`func NewQuerySuggestionList(suggestions []QuerySuggestion, ) *QuerySuggestionList`

NewQuerySuggestionList instantiates a new QuerySuggestionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySuggestionListWithDefaults

`func NewQuerySuggestionListWithDefaults() *QuerySuggestionList`

NewQuerySuggestionListWithDefaults instantiates a new QuerySuggestionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestions

`func (o *QuerySuggestionList) GetSuggestions() []QuerySuggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *QuerySuggestionList) GetSuggestionsOk() (*[]QuerySuggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *QuerySuggestionList) SetSuggestions(v []QuerySuggestion)`

SetSuggestions sets Suggestions field to given value.


### GetPerson

`func (o *QuerySuggestionList) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *QuerySuggestionList) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *QuerySuggestionList) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *QuerySuggestionList) HasPerson() bool`

HasPerson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


