# Sourcequery1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to **NullableString** | Limit to events matching this string | [optional] 
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
**Id** | **int32** |  | 
**Cohort** | Pointer to **NullableInt32** |  | [optional] 
**DistinctId** | Pointer to **NullableString** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**ActionId** | Pointer to **NullableInt32** | Show events matching a given action | [optional] 
**After** | Pointer to **NullableString** | Cursor for pagination. Contains the ordering value and session_id from the last record of the previous page. | [optional] 
**Before** | Pointer to **NullableString** | Only fetch sessions that started before this timestamp (default: &#39;+5s&#39;) | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**PersonId** | Pointer to **NullableString** | Person who performed the event | [optional] 
**Select** | **[]string** | Return a limited set of data. Will use default columns if empty. | 
**Source** | [**Source3**](Source3.md) |  | 
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
**Query** | **string** | Query to validate | 
**Values** | Pointer to **map[string]interface{}** | Constant values that can be referenced with the {placeholder} syntax in the query | [optional] 
**Variables** | Pointer to [**map[string]HogQLVariable**](HogQLVariable.md) | Variables to be subsituted into the query | [optional] 
**Debug** | Pointer to **NullableBool** | Enable more verbose output, usually run from the /debug page | [optional] 
**Globals** | Pointer to **map[string]interface{}** | Global values in scope | [optional] 
**Language** | [**HogLanguage**](HogLanguage.md) |  | 
**SourceQuery** | Pointer to [**NullableSourcequery**](Sourcequery.md) |  | [optional] [default to null]
**EndPosition** | **int32** | End position of the editor word | 
**StartPosition** | **int32** | Start position of the editor word | 
**DateRange** | [**DateRange**](DateRange.md) |  | 
**GroupBy** | [**[]SessionAttributionGroupBy**](SessionAttributionGroupBy.md) |  | 
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal**](Conversiongoal.md) |  | [optional] [default to null]
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**DraftConversionGoal** | Pointer to [**NullableDraftconversiongoal**](Draftconversiongoal.md) |  | [optional] [default to null]
**IncludeAllConversions** | Pointer to **NullableBool** | Include conversion goal rows even when they don&#39;t match campaign costs table | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**IntegrationFilter** | Pointer to [**IntegrationFilter**](IntegrationFilter.md) |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**BreakdownBy** | [**WebStatsBreakdown**](WebStatsBreakdown.md) |  | 
**IncludeBounceRate** | Pointer to **NullableBool** |  | [optional] 
**IncludeScrollDepth** | Pointer to **NullableBool** |  | [optional] 
**StripQueryParams** | Pointer to **NullableBool** |  | [optional] 
**Metric** | [**WebVitalsMetric**](WebVitalsMetric.md) |  | 
**Percentile** | [**WebVitalsPercentile**](WebVitalsPercentile.md) |  | 
**Thresholds** | **[]float32** |  | 
**SearchTerm** | Pointer to **NullableString** |  | [optional] 
**Metrics** | [**[]WebTrendsMetric**](WebTrendsMetric.md) |  | 
**Assignee** | Pointer to [**ErrorTrackingIssueAssignee**](ErrorTrackingIssueAssignee.md) |  | [optional] 
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
**Events** | **[]map[string]interface{}** |  | 
**ServiceNames** | **[]string** |  | 
**SeverityLevels** | [**[]LogSeverityLevel**](LogSeverityLevel.md) |  | 
**ExperimentId** | Pointer to **NullableInt32** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**FunnelsQuery** | [**FunnelsQuery**](FunnelsQuery.md) |  | 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**CountQuery** | [**TrendsQuery**](TrendsQuery.md) |  | 
**ExposureQuery** | Pointer to [**TrendsQuery**](TrendsQuery.md) |  | [optional] 
**AggregationGroupTypeIndex** | Pointer to **NullableInt32** | Groups aggregation | [optional] 
**CalendarHeatmapFilter** | Pointer to [**CalendarHeatmapFilter**](CalendarHeatmapFilter.md) |  | [optional] 
**DataColorTheme** | Pointer to **NullableFloat32** | Colors used in the insight&#39;s visualization | [optional] 
**SamplingFactor** | Pointer to **NullableFloat32** | Sampling rate | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CommentText** | Pointer to [**RecordingPropertyFilter**](RecordingPropertyFilter.md) |  | [optional] 
**ConsoleLogFilters** | Pointer to [**[]LogEntryPropertyFilter**](LogEntryPropertyFilter.md) |  | [optional] 
**DateFrom** | Pointer to **NullableString** |  | [optional] [default to "-3d"]
**DateTo** | Pointer to **NullableString** |  | [optional] 
**DistinctIds** | Pointer to **[]string** |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**HavingPredicates** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) |  | [optional] 
**Operand** | Pointer to [**FilterLogicalOperator**](FilterLogicalOperator.md) |  | [optional] 
**Order** | Pointer to [**RecordingOrder**](RecordingOrder.md) |  | [optional] 
**OrderDirection** | Pointer to [**RecordingOrderDirection**](RecordingOrderDirection.md) |  | [optional] 
**PersonUuid** | Pointer to **NullableString** |  | [optional] 
**SessionIds** | Pointer to **[]string** |  | [optional] 
**SessionRecordingId** | Pointer to **NullableString** | If provided, this recording will be fetched and prepended to the results, even if it doesn&#39;t match the filters | [optional] 
**UserModifiedFilters** | Pointer to **map[string]interface{}** |  | [optional] 
**ShowColumnConfigurator** | Pointer to **NullableBool** |  | [optional] 
**TraceId** | **string** |  | 
**Embedding** | **[]float32** |  | 
**EmbeddingVersion** | Pointer to **NullableFloat32** |  | [optional] 
**GroupKey** | Pointer to **NullableString** | Group key. Required with group_type_index for group queries. | [optional] 
**PersonId** | Pointer to **NullableString** | Person ID to fetch metrics for. Mutually exclusive with group parameters. | [optional] 

## Methods

### NewSourcequery1

`func NewSourcequery1(orderBy OrderBy3, properties []FixedpropertiesInner, id int32, select_ []string, source Source3, groupTypeIndex NullableInt32, breakdown []RevenueAnalyticsBreakdown, interval IntervalType, series []Series1Inner, query string, language HogLanguage, endPosition int32, startPosition int32, dateRange DateRange, groupBy []SessionAttributionGroupBy, breakdownBy WebStatsBreakdown, metric WebVitalsMetric, percentile WebVitalsPercentile, thresholds []float32, metrics []WebTrendsMetric, filterGroup PropertyGroupFilter, issueId string, volumeResolution int32, breakdownProperties []string, events []map[string]interface{}, serviceNames []string, severityLevels []LogSeverityLevel, funnelsQuery FunnelsQuery, countQuery TrendsQuery, traceId string, embedding []float32, ) *Sourcequery1`

NewSourcequery1 instantiates a new Sourcequery1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcequery1WithDefaults

`func NewSourcequery1WithDefaults() *Sourcequery1`

NewSourcequery1WithDefaults instantiates a new Sourcequery1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *Sourcequery1) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *Sourcequery1) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *Sourcequery1) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *Sourcequery1) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *Sourcequery1) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *Sourcequery1) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetEvent

`func (o *Sourcequery1) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Sourcequery1) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Sourcequery1) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Sourcequery1) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *Sourcequery1) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *Sourcequery1) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetFixedProperties

`func (o *Sourcequery1) GetFixedProperties() []FixedpropertiesInner1`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *Sourcequery1) GetFixedPropertiesOk() (*[]FixedpropertiesInner1, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *Sourcequery1) SetFixedProperties(v []FixedpropertiesInner1)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *Sourcequery1) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *Sourcequery1) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *Sourcequery1) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *Sourcequery1) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Sourcequery1) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Sourcequery1) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Sourcequery1) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *Sourcequery1) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Sourcequery1) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Sourcequery1) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Sourcequery1) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Sourcequery1) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Sourcequery1) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMath

`func (o *Sourcequery1) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *Sourcequery1) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *Sourcequery1) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *Sourcequery1) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *Sourcequery1) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *Sourcequery1) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *Sourcequery1) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *Sourcequery1) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *Sourcequery1) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *Sourcequery1) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *Sourcequery1) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *Sourcequery1) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *Sourcequery1) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *Sourcequery1) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *Sourcequery1) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *Sourcequery1) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *Sourcequery1) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *Sourcequery1) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *Sourcequery1) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *Sourcequery1) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *Sourcequery1) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *Sourcequery1) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *Sourcequery1) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *Sourcequery1) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *Sourcequery1) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *Sourcequery1) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *Sourcequery1) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *Sourcequery1) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *Sourcequery1) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *Sourcequery1) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *Sourcequery1) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *Sourcequery1) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *Sourcequery1) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *Sourcequery1) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *Sourcequery1) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *Sourcequery1) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *Sourcequery1) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *Sourcequery1) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *Sourcequery1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sourcequery1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sourcequery1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sourcequery1) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Sourcequery1) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Sourcequery1) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *Sourcequery1) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *Sourcequery1) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *Sourcequery1) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *Sourcequery1) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *Sourcequery1) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *Sourcequery1) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetOrderBy

`func (o *Sourcequery1) GetOrderBy() OrderBy3`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *Sourcequery1) GetOrderByOk() (*OrderBy3, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *Sourcequery1) SetOrderBy(v OrderBy3)`

SetOrderBy sets OrderBy field to given value.


### GetProperties

`func (o *Sourcequery1) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Sourcequery1) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Sourcequery1) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.


### SetPropertiesNil

`func (o *Sourcequery1) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Sourcequery1) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *Sourcequery1) GetResponse() UsageMetricsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Sourcequery1) GetResponseOk() (*UsageMetricsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Sourcequery1) SetResponse(v UsageMetricsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Sourcequery1) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetVersion

`func (o *Sourcequery1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Sourcequery1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Sourcequery1) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Sourcequery1) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Sourcequery1) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Sourcequery1) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *Sourcequery1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sourcequery1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sourcequery1) SetId(v int32)`

SetId sets Id field to given value.


### GetCohort

`func (o *Sourcequery1) GetCohort() int32`

GetCohort returns the Cohort field if non-nil, zero value otherwise.

### GetCohortOk

`func (o *Sourcequery1) GetCohortOk() (*int32, bool)`

GetCohortOk returns a tuple with the Cohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohort

`func (o *Sourcequery1) SetCohort(v int32)`

SetCohort sets Cohort field to given value.

### HasCohort

`func (o *Sourcequery1) HasCohort() bool`

HasCohort returns a boolean if a field has been set.

### SetCohortNil

`func (o *Sourcequery1) SetCohortNil(b bool)`

 SetCohortNil sets the value for Cohort to be an explicit nil

### UnsetCohort
`func (o *Sourcequery1) UnsetCohort()`

UnsetCohort ensures that no value is present for Cohort, not even an explicit nil
### GetDistinctId

`func (o *Sourcequery1) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *Sourcequery1) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *Sourcequery1) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *Sourcequery1) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### SetDistinctIdNil

`func (o *Sourcequery1) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *Sourcequery1) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetModifiers

`func (o *Sourcequery1) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Sourcequery1) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Sourcequery1) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Sourcequery1) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *Sourcequery1) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Sourcequery1) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Sourcequery1) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Sourcequery1) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *Sourcequery1) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *Sourcequery1) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSearch

`func (o *Sourcequery1) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *Sourcequery1) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *Sourcequery1) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *Sourcequery1) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *Sourcequery1) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *Sourcequery1) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetTags

`func (o *Sourcequery1) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Sourcequery1) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Sourcequery1) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Sourcequery1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetActionId

`func (o *Sourcequery1) GetActionId() int32`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *Sourcequery1) GetActionIdOk() (*int32, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *Sourcequery1) SetActionId(v int32)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *Sourcequery1) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### SetActionIdNil

`func (o *Sourcequery1) SetActionIdNil(b bool)`

 SetActionIdNil sets the value for ActionId to be an explicit nil

### UnsetActionId
`func (o *Sourcequery1) UnsetActionId()`

UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
### GetAfter

`func (o *Sourcequery1) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Sourcequery1) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Sourcequery1) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *Sourcequery1) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *Sourcequery1) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *Sourcequery1) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *Sourcequery1) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Sourcequery1) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Sourcequery1) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *Sourcequery1) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *Sourcequery1) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *Sourcequery1) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetFilterTestAccounts

`func (o *Sourcequery1) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *Sourcequery1) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *Sourcequery1) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *Sourcequery1) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *Sourcequery1) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *Sourcequery1) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetPersonId

`func (o *Sourcequery1) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Sourcequery1) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Sourcequery1) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Sourcequery1) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *Sourcequery1) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *Sourcequery1) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetSelect

`func (o *Sourcequery1) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *Sourcequery1) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *Sourcequery1) SetSelect(v []string)`

SetSelect sets Select field to given value.


### SetSelectNil

`func (o *Sourcequery1) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *Sourcequery1) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetSource

`func (o *Sourcequery1) GetSource() Source3`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Sourcequery1) GetSourceOk() (*Source3, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Sourcequery1) SetSource(v Source3)`

SetSource sets Source field to given value.


### GetWhere

`func (o *Sourcequery1) GetWhere() []string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *Sourcequery1) GetWhereOk() (*[]string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *Sourcequery1) SetWhere(v []string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *Sourcequery1) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### SetWhereNil

`func (o *Sourcequery1) SetWhereNil(b bool)`

 SetWhereNil sets the value for Where to be an explicit nil

### UnsetWhere
`func (o *Sourcequery1) UnsetWhere()`

UnsetWhere ensures that no value is present for Where, not even an explicit nil
### GetGroupTypeIndex

`func (o *Sourcequery1) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Sourcequery1) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Sourcequery1) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.


### SetGroupTypeIndexNil

`func (o *Sourcequery1) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Sourcequery1) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetBreakdown

`func (o *Sourcequery1) GetBreakdown() []RevenueAnalyticsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Sourcequery1) GetBreakdownOk() (*[]RevenueAnalyticsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Sourcequery1) SetBreakdown(v []RevenueAnalyticsBreakdown)`

SetBreakdown sets Breakdown field to given value.


### GetCompare

`func (o *Sourcequery1) GetCompare() Compare`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *Sourcequery1) GetCompareOk() (*Compare, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *Sourcequery1) SetCompare(v Compare)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *Sourcequery1) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### GetDay

`func (o *Sourcequery1) GetDay() Day`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Sourcequery1) GetDayOk() (*Day, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Sourcequery1) SetDay(v Day)`

SetDay sets Day field to given value.

### HasDay

`func (o *Sourcequery1) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *Sourcequery1) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *Sourcequery1) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetIncludeRecordings

