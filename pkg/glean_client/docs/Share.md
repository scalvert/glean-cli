# Share

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumDaysAgo** | **int32** | The number of days that has passed since the share happened | 
**Sharer** | Pointer to [**Person**](Person.md) |  | [optional] 
**SharingDocument** | Pointer to [**Document**](Document.md) |  | [optional] 

## Methods

### NewShare

`func NewShare(numDaysAgo int32, ) *Share`

NewShare instantiates a new Share object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareWithDefaults

`func NewShareWithDefaults() *Share`

NewShareWithDefaults instantiates a new Share object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumDaysAgo

`func (o *Share) GetNumDaysAgo() int32`

GetNumDaysAgo returns the NumDaysAgo field if non-nil, zero value otherwise.

### GetNumDaysAgoOk

`func (o *Share) GetNumDaysAgoOk() (*int32, bool)`

GetNumDaysAgoOk returns a tuple with the NumDaysAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDaysAgo

`func (o *Share) SetNumDaysAgo(v int32)`

SetNumDaysAgo sets NumDaysAgo field to given value.


### GetSharer

`func (o *Share) GetSharer() Person`

GetSharer returns the Sharer field if non-nil, zero value otherwise.

### GetSharerOk

`func (o *Share) GetSharerOk() (*Person, bool)`

GetSharerOk returns a tuple with the Sharer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharer

`func (o *Share) SetSharer(v Person)`

SetSharer sets Sharer field to given value.

### HasSharer

`func (o *Share) HasSharer() bool`

HasSharer returns a boolean if a field has been set.

### GetSharingDocument

`func (o *Share) GetSharingDocument() Document`

GetSharingDocument returns the SharingDocument field if non-nil, zero value otherwise.

### GetSharingDocumentOk

`func (o *Share) GetSharingDocumentOk() (*Document, bool)`

GetSharingDocumentOk returns a tuple with the SharingDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingDocument

`func (o *Share) SetSharingDocument(v Document)`

SetSharingDocument sets SharingDocument field to given value.

### HasSharingDocument

`func (o *Share) HasSharingDocument() bool`

HasSharingDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


