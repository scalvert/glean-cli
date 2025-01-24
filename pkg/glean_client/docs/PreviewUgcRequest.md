# PreviewUgcRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Draft** | Pointer to [**UgcDraft**](UgcDraft.md) |  | [optional] 
**DraftSpec** | Pointer to [**DocumentSpec**](DocumentSpec.md) |  | [optional] 
**Type** | Pointer to [**UgcType**](UgcType.md) |  | [optional] 

## Methods

### NewPreviewUgcRequest

`func NewPreviewUgcRequest() *PreviewUgcRequest`

NewPreviewUgcRequest instantiates a new PreviewUgcRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewUgcRequestWithDefaults

`func NewPreviewUgcRequestWithDefaults() *PreviewUgcRequest`

NewPreviewUgcRequestWithDefaults instantiates a new PreviewUgcRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraft

`func (o *PreviewUgcRequest) GetDraft() UgcDraft`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PreviewUgcRequest) GetDraftOk() (*UgcDraft, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PreviewUgcRequest) SetDraft(v UgcDraft)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *PreviewUgcRequest) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetDraftSpec

`func (o *PreviewUgcRequest) GetDraftSpec() DocumentSpec`

GetDraftSpec returns the DraftSpec field if non-nil, zero value otherwise.

### GetDraftSpecOk

`func (o *PreviewUgcRequest) GetDraftSpecOk() (*DocumentSpec, bool)`

GetDraftSpecOk returns a tuple with the DraftSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftSpec

`func (o *PreviewUgcRequest) SetDraftSpec(v DocumentSpec)`

SetDraftSpec sets DraftSpec field to given value.

### HasDraftSpec

`func (o *PreviewUgcRequest) HasDraftSpec() bool`

HasDraftSpec returns a boolean if a field has been set.

### GetType

`func (o *PreviewUgcRequest) GetType() UgcType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PreviewUgcRequest) GetTypeOk() (*UgcType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PreviewUgcRequest) SetType(v UgcType)`

SetType sets Type field to given value.

### HasType

`func (o *PreviewUgcRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


