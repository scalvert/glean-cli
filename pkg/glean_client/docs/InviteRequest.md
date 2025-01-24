# InviteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**CommunicationChannel**](CommunicationChannel.md) |  | [optional] 
**Template** | Pointer to [**CommunicationTemplate**](CommunicationTemplate.md) |  | [optional] 
**Recipients** | Pointer to [**[]Person**](Person.md) | The people who should receive this invite | [optional] 
**RecipientFilters** | Pointer to [**PeopleFilters**](PeopleFilters.md) |  | [optional] 

## Methods

### NewInviteRequest

`func NewInviteRequest() *InviteRequest`

NewInviteRequest instantiates a new InviteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteRequestWithDefaults

`func NewInviteRequestWithDefaults() *InviteRequest`

NewInviteRequestWithDefaults instantiates a new InviteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *InviteRequest) GetChannel() CommunicationChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InviteRequest) GetChannelOk() (*CommunicationChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InviteRequest) SetChannel(v CommunicationChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InviteRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTemplate

`func (o *InviteRequest) GetTemplate() CommunicationTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *InviteRequest) GetTemplateOk() (*CommunicationTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *InviteRequest) SetTemplate(v CommunicationTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *InviteRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRecipients

`func (o *InviteRequest) GetRecipients() []Person`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InviteRequest) GetRecipientsOk() (*[]Person, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InviteRequest) SetRecipients(v []Person)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InviteRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetRecipientFilters

`func (o *InviteRequest) GetRecipientFilters() PeopleFilters`

GetRecipientFilters returns the RecipientFilters field if non-nil, zero value otherwise.

### GetRecipientFiltersOk

`func (o *InviteRequest) GetRecipientFiltersOk() (*PeopleFilters, bool)`

GetRecipientFiltersOk returns a tuple with the RecipientFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientFilters

`func (o *InviteRequest) SetRecipientFilters(v PeopleFilters)`

SetRecipientFilters sets RecipientFilters field to given value.

### HasRecipientFilters

`func (o *InviteRequest) HasRecipientFilters() bool`

HasRecipientFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


