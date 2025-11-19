# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Uuid** | **string** |  | [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**AccessControl** | Pointer to **bool** |  | [optional] 
**Organization** | **string** |  | [readonly] 
**ProjectId** | **int64** |  | [readonly] 
**ApiToken** | **string** |  | [readonly] 
**SecretApiToken** | **NullableString** |  | [readonly] 
**SecretApiTokenBackup** | **NullableString** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**IngestedEvent** | **bool** |  | [readonly] 
**DefaultModifiers** | **map[string]interface{}** |  | [readonly] 
**PersonOnEventsQueryingEnabled** | **bool** |  | [readonly] 
**UserAccessLevel** | **NullableString** | The effective access level the user has for this object | [readonly] 
**AppUrls** | Pointer to **[]string** |  | [optional] 
**SlackIncomingWebhook** | Pointer to **NullableString** |  | [optional] 
**AnonymizeIps** | Pointer to **bool** |  | [optional] 
**CompletedSnippetOnboarding** | Pointer to **bool** |  | [optional] 
**TestAccountFilters** | Pointer to **interface{}** |  | [optional] 
**TestAccountFiltersDefaultChecked** | Pointer to **NullableBool** |  | [optional] 
**PathCleaningFilters** | Pointer to **interface{}** |  | [optional] 
**IsDemo** | Pointer to **bool** |  | [optional] 
**Timezone** | Pointer to [**TimezoneEnum**](TimezoneEnum.md) |  | [optional] 
**DataAttributes** | Pointer to **interface{}** |  | [optional] 
**PersonDisplayNameProperties** | Pointer to **[]string** |  | [optional] 
**CorrelationConfig** | Pointer to **interface{}** |  | [optional] 
**AutocaptureOptOut** | Pointer to **NullableBool** |  | [optional] 
**AutocaptureExceptionsOptIn** | Pointer to **NullableBool** |  | [optional] 
**AutocaptureWebVitalsOptIn** | Pointer to **NullableBool** |  | [optional] 
**AutocaptureWebVitalsAllowedMetrics** | Pointer to **interface{}** |  | [optional] 
**AutocaptureExceptionsErrorsToIgnore** | Pointer to **interface{}** |  | [optional] 
**CaptureConsoleLogOptIn** | Pointer to **NullableBool** |  | [optional] 
**CapturePerformanceOptIn** | Pointer to **NullableBool** |  | [optional] 
**SessionRecordingOptIn** | Pointer to **bool** |  | [optional] 
**SessionRecordingSampleRate** | Pointer to **NullableFloat64** |  | [optional] 
**SessionRecordingMinimumDurationMilliseconds** | Pointer to **NullableInt32** |  | [optional] 
**SessionRecordingLinkedFlag** | Pointer to **interface{}** |  | [optional] 
**SessionRecordingNetworkPayloadCaptureConfig** | Pointer to **interface{}** |  | [optional] 
**SessionRecordingMaskingConfig** | Pointer to **interface{}** |  | [optional] 
**SessionRecordingUrlTriggerConfig** | Pointer to **[]interface{}** |  | [optional] 
**SessionRecordingUrlBlocklistConfig** | Pointer to **[]interface{}** |  | [optional] 
**SessionRecordingEventTriggerConfig** | Pointer to **[]string** |  | [optional] 
**SessionRecordingTriggerMatchTypeConfig** | Pointer to **NullableString** |  | [optional] 
**SessionRecordingRetentionPeriod** | Pointer to [**SessionRecordingRetentionPeriodEnum**](SessionRecordingRetentionPeriodEnum.md) |  | [optional] 
**SessionReplayConfig** | Pointer to **interface{}** |  | [optional] 
**SurveyConfig** | Pointer to **interface{}** |  | [optional] 
**WeekStartDay** | Pointer to [**NullableWeekStartDayEnum**](WeekStartDayEnum.md) |  | [optional] 
**PrimaryDashboard** | Pointer to **NullableInt32** |  | [optional] 
**LiveEventsColumns** | Pointer to **[]string** |  | [optional] 
**RecordingDomains** | Pointer to **[]string** |  | [optional] 
**CookielessServerHashMode** | Pointer to [**NullableCookielessServerHashModeEnum**](CookielessServerHashModeEnum.md) |  | [optional] 
**HumanFriendlyComparisonPeriods** | Pointer to **NullableBool** |  | [optional] 
**InjectWebApps** | Pointer to **NullableBool** |  | [optional] 
**ExtraSettings** | Pointer to **interface{}** |  | [optional] 
**Modifiers** | Pointer to **interface{}** |  | [optional] 
**HasCompletedOnboardingFor** | Pointer to **interface{}** |  | [optional] 
**SurveysOptIn** | Pointer to **NullableBool** |  | [optional] 
**HeatmapsOptIn** | Pointer to **NullableBool** |  | [optional] 
**FlagsPersistenceDefault** | Pointer to **NullableBool** |  | [optional] 
**FeatureFlagConfirmationEnabled** | Pointer to **NullableBool** |  | [optional] 
**FeatureFlagConfirmationMessage** | Pointer to **NullableString** |  | [optional] 
**DefaultEvaluationEnvironmentsEnabled** | Pointer to **NullableBool** | Whether to automatically apply default evaluation environments to new feature flags | [optional] 
**CaptureDeadClicks** | Pointer to **NullableBool** |  | [optional] 
**DefaultDataTheme** | Pointer to **NullableInt32** |  | [optional] 
**RevenueAnalyticsConfig** | Pointer to [**TeamRevenueAnalyticsConfig**](TeamRevenueAnalyticsConfig.md) |  | [optional] 
**MarketingAnalyticsConfig** | Pointer to [**TeamMarketingAnalyticsConfig**](TeamMarketingAnalyticsConfig.md) |  | [optional] 
**OnboardingTasks** | Pointer to **interface{}** |  | [optional] 
**BaseCurrency** | Pointer to [**BaseCurrencyEnum**](BaseCurrencyEnum.md) |  | [optional] [default to USD]
**WebAnalyticsPreAggregatedTablesEnabled** | Pointer to **NullableBool** |  | [optional] 
**ExperimentRecalculationTime** | Pointer to **NullableString** | Time of day (UTC) when experiment metrics should be recalculated. If not set, uses the default recalculation time. | [optional] 
**ReceiveOrgLevelActivityLogs** | Pointer to **NullableBool** |  | [optional] 
**EffectiveMembershipLevel** | [**NullableEffectiveMembershipLevelEnum**](EffectiveMembershipLevelEnum.md) |  | [readonly] 
**HasGroupTypes** | **bool** |  | [readonly] 
**GroupTypes** | **[]map[string]interface{}** |  | [readonly] 
**LiveEventsToken** | **NullableString** |  | [readonly] 
**ProductIntents** | **string** |  | [readonly] 
**ManagedViewsets** | **string** |  | [readonly] 

## Methods

### NewTeam

`func NewTeam(id int32, uuid string, organization string, projectId int64, apiToken string, secretApiToken NullableString, secretApiTokenBackup NullableString, createdAt time.Time, updatedAt time.Time, ingestedEvent bool, defaultModifiers map[string]interface{}, personOnEventsQueryingEnabled bool, userAccessLevel NullableString, effectiveMembershipLevel NullableEffectiveMembershipLevelEnum, hasGroupTypes bool, groupTypes []map[string]interface{}, liveEventsToken NullableString, productIntents string, managedViewsets string, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v int32)`

SetId sets Id field to given value.


### GetUuid

`func (o *Team) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Team) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Team) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Team) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccessControl

`func (o *Team) GetAccessControl() bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *Team) GetAccessControlOk() (*bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *Team) SetAccessControl(v bool)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *Team) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetOrganization

`func (o *Team) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Team) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Team) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetProjectId

`func (o *Team) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Team) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Team) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetApiToken

`func (o *Team) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *Team) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *Team) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetSecretApiToken

`func (o *Team) GetSecretApiToken() string`

GetSecretApiToken returns the SecretApiToken field if non-nil, zero value otherwise.

### GetSecretApiTokenOk

`func (o *Team) GetSecretApiTokenOk() (*string, bool)`

GetSecretApiTokenOk returns a tuple with the SecretApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiToken

`func (o *Team) SetSecretApiToken(v string)`

SetSecretApiToken sets SecretApiToken field to given value.


### SetSecretApiTokenNil

`func (o *Team) SetSecretApiTokenNil(b bool)`

 SetSecretApiTokenNil sets the value for SecretApiToken to be an explicit nil

### UnsetSecretApiToken
`func (o *Team) UnsetSecretApiToken()`

UnsetSecretApiToken ensures that no value is present for SecretApiToken, not even an explicit nil
### GetSecretApiTokenBackup

`func (o *Team) GetSecretApiTokenBackup() string`

GetSecretApiTokenBackup returns the SecretApiTokenBackup field if non-nil, zero value otherwise.

### GetSecretApiTokenBackupOk

`func (o *Team) GetSecretApiTokenBackupOk() (*string, bool)`

GetSecretApiTokenBackupOk returns a tuple with the SecretApiTokenBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiTokenBackup

`func (o *Team) SetSecretApiTokenBackup(v string)`

SetSecretApiTokenBackup sets SecretApiTokenBackup field to given value.


### SetSecretApiTokenBackupNil

`func (o *Team) SetSecretApiTokenBackupNil(b bool)`

 SetSecretApiTokenBackupNil sets the value for SecretApiTokenBackup to be an explicit nil

### UnsetSecretApiTokenBackup
`func (o *Team) UnsetSecretApiTokenBackup()`

UnsetSecretApiTokenBackup ensures that no value is present for SecretApiTokenBackup, not even an explicit nil
### GetCreatedAt

`func (o *Team) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Team) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Team) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Team) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Team) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Team) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIngestedEvent

`func (o *Team) GetIngestedEvent() bool`

GetIngestedEvent returns the IngestedEvent field if non-nil, zero value otherwise.

### GetIngestedEventOk

`func (o *Team) GetIngestedEventOk() (*bool, bool)`

GetIngestedEventOk returns a tuple with the IngestedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEvent

`func (o *Team) SetIngestedEvent(v bool)`

SetIngestedEvent sets IngestedEvent field to given value.


