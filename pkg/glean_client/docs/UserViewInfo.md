# UserViewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocId** | Pointer to **string** | Unique Glean Document ID of the associated document. | [optional] 
**DocTitle** | Pointer to **string** | Title of associated document. | [optional] 
**DocUrl** | Pointer to **string** | URL of associated document. | [optional] 

## Methods

### NewUserViewInfo

`func NewUserViewInfo() *UserViewInfo`

NewUserViewInfo instantiates a new UserViewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserViewInfoWithDefaults

`func NewUserViewInfoWithDefaults() *UserViewInfo`

NewUserViewInfoWithDefaults instantiates a new UserViewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocId

`func (o *UserViewInfo) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *UserViewInfo) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *UserViewInfo) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *UserViewInfo) HasDocId() bool`

HasDocId returns a boolean if a field has been set.

### GetDocTitle

`func (o *UserViewInfo) GetDocTitle() string`

GetDocTitle returns the DocTitle field if non-nil, zero value otherwise.

### GetDocTitleOk

`func (o *UserViewInfo) GetDocTitleOk() (*string, bool)`

GetDocTitleOk returns a tuple with the DocTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocTitle

`func (o *UserViewInfo) SetDocTitle(v string)`

SetDocTitle sets DocTitle field to given value.

### HasDocTitle

`func (o *UserViewInfo) HasDocTitle() bool`

HasDocTitle returns a boolean if a field has been set.

### GetDocUrl

`func (o *UserViewInfo) GetDocUrl() string`

GetDocUrl returns the DocUrl field if non-nil, zero value otherwise.

### GetDocUrlOk

`func (o *UserViewInfo) GetDocUrlOk() (*string, bool)`

GetDocUrlOk returns a tuple with the DocUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocUrl

`func (o *UserViewInfo) SetDocUrl(v string)`

SetDocUrl sets DocUrl field to given value.

### HasDocUrl

`func (o *UserViewInfo) HasDocUrl() bool`

HasDocUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


