# PeopleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Person**](Person.md) | A Person for each ID in the request, each with PersonMetadata populated. | [optional] 
**RelatedDocuments** | Pointer to [**[]RelatedDocuments**](RelatedDocuments.md) | A list of documents related to this people response. This is only included if DOCUMENT_ACTIVITY is requested and only 1 person is included in the request. | [optional] 
**Errors** | Pointer to **[]string** | A list of IDs that could not be found. | [optional] 

## Methods

### NewPeopleResponse

`func NewPeopleResponse() *PeopleResponse`

NewPeopleResponse instantiates a new PeopleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleResponseWithDefaults

`func NewPeopleResponseWithDefaults() *PeopleResponse`

NewPeopleResponseWithDefaults instantiates a new PeopleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PeopleResponse) GetResults() []Person`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PeopleResponse) GetResultsOk() (*[]Person, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PeopleResponse) SetResults(v []Person)`

SetResults sets Results field to given value.

### HasResults

`func (o *PeopleResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetRelatedDocuments

`func (o *PeopleResponse) GetRelatedDocuments() []RelatedDocuments`

GetRelatedDocuments returns the RelatedDocuments field if non-nil, zero value otherwise.

### GetRelatedDocumentsOk

`func (o *PeopleResponse) GetRelatedDocumentsOk() (*[]RelatedDocuments, bool)`

GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocuments

`func (o *PeopleResponse) SetRelatedDocuments(v []RelatedDocuments)`

SetRelatedDocuments sets RelatedDocuments field to given value.

### HasRelatedDocuments

`func (o *PeopleResponse) HasRelatedDocuments() bool`

HasRelatedDocuments returns a boolean if a field has been set.

### GetErrors

`func (o *PeopleResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PeopleResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PeopleResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PeopleResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


