# ShortcutMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** | The time the shortcut was created in ISO format (ISO 8601). | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time the shortcut was updated in ISO format (ISO 8601). | [optional] 
**DestinationDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**IntermediateUrl** | Pointer to **string** | The URL from which the user is then redirected to the destination URL. Full replacement for https://go/&lt;inputAlias&gt;. | [optional] 
**ViewPrefix** | Pointer to **string** | The part of the shortcut preceding the input alias when used for showing shortcuts to users. Should end with \&quot;/\&quot;. e.g. \&quot;go/\&quot; for native shortcuts. | [optional] 
**IsExternal** | Pointer to **bool** | Indicates whether a shortcut is native or external. | [optional] 
**EditUrl** | Pointer to **string** | The URL using which the user can access the edit page of the shortcut. | [optional] 

## Methods

### NewShortcutMetadata

`func NewShortcutMetadata() *ShortcutMetadata`

NewShortcutMetadata instantiates a new ShortcutMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortcutMetadataWithDefaults

`func NewShortcutMetadataWithDefaults() *ShortcutMetadata`

NewShortcutMetadataWithDefaults instantiates a new ShortcutMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ShortcutMetadata) GetCreatedBy() Person`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ShortcutMetadata) GetCreatedByOk() (*Person, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ShortcutMetadata) SetCreatedBy(v Person)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ShortcutMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *ShortcutMetadata) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ShortcutMetadata) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ShortcutMetadata) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ShortcutMetadata) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ShortcutMetadata) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ShortcutMetadata) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ShortcutMetadata) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ShortcutMetadata) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdateTime

`func (o *ShortcutMetadata) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ShortcutMetadata) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ShortcutMetadata) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *ShortcutMetadata) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetDestinationDocument

`func (o *ShortcutMetadata) GetDestinationDocument() Document`

GetDestinationDocument returns the DestinationDocument field if non-nil, zero value otherwise.

### GetDestinationDocumentOk

`func (o *ShortcutMetadata) GetDestinationDocumentOk() (*Document, bool)`

GetDestinationDocumentOk returns a tuple with the DestinationDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDocument

`func (o *ShortcutMetadata) SetDestinationDocument(v Document)`

SetDestinationDocument sets DestinationDocument field to given value.

### HasDestinationDocument

`func (o *ShortcutMetadata) HasDestinationDocument() bool`

HasDestinationDocument returns a boolean if a field has been set.

### GetIntermediateUrl

`func (o *ShortcutMetadata) GetIntermediateUrl() string`

GetIntermediateUrl returns the IntermediateUrl field if non-nil, zero value otherwise.

### GetIntermediateUrlOk

`func (o *ShortcutMetadata) GetIntermediateUrlOk() (*string, bool)`

GetIntermediateUrlOk returns a tuple with the IntermediateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateUrl

`func (o *ShortcutMetadata) SetIntermediateUrl(v string)`

SetIntermediateUrl sets IntermediateUrl field to given value.

### HasIntermediateUrl

`func (o *ShortcutMetadata) HasIntermediateUrl() bool`

HasIntermediateUrl returns a boolean if a field has been set.

### GetViewPrefix

`func (o *ShortcutMetadata) GetViewPrefix() string`

GetViewPrefix returns the ViewPrefix field if non-nil, zero value otherwise.

### GetViewPrefixOk

`func (o *ShortcutMetadata) GetViewPrefixOk() (*string, bool)`

GetViewPrefixOk returns a tuple with the ViewPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewPrefix

`func (o *ShortcutMetadata) SetViewPrefix(v string)`

SetViewPrefix sets ViewPrefix field to given value.

### HasViewPrefix

`func (o *ShortcutMetadata) HasViewPrefix() bool`

HasViewPrefix returns a boolean if a field has been set.

### GetIsExternal

`func (o *ShortcutMetadata) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *ShortcutMetadata) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *ShortcutMetadata) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *ShortcutMetadata) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetEditUrl

`func (o *ShortcutMetadata) GetEditUrl() string`

GetEditUrl returns the EditUrl field if non-nil, zero value otherwise.

### GetEditUrlOk

`func (o *ShortcutMetadata) GetEditUrlOk() (*string, bool)`

GetEditUrlOk returns a tuple with the EditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditUrl

`func (o *ShortcutMetadata) SetEditUrl(v string)`

SetEditUrl sets EditUrl field to given value.

### HasEditUrl

`func (o *ShortcutMetadata) HasEditUrl() bool`

HasEditUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


