# CustomFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StrText** | Pointer to **string** | Text field for string value. | [optional] 
**UrlAnchor** | Pointer to **string** | Anchor text for hyperlink. | [optional] 
**UrlLink** | Pointer to **string** | Link for this URL. | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 

## Methods

### NewCustomFieldValue

`func NewCustomFieldValue() *CustomFieldValue`

NewCustomFieldValue instantiates a new CustomFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldValueWithDefaults

`func NewCustomFieldValueWithDefaults() *CustomFieldValue`

NewCustomFieldValueWithDefaults instantiates a new CustomFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrText

`func (o *CustomFieldValue) GetStrText() string`

GetStrText returns the StrText field if non-nil, zero value otherwise.

### GetStrTextOk

`func (o *CustomFieldValue) GetStrTextOk() (*string, bool)`

GetStrTextOk returns a tuple with the StrText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrText

`func (o *CustomFieldValue) SetStrText(v string)`

SetStrText sets StrText field to given value.

### HasStrText

`func (o *CustomFieldValue) HasStrText() bool`

HasStrText returns a boolean if a field has been set.

### GetUrlAnchor

`func (o *CustomFieldValue) GetUrlAnchor() string`

GetUrlAnchor returns the UrlAnchor field if non-nil, zero value otherwise.

### GetUrlAnchorOk

`func (o *CustomFieldValue) GetUrlAnchorOk() (*string, bool)`

GetUrlAnchorOk returns a tuple with the UrlAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlAnchor

`func (o *CustomFieldValue) SetUrlAnchor(v string)`

SetUrlAnchor sets UrlAnchor field to given value.

### HasUrlAnchor

`func (o *CustomFieldValue) HasUrlAnchor() bool`

HasUrlAnchor returns a boolean if a field has been set.

### GetUrlLink

`func (o *CustomFieldValue) GetUrlLink() string`

GetUrlLink returns the UrlLink field if non-nil, zero value otherwise.

### GetUrlLinkOk

`func (o *CustomFieldValue) GetUrlLinkOk() (*string, bool)`

GetUrlLinkOk returns a tuple with the UrlLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlLink

`func (o *CustomFieldValue) SetUrlLink(v string)`

SetUrlLink sets UrlLink field to given value.

### HasUrlLink

`func (o *CustomFieldValue) HasUrlLink() bool`

HasUrlLink returns a boolean if a field has been set.

### GetPerson

`func (o *CustomFieldValue) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *CustomFieldValue) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *CustomFieldValue) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *CustomFieldValue) HasPerson() bool`

HasPerson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


