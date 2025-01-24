# GetPinResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | Pointer to [**PinDocument**](PinDocument.md) |  | [optional] 

## Methods

### NewGetPinResponse

`func NewGetPinResponse() *GetPinResponse`

NewGetPinResponse instantiates a new GetPinResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPinResponseWithDefaults

`func NewGetPinResponseWithDefaults() *GetPinResponse`

NewGetPinResponseWithDefaults instantiates a new GetPinResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *GetPinResponse) GetPin() PinDocument`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *GetPinResponse) GetPinOk() (*PinDocument, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *GetPinResponse) SetPin(v PinDocument)`

SetPin sets Pin field to given value.

### HasPin

`func (o *GetPinResponse) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


