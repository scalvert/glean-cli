# AppResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datasource** | **string** | The app or other repository type this represents | 
**DocType** | Pointer to **string** | The datasource-specific type of the document (e.g. for Jira issues, this is the issue type such as Bug or Feature Request). | [optional] 
**MimeType** | Pointer to **string** | Mimetype is used to differentiate between sub applications from a datasource (e.g. Sheets, Docs from Gdrive) | [optional] 
**IconUrl** | Pointer to **string** | If there is available icon URL. | [optional] 

## Methods

### NewAppResult

`func NewAppResult(datasource string, ) *AppResult`

NewAppResult instantiates a new AppResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppResultWithDefaults

`func NewAppResultWithDefaults() *AppResult`

NewAppResultWithDefaults instantiates a new AppResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasource

`func (o *AppResult) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *AppResult) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *AppResult) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetDocType

`func (o *AppResult) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *AppResult) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *AppResult) SetDocType(v string)`

SetDocType sets DocType field to given value.

### HasDocType

`func (o *AppResult) HasDocType() bool`

HasDocType returns a boolean if a field has been set.

### GetMimeType

`func (o *AppResult) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AppResult) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AppResult) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AppResult) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetIconUrl

`func (o *AppResult) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *AppResult) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *AppResult) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *AppResult) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


