# PatchedTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**AccessControl** | Pointer to **bool** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] [readonly] 
**ProjectId** | Pointer to **int64** |  | [optional] [readonly] 
**ApiToken** | Pointer to **string** |  | [optional] [readonly] 
**SecretApiToken** | Pointer to **NullableString** |  | [optional] [readonly] 
**SecretApiTokenBackup** | Pointer to **NullableString** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**IngestedEvent** | Pointer to **bool** |  | [optional] [readonly] 
**DefaultModifiers** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**PersonOnEventsQueryingEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**UserAccessLevel** | Pointer to **NullableString** | The effective access level the user has for this object | [optional] [readonly] 
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
**EffectiveMembershipLevel** | Pointer to [**NullableEffectiveMembershipLevelEnum**](EffectiveMembershipLevelEnum.md) |  | [optional] [readonly] 
**HasGroupTypes** | Pointer to **bool** |  | [optional] [readonly] 
**GroupTypes** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**LiveEventsToken** | Pointer to **NullableString** |  | [optional] [readonly] 
**ProductIntents** | Pointer to **string** |  | [optional] [readonly] 
**ManagedViewsets** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedTeam

`func NewPatchedTeam() *PatchedTeam`

NewPatchedTeam instantiates a new PatchedTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTeamWithDefaults

`func NewPatchedTeamWithDefaults() *PatchedTeam`

NewPatchedTeamWithDefaults instantiates a new PatchedTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *PatchedTeam) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PatchedTeam) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PatchedTeam) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PatchedTeam) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *PatchedTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTeam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTeam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccessControl

`func (o *PatchedTeam) GetAccessControl() bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *PatchedTeam) GetAccessControlOk() (*bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *PatchedTeam) SetAccessControl(v bool)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *PatchedTeam) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetOrganization

`func (o *PatchedTeam) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PatchedTeam) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PatchedTeam) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PatchedTeam) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProjectId

`func (o *PatchedTeam) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PatchedTeam) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PatchedTeam) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PatchedTeam) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetApiToken

`func (o *PatchedTeam) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *PatchedTeam) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *PatchedTeam) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *PatchedTeam) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetSecretApiToken

`func (o *PatchedTeam) GetSecretApiToken() string`

GetSecretApiToken returns the SecretApiToken field if non-nil, zero value otherwise.

### GetSecretApiTokenOk

`func (o *PatchedTeam) GetSecretApiTokenOk() (*string, bool)`

GetSecretApiTokenOk returns a tuple with the SecretApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiToken

`func (o *PatchedTeam) SetSecretApiToken(v string)`

SetSecretApiToken sets SecretApiToken field to given value.

### HasSecretApiToken

`func (o *PatchedTeam) HasSecretApiToken() bool`

HasSecretApiToken returns a boolean if a field has been set.

### SetSecretApiTokenNil

`func (o *PatchedTeam) SetSecretApiTokenNil(b bool)`

 SetSecretApiTokenNil sets the value for SecretApiToken to be an explicit nil

### UnsetSecretApiToken
`func (o *PatchedTeam) UnsetSecretApiToken()`

UnsetSecretApiToken ensures that no value is present for SecretApiToken, not even an explicit nil
### GetSecretApiTokenBackup

`func (o *PatchedTeam) GetSecretApiTokenBackup() string`

GetSecretApiTokenBackup returns the SecretApiTokenBackup field if non-nil, zero value otherwise.

### GetSecretApiTokenBackupOk

`func (o *PatchedTeam) GetSecretApiTokenBackupOk() (*string, bool)`

GetSecretApiTokenBackupOk returns a tuple with the SecretApiTokenBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiTokenBackup

`func (o *PatchedTeam) SetSecretApiTokenBackup(v string)`

SetSecretApiTokenBackup sets SecretApiTokenBackup field to given value.

### HasSecretApiTokenBackup

`func (o *PatchedTeam) HasSecretApiTokenBackup() bool`

HasSecretApiTokenBackup returns a boolean if a field has been set.

### SetSecretApiTokenBackupNil

`func (o *PatchedTeam) SetSecretApiTokenBackupNil(b bool)`

 SetSecretApiTokenBackupNil sets the value for SecretApiTokenBackup to be an explicit nil

### UnsetSecretApiTokenBackup
`func (o *PatchedTeam) UnsetSecretApiTokenBackup()`

UnsetSecretApiTokenBackup ensures that no value is present for SecretApiTokenBackup, not even an explicit nil
### GetCreatedAt

`func (o *PatchedTeam) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedTeam) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedTeam) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedTeam) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedTeam) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedTeam) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedTeam) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedTeam) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIngestedEvent

`func (o *PatchedTeam) GetIngestedEvent() bool`

GetIngestedEvent returns the IngestedEvent field if non-nil, zero value otherwise.

### GetIngestedEventOk

`func (o *PatchedTeam) GetIngestedEventOk() (*bool, bool)`

GetIngestedEventOk returns a tuple with the IngestedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEvent

`func (o *PatchedTeam) SetIngestedEvent(v bool)`

SetIngestedEvent sets IngestedEvent field to given value.

### HasIngestedEvent

`func (o *PatchedTeam) HasIngestedEvent() bool`

HasIngestedEvent returns a boolean if a field has been set.

### GetDefaultModifiers

`func (o *PatchedTeam) GetDefaultModifiers() map[string]interface{}`

GetDefaultModifiers returns the DefaultModifiers field if non-nil, zero value otherwise.

### GetDefaultModifiersOk

`func (o *PatchedTeam) GetDefaultModifiersOk() (*map[string]interface{}, bool)`

GetDefaultModifiersOk returns a tuple with the DefaultModifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModifiers

`func (o *PatchedTeam) SetDefaultModifiers(v map[string]interface{})`

SetDefaultModifiers sets DefaultModifiers field to given value.

### HasDefaultModifiers

`func (o *PatchedTeam) HasDefaultModifiers() bool`

HasDefaultModifiers returns a boolean if a field has been set.

### GetPersonOnEventsQueryingEnabled

`func (o *PatchedTeam) GetPersonOnEventsQueryingEnabled() bool`

GetPersonOnEventsQueryingEnabled returns the PersonOnEventsQueryingEnabled field if non-nil, zero value otherwise.

### GetPersonOnEventsQueryingEnabledOk

`func (o *PatchedTeam) GetPersonOnEventsQueryingEnabledOk() (*bool, bool)`

GetPersonOnEventsQueryingEnabledOk returns a tuple with the PersonOnEventsQueryingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonOnEventsQueryingEnabled

`func (o *PatchedTeam) SetPersonOnEventsQueryingEnabled(v bool)`

SetPersonOnEventsQueryingEnabled sets PersonOnEventsQueryingEnabled field to given value.

### HasPersonOnEventsQueryingEnabled

`func (o *PatchedTeam) HasPersonOnEventsQueryingEnabled() bool`

HasPersonOnEventsQueryingEnabled returns a boolean if a field has been set.

### GetUserAccessLevel

`func (o *PatchedTeam) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *PatchedTeam) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *PatchedTeam) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.

### HasUserAccessLevel

`func (o *PatchedTeam) HasUserAccessLevel() bool`

HasUserAccessLevel returns a boolean if a field has been set.

### SetUserAccessLevelNil

