# RelatedDocuments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relation** | Pointer to **string** | How this document relates to the including entity. | [optional] 
**AssociatedEntityId** | Pointer to **string** | Which entity in the response that this entity relates to. Relevant when there are multiple entities associated with the response (such as merged customers) | [optional] 
**QuerySuggestion** | Pointer to [**QuerySuggestion**](QuerySuggestion.md) |  | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) | A truncated list of documents with this relation. TO BE DEPRECATED. | [optional] 
**Results** | Pointer to [**[]SearchResult**](SearchResult.md) | A truncated list of documents associated with this relation. To be used in favor of &#x60;documents&#x60; because it contains a trackingToken. | [optional] 

## Methods

### NewRelatedDocuments

`func NewRelatedDocuments() *RelatedDocuments`

NewRelatedDocuments instantiates a new RelatedDocuments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedDocumentsWithDefaults

`func NewRelatedDocumentsWithDefaults() *RelatedDocuments`

NewRelatedDocumentsWithDefaults instantiates a new RelatedDocuments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelation

`func (o *RelatedDocuments) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *RelatedDocuments) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *RelatedDocuments) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *RelatedDocuments) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetAssociatedEntityId

`func (o *RelatedDocuments) GetAssociatedEntityId() string`

GetAssociatedEntityId returns the AssociatedEntityId field if non-nil, zero value otherwise.

### GetAssociatedEntityIdOk

`func (o *RelatedDocuments) GetAssociatedEntityIdOk() (*string, bool)`

GetAssociatedEntityIdOk returns a tuple with the AssociatedEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedEntityId

`func (o *RelatedDocuments) SetAssociatedEntityId(v string)`

SetAssociatedEntityId sets AssociatedEntityId field to given value.

### HasAssociatedEntityId

`func (o *RelatedDocuments) HasAssociatedEntityId() bool`

HasAssociatedEntityId returns a boolean if a field has been set.

### GetQuerySuggestion

`func (o *RelatedDocuments) GetQuerySuggestion() QuerySuggestion`

GetQuerySuggestion returns the QuerySuggestion field if non-nil, zero value otherwise.

### GetQuerySuggestionOk

`func (o *RelatedDocuments) GetQuerySuggestionOk() (*QuerySuggestion, bool)`

GetQuerySuggestionOk returns a tuple with the QuerySuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySuggestion

`func (o *RelatedDocuments) SetQuerySuggestion(v QuerySuggestion)`

SetQuerySuggestion sets QuerySuggestion field to given value.

### HasQuerySuggestion

`func (o *RelatedDocuments) HasQuerySuggestion() bool`

HasQuerySuggestion returns a boolean if a field has been set.

### GetDocuments

`func (o *RelatedDocuments) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *RelatedDocuments) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *RelatedDocuments) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *RelatedDocuments) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetResults

`func (o *RelatedDocuments) GetResults() []SearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RelatedDocuments) GetResultsOk() (*[]SearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RelatedDocuments) SetResults(v []SearchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *RelatedDocuments) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


