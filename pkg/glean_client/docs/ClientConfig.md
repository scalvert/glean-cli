# ClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assistant** | Pointer to [**AssistantConfig**](AssistantConfig.md) |  | [optional] 
**Tools** | Pointer to [**ToolsConfig**](ToolsConfig.md) |  | [optional] 
**Shortcuts** | Pointer to [**ShortcutsConfig**](ShortcutsConfig.md) |  | [optional] 
**BadVersions** | Pointer to **[]string** | Known bad client versions that should force update themselves | [optional] 
**FeedPeopleCelebrationsEnabled** | Pointer to **bool** | Whether people celebrations is enabled or not for the instance | [optional] 
**FeedSuggestedEnabled** | Pointer to **bool** | Whether the suggested feed is enabled | [optional] 
**FeedTrendingEnabled** | Pointer to **bool** | Whether the trending feed is enabled | [optional] 
**FeedRecentsEnabled** | Pointer to **bool** | Whether the recents feed is enabled | [optional] 
**FeedMentionsEnabled** | Pointer to **bool** | Whether the mentions feed is enabled | [optional] 
**GptAgentEnabled** | Pointer to **bool** | Whether the GPT agent for Chat is enabled | [optional] 
**ChatHistoryEnabled** | Pointer to **bool** | Whether the chat history for Chat is enabled | [optional] 
**BoolValues** | Pointer to **map[string]bool** | A map of {string, boolean} pairs representing flags that globally guard conditional features. Omitted flags mean the client should use its default state. | [optional] 
**IntegerValues** | Pointer to **map[string]int32** | A map of {string, integer} pairs for client consumption. | [optional] 
**CompanyDisplayName** | Pointer to **string** | The user-facing name of the company owning the deployment | [optional] 
**CustomSerpMarkdown** | Pointer to **string** | A markdown string to be displayed on the search results page. Useful for outlinks to help pages. | [optional] 
**OnboardingQuery** | Pointer to **string** | A demonstrative query to show during new user onboarding | [optional] 
**IsOrgChartLinkVisible** | Pointer to **bool** | Determines whether the org chart link in the Directory panel is visible to all users. | [optional] 
**IsOrgChartAccessible** | Pointer to **bool** | Determines whether the org chart is accessible to all users, regardless of link visibility. Org chart can be accessible even if the org chart link in Directory is not visible. | [optional] 
**IsPeopleSetup** | Pointer to **bool** | Whether or not people data has been set up. | [optional] 
**IsPilotMode** | Pointer to **bool** | Whether or not the deployment is in pilot mode. | [optional] 
**WebAppUrl** | Pointer to **string** | URL the company uses to access the web app | [optional] 
**UserOutreach** | Pointer to [**UserOutreachConfig**](UserOutreachConfig.md) |  | [optional] 
**SearchLinkUrlTemplate** | Pointer to **string** | The URL to use for outbound links to Glean Search. Defaults to {webAppUrl}/search?q&#x3D;%s. | [optional] 
**ChatLinkUrlTemplate** | Pointer to **string** | The URL to use for outbound links to Glean Chat. Defaults to {webAppUrl}/chat. | [optional] 
**Themes** | Pointer to [**Themes**](Themes.md) |  | [optional] 
**Brandings** | Pointer to [**ClientConfigBrandings**](ClientConfigBrandings.md) |  | [optional] 
**GreetingFormat** | Pointer to **string** | Describes how to format the web app greeting. Possible format options include \\%t - timely greeting \\%n - the user&#39;s first name | [optional] 
**TaskSeeAllLabel** | Pointer to **string** | Label for the external link at the end of the Task card in order to guide user to the source. | [optional] 
**TaskSeeAllLink** | Pointer to **string** | Link used in conjunction with taskSeeAllLabel to redirect user to the task&#39;s source. | [optional] 
**ShortcutsPrefix** | Pointer to **string** | Company-wide custom prefix for Go Links. | [optional] 
**SsoCompanyProvider** | Pointer to **string** | SSO provider used by the company | [optional] 
**FeedbackCustomizations** | Pointer to [**FeedbackCustomizations**](FeedbackCustomizations.md) |  | [optional] 

