# RelatedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the related object | 
**Metadata** | Pointer to [**RelatedObjectMetadata**](RelatedObjectMetadata.md) |  | [optional] 

## Methods

### NewRelatedObject

`func NewRelatedObject(id string, ) *RelatedObject`

NewRelatedObject instantiates a new RelatedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedObjectWithDefaults

`func NewRelatedObjectWithDefaults() *RelatedObject`

NewRelatedObjectWithDefaults instantiates a new RelatedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelatedObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedObject) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *RelatedObject) GetMetadata() RelatedObjectMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RelatedObject) GetMetadataOk() (*RelatedObjectMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RelatedObject) SetMetadata(v RelatedObjectMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RelatedObject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


