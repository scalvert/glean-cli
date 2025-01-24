# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier. | 
**Domains** | Pointer to **[]string** | Link to company&#39;s associated website domains. | [optional] 
**Company** | [**Company**](Company.md) |  | 
**DocumentCounts** | Pointer to **map[string]int32** | A map of {string, int} pairs representing counts of each document type associated with this customer. | [optional] 
**Poc** | Pointer to [**[]Person**](Person.md) | A list of POC for company. | [optional] 
**Metadata** | Pointer to [**CustomerMetadata**](CustomerMetadata.md) |  | [optional] 
**MergedCustomers** | Pointer to [**[]Customer**](Customer.md) | A list of Customers. | [optional] 
**StartDate** | Pointer to **string** | The date when the interaction with customer started. | [optional] 
**ContractAnnualRevenue** | Pointer to **float64** | Average contract annual revenue with that customer. | [optional] 
**Notes** | Pointer to **string** | User facing (potentially generated) notes about company. | [optional] 

## Methods

### NewCustomer

`func NewCustomer(id string, company Company, ) *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.


### GetDomains

`func (o *Customer) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Customer) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Customer) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Customer) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetCompany

`func (o *Customer) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Customer) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Customer) SetCompany(v Company)`

SetCompany sets Company field to given value.


### GetDocumentCounts

`func (o *Customer) GetDocumentCounts() map[string]int32`

GetDocumentCounts returns the DocumentCounts field if non-nil, zero value otherwise.

### GetDocumentCountsOk

`func (o *Customer) GetDocumentCountsOk() (*map[string]int32, bool)`

GetDocumentCountsOk returns a tuple with the DocumentCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCounts

`func (o *Customer) SetDocumentCounts(v map[string]int32)`

SetDocumentCounts sets DocumentCounts field to given value.

### HasDocumentCounts

`func (o *Customer) HasDocumentCounts() bool`

HasDocumentCounts returns a boolean if a field has been set.

### GetPoc

`func (o *Customer) GetPoc() []Person`

GetPoc returns the Poc field if non-nil, zero value otherwise.

### GetPocOk

`func (o *Customer) GetPocOk() (*[]Person, bool)`

GetPocOk returns a tuple with the Poc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoc

`func (o *Customer) SetPoc(v []Person)`

SetPoc sets Poc field to given value.

### HasPoc

`func (o *Customer) HasPoc() bool`

HasPoc returns a boolean if a field has been set.

### GetMetadata

`func (o *Customer) GetMetadata() CustomerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Customer) GetMetadataOk() (*CustomerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Customer) SetMetadata(v CustomerMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Customer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMergedCustomers

`func (o *Customer) GetMergedCustomers() []Customer`

GetMergedCustomers returns the MergedCustomers field if non-nil, zero value otherwise.

### GetMergedCustomersOk

`func (o *Customer) GetMergedCustomersOk() (*[]Customer, bool)`

GetMergedCustomersOk returns a tuple with the MergedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedCustomers

`func (o *Customer) SetMergedCustomers(v []Customer)`

SetMergedCustomers sets MergedCustomers field to given value.

### HasMergedCustomers

`func (o *Customer) HasMergedCustomers() bool`

HasMergedCustomers returns a boolean if a field has been set.

### GetStartDate

`func (o *Customer) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Customer) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Customer) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Customer) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetContractAnnualRevenue

`func (o *Customer) GetContractAnnualRevenue() float64`

GetContractAnnualRevenue returns the ContractAnnualRevenue field if non-nil, zero value otherwise.

### GetContractAnnualRevenueOk

`func (o *Customer) GetContractAnnualRevenueOk() (*float64, bool)`

GetContractAnnualRevenueOk returns a tuple with the ContractAnnualRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAnnualRevenue

`func (o *Customer) SetContractAnnualRevenue(v float64)`

SetContractAnnualRevenue sets ContractAnnualRevenue field to given value.

### HasContractAnnualRevenue

`func (o *Customer) HasContractAnnualRevenue() bool`

HasContractAnnualRevenue returns a boolean if a field has been set.

### GetNotes

`func (o *Customer) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Customer) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Customer) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Customer) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


