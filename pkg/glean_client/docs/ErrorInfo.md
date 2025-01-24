# ErrorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BadGmailToken** | Pointer to **bool** | Indicates the gmail results could not be fetched due to bad token. | [optional] 
**BadOutlookToken** | Pointer to **bool** | Indicates the outlook results could not be fetched due to bad token. | [optional] 
**InvalidOperators** | Pointer to [**[]InvalidOperatorValueError**](InvalidOperatorValueError.md) | Indicates results could not be fetched due to invalid operators in the query. | [optional] 
**ErrorMessages** | Pointer to [**[]ErrorMessage**](ErrorMessage.md) |  | [optional] 

## Methods

### NewErrorInfo

`func NewErrorInfo() *ErrorInfo`

NewErrorInfo instantiates a new ErrorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorInfoWithDefaults

`func NewErrorInfoWithDefaults() *ErrorInfo`

NewErrorInfoWithDefaults instantiates a new ErrorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadGmailToken

`func (o *ErrorInfo) GetBadGmailToken() bool`

GetBadGmailToken returns the BadGmailToken field if non-nil, zero value otherwise.

### GetBadGmailTokenOk

`func (o *ErrorInfo) GetBadGmailTokenOk() (*bool, bool)`

GetBadGmailTokenOk returns a tuple with the BadGmailToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadGmailToken

`func (o *ErrorInfo) SetBadGmailToken(v bool)`

SetBadGmailToken sets BadGmailToken field to given value.

### HasBadGmailToken

`func (o *ErrorInfo) HasBadGmailToken() bool`

HasBadGmailToken returns a boolean if a field has been set.

### GetBadOutlookToken

`func (o *ErrorInfo) GetBadOutlookToken() bool`

GetBadOutlookToken returns the BadOutlookToken field if non-nil, zero value otherwise.

### GetBadOutlookTokenOk

`func (o *ErrorInfo) GetBadOutlookTokenOk() (*bool, bool)`

GetBadOutlookTokenOk returns a tuple with the BadOutlookToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadOutlookToken

`func (o *ErrorInfo) SetBadOutlookToken(v bool)`

SetBadOutlookToken sets BadOutlookToken field to given value.

### HasBadOutlookToken

`func (o *ErrorInfo) HasBadOutlookToken() bool`

HasBadOutlookToken returns a boolean if a field has been set.

### GetInvalidOperators

`func (o *ErrorInfo) GetInvalidOperators() []InvalidOperatorValueError`

GetInvalidOperators returns the InvalidOperators field if non-nil, zero value otherwise.

### GetInvalidOperatorsOk

`func (o *ErrorInfo) GetInvalidOperatorsOk() (*[]InvalidOperatorValueError, bool)`

GetInvalidOperatorsOk returns a tuple with the InvalidOperators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidOperators

`func (o *ErrorInfo) SetInvalidOperators(v []InvalidOperatorValueError)`

SetInvalidOperators sets InvalidOperators field to given value.

### HasInvalidOperators

`func (o *ErrorInfo) HasInvalidOperators() bool`

HasInvalidOperators returns a boolean if a field has been set.

### GetErrorMessages

`func (o *ErrorInfo) GetErrorMessages() []ErrorMessage`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *ErrorInfo) GetErrorMessagesOk() (*[]ErrorMessage, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *ErrorInfo) SetErrorMessages(v []ErrorMessage)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *ErrorInfo) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