`func (o *PatchedTeam) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *PatchedTeam) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetAppUrls

`func (o *PatchedTeam) GetAppUrls() []*string`

GetAppUrls returns the AppUrls field if non-nil, zero value otherwise.

### GetAppUrlsOk

`func (o *PatchedTeam) GetAppUrlsOk() (*[]*string, bool)`

GetAppUrlsOk returns a tuple with the AppUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrls

`func (o *PatchedTeam) SetAppUrls(v []*string)`

SetAppUrls sets AppUrls field to given value.

### HasAppUrls

`func (o *PatchedTeam) HasAppUrls() bool`

HasAppUrls returns a boolean if a field has been set.

### GetSlackIncomingWebhook

`func (o *PatchedTeam) GetSlackIncomingWebhook() string`

GetSlackIncomingWebhook returns the SlackIncomingWebhook field if non-nil, zero value otherwise.

### GetSlackIncomingWebhookOk

`func (o *PatchedTeam) GetSlackIncomingWebhookOk() (*string, bool)`

GetSlackIncomingWebhookOk returns a tuple with the SlackIncomingWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackIncomingWebhook

`func (o *PatchedTeam) SetSlackIncomingWebhook(v string)`

SetSlackIncomingWebhook sets SlackIncomingWebhook field to given value.

### HasSlackIncomingWebhook

`func (o *PatchedTeam) HasSlackIncomingWebhook() bool`

HasSlackIncomingWebhook returns a boolean if a field has been set.

### SetSlackIncomingWebhookNil

`func (o *PatchedTeam) SetSlackIncomingWebhookNil(b bool)`

 SetSlackIncomingWebhookNil sets the value for SlackIncomingWebhook to be an explicit nil

### UnsetSlackIncomingWebhook
`func (o *PatchedTeam) UnsetSlackIncomingWebhook()`

UnsetSlackIncomingWebhook ensures that no value is present for SlackIncomingWebhook, not even an explicit nil
### GetAnonymizeIps

`func (o *PatchedTeam) GetAnonymizeIps() bool`

GetAnonymizeIps returns the AnonymizeIps field if non-nil, zero value otherwise.

### GetAnonymizeIpsOk

`func (o *PatchedTeam) GetAnonymizeIpsOk() (*bool, bool)`

GetAnonymizeIpsOk returns a tuple with the AnonymizeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeIps

`func (o *PatchedTeam) SetAnonymizeIps(v bool)`

SetAnonymizeIps sets AnonymizeIps field to given value.

### HasAnonymizeIps

`func (o *PatchedTeam) HasAnonymizeIps() bool`

HasAnonymizeIps returns a boolean if a field has been set.

### GetCompletedSnippetOnboarding

`func (o *PatchedTeam) GetCompletedSnippetOnboarding() bool`

GetCompletedSnippetOnboarding returns the CompletedSnippetOnboarding field if non-nil, zero value otherwise.

### GetCompletedSnippetOnboardingOk

`func (o *PatchedTeam) GetCompletedSnippetOnboardingOk() (*bool, bool)`

GetCompletedSnippetOnboardingOk returns a tuple with the CompletedSnippetOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedSnippetOnboarding

`func (o *PatchedTeam) SetCompletedSnippetOnboarding(v bool)`

SetCompletedSnippetOnboarding sets CompletedSnippetOnboarding field to given value.

### HasCompletedSnippetOnboarding

`func (o *PatchedTeam) HasCompletedSnippetOnboarding() bool`

HasCompletedSnippetOnboarding returns a boolean if a field has been set.

### GetTestAccountFilters

`func (o *PatchedTeam) GetTestAccountFilters() interface{}`

GetTestAccountFilters returns the TestAccountFilters field if non-nil, zero value otherwise.

### GetTestAccountFiltersOk

`func (o *PatchedTeam) GetTestAccountFiltersOk() (*interface{}, bool)`

GetTestAccountFiltersOk returns a tuple with the TestAccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccountFilters

`func (o *PatchedTeam) SetTestAccountFilters(v interface{})`

SetTestAccountFilters sets TestAccountFilters field to given value.

### HasTestAccountFilters

`func (o *PatchedTeam) HasTestAccountFilters() bool`

HasTestAccountFilters returns a boolean if a field has been set.

### SetTestAccountFiltersNil

`func (o *PatchedTeam) SetTestAccountFiltersNil(b bool)`

 SetTestAccountFiltersNil sets the value for TestAccountFilters to be an explicit nil

### UnsetTestAccountFilters
`func (o *PatchedTeam) UnsetTestAccountFilters()`

UnsetTestAccountFilters ensures that no value is present for TestAccountFilters, not even an explicit nil
### GetTestAccountFiltersDefaultChecked

`func (o *PatchedTeam) GetTestAccountFiltersDefaultChecked() bool`

GetTestAccountFiltersDefaultChecked returns the TestAccountFiltersDefaultChecked field if non-nil, zero value otherwise.

### GetTestAccountFiltersDefaultCheckedOk

`func (o *PatchedTeam) GetTestAccountFiltersDefaultCheckedOk() (*bool, bool)`

GetTestAccountFiltersDefaultCheckedOk returns a tuple with the TestAccountFiltersDefaultChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccountFiltersDefaultChecked

`func (o *PatchedTeam) SetTestAccountFiltersDefaultChecked(v bool)`

SetTestAccountFiltersDefaultChecked sets TestAccountFiltersDefaultChecked field to given value.

### HasTestAccountFiltersDefaultChecked

`func (o *PatchedTeam) HasTestAccountFiltersDefaultChecked() bool`

HasTestAccountFiltersDefaultChecked returns a boolean if a field has been set.

### SetTestAccountFiltersDefaultCheckedNil

`func (o *PatchedTeam) SetTestAccountFiltersDefaultCheckedNil(b bool)`

 SetTestAccountFiltersDefaultCheckedNil sets the value for TestAccountFiltersDefaultChecked to be an explicit nil

### UnsetTestAccountFiltersDefaultChecked
`func (o *PatchedTeam) UnsetTestAccountFiltersDefaultChecked()`

UnsetTestAccountFiltersDefaultChecked ensures that no value is present for TestAccountFiltersDefaultChecked, not even an explicit nil
### GetPathCleaningFilters

`func (o *PatchedTeam) GetPathCleaningFilters() interface{}`

GetPathCleaningFilters returns the PathCleaningFilters field if non-nil, zero value otherwise.

### GetPathCleaningFiltersOk

`func (o *PatchedTeam) GetPathCleaningFiltersOk() (*interface{}, bool)`

GetPathCleaningFiltersOk returns a tuple with the PathCleaningFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathCleaningFilters

`func (o *PatchedTeam) SetPathCleaningFilters(v interface{})`

SetPathCleaningFilters sets PathCleaningFilters field to given value.

### HasPathCleaningFilters

`func (o *PatchedTeam) HasPathCleaningFilters() bool`

HasPathCleaningFilters returns a boolean if a field has been set.

### SetPathCleaningFiltersNil

`func (o *PatchedTeam) SetPathCleaningFiltersNil(b bool)`

 SetPathCleaningFiltersNil sets the value for PathCleaningFilters to be an explicit nil

### UnsetPathCleaningFilters
`func (o *PatchedTeam) UnsetPathCleaningFilters()`

UnsetPathCleaningFilters ensures that no value is present for PathCleaningFilters, not even an explicit nil
### GetIsDemo

`func (o *PatchedTeam) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *PatchedTeam) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *PatchedTeam) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *PatchedTeam) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetTimezone

`func (o *PatchedTeam) GetTimezone() TimezoneEnum`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *PatchedTeam) GetTimezoneOk() (*TimezoneEnum, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *PatchedTeam) SetTimezone(v TimezoneEnum)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *PatchedTeam) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDataAttributes

`func (o *PatchedTeam) GetDataAttributes() interface{}`

GetDataAttributes returns the DataAttributes field if non-nil, zero value otherwise.

### GetDataAttributesOk

`func (o *PatchedTeam) GetDataAttributesOk() (*interface{}, bool)`

GetDataAttributesOk returns a tuple with the DataAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAttributes

`func (o *PatchedTeam) SetDataAttributes(v interface{})`

SetDataAttributes sets DataAttributes field to given value.

### HasDataAttributes

`func (o *PatchedTeam) HasDataAttributes() bool`

HasDataAttributes returns a boolean if a field has been set.

### SetDataAttributesNil

`func (o *PatchedTeam) SetDataAttributesNil(b bool)`

 SetDataAttributesNil sets the value for DataAttributes to be an explicit nil

### UnsetDataAttributes
`func (o *PatchedTeam) UnsetDataAttributes()`

UnsetDataAttributes ensures that no value is present for DataAttributes, not even an explicit nil
### GetPersonDisplayNameProperties

`func (o *PatchedTeam) GetPersonDisplayNameProperties() []string`

GetPersonDisplayNameProperties returns the PersonDisplayNameProperties field if non-nil, zero value otherwise.

### GetPersonDisplayNamePropertiesOk

`func (o *PatchedTeam) GetPersonDisplayNamePropertiesOk() (*[]string, bool)`

GetPersonDisplayNamePropertiesOk returns a tuple with the PersonDisplayNameProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDisplayNameProperties

`func (o *PatchedTeam) SetPersonDisplayNameProperties(v []string)`

SetPersonDisplayNameProperties sets PersonDisplayNameProperties field to given value.

### HasPersonDisplayNameProperties

`func (o *PatchedTeam) HasPersonDisplayNameProperties() bool`

HasPersonDisplayNameProperties returns a boolean if a field has been set.

### SetPersonDisplayNamePropertiesNil

`func (o *PatchedTeam) SetPersonDisplayNamePropertiesNil(b bool)`

 SetPersonDisplayNamePropertiesNil sets the value for PersonDisplayNameProperties to be an explicit nil

### UnsetPersonDisplayNameProperties
`func (o *PatchedTeam) UnsetPersonDisplayNameProperties()`

UnsetPersonDisplayNameProperties ensures that no value is present for PersonDisplayNameProperties, not even an explicit nil
### GetCorrelationConfig

`func (o *PatchedTeam) GetCorrelationConfig() interface{}`

GetCorrelationConfig returns the CorrelationConfig field if non-nil, zero value otherwise.

### GetCorrelationConfigOk

`func (o *PatchedTeam) GetCorrelationConfigOk() (*interface{}, bool)`

GetCorrelationConfigOk returns a tuple with the CorrelationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationConfig

`func (o *PatchedTeam) SetCorrelationConfig(v interface{})`

SetCorrelationConfig sets CorrelationConfig field to given value.

### HasCorrelationConfig

`func (o *PatchedTeam) HasCorrelationConfig() bool`

HasCorrelationConfig returns a boolean if a field has been set.

### SetCorrelationConfigNil

`func (o *PatchedTeam) SetCorrelationConfigNil(b bool)`

 SetCorrelationConfigNil sets the value for CorrelationConfig to be an explicit nil

### UnsetCorrelationConfig
`func (o *PatchedTeam) UnsetCorrelationConfig()`

UnsetCorrelationConfig ensures that no value is present for CorrelationConfig, not even an explicit nil
### GetAutocaptureOptOut

`func (o *PatchedTeam) GetAutocaptureOptOut() bool`

GetAutocaptureOptOut returns the AutocaptureOptOut field if non-nil, zero value otherwise.

### GetAutocaptureOptOutOk

`func (o *PatchedTeam) GetAutocaptureOptOutOk() (*bool, bool)`

GetAutocaptureOptOutOk returns a tuple with the AutocaptureOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureOptOut

`func (o *PatchedTeam) SetAutocaptureOptOut(v bool)`

SetAutocaptureOptOut sets AutocaptureOptOut field to given value.

### HasAutocaptureOptOut

`func (o *PatchedTeam) HasAutocaptureOptOut() bool`

HasAutocaptureOptOut returns a boolean if a field has been set.

### SetAutocaptureOptOutNil

`func (o *PatchedTeam) SetAutocaptureOptOutNil(b bool)`

 SetAutocaptureOptOutNil sets the value for AutocaptureOptOut to be an explicit nil

### UnsetAutocaptureOptOut
`func (o *PatchedTeam) UnsetAutocaptureOptOut()`

UnsetAutocaptureOptOut ensures that no value is present for AutocaptureOptOut, not even an explicit nil
### GetAutocaptureExceptionsOptIn

`func (o *PatchedTeam) GetAutocaptureExceptionsOptIn() bool`

GetAutocaptureExceptionsOptIn returns the AutocaptureExceptionsOptIn field if non-nil, zero value otherwise.

### GetAutocaptureExceptionsOptInOk

`func (o *PatchedTeam) GetAutocaptureExceptionsOptInOk() (*bool, bool)`

GetAutocaptureExceptionsOptInOk returns a tuple with the AutocaptureExceptionsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureExceptionsOptIn

`func (o *PatchedTeam) SetAutocaptureExceptionsOptIn(v bool)`

SetAutocaptureExceptionsOptIn sets AutocaptureExceptionsOptIn field to given value.

### HasAutocaptureExceptionsOptIn

`func (o *PatchedTeam) HasAutocaptureExceptionsOptIn() bool`

HasAutocaptureExceptionsOptIn returns a boolean if a field has been set.

### SetAutocaptureExceptionsOptInNil

