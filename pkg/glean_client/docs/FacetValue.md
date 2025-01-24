# FacetValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StringValue** | Pointer to **string** | The value that should be set in the FacetFilter when applying this filter to a search request. | [optional] 
**IntegerValue** | Pointer to **int32** |  | [optional] 
**DisplayLabel** | Pointer to **string** | An optional user-friendly label to display in place of the facet value. | [optional] 
**IconConfig** | Pointer to [**IconConfig**](IconConfig.md) |  | [optional] 

## Methods

### NewFacetValue

`func NewFacetValue() *FacetValue`

NewFacetValue instantiates a new FacetValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetValueWithDefaults

`func NewFacetValueWithDefaults() *FacetValue`

NewFacetValueWithDefaults instantiates a new FacetValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStringValue

`func (o *FacetValue) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *FacetValue) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *FacetValue) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *FacetValue) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetIntegerValue

`func (o *FacetValue) GetIntegerValue() int32`

GetIntegerValue returns the IntegerValue field if non-nil, zero value otherwise.

### GetIntegerValueOk

`func (o *FacetValue) GetIntegerValueOk() (*int32, bool)`

GetIntegerValueOk returns a tuple with the IntegerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegerValue

`func (o *FacetValue) SetIntegerValue(v int32)`

SetIntegerValue sets IntegerValue field to given value.

### HasIntegerValue

`func (o *FacetValue) HasIntegerValue() bool`

HasIntegerValue returns a boolean if a field has been set.

### GetDisplayLabel

`func (o *FacetValue) GetDisplayLabel() string`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *FacetValue) GetDisplayLabelOk() (*string, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *FacetValue) SetDisplayLabel(v string)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *FacetValue) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### GetIconConfig

`func (o *FacetValue) GetIconConfig() IconConfig`

GetIconConfig returns the IconConfig field if non-nil, zero value otherwise.

### GetIconConfigOk

`func (o *FacetValue) GetIconConfigOk() (*IconConfig, bool)`

GetIconConfigOk returns a tuple with the IconConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconConfig

`func (o *FacetValue) SetIconConfig(v IconConfig)`

SetIconConfig sets IconConfig field to given value.

### HasIconConfig

`func (o *FacetValue) HasIconConfig() bool`

HasIconConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


