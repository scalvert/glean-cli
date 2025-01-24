# RestrictionFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentSpecs** | Pointer to [**[]DocumentSpec**](DocumentSpec.md) |  | [optional] 
**DatasourceInstances** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRestrictionFilters

`func NewRestrictionFilters() *RestrictionFilters`

NewRestrictionFilters instantiates a new RestrictionFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictionFiltersWithDefaults

`func NewRestrictionFiltersWithDefaults() *RestrictionFilters`

NewRestrictionFiltersWithDefaults instantiates a new RestrictionFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentSpecs

`func (o *RestrictionFilters) GetDocumentSpecs() []DocumentSpec`

GetDocumentSpecs returns the DocumentSpecs field if non-nil, zero value otherwise.

### GetDocumentSpecsOk

`func (o *RestrictionFilters) GetDocumentSpecsOk() (*[]DocumentSpec, bool)`

GetDocumentSpecsOk returns a tuple with the DocumentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSpecs

`func (o *RestrictionFilters) SetDocumentSpecs(v []DocumentSpec)`

SetDocumentSpecs sets DocumentSpecs field to given value.

### HasDocumentSpecs

`func (o *RestrictionFilters) HasDocumentSpecs() bool`

HasDocumentSpecs returns a boolean if a field has been set.

### GetDatasourceInstances

`func (o *RestrictionFilters) GetDatasourceInstances() []string`

GetDatasourceInstances returns the DatasourceInstances field if non-nil, zero value otherwise.

### GetDatasourceInstancesOk

`func (o *RestrictionFilters) GetDatasourceInstancesOk() (*[]string, bool)`

GetDatasourceInstancesOk returns a tuple with the DatasourceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstances

`func (o *RestrictionFilters) SetDatasourceInstances(v []string)`

SetDatasourceInstances sets DatasourceInstances field to given value.

### HasDatasourceInstances

`func (o *RestrictionFilters) HasDatasourceInstances() bool`

HasDatasourceInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