`func (o *PatchedTeam) SetAutocaptureExceptionsOptInNil(b bool)`

 SetAutocaptureExceptionsOptInNil sets the value for AutocaptureExceptionsOptIn to be an explicit nil

### UnsetAutocaptureExceptionsOptIn
`func (o *PatchedTeam) UnsetAutocaptureExceptionsOptIn()`

UnsetAutocaptureExceptionsOptIn ensures that no value is present for AutocaptureExceptionsOptIn, not even an explicit nil
### GetAutocaptureWebVitalsOptIn

`func (o *PatchedTeam) GetAutocaptureWebVitalsOptIn() bool`

GetAutocaptureWebVitalsOptIn returns the AutocaptureWebVitalsOptIn field if non-nil, zero value otherwise.

### GetAutocaptureWebVitalsOptInOk

`func (o *PatchedTeam) GetAutocaptureWebVitalsOptInOk() (*bool, bool)`

GetAutocaptureWebVitalsOptInOk returns a tuple with the AutocaptureWebVitalsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureWebVitalsOptIn

`func (o *PatchedTeam) SetAutocaptureWebVitalsOptIn(v bool)`

SetAutocaptureWebVitalsOptIn sets AutocaptureWebVitalsOptIn field to given value.

### HasAutocaptureWebVitalsOptIn

`func (o *PatchedTeam) HasAutocaptureWebVitalsOptIn() bool`

HasAutocaptureWebVitalsOptIn returns a boolean if a field has been set.

### SetAutocaptureWebVitalsOptInNil

`func (o *PatchedTeam) SetAutocaptureWebVitalsOptInNil(b bool)`

 SetAutocaptureWebVitalsOptInNil sets the value for AutocaptureWebVitalsOptIn to be an explicit nil

### UnsetAutocaptureWebVitalsOptIn
`func (o *PatchedTeam) UnsetAutocaptureWebVitalsOptIn()`

UnsetAutocaptureWebVitalsOptIn ensures that no value is present for AutocaptureWebVitalsOptIn, not even an explicit nil
### GetAutocaptureWebVitalsAllowedMetrics

`func (o *PatchedTeam) GetAutocaptureWebVitalsAllowedMetrics() interface{}`

GetAutocaptureWebVitalsAllowedMetrics returns the AutocaptureWebVitalsAllowedMetrics field if non-nil, zero value otherwise.

### GetAutocaptureWebVitalsAllowedMetricsOk

`func (o *PatchedTeam) GetAutocaptureWebVitalsAllowedMetricsOk() (*interface{}, bool)`

GetAutocaptureWebVitalsAllowedMetricsOk returns a tuple with the AutocaptureWebVitalsAllowedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureWebVitalsAllowedMetrics

`func (o *PatchedTeam) SetAutocaptureWebVitalsAllowedMetrics(v interface{})`

SetAutocaptureWebVitalsAllowedMetrics sets AutocaptureWebVitalsAllowedMetrics field to given value.

### HasAutocaptureWebVitalsAllowedMetrics

`func (o *PatchedTeam) HasAutocaptureWebVitalsAllowedMetrics() bool`

HasAutocaptureWebVitalsAllowedMetrics returns a boolean if a field has been set.

### SetAutocaptureWebVitalsAllowedMetricsNil

`func (o *PatchedTeam) SetAutocaptureWebVitalsAllowedMetricsNil(b bool)`

 SetAutocaptureWebVitalsAllowedMetricsNil sets the value for AutocaptureWebVitalsAllowedMetrics to be an explicit nil

### UnsetAutocaptureWebVitalsAllowedMetrics
`func (o *PatchedTeam) UnsetAutocaptureWebVitalsAllowedMetrics()`

UnsetAutocaptureWebVitalsAllowedMetrics ensures that no value is present for AutocaptureWebVitalsAllowedMetrics, not even an explicit nil
### GetAutocaptureExceptionsErrorsToIgnore

`func (o *PatchedTeam) GetAutocaptureExceptionsErrorsToIgnore() interface{}`

GetAutocaptureExceptionsErrorsToIgnore returns the AutocaptureExceptionsErrorsToIgnore field if non-nil, zero value otherwise.

### GetAutocaptureExceptionsErrorsToIgnoreOk

`func (o *PatchedTeam) GetAutocaptureExceptionsErrorsToIgnoreOk() (*interface{}, bool)`

GetAutocaptureExceptionsErrorsToIgnoreOk returns a tuple with the AutocaptureExceptionsErrorsToIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureExceptionsErrorsToIgnore

`func (o *PatchedTeam) SetAutocaptureExceptionsErrorsToIgnore(v interface{})`

SetAutocaptureExceptionsErrorsToIgnore sets AutocaptureExceptionsErrorsToIgnore field to given value.

### HasAutocaptureExceptionsErrorsToIgnore

`func (o *PatchedTeam) HasAutocaptureExceptionsErrorsToIgnore() bool`

HasAutocaptureExceptionsErrorsToIgnore returns a boolean if a field has been set.

### SetAutocaptureExceptionsErrorsToIgnoreNil

`func (o *PatchedTeam) SetAutocaptureExceptionsErrorsToIgnoreNil(b bool)`

 SetAutocaptureExceptionsErrorsToIgnoreNil sets the value for AutocaptureExceptionsErrorsToIgnore to be an explicit nil

### UnsetAutocaptureExceptionsErrorsToIgnore
`func (o *PatchedTeam) UnsetAutocaptureExceptionsErrorsToIgnore()`

UnsetAutocaptureExceptionsErrorsToIgnore ensures that no value is present for AutocaptureExceptionsErrorsToIgnore, not even an explicit nil
### GetCaptureConsoleLogOptIn

`func (o *PatchedTeam) GetCaptureConsoleLogOptIn() bool`

GetCaptureConsoleLogOptIn returns the CaptureConsoleLogOptIn field if non-nil, zero value otherwise.

### GetCaptureConsoleLogOptInOk

`func (o *PatchedTeam) GetCaptureConsoleLogOptInOk() (*bool, bool)`

GetCaptureConsoleLogOptInOk returns a tuple with the CaptureConsoleLogOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureConsoleLogOptIn

`func (o *PatchedTeam) SetCaptureConsoleLogOptIn(v bool)`

SetCaptureConsoleLogOptIn sets CaptureConsoleLogOptIn field to given value.

### HasCaptureConsoleLogOptIn

`func (o *PatchedTeam) HasCaptureConsoleLogOptIn() bool`

HasCaptureConsoleLogOptIn returns a boolean if a field has been set.

### SetCaptureConsoleLogOptInNil

`func (o *PatchedTeam) SetCaptureConsoleLogOptInNil(b bool)`

 SetCaptureConsoleLogOptInNil sets the value for CaptureConsoleLogOptIn to be an explicit nil

### UnsetCaptureConsoleLogOptIn
`func (o *PatchedTeam) UnsetCaptureConsoleLogOptIn()`

UnsetCaptureConsoleLogOptIn ensures that no value is present for CaptureConsoleLogOptIn, not even an explicit nil
### GetCapturePerformanceOptIn

`func (o *PatchedTeam) GetCapturePerformanceOptIn() bool`

GetCapturePerformanceOptIn returns the CapturePerformanceOptIn field if non-nil, zero value otherwise.

### GetCapturePerformanceOptInOk

`func (o *PatchedTeam) GetCapturePerformanceOptInOk() (*bool, bool)`

GetCapturePerformanceOptInOk returns a tuple with the CapturePerformanceOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturePerformanceOptIn

`func (o *PatchedTeam) SetCapturePerformanceOptIn(v bool)`

SetCapturePerformanceOptIn sets CapturePerformanceOptIn field to given value.

### HasCapturePerformanceOptIn

`func (o *PatchedTeam) HasCapturePerformanceOptIn() bool`

HasCapturePerformanceOptIn returns a boolean if a field has been set.

### SetCapturePerformanceOptInNil

`func (o *PatchedTeam) SetCapturePerformanceOptInNil(b bool)`

 SetCapturePerformanceOptInNil sets the value for CapturePerformanceOptIn to be an explicit nil

### UnsetCapturePerformanceOptIn
`func (o *PatchedTeam) UnsetCapturePerformanceOptIn()`

UnsetCapturePerformanceOptIn ensures that no value is present for CapturePerformanceOptIn, not even an explicit nil
### GetSessionRecordingOptIn

`func (o *PatchedTeam) GetSessionRecordingOptIn() bool`

GetSessionRecordingOptIn returns the SessionRecordingOptIn field if non-nil, zero value otherwise.

### GetSessionRecordingOptInOk

`func (o *PatchedTeam) GetSessionRecordingOptInOk() (*bool, bool)`

GetSessionRecordingOptInOk returns a tuple with the SessionRecordingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingOptIn

`func (o *PatchedTeam) SetSessionRecordingOptIn(v bool)`

SetSessionRecordingOptIn sets SessionRecordingOptIn field to given value.

### HasSessionRecordingOptIn

`func (o *PatchedTeam) HasSessionRecordingOptIn() bool`

HasSessionRecordingOptIn returns a boolean if a field has been set.

### GetSessionRecordingSampleRate

`func (o *PatchedTeam) GetSessionRecordingSampleRate() float64`

GetSessionRecordingSampleRate returns the SessionRecordingSampleRate field if non-nil, zero value otherwise.

### GetSessionRecordingSampleRateOk

`func (o *PatchedTeam) GetSessionRecordingSampleRateOk() (*float64, bool)`

GetSessionRecordingSampleRateOk returns a tuple with the SessionRecordingSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingSampleRate

`func (o *PatchedTeam) SetSessionRecordingSampleRate(v float64)`

SetSessionRecordingSampleRate sets SessionRecordingSampleRate field to given value.

### HasSessionRecordingSampleRate

`func (o *PatchedTeam) HasSessionRecordingSampleRate() bool`

HasSessionRecordingSampleRate returns a boolean if a field has been set.

### SetSessionRecordingSampleRateNil

`func (o *PatchedTeam) SetSessionRecordingSampleRateNil(b bool)`

 SetSessionRecordingSampleRateNil sets the value for SessionRecordingSampleRate to be an explicit nil

### UnsetSessionRecordingSampleRate
`func (o *PatchedTeam) UnsetSessionRecordingSampleRate()`

