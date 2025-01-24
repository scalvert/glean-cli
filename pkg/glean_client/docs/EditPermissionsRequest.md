# EditPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to **[]string** | The ids of the users whose permissions will be edited | [optional] 
**Permissions** | [**Permissions**](Permissions.md) |  | 

## Methods

### NewEditPermissionsRequest

`func NewEditPermissionsRequest(permissions Permissions, ) *EditPermissionsRequest`

NewEditPermissionsRequest instantiates a new EditPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditPermissionsRequestWithDefaults

`func NewEditPermissionsRequestWithDefaults() *EditPermissionsRequest`

NewEditPermissionsRequestWithDefaults instantiates a new EditPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *EditPermissionsRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *EditPermissionsRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *EditPermissionsRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *EditPermissionsRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetPermissions

`func (o *EditPermissionsRequest) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EditPermissionsRequest) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EditPermissionsRequest) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


