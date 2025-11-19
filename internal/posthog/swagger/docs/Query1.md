# Query1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to **NullableString** |  | [optional] 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner1**](FixedpropertiesInner1.md) | Currently only person filters supported. No filters for querying groups. See &#x60;filter_conditions()&#x60; in actor_strategies.py. | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "UsageMetricsQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Math** | Pointer to **NullableString** |  | [optional] 
**MathGroupTypeIndex** | Pointer to [**MathGroupTypeIndex**](MathGroupTypeIndex.md) |  | [optional] 
**MathHogql** | Pointer to **NullableString** |  | [optional] 
**MathMultiplier** | Pointer to **NullableFloat32** |  | [optional] 
**MathProperty** | Pointer to **NullableString** |  | [optional] 
**MathPropertyRevenueCurrency** | Pointer to [**RevenueCurrencyPropertyConfig**](RevenueCurrencyPropertyConfig.md) |  | [optional] 
**MathPropertyType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OptionalInFunnel** | Pointer to **NullableBool** |  | [optional] 
**OrderBy** | [**OrderBy3**](OrderBy3.md) |  | 
**Properties** | [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Properties configurable in the interface | 
**Response** | Pointer to [**UsageMetricsQueryResponse**](UsageMetricsQueryResponse.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**Id** | **string** |  | 
**Cohort** | Pointer to **NullableInt32** |  | [optional] 
**DistinctId** | Pointer to **NullableString** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**DistinctIdField** | **string** |  | 
**DwSourceType** | Pointer to **NullableString** |  | [optional] 
**IdField** | **string** |  | 
**TableName** | **string** |  | 
**TimestampField** | **string** |  | 
**ActionId** | Pointer to **NullableInt32** |  | [optional] 
**After** | Pointer to **NullableString** | Only fetch sessions that started after this timestamp (default: &#39;-24h&#39;) | [optional] 
**Before** | Pointer to **NullableString** | Only fetch sessions that started before this timestamp (default: &#39;+5s&#39;) | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**PersonId** | Pointer to **NullableString** | Person who performed the event | [optional] 
**Select** | **[]string** | Return a limited set of data. Will use default columns if empty. | 
**Source** | [**FunnelsActorsQuery**](FunnelsActorsQuery.md) |  | 
**Where** | Pointer to **[]string** | HogQL filters to apply on returned data | [optional] 
**GroupTypeIndex** | **NullableInt32** | Group type index. Required with group_key for group queries. | 
**Breakdown** | [**[]RevenueAnalyticsBreakdown**](RevenueAnalyticsBreakdown.md) |  | 
**Compare** | Pointer to [**Compare**](Compare.md) |  | [optional] 
**Day** | Pointer to [**NullableDay**](Day.md) |  | [optional] [default to null]
**IncludeRecordings** | Pointer to **NullableBool** |  | [optional] 
**Interval** | [**IntervalType**](IntervalType.md) |  | 
**Series** | [**[]Series1Inner**](Series1Inner.md) | Events and actions to include | 
**Status** | Pointer to [**Status2**](Status2.md) |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Explain** | Pointer to **NullableBool** |  | [optional] 
**Filters** | Pointer to [**Filters**](Filters.md) |  | [optional] 
**Query** | **map[string]interface{}** |  | 
**Values** | Pointer to **map[string]interface{}** | Constant values that can be referenced with the {placeholder} syntax in the query | [optional] 
**Variables** | Pointer to [**map[string]HogQLVariable**](HogQLVariable.md) | Variables to be substituted into the query | [optional] 
**Debug** | Pointer to **NullableBool** | Enable more verbose output, usually run from the /debug page | [optional] 
**Globals** | Pointer to **map[string]interface{}** | Global values in scope | [optional] 
**Language** | [**HogLanguage**](HogLanguage.md) |  | 
**SourceQuery** | Pointer to [**NullableSourcequery**](Sourcequery.md) |  | [optional] [default to null]
**EndPosition** | **int32** | End position of the editor word | 
**StartPosition** | **int32** | Start position of the editor word | 
**GroupBy** | [**RevenueAnalyticsTopCustomersGroupBy**](RevenueAnalyticsTopCustomersGroupBy.md) |  | 
**Assignee** | Pointer to [**ErrorTrackingIssueAssignee**](ErrorTrackingIssueAssignee.md) |  | [optional] 
**DateRange** | [**DateRange**](DateRange.md) |  | 
**FilterGroup** | [**PropertyGroupFilter**](PropertyGroupFilter.md) |  | 
**GroupKey** | Pointer to **NullableString** |  | [optional] 
**GroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 
**IssueId** | **string** |  | 
**OrderDirection** | Pointer to [**OrderDirection1**](OrderDirection1.md) |  | [optional] 
**RevenueEntity** | Pointer to [**RevenueEntity**](RevenueEntity.md) |  | [optional] 
**RevenuePeriod** | Pointer to [**RevenuePeriod**](RevenuePeriod.md) |  | [optional] 
**SearchQuery** | Pointer to **NullableString** |  | [optional] 
**VolumeResolution** | **int32** |  | 
**WithAggregations** | Pointer to **NullableBool** |  | [optional] 
**WithFirstEvent** | Pointer to **NullableBool** |  | [optional] 
**WithLastEvent** | Pointer to **NullableBool** |  | [optional] 
**MaxDistance** | Pointer to **NullableFloat32** |  | [optional] 
**ModelName** | Pointer to [**EmbeddingModelName**](EmbeddingModelName.md) |  | [optional] 
**Rendering** | Pointer to **NullableString** |  | [optional] 
**BreakdownProperties** | **[]string** |  | 
**MaxValuesPerProperty** | Pointer to **NullableInt32** |  | [optional] 
**Events** | **[]string** |  | 
**ExperimentId** | Pointer to **NullableInt32** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**FunnelsQuery** | [**FunnelsQuery**](FunnelsQuery.md) |  | 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**CountQuery** | [**TrendsQuery**](TrendsQuery.md) |  | 
**ExposureQuery** | Pointer to [**TrendsQuery**](TrendsQuery.md) |  | [optional] 
**Metric** | [**WebVitalsMetric**](WebVitalsMetric.md) |  | 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**ExperimentName** | **string** |  | 
**ExposureCriteria** | Pointer to [**ExperimentExposureCriteria**](ExperimentExposureCriteria.md) |  | [optional] 
**FeatureFlag** | **map[string]interface{}** |  | 
**Holdout** | Pointer to [**ExperimentHoldoutType**](ExperimentHoldoutType.md) |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**DistanceFunc** | [**DistanceFunc**](DistanceFunc.md) |  | 
**DocumentTypes** | **[]string** |  | 
**Model** | **string** |  | 
**OrderBy** | [**OrderBy**](OrderBy.md) |  | 
**OrderDirection** | [**OrderDirection**](OrderDirection.md) |  | 
**Origin** | [**EmbeddedDocument**](EmbeddedDocument.md) |  | 
**Products** | **[]string** |  | 
**Renderings** | **[]string** |  | 
**Threshold** | Pointer to **NullableFloat32** |  | [optional] 
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal**](Conversiongoal.md) |  | [optional] [default to null]
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**BreakdownBy** | [**WebStatsBreakdown**](WebStatsBreakdown.md) |  | 
**IncludeBounceRate** | Pointer to **NullableBool** |  | [optional] 
**IncludeScrollDepth** | Pointer to **NullableBool** |  | [optional] 
**StripQueryParams** | Pointer to **NullableBool** |  | [optional] 
**Percentile** | [**WebVitalsPercentile**](WebVitalsPercentile.md) |  | 
**Thresholds** | **[]float32** |  | 
**SearchTerm** | Pointer to **NullableString** |  | [optional] 
**DraftConversionGoal** | Pointer to [**NullableDraftconversiongoal**](Draftconversiongoal.md) |  | [optional] [default to null]
**IncludeAllConversions** | Pointer to **NullableBool** | Include conversion goal rows even when they don&#39;t match campaign costs table | [optional] 
**IntegrationFilter** | Pointer to [**IntegrationFilter**](IntegrationFilter.md) |  | [optional] 
**ChartSettings** | Pointer to [**ChartSettings**](ChartSettings.md) |  | [optional] 
**Display** | Pointer to [**ChartDisplayType**](ChartDisplayType.md) |  | [optional] 
**TableSettings** | Pointer to [**TableSettings**](TableSettings.md) |  | [optional] 
**AllowSorting** | Pointer to **NullableBool** | Can the user click on column headers to sort the table? (default: true) | [optional] 
**Columns** | Pointer to **[]string** | Columns shown in the table, unless the &#x60;source&#x60; provides them. | [optional] 
**Context** | Pointer to [**DataTableNodeViewPropsContext**](DataTableNodeViewPropsContext.md) |  | [optional] 
**DefaultColumns** | Pointer to **[]string** | Default columns to use when resetting column configuration | [optional] 
**Embedded** | Pointer to **NullableBool** | Query is embedded inside another bordered component | [optional] 
**Expandable** | Pointer to **NullableBool** | Can expand row to show raw event data (default: true) | [optional] 
**Full** | Pointer to **NullableBool** | Show with most visual options enabled. Used in insight scene. | [optional] 
**HiddenColumns** | Pointer to **[]string** | Columns that aren&#39;t shown in the table, even if in columns or returned data | [optional] 
**PinnedColumns** | Pointer to **[]string** | Columns that are sticky when scrolling horizontally | [optional] 
**PropertiesViaUrl** | Pointer to **NullableBool** | Link properties via the URL (default: false) | [optional] 
**ShowActions** | Pointer to **NullableBool** | Show the kebab menu at the end of the row | [optional] 
**ShowColumnConfigurator** | Pointer to **NullableBool** |  | [optional] 
**ShowDateRange** | Pointer to **NullableBool** | Show date range selector | [optional] 
**ShowElapsedTime** | Pointer to **NullableBool** | Show the time it takes to run a query | [optional] 
**ShowEventFilter** | Pointer to **NullableBool** | Include an event filter above the table (EventsNode only) | [optional] 
**ShowExport** | Pointer to **NullableBool** | Show the export button | [optional] 
**ShowHogQLEditor** | Pointer to **NullableBool** | Include a HogQL query editor above HogQL tables | [optional] 
**ShowOpenEditorButton** | Pointer to **NullableBool** | Show a button to open the current query as a new insight. (default: true) | [optional] 
**ShowPersistentColumnConfigurator** | Pointer to **NullableBool** | Show a button to configure and persist the table&#39;s default columns if possible | [optional] 
**ShowPropertyFilter** | Pointer to [**NullableShowpropertyfilter**](Showpropertyfilter.md) |  | [optional] [default to null]
**ShowReload** | Pointer to **NullableBool** | Show a reload button | [optional] 
**ShowResultsTable** | Pointer to **NullableBool** | Show a results table | [optional] 
**ShowSavedFilters** | Pointer to **NullableBool** | Show saved filters feature for this table (requires uniqueKey) | [optional] 
**ShowSavedQueries** | Pointer to **NullableBool** | Shows a list of saved queries | [optional] 
**ShowSearch** | Pointer to **NullableBool** | Include a free text search field (PersonsNode only) | [optional] 
**ShowTestAccountFilters** | Pointer to **NullableBool** | Show filter to exclude test accounts | [optional] 
**ShowTimings** | Pointer to **NullableBool** | Show a detailed query timing breakdown | [optional] 
**HidePersonsModal** | Pointer to **NullableBool** |  | [optional] 
**HideTooltipOnScroll** | Pointer to **NullableBool** |  | [optional] 
**ShortId** | **string** |  | 
**ShowCorrelationTable** | Pointer to **NullableBool** |  | [optional] 
**ShowFilters** | Pointer to **NullableBool** |  | [optional] 
**ShowHeader** | Pointer to **NullableBool** |  | [optional] 
**ShowLastComputation** | Pointer to **NullableBool** |  | [optional] 
**ShowLastComputationRefresh** | Pointer to **NullableBool** |  | [optional] 
**ShowResults** | Pointer to **NullableBool** |  | [optional] 
**ShowTable** | Pointer to **NullableBool** |  | [optional] 
**SuppressSessionAnalysisWarning** | Pointer to **NullableBool** |  | [optional] 
**VizSpecificOptions** | Pointer to [**VizSpecificOptions**](VizSpecificOptions.md) |  | [optional] 
**AggregationGroupTypeIndex** | Pointer to **NullableInt32** | Groups aggregation | [optional] 
**BreakdownFilter** | Pointer to [**BreakdownFilter**](BreakdownFilter.md) |  | [optional] 
**DataColorTheme** | Pointer to **NullableFloat32** | Colors used in the insight&#39;s visualization | [optional] 
**SamplingFactor** | Pointer to **NullableFloat32** | Sampling rate | [optional] 
**TrendsFilter** | Pointer to [**TrendsFilter**](TrendsFilter.md) |  | [optional] 
**FunnelsFilter** | Pointer to [**FunnelsFilter**](FunnelsFilter.md) |  | [optional] 
**RetentionFilter** | [**RetentionFilter**](RetentionFilter.md) |  | 
**FunnelPathsFilter** | Pointer to [**FunnelPathsFilter**](FunnelPathsFilter.md) |  | [optional] 
**PathsFilter** | [**PathsFilter**](PathsFilter.md) |  | 
**IntervalCount** | Pointer to **NullableInt32** | How many intervals comprise a period. Only used for cohorts, otherwise default 1. | [optional] 
**StickinessFilter** | Pointer to [**StickinessFilter**](StickinessFilter.md) |  | [optional] 
**LifecycleFilter** | Pointer to [**LifecycleFilter**](LifecycleFilter.md) |  | [optional] 
**FunnelCorrelationEventExcludePropertyNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationEventNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationExcludeEventNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationExcludeNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationNames** | Pointer to **[]string** |  | [optional] 
**FunnelCorrelationType** | [**FunnelCorrelationResultsType**](FunnelCorrelationResultsType.md) |  | 
**ServiceNames** | **[]string** |  | 
**SeverityLevels** | [**[]LogSeverityLevel**](LogSeverityLevel.md) |  | 
**MaxPropertyValues** | Pointer to **NullableInt32** |  | [optional] 
**TraceId** | **string** |  | 
**Embedding** | **[]float32** |  | 
**EmbeddingVersion** | Pointer to **NullableFloat32** |  | [optional] 
**GroupKey** | Pointer to **NullableString** | Group key. Required with group_type_index for group queries. | [optional] 
**PersonId** | Pointer to **NullableString** | Person ID to fetch metrics for. Mutually exclusive with group parameters. | [optional] 

## Methods

### NewQuery1

`func NewQuery1(orderBy OrderBy3, properties []FixedpropertiesInner, id string, distinctIdField string, idField string, tableName string, timestampField string, select_ []string, source FunnelsActorsQuery, groupTypeIndex NullableInt32, breakdown []RevenueAnalyticsBreakdown, interval IntervalType, series []Series1Inner, query map[string]interface{}, language HogLanguage, endPosition int32, startPosition int32, groupBy RevenueAnalyticsTopCustomersGroupBy, dateRange DateRange, filterGroup PropertyGroupFilter, issueId string, volumeResolution int32, breakdownProperties []string, events []string, funnelsQuery FunnelsQuery, countQuery TrendsQuery, metric WebVitalsMetric, experimentName string, featureFlag map[string]interface{}, distanceFunc DistanceFunc, documentTypes []string, model string, orderBy OrderBy, orderDirection OrderDirection, origin EmbeddedDocument, products []string, renderings []string, breakdownBy WebStatsBreakdown, percentile WebVitalsPercentile, thresholds []float32, shortId string, retentionFilter RetentionFilter, pathsFilter PathsFilter, funnelCorrelationType FunnelCorrelationResultsType, serviceNames []string, severityLevels []LogSeverityLevel, traceId string, embedding []float32, ) *Query1`

NewQuery1 instantiates a new Query1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuery1WithDefaults

`func NewQuery1WithDefaults() *Query1`

NewQuery1WithDefaults instantiates a new Query1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *Query1) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *Query1) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *Query1) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *Query1) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *Query1) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *Query1) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetEvent

