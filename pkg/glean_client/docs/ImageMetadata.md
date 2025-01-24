# ImageMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ImageType**](ImageType.md) |  | 
**Id** | Pointer to **string** | ID, if a specific entity/type is requested. The id may have different meaning for each type. For USER, it is user id For UGC, it is the id of the content For ICON, the doctype. | [optional] 
**Ds** | Pointer to **string** | A specific datasource for which an image is requested for. The ds may have different meaning for each type and can also be empty for some. For USER, it is empty or datasource the icon is asked for. For UGC, it is the UGC datasource. For ICON, it is datasource instance the icon is asked for. | [optional] 
**Cid** | Pointer to **string** | Content id to differentitate multiple images that can have the same type and datasource e.g. thumnail or image from content of UGC. It can also be empty. | [optional] 
**Ext** | Pointer to **string** | Extension the image is saved with. The extension data is deduced from content type for image uploads. | [optional] 

## Methods

### NewImageMetadata

`func NewImageMetadata(type_ ImageType, ) *ImageMetadata`

NewImageMetadata instantiates a new ImageMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageMetadataWithDefaults

`func NewImageMetadataWithDefaults() *ImageMetadata`

NewImageMetadataWithDefaults instantiates a new ImageMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImageMetadata) GetType() ImageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageMetadata) GetTypeOk() (*ImageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageMetadata) SetType(v ImageType)`

SetType sets Type field to given value.


### GetId

`func (o *ImageMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDs

`func (o *ImageMetadata) GetDs() string`

GetDs returns the Ds field if non-nil, zero value otherwise.

### GetDsOk

`func (o *ImageMetadata) GetDsOk() (*string, bool)`

GetDsOk returns a tuple with the Ds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDs

`func (o *ImageMetadata) SetDs(v string)`

SetDs sets Ds field to given value.

### HasDs

`func (o *ImageMetadata) HasDs() bool`

HasDs returns a boolean if a field has been set.

### GetCid

`func (o *ImageMetadata) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *ImageMetadata) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *ImageMetadata) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *ImageMetadata) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetExt

`func (o *ImageMetadata) GetExt() string`

GetExt returns the Ext field if non-nil, zero value otherwise.

### GetExtOk

`func (o *ImageMetadata) GetExtOk() (*string, bool)`

GetExtOk returns a tuple with the Ext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExt

`func (o *ImageMetadata) SetExt(v string)`

SetExt sets Ext field to given value.

### HasExt

`func (o *ImageMetadata) HasExt() bool`

HasExt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


