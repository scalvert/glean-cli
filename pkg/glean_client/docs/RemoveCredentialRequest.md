# RemoveCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | Pointer to **string** | the datasource the credential applies to | [optional] 
**DatasourceInstance** | Pointer to **string** | the datasource instance the credential applies to | [optional] 
**User** | Pointer to **string** | the user info (email or username for example) for the credential | [optional] 

## Methods

### NewRemoveCredentialRequest

`func NewRemoveCredentialRequest() *RemoveCredentialRequest`

NewRemoveCredentialRequest instantiates a new RemoveCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveCredentialRequestWithDefaults

`func NewRemoveCredentialRequestWithDefaults() *RemoveCredentialRequest`

NewRemoveCredentialRequestWithDefaults instantiates a new RemoveCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *RemoveCredentialRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *RemoveCredentialRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *RemoveCredentialRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *RemoveCredentialRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetDatasourceInstance

`func (o *RemoveCredentialRequest) GetDatasourceInstance() string`

GetDatasourceInstance returns the DatasourceInstance field if non-nil, zero value otherwise.

### GetDatasourceInstanceOk

`func (o *RemoveCredentialRequest) GetDatasourceInstanceOk() (*string, bool)`

GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstance

`func (o *RemoveCredentialRequest) SetDatasourceInstance(v string)`

SetDatasourceInstance sets DatasourceInstance field to given value.

### HasDatasourceInstance

`func (o *RemoveCredentialRequest) HasDatasourceInstance() bool`

HasDatasourceInstance returns a boolean if a field has been set.

### GetUser

`func (o *RemoveCredentialRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RemoveCredentialRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RemoveCredentialRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RemoveCredentialRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