### GetDefaultModifiers

`func (o *Team) GetDefaultModifiers() map[string]interface{}`

GetDefaultModifiers returns the DefaultModifiers field if non-nil, zero value otherwise.

### GetDefaultModifiersOk

`func (o *Team) GetDefaultModifiersOk() (*map[string]interface{}, bool)`

GetDefaultModifiersOk returns a tuple with the DefaultModifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModifiers

`func (o *Team) SetDefaultModifiers(v map[string]interface{})`

SetDefaultModifiers sets DefaultModifiers field to given value.


### GetPersonOnEventsQueryingEnabled

`func (o *Team) GetPersonOnEventsQueryingEnabled() bool`

GetPersonOnEventsQueryingEnabled returns the PersonOnEventsQueryingEnabled field if non-nil, zero value otherwise.

### GetPersonOnEventsQueryingEnabledOk

`func (o *Team) GetPersonOnEventsQueryingEnabledOk() (*bool, bool)`

GetPersonOnEventsQueryingEnabledOk returns a tuple with the PersonOnEventsQueryingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonOnEventsQueryingEnabled

`func (o *Team) SetPersonOnEventsQueryingEnabled(v bool)`

SetPersonOnEventsQueryingEnabled sets PersonOnEventsQueryingEnabled field to given value.


### GetUserAccessLevel

`func (o *Team) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *Team) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *Team) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.


### SetUserAccessLevelNil

`func (o *Team) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *Team) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetAppUrls

`func (o *Team) GetAppUrls() []*string`

GetAppUrls returns the AppUrls field if non-nil, zero value otherwise.

### GetAppUrlsOk

`func (o *Team) GetAppUrlsOk() (*[]*string, bool)`

GetAppUrlsOk returns a tuple with the AppUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrls

`func (o *Team) SetAppUrls(v []*string)`

SetAppUrls sets AppUrls field to given value.

### HasAppUrls

`func (o *Team) HasAppUrls() bool`

HasAppUrls returns a boolean if a field has been set.

### GetSlackIncomingWebhook

`func (o *Team) GetSlackIncomingWebhook() string`

GetSlackIncomingWebhook returns the SlackIncomingWebhook field if non-nil, zero value otherwise.

### GetSlackIncomingWebhookOk

`func (o *Team) GetSlackIncomingWebhookOk() (*string, bool)`

GetSlackIncomingWebhookOk returns a tuple with the SlackIncomingWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackIncomingWebhook

`func (o *Team) SetSlackIncomingWebhook(v string)`

SetSlackIncomingWebhook sets SlackIncomingWebhook field to given value.

### HasSlackIncomingWebhook

`func (o *Team) HasSlackIncomingWebhook() bool`

HasSlackIncomingWebhook returns a boolean if a field has been set.

### SetSlackIncomingWebhookNil

`func (o *Team) SetSlackIncomingWebhookNil(b bool)`

 SetSlackIncomingWebhookNil sets the value for SlackIncomingWebhook to be an explicit nil

### UnsetSlackIncomingWebhook
`func (o *Team) UnsetSlackIncomingWebhook()`

UnsetSlackIncomingWebhook ensures that no value is present for SlackIncomingWebhook, not even an explicit nil
### GetAnonymizeIps

`func (o *Team) GetAnonymizeIps() bool`

GetAnonymizeIps returns the AnonymizeIps field if non-nil, zero value otherwise.

### GetAnonymizeIpsOk

`func (o *Team) GetAnonymizeIpsOk() (*bool, bool)`

GetAnonymizeIpsOk returns a tuple with the AnonymizeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeIps

`func (o *Team) SetAnonymizeIps(v bool)`

SetAnonymizeIps sets AnonymizeIps field to given value.

### HasAnonymizeIps

`func (o *Team) HasAnonymizeIps() bool`

HasAnonymizeIps returns a boolean if a field has been set.

### GetCompletedSnippetOnboarding

`func (o *Team) GetCompletedSnippetOnboarding() bool`

GetCompletedSnippetOnboarding returns the CompletedSnippetOnboarding field if non-nil, zero value otherwise.

### GetCompletedSnippetOnboardingOk

`func (o *Team) GetCompletedSnippetOnboardingOk() (*bool, bool)`

GetCompletedSnippetOnboardingOk returns a tuple with the CompletedSnippetOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedSnippetOnboarding

`func (o *Team) SetCompletedSnippetOnboarding(v bool)`

SetCompletedSnippetOnboarding sets CompletedSnippetOnboarding field to given value.

### HasCompletedSnippetOnboarding

`func (o *Team) HasCompletedSnippetOnboarding() bool`

HasCompletedSnippetOnboarding returns a boolean if a field has been set.

### GetTestAccountFilters

`func (o *Team) GetTestAccountFilters() interface{}`

GetTestAccountFilters returns the TestAccountFilters field if non-nil, zero value otherwise.

### GetTestAccountFiltersOk

`func (o *Team) GetTestAccountFiltersOk() (*interface{}, bool)`

GetTestAccountFiltersOk returns a tuple with the TestAccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccountFilters

`func (o *Team) SetTestAccountFilters(v interface{})`

SetTestAccountFilters sets TestAccountFilters field to given value.

### HasTestAccountFilters

`func (o *Team) HasTestAccountFilters() bool`

HasTestAccountFilters returns a boolean if a field has been set.

### SetTestAccountFiltersNil

`func (o *Team) SetTestAccountFiltersNil(b bool)`

 SetTestAccountFiltersNil sets the value for TestAccountFilters to be an explicit nil

### UnsetTestAccountFilters
`func (o *Team) UnsetTestAccountFilters()`

UnsetTestAccountFilters ensures that no value is present for TestAccountFilters, not even an explicit nil
### GetTestAccountFiltersDefaultChecked

`func (o *Team) GetTestAccountFiltersDefaultChecked() bool`

GetTestAccountFiltersDefaultChecked returns the TestAccountFiltersDefaultChecked field if non-nil, zero value otherwise.

### GetTestAccountFiltersDefaultCheckedOk

`func (o *Team) GetTestAccountFiltersDefaultCheckedOk() (*bool, bool)`

GetTestAccountFiltersDefaultCheckedOk returns a tuple with the TestAccountFiltersDefaultChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccountFiltersDefaultChecked

`func (o *Team) SetTestAccountFiltersDefaultChecked(v bool)`

SetTestAccountFiltersDefaultChecked sets TestAccountFiltersDefaultChecked field to given value.

### HasTestAccountFiltersDefaultChecked

`func (o *Team) HasTestAccountFiltersDefaultChecked() bool`

HasTestAccountFiltersDefaultChecked returns a boolean if a field has been set.

### SetTestAccountFiltersDefaultCheckedNil

`func (o *Team) SetTestAccountFiltersDefaultCheckedNil(b bool)`

 SetTestAccountFiltersDefaultCheckedNil sets the value for TestAccountFiltersDefaultChecked to be an explicit nil

### UnsetTestAccountFiltersDefaultChecked
`func (o *Team) UnsetTestAccountFiltersDefaultChecked()`

UnsetTestAccountFiltersDefaultChecked ensures that no value is present for TestAccountFiltersDefaultChecked, not even an explicit nil
### GetPathCleaningFilters

`func (o *Team) GetPathCleaningFilters() interface{}`

GetPathCleaningFilters returns the PathCleaningFilters field if non-nil, zero value otherwise.

### GetPathCleaningFiltersOk

`func (o *Team) GetPathCleaningFiltersOk() (*interface{}, bool)`

GetPathCleaningFiltersOk returns a tuple with the PathCleaningFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathCleaningFilters

`func (o *Team) SetPathCleaningFilters(v interface{})`

SetPathCleaningFilters sets PathCleaningFilters field to given value.

### HasPathCleaningFilters

`func (o *Team) HasPathCleaningFilters() bool`

HasPathCleaningFilters returns a boolean if a field has been set.

### SetPathCleaningFiltersNil

`func (o *Team) SetPathCleaningFiltersNil(b bool)`

 SetPathCleaningFiltersNil sets the value for PathCleaningFilters to be an explicit nil

### UnsetPathCleaningFilters
`func (o *Team) UnsetPathCleaningFilters()`

UnsetPathCleaningFilters ensures that no value is present for PathCleaningFilters, not even an explicit nil
### GetIsDemo