`func (o *Sourcequery1) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *Sourcequery1) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *Sourcequery1) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *Sourcequery1) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *Sourcequery1) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *Sourcequery1) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetInterval

`func (o *Sourcequery1) GetInterval() IntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Sourcequery1) GetIntervalOk() (*IntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Sourcequery1) SetInterval(v IntervalType)`

SetInterval sets Interval field to given value.


### GetSeries

`func (o *Sourcequery1) GetSeries() []Series1Inner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Sourcequery1) GetSeriesOk() (*[]Series1Inner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Sourcequery1) SetSeries(v []Series1Inner)`

SetSeries sets Series field to given value.


### GetStatus

`func (o *Sourcequery1) GetStatus() Status2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Sourcequery1) GetStatusOk() (*Status2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Sourcequery1) SetStatus(v Status2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Sourcequery1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *Sourcequery1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Sourcequery1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Sourcequery1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Sourcequery1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Sourcequery1) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Sourcequery1) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetExplain

`func (o *Sourcequery1) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *Sourcequery1) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *Sourcequery1) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *Sourcequery1) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *Sourcequery1) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *Sourcequery1) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetFilters

`func (o *Sourcequery1) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Sourcequery1) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Sourcequery1) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Sourcequery1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetQuery

`func (o *Sourcequery1) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Sourcequery1) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Sourcequery1) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetValues

`func (o *Sourcequery1) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Sourcequery1) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Sourcequery1) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *Sourcequery1) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *Sourcequery1) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *Sourcequery1) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetVariables

`func (o *Sourcequery1) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Sourcequery1) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Sourcequery1) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Sourcequery1) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *Sourcequery1) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Sourcequery1) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetDebug

`func (o *Sourcequery1) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *Sourcequery1) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *Sourcequery1) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *Sourcequery1) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### SetDebugNil

`func (o *Sourcequery1) SetDebugNil(b bool)`

 SetDebugNil sets the value for Debug to be an explicit nil

### UnsetDebug
`func (o *Sourcequery1) UnsetDebug()`

UnsetDebug ensures that no value is present for Debug, not even an explicit nil
### GetGlobals

`func (o *Sourcequery1) GetGlobals() map[string]interface{}`

GetGlobals returns the Globals field if non-nil, zero value otherwise.

### GetGlobalsOk

`func (o *Sourcequery1) GetGlobalsOk() (*map[string]interface{}, bool)`

GetGlobalsOk returns a tuple with the Globals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobals

`func (o *Sourcequery1) SetGlobals(v map[string]interface{})`

SetGlobals sets Globals field to given value.

### HasGlobals

`func (o *Sourcequery1) HasGlobals() bool`

HasGlobals returns a boolean if a field has been set.

### SetGlobalsNil

`func (o *Sourcequery1) SetGlobalsNil(b bool)`

 SetGlobalsNil sets the value for Globals to be an explicit nil

### UnsetGlobals
`func (o *Sourcequery1) UnsetGlobals()`

UnsetGlobals ensures that no value is present for Globals, not even an explicit nil
### GetLanguage

`func (o *Sourcequery1) GetLanguage() HogLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Sourcequery1) GetLanguageOk() (*HogLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Sourcequery1) SetLanguage(v HogLanguage)`

SetLanguage sets Language field to given value.


### GetSourceQuery

`func (o *Sourcequery1) GetSourceQuery() Sourcequery`

GetSourceQuery returns the SourceQuery field if non-nil, zero value otherwise.

### GetSourceQueryOk

`func (o *Sourcequery1) GetSourceQueryOk() (*Sourcequery, bool)`

GetSourceQueryOk returns a tuple with the SourceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuery

`func (o *Sourcequery1) SetSourceQuery(v Sourcequery)`

SetSourceQuery sets SourceQuery field to given value.

### HasSourceQuery

`func (o *Sourcequery1) HasSourceQuery() bool`

HasSourceQuery returns a boolean if a field has been set.

### SetSourceQueryNil

`func (o *Sourcequery1) SetSourceQueryNil(b bool)`

 SetSourceQueryNil sets the value for SourceQuery to be an explicit nil

### UnsetSourceQuery
`func (o *Sourcequery1) UnsetSourceQuery()`

UnsetSourceQuery ensures that no value is present for SourceQuery, not even an explicit nil
### GetEndPosition

`func (o *Sourcequery1) GetEndPosition() int32`

GetEndPosition returns the EndPosition field if non-nil, zero value otherwise.

### GetEndPositionOk

`func (o *Sourcequery1) GetEndPositionOk() (*int32, bool)`

GetEndPositionOk returns a tuple with the EndPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPosition

`func (o *Sourcequery1) SetEndPosition(v int32)`

SetEndPosition sets EndPosition field to given value.


### GetStartPosition

`func (o *Sourcequery1) GetStartPosition() int32`

GetStartPosition returns the StartPosition field if non-nil, zero value otherwise.

### GetStartPositionOk

`func (o *Sourcequery1) GetStartPositionOk() (*int32, bool)`

GetStartPositionOk returns a tuple with the StartPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPosition

`func (o *Sourcequery1) SetStartPosition(v int32)`

SetStartPosition sets StartPosition field to given value.


### GetDateRange

`func (o *Sourcequery1) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Sourcequery1) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Sourcequery1) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetGroupBy

`func (o *Sourcequery1) GetGroupBy() []SessionAttributionGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Sourcequery1) GetGroupByOk() (*[]SessionAttributionGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Sourcequery1) SetGroupBy(v []SessionAttributionGroupBy)`

SetGroupBy sets GroupBy field to given value.


### GetCompareFilter

`func (o *Sourcequery1) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *Sourcequery1) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *Sourcequery1) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *Sourcequery1) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *Sourcequery1) GetConversionGoal() Conversiongoal`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *Sourcequery1) GetConversionGoalOk() (*Conversiongoal, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *Sourcequery1) SetConversionGoal(v Conversiongoal)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *Sourcequery1) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *Sourcequery1) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *Sourcequery1) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDoPathCleaning

`func (o *Sourcequery1) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *Sourcequery1) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *Sourcequery1) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *Sourcequery1) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *Sourcequery1) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *Sourcequery1) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetDraftConversionGoal

`func (o *Sourcequery1) GetDraftConversionGoal() Draftconversiongoal`

GetDraftConversionGoal returns the DraftConversionGoal field if non-nil, zero value otherwise.

### GetDraftConversionGoalOk

`func (o *Sourcequery1) GetDraftConversionGoalOk() (*Draftconversiongoal, bool)`

GetDraftConversionGoalOk returns a tuple with the DraftConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftConversionGoal

`func (o *Sourcequery1) SetDraftConversionGoal(v Draftconversiongoal)`

SetDraftConversionGoal sets DraftConversionGoal field to given value.

### HasDraftConversionGoal

`func (o *Sourcequery1) HasDraftConversionGoal() bool`

HasDraftConversionGoal returns a boolean if a field has been set.

### SetDraftConversionGoalNil

`func (o *Sourcequery1) SetDraftConversionGoalNil(b bool)`

 SetDraftConversionGoalNil sets the value for DraftConversionGoal to be an explicit nil

### UnsetDraftConversionGoal
`func (o *Sourcequery1) UnsetDraftConversionGoal()`

UnsetDraftConversionGoal ensures that no value is present for DraftConversionGoal, not even an explicit nil
### GetIncludeAllConversions

`func (o *Sourcequery1) GetIncludeAllConversions() bool`

GetIncludeAllConversions returns the IncludeAllConversions field if non-nil, zero value otherwise.

### GetIncludeAllConversionsOk

`func (o *Sourcequery1) GetIncludeAllConversionsOk() (*bool, bool)`

GetIncludeAllConversionsOk returns a tuple with the IncludeAllConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllConversions

`func (o *Sourcequery1) SetIncludeAllConversions(v bool)`

