# HogQLQueryModifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BounceRateDurationSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**BounceRatePageViewMode** | Pointer to [**BounceRatePageViewMode**](BounceRatePageViewMode.md) |  | [optional] 
**ConvertToProjectTimezone** | Pointer to **NullableBool** |  | [optional] 
**CustomChannelTypeRules** | Pointer to [**[]CustomChannelRule**](CustomChannelRule.md) |  | [optional] 
**DataWarehouseEventsModifiers** | Pointer to [**[]DataWarehouseEventsModifier**](DataWarehouseEventsModifier.md) |  | [optional] 
**Debug** | Pointer to **NullableBool** |  | [optional] 
**FormatCsvAllowDoubleQuotes** | Pointer to **NullableBool** |  | [optional] 
**InCohortVia** | Pointer to [**InCohortVia**](InCohortVia.md) |  | [optional] 
**MaterializationMode** | Pointer to [**MaterializationMode**](MaterializationMode.md) |  | [optional] 
**OptimizeJoinedFilters** | Pointer to **NullableBool** |  | [optional] 
**OptimizeProjections** | Pointer to **NullableBool** |  | [optional] 
**PersonsArgMaxVersion** | Pointer to [**PersonsArgMaxVersion**](PersonsArgMaxVersion.md) |  | [optional] 
**PersonsJoinMode** | Pointer to [**PersonsJoinMode**](PersonsJoinMode.md) |  | [optional] 
**PersonsOnEventsMode** | Pointer to [**PersonsOnEventsMode**](PersonsOnEventsMode.md) |  | [optional] 
**PropertyGroupsMode** | Pointer to [**PropertyGroupsMode**](PropertyGroupsMode.md) |  | [optional] 
**S3TableUseInvalidColumns** | Pointer to **NullableBool** |  | [optional] 
**SessionTableVersion** | Pointer to [**SessionTableVersion**](SessionTableVersion.md) |  | [optional] 
**SessionsV2JoinMode** | Pointer to [**SessionsV2JoinMode**](SessionsV2JoinMode.md) |  | [optional] 
**Timings** | Pointer to **NullableBool** |  | [optional] 
**UseMaterializedViews** | Pointer to **NullableBool** |  | [optional] 
**UsePreaggregatedTableTransforms** | Pointer to **NullableBool** | Try to automatically convert HogQL queries to use preaggregated tables at the AST level * | [optional] 
**UsePresortedEventsTable** | Pointer to **NullableBool** |  | [optional] 
**UseWebAnalyticsPreAggregatedTables** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewHogQLQueryModifiers

`func NewHogQLQueryModifiers() *HogQLQueryModifiers`

NewHogQLQueryModifiers instantiates a new HogQLQueryModifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLQueryModifiersWithDefaults

`func NewHogQLQueryModifiersWithDefaults() *HogQLQueryModifiers`

NewHogQLQueryModifiersWithDefaults instantiates a new HogQLQueryModifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounceRateDurationSeconds

`func (o *HogQLQueryModifiers) GetBounceRateDurationSeconds() float32`

GetBounceRateDurationSeconds returns the BounceRateDurationSeconds field if non-nil, zero value otherwise.

### GetBounceRateDurationSecondsOk

`func (o *HogQLQueryModifiers) GetBounceRateDurationSecondsOk() (*float32, bool)`

GetBounceRateDurationSecondsOk returns a tuple with the BounceRateDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceRateDurationSeconds

`func (o *HogQLQueryModifiers) SetBounceRateDurationSeconds(v float32)`

SetBounceRateDurationSeconds sets BounceRateDurationSeconds field to given value.

### HasBounceRateDurationSeconds

`func (o *HogQLQueryModifiers) HasBounceRateDurationSeconds() bool`

HasBounceRateDurationSeconds returns a boolean if a field has been set.

### SetBounceRateDurationSecondsNil

`func (o *HogQLQueryModifiers) SetBounceRateDurationSecondsNil(b bool)`

 SetBounceRateDurationSecondsNil sets the value for BounceRateDurationSeconds to be an explicit nil

### UnsetBounceRateDurationSeconds
`func (o *HogQLQueryModifiers) UnsetBounceRateDurationSeconds()`

