# CustomEntitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]CustomEntity**](CustomEntity.md) | A custom entity object | [optional] 
**Errors** | Pointer to **[]string** | A list of IDs that could not be found. | [optional] 

## Methods

### NewCustomEntitiesResponse

`func NewCustomEntitiesResponse() *CustomEntitiesResponse`

NewCustomEntitiesResponse instantiates a new CustomEntitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEntitiesResponseWithDefaults

`func NewCustomEntitiesResponseWithDefaults() *CustomEntitiesResponse`

NewCustomEntitiesResponseWithDefaults instantiates a new CustomEntitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CustomEntitiesResponse) GetResults() []CustomEntity`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CustomEntitiesResponse) GetResultsOk() (*[]CustomEntity, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CustomEntitiesResponse) SetResults(v []CustomEntity)`

SetResults sets Results field to given value.

### HasResults

`func (o *CustomEntitiesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetErrors

`func (o *CustomEntitiesResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CustomEntitiesResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CustomEntitiesResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CustomEntitiesResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


