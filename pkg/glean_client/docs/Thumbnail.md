# Thumbnail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhotoId** | Pointer to **string** | Photo id if the thumbnail is from splash. | [optional] 
**Url** | Pointer to **string** | Thumbnail URL. This can be user provided image and/or from downloaded images hosted by Glean. | [optional] 

## Methods

### NewThumbnail

`func NewThumbnail() *Thumbnail`

NewThumbnail instantiates a new Thumbnail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbnailWithDefaults

`func NewThumbnailWithDefaults() *Thumbnail`

NewThumbnailWithDefaults instantiates a new Thumbnail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhotoId

`func (o *Thumbnail) GetPhotoId() string`

GetPhotoId returns the PhotoId field if non-nil, zero value otherwise.

### GetPhotoIdOk

`func (o *Thumbnail) GetPhotoIdOk() (*string, bool)`

GetPhotoIdOk returns a tuple with the PhotoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoId

`func (o *Thumbnail) SetPhotoId(v string)`

SetPhotoId sets PhotoId field to given value.

### HasPhotoId

`func (o *Thumbnail) HasPhotoId() bool`

HasPhotoId returns a boolean if a field has been set.

### GetUrl

`func (o *Thumbnail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Thumbnail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Thumbnail) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Thumbnail) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