`func (o *Query1) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Query1) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Query1) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Query1) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *Query1) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *Query1) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetFixedProperties

`func (o *Query1) GetFixedProperties() []FixedpropertiesInner1`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *Query1) GetFixedPropertiesOk() (*[]FixedpropertiesInner1, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *Query1) SetFixedProperties(v []FixedpropertiesInner1)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *Query1) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *Query1) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *Query1) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *Query1) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Query1) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Query1) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Query1) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *Query1) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Query1) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Query1) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Query1) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Query1) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Query1) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMath

`func (o *Query1) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *Query1) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *Query1) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *Query1) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *Query1) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *Query1) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *Query1) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *Query1) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *Query1) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *Query1) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *Query1) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *Query1) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *Query1) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *Query1) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *Query1) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *Query1) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *Query1) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *Query1) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *Query1) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *Query1) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *Query1) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *Query1) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *Query1) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *Query1) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *Query1) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *Query1) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *Query1) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *Query1) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *Query1) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *Query1) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *Query1) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *Query1) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *Query1) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *Query1) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *Query1) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *Query1) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *Query1) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *Query1) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *Query1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Query1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Query1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Query1) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Query1) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Query1) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *Query1) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *Query1) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *Query1) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *Query1) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *Query1) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *Query1) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetOrderBy

`func (o *Query1) GetOrderBy() OrderBy3`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *Query1) GetOrderByOk() (*OrderBy3, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *Query1) SetOrderBy(v OrderBy3)`

SetOrderBy sets OrderBy field to given value.


### GetProperties

`func (o *Query1) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Query1) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Query1) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.


### SetPropertiesNil

`func (o *Query1) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Query1) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *Query1) GetResponse() UsageMetricsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Query1) GetResponseOk() (*UsageMetricsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Query1) SetResponse(v UsageMetricsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Query1) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetVersion

`func (o *Query1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Query1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Query1) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Query1) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Query1) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Query1) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *Query1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Query1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Query1) SetId(v string)`

SetId sets Id field to given value.


### GetCohort

`func (o *Query1) GetCohort() int32`

GetCohort returns the Cohort field if non-nil, zero value otherwise.

### GetCohortOk

`func (o *Query1) GetCohortOk() (*int32, bool)`

GetCohortOk returns a tuple with the Cohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohort

`func (o *Query1) SetCohort(v int32)`

SetCohort sets Cohort field to given value.

### HasCohort

`func (o *Query1) HasCohort() bool`

HasCohort returns a boolean if a field has been set.

### SetCohortNil

`func (o *Query1) SetCohortNil(b bool)`

 SetCohortNil sets the value for Cohort to be an explicit nil

### UnsetCohort
`func (o *Query1) UnsetCohort()`

UnsetCohort ensures that no value is present for Cohort, not even an explicit nil
### GetDistinctId

`func (o *Query1) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *Query1) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *Query1) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *Query1) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### SetDistinctIdNil

`func (o *Query1) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *Query1) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetModifiers

`func (o *Query1) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Query1) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Query1) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Query1) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *Query1) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Query1) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Query1) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Query1) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *Query1) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *Query1) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSearch

`func (o *Query1) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *Query1) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *Query1) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *Query1) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *Query1) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *Query1) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetTags

`func (o *Query1) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Query1) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Query1) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Query1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDistinctIdField

`func (o *Query1) GetDistinctIdField() string`

GetDistinctIdField returns the DistinctIdField field if non-nil, zero value otherwise.

### GetDistinctIdFieldOk

`func (o *Query1) GetDistinctIdFieldOk() (*string, bool)`

GetDistinctIdFieldOk returns a tuple with the DistinctIdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIdField

`func (o *Query1) SetDistinctIdField(v string)`

SetDistinctIdField sets DistinctIdField field to given value.


### GetDwSourceType

`func (o *Query1) GetDwSourceType() string`

GetDwSourceType returns the DwSourceType field if non-nil, zero value otherwise.

### GetDwSourceTypeOk

`func (o *Query1) GetDwSourceTypeOk() (*string, bool)`

GetDwSourceTypeOk returns a tuple with the DwSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwSourceType

`func (o *Query1) SetDwSourceType(v string)`

SetDwSourceType sets DwSourceType field to given value.

### HasDwSourceType

`func (o *Query1) HasDwSourceType() bool`

HasDwSourceType returns a boolean if a field has been set.

### SetDwSourceTypeNil

`func (o *Query1) SetDwSourceTypeNil(b bool)`

 SetDwSourceTypeNil sets the value for DwSourceType to be an explicit nil

### UnsetDwSourceType
`func (o *Query1) UnsetDwSourceType()`

UnsetDwSourceType ensures that no value is present for DwSourceType, not even an explicit nil
### GetIdField

`func (o *Query1) GetIdField() string`

GetIdField returns the IdField field if non-nil, zero value otherwise.

### GetIdFieldOk

`func (o *Query1) GetIdFieldOk() (*string, bool)`

GetIdFieldOk returns a tuple with the IdField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdField

`func (o *Query1) SetIdField(v string)`

SetIdField sets IdField field to given value.


### GetTableName

`func (o *Query1) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *Query1) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *Query1) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetTimestampField

`func (o *Query1) GetTimestampField() string`

GetTimestampField returns the TimestampField field if non-nil, zero value otherwise.

### GetTimestampFieldOk

`func (o *Query1) GetTimestampFieldOk() (*string, bool)`

GetTimestampFieldOk returns a tuple with the TimestampField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampField

`func (o *Query1) SetTimestampField(v string)`

SetTimestampField sets TimestampField field to given value.


### GetActionId

`func (o *Query1) GetActionId() int32`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *Query1) GetActionIdOk() (*int32, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *Query1) SetActionId(v int32)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *Query1) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### SetActionIdNil

`func (o *Query1) SetActionIdNil(b bool)`

 SetActionIdNil sets the value for ActionId to be an explicit nil

### UnsetActionId
`func (o *Query1) UnsetActionId()`

UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
### GetAfter

`func (o *Query1) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Query1) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Query1) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *Query1) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *Query1) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *Query1) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *Query1) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Query1) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Query1) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *Query1) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *Query1) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *Query1) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetFilterTestAccounts

`func (o *Query1) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *Query1) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *Query1) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *Query1) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *Query1) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *Query1) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetPersonId

`func (o *Query1) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Query1) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Query1) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Query1) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *Query1) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *Query1) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetSelect

`func (o *Query1) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *Query1) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *Query1) SetSelect(v []string)`

SetSelect sets Select field to given value.


### SetSelectNil

`func (o *Query1) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *Query1) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetSource

`func (o *Query1) GetSource() FunnelsActorsQuery`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Query1) GetSourceOk() (*FunnelsActorsQuery, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Query1) SetSource(v FunnelsActorsQuery)`

SetSource sets Source field to given value.


### GetWhere

`func (o *Query1) GetWhere() []string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *Query1) GetWhereOk() (*[]string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *Query1) SetWhere(v []string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *Query1) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### SetWhereNil

`func (o *Query1) SetWhereNil(b bool)`

 SetWhereNil sets the value for Where to be an explicit nil

### UnsetWhere
`func (o *Query1) UnsetWhere()`

UnsetWhere ensures that no value is present for Where, not even an explicit nil
### GetGroupTypeIndex

`func (o *Query1) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Query1) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Query1) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.


### SetGroupTypeIndexNil

`func (o *Query1) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Query1) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetBreakdown

`func (o *Query1) GetBreakdown() []RevenueAnalyticsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Query1) GetBreakdownOk() (*[]RevenueAnalyticsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Query1) SetBreakdown(v []RevenueAnalyticsBreakdown)`

SetBreakdown sets Breakdown field to given value.


### GetCompare

`func (o *Query1) GetCompare() Compare`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *Query1) GetCompareOk() (*Compare, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *Query1) SetCompare(v Compare)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *Query1) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### GetDay

`func (o *Query1) GetDay() Day`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Query1) GetDayOk() (*Day, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Query1) SetDay(v Day)`

SetDay sets Day field to given value.

### HasDay

`func (o *Query1) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *Query1) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *Query1) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetIncludeRecordings

`func (o *Query1) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *Query1) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *Query1) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *Query1) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *Query1) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *Query1) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetInterval

`func (o *Query1) GetInterval() IntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Query1) GetIntervalOk() (*IntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Query1) SetInterval(v IntervalType)`

SetInterval sets Interval field to given value.


### GetSeries

`func (o *Query1) GetSeries() []Series1Inner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Query1) GetSeriesOk() (*[]Series1Inner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Query1) SetSeries(v []Series1Inner)`

SetSeries sets Series field to given value.


### GetStatus

`func (o *Query1) GetStatus() Status2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Query1) GetStatusOk() (*Status2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Query1) SetStatus(v Status2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Query1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *Query1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Query1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Query1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Query1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Query1) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Query1) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetExplain

`func (o *Query1) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *Query1) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *Query1) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *Query1) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *Query1) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *Query1) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetFilters

`func (o *Query1) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Query1) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Query1) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Query1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetQuery

`func (o *Query1) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Query1) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Query1) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.


### GetValues

`func (o *Query1) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Query1) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Query1) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *Query1) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *Query1) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *Query1) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetVariables

`func (o *Query1) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Query1) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Query1) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Query1) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *Query1) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Query1) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetDebug

`func (o *Query1) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *Query1) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *Query1) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *Query1) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### SetDebugNil

`func (o *Query1) SetDebugNil(b bool)`

 SetDebugNil sets the value for Debug to be an explicit nil

### UnsetDebug
`func (o *Query1) UnsetDebug()`

UnsetDebug ensures that no value is present for Debug, not even an explicit nil
### GetGlobals

`func (o *Query1) GetGlobals() map[string]interface{}`

GetGlobals returns the Globals field if non-nil, zero value otherwise.

### GetGlobalsOk

`func (o *Query1) GetGlobalsOk() (*map[string]interface{}, bool)`

GetGlobalsOk returns a tuple with the Globals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobals

`func (o *Query1) SetGlobals(v map[string]interface{})`

SetGlobals sets Globals field to given value.

### HasGlobals

`func (o *Query1) HasGlobals() bool`

HasGlobals returns a boolean if a field has been set.

### SetGlobalsNil

`func (o *Query1) SetGlobalsNil(b bool)`

 SetGlobalsNil sets the value for Globals to be an explicit nil

### UnsetGlobals
`func (o *Query1) UnsetGlobals()`

UnsetGlobals ensures that no value is present for Globals, not even an explicit nil
### GetLanguage

`func (o *Query1) GetLanguage() HogLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Query1) GetLanguageOk() (*HogLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Query1) SetLanguage(v HogLanguage)`

SetLanguage sets Language field to given value.


### GetSourceQuery

`func (o *Query1) GetSourceQuery() Sourcequery`

GetSourceQuery returns the SourceQuery field if non-nil, zero value otherwise.

### GetSourceQueryOk

`func (o *Query1) GetSourceQueryOk() (*Sourcequery, bool)`

GetSourceQueryOk returns a tuple with the SourceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuery

`func (o *Query1) SetSourceQuery(v Sourcequery)`

SetSourceQuery sets SourceQuery field to given value.

### HasSourceQuery

`func (o *Query1) HasSourceQuery() bool`

HasSourceQuery returns a boolean if a field has been set.

### SetSourceQueryNil

`func (o *Query1) SetSourceQueryNil(b bool)`

 SetSourceQueryNil sets the value for SourceQuery to be an explicit nil

### UnsetSourceQuery
`func (o *Query1) UnsetSourceQuery()`

UnsetSourceQuery ensures that no value is present for SourceQuery, not even an explicit nil
### GetEndPosition

`func (o *Query1) GetEndPosition() int32`

GetEndPosition returns the EndPosition field if non-nil, zero value otherwise.

### GetEndPositionOk

`func (o *Query1) GetEndPositionOk() (*int32, bool)`

GetEndPositionOk returns a tuple with the EndPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPosition

`func (o *Query1) SetEndPosition(v int32)`

SetEndPosition sets EndPosition field to given value.


### GetStartPosition

`func (o *Query1) GetStartPosition() int32`

GetStartPosition returns the StartPosition field if non-nil, zero value otherwise.

### GetStartPositionOk

`func (o *Query1) GetStartPositionOk() (*int32, bool)`

GetStartPositionOk returns a tuple with the StartPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPosition

`func (o *Query1) SetStartPosition(v int32)`

SetStartPosition sets StartPosition field to given value.


### GetGroupBy

`func (o *Query1) GetGroupBy() RevenueAnalyticsTopCustomersGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Query1) GetGroupByOk() (*RevenueAnalyticsTopCustomersGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Query1) SetGroupBy(v RevenueAnalyticsTopCustomersGroupBy)`

SetGroupBy sets GroupBy field to given value.


### GetAssignee

`func (o *Query1) GetAssignee() ErrorTrackingIssueAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Query1) GetAssigneeOk() (*ErrorTrackingIssueAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Query1) SetAssignee(v ErrorTrackingIssueAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *Query1) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetDateRange

