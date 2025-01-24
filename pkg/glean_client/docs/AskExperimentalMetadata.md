# AskExperimentalMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryHasMentions** | Pointer to **bool** | Whether or not the query (i.e. the slack message) has a mention. | [optional] 
**QueryIsLengthAppropriate** | Pointer to **bool** | Whether or not the query (i.e. the slack message) is length appropriate. | [optional] 
**QueryIsAnswerable** | Pointer to **bool** | Whether or not the query (i.e. the slack message) has a question term. | [optional] 

## Methods

### NewAskExperimentalMetadata

`func NewAskExperimentalMetadata() *AskExperimentalMetadata`

NewAskExperimentalMetadata instantiates a new AskExperimentalMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskExperimentalMetadataWithDefaults

`func NewAskExperimentalMetadataWithDefaults() *AskExperimentalMetadata`

NewAskExperimentalMetadataWithDefaults instantiates a new AskExperimentalMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryHasMentions

`func (o *AskExperimentalMetadata) GetQueryHasMentions() bool`

GetQueryHasMentions returns the QueryHasMentions field if non-nil, zero value otherwise.

### GetQueryHasMentionsOk

`func (o *AskExperimentalMetadata) GetQueryHasMentionsOk() (*bool, bool)`

GetQueryHasMentionsOk returns a tuple with the QueryHasMentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryHasMentions

`func (o *AskExperimentalMetadata) SetQueryHasMentions(v bool)`

SetQueryHasMentions sets QueryHasMentions field to given value.

### HasQueryHasMentions

`func (o *AskExperimentalMetadata) HasQueryHasMentions() bool`

HasQueryHasMentions returns a boolean if a field has been set.

### GetQueryIsLengthAppropriate

`func (o *AskExperimentalMetadata) GetQueryIsLengthAppropriate() bool`

GetQueryIsLengthAppropriate returns the QueryIsLengthAppropriate field if non-nil, zero value otherwise.

### GetQueryIsLengthAppropriateOk

`func (o *AskExperimentalMetadata) GetQueryIsLengthAppropriateOk() (*bool, bool)`

GetQueryIsLengthAppropriateOk returns a tuple with the QueryIsLengthAppropriate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIsLengthAppropriate

`func (o *AskExperimentalMetadata) SetQueryIsLengthAppropriate(v bool)`

SetQueryIsLengthAppropriate sets QueryIsLengthAppropriate field to given value.

### HasQueryIsLengthAppropriate

`func (o *AskExperimentalMetadata) HasQueryIsLengthAppropriate() bool`

HasQueryIsLengthAppropriate returns a boolean if a field has been set.

### GetQueryIsAnswerable

`func (o *AskExperimentalMetadata) GetQueryIsAnswerable() bool`

GetQueryIsAnswerable returns the QueryIsAnswerable field if non-nil, zero value otherwise.

### GetQueryIsAnswerableOk

`func (o *AskExperimentalMetadata) GetQueryIsAnswerableOk() (*bool, bool)`

GetQueryIsAnswerableOk returns a tuple with the QueryIsAnswerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIsAnswerable

`func (o *AskExperimentalMetadata) SetQueryIsAnswerable(v bool)`

SetQueryIsAnswerable sets QueryIsAnswerable field to given value.

### HasQueryIsAnswerable

`func (o *AskExperimentalMetadata) HasQueryIsAnswerable() bool`

HasQueryIsAnswerable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


