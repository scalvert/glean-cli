# PersonObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name. | 
**ObfuscatedId** | **string** | An opaque identifier that can be used to request metadata for a Person. | 

## Methods

### NewPersonObject

`func NewPersonObject(name string, obfuscatedId string, ) *PersonObject`

NewPersonObject instantiates a new PersonObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonObjectWithDefaults

`func NewPersonObjectWithDefaults() *PersonObject`

NewPersonObjectWithDefaults instantiates a new PersonObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PersonObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonObject) SetName(v string)`

SetName sets Name field to given value.


### GetObfuscatedId

`func (o *PersonObject) GetObfuscatedId() string`

GetObfuscatedId returns the ObfuscatedId field if non-nil, zero value otherwise.

### GetObfuscatedIdOk

`func (o *PersonObject) GetObfuscatedIdOk() (*string, bool)`

GetObfuscatedIdOk returns a tuple with the ObfuscatedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedId

`func (o *PersonObject) SetObfuscatedId(v string)`

SetObfuscatedId sets ObfuscatedId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


