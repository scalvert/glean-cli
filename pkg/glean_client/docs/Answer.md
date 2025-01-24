# Answer

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
**Permissions** | Pointer to [**ObjectPermissions**](ObjectPermissions.md) |  | [optional] 
**CombinedAnswerText** | Pointer to [**StructuredText**](StructuredText.md) |  | [optional] 
**Likes** | Pointer to [**AnswerLikes**](AnswerLikes.md) |  | [optional] 
**Author** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** | The time the answer was created in ISO format (ISO 8601). | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time the answer was last updated in ISO format (ISO 8601). | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**Verification** | Pointer to [**Verification**](Verification.md) |  | [optional] 
**Board** | Pointer to [**AnswerBoard**](AnswerBoard.md) |  | [optional] 
**Collections** | Pointer to [**[]Collection**](Collection.md) | The collections to which the answer belongs. | [optional] 
**DocumentCategory** | Pointer to **string** | The document&#39;s document_category(.proto). | [optional] 
**SourceDocument** | Pointer to [**Document**](Document.md) |  | [optional] 

## Methods

### NewAnswer

`func NewAnswer(id int32, ) *Answer`

NewAnswer instantiates a new Answer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerWithDefaults

`func NewAnswerWithDefaults() *Answer`

NewAnswerWithDefaults instantiates a new Answer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Answer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Answer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Answer) SetId(v int32)`

SetId sets Id field to given value.


### GetDocId

`func (o *Answer) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *Answer) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *Answer) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *Answer) HasDocId() bool`

HasDocId returns a boolean if a field has been set.

### GetQuestion

`func (o *Answer) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *Answer) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *Answer) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *Answer) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetQuestionVariations

`func (o *Answer) GetQuestionVariations() []string`

GetQuestionVariations returns the QuestionVariations field if non-nil, zero value otherwise.

### GetQuestionVariationsOk

`func (o *Answer) GetQuestionVariationsOk() (*[]string, bool)`

GetQuestionVariationsOk returns a tuple with the QuestionVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionVariations

`func (o *Answer) SetQuestionVariations(v []string)`

SetQuestionVariations sets QuestionVariations field to given value.

### HasQuestionVariations

`func (o *Answer) HasQuestionVariations() bool`

HasQuestionVariations returns a boolean if a field has been set.

### GetBodyText

`func (o *Answer) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *Answer) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *Answer) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *Answer) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBoardId

`func (o *Answer) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *Answer) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *Answer) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *Answer) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *Answer) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *Answer) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *Answer) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *Answer) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetAddedRoles

`func (o *Answer) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *Answer) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *Answer) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *Answer) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *Answer) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *Answer) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *Answer) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *Answer) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetRoles

`func (o *Answer) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Answer) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Answer) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Answer) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSourceDocumentSpec

`func (o *Answer) GetSourceDocumentSpec() DocumentSpec`

GetSourceDocumentSpec returns the SourceDocumentSpec field if non-nil, zero value otherwise.

### GetSourceDocumentSpecOk

`func (o *Answer) GetSourceDocumentSpecOk() (*DocumentSpec, bool)`

GetSourceDocumentSpecOk returns a tuple with the SourceDocumentSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentSpec

`func (o *Answer) SetSourceDocumentSpec(v DocumentSpec)`

SetSourceDocumentSpec sets SourceDocumentSpec field to given value.

### HasSourceDocumentSpec

`func (o *Answer) HasSourceDocumentSpec() bool`

HasSourceDocumentSpec returns a boolean if a field has been set.

### GetSourceType

`func (o *Answer) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Answer) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Answer) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *Answer) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetPermissions

`func (o *Answer) GetPermissions() ObjectPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Answer) GetPermissionsOk() (*ObjectPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Answer) SetPermissions(v ObjectPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Answer) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCombinedAnswerText

`func (o *Answer) GetCombinedAnswerText() StructuredText`

GetCombinedAnswerText returns the CombinedAnswerText field if non-nil, zero value otherwise.

### GetCombinedAnswerTextOk

`func (o *Answer) GetCombinedAnswerTextOk() (*StructuredText, bool)`

GetCombinedAnswerTextOk returns a tuple with the CombinedAnswerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedAnswerText

`func (o *Answer) SetCombinedAnswerText(v StructuredText)`

SetCombinedAnswerText sets CombinedAnswerText field to given value.

### HasCombinedAnswerText

`func (o *Answer) HasCombinedAnswerText() bool`

HasCombinedAnswerText returns a boolean if a field has been set.

### GetLikes

`func (o *Answer) GetLikes() AnswerLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *Answer) GetLikesOk() (*AnswerLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *Answer) SetLikes(v AnswerLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *Answer) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetAuthor

`func (o *Answer) GetAuthor() Person`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Answer) GetAuthorOk() (*Person, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Answer) SetAuthor(v Person)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Answer) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreateTime

`func (o *Answer) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Answer) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Answer) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Answer) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Answer) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Answer) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Answer) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Answer) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Answer) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Answer) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Answer) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Answer) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVerification

`func (o *Answer) GetVerification() Verification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *Answer) GetVerificationOk() (*Verification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *Answer) SetVerification(v Verification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *Answer) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetBoard

`func (o *Answer) GetBoard() AnswerBoard`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *Answer) GetBoardOk() (*AnswerBoard, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *Answer) SetBoard(v AnswerBoard)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *Answer) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetCollections

`func (o *Answer) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *Answer) GetCollectionsOk() (*[]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *Answer) SetCollections(v []Collection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *Answer) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetDocumentCategory

`func (o *Answer) GetDocumentCategory() string`

GetDocumentCategory returns the DocumentCategory field if non-nil, zero value otherwise.

### GetDocumentCategoryOk

`func (o *Answer) GetDocumentCategoryOk() (*string, bool)`

GetDocumentCategoryOk returns a tuple with the DocumentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCategory

`func (o *Answer) SetDocumentCategory(v string)`

SetDocumentCategory sets DocumentCategory field to given value.

### HasDocumentCategory

`func (o *Answer) HasDocumentCategory() bool`

HasDocumentCategory returns a boolean if a field has been set.

### GetSourceDocument

`func (o *Answer) GetSourceDocument() Document`

GetSourceDocument returns the SourceDocument field if non-nil, zero value otherwise.

### GetSourceDocumentOk

`func (o *Answer) GetSourceDocumentOk() (*Document, bool)`

GetSourceDocumentOk returns a tuple with the SourceDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocument

`func (o *Answer) SetSourceDocument(v Document)`

SetSourceDocument sets SourceDocument field to given value.

### HasSourceDocument

`func (o *Answer) HasSourceDocument() bool`

HasSourceDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


