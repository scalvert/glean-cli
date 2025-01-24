# DocumentSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The title of the document section (e.g. the section header). | [optional] 
**Url** | Pointer to **string** | The permalink of the document section. | [optional] 

## Methods

### NewDocumentSection

`func NewDocumentSection() *DocumentSection`

NewDocumentSection instantiates a new DocumentSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSectionWithDefaults

`func NewDocumentSectionWithDefaults() *DocumentSection`

NewDocumentSectionWithDefaults instantiates a new DocumentSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DocumentSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentSection) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentSection) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *DocumentSection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DocumentSection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DocumentSection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DocumentSection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


