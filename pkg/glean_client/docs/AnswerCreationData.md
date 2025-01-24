# AnswerCreationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**CombinedAnswerText** | Pointer to [**StructuredTextMutableProperties**](StructuredTextMutableProperties.md) |  | [optional] 

## Methods

### NewAnswerCreationData

`func NewAnswerCreationData() *AnswerCreationData`

NewAnswerCreationData instantiates a new AnswerCreationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerCreationDataWithDefaults

`func NewAnswerCreationDataWithDefaults() *AnswerCreationData`

NewAnswerCreationDataWithDefaults instantiates a new AnswerCreationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *AnswerCreationData) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AnswerCreationData) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AnswerCreationData) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *AnswerCreationData) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetQuestionVariations

`func (o *AnswerCreationData) GetQuestionVariations() []string`

GetQuestionVariations returns the QuestionVariations field if non-nil, zero value otherwise.

### GetQuestionVariationsOk

`func (o *AnswerCreationData) GetQuestionVariationsOk() (*[]string, bool)`

GetQuestionVariationsOk returns a tuple with the QuestionVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionVariations

`func (o *AnswerCreationData) SetQuestionVariations(v []string)`

SetQuestionVariations sets QuestionVariations field to given value.

### HasQuestionVariations

`func (o *AnswerCreationData) HasQuestionVariations() bool`

HasQuestionVariations returns a boolean if a field has been set.

### GetBodyText

`func (o *AnswerCreationData) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *AnswerCreationData) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *AnswerCreationData) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *AnswerCreationData) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBoardId

`func (o *AnswerCreationData) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *AnswerCreationData) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *AnswerCreationData) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *AnswerCreationData) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *AnswerCreationData) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *AnswerCreationData) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *AnswerCreationData) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *AnswerCreationData) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetAddedRoles

`func (o *AnswerCreationData) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *AnswerCreationData) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *AnswerCreationData) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *AnswerCreationData) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *AnswerCreationData) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *AnswerCreationData) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *AnswerCreationData) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *AnswerCreationData) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetRoles

`func (o *AnswerCreationData) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AnswerCreationData) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AnswerCreationData) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AnswerCreationData) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSourceDocumentSpec

`func (o *AnswerCreationData) GetSourceDocumentSpec() DocumentSpec`

GetSourceDocumentSpec returns the SourceDocumentSpec field if non-nil, zero value otherwise.

### GetSourceDocumentSpecOk

`func (o *AnswerCreationData) GetSourceDocumentSpecOk() (*DocumentSpec, bool)`

GetSourceDocumentSpecOk returns a tuple with the SourceDocumentSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentSpec

`func (o *AnswerCreationData) SetSourceDocumentSpec(v DocumentSpec)`

SetSourceDocumentSpec sets SourceDocumentSpec field to given value.

### HasSourceDocumentSpec

`func (o *AnswerCreationData) HasSourceDocumentSpec() bool`

HasSourceDocumentSpec returns a boolean if a field has been set.

### GetSourceType

`func (o *AnswerCreationData) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AnswerCreationData) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AnswerCreationData) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AnswerCreationData) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetAddedCollections

`func (o *AnswerCreationData) GetAddedCollections() []int32`

GetAddedCollections returns the AddedCollections field if non-nil, zero value otherwise.

### GetAddedCollectionsOk

`func (o *AnswerCreationData) GetAddedCollectionsOk() (*[]int32, bool)`

GetAddedCollectionsOk returns a tuple with the AddedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCollections

`func (o *AnswerCreationData) SetAddedCollections(v []int32)`

SetAddedCollections sets AddedCollections field to given value.

### HasAddedCollections

`func (o *AnswerCreationData) HasAddedCollections() bool`

HasAddedCollections returns a boolean if a field has been set.

### GetCombinedAnswerText

`func (o *AnswerCreationData) GetCombinedAnswerText() StructuredTextMutableProperties`

GetCombinedAnswerText returns the CombinedAnswerText field if non-nil, zero value otherwise.

### GetCombinedAnswerTextOk

`func (o *AnswerCreationData) GetCombinedAnswerTextOk() (*StructuredTextMutableProperties, bool)`

GetCombinedAnswerTextOk returns a tuple with the CombinedAnswerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedAnswerText

`func (o *AnswerCreationData) SetCombinedAnswerText(v StructuredTextMutableProperties)`

SetCombinedAnswerText sets CombinedAnswerText field to given value.

### HasCombinedAnswerText

`func (o *AnswerCreationData) HasCombinedAnswerText() bool`

HasCombinedAnswerText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


