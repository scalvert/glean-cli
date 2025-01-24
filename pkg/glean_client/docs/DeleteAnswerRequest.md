# DeleteAnswerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The opaque ID of the Answer. | 
**DocId** | Pointer to **string** | Glean Document ID of the Answer. The Glean Document ID is supported for cases where the Answer ID isn&#39;t available. If both are available, using the Answer ID is preferred. | [optional] 

## Methods

### NewDeleteAnswerRequest

`func NewDeleteAnswerRequest(id int32, ) *DeleteAnswerRequest`

NewDeleteAnswerRequest instantiates a new DeleteAnswerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAnswerRequestWithDefaults

`func NewDeleteAnswerRequestWithDefaults() *DeleteAnswerRequest`

NewDeleteAnswerRequestWithDefaults instantiates a new DeleteAnswerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteAnswerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteAnswerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteAnswerRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetDocId

`func (o *DeleteAnswerRequest) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *DeleteAnswerRequest) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *DeleteAnswerRequest) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *DeleteAnswerRequest) HasDocId() bool`

HasDocId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


