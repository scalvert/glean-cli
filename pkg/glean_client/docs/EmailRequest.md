# EmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplate** | [**CommunicationTemplate**](CommunicationTemplate.md) |  | 
**AlertData** | Pointer to [**AlertData**](AlertData.md) |  | [optional] 
**Recipients** | Pointer to [**[]Person**](Person.md) | The people to send emails to | [optional] 
**RecipientFilters** | Pointer to [**PeopleFilters**](PeopleFilters.md) |  | [optional] 
**CompanyName** | Pointer to **string** | Name of the company. | [optional] 
**DatasourceInstance** | Pointer to **string** | The instance ID of the datasource (if any) | [optional] 
**Senders** | Pointer to [**[]Person**](Person.md) | The people who triggered this email | [optional] 
**WebAppUrl** | Pointer to **string** | The URL of the client triggering the request, as received in the ClientConfig | [optional] 
**ServerUrl** | Pointer to **string** | The URL of the QE instance the email request is processed by. | [optional] 
**UnsubscribeUrl** | Pointer to **string** | The URL to unsubscribe from emails. | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) | The documents this email request refers to | [optional] 
**Reasons** | Pointer to **[]string** | Reasons this email request was sent. Will be shown directly to end user. | [optional] 
**Blocks** | Pointer to **map[string][]map[string]interface{}** | For building complex email UIs, we use a block structure that dictates what we create in the UI | [optional] 
**Subjects** | Pointer to **map[string]string** | Mapping of recipientIds to the email subject they are to receive. Optional and only meant for templates with Sendgrid subject set to {{subject}} | [optional] 
**FeedbackPayload** | Pointer to [**EmailRequestFeedbackPayload**](EmailRequestFeedbackPayload.md) |  | [optional] 
**ChatFeedbackPayload** | Pointer to [**EmailRequestChatFeedbackPayload**](EmailRequestChatFeedbackPayload.md) |  | [optional] 

## Methods

### NewEmailRequest

`func NewEmailRequest(emailTemplate CommunicationTemplate, ) *EmailRequest`

NewEmailRequest instantiates a new EmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRequestWithDefaults

`func NewEmailRequestWithDefaults() *EmailRequest`

NewEmailRequestWithDefaults instantiates a new EmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplate

`func (o *EmailRequest) GetEmailTemplate() CommunicationTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *EmailRequest) GetEmailTemplateOk() (*CommunicationTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *EmailRequest) SetEmailTemplate(v CommunicationTemplate)`

SetEmailTemplate sets EmailTemplate field to given value.


### GetAlertData

`func (o *EmailRequest) GetAlertData() AlertData`

GetAlertData returns the AlertData field if non-nil, zero value otherwise.

### GetAlertDataOk

`func (o *EmailRequest) GetAlertDataOk() (*AlertData, bool)`

GetAlertDataOk returns a tuple with the AlertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertData

`func (o *EmailRequest) SetAlertData(v AlertData)`

SetAlertData sets AlertData field to given value.

### HasAlertData

`func (o *EmailRequest) HasAlertData() bool`

HasAlertData returns a boolean if a field has been set.

### GetRecipients

`func (o *EmailRequest) GetRecipients() []Person`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailRequest) GetRecipientsOk() (*[]Person, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailRequest) SetRecipients(v []Person)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *EmailRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetRecipientFilters

`func (o *EmailRequest) GetRecipientFilters() PeopleFilters`

GetRecipientFilters returns the RecipientFilters field if non-nil, zero value otherwise.

### GetRecipientFiltersOk

`func (o *EmailRequest) GetRecipientFiltersOk() (*PeopleFilters, bool)`

GetRecipientFiltersOk returns a tuple with the RecipientFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientFilters

`func (o *EmailRequest) SetRecipientFilters(v PeopleFilters)`

SetRecipientFilters sets RecipientFilters field to given value.

### HasRecipientFilters

`func (o *EmailRequest) HasRecipientFilters() bool`

HasRecipientFilters returns a boolean if a field has been set.

### GetCompanyName

`func (o *EmailRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *EmailRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *EmailRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *EmailRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetDatasourceInstance

`func (o *EmailRequest) GetDatasourceInstance() string`

GetDatasourceInstance returns the DatasourceInstance field if non-nil, zero value otherwise.

### GetDatasourceInstanceOk

`func (o *EmailRequest) GetDatasourceInstanceOk() (*string, bool)`

GetDatasourceInstanceOk returns a tuple with the DatasourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceInstance

`func (o *EmailRequest) SetDatasourceInstance(v string)`

SetDatasourceInstance sets DatasourceInstance field to given value.

### HasDatasourceInstance

`func (o *EmailRequest) HasDatasourceInstance() bool`

HasDatasourceInstance returns a boolean if a field has been set.

### GetSenders

`func (o *EmailRequest) GetSenders() []Person`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *EmailRequest) GetSendersOk() (*[]Person, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *EmailRequest) SetSenders(v []Person)`

