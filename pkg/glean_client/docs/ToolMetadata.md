# ToolMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of tool. | 
**Name** | **string** | Unique identifier for the tool. Name should be understandable by the LLM, and will be used to invoke a tool. | 
**DisplayName** | **string** | Human understandable name of the tool. Max 50 characters. | 
**ToolId** | Pointer to **string** | An opaque id which is unique identifier for the tool. | [optional] 
**DisplayDescription** | **string** | Description of the tool meant for a human. | 
**LogoUrl** | Pointer to **string** | URL used to fetch the logo. | [optional] 
**ObjectName** | Pointer to **string** | Name of the generated object. This will be used to indicate to the end user what the generated object contains. | [optional] 
**CreatedBy** | Pointer to [**PersonObject**](PersonObject.md) |  | [optional] 
**LastUpdatedBy** | Pointer to [**PersonObject**](PersonObject.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the tool was created in ISO format (ISO 8601) | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | The time the tool was last updated in ISO format (ISO 8601) | [optional] 
**Auth** | Pointer to [**AuthConfig**](AuthConfig.md) |  | [optional] 

## Methods

### NewToolMetadata

`func NewToolMetadata(type_ string, name string, displayName string, displayDescription string, ) *ToolMetadata`

NewToolMetadata instantiates a new ToolMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolMetadataWithDefaults

`func NewToolMetadataWithDefaults() *ToolMetadata`

NewToolMetadataWithDefaults instantiates a new ToolMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ToolMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *ToolMetadata) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ToolMetadata) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ToolMetadata) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetToolId

`func (o *ToolMetadata) GetToolId() string`

GetToolId returns the ToolId field if non-nil, zero value otherwise.

### GetToolIdOk

`func (o *ToolMetadata) GetToolIdOk() (*string, bool)`

GetToolIdOk returns a tuple with the ToolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolId

`func (o *ToolMetadata) SetToolId(v string)`

SetToolId sets ToolId field to given value.

### HasToolId

`func (o *ToolMetadata) HasToolId() bool`

HasToolId returns a boolean if a field has been set.

### GetDisplayDescription

`func (o *ToolMetadata) GetDisplayDescription() string`

GetDisplayDescription returns the DisplayDescription field if non-nil, zero value otherwise.

### GetDisplayDescriptionOk

`func (o *ToolMetadata) GetDisplayDescriptionOk() (*string, bool)`

GetDisplayDescriptionOk returns a tuple with the DisplayDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDescription

`func (o *ToolMetadata) SetDisplayDescription(v string)`

SetDisplayDescription sets DisplayDescription field to given value.


### GetLogoUrl

`func (o *ToolMetadata) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ToolMetadata) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ToolMetadata) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *ToolMetadata) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetObjectName

`func (o *ToolMetadata) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ToolMetadata) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ToolMetadata) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ToolMetadata) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ToolMetadata) GetCreatedBy() PersonObject`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ToolMetadata) GetCreatedByOk() (*PersonObject, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ToolMetadata) SetCreatedBy(v PersonObject)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ToolMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *ToolMetadata) GetLastUpdatedBy() PersonObject`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ToolMetadata) GetLastUpdatedByOk() (*PersonObject, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ToolMetadata) SetLastUpdatedBy(v PersonObject)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ToolMetadata) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ToolMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ToolMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *ToolMetadata) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ToolMetadata) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ToolMetadata) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ToolMetadata) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetAuth

`func (o *ToolMetadata) GetAuth() AuthConfig`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ToolMetadata) GetAuthOk() (*AuthConfig, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ToolMetadata) SetAuth(v AuthConfig)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ToolMetadata) HasAuth() bool`

HasAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


