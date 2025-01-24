# WriteActionParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the value (e.g., integer, string, etc.) | [optional] 
**Value** | Pointer to **string** | The value of the field. | [optional] 
**IsRequired** | Pointer to **bool** | Is the parameter a required field. | [optional] 
**Description** | Pointer to **string** | Description of the parameter. | [optional] 
**PossibleValues** | Pointer to [**[]PossibleValue**](PossibleValue.md) | Possible values that the parameter can take. | [optional] 

## Methods

### NewWriteActionParameter

`func NewWriteActionParameter() *WriteActionParameter`

NewWriteActionParameter instantiates a new WriteActionParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteActionParameterWithDefaults

`func NewWriteActionParameterWithDefaults() *WriteActionParameter`

NewWriteActionParameterWithDefaults instantiates a new WriteActionParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WriteActionParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WriteActionParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WriteActionParameter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WriteActionParameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *WriteActionParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WriteActionParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WriteActionParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WriteActionParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetIsRequired

`func (o *WriteActionParameter) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *WriteActionParameter) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *WriteActionParameter) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *WriteActionParameter) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetDescription

`func (o *WriteActionParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WriteActionParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WriteActionParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WriteActionParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPossibleValues

`func (o *WriteActionParameter) GetPossibleValues() []PossibleValue`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *WriteActionParameter) GetPossibleValuesOk() (*[]PossibleValue, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *WriteActionParameter) SetPossibleValues(v []PossibleValue)`

SetPossibleValues sets PossibleValues field to given value.

### HasPossibleValues

`func (o *WriteActionParameter) HasPossibleValues() bool`

HasPossibleValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