`func (o *Team) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *Team) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *Team) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *Team) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetTimezone

`func (o *Team) GetTimezone() TimezoneEnum`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Team) GetTimezoneOk() (*TimezoneEnum, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Team) SetTimezone(v TimezoneEnum)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Team) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDataAttributes

`func (o *Team) GetDataAttributes() interface{}`

GetDataAttributes returns the DataAttributes field if non-nil, zero value otherwise.

### GetDataAttributesOk

`func (o *Team) GetDataAttributesOk() (*interface{}, bool)`

GetDataAttributesOk returns a tuple with the DataAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAttributes

`func (o *Team) SetDataAttributes(v interface{})`

SetDataAttributes sets DataAttributes field to given value.

### HasDataAttributes

`func (o *Team) HasDataAttributes() bool`

HasDataAttributes returns a boolean if a field has been set.

### SetDataAttributesNil

`func (o *Team) SetDataAttributesNil(b bool)`

 SetDataAttributesNil sets the value for DataAttributes to be an explicit nil

### UnsetDataAttributes
`func (o *Team) UnsetDataAttributes()`

UnsetDataAttributes ensures that no value is present for DataAttributes, not even an explicit nil
### GetPersonDisplayNameProperties

`func (o *Team) GetPersonDisplayNameProperties() []string`

GetPersonDisplayNameProperties returns the PersonDisplayNameProperties field if non-nil, zero value otherwise.

### GetPersonDisplayNamePropertiesOk

`func (o *Team) GetPersonDisplayNamePropertiesOk() (*[]string, bool)`

GetPersonDisplayNamePropertiesOk returns a tuple with the PersonDisplayNameProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDisplayNameProperties

`func (o *Team) SetPersonDisplayNameProperties(v []string)`

SetPersonDisplayNameProperties sets PersonDisplayNameProperties field to given value.

### HasPersonDisplayNameProperties

`func (o *Team) HasPersonDisplayNameProperties() bool`

HasPersonDisplayNameProperties returns a boolean if a field has been set.

### SetPersonDisplayNamePropertiesNil

`func (o *Team) SetPersonDisplayNamePropertiesNil(b bool)`

 SetPersonDisplayNamePropertiesNil sets the value for PersonDisplayNameProperties to be an explicit nil

### UnsetPersonDisplayNameProperties
`func (o *Team) UnsetPersonDisplayNameProperties()`

UnsetPersonDisplayNameProperties ensures that no value is present for PersonDisplayNameProperties, not even an explicit nil
### GetCorrelationConfig

`func (o *Team) GetCorrelationConfig() interface{}`

GetCorrelationConfig returns the CorrelationConfig field if non-nil, zero value otherwise.

### GetCorrelationConfigOk

`func (o *Team) GetCorrelationConfigOk() (*interface{}, bool)`

GetCorrelationConfigOk returns a tuple with the CorrelationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationConfig

`func (o *Team) SetCorrelationConfig(v interface{})`

SetCorrelationConfig sets CorrelationConfig field to given value.

### HasCorrelationConfig

`func (o *Team) HasCorrelationConfig() bool`

HasCorrelationConfig returns a boolean if a field has been set.

### SetCorrelationConfigNil

`func (o *Team) SetCorrelationConfigNil(b bool)`

 SetCorrelationConfigNil sets the value for CorrelationConfig to be an explicit nil

### UnsetCorrelationConfig
`func (o *Team) UnsetCorrelationConfig()`

UnsetCorrelationConfig ensures that no value is present for CorrelationConfig, not even an explicit nil
### GetAutocaptureOptOut

`func (o *Team) GetAutocaptureOptOut() bool`

GetAutocaptureOptOut returns the AutocaptureOptOut field if non-nil, zero value otherwise.

### GetAutocaptureOptOutOk

`func (o *Team) GetAutocaptureOptOutOk() (*bool, bool)`

GetAutocaptureOptOutOk returns a tuple with the AutocaptureOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureOptOut

`func (o *Team) SetAutocaptureOptOut(v bool)`

SetAutocaptureOptOut sets AutocaptureOptOut field to given value.

### HasAutocaptureOptOut

`func (o *Team) HasAutocaptureOptOut() bool`

HasAutocaptureOptOut returns a boolean if a field has been set.

### SetAutocaptureOptOutNil

`func (o *Team) SetAutocaptureOptOutNil(b bool)`

 SetAutocaptureOptOutNil sets the value for AutocaptureOptOut to be an explicit nil

### UnsetAutocaptureOptOut
`func (o *Team) UnsetAutocaptureOptOut()`

UnsetAutocaptureOptOut ensures that no value is present for AutocaptureOptOut, not even an explicit nil
### GetAutocaptureExceptionsOptIn

`func (o *Team) GetAutocaptureExceptionsOptIn() bool`

GetAutocaptureExceptionsOptIn returns the AutocaptureExceptionsOptIn field if non-nil, zero value otherwise.

### GetAutocaptureExceptionsOptInOk

`func (o *Team) GetAutocaptureExceptionsOptInOk() (*bool, bool)`

GetAutocaptureExceptionsOptInOk returns a tuple with the AutocaptureExceptionsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureExceptionsOptIn

`func (o *Team) SetAutocaptureExceptionsOptIn(v bool)`

SetAutocaptureExceptionsOptIn sets AutocaptureExceptionsOptIn field to given value.

### HasAutocaptureExceptionsOptIn

`func (o *Team) HasAutocaptureExceptionsOptIn() bool`

HasAutocaptureExceptionsOptIn returns a boolean if a field has been set.

### SetAutocaptureExceptionsOptInNil

`func (o *Team) SetAutocaptureExceptionsOptInNil(b bool)`

 SetAutocaptureExceptionsOptInNil sets the value for AutocaptureExceptionsOptIn to be an explicit nil

### UnsetAutocaptureExceptionsOptIn
`func (o *Team) UnsetAutocaptureExceptionsOptIn()`

UnsetAutocaptureExceptionsOptIn ensures that no value is present for AutocaptureExceptionsOptIn, not even an explicit nil
### GetAutocaptureWebVitalsOptIn

`func (o *Team) GetAutocaptureWebVitalsOptIn() bool`

GetAutocaptureWebVitalsOptIn returns the AutocaptureWebVitalsOptIn field if non-nil, zero value otherwise.

### GetAutocaptureWebVitalsOptInOk

`func (o *Team) GetAutocaptureWebVitalsOptInOk() (*bool, bool)`

GetAutocaptureWebVitalsOptInOk returns a tuple with the AutocaptureWebVitalsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureWebVitalsOptIn

`func (o *Team) SetAutocaptureWebVitalsOptIn(v bool)`

SetAutocaptureWebVitalsOptIn sets AutocaptureWebVitalsOptIn field to given value.

### HasAutocaptureWebVitalsOptIn

`func (o *Team) HasAutocaptureWebVitalsOptIn() bool`

HasAutocaptureWebVitalsOptIn returns a boolean if a field has been set.

### SetAutocaptureWebVitalsOptInNil

`func (o *Team) SetAutocaptureWebVitalsOptInNil(b bool)`

 SetAutocaptureWebVitalsOptInNil sets the value for AutocaptureWebVitalsOptIn to be an explicit nil

### UnsetAutocaptureWebVitalsOptIn
`func (o *Team) UnsetAutocaptureWebVitalsOptIn()`

UnsetAutocaptureWebVitalsOptIn ensures that no value is present for AutocaptureWebVitalsOptIn, not even an explicit nil
### GetAutocaptureWebVitalsAllowedMetrics

`func (o *Team) GetAutocaptureWebVitalsAllowedMetrics() interface{}`

GetAutocaptureWebVitalsAllowedMetrics returns the AutocaptureWebVitalsAllowedMetrics field if non-nil, zero value otherwise.

### GetAutocaptureWebVitalsAllowedMetricsOk

`func (o *Team) GetAutocaptureWebVitalsAllowedMetricsOk() (*interface{}, bool)`

GetAutocaptureWebVitalsAllowedMetricsOk returns a tuple with the AutocaptureWebVitalsAllowedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureWebVitalsAllowedMetrics

`func (o *Team) SetAutocaptureWebVitalsAllowedMetrics(v interface{})`

SetAutocaptureWebVitalsAllowedMetrics sets AutocaptureWebVitalsAllowedMetrics field to given value.

### HasAutocaptureWebVitalsAllowedMetrics

`func (o *Team) HasAutocaptureWebVitalsAllowedMetrics() bool`

HasAutocaptureWebVitalsAllowedMetrics returns a boolean if a field has been set.

### SetAutocaptureWebVitalsAllowedMetricsNil

`func (o *Team) SetAutocaptureWebVitalsAllowedMetricsNil(b bool)`

 SetAutocaptureWebVitalsAllowedMetricsNil sets the value for AutocaptureWebVitalsAllowedMetrics to be an explicit nil

### UnsetAutocaptureWebVitalsAllowedMetrics
`func (o *Team) UnsetAutocaptureWebVitalsAllowedMetrics()`

UnsetAutocaptureWebVitalsAllowedMetrics ensures that no value is present for AutocaptureWebVitalsAllowedMetrics, not even an explicit nil
### GetAutocaptureExceptionsErrorsToIgnore

`func (o *Team) GetAutocaptureExceptionsErrorsToIgnore() interface{}`

GetAutocaptureExceptionsErrorsToIgnore returns the AutocaptureExceptionsErrorsToIgnore field if non-nil, zero value otherwise.

### GetAutocaptureExceptionsErrorsToIgnoreOk

`func (o *Team) GetAutocaptureExceptionsErrorsToIgnoreOk() (*interface{}, bool)`

GetAutocaptureExceptionsErrorsToIgnoreOk returns a tuple with the AutocaptureExceptionsErrorsToIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureExceptionsErrorsToIgnore

`func (o *Team) SetAutocaptureExceptionsErrorsToIgnore(v interface{})`

SetAutocaptureExceptionsErrorsToIgnore sets AutocaptureExceptionsErrorsToIgnore field to given value.

### HasAutocaptureExceptionsErrorsToIgnore

`func (o *Team) HasAutocaptureExceptionsErrorsToIgnore() bool`

HasAutocaptureExceptionsErrorsToIgnore returns a boolean if a field has been set.

### SetAutocaptureExceptionsErrorsToIgnoreNil

`func (o *Team) SetAutocaptureExceptionsErrorsToIgnoreNil(b bool)`

 SetAutocaptureExceptionsErrorsToIgnoreNil sets the value for AutocaptureExceptionsErrorsToIgnore to be an explicit nil

### UnsetAutocaptureExceptionsErrorsToIgnore
`func (o *Team) UnsetAutocaptureExceptionsErrorsToIgnore()`

UnsetAutocaptureExceptionsErrorsToIgnore ensures that no value is present for AutocaptureExceptionsErrorsToIgnore, not even an explicit nil
### GetCaptureConsoleLogOptIn

`func (o *Team) GetCaptureConsoleLogOptIn() bool`

GetCaptureConsoleLogOptIn returns the CaptureConsoleLogOptIn field if non-nil, zero value otherwise.

### GetCaptureConsoleLogOptInOk

`func (o *Team) GetCaptureConsoleLogOptInOk() (*bool, bool)`

GetCaptureConsoleLogOptInOk returns a tuple with the CaptureConsoleLogOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureConsoleLogOptIn

`func (o *Team) SetCaptureConsoleLogOptIn(v bool)`

SetCaptureConsoleLogOptIn sets CaptureConsoleLogOptIn field to given value.

### HasCaptureConsoleLogOptIn

`func (o *Team) HasCaptureConsoleLogOptIn() bool`

HasCaptureConsoleLogOptIn returns a boolean if a field has been set.

### SetCaptureConsoleLogOptInNil

`func (o *Team) SetCaptureConsoleLogOptInNil(b bool)`

 SetCaptureConsoleLogOptInNil sets the value for CaptureConsoleLogOptIn to be an explicit nil

### UnsetCaptureConsoleLogOptIn
`func (o *Team) UnsetCaptureConsoleLogOptIn()`

UnsetCaptureConsoleLogOptIn ensures that no value is present for CaptureConsoleLogOptIn, not even an explicit nil
### GetCapturePerformanceOptIn

`func (o *Team) GetCapturePerformanceOptIn() bool`

GetCapturePerformanceOptIn returns the CapturePerformanceOptIn field if non-nil, zero value otherwise.

### GetCapturePerformanceOptInOk

`func (o *Team) GetCapturePerformanceOptInOk() (*bool, bool)`

GetCapturePerformanceOptInOk returns a tuple with the CapturePerformanceOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturePerformanceOptIn

`func (o *Team) SetCapturePerformanceOptIn(v bool)`

SetCapturePerformanceOptIn sets CapturePerformanceOptIn field to given value.

### HasCapturePerformanceOptIn

`func (o *Team) HasCapturePerformanceOptIn() bool`

HasCapturePerformanceOptIn returns a boolean if a field has been set.

### SetCapturePerformanceOptInNil

`func (o *Team) SetCapturePerformanceOptInNil(b bool)`

 SetCapturePerformanceOptInNil sets the value for CapturePerformanceOptIn to be an explicit nil

### UnsetCapturePerformanceOptIn
`func (o *Team) UnsetCapturePerformanceOptIn()`

UnsetCapturePerformanceOptIn ensures that no value is present for CapturePerformanceOptIn, not even an explicit nil
### GetSessionRecordingOptIn

`func (o *Team) GetSessionRecordingOptIn() bool`

GetSessionRecordingOptIn returns the SessionRecordingOptIn field if non-nil, zero value otherwise.

### GetSessionRecordingOptInOk

`func (o *Team) GetSessionRecordingOptInOk() (*bool, bool)`

GetSessionRecordingOptInOk returns a tuple with the SessionRecordingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingOptIn

`func (o *Team) SetSessionRecordingOptIn(v bool)`

SetSessionRecordingOptIn sets SessionRecordingOptIn field to given value.

### HasSessionRecordingOptIn

`func (o *Team) HasSessionRecordingOptIn() bool`

HasSessionRecordingOptIn returns a boolean if a field has been set.

### GetSessionRecordingSampleRate

`func (o *Team) GetSessionRecordingSampleRate() float64`

GetSessionRecordingSampleRate returns the SessionRecordingSampleRate field if non-nil, zero value otherwise.

### GetSessionRecordingSampleRateOk

`func (o *Team) GetSessionRecordingSampleRateOk() (*float64, bool)`

GetSessionRecordingSampleRateOk returns a tuple with the SessionRecordingSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingSampleRate

`func (o *Team) SetSessionRecordingSampleRate(v float64)`

SetSessionRecordingSampleRate sets SessionRecordingSampleRate field to given value.

### HasSessionRecordingSampleRate

`func (o *Team) HasSessionRecordingSampleRate() bool`

HasSessionRecordingSampleRate returns a boolean if a field has been set.

### SetSessionRecordingSampleRateNil

`func (o *Team) SetSessionRecordingSampleRateNil(b bool)`

 SetSessionRecordingSampleRateNil sets the value for SessionRecordingSampleRate to be an explicit nil

### UnsetSessionRecordingSampleRate
`func (o *Team) UnsetSessionRecordingSampleRate()`

UnsetSessionRecordingSampleRate ensures that no value is present for SessionRecordingSampleRate, not even an explicit nil
### GetSessionRecordingMinimumDurationMilliseconds

`func (o *Team) GetSessionRecordingMinimumDurationMilliseconds() int32`

GetSessionRecordingMinimumDurationMilliseconds returns the SessionRecordingMinimumDurationMilliseconds field if non-nil, zero value otherwise.

### GetSessionRecordingMinimumDurationMillisecondsOk

`func (o *Team) GetSessionRecordingMinimumDurationMillisecondsOk() (*int32, bool)`

GetSessionRecordingMinimumDurationMillisecondsOk returns a tuple with the SessionRecordingMinimumDurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingMinimumDurationMilliseconds

`func (o *Team) SetSessionRecordingMinimumDurationMilliseconds(v int32)`

SetSessionRecordingMinimumDurationMilliseconds sets SessionRecordingMinimumDurationMilliseconds field to given value.

### HasSessionRecordingMinimumDurationMilliseconds

`func (o *Team) HasSessionRecordingMinimumDurationMilliseconds() bool`

HasSessionRecordingMinimumDurationMilliseconds returns a boolean if a field has been set.

### SetSessionRecordingMinimumDurationMillisecondsNil

`func (o *Team) SetSessionRecordingMinimumDurationMillisecondsNil(b bool)`

 SetSessionRecordingMinimumDurationMillisecondsNil sets the value for SessionRecordingMinimumDurationMilliseconds to be an explicit nil

### UnsetSessionRecordingMinimumDurationMilliseconds
`func (o *Team) UnsetSessionRecordingMinimumDurationMilliseconds()`

UnsetSessionRecordingMinimumDurationMilliseconds ensures that no value is present for SessionRecordingMinimumDurationMilliseconds, not even an explicit nil
### GetSessionRecordingLinkedFlag

`func (o *Team) GetSessionRecordingLinkedFlag() interface{}`

GetSessionRecordingLinkedFlag returns the SessionRecordingLinkedFlag field if non-nil, zero value otherwise.

### GetSessionRecordingLinkedFlagOk

`func (o *Team) GetSessionRecordingLinkedFlagOk() (*interface{}, bool)`

GetSessionRecordingLinkedFlagOk returns a tuple with the SessionRecordingLinkedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingLinkedFlag

`func (o *Team) SetSessionRecordingLinkedFlag(v interface{})`

SetSessionRecordingLinkedFlag sets SessionRecordingLinkedFlag field to given value.

### HasSessionRecordingLinkedFlag

`func (o *Team) HasSessionRecordingLinkedFlag() bool`

HasSessionRecordingLinkedFlag returns a boolean if a field has been set.

### SetSessionRecordingLinkedFlagNil

`func (o *Team) SetSessionRecordingLinkedFlagNil(b bool)`

 SetSessionRecordingLinkedFlagNil sets the value for SessionRecordingLinkedFlag to be an explicit nil

### UnsetSessionRecordingLinkedFlag
`func (o *Team) UnsetSessionRecordingLinkedFlag()`

UnsetSessionRecordingLinkedFlag ensures that no value is present for SessionRecordingLinkedFlag, not even an explicit nil
### GetSessionRecordingNetworkPayloadCaptureConfig

`func (o *Team) GetSessionRecordingNetworkPayloadCaptureConfig() interface{}`

GetSessionRecordingNetworkPayloadCaptureConfig returns the SessionRecordingNetworkPayloadCaptureConfig field if non-nil, zero value otherwise.

### GetSessionRecordingNetworkPayloadCaptureConfigOk

`func (o *Team) GetSessionRecordingNetworkPayloadCaptureConfigOk() (*interface{}, bool)`

GetSessionRecordingNetworkPayloadCaptureConfigOk returns a tuple with the SessionRecordingNetworkPayloadCaptureConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingNetworkPayloadCaptureConfig

`func (o *Team) SetSessionRecordingNetworkPayloadCaptureConfig(v interface{})`

SetSessionRecordingNetworkPayloadCaptureConfig sets SessionRecordingNetworkPayloadCaptureConfig field to given value.

### HasSessionRecordingNetworkPayloadCaptureConfig

`func (o *Team) HasSessionRecordingNetworkPayloadCaptureConfig() bool`

HasSessionRecordingNetworkPayloadCaptureConfig returns a boolean if a field has been set.

### SetSessionRecordingNetworkPayloadCaptureConfigNil

`func (o *Team) SetSessionRecordingNetworkPayloadCaptureConfigNil(b bool)`

 SetSessionRecordingNetworkPayloadCaptureConfigNil sets the value for SessionRecordingNetworkPayloadCaptureConfig to be an explicit nil

### UnsetSessionRecordingNetworkPayloadCaptureConfig
`func (o *Team) UnsetSessionRecordingNetworkPayloadCaptureConfig()`

UnsetSessionRecordingNetworkPayloadCaptureConfig ensures that no value is present for SessionRecordingNetworkPayloadCaptureConfig, not even an explicit nil
### GetSessionRecordingMaskingConfig

`func (o *Team) GetSessionRecordingMaskingConfig() interface{}`

GetSessionRecordingMaskingConfig returns the SessionRecordingMaskingConfig field if non-nil, zero value otherwise.

### GetSessionRecordingMaskingConfigOk

`func (o *Team) GetSessionRecordingMaskingConfigOk() (*interface{}, bool)`

GetSessionRecordingMaskingConfigOk returns a tuple with the SessionRecordingMaskingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingMaskingConfig

`func (o *Team) SetSessionRecordingMaskingConfig(v interface{})`

SetSessionRecordingMaskingConfig sets SessionRecordingMaskingConfig field to given value.

### HasSessionRecordingMaskingConfig

`func (o *Team) HasSessionRecordingMaskingConfig() bool`

HasSessionRecordingMaskingConfig returns a boolean if a field has been set.

### SetSessionRecordingMaskingConfigNil

`func (o *Team) SetSessionRecordingMaskingConfigNil(b bool)`

 SetSessionRecordingMaskingConfigNil sets the value for SessionRecordingMaskingConfig to be an explicit nil

### UnsetSessionRecordingMaskingConfig
`func (o *Team) UnsetSessionRecordingMaskingConfig()`

UnsetSessionRecordingMaskingConfig ensures that no value is present for SessionRecordingMaskingConfig, not even an explicit nil
### GetSessionRecordingUrlTriggerConfig

`func (o *Team) GetSessionRecordingUrlTriggerConfig() []*interface{}`

GetSessionRecordingUrlTriggerConfig returns the SessionRecordingUrlTriggerConfig field if non-nil, zero value otherwise.

### GetSessionRecordingUrlTriggerConfigOk

`func (o *Team) GetSessionRecordingUrlTriggerConfigOk() (*[]*interface{}, bool)`

GetSessionRecordingUrlTriggerConfigOk returns a tuple with the SessionRecordingUrlTriggerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingUrlTriggerConfig

`func (o *Team) SetSessionRecordingUrlTriggerConfig(v []*interface{})`

SetSessionRecordingUrlTriggerConfig sets SessionRecordingUrlTriggerConfig field to given value.

### HasSessionRecordingUrlTriggerConfig

`func (o *Team) HasSessionRecordingUrlTriggerConfig() bool`

HasSessionRecordingUrlTriggerConfig returns a boolean if a field has been set.

### SetSessionRecordingUrlTriggerConfigNil

`func (o *Team) SetSessionRecordingUrlTriggerConfigNil(b bool)`

 SetSessionRecordingUrlTriggerConfigNil sets the value for SessionRecordingUrlTriggerConfig to be an explicit nil

### UnsetSessionRecordingUrlTriggerConfig
`func (o *Team) UnsetSessionRecordingUrlTriggerConfig()`

UnsetSessionRecordingUrlTriggerConfig ensures that no value is present for SessionRecordingUrlTriggerConfig, not even an explicit nil
### GetSessionRecordingUrlBlocklistConfig

`func (o *Team) GetSessionRecordingUrlBlocklistConfig() []*interface{}`

GetSessionRecordingUrlBlocklistConfig returns the SessionRecordingUrlBlocklistConfig field if non-nil, zero value otherwise.

### GetSessionRecordingUrlBlocklistConfigOk

`func (o *Team) GetSessionRecordingUrlBlocklistConfigOk() (*[]*interface{}, bool)`

GetSessionRecordingUrlBlocklistConfigOk returns a tuple with the SessionRecordingUrlBlocklistConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingUrlBlocklistConfig

`func (o *Team) SetSessionRecordingUrlBlocklistConfig(v []*interface{})`

SetSessionRecordingUrlBlocklistConfig sets SessionRecordingUrlBlocklistConfig field to given value.

### HasSessionRecordingUrlBlocklistConfig

`func (o *Team) HasSessionRecordingUrlBlocklistConfig() bool`

HasSessionRecordingUrlBlocklistConfig returns a boolean if a field has been set.

### SetSessionRecordingUrlBlocklistConfigNil

`func (o *Team) SetSessionRecordingUrlBlocklistConfigNil(b bool)`

 SetSessionRecordingUrlBlocklistConfigNil sets the value for SessionRecordingUrlBlocklistConfig to be an explicit nil

### UnsetSessionRecordingUrlBlocklistConfig
`func (o *Team) UnsetSessionRecordingUrlBlocklistConfig()`

UnsetSessionRecordingUrlBlocklistConfig ensures that no value is present for SessionRecordingUrlBlocklistConfig, not even an explicit nil
### GetSessionRecordingEventTriggerConfig

`func (o *Team) GetSessionRecordingEventTriggerConfig() []*string`

GetSessionRecordingEventTriggerConfig returns the SessionRecordingEventTriggerConfig field if non-nil, zero value otherwise.

### GetSessionRecordingEventTriggerConfigOk

`func (o *Team) GetSessionRecordingEventTriggerConfigOk() (*[]*string, bool)`

GetSessionRecordingEventTriggerConfigOk returns a tuple with the SessionRecordingEventTriggerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingEventTriggerConfig

`func (o *Team) SetSessionRecordingEventTriggerConfig(v []*string)`

SetSessionRecordingEventTriggerConfig sets SessionRecordingEventTriggerConfig field to given value.

### HasSessionRecordingEventTriggerConfig

`func (o *Team) HasSessionRecordingEventTriggerConfig() bool`

HasSessionRecordingEventTriggerConfig returns a boolean if a field has been set.

### SetSessionRecordingEventTriggerConfigNil

`func (o *Team) SetSessionRecordingEventTriggerConfigNil(b bool)`

 SetSessionRecordingEventTriggerConfigNil sets the value for SessionRecordingEventTriggerConfig to be an explicit nil

### UnsetSessionRecordingEventTriggerConfig
`func (o *Team) UnsetSessionRecordingEventTriggerConfig()`

UnsetSessionRecordingEventTriggerConfig ensures that no value is present for SessionRecordingEventTriggerConfig, not even an explicit nil
### GetSessionRecordingTriggerMatchTypeConfig

`func (o *Team) GetSessionRecordingTriggerMatchTypeConfig() string`

GetSessionRecordingTriggerMatchTypeConfig returns the SessionRecordingTriggerMatchTypeConfig field if non-nil, zero value otherwise.

### GetSessionRecordingTriggerMatchTypeConfigOk

`func (o *Team) GetSessionRecordingTriggerMatchTypeConfigOk() (*string, bool)`

GetSessionRecordingTriggerMatchTypeConfigOk returns a tuple with the SessionRecordingTriggerMatchTypeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingTriggerMatchTypeConfig

`func (o *Team) SetSessionRecordingTriggerMatchTypeConfig(v string)`

SetSessionRecordingTriggerMatchTypeConfig sets SessionRecordingTriggerMatchTypeConfig field to given value.

### HasSessionRecordingTriggerMatchTypeConfig

`func (o *Team) HasSessionRecordingTriggerMatchTypeConfig() bool`

HasSessionRecordingTriggerMatchTypeConfig returns a boolean if a field has been set.

### SetSessionRecordingTriggerMatchTypeConfigNil

`func (o *Team) SetSessionRecordingTriggerMatchTypeConfigNil(b bool)`

 SetSessionRecordingTriggerMatchTypeConfigNil sets the value for SessionRecordingTriggerMatchTypeConfig to be an explicit nil

### UnsetSessionRecordingTriggerMatchTypeConfig
`func (o *Team) UnsetSessionRecordingTriggerMatchTypeConfig()`

UnsetSessionRecordingTriggerMatchTypeConfig ensures that no value is present for SessionRecordingTriggerMatchTypeConfig, not even an explicit nil
### GetSessionRecordingRetentionPeriod

`func (o *Team) GetSessionRecordingRetentionPeriod() SessionRecordingRetentionPeriodEnum`

GetSessionRecordingRetentionPeriod returns the SessionRecordingRetentionPeriod field if non-nil, zero value otherwise.

### GetSessionRecordingRetentionPeriodOk

`func (o *Team) GetSessionRecordingRetentionPeriodOk() (*SessionRecordingRetentionPeriodEnum, bool)`

GetSessionRecordingRetentionPeriodOk returns a tuple with the SessionRecordingRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingRetentionPeriod

`func (o *Team) SetSessionRecordingRetentionPeriod(v SessionRecordingRetentionPeriodEnum)`

SetSessionRecordingRetentionPeriod sets SessionRecordingRetentionPeriod field to given value.

### HasSessionRecordingRetentionPeriod

`func (o *Team) HasSessionRecordingRetentionPeriod() bool`

HasSessionRecordingRetentionPeriod returns a boolean if a field has been set.

### GetSessionReplayConfig

`func (o *Team) GetSessionReplayConfig() interface{}`

GetSessionReplayConfig returns the SessionReplayConfig field if non-nil, zero value otherwise.

### GetSessionReplayConfigOk

`func (o *Team) GetSessionReplayConfigOk() (*interface{}, bool)`

GetSessionReplayConfigOk returns a tuple with the SessionReplayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayConfig

`func (o *Team) SetSessionReplayConfig(v interface{})`

SetSessionReplayConfig sets SessionReplayConfig field to given value.

### HasSessionReplayConfig

`func (o *Team) HasSessionReplayConfig() bool`

HasSessionReplayConfig returns a boolean if a field has been set.

### SetSessionReplayConfigNil

`func (o *Team) SetSessionReplayConfigNil(b bool)`

 SetSessionReplayConfigNil sets the value for SessionReplayConfig to be an explicit nil

### UnsetSessionReplayConfig
`func (o *Team) UnsetSessionReplayConfig()`

UnsetSessionReplayConfig ensures that no value is present for SessionReplayConfig, not even an explicit nil
### GetSurveyConfig

`func (o *Team) GetSurveyConfig() interface{}`

GetSurveyConfig returns the SurveyConfig field if non-nil, zero value otherwise.

### GetSurveyConfigOk

`func (o *Team) GetSurveyConfigOk() (*interface{}, bool)`

GetSurveyConfigOk returns a tuple with the SurveyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyConfig

`func (o *Team) SetSurveyConfig(v interface{})`

SetSurveyConfig sets SurveyConfig field to given value.

### HasSurveyConfig

`func (o *Team) HasSurveyConfig() bool`

HasSurveyConfig returns a boolean if a field has been set.

### SetSurveyConfigNil

`func (o *Team) SetSurveyConfigNil(b bool)`

 SetSurveyConfigNil sets the value for SurveyConfig to be an explicit nil

### UnsetSurveyConfig
`func (o *Team) UnsetSurveyConfig()`

UnsetSurveyConfig ensures that no value is present for SurveyConfig, not even an explicit nil
### GetWeekStartDay

`func (o *Team) GetWeekStartDay() WeekStartDayEnum`

GetWeekStartDay returns the WeekStartDay field if non-nil, zero value otherwise.

### GetWeekStartDayOk

`func (o *Team) GetWeekStartDayOk() (*WeekStartDayEnum, bool)`

GetWeekStartDayOk returns a tuple with the WeekStartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekStartDay

`func (o *Team) SetWeekStartDay(v WeekStartDayEnum)`

SetWeekStartDay sets WeekStartDay field to given value.

### HasWeekStartDay

`func (o *Team) HasWeekStartDay() bool`

HasWeekStartDay returns a boolean if a field has been set.

### SetWeekStartDayNil

`func (o *Team) SetWeekStartDayNil(b bool)`

 SetWeekStartDayNil sets the value for WeekStartDay to be an explicit nil

### UnsetWeekStartDay
`func (o *Team) UnsetWeekStartDay()`

UnsetWeekStartDay ensures that no value is present for WeekStartDay, not even an explicit nil
### GetPrimaryDashboard

`func (o *Team) GetPrimaryDashboard() int32`

GetPrimaryDashboard returns the PrimaryDashboard field if non-nil, zero value otherwise.

### GetPrimaryDashboardOk

`func (o *Team) GetPrimaryDashboardOk() (*int32, bool)`

GetPrimaryDashboardOk returns a tuple with the PrimaryDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDashboard

`func (o *Team) SetPrimaryDashboard(v int32)`

SetPrimaryDashboard sets PrimaryDashboard field to given value.

### HasPrimaryDashboard

`func (o *Team) HasPrimaryDashboard() bool`

HasPrimaryDashboard returns a boolean if a field has been set.

### SetPrimaryDashboardNil

`func (o *Team) SetPrimaryDashboardNil(b bool)`

 SetPrimaryDashboardNil sets the value for PrimaryDashboard to be an explicit nil

### UnsetPrimaryDashboard
`func (o *Team) UnsetPrimaryDashboard()`

UnsetPrimaryDashboard ensures that no value is present for PrimaryDashboard, not even an explicit nil
### GetLiveEventsColumns

`func (o *Team) GetLiveEventsColumns() []string`

GetLiveEventsColumns returns the LiveEventsColumns field if non-nil, zero value otherwise.

### GetLiveEventsColumnsOk

`func (o *Team) GetLiveEventsColumnsOk() (*[]string, bool)`

GetLiveEventsColumnsOk returns a tuple with the LiveEventsColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEventsColumns

`func (o *Team) SetLiveEventsColumns(v []string)`

SetLiveEventsColumns sets LiveEventsColumns field to given value.

### HasLiveEventsColumns

`func (o *Team) HasLiveEventsColumns() bool`

HasLiveEventsColumns returns a boolean if a field has been set.

### SetLiveEventsColumnsNil

`func (o *Team) SetLiveEventsColumnsNil(b bool)`

 SetLiveEventsColumnsNil sets the value for LiveEventsColumns to be an explicit nil

### UnsetLiveEventsColumns
`func (o *Team) UnsetLiveEventsColumns()`

UnsetLiveEventsColumns ensures that no value is present for LiveEventsColumns, not even an explicit nil
### GetRecordingDomains

`func (o *Team) GetRecordingDomains() []*string`

GetRecordingDomains returns the RecordingDomains field if non-nil, zero value otherwise.

### GetRecordingDomainsOk

`func (o *Team) GetRecordingDomainsOk() (*[]*string, bool)`

GetRecordingDomainsOk returns a tuple with the RecordingDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingDomains

`func (o *Team) SetRecordingDomains(v []*string)`

SetRecordingDomains sets RecordingDomains field to given value.

### HasRecordingDomains

`func (o *Team) HasRecordingDomains() bool`

HasRecordingDomains returns a boolean if a field has been set.

### SetRecordingDomainsNil

`func (o *Team) SetRecordingDomainsNil(b bool)`

 SetRecordingDomainsNil sets the value for RecordingDomains to be an explicit nil

### UnsetRecordingDomains
`func (o *Team) UnsetRecordingDomains()`

UnsetRecordingDomains ensures that no value is present for RecordingDomains, not even an explicit nil
### GetCookielessServerHashMode

`func (o *Team) GetCookielessServerHashMode() CookielessServerHashModeEnum`

GetCookielessServerHashMode returns the CookielessServerHashMode field if non-nil, zero value otherwise.

### GetCookielessServerHashModeOk

`func (o *Team) GetCookielessServerHashModeOk() (*CookielessServerHashModeEnum, bool)`

GetCookielessServerHashModeOk returns a tuple with the CookielessServerHashMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookielessServerHashMode

`func (o *Team) SetCookielessServerHashMode(v CookielessServerHashModeEnum)`

SetCookielessServerHashMode sets CookielessServerHashMode field to given value.

### HasCookielessServerHashMode

`func (o *Team) HasCookielessServerHashMode() bool`

HasCookielessServerHashMode returns a boolean if a field has been set.

### SetCookielessServerHashModeNil

`func (o *Team) SetCookielessServerHashModeNil(b bool)`

 SetCookielessServerHashModeNil sets the value for CookielessServerHashMode to be an explicit nil

### UnsetCookielessServerHashMode
`func (o *Team) UnsetCookielessServerHashMode()`

UnsetCookielessServerHashMode ensures that no value is present for CookielessServerHashMode, not even an explicit nil
### GetHumanFriendlyComparisonPeriods

`func (o *Team) GetHumanFriendlyComparisonPeriods() bool`

GetHumanFriendlyComparisonPeriods returns the HumanFriendlyComparisonPeriods field if non-nil, zero value otherwise.

### GetHumanFriendlyComparisonPeriodsOk

`func (o *Team) GetHumanFriendlyComparisonPeriodsOk() (*bool, bool)`

GetHumanFriendlyComparisonPeriodsOk returns a tuple with the HumanFriendlyComparisonPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanFriendlyComparisonPeriods

`func (o *Team) SetHumanFriendlyComparisonPeriods(v bool)`

SetHumanFriendlyComparisonPeriods sets HumanFriendlyComparisonPeriods field to given value.

### HasHumanFriendlyComparisonPeriods

`func (o *Team) HasHumanFriendlyComparisonPeriods() bool`

HasHumanFriendlyComparisonPeriods returns a boolean if a field has been set.

### SetHumanFriendlyComparisonPeriodsNil

`func (o *Team) SetHumanFriendlyComparisonPeriodsNil(b bool)`

 SetHumanFriendlyComparisonPeriodsNil sets the value for HumanFriendlyComparisonPeriods to be an explicit nil

### UnsetHumanFriendlyComparisonPeriods
`func (o *Team) UnsetHumanFriendlyComparisonPeriods()`

UnsetHumanFriendlyComparisonPeriods ensures that no value is present for HumanFriendlyComparisonPeriods, not even an explicit nil
### GetInjectWebApps

`func (o *Team) GetInjectWebApps() bool`

GetInjectWebApps returns the InjectWebApps field if non-nil, zero value otherwise.

### GetInjectWebAppsOk

`func (o *Team) GetInjectWebAppsOk() (*bool, bool)`

GetInjectWebAppsOk returns a tuple with the InjectWebApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectWebApps

`func (o *Team) SetInjectWebApps(v bool)`

SetInjectWebApps sets InjectWebApps field to given value.

### HasInjectWebApps

`func (o *Team) HasInjectWebApps() bool`

HasInjectWebApps returns a boolean if a field has been set.

### SetInjectWebAppsNil

`func (o *Team) SetInjectWebAppsNil(b bool)`

 SetInjectWebAppsNil sets the value for InjectWebApps to be an explicit nil

### UnsetInjectWebApps
`func (o *Team) UnsetInjectWebApps()`

UnsetInjectWebApps ensures that no value is present for InjectWebApps, not even an explicit nil
### GetExtraSettings

`func (o *Team) GetExtraSettings() interface{}`

GetExtraSettings returns the ExtraSettings field if non-nil, zero value otherwise.

### GetExtraSettingsOk

`func (o *Team) GetExtraSettingsOk() (*interface{}, bool)`

GetExtraSettingsOk returns a tuple with the ExtraSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSettings

`func (o *Team) SetExtraSettings(v interface{})`

SetExtraSettings sets ExtraSettings field to given value.

### HasExtraSettings

`func (o *Team) HasExtraSettings() bool`

HasExtraSettings returns a boolean if a field has been set.

### SetExtraSettingsNil

`func (o *Team) SetExtraSettingsNil(b bool)`

 SetExtraSettingsNil sets the value for ExtraSettings to be an explicit nil

### UnsetExtraSettings
`func (o *Team) UnsetExtraSettings()`

UnsetExtraSettings ensures that no value is present for ExtraSettings, not even an explicit nil
### GetModifiers

`func (o *Team) GetModifiers() interface{}`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Team) GetModifiersOk() (*interface{}, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Team) SetModifiers(v interface{})`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Team) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### SetModifiersNil

