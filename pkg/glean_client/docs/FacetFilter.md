# FacetFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**[]FacetFilterValue**](FacetFilterValue.md) | Within a single FacetFilter, the values are to be treated like an OR. For example, fieldName type with values [EQUALS Presentation, EQUALS Spreadsheet] means we want to show a document if it&#39;s a Presentation OR a Spreadsheet. | [optional] 
**GroupName** | Pointer to **string** | Indicates the value of a facet, if any, that the given facet is grouped under. This is only used for nested facets, for example, fieldName could be owner and groupName would be Spreadsheet if showing all owners for spreadsheets as a nested facet. | [optional] 

## Methods

### NewFacetFilter

`func NewFacetFilter() *FacetFilter`

NewFacetFilter instantiates a new FacetFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetFilterWithDefaults

`func NewFacetFilterWithDefaults() *FacetFilter`

NewFacetFilterWithDefaults instantiates a new FacetFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *FacetFilter) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *FacetFilter) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *FacetFilter) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *FacetFilter) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetValues

`func (o *FacetFilter) GetValues() []FacetFilterValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FacetFilter) GetValuesOk() (*[]FacetFilterValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FacetFilter) SetValues(v []FacetFilterValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *FacetFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetGroupName

`func (o *FacetFilter) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *FacetFilter) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *FacetFilter) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *FacetFilter) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


