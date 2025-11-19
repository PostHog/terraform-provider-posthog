# PatchedProjectBackwardCompat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Organization** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ProductDescription** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**EffectiveMembershipLevel** | Pointer to [**NullableEffectiveMembershipLevelEnum**](EffectiveMembershipLevelEnum.md) |  | [optional] [readonly] 
**HasGroupTypes** | Pointer to **bool** |  | [optional] [readonly] 
**GroupTypes** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**LiveEventsToken** | Pointer to **NullableString** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**ApiToken** | Pointer to **string** |  | [optional] [readonly] 
**AppUrls** | Pointer to **[]string** |  | [optional] 
**SlackIncomingWebhook** | Pointer to **NullableString** |  | [optional] 
**AnonymizeIps** | Pointer to **bool** |  | [optional] 
**CompletedSnippetOnboarding** | Pointer to **bool** |  | [optional] 
**IngestedEvent** | Pointer to **bool** |  | [optional] [readonly] 
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
**SessionReplayConfig** | Pointer to **interface{}** |  | [optional] 
**SurveyConfig** | Pointer to **interface{}** |  | [optional] 
**AccessControl** | Pointer to **bool** |  | [optional] 
**WeekStartDay** | Pointer to [**NullableWeekStartDayEnum**](WeekStartDayEnum.md) |  | [optional] 
**PrimaryDashboard** | Pointer to **NullableInt32** |  | [optional] 
**LiveEventsColumns** | Pointer to **[]string** |  | [optional] 
**RecordingDomains** | Pointer to **[]string** |  | [optional] 
**PersonOnEventsQueryingEnabled** | Pointer to **string** |  | [optional] [readonly] 
**InjectWebApps** | Pointer to **NullableBool** |  | [optional] 
**ExtraSettings** | Pointer to **interface{}** |  | [optional] 
**Modifiers** | Pointer to **interface{}** |  | [optional] 
**DefaultModifiers** | Pointer to **string** |  | [optional] [readonly] 
**HasCompletedOnboardingFor** | Pointer to **interface{}** |  | [optional] 
**SurveysOptIn** | Pointer to **NullableBool** |  | [optional] 
**HeatmapsOptIn** | Pointer to **NullableBool** |  | [optional] 
**ProductIntents** | Pointer to **string** |  | [optional] [readonly] 
**FlagsPersistenceDefault** | Pointer to **NullableBool** |  | [optional] 
**SecretApiToken** | Pointer to **NullableString** |  | [optional] [readonly] 
**SecretApiTokenBackup** | Pointer to **NullableString** |  | [optional] [readonly] 
**ReceiveOrgLevelActivityLogs** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPatchedProjectBackwardCompat

`func NewPatchedProjectBackwardCompat() *PatchedProjectBackwardCompat`

NewPatchedProjectBackwardCompat instantiates a new PatchedProjectBackwardCompat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedProjectBackwardCompatWithDefaults

`func NewPatchedProjectBackwardCompatWithDefaults() *PatchedProjectBackwardCompat`

NewPatchedProjectBackwardCompatWithDefaults instantiates a new PatchedProjectBackwardCompat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedProjectBackwardCompat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedProjectBackwardCompat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedProjectBackwardCompat) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedProjectBackwardCompat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganization

`func (o *PatchedProjectBackwardCompat) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PatchedProjectBackwardCompat) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PatchedProjectBackwardCompat) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PatchedProjectBackwardCompat) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetName

`func (o *PatchedProjectBackwardCompat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedProjectBackwardCompat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedProjectBackwardCompat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedProjectBackwardCompat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductDescription

`func (o *PatchedProjectBackwardCompat) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *PatchedProjectBackwardCompat) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *PatchedProjectBackwardCompat) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *PatchedProjectBackwardCompat) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### SetProductDescriptionNil

`func (o *PatchedProjectBackwardCompat) SetProductDescriptionNil(b bool)`

 SetProductDescriptionNil sets the value for ProductDescription to be an explicit nil

### UnsetProductDescription
`func (o *PatchedProjectBackwardCompat) UnsetProductDescription()`

UnsetProductDescription ensures that no value is present for ProductDescription, not even an explicit nil
### GetCreatedAt

`func (o *PatchedProjectBackwardCompat) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedProjectBackwardCompat) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedProjectBackwardCompat) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedProjectBackwardCompat) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEffectiveMembershipLevel

`func (o *PatchedProjectBackwardCompat) GetEffectiveMembershipLevel() EffectiveMembershipLevelEnum`

GetEffectiveMembershipLevel returns the EffectiveMembershipLevel field if non-nil, zero value otherwise.

### GetEffectiveMembershipLevelOk

`func (o *PatchedProjectBackwardCompat) GetEffectiveMembershipLevelOk() (*EffectiveMembershipLevelEnum, bool)`

GetEffectiveMembershipLevelOk returns a tuple with the EffectiveMembershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveMembershipLevel

`func (o *PatchedProjectBackwardCompat) SetEffectiveMembershipLevel(v EffectiveMembershipLevelEnum)`

SetEffectiveMembershipLevel sets EffectiveMembershipLevel field to given value.

### HasEffectiveMembershipLevel

`func (o *PatchedProjectBackwardCompat) HasEffectiveMembershipLevel() bool`

HasEffectiveMembershipLevel returns a boolean if a field has been set.

### SetEffectiveMembershipLevelNil

`func (o *PatchedProjectBackwardCompat) SetEffectiveMembershipLevelNil(b bool)`

 SetEffectiveMembershipLevelNil sets the value for EffectiveMembershipLevel to be an explicit nil

### UnsetEffectiveMembershipLevel
`func (o *PatchedProjectBackwardCompat) UnsetEffectiveMembershipLevel()`

UnsetEffectiveMembershipLevel ensures that no value is present for EffectiveMembershipLevel, not even an explicit nil
### GetHasGroupTypes

`func (o *PatchedProjectBackwardCompat) GetHasGroupTypes() bool`

GetHasGroupTypes returns the HasGroupTypes field if non-nil, zero value otherwise.

### GetHasGroupTypesOk

`func (o *PatchedProjectBackwardCompat) GetHasGroupTypesOk() (*bool, bool)`

GetHasGroupTypesOk returns a tuple with the HasGroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGroupTypes

`func (o *PatchedProjectBackwardCompat) SetHasGroupTypes(v bool)`

SetHasGroupTypes sets HasGroupTypes field to given value.

### HasHasGroupTypes

`func (o *PatchedProjectBackwardCompat) HasHasGroupTypes() bool`

HasHasGroupTypes returns a boolean if a field has been set.

### GetGroupTypes

`func (o *PatchedProjectBackwardCompat) GetGroupTypes() []map[string]interface{}`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *PatchedProjectBackwardCompat) GetGroupTypesOk() (*[]map[string]interface{}, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypes

`func (o *PatchedProjectBackwardCompat) SetGroupTypes(v []map[string]interface{})`

SetGroupTypes sets GroupTypes field to given value.

### HasGroupTypes

`func (o *PatchedProjectBackwardCompat) HasGroupTypes() bool`

HasGroupTypes returns a boolean if a field has been set.

### GetLiveEventsToken

`func (o *PatchedProjectBackwardCompat) GetLiveEventsToken() string`

GetLiveEventsToken returns the LiveEventsToken field if non-nil, zero value otherwise.

### GetLiveEventsTokenOk

`func (o *PatchedProjectBackwardCompat) GetLiveEventsTokenOk() (*string, bool)`

GetLiveEventsTokenOk returns a tuple with the LiveEventsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEventsToken

`func (o *PatchedProjectBackwardCompat) SetLiveEventsToken(v string)`

SetLiveEventsToken sets LiveEventsToken field to given value.

### HasLiveEventsToken

`func (o *PatchedProjectBackwardCompat) HasLiveEventsToken() bool`

HasLiveEventsToken returns a boolean if a field has been set.

### SetLiveEventsTokenNil

`func (o *PatchedProjectBackwardCompat) SetLiveEventsTokenNil(b bool)`

 SetLiveEventsTokenNil sets the value for LiveEventsToken to be an explicit nil

### UnsetLiveEventsToken
`func (o *PatchedProjectBackwardCompat) UnsetLiveEventsToken()`

UnsetLiveEventsToken ensures that no value is present for LiveEventsToken, not even an explicit nil
### GetUpdatedAt

`func (o *PatchedProjectBackwardCompat) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedProjectBackwardCompat) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedProjectBackwardCompat) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedProjectBackwardCompat) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUuid

`func (o *PatchedProjectBackwardCompat) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PatchedProjectBackwardCompat) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PatchedProjectBackwardCompat) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PatchedProjectBackwardCompat) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetApiToken

