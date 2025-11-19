# ProjectBackwardCompat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Organization** | **string** |  | [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ProductDescription** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**EffectiveMembershipLevel** | [**NullableEffectiveMembershipLevelEnum**](EffectiveMembershipLevelEnum.md) |  | [readonly] 
**HasGroupTypes** | **bool** |  | [readonly] 
**GroupTypes** | **[]map[string]interface{}** |  | [readonly] 
**LiveEventsToken** | **NullableString** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Uuid** | **string** |  | [readonly] 
**ApiToken** | **string** |  | [readonly] 
**AppUrls** | Pointer to **[]string** |  | [optional] 
**SlackIncomingWebhook** | Pointer to **NullableString** |  | [optional] 
**AnonymizeIps** | Pointer to **bool** |  | [optional] 
**CompletedSnippetOnboarding** | Pointer to **bool** |  | [optional] 
**IngestedEvent** | **bool** |  | [readonly] 
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
**PersonOnEventsQueryingEnabled** | **string** |  | [readonly] 
**InjectWebApps** | Pointer to **NullableBool** |  | [optional] 
**ExtraSettings** | Pointer to **interface{}** |  | [optional] 
**Modifiers** | Pointer to **interface{}** |  | [optional] 
**DefaultModifiers** | **string** |  | [readonly] 
**HasCompletedOnboardingFor** | Pointer to **interface{}** |  | [optional] 
**SurveysOptIn** | Pointer to **NullableBool** |  | [optional] 
**HeatmapsOptIn** | Pointer to **NullableBool** |  | [optional] 
**ProductIntents** | **string** |  | [readonly] 
**FlagsPersistenceDefault** | Pointer to **NullableBool** |  | [optional] 
**SecretApiToken** | **NullableString** |  | [readonly] 
**SecretApiTokenBackup** | **NullableString** |  | [readonly] 
**ReceiveOrgLevelActivityLogs** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewProjectBackwardCompat

`func NewProjectBackwardCompat(id int32, organization string, createdAt time.Time, effectiveMembershipLevel NullableEffectiveMembershipLevelEnum, hasGroupTypes bool, groupTypes []map[string]interface{}, liveEventsToken NullableString, updatedAt time.Time, uuid string, apiToken string, ingestedEvent bool, personOnEventsQueryingEnabled string, defaultModifiers string, productIntents string, secretApiToken NullableString, secretApiTokenBackup NullableString, ) *ProjectBackwardCompat`

NewProjectBackwardCompat instantiates a new ProjectBackwardCompat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectBackwardCompatWithDefaults

`func NewProjectBackwardCompatWithDefaults() *ProjectBackwardCompat`

NewProjectBackwardCompatWithDefaults instantiates a new ProjectBackwardCompat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectBackwardCompat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectBackwardCompat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectBackwardCompat) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganization

`func (o *ProjectBackwardCompat) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectBackwardCompat) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectBackwardCompat) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetName

`func (o *ProjectBackwardCompat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectBackwardCompat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectBackwardCompat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectBackwardCompat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductDescription

`func (o *ProjectBackwardCompat) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *ProjectBackwardCompat) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *ProjectBackwardCompat) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *ProjectBackwardCompat) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### SetProductDescriptionNil

`func (o *ProjectBackwardCompat) SetProductDescriptionNil(b bool)`

 SetProductDescriptionNil sets the value for ProductDescription to be an explicit nil

### UnsetProductDescription
`func (o *ProjectBackwardCompat) UnsetProductDescription()`

UnsetProductDescription ensures that no value is present for ProductDescription, not even an explicit nil
### GetCreatedAt

`func (o *ProjectBackwardCompat) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectBackwardCompat) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectBackwardCompat) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEffectiveMembershipLevel

`func (o *ProjectBackwardCompat) GetEffectiveMembershipLevel() EffectiveMembershipLevelEnum`

GetEffectiveMembershipLevel returns the EffectiveMembershipLevel field if non-nil, zero value otherwise.

### GetEffectiveMembershipLevelOk

`func (o *ProjectBackwardCompat) GetEffectiveMembershipLevelOk() (*EffectiveMembershipLevelEnum, bool)`

GetEffectiveMembershipLevelOk returns a tuple with the EffectiveMembershipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveMembershipLevel

`func (o *ProjectBackwardCompat) SetEffectiveMembershipLevel(v EffectiveMembershipLevelEnum)`

SetEffectiveMembershipLevel sets EffectiveMembershipLevel field to given value.


### SetEffectiveMembershipLevelNil

`func (o *ProjectBackwardCompat) SetEffectiveMembershipLevelNil(b bool)`

 SetEffectiveMembershipLevelNil sets the value for EffectiveMembershipLevel to be an explicit nil

### UnsetEffectiveMembershipLevel
`func (o *ProjectBackwardCompat) UnsetEffectiveMembershipLevel()`

UnsetEffectiveMembershipLevel ensures that no value is present for EffectiveMembershipLevel, not even an explicit nil
### GetHasGroupTypes

`func (o *ProjectBackwardCompat) GetHasGroupTypes() bool`

GetHasGroupTypes returns the HasGroupTypes field if non-nil, zero value otherwise.

### GetHasGroupTypesOk

`func (o *ProjectBackwardCompat) GetHasGroupTypesOk() (*bool, bool)`

GetHasGroupTypesOk returns a tuple with the HasGroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGroupTypes

`func (o *ProjectBackwardCompat) SetHasGroupTypes(v bool)`

SetHasGroupTypes sets HasGroupTypes field to given value.


### GetGroupTypes

`func (o *ProjectBackwardCompat) GetGroupTypes() []map[string]interface{}`

GetGroupTypes returns the GroupTypes field if non-nil, zero value otherwise.

### GetGroupTypesOk

`func (o *ProjectBackwardCompat) GetGroupTypesOk() (*[]map[string]interface{}, bool)`

GetGroupTypesOk returns a tuple with the GroupTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypes

`func (o *ProjectBackwardCompat) SetGroupTypes(v []map[string]interface{})`

SetGroupTypes sets GroupTypes field to given value.


### GetLiveEventsToken

`func (o *ProjectBackwardCompat) GetLiveEventsToken() string`

GetLiveEventsToken returns the LiveEventsToken field if non-nil, zero value otherwise.

### GetLiveEventsTokenOk

`func (o *ProjectBackwardCompat) GetLiveEventsTokenOk() (*string, bool)`

GetLiveEventsTokenOk returns a tuple with the LiveEventsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEventsToken

`func (o *ProjectBackwardCompat) SetLiveEventsToken(v string)`

SetLiveEventsToken sets LiveEventsToken field to given value.


### SetLiveEventsTokenNil

`func (o *ProjectBackwardCompat) SetLiveEventsTokenNil(b bool)`

 SetLiveEventsTokenNil sets the value for LiveEventsToken to be an explicit nil

### UnsetLiveEventsToken
`func (o *ProjectBackwardCompat) UnsetLiveEventsToken()`

UnsetLiveEventsToken ensures that no value is present for LiveEventsToken, not even an explicit nil
### GetUpdatedAt

`func (o *ProjectBackwardCompat) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectBackwardCompat) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectBackwardCompat) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUuid

`func (o *ProjectBackwardCompat) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectBackwardCompat) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectBackwardCompat) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetApiToken

`func (o *ProjectBackwardCompat) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ProjectBackwardCompat) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ProjectBackwardCompat) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetAppUrls

`func (o *ProjectBackwardCompat) GetAppUrls() []*string`

GetAppUrls returns the AppUrls field if non-nil, zero value otherwise.

### GetAppUrlsOk

`func (o *ProjectBackwardCompat) GetAppUrlsOk() (*[]*string, bool)`

GetAppUrlsOk returns a tuple with the AppUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrls

`func (o *ProjectBackwardCompat) SetAppUrls(v []*string)`

SetAppUrls sets AppUrls field to given value.

### HasAppUrls

`func (o *ProjectBackwardCompat) HasAppUrls() bool`

HasAppUrls returns a boolean if a field has been set.

### GetSlackIncomingWebhook

`func (o *ProjectBackwardCompat) GetSlackIncomingWebhook() string`

GetSlackIncomingWebhook returns the SlackIncomingWebhook field if non-nil, zero value otherwise.

### GetSlackIncomingWebhookOk

`func (o *ProjectBackwardCompat) GetSlackIncomingWebhookOk() (*string, bool)`

GetSlackIncomingWebhookOk returns a tuple with the SlackIncomingWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackIncomingWebhook

`func (o *ProjectBackwardCompat) SetSlackIncomingWebhook(v string)`

SetSlackIncomingWebhook sets SlackIncomingWebhook field to given value.

### HasSlackIncomingWebhook

`func (o *ProjectBackwardCompat) HasSlackIncomingWebhook() bool`

HasSlackIncomingWebhook returns a boolean if a field has been set.

### SetSlackIncomingWebhookNil

`func (o *ProjectBackwardCompat) SetSlackIncomingWebhookNil(b bool)`

 SetSlackIncomingWebhookNil sets the value for SlackIncomingWebhook to be an explicit nil

### UnsetSlackIncomingWebhook
`func (o *ProjectBackwardCompat) UnsetSlackIncomingWebhook()`

UnsetSlackIncomingWebhook ensures that no value is present for SlackIncomingWebhook, not even an explicit nil
### GetAnonymizeIps

`func (o *ProjectBackwardCompat) GetAnonymizeIps() bool`

GetAnonymizeIps returns the AnonymizeIps field if non-nil, zero value otherwise.

### GetAnonymizeIpsOk

`func (o *ProjectBackwardCompat) GetAnonymizeIpsOk() (*bool, bool)`

GetAnonymizeIpsOk returns a tuple with the AnonymizeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeIps

`func (o *ProjectBackwardCompat) SetAnonymizeIps(v bool)`

SetAnonymizeIps sets AnonymizeIps field to given value.

### HasAnonymizeIps

`func (o *ProjectBackwardCompat) HasAnonymizeIps() bool`

HasAnonymizeIps returns a boolean if a field has been set.

### GetCompletedSnippetOnboarding

`func (o *ProjectBackwardCompat) GetCompletedSnippetOnboarding() bool`

GetCompletedSnippetOnboarding returns the CompletedSnippetOnboarding field if non-nil, zero value otherwise.

### GetCompletedSnippetOnboardingOk

`func (o *ProjectBackwardCompat) GetCompletedSnippetOnboardingOk() (*bool, bool)`

GetCompletedSnippetOnboardingOk returns a tuple with the CompletedSnippetOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedSnippetOnboarding

`func (o *ProjectBackwardCompat) SetCompletedSnippetOnboarding(v bool)`

SetCompletedSnippetOnboarding sets CompletedSnippetOnboarding field to given value.

### HasCompletedSnippetOnboarding

`func (o *ProjectBackwardCompat) HasCompletedSnippetOnboarding() bool`

HasCompletedSnippetOnboarding returns a boolean if a field has been set.

### GetIngestedEvent

`func (o *ProjectBackwardCompat) GetIngestedEvent() bool`

GetIngestedEvent returns the IngestedEvent field if non-nil, zero value otherwise.

### GetIngestedEventOk

`func (o *ProjectBackwardCompat) GetIngestedEventOk() (*bool, bool)`

GetIngestedEventOk returns a tuple with the IngestedEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedEvent

`func (o *ProjectBackwardCompat) SetIngestedEvent(v bool)`

SetIngestedEvent sets IngestedEvent field to given value.


### GetTestAccountFilters

`func (o *ProjectBackwardCompat) GetTestAccountFilters() interface{}`

GetTestAccountFilters returns the TestAccountFilters field if non-nil, zero value otherwise.

### GetTestAccountFiltersOk

`func (o *ProjectBackwardCompat) GetTestAccountFiltersOk() (*interface{}, bool)`

GetTestAccountFiltersOk returns a tuple with the TestAccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccountFilters

`func (o *ProjectBackwardCompat) SetTestAccountFilters(v interface{})`

SetTestAccountFilters sets TestAccountFilters field to given value.

### HasTestAccountFilters

`func (o *ProjectBackwardCompat) HasTestAccountFilters() bool`

HasTestAccountFilters returns a boolean if a field has been set.

### SetTestAccountFiltersNil

`func (o *ProjectBackwardCompat) SetTestAccountFiltersNil(b bool)`

 SetTestAccountFiltersNil sets the value for TestAccountFilters to be an explicit nil

### UnsetTestAccountFilters
`func (o *ProjectBackwardCompat) UnsetTestAccountFilters()`

UnsetTestAccountFilters ensures that no value is present for TestAccountFilters, not even an explicit nil
### GetTestAccountFiltersDefaultChecked

`func (o *ProjectBackwardCompat) GetTestAccountFiltersDefaultChecked() bool`

GetTestAccountFiltersDefaultChecked returns the TestAccountFiltersDefaultChecked field if non-nil, zero value otherwise.

### GetTestAccountFiltersDefaultCheckedOk

`func (o *ProjectBackwardCompat) GetTestAccountFiltersDefaultCheckedOk() (*bool, bool)`

GetTestAccountFiltersDefaultCheckedOk returns a tuple with the TestAccountFiltersDefaultChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccountFiltersDefaultChecked

`func (o *ProjectBackwardCompat) SetTestAccountFiltersDefaultChecked(v bool)`

SetTestAccountFiltersDefaultChecked sets TestAccountFiltersDefaultChecked field to given value.

### HasTestAccountFiltersDefaultChecked

`func (o *ProjectBackwardCompat) HasTestAccountFiltersDefaultChecked() bool`

HasTestAccountFiltersDefaultChecked returns a boolean if a field has been set.

### SetTestAccountFiltersDefaultCheckedNil

`func (o *ProjectBackwardCompat) SetTestAccountFiltersDefaultCheckedNil(b bool)`

 SetTestAccountFiltersDefaultCheckedNil sets the value for TestAccountFiltersDefaultChecked to be an explicit nil

### UnsetTestAccountFiltersDefaultChecked
`func (o *ProjectBackwardCompat) UnsetTestAccountFiltersDefaultChecked()`

UnsetTestAccountFiltersDefaultChecked ensures that no value is present for TestAccountFiltersDefaultChecked, not even an explicit nil
### GetPathCleaningFilters

`func (o *ProjectBackwardCompat) GetPathCleaningFilters() interface{}`

GetPathCleaningFilters returns the PathCleaningFilters field if non-nil, zero value otherwise.

### GetPathCleaningFiltersOk

`func (o *ProjectBackwardCompat) GetPathCleaningFiltersOk() (*interface{}, bool)`

GetPathCleaningFiltersOk returns a tuple with the PathCleaningFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathCleaningFilters

`func (o *ProjectBackwardCompat) SetPathCleaningFilters(v interface{})`

SetPathCleaningFilters sets PathCleaningFilters field to given value.

### HasPathCleaningFilters

`func (o *ProjectBackwardCompat) HasPathCleaningFilters() bool`

HasPathCleaningFilters returns a boolean if a field has been set.

### SetPathCleaningFiltersNil

`func (o *ProjectBackwardCompat) SetPathCleaningFiltersNil(b bool)`

 SetPathCleaningFiltersNil sets the value for PathCleaningFilters to be an explicit nil

### UnsetPathCleaningFilters
`func (o *ProjectBackwardCompat) UnsetPathCleaningFilters()`

UnsetPathCleaningFilters ensures that no value is present for PathCleaningFilters, not even an explicit nil
### GetIsDemo

`func (o *ProjectBackwardCompat) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *ProjectBackwardCompat) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *ProjectBackwardCompat) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *ProjectBackwardCompat) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetTimezone

`func (o *ProjectBackwardCompat) GetTimezone() TimezoneEnum`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ProjectBackwardCompat) GetTimezoneOk() (*TimezoneEnum, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ProjectBackwardCompat) SetTimezone(v TimezoneEnum)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ProjectBackwardCompat) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDataAttributes

`func (o *ProjectBackwardCompat) GetDataAttributes() interface{}`

GetDataAttributes returns the DataAttributes field if non-nil, zero value otherwise.

### GetDataAttributesOk

`func (o *ProjectBackwardCompat) GetDataAttributesOk() (*interface{}, bool)`

GetDataAttributesOk returns a tuple with the DataAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAttributes

`func (o *ProjectBackwardCompat) SetDataAttributes(v interface{})`

SetDataAttributes sets DataAttributes field to given value.

### HasDataAttributes

`func (o *ProjectBackwardCompat) HasDataAttributes() bool`

HasDataAttributes returns a boolean if a field has been set.

### SetDataAttributesNil

`func (o *ProjectBackwardCompat) SetDataAttributesNil(b bool)`

 SetDataAttributesNil sets the value for DataAttributes to be an explicit nil

### UnsetDataAttributes
`func (o *ProjectBackwardCompat) UnsetDataAttributes()`

UnsetDataAttributes ensures that no value is present for DataAttributes, not even an explicit nil
### GetPersonDisplayNameProperties

`func (o *ProjectBackwardCompat) GetPersonDisplayNameProperties() []string`

GetPersonDisplayNameProperties returns the PersonDisplayNameProperties field if non-nil, zero value otherwise.

### GetPersonDisplayNamePropertiesOk

`func (o *ProjectBackwardCompat) GetPersonDisplayNamePropertiesOk() (*[]string, bool)`

GetPersonDisplayNamePropertiesOk returns a tuple with the PersonDisplayNameProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonDisplayNameProperties

`func (o *ProjectBackwardCompat) SetPersonDisplayNameProperties(v []string)`

SetPersonDisplayNameProperties sets PersonDisplayNameProperties field to given value.

### HasPersonDisplayNameProperties

`func (o *ProjectBackwardCompat) HasPersonDisplayNameProperties() bool`

HasPersonDisplayNameProperties returns a boolean if a field has been set.

### SetPersonDisplayNamePropertiesNil

`func (o *ProjectBackwardCompat) SetPersonDisplayNamePropertiesNil(b bool)`

 SetPersonDisplayNamePropertiesNil sets the value for PersonDisplayNameProperties to be an explicit nil

### UnsetPersonDisplayNameProperties
`func (o *ProjectBackwardCompat) UnsetPersonDisplayNameProperties()`

UnsetPersonDisplayNameProperties ensures that no value is present for PersonDisplayNameProperties, not even an explicit nil
### GetCorrelationConfig

`func (o *ProjectBackwardCompat) GetCorrelationConfig() interface{}`

GetCorrelationConfig returns the CorrelationConfig field if non-nil, zero value otherwise.

### GetCorrelationConfigOk

`func (o *ProjectBackwardCompat) GetCorrelationConfigOk() (*interface{}, bool)`

GetCorrelationConfigOk returns a tuple with the CorrelationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationConfig

`func (o *ProjectBackwardCompat) SetCorrelationConfig(v interface{})`

SetCorrelationConfig sets CorrelationConfig field to given value.

### HasCorrelationConfig

`func (o *ProjectBackwardCompat) HasCorrelationConfig() bool`

HasCorrelationConfig returns a boolean if a field has been set.

### SetCorrelationConfigNil

`func (o *ProjectBackwardCompat) SetCorrelationConfigNil(b bool)`

 SetCorrelationConfigNil sets the value for CorrelationConfig to be an explicit nil

### UnsetCorrelationConfig
`func (o *ProjectBackwardCompat) UnsetCorrelationConfig()`

UnsetCorrelationConfig ensures that no value is present for CorrelationConfig, not even an explicit nil
### GetAutocaptureOptOut

`func (o *ProjectBackwardCompat) GetAutocaptureOptOut() bool`

GetAutocaptureOptOut returns the AutocaptureOptOut field if non-nil, zero value otherwise.

### GetAutocaptureOptOutOk

`func (o *ProjectBackwardCompat) GetAutocaptureOptOutOk() (*bool, bool)`

GetAutocaptureOptOutOk returns a tuple with the AutocaptureOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureOptOut

`func (o *ProjectBackwardCompat) SetAutocaptureOptOut(v bool)`

SetAutocaptureOptOut sets AutocaptureOptOut field to given value.

### HasAutocaptureOptOut

`func (o *ProjectBackwardCompat) HasAutocaptureOptOut() bool`

HasAutocaptureOptOut returns a boolean if a field has been set.

### SetAutocaptureOptOutNil

`func (o *ProjectBackwardCompat) SetAutocaptureOptOutNil(b bool)`

 SetAutocaptureOptOutNil sets the value for AutocaptureOptOut to be an explicit nil

### UnsetAutocaptureOptOut
`func (o *ProjectBackwardCompat) UnsetAutocaptureOptOut()`

UnsetAutocaptureOptOut ensures that no value is present for AutocaptureOptOut, not even an explicit nil
### GetAutocaptureExceptionsOptIn

`func (o *ProjectBackwardCompat) GetAutocaptureExceptionsOptIn() bool`

GetAutocaptureExceptionsOptIn returns the AutocaptureExceptionsOptIn field if non-nil, zero value otherwise.

### GetAutocaptureExceptionsOptInOk

`func (o *ProjectBackwardCompat) GetAutocaptureExceptionsOptInOk() (*bool, bool)`

GetAutocaptureExceptionsOptInOk returns a tuple with the AutocaptureExceptionsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureExceptionsOptIn

`func (o *ProjectBackwardCompat) SetAutocaptureExceptionsOptIn(v bool)`

SetAutocaptureExceptionsOptIn sets AutocaptureExceptionsOptIn field to given value.

### HasAutocaptureExceptionsOptIn

`func (o *ProjectBackwardCompat) HasAutocaptureExceptionsOptIn() bool`

HasAutocaptureExceptionsOptIn returns a boolean if a field has been set.

### SetAutocaptureExceptionsOptInNil

`func (o *ProjectBackwardCompat) SetAutocaptureExceptionsOptInNil(b bool)`

 SetAutocaptureExceptionsOptInNil sets the value for AutocaptureExceptionsOptIn to be an explicit nil

### UnsetAutocaptureExceptionsOptIn
`func (o *ProjectBackwardCompat) UnsetAutocaptureExceptionsOptIn()`

UnsetAutocaptureExceptionsOptIn ensures that no value is present for AutocaptureExceptionsOptIn, not even an explicit nil
### GetAutocaptureWebVitalsOptIn

`func (o *ProjectBackwardCompat) GetAutocaptureWebVitalsOptIn() bool`