SetIncludeAllConversions sets IncludeAllConversions field to given value.

### HasIncludeAllConversions

`func (o *Sourcequery1) HasIncludeAllConversions() bool`

HasIncludeAllConversions returns a boolean if a field has been set.

### SetIncludeAllConversionsNil

`func (o *Sourcequery1) SetIncludeAllConversionsNil(b bool)`

 SetIncludeAllConversionsNil sets the value for IncludeAllConversions to be an explicit nil

### UnsetIncludeAllConversions
`func (o *Sourcequery1) UnsetIncludeAllConversions()`

UnsetIncludeAllConversions ensures that no value is present for IncludeAllConversions, not even an explicit nil
### GetIncludeRevenue

`func (o *Sourcequery1) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *Sourcequery1) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *Sourcequery1) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *Sourcequery1) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *Sourcequery1) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *Sourcequery1) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetIntegrationFilter

`func (o *Sourcequery1) GetIntegrationFilter() IntegrationFilter`

GetIntegrationFilter returns the IntegrationFilter field if non-nil, zero value otherwise.

### GetIntegrationFilterOk

`func (o *Sourcequery1) GetIntegrationFilterOk() (*IntegrationFilter, bool)`

GetIntegrationFilterOk returns a tuple with the IntegrationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationFilter

`func (o *Sourcequery1) SetIntegrationFilter(v IntegrationFilter)`

SetIntegrationFilter sets IntegrationFilter field to given value.

### HasIntegrationFilter

`func (o *Sourcequery1) HasIntegrationFilter() bool`

HasIntegrationFilter returns a boolean if a field has been set.

### GetSampling

`func (o *Sourcequery1) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *Sourcequery1) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *Sourcequery1) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *Sourcequery1) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *Sourcequery1) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *Sourcequery1) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *Sourcequery1) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *Sourcequery1) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *Sourcequery1) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *Sourcequery1) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetBreakdownBy

`func (o *Sourcequery1) GetBreakdownBy() WebStatsBreakdown`

GetBreakdownBy returns the BreakdownBy field if non-nil, zero value otherwise.

### GetBreakdownByOk

`func (o *Sourcequery1) GetBreakdownByOk() (*WebStatsBreakdown, bool)`

GetBreakdownByOk returns a tuple with the BreakdownBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownBy

`func (o *Sourcequery1) SetBreakdownBy(v WebStatsBreakdown)`

SetBreakdownBy sets BreakdownBy field to given value.


### GetIncludeBounceRate

`func (o *Sourcequery1) GetIncludeBounceRate() bool`

GetIncludeBounceRate returns the IncludeBounceRate field if non-nil, zero value otherwise.

### GetIncludeBounceRateOk

`func (o *Sourcequery1) GetIncludeBounceRateOk() (*bool, bool)`

GetIncludeBounceRateOk returns a tuple with the IncludeBounceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBounceRate

`func (o *Sourcequery1) SetIncludeBounceRate(v bool)`

SetIncludeBounceRate sets IncludeBounceRate field to given value.

### HasIncludeBounceRate

`func (o *Sourcequery1) HasIncludeBounceRate() bool`

HasIncludeBounceRate returns a boolean if a field has been set.

### SetIncludeBounceRateNil

`func (o *Sourcequery1) SetIncludeBounceRateNil(b bool)`

 SetIncludeBounceRateNil sets the value for IncludeBounceRate to be an explicit nil

### UnsetIncludeBounceRate
`func (o *Sourcequery1) UnsetIncludeBounceRate()`

UnsetIncludeBounceRate ensures that no value is present for IncludeBounceRate, not even an explicit nil
### GetIncludeScrollDepth

`func (o *Sourcequery1) GetIncludeScrollDepth() bool`

GetIncludeScrollDepth returns the IncludeScrollDepth field if non-nil, zero value otherwise.

### GetIncludeScrollDepthOk

`func (o *Sourcequery1) GetIncludeScrollDepthOk() (*bool, bool)`

GetIncludeScrollDepthOk returns a tuple with the IncludeScrollDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeScrollDepth

`func (o *Sourcequery1) SetIncludeScrollDepth(v bool)`

SetIncludeScrollDepth sets IncludeScrollDepth field to given value.

### HasIncludeScrollDepth

`func (o *Sourcequery1) HasIncludeScrollDepth() bool`

HasIncludeScrollDepth returns a boolean if a field has been set.

### SetIncludeScrollDepthNil

`func (o *Sourcequery1) SetIncludeScrollDepthNil(b bool)`

 SetIncludeScrollDepthNil sets the value for IncludeScrollDepth to be an explicit nil

### UnsetIncludeScrollDepth
`func (o *Sourcequery1) UnsetIncludeScrollDepth()`

UnsetIncludeScrollDepth ensures that no value is present for IncludeScrollDepth, not even an explicit nil
### GetStripQueryParams

`func (o *Sourcequery1) GetStripQueryParams() bool`

GetStripQueryParams returns the StripQueryParams field if non-nil, zero value otherwise.

### GetStripQueryParamsOk

`func (o *Sourcequery1) GetStripQueryParamsOk() (*bool, bool)`

GetStripQueryParamsOk returns a tuple with the StripQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripQueryParams

`func (o *Sourcequery1) SetStripQueryParams(v bool)`

SetStripQueryParams sets StripQueryParams field to given value.

### HasStripQueryParams

`func (o *Sourcequery1) HasStripQueryParams() bool`

HasStripQueryParams returns a boolean if a field has been set.

### SetStripQueryParamsNil

`func (o *Sourcequery1) SetStripQueryParamsNil(b bool)`

 SetStripQueryParamsNil sets the value for StripQueryParams to be an explicit nil

### UnsetStripQueryParams
`func (o *Sourcequery1) UnsetStripQueryParams()`

UnsetStripQueryParams ensures that no value is present for StripQueryParams, not even an explicit nil
### GetMetric

`func (o *Sourcequery1) GetMetric() WebVitalsMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Sourcequery1) GetMetricOk() (*WebVitalsMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Sourcequery1) SetMetric(v WebVitalsMetric)`

SetMetric sets Metric field to given value.


### GetPercentile

`func (o *Sourcequery1) GetPercentile() WebVitalsPercentile`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *Sourcequery1) GetPercentileOk() (*WebVitalsPercentile, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *Sourcequery1) SetPercentile(v WebVitalsPercentile)`

SetPercentile sets Percentile field to given value.


### GetThresholds

`func (o *Sourcequery1) GetThresholds() []float32`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Sourcequery1) GetThresholdsOk() (*[]float32, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Sourcequery1) SetThresholds(v []float32)`

SetThresholds sets Thresholds field to given value.


### GetSearchTerm

`func (o *Sourcequery1) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *Sourcequery1) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *Sourcequery1) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *Sourcequery1) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### SetSearchTermNil

`func (o *Sourcequery1) SetSearchTermNil(b bool)`

 SetSearchTermNil sets the value for SearchTerm to be an explicit nil

### UnsetSearchTerm
`func (o *Sourcequery1) UnsetSearchTerm()`

UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil
### GetMetrics

`func (o *Sourcequery1) GetMetrics() []WebTrendsMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Sourcequery1) GetMetricsOk() (*[]WebTrendsMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Sourcequery1) SetMetrics(v []WebTrendsMetric)`

SetMetrics sets Metrics field to given value.


### GetAssignee

`func (o *Sourcequery1) GetAssignee() ErrorTrackingIssueAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Sourcequery1) GetAssigneeOk() (*ErrorTrackingIssueAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Sourcequery1) SetAssignee(v ErrorTrackingIssueAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *Sourcequery1) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetFilterGroup

`func (o *Sourcequery1) GetFilterGroup() PropertyGroupFilter`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *Sourcequery1) GetFilterGroupOk() (*PropertyGroupFilter, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *Sourcequery1) SetFilterGroup(v PropertyGroupFilter)`

SetFilterGroup sets FilterGroup field to given value.


### GetGroupKey

`func (o *Sourcequery1) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *Sourcequery1) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *Sourcequery1) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *Sourcequery1) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### SetGroupKeyNil

`func (o *Sourcequery1) SetGroupKeyNil(b bool)`

 SetGroupKeyNil sets the value for GroupKey to be an explicit nil

### UnsetGroupKey
`func (o *Sourcequery1) UnsetGroupKey()`

UnsetGroupKey ensures that no value is present for GroupKey, not even an explicit nil
### GetGroupTypeIndex

`func (o *Sourcequery1) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Sourcequery1) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Sourcequery1) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *Sourcequery1) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *Sourcequery1) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Sourcequery1) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetIssueId

`func (o *Sourcequery1) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Sourcequery1) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Sourcequery1) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.


### GetOrderDirection

`func (o *Sourcequery1) GetOrderDirection() OrderDirection1`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *Sourcequery1) GetOrderDirectionOk() (*OrderDirection1, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *Sourcequery1) SetOrderDirection(v OrderDirection1)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *Sourcequery1) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetRevenueEntity

`func (o *Sourcequery1) GetRevenueEntity() RevenueEntity`

GetRevenueEntity returns the RevenueEntity field if non-nil, zero value otherwise.

### GetRevenueEntityOk

`func (o *Sourcequery1) GetRevenueEntityOk() (*RevenueEntity, bool)`

GetRevenueEntityOk returns a tuple with the RevenueEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueEntity

`func (o *Sourcequery1) SetRevenueEntity(v RevenueEntity)`

SetRevenueEntity sets RevenueEntity field to given value.

### HasRevenueEntity

`func (o *Sourcequery1) HasRevenueEntity() bool`

HasRevenueEntity returns a boolean if a field has been set.

### GetRevenuePeriod

`func (o *Sourcequery1) GetRevenuePeriod() RevenuePeriod`

GetRevenuePeriod returns the RevenuePeriod field if non-nil, zero value otherwise.

### GetRevenuePeriodOk

`func (o *Sourcequery1) GetRevenuePeriodOk() (*RevenuePeriod, bool)`

GetRevenuePeriodOk returns a tuple with the RevenuePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePeriod

`func (o *Sourcequery1) SetRevenuePeriod(v RevenuePeriod)`

SetRevenuePeriod sets RevenuePeriod field to given value.

### HasRevenuePeriod

`func (o *Sourcequery1) HasRevenuePeriod() bool`

HasRevenuePeriod returns a boolean if a field has been set.

### GetSearchQuery

`func (o *Sourcequery1) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *Sourcequery1) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *Sourcequery1) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *Sourcequery1) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### SetSearchQueryNil

`func (o *Sourcequery1) SetSearchQueryNil(b bool)`

 SetSearchQueryNil sets the value for SearchQuery to be an explicit nil

### UnsetSearchQuery
`func (o *Sourcequery1) UnsetSearchQuery()`

UnsetSearchQuery ensures that no value is present for SearchQuery, not even an explicit nil
### GetVolumeResolution

`func (o *Sourcequery1) GetVolumeResolution() int32`

GetVolumeResolution returns the VolumeResolution field if non-nil, zero value otherwise.

### GetVolumeResolutionOk

`func (o *Sourcequery1) GetVolumeResolutionOk() (*int32, bool)`

GetVolumeResolutionOk returns a tuple with the VolumeResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeResolution

`func (o *Sourcequery1) SetVolumeResolution(v int32)`

SetVolumeResolution sets VolumeResolution field to given value.


### GetWithAggregations

`func (o *Sourcequery1) GetWithAggregations() bool`

GetWithAggregations returns the WithAggregations field if non-nil, zero value otherwise.

### GetWithAggregationsOk

`func (o *Sourcequery1) GetWithAggregationsOk() (*bool, bool)`

GetWithAggregationsOk returns a tuple with the WithAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAggregations

`func (o *Sourcequery1) SetWithAggregations(v bool)`

SetWithAggregations sets WithAggregations field to given value.

### HasWithAggregations

`func (o *Sourcequery1) HasWithAggregations() bool`

HasWithAggregations returns a boolean if a field has been set.

### SetWithAggregationsNil

`func (o *Sourcequery1) SetWithAggregationsNil(b bool)`

 SetWithAggregationsNil sets the value for WithAggregations to be an explicit nil

### UnsetWithAggregations
`func (o *Sourcequery1) UnsetWithAggregations()`

UnsetWithAggregations ensures that no value is present for WithAggregations, not even an explicit nil
### GetWithFirstEvent

`func (o *Sourcequery1) GetWithFirstEvent() bool`

GetWithFirstEvent returns the WithFirstEvent field if non-nil, zero value otherwise.

### GetWithFirstEventOk

`func (o *Sourcequery1) GetWithFirstEventOk() (*bool, bool)`

GetWithFirstEventOk returns a tuple with the WithFirstEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithFirstEvent

`func (o *Sourcequery1) SetWithFirstEvent(v bool)`

SetWithFirstEvent sets WithFirstEvent field to given value.

### HasWithFirstEvent

`func (o *Sourcequery1) HasWithFirstEvent() bool`

HasWithFirstEvent returns a boolean if a field has been set.

### SetWithFirstEventNil

`func (o *Sourcequery1) SetWithFirstEventNil(b bool)`

 SetWithFirstEventNil sets the value for WithFirstEvent to be an explicit nil

### UnsetWithFirstEvent
`func (o *Sourcequery1) UnsetWithFirstEvent()`

UnsetWithFirstEvent ensures that no value is present for WithFirstEvent, not even an explicit nil
### GetWithLastEvent

`func (o *Sourcequery1) GetWithLastEvent() bool`

GetWithLastEvent returns the WithLastEvent field if non-nil, zero value otherwise.

### GetWithLastEventOk

`func (o *Sourcequery1) GetWithLastEventOk() (*bool, bool)`

GetWithLastEventOk returns a tuple with the WithLastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithLastEvent

`func (o *Sourcequery1) SetWithLastEvent(v bool)`

SetWithLastEvent sets WithLastEvent field to given value.

### HasWithLastEvent

`func (o *Sourcequery1) HasWithLastEvent() bool`

HasWithLastEvent returns a boolean if a field has been set.

### SetWithLastEventNil

`func (o *Sourcequery1) SetWithLastEventNil(b bool)`

 SetWithLastEventNil sets the value for WithLastEvent to be an explicit nil

### UnsetWithLastEvent
`func (o *Sourcequery1) UnsetWithLastEvent()`

UnsetWithLastEvent ensures that no value is present for WithLastEvent, not even an explicit nil
### GetMaxDistance

`func (o *Sourcequery1) GetMaxDistance() float32`

GetMaxDistance returns the MaxDistance field if non-nil, zero value otherwise.

### GetMaxDistanceOk

`func (o *Sourcequery1) GetMaxDistanceOk() (*float32, bool)`

GetMaxDistanceOk returns a tuple with the MaxDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistance

`func (o *Sourcequery1) SetMaxDistance(v float32)`

SetMaxDistance sets MaxDistance field to given value.

### HasMaxDistance

`func (o *Sourcequery1) HasMaxDistance() bool`

HasMaxDistance returns a boolean if a field has been set.

### SetMaxDistanceNil

`func (o *Sourcequery1) SetMaxDistanceNil(b bool)`

 SetMaxDistanceNil sets the value for MaxDistance to be an explicit nil

### UnsetMaxDistance
`func (o *Sourcequery1) UnsetMaxDistance()`

UnsetMaxDistance ensures that no value is present for MaxDistance, not even an explicit nil
### GetModelName

`func (o *Sourcequery1) GetModelName() EmbeddingModelName`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Sourcequery1) GetModelNameOk() (*EmbeddingModelName, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Sourcequery1) SetModelName(v EmbeddingModelName)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *Sourcequery1) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetRendering

`func (o *Sourcequery1) GetRendering() string`

GetRendering returns the Rendering field if non-nil, zero value otherwise.

### GetRenderingOk

`func (o *Sourcequery1) GetRenderingOk() (*string, bool)`

GetRenderingOk returns a tuple with the Rendering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendering

`func (o *Sourcequery1) SetRendering(v string)`

SetRendering sets Rendering field to given value.

### HasRendering

`func (o *Sourcequery1) HasRendering() bool`

HasRendering returns a boolean if a field has been set.

### SetRenderingNil

`func (o *Sourcequery1) SetRenderingNil(b bool)`

 SetRenderingNil sets the value for Rendering to be an explicit nil

### UnsetRendering
`func (o *Sourcequery1) UnsetRendering()`

UnsetRendering ensures that no value is present for Rendering, not even an explicit nil
### GetBreakdownProperties

`func (o *Sourcequery1) GetBreakdownProperties() []string`

GetBreakdownProperties returns the BreakdownProperties field if non-nil, zero value otherwise.

### GetBreakdownPropertiesOk

`func (o *Sourcequery1) GetBreakdownPropertiesOk() (*[]string, bool)`

GetBreakdownPropertiesOk returns a tuple with the BreakdownProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownProperties

`func (o *Sourcequery1) SetBreakdownProperties(v []string)`

SetBreakdownProperties sets BreakdownProperties field to given value.


### GetMaxValuesPerProperty

`func (o *Sourcequery1) GetMaxValuesPerProperty() int32`

GetMaxValuesPerProperty returns the MaxValuesPerProperty field if non-nil, zero value otherwise.

### GetMaxValuesPerPropertyOk

`func (o *Sourcequery1) GetMaxValuesPerPropertyOk() (*int32, bool)`

GetMaxValuesPerPropertyOk returns a tuple with the MaxValuesPerProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValuesPerProperty

`func (o *Sourcequery1) SetMaxValuesPerProperty(v int32)`

SetMaxValuesPerProperty sets MaxValuesPerProperty field to given value.

### HasMaxValuesPerProperty

`func (o *Sourcequery1) HasMaxValuesPerProperty() bool`

HasMaxValuesPerProperty returns a boolean if a field has been set.

### SetMaxValuesPerPropertyNil

`func (o *Sourcequery1) SetMaxValuesPerPropertyNil(b bool)`

 SetMaxValuesPerPropertyNil sets the value for MaxValuesPerProperty to be an explicit nil

### UnsetMaxValuesPerProperty
`func (o *Sourcequery1) UnsetMaxValuesPerProperty()`

UnsetMaxValuesPerProperty ensures that no value is present for MaxValuesPerProperty, not even an explicit nil
### GetEvents

`func (o *Sourcequery1) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Sourcequery1) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Sourcequery1) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.


