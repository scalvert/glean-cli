# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserID** | Pointer to **string** | An opaque user ID for the claimed authority (i.e., the actas param, or the origid if actas is not specified). | [optional] 
**OrigID** | Pointer to **string** | An opaque user ID for the authenticated user (ignores actas). | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserID

`func (o *User) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *User) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *User) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *User) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetOrigID

`func (o *User) GetOrigID() string`

GetOrigID returns the OrigID field if non-nil, zero value otherwise.

### GetOrigIDOk

`func (o *User) GetOrigIDOk() (*string, bool)`

GetOrigIDOk returns a tuple with the OrigID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigID

`func (o *User) SetOrigID(v string)`

SetOrigID sets OrigID field to given value.

### HasOrigID

`func (o *User) HasOrigID() bool`

HasOrigID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


