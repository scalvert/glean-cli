# PersonMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** | The first name of the person | [optional] 
**LastName** | Pointer to **string** | The last name of the person | [optional] 
**Title** | Pointer to **string** | Job title. | [optional] 
**BusinessUnit** | Pointer to **string** | Typically the highest level organizational unit; generally applies to bigger companies with multiple distinct businesses. | [optional] 
**Department** | Pointer to **string** | An organizational unit where everyone has a similar task, e.g. &#x60;Engineering&#x60;. | [optional] 
**Teams** | Pointer to [**[]PersonTeam**](PersonTeam.md) | Info about the employee&#39;s team(s). | [optional] 
**DepartmentCount** | Pointer to **int32** | The number of people in this person&#39;s department. | [optional] 
**Email** | Pointer to **string** | The user&#39;s primary email address | [optional] 
**AliasEmails** | Pointer to **[]string** | Additional email addresses of this user beyond the primary, if any. | [optional] 
**Location** | Pointer to **string** | User facing string representing the person&#39;s location. | [optional] 
**StructuredLocation** | Pointer to [**StructuredLocation**](StructuredLocation.md) |  | [optional] 
**ExternalProfileLink** | Pointer to **string** | Link to a customer&#39;s internal profile page. This is set to &#39;#&#39; when no link is desired. | [optional] 
**Manager** | Pointer to [**Person**](Person.md) |  | [optional] 
**ManagementChain** | Pointer to [**[]Person**](Person.md) | The chain of reporting in the company as far up as it goes. The last entry is this person&#39;s direct manager. | [optional] 
**Phone** | Pointer to **string** | Phone number as a number string. | [optional] 
**PhotoUrl** | Pointer to **string** | The URL of the person&#39;s avatar. Public, glean-authenticated and Base64 encoded data URLs are all valid (but not third-party-authenticated URLs). | [optional] 
**UneditedPhotoUrl** | Pointer to **string** | The original photo URL of the person&#39;s avatar before any edits they made are applied | [optional] 
**BannerUrl** | Pointer to **string** | The URL of the person&#39;s banner photo. | [optional] 
**Reports** | Pointer to [**[]Person**](Person.md) |  | [optional] 
**StartDate** | Pointer to **string** | The date when the employee started. | [optional] 
**EndDate** | Pointer to **string** | If a former employee, the last date of employment. | [optional] 
**Bio** | Pointer to **string** | Short biography or mission statement of the employee. | [optional] 
**Pronoun** | Pointer to **string** | She/her, He/his or other pronoun. | [optional] 
**OrgSizeCount** | Pointer to **int32** | The total recursive size of the people reporting to this person, or 1 | [optional] 
**DirectReportsCount** | Pointer to **int32** | The total number of people who directly report to this person, or 0 | [optional] 
**PreferredName** | Pointer to **string** | The preferred name of the person, or a nickname. | [optional] 
**SocialNetwork** | Pointer to [**[]SocialNetwork**](SocialNetwork.md) | List of social network profiles. | [optional] 
**DatasourceProfile** | Pointer to [**[]DatasourceProfile**](DatasourceProfile.md) | List of profiles this user has in different datasources / tools that they use. | [optional] 
**QuerySuggestions** | Pointer to [**QuerySuggestionList**](QuerySuggestionList.md) |  | [optional] 
**PeopleDistance** | Pointer to [**[]PersonDistance**](PersonDistance.md) | List of people and distances to those people from this person. Optionally with metadata. | [optional] 
**InviteInfo** | Pointer to [**InviteInfo**](InviteInfo.md) |  | [optional] 
**IsSignedUp** | Pointer to **bool** | Whether the user has signed into Glean at least once. | [optional] 
**LastExtensionUse** | Pointer to **time.Time** | The last time the user has used the Glean extension in ISO 8601 format. | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldData**](CustomFieldData.md) | User customizable fields for additional people information. | [optional] 
**LoggingId** | Pointer to **string** | The logging id of the person used in scrubbed logs, tracking GA metrics. | [optional] 
**StartDatePercentile** | Pointer to **float32** | Percentage of the company that started strictly after this person. Between [0,100). | [optional] 
**BusyEvents** | Pointer to [**[]AnonymousEvent**](AnonymousEvent.md) | Intervals of busy time for this person, along with the type of event they&#39;re busy with. | [optional] 
**ProfileBoolSettings** | Pointer to **map[string]bool** | flag settings to indicate user profile settings for certain items | [optional] 
**Badges** | Pointer to [**[]Badge**](Badge.md) | The badges that a user has earned over their lifetime. | [optional] 
**IsOrgRoot** | Pointer to **bool** | Whether this person is a \&quot;root\&quot; node in their organization&#39;s hierarchy. | [optional] 

