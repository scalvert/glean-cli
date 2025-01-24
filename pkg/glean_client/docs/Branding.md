# Branding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyBackgroundImageName** | Pointer to **string** | User facing company background image to be displayed to users on the home page. | [optional] 
**CompanyLogoUrl** | Pointer to **string** | An image URL pointing to a custom logo that should be displayed to users. Must be square and recognizable down to 40x40px. SVG images with transparent background are preferred. | [optional] 
**CompanyWideLogoUrl** | Pointer to **string** | An image URL pointing to a wide format custom logo that should be displayed to users. Should be rectangular and recognizable at 40px height, and aspect ratio should be between 2:1 and 4:1. SVG images with transparent background are preferred. | [optional] 
**CompanyLogoBackgroundColor** | Pointer to **string** | A hex RGB color to display behind custom logo (e.g. &#39;#ff4080&#39;). | [optional] 

## Methods

### NewBranding

`func NewBranding() *Branding`

NewBranding instantiates a new Branding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingWithDefaults

`func NewBrandingWithDefaults() *Branding`

NewBrandingWithDefaults instantiates a new Branding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyBackgroundImageName

`func (o *Branding) GetCompanyBackgroundImageName() string`

GetCompanyBackgroundImageName returns the CompanyBackgroundImageName field if non-nil, zero value otherwise.

### GetCompanyBackgroundImageNameOk

`func (o *Branding) GetCompanyBackgroundImageNameOk() (*string, bool)`

GetCompanyBackgroundImageNameOk returns a tuple with the CompanyBackgroundImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyBackgroundImageName

`func (o *Branding) SetCompanyBackgroundImageName(v string)`

SetCompanyBackgroundImageName sets CompanyBackgroundImageName field to given value.

### HasCompanyBackgroundImageName

`func (o *Branding) HasCompanyBackgroundImageName() bool`

HasCompanyBackgroundImageName returns a boolean if a field has been set.

### GetCompanyLogoUrl

`func (o *Branding) GetCompanyLogoUrl() string`

GetCompanyLogoUrl returns the CompanyLogoUrl field if non-nil, zero value otherwise.

### GetCompanyLogoUrlOk

`func (o *Branding) GetCompanyLogoUrlOk() (*string, bool)`

GetCompanyLogoUrlOk returns a tuple with the CompanyLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLogoUrl

`func (o *Branding) SetCompanyLogoUrl(v string)`

SetCompanyLogoUrl sets CompanyLogoUrl field to given value.

### HasCompanyLogoUrl

`func (o *Branding) HasCompanyLogoUrl() bool`

HasCompanyLogoUrl returns a boolean if a field has been set.

### GetCompanyWideLogoUrl

`func (o *Branding) GetCompanyWideLogoUrl() string`

GetCompanyWideLogoUrl returns the CompanyWideLogoUrl field if non-nil, zero value otherwise.

### GetCompanyWideLogoUrlOk

`func (o *Branding) GetCompanyWideLogoUrlOk() (*string, bool)`

GetCompanyWideLogoUrlOk returns a tuple with the CompanyWideLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyWideLogoUrl

`func (o *Branding) SetCompanyWideLogoUrl(v string)`

SetCompanyWideLogoUrl sets CompanyWideLogoUrl field to given value.

### HasCompanyWideLogoUrl

`func (o *Branding) HasCompanyWideLogoUrl() bool`

HasCompanyWideLogoUrl returns a boolean if a field has been set.

### GetCompanyLogoBackgroundColor

`func (o *Branding) GetCompanyLogoBackgroundColor() string`

GetCompanyLogoBackgroundColor returns the CompanyLogoBackgroundColor field if non-nil, zero value otherwise.

### GetCompanyLogoBackgroundColorOk

`func (o *Branding) GetCompanyLogoBackgroundColorOk() (*string, bool)`

GetCompanyLogoBackgroundColorOk returns a tuple with the CompanyLogoBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLogoBackgroundColor

`func (o *Branding) SetCompanyLogoBackgroundColor(v string)`

SetCompanyLogoBackgroundColor sets CompanyLogoBackgroundColor field to given value.

### HasCompanyLogoBackgroundColor

`func (o *Branding) HasCompanyLogoBackgroundColor() bool`

HasCompanyLogoBackgroundColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