UnsetSessionRecordingSampleRate ensures that no value is present for SessionRecordingSampleRate, not even an explicit nil
### GetSessionRecordingMinimumDurationMilliseconds

`func (o *PatchedTeam) GetSessionRecordingMinimumDurationMilliseconds() int32`

GetSessionRecordingMinimumDurationMilliseconds returns the SessionRecordingMinimumDurationMilliseconds field if non-nil, zero value otherwise.

### GetSessionRecordingMinimumDurationMillisecondsOk

`func (o *PatchedTeam) GetSessionRecordingMinimumDurationMillisecondsOk() (*int32, bool)`

GetSessionRecordingMinimumDurationMillisecondsOk returns a tuple with the SessionRecordingMinimumDurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingMinimumDurationMilliseconds

`func (o *PatchedTeam) SetSessionRecordingMinimumDurationMilliseconds(v int32)`

SetSessionRecordingMinimumDurationMilliseconds sets SessionRecordingMinimumDurationMilliseconds field to given value.

### HasSessionRecordingMinimumDurationMilliseconds

`func (o *PatchedTeam) HasSessionRecordingMinimumDurationMilliseconds() bool`

HasSessionRecordingMinimumDurationMilliseconds returns a boolean if a field has been set.

### SetSessionRecordingMinimumDurationMillisecondsNil

`func (o *PatchedTeam) SetSessionRecordingMinimumDurationMillisecondsNil(b bool)`

 SetSessionRecordingMinimumDurationMillisecondsNil sets the value for SessionRecordingMinimumDurationMilliseconds to be an explicit nil

### UnsetSessionRecordingMinimumDurationMilliseconds
`func (o *PatchedTeam) UnsetSessionRecordingMinimumDurationMilliseconds()`

UnsetSessionRecordingMinimumDurationMilliseconds ensures that no value is present for SessionRecordingMinimumDurationMilliseconds, not even an explicit nil
### GetSessionRecordingLinkedFlag

`func (o *PatchedTeam) GetSessionRecordingLinkedFlag() interface{}`

GetSessionRecordingLinkedFlag returns the SessionRecordingLinkedFlag field if non-nil, zero value otherwise.

### GetSessionRecordingLinkedFlagOk

`func (o *PatchedTeam) GetSessionRecordingLinkedFlagOk() (*interface{}, bool)`

GetSessionRecordingLinkedFlagOk returns a tuple with the SessionRecordingLinkedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingLinkedFlag

`func (o *PatchedTeam) SetSessionRecordingLinkedFlag(v interface{})`

SetSessionRecordingLinkedFlag sets SessionRecordingLinkedFlag field to given value.

### HasSessionRecordingLinkedFlag

`func (o *PatchedTeam) HasSessionRecordingLinkedFlag() bool`

HasSessionRecordingLinkedFlag returns a boolean if a field has been set.

### SetSessionRecordingLinkedFlagNil

`func (o *PatchedTeam) SetSessionRecordingLinkedFlagNil(b bool)`

 SetSessionRecordingLinkedFlagNil sets the value for SessionRecordingLinkedFlag to be an explicit nil

### UnsetSessionRecordingLinkedFlag
`func (o *PatchedTeam) UnsetSessionRecordingLinkedFlag()`

UnsetSessionRecordingLinkedFlag ensures that no value is present for SessionRecordingLinkedFlag, not even an explicit nil
### GetSessionRecordingNetworkPayloadCaptureConfig

`func (o *PatchedTeam) GetSessionRecordingNetworkPayloadCaptureConfig() interface{}`

GetSessionRecordingNetworkPayloadCaptureConfig returns the SessionRecordingNetworkPayloadCaptureConfig field if non-nil, zero value otherwise.

### GetSessionRecordingNetworkPayloadCaptureConfigOk

`func (o *PatchedTeam) GetSessionRecordingNetworkPayloadCaptureConfigOk() (*interface{}, bool)`

GetSessionRecordingNetworkPayloadCaptureConfigOk returns a tuple with the SessionRecordingNetworkPayloadCaptureConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingNetworkPayloadCaptureConfig

`func (o *PatchedTeam) SetSessionRecordingNetworkPayloadCaptureConfig(v interface{})`

SetSessionRecordingNetworkPayloadCaptureConfig sets SessionRecordingNetworkPayloadCaptureConfig field to given value.

### HasSessionRecordingNetworkPayloadCaptureConfig

`func (o *PatchedTeam) HasSessionRecordingNetworkPayloadCaptureConfig() bool`

HasSessionRecordingNetworkPayloadCaptureConfig returns a boolean if a field has been set.

### SetSessionRecordingNetworkPayloadCaptureConfigNil

`func (o *PatchedTeam) SetSessionRecordingNetworkPayloadCaptureConfigNil(b bool)`

 SetSessionRecordingNetworkPayloadCaptureConfigNil sets the value for SessionRecordingNetworkPayloadCaptureConfig to be an explicit nil

### UnsetSessionRecordingNetworkPayloadCaptureConfig
`func (o *PatchedTeam) UnsetSessionRecordingNetworkPayloadCaptureConfig()`

UnsetSessionRecordingNetworkPayloadCaptureConfig ensures that no value is present for SessionRecordingNetworkPayloadCaptureConfig, not even an explicit nil
### GetSessionRecordingMaskingConfig

`func (o *PatchedTeam) GetSessionRecordingMaskingConfig() interface{}`

GetSessionRecordingMaskingConfig returns the SessionRecordingMaskingConfig field if non-nil, zero value otherwise.

### GetSessionRecordingMaskingConfigOk

`func (o *PatchedTeam) GetSessionRecordingMaskingConfigOk() (*interface{}, bool)`

GetSessionRecordingMaskingConfigOk returns a tuple with the SessionRecordingMaskingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingMaskingConfig

`func (o *PatchedTeam) SetSessionRecordingMaskingConfig(v interface{})`

SetSessionRecordingMaskingConfig sets SessionRecordingMaskingConfig field to given value.

### HasSessionRecordingMaskingConfig

`func (o *PatchedTeam) HasSessionRecordingMaskingConfig() bool`

HasSessionRecordingMaskingConfig returns a boolean if a field has been set.

### SetSessionRecordingMaskingConfigNil

`func (o *PatchedTeam) SetSessionRecordingMaskingConfigNil(b bool)`

 SetSessionRecordingMaskingConfigNil sets the value for SessionRecordingMaskingConfig to be an explicit nil

### UnsetSessionRecordingMaskingConfig
`func (o *PatchedTeam) UnsetSessionRecordingMaskingConfig()`

UnsetSessionRecordingMaskingConfig ensures that no value is present for SessionRecordingMaskingConfig, not even an explicit nil
### GetSessionRecordingUrlTriggerConfig

`func (o *PatchedTeam) GetSessionRecordingUrlTriggerConfig() []*interface{}`

GetSessionRecordingUrlTriggerConfig returns the SessionRecordingUrlTriggerConfig field if non-nil, zero value otherwise.

### GetSessionRecordingUrlTriggerConfigOk

`func (o *PatchedTeam) GetSessionRecordingUrlTriggerConfigOk() (*[]*interface{}, bool)`

GetSessionRecordingUrlTriggerConfigOk returns a tuple with the SessionRecordingUrlTriggerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingUrlTriggerConfig

`func (o *PatchedTeam) SetSessionRecordingUrlTriggerConfig(v []*interface{})`

SetSessionRecordingUrlTriggerConfig sets SessionRecordingUrlTriggerConfig field to given value.

### HasSessionRecordingUrlTriggerConfig

`func (o *PatchedTeam) HasSessionRecordingUrlTriggerConfig() bool`

HasSessionRecordingUrlTriggerConfig returns a boolean if a field has been set.

### SetSessionRecordingUrlTriggerConfigNil

`func (o *PatchedTeam) SetSessionRecordingUrlTriggerConfigNil(b bool)`

 SetSessionRecordingUrlTriggerConfigNil sets the value for SessionRecordingUrlTriggerConfig to be an explicit nil

### UnsetSessionRecordingUrlTriggerConfig
`func (o *PatchedTeam) UnsetSessionRecordingUrlTriggerConfig()`

UnsetSessionRecordingUrlTriggerConfig ensures that no value is present for SessionRecordingUrlTriggerConfig, not even an explicit nil
### GetSessionRecordingUrlBlocklistConfig

`func (o *PatchedTeam) GetSessionRecordingUrlBlocklistConfig() []*interface{}`

GetSessionRecordingUrlBlocklistConfig returns the SessionRecordingUrlBlocklistConfig field if non-nil, zero value otherwise.

### GetSessionRecordingUrlBlocklistConfigOk

`func (o *PatchedTeam) GetSessionRecordingUrlBlocklistConfigOk() (*[]*interface{}, bool)`

GetSessionRecordingUrlBlocklistConfigOk returns a tuple with the SessionRecordingUrlBlocklistConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingUrlBlocklistConfig

`func (o *PatchedTeam) SetSessionRecordingUrlBlocklistConfig(v []*interface{})`

SetSessionRecordingUrlBlocklistConfig sets SessionRecordingUrlBlocklistConfig field to given value.

### HasSessionRecordingUrlBlocklistConfig

`func (o *PatchedTeam) HasSessionRecordingUrlBlocklistConfig() bool`

HasSessionRecordingUrlBlocklistConfig returns a boolean if a field has been set.

### SetSessionRecordingUrlBlocklistConfigNil

`func (o *PatchedTeam) SetSessionRecordingUrlBlocklistConfigNil(b bool)`

 SetSessionRecordingUrlBlocklistConfigNil sets the value for SessionRecordingUrlBlocklistConfig to be an explicit nil

### UnsetSessionRecordingUrlBlocklistConfig
`func (o *PatchedTeam) UnsetSessionRecordingUrlBlocklistConfig()`

UnsetSessionRecordingUrlBlocklistConfig ensures that no value is present for SessionRecordingUrlBlocklistConfig, not even an explicit nil
### GetSessionRecordingEventTriggerConfig

`func (o *PatchedTeam) GetSessionRecordingEventTriggerConfig() []*string`

GetSessionRecordingEventTriggerConfig returns the SessionRecordingEventTriggerConfig field if non-nil, zero value otherwise.

### GetSessionRecordingEventTriggerConfigOk

`func (o *PatchedTeam) GetSessionRecordingEventTriggerConfigOk() (*[]*string, bool)`