`func (o *Query1) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Query1) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Query1) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetFilterGroup

`func (o *Query1) GetFilterGroup() PropertyGroupFilter`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *Query1) GetFilterGroupOk() (*PropertyGroupFilter, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *Query1) SetFilterGroup(v PropertyGroupFilter)`

SetFilterGroup sets FilterGroup field to given value.


### GetGroupKey

`func (o *Query1) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *Query1) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *Query1) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *Query1) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### SetGroupKeyNil

`func (o *Query1) SetGroupKeyNil(b bool)`

 SetGroupKeyNil sets the value for GroupKey to be an explicit nil

### UnsetGroupKey
`func (o *Query1) UnsetGroupKey()`

UnsetGroupKey ensures that no value is present for GroupKey, not even an explicit nil
### GetGroupTypeIndex

`func (o *Query1) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Query1) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Query1) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *Query1) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *Query1) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Query1) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetIssueId

`func (o *Query1) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Query1) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Query1) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.


### GetOrderDirection

`func (o *Query1) GetOrderDirection() OrderDirection1`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *Query1) GetOrderDirectionOk() (*OrderDirection1, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *Query1) SetOrderDirection(v OrderDirection1)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *Query1) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetRevenueEntity

`func (o *Query1) GetRevenueEntity() RevenueEntity`

GetRevenueEntity returns the RevenueEntity field if non-nil, zero value otherwise.

### GetRevenueEntityOk

`func (o *Query1) GetRevenueEntityOk() (*RevenueEntity, bool)`

GetRevenueEntityOk returns a tuple with the RevenueEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueEntity

`func (o *Query1) SetRevenueEntity(v RevenueEntity)`

SetRevenueEntity sets RevenueEntity field to given value.

### HasRevenueEntity

`func (o *Query1) HasRevenueEntity() bool`

HasRevenueEntity returns a boolean if a field has been set.

### GetRevenuePeriod

`func (o *Query1) GetRevenuePeriod() RevenuePeriod`

GetRevenuePeriod returns the RevenuePeriod field if non-nil, zero value otherwise.

### GetRevenuePeriodOk

`func (o *Query1) GetRevenuePeriodOk() (*RevenuePeriod, bool)`

GetRevenuePeriodOk returns a tuple with the RevenuePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePeriod

`func (o *Query1) SetRevenuePeriod(v RevenuePeriod)`

SetRevenuePeriod sets RevenuePeriod field to given value.

### HasRevenuePeriod

`func (o *Query1) HasRevenuePeriod() bool`

HasRevenuePeriod returns a boolean if a field has been set.

### GetSearchQuery

`func (o *Query1) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *Query1) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *Query1) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *Query1) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### SetSearchQueryNil

`func (o *Query1) SetSearchQueryNil(b bool)`

 SetSearchQueryNil sets the value for SearchQuery to be an explicit nil

### UnsetSearchQuery
`func (o *Query1) UnsetSearchQuery()`

UnsetSearchQuery ensures that no value is present for SearchQuery, not even an explicit nil
### GetVolumeResolution

`func (o *Query1) GetVolumeResolution() int32`

GetVolumeResolution returns the VolumeResolution field if non-nil, zero value otherwise.

### GetVolumeResolutionOk

`func (o *Query1) GetVolumeResolutionOk() (*int32, bool)`

GetVolumeResolutionOk returns a tuple with the VolumeResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeResolution

`func (o *Query1) SetVolumeResolution(v int32)`

SetVolumeResolution sets VolumeResolution field to given value.


### GetWithAggregations

`func (o *Query1) GetWithAggregations() bool`

GetWithAggregations returns the WithAggregations field if non-nil, zero value otherwise.

### GetWithAggregationsOk

`func (o *Query1) GetWithAggregationsOk() (*bool, bool)`

GetWithAggregationsOk returns a tuple with the WithAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAggregations

`func (o *Query1) SetWithAggregations(v bool)`

SetWithAggregations sets WithAggregations field to given value.

### HasWithAggregations

`func (o *Query1) HasWithAggregations() bool`

HasWithAggregations returns a boolean if a field has been set.

### SetWithAggregationsNil

`func (o *Query1) SetWithAggregationsNil(b bool)`

 SetWithAggregationsNil sets the value for WithAggregations to be an explicit nil

### UnsetWithAggregations
`func (o *Query1) UnsetWithAggregations()`

UnsetWithAggregations ensures that no value is present for WithAggregations, not even an explicit nil
### GetWithFirstEvent

`func (o *Query1) GetWithFirstEvent() bool`

GetWithFirstEvent returns the WithFirstEvent field if non-nil, zero value otherwise.

### GetWithFirstEventOk

`func (o *Query1) GetWithFirstEventOk() (*bool, bool)`

GetWithFirstEventOk returns a tuple with the WithFirstEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithFirstEvent

`func (o *Query1) SetWithFirstEvent(v bool)`

SetWithFirstEvent sets WithFirstEvent field to given value.

### HasWithFirstEvent

`func (o *Query1) HasWithFirstEvent() bool`

HasWithFirstEvent returns a boolean if a field has been set.

### SetWithFirstEventNil

`func (o *Query1) SetWithFirstEventNil(b bool)`

 SetWithFirstEventNil sets the value for WithFirstEvent to be an explicit nil

### UnsetWithFirstEvent
`func (o *Query1) UnsetWithFirstEvent()`

UnsetWithFirstEvent ensures that no value is present for WithFirstEvent, not even an explicit nil
### GetWithLastEvent

`func (o *Query1) GetWithLastEvent() bool`

GetWithLastEvent returns the WithLastEvent field if non-nil, zero value otherwise.

### GetWithLastEventOk

`func (o *Query1) GetWithLastEventOk() (*bool, bool)`

GetWithLastEventOk returns a tuple with the WithLastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithLastEvent

`func (o *Query1) SetWithLastEvent(v bool)`

SetWithLastEvent sets WithLastEvent field to given value.

### HasWithLastEvent

`func (o *Query1) HasWithLastEvent() bool`

HasWithLastEvent returns a boolean if a field has been set.

### SetWithLastEventNil

`func (o *Query1) SetWithLastEventNil(b bool)`

 SetWithLastEventNil sets the value for WithLastEvent to be an explicit nil

### UnsetWithLastEvent
`func (o *Query1) UnsetWithLastEvent()`

UnsetWithLastEvent ensures that no value is present for WithLastEvent, not even an explicit nil
### GetMaxDistance

`func (o *Query1) GetMaxDistance() float32`

GetMaxDistance returns the MaxDistance field if non-nil, zero value otherwise.

### GetMaxDistanceOk

`func (o *Query1) GetMaxDistanceOk() (*float32, bool)`

GetMaxDistanceOk returns a tuple with the MaxDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistance

`func (o *Query1) SetMaxDistance(v float32)`

SetMaxDistance sets MaxDistance field to given value.

### HasMaxDistance

`func (o *Query1) HasMaxDistance() bool`

HasMaxDistance returns a boolean if a field has been set.

### SetMaxDistanceNil

`func (o *Query1) SetMaxDistanceNil(b bool)`

 SetMaxDistanceNil sets the value for MaxDistance to be an explicit nil

### UnsetMaxDistance
`func (o *Query1) UnsetMaxDistance()`

UnsetMaxDistance ensures that no value is present for MaxDistance, not even an explicit nil
### GetModelName

`func (o *Query1) GetModelName() EmbeddingModelName`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Query1) GetModelNameOk() (*EmbeddingModelName, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Query1) SetModelName(v EmbeddingModelName)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *Query1) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetRendering

`func (o *Query1) GetRendering() string`

GetRendering returns the Rendering field if non-nil, zero value otherwise.

### GetRenderingOk

`func (o *Query1) GetRenderingOk() (*string, bool)`

GetRenderingOk returns a tuple with the Rendering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendering

`func (o *Query1) SetRendering(v string)`

SetRendering sets Rendering field to given value.

### HasRendering

`func (o *Query1) HasRendering() bool`

HasRendering returns a boolean if a field has been set.

### SetRenderingNil

`func (o *Query1) SetRenderingNil(b bool)`

 SetRenderingNil sets the value for Rendering to be an explicit nil

### UnsetRendering
`func (o *Query1) UnsetRendering()`

UnsetRendering ensures that no value is present for Rendering, not even an explicit nil
### GetBreakdownProperties

`func (o *Query1) GetBreakdownProperties() []string`

GetBreakdownProperties returns the BreakdownProperties field if non-nil, zero value otherwise.

### GetBreakdownPropertiesOk

`func (o *Query1) GetBreakdownPropertiesOk() (*[]string, bool)`

GetBreakdownPropertiesOk returns a tuple with the BreakdownProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownProperties

`func (o *Query1) SetBreakdownProperties(v []string)`

SetBreakdownProperties sets BreakdownProperties field to given value.


### GetMaxValuesPerProperty

`func (o *Query1) GetMaxValuesPerProperty() int32`

GetMaxValuesPerProperty returns the MaxValuesPerProperty field if non-nil, zero value otherwise.

### GetMaxValuesPerPropertyOk

`func (o *Query1) GetMaxValuesPerPropertyOk() (*int32, bool)`

GetMaxValuesPerPropertyOk returns a tuple with the MaxValuesPerProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValuesPerProperty

`func (o *Query1) SetMaxValuesPerProperty(v int32)`

SetMaxValuesPerProperty sets MaxValuesPerProperty field to given value.

### HasMaxValuesPerProperty

`func (o *Query1) HasMaxValuesPerProperty() bool`

HasMaxValuesPerProperty returns a boolean if a field has been set.

### SetMaxValuesPerPropertyNil

`func (o *Query1) SetMaxValuesPerPropertyNil(b bool)`

 SetMaxValuesPerPropertyNil sets the value for MaxValuesPerProperty to be an explicit nil

### UnsetMaxValuesPerProperty
`func (o *Query1) UnsetMaxValuesPerProperty()`

UnsetMaxValuesPerProperty ensures that no value is present for MaxValuesPerProperty, not even an explicit nil
### GetEvents

`func (o *Query1) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Query1) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Query1) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetExperimentId

`func (o *Query1) GetExperimentId() int32`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *Query1) GetExperimentIdOk() (*int32, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *Query1) SetExperimentId(v int32)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *Query1) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### SetExperimentIdNil

`func (o *Query1) SetExperimentIdNil(b bool)`

 SetExperimentIdNil sets the value for ExperimentId to be an explicit nil

### UnsetExperimentId
`func (o *Query1) UnsetExperimentId()`

UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil
### GetFingerprint

`func (o *Query1) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Query1) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Query1) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Query1) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *Query1) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Query1) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetFunnelsQuery

`func (o *Query1) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *Query1) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *Query1) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.


### GetUuid

`func (o *Query1) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Query1) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Query1) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Query1) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Query1) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Query1) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetCountQuery

`func (o *Query1) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *Query1) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *Query1) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.


### GetExposureQuery

`func (o *Query1) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *Query1) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *Query1) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *Query1) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetMetric

`func (o *Query1) GetMetric() WebVitalsMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Query1) GetMetricOk() (*WebVitalsMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Query1) SetMetric(v WebVitalsMetric)`

SetMetric sets Metric field to given value.


### GetEndDate

