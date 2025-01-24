# GetDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**map[string]DocumentOrError**](DocumentOrError.md) | The document details or the error if document is not found. | [optional] 

## Methods

### NewGetDocumentsResponse

`func NewGetDocumentsResponse() *GetDocumentsResponse`

NewGetDocumentsResponse instantiates a new GetDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentsResponseWithDefaults

`func NewGetDocumentsResponseWithDefaults() *GetDocumentsResponse`

NewGetDocumentsResponseWithDefaults instantiates a new GetDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *GetDocumentsResponse) GetDocuments() map[string]DocumentOrError`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *GetDocumentsResponse) GetDocumentsOk() (*map[string]DocumentOrError, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *GetDocumentsResponse) SetDocuments(v map[string]DocumentOrError)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *GetDocumentsResponse) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


