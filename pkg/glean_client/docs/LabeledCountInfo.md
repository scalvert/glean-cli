# LabeledCountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label for the included count information. | 
**CountInfo** | Pointer to [**[]CountInfo**](CountInfo.md) | List of data points for counts for a given date period. | [optional] 

## Methods

### NewLabeledCountInfo

`func NewLabeledCountInfo(label string, ) *LabeledCountInfo`

NewLabeledCountInfo instantiates a new LabeledCountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabeledCountInfoWithDefaults

`func NewLabeledCountInfoWithDefaults() *LabeledCountInfo`

NewLabeledCountInfoWithDefaults instantiates a new LabeledCountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LabeledCountInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LabeledCountInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LabeledCountInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCountInfo

`func (o *LabeledCountInfo) GetCountInfo() []CountInfo`

GetCountInfo returns the CountInfo field if non-nil, zero value otherwise.

### GetCountInfoOk

`func (o *LabeledCountInfo) GetCountInfoOk() (*[]CountInfo, bool)`

GetCountInfoOk returns a tuple with the CountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountInfo

`func (o *LabeledCountInfo) SetCountInfo(v []CountInfo)`

SetCountInfo sets CountInfo field to given value.

### HasCountInfo

`func (o *LabeledCountInfo) HasCountInfo() bool`

HasCountInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