`func (o *Query1) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Query1) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Query1) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Query1) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Query1) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Query1) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetExperimentName

`func (o *Query1) GetExperimentName() string`

GetExperimentName returns the ExperimentName field if non-nil, zero value otherwise.

### GetExperimentNameOk

`func (o *Query1) GetExperimentNameOk() (*string, bool)`

GetExperimentNameOk returns a tuple with the ExperimentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentName

`func (o *Query1) SetExperimentName(v string)`

SetExperimentName sets ExperimentName field to given value.


### GetExposureCriteria

`func (o *Query1) GetExposureCriteria() ExperimentExposureCriteria`

GetExposureCriteria returns the ExposureCriteria field if non-nil, zero value otherwise.

### GetExposureCriteriaOk

`func (o *Query1) GetExposureCriteriaOk() (*ExperimentExposureCriteria, bool)`

GetExposureCriteriaOk returns a tuple with the ExposureCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureCriteria

`func (o *Query1) SetExposureCriteria(v ExperimentExposureCriteria)`

SetExposureCriteria sets ExposureCriteria field to given value.

### HasExposureCriteria

`func (o *Query1) HasExposureCriteria() bool`

HasExposureCriteria returns a boolean if a field has been set.

### GetFeatureFlag

`func (o *Query1) GetFeatureFlag() map[string]interface{}`

GetFeatureFlag returns the FeatureFlag field if non-nil, zero value otherwise.

### GetFeatureFlagOk

`func (o *Query1) GetFeatureFlagOk() (*map[string]interface{}, bool)`

GetFeatureFlagOk returns a tuple with the FeatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlag

`func (o *Query1) SetFeatureFlag(v map[string]interface{})`

SetFeatureFlag sets FeatureFlag field to given value.


### GetHoldout

`func (o *Query1) GetHoldout() ExperimentHoldoutType`

GetHoldout returns the Holdout field if non-nil, zero value otherwise.

### GetHoldoutOk

`func (o *Query1) GetHoldoutOk() (*ExperimentHoldoutType, bool)`

GetHoldoutOk returns a tuple with the Holdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldout

`func (o *Query1) SetHoldout(v ExperimentHoldoutType)`

SetHoldout sets Holdout field to given value.

### HasHoldout

`func (o *Query1) HasHoldout() bool`

HasHoldout returns a boolean if a field has been set.

### GetStartDate

`func (o *Query1) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Query1) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Query1) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Query1) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Query1) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Query1) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetDistanceFunc

`func (o *Query1) GetDistanceFunc() DistanceFunc`

GetDistanceFunc returns the DistanceFunc field if non-nil, zero value otherwise.

### GetDistanceFuncOk

`func (o *Query1) GetDistanceFuncOk() (*DistanceFunc, bool)`

GetDistanceFuncOk returns a tuple with the DistanceFunc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceFunc

`func (o *Query1) SetDistanceFunc(v DistanceFunc)`

SetDistanceFunc sets DistanceFunc field to given value.


### GetDocumentTypes

`func (o *Query1) GetDocumentTypes() []string`

GetDocumentTypes returns the DocumentTypes field if non-nil, zero value otherwise.

### GetDocumentTypesOk

`func (o *Query1) GetDocumentTypesOk() (*[]string, bool)`

GetDocumentTypesOk returns a tuple with the DocumentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypes

`func (o *Query1) SetDocumentTypes(v []string)`

SetDocumentTypes sets DocumentTypes field to given value.


### GetModel

`func (o *Query1) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Query1) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Query1) SetModel(v string)`

SetModel sets Model field to given value.


### GetOrderBy

`func (o *Query1) GetOrderBy() OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *Query1) GetOrderByOk() (*OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *Query1) SetOrderBy(v OrderBy)`

SetOrderBy sets OrderBy field to given value.


### GetOrderDirection

`func (o *Query1) GetOrderDirection() OrderDirection`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *Query1) GetOrderDirectionOk() (*OrderDirection, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *Query1) SetOrderDirection(v OrderDirection)`

SetOrderDirection sets OrderDirection field to given value.


### GetOrigin

`func (o *Query1) GetOrigin() EmbeddedDocument`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Query1) GetOriginOk() (*EmbeddedDocument, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Query1) SetOrigin(v EmbeddedDocument)`

SetOrigin sets Origin field to given value.


### GetProducts

`func (o *Query1) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Query1) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Query1) SetProducts(v []string)`

SetProducts sets Products field to given value.


### GetRenderings

`func (o *Query1) GetRenderings() []string`

GetRenderings returns the Renderings field if non-nil, zero value otherwise.

### GetRenderingsOk

`func (o *Query1) GetRenderingsOk() (*[]string, bool)`

GetRenderingsOk returns a tuple with the Renderings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderings

`func (o *Query1) SetRenderings(v []string)`

SetRenderings sets Renderings field to given value.


### GetThreshold

`func (o *Query1) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Query1) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Query1) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *Query1) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### SetThresholdNil

`func (o *Query1) SetThresholdNil(b bool)`

 SetThresholdNil sets the value for Threshold to be an explicit nil

### UnsetThreshold
`func (o *Query1) UnsetThreshold()`

UnsetThreshold ensures that no value is present for Threshold, not even an explicit nil
### GetCompareFilter

`func (o *Query1) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *Query1) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *Query1) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *Query1) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *Query1) GetConversionGoal() Conversiongoal`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *Query1) GetConversionGoalOk() (*Conversiongoal, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *Query1) SetConversionGoal(v Conversiongoal)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *Query1) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *Query1) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *Query1) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDoPathCleaning

`func (o *Query1) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *Query1) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *Query1) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *Query1) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *Query1) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *Query1) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetIncludeRevenue

`func (o *Query1) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *Query1) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *Query1) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *Query1) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *Query1) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *Query1) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetSampling

`func (o *Query1) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *Query1) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *Query1) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *Query1) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *Query1) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *Query1) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *Query1) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *Query1) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *Query1) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *Query1) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetBreakdownBy

`func (o *Query1) GetBreakdownBy() WebStatsBreakdown`

GetBreakdownBy returns the BreakdownBy field if non-nil, zero value otherwise.

### GetBreakdownByOk

`func (o *Query1) GetBreakdownByOk() (*WebStatsBreakdown, bool)`

GetBreakdownByOk returns a tuple with the BreakdownBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownBy

`func (o *Query1) SetBreakdownBy(v WebStatsBreakdown)`

SetBreakdownBy sets BreakdownBy field to given value.


### GetIncludeBounceRate

`func (o *Query1) GetIncludeBounceRate() bool`

GetIncludeBounceRate returns the IncludeBounceRate field if non-nil, zero value otherwise.

### GetIncludeBounceRateOk

`func (o *Query1) GetIncludeBounceRateOk() (*bool, bool)`

GetIncludeBounceRateOk returns a tuple with the IncludeBounceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBounceRate

`func (o *Query1) SetIncludeBounceRate(v bool)`

SetIncludeBounceRate sets IncludeBounceRate field to given value.

### HasIncludeBounceRate

`func (o *Query1) HasIncludeBounceRate() bool`

HasIncludeBounceRate returns a boolean if a field has been set.

### SetIncludeBounceRateNil

`func (o *Query1) SetIncludeBounceRateNil(b bool)`

 SetIncludeBounceRateNil sets the value for IncludeBounceRate to be an explicit nil

### UnsetIncludeBounceRate
`func (o *Query1) UnsetIncludeBounceRate()`

UnsetIncludeBounceRate ensures that no value is present for IncludeBounceRate, not even an explicit nil
### GetIncludeScrollDepth

`func (o *Query1) GetIncludeScrollDepth() bool`

GetIncludeScrollDepth returns the IncludeScrollDepth field if non-nil, zero value otherwise.

### GetIncludeScrollDepthOk

`func (o *Query1) GetIncludeScrollDepthOk() (*bool, bool)`

GetIncludeScrollDepthOk returns a tuple with the IncludeScrollDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeScrollDepth

`func (o *Query1) SetIncludeScrollDepth(v bool)`

SetIncludeScrollDepth sets IncludeScrollDepth field to given value.

### HasIncludeScrollDepth

`func (o *Query1) HasIncludeScrollDepth() bool`

HasIncludeScrollDepth returns a boolean if a field has been set.

### SetIncludeScrollDepthNil

`func (o *Query1) SetIncludeScrollDepthNil(b bool)`

 SetIncludeScrollDepthNil sets the value for IncludeScrollDepth to be an explicit nil

### UnsetIncludeScrollDepth
`func (o *Query1) UnsetIncludeScrollDepth()`

UnsetIncludeScrollDepth ensures that no value is present for IncludeScrollDepth, not even an explicit nil
### GetStripQueryParams

`func (o *Query1) GetStripQueryParams() bool`

GetStripQueryParams returns the StripQueryParams field if non-nil, zero value otherwise.

### GetStripQueryParamsOk

`func (o *Query1) GetStripQueryParamsOk() (*bool, bool)`

GetStripQueryParamsOk returns a tuple with the StripQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripQueryParams

`func (o *Query1) SetStripQueryParams(v bool)`

SetStripQueryParams sets StripQueryParams field to given value.

### HasStripQueryParams

`func (o *Query1) HasStripQueryParams() bool`

HasStripQueryParams returns a boolean if a field has been set.

### SetStripQueryParamsNil

`func (o *Query1) SetStripQueryParamsNil(b bool)`

 SetStripQueryParamsNil sets the value for StripQueryParams to be an explicit nil

### UnsetStripQueryParams
`func (o *Query1) UnsetStripQueryParams()`

UnsetStripQueryParams ensures that no value is present for StripQueryParams, not even an explicit nil
### GetPercentile

`func (o *Query1) GetPercentile() WebVitalsPercentile`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *Query1) GetPercentileOk() (*WebVitalsPercentile, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *Query1) SetPercentile(v WebVitalsPercentile)`

SetPercentile sets Percentile field to given value.


### GetThresholds

`func (o *Query1) GetThresholds() []float32`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Query1) GetThresholdsOk() (*[]float32, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Query1) SetThresholds(v []float32)`

SetThresholds sets Thresholds field to given value.


### GetSearchTerm

`func (o *Query1) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *Query1) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *Query1) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *Query1) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### SetSearchTermNil

`func (o *Query1) SetSearchTermNil(b bool)`

 SetSearchTermNil sets the value for SearchTerm to be an explicit nil

### UnsetSearchTerm
`func (o *Query1) UnsetSearchTerm()`

UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil
### GetDraftConversionGoal

`func (o *Query1) GetDraftConversionGoal() Draftconversiongoal`

GetDraftConversionGoal returns the DraftConversionGoal field if non-nil, zero value otherwise.

### GetDraftConversionGoalOk

`func (o *Query1) GetDraftConversionGoalOk() (*Draftconversiongoal, bool)`

GetDraftConversionGoalOk returns a tuple with the DraftConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftConversionGoal

`func (o *Query1) SetDraftConversionGoal(v Draftconversiongoal)`

SetDraftConversionGoal sets DraftConversionGoal field to given value.

### HasDraftConversionGoal

`func (o *Query1) HasDraftConversionGoal() bool`

HasDraftConversionGoal returns a boolean if a field has been set.

### SetDraftConversionGoalNil

`func (o *Query1) SetDraftConversionGoalNil(b bool)`

 SetDraftConversionGoalNil sets the value for DraftConversionGoal to be an explicit nil

### UnsetDraftConversionGoal
`func (o *Query1) UnsetDraftConversionGoal()`

UnsetDraftConversionGoal ensures that no value is present for DraftConversionGoal, not even an explicit nil
### GetIncludeAllConversions

`func (o *Query1) GetIncludeAllConversions() bool`

GetIncludeAllConversions returns the IncludeAllConversions field if non-nil, zero value otherwise.

### GetIncludeAllConversionsOk

`func (o *Query1) GetIncludeAllConversionsOk() (*bool, bool)`

GetIncludeAllConversionsOk returns a tuple with the IncludeAllConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllConversions

`func (o *Query1) SetIncludeAllConversions(v bool)`

SetIncludeAllConversions sets IncludeAllConversions field to given value.

### HasIncludeAllConversions

`func (o *Query1) HasIncludeAllConversions() bool`

HasIncludeAllConversions returns a boolean if a field has been set.

### SetIncludeAllConversionsNil

`func (o *Query1) SetIncludeAllConversionsNil(b bool)`

 SetIncludeAllConversionsNil sets the value for IncludeAllConversions to be an explicit nil

### UnsetIncludeAllConversions
`func (o *Query1) UnsetIncludeAllConversions()`

UnsetIncludeAllConversions ensures that no value is present for IncludeAllConversions, not even an explicit nil
### GetIntegrationFilter

`func (o *Query1) GetIntegrationFilter() IntegrationFilter`

GetIntegrationFilter returns the IntegrationFilter field if non-nil, zero value otherwise.

### GetIntegrationFilterOk

`func (o *Query1) GetIntegrationFilterOk() (*IntegrationFilter, bool)`

GetIntegrationFilterOk returns a tuple with the IntegrationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationFilter

`func (o *Query1) SetIntegrationFilter(v IntegrationFilter)`

SetIntegrationFilter sets IntegrationFilter field to given value.

### HasIntegrationFilter

`func (o *Query1) HasIntegrationFilter() bool`

HasIntegrationFilter returns a boolean if a field has been set.

### GetChartSettings

`func (o *Query1) GetChartSettings() ChartSettings`

GetChartSettings returns the ChartSettings field if non-nil, zero value otherwise.

### GetChartSettingsOk

`func (o *Query1) GetChartSettingsOk() (*ChartSettings, bool)`

GetChartSettingsOk returns a tuple with the ChartSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartSettings

`func (o *Query1) SetChartSettings(v ChartSettings)`

SetChartSettings sets ChartSettings field to given value.

### HasChartSettings

`func (o *Query1) HasChartSettings() bool`

HasChartSettings returns a boolean if a field has been set.

### GetDisplay

`func (o *Query1) GetDisplay() ChartDisplayType`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Query1) GetDisplayOk() (*ChartDisplayType, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Query1) SetDisplay(v ChartDisplayType)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Query1) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetTableSettings

`func (o *Query1) GetTableSettings() TableSettings`

GetTableSettings returns the TableSettings field if non-nil, zero value otherwise.

### GetTableSettingsOk

`func (o *Query1) GetTableSettingsOk() (*TableSettings, bool)`

GetTableSettingsOk returns a tuple with the TableSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableSettings

`func (o *Query1) SetTableSettings(v TableSettings)`

SetTableSettings sets TableSettings field to given value.

### HasTableSettings

`func (o *Query1) HasTableSettings() bool`

HasTableSettings returns a boolean if a field has been set.

### GetAllowSorting

`func (o *Query1) GetAllowSorting() bool`

GetAllowSorting returns the AllowSorting field if non-nil, zero value otherwise.

### GetAllowSortingOk

`func (o *Query1) GetAllowSortingOk() (*bool, bool)`

GetAllowSortingOk returns a tuple with the AllowSorting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSorting

`func (o *Query1) SetAllowSorting(v bool)`

SetAllowSorting sets AllowSorting field to given value.

### HasAllowSorting

`func (o *Query1) HasAllowSorting() bool`

HasAllowSorting returns a boolean if a field has been set.

### SetAllowSortingNil

`func (o *Query1) SetAllowSortingNil(b bool)`

 SetAllowSortingNil sets the value for AllowSorting to be an explicit nil

### UnsetAllowSorting
`func (o *Query1) UnsetAllowSorting()`

UnsetAllowSorting ensures that no value is present for AllowSorting, not even an explicit nil
### GetColumns

`func (o *Query1) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Query1) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Query1) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Query1) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *Query1) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *Query1) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetContext

