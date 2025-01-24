# GetDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentSpecs** | [**[]DocumentSpec**](DocumentSpec.md) | The specification for the documents to be retrieved. | 
**IncludeFields** | Pointer to **[]string** | List of Document fields to return (that aren&#39;t returned by default) | [optional] 

## Methods

### NewGetDocumentsRequest

`func NewGetDocumentsRequest(documentSpecs []DocumentSpec, ) *GetDocumentsRequest`

NewGetDocumentsRequest instantiates a new GetDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentsRequestWithDefaults

`func NewGetDocumentsRequestWithDefaults() *GetDocumentsRequest`

NewGetDocumentsRequestWithDefaults instantiates a new GetDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentSpecs

`func (o *GetDocumentsRequest) GetDocumentSpecs() []DocumentSpec`

GetDocumentSpecs returns the DocumentSpecs field if non-nil, zero value otherwise.

### GetDocumentSpecsOk

`func (o *GetDocumentsRequest) GetDocumentSpecsOk() (*[]DocumentSpec, bool)`

GetDocumentSpecsOk returns a tuple with the DocumentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSpecs

`func (o *GetDocumentsRequest) SetDocumentSpecs(v []DocumentSpec)`

SetDocumentSpecs sets DocumentSpecs field to given value.


### GetIncludeFields

`func (o *GetDocumentsRequest) GetIncludeFields() []string`

GetIncludeFields returns the IncludeFields field if non-nil, zero value otherwise.

### GetIncludeFieldsOk

`func (o *GetDocumentsRequest) GetIncludeFieldsOk() (*[]string, bool)`

GetIncludeFieldsOk returns a tuple with the IncludeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFields

`func (o *GetDocumentsRequest) SetIncludeFields(v []string)`

SetIncludeFields sets IncludeFields field to given value.

### HasIncludeFields

`func (o *GetDocumentsRequest) HasIncludeFields() bool`

HasIncludeFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


