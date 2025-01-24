# PersonDistance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name. | 
**ObfuscatedId** | **string** | An opaque identifier that can be used to request metadata for a Person. | 
**Distance** | **float32** | Distance to person, refer to PeopleDistance pipeline on interpretation of the value. | 

## Methods

### NewPersonDistance

`func NewPersonDistance(name string, obfuscatedId string, distance float32, ) *PersonDistance`

NewPersonDistance instantiates a new PersonDistance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonDistanceWithDefaults

`func NewPersonDistanceWithDefaults() *PersonDistance`

NewPersonDistanceWithDefaults instantiates a new PersonDistance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PersonDistance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonDistance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonDistance) SetName(v string)`

SetName sets Name field to given value.


### GetObfuscatedId

`func (o *PersonDistance) GetObfuscatedId() string`

GetObfuscatedId returns the ObfuscatedId field if non-nil, zero value otherwise.

### GetObfuscatedIdOk

`func (o *PersonDistance) GetObfuscatedIdOk() (*string, bool)`

GetObfuscatedIdOk returns a tuple with the ObfuscatedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedId

`func (o *PersonDistance) SetObfuscatedId(v string)`

SetObfuscatedId sets ObfuscatedId field to given value.


### GetDistance

`func (o *PersonDistance) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *PersonDistance) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *PersonDistance) SetDistance(v float32)`

SetDistance sets Distance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