`func (o *PatchedProjectBackwardCompat) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *PatchedProjectBackwardCompat) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *PatchedProjectBackwardCompat) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *PatchedProjectBackwardCompat) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetAppUrls

`func (o *PatchedProjectBackwardCompat) GetAppUrls() []*string`

GetAppUrls returns the AppUrls field if non-nil, zero value otherwise.

### GetAppUrlsOk

`func (o *PatchedProjectBackwardCompat) GetAppUrlsOk() (*[]*string, bool)`

GetAppUrlsOk returns a tuple with the AppUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrls

`func (o *PatchedProjectBackwardCompat) SetAppUrls(v []*string)`

SetAppUrls sets AppUrls field to given value.

### HasAppUrls

`func (o *PatchedProjectBackwardCompat) HasAppUrls() bool`

HasAppUrls returns a boolean if a field has been set.

### GetSlackIncomingWebhook

`func (o *PatchedProjectBackwardCompat) GetSlackIncomingWebhook() string`

GetSlackIncomingWebhook returns the SlackIncomingWebhook field if non-nil, zero value otherwise.

### GetSlackIncomingWebhookOk

`func (o *PatchedProjectBackwardCompat) GetSlackIncomingWebhookOk() (*string, bool)`

GetSlackIncomingWebhookOk returns a tuple with the SlackIncomingWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackIncomingWebhook

`func (o *PatchedProjectBackwardCompat) SetSlackIncomingWebhook(v string)`

SetSlackIncomingWebhook sets SlackIncomingWebhook field to given value.

### HasSlackIncomingWebhook

`func (o *PatchedProjectBackwardCompat) HasSlackIncomingWebhook() bool`

HasSlackIncomingWebhook returns a boolean if a field has been set.

### SetSlackIncomingWebhookNil

`func (o *PatchedProjectBackwardCompat) SetSlackIncomingWebhookNil(b bool)`

 SetSlackIncomingWebhookNil sets the value for SlackIncomingWebhook to be an explicit nil

### UnsetSlackIncomingWebhook
`func (o *PatchedProjectBackwardCompat) UnsetSlackIncomingWebhook()`

UnsetSlackIncomingWebhook ensures that no value is present for SlackIncomingWebhook, not even an explicit nil
### GetAnonymizeIps

`func (o *PatchedProjectBackwardCompat) GetAnonymizeIps() bool`

GetAnonymizeIps returns the AnonymizeIps field if non-nil, zero value otherwise.

### GetAnonymizeIpsOk

`func (o *PatchedProjectBackwardCompat) GetAnonymizeIpsOk() (*bool, bool)`

GetAnonymizeIpsOk returns a tuple with the AnonymizeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeIps

`func (o *PatchedProjectBackwardCompat) SetAnonymizeIps(v bool)`

SetAnonymizeIps sets AnonymizeIps field to given value.

### HasAnonymizeIps

`func (o *PatchedProjectBackwardCompat) HasAnonymizeIps() bool`

HasAnonymizeIps returns a boolean if a field has been set.

### GetCompletedSnippetOnboarding

`func (o *PatchedProjectBackwardCompat) GetCompletedSnippetOnboarding() bool`

GetCompletedSnippetOnboarding returns the CompletedSnippetOnboarding field if non-nil, zero value otherwise.

### GetCompletedSnippetOnboardingOk

`func (o *PatchedProjectBackwardCompat) GetCompletedSnippetOnboardingOk() (*bool, bool)`

GetCompletedSnippetOnboardingOk returns a tuple with the CompletedSnippetOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedSnippetOnboarding

`func (o *PatchedProjectBackwardCompat) SetCompletedSnippetOnboarding(v bool)`

SetCompletedSnippetOnboarding sets CompletedSnippetOnboarding field to given value.

### HasCompletedSnippetOnboarding

`func (o *PatchedProjectBackwardCompat) HasCompletedSnippetOnboarding() bool`

HasCompletedSnippetOnboarding returns a boolean if a field has been set.

### GetIngestedEvent

`func (o *PatchedProjectBackwardCompat) GetIngestedEvent() bool`

GetIngestedEvent returns the IngestedEvent field if non-nil, zero value otherwise.

### GetIngestedEventOk

`func (o *PatchedProjectBackwardCompat) GetIngestedEventOk() (*bool, bool)`

GetIngestedEventOk returns a tuple with the IngestedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEvent

`func (o *PatchedProjectBackwardCompat) SetIngestedEvent(v bool)`

SetIngestedEvent sets IngestedEvent field to given value.

### HasIngestedEvent

`func (o *PatchedProjectBackwardCompat) HasIngestedEvent() bool`

HasIngestedEvent returns a boolean if a field has been set.

### GetTestAccountFilters

`func (o *PatchedProjectBackwardCompat) GetTestAccountFilters() interface{}`

GetTestAccountFilters returns the TestAccountFilters field if non-nil, zero value otherwise.

### GetTestAccountFiltersOk

`func (o *PatchedProjectBackwardCompat) GetTestAccountFiltersOk() (*interface{}, bool)`

GetTestAccountFiltersOk returns a tuple with the TestAccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccountFilters

`func (o *PatchedProjectBackwardCompat) SetTestAccountFilters(v interface{})`

SetTestAccountFilters sets TestAccountFilters field to given value.

### HasTestAccountFilters

`func (o *PatchedProjectBackwardCompat) HasTestAccountFilters() bool`

HasTestAccountFilters returns a boolean if a field has been set.

### SetTestAccountFiltersNil

`func (o *PatchedProjectBackwardCompat) SetTestAccountFiltersNil(b bool)`

 SetTestAccountFiltersNil sets the value for TestAccountFilters to be an explicit nil

### UnsetTestAccountFilters
`func (o *PatchedProjectBackwardCompat) UnsetTestAccountFilters()`

UnsetTestAccountFilters ensures that no value is present for TestAccountFilters, not even an explicit nil
### GetTestAccountFiltersDefaultChecked

`func (o *PatchedProjectBackwardCompat) GetTestAccountFiltersDefaultChecked() bool`

GetTestAccountFiltersDefaultChecked returns the TestAccountFiltersDefaultChecked field if non-nil, zero value otherwise.

### GetTestAccountFiltersDefaultCheckedOk

`func (o *PatchedProjectBackwardCompat) GetTestAccountFiltersDefaultCheckedOk() (*bool, bool)`

GetTestAccountFiltersDefaultCheckedOk returns a tuple with the TestAccountFiltersDefaultChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccountFiltersDefaultChecked

`func (o *PatchedProjectBackwardCompat) SetTestAccountFiltersDefaultChecked(v bool)`

SetTestAccountFiltersDefaultChecked sets TestAccountFiltersDefaultChecked field to given value.

### HasTestAccountFiltersDefaultChecked

`func (o *PatchedProjectBackwardCompat) HasTestAccountFiltersDefaultChecked() bool`

HasTestAccountFiltersDefaultChecked returns a boolean if a field has been set.

### SetTestAccountFiltersDefaultCheckedNil

`func (o *PatchedProjectBackwardCompat) SetTestAccountFiltersDefaultCheckedNil(b bool)`

 SetTestAccountFiltersDefaultCheckedNil sets the value for TestAccountFiltersDefaultChecked to be an explicit nil

### UnsetTestAccountFiltersDefaultChecked
`func (o *PatchedProjectBackwardCompat) UnsetTestAccountFiltersDefaultChecked()`

UnsetTestAccountFiltersDefaultChecked ensures that no value is present for TestAccountFiltersDefaultChecked, not even an explicit nil
### GetPathCleaningFilters

`func (o *PatchedProjectBackwardCompat) GetPathCleaningFilters() interface{}`

GetPathCleaningFilters returns the PathCleaningFilters field if non-nil, zero value otherwise.

### GetPathCleaningFiltersOk

`func (o *PatchedProjectBackwardCompat) GetPathCleaningFiltersOk() (*interface{}, bool)`

GetPathCleaningFiltersOk returns a tuple with the PathCleaningFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathCleaningFilters

`func (o *PatchedProjectBackwardCompat) SetPathCleaningFilters(v interface{})`

SetPathCleaningFilters sets PathCleaningFilters field to given value.

### HasPathCleaningFilters

`func (o *PatchedProjectBackwardCompat) HasPathCleaningFilters() bool`

HasPathCleaningFilters returns a boolean if a field has been set.

### SetPathCleaningFiltersNil

`func (o *PatchedProjectBackwardCompat) SetPathCleaningFiltersNil(b bool)`

 SetPathCleaningFiltersNil sets the value for PathCleaningFilters to be an explicit nil

### UnsetPathCleaningFilters
`func (o *PatchedProjectBackwardCompat) UnsetPathCleaningFilters()`

UnsetPathCleaningFilters ensures that no value is present for PathCleaningFilters, not even an explicit nil
### GetIsDemo

`func (o *PatchedProjectBackwardCompat) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *PatchedProjectBackwardCompat) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *PatchedProjectBackwardCompat) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *PatchedProjectBackwardCompat) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetTimezone

`func (o *PatchedProjectBackwardCompat) GetTimezone() TimezoneEnum`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *PatchedProjectBackwardCompat) GetTimezoneOk() (*TimezoneEnum, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *PatchedProjectBackwardCompat) SetTimezone(v TimezoneEnum)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *PatchedProjectBackwardCompat) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDataAttributes

`func (o *PatchedProjectBackwardCompat) GetDataAttributes() interface{}`

GetDataAttributes returns the DataAttributes field if non-nil, zero value otherwise.

### GetDataAttributesOk

`func (o *PatchedProjectBackwardCompat) GetDataAttributesOk() (*interface{}, bool)`

GetDataAttributesOk returns a tuple with the DataAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAttributes

`func (o *PatchedProjectBackwardCompat) SetDataAttributes(v interface{})`

SetDataAttributes sets DataAttributes field to given value.

### HasDataAttributes

`func (o *PatchedProjectBackwardCompat) HasDataAttributes() bool`

HasDataAttributes returns a boolean if a field has been set.

### SetDataAttributesNil

`func (o *PatchedProjectBackwardCompat) SetDataAttributesNil(b bool)`

 SetDataAttributesNil sets the value for DataAttributes to be an explicit nil

### UnsetDataAttributes
`func (o *PatchedProjectBackwardCompat) UnsetDataAttributes()`

UnsetDataAttributes ensures that no value is present for DataAttributes, not even an explicit nil
### GetPersonDisplayNameProperties

`func (o *PatchedProjectBackwardCompat) GetPersonDisplayNameProperties() []string`

GetPersonDisplayNameProperties returns the PersonDisplayNameProperties field if non-nil, zero value otherwise.

### GetPersonDisplayNamePropertiesOk

`func (o *PatchedProjectBackwardCompat) GetPersonDisplayNamePropertiesOk() (*[]string, bool)`

GetPersonDisplayNamePropertiesOk returns a tuple with the PersonDisplayNameProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDisplayNameProperties

`func (o *PatchedProjectBackwardCompat) SetPersonDisplayNameProperties(v []string)`

SetPersonDisplayNameProperties sets PersonDisplayNameProperties field to given value.

### HasPersonDisplayNameProperties

`func (o *PatchedProjectBackwardCompat) HasPersonDisplayNameProperties() bool`

HasPersonDisplayNameProperties returns a boolean if a field has been set.

### SetPersonDisplayNamePropertiesNil

`func (o *PatchedProjectBackwardCompat) SetPersonDisplayNamePropertiesNil(b bool)`

 SetPersonDisplayNamePropertiesNil sets the value for PersonDisplayNameProperties to be an explicit nil

### UnsetPersonDisplayNameProperties
`func (o *PatchedProjectBackwardCompat) UnsetPersonDisplayNameProperties()`

UnsetPersonDisplayNameProperties ensures that no value is present for PersonDisplayNameProperties, not even an explicit nil
### GetCorrelationConfig

`func (o *PatchedProjectBackwardCompat) GetCorrelationConfig() interface{}`

GetCorrelationConfig returns the CorrelationConfig field if non-nil, zero value otherwise.

### GetCorrelationConfigOk

`func (o *PatchedProjectBackwardCompat) GetCorrelationConfigOk() (*interface{}, bool)`

GetCorrelationConfigOk returns a tuple with the CorrelationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationConfig

`func (o *PatchedProjectBackwardCompat) SetCorrelationConfig(v interface{})`

SetCorrelationConfig sets CorrelationConfig field to given value.

### HasCorrelationConfig

`func (o *PatchedProjectBackwardCompat) HasCorrelationConfig() bool`

HasCorrelationConfig returns a boolean if a field has been set.

### SetCorrelationConfigNil

`func (o *PatchedProjectBackwardCompat) SetCorrelationConfigNil(b bool)`

 SetCorrelationConfigNil sets the value for CorrelationConfig to be an explicit nil

### UnsetCorrelationConfig
`func (o *PatchedProjectBackwardCompat) UnsetCorrelationConfig()`

UnsetCorrelationConfig ensures that no value is present for CorrelationConfig, not even an explicit nil
### GetAutocaptureOptOut

`func (o *PatchedProjectBackwardCompat) GetAutocaptureOptOut() bool`

GetAutocaptureOptOut returns the AutocaptureOptOut field if non-nil, zero value otherwise.

### GetAutocaptureOptOutOk

`func (o *PatchedProjectBackwardCompat) GetAutocaptureOptOutOk() (*bool, bool)`

GetAutocaptureOptOutOk returns a tuple with the AutocaptureOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureOptOut

`func (o *PatchedProjectBackwardCompat) SetAutocaptureOptOut(v bool)`

SetAutocaptureOptOut sets AutocaptureOptOut field to given value.

### HasAutocaptureOptOut

`func (o *PatchedProjectBackwardCompat) HasAutocaptureOptOut() bool`

HasAutocaptureOptOut returns a boolean if a field has been set.

### SetAutocaptureOptOutNil

`func (o *PatchedProjectBackwardCompat) SetAutocaptureOptOutNil(b bool)`

 SetAutocaptureOptOutNil sets the value for AutocaptureOptOut to be an explicit nil

### UnsetAutocaptureOptOut
`func (o *PatchedProjectBackwardCompat) UnsetAutocaptureOptOut()`

UnsetAutocaptureOptOut ensures that no value is present for AutocaptureOptOut, not even an explicit nil
### GetAutocaptureExceptionsOptIn

`func (o *PatchedProjectBackwardCompat) GetAutocaptureExceptionsOptIn() bool`

GetAutocaptureExceptionsOptIn returns the AutocaptureExceptionsOptIn field if non-nil, zero value otherwise.

### GetAutocaptureExceptionsOptInOk

`func (o *PatchedProjectBackwardCompat) GetAutocaptureExceptionsOptInOk() (*bool, bool)`

GetAutocaptureExceptionsOptInOk returns a tuple with the AutocaptureExceptionsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureExceptionsOptIn

`func (o *PatchedProjectBackwardCompat) SetAutocaptureExceptionsOptIn(v bool)`

