# CustomEntityMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomData** | Pointer to [**map[string]CustomDataValue**](CustomDataValue.md) | Custom fields specific to individual datasources | [optional] 

## Methods

### NewCustomEntityMetadata

`func NewCustomEntityMetadata() *CustomEntityMetadata`

NewCustomEntityMetadata instantiates a new CustomEntityMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntityMetadataWithDefaults

`func NewCustomEntityMetadataWithDefaults() *CustomEntityMetadata`

NewCustomEntityMetadataWithDefaults instantiates a new CustomEntityMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomData

`func (o *CustomEntityMetadata) GetCustomData() map[string]CustomDataValue`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *CustomEntityMetadata) GetCustomDataOk() (*map[string]CustomDataValue, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *CustomEntityMetadata) SetCustomData(v map[string]CustomDataValue)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *CustomEntityMetadata) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