GetSessionRecordingEventTriggerConfigOk returns a tuple with the SessionRecordingEventTriggerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingEventTriggerConfig

`func (o *PatchedTeam) SetSessionRecordingEventTriggerConfig(v []*string)`

SetSessionRecordingEventTriggerConfig sets SessionRecordingEventTriggerConfig field to given value.

### HasSessionRecordingEventTriggerConfig

`func (o *PatchedTeam) HasSessionRecordingEventTriggerConfig() bool`

HasSessionRecordingEventTriggerConfig returns a boolean if a field has been set.

### SetSessionRecordingEventTriggerConfigNil

`func (o *PatchedTeam) SetSessionRecordingEventTriggerConfigNil(b bool)`

 SetSessionRecordingEventTriggerConfigNil sets the value for SessionRecordingEventTriggerConfig to be an explicit nil

### UnsetSessionRecordingEventTriggerConfig
`func (o *PatchedTeam) UnsetSessionRecordingEventTriggerConfig()`

UnsetSessionRecordingEventTriggerConfig ensures that no value is present for SessionRecordingEventTriggerConfig, not even an explicit nil
### GetSessionRecordingTriggerMatchTypeConfig

`func (o *PatchedTeam) GetSessionRecordingTriggerMatchTypeConfig() string`

GetSessionRecordingTriggerMatchTypeConfig returns the SessionRecordingTriggerMatchTypeConfig field if non-nil, zero value otherwise.

### GetSessionRecordingTriggerMatchTypeConfigOk

`func (o *PatchedTeam) GetSessionRecordingTriggerMatchTypeConfigOk() (*string, bool)`

GetSessionRecordingTriggerMatchTypeConfigOk returns a tuple with the SessionRecordingTriggerMatchTypeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingTriggerMatchTypeConfig

`func (o *PatchedTeam) SetSessionRecordingTriggerMatchTypeConfig(v string)`

SetSessionRecordingTriggerMatchTypeConfig sets SessionRecordingTriggerMatchTypeConfig field to given value.

### HasSessionRecordingTriggerMatchTypeConfig

`func (o *PatchedTeam) HasSessionRecordingTriggerMatchTypeConfig() bool`

HasSessionRecordingTriggerMatchTypeConfig returns a boolean if a field has been set.

### SetSessionRecordingTriggerMatchTypeConfigNil

`func (o *PatchedTeam) SetSessionRecordingTriggerMatchTypeConfigNil(b bool)`

 SetSessionRecordingTriggerMatchTypeConfigNil sets the value for SessionRecordingTriggerMatchTypeConfig to be an explicit nil

### UnsetSessionRecordingTriggerMatchTypeConfig
`func (o *PatchedTeam) UnsetSessionRecordingTriggerMatchTypeConfig()`

UnsetSessionRecordingTriggerMatchTypeConfig ensures that no value is present for SessionRecordingTriggerMatchTypeConfig, not even an explicit nil
### GetSessionRecordingRetentionPeriod

`func (o *PatchedTeam) GetSessionRecordingRetentionPeriod() SessionRecordingRetentionPeriodEnum`

GetSessionRecordingRetentionPeriod returns the SessionRecordingRetentionPeriod field if non-nil, zero value otherwise.

### GetSessionRecordingRetentionPeriodOk

`func (o *PatchedTeam) GetSessionRecordingRetentionPeriodOk() (*SessionRecordingRetentionPeriodEnum, bool)`

GetSessionRecordingRetentionPeriodOk returns a tuple with the SessionRecordingRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingRetentionPeriod

`func (o *PatchedTeam) SetSessionRecordingRetentionPeriod(v SessionRecordingRetentionPeriodEnum)`

SetSessionRecordingRetentionPeriod sets SessionRecordingRetentionPeriod field to given value.

### HasSessionRecordingRetentionPeriod

`func (o *PatchedTeam) HasSessionRecordingRetentionPeriod() bool`

HasSessionRecordingRetentionPeriod returns a boolean if a field has been set.

### GetSessionReplayConfig

`func (o *PatchedTeam) GetSessionReplayConfig() interface{}`

GetSessionReplayConfig returns the SessionReplayConfig field if non-nil, zero value otherwise.

### GetSessionReplayConfigOk

`func (o *PatchedTeam) GetSessionReplayConfigOk() (*interface{}, bool)`

GetSessionReplayConfigOk returns a tuple with the SessionReplayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayConfig

`func (o *PatchedTeam) SetSessionReplayConfig(v interface{})`

SetSessionReplayConfig sets SessionReplayConfig field to given value.

### HasSessionReplayConfig

`func (o *PatchedTeam) HasSessionReplayConfig() bool`

HasSessionReplayConfig returns a boolean if a field has been set.

### SetSessionReplayConfigNil

`func (o *PatchedTeam) SetSessionReplayConfigNil(b bool)`

 SetSessionReplayConfigNil sets the value for SessionReplayConfig to be an explicit nil

### UnsetSessionReplayConfig
`func (o *PatchedTeam) UnsetSessionReplayConfig()`

UnsetSessionReplayConfig ensures that no value is present for SessionReplayConfig, not even an explicit nil
### GetSurveyConfig

`func (o *PatchedTeam) GetSurveyConfig() interface{}`

GetSurveyConfig returns the SurveyConfig field if non-nil, zero value otherwise.

### GetSurveyConfigOk

`func (o *PatchedTeam) GetSurveyConfigOk() (*interface{}, bool)`

GetSurveyConfigOk returns a tuple with the SurveyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyConfig

`func (o *PatchedTeam) SetSurveyConfig(v interface{})`

SetSurveyConfig sets SurveyConfig field to given value.

### HasSurveyConfig

`func (o *PatchedTeam) HasSurveyConfig() bool`

HasSurveyConfig returns a boolean if a field has been set.

### SetSurveyConfigNil

`func (o *PatchedTeam) SetSurveyConfigNil(b bool)`

 SetSurveyConfigNil sets the value for SurveyConfig to be an explicit nil

### UnsetSurveyConfig
`func (o *PatchedTeam) UnsetSurveyConfig()`

UnsetSurveyConfig ensures that no value is present for SurveyConfig, not even an explicit nil
### GetWeekStartDay

`func (o *PatchedTeam) GetWeekStartDay() WeekStartDayEnum`

GetWeekStartDay returns the WeekStartDay field if non-nil, zero value otherwise.

### GetWeekStartDayOk

`func (o *PatchedTeam) GetWeekStartDayOk() (*WeekStartDayEnum, bool)`

GetWeekStartDayOk returns a tuple with the WeekStartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekStartDay

`func (o *PatchedTeam) SetWeekStartDay(v WeekStartDayEnum)`

SetWeekStartDay sets WeekStartDay field to given value.

### HasWeekStartDay

`func (o *PatchedTeam) HasWeekStartDay() bool`

HasWeekStartDay returns a boolean if a field has been set.

### SetWeekStartDayNil

`func (o *PatchedTeam) SetWeekStartDayNil(b bool)`

 SetWeekStartDayNil sets the value for WeekStartDay to be an explicit nil

### UnsetWeekStartDay
`func (o *PatchedTeam) UnsetWeekStartDay()`

UnsetWeekStartDay ensures that no value is present for WeekStartDay, not even an explicit nil
### GetPrimaryDashboard

`func (o *PatchedTeam) GetPrimaryDashboard() int32`

GetPrimaryDashboard returns the PrimaryDashboard field if non-nil, zero value otherwise.

### GetPrimaryDashboardOk

`func (o *PatchedTeam) GetPrimaryDashboardOk() (*int32, bool)`

GetPrimaryDashboardOk returns a tuple with the PrimaryDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDashboard

`func (o *PatchedTeam) SetPrimaryDashboard(v int32)`

SetPrimaryDashboard sets PrimaryDashboard field to given value.

### HasPrimaryDashboard

`func (o *PatchedTeam) HasPrimaryDashboard() bool`

HasPrimaryDashboard returns a boolean if a field has been set.

### SetPrimaryDashboardNil

`func (o *PatchedTeam) SetPrimaryDashboardNil(b bool)`

 SetPrimaryDashboardNil sets the value for PrimaryDashboard to be an explicit nil

### UnsetPrimaryDashboard
`func (o *PatchedTeam) UnsetPrimaryDashboard()`

UnsetPrimaryDashboard ensures that no value is present for PrimaryDashboard, not even an explicit nil
### GetLiveEventsColumns

`func (o *PatchedTeam) GetLiveEventsColumns() []string`

GetLiveEventsColumns returns the LiveEventsColumns field if non-nil, zero value otherwise.

### GetLiveEventsColumnsOk

`func (o *PatchedTeam) GetLiveEventsColumnsOk() (*[]string, bool)`

GetLiveEventsColumnsOk returns a tuple with the LiveEventsColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEventsColumns

`func (o *PatchedTeam) SetLiveEventsColumns(v []string)`

SetLiveEventsColumns sets LiveEventsColumns field to given value.

### HasLiveEventsColumns

`func (o *PatchedTeam) HasLiveEventsColumns() bool`

HasLiveEventsColumns returns a boolean if a field has been set.

### SetLiveEventsColumnsNil

`func (o *PatchedTeam) SetLiveEventsColumnsNil(b bool)`

 SetLiveEventsColumnsNil sets the value for LiveEventsColumns to be an explicit nil

### UnsetLiveEventsColumns
`func (o *PatchedTeam) UnsetLiveEventsColumns()`

UnsetLiveEventsColumns ensures that no value is present for LiveEventsColumns, not even an explicit nil
### GetRecordingDomains

`func (o *PatchedTeam) GetRecordingDomains() []*string`

GetRecordingDomains returns the RecordingDomains field if non-nil, zero value otherwise.

### GetRecordingDomainsOk

`func (o *PatchedTeam) GetRecordingDomainsOk() (*[]*string, bool)`

GetRecordingDomainsOk returns a tuple with the RecordingDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingDomains

`func (o *PatchedTeam) SetRecordingDomains(v []*string)`

SetRecordingDomains sets RecordingDomains field to given value.

### HasRecordingDomains

`func (o *PatchedTeam) HasRecordingDomains() bool`

HasRecordingDomains returns a boolean if a field has been set.

### SetRecordingDomainsNil