SetAutocaptureExceptionsOptIn sets AutocaptureExceptionsOptIn field to given value.

### HasAutocaptureExceptionsOptIn

`func (o *PatchedProjectBackwardCompat) HasAutocaptureExceptionsOptIn() bool`

HasAutocaptureExceptionsOptIn returns a boolean if a field has been set.

### SetAutocaptureExceptionsOptInNil

`func (o *PatchedProjectBackwardCompat) SetAutocaptureExceptionsOptInNil(b bool)`

 SetAutocaptureExceptionsOptInNil sets the value for AutocaptureExceptionsOptIn to be an explicit nil

### UnsetAutocaptureExceptionsOptIn
`func (o *PatchedProjectBackwardCompat) UnsetAutocaptureExceptionsOptIn()`

UnsetAutocaptureExceptionsOptIn ensures that no value is present for AutocaptureExceptionsOptIn, not even an explicit nil
### GetAutocaptureWebVitalsOptIn

`func (o *PatchedProjectBackwardCompat) GetAutocaptureWebVitalsOptIn() bool`

GetAutocaptureWebVitalsOptIn returns the AutocaptureWebVitalsOptIn field if non-nil, zero value otherwise.

### GetAutocaptureWebVitalsOptInOk

`func (o *PatchedProjectBackwardCompat) GetAutocaptureWebVitalsOptInOk() (*bool, bool)`

GetAutocaptureWebVitalsOptInOk returns a tuple with the AutocaptureWebVitalsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureWebVitalsOptIn

`func (o *PatchedProjectBackwardCompat) SetAutocaptureWebVitalsOptIn(v bool)`

SetAutocaptureWebVitalsOptIn sets AutocaptureWebVitalsOptIn field to given value.

### HasAutocaptureWebVitalsOptIn

`func (o *PatchedProjectBackwardCompat) HasAutocaptureWebVitalsOptIn() bool`

HasAutocaptureWebVitalsOptIn returns a boolean if a field has been set.

### SetAutocaptureWebVitalsOptInNil

`func (o *PatchedProjectBackwardCompat) SetAutocaptureWebVitalsOptInNil(b bool)`

 SetAutocaptureWebVitalsOptInNil sets the value for AutocaptureWebVitalsOptIn to be an explicit nil

### UnsetAutocaptureWebVitalsOptIn
`func (o *PatchedProjectBackwardCompat) UnsetAutocaptureWebVitalsOptIn()`

UnsetAutocaptureWebVitalsOptIn ensures that no value is present for AutocaptureWebVitalsOptIn, not even an explicit nil
### GetAutocaptureWebVitalsAllowedMetrics

`func (o *PatchedProjectBackwardCompat) GetAutocaptureWebVitalsAllowedMetrics() interface{}`

GetAutocaptureWebVitalsAllowedMetrics returns the AutocaptureWebVitalsAllowedMetrics field if non-nil, zero value otherwise.

### GetAutocaptureWebVitalsAllowedMetricsOk

`func (o *PatchedProjectBackwardCompat) GetAutocaptureWebVitalsAllowedMetricsOk() (*interface{}, bool)`

GetAutocaptureWebVitalsAllowedMetricsOk returns a tuple with the AutocaptureWebVitalsAllowedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureWebVitalsAllowedMetrics

`func (o *PatchedProjectBackwardCompat) SetAutocaptureWebVitalsAllowedMetrics(v interface{})`

SetAutocaptureWebVitalsAllowedMetrics sets AutocaptureWebVitalsAllowedMetrics field to given value.

### HasAutocaptureWebVitalsAllowedMetrics

`func (o *PatchedProjectBackwardCompat) HasAutocaptureWebVitalsAllowedMetrics() bool`

HasAutocaptureWebVitalsAllowedMetrics returns a boolean if a field has been set.

### SetAutocaptureWebVitalsAllowedMetricsNil

`func (o *PatchedProjectBackwardCompat) SetAutocaptureWebVitalsAllowedMetricsNil(b bool)`

 SetAutocaptureWebVitalsAllowedMetricsNil sets the value for AutocaptureWebVitalsAllowedMetrics to be an explicit nil

### UnsetAutocaptureWebVitalsAllowedMetrics
`func (o *PatchedProjectBackwardCompat) UnsetAutocaptureWebVitalsAllowedMetrics()`

UnsetAutocaptureWebVitalsAllowedMetrics ensures that no value is present for AutocaptureWebVitalsAllowedMetrics, not even an explicit nil
### GetAutocaptureExceptionsErrorsToIgnore

`func (o *PatchedProjectBackwardCompat) GetAutocaptureExceptionsErrorsToIgnore() interface{}`

GetAutocaptureExceptionsErrorsToIgnore returns the AutocaptureExceptionsErrorsToIgnore field if non-nil, zero value otherwise.

### GetAutocaptureExceptionsErrorsToIgnoreOk

`func (o *PatchedProjectBackwardCompat) GetAutocaptureExceptionsErrorsToIgnoreOk() (*interface{}, bool)`

GetAutocaptureExceptionsErrorsToIgnoreOk returns a tuple with the AutocaptureExceptionsErrorsToIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureExceptionsErrorsToIgnore

`func (o *PatchedProjectBackwardCompat) SetAutocaptureExceptionsErrorsToIgnore(v interface{})`

SetAutocaptureExceptionsErrorsToIgnore sets AutocaptureExceptionsErrorsToIgnore field to given value.

### HasAutocaptureExceptionsErrorsToIgnore

`func (o *PatchedProjectBackwardCompat) HasAutocaptureExceptionsErrorsToIgnore() bool`

HasAutocaptureExceptionsErrorsToIgnore returns a boolean if a field has been set.

### SetAutocaptureExceptionsErrorsToIgnoreNil

`func (o *PatchedProjectBackwardCompat) SetAutocaptureExceptionsErrorsToIgnoreNil(b bool)`

 SetAutocaptureExceptionsErrorsToIgnoreNil sets the value for AutocaptureExceptionsErrorsToIgnore to be an explicit nil

### UnsetAutocaptureExceptionsErrorsToIgnore
`func (o *PatchedProjectBackwardCompat) UnsetAutocaptureExceptionsErrorsToIgnore()`

UnsetAutocaptureExceptionsErrorsToIgnore ensures that no value is present for AutocaptureExceptionsErrorsToIgnore, not even an explicit nil
### GetCaptureConsoleLogOptIn

`func (o *PatchedProjectBackwardCompat) GetCaptureConsoleLogOptIn() bool`

GetCaptureConsoleLogOptIn returns the CaptureConsoleLogOptIn field if non-nil, zero value otherwise.

### GetCaptureConsoleLogOptInOk

`func (o *PatchedProjectBackwardCompat) GetCaptureConsoleLogOptInOk() (*bool, bool)`

GetCaptureConsoleLogOptInOk returns a tuple with the CaptureConsoleLogOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureConsoleLogOptIn

`func (o *PatchedProjectBackwardCompat) SetCaptureConsoleLogOptIn(v bool)`

SetCaptureConsoleLogOptIn sets CaptureConsoleLogOptIn field to given value.

### HasCaptureConsoleLogOptIn

`func (o *PatchedProjectBackwardCompat) HasCaptureConsoleLogOptIn() bool`

HasCaptureConsoleLogOptIn returns a boolean if a field has been set.

### SetCaptureConsoleLogOptInNil

`func (o *PatchedProjectBackwardCompat) SetCaptureConsoleLogOptInNil(b bool)`

 SetCaptureConsoleLogOptInNil sets the value for CaptureConsoleLogOptIn to be an explicit nil

### UnsetCaptureConsoleLogOptIn
`func (o *PatchedProjectBackwardCompat) UnsetCaptureConsoleLogOptIn()`

UnsetCaptureConsoleLogOptIn ensures that no value is present for CaptureConsoleLogOptIn, not even an explicit nil
### GetCapturePerformanceOptIn

`func (o *PatchedProjectBackwardCompat) GetCapturePerformanceOptIn() bool`

