# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User-friendly display name. | 
**ProfileUrl** | Pointer to **string** | Link to internal company company profile. | [optional] 
**WebsiteUrls** | Pointer to **[]string** | Link to company&#39;s associated websites. | [optional] 
**LogoUrl** | Pointer to **string** | The URL of the company&#39;s logo. Public, Glean-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs). | [optional] 
**Location** | Pointer to **string** | User facing string representing the company&#39;s location. | [optional] 
**Phone** | Pointer to **string** | Phone number as a number string. | [optional] 
**Fax** | Pointer to **string** | Fax number as a number string. | [optional] 
**Industry** | Pointer to **string** | User facing string representing the company&#39;s industry. | [optional] 
**AnnualRevenue** | Pointer to **float64** | Average company&#39;s annual revenue for reference. | [optional] 
**NumberOfEmployees** | Pointer to **int64** | Average company&#39;s number of employees for reference. | [optional] 
**StockSymbol** | Pointer to **string** | Company&#39;s stock symbol if company is public. | [optional] 
**FoundedDate** | Pointer to **string** | The date when the company was founded. | [optional] 
**About** | Pointer to **string** | User facing description of company. | [optional] 

## Methods

### NewCompany

`func NewCompany(name string, ) *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.


### GetProfileUrl

`func (o *Company) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *Company) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *Company) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *Company) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### GetWebsiteUrls

`func (o *Company) GetWebsiteUrls() []string`

GetWebsiteUrls returns the WebsiteUrls field if non-nil, zero value otherwise.

### GetWebsiteUrlsOk

`func (o *Company) GetWebsiteUrlsOk() (*[]string, bool)`

GetWebsiteUrlsOk returns a tuple with the WebsiteUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrls

`func (o *Company) SetWebsiteUrls(v []string)`

SetWebsiteUrls sets WebsiteUrls field to given value.

### HasWebsiteUrls

`func (o *Company) HasWebsiteUrls() bool`

HasWebsiteUrls returns a boolean if a field has been set.

### GetLogoUrl

`func (o *Company) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Company) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Company) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Company) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetLocation

`func (o *Company) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Company) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Company) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Company) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPhone

`func (o *Company) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Company) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Company) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Company) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFax

`func (o *Company) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Company) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Company) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Company) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetIndustry

`func (o *Company) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *Company) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *Company) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *Company) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetAnnualRevenue

`func (o *Company) GetAnnualRevenue() float64`

GetAnnualRevenue returns the AnnualRevenue field if non-nil, zero value otherwise.

### GetAnnualRevenueOk

`func (o *Company) GetAnnualRevenueOk() (*float64, bool)`

GetAnnualRevenueOk returns a tuple with the AnnualRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualRevenue

`func (o *Company) SetAnnualRevenue(v float64)`

SetAnnualRevenue sets AnnualRevenue field to given value.

### HasAnnualRevenue

`func (o *Company) HasAnnualRevenue() bool`

HasAnnualRevenue returns a boolean if a field has been set.

### GetNumberOfEmployees

`func (o *Company) GetNumberOfEmployees() int64`

GetNumberOfEmployees returns the NumberOfEmployees field if non-nil, zero value otherwise.

### GetNumberOfEmployeesOk

`func (o *Company) GetNumberOfEmployeesOk() (*int64, bool)`

GetNumberOfEmployeesOk returns a tuple with the NumberOfEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEmployees

`func (o *Company) SetNumberOfEmployees(v int64)`

SetNumberOfEmployees sets NumberOfEmployees field to given value.

### HasNumberOfEmployees

`func (o *Company) HasNumberOfEmployees() bool`

HasNumberOfEmployees returns a boolean if a field has been set.

### GetStockSymbol

`func (o *Company) GetStockSymbol() string`

GetStockSymbol returns the StockSymbol field if non-nil, zero value otherwise.

### GetStockSymbolOk

`func (o *Company) GetStockSymbolOk() (*string, bool)`

GetStockSymbolOk returns a tuple with the StockSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockSymbol

`func (o *Company) SetStockSymbol(v string)`

SetStockSymbol sets StockSymbol field to given value.

### HasStockSymbol

`func (o *Company) HasStockSymbol() bool`

HasStockSymbol returns a boolean if a field has been set.

### GetFoundedDate

`func (o *Company) GetFoundedDate() string`

GetFoundedDate returns the FoundedDate field if non-nil, zero value otherwise.

### GetFoundedDateOk

`func (o *Company) GetFoundedDateOk() (*string, bool)`

GetFoundedDateOk returns a tuple with the FoundedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundedDate

`func (o *Company) SetFoundedDate(v string)`

SetFoundedDate sets FoundedDate field to given value.

### HasFoundedDate

`func (o *Company) HasFoundedDate() bool`

HasFoundedDate returns a boolean if a field has been set.

### GetAbout

`func (o *Company) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *Company) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *Company) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *Company) HasAbout() bool`

HasAbout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


