# ClientAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the client action. | 
**Quicklink** | Pointer to [**Quicklink**](Quicklink.md) |  | [optional] 
**DestinationUrl** | Pointer to **string** | Specific URL if action requires a destination URL to complete. Has precedence over action context. | [optional] 

## Methods

### NewClientAction

`func NewClientAction(type_ string, ) *ClientAction`

NewClientAction instantiates a new ClientAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientActionWithDefaults

`func NewClientActionWithDefaults() *ClientAction`

NewClientActionWithDefaults instantiates a new ClientAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClientAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientAction) SetType(v string)`

SetType sets Type field to given value.


### GetQuicklink

`func (o *ClientAction) GetQuicklink() Quicklink`

GetQuicklink returns the Quicklink field if non-nil, zero value otherwise.

### GetQuicklinkOk

`func (o *ClientAction) GetQuicklinkOk() (*Quicklink, bool)`

GetQuicklinkOk returns a tuple with the Quicklink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuicklink

`func (o *ClientAction) SetQuicklink(v Quicklink)`

SetQuicklink sets Quicklink field to given value.

### HasQuicklink

`func (o *ClientAction) HasQuicklink() bool`

HasQuicklink returns a boolean if a field has been set.

### GetDestinationUrl

`func (o *ClientAction) GetDestinationUrl() string`

GetDestinationUrl returns the DestinationUrl field if non-nil, zero value otherwise.

### GetDestinationUrlOk

`func (o *ClientAction) GetDestinationUrlOk() (*string, bool)`

GetDestinationUrlOk returns a tuple with the DestinationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUrl

`func (o *ClientAction) SetDestinationUrl(v string)`

SetDestinationUrl sets DestinationUrl field to given value.

### HasDestinationUrl

`func (o *ClientAction) HasDestinationUrl() bool`

HasDestinationUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


