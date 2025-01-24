# EventClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**EventClassificationName**](EventClassificationName.md) |  | [optional] 
**Strategies** | Pointer to [**[]EventStrategyName**](EventStrategyName.md) |  | [optional] 

## Methods

### NewEventClassification

`func NewEventClassification() *EventClassification`

NewEventClassification instantiates a new EventClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventClassificationWithDefaults

`func NewEventClassificationWithDefaults() *EventClassification`

NewEventClassificationWithDefaults instantiates a new EventClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventClassification) GetName() EventClassificationName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventClassification) GetNameOk() (*EventClassificationName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventClassification) SetName(v EventClassificationName)`

SetName sets Name field to given value.

### HasName

`func (o *EventClassification) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStrategies

`func (o *EventClassification) GetStrategies() []EventStrategyName`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *EventClassification) GetStrategiesOk() (*[]EventStrategyName, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *EventClassification) SetStrategies(v []EventStrategyName)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *EventClassification) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