GetAutocaptureWebVitalsOptIn returns the AutocaptureWebVitalsOptIn field if non-nil, zero value otherwise.

### GetAutocaptureWebVitalsOptInOk

`func (o *ProjectBackwardCompat) GetAutocaptureWebVitalsOptInOk() (*bool, bool)`

GetAutocaptureWebVitalsOptInOk returns a tuple with the AutocaptureWebVitalsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureWebVitalsOptIn

`func (o *ProjectBackwardCompat) SetAutocaptureWebVitalsOptIn(v bool)`

SetAutocaptureWebVitalsOptIn sets AutocaptureWebVitalsOptIn field to given value.

### HasAutocaptureWebVitalsOptIn

`func (o *ProjectBackwardCompat) HasAutocaptureWebVitalsOptIn() bool`

HasAutocaptureWebVitalsOptIn returns a boolean if a field has been set.

### SetAutocaptureWebVitalsOptInNil

`func (o *ProjectBackwardCompat) SetAutocaptureWebVitalsOptInNil(b bool)`

 SetAutocaptureWebVitalsOptInNil sets the value for AutocaptureWebVitalsOptIn to be an explicit nil

### UnsetAutocaptureWebVitalsOptIn
`func (o *ProjectBackwardCompat) UnsetAutocaptureWebVitalsOptIn()`

UnsetAutocaptureWebVitalsOptIn ensures that no value is present for AutocaptureWebVitalsOptIn, not even an explicit nil
### GetAutocaptureWebVitalsAllowedMetrics

`func (o *ProjectBackwardCompat) GetAutocaptureWebVitalsAllowedMetrics() interface{}`

GetAutocaptureWebVitalsAllowedMetrics returns the AutocaptureWebVitalsAllowedMetrics field if non-nil, zero value otherwise.

### GetAutocaptureWebVitalsAllowedMetricsOk

`func (o *ProjectBackwardCompat) GetAutocaptureWebVitalsAllowedMetricsOk() (*interface{}, bool)`

GetAutocaptureWebVitalsAllowedMetricsOk returns a tuple with the AutocaptureWebVitalsAllowedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureWebVitalsAllowedMetrics

`func (o *ProjectBackwardCompat) SetAutocaptureWebVitalsAllowedMetrics(v interface{})`

SetAutocaptureWebVitalsAllowedMetrics sets AutocaptureWebVitalsAllowedMetrics field to given value.

### HasAutocaptureWebVitalsAllowedMetrics

`func (o *ProjectBackwardCompat) HasAutocaptureWebVitalsAllowedMetrics() bool`

HasAutocaptureWebVitalsAllowedMetrics returns a boolean if a field has been set.

### SetAutocaptureWebVitalsAllowedMetricsNil

`func (o *ProjectBackwardCompat) SetAutocaptureWebVitalsAllowedMetricsNil(b bool)`

 SetAutocaptureWebVitalsAllowedMetricsNil sets the value for AutocaptureWebVitalsAllowedMetrics to be an explicit nil

### UnsetAutocaptureWebVitalsAllowedMetrics
`func (o *ProjectBackwardCompat) UnsetAutocaptureWebVitalsAllowedMetrics()`

UnsetAutocaptureWebVitalsAllowedMetrics ensures that no value is present for AutocaptureWebVitalsAllowedMetrics, not even an explicit nil
### GetAutocaptureExceptionsErrorsToIgnore

`func (o *ProjectBackwardCompat) GetAutocaptureExceptionsErrorsToIgnore() interface{}`

GetAutocaptureExceptionsErrorsToIgnore returns the AutocaptureExceptionsErrorsToIgnore field if non-nil, zero value otherwise.

### GetAutocaptureExceptionsErrorsToIgnoreOk

`func (o *ProjectBackwardCompat) GetAutocaptureExceptionsErrorsToIgnoreOk() (*interface{}, bool)`

GetAutocaptureExceptionsErrorsToIgnoreOk returns a tuple with the AutocaptureExceptionsErrorsToIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocaptureExceptionsErrorsToIgnore

`func (o *ProjectBackwardCompat) SetAutocaptureExceptionsErrorsToIgnore(v interface{})`

SetAutocaptureExceptionsErrorsToIgnore sets AutocaptureExceptionsErrorsToIgnore field to given value.

### HasAutocaptureExceptionsErrorsToIgnore

`func (o *ProjectBackwardCompat) HasAutocaptureExceptionsErrorsToIgnore() bool`

HasAutocaptureExceptionsErrorsToIgnore returns a boolean if a field has been set.

### SetAutocaptureExceptionsErrorsToIgnoreNil

`func (o *ProjectBackwardCompat) SetAutocaptureExceptionsErrorsToIgnoreNil(b bool)`

 SetAutocaptureExceptionsErrorsToIgnoreNil sets the value for AutocaptureExceptionsErrorsToIgnore to be an explicit nil

### UnsetAutocaptureExceptionsErrorsToIgnore
`func (o *ProjectBackwardCompat) UnsetAutocaptureExceptionsErrorsToIgnore()`

UnsetAutocaptureExceptionsErrorsToIgnore ensures that no value is present for AutocaptureExceptionsErrorsToIgnore, not even an explicit nil
### GetCaptureConsoleLogOptIn

`func (o *ProjectBackwardCompat) GetCaptureConsoleLogOptIn() bool`

GetCaptureConsoleLogOptIn returns the CaptureConsoleLogOptIn field if non-nil, zero value otherwise.

### GetCaptureConsoleLogOptInOk

`func (o *ProjectBackwardCompat) GetCaptureConsoleLogOptInOk() (*bool, bool)`

GetCaptureConsoleLogOptInOk returns a tuple with the CaptureConsoleLogOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureConsoleLogOptIn

`func (o *ProjectBackwardCompat) SetCaptureConsoleLogOptIn(v bool)`

SetCaptureConsoleLogOptIn sets CaptureConsoleLogOptIn field to given value.

### HasCaptureConsoleLogOptIn

`func (o *ProjectBackwardCompat) HasCaptureConsoleLogOptIn() bool`

HasCaptureConsoleLogOptIn returns a boolean if a field has been set.

### SetCaptureConsoleLogOptInNil

`func (o *ProjectBackwardCompat) SetCaptureConsoleLogOptInNil(b bool)`

 SetCaptureConsoleLogOptInNil sets the value for CaptureConsoleLogOptIn to be an explicit nil

### UnsetCaptureConsoleLogOptIn
`func (o *ProjectBackwardCompat) UnsetCaptureConsoleLogOptIn()`

UnsetCaptureConsoleLogOptIn ensures that no value is present for CaptureConsoleLogOptIn, not even an explicit nil
### GetCapturePerformanceOptIn

`func (o *ProjectBackwardCompat) GetCapturePerformanceOptIn() bool`

GetCapturePerformanceOptIn returns the CapturePerformanceOptIn field if non-nil, zero value otherwise.

### GetCapturePerformanceOptInOk

`func (o *ProjectBackwardCompat) GetCapturePerformanceOptInOk() (*bool, bool)`

GetCapturePerformanceOptInOk returns a tuple with the CapturePerformanceOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturePerformanceOptIn

`func (o *ProjectBackwardCompat) SetCapturePerformanceOptIn(v bool)`

SetCapturePerformanceOptIn sets CapturePerformanceOptIn field to given value.

### HasCapturePerformanceOptIn

`func (o *ProjectBackwardCompat) HasCapturePerformanceOptIn() bool`

HasCapturePerformanceOptIn returns a boolean if a field has been set.

### SetCapturePerformanceOptInNil

