# GeneratedAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StrategyName** | Pointer to [**EventStrategyName**](EventStrategyName.md) |  | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) |  | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Customer** | Pointer to [**Customer**](Customer.md) |  | [optional] 
**ExternalLinks** | Pointer to [**[]StructuredLink**](StructuredLink.md) | A list of links to external sources outside of Glean. | [optional] 
**Content** | Pointer to [**[]GeneratedAttachmentContent**](GeneratedAttachmentContent.md) |  | [optional] 

## Methods

### NewGeneratedAttachment

`func NewGeneratedAttachment() *GeneratedAttachment`

NewGeneratedAttachment instantiates a new GeneratedAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratedAttachmentWithDefaults

`func NewGeneratedAttachmentWithDefaults() *GeneratedAttachment`

NewGeneratedAttachmentWithDefaults instantiates a new GeneratedAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategyName

`func (o *GeneratedAttachment) GetStrategyName() EventStrategyName`

GetStrategyName returns the StrategyName field if non-nil, zero value otherwise.

### GetStrategyNameOk

`func (o *GeneratedAttachment) GetStrategyNameOk() (*EventStrategyName, bool)`

GetStrategyNameOk returns a tuple with the StrategyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyName

`func (o *GeneratedAttachment) SetStrategyName(v EventStrategyName)`

SetStrategyName sets StrategyName field to given value.

### HasStrategyName

`func (o *GeneratedAttachment) HasStrategyName() bool`

HasStrategyName returns a boolean if a field has been set.

### GetDocuments

`func (o *GeneratedAttachment) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *GeneratedAttachment) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *GeneratedAttachment) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *GeneratedAttachment) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetPerson

`func (o *GeneratedAttachment) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *GeneratedAttachment) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *GeneratedAttachment) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *GeneratedAttachment) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetCustomer

`func (o *GeneratedAttachment) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GeneratedAttachment) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GeneratedAttachment) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *GeneratedAttachment) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetExternalLinks

`func (o *GeneratedAttachment) GetExternalLinks() []StructuredLink`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *GeneratedAttachment) GetExternalLinksOk() (*[]StructuredLink, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *GeneratedAttachment) SetExternalLinks(v []StructuredLink)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *GeneratedAttachment) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetContent

`func (o *GeneratedAttachment) GetContent() []GeneratedAttachmentContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GeneratedAttachment) GetContentOk() (*[]GeneratedAttachmentContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GeneratedAttachment) SetContent(v []GeneratedAttachmentContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *GeneratedAttachment) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