`func (o *PatchedTeam) SetRecordingDomainsNil(b bool)`

 SetRecordingDomainsNil sets the value for RecordingDomains to be an explicit nil

### UnsetRecordingDomains
`func (o *PatchedTeam) UnsetRecordingDomains()`

UnsetRecordingDomains ensures that no value is present for RecordingDomains, not even an explicit nil
### GetCookielessServerHashMode

`func (o *PatchedTeam) GetCookielessServerHashMode() CookielessServerHashModeEnum`

GetCookielessServerHashMode returns the CookielessServerHashMode field if non-nil, zero value otherwise.

### GetCookielessServerHashModeOk

`func (o *PatchedTeam) GetCookielessServerHashModeOk() (*CookielessServerHashModeEnum, bool)`

GetCookielessServerHashModeOk returns a tuple with the CookielessServerHashMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookielessServerHashMode

`func (o *PatchedTeam) SetCookielessServerHashMode(v CookielessServerHashModeEnum)`

SetCookielessServerHashMode sets CookielessServerHashMode field to given value.

### HasCookielessServerHashMode

`func (o *PatchedTeam) HasCookielessServerHashMode() bool`

HasCookielessServerHashMode returns a boolean if a field has been set.

### SetCookielessServerHashModeNil

`func (o *PatchedTeam) SetCookielessServerHashModeNil(b bool)`

 SetCookielessServerHashModeNil sets the value for CookielessServerHashMode to be an explicit nil

### UnsetCookielessServerHashMode
`func (o *PatchedTeam) UnsetCookielessServerHashMode()`

UnsetCookielessServerHashMode ensures that no value is present for CookielessServerHashMode, not even an explicit nil
### GetHumanFriendlyComparisonPeriods

`func (o *PatchedTeam) GetHumanFriendlyComparisonPeriods() bool`

GetHumanFriendlyComparisonPeriods returns the HumanFriendlyComparisonPeriods field if non-nil, zero value otherwise.

### GetHumanFriendlyComparisonPeriodsOk

`func (o *PatchedTeam) GetHumanFriendlyComparisonPeriodsOk() (*bool, bool)`

GetHumanFriendlyComparisonPeriodsOk returns a tuple with the HumanFriendlyComparisonPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanFriendlyComparisonPeriods

`func (o *PatchedTeam) SetHumanFriendlyComparisonPeriods(v bool)`

SetHumanFriendlyComparisonPeriods sets HumanFriendlyComparisonPeriods field to given value.

### HasHumanFriendlyComparisonPeriods

`func (o *PatchedTeam) HasHumanFriendlyComparisonPeriods() bool`

HasHumanFriendlyComparisonPeriods returns a boolean if a field has been set.

### SetHumanFriendlyComparisonPeriodsNil

`func (o *PatchedTeam) SetHumanFriendlyComparisonPeriodsNil(b bool)`

 SetHumanFriendlyComparisonPeriodsNil sets the value for HumanFriendlyComparisonPeriods to be an explicit nil

### UnsetHumanFriendlyComparisonPeriods
`func (o *PatchedTeam) UnsetHumanFriendlyComparisonPeriods()`

UnsetHumanFriendlyComparisonPeriods ensures that no value is present for HumanFriendlyComparisonPeriods, not even an explicit nil
### GetInjectWebApps

`func (o *PatchedTeam) GetInjectWebApps() bool`

GetInjectWebApps returns the InjectWebApps field if non-nil, zero value otherwise.

### GetInjectWebAppsOk

`func (o *PatchedTeam) GetInjectWebAppsOk() (*bool, bool)`

GetInjectWebAppsOk returns a tuple with the InjectWebApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectWebApps

`func (o *PatchedTeam) SetInjectWebApps(v bool)`

SetInjectWebApps sets InjectWebApps field to given value.

### HasInjectWebApps

`func (o *PatchedTeam) HasInjectWebApps() bool`

HasInjectWebApps returns a boolean if a field has been set.

### SetInjectWebAppsNil

`func (o *PatchedTeam) SetInjectWebAppsNil(b bool)`

 SetInjectWebAppsNil sets the value for InjectWebApps to be an explicit nil

### UnsetInjectWebApps
`func (o *PatchedTeam) UnsetInjectWebApps()`

UnsetInjectWebApps ensures that no value is present for InjectWebApps, not even an explicit nil
### GetExtraSettings

`func (o *PatchedTeam) GetExtraSettings() interface{}`

GetExtraSettings returns the ExtraSettings field if non-nil, zero value otherwise.

### GetExtraSettingsOk

`func (o *PatchedTeam) GetExtraSettingsOk() (*interface{}, bool)`

GetExtraSettingsOk returns a tuple with the ExtraSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSettings

`func (o *PatchedTeam) SetExtraSettings(v interface{})`

SetExtraSettings sets ExtraSettings field to given value.

### HasExtraSettings

`func (o *PatchedTeam) HasExtraSettings() bool`

HasExtraSettings returns a boolean if a field has been set.

### SetExtraSettingsNil

`func (o *PatchedTeam) SetExtraSettingsNil(b bool)`

 SetExtraSettingsNil sets the value for ExtraSettings to be an explicit nil

### UnsetExtraSettings
`func (o *PatchedTeam) UnsetExtraSettings()`

UnsetExtraSettings ensures that no value is present for ExtraSettings, not even an explicit nil
### GetModifiers