`func (o *ProjectBackwardCompat) SetCapturePerformanceOptInNil(b bool)`

 SetCapturePerformanceOptInNil sets the value for CapturePerformanceOptIn to be an explicit nil

### UnsetCapturePerformanceOptIn
`func (o *ProjectBackwardCompat) UnsetCapturePerformanceOptIn()`

UnsetCapturePerformanceOptIn ensures that no value is present for CapturePerformanceOptIn, not even an explicit nil
### GetSessionRecordingOptIn

`func (o *ProjectBackwardCompat) GetSessionRecordingOptIn() bool`

GetSessionRecordingOptIn returns the SessionRecordingOptIn field if non-nil, zero value otherwise.

### GetSessionRecordingOptInOk

`func (o *ProjectBackwardCompat) GetSessionRecordingOptInOk() (*bool, bool)`

GetSessionRecordingOptInOk returns a tuple with the SessionRecordingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingOptIn

`func (o *ProjectBackwardCompat) SetSessionRecordingOptIn(v bool)`

SetSessionRecordingOptIn sets SessionRecordingOptIn field to given value.

### HasSessionRecordingOptIn

`func (o *ProjectBackwardCompat) HasSessionRecordingOptIn() bool`

HasSessionRecordingOptIn returns a boolean if a field has been set.

### GetSessionRecordingSampleRate

`func (o *ProjectBackwardCompat) GetSessionRecordingSampleRate() float64`

GetSessionRecordingSampleRate returns the SessionRecordingSampleRate field if non-nil, zero value otherwise.

### GetSessionRecordingSampleRateOk

`func (o *ProjectBackwardCompat) GetSessionRecordingSampleRateOk() (*float64, bool)`

GetSessionRecordingSampleRateOk returns a tuple with the SessionRecordingSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingSampleRate

`func (o *ProjectBackwardCompat) SetSessionRecordingSampleRate(v float64)`

SetSessionRecordingSampleRate sets SessionRecordingSampleRate field to given value.

### HasSessionRecordingSampleRate

`func (o *ProjectBackwardCompat) HasSessionRecordingSampleRate() bool`

HasSessionRecordingSampleRate returns a boolean if a field has been set.

### SetSessionRecordingSampleRateNil

`func (o *ProjectBackwardCompat) SetSessionRecordingSampleRateNil(b bool)`

 SetSessionRecordingSampleRateNil sets the value for SessionRecordingSampleRate to be an explicit nil

### UnsetSessionRecordingSampleRate
`func (o *ProjectBackwardCompat) UnsetSessionRecordingSampleRate()`

UnsetSessionRecordingSampleRate ensures that no value is present for SessionRecordingSampleRate, not even an explicit nil
### GetSessionRecordingMinimumDurationMilliseconds

`func (o *ProjectBackwardCompat) GetSessionRecordingMinimumDurationMilliseconds() int32`

GetSessionRecordingMinimumDurationMilliseconds returns the SessionRecordingMinimumDurationMilliseconds field if non-nil, zero value otherwise.

### GetSessionRecordingMinimumDurationMillisecondsOk

`func (o *ProjectBackwardCompat) GetSessionRecordingMinimumDurationMillisecondsOk() (*int32, bool)`

GetSessionRecordingMinimumDurationMillisecondsOk returns a tuple with the SessionRecordingMinimumDurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingMinimumDurationMilliseconds

`func (o *ProjectBackwardCompat) SetSessionRecordingMinimumDurationMilliseconds(v int32)`

SetSessionRecordingMinimumDurationMilliseconds sets SessionRecordingMinimumDurationMilliseconds field to given value.

### HasSessionRecordingMinimumDurationMilliseconds

`func (o *ProjectBackwardCompat) HasSessionRecordingMinimumDurationMilliseconds() bool`

HasSessionRecordingMinimumDurationMilliseconds returns a boolean if a field has been set.

### SetSessionRecordingMinimumDurationMillisecondsNil

`func (o *ProjectBackwardCompat) SetSessionRecordingMinimumDurationMillisecondsNil(b bool)`

 SetSessionRecordingMinimumDurationMillisecondsNil sets the value for SessionRecordingMinimumDurationMilliseconds to be an explicit nil

### UnsetSessionRecordingMinimumDurationMilliseconds
`func (o *ProjectBackwardCompat) UnsetSessionRecordingMinimumDurationMilliseconds()`

UnsetSessionRecordingMinimumDurationMilliseconds ensures that no value is present for SessionRecordingMinimumDurationMilliseconds, not even an explicit nil
### GetSessionRecordingLinkedFlag

`func (o *ProjectBackwardCompat) GetSessionRecordingLinkedFlag() interface{}`

GetSessionRecordingLinkedFlag returns the SessionRecordingLinkedFlag field if non-nil, zero value otherwise.

### GetSessionRecordingLinkedFlagOk

`func (o *ProjectBackwardCompat) GetSessionRecordingLinkedFlagOk() (*interface{}, bool)`

GetSessionRecordingLinkedFlagOk returns a tuple with the SessionRecordingLinkedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingLinkedFlag

`func (o *ProjectBackwardCompat) SetSessionRecordingLinkedFlag(v interface{})`

SetSessionRecordingLinkedFlag sets SessionRecordingLinkedFlag field to given value.

### HasSessionRecordingLinkedFlag

`func (o *ProjectBackwardCompat) HasSessionRecordingLinkedFlag() bool`

HasSessionRecordingLinkedFlag returns a boolean if a field has been set.

### SetSessionRecordingLinkedFlagNil

`func (o *ProjectBackwardCompat) SetSessionRecordingLinkedFlagNil(b bool)`

 SetSessionRecordingLinkedFlagNil sets the value for SessionRecordingLinkedFlag to be an explicit nil

### UnsetSessionRecordingLinkedFlag
`func (o *ProjectBackwardCompat) UnsetSessionRecordingLinkedFlag()`

UnsetSessionRecordingLinkedFlag ensures that no value is present for SessionRecordingLinkedFlag, not even an explicit nil
### GetSessionRecordingNetworkPayloadCaptureConfig

`func (o *ProjectBackwardCompat) GetSessionRecordingNetworkPayloadCaptureConfig() interface{}`

GetSessionRecordingNetworkPayloadCaptureConfig returns the SessionRecordingNetworkPayloadCaptureConfig field if non-nil, zero value otherwise.

### GetSessionRecordingNetworkPayloadCaptureConfigOk

`func (o *ProjectBackwardCompat) GetSessionRecordingNetworkPayloadCaptureConfigOk() (*interface{}, bool)`

GetSessionRecordingNetworkPayloadCaptureConfigOk returns a tuple with the SessionRecordingNetworkPayloadCaptureConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingNetworkPayloadCaptureConfig

`func (o *ProjectBackwardCompat) SetSessionRecordingNetworkPayloadCaptureConfig(v interface{})`

SetSessionRecordingNetworkPayloadCaptureConfig sets SessionRecordingNetworkPayloadCaptureConfig field to given value.

### HasSessionRecordingNetworkPayloadCaptureConfig

`func (o *ProjectBackwardCompat) HasSessionRecordingNetworkPayloadCaptureConfig() bool`

HasSessionRecordingNetworkPayloadCaptureConfig returns a boolean if a field has been set.

### SetSessionRecordingNetworkPayloadCaptureConfigNil

`func (o *ProjectBackwardCompat) SetSessionRecordingNetworkPayloadCaptureConfigNil(b bool)`

 SetSessionRecordingNetworkPayloadCaptureConfigNil sets the value for SessionRecordingNetworkPayloadCaptureConfig to be an explicit nil

