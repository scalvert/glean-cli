# ChatMessageFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructuredResults** | Pointer to [**[]StructuredResult**](StructuredResult.md) | An array of entities in the work graph retrieved via a data request. | [optional] 
**TrackingToken** | Pointer to **string** | An opaque token that represents this particular result in this particular query. To be used for /feedback reporting. | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**QuerySuggestion** | Pointer to [**QuerySuggestion**](QuerySuggestion.md) |  | [optional] 
**WriteAction** | Pointer to [**WriteAction**](WriteAction.md) |  | [optional] 

## Methods

### NewChatMessageFragment

`func NewChatMessageFragment() *ChatMessageFragment`

NewChatMessageFragment instantiates a new ChatMessageFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatMessageFragmentWithDefaults

`func NewChatMessageFragmentWithDefaults() *ChatMessageFragment`

NewChatMessageFragmentWithDefaults instantiates a new ChatMessageFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructuredResults

`func (o *ChatMessageFragment) GetStructuredResults() []StructuredResult`

GetStructuredResults returns the StructuredResults field if non-nil, zero value otherwise.

### GetStructuredResultsOk

`func (o *ChatMessageFragment) GetStructuredResultsOk() (*[]StructuredResult, bool)`

GetStructuredResultsOk returns a tuple with the StructuredResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredResults

`func (o *ChatMessageFragment) SetStructuredResults(v []StructuredResult)`

SetStructuredResults sets StructuredResults field to given value.

### HasStructuredResults

`func (o *ChatMessageFragment) HasStructuredResults() bool`

HasStructuredResults returns a boolean if a field has been set.

### GetTrackingToken

`func (o *ChatMessageFragment) GetTrackingToken() string`

GetTrackingToken returns the TrackingToken field if non-nil, zero value otherwise.

### GetTrackingTokenOk

`func (o *ChatMessageFragment) GetTrackingTokenOk() (*string, bool)`

GetTrackingTokenOk returns a tuple with the TrackingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingToken

`func (o *ChatMessageFragment) SetTrackingToken(v string)`

SetTrackingToken sets TrackingToken field to given value.

### HasTrackingToken

`func (o *ChatMessageFragment) HasTrackingToken() bool`

HasTrackingToken returns a boolean if a field has been set.

### GetText

`func (o *ChatMessageFragment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChatMessageFragment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChatMessageFragment) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ChatMessageFragment) HasText() bool`

HasText returns a boolean if a field has been set.

### GetQuerySuggestion

`func (o *ChatMessageFragment) GetQuerySuggestion() QuerySuggestion`

GetQuerySuggestion returns the QuerySuggestion field if non-nil, zero value otherwise.

### GetQuerySuggestionOk

`func (o *ChatMessageFragment) GetQuerySuggestionOk() (*QuerySuggestion, bool)`

GetQuerySuggestionOk returns a tuple with the QuerySuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySuggestion

`func (o *ChatMessageFragment) SetQuerySuggestion(v QuerySuggestion)`

SetQuerySuggestion sets QuerySuggestion field to given value.

### HasQuerySuggestion

`func (o *ChatMessageFragment) HasQuerySuggestion() bool`

HasQuerySuggestion returns a boolean if a field has been set.

### GetWriteAction

`func (o *ChatMessageFragment) GetWriteAction() WriteAction`

GetWriteAction returns the WriteAction field if non-nil, zero value otherwise.

### GetWriteActionOk

`func (o *ChatMessageFragment) GetWriteActionOk() (*WriteAction, bool)`

GetWriteActionOk returns a tuple with the WriteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAction

`func (o *ChatMessageFragment) SetWriteAction(v WriteAction)`

SetWriteAction sets WriteAction field to given value.

### HasWriteAction

`func (o *ChatMessageFragment) HasWriteAction() bool`

HasWriteAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


