# TeamEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | An email address | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "OTHER"]
**IsUserGenerated** | Pointer to **bool** | true iff the email was manually added by a user from within Glean (aka not from the original data source) | [optional] 

## Methods

### NewTeamEmail

`func NewTeamEmail() *TeamEmail`

NewTeamEmail instantiates a new TeamEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamEmailWithDefaults

`func NewTeamEmailWithDefaults() *TeamEmail`

NewTeamEmailWithDefaults instantiates a new TeamEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *TeamEmail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TeamEmail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TeamEmail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TeamEmail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetType

`func (o *TeamEmail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TeamEmail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TeamEmail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TeamEmail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsUserGenerated

`func (o *TeamEmail) GetIsUserGenerated() bool`

GetIsUserGenerated returns the IsUserGenerated field if non-nil, zero value otherwise.

### GetIsUserGeneratedOk

`func (o *TeamEmail) GetIsUserGeneratedOk() (*bool, bool)`

GetIsUserGeneratedOk returns a tuple with the IsUserGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserGenerated

`func (o *TeamEmail) SetIsUserGenerated(v bool)`

SetIsUserGenerated sets IsUserGenerated field to given value.

### HasIsUserGenerated

`func (o *TeamEmail) HasIsUserGenerated() bool`

HasIsUserGenerated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