## Methods

### NewClientConfig

`func NewClientConfig() *ClientConfig`

NewClientConfig instantiates a new ClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfigWithDefaults

`func NewClientConfigWithDefaults() *ClientConfig`

NewClientConfigWithDefaults instantiates a new ClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistant

`func (o *ClientConfig) GetAssistant() AssistantConfig`

GetAssistant returns the Assistant field if non-nil, zero value otherwise.

### GetAssistantOk

`func (o *ClientConfig) GetAssistantOk() (*AssistantConfig, bool)`

GetAssistantOk returns a tuple with the Assistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistant

`func (o *ClientConfig) SetAssistant(v AssistantConfig)`

SetAssistant sets Assistant field to given value.

### HasAssistant

`func (o *ClientConfig) HasAssistant() bool`

HasAssistant returns a boolean if a field has been set.

### GetTools

`func (o *ClientConfig) GetTools() ToolsConfig`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ClientConfig) GetToolsOk() (*ToolsConfig, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ClientConfig) SetTools(v ToolsConfig)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ClientConfig) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetShortcuts

`func (o *ClientConfig) GetShortcuts() ShortcutsConfig`

GetShortcuts returns the Shortcuts field if non-nil, zero value otherwise.

### GetShortcutsOk

`func (o *ClientConfig) GetShortcutsOk() (*ShortcutsConfig, bool)`

GetShortcutsOk returns a tuple with the Shortcuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcuts

`func (o *ClientConfig) SetShortcuts(v ShortcutsConfig)`

SetShortcuts sets Shortcuts field to given value.

### HasShortcuts

`func (o *ClientConfig) HasShortcuts() bool`

HasShortcuts returns a boolean if a field has been set.

### GetBadVersions

`func (o *ClientConfig) GetBadVersions() []string`

GetBadVersions returns the BadVersions field if non-nil, zero value otherwise.

### GetBadVersionsOk

`func (o *ClientConfig) GetBadVersionsOk() (*[]string, bool)`

GetBadVersionsOk returns a tuple with the BadVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadVersions

`func (o *ClientConfig) SetBadVersions(v []string)`

SetBadVersions sets BadVersions field to given value.

### HasBadVersions

`func (o *ClientConfig) HasBadVersions() bool`

HasBadVersions returns a boolean if a field has been set.

### GetFeedPeopleCelebrationsEnabled

`func (o *ClientConfig) GetFeedPeopleCelebrationsEnabled() bool`

GetFeedPeopleCelebrationsEnabled returns the FeedPeopleCelebrationsEnabled field if non-nil, zero value otherwise.

### GetFeedPeopleCelebrationsEnabledOk

`func (o *ClientConfig) GetFeedPeopleCelebrationsEnabledOk() (*bool, bool)`

GetFeedPeopleCelebrationsEnabledOk returns a tuple with the FeedPeopleCelebrationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedPeopleCelebrationsEnabled

`func (o *ClientConfig) SetFeedPeopleCelebrationsEnabled(v bool)`

SetFeedPeopleCelebrationsEnabled sets FeedPeopleCelebrationsEnabled field to given value.

### HasFeedPeopleCelebrationsEnabled

`func (o *ClientConfig) HasFeedPeopleCelebrationsEnabled() bool`

HasFeedPeopleCelebrationsEnabled returns a boolean if a field has been set.

### GetFeedSuggestedEnabled

`func (o *ClientConfig) GetFeedSuggestedEnabled() bool`

GetFeedSuggestedEnabled returns the FeedSuggestedEnabled field if non-nil, zero value otherwise.

### GetFeedSuggestedEnabledOk

`func (o *ClientConfig) GetFeedSuggestedEnabledOk() (*bool, bool)`

