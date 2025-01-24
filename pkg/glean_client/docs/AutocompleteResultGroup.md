# AutocompleteResultGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartIndex** | Pointer to **int32** | The inclusive start index of the range. | [optional] 
**EndIndex** | Pointer to **int32** | The exclusive end index of the range. | [optional] 
**Title** | Pointer to **string** | The title of the result group to be displayed. Empty means no title. | [optional] 

## Methods

### NewAutocompleteResultGroup

`func NewAutocompleteResultGroup() *AutocompleteResultGroup`

NewAutocompleteResultGroup instantiates a new AutocompleteResultGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutocompleteResultGroupWithDefaults

`func NewAutocompleteResultGroupWithDefaults() *AutocompleteResultGroup`

NewAutocompleteResultGroupWithDefaults instantiates a new AutocompleteResultGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartIndex

`func (o *AutocompleteResultGroup) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *AutocompleteResultGroup) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *AutocompleteResultGroup) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *AutocompleteResultGroup) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetEndIndex

`func (o *AutocompleteResultGroup) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *AutocompleteResultGroup) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *AutocompleteResultGroup) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *AutocompleteResultGroup) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetTitle

`func (o *AutocompleteResultGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutocompleteResultGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutocompleteResultGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutocompleteResultGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


