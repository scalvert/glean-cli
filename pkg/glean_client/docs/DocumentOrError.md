# DocumentOrError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Glean Document ID. | [optional] 
**Datasource** | Pointer to **string** | The app or other repository type from which the document was extracted | [optional] 
**ConnectorType** | Pointer to [**ConnectorType**](ConnectorType.md) |  | [optional] 
**DocType** | Pointer to **string** | The datasource-specific type of the document (e.g. for Jira issues, this is the issue type such as Bug or Feature Request). | [optional] 
**Content** | Pointer to [**DocumentContent**](DocumentContent.md) |  | [optional] 
**ContainerDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**ParentDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**Title** | Pointer to **string** | The title of the document. | [optional] 
**Url** | Pointer to **string** | A permalink for the document. | [optional] 
**Metadata** | Pointer to [**DocumentMetadata**](DocumentMetadata.md) |  | [optional] 
**Sections** | Pointer to [**[]DocumentSection**](DocumentSection.md) | A list of content sub-sections in the document, e.g. text blocks with different headings in a Drive doc or Confluence page. | [optional] 
**Error** | Pointer to **string** | The text for error, reason. | [optional] 

## Methods

### NewDocumentOrError

`func NewDocumentOrError() *DocumentOrError`

NewDocumentOrError instantiates a new DocumentOrError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentOrErrorWithDefaults

`func NewDocumentOrErrorWithDefaults() *DocumentOrError`

NewDocumentOrErrorWithDefaults instantiates a new DocumentOrError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentOrError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentOrError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentOrError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentOrError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDatasource

`func (o *DocumentOrError) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *DocumentOrError) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *DocumentOrError) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *DocumentOrError) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetConnectorType

`func (o *DocumentOrError) GetConnectorType() ConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *DocumentOrError) GetConnectorTypeOk() (*ConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *DocumentOrError) SetConnectorType(v ConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *DocumentOrError) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetDocType

`func (o *DocumentOrError) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *DocumentOrError) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *DocumentOrError) SetDocType(v string)`

SetDocType sets DocType field to given value.

### HasDocType

`func (o *DocumentOrError) HasDocType() bool`

HasDocType returns a boolean if a field has been set.

### GetContent

`func (o *DocumentOrError) GetContent() DocumentContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DocumentOrError) GetContentOk() (*DocumentContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DocumentOrError) SetContent(v DocumentContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *DocumentOrError) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContainerDocument

`func (o *DocumentOrError) GetContainerDocument() Document`

GetContainerDocument returns the ContainerDocument field if non-nil, zero value otherwise.

### GetContainerDocumentOk

`func (o *DocumentOrError) GetContainerDocumentOk() (*Document, bool)`

GetContainerDocumentOk returns a tuple with the ContainerDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDocument

`func (o *DocumentOrError) SetContainerDocument(v Document)`

SetContainerDocument sets ContainerDocument field to given value.

### HasContainerDocument

`func (o *DocumentOrError) HasContainerDocument() bool`

HasContainerDocument returns a boolean if a field has been set.

### GetParentDocument

`func (o *DocumentOrError) GetParentDocument() Document`

GetParentDocument returns the ParentDocument field if non-nil, zero value otherwise.

### GetParentDocumentOk

`func (o *DocumentOrError) GetParentDocumentOk() (*Document, bool)`

GetParentDocumentOk returns a tuple with the ParentDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDocument

`func (o *DocumentOrError) SetParentDocument(v Document)`

SetParentDocument sets ParentDocument field to given value.

### HasParentDocument

`func (o *DocumentOrError) HasParentDocument() bool`

HasParentDocument returns a boolean if a field has been set.

### GetTitle

`func (o *DocumentOrError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentOrError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentOrError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentOrError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *DocumentOrError) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DocumentOrError) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DocumentOrError) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DocumentOrError) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *DocumentOrError) GetMetadata() DocumentMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DocumentOrError) GetMetadataOk() (*DocumentMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DocumentOrError) SetMetadata(v DocumentMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DocumentOrError) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSections

`func (o *DocumentOrError) GetSections() []DocumentSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *DocumentOrError) GetSectionsOk() (*[]DocumentSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *DocumentOrError) SetSections(v []DocumentSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *DocumentOrError) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetError

`func (o *DocumentOrError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DocumentOrError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DocumentOrError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DocumentOrError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