GetFeedSuggestedEnabledOk returns a tuple with the FeedSuggestedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedSuggestedEnabled

`func (o *ClientConfig) SetFeedSuggestedEnabled(v bool)`

SetFeedSuggestedEnabled sets FeedSuggestedEnabled field to given value.

### HasFeedSuggestedEnabled

`func (o *ClientConfig) HasFeedSuggestedEnabled() bool`

HasFeedSuggestedEnabled returns a boolean if a field has been set.

### GetFeedTrendingEnabled

`func (o *ClientConfig) GetFeedTrendingEnabled() bool`

GetFeedTrendingEnabled returns the FeedTrendingEnabled field if non-nil, zero value otherwise.

### GetFeedTrendingEnabledOk

`func (o *ClientConfig) GetFeedTrendingEnabledOk() (*bool, bool)`

GetFeedTrendingEnabledOk returns a tuple with the FeedTrendingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedTrendingEnabled

`func (o *ClientConfig) SetFeedTrendingEnabled(v bool)`

SetFeedTrendingEnabled sets FeedTrendingEnabled field to given value.

### HasFeedTrendingEnabled

`func (o *ClientConfig) HasFeedTrendingEnabled() bool`

HasFeedTrendingEnabled returns a boolean if a field has been set.

### GetFeedRecentsEnabled

`func (o *ClientConfig) GetFeedRecentsEnabled() bool`

GetFeedRecentsEnabled returns the FeedRecentsEnabled field if non-nil, zero value otherwise.

### GetFeedRecentsEnabledOk

`func (o *ClientConfig) GetFeedRecentsEnabledOk() (*bool, bool)`

GetFeedRecentsEnabledOk returns a tuple with the FeedRecentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedRecentsEnabled

`func (o *ClientConfig) SetFeedRecentsEnabled(v bool)`

SetFeedRecentsEnabled sets FeedRecentsEnabled field to given value.

### HasFeedRecentsEnabled

`func (o *ClientConfig) HasFeedRecentsEnabled() bool`

HasFeedRecentsEnabled returns a boolean if a field has been set.

### GetFeedMentionsEnabled

`func (o *ClientConfig) GetFeedMentionsEnabled() bool`

GetFeedMentionsEnabled returns the FeedMentionsEnabled field if non-nil, zero value otherwise.

### GetFeedMentionsEnabledOk

`func (o *ClientConfig) GetFeedMentionsEnabledOk() (*bool, bool)`

GetFeedMentionsEnabledOk returns a tuple with the FeedMentionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedMentionsEnabled

`func (o *ClientConfig) SetFeedMentionsEnabled(v bool)`

SetFeedMentionsEnabled sets FeedMentionsEnabled field to given value.

### HasFeedMentionsEnabled

`func (o *ClientConfig) HasFeedMentionsEnabled() bool`

HasFeedMentionsEnabled returns a boolean if a field has been set.

### GetGptAgentEnabled

`func (o *ClientConfig) GetGptAgentEnabled() bool`

GetGptAgentEnabled returns the GptAgentEnabled field if non-nil, zero value otherwise.

### GetGptAgentEnabledOk

`func (o *ClientConfig) GetGptAgentEnabledOk() (*bool, bool)`

GetGptAgentEnabledOk returns a tuple with the GptAgentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGptAgentEnabled

`func (o *ClientConfig) SetGptAgentEnabled(v bool)`

SetGptAgentEnabled sets GptAgentEnabled field to given value.

### HasGptAgentEnabled

`func (o *ClientConfig) HasGptAgentEnabled() bool`

HasGptAgentEnabled returns a boolean if a field has been set.

### GetChatHistoryEnabled

`func (o *ClientConfig) GetChatHistoryEnabled() bool`

GetChatHistoryEnabled returns the ChatHistoryEnabled field if non-nil, zero value otherwise.

### GetChatHistoryEnabledOk

