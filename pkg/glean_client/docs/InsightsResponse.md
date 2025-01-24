# InsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeseries** | Pointer to [**[]LabeledCountInfo**](LabeledCountInfo.md) | List of timeseries to make charts (if applicable). | [optional] 
**Users** | Pointer to [**UserInsightsResponse**](UserInsightsResponse.md) |  | [optional] 
**Content** | Pointer to [**ContentInsightsResponse**](ContentInsightsResponse.md) |  | [optional] 
**Queries** | Pointer to [**QueryInsightsResponse**](QueryInsightsResponse.md) |  | [optional] 
**Collections** | Pointer to [**ContentInsightsResponse**](ContentInsightsResponse.md) |  | [optional] 
**CollectionsV2** | Pointer to [**ContentInsightsResponse**](ContentInsightsResponse.md) |  | [optional] 
**Shortcuts** | Pointer to [**ShortcutInsightsResponse**](ShortcutInsightsResponse.md) |  | [optional] 
**Announcements** | Pointer to [**ContentInsightsResponse**](ContentInsightsResponse.md) |  | [optional] 
**Answers** | Pointer to [**ContentInsightsResponse**](ContentInsightsResponse.md) |  | [optional] 
**Ai** | Pointer to [**AiInsightsResponse**](AiInsightsResponse.md) |  | [optional] 
**AiApps** | Pointer to [**AiAppsInsightsResponse**](AiAppsInsightsResponse.md) |  | [optional] 
**Departments** | Pointer to **[]string** | list of all departments. | [optional] 

## Methods

### NewInsightsResponse

`func NewInsightsResponse() *InsightsResponse`

NewInsightsResponse instantiates a new InsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsResponseWithDefaults

`func NewInsightsResponseWithDefaults() *InsightsResponse`

NewInsightsResponseWithDefaults instantiates a new InsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeseries

`func (o *InsightsResponse) GetTimeseries() []LabeledCountInfo`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *InsightsResponse) GetTimeseriesOk() (*[]LabeledCountInfo, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *InsightsResponse) SetTimeseries(v []LabeledCountInfo)`

SetTimeseries sets Timeseries field to given value.

### HasTimeseries

`func (o *InsightsResponse) HasTimeseries() bool`

HasTimeseries returns a boolean if a field has been set.

### GetUsers

`func (o *InsightsResponse) GetUsers() UserInsightsResponse`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InsightsResponse) GetUsersOk() (*UserInsightsResponse, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InsightsResponse) SetUsers(v UserInsightsResponse)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *InsightsResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetContent

`func (o *InsightsResponse) GetContent() ContentInsightsResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InsightsResponse) GetContentOk() (*ContentInsightsResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InsightsResponse) SetContent(v ContentInsightsResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *InsightsResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetQueries

`func (o *InsightsResponse) GetQueries() QueryInsightsResponse`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *InsightsResponse) GetQueriesOk() (*QueryInsightsResponse, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *InsightsResponse) SetQueries(v QueryInsightsResponse)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *InsightsResponse) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetCollections

`func (o *InsightsResponse) GetCollections() ContentInsightsResponse`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *InsightsResponse) GetCollectionsOk() (*ContentInsightsResponse, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *InsightsResponse) SetCollections(v ContentInsightsResponse)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *InsightsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetCollectionsV2

`func (o *InsightsResponse) GetCollectionsV2() ContentInsightsResponse`

GetCollectionsV2 returns the CollectionsV2 field if non-nil, zero value otherwise.

### GetCollectionsV2Ok

`func (o *InsightsResponse) GetCollectionsV2Ok() (*ContentInsightsResponse, bool)`

GetCollectionsV2Ok returns a tuple with the CollectionsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionsV2

`func (o *InsightsResponse) SetCollectionsV2(v ContentInsightsResponse)`

SetCollectionsV2 sets CollectionsV2 field to given value.

### HasCollectionsV2

`func (o *InsightsResponse) HasCollectionsV2() bool`

HasCollectionsV2 returns a boolean if a field has been set.

### GetShortcuts

`func (o *InsightsResponse) GetShortcuts() ShortcutInsightsResponse`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *InsightsResponse) GetShortcutsOk() (*ShortcutInsightsResponse, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *InsightsResponse) SetShortcuts(v ShortcutInsightsResponse)`

SetShortcuts sets Shortcuts field to given value.

### HasShortcuts

`func (o *InsightsResponse) HasShortcuts() bool`

HasShortcuts returns a boolean if a field has been set.

### GetAnnouncements

`func (o *InsightsResponse) GetAnnouncements() ContentInsightsResponse`

GetAnnouncements returns the Announcements field if non-nil, zero value otherwise.

### GetAnnouncementsOk

`func (o *InsightsResponse) GetAnnouncementsOk() (*ContentInsightsResponse, bool)`

GetAnnouncementsOk returns a tuple with the Announcements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncements

`func (o *InsightsResponse) SetAnnouncements(v ContentInsightsResponse)`

SetAnnouncements sets Announcements field to given value.

### HasAnnouncements

`func (o *InsightsResponse) HasAnnouncements() bool`

HasAnnouncements returns a boolean if a field has been set.

### GetAnswers

`func (o *InsightsResponse) GetAnswers() ContentInsightsResponse`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *InsightsResponse) GetAnswersOk() (*ContentInsightsResponse, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *InsightsResponse) SetAnswers(v ContentInsightsResponse)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *InsightsResponse) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetAi

`func (o *InsightsResponse) GetAi() AiInsightsResponse`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *InsightsResponse) GetAiOk() (*AiInsightsResponse, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *InsightsResponse) SetAi(v AiInsightsResponse)`

SetAi sets Ai field to given value.

### HasAi

`func (o *InsightsResponse) HasAi() bool`

HasAi returns a boolean if a field has been set.

### GetAiApps

`func (o *InsightsResponse) GetAiApps() AiAppsInsightsResponse`

GetAiApps returns the AiApps field if non-nil, zero value otherwise.

### GetAiAppsOk

`func (o *InsightsResponse) GetAiAppsOk() (*AiAppsInsightsResponse, bool)`

GetAiAppsOk returns a tuple with the AiApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiApps

`func (o *InsightsResponse) SetAiApps(v AiAppsInsightsResponse)`

SetAiApps sets AiApps field to given value.

### HasAiApps

`func (o *InsightsResponse) HasAiApps() bool`

HasAiApps returns a boolean if a field has been set.

### GetDepartments

`func (o *InsightsResponse) GetDepartments() []string`

GetDepartments returns the Departments field if non-nil, zero value otherwise.

### GetDepartmentsOk

`func (o *InsightsResponse) GetDepartmentsOk() (*[]string, bool)`

GetDepartmentsOk returns a tuple with the Departments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartments

`func (o *InsightsResponse) SetDepartments(v []string)`

SetDepartments sets Departments field to given value.

### HasDepartments

`func (o *InsightsResponse) HasDepartments() bool`

HasDepartments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


