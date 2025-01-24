# OperatorMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsCustom** | Pointer to **bool** | Whether this operator is supported by default or something that was created within a workplace app (e.g. custom jira field). | [optional] 
**OperatorType** | Pointer to **string** |  | [optional] 
**HelpText** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to [**[]OperatorScope**](OperatorScope.md) |  | [optional] 

## Methods

### NewOperatorMetadata

`func NewOperatorMetadata(name string, ) *OperatorMetadata`

NewOperatorMetadata instantiates a new OperatorMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorMetadataWithDefaults

`func NewOperatorMetadataWithDefaults() *OperatorMetadata`

NewOperatorMetadataWithDefaults instantiates a new OperatorMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OperatorMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperatorMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperatorMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetIsCustom

`func (o *OperatorMetadata) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *OperatorMetadata) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *OperatorMetadata) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.

### HasIsCustom

`func (o *OperatorMetadata) HasIsCustom() bool`

HasIsCustom returns a boolean if a field has been set.

### GetOperatorType

`func (o *OperatorMetadata) GetOperatorType() string`

GetOperatorType returns the OperatorType field if non-nil, zero value otherwise.

### GetOperatorTypeOk

`func (o *OperatorMetadata) GetOperatorTypeOk() (*string, bool)`

GetOperatorTypeOk returns a tuple with the OperatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorType

`func (o *OperatorMetadata) SetOperatorType(v string)`

SetOperatorType sets OperatorType field to given value.

### HasOperatorType

`func (o *OperatorMetadata) HasOperatorType() bool`

HasOperatorType returns a boolean if a field has been set.

### GetHelpText

`func (o *OperatorMetadata) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *OperatorMetadata) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *OperatorMetadata) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *OperatorMetadata) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetScopes

`func (o *OperatorMetadata) GetScopes() []OperatorScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OperatorMetadata) GetScopesOk() (*[]OperatorScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OperatorMetadata) SetScopes(v []OperatorScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OperatorMetadata) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


