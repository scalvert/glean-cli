# DocumentInteractions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumComments** | Pointer to **int32** | The count of comments (thread replies in the case of slack). | [optional] 
**NumReactions** | Pointer to **int32** | The count of reactions on the document. | [optional] 
**Reactions** | Pointer to **[]string** | To be deprecated in favor of reacts. A (potentially non-exhaustive) list of reactions for the document. | [optional] 
**Reacts** | Pointer to [**[]Reaction**](Reaction.md) |  | [optional] 
**Shares** | Pointer to [**[]Share**](Share.md) | Describes instances of someone posting a link to this document in one of our indexed datasources. | [optional] 
**VisitorCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 

## Methods

### NewDocumentInteractions

`func NewDocumentInteractions() *DocumentInteractions`

NewDocumentInteractions instantiates a new DocumentInteractions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentInteractionsWithDefaults

`func NewDocumentInteractionsWithDefaults() *DocumentInteractions`

NewDocumentInteractionsWithDefaults instantiates a new DocumentInteractions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumComments

`func (o *DocumentInteractions) GetNumComments() int32`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *DocumentInteractions) GetNumCommentsOk() (*int32, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *DocumentInteractions) SetNumComments(v int32)`

SetNumComments sets NumComments field to given value.

### HasNumComments

`func (o *DocumentInteractions) HasNumComments() bool`

HasNumComments returns a boolean if a field has been set.

### GetNumReactions

`func (o *DocumentInteractions) GetNumReactions() int32`

GetNumReactions returns the NumReactions field if non-nil, zero value otherwise.

### GetNumReactionsOk

`func (o *DocumentInteractions) GetNumReactionsOk() (*int32, bool)`

GetNumReactionsOk returns a tuple with the NumReactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReactions

`func (o *DocumentInteractions) SetNumReactions(v int32)`

SetNumReactions sets NumReactions field to given value.

### HasNumReactions

`func (o *DocumentInteractions) HasNumReactions() bool`

HasNumReactions returns a boolean if a field has been set.

### GetReactions

`func (o *DocumentInteractions) GetReactions() []string`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *DocumentInteractions) GetReactionsOk() (*[]string, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *DocumentInteractions) SetReactions(v []string)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *DocumentInteractions) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetReacts

`func (o *DocumentInteractions) GetReacts() []Reaction`

GetReacts returns the Reacts field if non-nil, zero value otherwise.

### GetReactsOk

`func (o *DocumentInteractions) GetReactsOk() (*[]Reaction, bool)`

GetReactsOk returns a tuple with the Reacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReacts

`func (o *DocumentInteractions) SetReacts(v []Reaction)`

SetReacts sets Reacts field to given value.

### HasReacts

`func (o *DocumentInteractions) HasReacts() bool`

HasReacts returns a boolean if a field has been set.

### GetShares

`func (o *DocumentInteractions) GetShares() []Share`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *DocumentInteractions) GetSharesOk() (*[]Share, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *DocumentInteractions) SetShares(v []Share)`

SetShares sets Shares field to given value.

### HasShares

`func (o *DocumentInteractions) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetVisitorCount

`func (o *DocumentInteractions) GetVisitorCount() CountInfo`

GetVisitorCount returns the VisitorCount field if non-nil, zero value otherwise.

### GetVisitorCountOk

`func (o *DocumentInteractions) GetVisitorCountOk() (*CountInfo, bool)`

GetVisitorCountOk returns a tuple with the VisitorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorCount

`func (o *DocumentInteractions) SetVisitorCount(v CountInfo)`

SetVisitorCount sets VisitorCount field to given value.

### HasVisitorCount

`func (o *DocumentInteractions) HasVisitorCount() bool`

HasVisitorCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


