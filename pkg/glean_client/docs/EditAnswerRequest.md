# EditAnswerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The opaque ID of the Answer. | 
**DocId** | Pointer to **string** | Glean Document ID of the Answer. The Glean Document ID is supported for cases where the Answer ID isn&#39;t available. If both are available, using the Answer ID is preferred. | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**QuestionVariations** | Pointer to **[]string** | Additional ways of phrasing this question. | [optional] 
**BodyText** | Pointer to **string** | The plain text answer to the question. | [optional] 
**BoardId** | Pointer to **int32** | The parent board ID of this Answer, or 0 if it&#39;s a floating Answer. | [optional] 
**AudienceFilters** | Pointer to [**[]FacetFilter**](FacetFilter.md) | Filters which restrict who should see the answer. Values are taken from the corresponding filters in people search. | [optional] 
**AddedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the answer added by the owner. | [optional] 
**RemovedRoles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of user roles for the answer removed by the owner. | [optional] 
**Roles** | Pointer to [**[]UserRoleSpecification**](UserRoleSpecification.md) | A list of roles for this answer explicitly granted by an owner, editor, or admin. | [optional] 
**SourceDocumentSpec** | Pointer to [**DocumentSpec**](DocumentSpec.md) |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**AddedCollections** | Pointer to **[]int32** | IDs of Collections to which a document is added. | [optional] 
**RemovedCollections** | Pointer to **[]int32** | IDs of Collections from which a document is removed. | [optional] 
**CombinedAnswerText** | Pointer to [**StructuredTextMutableProperties**](StructuredTextMutableProperties.md) |  | [optional] 

## Methods

### NewEditAnswerRequest

`func NewEditAnswerRequest(id int32, ) *EditAnswerRequest`

NewEditAnswerRequest instantiates a new EditAnswerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAnswerRequestWithDefaults

`func NewEditAnswerRequestWithDefaults() *EditAnswerRequest`

NewEditAnswerRequestWithDefaults instantiates a new EditAnswerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditAnswerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditAnswerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditAnswerRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDocId

`func (o *EditAnswerRequest) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *EditAnswerRequest) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *EditAnswerRequest) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *EditAnswerRequest) HasDocId() bool`

HasDocId returns a boolean if a field has been set.

### GetQuestion

`func (o *EditAnswerRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *EditAnswerRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *EditAnswerRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *EditAnswerRequest) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetQuestionVariations

`func (o *EditAnswerRequest) GetQuestionVariations() []string`

GetQuestionVariations returns the QuestionVariations field if non-nil, zero value otherwise.

### GetQuestionVariationsOk

`func (o *EditAnswerRequest) GetQuestionVariationsOk() (*[]string, bool)`

GetQuestionVariationsOk returns a tuple with the QuestionVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionVariations

`func (o *EditAnswerRequest) SetQuestionVariations(v []string)`

SetQuestionVariations sets QuestionVariations field to given value.

### HasQuestionVariations

`func (o *EditAnswerRequest) HasQuestionVariations() bool`

HasQuestionVariations returns a boolean if a field has been set.

### GetBodyText

`func (o *EditAnswerRequest) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *EditAnswerRequest) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *EditAnswerRequest) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *EditAnswerRequest) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBoardId

`func (o *EditAnswerRequest) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *EditAnswerRequest) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *EditAnswerRequest) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *EditAnswerRequest) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *EditAnswerRequest) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *EditAnswerRequest) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *EditAnswerRequest) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *EditAnswerRequest) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetAddedRoles

`func (o *EditAnswerRequest) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *EditAnswerRequest) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *EditAnswerRequest) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *EditAnswerRequest) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *EditAnswerRequest) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *EditAnswerRequest) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *EditAnswerRequest) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *EditAnswerRequest) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetRoles

`func (o *EditAnswerRequest) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *EditAnswerRequest) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *EditAnswerRequest) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *EditAnswerRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSourceDocumentSpec

`func (o *EditAnswerRequest) GetSourceDocumentSpec() DocumentSpec`

GetSourceDocumentSpec returns the SourceDocumentSpec field if non-nil, zero value otherwise.

### GetSourceDocumentSpecOk

`func (o *EditAnswerRequest) GetSourceDocumentSpecOk() (*DocumentSpec, bool)`

GetSourceDocumentSpecOk returns a tuple with the SourceDocumentSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentSpec

`func (o *EditAnswerRequest) SetSourceDocumentSpec(v DocumentSpec)`

SetSourceDocumentSpec sets SourceDocumentSpec field to given value.

### HasSourceDocumentSpec

`func (o *EditAnswerRequest) HasSourceDocumentSpec() bool`

HasSourceDocumentSpec returns a boolean if a field has been set.

### GetSourceType

`func (o *EditAnswerRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *EditAnswerRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *EditAnswerRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *EditAnswerRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetAddedCollections

`func (o *EditAnswerRequest) GetAddedCollections() []int32`

GetAddedCollections returns the AddedCollections field if non-nil, zero value otherwise.

### GetAddedCollectionsOk

`func (o *EditAnswerRequest) GetAddedCollectionsOk() (*[]int32, bool)`

GetAddedCollectionsOk returns a tuple with the AddedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCollections

`func (o *EditAnswerRequest) SetAddedCollections(v []int32)`

SetAddedCollections sets AddedCollections field to given value.

### HasAddedCollections

`func (o *EditAnswerRequest) HasAddedCollections() bool`

HasAddedCollections returns a boolean if a field has been set.

### GetRemovedCollections

`func (o *EditAnswerRequest) GetRemovedCollections() []int32`

GetRemovedCollections returns the RemovedCollections field if non-nil, zero value otherwise.

### GetRemovedCollectionsOk

`func (o *EditAnswerRequest) GetRemovedCollectionsOk() (*[]int32, bool)`

GetRemovedCollectionsOk returns a tuple with the RemovedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedCollections

`func (o *EditAnswerRequest) SetRemovedCollections(v []int32)`

SetRemovedCollections sets RemovedCollections field to given value.

### HasRemovedCollections

`func (o *EditAnswerRequest) HasRemovedCollections() bool`

HasRemovedCollections returns a boolean if a field has been set.

### GetCombinedAnswerText

`func (o *EditAnswerRequest) GetCombinedAnswerText() StructuredTextMutableProperties`

GetCombinedAnswerText returns the CombinedAnswerText field if non-nil, zero value otherwise.

### GetCombinedAnswerTextOk

`func (o *EditAnswerRequest) GetCombinedAnswerTextOk() (*StructuredTextMutableProperties, bool)`

GetCombinedAnswerTextOk returns a tuple with the CombinedAnswerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedAnswerText

`func (o *EditAnswerRequest) SetCombinedAnswerText(v StructuredTextMutableProperties)`

SetCombinedAnswerText sets CombinedAnswerText field to given value.

### HasCombinedAnswerText

`func (o *EditAnswerRequest) HasCombinedAnswerText() bool`

HasCombinedAnswerText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


