# ListAnswerBoardsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WithAudience** | Pointer to **bool** | Whether to include the audience filters with the listed Answer Boards. | [optional] 
**WithRoles** | Pointer to **bool** | Whether to include the editor roles with the listed Answer Boards. | [optional] 

## Methods

### NewListAnswerBoardsRequest

`func NewListAnswerBoardsRequest() *ListAnswerBoardsRequest`

NewListAnswerBoardsRequest instantiates a new ListAnswerBoardsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnswerBoardsRequestWithDefaults

`func NewListAnswerBoardsRequestWithDefaults() *ListAnswerBoardsRequest`

NewListAnswerBoardsRequestWithDefaults instantiates a new ListAnswerBoardsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithAudience

`func (o *ListAnswerBoardsRequest) GetWithAudience() bool`

GetWithAudience returns the WithAudience field if non-nil, zero value otherwise.

### GetWithAudienceOk

`func (o *ListAnswerBoardsRequest) GetWithAudienceOk() (*bool, bool)`

GetWithAudienceOk returns a tuple with the WithAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAudience

`func (o *ListAnswerBoardsRequest) SetWithAudience(v bool)`

SetWithAudience sets WithAudience field to given value.

### HasWithAudience

`func (o *ListAnswerBoardsRequest) HasWithAudience() bool`

HasWithAudience returns a boolean if a field has been set.

### GetWithRoles

`func (o *ListAnswerBoardsRequest) GetWithRoles() bool`

GetWithRoles returns the WithRoles field if non-nil, zero value otherwise.

### GetWithRolesOk

`func (o *ListAnswerBoardsRequest) GetWithRolesOk() (*bool, bool)`

GetWithRolesOk returns a tuple with the WithRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithRoles

`func (o *ListAnswerBoardsRequest) SetWithRoles(v bool)`

SetWithRoles sets WithRoles field to given value.

### HasWithRoles

`func (o *ListAnswerBoardsRequest) HasWithRoles() bool`

HasWithRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