`func (o *ClientConfig) GetChatHistoryEnabledOk() (*bool, bool)`

GetChatHistoryEnabledOk returns a tuple with the ChatHistoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatHistoryEnabled

`func (o *ClientConfig) SetChatHistoryEnabled(v bool)`

SetChatHistoryEnabled sets ChatHistoryEnabled field to given value.

### HasChatHistoryEnabled

`func (o *ClientConfig) HasChatHistoryEnabled() bool`

HasChatHistoryEnabled returns a boolean if a field has been set.

### GetBoolValues

`func (o *ClientConfig) GetBoolValues() map[string]bool`

GetBoolValues returns the BoolValues field if non-nil, zero value otherwise.

### GetBoolValuesOk

`func (o *ClientConfig) GetBoolValuesOk() (*map[string]bool, bool)`

GetBoolValuesOk returns a tuple with the BoolValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValues

`func (o *ClientConfig) SetBoolValues(v map[string]bool)`

SetBoolValues sets BoolValues field to given value.

### HasBoolValues

`func (o *ClientConfig) HasBoolValues() bool`

HasBoolValues returns a boolean if a field has been set.

### GetIntegerValues

`func (o *ClientConfig) GetIntegerValues() map[string]int32`

GetIntegerValues returns the IntegerValues field if non-nil, zero value otherwise.

### GetIntegerValuesOk

`func (o *ClientConfig) GetIntegerValuesOk() (*map[string]int32, bool)`

GetIntegerValuesOk returns a tuple with the IntegerValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegerValues

`func (o *ClientConfig) SetIntegerValues(v map[string]int32)`

SetIntegerValues sets IntegerValues field to given value.

### HasIntegerValues

`func (o *ClientConfig) HasIntegerValues() bool`

HasIntegerValues returns a boolean if a field has been set.

### GetCompanyDisplayName

`func (o *ClientConfig) GetCompanyDisplayName() string`

GetCompanyDisplayName returns the CompanyDisplayName field if non-nil, zero value otherwise.

### GetCompanyDisplayNameOk

`func (o *ClientConfig) GetCompanyDisplayNameOk() (*string, bool)`

GetCompanyDisplayNameOk returns a tuple with the CompanyDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDisplayName

`func (o *ClientConfig) SetCompanyDisplayName(v string)`

SetCompanyDisplayName sets CompanyDisplayName field to given value.

### HasCompanyDisplayName

`func (o *ClientConfig) HasCompanyDisplayName() bool`

HasCompanyDisplayName returns a boolean if a field has been set.

### GetCustomSerpMarkdown

`func (o *ClientConfig) GetCustomSerpMarkdown() string`

GetCustomSerpMarkdown returns the CustomSerpMarkdown field if non-nil, zero value otherwise.

### GetCustomSerpMarkdownOk

`func (o *ClientConfig) GetCustomSerpMarkdownOk() (*string, bool)`

GetCustomSerpMarkdownOk returns a tuple with the CustomSerpMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSerpMarkdown

`func (o *ClientConfig) SetCustomSerpMarkdown(v string)`

SetCustomSerpMarkdown sets CustomSerpMarkdown field to given value.

### HasCustomSerpMarkdown

`func (o *ClientConfig) HasCustomSerpMarkdown() bool`

HasCustomSerpMarkdown returns a boolean if a field has been set.

### GetOnboardingQuery

`func (o *ClientConfig) GetOnboardingQuery() string`

GetOnboardingQuery returns the OnboardingQuery field if non-nil, zero value otherwise.

### GetOnboardingQueryOk

`func (o *ClientConfig) GetOnboardingQueryOk() (*string, bool)`

GetOnboardingQueryOk returns a tuple with the OnboardingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingQuery

`func (o *ClientConfig) SetOnboardingQuery(v string)`

SetOnboardingQuery sets OnboardingQuery field to given value.

### HasOnboardingQuery

