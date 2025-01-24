# ViewerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | DEPRECATED - use permissions instead. Viewer&#39;s role on the specific document. | [optional] 
**LastViewedTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewViewerInfo

`func NewViewerInfo() *ViewerInfo`

NewViewerInfo instantiates a new ViewerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewerInfoWithDefaults

`func NewViewerInfoWithDefaults() *ViewerInfo`

NewViewerInfoWithDefaults instantiates a new ViewerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ViewerInfo) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ViewerInfo) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ViewerInfo) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ViewerInfo) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetLastViewedTime

`func (o *ViewerInfo) GetLastViewedTime() time.Time`

GetLastViewedTime returns the LastViewedTime field if non-nil, zero value otherwise.

### GetLastViewedTimeOk

`func (o *ViewerInfo) GetLastViewedTimeOk() (*time.Time, bool)`

GetLastViewedTimeOk returns a tuple with the LastViewedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedTime

`func (o *ViewerInfo) SetLastViewedTime(v time.Time)`

SetLastViewedTime sets LastViewedTime field to given value.

### HasLastViewedTime

`func (o *ViewerInfo) HasLastViewedTime() bool`

HasLastViewedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


