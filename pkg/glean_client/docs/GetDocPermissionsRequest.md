# GetDocPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | The Glean Document ID to retrieve permissions for. | [optional] 

## Methods

### NewGetDocPermissionsRequest

`func NewGetDocPermissionsRequest() *GetDocPermissionsRequest`

NewGetDocPermissionsRequest instantiates a new GetDocPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocPermissionsRequestWithDefaults

`func NewGetDocPermissionsRequestWithDefaults() *GetDocPermissionsRequest`

NewGetDocPermissionsRequestWithDefaults instantiates a new GetDocPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *GetDocPermissionsRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *GetDocPermissionsRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *GetDocPermissionsRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *GetDocPermissionsRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