`func (o *Team) SetModifiersNil(b bool)`

 SetModifiersNil sets the value for Modifiers to be an explicit nil

### UnsetModifiers
`func (o *Team) UnsetModifiers()`

UnsetModifiers ensures that no value is present for Modifiers, not even an explicit nil
### GetHasCompletedOnboardingFor

`func (o *Team) GetHasCompletedOnboardingFor() interface{}`

GetHasCompletedOnboardingFor returns the HasCompletedOnboardingFor field if non-nil, zero value otherwise.

### GetHasCompletedOnboardingForOk

`func (o *Team) GetHasCompletedOnboardingForOk() (*interface{}, bool)`

GetHasCompletedOnboardingForOk returns a tuple with the HasCompletedOnboardingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompletedOnboardingFor

`func (o *Team) SetHasCompletedOnboardingFor(v interface{})`

SetHasCompletedOnboardingFor sets HasCompletedOnboardingFor field to given value.

### HasHasCompletedOnboardingFor

`func (o *Team) HasHasCompletedOnboardingFor() bool`

HasHasCompletedOnboardingFor returns a boolean if a field has been set.

### SetHasCompletedOnboardingForNil

`func (o *Team) SetHasCompletedOnboardingForNil(b bool)`

 SetHasCompletedOnboardingForNil sets the value for HasCompletedOnboardingFor to be an explicit nil

