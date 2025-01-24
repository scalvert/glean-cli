# ListCollectionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeAudience** | Pointer to **bool** | Whether to include the audience filters with the listed Collections. | [optional] 
**IncludeRoles** | Pointer to **bool** | Whether to include the editor roles with the listed Collections. | [optional] 
**AllowedDatasource** | Pointer to **string** | The datasource type this Collection can hold. ANSWERS - for Collections representing answer boards | [optional] 

## Methods

### NewListCollectionsRequest

`func NewListCollectionsRequest() *ListCollectionsRequest`

NewListCollectionsRequest instantiates a new ListCollectionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCollectionsRequestWithDefaults

`func NewListCollectionsRequestWithDefaults() *ListCollectionsRequest`

NewListCollectionsRequestWithDefaults instantiates a new ListCollectionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeAudience

`func (o *ListCollectionsRequest) GetIncludeAudience() bool`

GetIncludeAudience returns the IncludeAudience field if non-nil, zero value otherwise.

### GetIncludeAudienceOk

`func (o *ListCollectionsRequest) GetIncludeAudienceOk() (*bool, bool)`

GetIncludeAudienceOk returns a tuple with the IncludeAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAudience

`func (o *ListCollectionsRequest) SetIncludeAudience(v bool)`

SetIncludeAudience sets IncludeAudience field to given value.

### HasIncludeAudience

`func (o *ListCollectionsRequest) HasIncludeAudience() bool`

HasIncludeAudience returns a boolean if a field has been set.

### GetIncludeRoles

`func (o *ListCollectionsRequest) GetIncludeRoles() bool`

GetIncludeRoles returns the IncludeRoles field if non-nil, zero value otherwise.

### GetIncludeRolesOk

`func (o *ListCollectionsRequest) GetIncludeRolesOk() (*bool, bool)`

GetIncludeRolesOk returns a tuple with the IncludeRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRoles

`func (o *ListCollectionsRequest) SetIncludeRoles(v bool)`

SetIncludeRoles sets IncludeRoles field to given value.

### HasIncludeRoles

`func (o *ListCollectionsRequest) HasIncludeRoles() bool`

HasIncludeRoles returns a boolean if a field has been set.

### GetAllowedDatasource

`func (o *ListCollectionsRequest) GetAllowedDatasource() string`

GetAllowedDatasource returns the AllowedDatasource field if non-nil, zero value otherwise.

### GetAllowedDatasourceOk

`func (o *ListCollectionsRequest) GetAllowedDatasourceOk() (*string, bool)`

GetAllowedDatasourceOk returns a tuple with the AllowedDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDatasource

`func (o *ListCollectionsRequest) SetAllowedDatasource(v string)`

SetAllowedDatasource sets AllowedDatasource field to given value.

### HasAllowedDatasource

`func (o *ListCollectionsRequest) HasAllowedDatasource() bool`

HasAllowedDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


