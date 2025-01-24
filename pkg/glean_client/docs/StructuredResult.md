# StructuredResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**Document**](Document.md) |  | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Customer** | Pointer to [**Customer**](Customer.md) |  | [optional] 
**Team** | Pointer to [**Team**](Team.md) |  | [optional] 
**CustomEntity** | Pointer to [**CustomEntity**](CustomEntity.md) |  | [optional] 
**Answer** | Pointer to [**Answer**](Answer.md) |  | [optional] 
**ExtractedQnA** | Pointer to [**ExtractedQnA**](ExtractedQnA.md) |  | [optional] 
**App** | Pointer to [**AppResult**](AppResult.md) |  | [optional] 
**Collection** | Pointer to [**Collection**](Collection.md) |  | [optional] 
**AnswerBoard** | Pointer to [**AnswerBoard**](AnswerBoard.md) |  | [optional] 
**Code** | Pointer to [**Code**](Code.md) |  | [optional] 
**Shortcut** | Pointer to [**Shortcut**](Shortcut.md) |  | [optional] 
**QuerySuggestions** | Pointer to [**QuerySuggestionList**](QuerySuggestionList.md) |  | [optional] 
**RelatedDocuments** | Pointer to [**[]RelatedDocuments**](RelatedDocuments.md) | A list of documents related to this structured result. | [optional] 
**RelatedQuestion** | Pointer to [**RelatedQuestion**](RelatedQuestion.md) |  | [optional] 
**Snippets** | Pointer to [**[]SearchResultSnippet**](SearchResultSnippet.md) | Any snippets associated to the populated object. | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular result in this particular query. To be used for /feedback reporting. | [optional] 
**Prominence** | Pointer to **string** | The level of visual distinction that should be given to a result. | [optional] 
**Source** | Pointer to **string** | Source context for this result. Possible values depend on the result type. | [optional] 

## Methods

### NewStructuredResult

`func NewStructuredResult() *StructuredResult`

NewStructuredResult instantiates a new StructuredResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuredResultWithDefaults

`func NewStructuredResultWithDefaults() *StructuredResult`

NewStructuredResultWithDefaults instantiates a new StructuredResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *StructuredResult) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *StructuredResult) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *StructuredResult) SetDocument(v Document)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *StructuredResult) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetPerson

`func (o *StructuredResult) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *StructuredResult) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *StructuredResult) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *StructuredResult) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetCustomer

`func (o *StructuredResult) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *StructuredResult) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *StructuredResult) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *StructuredResult) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetTeam

`func (o *StructuredResult) GetTeam() Team`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *StructuredResult) GetTeamOk() (*Team, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *StructuredResult) SetTeam(v Team)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *StructuredResult) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetCustomEntity

`func (o *StructuredResult) GetCustomEntity() CustomEntity`

GetCustomEntity returns the CustomEntity field if non-nil, zero value otherwise.

### GetCustomEntityOk

`func (o *StructuredResult) GetCustomEntityOk() (*CustomEntity, bool)`

GetCustomEntityOk returns a tuple with the CustomEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEntity

`func (o *StructuredResult) SetCustomEntity(v CustomEntity)`

SetCustomEntity sets CustomEntity field to given value.

### HasCustomEntity

`func (o *StructuredResult) HasCustomEntity() bool`

HasCustomEntity returns a boolean if a field has been set.

### GetAnswer

`func (o *StructuredResult) GetAnswer() Answer`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *StructuredResult) GetAnswerOk() (*Answer, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *StructuredResult) SetAnswer(v Answer)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *StructuredResult) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetExtractedQnA

`func (o *StructuredResult) GetExtractedQnA() ExtractedQnA`

GetExtractedQnA returns the ExtractedQnA field if non-nil, zero value otherwise.

### GetExtractedQnAOk

`func (o *StructuredResult) GetExtractedQnAOk() (*ExtractedQnA, bool)`

GetExtractedQnAOk returns a tuple with the ExtractedQnA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedQnA

`func (o *StructuredResult) SetExtractedQnA(v ExtractedQnA)`

SetExtractedQnA sets ExtractedQnA field to given value.

### HasExtractedQnA

`func (o *StructuredResult) HasExtractedQnA() bool`

HasExtractedQnA returns a boolean if a field has been set.

### GetApp

`func (o *StructuredResult) GetApp() AppResult`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *StructuredResult) GetAppOk() (*AppResult, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *StructuredResult) SetApp(v AppResult)`

SetApp sets App field to given value.

### HasApp

`func (o *StructuredResult) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCollection

