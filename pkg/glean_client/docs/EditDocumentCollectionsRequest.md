# EditDocumentCollectionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedCollections** | Pointer to **[]int32** | IDs of Collections to which a document is added. | [optional] 
**RemovedCollections** | Pointer to **[]int32** | IDs of Collections from which a document is removed. | [optional] 
**DocumentId** | Pointer to **string** | The Glean Document ID of the item being added to or removed from Collections if it&#39;s an indexed document. | [optional] 
**Url** | Pointer to **string** | The URL of the item being added to or removed from Collections. | [optional] 
**Name** | Pointer to **string** | Custom title of the document if adding a non-indexed URL. | [optional] 
**Description** | Pointer to **string** | The description of this CollectionItem. | [optional] 

## Methods

### NewEditDocumentCollectionsRequest

`func NewEditDocumentCollectionsRequest() *EditDocumentCollectionsRequest`

NewEditDocumentCollectionsRequest instantiates a new EditDocumentCollectionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDocumentCollectionsRequestWithDefaults

`func NewEditDocumentCollectionsRequestWithDefaults() *EditDocumentCollectionsRequest`

NewEditDocumentCollectionsRequestWithDefaults instantiates a new EditDocumentCollectionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedCollections

`func (o *EditDocumentCollectionsRequest) GetAddedCollections() []int32`

GetAddedCollections returns the AddedCollections field if non-nil, zero value otherwise.

### GetAddedCollectionsOk

`func (o *EditDocumentCollectionsRequest) GetAddedCollectionsOk() (*[]int32, bool)`

GetAddedCollectionsOk returns a tuple with the AddedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCollections

`func (o *EditDocumentCollectionsRequest) SetAddedCollections(v []int32)`

SetAddedCollections sets AddedCollections field to given value.

### HasAddedCollections

`func (o *EditDocumentCollectionsRequest) HasAddedCollections() bool`

HasAddedCollections returns a boolean if a field has been set.

### GetRemovedCollections

`func (o *EditDocumentCollectionsRequest) GetRemovedCollections() []int32`

GetRemovedCollections returns the RemovedCollections field if non-nil, zero value otherwise.

### GetRemovedCollectionsOk

`func (o *EditDocumentCollectionsRequest) GetRemovedCollectionsOk() (*[]int32, bool)`

GetRemovedCollectionsOk returns a tuple with the RemovedCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedCollections

`func (o *EditDocumentCollectionsRequest) SetRemovedCollections(v []int32)`

SetRemovedCollections sets RemovedCollections field to given value.

### HasRemovedCollections

`func (o *EditDocumentCollectionsRequest) HasRemovedCollections() bool`

HasRemovedCollections returns a boolean if a field has been set.

### GetDocumentId

`func (o *EditDocumentCollectionsRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *EditDocumentCollectionsRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *EditDocumentCollectionsRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *EditDocumentCollectionsRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetUrl

`func (o *EditDocumentCollectionsRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EditDocumentCollectionsRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EditDocumentCollectionsRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EditDocumentCollectionsRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *EditDocumentCollectionsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditDocumentCollectionsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditDocumentCollectionsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditDocumentCollectionsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EditDocumentCollectionsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditDocumentCollectionsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditDocumentCollectionsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditDocumentCollectionsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


