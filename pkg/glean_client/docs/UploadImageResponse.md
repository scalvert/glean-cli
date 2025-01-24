# UploadImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL of the uploaded image. | 
**Metadata** | Pointer to [**ImageMetadata**](ImageMetadata.md) |  | [optional] 

## Methods

### NewUploadImageResponse

`func NewUploadImageResponse(url string, ) *UploadImageResponse`

NewUploadImageResponse instantiates a new UploadImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadImageResponseWithDefaults

`func NewUploadImageResponseWithDefaults() *UploadImageResponse`

NewUploadImageResponseWithDefaults instantiates a new UploadImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UploadImageResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadImageResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadImageResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMetadata

`func (o *UploadImageResponse) GetMetadata() ImageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UploadImageResponse) GetMetadataOk() (*ImageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UploadImageResponse) SetMetadata(v ImageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UploadImageResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


