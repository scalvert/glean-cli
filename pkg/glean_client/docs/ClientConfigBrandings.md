# ClientConfigBrandings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Light** | Pointer to [**Branding**](Branding.md) |  | [optional] 
**Dark** | Pointer to [**Branding**](Branding.md) |  | [optional] 
**ProductTerms** | Pointer to [**ProductTerms**](ProductTerms.md) |  | [optional] 

## Methods

### NewClientConfigBrandings

`func NewClientConfigBrandings() *ClientConfigBrandings`

NewClientConfigBrandings instantiates a new ClientConfigBrandings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfigBrandingsWithDefaults

`func NewClientConfigBrandingsWithDefaults() *ClientConfigBrandings`

NewClientConfigBrandingsWithDefaults instantiates a new ClientConfigBrandings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLight

`func (o *ClientConfigBrandings) GetLight() Branding`

GetLight returns the Light field if non-nil, zero value otherwise.

### GetLightOk

`func (o *ClientConfigBrandings) GetLightOk() (*Branding, bool)`

GetLightOk returns a tuple with the Light field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLight

`func (o *ClientConfigBrandings) SetLight(v Branding)`

SetLight sets Light field to given value.

### HasLight

`func (o *ClientConfigBrandings) HasLight() bool`

HasLight returns a boolean if a field has been set.

### GetDark

`func (o *ClientConfigBrandings) GetDark() Branding`

GetDark returns the Dark field if non-nil, zero value otherwise.

### GetDarkOk

`func (o *ClientConfigBrandings) GetDarkOk() (*Branding, bool)`

GetDarkOk returns a tuple with the Dark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDark

`func (o *ClientConfigBrandings) SetDark(v Branding)`

SetDark sets Dark field to given value.

### HasDark

`func (o *ClientConfigBrandings) HasDark() bool`

HasDark returns a boolean if a field has been set.

### GetProductTerms

`func (o *ClientConfigBrandings) GetProductTerms() ProductTerms`

GetProductTerms returns the ProductTerms field if non-nil, zero value otherwise.

### GetProductTermsOk

`func (o *ClientConfigBrandings) GetProductTermsOk() (*ProductTerms, bool)`

GetProductTermsOk returns a tuple with the ProductTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTerms

`func (o *ClientConfigBrandings) SetProductTerms(v ProductTerms)`

SetProductTerms sets ProductTerms field to given value.

### HasProductTerms

`func (o *ClientConfigBrandings) HasProductTerms() bool`

HasProductTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