`func (o *ClientConfig) HasOnboardingQuery() bool`

HasOnboardingQuery returns a boolean if a field has been set.

### GetIsOrgChartLinkVisible

`func (o *ClientConfig) GetIsOrgChartLinkVisible() bool`

GetIsOrgChartLinkVisible returns the IsOrgChartLinkVisible field if non-nil, zero value otherwise.

### GetIsOrgChartLinkVisibleOk

`func (o *ClientConfig) GetIsOrgChartLinkVisibleOk() (*bool, bool)`

GetIsOrgChartLinkVisibleOk returns a tuple with the IsOrgChartLinkVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgChartLinkVisible

`func (o *ClientConfig) SetIsOrgChartLinkVisible(v bool)`

SetIsOrgChartLinkVisible sets IsOrgChartLinkVisible field to given value.

### HasIsOrgChartLinkVisible

`func (o *ClientConfig) HasIsOrgChartLinkVisible() bool`

HasIsOrgChartLinkVisible returns a boolean if a field has been set.

### GetIsOrgChartAccessible

`func (o *ClientConfig) GetIsOrgChartAccessible() bool`

GetIsOrgChartAccessible returns the IsOrgChartAccessible field if non-nil, zero value otherwise.

### GetIsOrgChartAccessibleOk

`func (o *ClientConfig) GetIsOrgChartAccessibleOk() (*bool, bool)`

GetIsOrgChartAccessibleOk returns a tuple with the IsOrgChartAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrgChartAccessible

`func (o *ClientConfig) SetIsOrgChartAccessible(v bool)`

SetIsOrgChartAccessible sets IsOrgChartAccessible field to given value.

### HasIsOrgChartAccessible

`func (o *ClientConfig) HasIsOrgChartAccessible() bool`

HasIsOrgChartAccessible returns a boolean if a field has been set.

### GetIsPeopleSetup

`func (o *ClientConfig) GetIsPeopleSetup() bool`

GetIsPeopleSetup returns the IsPeopleSetup field if non-nil, zero value otherwise.

### GetIsPeopleSetupOk

`func (o *ClientConfig) GetIsPeopleSetupOk() (*bool, bool)`

GetIsPeopleSetupOk returns a tuple with the IsPeopleSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPeopleSetup

`func (o *ClientConfig) SetIsPeopleSetup(v bool)`

SetIsPeopleSetup sets IsPeopleSetup field to given value.

### HasIsPeopleSetup

`func (o *ClientConfig) HasIsPeopleSetup() bool`

HasIsPeopleSetup returns a boolean if a field has been set.

### GetIsPilotMode

`func (o *ClientConfig) GetIsPilotMode() bool`

GetIsPilotMode returns the IsPilotMode field if non-nil, zero value otherwise.

### GetIsPilotModeOk

`func (o *ClientConfig) GetIsPilotModeOk() (*bool, bool)`

GetIsPilotModeOk returns a tuple with the IsPilotMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPilotMode

`func (o *ClientConfig) SetIsPilotMode(v bool)`

SetIsPilotMode sets IsPilotMode field to given value.

### HasIsPilotMode

`func (o *ClientConfig) HasIsPilotMode() bool`

HasIsPilotMode returns a boolean if a field has been set.

### GetWebAppUrl

`func (o *ClientConfig) GetWebAppUrl() string`

GetWebAppUrl returns the WebAppUrl field if non-nil, zero value otherwise.

### GetWebAppUrlOk

`func (o *ClientConfig) GetWebAppUrlOk() (*string, bool)`

GetWebAppUrlOk returns a tuple with the WebAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAppUrl

`func (o *ClientConfig) SetWebAppUrl(v string)`

SetWebAppUrl sets WebAppUrl field to given value.

### HasWebAppUrl

`func (o *ClientConfig) HasWebAppUrl() bool`

HasWebAppUrl returns a boolean if a field has been set.

### GetUserOutreach