`func (o *PatchedTeam) GetModifiers() interface{}`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *PatchedTeam) GetModifiersOk() (*interface{}, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *PatchedTeam) SetModifiers(v interface{})`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *PatchedTeam) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### SetModifiersNil

`func (o *PatchedTeam) SetModifiersNil(b bool)`

 SetModifiersNil sets the value for Modifiers to be an explicit nil

### UnsetModifiers
`func (o *PatchedTeam) UnsetModifiers()`

UnsetModifiers ensures that no value is present for Modifiers, not even an explicit nil
### GetHasCompletedOnboardingFor

`func (o *PatchedTeam) GetHasCompletedOnboardingFor() interface{}`

GetHasCompletedOnboardingFor returns the HasCompletedOnboardingFor field if non-nil, zero value otherwise.

### GetHasCompletedOnboardingForOk

`func (o *PatchedTeam) GetHasCompletedOnboardingForOk() (*interface{}, bool)`

GetHasCompletedOnboardingForOk returns a tuple with the HasCompletedOnboardingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompletedOnboardingFor

`func (o *PatchedTeam) SetHasCompletedOnboardingFor(v interface{})`

SetHasCompletedOnboardingFor sets HasCompletedOnboardingFor field to given value.

### HasHasCompletedOnboardingFor

`func (o *PatchedTeam) HasHasCompletedOnboardingFor() bool`

HasHasCompletedOnboardingFor returns a boolean if a field has been set.

### SetHasCompletedOnboardingForNil

`func (o *PatchedTeam) SetHasCompletedOnboardingForNil(b bool)`

 SetHasCompletedOnboardingForNil sets the value for HasCompletedOnboardingFor to be an explicit nil

### UnsetHasCompletedOnboardingFor
`func (o *PatchedTeam) UnsetHasCompletedOnboardingFor()`

UnsetHasCompletedOnboardingFor ensures that no value is present for HasCompletedOnboardingFor, not even an explicit nil
### GetSurveysOptIn

`func (o *PatchedTeam) GetSurveysOptIn() bool`

GetSurveysOptIn returns the SurveysOptIn field if non-nil, zero value otherwise.

### GetSurveysOptInOk

`func (o *PatchedTeam) GetSurveysOptInOk() (*bool, bool)`

GetSurveysOptInOk returns a tuple with the SurveysOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveysOptIn

`func (o *PatchedTeam) SetSurveysOptIn(v bool)`

SetSurveysOptIn sets SurveysOptIn field to given value.

### HasSurveysOptIn

`func (o *PatchedTeam) HasSurveysOptIn() bool`

HasSurveysOptIn returns a boolean if a field has been set.

### SetSurveysOptInNil

`func (o *PatchedTeam) SetSurveysOptInNil(b bool)`

 SetSurveysOptInNil sets the value for SurveysOptIn to be an explicit nil

### UnsetSurveysOptIn
`func (o *PatchedTeam) UnsetSurveysOptIn()`

UnsetSurveysOptIn ensures that no value is present for SurveysOptIn, not even an explicit nil
### GetHeatmapsOptIn

`func (o *PatchedTeam) GetHeatmapsOptIn() bool`

GetHeatmapsOptIn returns the HeatmapsOptIn field if non-nil, zero value otherwise.

### GetHeatmapsOptInOk

`func (o *PatchedTeam) GetHeatmapsOptInOk() (*bool, bool)`

GetHeatmapsOptInOk returns a tuple with the HeatmapsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapsOptIn

`func (o *PatchedTeam) SetHeatmapsOptIn(v bool)`

SetHeatmapsOptIn sets HeatmapsOptIn field to given value.

### HasHeatmapsOptIn

`func (o *PatchedTeam) HasHeatmapsOptIn() bool`

HasHeatmapsOptIn returns a boolean if a field has been set.

### SetHeatmapsOptInNil

`func (o *PatchedTeam) SetHeatmapsOptInNil(b bool)`

 SetHeatmapsOptInNil sets the value for HeatmapsOptIn to be an explicit nil

### UnsetHeatmapsOptIn
`func (o *PatchedTeam) UnsetHeatmapsOptIn()`

UnsetHeatmapsOptIn ensures that no value is present for HeatmapsOptIn, not even an explicit nil
### GetFlagsPersistenceDefault

`func (o *PatchedTeam) GetFlagsPersistenceDefault() bool`

GetFlagsPersistenceDefault returns the FlagsPersistenceDefault field if non-nil, zero value otherwise.

### GetFlagsPersistenceDefaultOk

`func (o *PatchedTeam) GetFlagsPersistenceDefaultOk() (*bool, bool)`

GetFlagsPersistenceDefaultOk returns a tuple with the FlagsPersistenceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsPersistenceDefault

`func (o *PatchedTeam) SetFlagsPersistenceDefault(v bool)`

SetFlagsPersistenceDefault sets FlagsPersistenceDefault field to given value.

### HasFlagsPersistenceDefault

`func (o *PatchedTeam) HasFlagsPersistenceDefault() bool`

HasFlagsPersistenceDefault returns a boolean if a field has been set.

### SetFlagsPersistenceDefaultNil

`func (o *PatchedTeam) SetFlagsPersistenceDefaultNil(b bool)`

 SetFlagsPersistenceDefaultNil sets the value for FlagsPersistenceDefault to be an explicit nil

### UnsetFlagsPersistenceDefault
`func (o *PatchedTeam) UnsetFlagsPersistenceDefault()`

UnsetFlagsPersistenceDefault ensures that no value is present for FlagsPersistenceDefault, not even an explicit nil
### GetFeatureFlagConfirmationEnabled

`func (o *PatchedTeam) GetFeatureFlagConfirmationEnabled() bool`

GetFeatureFlagConfirmationEnabled returns the FeatureFlagConfirmationEnabled field if non-nil, zero value otherwise.

### GetFeatureFlagConfirmationEnabledOk

`func (o *PatchedTeam) GetFeatureFlagConfirmationEnabledOk() (*bool, bool)`

GetFeatureFlagConfirmationEnabledOk returns a tuple with the FeatureFlagConfirmationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagConfirmationEnabled

`func (o *PatchedTeam) SetFeatureFlagConfirmationEnabled(v bool)`

SetFeatureFlagConfirmationEnabled sets FeatureFlagConfirmationEnabled field to given value.

### HasFeatureFlagConfirmationEnabled

`func (o *PatchedTeam) HasFeatureFlagConfirmationEnabled() bool`

HasFeatureFlagConfirmationEnabled returns a boolean if a field has been set.

### SetFeatureFlagConfirmationEnabledNil

`func (o *PatchedTeam) SetFeatureFlagConfirmationEnabledNil(b bool)`

 SetFeatureFlagConfirmationEnabledNil sets the value for FeatureFlagConfirmationEnabled to be an explicit nil

### UnsetFeatureFlagConfirmationEnabled
`func (o *PatchedTeam) UnsetFeatureFlagConfirmationEnabled()`

UnsetFeatureFlagConfirmationEnabled ensures that no value is present for FeatureFlagConfirmationEnabled, not even an explicit nil
### GetFeatureFlagConfirmationMessage

`func (o *PatchedTeam) GetFeatureFlagConfirmationMessage() string`

GetFeatureFlagConfirmationMessage returns the FeatureFlagConfirmationMessage field if non-nil, zero value otherwise.

### GetFeatureFlagConfirmationMessageOk

`func (o *PatchedTeam) GetFeatureFlagConfirmationMessageOk() (*string, bool)`

GetFeatureFlagConfirmationMessageOk returns a tuple with the FeatureFlagConfirmationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagConfirmationMessage

`func (o *PatchedTeam) SetFeatureFlagConfirmationMessage(v string)`

SetFeatureFlagConfirmationMessage sets FeatureFlagConfirmationMessage field to given value.

### HasFeatureFlagConfirmationMessage

`func (o *PatchedTeam) HasFeatureFlagConfirmationMessage() bool`

HasFeatureFlagConfirmationMessage returns a boolean if a field has been set.

### SetFeatureFlagConfirmationMessageNil

`func (o *PatchedTeam) SetFeatureFlagConfirmationMessageNil(b bool)`

 SetFeatureFlagConfirmationMessageNil sets the value for FeatureFlagConfirmationMessage to be an explicit nil

### UnsetFeatureFlagConfirmationMessage
`func (o *PatchedTeam) UnsetFeatureFlagConfirmationMessage()`

UnsetFeatureFlagConfirmationMessage ensures that no value is present for FeatureFlagConfirmationMessage, not even an explicit nil
### GetDefaultEvaluationEnvironmentsEnabled

`func (o *PatchedTeam) GetDefaultEvaluationEnvironmentsEnabled() bool`

GetDefaultEvaluationEnvironmentsEnabled returns the DefaultEvaluationEnvironmentsEnabled field if non-nil, zero value otherwise.

### GetDefaultEvaluationEnvironmentsEnabledOk

`func (o *PatchedTeam) GetDefaultEvaluationEnvironmentsEnabledOk() (*bool, bool)`

GetDefaultEvaluationEnvironmentsEnabledOk returns a tuple with the DefaultEvaluationEnvironmentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEvaluationEnvironmentsEnabled

`func (o *PatchedTeam) SetDefaultEvaluationEnvironmentsEnabled(v bool)`

SetDefaultEvaluationEnvironmentsEnabled sets DefaultEvaluationEnvironmentsEnabled field to given value.

### HasDefaultEvaluationEnvironmentsEnabled

`func (o *PatchedTeam) HasDefaultEvaluationEnvironmentsEnabled() bool`

HasDefaultEvaluationEnvironmentsEnabled returns a boolean if a field has been set.

### SetDefaultEvaluationEnvironmentsEnabledNil

`func (o *PatchedTeam) SetDefaultEvaluationEnvironmentsEnabledNil(b bool)`

 SetDefaultEvaluationEnvironmentsEnabledNil sets the value for DefaultEvaluationEnvironmentsEnabled to be an explicit nil

### UnsetDefaultEvaluationEnvironmentsEnabled
`func (o *PatchedTeam) UnsetDefaultEvaluationEnvironmentsEnabled()`

UnsetDefaultEvaluationEnvironmentsEnabled ensures that no value is present for DefaultEvaluationEnvironmentsEnabled, not even an explicit nil
### GetCaptureDeadClicks

`func (o *PatchedTeam) GetCaptureDeadClicks() bool`

GetCaptureDeadClicks returns the CaptureDeadClicks field if non-nil, zero value otherwise.

### GetCaptureDeadClicksOk

`func (o *PatchedTeam) GetCaptureDeadClicksOk() (*bool, bool)`

GetCaptureDeadClicksOk returns a tuple with the CaptureDeadClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDeadClicks

`func (o *PatchedTeam) SetCaptureDeadClicks(v bool)`

SetCaptureDeadClicks sets CaptureDeadClicks field to given value.

### HasCaptureDeadClicks

`func (o *PatchedTeam) HasCaptureDeadClicks() bool`

HasCaptureDeadClicks returns a boolean if a field has been set.

### SetCaptureDeadClicksNil

`func (o *PatchedTeam) SetCaptureDeadClicksNil(b bool)`

 SetCaptureDeadClicksNil sets the value for CaptureDeadClicks to be an explicit nil

### UnsetCaptureDeadClicks
`func (o *PatchedTeam) UnsetCaptureDeadClicks()`

UnsetCaptureDeadClicks ensures that no value is present for CaptureDeadClicks, not even an explicit nil
### GetDefaultDataTheme

`func (o *PatchedTeam) GetDefaultDataTheme() int32`

GetDefaultDataTheme returns the DefaultDataTheme field if non-nil, zero value otherwise.

### GetDefaultDataThemeOk

`func (o *PatchedTeam) GetDefaultDataThemeOk() (*int32, bool)`

GetDefaultDataThemeOk returns a tuple with the DefaultDataTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDataTheme

`func (o *PatchedTeam) SetDefaultDataTheme(v int32)`

SetDefaultDataTheme sets DefaultDataTheme field to given value.

### HasDefaultDataTheme

`func (o *PatchedTeam) HasDefaultDataTheme() bool`

HasDefaultDataTheme returns a boolean if a field has been set.

### SetDefaultDataThemeNil

`func (o *PatchedTeam) SetDefaultDataThemeNil(b bool)`

 SetDefaultDataThemeNil sets the value for DefaultDataTheme to be an explicit nil

### UnsetDefaultDataTheme
`func (o *PatchedTeam) UnsetDefaultDataTheme()`

UnsetDefaultDataTheme ensures that no value is present for DefaultDataTheme, not even an explicit nil
### GetRevenueAnalyticsConfig

`func (o *PatchedTeam) GetRevenueAnalyticsConfig() TeamRevenueAnalyticsConfig`

GetRevenueAnalyticsConfig returns the RevenueAnalyticsConfig field if non-nil, zero value otherwise.

### GetRevenueAnalyticsConfigOk

`func (o *PatchedTeam) GetRevenueAnalyticsConfigOk() (*TeamRevenueAnalyticsConfig, bool)`

GetRevenueAnalyticsConfigOk returns a tuple with the RevenueAnalyticsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAnalyticsConfig

`func (o *PatchedTeam) SetRevenueAnalyticsConfig(v TeamRevenueAnalyticsConfig)`

SetRevenueAnalyticsConfig sets RevenueAnalyticsConfig field to given value.

### HasRevenueAnalyticsConfig

`func (o *PatchedTeam) HasRevenueAnalyticsConfig() bool`

HasRevenueAnalyticsConfig returns a boolean if a field has been set.

### GetMarketingAnalyticsConfig

`func (o *PatchedTeam) GetMarketingAnalyticsConfig() TeamMarketingAnalyticsConfig`

GetMarketingAnalyticsConfig returns the MarketingAnalyticsConfig field if non-nil, zero value otherwise.

### GetMarketingAnalyticsConfigOk

`func (o *PatchedTeam) GetMarketingAnalyticsConfigOk() (*TeamMarketingAnalyticsConfig, bool)`

GetMarketingAnalyticsConfigOk returns a tuple with the MarketingAnalyticsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingAnalyticsConfig

`func (o *PatchedTeam) SetMarketingAnalyticsConfig(v TeamMarketingAnalyticsConfig)`

SetMarketingAnalyticsConfig sets MarketingAnalyticsConfig field to given value.

### HasMarketingAnalyticsConfig

`func (o *PatchedTeam) HasMarketingAnalyticsConfig() bool`

HasMarketingAnalyticsConfig returns a boolean if a field has been set.

### GetOnboardingTasks

`func (o *PatchedTeam) GetOnboardingTasks() interface{}`

GetOnboardingTasks returns the OnboardingTasks field if non-nil, zero value otherwise.

### GetOnboardingTasksOk

`func (o *PatchedTeam) GetOnboardingTasksOk() (*interface{}, bool)`

GetOnboardingTasksOk returns a tuple with the OnboardingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingTasks

`func (o *PatchedTeam) SetOnboardingTasks(v interface{})`

SetOnboardingTasks sets OnboardingTasks field to given value.

### HasOnboardingTasks

`func (o *PatchedTeam) HasOnboardingTasks() bool`

HasOnboardingTasks returns a boolean if a field has been set.

### SetOnboardingTasksNil

`func (o *PatchedTeam) SetOnboardingTasksNil(b bool)`

 SetOnboardingTasksNil sets the value for OnboardingTasks to be an explicit nil

### UnsetOnboardingTasks
`func (o *PatchedTeam) UnsetOnboardingTasks()`

UnsetOnboardingTasks ensures that no value is present for OnboardingTasks, not even an explicit nil
### GetBaseCurrency

`func (o *PatchedTeam) GetBaseCurrency() BaseCurrencyEnum`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *PatchedTeam) GetBaseCurrencyOk() (*BaseCurrencyEnum, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *PatchedTeam) SetBaseCurrency(v BaseCurrencyEnum)`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *PatchedTeam) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### GetWebAnalyticsPreAggregatedTablesEnabled

