# CustomFieldData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A user-facing label for this field. | 
**Values** | [**[]CustomFieldValue**](CustomFieldValue.md) |  | 
**Displayable** | **bool** | Determines whether the client should display this custom field | [default to true]

## Methods

### NewCustomFieldData

`func NewCustomFieldData(label string, values []CustomFieldValue, displayable bool, ) *CustomFieldData`

NewCustomFieldData instantiates a new CustomFieldData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldDataWithDefaults

`func NewCustomFieldDataWithDefaults() *CustomFieldData`

NewCustomFieldDataWithDefaults instantiates a new CustomFieldData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CustomFieldData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomFieldData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomFieldData) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValues

`func (o *CustomFieldData) GetValues() []CustomFieldValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CustomFieldData) GetValuesOk() (*[]CustomFieldValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CustomFieldData) SetValues(v []CustomFieldValue)`

SetValues sets Values field to given value.


### GetDisplayable

`func (o *CustomFieldData) GetDisplayable() bool`

GetDisplayable returns the Displayable field if non-nil, zero value otherwise.

### GetDisplayableOk

`func (o *CustomFieldData) GetDisplayableOk() (*bool, bool)`

GetDisplayableOk returns a tuple with the Displayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayable

`func (o *CustomFieldData) SetDisplayable(v bool)`

SetDisplayable sets Displayable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


