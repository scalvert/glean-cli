# CustomDataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayLabel** | Pointer to **string** |  | [optional] 
**StringValue** | Pointer to **string** |  | [optional] 
**StringListValue** | Pointer to **[]string** | list of strings for multi-value properties | [optional] 
**NumberValue** | Pointer to **float32** |  | [optional] 

## Methods

### NewCustomDataValue

`func NewCustomDataValue() *CustomDataValue`

NewCustomDataValue instantiates a new CustomDataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDataValueWithDefaults

`func NewCustomDataValueWithDefaults() *CustomDataValue`

NewCustomDataValueWithDefaults instantiates a new CustomDataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayLabel

`func (o *CustomDataValue) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *CustomDataValue) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *CustomDataValue) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *CustomDataValue) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetStringValue

`func (o *CustomDataValue) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *CustomDataValue) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *CustomDataValue) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *CustomDataValue) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetStringListValue

`func (o *CustomDataValue) GetStringListValue() []string`

GetStringListValue returns the StringListValue field if non-nil, zero value otherwise.

### GetStringListValueOk

`func (o *CustomDataValue) GetStringListValueOk() (*[]string, bool)`

GetStringListValueOk returns a tuple with the StringListValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringListValue

`func (o *CustomDataValue) SetStringListValue(v []string)`

SetStringListValue sets StringListValue field to given value.

### HasStringListValue

`func (o *CustomDataValue) HasStringListValue() bool`

HasStringListValue returns a boolean if a field has been set.

### GetNumberValue

`func (o *CustomDataValue) GetNumberValue() float32`

GetNumberValue returns the NumberValue field if non-nil, zero value otherwise.

### GetNumberValueOk

`func (o *CustomDataValue) GetNumberValueOk() (*float32, bool)`

GetNumberValueOk returns a tuple with the NumberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValue

`func (o *CustomDataValue) SetNumberValue(v float32)`

SetNumberValue sets NumberValue field to given value.

### HasNumberValue

`func (o *CustomDataValue) HasNumberValue() bool`

HasNumberValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


