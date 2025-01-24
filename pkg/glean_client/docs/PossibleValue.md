# PossibleValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Possible value | [optional] 
**Label** | Pointer to **string** | User-friendly label associated with the value | [optional] 

## Methods

### NewPossibleValue

`func NewPossibleValue() *PossibleValue`

NewPossibleValue instantiates a new PossibleValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPossibleValueWithDefaults

`func NewPossibleValueWithDefaults() *PossibleValue`

NewPossibleValueWithDefaults instantiates a new PossibleValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PossibleValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PossibleValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PossibleValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PossibleValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *PossibleValue) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PossibleValue) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PossibleValue) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PossibleValue) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


