# FeedbackCustomizations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultChannels** | Pointer to [**[]FeedbackChannel**](FeedbackChannel.md) | The channels to which feedback will be sent for any feature that does not have specific configuration. | [optional] 
**FeatureChannels** | Pointer to [**map[string][]FeedbackChannel**](array.md) | The channels to which feedback will be sent for individual features. The keys of the map will match the values in FeedbackFeature. Features not present in the map should use defaultChannels. | [optional] 
**Disclaimer** | Pointer to **string** | A custom message shown to users during the in-product feedback flow, e.g. to warn users against sending sensitive or personally-identifying information. | [optional] 
**CompanyPrivacyPolicyLink** | Pointer to **string** | An optional link to a privacy policy provided by the users&#39; company that will be shown to them during the in-product feedback flow if their company will receive their feedback. Glean&#39;s policy will also be shown if Glean is receiving the feedback. | [optional] 
**SupportMessage** | Pointer to **string** | User visible text shown when seeking support to guide them to their company&#39;s internal support page when appropriate | [optional] 
**SupportLinkText** | Pointer to **string** | User visible text that will link to the user&#39;s company&#39;s internal support page | [optional] 
**SupportLink** | Pointer to **string** | URL to the user&#39;s company&#39;s internal suport page | [optional] 

## Methods

### NewFeedbackCustomizations

`func NewFeedbackCustomizations() *FeedbackCustomizations`

NewFeedbackCustomizations instantiates a new FeedbackCustomizations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackCustomizationsWithDefaults

`func NewFeedbackCustomizationsWithDefaults() *FeedbackCustomizations`

NewFeedbackCustomizationsWithDefaults instantiates a new FeedbackCustomizations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultChannels

`func (o *FeedbackCustomizations) GetDefaultChannels() []FeedbackChannel`

GetDefaultChannels returns the DefaultChannels field if non-nil, zero value otherwise.

### GetDefaultChannelsOk

`func (o *FeedbackCustomizations) GetDefaultChannelsOk() (*[]FeedbackChannel, bool)`

GetDefaultChannelsOk returns a tuple with the DefaultChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannels

`func (o *FeedbackCustomizations) SetDefaultChannels(v []FeedbackChannel)`

SetDefaultChannels sets DefaultChannels field to given value.

### HasDefaultChannels

`func (o *FeedbackCustomizations) HasDefaultChannels() bool`

HasDefaultChannels returns a boolean if a field has been set.

### GetFeatureChannels

`func (o *FeedbackCustomizations) GetFeatureChannels() map[string][]FeedbackChannel`

GetFeatureChannels returns the FeatureChannels field if non-nil, zero value otherwise.

### GetFeatureChannelsOk

`func (o *FeedbackCustomizations) GetFeatureChannelsOk() (*map[string][]FeedbackChannel, bool)`

GetFeatureChannelsOk returns a tuple with the FeatureChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureChannels

`func (o *FeedbackCustomizations) SetFeatureChannels(v map[string][]FeedbackChannel)`

SetFeatureChannels sets FeatureChannels field to given value.

### HasFeatureChannels

`func (o *FeedbackCustomizations) HasFeatureChannels() bool`

HasFeatureChannels returns a boolean if a field has been set.

### GetDisclaimer

`func (o *FeedbackCustomizations) GetDisclaimer() string`

GetDisclaimer returns the Disclaimer field if non-nil, zero value otherwise.

### GetDisclaimerOk

`func (o *FeedbackCustomizations) GetDisclaimerOk() (*string, bool)`

GetDisclaimerOk returns a tuple with the Disclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclaimer

`func (o *FeedbackCustomizations) SetDisclaimer(v string)`

SetDisclaimer sets Disclaimer field to given value.

### HasDisclaimer

`func (o *FeedbackCustomizations) HasDisclaimer() bool`

HasDisclaimer returns a boolean if a field has been set.

### GetCompanyPrivacyPolicyLink

`func (o *FeedbackCustomizations) GetCompanyPrivacyPolicyLink() string`

GetCompanyPrivacyPolicyLink returns the CompanyPrivacyPolicyLink field if non-nil, zero value otherwise.

### GetCompanyPrivacyPolicyLinkOk

`func (o *FeedbackCustomizations) GetCompanyPrivacyPolicyLinkOk() (*string, bool)`

GetCompanyPrivacyPolicyLinkOk returns a tuple with the CompanyPrivacyPolicyLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrivacyPolicyLink

`func (o *FeedbackCustomizations) SetCompanyPrivacyPolicyLink(v string)`

SetCompanyPrivacyPolicyLink sets CompanyPrivacyPolicyLink field to given value.

### HasCompanyPrivacyPolicyLink

`func (o *FeedbackCustomizations) HasCompanyPrivacyPolicyLink() bool`

HasCompanyPrivacyPolicyLink returns a boolean if a field has been set.

### GetSupportMessage

`func (o *FeedbackCustomizations) GetSupportMessage() string`

GetSupportMessage returns the SupportMessage field if non-nil, zero value otherwise.

### GetSupportMessageOk

`func (o *FeedbackCustomizations) GetSupportMessageOk() (*string, bool)`

GetSupportMessageOk returns a tuple with the SupportMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMessage

`func (o *FeedbackCustomizations) SetSupportMessage(v string)`

SetSupportMessage sets SupportMessage field to given value.

### HasSupportMessage

`func (o *FeedbackCustomizations) HasSupportMessage() bool`

HasSupportMessage returns a boolean if a field has been set.

### GetSupportLinkText

`func (o *FeedbackCustomizations) GetSupportLinkText() string`

GetSupportLinkText returns the SupportLinkText field if non-nil, zero value otherwise.

### GetSupportLinkTextOk

`func (o *FeedbackCustomizations) GetSupportLinkTextOk() (*string, bool)`

GetSupportLinkTextOk returns a tuple with the SupportLinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLinkText

`func (o *FeedbackCustomizations) SetSupportLinkText(v string)`

SetSupportLinkText sets SupportLinkText field to given value.

### HasSupportLinkText

`func (o *FeedbackCustomizations) HasSupportLinkText() bool`

HasSupportLinkText returns a boolean if a field has been set.

### GetSupportLink

`func (o *FeedbackCustomizations) GetSupportLink() string`

GetSupportLink returns the SupportLink field if non-nil, zero value otherwise.

### GetSupportLinkOk

`func (o *FeedbackCustomizations) GetSupportLinkOk() (*string, bool)`

GetSupportLinkOk returns a tuple with the SupportLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLink

`func (o *FeedbackCustomizations) SetSupportLink(v string)`

SetSupportLink sets SupportLink field to given value.

### HasSupportLink

`func (o *FeedbackCustomizations) HasSupportLink() bool`

HasSupportLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