GetCapturePerformanceOptIn returns the CapturePerformanceOptIn field if non-nil, zero value otherwise.

### GetCapturePerformanceOptInOk

`func (o *PatchedProjectBackwardCompat) GetCapturePerformanceOptInOk() (*bool, bool)`

GetCapturePerformanceOptInOk returns a tuple with the CapturePerformanceOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturePerformanceOptIn

`func (o *PatchedProjectBackwardCompat) SetCapturePerformanceOptIn(v bool)`

SetCapturePerformanceOptIn sets CapturePerformanceOptIn field to given value.

### HasCapturePerformanceOptIn

`func (o *PatchedProjectBackwardCompat) HasCapturePerformanceOptIn() bool`

HasCapturePerformanceOptIn returns a boolean if a field has been set.

### SetCapturePerformanceOptInNil

`func (o *PatchedProjectBackwardCompat) SetCapturePerformanceOptInNil(b bool)`

 SetCapturePerformanceOptInNil sets the value for CapturePerformanceOptIn to be an explicit nil

### UnsetCapturePerformanceOptIn
`func (o *PatchedProjectBackwardCompat) UnsetCapturePerformanceOptIn()`

UnsetCapturePerformanceOptIn ensures that no value is present for CapturePerformanceOptIn, not even an explicit nil
### GetSessionRecordingOptIn

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingOptIn() bool`

GetSessionRecordingOptIn returns the SessionRecordingOptIn field if non-nil, zero value otherwise.

### GetSessionRecordingOptInOk

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingOptInOk() (*bool, bool)`

GetSessionRecordingOptInOk returns a tuple with the SessionRecordingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingOptIn

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingOptIn(v bool)`

SetSessionRecordingOptIn sets SessionRecordingOptIn field to given value.

### HasSessionRecordingOptIn

`func (o *PatchedProjectBackwardCompat) HasSessionRecordingOptIn() bool`

HasSessionRecordingOptIn returns a boolean if a field has been set.

### GetSessionRecordingSampleRate

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingSampleRate() float64`

GetSessionRecordingSampleRate returns the SessionRecordingSampleRate field if non-nil, zero value otherwise.

### GetSessionRecordingSampleRateOk

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingSampleRateOk() (*float64, bool)`

GetSessionRecordingSampleRateOk returns a tuple with the SessionRecordingSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingSampleRate

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingSampleRate(v float64)`

SetSessionRecordingSampleRate sets SessionRecordingSampleRate field to given value.

### HasSessionRecordingSampleRate

`func (o *PatchedProjectBackwardCompat) HasSessionRecordingSampleRate() bool`

HasSessionRecordingSampleRate returns a boolean if a field has been set.

### SetSessionRecordingSampleRateNil

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingSampleRateNil(b bool)`

 SetSessionRecordingSampleRateNil sets the value for SessionRecordingSampleRate to be an explicit nil

### UnsetSessionRecordingSampleRate
`func (o *PatchedProjectBackwardCompat) UnsetSessionRecordingSampleRate()`

UnsetSessionRecordingSampleRate ensures that no value is present for SessionRecordingSampleRate, not even an explicit nil
### GetSessionRecordingMinimumDurationMilliseconds

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingMinimumDurationMilliseconds() int32`

GetSessionRecordingMinimumDurationMilliseconds returns the SessionRecordingMinimumDurationMilliseconds field if non-nil, zero value otherwise.

### GetSessionRecordingMinimumDurationMillisecondsOk

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingMinimumDurationMillisecondsOk() (*int32, bool)`

GetSessionRecordingMinimumDurationMillisecondsOk returns a tuple with the SessionRecordingMinimumDurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingMinimumDurationMilliseconds

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingMinimumDurationMilliseconds(v int32)`

SetSessionRecordingMinimumDurationMilliseconds sets SessionRecordingMinimumDurationMilliseconds field to given value.

### HasSessionRecordingMinimumDurationMilliseconds

`func (o *PatchedProjectBackwardCompat) HasSessionRecordingMinimumDurationMilliseconds() bool`

HasSessionRecordingMinimumDurationMilliseconds returns a boolean if a field has been set.

### SetSessionRecordingMinimumDurationMillisecondsNil

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingMinimumDurationMillisecondsNil(b bool)`

 SetSessionRecordingMinimumDurationMillisecondsNil sets the value for SessionRecordingMinimumDurationMilliseconds to be an explicit nil

### UnsetSessionRecordingMinimumDurationMilliseconds
`func (o *PatchedProjectBackwardCompat) UnsetSessionRecordingMinimumDurationMilliseconds()`

UnsetSessionRecordingMinimumDurationMilliseconds ensures that no value is present for SessionRecordingMinimumDurationMilliseconds, not even an explicit nil
### GetSessionRecordingLinkedFlag

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingLinkedFlag() interface{}`

GetSessionRecordingLinkedFlag returns the SessionRecordingLinkedFlag field if non-nil, zero value otherwise.

### GetSessionRecordingLinkedFlagOk

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingLinkedFlagOk() (*interface{}, bool)`

GetSessionRecordingLinkedFlagOk returns a tuple with the SessionRecordingLinkedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingLinkedFlag

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingLinkedFlag(v interface{})`

SetSessionRecordingLinkedFlag sets SessionRecordingLinkedFlag field to given value.

### HasSessionRecordingLinkedFlag

`func (o *PatchedProjectBackwardCompat) HasSessionRecordingLinkedFlag() bool`

HasSessionRecordingLinkedFlag returns a boolean if a field has been set.

### SetSessionRecordingLinkedFlagNil

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingLinkedFlagNil(b bool)`

 SetSessionRecordingLinkedFlagNil sets the value for SessionRecordingLinkedFlag to be an explicit nil

### UnsetSessionRecordingLinkedFlag
`func (o *PatchedProjectBackwardCompat) UnsetSessionRecordingLinkedFlag()`

UnsetSessionRecordingLinkedFlag ensures that no value is present for SessionRecordingLinkedFlag, not even an explicit nil
### GetSessionRecordingNetworkPayloadCaptureConfig

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingNetworkPayloadCaptureConfig() interface{}`

GetSessionRecordingNetworkPayloadCaptureConfig returns the SessionRecordingNetworkPayloadCaptureConfig field if non-nil, zero value otherwise.

### GetSessionRecordingNetworkPayloadCaptureConfigOk

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingNetworkPayloadCaptureConfigOk() (*interface{}, bool)`

GetSessionRecordingNetworkPayloadCaptureConfigOk returns a tuple with the SessionRecordingNetworkPayloadCaptureConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingNetworkPayloadCaptureConfig

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingNetworkPayloadCaptureConfig(v interface{})`

SetSessionRecordingNetworkPayloadCaptureConfig sets SessionRecordingNetworkPayloadCaptureConfig field to given value.

### HasSessionRecordingNetworkPayloadCaptureConfig

`func (o *PatchedProjectBackwardCompat) HasSessionRecordingNetworkPayloadCaptureConfig() bool`

HasSessionRecordingNetworkPayloadCaptureConfig returns a boolean if a field has been set.

### SetSessionRecordingNetworkPayloadCaptureConfigNil

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingNetworkPayloadCaptureConfigNil(b bool)`

 SetSessionRecordingNetworkPayloadCaptureConfigNil sets the value for SessionRecordingNetworkPayloadCaptureConfig to be an explicit nil

### UnsetSessionRecordingNetworkPayloadCaptureConfig
`func (o *PatchedProjectBackwardCompat) UnsetSessionRecordingNetworkPayloadCaptureConfig()`

UnsetSessionRecordingNetworkPayloadCaptureConfig ensures that no value is present for SessionRecordingNetworkPayloadCaptureConfig, not even an explicit nil
### GetSessionRecordingMaskingConfig

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingMaskingConfig() interface{}`

GetSessionRecordingMaskingConfig returns the SessionRecordingMaskingConfig field if non-nil, zero value otherwise.

