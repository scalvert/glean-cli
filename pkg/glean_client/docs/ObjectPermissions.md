# ObjectPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Write** | Pointer to [**WritePermission**](WritePermission.md) |  | [optional] 

## Methods

### NewObjectPermissions

`func NewObjectPermissions() *ObjectPermissions`

NewObjectPermissions instantiates a new ObjectPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPermissionsWithDefaults

`func NewObjectPermissionsWithDefaults() *ObjectPermissions`

NewObjectPermissionsWithDefaults instantiates a new ObjectPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWrite

`func (o *ObjectPermissions) GetWrite() WritePermission`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *ObjectPermissions) GetWriteOk() (*WritePermission, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *ObjectPermissions) SetWrite(v WritePermission)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *ObjectPermissions) HasWrite() bool`

HasWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


