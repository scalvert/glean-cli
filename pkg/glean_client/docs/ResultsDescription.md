# ResultsDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Textual description of the results. Can be shown at the top of SERP, e.g. &#39;People who write about this topic&#39; for experts in people tab. | [optional] 
**IconConfig** | Pointer to [**IconConfig**](IconConfig.md) |  | [optional] 

## Methods

### NewResultsDescription

`func NewResultsDescription() *ResultsDescription`

NewResultsDescription instantiates a new ResultsDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsDescriptionWithDefaults

`func NewResultsDescriptionWithDefaults() *ResultsDescription`

NewResultsDescriptionWithDefaults instantiates a new ResultsDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ResultsDescription) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ResultsDescription) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ResultsDescription) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ResultsDescription) HasText() bool`

HasText returns a boolean if a field has been set.

### GetIconConfig

`func (o *ResultsDescription) GetIconConfig() IconConfig`

GetIconConfig returns the IconConfig field if non-nil, zero value otherwise.

### GetIconConfigOk

`func (o *ResultsDescription) GetIconConfigOk() (*IconConfig, bool)`

GetIconConfigOk returns a tuple with the IconConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconConfig

`func (o *ResultsDescription) SetIconConfig(v IconConfig)`

SetIconConfig sets IconConfig field to given value.

### HasIconConfig

`func (o *ResultsDescription) HasIconConfig() bool`

HasIconConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


