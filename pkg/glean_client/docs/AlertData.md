# AlertData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the admin alert | [optional] 
**TriggeredTime** | Pointer to **time.Time** | The time that the alert was triggered | [optional] 
**ProjectName** | Pointer to **string** | Human readable name of the project instance | [optional] 
**HelpLink** | Pointer to **string** | Help link for the alert that the admin can reference | [optional] 
**Datasource** | Pointer to **string** | Datasource that the alert is related to (possibly null) | [optional] 

## Methods

### NewAlertData

`func NewAlertData() *AlertData`

NewAlertData instantiates a new AlertData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDataWithDefaults

`func NewAlertDataWithDefaults() *AlertData`

NewAlertDataWithDefaults instantiates a new AlertData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTriggeredTime

`func (o *AlertData) GetTriggeredTime() time.Time`

GetTriggeredTime returns the TriggeredTime field if non-nil, zero value otherwise.

### GetTriggeredTimeOk

`func (o *AlertData) GetTriggeredTimeOk() (*time.Time, bool)`

GetTriggeredTimeOk returns a tuple with the TriggeredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredTime

`func (o *AlertData) SetTriggeredTime(v time.Time)`

SetTriggeredTime sets TriggeredTime field to given value.

### HasTriggeredTime

`func (o *AlertData) HasTriggeredTime() bool`

HasTriggeredTime returns a boolean if a field has been set.

### GetProjectName

`func (o *AlertData) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *AlertData) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *AlertData) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *AlertData) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetHelpLink

`func (o *AlertData) GetHelpLink() string`

GetHelpLink returns the HelpLink field if non-nil, zero value otherwise.

### GetHelpLinkOk

`func (o *AlertData) GetHelpLinkOk() (*string, bool)`

GetHelpLinkOk returns a tuple with the HelpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpLink

`func (o *AlertData) SetHelpLink(v string)`

SetHelpLink sets HelpLink field to given value.

### HasHelpLink

`func (o *AlertData) HasHelpLink() bool`

HasHelpLink returns a boolean if a field has been set.

### GetDatasource

`func (o *AlertData) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *AlertData) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *AlertData) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *AlertData) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


