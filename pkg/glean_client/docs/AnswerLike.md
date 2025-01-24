# AnswerLike

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** | The time the user liked the answer in ISO format (ISO 8601). | [optional] 

## Methods

### NewAnswerLike

`func NewAnswerLike() *AnswerLike`

NewAnswerLike instantiates a new AnswerLike object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerLikeWithDefaults

`func NewAnswerLikeWithDefaults() *AnswerLike`

NewAnswerLikeWithDefaults instantiates a new AnswerLike object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AnswerLike) GetUser() Person`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AnswerLike) GetUserOk() (*Person, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AnswerLike) SetUser(v Person)`

SetUser sets User field to given value.

### HasUser

`func (o *AnswerLike) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCreateTime

`func (o *AnswerLike) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AnswerLike) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AnswerLike) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AnswerLike) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