## Methods

### NewPersonMetadata

`func NewPersonMetadata() *PersonMetadata`

NewPersonMetadata instantiates a new PersonMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonMetadataWithDefaults

`func NewPersonMetadataWithDefaults() *PersonMetadata`

NewPersonMetadataWithDefaults instantiates a new PersonMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PersonMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PersonMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PersonMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PersonMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFirstName

`func (o *PersonMetadata) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PersonMetadata) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PersonMetadata) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PersonMetadata) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PersonMetadata) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PersonMetadata) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PersonMetadata) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PersonMetadata) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetTitle

`func (o *PersonMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PersonMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PersonMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PersonMetadata) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBusinessUnit

`func (o *PersonMetadata) GetBusinessUnit() string`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *PersonMetadata) GetBusinessUnitOk() (*string, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *PersonMetadata) SetBusinessUnit(v string)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *PersonMetadata) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetDepartment

`func (o *PersonMetadata) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *PersonMetadata) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *PersonMetadata) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *PersonMetadata) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetTeams

`func (o *PersonMetadata) GetTeams() []PersonTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *PersonMetadata) GetTeamsOk() (*[]PersonTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *PersonMetadata) SetTeams(v []PersonTeam)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *PersonMetadata) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetDepartmentCount

`func (o *PersonMetadata) GetDepartmentCount() int32`

GetDepartmentCount returns the DepartmentCount field if non-nil, zero value otherwise.

### GetDepartmentCountOk

`func (o *PersonMetadata) GetDepartmentCountOk() (*int32, bool)`

GetDepartmentCountOk returns a tuple with the DepartmentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentCount

`func (o *PersonMetadata) SetDepartmentCount(v int32)`

SetDepartmentCount sets DepartmentCount field to given value.

### HasDepartmentCount

`func (o *PersonMetadata) HasDepartmentCount() bool`

HasDepartmentCount returns a boolean if a field has been set.

### GetEmail

`func (o *PersonMetadata) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PersonMetadata) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PersonMetadata) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PersonMetadata) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAliasEmails

`func (o *PersonMetadata) GetAliasEmails() []string`

GetAliasEmails returns the AliasEmails field if non-nil, zero value otherwise.

### GetAliasEmailsOk

`func (o *PersonMetadata) GetAliasEmailsOk() (*[]string, bool)`

GetAliasEmailsOk returns a tuple with the AliasEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasEmails

`func (o *PersonMetadata) SetAliasEmails(v []string)`

SetAliasEmails sets AliasEmails field to given value.

### HasAliasEmails

`func (o *PersonMetadata) HasAliasEmails() bool`

HasAliasEmails returns a boolean if a field has been set.

### GetLocation

`func (o *PersonMetadata) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PersonMetadata) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PersonMetadata) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PersonMetadata) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStructuredLocation

`func (o *PersonMetadata) GetStructuredLocation() StructuredLocation`

GetStructuredLocation returns the StructuredLocation field if non-nil, zero value otherwise.

### GetStructuredLocationOk

`func (o *PersonMetadata) GetStructuredLocationOk() (*StructuredLocation, bool)`

GetStructuredLocationOk returns a tuple with the StructuredLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredLocation

`func (o *PersonMetadata) SetStructuredLocation(v StructuredLocation)`

SetStructuredLocation sets StructuredLocation field to given value.

### HasStructuredLocation

`func (o *PersonMetadata) HasStructuredLocation() bool`

HasStructuredLocation returns a boolean if a field has been set.

### GetExternalProfileLink

`func (o *PersonMetadata) GetExternalProfileLink() string`

GetExternalProfileLink returns the ExternalProfileLink field if non-nil, zero value otherwise.

### GetExternalProfileLinkOk

`func (o *PersonMetadata) GetExternalProfileLinkOk() (*string, bool)`

GetExternalProfileLinkOk returns a tuple with the ExternalProfileLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProfileLink

`func (o *PersonMetadata) SetExternalProfileLink(v string)`

SetExternalProfileLink sets ExternalProfileLink field to given value.

### HasExternalProfileLink

`func (o *PersonMetadata) HasExternalProfileLink() bool`

HasExternalProfileLink returns a boolean if a field has been set.

### GetManager

`func (o *PersonMetadata) GetManager() Person`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *PersonMetadata) GetManagerOk() (*Person, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *PersonMetadata) SetManager(v Person)`

SetManager sets Manager field to given value.

### HasManager

`func (o *PersonMetadata) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetManagementChain

`func (o *PersonMetadata) GetManagementChain() []Person`

GetManagementChain returns the ManagementChain field if non-nil, zero value otherwise.

### GetManagementChainOk

`func (o *PersonMetadata) GetManagementChainOk() (*[]Person, bool)`

GetManagementChainOk returns a tuple with the ManagementChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementChain

`func (o *PersonMetadata) SetManagementChain(v []Person)`

SetManagementChain sets ManagementChain field to given value.

### HasManagementChain

`func (o *PersonMetadata) HasManagementChain() bool`

HasManagementChain returns a boolean if a field has been set.

### GetPhone

`func (o *PersonMetadata) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PersonMetadata) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PersonMetadata) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PersonMetadata) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhotoUrl

`func (o *PersonMetadata) GetPhotoUrl() string`

GetPhotoUrl returns the PhotoUrl field if non-nil, zero value otherwise.

### GetPhotoUrlOk

`func (o *PersonMetadata) GetPhotoUrlOk() (*string, bool)`

GetPhotoUrlOk returns a tuple with the PhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrl

`func (o *PersonMetadata) SetPhotoUrl(v string)`

SetPhotoUrl sets PhotoUrl field to given value.

### HasPhotoUrl

`func (o *PersonMetadata) HasPhotoUrl() bool`

HasPhotoUrl returns a boolean if a field has been set.

### GetUneditedPhotoUrl

`func (o *PersonMetadata) GetUneditedPhotoUrl() string`

GetUneditedPhotoUrl returns the UneditedPhotoUrl field if non-nil, zero value otherwise.

### GetUneditedPhotoUrlOk

`func (o *PersonMetadata) GetUneditedPhotoUrlOk() (*string, bool)`

GetUneditedPhotoUrlOk returns a tuple with the UneditedPhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUneditedPhotoUrl

`func (o *PersonMetadata) SetUneditedPhotoUrl(v string)`

SetUneditedPhotoUrl sets UneditedPhotoUrl field to given value.

### HasUneditedPhotoUrl

`func (o *PersonMetadata) HasUneditedPhotoUrl() bool`

HasUneditedPhotoUrl returns a boolean if a field has been set.

### GetBannerUrl

`func (o *PersonMetadata) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *PersonMetadata) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *PersonMetadata) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *PersonMetadata) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### GetReports

`func (o *PersonMetadata) GetReports() []Person`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *PersonMetadata) GetReportsOk() (*[]Person, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *PersonMetadata) SetReports(v []Person)`

SetReports sets Reports field to given value.

### HasReports

`func (o *PersonMetadata) HasReports() bool`

HasReports returns a boolean if a field has been set.

### GetStartDate

`func (o *PersonMetadata) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PersonMetadata) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PersonMetadata) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PersonMetadata) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PersonMetadata) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PersonMetadata) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PersonMetadata) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PersonMetadata) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetBio

