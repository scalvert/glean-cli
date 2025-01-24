# ListPinsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pins** | [**[]PinDocument**](PinDocument.md) | List of pinned documents. | 

## Methods

### NewListPinsResponse

`func NewListPinsResponse(pins []PinDocument, ) *ListPinsResponse`

NewListPinsResponse instantiates a new ListPinsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPinsResponseWithDefaults

`func NewListPinsResponseWithDefaults() *ListPinsResponse`

NewListPinsResponseWithDefaults instantiates a new ListPinsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPins

`func (o *ListPinsResponse) GetPins() []PinDocument`

GetPins returns the Pins field if non-nil, zero value otherwise.

### GetPinsOk

`func (o *ListPinsResponse) GetPinsOk() (*[]PinDocument, bool)`

GetPinsOk returns a tuple with the Pins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPins

`func (o *ListPinsResponse) SetPins(v []PinDocument)`

SetPins sets Pins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


