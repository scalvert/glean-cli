# ListCollectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | [**[]Collection**](Collection.md) | List of all Collections, no Collection items are fetched. | 

## Methods

### NewListCollectionsResponse

`func NewListCollectionsResponse(collections []Collection, ) *ListCollectionsResponse`

NewListCollectionsResponse instantiates a new ListCollectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCollectionsResponseWithDefaults

`func NewListCollectionsResponseWithDefaults() *ListCollectionsResponse`

NewListCollectionsResponseWithDefaults instantiates a new ListCollectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *ListCollectionsResponse) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ListCollectionsResponse) GetCollectionsOk() (*[]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ListCollectionsResponse) SetCollections(v []Collection)`

SetCollections sets Collections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


