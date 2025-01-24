# GetCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**RootCollection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular Collection. To be used for &#x60;/feedback&#x60; reporting. | [optional] 
**Error** | Pointer to [**CollectionError**](CollectionError.md) |  | [optional] 

## Methods

### NewGetCollectionResponse

`func NewGetCollectionResponse() *GetCollectionResponse`

NewGetCollectionResponse instantiates a new GetCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionResponseWithDefaults

`func NewGetCollectionResponseWithDefaults() *GetCollectionResponse`

NewGetCollectionResponseWithDefaults instantiates a new GetCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *GetCollectionResponse) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *GetCollectionResponse) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *GetCollectionResponse) SetCollection(v Collection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *GetCollectionResponse) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetRootCollection

`func (o *GetCollectionResponse) GetRootCollection() Collection`

GetRootCollection returns the RootCollection field if non-nil, zero value otherwise.

### GetRootCollectionOk

`func (o *GetCollectionResponse) GetRootCollectionOk() (*Collection, bool)`

GetRootCollectionOk returns a tuple with the RootCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCollection

`func (o *GetCollectionResponse) SetRootCollection(v Collection)`

SetRootCollection sets RootCollection field to given value.

### HasRootCollection

`func (o *GetCollectionResponse) HasRootCollection() bool`

HasRootCollection returns a boolean if a field has been set.

### GetTrackingToken

`func (o *GetCollectionResponse) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *GetCollectionResponse) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *GetCollectionResponse) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *GetCollectionResponse) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetError

`func (o *GetCollectionResponse) GetError() CollectionError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetCollectionResponse) GetErrorOk() (*CollectionError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetCollectionResponse) SetError(v CollectionError)`

SetError sets Error field to given value.

### HasError

`func (o *GetCollectionResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


