# ClientData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastRefreshTimestamp** | Pointer to **int32** | Server unix timestamp of the last refresh/request (in seconds since epoch UTC). | [optional] 

## Methods

### NewClientData

`func NewClientData() *ClientData`

NewClientData instantiates a new ClientData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDataWithDefaults

`func NewClientDataWithDefaults() *ClientData`

NewClientDataWithDefaults instantiates a new ClientData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastRefreshTimestamp

`func (o *ClientData) GetLastRefreshTimestamp() int32`

GetLastRefreshTimestamp returns the LastRefreshTimestamp field if non-nil, zero value otherwise.

### GetLastRefreshTimestampOk

`func (o *ClientData) GetLastRefreshTimestampOk() (*int32, bool)`

GetLastRefreshTimestampOk returns a tuple with the LastRefreshTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshTimestamp

`func (o *ClientData) SetLastRefreshTimestamp(v int32)`

SetLastRefreshTimestamp sets LastRefreshTimestamp field to given value.

### HasLastRefreshTimestamp

`func (o *ClientData) HasLastRefreshTimestamp() bool`

HasLastRefreshTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