`func (o *ClientConfig) GetUserOutreach() UserOutreachConfig`

GetUserOutreach returns the UserOutreach field if non-nil, zero value otherwise.

### GetUserOutreachOk

`func (o *ClientConfig) GetUserOutreachOk() (*UserOutreachConfig, bool)`

GetUserOutreachOk returns a tuple with the UserOutreach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOutreach

`func (o *ClientConfig) SetUserOutreach(v UserOutreachConfig)`

SetUserOutreach sets UserOutreach field to given value.

### HasUserOutreach

`func (o *ClientConfig) HasUserOutreach() bool`

HasUserOutreach returns a boolean if a field has been set.

### GetSearchLinkUrlTemplate

`func (o *ClientConfig) GetSearchLinkUrlTemplate() string`

GetSearchLinkUrlTemplate returns the SearchLinkUrlTemplate field if non-nil, zero value otherwise.

### GetSearchLinkUrlTemplateOk

`func (o *ClientConfig) GetSearchLinkUrlTemplateOk() (*string, bool)`

GetSearchLinkUrlTemplateOk returns a tuple with the SearchLinkUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchLinkUrlTemplate

`func (o *ClientConfig) SetSearchLinkUrlTemplate(v string)`

SetSearchLinkUrlTemplate sets SearchLinkUrlTemplate field to given value.

### HasSearchLinkUrlTemplate

`func (o *ClientConfig) HasSearchLinkUrlTemplate() bool`

HasSearchLinkUrlTemplate returns a boolean if a field has been set.

### GetChatLinkUrlTemplate

`func (o *ClientConfig) GetChatLinkUrlTemplate() string`

GetChatLinkUrlTemplate returns the ChatLinkUrlTemplate field if non-nil, zero value otherwise.

### GetChatLinkUrlTemplateOk

`func (o *ClientConfig) GetChatLinkUrlTemplateOk() (*string, bool)`

GetChatLinkUrlTemplateOk returns a tuple with the ChatLinkUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatLinkUrlTemplate

`func (o *ClientConfig) SetChatLinkUrlTemplate(v string)`

SetChatLinkUrlTemplate sets ChatLinkUrlTemplate field to given value.

### HasChatLinkUrlTemplate

`func (o *ClientConfig) HasChatLinkUrlTemplate() bool`

HasChatLinkUrlTemplate returns a boolean if a field has been set.

### GetThemes

