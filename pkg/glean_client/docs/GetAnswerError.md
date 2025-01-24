# GetAnswerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | Pointer to **string** |  | [optional] 
**AnswerAuthor** | Pointer to [**Person**](Person.md) |  | [optional] 

## Methods

### NewGetAnswerError

`func NewGetAnswerError() *GetAnswerError`

NewGetAnswerError instantiates a new GetAnswerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnswerErrorWithDefaults

`func NewGetAnswerErrorWithDefaults() *GetAnswerError`

NewGetAnswerErrorWithDefaults instantiates a new GetAnswerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *GetAnswerError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *GetAnswerError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *GetAnswerError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *GetAnswerError) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetAnswerAuthor

`func (o *GetAnswerError) GetAnswerAuthor() Person`

GetAnswerAuthor returns the AnswerAuthor field if non-nil, zero value otherwise.

### GetAnswerAuthorOk

`func (o *GetAnswerError) GetAnswerAuthorOk() (*Person, bool)`

GetAnswerAuthorOk returns a tuple with the AnswerAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerAuthor

`func (o *GetAnswerError) SetAnswerAuthor(v Person)`

SetAnswerAuthor sets AnswerAuthor field to given value.

### HasAnswerAuthor

`func (o *GetAnswerError) HasAnswerAuthor() bool`

HasAnswerAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


