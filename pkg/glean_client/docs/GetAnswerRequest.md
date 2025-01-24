# GetAnswerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The opaque ID of the Answer. | [optional] 
**DocId** | Pointer to **string** | Glean Document ID of the Answer. The Glean Document ID is supported for cases where the Answer ID isn&#39;t available. If both are available, using the Answer ID is preferred. | [optional] 

## Methods

### NewGetAnswerRequest

`func NewGetAnswerRequest() *GetAnswerRequest`

NewGetAnswerRequest instantiates a new GetAnswerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnswerRequestWithDefaults

`func NewGetAnswerRequestWithDefaults() *GetAnswerRequest`

NewGetAnswerRequestWithDefaults instantiates a new GetAnswerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAnswerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAnswerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAnswerRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAnswerRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocId

`func (o *GetAnswerRequest) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *GetAnswerRequest) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *GetAnswerRequest) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *GetAnswerRequest) HasDocId() bool`

HasDocId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