### UnsetHasCompletedOnboardingFor
`func (o *Team) UnsetHasCompletedOnboardingFor()`

UnsetHasCompletedOnboardingFor ensures that no value is present for HasCompletedOnboardingFor, not even an explicit nil
### GetSurveysOptIn

`func (o *Team) GetSurveysOptIn() bool`

GetSurveysOptIn returns the SurveysOptIn field if non-nil, zero value otherwise.

### GetSurveysOptInOk

`func (o *Team) GetSurveysOptInOk() (*bool, bool)`

GetSurveysOptInOk returns a tuple with the SurveysOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveysOptIn

`func (o *Team) SetSurveysOptIn(v bool)`

SetSurveysOptIn sets SurveysOptIn field to given value.

### HasSurveysOptIn

`func (o *Team) HasSurveysOptIn() bool`

HasSurveysOptIn returns a boolean if a field has been set.

### SetSurveysOptInNil

`func (o *Team) SetSurveysOptInNil(b bool)`

 SetSurveysOptInNil sets the value for SurveysOptIn to be an explicit nil

### UnsetSurveysOptIn
`func (o *Team) UnsetSurveysOptIn()`

UnsetSurveysOptIn ensures that no value is present for SurveysOptIn, not even an explicit nil
### GetHeatmapsOptIn

