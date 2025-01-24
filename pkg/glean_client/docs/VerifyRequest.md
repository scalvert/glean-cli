# VerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** | The document which is verified. | 
**Action** | Pointer to **string** | The verification action requested. | [optional] 

## Methods

### NewVerifyRequest

`func NewVerifyRequest(documentId string, ) *VerifyRequest`

NewVerifyRequest instantiates a new VerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyRequestWithDefaults

`func NewVerifyRequestWithDefaults() *VerifyRequest`

NewVerifyRequestWithDefaults instantiates a new VerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *VerifyRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *VerifyRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *VerifyRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetAction

`func (o *VerifyRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VerifyRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VerifyRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VerifyRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