`func (o *PersonMetadata) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *PersonMetadata) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *PersonMetadata) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *PersonMetadata) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetPronoun

`func (o *PersonMetadata) GetPronoun() string`

GetPronoun returns the Pronoun field if non-nil, zero value otherwise.

### GetPronounOk

`func (o *PersonMetadata) GetPronounOk() (*string, bool)`

GetPronounOk returns a tuple with the Pronoun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronoun

`func (o *PersonMetadata) SetPronoun(v string)`

SetPronoun sets Pronoun field to given value.

### HasPronoun

`func (o *PersonMetadata) HasPronoun() bool`

HasPronoun returns a boolean if a field has been set.

### GetOrgSizeCount

`func (o *PersonMetadata) GetOrgSizeCount() int32`

GetOrgSizeCount returns the OrgSizeCount field if non-nil, zero value otherwise.

### GetOrgSizeCountOk

`func (o *PersonMetadata) GetOrgSizeCountOk() (*int32, bool)`

GetOrgSizeCountOk returns a tuple with the OrgSizeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSizeCount

`func (o *PersonMetadata) SetOrgSizeCount(v int32)`

SetOrgSizeCount sets OrgSizeCount field to given value.

### HasOrgSizeCount

`func (o *PersonMetadata) HasOrgSizeCount() bool`

HasOrgSizeCount returns a boolean if a field has been set.

### GetDirectReportsCount

`func (o *PersonMetadata) GetDirectReportsCount() int32`

GetDirectReportsCount returns the DirectReportsCount field if non-nil, zero value otherwise.

### GetDirectReportsCountOk

`func (o *PersonMetadata) GetDirectReportsCountOk() (*int32, bool)`

GetDirectReportsCountOk returns a tuple with the DirectReportsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectReportsCount

`func (o *PersonMetadata) SetDirectReportsCount(v int32)`

SetDirectReportsCount sets DirectReportsCount field to given value.

### HasDirectReportsCount

`func (o *PersonMetadata) HasDirectReportsCount() bool`

HasDirectReportsCount returns a boolean if a field has been set.

### GetPreferredName

`func (o *PersonMetadata) GetPreferredName() string`

GetPreferredName returns the PreferredName field if non-nil, zero value otherwise.

### GetPreferredNameOk

`func (o *PersonMetadata) GetPreferredNameOk() (*string, bool)`

GetPreferredNameOk returns a tuple with the PreferredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredName

`func (o *PersonMetadata) SetPreferredName(v string)`

SetPreferredName sets PreferredName field to given value.

### HasPreferredName

`func (o *PersonMetadata) HasPreferredName() bool`

HasPreferredName returns a boolean if a field has been set.

### GetSocialNetwork

`func (o *PersonMetadata) GetSocialNetwork() []SocialNetwork`

GetSocialNetwork returns the SocialNetwork field if non-nil, zero value otherwise.

### GetSocialNetworkOk

`func (o *PersonMetadata) GetSocialNetworkOk() (*[]SocialNetwork, bool)`

GetSocialNetworkOk returns a tuple with the SocialNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialNetwork

`func (o *PersonMetadata) SetSocialNetwork(v []SocialNetwork)`

SetSocialNetwork sets SocialNetwork field to given value.

### HasSocialNetwork

`func (o *PersonMetadata) HasSocialNetwork() bool`

HasSocialNetwork returns a boolean if a field has been set.

### GetDatasourceProfile

`func (o *PersonMetadata) GetDatasourceProfile() []DatasourceProfile`

GetDatasourceProfile returns the DatasourceProfile field if non-nil, zero value otherwise.

### GetDatasourceProfileOk

`func (o *PersonMetadata) GetDatasourceProfileOk() (*[]DatasourceProfile, bool)`

GetDatasourceProfileOk returns a tuple with the DatasourceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceProfile

`func (o *PersonMetadata) SetDatasourceProfile(v []DatasourceProfile)`

SetDatasourceProfile sets DatasourceProfile field to given value.

### HasDatasourceProfile

`func (o *PersonMetadata) HasDatasourceProfile() bool`

HasDatasourceProfile returns a boolean if a field has been set.

### GetQuerySuggestions

`func (o *PersonMetadata) GetQuerySuggestions() QuerySuggestionList`

GetQuerySuggestions returns the QuerySuggestions field if non-nil, zero value otherwise.

### GetQuerySuggestionsOk

`func (o *PersonMetadata) GetQuerySuggestionsOk() (*QuerySuggestionList, bool)`

GetQuerySuggestionsOk returns a tuple with the QuerySuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySuggestions

`func (o *PersonMetadata) SetQuerySuggestions(v QuerySuggestionList)`

SetQuerySuggestions sets QuerySuggestions field to given value.

### HasQuerySuggestions

`func (o *PersonMetadata) HasQuerySuggestions() bool`

HasQuerySuggestions returns a boolean if a field has been set.

### GetPeopleDistance

`func (o *PersonMetadata) GetPeopleDistance() []PersonDistance`

GetPeopleDistance returns the PeopleDistance field if non-nil, zero value otherwise.

### GetPeopleDistanceOk

`func (o *PersonMetadata) GetPeopleDistanceOk() (*[]PersonDistance, bool)`

GetPeopleDistanceOk returns a tuple with the PeopleDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeopleDistance

`func (o *PersonMetadata) SetPeopleDistance(v []PersonDistance)`

SetPeopleDistance sets PeopleDistance field to given value.

### HasPeopleDistance

`func (o *PersonMetadata) HasPeopleDistance() bool`

HasPeopleDistance returns a boolean if a field has been set.

### GetInviteInfo

`func (o *PersonMetadata) GetInviteInfo() InviteInfo`

GetInviteInfo returns the InviteInfo field if non-nil, zero value otherwise.

### GetInviteInfoOk

`func (o *PersonMetadata) GetInviteInfoOk() (*InviteInfo, bool)`

GetInviteInfoOk returns a tuple with the InviteInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteInfo

`func (o *PersonMetadata) SetInviteInfo(v InviteInfo)`

SetInviteInfo sets InviteInfo field to given value.

### HasInviteInfo

`func (o *PersonMetadata) HasInviteInfo() bool`

HasInviteInfo returns a boolean if a field has been set.

### GetIsSignedUp

`func (o *PersonMetadata) GetIsSignedUp() bool`

GetIsSignedUp returns the IsSignedUp field if non-nil, zero value otherwise.

### GetIsSignedUpOk

`func (o *PersonMetadata) GetIsSignedUpOk() (*bool, bool)`

GetIsSignedUpOk returns a tuple with the IsSignedUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSignedUp

`func (o *PersonMetadata) SetIsSignedUp(v bool)`

SetIsSignedUp sets IsSignedUp field to given value.

### HasIsSignedUp

`func (o *PersonMetadata) HasIsSignedUp() bool`

HasIsSignedUp returns a boolean if a field has been set.

### GetLastExtensionUse

`func (o *PersonMetadata) GetLastExtensionUse() time.Time`

GetLastExtensionUse returns the LastExtensionUse field if non-nil, zero value otherwise.

### GetLastExtensionUseOk

`func (o *PersonMetadata) GetLastExtensionUseOk() (*time.Time, bool)`

GetLastExtensionUseOk returns a tuple with the LastExtensionUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExtensionUse

`func (o *PersonMetadata) SetLastExtensionUse(v time.Time)`

SetLastExtensionUse sets LastExtensionUse field to given value.

### HasLastExtensionUse

`func (o *PersonMetadata) HasLastExtensionUse() bool`

HasLastExtensionUse returns a boolean if a field has been set.

### GetPermissions

`func (o *PersonMetadata) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PersonMetadata) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PersonMetadata) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PersonMetadata) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCustomFields

