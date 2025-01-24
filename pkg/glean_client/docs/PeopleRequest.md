# PeopleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimezoneOffset** | Pointer to **int32** | The offset of the client&#39;s timezone in minutes from UTC. e.g. PDT is -420 because it&#39;s 7 hours behind UTC. | [optional] 
**ObfuscatedIds** | Pointer to **[]string** | The Person IDs to retrieve. If no IDs are requested, the current user&#39;s details are returned. | [optional] 
**EmailIds** | Pointer to **[]string** | The email IDs to retrieve. The result is the deduplicated union of emailIds and obfuscatedIds. | [optional] 
**IncludeFields** | Pointer to **[]string** | List of PersonMetadata fields to return (that aren&#39;t returned by default) | [optional] 
**IncludeTypes** | Pointer to **[]string** | The types of people entities to include in the response in addition to those returned by default. | [optional] 
**Source** | Pointer to **string** | A string denoting the search surface from which the endpoint is called. | [optional] 

## Methods

### NewPeopleRequest

`func NewPeopleRequest() *PeopleRequest`

NewPeopleRequest instantiates a new PeopleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeopleRequestWithDefaults

`func NewPeopleRequestWithDefaults() *PeopleRequest`

NewPeopleRequestWithDefaults instantiates a new PeopleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezoneOffset

`func (o *PeopleRequest) GetTimezoneOffset() int32`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *PeopleRequest) GetTimezoneOffsetOk() (*int32, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *PeopleRequest) SetTimezoneOffset(v int32)`

SetTimezoneOffset sets TimezoneOffset field to given value.

### HasTimezoneOffset

`func (o *PeopleRequest) HasTimezoneOffset() bool`

HasTimezoneOffset returns a boolean if a field has been set.

### GetObfuscatedIds

`func (o *PeopleRequest) GetObfuscatedIds() []string`

GetObfuscatedIds returns the ObfuscatedIds field if non-nil, zero value otherwise.

### GetObfuscatedIdsOk

`func (o *PeopleRequest) GetObfuscatedIdsOk() (*[]string, bool)`

GetObfuscatedIdsOk returns a tuple with the ObfuscatedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedIds

`func (o *PeopleRequest) SetObfuscatedIds(v []string)`

SetObfuscatedIds sets ObfuscatedIds field to given value.

### HasObfuscatedIds

`func (o *PeopleRequest) HasObfuscatedIds() bool`

HasObfuscatedIds returns a boolean if a field has been set.

### GetEmailIds

`func (o *PeopleRequest) GetEmailIds() []string`

GetEmailIds returns the EmailIds field if non-nil, zero value otherwise.

### GetEmailIdsOk

`func (o *PeopleRequest) GetEmailIdsOk() (*[]string, bool)`

GetEmailIdsOk returns a tuple with the EmailIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIds

`func (o *PeopleRequest) SetEmailIds(v []string)`

SetEmailIds sets EmailIds field to given value.

### HasEmailIds

`func (o *PeopleRequest) HasEmailIds() bool`

HasEmailIds returns a boolean if a field has been set.

### GetIncludeFields

`func (o *PeopleRequest) GetIncludeFields() []string`

GetIncludeFields returns the IncludeFields field if non-nil, zero value otherwise.

### GetIncludeFieldsOk

`func (o *PeopleRequest) GetIncludeFieldsOk() (*[]string, bool)`

GetIncludeFieldsOk returns a tuple with the IncludeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFields

`func (o *PeopleRequest) SetIncludeFields(v []string)`

SetIncludeFields sets IncludeFields field to given value.

### HasIncludeFields

`func (o *PeopleRequest) HasIncludeFields() bool`

HasIncludeFields returns a boolean if a field has been set.

### GetIncludeTypes

`func (o *PeopleRequest) GetIncludeTypes() []string`

GetIncludeTypes returns the IncludeTypes field if non-nil, zero value otherwise.

### GetIncludeTypesOk

`func (o *PeopleRequest) GetIncludeTypesOk() (*[]string, bool)`

GetIncludeTypesOk returns a tuple with the IncludeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTypes

`func (o *PeopleRequest) SetIncludeTypes(v []string)`

SetIncludeTypes sets IncludeTypes field to given value.

### HasIncludeTypes

`func (o *PeopleRequest) HasIncludeTypes() bool`

HasIncludeTypes returns a boolean if a field has been set.

### GetSource

`func (o *PeopleRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PeopleRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PeopleRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PeopleRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