### UnsetSessionRecordingNetworkPayloadCaptureConfig
`func (o *ProjectBackwardCompat) UnsetSessionRecordingNetworkPayloadCaptureConfig()`

UnsetSessionRecordingNetworkPayloadCaptureConfig ensures that no value is present for SessionRecordingNetworkPayloadCaptureConfig, not even an explicit nil
### GetSessionRecordingMaskingConfig

`func (o *ProjectBackwardCompat) GetSessionRecordingMaskingConfig() interface{}`

GetSessionRecordingMaskingConfig returns the SessionRecordingMaskingConfig field if non-nil, zero value otherwise.

### GetSessionRecordingMaskingConfigOk

`func (o *ProjectBackwardCompat) GetSessionRecordingMaskingConfigOk() (*interface{}, bool)`

GetSessionRecordingMaskingConfigOk returns a tuple with the SessionRecordingMaskingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingMaskingConfig

`func (o *ProjectBackwardCompat) SetSessionRecordingMaskingConfig(v interface{})`

SetSessionRecordingMaskingConfig sets SessionRecordingMaskingConfig field to given value.

### HasSessionRecordingMaskingConfig

`func (o *ProjectBackwardCompat) HasSessionRecordingMaskingConfig() bool`

HasSessionRecordingMaskingConfig returns a boolean if a field has been set.

### SetSessionRecordingMaskingConfigNil

`func (o *ProjectBackwardCompat) SetSessionRecordingMaskingConfigNil(b bool)`

 SetSessionRecordingMaskingConfigNil sets the value for SessionRecordingMaskingConfig to be an explicit nil

### UnsetSessionRecordingMaskingConfig
`func (o *ProjectBackwardCompat) UnsetSessionRecordingMaskingConfig()`

UnsetSessionRecordingMaskingConfig ensures that no value is present for SessionRecordingMaskingConfig, not even an explicit nil
### GetSessionReplayConfig

`func (o *ProjectBackwardCompat) GetSessionReplayConfig() interface{}`

GetSessionReplayConfig returns the SessionReplayConfig field if non-nil, zero value otherwise.

### GetSessionReplayConfigOk

`func (o *ProjectBackwardCompat) GetSessionReplayConfigOk() (*interface{}, bool)`

GetSessionReplayConfigOk returns a tuple with the SessionReplayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayConfig

`func (o *ProjectBackwardCompat) SetSessionReplayConfig(v interface{})`

SetSessionReplayConfig sets SessionReplayConfig field to given value.

### HasSessionReplayConfig

`func (o *ProjectBackwardCompat) HasSessionReplayConfig() bool`

HasSessionReplayConfig returns a boolean if a field has been set.

### SetSessionReplayConfigNil

`func (o *ProjectBackwardCompat) SetSessionReplayConfigNil(b bool)`

 SetSessionReplayConfigNil sets the value for SessionReplayConfig to be an explicit nil

### UnsetSessionReplayConfig
`func (o *ProjectBackwardCompat) UnsetSessionReplayConfig()`

UnsetSessionReplayConfig ensures that no value is present for SessionReplayConfig, not even an explicit nil
### GetSurveyConfig

`func (o *ProjectBackwardCompat) GetSurveyConfig() interface{}`

GetSurveyConfig returns the SurveyConfig field if non-nil, zero value otherwise.

### GetSurveyConfigOk

`func (o *ProjectBackwardCompat) GetSurveyConfigOk() (*interface{}, bool)`

GetSurveyConfigOk returns a tuple with the SurveyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyConfig

`func (o *ProjectBackwardCompat) SetSurveyConfig(v interface{})`

SetSurveyConfig sets SurveyConfig field to given value.

### HasSurveyConfig

`func (o *ProjectBackwardCompat) HasSurveyConfig() bool`

HasSurveyConfig returns a boolean if a field has been set.

### SetSurveyConfigNil

`func (o *ProjectBackwardCompat) SetSurveyConfigNil(b bool)`

 SetSurveyConfigNil sets the value for SurveyConfig to be an explicit nil

### UnsetSurveyConfig
`func (o *ProjectBackwardCompat) UnsetSurveyConfig()`

UnsetSurveyConfig ensures that no value is present for SurveyConfig, not even an explicit nil
### GetAccessControl