`func (o *PersonMetadata) GetCustomFields() []CustomFieldData`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PersonMetadata) GetCustomFieldsOk() (*[]CustomFieldData, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PersonMetadata) SetCustomFields(v []CustomFieldData)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PersonMetadata) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLoggingId

`func (o *PersonMetadata) GetLoggingId() string`

GetLoggingId returns the LoggingId field if non-nil, zero value otherwise.

### GetLoggingIdOk

`func (o *PersonMetadata) GetLoggingIdOk() (*string, bool)`

GetLoggingIdOk returns a tuple with the LoggingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingId

`func (o *PersonMetadata) SetLoggingId(v string)`

SetLoggingId sets LoggingId field to given value.

### HasLoggingId

`func (o *PersonMetadata) HasLoggingId() bool`

HasLoggingId returns a boolean if a field has been set.

### GetStartDatePercentile

`func (o *PersonMetadata) GetStartDatePercentile() float32`

GetStartDatePercentile returns the StartDatePercentile field if non-nil, zero value otherwise.

### GetStartDatePercentileOk

`func (o *PersonMetadata) GetStartDatePercentileOk() (*float32, bool)`

GetStartDatePercentileOk returns a tuple with the StartDatePercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDatePercentile

`func (o *PersonMetadata) SetStartDatePercentile(v float32)`

