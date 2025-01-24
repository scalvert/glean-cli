# TeamCreationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | For people field teams, the field value, e.g. ENGINEERING. Otherwise, a Glean Document ID used to identify the team. | [optional] 
**PeopleField** | Pointer to **string** | if the data source is people fields, then this is the field name (otherwise it&#39;s ignored) | [optional] 
**Datasource** | Pointer to **string** | what data source this team comes from, e.g. GDRIVE | [optional] 
**CreatedFrom** | Pointer to **string** | If the team is from a doc (i.e. not from a people field), this is the doc title, e.g. for Slack channels, the channel name. Otherwise, it&#39;s ignored. | [optional] 

## Methods

### NewTeamCreationData

`func NewTeamCreationData() *TeamCreationData`

NewTeamCreationData instantiates a new TeamCreationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamCreationDataWithDefaults

`func NewTeamCreationDataWithDefaults() *TeamCreationData`

NewTeamCreationDataWithDefaults instantiates a new TeamCreationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamCreationData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamCreationData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamCreationData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamCreationData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPeopleField

`func (o *TeamCreationData) GetPeopleField() string`

GetPeopleField returns the PeopleField field if non-nil, zero value otherwise.

### GetPeopleFieldOk

`func (o *TeamCreationData) GetPeopleFieldOk() (*string, bool)`

GetPeopleFieldOk returns a tuple with the PeopleField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeopleField

`func (o *TeamCreationData) SetPeopleField(v string)`

SetPeopleField sets PeopleField field to given value.

### HasPeopleField

`func (o *TeamCreationData) HasPeopleField() bool`

HasPeopleField returns a boolean if a field has been set.

### GetDatasource

`func (o *TeamCreationData) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *TeamCreationData) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *TeamCreationData) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *TeamCreationData) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetCreatedFrom

`func (o *TeamCreationData) GetCreatedFrom() string`

GetCreatedFrom returns the CreatedFrom field if non-nil, zero value otherwise.

### GetCreatedFromOk

`func (o *TeamCreationData) GetCreatedFromOk() (*string, bool)`

GetCreatedFromOk returns a tuple with the CreatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFrom

`func (o *TeamCreationData) SetCreatedFrom(v string)`

SetCreatedFrom sets CreatedFrom field to given value.

### HasCreatedFrom

`func (o *TeamCreationData) HasCreatedFrom() bool`

HasCreatedFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