### SetEventsNil

`func (o *Sourcequery1) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *Sourcequery1) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetServiceNames

`func (o *Sourcequery1) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *Sourcequery1) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *Sourcequery1) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.


### GetSeverityLevels

`func (o *Sourcequery1) GetSeverityLevels() []LogSeverityLevel`

GetSeverityLevels returns the SeverityLevels field if non-nil, zero value otherwise.

### GetSeverityLevelsOk

`func (o *Sourcequery1) GetSeverityLevelsOk() (*[]LogSeverityLevel, bool)`

GetSeverityLevelsOk returns a tuple with the SeverityLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevels

`func (o *Sourcequery1) SetSeverityLevels(v []LogSeverityLevel)`

SetSeverityLevels sets SeverityLevels field to given value.


### GetExperimentId

`func (o *Sourcequery1) GetExperimentId() int32`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *Sourcequery1) GetExperimentIdOk() (*int32, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *Sourcequery1) SetExperimentId(v int32)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *Sourcequery1) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### SetExperimentIdNil

`func (o *Sourcequery1) SetExperimentIdNil(b bool)`

 SetExperimentIdNil sets the value for ExperimentId to be an explicit nil

### UnsetExperimentId
`func (o *Sourcequery1) UnsetExperimentId()`

UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil
### GetFingerprint

`func (o *Sourcequery1) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Sourcequery1) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Sourcequery1) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Sourcequery1) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *Sourcequery1) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Sourcequery1) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetFunnelsQuery

`func (o *Sourcequery1) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *Sourcequery1) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *Sourcequery1) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.


### GetUuid

`func (o *Sourcequery1) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Sourcequery1) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Sourcequery1) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Sourcequery1) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Sourcequery1) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Sourcequery1) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetCountQuery

`func (o *Sourcequery1) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *Sourcequery1) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *Sourcequery1) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.


### GetExposureQuery

