# AssistantConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatBannerText** | Pointer to **string** | Disclaimer message to be displayed as a banner on top of chat. This could be in markdown format with \&quot;\\n\&quot; between each line. | [optional] 
**ChatBoxDisclaimer** | Pointer to **string** | Disclaimer message to be displayed below the chat box. This could be in markdown format. | [optional] 
**ChatLinkUrlTemplate** | Pointer to **string** | The URL to use for outbound links to Glean Chat. Defaults to {webAppUrl}/chat. | [optional] 
**ChatStarterHeader** | Pointer to **string** | Label for the chat header during initial state. | [optional] 
**ChatStarterSubheader** | Pointer to **string** | Label for the chat subheader during initial state. | [optional] 
**AgentClientConfigs** | Pointer to [**[]AgentClientConfig**](AgentClientConfig.md) |  | [optional] 
**RedlistedDatasources** | Pointer to **[]string** | A list of datasources that are disabled in Chat | [optional] 
**GreenlistedDatasourceInstances** | Pointer to **[]string** | A list of datasources that are always visible in Chat | [optional] 
**GptAgentEnabled** | Pointer to **bool** | Whether the GPT agent (general mode) for Chat is enabled | [optional] 
**ChatHistoryEnabled** | Pointer to **bool** | Whether the chat history for Chat is enabled for the deployment | [optional] 
**ChatGuideUrl** | Pointer to **string** | Redirect URL for \&quot;Chat guide\&quot; in the default chat starter subheader | [optional] 

## Methods

### NewAssistantConfig

`func NewAssistantConfig() *AssistantConfig`

NewAssistantConfig instantiates a new AssistantConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantConfigWithDefaults

`func NewAssistantConfigWithDefaults() *AssistantConfig`

NewAssistantConfigWithDefaults instantiates a new AssistantConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatBannerText

`func (o *AssistantConfig) GetChatBannerText() string`

GetChatBannerText returns the ChatBannerText field if non-nil, zero value otherwise.

### GetChatBannerTextOk

`func (o *AssistantConfig) GetChatBannerTextOk() (*string, bool)`

GetChatBannerTextOk returns a tuple with the ChatBannerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatBannerText

`func (o *AssistantConfig) SetChatBannerText(v string)`

SetChatBannerText sets ChatBannerText field to given value.

### HasChatBannerText

`func (o *AssistantConfig) HasChatBannerText() bool`

HasChatBannerText returns a boolean if a field has been set.

### GetChatBoxDisclaimer

`func (o *AssistantConfig) GetChatBoxDisclaimer() string`

GetChatBoxDisclaimer returns the ChatBoxDisclaimer field if non-nil, zero value otherwise.

### GetChatBoxDisclaimerOk

`func (o *AssistantConfig) GetChatBoxDisclaimerOk() (*string, bool)`

GetChatBoxDisclaimerOk returns a tuple with the ChatBoxDisclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatBoxDisclaimer

`func (o *AssistantConfig) SetChatBoxDisclaimer(v string)`

SetChatBoxDisclaimer sets ChatBoxDisclaimer field to given value.

### HasChatBoxDisclaimer

`func (o *AssistantConfig) HasChatBoxDisclaimer() bool`

HasChatBoxDisclaimer returns a boolean if a field has been set.

### GetChatLinkUrlTemplate

`func (o *AssistantConfig) GetChatLinkUrlTemplate() string`

GetChatLinkUrlTemplate returns the ChatLinkUrlTemplate field if non-nil, zero value otherwise.

### GetChatLinkUrlTemplateOk

`func (o *AssistantConfig) GetChatLinkUrlTemplateOk() (*string, bool)`

GetChatLinkUrlTemplateOk returns a tuple with the ChatLinkUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatLinkUrlTemplate

`func (o *AssistantConfig) SetChatLinkUrlTemplate(v string)`

SetChatLinkUrlTemplate sets ChatLinkUrlTemplate field to given value.

### HasChatLinkUrlTemplate

`func (o *AssistantConfig) HasChatLinkUrlTemplate() bool`

HasChatLinkUrlTemplate returns a boolean if a field has been set.

### GetChatStarterHeader

`func (o *AssistantConfig) GetChatStarterHeader() string`

GetChatStarterHeader returns the ChatStarterHeader field if non-nil, zero value otherwise.

### GetChatStarterHeaderOk

`func (o *AssistantConfig) GetChatStarterHeaderOk() (*string, bool)`

GetChatStarterHeaderOk returns a tuple with the ChatStarterHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatStarterHeader

`func (o *AssistantConfig) SetChatStarterHeader(v string)`

SetChatStarterHeader sets ChatStarterHeader field to given value.

### HasChatStarterHeader

`func (o *AssistantConfig) HasChatStarterHeader() bool`

HasChatStarterHeader returns a boolean if a field has been set.

### GetChatStarterSubheader

`func (o *AssistantConfig) GetChatStarterSubheader() string`

GetChatStarterSubheader returns the ChatStarterSubheader field if non-nil, zero value otherwise.

### GetChatStarterSubheaderOk

`func (o *AssistantConfig) GetChatStarterSubheaderOk() (*string, bool)`

