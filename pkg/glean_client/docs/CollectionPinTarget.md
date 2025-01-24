# CollectionPinTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**CollectionPinnableCategories**](CollectionPinnableCategories.md) |  | 
**Value** | Pointer to **string** | Optional. If category supports values, then the additional value for the category e.g. department name for DEPARTMENT_RESOURCE, team name/id for TEAM_RESOURCE and so on. | [optional] 
**Target** | Pointer to [**CollectionPinnableTargets**](CollectionPinnableTargets.md) |  | [optional] 

## Methods

### NewCollectionPinTarget

`func NewCollectionPinTarget(category CollectionPinnableCategories, ) *CollectionPinTarget`

NewCollectionPinTarget instantiates a new CollectionPinTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionPinTargetWithDefaults

`func NewCollectionPinTargetWithDefaults() *CollectionPinTarget`

NewCollectionPinTargetWithDefaults instantiates a new CollectionPinTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CollectionPinTarget) GetCategory() CollectionPinnableCategories`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CollectionPinTarget) GetCategoryOk() (*CollectionPinnableCategories, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CollectionPinTarget) SetCategory(v CollectionPinnableCategories)`

SetCategory sets Category field to given value.


### GetValue

`func (o *CollectionPinTarget) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CollectionPinTarget) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CollectionPinTarget) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CollectionPinTarget) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTarget

`func (o *CollectionPinTarget) GetTarget() CollectionPinnableTargets`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CollectionPinTarget) GetTargetOk() (*CollectionPinnableTargets, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CollectionPinTarget) SetTarget(v CollectionPinnableTargets)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CollectionPinTarget) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