`func (o *Team) GetHeatmapsOptIn() bool`

GetHeatmapsOptIn returns the HeatmapsOptIn field if non-nil, zero value otherwise.

### GetHeatmapsOptInOk

`func (o *Team) GetHeatmapsOptInOk() (*bool, bool)`

GetHeatmapsOptInOk returns a tuple with the HeatmapsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapsOptIn

`func (o *Team) SetHeatmapsOptIn(v bool)`

SetHeatmapsOptIn sets HeatmapsOptIn field to given value.

### HasHeatmapsOptIn

`func (o *Team) HasHeatmapsOptIn() bool`

HasHeatmapsOptIn returns a boolean if a field has been set.

### SetHeatmapsOptInNil

`func (o *Team) SetHeatmapsOptInNil(b bool)`

 SetHeatmapsOptInNil sets the value for HeatmapsOptIn to be an explicit nil

### UnsetHeatmapsOptIn
`func (o *Team) UnsetHeatmapsOptIn()`

UnsetHeatmapsOptIn ensures that no value is present for HeatmapsOptIn, not even an explicit nil
### GetFlagsPersistenceDefault

`func (o *Team) GetFlagsPersistenceDefault() bool`

GetFlagsPersistenceDefault returns the FlagsPersistenceDefault field if non-nil, zero value otherwise.