`func (o *Sourcequery1) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *Sourcequery1) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *Sourcequery1) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *Sourcequery1) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetAggregationGroupTypeIndex

`func (o *Sourcequery1) GetAggregationGroupTypeIndex() int32`

GetAggregationGroupTypeIndex returns the AggregationGroupTypeIndex field if non-nil, zero value otherwise.

### GetAggregationGroupTypeIndexOk

`func (o *Sourcequery1) GetAggregationGroupTypeIndexOk() (*int32, bool)`

GetAggregationGroupTypeIndexOk returns a tuple with the AggregationGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationGroupTypeIndex

`func (o *Sourcequery1) SetAggregationGroupTypeIndex(v int32)`

SetAggregationGroupTypeIndex sets AggregationGroupTypeIndex field to given value.

### HasAggregationGroupTypeIndex

`func (o *Sourcequery1) HasAggregationGroupTypeIndex() bool`

HasAggregationGroupTypeIndex returns a boolean if a field has been set.

### SetAggregationGroupTypeIndexNil

`func (o *Sourcequery1) SetAggregationGroupTypeIndexNil(b bool)`

 SetAggregationGroupTypeIndexNil sets the value for AggregationGroupTypeIndex to be an explicit nil

### UnsetAggregationGroupTypeIndex
`func (o *Sourcequery1) UnsetAggregationGroupTypeIndex()`

UnsetAggregationGroupTypeIndex ensures that no value is present for AggregationGroupTypeIndex, not even an explicit nil
### GetCalendarHeatmapFilter

`func (o *Sourcequery1) GetCalendarHeatmapFilter() CalendarHeatmapFilter`

GetCalendarHeatmapFilter returns the CalendarHeatmapFilter field if non-nil, zero value otherwise.

### GetCalendarHeatmapFilterOk

`func (o *Sourcequery1) GetCalendarHeatmapFilterOk() (*CalendarHeatmapFilter, bool)`

GetCalendarHeatmapFilterOk returns a tuple with the CalendarHeatmapFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarHeatmapFilter

`func (o *Sourcequery1) SetCalendarHeatmapFilter(v CalendarHeatmapFilter)`

SetCalendarHeatmapFilter sets CalendarHeatmapFilter field to given value.

### HasCalendarHeatmapFilter

`func (o *Sourcequery1) HasCalendarHeatmapFilter() bool`

HasCalendarHeatmapFilter returns a boolean if a field has been set.

### GetDataColorTheme

`func (o *Sourcequery1) GetDataColorTheme() float32`

GetDataColorTheme returns the DataColorTheme field if non-nil, zero value otherwise.

### GetDataColorThemeOk

`func (o *Sourcequery1) GetDataColorThemeOk() (*float32, bool)`

GetDataColorThemeOk returns a tuple with the DataColorTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorTheme

`func (o *Sourcequery1) SetDataColorTheme(v float32)`

SetDataColorTheme sets DataColorTheme field to given value.

### HasDataColorTheme

`func (o *Sourcequery1) HasDataColorTheme() bool`

HasDataColorTheme returns a boolean if a field has been set.

### SetDataColorThemeNil

`func (o *Sourcequery1) SetDataColorThemeNil(b bool)`

 SetDataColorThemeNil sets the value for DataColorTheme to be an explicit nil

### UnsetDataColorTheme
`func (o *Sourcequery1) UnsetDataColorTheme()`

UnsetDataColorTheme ensures that no value is present for DataColorTheme, not even an explicit nil
### GetSamplingFactor

`func (o *Sourcequery1) GetSamplingFactor() float32`

GetSamplingFactor returns the SamplingFactor field if non-nil, zero value otherwise.

### GetSamplingFactorOk

`func (o *Sourcequery1) GetSamplingFactorOk() (*float32, bool)`

GetSamplingFactorOk returns a tuple with the SamplingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingFactor

`func (o *Sourcequery1) SetSamplingFactor(v float32)`

SetSamplingFactor sets SamplingFactor field to given value.

### HasSamplingFactor

`func (o *Sourcequery1) HasSamplingFactor() bool`

HasSamplingFactor returns a boolean if a field has been set.

### SetSamplingFactorNil

`func (o *Sourcequery1) SetSamplingFactorNil(b bool)`

 SetSamplingFactorNil sets the value for SamplingFactor to be an explicit nil

### UnsetSamplingFactor
`func (o *Sourcequery1) UnsetSamplingFactor()`

UnsetSamplingFactor ensures that no value is present for SamplingFactor, not even an explicit nil
### GetActions

`func (o *Sourcequery1) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Sourcequery1) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Sourcequery1) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *Sourcequery1) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *Sourcequery1) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *Sourcequery1) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetCommentText

`func (o *Sourcequery1) GetCommentText() RecordingPropertyFilter`

GetCommentText returns the CommentText field if non-nil, zero value otherwise.

### GetCommentTextOk

`func (o *Sourcequery1) GetCommentTextOk() (*RecordingPropertyFilter, bool)`

GetCommentTextOk returns a tuple with the CommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentText

`func (o *Sourcequery1) SetCommentText(v RecordingPropertyFilter)`

SetCommentText sets CommentText field to given value.

### HasCommentText

`func (o *Sourcequery1) HasCommentText() bool`

HasCommentText returns a boolean if a field has been set.

### GetConsoleLogFilters

`func (o *Sourcequery1) GetConsoleLogFilters() []LogEntryPropertyFilter`

GetConsoleLogFilters returns the ConsoleLogFilters field if non-nil, zero value otherwise.

### GetConsoleLogFiltersOk

`func (o *Sourcequery1) GetConsoleLogFiltersOk() (*[]LogEntryPropertyFilter, bool)`

GetConsoleLogFiltersOk returns a tuple with the ConsoleLogFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogFilters

`func (o *Sourcequery1) SetConsoleLogFilters(v []LogEntryPropertyFilter)`

SetConsoleLogFilters sets ConsoleLogFilters field to given value.

### HasConsoleLogFilters

`func (o *Sourcequery1) HasConsoleLogFilters() bool`

HasConsoleLogFilters returns a boolean if a field has been set.

### SetConsoleLogFiltersNil

`func (o *Sourcequery1) SetConsoleLogFiltersNil(b bool)`

 SetConsoleLogFiltersNil sets the value for ConsoleLogFilters to be an explicit nil

### UnsetConsoleLogFilters
`func (o *Sourcequery1) UnsetConsoleLogFilters()`

UnsetConsoleLogFilters ensures that no value is present for ConsoleLogFilters, not even an explicit nil
### GetDateFrom

`func (o *Sourcequery1) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *Sourcequery1) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *Sourcequery1) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *Sourcequery1) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *Sourcequery1) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *Sourcequery1) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *Sourcequery1) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *Sourcequery1) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *Sourcequery1) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *Sourcequery1) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *Sourcequery1) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *Sourcequery1) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil
### GetDistinctIds

`func (o *Sourcequery1) GetDistinctIds() []string`

GetDistinctIds returns the DistinctIds field if non-nil, zero value otherwise.

### GetDistinctIdsOk

`func (o *Sourcequery1) GetDistinctIdsOk() (*[]string, bool)`

GetDistinctIdsOk returns a tuple with the DistinctIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIds

`func (o *Sourcequery1) SetDistinctIds(v []string)`

SetDistinctIds sets DistinctIds field to given value.

### HasDistinctIds

`func (o *Sourcequery1) HasDistinctIds() bool`

HasDistinctIds returns a boolean if a field has been set.

### SetDistinctIdsNil

`func (o *Sourcequery1) SetDistinctIdsNil(b bool)`

 SetDistinctIdsNil sets the value for DistinctIds to be an explicit nil

### UnsetDistinctIds
`func (o *Sourcequery1) UnsetDistinctIds()`

UnsetDistinctIds ensures that no value is present for DistinctIds, not even an explicit nil
### GetFilterTestAccounts

`func (o *Sourcequery1) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *Sourcequery1) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *Sourcequery1) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *Sourcequery1) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *Sourcequery1) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *Sourcequery1) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetHavingPredicates

`func (o *Sourcequery1) GetHavingPredicates() []FixedpropertiesInner`

GetHavingPredicates returns the HavingPredicates field if non-nil, zero value otherwise.

### GetHavingPredicatesOk

`func (o *Sourcequery1) GetHavingPredicatesOk() (*[]FixedpropertiesInner, bool)`

GetHavingPredicatesOk returns a tuple with the HavingPredicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHavingPredicates

`func (o *Sourcequery1) SetHavingPredicates(v []FixedpropertiesInner)`

SetHavingPredicates sets HavingPredicates field to given value.

### HasHavingPredicates

`func (o *Sourcequery1) HasHavingPredicates() bool`

HasHavingPredicates returns a boolean if a field has been set.

### SetHavingPredicatesNil

`func (o *Sourcequery1) SetHavingPredicatesNil(b bool)`

 SetHavingPredicatesNil sets the value for HavingPredicates to be an explicit nil

### UnsetHavingPredicates
`func (o *Sourcequery1) UnsetHavingPredicates()`

UnsetHavingPredicates ensures that no value is present for HavingPredicates, not even an explicit nil
### GetOperand

`func (o *Sourcequery1) GetOperand() FilterLogicalOperator`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *Sourcequery1) GetOperandOk() (*FilterLogicalOperator, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *Sourcequery1) SetOperand(v FilterLogicalOperator)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *Sourcequery1) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetOrder

`func (o *Sourcequery1) GetOrder() RecordingOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Sourcequery1) GetOrderOk() (*RecordingOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Sourcequery1) SetOrder(v RecordingOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Sourcequery1) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderDirection

`func (o *Sourcequery1) GetOrderDirection() RecordingOrderDirection`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *Sourcequery1) GetOrderDirectionOk() (*RecordingOrderDirection, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *Sourcequery1) SetOrderDirection(v RecordingOrderDirection)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *Sourcequery1) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetPersonUuid

`func (o *Sourcequery1) GetPersonUuid() string`

GetPersonUuid returns the PersonUuid field if non-nil, zero value otherwise.

### GetPersonUuidOk

`func (o *Sourcequery1) GetPersonUuidOk() (*string, bool)`

GetPersonUuidOk returns a tuple with the PersonUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonUuid

`func (o *Sourcequery1) SetPersonUuid(v string)`

SetPersonUuid sets PersonUuid field to given value.

### HasPersonUuid

`func (o *Sourcequery1) HasPersonUuid() bool`

HasPersonUuid returns a boolean if a field has been set.

### SetPersonUuidNil

`func (o *Sourcequery1) SetPersonUuidNil(b bool)`

 SetPersonUuidNil sets the value for PersonUuid to be an explicit nil

### UnsetPersonUuid
`func (o *Sourcequery1) UnsetPersonUuid()`

UnsetPersonUuid ensures that no value is present for PersonUuid, not even an explicit nil
### GetSessionIds

`func (o *Sourcequery1) GetSessionIds() []string`

GetSessionIds returns the SessionIds field if non-nil, zero value otherwise.

### GetSessionIdsOk

`func (o *Sourcequery1) GetSessionIdsOk() (*[]string, bool)`

GetSessionIdsOk returns a tuple with the SessionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIds

`func (o *Sourcequery1) SetSessionIds(v []string)`

SetSessionIds sets SessionIds field to given value.

### HasSessionIds

`func (o *Sourcequery1) HasSessionIds() bool`

HasSessionIds returns a boolean if a field has been set.

### SetSessionIdsNil

`func (o *Sourcequery1) SetSessionIdsNil(b bool)`

 SetSessionIdsNil sets the value for SessionIds to be an explicit nil

### UnsetSessionIds
`func (o *Sourcequery1) UnsetSessionIds()`

UnsetSessionIds ensures that no value is present for SessionIds, not even an explicit nil
### GetSessionRecordingId

`func (o *Sourcequery1) GetSessionRecordingId() string`

GetSessionRecordingId returns the SessionRecordingId field if non-nil, zero value otherwise.

### GetSessionRecordingIdOk

`func (o *Sourcequery1) GetSessionRecordingIdOk() (*string, bool)`

GetSessionRecordingIdOk returns a tuple with the SessionRecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingId

`func (o *Sourcequery1) SetSessionRecordingId(v string)`

SetSessionRecordingId sets SessionRecordingId field to given value.

### HasSessionRecordingId

`func (o *Sourcequery1) HasSessionRecordingId() bool`

HasSessionRecordingId returns a boolean if a field has been set.

### SetSessionRecordingIdNil

`func (o *Sourcequery1) SetSessionRecordingIdNil(b bool)`

 SetSessionRecordingIdNil sets the value for SessionRecordingId to be an explicit nil

### UnsetSessionRecordingId
`func (o *Sourcequery1) UnsetSessionRecordingId()`

UnsetSessionRecordingId ensures that no value is present for SessionRecordingId, not even an explicit nil
### GetUserModifiedFilters

`func (o *Sourcequery1) GetUserModifiedFilters() map[string]interface{}`

GetUserModifiedFilters returns the UserModifiedFilters field if non-nil, zero value otherwise.

### GetUserModifiedFiltersOk

`func (o *Sourcequery1) GetUserModifiedFiltersOk() (*map[string]interface{}, bool)`

GetUserModifiedFiltersOk returns a tuple with the UserModifiedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserModifiedFilters

`func (o *Sourcequery1) SetUserModifiedFilters(v map[string]interface{})`

SetUserModifiedFilters sets UserModifiedFilters field to given value.

### HasUserModifiedFilters

`func (o *Sourcequery1) HasUserModifiedFilters() bool`

HasUserModifiedFilters returns a boolean if a field has been set.

### SetUserModifiedFiltersNil

`func (o *Sourcequery1) SetUserModifiedFiltersNil(b bool)`

 SetUserModifiedFiltersNil sets the value for UserModifiedFilters to be an explicit nil

### UnsetUserModifiedFilters
`func (o *Sourcequery1) UnsetUserModifiedFilters()`

UnsetUserModifiedFilters ensures that no value is present for UserModifiedFilters, not even an explicit nil
### GetShowColumnConfigurator

`func (o *Sourcequery1) GetShowColumnConfigurator() bool`

GetShowColumnConfigurator returns the ShowColumnConfigurator field if non-nil, zero value otherwise.

### GetShowColumnConfiguratorOk

`func (o *Sourcequery1) GetShowColumnConfiguratorOk() (*bool, bool)`

GetShowColumnConfiguratorOk returns a tuple with the ShowColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowColumnConfigurator

`func (o *Sourcequery1) SetShowColumnConfigurator(v bool)`

SetShowColumnConfigurator sets ShowColumnConfigurator field to given value.

### HasShowColumnConfigurator

`func (o *Sourcequery1) HasShowColumnConfigurator() bool`

HasShowColumnConfigurator returns a boolean if a field has been set.

### SetShowColumnConfiguratorNil

`func (o *Sourcequery1) SetShowColumnConfiguratorNil(b bool)`

 SetShowColumnConfiguratorNil sets the value for ShowColumnConfigurator to be an explicit nil

### UnsetShowColumnConfigurator
`func (o *Sourcequery1) UnsetShowColumnConfigurator()`

UnsetShowColumnConfigurator ensures that no value is present for ShowColumnConfigurator, not even an explicit nil
### GetTraceId

`func (o *Sourcequery1) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *Sourcequery1) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *Sourcequery1) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetEmbedding

`func (o *Sourcequery1) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *Sourcequery1) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *Sourcequery1) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.


### GetEmbeddingVersion

`func (o *Sourcequery1) GetEmbeddingVersion() float32`

GetEmbeddingVersion returns the EmbeddingVersion field if non-nil, zero value otherwise.

### GetEmbeddingVersionOk

`func (o *Sourcequery1) GetEmbeddingVersionOk() (*float32, bool)`

GetEmbeddingVersionOk returns a tuple with the EmbeddingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingVersion

`func (o *Sourcequery1) SetEmbeddingVersion(v float32)`

SetEmbeddingVersion sets EmbeddingVersion field to given value.

### HasEmbeddingVersion

`func (o *Sourcequery1) HasEmbeddingVersion() bool`

HasEmbeddingVersion returns a boolean if a field has been set.

### SetEmbeddingVersionNil

`func (o *Sourcequery1) SetEmbeddingVersionNil(b bool)`

 SetEmbeddingVersionNil sets the value for EmbeddingVersion to be an explicit nil

### UnsetEmbeddingVersion
`func (o *Sourcequery1) UnsetEmbeddingVersion()`

UnsetEmbeddingVersion ensures that no value is present for EmbeddingVersion, not even an explicit nil
### GetGroupKey

`func (o *Sourcequery1) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *Sourcequery1) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *Sourcequery1) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *Sourcequery1) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### SetGroupKeyNil

`func (o *Sourcequery1) SetGroupKeyNil(b bool)`

 SetGroupKeyNil sets the value for GroupKey to be an explicit nil

### UnsetGroupKey
`func (o *Sourcequery1) UnsetGroupKey()`

UnsetGroupKey ensures that no value is present for GroupKey, not even an explicit nil
### GetPersonId

`func (o *Sourcequery1) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Sourcequery1) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Sourcequery1) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Sourcequery1) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *Sourcequery1) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *Sourcequery1) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


