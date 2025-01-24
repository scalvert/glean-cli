# DisplayableListConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | defines how to render this particular displayable list card | [optional] 
**Title** | Pointer to **string** | Primary title for the list. | [optional] 
**Enabled** | Pointer to **bool** | Whether the list should be shown to the user | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should should see displayable list | [optional] 
**ItemUIConfig** | Pointer to [**DisplayableListItemUIConfig**](DisplayableListItemUIConfig.md) |  | [optional] 

## Methods

### NewDisplayableListConfig

`func NewDisplayableListConfig() *DisplayableListConfig`

NewDisplayableListConfig instantiates a new DisplayableListConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayableListConfigWithDefaults

`func NewDisplayableListConfigWithDefaults() *DisplayableListConfig`

NewDisplayableListConfigWithDefaults instantiates a new DisplayableListConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *DisplayableListConfig) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DisplayableListConfig) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DisplayableListConfig) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DisplayableListConfig) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTitle

`func (o *DisplayableListConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DisplayableListConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DisplayableListConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DisplayableListConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEnabled

`func (o *DisplayableListConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DisplayableListConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DisplayableListConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DisplayableListConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *DisplayableListConfig) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *DisplayableListConfig) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *DisplayableListConfig) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *DisplayableListConfig) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetItemUIConfig

`func (o *DisplayableListConfig) GetItemUIConfig() DisplayableListItemUIConfig`

GetItemUIConfig returns the ItemUIConfig field if non-nil, zero value otherwise.

### GetItemUIConfigOk

`func (o *DisplayableListConfig) GetItemUIConfigOk() (*DisplayableListItemUIConfig, bool)`

GetItemUIConfigOk returns a tuple with the ItemUIConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUIConfig

`func (o *DisplayableListConfig) SetItemUIConfig(v DisplayableListItemUIConfig)`

SetItemUIConfig sets ItemUIConfig field to given value.

### HasItemUIConfig

`func (o *DisplayableListConfig) HasItemUIConfig() bool`

HasItemUIConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


