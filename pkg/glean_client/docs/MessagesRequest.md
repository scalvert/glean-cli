# MessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdType** | **string** | Type of the id in the incoming request. | 
**Id** | **string** | ID corresponding to the requested idType. Note that channel and threads are represented by the underlying datasource&#39;s ID and conversations are represented by their document&#39;s ID. | 
**WorkspaceId** | Pointer to **string** | Id for the for the workspace in case of multiple workspaces. | [optional] 
**Direction** | Pointer to **string** | The direction of the results asked with respect to the reference timestamp. Missing field defaults to OLDER. | [optional] 
**TimestampMillis** | **int64** | Timestamp in millis of the reference message. | 
**IncludeRootMessage** | Pointer to **bool** | Whether to include root message in response. | [optional] 
**Datasource** | Pointer to **string** | The type of the data source. Missing field defaults to SLACK. | [optional] 
**DatasourceInstanceDisplayName** | Pointer to **string** | The datasource instance display name from which the document was extracted. This is used for appinstance facet filter for datasources that support multiple instances. | [optional] 

## Methods

### NewMessagesRequest

`func NewMessagesRequest(idType string, id string, timestampMillis int64, ) *MessagesRequest`

NewMessagesRequest instantiates a new MessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesRequestWithDefaults

`func NewMessagesRequestWithDefaults() *MessagesRequest`

NewMessagesRequestWithDefaults instantiates a new MessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdType

`func (o *MessagesRequest) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *MessagesRequest) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *MessagesRequest) SetIdType(v string)`

SetIdType sets IdType field to given value.


### GetId

`func (o *MessagesRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessagesRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessagesRequest) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *MessagesRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *MessagesRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *MessagesRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *MessagesRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetDirection

`func (o *MessagesRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MessagesRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MessagesRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MessagesRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTimestampMillis

`func (o *MessagesRequest) GetTimestampMillis() int64`

GetTimestampMillis returns the TimestampMillis field if non-nil, zero value otherwise.

### GetTimestampMillisOk

`func (o *MessagesRequest) GetTimestampMillisOk() (*int64, bool)`

GetTimestampMillisOk returns a tuple with the TimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampMillis

`func (o *MessagesRequest) SetTimestampMillis(v int64)`

SetTimestampMillis sets TimestampMillis field to given value.


### GetIncludeRootMessage

`func (o *MessagesRequest) GetIncludeRootMessage() bool`

GetIncludeRootMessage returns the IncludeRootMessage field if non-nil, zero value otherwise.

### GetIncludeRootMessageOk

`func (o *MessagesRequest) GetIncludeRootMessageOk() (*bool, bool)`

GetIncludeRootMessageOk returns a tuple with the IncludeRootMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRootMessage

`func (o *MessagesRequest) SetIncludeRootMessage(v bool)`

SetIncludeRootMessage sets IncludeRootMessage field to given value.

### HasIncludeRootMessage

`func (o *MessagesRequest) HasIncludeRootMessage() bool`

HasIncludeRootMessage returns a boolean if a field has been set.

### GetDatasource

`func (o *MessagesRequest) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *MessagesRequest) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *MessagesRequest) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *MessagesRequest) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetDatasourceInstanceDisplayName

`func (o *MessagesRequest) GetDatasourceInstanceDisplayName() string`

GetDatasourceInstanceDisplayName returns the DatasourceInstanceDisplayName field if non-nil, zero value otherwise.

### GetDatasourceInstanceDisplayNameOk

`func (o *MessagesRequest) GetDatasourceInstanceDisplayNameOk() (*string, bool)`

GetDatasourceInstanceDisplayNameOk returns a tuple with the DatasourceInstanceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstanceDisplayName

`func (o *MessagesRequest) SetDatasourceInstanceDisplayName(v string)`

SetDatasourceInstanceDisplayName sets DatasourceInstanceDisplayName field to given value.

### HasDatasourceInstanceDisplayName

`func (o *MessagesRequest) HasDatasourceInstanceDisplayName() bool`

HasDatasourceInstanceDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


