# CreateAuthTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | An authentication token that can be passed to any endpoint via Bearer Authentication | 
**ExpirationTime** | **int64** | Unix timestamp for when this token expires (in seconds since epoch UTC). | 

## Methods

### NewCreateAuthTokenResponse

`func NewCreateAuthTokenResponse(token string, expirationTime int64, ) *CreateAuthTokenResponse`

NewCreateAuthTokenResponse instantiates a new CreateAuthTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthTokenResponseWithDefaults

`func NewCreateAuthTokenResponseWithDefaults() *CreateAuthTokenResponse`

NewCreateAuthTokenResponseWithDefaults instantiates a new CreateAuthTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CreateAuthTokenResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAuthTokenResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAuthTokenResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpirationTime

`func (o *CreateAuthTokenResponse) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CreateAuthTokenResponse) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CreateAuthTokenResponse) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


