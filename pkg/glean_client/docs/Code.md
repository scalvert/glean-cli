# Code

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoName** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FileUrl** | Pointer to **string** |  | [optional] 
**Lines** | Pointer to [**[]CodeLine**](CodeLine.md) |  | [optional] 
**IsLastMatch** | Pointer to **bool** | Last file match for a repo | [optional] 

## Methods

### NewCode

`func NewCode() *Code`

NewCode instantiates a new Code object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeWithDefaults

`func NewCodeWithDefaults() *Code`

NewCodeWithDefaults instantiates a new Code object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoName

`func (o *Code) GetRepoName() string`

GetRepoName returns the RepoName field if non-nil, zero value otherwise.

### GetRepoNameOk

`func (o *Code) GetRepoNameOk() (*string, bool)`

GetRepoNameOk returns a tuple with the RepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoName

`func (o *Code) SetRepoName(v string)`

SetRepoName sets RepoName field to given value.

### HasRepoName

`func (o *Code) HasRepoName() bool`

HasRepoName returns a boolean if a field has been set.

### GetFileName

`func (o *Code) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Code) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Code) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Code) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileUrl

`func (o *Code) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *Code) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *Code) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *Code) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetLines

`func (o *Code) GetLines() []CodeLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Code) GetLinesOk() (*[]CodeLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Code) SetLines(v []CodeLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Code) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetIsLastMatch

`func (o *Code) GetIsLastMatch() bool`

GetIsLastMatch returns the IsLastMatch field if non-nil, zero value otherwise.

### GetIsLastMatchOk

`func (o *Code) GetIsLastMatchOk() (*bool, bool)`

GetIsLastMatchOk returns a tuple with the IsLastMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastMatch

`func (o *Code) SetIsLastMatch(v bool)`

SetIsLastMatch sets IsLastMatch field to given value.

### HasIsLastMatch

`func (o *Code) HasIsLastMatch() bool`

HasIsLastMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


