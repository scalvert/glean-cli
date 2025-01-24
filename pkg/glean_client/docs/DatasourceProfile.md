# DatasourceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | **string** | The datasource the profile is of. | 
**Handle** | **string** | The display name of the entity in the given datasource. | 
**Url** | Pointer to **string** | URL to view the entity&#39;s profile. | [optional] 
**NativeAppUrl** | Pointer to **string** | A deep link, if available, into the datasource&#39;s native application for the entity&#39;s platform (i.e. slack://...). | [optional] 
**IsUserGenerated** | Pointer to **bool** | For internal use only. True iff the data source profile was manually added by a user from within Glean (aka not from the original data source) | [optional] 

## Methods

### NewDatasourceProfile

`func NewDatasourceProfile(datasource string, handle string, ) *DatasourceProfile`

NewDatasourceProfile instantiates a new DatasourceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceProfileWithDefaults

`func NewDatasourceProfileWithDefaults() *DatasourceProfile`

NewDatasourceProfileWithDefaults instantiates a new DatasourceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *DatasourceProfile) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *DatasourceProfile) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *DatasourceProfile) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetHandle

`func (o *DatasourceProfile) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *DatasourceProfile) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *DatasourceProfile) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetUrl

`func (o *DatasourceProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DatasourceProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DatasourceProfile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DatasourceProfile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNativeAppUrl

`func (o *DatasourceProfile) GetNativeAppUrl() string`

GetNativeAppUrl returns the NativeAppUrl field if non-nil, zero value otherwise.

### GetNativeAppUrlOk

`func (o *DatasourceProfile) GetNativeAppUrlOk() (*string, bool)`

GetNativeAppUrlOk returns a tuple with the NativeAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeAppUrl

`func (o *DatasourceProfile) SetNativeAppUrl(v string)`

SetNativeAppUrl sets NativeAppUrl field to given value.

### HasNativeAppUrl

`func (o *DatasourceProfile) HasNativeAppUrl() bool`

HasNativeAppUrl returns a boolean if a field has been set.

### GetIsUserGenerated

`func (o *DatasourceProfile) GetIsUserGenerated() bool`

GetIsUserGenerated returns the IsUserGenerated field if non-nil, zero value otherwise.

### GetIsUserGeneratedOk

`func (o *DatasourceProfile) GetIsUserGeneratedOk() (*bool, bool)`

GetIsUserGeneratedOk returns a tuple with the IsUserGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserGenerated

`func (o *DatasourceProfile) SetIsUserGenerated(v bool)`

SetIsUserGenerated sets IsUserGenerated field to given value.

### HasIsUserGenerated

`func (o *DatasourceProfile) HasIsUserGenerated() bool`

HasIsUserGenerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


