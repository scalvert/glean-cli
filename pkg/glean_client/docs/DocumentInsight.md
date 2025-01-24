# DocumentInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | [**Document**](Document.md) |  | 
**ViewCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 
**VisitorCount** | Pointer to [**CountInfo**](CountInfo.md) |  | [optional] 

## Methods

### NewDocumentInsight

`func NewDocumentInsight(document Document, ) *DocumentInsight`

NewDocumentInsight instantiates a new DocumentInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentInsightWithDefaults

`func NewDocumentInsightWithDefaults() *DocumentInsight`

NewDocumentInsightWithDefaults instantiates a new DocumentInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *DocumentInsight) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *DocumentInsight) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *DocumentInsight) SetDocument(v Document)`

SetDocument sets Document field to given value.


### GetViewCount

`func (o *DocumentInsight) GetViewCount() CountInfo`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *DocumentInsight) GetViewCountOk() (*CountInfo, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *DocumentInsight) SetViewCount(v CountInfo)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *DocumentInsight) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### GetVisitorCount

`func (o *DocumentInsight) GetVisitorCount() CountInfo`

GetVisitorCount returns the VisitorCount field if non-nil, zero value otherwise.

### GetVisitorCountOk

`func (o *DocumentInsight) GetVisitorCountOk() (*CountInfo, bool)`

GetVisitorCountOk returns a tuple with the VisitorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorCount

`func (o *DocumentInsight) SetVisitorCount(v CountInfo)`

SetVisitorCount sets VisitorCount field to given value.

### HasVisitorCount

`func (o *DocumentInsight) HasVisitorCount() bool`

HasVisitorCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