`func (o *PatchedTeam) GetWebAnalyticsPreAggregatedTablesEnabled() bool`

GetWebAnalyticsPreAggregatedTablesEnabled returns the WebAnalyticsPreAggregatedTablesEnabled field if non-nil, zero value otherwise.

### GetWebAnalyticsPreAggregatedTablesEnabledOk

`func (o *PatchedTeam) GetWebAnalyticsPreAggregatedTablesEnabledOk() (*bool, bool)`

GetWebAnalyticsPreAggregatedTablesEnabledOk returns a tuple with the WebAnalyticsPreAggregatedTablesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAnalyticsPreAggregatedTablesEnabled

`func (o *PatchedTeam) SetWebAnalyticsPreAggregatedTablesEnabled(v bool)`

SetWebAnalyticsPreAggregatedTablesEnabled sets WebAnalyticsPreAggregatedTablesEnabled field to given value.

### HasWebAnalyticsPreAggregatedTablesEnabled

`func (o *PatchedTeam) HasWebAnalyticsPreAggregatedTablesEnabled() bool`

HasWebAnalyticsPreAggregatedTablesEnabled returns a boolean if a field has been set.

### SetWebAnalyticsPreAggregatedTablesEnabledNil

`func (o *PatchedTeam) SetWebAnalyticsPreAggregatedTablesEnabledNil(b bool)`

 SetWebAnalyticsPreAggregatedTablesEnabledNil sets the value for WebAnalyticsPreAggregatedTablesEnabled to be an explicit nil

### UnsetWebAnalyticsPreAggregatedTablesEnabled
`func (o *PatchedTeam) UnsetWebAnalyticsPreAggregatedTablesEnabled()`

UnsetWebAnalyticsPreAggregatedTablesEnabled ensures that no value is present for WebAnalyticsPreAggregatedTablesEnabled, not even an explicit nil
### GetExperimentRecalculationTime

`func (o *PatchedTeam) GetExperimentRecalculationTime() string`

GetExperimentRecalculationTime returns the ExperimentRecalculationTime field if non-nil, zero value otherwise.

### GetExperimentRecalculationTimeOk

`func (o *PatchedTeam) GetExperimentRecalculationTimeOk() (*string, bool)`

GetExperimentRecalculationTimeOk returns a tuple with the ExperimentRecalculationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentRecalculationTime

`func (o *PatchedTeam) SetExperimentRecalculationTime(v string)`

SetExperimentRecalculationTime sets ExperimentRecalculationTime field to given value.

### HasExperimentRecalculationTime

`func (o *PatchedTeam) HasExperimentRecalculationTime() bool`

HasExperimentRecalculationTime returns a boolean if a field has been set.

### SetExperimentRecalculationTimeNil

`func (o *PatchedTeam) SetExperimentRecalculationTimeNil(b bool)`

 SetExperimentRecalculationTimeNil sets the value for ExperimentRecalculationTime to be an explicit nil

### UnsetExperimentRecalculationTime
`func (o *PatchedTeam) UnsetExperimentRecalculationTime()`

UnsetExperimentRecalculationTime ensures that no value is present for ExperimentRecalculationTime, not even an explicit nil
### GetReceiveOrgLevelActivityLogs

`func (o *PatchedTeam) GetReceiveOrgLevelActivityLogs() bool`

GetReceiveOrgLevelActivityLogs returns the ReceiveOrgLevelActivityLogs field if non-nil, zero value otherwise.

### GetReceiveOrgLevelActivityLogsOk

`func (o *PatchedTeam) GetReceiveOrgLevelActivityLogsOk() (*bool, bool)`

GetReceiveOrgLevelActivityLogsOk returns a tuple with the ReceiveOrgLevelActivityLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveOrgLevelActivityLogs

`func (o *PatchedTeam) SetReceiveOrgLevelActivityLogs(v bool)`

SetReceiveOrgLevelActivityLogs sets ReceiveOrgLevelActivityLogs field to given value.

### HasReceiveOrgLevelActivityLogs

`func (o *PatchedTeam) HasReceiveOrgLevelActivityLogs() bool`

HasReceiveOrgLevelActivityLogs returns a boolean if a field has been set.

### SetReceiveOrgLevelActivityLogsNil

`func (o *PatchedTeam) SetReceiveOrgLevelActivityLogsNil(b bool)`

 SetReceiveOrgLevelActivityLogsNil sets the value for ReceiveOrgLevelActivityLogs to be an explicit nil

### UnsetReceiveOrgLevelActivityLogs
`func (o *PatchedTeam) UnsetReceiveOrgLevelActivityLogs()`

UnsetReceiveOrgLevelActivityLogs ensures that no value is present for ReceiveOrgLevelActivityLogs, not even an explicit nil
### GetEffectiveMembershipLevel

`func (o *PatchedTeam) GetEffectiveMembershipLevel() EffectiveMembershipLevelEnum`

GetEffectiveMembershipLevel returns the EffectiveMembershipLevel field if non-nil, zero value otherwise.

### GetEffectiveMembershipLevelOk

`func (o *PatchedTeam) GetEffectiveMembershipLevelOk() (*EffectiveMembershipLevelEnum, bool)`

GetEffectiveMembershipLevelOk returns a tuple with the EffectiveMembershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveMembershipLevel

`func (o *PatchedTeam) SetEffectiveMembershipLevel(v EffectiveMembershipLevelEnum)`

SetEffectiveMembershipLevel sets EffectiveMembershipLevel field to given value.

### HasEffectiveMembershipLevel

`func (o *PatchedTeam) HasEffectiveMembershipLevel() bool`

HasEffectiveMembershipLevel returns a boolean if a field has been set.

### SetEffectiveMembershipLevelNil

`func (o *PatchedTeam) SetEffectiveMembershipLevelNil(b bool)`

 SetEffectiveMembershipLevelNil sets the value for EffectiveMembershipLevel to be an explicit nil

### UnsetEffectiveMembershipLevel
`func (o *PatchedTeam) UnsetEffectiveMembershipLevel()`

UnsetEffectiveMembershipLevel ensures that no value is present for EffectiveMembershipLevel, not even an explicit nil
### GetHasGroupTypes

`func (o *PatchedTeam) GetHasGroupTypes() bool`

GetHasGroupTypes returns the HasGroupTypes field if non-nil, zero value otherwise.

### GetHasGroupTypesOk

`func (o *PatchedTeam) GetHasGroupTypesOk() (*bool, bool)`

GetHasGroupTypesOk returns a tuple with the HasGroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGroupTypes

`func (o *PatchedTeam) SetHasGroupTypes(v bool)`

SetHasGroupTypes sets HasGroupTypes field to given value.

### HasHasGroupTypes

`func (o *PatchedTeam) HasHasGroupTypes() bool`

HasHasGroupTypes returns a boolean if a field has been set.

### GetGroupTypes

`func (o *PatchedTeam) GetGroupTypes() []map[string]interface{}`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *PatchedTeam) GetGroupTypesOk() (*[]map[string]interface{}, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypes

`func (o *PatchedTeam) SetGroupTypes(v []map[string]interface{})`

SetGroupTypes sets GroupTypes field to given value.

### HasGroupTypes

`func (o *PatchedTeam) HasGroupTypes() bool`

HasGroupTypes returns a boolean if a field has been set.

### GetLiveEventsToken

`func (o *PatchedTeam) GetLiveEventsToken() string`

GetLiveEventsToken returns the LiveEventsToken field if non-nil, zero value otherwise.

### GetLiveEventsTokenOk

`func (o *PatchedTeam) GetLiveEventsTokenOk() (*string, bool)`

GetLiveEventsTokenOk returns a tuple with the LiveEventsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEventsToken

`func (o *PatchedTeam) SetLiveEventsToken(v string)`

SetLiveEventsToken sets LiveEventsToken field to given value.

### HasLiveEventsToken

`func (o *PatchedTeam) HasLiveEventsToken() bool`

HasLiveEventsToken returns a boolean if a field has been set.

### SetLiveEventsTokenNil

`func (o *PatchedTeam) SetLiveEventsTokenNil(b bool)`

 SetLiveEventsTokenNil sets the value for LiveEventsToken to be an explicit nil

### UnsetLiveEventsToken
`func (o *PatchedTeam) UnsetLiveEventsToken()`

UnsetLiveEventsToken ensures that no value is present for LiveEventsToken, not even an explicit nil
### GetProductIntents

`func (o *PatchedTeam) GetProductIntents() string`

GetProductIntents returns the ProductIntents field if non-nil, zero value otherwise.

### GetProductIntentsOk

`func (o *PatchedTeam) GetProductIntentsOk() (*string, bool)`

GetProductIntentsOk returns a tuple with the ProductIntents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIntents

`func (o *PatchedTeam) SetProductIntents(v string)`

SetProductIntents sets ProductIntents field to given value.

### HasProductIntents

`func (o *PatchedTeam) HasProductIntents() bool`

HasProductIntents returns a boolean if a field has been set.

### GetManagedViewsets

`func (o *PatchedTeam) GetManagedViewsets() string`

GetManagedViewsets returns the ManagedViewsets field if non-nil, zero value otherwise.

### GetManagedViewsetsOk

`func (o *PatchedTeam) GetManagedViewsetsOk() (*string, bool)`

GetManagedViewsetsOk returns a tuple with the ManagedViewsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedViewsets

`func (o *PatchedTeam) SetManagedViewsets(v string)`

SetManagedViewsets sets ManagedViewsets field to given value.

### HasManagedViewsets

`func (o *PatchedTeam) HasManagedViewsets() bool`

HasManagedViewsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