UnsetBounceRateDurationSeconds ensures that no value is present for BounceRateDurationSeconds, not even an explicit nil
### GetBounceRatePageViewMode

`func (o *HogQLQueryModifiers) GetBounceRatePageViewMode() BounceRatePageViewMode`

GetBounceRatePageViewMode returns the BounceRatePageViewMode field if non-nil, zero value otherwise.

### GetBounceRatePageViewModeOk

`func (o *HogQLQueryModifiers) GetBounceRatePageViewModeOk() (*BounceRatePageViewMode, bool)`

GetBounceRatePageViewModeOk returns a tuple with the BounceRatePageViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceRatePageViewMode

`func (o *HogQLQueryModifiers) SetBounceRatePageViewMode(v BounceRatePageViewMode)`

SetBounceRatePageViewMode sets BounceRatePageViewMode field to given value.

### HasBounceRatePageViewMode

`func (o *HogQLQueryModifiers) HasBounceRatePageViewMode() bool`

HasBounceRatePageViewMode returns a boolean if a field has been set.

### GetConvertToProjectTimezone

`func (o *HogQLQueryModifiers) GetConvertToProjectTimezone() bool`

GetConvertToProjectTimezone returns the ConvertToProjectTimezone field if non-nil, zero value otherwise.

### GetConvertToProjectTimezoneOk

`func (o *HogQLQueryModifiers) GetConvertToProjectTimezoneOk() (*bool, bool)`

GetConvertToProjectTimezoneOk returns a tuple with the ConvertToProjectTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertToProjectTimezone

`func (o *HogQLQueryModifiers) SetConvertToProjectTimezone(v bool)`

SetConvertToProjectTimezone sets ConvertToProjectTimezone field to given value.

### HasConvertToProjectTimezone

`func (o *HogQLQueryModifiers) HasConvertToProjectTimezone() bool`

HasConvertToProjectTimezone returns a boolean if a field has been set.

### SetConvertToProjectTimezoneNil

`func (o *HogQLQueryModifiers) SetConvertToProjectTimezoneNil(b bool)`

 SetConvertToProjectTimezoneNil sets the value for ConvertToProjectTimezone to be an explicit nil

### UnsetConvertToProjectTimezone
`func (o *HogQLQueryModifiers) UnsetConvertToProjectTimezone()`

UnsetConvertToProjectTimezone ensures that no value is present for ConvertToProjectTimezone, not even an explicit nil
### GetCustomChannelTypeRules

`func (o *HogQLQueryModifiers) GetCustomChannelTypeRules() []CustomChannelRule`

GetCustomChannelTypeRules returns the CustomChannelTypeRules field if non-nil, zero value otherwise.

### GetCustomChannelTypeRulesOk

`func (o *HogQLQueryModifiers) GetCustomChannelTypeRulesOk() (*[]CustomChannelRule, bool)`

GetCustomChannelTypeRulesOk returns a tuple with the CustomChannelTypeRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomChannelTypeRules

`func (o *HogQLQueryModifiers) SetCustomChannelTypeRules(v []CustomChannelRule)`

SetCustomChannelTypeRules sets CustomChannelTypeRules field to given value.

### HasCustomChannelTypeRules

`func (o *HogQLQueryModifiers) HasCustomChannelTypeRules() bool`

HasCustomChannelTypeRules returns a boolean if a field has been set.

### SetCustomChannelTypeRulesNil

`func (o *HogQLQueryModifiers) SetCustomChannelTypeRulesNil(b bool)`

 SetCustomChannelTypeRulesNil sets the value for CustomChannelTypeRules to be an explicit nil

### UnsetCustomChannelTypeRules
`func (o *HogQLQueryModifiers) UnsetCustomChannelTypeRules()`

UnsetCustomChannelTypeRules ensures that no value is present for CustomChannelTypeRules, not even an explicit nil
### GetDataWarehouseEventsModifiers

`func (o *HogQLQueryModifiers) GetDataWarehouseEventsModifiers() []DataWarehouseEventsModifier`

GetDataWarehouseEventsModifiers returns the DataWarehouseEventsModifiers field if non-nil, zero value otherwise.

### GetDataWarehouseEventsModifiersOk

`func (o *HogQLQueryModifiers) GetDataWarehouseEventsModifiersOk() (*[]DataWarehouseEventsModifier, bool)`

