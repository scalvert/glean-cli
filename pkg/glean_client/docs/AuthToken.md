# AuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**Datasource** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**AuthUser** | Pointer to **string** | Used by Google to indicate the index of the logged in user. Useful for generating hyperlinks that support multilogin. | [optional] 
**Expiration** | Pointer to **int64** | Unix timestamp when this token expires (in seconds since epoch UTC). | [optional] 

## Methods

### NewAuthToken

`func NewAuthToken(accessToken string, datasource string, ) *AuthToken`

NewAuthToken instantiates a new AuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenWithDefaults

`func NewAuthTokenWithDefaults() *AuthToken`

NewAuthTokenWithDefaults instantiates a new AuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AuthToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetDatasource

`func (o *AuthToken) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *AuthToken) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *AuthToken) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetScope

`func (o *AuthToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AuthToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AuthToken) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AuthToken) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTokenType

`func (o *AuthToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AuthToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AuthToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AuthToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetAuthUser

`func (o *AuthToken) GetAuthUser() string`

GetAuthUser returns the AuthUser field if non-nil, zero value otherwise.

### GetAuthUserOk

`func (o *AuthToken) GetAuthUserOk() (*string, bool)`

GetAuthUserOk returns a tuple with the AuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUser

`func (o *AuthToken) SetAuthUser(v string)`

SetAuthUser sets AuthUser field to given value.

### HasAuthUser

`func (o *AuthToken) HasAuthUser() bool`

HasAuthUser returns a boolean if a field has been set.

### GetExpiration

`func (o *AuthToken) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AuthToken) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AuthToken) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AuthToken) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


