# SocialNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Possible values are \&quot;twitter\&quot;, \&quot;linkedin\&quot;. | 
**ProfileName** | Pointer to **string** | Human-readable profile name. | [optional] 
**ProfileUrl** | **string** | Link to profile. | 

## Methods

### NewSocialNetwork

`func NewSocialNetwork(name string, profileUrl string, ) *SocialNetwork`

NewSocialNetwork instantiates a new SocialNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialNetworkWithDefaults

`func NewSocialNetworkWithDefaults() *SocialNetwork`

NewSocialNetworkWithDefaults instantiates a new SocialNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SocialNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SocialNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SocialNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetProfileName

`func (o *SocialNetwork) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *SocialNetwork) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *SocialNetwork) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *SocialNetwork) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetProfileUrl

`func (o *SocialNetwork) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *SocialNetwork) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *SocialNetwork) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


