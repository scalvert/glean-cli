# AnswerLikes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LikedBy** | [**[]AnswerLike**](AnswerLike.md) |  | 
**LikedByUser** | **bool** | Whether the user in context liked the answer. | 
**NumLikes** | **int32** | The total number of likes for the answer. | 

## Methods

### NewAnswerLikes

`func NewAnswerLikes(likedBy []AnswerLike, likedByUser bool, numLikes int32, ) *AnswerLikes`

NewAnswerLikes instantiates a new AnswerLikes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerLikesWithDefaults

`func NewAnswerLikesWithDefaults() *AnswerLikes`

NewAnswerLikesWithDefaults instantiates a new AnswerLikes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLikedBy

`func (o *AnswerLikes) GetLikedBy() []AnswerLike`

GetLikedBy returns the LikedBy field if non-nil, zero value otherwise.

### GetLikedByOk

`func (o *AnswerLikes) GetLikedByOk() (*[]AnswerLike, bool)`

GetLikedByOk returns a tuple with the LikedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikedBy

`func (o *AnswerLikes) SetLikedBy(v []AnswerLike)`

SetLikedBy sets LikedBy field to given value.


### GetLikedByUser

`func (o *AnswerLikes) GetLikedByUser() bool`

GetLikedByUser returns the LikedByUser field if non-nil, zero value otherwise.

### GetLikedByUserOk

`func (o *AnswerLikes) GetLikedByUserOk() (*bool, bool)`

GetLikedByUserOk returns a tuple with the LikedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikedByUser

`func (o *AnswerLikes) SetLikedByUser(v bool)`

SetLikedByUser sets LikedByUser field to given value.


### GetNumLikes

`func (o *AnswerLikes) GetNumLikes() int32`

GetNumLikes returns the NumLikes field if non-nil, zero value otherwise.

### GetNumLikesOk

`func (o *AnswerLikes) GetNumLikesOk() (*int32, bool)`

GetNumLikesOk returns a tuple with the NumLikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLikes

`func (o *AnswerLikes) SetNumLikes(v int32)`

SetNumLikes sets NumLikes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


