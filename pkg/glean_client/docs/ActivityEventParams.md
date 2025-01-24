# ActivityEventParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyContent** | Pointer to **string** | The HTML content of the page body. | [optional] 
**DatasourceInstance** | Pointer to **string** | The full datasource instance name inferred from the URL of the event | [optional] 
**Datasource** | Pointer to **string** | The datasource without the instance inferred from the URL of the event | [optional] 
**InstanceOnlyName** | Pointer to **string** | The instance only name of the datasource instance, e.g. 1 for jira_1, inferred from the URL of the event | [optional] 
**Duration** | Pointer to **int32** | Length in seconds of the activity. For VIEWS, this represents the amount the page was visible in the foreground. | [optional] 
**Query** | Pointer to **string** | The user&#39;s search query associated with a SEARCH. | [optional] 
**Referrer** | Pointer to **string** | The referring URL of the VIEW or SEARCH. | [optional] 
**Title** | Pointer to **string** | The page title associated with the URL of the event | [optional] 
**Truncated** | Pointer to **bool** | Indicates that the parameters are incomplete and more parameters may be sent with the same action+timestamp+URL in the future. This is used for sending the duration when a &#x60;VIEW&#x60; is finished. | [optional] 

## Methods

### NewActivityEventParams

`func NewActivityEventParams() *ActivityEventParams`

NewActivityEventParams instantiates a new ActivityEventParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityEventParamsWithDefaults

`func NewActivityEventParamsWithDefaults() *ActivityEventParams`

NewActivityEventParamsWithDefaults instantiates a new ActivityEventParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyContent

`func (o *ActivityEventParams) GetBodyContent() string`

GetBodyContent returns the BodyContent field if non-nil, zero value otherwise.

### GetBodyContentOk

`func (o *ActivityEventParams) GetBodyContentOk() (*string, bool)`

GetBodyContentOk returns a tuple with the BodyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContent

`func (o *ActivityEventParams) SetBodyContent(v string)`

SetBodyContent sets BodyContent field to given value.

### HasBodyContent

`func (o *ActivityEventParams) HasBodyContent() bool`

HasBodyContent returns a boolean if a field has been set.

### GetDatasourceInstance

`func (o *ActivityEventParams) GetDatasourceInstance() string`

GetDatasourceInstance returns the DatasourceInstance field if non-nil, zero value otherwise.

### GetDatasourceInstanceOk

`func (o *ActivityEventParams) GetDatasourceInstanceOk() (*string, bool)`

GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstance

`func (o *ActivityEventParams) SetDatasourceInstance(v string)`

SetDatasourceInstance sets DatasourceInstance field to given value.

### HasDatasourceInstance

`func (o *ActivityEventParams) HasDatasourceInstance() bool`

HasDatasourceInstance returns a boolean if a field has been set.

### GetDatasource

`func (o *ActivityEventParams) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *ActivityEventParams) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *ActivityEventParams) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *ActivityEventParams) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetInstanceOnlyName

`func (o *ActivityEventParams) GetInstanceOnlyName() string`

GetInstanceOnlyName returns the InstanceOnlyName field if non-nil, zero value otherwise.

### GetInstanceOnlyNameOk

`func (o *ActivityEventParams) GetInstanceOnlyNameOk() (*string, bool)`

GetInstanceOnlyNameOk returns a tuple with the InstanceOnlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceOnlyName

`func (o *ActivityEventParams) SetInstanceOnlyName(v string)`

SetInstanceOnlyName sets InstanceOnlyName field to given value.

### HasInstanceOnlyName

`func (o *ActivityEventParams) HasInstanceOnlyName() bool`

HasInstanceOnlyName returns a boolean if a field has been set.

### GetDuration

`func (o *ActivityEventParams) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ActivityEventParams) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ActivityEventParams) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ActivityEventParams) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetQuery

`func (o *ActivityEventParams) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ActivityEventParams) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ActivityEventParams) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ActivityEventParams) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetReferrer

`func (o *ActivityEventParams) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *ActivityEventParams) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *ActivityEventParams) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *ActivityEventParams) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetTitle

`func (o *ActivityEventParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ActivityEventParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ActivityEventParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ActivityEventParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTruncated

`func (o *ActivityEventParams) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ActivityEventParams) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ActivityEventParams) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *ActivityEventParams) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