`func (o *Query1) GetContext() DataTableNodeViewPropsContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Query1) GetContextOk() (*DataTableNodeViewPropsContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Query1) SetContext(v DataTableNodeViewPropsContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *Query1) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDefaultColumns

`func (o *Query1) GetDefaultColumns() []string`

GetDefaultColumns returns the DefaultColumns field if non-nil, zero value otherwise.

### GetDefaultColumnsOk

`func (o *Query1) GetDefaultColumnsOk() (*[]string, bool)`

GetDefaultColumnsOk returns a tuple with the DefaultColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColumns

`func (o *Query1) SetDefaultColumns(v []string)`

SetDefaultColumns sets DefaultColumns field to given value.

### HasDefaultColumns

`func (o *Query1) HasDefaultColumns() bool`

HasDefaultColumns returns a boolean if a field has been set.

### SetDefaultColumnsNil

`func (o *Query1) SetDefaultColumnsNil(b bool)`

 SetDefaultColumnsNil sets the value for DefaultColumns to be an explicit nil

### UnsetDefaultColumns
`func (o *Query1) UnsetDefaultColumns()`

UnsetDefaultColumns ensures that no value is present for DefaultColumns, not even an explicit nil
### GetEmbedded

`func (o *Query1) GetEmbedded() bool`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Query1) GetEmbeddedOk() (*bool, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Query1) SetEmbedded(v bool)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *Query1) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### SetEmbeddedNil

`func (o *Query1) SetEmbeddedNil(b bool)`

 SetEmbeddedNil sets the value for Embedded to be an explicit nil

### UnsetEmbedded
`func (o *Query1) UnsetEmbedded()`

UnsetEmbedded ensures that no value is present for Embedded, not even an explicit nil
### GetExpandable

`func (o *Query1) GetExpandable() bool`

GetExpandable returns the Expandable field if non-nil, zero value otherwise.

### GetExpandableOk

`func (o *Query1) GetExpandableOk() (*bool, bool)`

GetExpandableOk returns a tuple with the Expandable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandable

`func (o *Query1) SetExpandable(v bool)`

SetExpandable sets Expandable field to given value.

### HasExpandable

`func (o *Query1) HasExpandable() bool`

HasExpandable returns a boolean if a field has been set.

### SetExpandableNil

`func (o *Query1) SetExpandableNil(b bool)`

 SetExpandableNil sets the value for Expandable to be an explicit nil

### UnsetExpandable
`func (o *Query1) UnsetExpandable()`

UnsetExpandable ensures that no value is present for Expandable, not even an explicit nil
### GetFull

`func (o *Query1) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *Query1) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *Query1) SetFull(v bool)`

SetFull sets Full field to given value.

### HasFull

`func (o *Query1) HasFull() bool`

HasFull returns a boolean if a field has been set.

### SetFullNil

`func (o *Query1) SetFullNil(b bool)`

 SetFullNil sets the value for Full to be an explicit nil

### UnsetFull
`func (o *Query1) UnsetFull()`

UnsetFull ensures that no value is present for Full, not even an explicit nil
### GetHiddenColumns

`func (o *Query1) GetHiddenColumns() []string`

GetHiddenColumns returns the HiddenColumns field if non-nil, zero value otherwise.

### GetHiddenColumnsOk

`func (o *Query1) GetHiddenColumnsOk() (*[]string, bool)`

GetHiddenColumnsOk returns a tuple with the HiddenColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenColumns

`func (o *Query1) SetHiddenColumns(v []string)`

SetHiddenColumns sets HiddenColumns field to given value.

### HasHiddenColumns

`func (o *Query1) HasHiddenColumns() bool`

HasHiddenColumns returns a boolean if a field has been set.

### SetHiddenColumnsNil

`func (o *Query1) SetHiddenColumnsNil(b bool)`

 SetHiddenColumnsNil sets the value for HiddenColumns to be an explicit nil

### UnsetHiddenColumns
`func (o *Query1) UnsetHiddenColumns()`

UnsetHiddenColumns ensures that no value is present for HiddenColumns, not even an explicit nil
### GetPinnedColumns

`func (o *Query1) GetPinnedColumns() []string`

GetPinnedColumns returns the PinnedColumns field if non-nil, zero value otherwise.

### GetPinnedColumnsOk

`func (o *Query1) GetPinnedColumnsOk() (*[]string, bool)`

GetPinnedColumnsOk returns a tuple with the PinnedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedColumns

`func (o *Query1) SetPinnedColumns(v []string)`

SetPinnedColumns sets PinnedColumns field to given value.

### HasPinnedColumns

`func (o *Query1) HasPinnedColumns() bool`

HasPinnedColumns returns a boolean if a field has been set.

### SetPinnedColumnsNil

`func (o *Query1) SetPinnedColumnsNil(b bool)`

 SetPinnedColumnsNil sets the value for PinnedColumns to be an explicit nil

### UnsetPinnedColumns
`func (o *Query1) UnsetPinnedColumns()`

UnsetPinnedColumns ensures that no value is present for PinnedColumns, not even an explicit nil
### GetPropertiesViaUrl

`func (o *Query1) GetPropertiesViaUrl() bool`

GetPropertiesViaUrl returns the PropertiesViaUrl field if non-nil, zero value otherwise.

### GetPropertiesViaUrlOk

`func (o *Query1) GetPropertiesViaUrlOk() (*bool, bool)`

GetPropertiesViaUrlOk returns a tuple with the PropertiesViaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesViaUrl

`func (o *Query1) SetPropertiesViaUrl(v bool)`

SetPropertiesViaUrl sets PropertiesViaUrl field to given value.

### HasPropertiesViaUrl

`func (o *Query1) HasPropertiesViaUrl() bool`

HasPropertiesViaUrl returns a boolean if a field has been set.

### SetPropertiesViaUrlNil

`func (o *Query1) SetPropertiesViaUrlNil(b bool)`

 SetPropertiesViaUrlNil sets the value for PropertiesViaUrl to be an explicit nil

### UnsetPropertiesViaUrl
`func (o *Query1) UnsetPropertiesViaUrl()`

UnsetPropertiesViaUrl ensures that no value is present for PropertiesViaUrl, not even an explicit nil
### GetShowActions

`func (o *Query1) GetShowActions() bool`

GetShowActions returns the ShowActions field if non-nil, zero value otherwise.

### GetShowActionsOk

`func (o *Query1) GetShowActionsOk() (*bool, bool)`

GetShowActionsOk returns a tuple with the ShowActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowActions

`func (o *Query1) SetShowActions(v bool)`

SetShowActions sets ShowActions field to given value.

### HasShowActions

`func (o *Query1) HasShowActions() bool`

HasShowActions returns a boolean if a field has been set.

### SetShowActionsNil

`func (o *Query1) SetShowActionsNil(b bool)`

 SetShowActionsNil sets the value for ShowActions to be an explicit nil

### UnsetShowActions
`func (o *Query1) UnsetShowActions()`

UnsetShowActions ensures that no value is present for ShowActions, not even an explicit nil
### GetShowColumnConfigurator

`func (o *Query1) GetShowColumnConfigurator() bool`

GetShowColumnConfigurator returns the ShowColumnConfigurator field if non-nil, zero value otherwise.

### GetShowColumnConfiguratorOk

`func (o *Query1) GetShowColumnConfiguratorOk() (*bool, bool)`

GetShowColumnConfiguratorOk returns a tuple with the ShowColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowColumnConfigurator

`func (o *Query1) SetShowColumnConfigurator(v bool)`

SetShowColumnConfigurator sets ShowColumnConfigurator field to given value.

### HasShowColumnConfigurator

`func (o *Query1) HasShowColumnConfigurator() bool`

HasShowColumnConfigurator returns a boolean if a field has been set.

### SetShowColumnConfiguratorNil

`func (o *Query1) SetShowColumnConfiguratorNil(b bool)`

 SetShowColumnConfiguratorNil sets the value for ShowColumnConfigurator to be an explicit nil

### UnsetShowColumnConfigurator
`func (o *Query1) UnsetShowColumnConfigurator()`

UnsetShowColumnConfigurator ensures that no value is present for ShowColumnConfigurator, not even an explicit nil
### GetShowDateRange

`func (o *Query1) GetShowDateRange() bool`

GetShowDateRange returns the ShowDateRange field if non-nil, zero value otherwise.

### GetShowDateRangeOk

`func (o *Query1) GetShowDateRangeOk() (*bool, bool)`

GetShowDateRangeOk returns a tuple with the ShowDateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDateRange

`func (o *Query1) SetShowDateRange(v bool)`

SetShowDateRange sets ShowDateRange field to given value.

### HasShowDateRange

`func (o *Query1) HasShowDateRange() bool`

HasShowDateRange returns a boolean if a field has been set.

### SetShowDateRangeNil

`func (o *Query1) SetShowDateRangeNil(b bool)`

 SetShowDateRangeNil sets the value for ShowDateRange to be an explicit nil

### UnsetShowDateRange
`func (o *Query1) UnsetShowDateRange()`

UnsetShowDateRange ensures that no value is present for ShowDateRange, not even an explicit nil
### GetShowElapsedTime

`func (o *Query1) GetShowElapsedTime() bool`

GetShowElapsedTime returns the ShowElapsedTime field if non-nil, zero value otherwise.

### GetShowElapsedTimeOk

`func (o *Query1) GetShowElapsedTimeOk() (*bool, bool)`

GetShowElapsedTimeOk returns a tuple with the ShowElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowElapsedTime

`func (o *Query1) SetShowElapsedTime(v bool)`

SetShowElapsedTime sets ShowElapsedTime field to given value.

### HasShowElapsedTime

`func (o *Query1) HasShowElapsedTime() bool`

HasShowElapsedTime returns a boolean if a field has been set.

### SetShowElapsedTimeNil

`func (o *Query1) SetShowElapsedTimeNil(b bool)`

 SetShowElapsedTimeNil sets the value for ShowElapsedTime to be an explicit nil

### UnsetShowElapsedTime
`func (o *Query1) UnsetShowElapsedTime()`

UnsetShowElapsedTime ensures that no value is present for ShowElapsedTime, not even an explicit nil
### GetShowEventFilter

`func (o *Query1) GetShowEventFilter() bool`

GetShowEventFilter returns the ShowEventFilter field if non-nil, zero value otherwise.

### GetShowEventFilterOk

`func (o *Query1) GetShowEventFilterOk() (*bool, bool)`

GetShowEventFilterOk returns a tuple with the ShowEventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEventFilter

`func (o *Query1) SetShowEventFilter(v bool)`

SetShowEventFilter sets ShowEventFilter field to given value.

### HasShowEventFilter

`func (o *Query1) HasShowEventFilter() bool`

HasShowEventFilter returns a boolean if a field has been set.

### SetShowEventFilterNil

`func (o *Query1) SetShowEventFilterNil(b bool)`

 SetShowEventFilterNil sets the value for ShowEventFilter to be an explicit nil

### UnsetShowEventFilter
`func (o *Query1) UnsetShowEventFilter()`

UnsetShowEventFilter ensures that no value is present for ShowEventFilter, not even an explicit nil
### GetShowExport

`func (o *Query1) GetShowExport() bool`

GetShowExport returns the ShowExport field if non-nil, zero value otherwise.

### GetShowExportOk

`func (o *Query1) GetShowExportOk() (*bool, bool)`

GetShowExportOk returns a tuple with the ShowExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowExport

`func (o *Query1) SetShowExport(v bool)`

SetShowExport sets ShowExport field to given value.

### HasShowExport

`func (o *Query1) HasShowExport() bool`

HasShowExport returns a boolean if a field has been set.

### SetShowExportNil

`func (o *Query1) SetShowExportNil(b bool)`

 SetShowExportNil sets the value for ShowExport to be an explicit nil

### UnsetShowExport
`func (o *Query1) UnsetShowExport()`

UnsetShowExport ensures that no value is present for ShowExport, not even an explicit nil
### GetShowHogQLEditor

`func (o *Query1) GetShowHogQLEditor() bool`

GetShowHogQLEditor returns the ShowHogQLEditor field if non-nil, zero value otherwise.

### GetShowHogQLEditorOk

`func (o *Query1) GetShowHogQLEditorOk() (*bool, bool)`

GetShowHogQLEditorOk returns a tuple with the ShowHogQLEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHogQLEditor

`func (o *Query1) SetShowHogQLEditor(v bool)`

SetShowHogQLEditor sets ShowHogQLEditor field to given value.

### HasShowHogQLEditor

`func (o *Query1) HasShowHogQLEditor() bool`

HasShowHogQLEditor returns a boolean if a field has been set.

### SetShowHogQLEditorNil

`func (o *Query1) SetShowHogQLEditorNil(b bool)`

 SetShowHogQLEditorNil sets the value for ShowHogQLEditor to be an explicit nil

### UnsetShowHogQLEditor
`func (o *Query1) UnsetShowHogQLEditor()`

UnsetShowHogQLEditor ensures that no value is present for ShowHogQLEditor, not even an explicit nil
### GetShowOpenEditorButton

`func (o *Query1) GetShowOpenEditorButton() bool`

GetShowOpenEditorButton returns the ShowOpenEditorButton field if non-nil, zero value otherwise.

### GetShowOpenEditorButtonOk

`func (o *Query1) GetShowOpenEditorButtonOk() (*bool, bool)`

GetShowOpenEditorButtonOk returns a tuple with the ShowOpenEditorButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOpenEditorButton

`func (o *Query1) SetShowOpenEditorButton(v bool)`

SetShowOpenEditorButton sets ShowOpenEditorButton field to given value.

### HasShowOpenEditorButton

`func (o *Query1) HasShowOpenEditorButton() bool`

HasShowOpenEditorButton returns a boolean if a field has been set.

### SetShowOpenEditorButtonNil

`func (o *Query1) SetShowOpenEditorButtonNil(b bool)`

 SetShowOpenEditorButtonNil sets the value for ShowOpenEditorButton to be an explicit nil

### UnsetShowOpenEditorButton
`func (o *Query1) UnsetShowOpenEditorButton()`

UnsetShowOpenEditorButton ensures that no value is present for ShowOpenEditorButton, not even an explicit nil
### GetShowPersistentColumnConfigurator

`func (o *Query1) GetShowPersistentColumnConfigurator() bool`

GetShowPersistentColumnConfigurator returns the ShowPersistentColumnConfigurator field if non-nil, zero value otherwise.

### GetShowPersistentColumnConfiguratorOk

`func (o *Query1) GetShowPersistentColumnConfiguratorOk() (*bool, bool)`

GetShowPersistentColumnConfiguratorOk returns a tuple with the ShowPersistentColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPersistentColumnConfigurator

`func (o *Query1) SetShowPersistentColumnConfigurator(v bool)`

SetShowPersistentColumnConfigurator sets ShowPersistentColumnConfigurator field to given value.

### HasShowPersistentColumnConfigurator

`func (o *Query1) HasShowPersistentColumnConfigurator() bool`

HasShowPersistentColumnConfigurator returns a boolean if a field has been set.

### SetShowPersistentColumnConfiguratorNil

`func (o *Query1) SetShowPersistentColumnConfiguratorNil(b bool)`

 SetShowPersistentColumnConfiguratorNil sets the value for ShowPersistentColumnConfigurator to be an explicit nil

### UnsetShowPersistentColumnConfigurator
`func (o *Query1) UnsetShowPersistentColumnConfigurator()`

UnsetShowPersistentColumnConfigurator ensures that no value is present for ShowPersistentColumnConfigurator, not even an explicit nil
### GetShowPropertyFilter

`func (o *Query1) GetShowPropertyFilter() Showpropertyfilter`

GetShowPropertyFilter returns the ShowPropertyFilter field if non-nil, zero value otherwise.

### GetShowPropertyFilterOk

`func (o *Query1) GetShowPropertyFilterOk() (*Showpropertyfilter, bool)`

GetShowPropertyFilterOk returns a tuple with the ShowPropertyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPropertyFilter

`func (o *Query1) SetShowPropertyFilter(v Showpropertyfilter)`

SetShowPropertyFilter sets ShowPropertyFilter field to given value.

### HasShowPropertyFilter

`func (o *Query1) HasShowPropertyFilter() bool`

HasShowPropertyFilter returns a boolean if a field has been set.

### SetShowPropertyFilterNil

`func (o *Query1) SetShowPropertyFilterNil(b bool)`

 SetShowPropertyFilterNil sets the value for ShowPropertyFilter to be an explicit nil

### UnsetShowPropertyFilter
`func (o *Query1) UnsetShowPropertyFilter()`

UnsetShowPropertyFilter ensures that no value is present for ShowPropertyFilter, not even an explicit nil
### GetShowReload

`func (o *Query1) GetShowReload() bool`

GetShowReload returns the ShowReload field if non-nil, zero value otherwise.

### GetShowReloadOk

`func (o *Query1) GetShowReloadOk() (*bool, bool)`

GetShowReloadOk returns a tuple with the ShowReload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowReload

`func (o *Query1) SetShowReload(v bool)`

SetShowReload sets ShowReload field to given value.

### HasShowReload

`func (o *Query1) HasShowReload() bool`

HasShowReload returns a boolean if a field has been set.

### SetShowReloadNil

`func (o *Query1) SetShowReloadNil(b bool)`

 SetShowReloadNil sets the value for ShowReload to be an explicit nil

### UnsetShowReload
`func (o *Query1) UnsetShowReload()`

UnsetShowReload ensures that no value is present for ShowReload, not even an explicit nil
### GetShowResultsTable

`func (o *Query1) GetShowResultsTable() bool`

GetShowResultsTable returns the ShowResultsTable field if non-nil, zero value otherwise.

### GetShowResultsTableOk

`func (o *Query1) GetShowResultsTableOk() (*bool, bool)`

GetShowResultsTableOk returns a tuple with the ShowResultsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResultsTable

`func (o *Query1) SetShowResultsTable(v bool)`

SetShowResultsTable sets ShowResultsTable field to given value.

### HasShowResultsTable

`func (o *Query1) HasShowResultsTable() bool`

HasShowResultsTable returns a boolean if a field has been set.

### SetShowResultsTableNil

`func (o *Query1) SetShowResultsTableNil(b bool)`

 SetShowResultsTableNil sets the value for ShowResultsTable to be an explicit nil

### UnsetShowResultsTable
`func (o *Query1) UnsetShowResultsTable()`

UnsetShowResultsTable ensures that no value is present for ShowResultsTable, not even an explicit nil
### GetShowSavedFilters

`func (o *Query1) GetShowSavedFilters() bool`

GetShowSavedFilters returns the ShowSavedFilters field if non-nil, zero value otherwise.

### GetShowSavedFiltersOk

`func (o *Query1) GetShowSavedFiltersOk() (*bool, bool)`

GetShowSavedFiltersOk returns a tuple with the ShowSavedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSavedFilters

`func (o *Query1) SetShowSavedFilters(v bool)`

SetShowSavedFilters sets ShowSavedFilters field to given value.

### HasShowSavedFilters

`func (o *Query1) HasShowSavedFilters() bool`

HasShowSavedFilters returns a boolean if a field has been set.

### SetShowSavedFiltersNil

`func (o *Query1) SetShowSavedFiltersNil(b bool)`

 SetShowSavedFiltersNil sets the value for ShowSavedFilters to be an explicit nil

### UnsetShowSavedFilters
`func (o *Query1) UnsetShowSavedFilters()`

UnsetShowSavedFilters ensures that no value is present for ShowSavedFilters, not even an explicit nil
### GetShowSavedQueries

`func (o *Query1) GetShowSavedQueries() bool`

GetShowSavedQueries returns the ShowSavedQueries field if non-nil, zero value otherwise.

### GetShowSavedQueriesOk

`func (o *Query1) GetShowSavedQueriesOk() (*bool, bool)`

GetShowSavedQueriesOk returns a tuple with the ShowSavedQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSavedQueries

`func (o *Query1) SetShowSavedQueries(v bool)`

SetShowSavedQueries sets ShowSavedQueries field to given value.

### HasShowSavedQueries

`func (o *Query1) HasShowSavedQueries() bool`

HasShowSavedQueries returns a boolean if a field has been set.

### SetShowSavedQueriesNil

`func (o *Query1) SetShowSavedQueriesNil(b bool)`

 SetShowSavedQueriesNil sets the value for ShowSavedQueries to be an explicit nil

### UnsetShowSavedQueries
`func (o *Query1) UnsetShowSavedQueries()`

UnsetShowSavedQueries ensures that no value is present for ShowSavedQueries, not even an explicit nil
### GetShowSearch

`func (o *Query1) GetShowSearch() bool`

GetShowSearch returns the ShowSearch field if non-nil, zero value otherwise.

### GetShowSearchOk

`func (o *Query1) GetShowSearchOk() (*bool, bool)`

GetShowSearchOk returns a tuple with the ShowSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSearch

`func (o *Query1) SetShowSearch(v bool)`

SetShowSearch sets ShowSearch field to given value.

### HasShowSearch

`func (o *Query1) HasShowSearch() bool`

HasShowSearch returns a boolean if a field has been set.

### SetShowSearchNil

`func (o *Query1) SetShowSearchNil(b bool)`

 SetShowSearchNil sets the value for ShowSearch to be an explicit nil

### UnsetShowSearch
`func (o *Query1) UnsetShowSearch()`

UnsetShowSearch ensures that no value is present for ShowSearch, not even an explicit nil
### GetShowTestAccountFilters

`func (o *Query1) GetShowTestAccountFilters() bool`

GetShowTestAccountFilters returns the ShowTestAccountFilters field if non-nil, zero value otherwise.

### GetShowTestAccountFiltersOk

`func (o *Query1) GetShowTestAccountFiltersOk() (*bool, bool)`

GetShowTestAccountFiltersOk returns a tuple with the ShowTestAccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTestAccountFilters

`func (o *Query1) SetShowTestAccountFilters(v bool)`

SetShowTestAccountFilters sets ShowTestAccountFilters field to given value.

### HasShowTestAccountFilters

`func (o *Query1) HasShowTestAccountFilters() bool`

HasShowTestAccountFilters returns a boolean if a field has been set.

### SetShowTestAccountFiltersNil

`func (o *Query1) SetShowTestAccountFiltersNil(b bool)`

 SetShowTestAccountFiltersNil sets the value for ShowTestAccountFilters to be an explicit nil

### UnsetShowTestAccountFilters
`func (o *Query1) UnsetShowTestAccountFilters()`

UnsetShowTestAccountFilters ensures that no value is present for ShowTestAccountFilters, not even an explicit nil
### GetShowTimings

`func (o *Query1) GetShowTimings() bool`

GetShowTimings returns the ShowTimings field if non-nil, zero value otherwise.

### GetShowTimingsOk

`func (o *Query1) GetShowTimingsOk() (*bool, bool)`

GetShowTimingsOk returns a tuple with the ShowTimings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTimings

`func (o *Query1) SetShowTimings(v bool)`

SetShowTimings sets ShowTimings field to given value.

### HasShowTimings

`func (o *Query1) HasShowTimings() bool`

HasShowTimings returns a boolean if a field has been set.

### SetShowTimingsNil

`func (o *Query1) SetShowTimingsNil(b bool)`

 SetShowTimingsNil sets the value for ShowTimings to be an explicit nil

### UnsetShowTimings
`func (o *Query1) UnsetShowTimings()`

UnsetShowTimings ensures that no value is present for ShowTimings, not even an explicit nil
### GetHidePersonsModal

`func (o *Query1) GetHidePersonsModal() bool`

GetHidePersonsModal returns the HidePersonsModal field if non-nil, zero value otherwise.

### GetHidePersonsModalOk

`func (o *Query1) GetHidePersonsModalOk() (*bool, bool)`

GetHidePersonsModalOk returns a tuple with the HidePersonsModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePersonsModal

`func (o *Query1) SetHidePersonsModal(v bool)`

SetHidePersonsModal sets HidePersonsModal field to given value.

### HasHidePersonsModal

`func (o *Query1) HasHidePersonsModal() bool`

HasHidePersonsModal returns a boolean if a field has been set.

### SetHidePersonsModalNil

`func (o *Query1) SetHidePersonsModalNil(b bool)`

 SetHidePersonsModalNil sets the value for HidePersonsModal to be an explicit nil

### UnsetHidePersonsModal
`func (o *Query1) UnsetHidePersonsModal()`

UnsetHidePersonsModal ensures that no value is present for HidePersonsModal, not even an explicit nil
### GetHideTooltipOnScroll

`func (o *Query1) GetHideTooltipOnScroll() bool`

GetHideTooltipOnScroll returns the HideTooltipOnScroll field if non-nil, zero value otherwise.

### GetHideTooltipOnScrollOk

`func (o *Query1) GetHideTooltipOnScrollOk() (*bool, bool)`

GetHideTooltipOnScrollOk returns a tuple with the HideTooltipOnScroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTooltipOnScroll

`func (o *Query1) SetHideTooltipOnScroll(v bool)`

SetHideTooltipOnScroll sets HideTooltipOnScroll field to given value.

### HasHideTooltipOnScroll

`func (o *Query1) HasHideTooltipOnScroll() bool`

HasHideTooltipOnScroll returns a boolean if a field has been set.

### SetHideTooltipOnScrollNil

`func (o *Query1) SetHideTooltipOnScrollNil(b bool)`

 SetHideTooltipOnScrollNil sets the value for HideTooltipOnScroll to be an explicit nil

### UnsetHideTooltipOnScroll
`func (o *Query1) UnsetHideTooltipOnScroll()`

UnsetHideTooltipOnScroll ensures that no value is present for HideTooltipOnScroll, not even an explicit nil
### GetShortId

`func (o *Query1) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *Query1) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *Query1) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetShowCorrelationTable

