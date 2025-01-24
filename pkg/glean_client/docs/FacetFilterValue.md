# FacetFilterValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**RelationType** | Pointer to **string** |  | [optional] 
**IsNegated** | Pointer to **bool** | DEPRECATED - please use relationType instead | [optional] 

## Methods

### NewFacetFilterValue

`func NewFacetFilterValue() *FacetFilterValue`

NewFacetFilterValue instantiates a new FacetFilterValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetFilterValueWithDefaults

`func NewFacetFilterValueWithDefaults() *FacetFilterValue`

NewFacetFilterValueWithDefaults instantiates a new FacetFilterValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *FacetFilterValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FacetFilterValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FacetFilterValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FacetFilterValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRelationType

`func (o *FacetFilterValue) GetRelationType() string`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *FacetFilterValue) GetRelationTypeOk() (*string, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *FacetFilterValue) SetRelationType(v string)`

SetRelationType sets RelationType field to given value.

### HasRelationType

`func (o *FacetFilterValue) HasRelationType() bool`

HasRelationType returns a boolean if a field has been set.

### GetIsNegated

`func (o *FacetFilterValue) GetIsNegated() bool`

GetIsNegated returns the IsNegated field if non-nil, zero value otherwise.

### GetIsNegatedOk

`func (o *FacetFilterValue) GetIsNegatedOk() (*bool, bool)`

GetIsNegatedOk returns a tuple with the IsNegated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNegated

`func (o *FacetFilterValue) SetIsNegated(v bool)`

SetIsNegated sets IsNegated field to given value.

### HasIsNegated

`func (o *FacetFilterValue) HasIsNegated() bool`

HasIsNegated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


