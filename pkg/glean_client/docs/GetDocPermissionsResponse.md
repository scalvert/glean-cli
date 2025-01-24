# GetDocPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUserEmails** | Pointer to **[]string** | A list of emails of users who have access to the document. If the document is visible to all Glean users, a list with only a single value of &#39;VISIBLE_TO_ALL&#39;. | [optional] 

## Methods

### NewGetDocPermissionsResponse

`func NewGetDocPermissionsResponse() *GetDocPermissionsResponse`

NewGetDocPermissionsResponse instantiates a new GetDocPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocPermissionsResponseWithDefaults

`func NewGetDocPermissionsResponseWithDefaults() *GetDocPermissionsResponse`

NewGetDocPermissionsResponseWithDefaults instantiates a new GetDocPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUserEmails

`func (o *GetDocPermissionsResponse) GetAllowedUserEmails() []string`

GetAllowedUserEmails returns the AllowedUserEmails field if non-nil, zero value otherwise.

### GetAllowedUserEmailsOk

`func (o *GetDocPermissionsResponse) GetAllowedUserEmailsOk() (*[]string, bool)`

GetAllowedUserEmailsOk returns a tuple with the AllowedUserEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUserEmails

`func (o *GetDocPermissionsResponse) SetAllowedUserEmails(v []string)`

SetAllowedUserEmails sets AllowedUserEmails field to given value.

### HasAllowedUserEmails

`func (o *GetDocPermissionsResponse) HasAllowedUserEmails() bool`

HasAllowedUserEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