### GetFlagsPersistenceDefaultOk

`func (o *Team) GetFlagsPersistenceDefaultOk() (*bool, bool)`

GetFlagsPersistenceDefaultOk returns a tuple with the FlagsPersistenceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsPersistenceDefault

`func (o *Team) SetFlagsPersistenceDefault(v bool)`

SetFlagsPersistenceDefault sets FlagsPersistenceDefault field to given value.

### HasFlagsPersistenceDefault

`func (o *Team) HasFlagsPersistenceDefault() bool`

HasFlagsPersistenceDefault returns a boolean if a field has been set.

### SetFlagsPersistenceDefaultNil

`func (o *Team) SetFlagsPersistenceDefaultNil(b bool)`

 SetFlagsPersistenceDefaultNil sets the value for FlagsPersistenceDefault to be an explicit nil

### UnsetFlagsPersistenceDefault
`func (o *Team) UnsetFlagsPersistenceDefault()`

UnsetFlagsPersistenceDefault ensures that no value is present for FlagsPersistenceDefault, not even an explicit nil
### GetFeatureFlagConfirmationEnabled

`func (o *Team) GetFeatureFlagConfirmationEnabled() bool`

GetFeatureFlagConfirmationEnabled returns the FeatureFlagConfirmationEnabled field if non-nil, zero value otherwise.

### GetFeatureFlagConfirmationEnabledOk

`func (o *Team) GetFeatureFlagConfirmationEnabledOk() (*bool, bool)`

GetFeatureFlagConfirmationEnabledOk returns a tuple with the FeatureFlagConfirmationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagConfirmationEnabled

`func (o *Team) SetFeatureFlagConfirmationEnabled(v bool)`

SetFeatureFlagConfirmationEnabled sets FeatureFlagConfirmationEnabled field to given value.

### HasFeatureFlagConfirmationEnabled

`func (o *Team) HasFeatureFlagConfirmationEnabled() bool`

HasFeatureFlagConfirmationEnabled returns a boolean if a field has been set.

### SetFeatureFlagConfirmationEnabledNil

`func (o *Team) SetFeatureFlagConfirmationEnabledNil(b bool)`

 SetFeatureFlagConfirmationEnabledNil sets the value for FeatureFlagConfirmationEnabled to be an explicit nil

### UnsetFeatureFlagConfirmationEnabled
`func (o *Team) UnsetFeatureFlagConfirmationEnabled()`

UnsetFeatureFlagConfirmationEnabled ensures that no value is present for FeatureFlagConfirmationEnabled, not even an explicit nil
### GetFeatureFlagConfirmationMessage

`func (o *Team) GetFeatureFlagConfirmationMessage() string`

GetFeatureFlagConfirmationMessage returns the FeatureFlagConfirmationMessage field if non-nil, zero value otherwise.

### GetFeatureFlagConfirmationMessageOk

`func (o *Team) GetFeatureFlagConfirmationMessageOk() (*string, bool)`

GetFeatureFlagConfirmationMessageOk returns a tuple with the FeatureFlagConfirmationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagConfirmationMessage

`func (o *Team) SetFeatureFlagConfirmationMessage(v string)`

SetFeatureFlagConfirmationMessage sets FeatureFlagConfirmationMessage field to given value.

### HasFeatureFlagConfirmationMessage

`func (o *Team) HasFeatureFlagConfirmationMessage() bool`

HasFeatureFlagConfirmationMessage returns a boolean if a field has been set.

### SetFeatureFlagConfirmationMessageNil

`func (o *Team) SetFeatureFlagConfirmationMessageNil(b bool)`

 SetFeatureFlagConfirmationMessageNil sets the value for FeatureFlagConfirmationMessage to be an explicit nil

### UnsetFeatureFlagConfirmationMessage
`func (o *Team) UnsetFeatureFlagConfirmationMessage()`

UnsetFeatureFlagConfirmationMessage ensures that no value is present for FeatureFlagConfirmationMessage, not even an explicit nil
### GetDefaultEvaluationEnvironmentsEnabled

`func (o *Team) GetDefaultEvaluationEnvironmentsEnabled() bool`

GetDefaultEvaluationEnvironmentsEnabled returns the DefaultEvaluationEnvironmentsEnabled field if non-nil, zero value otherwise.

### GetDefaultEvaluationEnvironmentsEnabledOk

`func (o *Team) GetDefaultEvaluationEnvironmentsEnabledOk() (*bool, bool)`

GetDefaultEvaluationEnvironmentsEnabledOk returns a tuple with the DefaultEvaluationEnvironmentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEvaluationEnvironmentsEnabled

`func (o *Team) SetDefaultEvaluationEnvironmentsEnabled(v bool)`

SetDefaultEvaluationEnvironmentsEnabled sets DefaultEvaluationEnvironmentsEnabled field to given value.

### HasDefaultEvaluationEnvironmentsEnabled

`func (o *Team) HasDefaultEvaluationEnvironmentsEnabled() bool`

HasDefaultEvaluationEnvironmentsEnabled returns a boolean if a field has been set.

### SetDefaultEvaluationEnvironmentsEnabledNil

`func (o *Team) SetDefaultEvaluationEnvironmentsEnabledNil(b bool)`

 SetDefaultEvaluationEnvironmentsEnabledNil sets the value for DefaultEvaluationEnvironmentsEnabled to be an explicit nil

### UnsetDefaultEvaluationEnvironmentsEnabled
`func (o *Team) UnsetDefaultEvaluationEnvironmentsEnabled()`

UnsetDefaultEvaluationEnvironmentsEnabled ensures that no value is present for DefaultEvaluationEnvironmentsEnabled, not even an explicit nil
### GetCaptureDeadClicks

`func (o *Team) GetCaptureDeadClicks() bool`

GetCaptureDeadClicks returns the CaptureDeadClicks field if non-nil, zero value otherwise.

### GetCaptureDeadClicksOk

`func (o *Team) GetCaptureDeadClicksOk() (*bool, bool)`

GetCaptureDeadClicksOk returns a tuple with the CaptureDeadClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDeadClicks

`func (o *Team) SetCaptureDeadClicks(v bool)`

SetCaptureDeadClicks sets CaptureDeadClicks field to given value.

### HasCaptureDeadClicks

`func (o *Team) HasCaptureDeadClicks() bool`

HasCaptureDeadClicks returns a boolean if a field has been set.

### SetCaptureDeadClicksNil

`func (o *Team) SetCaptureDeadClicksNil(b bool)`

 SetCaptureDeadClicksNil sets the value for CaptureDeadClicks to be an explicit nil

### UnsetCaptureDeadClicks
`func (o *Team) UnsetCaptureDeadClicks()`

UnsetCaptureDeadClicks ensures that no value is present for CaptureDeadClicks, not even an explicit nil
### GetDefaultDataTheme

`func (o *Team) GetDefaultDataTheme() int32`

GetDefaultDataTheme returns the DefaultDataTheme field if non-nil, zero value otherwise.

### GetDefaultDataThemeOk

`func (o *Team) GetDefaultDataThemeOk() (*int32, bool)`

GetDefaultDataThemeOk returns a tuple with the DefaultDataTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDataTheme

`func (o *Team) SetDefaultDataTheme(v int32)`

SetDefaultDataTheme sets DefaultDataTheme field to given value.

### HasDefaultDataTheme

`func (o *Team) HasDefaultDataTheme() bool`

HasDefaultDataTheme returns a boolean if a field has been set.

### SetDefaultDataThemeNil

`func (o *Team) SetDefaultDataThemeNil(b bool)`

 SetDefaultDataThemeNil sets the value for DefaultDataTheme to be an explicit nil

### UnsetDefaultDataTheme
`func (o *Team) UnsetDefaultDataTheme()`

UnsetDefaultDataTheme ensures that no value is present for DefaultDataTheme, not even an explicit nil
### GetRevenueAnalyticsConfig

`func (o *Team) GetRevenueAnalyticsConfig() TeamRevenueAnalyticsConfig`

GetRevenueAnalyticsConfig returns the RevenueAnalyticsConfig field if non-nil, zero value otherwise.

### GetRevenueAnalyticsConfigOk

`func (o *Team) GetRevenueAnalyticsConfigOk() (*TeamRevenueAnalyticsConfig, bool)`

GetRevenueAnalyticsConfigOk returns a tuple with the RevenueAnalyticsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAnalyticsConfig

`func (o *Team) SetRevenueAnalyticsConfig(v TeamRevenueAnalyticsConfig)`

SetRevenueAnalyticsConfig sets RevenueAnalyticsConfig field to given value.

### HasRevenueAnalyticsConfig

`func (o *Team) HasRevenueAnalyticsConfig() bool`

HasRevenueAnalyticsConfig returns a boolean if a field has been set.

### GetMarketingAnalyticsConfig

`func (o *Team) GetMarketingAnalyticsConfig() TeamMarketingAnalyticsConfig`

GetMarketingAnalyticsConfig returns the MarketingAnalyticsConfig field if non-nil, zero value otherwise.

