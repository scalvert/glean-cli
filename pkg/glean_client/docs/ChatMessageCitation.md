# ChatMessageCitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular result in this particular ChatMessage. To be used for /feedback reporting. | [optional] 
**SourceDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**SourcePerson** | Pointer to [**Person**](Person.md) |  | [optional] 
**ReferenceRanges** | Pointer to [**[]ReferenceRange**](ReferenceRange.md) | Each reference range and its corresponding snippets | [optional] 

## Methods

### NewChatMessageCitation

`func NewChatMessageCitation() *ChatMessageCitation`

NewChatMessageCitation instantiates a new ChatMessageCitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageCitationWithDefaults

`func NewChatMessageCitationWithDefaults() *ChatMessageCitation`

NewChatMessageCitationWithDefaults instantiates a new ChatMessageCitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingToken

`func (o *ChatMessageCitation) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *ChatMessageCitation) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *ChatMessageCitation) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *ChatMessageCitation) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetSourceDocument

`func (o *ChatMessageCitation) GetSourceDocument() Document`

GetSourceDocument returns the SourceDocument field if non-nil, zero value otherwise.

### GetSourceDocumentOk

`func (o *ChatMessageCitation) GetSourceDocumentOk() (*Document, bool)`

GetSourceDocumentOk returns a tuple with the SourceDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocument

`func (o *ChatMessageCitation) SetSourceDocument(v Document)`

SetSourceDocument sets SourceDocument field to given value.

### HasSourceDocument

`func (o *ChatMessageCitation) HasSourceDocument() bool`

HasSourceDocument returns a boolean if a field has been set.

### GetSourcePerson

`func (o *ChatMessageCitation) GetSourcePerson() Person`

GetSourcePerson returns the SourcePerson field if non-nil, zero value otherwise.

### GetSourcePersonOk

`func (o *ChatMessageCitation) GetSourcePersonOk() (*Person, bool)`

GetSourcePersonOk returns a tuple with the SourcePerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePerson

`func (o *ChatMessageCitation) SetSourcePerson(v Person)`

SetSourcePerson sets SourcePerson field to given value.

### HasSourcePerson

`func (o *ChatMessageCitation) HasSourcePerson() bool`

HasSourcePerson returns a boolean if a field has been set.

### GetReferenceRanges

`func (o *ChatMessageCitation) GetReferenceRanges() []ReferenceRange`

GetReferenceRanges returns the ReferenceRanges field if non-nil, zero value otherwise.

### GetReferenceRangesOk

`func (o *ChatMessageCitation) GetReferenceRangesOk() (*[]ReferenceRange, bool)`

GetReferenceRangesOk returns a tuple with the ReferenceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceRanges

`func (o *ChatMessageCitation) SetReferenceRanges(v []ReferenceRange)`

SetReferenceRanges sets ReferenceRanges field to given value.

### HasReferenceRanges

`func (o *ChatMessageCitation) HasReferenceRanges() bool`

HasReferenceRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


