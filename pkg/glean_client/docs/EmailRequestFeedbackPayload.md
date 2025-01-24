# EmailRequestFeedbackPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueType** | Pointer to **string** | The type of issue being reported, e.g. RESULT_MISSING or OTHER for search feedback. | [optional] 
**Comments** | Pointer to **string** | Additional freeform comments provided by the reporter. | [optional] 
**Url** | Pointer to **string** | The URL the reporter was on when feedback was sent. | [optional] 
**Query** | Pointer to **string** | The query the reporter tried when feedback was sent. | [optional] 
**CustomJson** | Pointer to **string** | Arbitrary email param payloads from 3P-customer widgets. Prefer the structured fields when possible. | [optional] 

## Methods

### NewEmailRequestFeedbackPayload

`func NewEmailRequestFeedbackPayload() *EmailRequestFeedbackPayload`

NewEmailRequestFeedbackPayload instantiates a new EmailRequestFeedbackPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRequestFeedbackPayloadWithDefaults

`func NewEmailRequestFeedbackPayloadWithDefaults() *EmailRequestFeedbackPayload`

NewEmailRequestFeedbackPayloadWithDefaults instantiates a new EmailRequestFeedbackPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueType

`func (o *EmailRequestFeedbackPayload) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *EmailRequestFeedbackPayload) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *EmailRequestFeedbackPayload) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *EmailRequestFeedbackPayload) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetComments

`func (o *EmailRequestFeedbackPayload) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *EmailRequestFeedbackPayload) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *EmailRequestFeedbackPayload) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *EmailRequestFeedbackPayload) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetUrl

`func (o *EmailRequestFeedbackPayload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmailRequestFeedbackPayload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmailRequestFeedbackPayload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EmailRequestFeedbackPayload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetQuery

`func (o *EmailRequestFeedbackPayload) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EmailRequestFeedbackPayload) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EmailRequestFeedbackPayload) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EmailRequestFeedbackPayload) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetCustomJson

`func (o *EmailRequestFeedbackPayload) GetCustomJson() string`

GetCustomJson returns the CustomJson field if non-nil, zero value otherwise.

### GetCustomJsonOk

`func (o *EmailRequestFeedbackPayload) GetCustomJsonOk() (*string, bool)`

GetCustomJsonOk returns a tuple with the CustomJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomJson

`func (o *EmailRequestFeedbackPayload) SetCustomJson(v string)`

SetCustomJson sets CustomJson field to given value.

### HasCustomJson

`func (o *EmailRequestFeedbackPayload) HasCustomJson() bool`

HasCustomJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