`func (o *StructuredResult) GetCollection() Collection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *StructuredResult) GetCollectionOk() (*Collection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *StructuredResult) SetCollection(v Collection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *StructuredResult) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetAnswerBoard

`func (o *StructuredResult) GetAnswerBoard() AnswerBoard`

GetAnswerBoard returns the AnswerBoard field if non-nil, zero value otherwise.

### GetAnswerBoardOk

`func (o *StructuredResult) GetAnswerBoardOk() (*AnswerBoard, bool)`

GetAnswerBoardOk returns a tuple with the AnswerBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerBoard

`func (o *StructuredResult) SetAnswerBoard(v AnswerBoard)`

SetAnswerBoard sets AnswerBoard field to given value.

### HasAnswerBoard

`func (o *StructuredResult) HasAnswerBoard() bool`

HasAnswerBoard returns a boolean if a field has been set.

### GetCode

`func (o *StructuredResult) GetCode() Code`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StructuredResult) GetCodeOk() (*Code, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StructuredResult) SetCode(v Code)`

SetCode sets Code field to given value.

### HasCode

`func (o *StructuredResult) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetShortcut

`func (o *StructuredResult) GetShortcut() Shortcut`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *StructuredResult) GetShortcutOk() (*Shortcut, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *StructuredResult) SetShortcut(v Shortcut)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *StructuredResult) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetQuerySuggestions

`func (o *StructuredResult) GetQuerySuggestions() QuerySuggestionList`

GetQuerySuggestions returns the QuerySuggestions field if non-nil, zero value otherwise.

### GetQuerySuggestionsOk

`func (o *StructuredResult) GetQuerySuggestionsOk() (*QuerySuggestionList, bool)`

GetQuerySuggestionsOk returns a tuple with the QuerySuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySuggestions

`func (o *StructuredResult) SetQuerySuggestions(v QuerySuggestionList)`

SetQuerySuggestions sets QuerySuggestions field to given value.

### HasQuerySuggestions

`func (o *StructuredResult) HasQuerySuggestions() bool`

HasQuerySuggestions returns a boolean if a field has been set.

### GetRelatedDocuments

`func (o *StructuredResult) GetRelatedDocuments() []RelatedDocuments`

GetRelatedDocuments returns the RelatedDocuments field if non-nil, zero value otherwise.

### GetRelatedDocumentsOk

`func (o *StructuredResult) GetRelatedDocumentsOk() (*[]RelatedDocuments, bool)`

GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocuments

`func (o *StructuredResult) SetRelatedDocuments(v []RelatedDocuments)`

SetRelatedDocuments sets RelatedDocuments field to given value.

### HasRelatedDocuments

`func (o *StructuredResult) HasRelatedDocuments() bool`

HasRelatedDocuments returns a boolean if a field has been set.

### GetRelatedQuestion

`func (o *StructuredResult) GetRelatedQuestion() RelatedQuestion`

GetRelatedQuestion returns the RelatedQuestion field if non-nil, zero value otherwise.

### GetRelatedQuestionOk

`func (o *StructuredResult) GetRelatedQuestionOk() (*RelatedQuestion, bool)`

GetRelatedQuestionOk returns a tuple with the RelatedQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedQuestion

`func (o *StructuredResult) SetRelatedQuestion(v RelatedQuestion)`

SetRelatedQuestion sets RelatedQuestion field to given value.

### HasRelatedQuestion

`func (o *StructuredResult) HasRelatedQuestion() bool`

HasRelatedQuestion returns a boolean if a field has been set.

### GetSnippets

`func (o *StructuredResult) GetSnippets() []SearchResultSnippet`

GetSnippets returns the Snippets field if non-nil, zero value otherwise.

### GetSnippetsOk

`func (o *StructuredResult) GetSnippetsOk() (*[]SearchResultSnippet, bool)`

GetSnippetsOk returns a tuple with the Snippets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippets

`func (o *StructuredResult) SetSnippets(v []SearchResultSnippet)`

SetSnippets sets Snippets field to given value.

### HasSnippets

`func (o *StructuredResult) HasSnippets() bool`

HasSnippets returns a boolean if a field has been set.

### GetTrackingToken

`func (o *StructuredResult) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *StructuredResult) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *StructuredResult) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *StructuredResult) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetProminence

`func (o *StructuredResult) GetProminence() string`

GetProminence returns the Prominence field if non-nil, zero value otherwise.

### GetProminenceOk

`func (o *StructuredResult) GetProminenceOk() (*string, bool)`

GetProminenceOk returns a tuple with the Prominence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProminence

`func (o *StructuredResult) SetProminence(v string)`

SetProminence sets Prominence field to given value.

### HasProminence

`func (o *StructuredResult) HasProminence() bool`

HasProminence returns a boolean if a field has been set.

### GetSource

`func (o *StructuredResult) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StructuredResult) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StructuredResult) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StructuredResult) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


