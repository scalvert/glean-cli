# StructuredLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeskLocation** | Pointer to **string** | Desk number. | [optional] 
**Timezone** | Pointer to **string** | Location&#39;s timezone, e.g. UTC, PST. | [optional] 
**Address** | Pointer to **string** | Office address or name. | [optional] 
**City** | Pointer to **string** | Name of the city. | [optional] 
**State** | Pointer to **string** | State code. | [optional] 
**Region** | Pointer to **string** | Region information, e.g. NORAM, APAC. | [optional] 
**ZipCode** | Pointer to **string** | ZIP Code for the address. | [optional] 
**Country** | Pointer to **string** | Country name. | [optional] 
**CountryCode** | Pointer to **string** | Alpha-2 or Alpha-3 ISO 3166 country code, e.g. US or USA. | [optional] 

## Methods

### NewStructuredLocation

`func NewStructuredLocation() *StructuredLocation`

NewStructuredLocation instantiates a new StructuredLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuredLocationWithDefaults

`func NewStructuredLocationWithDefaults() *StructuredLocation`

NewStructuredLocationWithDefaults instantiates a new StructuredLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeskLocation

`func (o *StructuredLocation) GetDeskLocation() string`

GetDeskLocation returns the DeskLocation field if non-nil, zero value otherwise.

### GetDeskLocationOk

`func (o *StructuredLocation) GetDeskLocationOk() (*string, bool)`

GetDeskLocationOk returns a tuple with the DeskLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeskLocation

`func (o *StructuredLocation) SetDeskLocation(v string)`

SetDeskLocation sets DeskLocation field to given value.

### HasDeskLocation

`func (o *StructuredLocation) HasDeskLocation() bool`

HasDeskLocation returns a boolean if a field has been set.

### GetTimezone

`func (o *StructuredLocation) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *StructuredLocation) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *StructuredLocation) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *StructuredLocation) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAddress

`func (o *StructuredLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StructuredLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StructuredLocation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StructuredLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *StructuredLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *StructuredLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *StructuredLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *StructuredLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *StructuredLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StructuredLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StructuredLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StructuredLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRegion

`func (o *StructuredLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StructuredLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StructuredLocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StructuredLocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetZipCode

`func (o *StructuredLocation) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *StructuredLocation) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *StructuredLocation) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *StructuredLocation) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCountry

`func (o *StructuredLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StructuredLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StructuredLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *StructuredLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *StructuredLocation) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *StructuredLocation) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *StructuredLocation) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *StructuredLocation) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


