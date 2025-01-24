# Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | The verification state for the document. | 
**Metadata** | Pointer to [**VerificationMetadata**](VerificationMetadata.md) |  | [optional] 

## Methods

### NewVerification

`func NewVerification(state string, ) *Verification`

NewVerification instantiates a new Verification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationWithDefaults

`func NewVerificationWithDefaults() *Verification`

NewVerificationWithDefaults instantiates a new Verification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *Verification) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Verification) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Verification) SetState(v string)`

SetState sets State field to given value.


### GetMetadata

`func (o *Verification) GetMetadata() VerificationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Verification) GetMetadataOk() (*VerificationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Verification) SetMetadata(v VerificationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Verification) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


