# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name. | 
**ObfuscatedId** | **string** | An opaque identifier that can be used to request metadata for a Person. | 
**RelatedDocuments** | Pointer to [**[]RelatedDocuments**](RelatedDocuments.md) | A list of documents related to this person. | [optional] 
**Metadata** | Pointer to [**PersonMetadata**](PersonMetadata.md) |  | [optional] 

## Methods

### NewPerson

`func NewPerson(name string, obfuscatedId string, ) *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Person) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Person) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Person) SetName(v string)`

SetName sets Name field to given value.


### GetObfuscatedId

`func (o *Person) GetObfuscatedId() string`

GetObfuscatedId returns the ObfuscatedId field if non-nil, zero value otherwise.

### GetObfuscatedIdOk

`func (o *Person) GetObfuscatedIdOk() (*string, bool)`

GetObfuscatedIdOk returns a tuple with the ObfuscatedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedId

`func (o *Person) SetObfuscatedId(v string)`

SetObfuscatedId sets ObfuscatedId field to given value.


### GetRelatedDocuments

`func (o *Person) GetRelatedDocuments() []RelatedDocuments`

GetRelatedDocuments returns the RelatedDocuments field if non-nil, zero value otherwise.

### GetRelatedDocumentsOk

`func (o *Person) GetRelatedDocumentsOk() (*[]RelatedDocuments, bool)`

GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocuments

`func (o *Person) SetRelatedDocuments(v []RelatedDocuments)`

SetRelatedDocuments sets RelatedDocuments field to given value.

### HasRelatedDocuments

`func (o *Person) HasRelatedDocuments() bool`

HasRelatedDocuments returns a boolean if a field has been set.

### GetMetadata

`func (o *Person) GetMetadata() PersonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Person) GetMetadataOk() (*PersonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Person) SetMetadata(v PersonMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Person) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


