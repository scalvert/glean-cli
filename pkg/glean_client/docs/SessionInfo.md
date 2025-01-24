# SessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionTrackingToken** | Pointer to **string** | A unique token for this session. A new session (and token) is created when the user issues a request from a new tab or when our server hasn&#39;t seen activity for more than 10 minutes from a tab. | [optional] 
**TabId** | Pointer to **string** | A unique id for all requests a user makes from a given tab, no matter how far apart. A new tab id is only generated when a user issues a request from a new tab. | [optional] 
**LastSeen** | Pointer to **time.Time** | The last time the server saw this token. | [optional] 
**LastQuery** | Pointer to **string** | The last query seen by the server. | [optional] 

## Methods

### NewSessionInfo

`func NewSessionInfo() *SessionInfo`

NewSessionInfo instantiates a new SessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionInfoWithDefaults

`func NewSessionInfoWithDefaults() *SessionInfo`

NewSessionInfoWithDefaults instantiates a new SessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionTrackingToken

`func (o *SessionInfo) GetSessionTrackingToken() string`

GetSessionTrackingToken returns the SessionTrackingToken field if non-nil, zero value otherwise.

### GetSessionTrackingTokenOk

`func (o *SessionInfo) GetSessionTrackingTokenOk() (*string, bool)`

GetSessionTrackingTokenOk returns a tuple with the SessionTrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTrackingToken

`func (o *SessionInfo) SetSessionTrackingToken(v string)`

SetSessionTrackingToken sets SessionTrackingToken field to given value.

### HasSessionTrackingToken

`func (o *SessionInfo) HasSessionTrackingToken() bool`

HasSessionTrackingToken returns a boolean if a field has been set.

### GetTabId

`func (o *SessionInfo) GetTabId() string`

GetTabId returns the TabId field if non-nil, zero value otherwise.

### GetTabIdOk

`func (o *SessionInfo) GetTabIdOk() (*string, bool)`

GetTabIdOk returns a tuple with the TabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabId

`func (o *SessionInfo) SetTabId(v string)`

SetTabId sets TabId field to given value.

### HasTabId

`func (o *SessionInfo) HasTabId() bool`

HasTabId returns a boolean if a field has been set.

### GetLastSeen

`func (o *SessionInfo) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *SessionInfo) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *SessionInfo) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *SessionInfo) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastQuery

`func (o *SessionInfo) GetLastQuery() string`

GetLastQuery returns the LastQuery field if non-nil, zero value otherwise.

### GetLastQueryOk

`func (o *SessionInfo) GetLastQueryOk() (*string, bool)`

GetLastQueryOk returns a tuple with the LastQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQuery

`func (o *SessionInfo) SetLastQuery(v string)`

SetLastQuery sets LastQuery field to given value.

### HasLastQuery

`func (o *SessionInfo) HasLastQuery() bool`

HasLastQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


