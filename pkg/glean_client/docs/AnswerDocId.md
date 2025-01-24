# AnswerDocId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocId** | Pointer to **string** | Glean Document ID of the Answer. The Glean Document ID is supported for cases where the Answer ID isn&#39;t available. If both are available, using the Answer ID is preferred. | [optional] 

## Methods

### NewAnswerDocId

`func NewAnswerDocId() *AnswerDocId`

NewAnswerDocId instantiates a new AnswerDocId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerDocIdWithDefaults

`func NewAnswerDocIdWithDefaults() *AnswerDocId`

NewAnswerDocIdWithDefaults instantiates a new AnswerDocId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocId

`func (o *AnswerDocId) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *AnswerDocId) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *AnswerDocId) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *AnswerDocId) HasDocId() bool`

HasDocId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


