# PreviewStructuredTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StructuredText** | [**StructuredText**](StructuredText.md) |  | 
**DocsInaccessibleToUser** | Pointer to **[]string** | A list of links the user doesn&#39;t have access to. | [optional] 
**CombinedAnswerText** | [**StructuredText**](StructuredText.md) |  | 

## Methods

### NewPreviewStructuredTextResponse

`func NewPreviewStructuredTextResponse(structuredText StructuredText, combinedAnswerText StructuredText, ) *PreviewStructuredTextResponse`

NewPreviewStructuredTextResponse instantiates a new PreviewStructuredTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewStructuredTextResponseWithDefaults

`func NewPreviewStructuredTextResponseWithDefaults() *PreviewStructuredTextResponse`

NewPreviewStructuredTextResponseWithDefaults instantiates a new PreviewStructuredTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStructuredText

`func (o *PreviewStructuredTextResponse) GetStructuredText() StructuredText`

GetStructuredText returns the StructuredText field if non-nil, zero value otherwise.

### GetStructuredTextOk

`func (o *PreviewStructuredTextResponse) GetStructuredTextOk() (*StructuredText, bool)`

GetStructuredTextOk returns a tuple with the StructuredText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredText

`func (o *PreviewStructuredTextResponse) SetStructuredText(v StructuredText)`

SetStructuredText sets StructuredText field to given value.


### GetDocsInaccessibleToUser

`func (o *PreviewStructuredTextResponse) GetDocsInaccessibleToUser() []string`

GetDocsInaccessibleToUser returns the DocsInaccessibleToUser field if non-nil, zero value otherwise.

### GetDocsInaccessibleToUserOk

`func (o *PreviewStructuredTextResponse) GetDocsInaccessibleToUserOk() (*[]string, bool)`

GetDocsInaccessibleToUserOk returns a tuple with the DocsInaccessibleToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsInaccessibleToUser

`func (o *PreviewStructuredTextResponse) SetDocsInaccessibleToUser(v []string)`

SetDocsInaccessibleToUser sets DocsInaccessibleToUser field to given value.

### HasDocsInaccessibleToUser

`func (o *PreviewStructuredTextResponse) HasDocsInaccessibleToUser() bool`

HasDocsInaccessibleToUser returns a boolean if a field has been set.

### GetCombinedAnswerText

`func (o *PreviewStructuredTextResponse) GetCombinedAnswerText() StructuredText`

GetCombinedAnswerText returns the CombinedAnswerText field if non-nil, zero value otherwise.

### GetCombinedAnswerTextOk

`func (o *PreviewStructuredTextResponse) GetCombinedAnswerTextOk() (*StructuredText, bool)`

GetCombinedAnswerTextOk returns a tuple with the CombinedAnswerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedAnswerText

`func (o *PreviewStructuredTextResponse) SetCombinedAnswerText(v StructuredText)`

SetCombinedAnswerText sets CombinedAnswerText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