`func (o *Query1) GetShowCorrelationTable() bool`

GetShowCorrelationTable returns the ShowCorrelationTable field if non-nil, zero value otherwise.

### GetShowCorrelationTableOk

`func (o *Query1) GetShowCorrelationTableOk() (*bool, bool)`

GetShowCorrelationTableOk returns a tuple with the ShowCorrelationTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCorrelationTable

`func (o *Query1) SetShowCorrelationTable(v bool)`

SetShowCorrelationTable sets ShowCorrelationTable field to given value.

### HasShowCorrelationTable

`func (o *Query1) HasShowCorrelationTable() bool`

HasShowCorrelationTable returns a boolean if a field has been set.

### SetShowCorrelationTableNil

`func (o *Query1) SetShowCorrelationTableNil(b bool)`

 SetShowCorrelationTableNil sets the value for ShowCorrelationTable to be an explicit nil

### UnsetShowCorrelationTable
`func (o *Query1) UnsetShowCorrelationTable()`

UnsetShowCorrelationTable ensures that no value is present for ShowCorrelationTable, not even an explicit nil
### GetShowFilters

`func (o *Query1) GetShowFilters() bool`

GetShowFilters returns the ShowFilters field if non-nil, zero value otherwise.

### GetShowFiltersOk

`func (o *Query1) GetShowFiltersOk() (*bool, bool)`

