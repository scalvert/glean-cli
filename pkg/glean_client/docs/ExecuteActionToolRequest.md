# ExecuteActionToolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the tool. | 
**Parameters** | Pointer to [**map[string]WriteActionParameter**](WriteActionParameter.md) | The parameters to be passed to the tool for action. | [optional] 

## Methods

### NewExecuteActionToolRequest

`func NewExecuteActionToolRequest(name string, ) *ExecuteActionToolRequest`

NewExecuteActionToolRequest instantiates a new ExecuteActionToolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteActionToolRequestWithDefaults

`func NewExecuteActionToolRequestWithDefaults() *ExecuteActionToolRequest`

NewExecuteActionToolRequestWithDefaults instantiates a new ExecuteActionToolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExecuteActionToolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecuteActionToolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecuteActionToolRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *ExecuteActionToolRequest) GetParameters() map[string]WriteActionParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ExecuteActionToolRequest) GetParametersOk() (*map[string]WriteActionParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ExecuteActionToolRequest) SetParameters(v map[string]WriteActionParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ExecuteActionToolRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