### GetSessionRecordingMaskingConfigOk

`func (o *PatchedProjectBackwardCompat) GetSessionRecordingMaskingConfigOk() (*interface{}, bool)`

GetSessionRecordingMaskingConfigOk returns a tuple with the SessionRecordingMaskingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingMaskingConfig

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingMaskingConfig(v interface{})`

SetSessionRecordingMaskingConfig sets SessionRecordingMaskingConfig field to given value.

### HasSessionRecordingMaskingConfig

`func (o *PatchedProjectBackwardCompat) HasSessionRecordingMaskingConfig() bool`

HasSessionRecordingMaskingConfig returns a boolean if a field has been set.

### SetSessionRecordingMaskingConfigNil

`func (o *PatchedProjectBackwardCompat) SetSessionRecordingMaskingConfigNil(b bool)`

 SetSessionRecordingMaskingConfigNil sets the value for SessionRecordingMaskingConfig to be an explicit nil

### UnsetSessionRecordingMaskingConfig
`func (o *PatchedProjectBackwardCompat) UnsetSessionRecordingMaskingConfig()`

UnsetSessionRecordingMaskingConfig ensures that no value is present for SessionRecordingMaskingConfig, not even an explicit nil
### GetSessionReplayConfig

`func (o *PatchedProjectBackwardCompat) GetSessionReplayConfig() interface{}`

GetSessionReplayConfig returns the SessionReplayConfig field if non-nil, zero value otherwise.

### GetSessionReplayConfigOk

`func (o *PatchedProjectBackwardCompat) GetSessionReplayConfigOk() (*interface{}, bool)`

GetSessionReplayConfigOk returns a tuple with the SessionReplayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayConfig

`func (o *PatchedProjectBackwardCompat) SetSessionReplayConfig(v interface{})`

SetSessionReplayConfig sets SessionReplayConfig field to given value.

### HasSessionReplayConfig

`func (o *PatchedProjectBackwardCompat) HasSessionReplayConfig() bool`

HasSessionReplayConfig returns a boolean if a field has been set.

### SetSessionReplayConfigNil

`func (o *PatchedProjectBackwardCompat) SetSessionReplayConfigNil(b bool)`

 SetSessionReplayConfigNil sets the value for SessionReplayConfig to be an explicit nil

### UnsetSessionReplayConfig
`func (o *PatchedProjectBackwardCompat) UnsetSessionReplayConfig()`

UnsetSessionReplayConfig ensures that no value is present for SessionReplayConfig, not even an explicit nil
### GetSurveyConfig

`func (o *PatchedProjectBackwardCompat) GetSurveyConfig() interface{}`

GetSurveyConfig returns the SurveyConfig field if non-nil, zero value otherwise.

### GetSurveyConfigOk

`func (o *PatchedProjectBackwardCompat) GetSurveyConfigOk() (*interface{}, bool)`

GetSurveyConfigOk returns a tuple with the SurveyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyConfig

`func (o *PatchedProjectBackwardCompat) SetSurveyConfig(v interface{})`

SetSurveyConfig sets SurveyConfig field to given value.

### HasSurveyConfig

`func (o *PatchedProjectBackwardCompat) HasSurveyConfig() bool`

HasSurveyConfig returns a boolean if a field has been set.

### SetSurveyConfigNil

`func (o *PatchedProjectBackwardCompat) SetSurveyConfigNil(b bool)`

 SetSurveyConfigNil sets the value for SurveyConfig to be an explicit nil

### UnsetSurveyConfig
`func (o *PatchedProjectBackwardCompat) UnsetSurveyConfig()`

UnsetSurveyConfig ensures that no value is present for SurveyConfig, not even an explicit nil
### GetAccessControl

`func (o *PatchedProjectBackwardCompat) GetAccessControl() bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *PatchedProjectBackwardCompat) GetAccessControlOk() (*bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *PatchedProjectBackwardCompat) SetAccessControl(v bool)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *PatchedProjectBackwardCompat) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetWeekStartDay

`func (o *PatchedProjectBackwardCompat) GetWeekStartDay() WeekStartDayEnum`

GetWeekStartDay returns the WeekStartDay field if non-nil, zero value otherwise.

### GetWeekStartDayOk

`func (o *PatchedProjectBackwardCompat) GetWeekStartDayOk() (*WeekStartDayEnum, bool)`

GetWeekStartDayOk returns a tuple with the WeekStartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekStartDay

`func (o *PatchedProjectBackwardCompat) SetWeekStartDay(v WeekStartDayEnum)`

SetWeekStartDay sets WeekStartDay field to given value.

### HasWeekStartDay

`func (o *PatchedProjectBackwardCompat) HasWeekStartDay() bool`

HasWeekStartDay returns a boolean if a field has been set.

### SetWeekStartDayNil

`func (o *PatchedProjectBackwardCompat) SetWeekStartDayNil(b bool)`

 SetWeekStartDayNil sets the value for WeekStartDay to be an explicit nil

### UnsetWeekStartDay
`func (o *PatchedProjectBackwardCompat) UnsetWeekStartDay()`

UnsetWeekStartDay ensures that no value is present for WeekStartDay, not even an explicit nil
### GetPrimaryDashboard

`func (o *PatchedProjectBackwardCompat) GetPrimaryDashboard() int32`

GetPrimaryDashboard returns the PrimaryDashboard field if non-nil, zero value otherwise.

### GetPrimaryDashboardOk

`func (o *PatchedProjectBackwardCompat) GetPrimaryDashboardOk() (*int32, bool)`

GetPrimaryDashboardOk returns a tuple with the PrimaryDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDashboard

`func (o *PatchedProjectBackwardCompat) SetPrimaryDashboard(v int32)`

SetPrimaryDashboard sets PrimaryDashboard field to given value.

### HasPrimaryDashboard

`func (o *PatchedProjectBackwardCompat) HasPrimaryDashboard() bool`

HasPrimaryDashboard returns a boolean if a field has been set.

### SetPrimaryDashboardNil

`func (o *PatchedProjectBackwardCompat) SetPrimaryDashboardNil(b bool)`

 SetPrimaryDashboardNil sets the value for PrimaryDashboard to be an explicit nil

### UnsetPrimaryDashboard
`func (o *PatchedProjectBackwardCompat) UnsetPrimaryDashboard()`

UnsetPrimaryDashboard ensures that no value is present for PrimaryDashboard, not even an explicit nil
### GetLiveEventsColumns

`func (o *PatchedProjectBackwardCompat) GetLiveEventsColumns() []string`

GetLiveEventsColumns returns the LiveEventsColumns field if non-nil, zero value otherwise.

### GetLiveEventsColumnsOk

`func (o *PatchedProjectBackwardCompat) GetLiveEventsColumnsOk() (*[]string, bool)`

GetLiveEventsColumnsOk returns a tuple with the LiveEventsColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEventsColumns

`func (o *PatchedProjectBackwardCompat) SetLiveEventsColumns(v []string)`

SetLiveEventsColumns sets LiveEventsColumns field to given value.

### HasLiveEventsColumns

`func (o *PatchedProjectBackwardCompat) HasLiveEventsColumns() bool`

HasLiveEventsColumns returns a boolean if a field has been set.

### SetLiveEventsColumnsNil

`func (o *PatchedProjectBackwardCompat) SetLiveEventsColumnsNil(b bool)`

 SetLiveEventsColumnsNil sets the value for LiveEventsColumns to be an explicit nil

### UnsetLiveEventsColumns
`func (o *PatchedProjectBackwardCompat) UnsetLiveEventsColumns()`

UnsetLiveEventsColumns ensures that no value is present for LiveEventsColumns, not even an explicit nil
### GetRecordingDomains

`func (o *PatchedProjectBackwardCompat) GetRecordingDomains() []*string`

GetRecordingDomains returns the RecordingDomains field if non-nil, zero value otherwise.

### GetRecordingDomainsOk