GetShowFiltersOk returns a tuple with the ShowFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowFilters

`func (o *Query1) SetShowFilters(v bool)`

SetShowFilters sets ShowFilters field to given value.

### HasShowFilters

`func (o *Query1) HasShowFilters() bool`

HasShowFilters returns a boolean if a field has been set.

### SetShowFiltersNil

`func (o *Query1) SetShowFiltersNil(b bool)`

 SetShowFiltersNil sets the value for ShowFilters to be an explicit nil

### UnsetShowFilters
`func (o *Query1) UnsetShowFilters()`

UnsetShowFilters ensures that no value is present for ShowFilters, not even an explicit nil
### GetShowHeader

`func (o *Query1) GetShowHeader() bool`

GetShowHeader returns the ShowHeader field if non-nil, zero value otherwise.

### GetShowHeaderOk

`func (o *Query1) GetShowHeaderOk() (*bool, bool)`

GetShowHeaderOk returns a tuple with the ShowHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHeader

`func (o *Query1) SetShowHeader(v bool)`

SetShowHeader sets ShowHeader field to given value.

### HasShowHeader

`func (o *Query1) HasShowHeader() bool`

HasShowHeader returns a boolean if a field has been set.

### SetShowHeaderNil

`func (o *Query1) SetShowHeaderNil(b bool)`

 SetShowHeaderNil sets the value for ShowHeader to be an explicit nil

### UnsetShowHeader
`func (o *Query1) UnsetShowHeader()`

UnsetShowHeader ensures that no value is present for ShowHeader, not even an explicit nil
### GetShowLastComputation

`func (o *Query1) GetShowLastComputation() bool`

GetShowLastComputation returns the ShowLastComputation field if non-nil, zero value otherwise.

### GetShowLastComputationOk

`func (o *Query1) GetShowLastComputationOk() (*bool, bool)`

GetShowLastComputationOk returns a tuple with the ShowLastComputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastComputation

`func (o *Query1) SetShowLastComputation(v bool)`

SetShowLastComputation sets ShowLastComputation field to given value.

### HasShowLastComputation

`func (o *Query1) HasShowLastComputation() bool`

HasShowLastComputation returns a boolean if a field has been set.

### SetShowLastComputationNil

`func (o *Query1) SetShowLastComputationNil(b bool)`

 SetShowLastComputationNil sets the value for ShowLastComputation to be an explicit nil

### UnsetShowLastComputation
`func (o *Query1) UnsetShowLastComputation()`

UnsetShowLastComputation ensures that no value is present for ShowLastComputation, not even an explicit nil
### GetShowLastComputationRefresh

`func (o *Query1) GetShowLastComputationRefresh() bool`

GetShowLastComputationRefresh returns the ShowLastComputationRefresh field if non-nil, zero value otherwise.

### GetShowLastComputationRefreshOk

`func (o *Query1) GetShowLastComputationRefreshOk() (*bool, bool)`

GetShowLastComputationRefreshOk returns a tuple with the ShowLastComputationRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLastComputationRefresh

`func (o *Query1) SetShowLastComputationRefresh(v bool)`

SetShowLastComputationRefresh sets ShowLastComputationRefresh field to given value.

### HasShowLastComputationRefresh

`func (o *Query1) HasShowLastComputationRefresh() bool`

HasShowLastComputationRefresh returns a boolean if a field has been set.

### SetShowLastComputationRefreshNil

`func (o *Query1) SetShowLastComputationRefreshNil(b bool)`

 SetShowLastComputationRefreshNil sets the value for ShowLastComputationRefresh to be an explicit nil

### UnsetShowLastComputationRefresh
`func (o *Query1) UnsetShowLastComputationRefresh()`

UnsetShowLastComputationRefresh ensures that no value is present for ShowLastComputationRefresh, not even an explicit nil
### GetShowResults

`func (o *Query1) GetShowResults() bool`

GetShowResults returns the ShowResults field if non-nil, zero value otherwise.

### GetShowResultsOk

`func (o *Query1) GetShowResultsOk() (*bool, bool)`

GetShowResultsOk returns a tuple with the ShowResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResults

`func (o *Query1) SetShowResults(v bool)`

SetShowResults sets ShowResults field to given value.

### HasShowResults

`func (o *Query1) HasShowResults() bool`

HasShowResults returns a boolean if a field has been set.

### SetShowResultsNil

`func (o *Query1) SetShowResultsNil(b bool)`

 SetShowResultsNil sets the value for ShowResults to be an explicit nil

### UnsetShowResults
`func (o *Query1) UnsetShowResults()`

UnsetShowResults ensures that no value is present for ShowResults, not even an explicit nil
### GetShowTable

`func (o *Query1) GetShowTable() bool`

GetShowTable returns the ShowTable field if non-nil, zero value otherwise.

### GetShowTableOk

`func (o *Query1) GetShowTableOk() (*bool, bool)`

GetShowTableOk returns a tuple with the ShowTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTable

`func (o *Query1) SetShowTable(v bool)`

SetShowTable sets ShowTable field to given value.

### HasShowTable

`func (o *Query1) HasShowTable() bool`

HasShowTable returns a boolean if a field has been set.

### SetShowTableNil

`func (o *Query1) SetShowTableNil(b bool)`

 SetShowTableNil sets the value for ShowTable to be an explicit nil

### UnsetShowTable
`func (o *Query1) UnsetShowTable()`

UnsetShowTable ensures that no value is present for ShowTable, not even an explicit nil
### GetSuppressSessionAnalysisWarning

`func (o *Query1) GetSuppressSessionAnalysisWarning() bool`

GetSuppressSessionAnalysisWarning returns the SuppressSessionAnalysisWarning field if non-nil, zero value otherwise.

### GetSuppressSessionAnalysisWarningOk

`func (o *Query1) GetSuppressSessionAnalysisWarningOk() (*bool, bool)`

GetSuppressSessionAnalysisWarningOk returns a tuple with the SuppressSessionAnalysisWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressSessionAnalysisWarning

`func (o *Query1) SetSuppressSessionAnalysisWarning(v bool)`

SetSuppressSessionAnalysisWarning sets SuppressSessionAnalysisWarning field to given value.

### HasSuppressSessionAnalysisWarning

`func (o *Query1) HasSuppressSessionAnalysisWarning() bool`

HasSuppressSessionAnalysisWarning returns a boolean if a field has been set.

### SetSuppressSessionAnalysisWarningNil

`func (o *Query1) SetSuppressSessionAnalysisWarningNil(b bool)`

 SetSuppressSessionAnalysisWarningNil sets the value for SuppressSessionAnalysisWarning to be an explicit nil

### UnsetSuppressSessionAnalysisWarning
`func (o *Query1) UnsetSuppressSessionAnalysisWarning()`

UnsetSuppressSessionAnalysisWarning ensures that no value is present for SuppressSessionAnalysisWarning, not even an explicit nil
### GetVizSpecificOptions

`func (o *Query1) GetVizSpecificOptions() VizSpecificOptions`

GetVizSpecificOptions returns the VizSpecificOptions field if non-nil, zero value otherwise.

### GetVizSpecificOptionsOk

`func (o *Query1) GetVizSpecificOptionsOk() (*VizSpecificOptions, bool)`

GetVizSpecificOptionsOk returns a tuple with the VizSpecificOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVizSpecificOptions

`func (o *Query1) SetVizSpecificOptions(v VizSpecificOptions)`

SetVizSpecificOptions sets VizSpecificOptions field to given value.

### HasVizSpecificOptions

`func (o *Query1) HasVizSpecificOptions() bool`

HasVizSpecificOptions returns a boolean if a field has been set.

### GetAggregationGroupTypeIndex

`func (o *Query1) GetAggregationGroupTypeIndex() int32`

GetAggregationGroupTypeIndex returns the AggregationGroupTypeIndex field if non-nil, zero value otherwise.

### GetAggregationGroupTypeIndexOk

`func (o *Query1) GetAggregationGroupTypeIndexOk() (*int32, bool)`

GetAggregationGroupTypeIndexOk returns a tuple with the AggregationGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationGroupTypeIndex

`func (o *Query1) SetAggregationGroupTypeIndex(v int32)`

SetAggregationGroupTypeIndex sets AggregationGroupTypeIndex field to given value.

### HasAggregationGroupTypeIndex

`func (o *Query1) HasAggregationGroupTypeIndex() bool`

HasAggregationGroupTypeIndex returns a boolean if a field has been set.

### SetAggregationGroupTypeIndexNil

`func (o *Query1) SetAggregationGroupTypeIndexNil(b bool)`

 SetAggregationGroupTypeIndexNil sets the value for AggregationGroupTypeIndex to be an explicit nil

### UnsetAggregationGroupTypeIndex
`func (o *Query1) UnsetAggregationGroupTypeIndex()`

UnsetAggregationGroupTypeIndex ensures that no value is present for AggregationGroupTypeIndex, not even an explicit nil
### GetBreakdownFilter

`func (o *Query1) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *Query1) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *Query1) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *Query1) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetDataColorTheme

`func (o *Query1) GetDataColorTheme() float32`

GetDataColorTheme returns the DataColorTheme field if non-nil, zero value otherwise.

### GetDataColorThemeOk

`func (o *Query1) GetDataColorThemeOk() (*float32, bool)`

GetDataColorThemeOk returns a tuple with the DataColorTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorTheme

`func (o *Query1) SetDataColorTheme(v float32)`

SetDataColorTheme sets DataColorTheme field to given value.

### HasDataColorTheme

`func (o *Query1) HasDataColorTheme() bool`

HasDataColorTheme returns a boolean if a field has been set.

### SetDataColorThemeNil

`func (o *Query1) SetDataColorThemeNil(b bool)`

 SetDataColorThemeNil sets the value for DataColorTheme to be an explicit nil

### UnsetDataColorTheme
`func (o *Query1) UnsetDataColorTheme()`

UnsetDataColorTheme ensures that no value is present for DataColorTheme, not even an explicit nil
### GetSamplingFactor

`func (o *Query1) GetSamplingFactor() float32`

GetSamplingFactor returns the SamplingFactor field if non-nil, zero value otherwise.

### GetSamplingFactorOk

`func (o *Query1) GetSamplingFactorOk() (*float32, bool)`

GetSamplingFactorOk returns a tuple with the SamplingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingFactor

`func (o *Query1) SetSamplingFactor(v float32)`

SetSamplingFactor sets SamplingFactor field to given value.

### HasSamplingFactor

`func (o *Query1) HasSamplingFactor() bool`

HasSamplingFactor returns a boolean if a field has been set.

### SetSamplingFactorNil

`func (o *Query1) SetSamplingFactorNil(b bool)`

 SetSamplingFactorNil sets the value for SamplingFactor to be an explicit nil

### UnsetSamplingFactor
`func (o *Query1) UnsetSamplingFactor()`

UnsetSamplingFactor ensures that no value is present for SamplingFactor, not even an explicit nil
### GetTrendsFilter

`func (o *Query1) GetTrendsFilter() TrendsFilter`

GetTrendsFilter returns the TrendsFilter field if non-nil, zero value otherwise.

### GetTrendsFilterOk

`func (o *Query1) GetTrendsFilterOk() (*TrendsFilter, bool)`

GetTrendsFilterOk returns a tuple with the TrendsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendsFilter

`func (o *Query1) SetTrendsFilter(v TrendsFilter)`

SetTrendsFilter sets TrendsFilter field to given value.

### HasTrendsFilter

`func (o *Query1) HasTrendsFilter() bool`

HasTrendsFilter returns a boolean if a field has been set.

### GetFunnelsFilter

`func (o *Query1) GetFunnelsFilter() FunnelsFilter`

GetFunnelsFilter returns the FunnelsFilter field if non-nil, zero value otherwise.

### GetFunnelsFilterOk

`func (o *Query1) GetFunnelsFilterOk() (*FunnelsFilter, bool)`

GetFunnelsFilterOk returns a tuple with the FunnelsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsFilter

`func (o *Query1) SetFunnelsFilter(v FunnelsFilter)`

SetFunnelsFilter sets FunnelsFilter field to given value.

### HasFunnelsFilter

`func (o *Query1) HasFunnelsFilter() bool`

HasFunnelsFilter returns a boolean if a field has been set.

### GetRetentionFilter

`func (o *Query1) GetRetentionFilter() RetentionFilter`

GetRetentionFilter returns the RetentionFilter field if non-nil, zero value otherwise.

### GetRetentionFilterOk

`func (o *Query1) GetRetentionFilterOk() (*RetentionFilter, bool)`

GetRetentionFilterOk returns a tuple with the RetentionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionFilter

`func (o *Query1) SetRetentionFilter(v RetentionFilter)`

SetRetentionFilter sets RetentionFilter field to given value.


### GetFunnelPathsFilter

`func (o *Query1) GetFunnelPathsFilter() FunnelPathsFilter`

GetFunnelPathsFilter returns the FunnelPathsFilter field if non-nil, zero value otherwise.

### GetFunnelPathsFilterOk

`func (o *Query1) GetFunnelPathsFilterOk() (*FunnelPathsFilter, bool)`

GetFunnelPathsFilterOk returns a tuple with the FunnelPathsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelPathsFilter

`func (o *Query1) SetFunnelPathsFilter(v FunnelPathsFilter)`

SetFunnelPathsFilter sets FunnelPathsFilter field to given value.

### HasFunnelPathsFilter

`func (o *Query1) HasFunnelPathsFilter() bool`

HasFunnelPathsFilter returns a boolean if a field has been set.

### GetPathsFilter

`func (o *Query1) GetPathsFilter() PathsFilter`

GetPathsFilter returns the PathsFilter field if non-nil, zero value otherwise.

### GetPathsFilterOk

`func (o *Query1) GetPathsFilterOk() (*PathsFilter, bool)`

GetPathsFilterOk returns a tuple with the PathsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathsFilter

`func (o *Query1) SetPathsFilter(v PathsFilter)`

SetPathsFilter sets PathsFilter field to given value.


### GetIntervalCount

`func (o *Query1) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *Query1) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *Query1) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *Query1) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *Query1) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *Query1) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetStickinessFilter

`func (o *Query1) GetStickinessFilter() StickinessFilter`

GetStickinessFilter returns the StickinessFilter field if non-nil, zero value otherwise.

### GetStickinessFilterOk

`func (o *Query1) GetStickinessFilterOk() (*StickinessFilter, bool)`

GetStickinessFilterOk returns a tuple with the StickinessFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickinessFilter

`func (o *Query1) SetStickinessFilter(v StickinessFilter)`

SetStickinessFilter sets StickinessFilter field to given value.

### HasStickinessFilter

`func (o *Query1) HasStickinessFilter() bool`

HasStickinessFilter returns a boolean if a field has been set.

### GetLifecycleFilter

`func (o *Query1) GetLifecycleFilter() LifecycleFilter`

GetLifecycleFilter returns the LifecycleFilter field if non-nil, zero value otherwise.

### GetLifecycleFilterOk

`func (o *Query1) GetLifecycleFilterOk() (*LifecycleFilter, bool)`

GetLifecycleFilterOk returns a tuple with the LifecycleFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleFilter

`func (o *Query1) SetLifecycleFilter(v LifecycleFilter)`

SetLifecycleFilter sets LifecycleFilter field to given value.

### HasLifecycleFilter

`func (o *Query1) HasLifecycleFilter() bool`

HasLifecycleFilter returns a boolean if a field has been set.

### GetFunnelCorrelationEventExcludePropertyNames

`func (o *Query1) GetFunnelCorrelationEventExcludePropertyNames() []string`

GetFunnelCorrelationEventExcludePropertyNames returns the FunnelCorrelationEventExcludePropertyNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationEventExcludePropertyNamesOk

`func (o *Query1) GetFunnelCorrelationEventExcludePropertyNamesOk() (*[]string, bool)`

GetFunnelCorrelationEventExcludePropertyNamesOk returns a tuple with the FunnelCorrelationEventExcludePropertyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationEventExcludePropertyNames

`func (o *Query1) SetFunnelCorrelationEventExcludePropertyNames(v []string)`

SetFunnelCorrelationEventExcludePropertyNames sets FunnelCorrelationEventExcludePropertyNames field to given value.

### HasFunnelCorrelationEventExcludePropertyNames

`func (o *Query1) HasFunnelCorrelationEventExcludePropertyNames() bool`

HasFunnelCorrelationEventExcludePropertyNames returns a boolean if a field has been set.

### SetFunnelCorrelationEventExcludePropertyNamesNil

`func (o *Query1) SetFunnelCorrelationEventExcludePropertyNamesNil(b bool)`

 SetFunnelCorrelationEventExcludePropertyNamesNil sets the value for FunnelCorrelationEventExcludePropertyNames to be an explicit nil

### UnsetFunnelCorrelationEventExcludePropertyNames
`func (o *Query1) UnsetFunnelCorrelationEventExcludePropertyNames()`

UnsetFunnelCorrelationEventExcludePropertyNames ensures that no value is present for FunnelCorrelationEventExcludePropertyNames, not even an explicit nil
### GetFunnelCorrelationEventNames

`func (o *Query1) GetFunnelCorrelationEventNames() []string`

GetFunnelCorrelationEventNames returns the FunnelCorrelationEventNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationEventNamesOk

`func (o *Query1) GetFunnelCorrelationEventNamesOk() (*[]string, bool)`

GetFunnelCorrelationEventNamesOk returns a tuple with the FunnelCorrelationEventNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationEventNames

`func (o *Query1) SetFunnelCorrelationEventNames(v []string)`

SetFunnelCorrelationEventNames sets FunnelCorrelationEventNames field to given value.

### HasFunnelCorrelationEventNames

`func (o *Query1) HasFunnelCorrelationEventNames() bool`

HasFunnelCorrelationEventNames returns a boolean if a field has been set.

### SetFunnelCorrelationEventNamesNil

`func (o *Query1) SetFunnelCorrelationEventNamesNil(b bool)`

 SetFunnelCorrelationEventNamesNil sets the value for FunnelCorrelationEventNames to be an explicit nil

### UnsetFunnelCorrelationEventNames
`func (o *Query1) UnsetFunnelCorrelationEventNames()`

UnsetFunnelCorrelationEventNames ensures that no value is present for FunnelCorrelationEventNames, not even an explicit nil
### GetFunnelCorrelationExcludeEventNames

`func (o *Query1) GetFunnelCorrelationExcludeEventNames() []string`

GetFunnelCorrelationExcludeEventNames returns the FunnelCorrelationExcludeEventNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationExcludeEventNamesOk

`func (o *Query1) GetFunnelCorrelationExcludeEventNamesOk() (*[]string, bool)`

GetFunnelCorrelationExcludeEventNamesOk returns a tuple with the FunnelCorrelationExcludeEventNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationExcludeEventNames

`func (o *Query1) SetFunnelCorrelationExcludeEventNames(v []string)`

SetFunnelCorrelationExcludeEventNames sets FunnelCorrelationExcludeEventNames field to given value.

### HasFunnelCorrelationExcludeEventNames

`func (o *Query1) HasFunnelCorrelationExcludeEventNames() bool`

HasFunnelCorrelationExcludeEventNames returns a boolean if a field has been set.

### SetFunnelCorrelationExcludeEventNamesNil

`func (o *Query1) SetFunnelCorrelationExcludeEventNamesNil(b bool)`

 SetFunnelCorrelationExcludeEventNamesNil sets the value for FunnelCorrelationExcludeEventNames to be an explicit nil

### UnsetFunnelCorrelationExcludeEventNames
`func (o *Query1) UnsetFunnelCorrelationExcludeEventNames()`

UnsetFunnelCorrelationExcludeEventNames ensures that no value is present for FunnelCorrelationExcludeEventNames, not even an explicit nil
### GetFunnelCorrelationExcludeNames

`func (o *Query1) GetFunnelCorrelationExcludeNames() []string`

GetFunnelCorrelationExcludeNames returns the FunnelCorrelationExcludeNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationExcludeNamesOk

`func (o *Query1) GetFunnelCorrelationExcludeNamesOk() (*[]string, bool)`

GetFunnelCorrelationExcludeNamesOk returns a tuple with the FunnelCorrelationExcludeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationExcludeNames

`func (o *Query1) SetFunnelCorrelationExcludeNames(v []string)`

SetFunnelCorrelationExcludeNames sets FunnelCorrelationExcludeNames field to given value.

### HasFunnelCorrelationExcludeNames

`func (o *Query1) HasFunnelCorrelationExcludeNames() bool`

HasFunnelCorrelationExcludeNames returns a boolean if a field has been set.

### SetFunnelCorrelationExcludeNamesNil

`func (o *Query1) SetFunnelCorrelationExcludeNamesNil(b bool)`

 SetFunnelCorrelationExcludeNamesNil sets the value for FunnelCorrelationExcludeNames to be an explicit nil

### UnsetFunnelCorrelationExcludeNames
`func (o *Query1) UnsetFunnelCorrelationExcludeNames()`

UnsetFunnelCorrelationExcludeNames ensures that no value is present for FunnelCorrelationExcludeNames, not even an explicit nil
### GetFunnelCorrelationNames

`func (o *Query1) GetFunnelCorrelationNames() []string`

GetFunnelCorrelationNames returns the FunnelCorrelationNames field if non-nil, zero value otherwise.

### GetFunnelCorrelationNamesOk

`func (o *Query1) GetFunnelCorrelationNamesOk() (*[]string, bool)`

GetFunnelCorrelationNamesOk returns a tuple with the FunnelCorrelationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationNames

`func (o *Query1) SetFunnelCorrelationNames(v []string)`

SetFunnelCorrelationNames sets FunnelCorrelationNames field to given value.

### HasFunnelCorrelationNames

`func (o *Query1) HasFunnelCorrelationNames() bool`

HasFunnelCorrelationNames returns a boolean if a field has been set.

### SetFunnelCorrelationNamesNil

`func (o *Query1) SetFunnelCorrelationNamesNil(b bool)`

 SetFunnelCorrelationNamesNil sets the value for FunnelCorrelationNames to be an explicit nil

### UnsetFunnelCorrelationNames
`func (o *Query1) UnsetFunnelCorrelationNames()`

UnsetFunnelCorrelationNames ensures that no value is present for FunnelCorrelationNames, not even an explicit nil
### GetFunnelCorrelationType

`func (o *Query1) GetFunnelCorrelationType() FunnelCorrelationResultsType`

GetFunnelCorrelationType returns the FunnelCorrelationType field if non-nil, zero value otherwise.

### GetFunnelCorrelationTypeOk

`func (o *Query1) GetFunnelCorrelationTypeOk() (*FunnelCorrelationResultsType, bool)`

GetFunnelCorrelationTypeOk returns a tuple with the FunnelCorrelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelCorrelationType

`func (o *Query1) SetFunnelCorrelationType(v FunnelCorrelationResultsType)`

SetFunnelCorrelationType sets FunnelCorrelationType field to given value.


### GetServiceNames

`func (o *Query1) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *Query1) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *Query1) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.


### GetSeverityLevels

`func (o *Query1) GetSeverityLevels() []LogSeverityLevel`

GetSeverityLevels returns the SeverityLevels field if non-nil, zero value otherwise.

### GetSeverityLevelsOk

`func (o *Query1) GetSeverityLevelsOk() (*[]LogSeverityLevel, bool)`

GetSeverityLevelsOk returns a tuple with the SeverityLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevels

`func (o *Query1) SetSeverityLevels(v []LogSeverityLevel)`

SetSeverityLevels sets SeverityLevels field to given value.


### GetMaxPropertyValues

`func (o *Query1) GetMaxPropertyValues() int32`

GetMaxPropertyValues returns the MaxPropertyValues field if non-nil, zero value otherwise.

### GetMaxPropertyValuesOk

`func (o *Query1) GetMaxPropertyValuesOk() (*int32, bool)`

GetMaxPropertyValuesOk returns a tuple with the MaxPropertyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPropertyValues

`func (o *Query1) SetMaxPropertyValues(v int32)`

SetMaxPropertyValues sets MaxPropertyValues field to given value.

### HasMaxPropertyValues

`func (o *Query1) HasMaxPropertyValues() bool`

HasMaxPropertyValues returns a boolean if a field has been set.

### SetMaxPropertyValuesNil

`func (o *Query1) SetMaxPropertyValuesNil(b bool)`

 SetMaxPropertyValuesNil sets the value for MaxPropertyValues to be an explicit nil

### UnsetMaxPropertyValues
`func (o *Query1) UnsetMaxPropertyValues()`

UnsetMaxPropertyValues ensures that no value is present for MaxPropertyValues, not even an explicit nil
### GetTraceId

`func (o *Query1) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *Query1) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *Query1) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetEmbedding

`func (o *Query1) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *Query1) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *Query1) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.


### GetEmbeddingVersion

`func (o *Query1) GetEmbeddingVersion() float32`

GetEmbeddingVersion returns the EmbeddingVersion field if non-nil, zero value otherwise.

### GetEmbeddingVersionOk

`func (o *Query1) GetEmbeddingVersionOk() (*float32, bool)`

GetEmbeddingVersionOk returns a tuple with the EmbeddingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingVersion

`func (o *Query1) SetEmbeddingVersion(v float32)`

SetEmbeddingVersion sets EmbeddingVersion field to given value.

### HasEmbeddingVersion

`func (o *Query1) HasEmbeddingVersion() bool`

HasEmbeddingVersion returns a boolean if a field has been set.

### SetEmbeddingVersionNil

`func (o *Query1) SetEmbeddingVersionNil(b bool)`

 SetEmbeddingVersionNil sets the value for EmbeddingVersion to be an explicit nil

### UnsetEmbeddingVersion
`func (o *Query1) UnsetEmbeddingVersion()`

UnsetEmbeddingVersion ensures that no value is present for EmbeddingVersion, not even an explicit nil
### GetGroupKey

`func (o *Query1) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *Query1) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *Query1) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *Query1) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### SetGroupKeyNil

`func (o *Query1) SetGroupKeyNil(b bool)`

 SetGroupKeyNil sets the value for GroupKey to be an explicit nil

### UnsetGroupKey
`func (o *Query1) UnsetGroupKey()`

UnsetGroupKey ensures that no value is present for GroupKey, not even an explicit nil
### GetPersonId

`func (o *Query1) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Query1) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Query1) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Query1) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *Query1) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *Query1) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


