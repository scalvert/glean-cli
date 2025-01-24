# AnswerMutableProperties

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

## Methods

### NewAnswerMutableProperties

`func NewAnswerMutableProperties() *AnswerMutableProperties`

NewAnswerMutableProperties instantiates a new AnswerMutableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerMutablePropertiesWithDefaults

`func NewAnswerMutablePropertiesWithDefaults() *AnswerMutableProperties`

NewAnswerMutablePropertiesWithDefaults instantiates a new AnswerMutableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *AnswerMutableProperties) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AnswerMutableProperties) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AnswerMutableProperties) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *AnswerMutableProperties) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetQuestionVariations

`func (o *AnswerMutableProperties) GetQuestionVariations() []string`

GetQuestionVariations returns the QuestionVariations field if non-nil, zero value otherwise.

### GetQuestionVariationsOk

`func (o *AnswerMutableProperties) GetQuestionVariationsOk() (*[]string, bool)`

GetQuestionVariationsOk returns a tuple with the QuestionVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionVariations

`func (o *AnswerMutableProperties) SetQuestionVariations(v []string)`

SetQuestionVariations sets QuestionVariations field to given value.

### HasQuestionVariations

`func (o *AnswerMutableProperties) HasQuestionVariations() bool`

HasQuestionVariations returns a boolean if a field has been set.

### GetBodyText

`func (o *AnswerMutableProperties) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *AnswerMutableProperties) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *AnswerMutableProperties) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *AnswerMutableProperties) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetBoardId

`func (o *AnswerMutableProperties) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *AnswerMutableProperties) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *AnswerMutableProperties) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *AnswerMutableProperties) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetAudienceFilters

`func (o *AnswerMutableProperties) GetAudienceFilters() []FacetFilter`

GetAudienceFilters returns the AudienceFilters field if non-nil, zero value otherwise.

### GetAudienceFiltersOk

`func (o *AnswerMutableProperties) GetAudienceFiltersOk() (*[]FacetFilter, bool)`

GetAudienceFiltersOk returns a tuple with the AudienceFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceFilters

`func (o *AnswerMutableProperties) SetAudienceFilters(v []FacetFilter)`

SetAudienceFilters sets AudienceFilters field to given value.

### HasAudienceFilters

`func (o *AnswerMutableProperties) HasAudienceFilters() bool`

HasAudienceFilters returns a boolean if a field has been set.

### GetAddedRoles

`func (o *AnswerMutableProperties) GetAddedRoles() []UserRoleSpecification`

GetAddedRoles returns the AddedRoles field if non-nil, zero value otherwise.

### GetAddedRolesOk

`func (o *AnswerMutableProperties) GetAddedRolesOk() (*[]UserRoleSpecification, bool)`

GetAddedRolesOk returns a tuple with the AddedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedRoles

`func (o *AnswerMutableProperties) SetAddedRoles(v []UserRoleSpecification)`

SetAddedRoles sets AddedRoles field to given value.

### HasAddedRoles

`func (o *AnswerMutableProperties) HasAddedRoles() bool`

HasAddedRoles returns a boolean if a field has been set.

### GetRemovedRoles

`func (o *AnswerMutableProperties) GetRemovedRoles() []UserRoleSpecification`

GetRemovedRoles returns the RemovedRoles field if non-nil, zero value otherwise.

### GetRemovedRolesOk

`func (o *AnswerMutableProperties) GetRemovedRolesOk() (*[]UserRoleSpecification, bool)`

GetRemovedRolesOk returns a tuple with the RemovedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedRoles

`func (o *AnswerMutableProperties) SetRemovedRoles(v []UserRoleSpecification)`

SetRemovedRoles sets RemovedRoles field to given value.

### HasRemovedRoles

`func (o *AnswerMutableProperties) HasRemovedRoles() bool`

HasRemovedRoles returns a boolean if a field has been set.

### GetRoles

`func (o *AnswerMutableProperties) GetRoles() []UserRoleSpecification`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AnswerMutableProperties) GetRolesOk() (*[]UserRoleSpecification, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AnswerMutableProperties) SetRoles(v []UserRoleSpecification)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AnswerMutableProperties) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSourceDocumentSpec

`func (o *AnswerMutableProperties) GetSourceDocumentSpec() DocumentSpec`

GetSourceDocumentSpec returns the SourceDocumentSpec field if non-nil, zero value otherwise.

### GetSourceDocumentSpecOk

`func (o *AnswerMutableProperties) GetSourceDocumentSpecOk() (*DocumentSpec, bool)`

GetSourceDocumentSpecOk returns a tuple with the SourceDocumentSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocumentSpec

`func (o *AnswerMutableProperties) SetSourceDocumentSpec(v DocumentSpec)`

SetSourceDocumentSpec sets SourceDocumentSpec field to given value.

### HasSourceDocumentSpec

`func (o *AnswerMutableProperties) HasSourceDocumentSpec() bool`

HasSourceDocumentSpec returns a boolean if a field has been set.

### GetSourceType

`func (o *AnswerMutableProperties) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AnswerMutableProperties) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AnswerMutableProperties) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AnswerMutableProperties) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