SetStartDatePercentile sets StartDatePercentile field to given value.

### HasStartDatePercentile

`func (o *PersonMetadata) HasStartDatePercentile() bool`

HasStartDatePercentile returns a boolean if a field has been set.

### GetBusyEvents

`func (o *PersonMetadata) GetBusyEvents() []AnonymousEvent`

GetBusyEvents returns the BusyEvents field if non-nil, zero value otherwise.

### GetBusyEventsOk

`func (o *PersonMetadata) GetBusyEventsOk() (*[]AnonymousEvent, bool)`

GetBusyEventsOk returns a tuple with the BusyEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyEvents

`func (o *PersonMetadata) SetBusyEvents(v []AnonymousEvent)`

SetBusyEvents sets BusyEvents field to given value.

### HasBusyEvents

`func (o *PersonMetadata) HasBusyEvents() bool`

HasBusyEvents returns a boolean if a field has been set.

### GetProfileBoolSettings

`func (o *PersonMetadata) GetProfileBoolSettings() map[string]bool`

GetProfileBoolSettings returns the ProfileBoolSettings field if non-nil, zero value otherwise.

### GetProfileBoolSettingsOk

`func (o *PersonMetadata) GetProfileBoolSettingsOk() (*map[string]bool, bool)`

GetProfileBoolSettingsOk returns a tuple with the ProfileBoolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileBoolSettings

`func (o *PersonMetadata) SetProfileBoolSettings(v map[string]bool)`

SetProfileBoolSettings sets ProfileBoolSettings field to given value.

### HasProfileBoolSettings

`func (o *PersonMetadata) HasProfileBoolSettings() bool`

HasProfileBoolSettings returns a boolean if a field has been set.

### GetBadges

`func (o *PersonMetadata) GetBadges() []Badge`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *PersonMetadata) GetBadgesOk() (*[]Badge, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *PersonMetadata) SetBadges(v []Badge)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *PersonMetadata) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetIsOrgRoot

`func (o *PersonMetadata) GetIsOrgRoot() bool`

GetIsOrgRoot returns the IsOrgRoot field if non-nil, zero value otherwise.

### GetIsOrgRootOk

`func (o *PersonMetadata) GetIsOrgRootOk() (*bool, bool)`

GetIsOrgRootOk returns a tuple with the IsOrgRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgRoot

`func (o *PersonMetadata) SetIsOrgRoot(v bool)`

SetIsOrgRoot sets IsOrgRoot field to given value.

### HasIsOrgRoot

`func (o *PersonMetadata) HasIsOrgRoot() bool`

HasIsOrgRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