SetSenders sets Senders field to given value.

### HasSenders

`func (o *EmailRequest) HasSenders() bool`

HasSenders returns a boolean if a field has been set.

### GetWebAppUrl

`func (o *EmailRequest) GetWebAppUrl() string`

GetWebAppUrl returns the WebAppUrl field if non-nil, zero value otherwise.

### GetWebAppUrlOk

`func (o *EmailRequest) GetWebAppUrlOk() (*string, bool)`

GetWebAppUrlOk returns a tuple with the WebAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppUrl

`func (o *EmailRequest) SetWebAppUrl(v string)`

SetWebAppUrl sets WebAppUrl field to given value.

### HasWebAppUrl

`func (o *EmailRequest) HasWebAppUrl() bool`

HasWebAppUrl returns a boolean if a field has been set.

### GetServerUrl

`func (o *EmailRequest) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *EmailRequest) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *EmailRequest) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *EmailRequest) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetUnsubscribeUrl

`func (o *EmailRequest) GetUnsubscribeUrl() string`

GetUnsubscribeUrl returns the UnsubscribeUrl field if non-nil, zero value otherwise.

### GetUnsubscribeUrlOk

`func (o *EmailRequest) GetUnsubscribeUrlOk() (*string, bool)`

GetUnsubscribeUrlOk returns a tuple with the UnsubscribeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeUrl

`func (o *EmailRequest) SetUnsubscribeUrl(v string)`

SetUnsubscribeUrl sets UnsubscribeUrl field to given value.

### HasUnsubscribeUrl

`func (o *EmailRequest) HasUnsubscribeUrl() bool`

HasUnsubscribeUrl returns a boolean if a field has been set.

### GetDocuments

`func (o *EmailRequest) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *EmailRequest) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *EmailRequest) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *EmailRequest) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetReasons

`func (o *EmailRequest) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *EmailRequest) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *EmailRequest) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *EmailRequest) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetBlocks

`func (o *EmailRequest) GetBlocks() map[string][]map[string]interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *EmailRequest) GetBlocksOk() (*map[string][]map[string]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *EmailRequest) SetBlocks(v map[string][]map[string]interface{})`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *EmailRequest) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetSubjects

`func (o *EmailRequest) GetSubjects() map[string]string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *EmailRequest) GetSubjectsOk() (*map[string]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *EmailRequest) SetSubjects(v map[string]string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *EmailRequest) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetFeedbackPayload

`func (o *EmailRequest) GetFeedbackPayload() EmailRequestFeedbackPayload`

GetFeedbackPayload returns the FeedbackPayload field if non-nil, zero value otherwise.

### GetFeedbackPayloadOk

`func (o *EmailRequest) GetFeedbackPayloadOk() (*EmailRequestFeedbackPayload, bool)`

GetFeedbackPayloadOk returns a tuple with the FeedbackPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackPayload

`func (o *EmailRequest) SetFeedbackPayload(v EmailRequestFeedbackPayload)`

SetFeedbackPayload sets FeedbackPayload field to given value.

### HasFeedbackPayload

`func (o *EmailRequest) HasFeedbackPayload() bool`

HasFeedbackPayload returns a boolean if a field has been set.

### GetChatFeedbackPayload

`func (o *EmailRequest) GetChatFeedbackPayload() EmailRequestChatFeedbackPayload`

GetChatFeedbackPayload returns the ChatFeedbackPayload field if non-nil, zero value otherwise.

### GetChatFeedbackPayloadOk

`func (o *EmailRequest) GetChatFeedbackPayloadOk() (*EmailRequestChatFeedbackPayload, bool)`

GetChatFeedbackPayloadOk returns a tuple with the ChatFeedbackPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatFeedbackPayload

`func (o *EmailRequest) SetChatFeedbackPayload(v EmailRequestChatFeedbackPayload)`

SetChatFeedbackPayload sets ChatFeedbackPayload field to given value.

### HasChatFeedbackPayload

`func (o *EmailRequest) HasChatFeedbackPayload() bool`

HasChatFeedbackPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


