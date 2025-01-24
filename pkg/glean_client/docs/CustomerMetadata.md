# CustomerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasourceId** | Pointer to **string** | The user visible id of the salesforce customer account. | [optional] 
**CustomData** | Pointer to [**map[string]CustomDataValue**](CustomDataValue.md) | Custom fields specific to individual datasources | [optional] 

## Methods

### NewCustomerMetadata

`func NewCustomerMetadata() *CustomerMetadata`

NewCustomerMetadata instantiates a new CustomerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerMetadataWithDefaults

`func NewCustomerMetadataWithDefaults() *CustomerMetadata`

NewCustomerMetadataWithDefaults instantiates a new CustomerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasourceId

`func (o *CustomerMetadata) GetDatasourceId() string`

GetDatasourceId returns the DatasourceId field if non-nil, zero value otherwise.

### GetDatasourceIdOk

`func (o *CustomerMetadata) GetDatasourceIdOk() (*string, bool)`

GetDatasourceIdOk returns a tuple with the DatasourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceId

`func (o *CustomerMetadata) SetDatasourceId(v string)`

SetDatasourceId sets DatasourceId field to given value.

### HasDatasourceId

`func (o *CustomerMetadata) HasDatasourceId() bool`

HasDatasourceId returns a boolean if a field has been set.

### GetCustomData

`func (o *CustomerMetadata) GetCustomData() map[string]CustomDataValue`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *CustomerMetadata) GetCustomDataOk() (*map[string]CustomDataValue, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *CustomerMetadata) SetCustomData(v map[string]CustomDataValue)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *CustomerMetadata) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


