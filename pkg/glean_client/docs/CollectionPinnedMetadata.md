# CollectionPinnedMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingPins** | Pointer to [**[]CollectionPinTarget**](CollectionPinTarget.md) | List of targets this Collection is pinned to. | [optional] 
**EligiblePins** | Pointer to [**[]CollectionPinMetadata**](CollectionPinMetadata.md) | List of targets this Collection can be pinned to, excluding the targets this Collection is already pinned to. We also include Collection ID already is pinned to each eligible target, which will be 0 if the target has no pinned Collection. | [optional] 

## Methods

### NewCollectionPinnedMetadata

`func NewCollectionPinnedMetadata() *CollectionPinnedMetadata`

NewCollectionPinnedMetadata instantiates a new CollectionPinnedMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionPinnedMetadataWithDefaults

`func NewCollectionPinnedMetadataWithDefaults() *CollectionPinnedMetadata`

NewCollectionPinnedMetadataWithDefaults instantiates a new CollectionPinnedMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingPins

`func (o *CollectionPinnedMetadata) GetExistingPins() []CollectionPinTarget`

GetExistingPins returns the ExistingPins field if non-nil, zero value otherwise.

### GetExistingPinsOk

`func (o *CollectionPinnedMetadata) GetExistingPinsOk() (*[]CollectionPinTarget, bool)`

GetExistingPinsOk returns a tuple with the ExistingPins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingPins

`func (o *CollectionPinnedMetadata) SetExistingPins(v []CollectionPinTarget)`

SetExistingPins sets ExistingPins field to given value.

### HasExistingPins

`func (o *CollectionPinnedMetadata) HasExistingPins() bool`

HasExistingPins returns a boolean if a field has been set.

### GetEligiblePins

`func (o *CollectionPinnedMetadata) GetEligiblePins() []CollectionPinMetadata`

GetEligiblePins returns the EligiblePins field if non-nil, zero value otherwise.

### GetEligiblePinsOk

`func (o *CollectionPinnedMetadata) GetEligiblePinsOk() (*[]CollectionPinMetadata, bool)`

GetEligiblePinsOk returns a tuple with the EligiblePins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligiblePins

`func (o *CollectionPinnedMetadata) SetEligiblePins(v []CollectionPinMetadata)`

SetEligiblePins sets EligiblePins field to given value.

### HasEligiblePins

`func (o *CollectionPinnedMetadata) HasEligiblePins() bool`

HasEligiblePins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


