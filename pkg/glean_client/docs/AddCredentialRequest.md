# AddCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | Pointer to **string** | the datasource the credential applies to | [optional] 
**DatasourceInstance** | Pointer to **string** | the datasource instance the credential applies to | [optional] 
**User** | Pointer to **string** | the user info (email or username for example) for the credential | [optional] 
**Token** | Pointer to **string** | the token part of the credential (password, apiToken etc) | [optional] 
**Metadata** | Pointer to **string** | any metadata associated with the user credential | [optional] 

## Methods

### NewAddCredentialRequest

`func NewAddCredentialRequest() *AddCredentialRequest`

NewAddCredentialRequest instantiates a new AddCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCredentialRequestWithDefaults

`func NewAddCredentialRequestWithDefaults() *AddCredentialRequest`

NewAddCredentialRequestWithDefaults instantiates a new AddCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *AddCredentialRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *AddCredentialRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *AddCredentialRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *AddCredentialRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetDatasourceInstance

`func (o *AddCredentialRequest) GetDatasourceInstance() string`

GetDatasourceInstance returns the DatasourceInstance field if non-nil, zero value otherwise.

### GetDatasourceInstanceOk

`func (o *AddCredentialRequest) GetDatasourceInstanceOk() (*string, bool)`

GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstance

`func (o *AddCredentialRequest) SetDatasourceInstance(v string)`

SetDatasourceInstance sets DatasourceInstance field to given value.

### HasDatasourceInstance

`func (o *AddCredentialRequest) HasDatasourceInstance() bool`

HasDatasourceInstance returns a boolean if a field has been set.

### GetUser

`func (o *AddCredentialRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AddCredentialRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AddCredentialRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AddCredentialRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetToken

`func (o *AddCredentialRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddCredentialRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddCredentialRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AddCredentialRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetMetadata

`func (o *AddCredentialRequest) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddCredentialRequest) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddCredentialRequest) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddCredentialRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