GetDataWarehouseEventsModifiersOk returns a tuple with the DataWarehouseEventsModifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWarehouseEventsModifiers

`func (o *HogQLQueryModifiers) SetDataWarehouseEventsModifiers(v []DataWarehouseEventsModifier)`

SetDataWarehouseEventsModifiers sets DataWarehouseEventsModifiers field to given value.

### HasDataWarehouseEventsModifiers

`func (o *HogQLQueryModifiers) HasDataWarehouseEventsModifiers() bool`

HasDataWarehouseEventsModifiers returns a boolean if a field has been set.

### SetDataWarehouseEventsModifiersNil

`func (o *HogQLQueryModifiers) SetDataWarehouseEventsModifiersNil(b bool)`

 SetDataWarehouseEventsModifiersNil sets the value for DataWarehouseEventsModifiers to be an explicit nil

### UnsetDataWarehouseEventsModifiers
`func (o *HogQLQueryModifiers) UnsetDataWarehouseEventsModifiers()`

UnsetDataWarehouseEventsModifiers ensures that no value is present for DataWarehouseEventsModifiers, not even an explicit nil
### GetDebug

`func (o *HogQLQueryModifiers) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *HogQLQueryModifiers) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *HogQLQueryModifiers) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *HogQLQueryModifiers) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### SetDebugNil

`func (o *HogQLQueryModifiers) SetDebugNil(b bool)`

 SetDebugNil sets the value for Debug to be an explicit nil

### UnsetDebug
`func (o *HogQLQueryModifiers) UnsetDebug()`

UnsetDebug ensures that no value is present for Debug, not even an explicit nil
### GetFormatCsvAllowDoubleQuotes

`func (o *HogQLQueryModifiers) GetFormatCsvAllowDoubleQuotes() bool`

GetFormatCsvAllowDoubleQuotes returns the FormatCsvAllowDoubleQuotes field if non-nil, zero value otherwise.

### GetFormatCsvAllowDoubleQuotesOk

`func (o *HogQLQueryModifiers) GetFormatCsvAllowDoubleQuotesOk() (*bool, bool)`

GetFormatCsvAllowDoubleQuotesOk returns a tuple with the FormatCsvAllowDoubleQuotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatCsvAllowDoubleQuotes

`func (o *HogQLQueryModifiers) SetFormatCsvAllowDoubleQuotes(v bool)`

SetFormatCsvAllowDoubleQuotes sets FormatCsvAllowDoubleQuotes field to given value.

### HasFormatCsvAllowDoubleQuotes

`func (o *HogQLQueryModifiers) HasFormatCsvAllowDoubleQuotes() bool`

HasFormatCsvAllowDoubleQuotes returns a boolean if a field has been set.

### SetFormatCsvAllowDoubleQuotesNil

`func (o *HogQLQueryModifiers) SetFormatCsvAllowDoubleQuotesNil(b bool)`

 SetFormatCsvAllowDoubleQuotesNil sets the value for FormatCsvAllowDoubleQuotes to be an explicit nil

### UnsetFormatCsvAllowDoubleQuotes
`func (o *HogQLQueryModifiers) UnsetFormatCsvAllowDoubleQuotes()`

UnsetFormatCsvAllowDoubleQuotes ensures that no value is present for FormatCsvAllowDoubleQuotes, not even an explicit nil
### GetInCohortVia

`func (o *HogQLQueryModifiers) GetInCohortVia() InCohortVia`

GetInCohortVia returns the InCohortVia field if non-nil, zero value otherwise.

### GetInCohortViaOk

`func (o *HogQLQueryModifiers) GetInCohortViaOk() (*InCohortVia, bool)`

GetInCohortViaOk returns a tuple with the InCohortVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCohortVia

`func (o *HogQLQueryModifiers) SetInCohortVia(v InCohortVia)`

SetInCohortVia sets InCohortVia field to given value.

### HasInCohortVia

`func (o *HogQLQueryModifiers) HasInCohortVia() bool`

HasInCohortVia returns a boolean if a field has been set.

### GetMaterializationMode

`func (o *HogQLQueryModifiers) GetMaterializationMode() MaterializationMode`

GetMaterializationMode returns the MaterializationMode field if non-nil, zero value otherwise.

### GetMaterializationModeOk

`func (o *HogQLQueryModifiers) GetMaterializationModeOk() (*MaterializationMode, bool)`

GetMaterializationModeOk returns a tuple with the MaterializationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterializationMode

`func (o *HogQLQueryModifiers) SetMaterializationMode(v MaterializationMode)`

SetMaterializationMode sets MaterializationMode field to given value.

### HasMaterializationMode

`func (o *HogQLQueryModifiers) HasMaterializationMode() bool`

HasMaterializationMode returns a boolean if a field has been set.

### GetOptimizeJoinedFilters

`func (o *HogQLQueryModifiers) GetOptimizeJoinedFilters() bool`

GetOptimizeJoinedFilters returns the OptimizeJoinedFilters field if non-nil, zero value otherwise.

### GetOptimizeJoinedFiltersOk

`func (o *HogQLQueryModifiers) GetOptimizeJoinedFiltersOk() (*bool, bool)`

GetOptimizeJoinedFiltersOk returns a tuple with the OptimizeJoinedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeJoinedFilters

`func (o *HogQLQueryModifiers) SetOptimizeJoinedFilters(v bool)`

SetOptimizeJoinedFilters sets OptimizeJoinedFilters field to given value.

### HasOptimizeJoinedFilters

`func (o *HogQLQueryModifiers) HasOptimizeJoinedFilters() bool`

HasOptimizeJoinedFilters returns a boolean if a field has been set.

### SetOptimizeJoinedFiltersNil

`func (o *HogQLQueryModifiers) SetOptimizeJoinedFiltersNil(b bool)`

 SetOptimizeJoinedFiltersNil sets the value for OptimizeJoinedFilters to be an explicit nil

### UnsetOptimizeJoinedFilters
`func (o *HogQLQueryModifiers) UnsetOptimizeJoinedFilters()`

UnsetOptimizeJoinedFilters ensures that no value is present for OptimizeJoinedFilters, not even an explicit nil
### GetOptimizeProjections

`func (o *HogQLQueryModifiers) GetOptimizeProjections() bool`

GetOptimizeProjections returns the OptimizeProjections field if non-nil, zero value otherwise.

### GetOptimizeProjectionsOk

`func (o *HogQLQueryModifiers) GetOptimizeProjectionsOk() (*bool, bool)`

GetOptimizeProjectionsOk returns a tuple with the OptimizeProjections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeProjections

`func (o *HogQLQueryModifiers) SetOptimizeProjections(v bool)`

SetOptimizeProjections sets OptimizeProjections field to given value.

### HasOptimizeProjections

`func (o *HogQLQueryModifiers) HasOptimizeProjections() bool`

HasOptimizeProjections returns a boolean if a field has been set.

### SetOptimizeProjectionsNil

`func (o *HogQLQueryModifiers) SetOptimizeProjectionsNil(b bool)`

 SetOptimizeProjectionsNil sets the value for OptimizeProjections to be an explicit nil

### UnsetOptimizeProjections
`func (o *HogQLQueryModifiers) UnsetOptimizeProjections()`

UnsetOptimizeProjections ensures that no value is present for OptimizeProjections, not even an explicit nil
### GetPersonsArgMaxVersion

`func (o *HogQLQueryModifiers) GetPersonsArgMaxVersion() PersonsArgMaxVersion`

GetPersonsArgMaxVersion returns the PersonsArgMaxVersion field if non-nil, zero value otherwise.

### GetPersonsArgMaxVersionOk

`func (o *HogQLQueryModifiers) GetPersonsArgMaxVersionOk() (*PersonsArgMaxVersion, bool)`

GetPersonsArgMaxVersionOk returns a tuple with the PersonsArgMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonsArgMaxVersion

`func (o *HogQLQueryModifiers) SetPersonsArgMaxVersion(v PersonsArgMaxVersion)`

SetPersonsArgMaxVersion sets PersonsArgMaxVersion field to given value.

### HasPersonsArgMaxVersion

`func (o *HogQLQueryModifiers) HasPersonsArgMaxVersion() bool`

HasPersonsArgMaxVersion returns a boolean if a field has been set.

### GetPersonsJoinMode

`func (o *HogQLQueryModifiers) GetPersonsJoinMode() PersonsJoinMode`

GetPersonsJoinMode returns the PersonsJoinMode field if non-nil, zero value otherwise.

### GetPersonsJoinModeOk

`func (o *HogQLQueryModifiers) GetPersonsJoinModeOk() (*PersonsJoinMode, bool)`

GetPersonsJoinModeOk returns a tuple with the PersonsJoinMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonsJoinMode

`func (o *HogQLQueryModifiers) SetPersonsJoinMode(v PersonsJoinMode)`

SetPersonsJoinMode sets PersonsJoinMode field to given value.

### HasPersonsJoinMode

`func (o *HogQLQueryModifiers) HasPersonsJoinMode() bool`

HasPersonsJoinMode returns a boolean if a field has been set.

### GetPersonsOnEventsMode

`func (o *HogQLQueryModifiers) GetPersonsOnEventsMode() PersonsOnEventsMode`

GetPersonsOnEventsMode returns the PersonsOnEventsMode field if non-nil, zero value otherwise.

### GetPersonsOnEventsModeOk

`func (o *HogQLQueryModifiers) GetPersonsOnEventsModeOk() (*PersonsOnEventsMode, bool)`

GetPersonsOnEventsModeOk returns a tuple with the PersonsOnEventsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonsOnEventsMode

`func (o *HogQLQueryModifiers) SetPersonsOnEventsMode(v PersonsOnEventsMode)`

SetPersonsOnEventsMode sets PersonsOnEventsMode field to given value.

### HasPersonsOnEventsMode

`func (o *HogQLQueryModifiers) HasPersonsOnEventsMode() bool`

HasPersonsOnEventsMode returns a boolean if a field has been set.

### GetPropertyGroupsMode

`func (o *HogQLQueryModifiers) GetPropertyGroupsMode() PropertyGroupsMode`

GetPropertyGroupsMode returns the PropertyGroupsMode field if non-nil, zero value otherwise.

### GetPropertyGroupsModeOk

`func (o *HogQLQueryModifiers) GetPropertyGroupsModeOk() (*PropertyGroupsMode, bool)`

GetPropertyGroupsModeOk returns a tuple with the PropertyGroupsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyGroupsMode

`func (o *HogQLQueryModifiers) SetPropertyGroupsMode(v PropertyGroupsMode)`

SetPropertyGroupsMode sets PropertyGroupsMode field to given value.

### HasPropertyGroupsMode

`func (o *HogQLQueryModifiers) HasPropertyGroupsMode() bool`

HasPropertyGroupsMode returns a boolean if a field has been set.

### GetS3TableUseInvalidColumns

`func (o *HogQLQueryModifiers) GetS3TableUseInvalidColumns() bool`

GetS3TableUseInvalidColumns returns the S3TableUseInvalidColumns field if non-nil, zero value otherwise.

### GetS3TableUseInvalidColumnsOk

`func (o *HogQLQueryModifiers) GetS3TableUseInvalidColumnsOk() (*bool, bool)`

GetS3TableUseInvalidColumnsOk returns a tuple with the S3TableUseInvalidColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3TableUseInvalidColumns

`func (o *HogQLQueryModifiers) SetS3TableUseInvalidColumns(v bool)`

SetS3TableUseInvalidColumns sets S3TableUseInvalidColumns field to given value.

### HasS3TableUseInvalidColumns

`func (o *HogQLQueryModifiers) HasS3TableUseInvalidColumns() bool`

HasS3TableUseInvalidColumns returns a boolean if a field has been set.

### SetS3TableUseInvalidColumnsNil

`func (o *HogQLQueryModifiers) SetS3TableUseInvalidColumnsNil(b bool)`

 SetS3TableUseInvalidColumnsNil sets the value for S3TableUseInvalidColumns to be an explicit nil

### UnsetS3TableUseInvalidColumns
`func (o *HogQLQueryModifiers) UnsetS3TableUseInvalidColumns()`

UnsetS3TableUseInvalidColumns ensures that no value is present for S3TableUseInvalidColumns, not even an explicit nil
### GetSessionTableVersion

`func (o *HogQLQueryModifiers) GetSessionTableVersion() SessionTableVersion`

GetSessionTableVersion returns the SessionTableVersion field if non-nil, zero value otherwise.

### GetSessionTableVersionOk

`func (o *HogQLQueryModifiers) GetSessionTableVersionOk() (*SessionTableVersion, bool)`

GetSessionTableVersionOk returns a tuple with the SessionTableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTableVersion

`func (o *HogQLQueryModifiers) SetSessionTableVersion(v SessionTableVersion)`

SetSessionTableVersion sets SessionTableVersion field to given value.

### HasSessionTableVersion

`func (o *HogQLQueryModifiers) HasSessionTableVersion() bool`

HasSessionTableVersion returns a boolean if a field has been set.

### GetSessionsV2JoinMode

`func (o *HogQLQueryModifiers) GetSessionsV2JoinMode() SessionsV2JoinMode`

GetSessionsV2JoinMode returns the SessionsV2JoinMode field if non-nil, zero value otherwise.

### GetSessionsV2JoinModeOk

`func (o *HogQLQueryModifiers) GetSessionsV2JoinModeOk() (*SessionsV2JoinMode, bool)`

GetSessionsV2JoinModeOk returns a tuple with the SessionsV2JoinMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsV2JoinMode

`func (o *HogQLQueryModifiers) SetSessionsV2JoinMode(v SessionsV2JoinMode)`

SetSessionsV2JoinMode sets SessionsV2JoinMode field to given value.

### HasSessionsV2JoinMode

`func (o *HogQLQueryModifiers) HasSessionsV2JoinMode() bool`

HasSessionsV2JoinMode returns a boolean if a field has been set.

### GetTimings

`func (o *HogQLQueryModifiers) GetTimings() bool`

GetTimings returns the Timings field if non-nil, zero value otherwise.

### GetTimingsOk

`func (o *HogQLQueryModifiers) GetTimingsOk() (*bool, bool)`

GetTimingsOk returns a tuple with the Timings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimings

`func (o *HogQLQueryModifiers) SetTimings(v bool)`

SetTimings sets Timings field to given value.

### HasTimings

`func (o *HogQLQueryModifiers) HasTimings() bool`

HasTimings returns a boolean if a field has been set.

### SetTimingsNil

`func (o *HogQLQueryModifiers) SetTimingsNil(b bool)`

 SetTimingsNil sets the value for Timings to be an explicit nil

### UnsetTimings
`func (o *HogQLQueryModifiers) UnsetTimings()`

UnsetTimings ensures that no value is present for Timings, not even an explicit nil
### GetUseMaterializedViews

`func (o *HogQLQueryModifiers) GetUseMaterializedViews() bool`

GetUseMaterializedViews returns the UseMaterializedViews field if non-nil, zero value otherwise.

### GetUseMaterializedViewsOk

`func (o *HogQLQueryModifiers) GetUseMaterializedViewsOk() (*bool, bool)`

GetUseMaterializedViewsOk returns a tuple with the UseMaterializedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaterializedViews

`func (o *HogQLQueryModifiers) SetUseMaterializedViews(v bool)`

SetUseMaterializedViews sets UseMaterializedViews field to given value.

### HasUseMaterializedViews

`func (o *HogQLQueryModifiers) HasUseMaterializedViews() bool`

HasUseMaterializedViews returns a boolean if a field has been set.

### SetUseMaterializedViewsNil

`func (o *HogQLQueryModifiers) SetUseMaterializedViewsNil(b bool)`

 SetUseMaterializedViewsNil sets the value for UseMaterializedViews to be an explicit nil

### UnsetUseMaterializedViews
`func (o *HogQLQueryModifiers) UnsetUseMaterializedViews()`

UnsetUseMaterializedViews ensures that no value is present for UseMaterializedViews, not even an explicit nil
### GetUsePreaggregatedTableTransforms

`func (o *HogQLQueryModifiers) GetUsePreaggregatedTableTransforms() bool`

GetUsePreaggregatedTableTransforms returns the UsePreaggregatedTableTransforms field if non-nil, zero value otherwise.

### GetUsePreaggregatedTableTransformsOk

`func (o *HogQLQueryModifiers) GetUsePreaggregatedTableTransformsOk() (*bool, bool)`

GetUsePreaggregatedTableTransformsOk returns a tuple with the UsePreaggregatedTableTransforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreaggregatedTableTransforms

`func (o *HogQLQueryModifiers) SetUsePreaggregatedTableTransforms(v bool)`

SetUsePreaggregatedTableTransforms sets UsePreaggregatedTableTransforms field to given value.

### HasUsePreaggregatedTableTransforms

`func (o *HogQLQueryModifiers) HasUsePreaggregatedTableTransforms() bool`

HasUsePreaggregatedTableTransforms returns a boolean if a field has been set.

### SetUsePreaggregatedTableTransformsNil

`func (o *HogQLQueryModifiers) SetUsePreaggregatedTableTransformsNil(b bool)`

 SetUsePreaggregatedTableTransformsNil sets the value for UsePreaggregatedTableTransforms to be an explicit nil

### UnsetUsePreaggregatedTableTransforms
`func (o *HogQLQueryModifiers) UnsetUsePreaggregatedTableTransforms()`

UnsetUsePreaggregatedTableTransforms ensures that no value is present for UsePreaggregatedTableTransforms, not even an explicit nil
### GetUsePresortedEventsTable

`func (o *HogQLQueryModifiers) GetUsePresortedEventsTable() bool`

GetUsePresortedEventsTable returns the UsePresortedEventsTable field if non-nil, zero value otherwise.

### GetUsePresortedEventsTableOk

`func (o *HogQLQueryModifiers) GetUsePresortedEventsTableOk() (*bool, bool)`

GetUsePresortedEventsTableOk returns a tuple with the UsePresortedEventsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePresortedEventsTable

`func (o *HogQLQueryModifiers) SetUsePresortedEventsTable(v bool)`

SetUsePresortedEventsTable sets UsePresortedEventsTable field to given value.

### HasUsePresortedEventsTable

`func (o *HogQLQueryModifiers) HasUsePresortedEventsTable() bool`

HasUsePresortedEventsTable returns a boolean if a field has been set.

### SetUsePresortedEventsTableNil

`func (o *HogQLQueryModifiers) SetUsePresortedEventsTableNil(b bool)`

 SetUsePresortedEventsTableNil sets the value for UsePresortedEventsTable to be an explicit nil

### UnsetUsePresortedEventsTable
`func (o *HogQLQueryModifiers) UnsetUsePresortedEventsTable()`

UnsetUsePresortedEventsTable ensures that no value is present for UsePresortedEventsTable, not even an explicit nil
### GetUseWebAnalyticsPreAggregatedTables

`func (o *HogQLQueryModifiers) GetUseWebAnalyticsPreAggregatedTables() bool`

GetUseWebAnalyticsPreAggregatedTables returns the UseWebAnalyticsPreAggregatedTables field if non-nil, zero value otherwise.

### GetUseWebAnalyticsPreAggregatedTablesOk

`func (o *HogQLQueryModifiers) GetUseWebAnalyticsPreAggregatedTablesOk() (*bool, bool)`

GetUseWebAnalyticsPreAggregatedTablesOk returns a tuple with the UseWebAnalyticsPreAggregatedTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWebAnalyticsPreAggregatedTables

`func (o *HogQLQueryModifiers) SetUseWebAnalyticsPreAggregatedTables(v bool)`

SetUseWebAnalyticsPreAggregatedTables sets UseWebAnalyticsPreAggregatedTables field to given value.

### HasUseWebAnalyticsPreAggregatedTables

`func (o *HogQLQueryModifiers) HasUseWebAnalyticsPreAggregatedTables() bool`

HasUseWebAnalyticsPreAggregatedTables returns a boolean if a field has been set.

### SetUseWebAnalyticsPreAggregatedTablesNil

`func (o *HogQLQueryModifiers) SetUseWebAnalyticsPreAggregatedTablesNil(b bool)`

 SetUseWebAnalyticsPreAggregatedTablesNil sets the value for UseWebAnalyticsPreAggregatedTables to be an explicit nil

### UnsetUseWebAnalyticsPreAggregatedTables
`func (o *HogQLQueryModifiers) UnsetUseWebAnalyticsPreAggregatedTables()`

UnsetUseWebAnalyticsPreAggregatedTables ensures that no value is present for UseWebAnalyticsPreAggregatedTables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