### GetMarketingAnalyticsConfigOk

`func (o *Team) GetMarketingAnalyticsConfigOk() (*TeamMarketingAnalyticsConfig, bool)`

GetMarketingAnalyticsConfigOk returns a tuple with the MarketingAnalyticsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingAnalyticsConfig

`func (o *Team) SetMarketingAnalyticsConfig(v TeamMarketingAnalyticsConfig)`

SetMarketingAnalyticsConfig sets MarketingAnalyticsConfig field to given value.

### HasMarketingAnalyticsConfig

`func (o *Team) HasMarketingAnalyticsConfig() bool`

HasMarketingAnalyticsConfig returns a boolean if a field has been set.

### GetOnboardingTasks

`func (o *Team) GetOnboardingTasks() interface{}`

GetOnboardingTasks returns the OnboardingTasks field if non-nil, zero value otherwise.

### GetOnboardingTasksOk

`func (o *Team) GetOnboardingTasksOk() (*interface{}, bool)`

GetOnboardingTasksOk returns a tuple with the OnboardingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingTasks

`func (o *Team) SetOnboardingTasks(v interface{})`

SetOnboardingTasks sets OnboardingTasks field to given value.

### HasOnboardingTasks

`func (o *Team) HasOnboardingTasks() bool`

HasOnboardingTasks returns a boolean if a field has been set.

### SetOnboardingTasksNil

`func (o *Team) SetOnboardingTasksNil(b bool)`

 SetOnboardingTasksNil sets the value for OnboardingTasks to be an explicit nil

### UnsetOnboardingTasks
`func (o *Team) UnsetOnboardingTasks()`

UnsetOnboardingTasks ensures that no value is present for OnboardingTasks, not even an explicit nil
### GetBaseCurrency

`func (o *Team) GetBaseCurrency() BaseCurrencyEnum`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *Team) GetBaseCurrencyOk() (*BaseCurrencyEnum, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *Team) SetBaseCurrency(v BaseCurrencyEnum)`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *Team) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### GetWebAnalyticsPreAggregatedTablesEnabled

`func (o *Team) GetWebAnalyticsPreAggregatedTablesEnabled() bool`

GetWebAnalyticsPreAggregatedTablesEnabled returns the WebAnalyticsPreAggregatedTablesEnabled field if non-nil, zero value otherwise.

### GetWebAnalyticsPreAggregatedTablesEnabledOk

`func (o *Team) GetWebAnalyticsPreAggregatedTablesEnabledOk() (*bool, bool)`

GetWebAnalyticsPreAggregatedTablesEnabledOk returns a tuple with the WebAnalyticsPreAggregatedTablesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAnalyticsPreAggregatedTablesEnabled

`func (o *Team) SetWebAnalyticsPreAggregatedTablesEnabled(v bool)`

SetWebAnalyticsPreAggregatedTablesEnabled sets WebAnalyticsPreAggregatedTablesEnabled field to given value.

### HasWebAnalyticsPreAggregatedTablesEnabled

`func (o *Team) HasWebAnalyticsPreAggregatedTablesEnabled() bool`

HasWebAnalyticsPreAggregatedTablesEnabled returns a boolean if a field has been set.

### SetWebAnalyticsPreAggregatedTablesEnabledNil

`func (o *Team) SetWebAnalyticsPreAggregatedTablesEnabledNil(b bool)`

 SetWebAnalyticsPreAggregatedTablesEnabledNil sets the value for WebAnalyticsPreAggregatedTablesEnabled to be an explicit nil

### UnsetWebAnalyticsPreAggregatedTablesEnabled
`func (o *Team) UnsetWebAnalyticsPreAggregatedTablesEnabled()`

UnsetWebAnalyticsPreAggregatedTablesEnabled ensures that no value is present for WebAnalyticsPreAggregatedTablesEnabled, not even an explicit nil
### GetExperimentRecalculationTime

`func (o *Team) GetExperimentRecalculationTime() string`

GetExperimentRecalculationTime returns the ExperimentRecalculationTime field if non-nil, zero value otherwise.

### GetExperimentRecalculationTimeOk

`func (o *Team) GetExperimentRecalculationTimeOk() (*string, bool)`

GetExperimentRecalculationTimeOk returns a tuple with the ExperimentRecalculationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentRecalculationTime

`func (o *Team) SetExperimentRecalculationTime(v string)`

SetExperimentRecalculationTime sets ExperimentRecalculationTime field to given value.

### HasExperimentRecalculationTime

`func (o *Team) HasExperimentRecalculationTime() bool`

HasExperimentRecalculationTime returns a boolean if a field has been set.

### SetExperimentRecalculationTimeNil

`func (o *Team) SetExperimentRecalculationTimeNil(b bool)`

 SetExperimentRecalculationTimeNil sets the value for ExperimentRecalculationTime to be an explicit nil

### UnsetExperimentRecalculationTime
`func (o *Team) UnsetExperimentRecalculationTime()`

UnsetExperimentRecalculationTime ensures that no value is present for ExperimentRecalculationTime, not even an explicit nil
### GetReceiveOrgLevelActivityLogs

`func (o *Team) GetReceiveOrgLevelActivityLogs() bool`

GetReceiveOrgLevelActivityLogs returns the ReceiveOrgLevelActivityLogs field if non-nil, zero value otherwise.

### GetReceiveOrgLevelActivityLogsOk

`func (o *Team) GetReceiveOrgLevelActivityLogsOk() (*bool, bool)`

GetReceiveOrgLevelActivityLogsOk returns a tuple with the ReceiveOrgLevelActivityLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveOrgLevelActivityLogs

`func (o *Team) SetReceiveOrgLevelActivityLogs(v bool)`

SetReceiveOrgLevelActivityLogs sets ReceiveOrgLevelActivityLogs field to given value.

### HasReceiveOrgLevelActivityLogs

`func (o *Team) HasReceiveOrgLevelActivityLogs() bool`

HasReceiveOrgLevelActivityLogs returns a boolean if a field has been set.

### SetReceiveOrgLevelActivityLogsNil

`func (o *Team) SetReceiveOrgLevelActivityLogsNil(b bool)`

 SetReceiveOrgLevelActivityLogsNil sets the value for ReceiveOrgLevelActivityLogs to be an explicit nil

### UnsetReceiveOrgLevelActivityLogs
`func (o *Team) UnsetReceiveOrgLevelActivityLogs()`

UnsetReceiveOrgLevelActivityLogs ensures that no value is present for ReceiveOrgLevelActivityLogs, not even an explicit nil
### GetEffectiveMembershipLevel

`func (o *Team) GetEffectiveMembershipLevel() EffectiveMembershipLevelEnum`

GetEffectiveMembershipLevel returns the EffectiveMembershipLevel field if non-nil, zero value otherwise.

### GetEffectiveMembershipLevelOk

`func (o *Team) GetEffectiveMembershipLevelOk() (*EffectiveMembershipLevelEnum, bool)`

GetEffectiveMembershipLevelOk returns a tuple with the EffectiveMembershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveMembershipLevel

`func (o *Team) SetEffectiveMembershipLevel(v EffectiveMembershipLevelEnum)`

SetEffectiveMembershipLevel sets EffectiveMembershipLevel field to given value.


### SetEffectiveMembershipLevelNil

`func (o *Team) SetEffectiveMembershipLevelNil(b bool)`

 SetEffectiveMembershipLevelNil sets the value for EffectiveMembershipLevel to be an explicit nil

### UnsetEffectiveMembershipLevel
`func (o *Team) UnsetEffectiveMembershipLevel()`

UnsetEffectiveMembershipLevel ensures that no value is present for EffectiveMembershipLevel, not even an explicit nil
### GetHasGroupTypes

`func (o *Team) GetHasGroupTypes() bool`

GetHasGroupTypes returns the HasGroupTypes field if non-nil, zero value otherwise.

### GetHasGroupTypesOk

`func (o *Team) GetHasGroupTypesOk() (*bool, bool)`

GetHasGroupTypesOk returns a tuple with the HasGroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGroupTypes

`func (o *Team) SetHasGroupTypes(v bool)`

SetHasGroupTypes sets HasGroupTypes field to given value.


### GetGroupTypes

`func (o *Team) GetGroupTypes() []map[string]interface{}`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *Team) GetGroupTypesOk() (*[]map[string]interface{}, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypes

`func (o *Team) SetGroupTypes(v []map[string]interface{})`

SetGroupTypes sets GroupTypes field to given value.


### GetLiveEventsToken

`func (o *Team) GetLiveEventsToken() string`

GetLiveEventsToken returns the LiveEventsToken field if non-nil, zero value otherwise.

### GetLiveEventsTokenOk

`func (o *Team) GetLiveEventsTokenOk() (*string, bool)`

GetLiveEventsTokenOk returns a tuple with the LiveEventsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEventsToken

`func (o *Team) SetLiveEventsToken(v string)`

SetLiveEventsToken sets LiveEventsToken field to given value.


### SetLiveEventsTokenNil

`func (o *Team) SetLiveEventsTokenNil(b bool)`

 SetLiveEventsTokenNil sets the value for LiveEventsToken to be an explicit nil

### UnsetLiveEventsToken
`func (o *Team) UnsetLiveEventsToken()`

UnsetLiveEventsToken ensures that no value is present for LiveEventsToken, not even an explicit nil
### GetProductIntents

`func (o *Team) GetProductIntents() string`

GetProductIntents returns the ProductIntents field if non-nil, zero value otherwise.

### GetProductIntentsOk

`func (o *Team) GetProductIntentsOk() (*string, bool)`

GetProductIntentsOk returns a tuple with the ProductIntents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIntents

`func (o *Team) SetProductIntents(v string)`

SetProductIntents sets ProductIntents field to given value.


### GetManagedViewsets

`func (o *Team) GetManagedViewsets() string`

GetManagedViewsets returns the ManagedViewsets field if non-nil, zero value otherwise.

### GetManagedViewsetsOk

`func (o *Team) GetManagedViewsetsOk() (*string, bool)`

GetManagedViewsetsOk returns a tuple with the ManagedViewsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedViewsets

`func (o *Team) SetManagedViewsets(v string)`

SetManagedViewsets sets ManagedViewsets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


