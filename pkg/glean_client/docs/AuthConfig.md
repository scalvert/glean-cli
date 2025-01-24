# AuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOnPrem** | Pointer to **bool** | Whether or not this tool is hosted on-premise. | [optional] 
**Type** | Pointer to **string** | The type of authentication being used. Use &#39;OAUTH_*&#39; when Glean calls an external API (e.g., Jira) on behalf of a user to obtain an OAuth token. &#39;OAUTH_ADMIN&#39; utilizes an admin token for external API calls on behalf all users. &#39;OAUTH_USER&#39; uses individual user tokens for external API calls. | [optional] 
**Status** | Pointer to **string** | Auth status of the tool. | [optional] 
**ClientUrl** | Pointer to **string** | The URL where users will be directed to start the OAuth flow. | [optional] 
**Scopes** | Pointer to **[]string** | A list of strings denoting the different scopes or access levels required by the tool. | [optional] 
**AuthorizationUrl** | Pointer to **string** | The OAuth provider&#39;s endpoint, where access tokens are requested. | [optional] 

## Methods

### NewAuthConfig

`func NewAuthConfig() *AuthConfig`

NewAuthConfig instantiates a new AuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthConfigWithDefaults

`func NewAuthConfigWithDefaults() *AuthConfig`

NewAuthConfigWithDefaults instantiates a new AuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOnPrem

`func (o *AuthConfig) GetIsOnPrem() bool`

GetIsOnPrem returns the IsOnPrem field if non-nil, zero value otherwise.

### GetIsOnPremOk

`func (o *AuthConfig) GetIsOnPremOk() (*bool, bool)`

GetIsOnPremOk returns a tuple with the IsOnPrem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnPrem

`func (o *AuthConfig) SetIsOnPrem(v bool)`

SetIsOnPrem sets IsOnPrem field to given value.

### HasIsOnPrem

`func (o *AuthConfig) HasIsOnPrem() bool`

HasIsOnPrem returns a boolean if a field has been set.

### GetType

`func (o *AuthConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *AuthConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClientUrl

`func (o *AuthConfig) GetClientUrl() string`

GetClientUrl returns the ClientUrl field if non-nil, zero value otherwise.

### GetClientUrlOk

`func (o *AuthConfig) GetClientUrlOk() (*string, bool)`

GetClientUrlOk returns a tuple with the ClientUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUrl

`func (o *AuthConfig) SetClientUrl(v string)`

SetClientUrl sets ClientUrl field to given value.

### HasClientUrl

`func (o *AuthConfig) HasClientUrl() bool`

HasClientUrl returns a boolean if a field has been set.

### GetScopes

`func (o *AuthConfig) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthConfig) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthConfig) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthConfig) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetAuthorizationUrl

`func (o *AuthConfig) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *AuthConfig) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *AuthConfig) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.

### HasAuthorizationUrl

`func (o *AuthConfig) HasAuthorizationUrl() bool`

HasAuthorizationUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


