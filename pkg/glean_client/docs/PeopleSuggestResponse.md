# PeopleSuggestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestions** | Pointer to [**[]PersonSuggestionList**](PersonSuggestionList.md) | Information about people suggestions for asked categories. | [optional] 

## Methods

### NewPeopleSuggestResponse

`func NewPeopleSuggestResponse() *PeopleSuggestResponse`

NewPeopleSuggestResponse instantiates a new PeopleSuggestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleSuggestResponseWithDefaults

`func NewPeopleSuggestResponseWithDefaults() *PeopleSuggestResponse`

NewPeopleSuggestResponseWithDefaults instantiates a new PeopleSuggestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestions

`func (o *PeopleSuggestResponse) GetSuggestions() []PersonSuggestionList`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *PeopleSuggestResponse) GetSuggestionsOk() (*[]PersonSuggestionList, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *PeopleSuggestResponse) SetSuggestions(v []PersonSuggestionList)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *PeopleSuggestResponse) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