GetChatStarterSubheaderOk returns a tuple with the ChatStarterSubheader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatStarterSubheader

`func (o *AssistantConfig) SetChatStarterSubheader(v string)`

SetChatStarterSubheader sets ChatStarterSubheader field to given value.

### HasChatStarterSubheader

`func (o *AssistantConfig) HasChatStarterSubheader() bool`

HasChatStarterSubheader returns a boolean if a field has been set.

### GetAgentClientConfigs

`func (o *AssistantConfig) GetAgentClientConfigs() []AgentClientConfig`

GetAgentClientConfigs returns the AgentClientConfigs field if non-nil, zero value otherwise.

### GetAgentClientConfigsOk

`func (o *AssistantConfig) GetAgentClientConfigsOk() (*[]AgentClientConfig, bool)`

GetAgentClientConfigsOk returns a tuple with the AgentClientConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentClientConfigs

`func (o *AssistantConfig) SetAgentClientConfigs(v []AgentClientConfig)`

SetAgentClientConfigs sets AgentClientConfigs field to given value.

### HasAgentClientConfigs

`func (o *AssistantConfig) HasAgentClientConfigs() bool`

HasAgentClientConfigs returns a boolean if a field has been set.

### GetRedlistedDatasources

`func (o *AssistantConfig) GetRedlistedDatasources() []string`

GetRedlistedDatasources returns the RedlistedDatasources field if non-nil, zero value otherwise.

### GetRedlistedDatasourcesOk

`func (o *AssistantConfig) GetRedlistedDatasourcesOk() (*[]string, bool)`

GetRedlistedDatasourcesOk returns a tuple with the RedlistedDatasources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedlistedDatasources

`func (o *AssistantConfig) SetRedlistedDatasources(v []string)`

SetRedlistedDatasources sets RedlistedDatasources field to given value.

### HasRedlistedDatasources

`func (o *AssistantConfig) HasRedlistedDatasources() bool`

HasRedlistedDatasources returns a boolean if a field has been set.

### GetGreenlistedDatasourceInstances

`func (o *AssistantConfig) GetGreenlistedDatasourceInstances() []string`

GetGreenlistedDatasourceInstances returns the GreenlistedDatasourceInstances field if non-nil, zero value otherwise.

### GetGreenlistedDatasourceInstancesOk

`func (o *AssistantConfig) GetGreenlistedDatasourceInstancesOk() (*[]string, bool)`

GetGreenlistedDatasourceInstancesOk returns a tuple with the GreenlistedDatasourceInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreenlistedDatasourceInstances

`func (o *AssistantConfig) SetGreenlistedDatasourceInstances(v []string)`

SetGreenlistedDatasourceInstances sets GreenlistedDatasourceInstances field to given value.

### HasGreenlistedDatasourceInstances

`func (o *AssistantConfig) HasGreenlistedDatasourceInstances() bool`

HasGreenlistedDatasourceInstances returns a boolean if a field has been set.

### GetGptAgentEnabled

`func (o *AssistantConfig) GetGptAgentEnabled() bool`

GetGptAgentEnabled returns the GptAgentEnabled field if non-nil, zero value otherwise.

### GetGptAgentEnabledOk

`func (o *AssistantConfig) GetGptAgentEnabledOk() (*bool, bool)`

GetGptAgentEnabledOk returns a tuple with the GptAgentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGptAgentEnabled

`func (o *AssistantConfig) SetGptAgentEnabled(v bool)`

SetGptAgentEnabled sets GptAgentEnabled field to given value.

### HasGptAgentEnabled

`func (o *AssistantConfig) HasGptAgentEnabled() bool`

HasGptAgentEnabled returns a boolean if a field has been set.

### GetChatHistoryEnabled

`func (o *AssistantConfig) GetChatHistoryEnabled() bool`

GetChatHistoryEnabled returns the ChatHistoryEnabled field if non-nil, zero value otherwise.

### GetChatHistoryEnabledOk

`func (o *AssistantConfig) GetChatHistoryEnabledOk() (*bool, bool)`

GetChatHistoryEnabledOk returns a tuple with the ChatHistoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatHistoryEnabled

`func (o *AssistantConfig) SetChatHistoryEnabled(v bool)`

SetChatHistoryEnabled sets ChatHistoryEnabled field to given value.

### HasChatHistoryEnabled

`func (o *AssistantConfig) HasChatHistoryEnabled() bool`

HasChatHistoryEnabled returns a boolean if a field has been set.

### GetChatGuideUrl

`func (o *AssistantConfig) GetChatGuideUrl() string`

GetChatGuideUrl returns the ChatGuideUrl field if non-nil, zero value otherwise.

### GetChatGuideUrlOk

`func (o *AssistantConfig) GetChatGuideUrlOk() (*string, bool)`

GetChatGuideUrlOk returns a tuple with the ChatGuideUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatGuideUrl

`func (o *AssistantConfig) SetChatGuideUrl(v string)`

SetChatGuideUrl sets ChatGuideUrl field to given value.

### HasChatGuideUrl

`func (o *AssistantConfig) HasChatGuideUrl() bool`

HasChatGuideUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