`func (o *PatchedProjectBackwardCompat) GetRecordingDomainsOk() (*[]*string, bool)`

GetRecordingDomainsOk returns a tuple with the RecordingDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingDomains

`func (o *PatchedProjectBackwardCompat) SetRecordingDomains(v []*string)`

SetRecordingDomains sets RecordingDomains field to given value.

### HasRecordingDomains

`func (o *PatchedProjectBackwardCompat) HasRecordingDomains() bool`

HasRecordingDomains returns a boolean if a field has been set.

### SetRecordingDomainsNil

`func (o *PatchedProjectBackwardCompat) SetRecordingDomainsNil(b bool)`

 SetRecordingDomainsNil sets the value for RecordingDomains to be an explicit nil

### UnsetRecordingDomains
`func (o *PatchedProjectBackwardCompat) UnsetRecordingDomains()`

UnsetRecordingDomains ensures that no value is present for RecordingDomains, not even an explicit nil
### GetPersonOnEventsQueryingEnabled

`func (o *PatchedProjectBackwardCompat) GetPersonOnEventsQueryingEnabled() string`

GetPersonOnEventsQueryingEnabled returns the PersonOnEventsQueryingEnabled field if non-nil, zero value otherwise.

### GetPersonOnEventsQueryingEnabledOk

`func (o *PatchedProjectBackwardCompat) GetPersonOnEventsQueryingEnabledOk() (*string, bool)`

GetPersonOnEventsQueryingEnabledOk returns a tuple with the PersonOnEventsQueryingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonOnEventsQueryingEnabled

`func (o *PatchedProjectBackwardCompat) SetPersonOnEventsQueryingEnabled(v string)`

SetPersonOnEventsQueryingEnabled sets PersonOnEventsQueryingEnabled field to given value.

### HasPersonOnEventsQueryingEnabled

`func (o *PatchedProjectBackwardCompat) HasPersonOnEventsQueryingEnabled() bool`

HasPersonOnEventsQueryingEnabled returns a boolean if a field has been set.

### GetInjectWebApps

`func (o *PatchedProjectBackwardCompat) GetInjectWebApps() bool`

GetInjectWebApps returns the InjectWebApps field if non-nil, zero value otherwise.

### GetInjectWebAppsOk

`func (o *PatchedProjectBackwardCompat) GetInjectWebAppsOk() (*bool, bool)`

GetInjectWebAppsOk returns a tuple with the InjectWebApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectWebApps

`func (o *PatchedProjectBackwardCompat) SetInjectWebApps(v bool)`

SetInjectWebApps sets InjectWebApps field to given value.

### HasInjectWebApps

`func (o *PatchedProjectBackwardCompat) HasInjectWebApps() bool`

HasInjectWebApps returns a boolean if a field has been set.

### SetInjectWebAppsNil

`func (o *PatchedProjectBackwardCompat) SetInjectWebAppsNil(b bool)`

 SetInjectWebAppsNil sets the value for InjectWebApps to be an explicit nil

### UnsetInjectWebApps
`func (o *PatchedProjectBackwardCompat) UnsetInjectWebApps()`

UnsetInjectWebApps ensures that no value is present for InjectWebApps, not even an explicit nil
### GetExtraSettings

`func (o *PatchedProjectBackwardCompat) GetExtraSettings() interface{}`

GetExtraSettings returns the ExtraSettings field if non-nil, zero value otherwise.

### GetExtraSettingsOk

`func (o *PatchedProjectBackwardCompat) GetExtraSettingsOk() (*interface{}, bool)`

GetExtraSettingsOk returns a tuple with the ExtraSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSettings

`func (o *PatchedProjectBackwardCompat) SetExtraSettings(v interface{})`

SetExtraSettings sets ExtraSettings field to given value.

### HasExtraSettings

`func (o *PatchedProjectBackwardCompat) HasExtraSettings() bool`

HasExtraSettings returns a boolean if a field has been set.

### SetExtraSettingsNil

`func (o *PatchedProjectBackwardCompat) SetExtraSettingsNil(b bool)`

 SetExtraSettingsNil sets the value for ExtraSettings to be an explicit nil

### UnsetExtraSettings
`func (o *PatchedProjectBackwardCompat) UnsetExtraSettings()`

UnsetExtraSettings ensures that no value is present for ExtraSettings, not even an explicit nil
### GetModifiers

