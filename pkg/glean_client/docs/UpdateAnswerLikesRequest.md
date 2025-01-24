# UpdateAnswerLikesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerId** | **int32** | The opaque id of the answer to like. | 
**AnswerDocId** | Pointer to **string** | Glean Document ID of the Answer. The Glean Document ID is supported for cases where the Answer ID is unavailable. If both are available, using the Answer ID is preferred. | [optional] 
**Action** | **string** |  | 

## Methods

### NewUpdateAnswerLikesRequest

`func NewUpdateAnswerLikesRequest(answerId int32, action string, ) *UpdateAnswerLikesRequest`

NewUpdateAnswerLikesRequest instantiates a new UpdateAnswerLikesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAnswerLikesRequestWithDefaults

`func NewUpdateAnswerLikesRequestWithDefaults() *UpdateAnswerLikesRequest`

NewUpdateAnswerLikesRequestWithDefaults instantiates a new UpdateAnswerLikesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerId

`func (o *UpdateAnswerLikesRequest) GetAnswerId() int32`

GetAnswerId returns the AnswerId field if non-nil, zero value otherwise.

### GetAnswerIdOk

`func (o *UpdateAnswerLikesRequest) GetAnswerIdOk() (*int32, bool)`

GetAnswerIdOk returns a tuple with the AnswerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerId

`func (o *UpdateAnswerLikesRequest) SetAnswerId(v int32)`

SetAnswerId sets AnswerId field to given value.


### GetAnswerDocId

`func (o *UpdateAnswerLikesRequest) GetAnswerDocId() string`

GetAnswerDocId returns the AnswerDocId field if non-nil, zero value otherwise.

### GetAnswerDocIdOk

`func (o *UpdateAnswerLikesRequest) GetAnswerDocIdOk() (*string, bool)`

GetAnswerDocIdOk returns a tuple with the AnswerDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerDocId

`func (o *UpdateAnswerLikesRequest) SetAnswerDocId(v string)`

SetAnswerDocId sets AnswerDocId field to given value.

### HasAnswerDocId

`func (o *UpdateAnswerLikesRequest) HasAnswerDocId() bool`

HasAnswerDocId returns a boolean if a field has been set.

### GetAction

`func (o *UpdateAnswerLikesRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateAnswerLikesRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateAnswerLikesRequest) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