`func (o *ProjectBackwardCompat) GetAccessControl() bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *ProjectBackwardCompat) GetAccessControlOk() (*bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *ProjectBackwardCompat) SetAccessControl(v bool)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *ProjectBackwardCompat) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetWeekStartDay

`func (o *ProjectBackwardCompat) GetWeekStartDay() WeekStartDayEnum`

GetWeekStartDay returns the WeekStartDay field if non-nil, zero value otherwise.

### GetWeekStartDayOk

`func (o *ProjectBackwardCompat) GetWeekStartDayOk() (*WeekStartDayEnum, bool)`

GetWeekStartDayOk returns a tuple with the WeekStartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekStartDay

`func (o *ProjectBackwardCompat) SetWeekStartDay(v WeekStartDayEnum)`

SetWeekStartDay sets WeekStartDay field to given value.

### HasWeekStartDay

`func (o *ProjectBackwardCompat) HasWeekStartDay() bool`

HasWeekStartDay returns a boolean if a field has been set.

### SetWeekStartDayNil

`func (o *ProjectBackwardCompat) SetWeekStartDayNil(b bool)`

 SetWeekStartDayNil sets the value for WeekStartDay to be an explicit nil

### UnsetWeekStartDay
`func (o *ProjectBackwardCompat) UnsetWeekStartDay()`

UnsetWeekStartDay ensures that no value is present for WeekStartDay, not even an explicit nil
### GetPrimaryDashboard

`func (o *ProjectBackwardCompat) GetPrimaryDashboard() int32`

GetPrimaryDashboard returns the PrimaryDashboard field if non-nil, zero value otherwise.

### GetPrimaryDashboardOk

`func (o *ProjectBackwardCompat) GetPrimaryDashboardOk() (*int32, bool)`

GetPrimaryDashboardOk returns a tuple with the PrimaryDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDashboard

`func (o *ProjectBackwardCompat) SetPrimaryDashboard(v int32)`

SetPrimaryDashboard sets PrimaryDashboard field to given value.

### HasPrimaryDashboard

`func (o *ProjectBackwardCompat) HasPrimaryDashboard() bool`

HasPrimaryDashboard returns a boolean if a field has been set.

### SetPrimaryDashboardNil

`func (o *ProjectBackwardCompat) SetPrimaryDashboardNil(b bool)`

 SetPrimaryDashboardNil sets the value for PrimaryDashboard to be an explicit nil

### UnsetPrimaryDashboard
`func (o *ProjectBackwardCompat) UnsetPrimaryDashboard()`

UnsetPrimaryDashboard ensures that no value is present for PrimaryDashboard, not even an explicit nil
### GetLiveEventsColumns

`func (o *ProjectBackwardCompat) GetLiveEventsColumns() []string`

GetLiveEventsColumns returns the LiveEventsColumns field if non-nil, zero value otherwise.

### GetLiveEventsColumnsOk

`func (o *ProjectBackwardCompat) GetLiveEventsColumnsOk() (*[]string, bool)`

GetLiveEventsColumnsOk returns a tuple with the LiveEventsColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveEventsColumns

`func (o *ProjectBackwardCompat) SetLiveEventsColumns(v []string)`

SetLiveEventsColumns sets LiveEventsColumns field to given value.

### HasLiveEventsColumns

`func (o *ProjectBackwardCompat) HasLiveEventsColumns() bool`

HasLiveEventsColumns returns a boolean if a field has been set.

### SetLiveEventsColumnsNil

`func (o *ProjectBackwardCompat) SetLiveEventsColumnsNil(b bool)`

 SetLiveEventsColumnsNil sets the value for LiveEventsColumns to be an explicit nil

### UnsetLiveEventsColumns
`func (o *ProjectBackwardCompat) UnsetLiveEventsColumns()`

UnsetLiveEventsColumns ensures that no value is present for LiveEventsColumns, not even an explicit nil
### GetRecordingDomains

`func (o *ProjectBackwardCompat) GetRecordingDomains() []*string`

GetRecordingDomains returns the RecordingDomains field if non-nil, zero value otherwise.

### GetRecordingDomainsOk

`func (o *ProjectBackwardCompat) GetRecordingDomainsOk() (*[]*string, bool)`

GetRecordingDomainsOk returns a tuple with the RecordingDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingDomains

`func (o *ProjectBackwardCompat) SetRecordingDomains(v []*string)`

SetRecordingDomains sets RecordingDomains field to given value.

### HasRecordingDomains

`func (o *ProjectBackwardCompat) HasRecordingDomains() bool`

HasRecordingDomains returns a boolean if a field has been set.

### SetRecordingDomainsNil

`func (o *ProjectBackwardCompat) SetRecordingDomainsNil(b bool)`

 SetRecordingDomainsNil sets the value for RecordingDomains to be an explicit nil

### UnsetRecordingDomains
`func (o *ProjectBackwardCompat) UnsetRecordingDomains()`

UnsetRecordingDomains ensures that no value is present for RecordingDomains, not even an explicit nil
### GetPersonOnEventsQueryingEnabled

`func (o *ProjectBackwardCompat) GetPersonOnEventsQueryingEnabled() string`

GetPersonOnEventsQueryingEnabled returns the PersonOnEventsQueryingEnabled field if non-nil, zero value otherwise.

### GetPersonOnEventsQueryingEnabledOk

`func (o *ProjectBackwardCompat) GetPersonOnEventsQueryingEnabledOk() (*string, bool)`

GetPersonOnEventsQueryingEnabledOk returns a tuple with the PersonOnEventsQueryingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonOnEventsQueryingEnabled

`func (o *ProjectBackwardCompat) SetPersonOnEventsQueryingEnabled(v string)`

SetPersonOnEventsQueryingEnabled sets PersonOnEventsQueryingEnabled field to given value.


### GetInjectWebApps

`func (o *ProjectBackwardCompat) GetInjectWebApps() bool`

GetInjectWebApps returns the InjectWebApps field if non-nil, zero value otherwise.

### GetInjectWebAppsOk

`func (o *ProjectBackwardCompat) GetInjectWebAppsOk() (*bool, bool)`

GetInjectWebAppsOk returns a tuple with the InjectWebApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjectWebApps

`func (o *ProjectBackwardCompat) SetInjectWebApps(v bool)`

SetInjectWebApps sets InjectWebApps field to given value.

### HasInjectWebApps

`func (o *ProjectBackwardCompat) HasInjectWebApps() bool`

HasInjectWebApps returns a boolean if a field has been set.

### SetInjectWebAppsNil

`func (o *ProjectBackwardCompat) SetInjectWebAppsNil(b bool)`

 SetInjectWebAppsNil sets the value for InjectWebApps to be an explicit nil

### UnsetInjectWebApps
`func (o *ProjectBackwardCompat) UnsetInjectWebApps()`

UnsetInjectWebApps ensures that no value is present for InjectWebApps, not even an explicit nil
### GetExtraSettings

`func (o *ProjectBackwardCompat) GetExtraSettings() interface{}`

GetExtraSettings returns the ExtraSettings field if non-nil, zero value otherwise.

### GetExtraSettingsOk

`func (o *ProjectBackwardCompat) GetExtraSettingsOk() (*interface{}, bool)`

GetExtraSettingsOk returns a tuple with the ExtraSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSettings

`func (o *ProjectBackwardCompat) SetExtraSettings(v interface{})`

SetExtraSettings sets ExtraSettings field to given value.

### HasExtraSettings

`func (o *ProjectBackwardCompat) HasExtraSettings() bool`

HasExtraSettings returns a boolean if a field has been set.

### SetExtraSettingsNil

`func (o *ProjectBackwardCompat) SetExtraSettingsNil(b bool)`

 SetExtraSettingsNil sets the value for ExtraSettings to be an explicit nil

### UnsetExtraSettings
`func (o *ProjectBackwardCompat) UnsetExtraSettings()`

UnsetExtraSettings ensures that no value is present for ExtraSettings, not even an explicit nil
### GetModifiers

`func (o *ProjectBackwardCompat) GetModifiers() interface{}`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ProjectBackwardCompat) GetModifiersOk() (*interface{}, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ProjectBackwardCompat) SetModifiers(v interface{})`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ProjectBackwardCompat) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### SetModifiersNil

`func (o *ProjectBackwardCompat) SetModifiersNil(b bool)`

 SetModifiersNil sets the value for Modifiers to be an explicit nil

### UnsetModifiers
`func (o *ProjectBackwardCompat) UnsetModifiers()`

UnsetModifiers ensures that no value is present for Modifiers, not even an explicit nil
### GetDefaultModifiers

`func (o *ProjectBackwardCompat) GetDefaultModifiers() string`

GetDefaultModifiers returns the DefaultModifiers field if non-nil, zero value otherwise.

### GetDefaultModifiersOk

`func (o *ProjectBackwardCompat) GetDefaultModifiersOk() (*string, bool)`

GetDefaultModifiersOk returns a tuple with the DefaultModifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModifiers

`func (o *ProjectBackwardCompat) SetDefaultModifiers(v string)`

SetDefaultModifiers sets DefaultModifiers field to given value.


### GetHasCompletedOnboardingFor

`func (o *ProjectBackwardCompat) GetHasCompletedOnboardingFor() interface{}`

GetHasCompletedOnboardingFor returns the HasCompletedOnboardingFor field if non-nil, zero value otherwise.

### GetHasCompletedOnboardingForOk

`func (o *ProjectBackwardCompat) GetHasCompletedOnboardingForOk() (*interface{}, bool)`

GetHasCompletedOnboardingForOk returns a tuple with the HasCompletedOnboardingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompletedOnboardingFor

`func (o *ProjectBackwardCompat) SetHasCompletedOnboardingFor(v interface{})`

SetHasCompletedOnboardingFor sets HasCompletedOnboardingFor field to given value.

### HasHasCompletedOnboardingFor

`func (o *ProjectBackwardCompat) HasHasCompletedOnboardingFor() bool`

HasHasCompletedOnboardingFor returns a boolean if a field has been set.

### SetHasCompletedOnboardingForNil

`func (o *ProjectBackwardCompat) SetHasCompletedOnboardingForNil(b bool)`

 SetHasCompletedOnboardingForNil sets the value for HasCompletedOnboardingFor to be an explicit nil

### UnsetHasCompletedOnboardingFor
`func (o *ProjectBackwardCompat) UnsetHasCompletedOnboardingFor()`

UnsetHasCompletedOnboardingFor ensures that no value is present for HasCompletedOnboardingFor, not even an explicit nil
### GetSurveysOptIn

`func (o *ProjectBackwardCompat) GetSurveysOptIn() bool`

GetSurveysOptIn returns the SurveysOptIn field if non-nil, zero value otherwise.

### GetSurveysOptInOk

`func (o *ProjectBackwardCompat) GetSurveysOptInOk() (*bool, bool)`

GetSurveysOptInOk returns a tuple with the SurveysOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveysOptIn

`func (o *ProjectBackwardCompat) SetSurveysOptIn(v bool)`

SetSurveysOptIn sets SurveysOptIn field to given value.

### HasSurveysOptIn

`func (o *ProjectBackwardCompat) HasSurveysOptIn() bool`

HasSurveysOptIn returns a boolean if a field has been set.

### SetSurveysOptInNil

`func (o *ProjectBackwardCompat) SetSurveysOptInNil(b bool)`

 SetSurveysOptInNil sets the value for SurveysOptIn to be an explicit nil

### UnsetSurveysOptIn
`func (o *ProjectBackwardCompat) UnsetSurveysOptIn()`

UnsetSurveysOptIn ensures that no value is present for SurveysOptIn, not even an explicit nil
### GetHeatmapsOptIn

`func (o *ProjectBackwardCompat) GetHeatmapsOptIn() bool`

GetHeatmapsOptIn returns the HeatmapsOptIn field if non-nil, zero value otherwise.

### GetHeatmapsOptInOk

`func (o *ProjectBackwardCompat) GetHeatmapsOptInOk() (*bool, bool)`

GetHeatmapsOptInOk returns a tuple with the HeatmapsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapsOptIn

`func (o *ProjectBackwardCompat) SetHeatmapsOptIn(v bool)`

SetHeatmapsOptIn sets HeatmapsOptIn field to given value.

### HasHeatmapsOptIn

`func (o *ProjectBackwardCompat) HasHeatmapsOptIn() bool`

HasHeatmapsOptIn returns a boolean if a field has been set.

### SetHeatmapsOptInNil

`func (o *ProjectBackwardCompat) SetHeatmapsOptInNil(b bool)`

 SetHeatmapsOptInNil sets the value for HeatmapsOptIn to be an explicit nil

### UnsetHeatmapsOptIn
`func (o *ProjectBackwardCompat) UnsetHeatmapsOptIn()`

UnsetHeatmapsOptIn ensures that no value is present for HeatmapsOptIn, not even an explicit nil
### GetProductIntents

`func (o *ProjectBackwardCompat) GetProductIntents() string`

GetProductIntents returns the ProductIntents field if non-nil, zero value otherwise.

### GetProductIntentsOk

`func (o *ProjectBackwardCompat) GetProductIntentsOk() (*string, bool)`

GetProductIntentsOk returns a tuple with the ProductIntents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIntents

`func (o *ProjectBackwardCompat) SetProductIntents(v string)`

SetProductIntents sets ProductIntents field to given value.


### GetFlagsPersistenceDefault

`func (o *ProjectBackwardCompat) GetFlagsPersistenceDefault() bool`

GetFlagsPersistenceDefault returns the FlagsPersistenceDefault field if non-nil, zero value otherwise.

### GetFlagsPersistenceDefaultOk

`func (o *ProjectBackwardCompat) GetFlagsPersistenceDefaultOk() (*bool, bool)`

GetFlagsPersistenceDefaultOk returns a tuple with the FlagsPersistenceDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagsPersistenceDefault

`func (o *ProjectBackwardCompat) SetFlagsPersistenceDefault(v bool)`

SetFlagsPersistenceDefault sets FlagsPersistenceDefault field to given value.

### HasFlagsPersistenceDefault

`func (o *ProjectBackwardCompat) HasFlagsPersistenceDefault() bool`

HasFlagsPersistenceDefault returns a boolean if a field has been set.

### SetFlagsPersistenceDefaultNil

`func (o *ProjectBackwardCompat) SetFlagsPersistenceDefaultNil(b bool)`

 SetFlagsPersistenceDefaultNil sets the value for FlagsPersistenceDefault to be an explicit nil

### UnsetFlagsPersistenceDefault
`func (o *ProjectBackwardCompat) UnsetFlagsPersistenceDefault()`

UnsetFlagsPersistenceDefault ensures that no value is present for FlagsPersistenceDefault, not even an explicit nil
### GetSecretApiToken

`func (o *ProjectBackwardCompat) GetSecretApiToken() string`

GetSecretApiToken returns the SecretApiToken field if non-nil, zero value otherwise.

### GetSecretApiTokenOk

`func (o *ProjectBackwardCompat) GetSecretApiTokenOk() (*string, bool)`

GetSecretApiTokenOk returns a tuple with the SecretApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiToken

`func (o *ProjectBackwardCompat) SetSecretApiToken(v string)`

SetSecretApiToken sets SecretApiToken field to given value.


### SetSecretApiTokenNil

`func (o *ProjectBackwardCompat) SetSecretApiTokenNil(b bool)`

 SetSecretApiTokenNil sets the value for SecretApiToken to be an explicit nil

### UnsetSecretApiToken
`func (o *ProjectBackwardCompat) UnsetSecretApiToken()`

UnsetSecretApiToken ensures that no value is present for SecretApiToken, not even an explicit nil
### GetSecretApiTokenBackup

`func (o *ProjectBackwardCompat) GetSecretApiTokenBackup() string`

GetSecretApiTokenBackup returns the SecretApiTokenBackup field if non-nil, zero value otherwise.

### GetSecretApiTokenBackupOk

`func (o *ProjectBackwardCompat) GetSecretApiTokenBackupOk() (*string, bool)`

GetSecretApiTokenBackupOk returns a tuple with the SecretApiTokenBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiTokenBackup

`func (o *ProjectBackwardCompat) SetSecretApiTokenBackup(v string)`

SetSecretApiTokenBackup sets SecretApiTokenBackup field to given value.


### SetSecretApiTokenBackupNil

`func (o *ProjectBackwardCompat) SetSecretApiTokenBackupNil(b bool)`

 SetSecretApiTokenBackupNil sets the value for SecretApiTokenBackup to be an explicit nil

### UnsetSecretApiTokenBackup
`func (o *ProjectBackwardCompat) UnsetSecretApiTokenBackup()`

UnsetSecretApiTokenBackup ensures that no value is present for SecretApiTokenBackup, not even an explicit nil
### GetReceiveOrgLevelActivityLogs

`func (o *ProjectBackwardCompat) GetReceiveOrgLevelActivityLogs() bool`

GetReceiveOrgLevelActivityLogs returns the ReceiveOrgLevelActivityLogs field if non-nil, zero value otherwise.

### GetReceiveOrgLevelActivityLogsOk

`func (o *ProjectBackwardCompat) GetReceiveOrgLevelActivityLogsOk() (*bool, bool)`

GetReceiveOrgLevelActivityLogsOk returns a tuple with the ReceiveOrgLevelActivityLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveOrgLevelActivityLogs

`func (o *ProjectBackwardCompat) SetReceiveOrgLevelActivityLogs(v bool)`

SetReceiveOrgLevelActivityLogs sets ReceiveOrgLevelActivityLogs field to given value.

### HasReceiveOrgLevelActivityLogs

`func (o *ProjectBackwardCompat) HasReceiveOrgLevelActivityLogs() bool`

HasReceiveOrgLevelActivityLogs returns a boolean if a field has been set.

### SetReceiveOrgLevelActivityLogsNil

`func (o *ProjectBackwardCompat) SetReceiveOrgLevelActivityLogsNil(b bool)`

 SetReceiveOrgLevelActivityLogsNil sets the value for ReceiveOrgLevelActivityLogs to be an explicit nil

### UnsetReceiveOrgLevelActivityLogs
`func (o *ProjectBackwardCompat) UnsetReceiveOrgLevelActivityLogs()`

UnsetReceiveOrgLevelActivityLogs ensures that no value is present for ReceiveOrgLevelActivityLogs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