`func (o *ClientConfig) GetThemes() Themes`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *ClientConfig) GetThemesOk() (*Themes, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *ClientConfig) SetThemes(v Themes)`

SetThemes sets Themes field to given value.

### HasThemes

`func (o *ClientConfig) HasThemes() bool`

HasThemes returns a boolean if a field has been set.

### GetBrandings

`func (o *ClientConfig) GetBrandings() ClientConfigBrandings`

GetBrandings returns the Brandings field if non-nil, zero value otherwise.

### GetBrandingsOk

`func (o *ClientConfig) GetBrandingsOk() (*ClientConfigBrandings, bool)`

GetBrandingsOk returns a tuple with the Brandings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandings

`func (o *ClientConfig) SetBrandings(v ClientConfigBrandings)`

SetBrandings sets Brandings field to given value.

### HasBrandings

`func (o *ClientConfig) HasBrandings() bool`

HasBrandings returns a boolean if a field has been set.

### GetGreetingFormat

`func (o *ClientConfig) GetGreetingFormat() string`

GetGreetingFormat returns the GreetingFormat field if non-nil, zero value otherwise.

### GetGreetingFormatOk

`func (o *ClientConfig) GetGreetingFormatOk() (*string, bool)`

GetGreetingFormatOk returns a tuple with the GreetingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingFormat

`func (o *ClientConfig) SetGreetingFormat(v string)`

SetGreetingFormat sets GreetingFormat field to given value.

### HasGreetingFormat

`func (o *ClientConfig) HasGreetingFormat() bool`

HasGreetingFormat returns a boolean if a field has been set.

### GetTaskSeeAllLabel

`func (o *ClientConfig) GetTaskSeeAllLabel() string`

GetTaskSeeAllLabel returns the TaskSeeAllLabel field if non-nil, zero value otherwise.

### GetTaskSeeAllLabelOk

`func (o *ClientConfig) GetTaskSeeAllLabelOk() (*string, bool)`

GetTaskSeeAllLabelOk returns a tuple with the TaskSeeAllLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSeeAllLabel

`func (o *ClientConfig) SetTaskSeeAllLabel(v string)`

SetTaskSeeAllLabel sets TaskSeeAllLabel field to given value.

### HasTaskSeeAllLabel

`func (o *ClientConfig) HasTaskSeeAllLabel() bool`

HasTaskSeeAllLabel returns a boolean if a field has been set.

### GetTaskSeeAllLink

`func (o *ClientConfig) GetTaskSeeAllLink() string`

GetTaskSeeAllLink returns the TaskSeeAllLink field if non-nil, zero value otherwise.

### GetTaskSeeAllLinkOk

`func (o *ClientConfig) GetTaskSeeAllLinkOk() (*string, bool)`

GetTaskSeeAllLinkOk returns a tuple with the TaskSeeAllLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSeeAllLink

`func (o *ClientConfig) SetTaskSeeAllLink(v string)`

SetTaskSeeAllLink sets TaskSeeAllLink field to given value.

### HasTaskSeeAllLink

`func (o *ClientConfig) HasTaskSeeAllLink() bool`

HasTaskSeeAllLink returns a boolean if a field has been set.

### GetShortcutsPrefix

`func (o *ClientConfig) GetShortcutsPrefix() string`

GetShortcutsPrefix returns the ShortcutsPrefix field if non-nil, zero value otherwise.

### GetShortcutsPrefixOk

`func (o *ClientConfig) GetShortcutsPrefixOk() (*string, bool)`

GetShortcutsPrefixOk returns a tuple with the ShortcutsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutsPrefix

`func (o *ClientConfig) SetShortcutsPrefix(v string)`

SetShortcutsPrefix sets ShortcutsPrefix field to given value.

### HasShortcutsPrefix

`func (o *ClientConfig) HasShortcutsPrefix() bool`

HasShortcutsPrefix returns a boolean if a field has been set.

### GetSsoCompanyProvider

`func (o *ClientConfig) GetSsoCompanyProvider() string`

GetSsoCompanyProvider returns the SsoCompanyProvider field if non-nil, zero value otherwise.

### GetSsoCompanyProviderOk

`func (o *ClientConfig) GetSsoCompanyProviderOk() (*string, bool)`

GetSsoCompanyProviderOk returns a tuple with the SsoCompanyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoCompanyProvider

`func (o *ClientConfig) SetSsoCompanyProvider(v string)`

SetSsoCompanyProvider sets SsoCompanyProvider field to given value.

### HasSsoCompanyProvider

`func (o *ClientConfig) HasSsoCompanyProvider() bool`

HasSsoCompanyProvider returns a boolean if a field has been set.

### GetFeedbackCustomizations

`func (o *ClientConfig) GetFeedbackCustomizations() FeedbackCustomizations`

GetFeedbackCustomizations returns the FeedbackCustomizations field if non-nil, zero value otherwise.

### GetFeedbackCustomizationsOk

`func (o *ClientConfig) GetFeedbackCustomizationsOk() (*FeedbackCustomizations, bool)`

GetFeedbackCustomizationsOk returns a tuple with the FeedbackCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackCustomizations

`func (o *ClientConfig) SetFeedbackCustomizations(v FeedbackCustomizations)`

SetFeedbackCustomizations sets FeedbackCustomizations field to given value.

### HasFeedbackCustomizations

`func (o *ClientConfig) HasFeedbackCustomizations() bool`

HasFeedbackCustomizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


