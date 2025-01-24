# AskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectOnly** | Pointer to **bool** | Whether to apply only question detection and not answering. | [optional] 
**AskExperimentalMetadata** | Pointer to [**AskExperimentalMetadata**](AskExperimentalMetadata.md) |  | [optional] 
**SearchRequest** | [**SearchRequest**](SearchRequest.md) |  | 
**ExcludedDocumentSpecs** | Pointer to [**[]DocumentSpec**](DocumentSpec.md) | A list of Glean Document IDs to be excluded when retrieving documents. Note that, currently, it only supports exclusion of one Glean Documnet ID based spec. If multiple specifications are provided only the first Glean Document ID based spec is excluded and the remaining specs are ignored. | [optional] 
**Operators** | Pointer to **string** | Search operators to append to the query | [optional] 
**Backend** | Pointer to **string** | Which backend to use to fulfill the requests. | [optional] 
**ChatApplicationId** | Pointer to **string** | The ID of the application this request originates from, used to determine the configuration of underlying chat processes when invoking the CHAT backend. This should correspond to the ID set during admin setup. If not specified, the default chat experience will be used. | [optional] 

## Methods

### NewAskRequest

`func NewAskRequest(searchRequest SearchRequest, ) *AskRequest`

NewAskRequest instantiates a new AskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskRequestWithDefaults

`func NewAskRequestWithDefaults() *AskRequest`

NewAskRequestWithDefaults instantiates a new AskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectOnly

`func (o *AskRequest) GetDetectOnly() bool`

GetDetectOnly returns the DetectOnly field if non-nil, zero value otherwise.

### GetDetectOnlyOk

`func (o *AskRequest) GetDetectOnlyOk() (*bool, bool)`

GetDetectOnlyOk returns a tuple with the DetectOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectOnly

`func (o *AskRequest) SetDetectOnly(v bool)`

SetDetectOnly sets DetectOnly field to given value.

### HasDetectOnly

`func (o *AskRequest) HasDetectOnly() bool`

HasDetectOnly returns a boolean if a field has been set.

### GetAskExperimentalMetadata

`func (o *AskRequest) GetAskExperimentalMetadata() AskExperimentalMetadata`

GetAskExperimentalMetadata returns the AskExperimentalMetadata field if non-nil, zero value otherwise.

### GetAskExperimentalMetadataOk

`func (o *AskRequest) GetAskExperimentalMetadataOk() (*AskExperimentalMetadata, bool)`

GetAskExperimentalMetadataOk returns a tuple with the AskExperimentalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskExperimentalMetadata

`func (o *AskRequest) SetAskExperimentalMetadata(v AskExperimentalMetadata)`

SetAskExperimentalMetadata sets AskExperimentalMetadata field to given value.

### HasAskExperimentalMetadata

`func (o *AskRequest) HasAskExperimentalMetadata() bool`

HasAskExperimentalMetadata returns a boolean if a field has been set.

### GetSearchRequest

`func (o *AskRequest) GetSearchRequest() SearchRequest`

GetSearchRequest returns the SearchRequest field if non-nil, zero value otherwise.

### GetSearchRequestOk

`func (o *AskRequest) GetSearchRequestOk() (*SearchRequest, bool)`

GetSearchRequestOk returns a tuple with the SearchRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchRequest

`func (o *AskRequest) SetSearchRequest(v SearchRequest)`

SetSearchRequest sets SearchRequest field to given value.


### GetExcludedDocumentSpecs

`func (o *AskRequest) GetExcludedDocumentSpecs() []DocumentSpec`

GetExcludedDocumentSpecs returns the ExcludedDocumentSpecs field if non-nil, zero value otherwise.

### GetExcludedDocumentSpecsOk

`func (o *AskRequest) GetExcludedDocumentSpecsOk() (*[]DocumentSpec, bool)`

GetExcludedDocumentSpecsOk returns a tuple with the ExcludedDocumentSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedDocumentSpecs

`func (o *AskRequest) SetExcludedDocumentSpecs(v []DocumentSpec)`

SetExcludedDocumentSpecs sets ExcludedDocumentSpecs field to given value.

### HasExcludedDocumentSpecs

`func (o *AskRequest) HasExcludedDocumentSpecs() bool`

HasExcludedDocumentSpecs returns a boolean if a field has been set.

### GetOperators

`func (o *AskRequest) GetOperators() string`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *AskRequest) GetOperatorsOk() (*string, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *AskRequest) SetOperators(v string)`

SetOperators sets Operators field to given value.

### HasOperators

`func (o *AskRequest) HasOperators() bool`

HasOperators returns a boolean if a field has been set.

### GetBackend

`func (o *AskRequest) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *AskRequest) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *AskRequest) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *AskRequest) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetChatApplicationId

`func (o *AskRequest) GetChatApplicationId() string`

GetChatApplicationId returns the ChatApplicationId field if non-nil, zero value otherwise.

### GetChatApplicationIdOk

`func (o *AskRequest) GetChatApplicationIdOk() (*string, bool)`

GetChatApplicationIdOk returns a tuple with the ChatApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatApplicationId

`func (o *AskRequest) SetChatApplicationId(v string)`

SetChatApplicationId sets ChatApplicationId field to given value.

### HasChatApplicationId

`func (o *AskRequest) HasChatApplicationId() bool`

HasChatApplicationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