`func (o *PatchedProjectBackwardCompat) GetModifiers() interface{}`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *PatchedProjectBackwardCompat) GetModifiersOk() (*interface{}, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *PatchedProjectBackwardCompat) SetModifiers(v interface{})`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *PatchedProjectBackwardCompat) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### SetModifiersNil

`func (o *PatchedProjectBackwardCompat) SetModifiersNil(b bool)`

 SetModifiersNil sets the value for Modifiers to be an explicit nil

### UnsetModifiers
`func (o *PatchedProjectBackwardCompat) UnsetModifiers()`

UnsetModifiers ensures that no value is present for Modifiers, not even an explicit nil
### GetDefaultModifiers

`func (o *PatchedProjectBackwardCompat) GetDefaultModifiers() string`

GetDefaultModifiers returns the DefaultModifiers field if non-nil, zero value otherwise.

### GetDefaultModifiersOk

`func (o *PatchedProjectBackwardCompat) GetDefaultModifiersOk() (*string, bool)`

GetDefaultModifiersOk returns a tuple with the DefaultModifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModifiers

`func (o *PatchedProjectBackwardCompat) SetDefaultModifiers(v string)`

SetDefaultModifiers sets DefaultModifiers field to given value.

### HasDefaultModifiers

`func (o *PatchedProjectBackwardCompat) HasDefaultModifiers() bool`

HasDefaultModifiers returns a boolean if a field has been set.

### GetHasCompletedOnboardingFor

`func (o *PatchedProjectBackwardCompat) GetHasCompletedOnboardingFor() interface{}`

GetHasCompletedOnboardingFor returns the HasCompletedOnboardingFor field if non-nil, zero value otherwise.

### GetHasCompletedOnboardingForOk

`func (o *PatchedProjectBackwardCompat) GetHasCompletedOnboardingForOk() (*interface{}, bool)`

GetHasCompletedOnboardingForOk returns a tuple with the HasCompletedOnboardingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompletedOnboardingFor

`func (o *PatchedProjectBackwardCompat) SetHasCompletedOnboardingFor(v interface{})`

SetHasCompletedOnboardingFor sets HasCompletedOnboardingFor field to given value.

### HasHasCompletedOnboardingFor

`func (o *PatchedProjectBackwardCompat) HasHasCompletedOnboardingFor() bool`

HasHasCompletedOnboardingFor returns a boolean if a field has been set.

### SetHasCompletedOnboardingForNil

`func (o *PatchedProjectBackwardCompat) SetHasCompletedOnboardingForNil(b bool)`

 SetHasCompletedOnboardingForNil sets the value for HasCompletedOnboardingFor to be an explicit nil

### UnsetHasCompletedOnboardingFor
`func (o *PatchedProjectBackwardCompat) UnsetHasCompletedOnboardingFor()`

UnsetHasCompletedOnboardingFor ensures that no value is present for HasCompletedOnboardingFor, not even an explicit nil
### GetSurveysOptIn

`func (o *PatchedProjectBackwardCompat) GetSurveysOptIn() bool`

GetSurveysOptIn returns the SurveysOptIn field if non-nil, zero value otherwise.

### GetSurveysOptInOk

`func (o *PatchedProjectBackwardCompat) GetSurveysOptInOk() (*bool, bool)`

GetSurveysOptInOk returns a tuple with the SurveysOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveysOptIn

`func (o *PatchedProjectBackwardCompat) SetSurveysOptIn(v bool)`

SetSurveysOptIn sets SurveysOptIn field to given value.

### HasSurveysOptIn

`func (o *PatchedProjectBackwardCompat) HasSurveysOptIn() bool`

HasSurveysOptIn returns a boolean if a field has been set.

### SetSurveysOptInNil

`func (o *PatchedProjectBackwardCompat) SetSurveysOptInNil(b bool)`

 SetSurveysOptInNil sets the value for SurveysOptIn to be an explicit nil

### UnsetSurveysOptIn
`func (o *PatchedProjectBackwardCompat) UnsetSurveysOptIn()`

UnsetSurveysOptIn ensures that no value is present for SurveysOptIn, not even an explicit nil
### GetHeatmapsOptIn

`func (o *PatchedProjectBackwardCompat) GetHeatmapsOptIn() bool`

GetHeatmapsOptIn returns the HeatmapsOptIn field if non-nil, zero value otherwise.

### GetHeatmapsOptInOk

`func (o *PatchedProjectBackwardCompat) GetHeatmapsOptInOk() (*bool, bool)`

GetHeatmapsOptInOk returns a tuple with the HeatmapsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapsOptIn

`func (o *PatchedProjectBackwardCompat) SetHeatmapsOptIn(v bool)`

SetHeatmapsOptIn sets HeatmapsOptIn field to given value.

### HasHeatmapsOptIn

`func (o *PatchedProjectBackwardCompat) HasHeatmapsOptIn() bool`

HasHeatmapsOptIn returns a boolean if a field has been set.

### SetHeatmapsOptInNil

`func (o *PatchedProjectBackwardCompat) SetHeatmapsOptInNil(b bool)`

 SetHeatmapsOptInNil sets the value for HeatmapsOptIn to be an explicit nil

### UnsetHeatmapsOptIn
`func (o *PatchedProjectBackwardCompat) UnsetHeatmapsOptIn()`

UnsetHeatmapsOptIn ensures that no value is present for HeatmapsOptIn, not even an explicit nil
### GetProductIntents

`func (o *PatchedProjectBackwardCompat) GetProductIntents() string`

GetProductIntents returns the ProductIntents field if non-nil, zero value otherwise.

### GetProductIntentsOk

`func (o *PatchedProjectBackwardCompat) GetProductIntentsOk() (*string, bool)`

GetProductIntentsOk returns a tuple with the ProductIntents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIntents

`func (o *PatchedProjectBackwardCompat) SetProductIntents(v string)`

SetProductIntents sets ProductIntents field to given value.

### HasProductIntents

`func (o *PatchedProjectBackwardCompat) HasProductIntents() bool`

HasProductIntents returns a boolean if a field has been set.

### GetFlagsPersistenceDefault

`func (o *PatchedProjectBackwardCompat) GetFlagsPersistenceDefault() bool`

GetFlagsPersistenceDefault returns the FlagsPersistenceDefault field if non-nil, zero value otherwise.

### GetFlagsPersistenceDefaultOk

`func (o *PatchedProjectBackwardCompat) GetFlagsPersistenceDefaultOk() (*bool, bool)`

GetFlagsPersistenceDefaultOk returns a tuple with the FlagsPersistenceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsPersistenceDefault

`func (o *PatchedProjectBackwardCompat) SetFlagsPersistenceDefault(v bool)`

SetFlagsPersistenceDefault sets FlagsPersistenceDefault field to given value.

### HasFlagsPersistenceDefault

`func (o *PatchedProjectBackwardCompat) HasFlagsPersistenceDefault() bool`

HasFlagsPersistenceDefault returns a boolean if a field has been set.

### SetFlagsPersistenceDefaultNil

`func (o *PatchedProjectBackwardCompat) SetFlagsPersistenceDefaultNil(b bool)`

 SetFlagsPersistenceDefaultNil sets the value for FlagsPersistenceDefault to be an explicit nil

### UnsetFlagsPersistenceDefault
`func (o *PatchedProjectBackwardCompat) UnsetFlagsPersistenceDefault()`

UnsetFlagsPersistenceDefault ensures that no value is present for FlagsPersistenceDefault, not even an explicit nil
### GetSecretApiToken

`func (o *PatchedProjectBackwardCompat) GetSecretApiToken() string`

GetSecretApiToken returns the SecretApiToken field if non-nil, zero value otherwise.

### GetSecretApiTokenOk

`func (o *PatchedProjectBackwardCompat) GetSecretApiTokenOk() (*string, bool)`

GetSecretApiTokenOk returns a tuple with the SecretApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiToken

`func (o *PatchedProjectBackwardCompat) SetSecretApiToken(v string)`

SetSecretApiToken sets SecretApiToken field to given value.

### HasSecretApiToken

`func (o *PatchedProjectBackwardCompat) HasSecretApiToken() bool`

HasSecretApiToken returns a boolean if a field has been set.

### SetSecretApiTokenNil

`func (o *PatchedProjectBackwardCompat) SetSecretApiTokenNil(b bool)`

 SetSecretApiTokenNil sets the value for SecretApiToken to be an explicit nil

### UnsetSecretApiToken
`func (o *PatchedProjectBackwardCompat) UnsetSecretApiToken()`

UnsetSecretApiToken ensures that no value is present for SecretApiToken, not even an explicit nil
### GetSecretApiTokenBackup

`func (o *PatchedProjectBackwardCompat) GetSecretApiTokenBackup() string`

GetSecretApiTokenBackup returns the SecretApiTokenBackup field if non-nil, zero value otherwise.

### GetSecretApiTokenBackupOk

`func (o *PatchedProjectBackwardCompat) GetSecretApiTokenBackupOk() (*string, bool)`

GetSecretApiTokenBackupOk returns a tuple with the SecretApiTokenBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiTokenBackup

`func (o *PatchedProjectBackwardCompat) SetSecretApiTokenBackup(v string)`

SetSecretApiTokenBackup sets SecretApiTokenBackup field to given value.

### HasSecretApiTokenBackup

`func (o *PatchedProjectBackwardCompat) HasSecretApiTokenBackup() bool`

HasSecretApiTokenBackup returns a boolean if a field has been set.

### SetSecretApiTokenBackupNil

`func (o *PatchedProjectBackwardCompat) SetSecretApiTokenBackupNil(b bool)`

 SetSecretApiTokenBackupNil sets the value for SecretApiTokenBackup to be an explicit nil

### UnsetSecretApiTokenBackup
`func (o *PatchedProjectBackwardCompat) UnsetSecretApiTokenBackup()`

UnsetSecretApiTokenBackup ensures that no value is present for SecretApiTokenBackup, not even an explicit nil
### GetReceiveOrgLevelActivityLogs

`func (o *PatchedProjectBackwardCompat) GetReceiveOrgLevelActivityLogs() bool`

GetReceiveOrgLevelActivityLogs returns the ReceiveOrgLevelActivityLogs field if non-nil, zero value otherwise.

### GetReceiveOrgLevelActivityLogsOk

`func (o *PatchedProjectBackwardCompat) GetReceiveOrgLevelActivityLogsOk() (*bool, bool)`

GetReceiveOrgLevelActivityLogsOk returns a tuple with the ReceiveOrgLevelActivityLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveOrgLevelActivityLogs

`func (o *PatchedProjectBackwardCompat) SetReceiveOrgLevelActivityLogs(v bool)`

SetReceiveOrgLevelActivityLogs sets ReceiveOrgLevelActivityLogs field to given value.

### HasReceiveOrgLevelActivityLogs

`func (o *PatchedProjectBackwardCompat) HasReceiveOrgLevelActivityLogs() bool`

HasReceiveOrgLevelActivityLogs returns a boolean if a field has been set.

### SetReceiveOrgLevelActivityLogsNil

`func (o *PatchedProjectBackwardCompat) SetReceiveOrgLevelActivityLogsNil(b bool)`

 SetReceiveOrgLevelActivityLogsNil sets the value for ReceiveOrgLevelActivityLogs to be an explicit nil

### UnsetReceiveOrgLevelActivityLogs
`func (o *PatchedProjectBackwardCompat) UnsetReceiveOrgLevelActivityLogs()`

UnsetReceiveOrgLevelActivityLogs ensures that no value is present for ReceiveOrgLevelActivityLogs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


