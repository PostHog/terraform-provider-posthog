# Sourcequery

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

### NewSourcequery

`func NewSourcequery(orderBy OrderBy3, properties []FixedpropertiesInner, id int32, select_ []string, source Source3, groupTypeIndex NullableInt32, breakdown []RevenueAnalyticsBreakdown, interval IntervalType, series []Series1Inner, query string, language HogLanguage, endPosition int32, startPosition int32, dateRange DateRange, groupBy []SessionAttributionGroupBy, breakdownBy WebStatsBreakdown, metric WebVitalsMetric, percentile WebVitalsPercentile, thresholds []float32, metrics []WebTrendsMetric, filterGroup PropertyGroupFilter, issueId string, volumeResolution int32, breakdownProperties []string, events []map[string]interface{}, serviceNames []string, severityLevels []LogSeverityLevel, funnelsQuery FunnelsQuery, countQuery TrendsQuery, traceId string, embedding []float32, ) *Sourcequery`

NewSourcequery instantiates a new Sourcequery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcequeryWithDefaults

`func NewSourcequeryWithDefaults() *Sourcequery`

NewSourcequeryWithDefaults instantiates a new Sourcequery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *Sourcequery) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *Sourcequery) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *Sourcequery) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *Sourcequery) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *Sourcequery) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *Sourcequery) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetEvent

`func (o *Sourcequery) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Sourcequery) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Sourcequery) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Sourcequery) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *Sourcequery) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *Sourcequery) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetFixedProperties

`func (o *Sourcequery) GetFixedProperties() []FixedpropertiesInner1`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *Sourcequery) GetFixedPropertiesOk() (*[]FixedpropertiesInner1, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *Sourcequery) SetFixedProperties(v []FixedpropertiesInner1)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *Sourcequery) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *Sourcequery) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *Sourcequery) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *Sourcequery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Sourcequery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Sourcequery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Sourcequery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *Sourcequery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Sourcequery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Sourcequery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Sourcequery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Sourcequery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Sourcequery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMath

`func (o *Sourcequery) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *Sourcequery) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *Sourcequery) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *Sourcequery) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *Sourcequery) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *Sourcequery) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *Sourcequery) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *Sourcequery) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *Sourcequery) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *Sourcequery) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *Sourcequery) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *Sourcequery) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *Sourcequery) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *Sourcequery) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *Sourcequery) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *Sourcequery) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *Sourcequery) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *Sourcequery) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *Sourcequery) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *Sourcequery) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *Sourcequery) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *Sourcequery) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *Sourcequery) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *Sourcequery) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *Sourcequery) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *Sourcequery) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *Sourcequery) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *Sourcequery) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *Sourcequery) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *Sourcequery) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *Sourcequery) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *Sourcequery) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *Sourcequery) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *Sourcequery) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *Sourcequery) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *Sourcequery) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *Sourcequery) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *Sourcequery) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *Sourcequery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sourcequery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sourcequery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sourcequery) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Sourcequery) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Sourcequery) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *Sourcequery) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *Sourcequery) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *Sourcequery) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *Sourcequery) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *Sourcequery) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *Sourcequery) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetOrderBy

`func (o *Sourcequery) GetOrderBy() OrderBy3`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *Sourcequery) GetOrderByOk() (*OrderBy3, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *Sourcequery) SetOrderBy(v OrderBy3)`

SetOrderBy sets OrderBy field to given value.


### GetProperties

`func (o *Sourcequery) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Sourcequery) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Sourcequery) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.


### SetPropertiesNil

`func (o *Sourcequery) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Sourcequery) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *Sourcequery) GetResponse() UsageMetricsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Sourcequery) GetResponseOk() (*UsageMetricsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Sourcequery) SetResponse(v UsageMetricsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Sourcequery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetVersion

`func (o *Sourcequery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Sourcequery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Sourcequery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Sourcequery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Sourcequery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Sourcequery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *Sourcequery) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Sourcequery) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Sourcequery) SetId(v int32)`

SetId sets Id field to given value.


### GetCohort

`func (o *Sourcequery) GetCohort() int32`

GetCohort returns the Cohort field if non-nil, zero value otherwise.

### GetCohortOk

`func (o *Sourcequery) GetCohortOk() (*int32, bool)`

GetCohortOk returns a tuple with the Cohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohort

`func (o *Sourcequery) SetCohort(v int32)`

SetCohort sets Cohort field to given value.

### HasCohort

`func (o *Sourcequery) HasCohort() bool`

HasCohort returns a boolean if a field has been set.

### SetCohortNil

`func (o *Sourcequery) SetCohortNil(b bool)`

 SetCohortNil sets the value for Cohort to be an explicit nil

### UnsetCohort
`func (o *Sourcequery) UnsetCohort()`

UnsetCohort ensures that no value is present for Cohort, not even an explicit nil
### GetDistinctId

`func (o *Sourcequery) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *Sourcequery) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *Sourcequery) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *Sourcequery) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### SetDistinctIdNil

`func (o *Sourcequery) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *Sourcequery) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetModifiers

`func (o *Sourcequery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Sourcequery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Sourcequery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Sourcequery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *Sourcequery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Sourcequery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Sourcequery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Sourcequery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *Sourcequery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *Sourcequery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSearch

`func (o *Sourcequery) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *Sourcequery) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *Sourcequery) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *Sourcequery) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *Sourcequery) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *Sourcequery) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetTags

`func (o *Sourcequery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Sourcequery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Sourcequery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Sourcequery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetActionId

`func (o *Sourcequery) GetActionId() int32`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *Sourcequery) GetActionIdOk() (*int32, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *Sourcequery) SetActionId(v int32)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *Sourcequery) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### SetActionIdNil

`func (o *Sourcequery) SetActionIdNil(b bool)`

 SetActionIdNil sets the value for ActionId to be an explicit nil

### UnsetActionId
`func (o *Sourcequery) UnsetActionId()`

UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
### GetAfter

`func (o *Sourcequery) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Sourcequery) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Sourcequery) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *Sourcequery) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *Sourcequery) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *Sourcequery) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *Sourcequery) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Sourcequery) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Sourcequery) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *Sourcequery) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *Sourcequery) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *Sourcequery) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetFilterTestAccounts

`func (o *Sourcequery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *Sourcequery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *Sourcequery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *Sourcequery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *Sourcequery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *Sourcequery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetPersonId

`func (o *Sourcequery) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Sourcequery) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Sourcequery) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Sourcequery) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *Sourcequery) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *Sourcequery) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetSelect

`func (o *Sourcequery) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *Sourcequery) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *Sourcequery) SetSelect(v []string)`

SetSelect sets Select field to given value.


### SetSelectNil

`func (o *Sourcequery) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *Sourcequery) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetSource

`func (o *Sourcequery) GetSource() Source3`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Sourcequery) GetSourceOk() (*Source3, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Sourcequery) SetSource(v Source3)`

SetSource sets Source field to given value.


### GetWhere

`func (o *Sourcequery) GetWhere() []string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *Sourcequery) GetWhereOk() (*[]string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *Sourcequery) SetWhere(v []string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *Sourcequery) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### SetWhereNil

`func (o *Sourcequery) SetWhereNil(b bool)`

 SetWhereNil sets the value for Where to be an explicit nil

### UnsetWhere
`func (o *Sourcequery) UnsetWhere()`

UnsetWhere ensures that no value is present for Where, not even an explicit nil
### GetGroupTypeIndex

`func (o *Sourcequery) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Sourcequery) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Sourcequery) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.


### SetGroupTypeIndexNil

`func (o *Sourcequery) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Sourcequery) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetBreakdown

`func (o *Sourcequery) GetBreakdown() []RevenueAnalyticsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Sourcequery) GetBreakdownOk() (*[]RevenueAnalyticsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Sourcequery) SetBreakdown(v []RevenueAnalyticsBreakdown)`

SetBreakdown sets Breakdown field to given value.


### GetCompare

`func (o *Sourcequery) GetCompare() Compare`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *Sourcequery) GetCompareOk() (*Compare, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *Sourcequery) SetCompare(v Compare)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *Sourcequery) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### GetDay

`func (o *Sourcequery) GetDay() Day`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Sourcequery) GetDayOk() (*Day, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Sourcequery) SetDay(v Day)`

SetDay sets Day field to given value.

### HasDay

`func (o *Sourcequery) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *Sourcequery) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *Sourcequery) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetIncludeRecordings

`func (o *Sourcequery) GetIncludeRecordings() bool`

GetIncludeRecordings returns the IncludeRecordings field if non-nil, zero value otherwise.

### GetIncludeRecordingsOk

`func (o *Sourcequery) GetIncludeRecordingsOk() (*bool, bool)`

GetIncludeRecordingsOk returns a tuple with the IncludeRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordings

`func (o *Sourcequery) SetIncludeRecordings(v bool)`

SetIncludeRecordings sets IncludeRecordings field to given value.

### HasIncludeRecordings

`func (o *Sourcequery) HasIncludeRecordings() bool`

HasIncludeRecordings returns a boolean if a field has been set.

### SetIncludeRecordingsNil

`func (o *Sourcequery) SetIncludeRecordingsNil(b bool)`

 SetIncludeRecordingsNil sets the value for IncludeRecordings to be an explicit nil

### UnsetIncludeRecordings
`func (o *Sourcequery) UnsetIncludeRecordings()`

UnsetIncludeRecordings ensures that no value is present for IncludeRecordings, not even an explicit nil
### GetInterval

`func (o *Sourcequery) GetInterval() IntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Sourcequery) GetIntervalOk() (*IntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Sourcequery) SetInterval(v IntervalType)`

SetInterval sets Interval field to given value.


### GetSeries

`func (o *Sourcequery) GetSeries() []Series1Inner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Sourcequery) GetSeriesOk() (*[]Series1Inner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Sourcequery) SetSeries(v []Series1Inner)`

SetSeries sets Series field to given value.


### GetStatus

`func (o *Sourcequery) GetStatus() Status2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Sourcequery) GetStatusOk() (*Status2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Sourcequery) SetStatus(v Status2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Sourcequery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *Sourcequery) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Sourcequery) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Sourcequery) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Sourcequery) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Sourcequery) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Sourcequery) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetExplain

`func (o *Sourcequery) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *Sourcequery) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *Sourcequery) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *Sourcequery) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *Sourcequery) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *Sourcequery) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetFilters

`func (o *Sourcequery) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Sourcequery) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Sourcequery) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Sourcequery) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetQuery

`func (o *Sourcequery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Sourcequery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Sourcequery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetValues

`func (o *Sourcequery) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Sourcequery) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Sourcequery) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *Sourcequery) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *Sourcequery) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *Sourcequery) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetVariables

`func (o *Sourcequery) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Sourcequery) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Sourcequery) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Sourcequery) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *Sourcequery) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Sourcequery) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetDebug

`func (o *Sourcequery) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *Sourcequery) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *Sourcequery) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *Sourcequery) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### SetDebugNil

`func (o *Sourcequery) SetDebugNil(b bool)`

 SetDebugNil sets the value for Debug to be an explicit nil

### UnsetDebug
`func (o *Sourcequery) UnsetDebug()`

UnsetDebug ensures that no value is present for Debug, not even an explicit nil
### GetGlobals

`func (o *Sourcequery) GetGlobals() map[string]interface{}`

GetGlobals returns the Globals field if non-nil, zero value otherwise.

### GetGlobalsOk

`func (o *Sourcequery) GetGlobalsOk() (*map[string]interface{}, bool)`

GetGlobalsOk returns a tuple with the Globals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobals

`func (o *Sourcequery) SetGlobals(v map[string]interface{})`

SetGlobals sets Globals field to given value.

### HasGlobals

`func (o *Sourcequery) HasGlobals() bool`

HasGlobals returns a boolean if a field has been set.

### SetGlobalsNil

`func (o *Sourcequery) SetGlobalsNil(b bool)`

 SetGlobalsNil sets the value for Globals to be an explicit nil

### UnsetGlobals
`func (o *Sourcequery) UnsetGlobals()`

UnsetGlobals ensures that no value is present for Globals, not even an explicit nil
### GetLanguage

`func (o *Sourcequery) GetLanguage() HogLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Sourcequery) GetLanguageOk() (*HogLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Sourcequery) SetLanguage(v HogLanguage)`

SetLanguage sets Language field to given value.


### GetSourceQuery

`func (o *Sourcequery) GetSourceQuery() Sourcequery`

GetSourceQuery returns the SourceQuery field if non-nil, zero value otherwise.

### GetSourceQueryOk

`func (o *Sourcequery) GetSourceQueryOk() (*Sourcequery, bool)`

GetSourceQueryOk returns a tuple with the SourceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuery

`func (o *Sourcequery) SetSourceQuery(v Sourcequery)`

SetSourceQuery sets SourceQuery field to given value.

### HasSourceQuery

`func (o *Sourcequery) HasSourceQuery() bool`

HasSourceQuery returns a boolean if a field has been set.

### SetSourceQueryNil

`func (o *Sourcequery) SetSourceQueryNil(b bool)`

 SetSourceQueryNil sets the value for SourceQuery to be an explicit nil

### UnsetSourceQuery
`func (o *Sourcequery) UnsetSourceQuery()`

UnsetSourceQuery ensures that no value is present for SourceQuery, not even an explicit nil
### GetEndPosition

`func (o *Sourcequery) GetEndPosition() int32`

GetEndPosition returns the EndPosition field if non-nil, zero value otherwise.

### GetEndPositionOk

`func (o *Sourcequery) GetEndPositionOk() (*int32, bool)`

GetEndPositionOk returns a tuple with the EndPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPosition

`func (o *Sourcequery) SetEndPosition(v int32)`

SetEndPosition sets EndPosition field to given value.


### GetStartPosition

`func (o *Sourcequery) GetStartPosition() int32`

GetStartPosition returns the StartPosition field if non-nil, zero value otherwise.

### GetStartPositionOk

`func (o *Sourcequery) GetStartPositionOk() (*int32, bool)`

GetStartPositionOk returns a tuple with the StartPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPosition

`func (o *Sourcequery) SetStartPosition(v int32)`

SetStartPosition sets StartPosition field to given value.


### GetDateRange

`func (o *Sourcequery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Sourcequery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Sourcequery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetGroupBy

`func (o *Sourcequery) GetGroupBy() []SessionAttributionGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Sourcequery) GetGroupByOk() (*[]SessionAttributionGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Sourcequery) SetGroupBy(v []SessionAttributionGroupBy)`

SetGroupBy sets GroupBy field to given value.


### GetCompareFilter

`func (o *Sourcequery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *Sourcequery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *Sourcequery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *Sourcequery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *Sourcequery) GetConversionGoal() Conversiongoal`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *Sourcequery) GetConversionGoalOk() (*Conversiongoal, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *Sourcequery) SetConversionGoal(v Conversiongoal)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *Sourcequery) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *Sourcequery) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *Sourcequery) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDoPathCleaning

`func (o *Sourcequery) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *Sourcequery) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *Sourcequery) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *Sourcequery) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *Sourcequery) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *Sourcequery) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetDraftConversionGoal

`func (o *Sourcequery) GetDraftConversionGoal() Draftconversiongoal`

GetDraftConversionGoal returns the DraftConversionGoal field if non-nil, zero value otherwise.

### GetDraftConversionGoalOk

`func (o *Sourcequery) GetDraftConversionGoalOk() (*Draftconversiongoal, bool)`

GetDraftConversionGoalOk returns a tuple with the DraftConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftConversionGoal

`func (o *Sourcequery) SetDraftConversionGoal(v Draftconversiongoal)`

SetDraftConversionGoal sets DraftConversionGoal field to given value.

### HasDraftConversionGoal

`func (o *Sourcequery) HasDraftConversionGoal() bool`

HasDraftConversionGoal returns a boolean if a field has been set.

### SetDraftConversionGoalNil

`func (o *Sourcequery) SetDraftConversionGoalNil(b bool)`

 SetDraftConversionGoalNil sets the value for DraftConversionGoal to be an explicit nil

### UnsetDraftConversionGoal
`func (o *Sourcequery) UnsetDraftConversionGoal()`

UnsetDraftConversionGoal ensures that no value is present for DraftConversionGoal, not even an explicit nil
### GetIncludeAllConversions

`func (o *Sourcequery) GetIncludeAllConversions() bool`

GetIncludeAllConversions returns the IncludeAllConversions field if non-nil, zero value otherwise.

### GetIncludeAllConversionsOk

`func (o *Sourcequery) GetIncludeAllConversionsOk() (*bool, bool)`

GetIncludeAllConversionsOk returns a tuple with the IncludeAllConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllConversions

`func (o *Sourcequery) SetIncludeAllConversions(v bool)`

SetIncludeAllConversions sets IncludeAllConversions field to given value.

### HasIncludeAllConversions

`func (o *Sourcequery) HasIncludeAllConversions() bool`

HasIncludeAllConversions returns a boolean if a field has been set.

### SetIncludeAllConversionsNil

`func (o *Sourcequery) SetIncludeAllConversionsNil(b bool)`

 SetIncludeAllConversionsNil sets the value for IncludeAllConversions to be an explicit nil

### UnsetIncludeAllConversions
`func (o *Sourcequery) UnsetIncludeAllConversions()`

UnsetIncludeAllConversions ensures that no value is present for IncludeAllConversions, not even an explicit nil
### GetIncludeRevenue

`func (o *Sourcequery) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *Sourcequery) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *Sourcequery) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *Sourcequery) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *Sourcequery) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *Sourcequery) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetIntegrationFilter

`func (o *Sourcequery) GetIntegrationFilter() IntegrationFilter`

GetIntegrationFilter returns the IntegrationFilter field if non-nil, zero value otherwise.

### GetIntegrationFilterOk

`func (o *Sourcequery) GetIntegrationFilterOk() (*IntegrationFilter, bool)`

GetIntegrationFilterOk returns a tuple with the IntegrationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationFilter

`func (o *Sourcequery) SetIntegrationFilter(v IntegrationFilter)`

SetIntegrationFilter sets IntegrationFilter field to given value.

### HasIntegrationFilter

`func (o *Sourcequery) HasIntegrationFilter() bool`

HasIntegrationFilter returns a boolean if a field has been set.

### GetSampling

`func (o *Sourcequery) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *Sourcequery) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *Sourcequery) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *Sourcequery) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *Sourcequery) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *Sourcequery) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *Sourcequery) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *Sourcequery) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *Sourcequery) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *Sourcequery) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetBreakdownBy

`func (o *Sourcequery) GetBreakdownBy() WebStatsBreakdown`

GetBreakdownBy returns the BreakdownBy field if non-nil, zero value otherwise.

### GetBreakdownByOk

`func (o *Sourcequery) GetBreakdownByOk() (*WebStatsBreakdown, bool)`

GetBreakdownByOk returns a tuple with the BreakdownBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownBy

`func (o *Sourcequery) SetBreakdownBy(v WebStatsBreakdown)`

SetBreakdownBy sets BreakdownBy field to given value.


### GetIncludeBounceRate

`func (o *Sourcequery) GetIncludeBounceRate() bool`

GetIncludeBounceRate returns the IncludeBounceRate field if non-nil, zero value otherwise.

### GetIncludeBounceRateOk

`func (o *Sourcequery) GetIncludeBounceRateOk() (*bool, bool)`

GetIncludeBounceRateOk returns a tuple with the IncludeBounceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBounceRate

`func (o *Sourcequery) SetIncludeBounceRate(v bool)`

SetIncludeBounceRate sets IncludeBounceRate field to given value.

### HasIncludeBounceRate

`func (o *Sourcequery) HasIncludeBounceRate() bool`

HasIncludeBounceRate returns a boolean if a field has been set.

### SetIncludeBounceRateNil

`func (o *Sourcequery) SetIncludeBounceRateNil(b bool)`

 SetIncludeBounceRateNil sets the value for IncludeBounceRate to be an explicit nil

### UnsetIncludeBounceRate
`func (o *Sourcequery) UnsetIncludeBounceRate()`

UnsetIncludeBounceRate ensures that no value is present for IncludeBounceRate, not even an explicit nil
### GetIncludeScrollDepth

`func (o *Sourcequery) GetIncludeScrollDepth() bool`

GetIncludeScrollDepth returns the IncludeScrollDepth field if non-nil, zero value otherwise.

### GetIncludeScrollDepthOk

`func (o *Sourcequery) GetIncludeScrollDepthOk() (*bool, bool)`

GetIncludeScrollDepthOk returns a tuple with the IncludeScrollDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeScrollDepth

`func (o *Sourcequery) SetIncludeScrollDepth(v bool)`

SetIncludeScrollDepth sets IncludeScrollDepth field to given value.

### HasIncludeScrollDepth

`func (o *Sourcequery) HasIncludeScrollDepth() bool`

HasIncludeScrollDepth returns a boolean if a field has been set.

### SetIncludeScrollDepthNil

`func (o *Sourcequery) SetIncludeScrollDepthNil(b bool)`

 SetIncludeScrollDepthNil sets the value for IncludeScrollDepth to be an explicit nil

### UnsetIncludeScrollDepth
`func (o *Sourcequery) UnsetIncludeScrollDepth()`

UnsetIncludeScrollDepth ensures that no value is present for IncludeScrollDepth, not even an explicit nil
### GetStripQueryParams

`func (o *Sourcequery) GetStripQueryParams() bool`

GetStripQueryParams returns the StripQueryParams field if non-nil, zero value otherwise.

### GetStripQueryParamsOk

`func (o *Sourcequery) GetStripQueryParamsOk() (*bool, bool)`

GetStripQueryParamsOk returns a tuple with the StripQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripQueryParams

`func (o *Sourcequery) SetStripQueryParams(v bool)`

SetStripQueryParams sets StripQueryParams field to given value.

### HasStripQueryParams

`func (o *Sourcequery) HasStripQueryParams() bool`

HasStripQueryParams returns a boolean if a field has been set.

### SetStripQueryParamsNil

`func (o *Sourcequery) SetStripQueryParamsNil(b bool)`

 SetStripQueryParamsNil sets the value for StripQueryParams to be an explicit nil

### UnsetStripQueryParams
`func (o *Sourcequery) UnsetStripQueryParams()`

UnsetStripQueryParams ensures that no value is present for StripQueryParams, not even an explicit nil
### GetMetric

`func (o *Sourcequery) GetMetric() WebVitalsMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Sourcequery) GetMetricOk() (*WebVitalsMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Sourcequery) SetMetric(v WebVitalsMetric)`

SetMetric sets Metric field to given value.


### GetPercentile

`func (o *Sourcequery) GetPercentile() WebVitalsPercentile`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *Sourcequery) GetPercentileOk() (*WebVitalsPercentile, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *Sourcequery) SetPercentile(v WebVitalsPercentile)`

SetPercentile sets Percentile field to given value.


### GetThresholds

`func (o *Sourcequery) GetThresholds() []float32`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Sourcequery) GetThresholdsOk() (*[]float32, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Sourcequery) SetThresholds(v []float32)`

SetThresholds sets Thresholds field to given value.


### GetSearchTerm

`func (o *Sourcequery) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *Sourcequery) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *Sourcequery) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *Sourcequery) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### SetSearchTermNil

`func (o *Sourcequery) SetSearchTermNil(b bool)`

 SetSearchTermNil sets the value for SearchTerm to be an explicit nil

### UnsetSearchTerm
`func (o *Sourcequery) UnsetSearchTerm()`

UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil
### GetMetrics

`func (o *Sourcequery) GetMetrics() []WebTrendsMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Sourcequery) GetMetricsOk() (*[]WebTrendsMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Sourcequery) SetMetrics(v []WebTrendsMetric)`

SetMetrics sets Metrics field to given value.


### GetAssignee

`func (o *Sourcequery) GetAssignee() ErrorTrackingIssueAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Sourcequery) GetAssigneeOk() (*ErrorTrackingIssueAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Sourcequery) SetAssignee(v ErrorTrackingIssueAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *Sourcequery) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetFilterGroup

`func (o *Sourcequery) GetFilterGroup() PropertyGroupFilter`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *Sourcequery) GetFilterGroupOk() (*PropertyGroupFilter, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *Sourcequery) SetFilterGroup(v PropertyGroupFilter)`

SetFilterGroup sets FilterGroup field to given value.


### GetGroupKey

`func (o *Sourcequery) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *Sourcequery) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *Sourcequery) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *Sourcequery) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### SetGroupKeyNil

`func (o *Sourcequery) SetGroupKeyNil(b bool)`

 SetGroupKeyNil sets the value for GroupKey to be an explicit nil

### UnsetGroupKey
`func (o *Sourcequery) UnsetGroupKey()`

UnsetGroupKey ensures that no value is present for GroupKey, not even an explicit nil
### GetGroupTypeIndex

`func (o *Sourcequery) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Sourcequery) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Sourcequery) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *Sourcequery) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *Sourcequery) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Sourcequery) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetIssueId

`func (o *Sourcequery) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Sourcequery) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Sourcequery) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.


### GetOrderDirection

`func (o *Sourcequery) GetOrderDirection() OrderDirection1`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *Sourcequery) GetOrderDirectionOk() (*OrderDirection1, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *Sourcequery) SetOrderDirection(v OrderDirection1)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *Sourcequery) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetRevenueEntity

`func (o *Sourcequery) GetRevenueEntity() RevenueEntity`

GetRevenueEntity returns the RevenueEntity field if non-nil, zero value otherwise.

### GetRevenueEntityOk

`func (o *Sourcequery) GetRevenueEntityOk() (*RevenueEntity, bool)`

GetRevenueEntityOk returns a tuple with the RevenueEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueEntity

`func (o *Sourcequery) SetRevenueEntity(v RevenueEntity)`

SetRevenueEntity sets RevenueEntity field to given value.

### HasRevenueEntity

`func (o *Sourcequery) HasRevenueEntity() bool`

HasRevenueEntity returns a boolean if a field has been set.

### GetRevenuePeriod

`func (o *Sourcequery) GetRevenuePeriod() RevenuePeriod`

GetRevenuePeriod returns the RevenuePeriod field if non-nil, zero value otherwise.

### GetRevenuePeriodOk

`func (o *Sourcequery) GetRevenuePeriodOk() (*RevenuePeriod, bool)`

GetRevenuePeriodOk returns a tuple with the RevenuePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePeriod

`func (o *Sourcequery) SetRevenuePeriod(v RevenuePeriod)`

SetRevenuePeriod sets RevenuePeriod field to given value.

### HasRevenuePeriod

`func (o *Sourcequery) HasRevenuePeriod() bool`

HasRevenuePeriod returns a boolean if a field has been set.

### GetSearchQuery

`func (o *Sourcequery) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *Sourcequery) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *Sourcequery) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *Sourcequery) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### SetSearchQueryNil

`func (o *Sourcequery) SetSearchQueryNil(b bool)`

 SetSearchQueryNil sets the value for SearchQuery to be an explicit nil

### UnsetSearchQuery
`func (o *Sourcequery) UnsetSearchQuery()`

UnsetSearchQuery ensures that no value is present for SearchQuery, not even an explicit nil
### GetVolumeResolution

`func (o *Sourcequery) GetVolumeResolution() int32`

GetVolumeResolution returns the VolumeResolution field if non-nil, zero value otherwise.

### GetVolumeResolutionOk

`func (o *Sourcequery) GetVolumeResolutionOk() (*int32, bool)`

GetVolumeResolutionOk returns a tuple with the VolumeResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeResolution

`func (o *Sourcequery) SetVolumeResolution(v int32)`

SetVolumeResolution sets VolumeResolution field to given value.


### GetWithAggregations

`func (o *Sourcequery) GetWithAggregations() bool`

GetWithAggregations returns the WithAggregations field if non-nil, zero value otherwise.

### GetWithAggregationsOk

`func (o *Sourcequery) GetWithAggregationsOk() (*bool, bool)`

GetWithAggregationsOk returns a tuple with the WithAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAggregations

`func (o *Sourcequery) SetWithAggregations(v bool)`

SetWithAggregations sets WithAggregations field to given value.

### HasWithAggregations

`func (o *Sourcequery) HasWithAggregations() bool`

HasWithAggregations returns a boolean if a field has been set.

### SetWithAggregationsNil

`func (o *Sourcequery) SetWithAggregationsNil(b bool)`

 SetWithAggregationsNil sets the value for WithAggregations to be an explicit nil

### UnsetWithAggregations
`func (o *Sourcequery) UnsetWithAggregations()`

UnsetWithAggregations ensures that no value is present for WithAggregations, not even an explicit nil
### GetWithFirstEvent

`func (o *Sourcequery) GetWithFirstEvent() bool`

GetWithFirstEvent returns the WithFirstEvent field if non-nil, zero value otherwise.

### GetWithFirstEventOk

`func (o *Sourcequery) GetWithFirstEventOk() (*bool, bool)`

GetWithFirstEventOk returns a tuple with the WithFirstEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithFirstEvent

`func (o *Sourcequery) SetWithFirstEvent(v bool)`

SetWithFirstEvent sets WithFirstEvent field to given value.

### HasWithFirstEvent

`func (o *Sourcequery) HasWithFirstEvent() bool`

HasWithFirstEvent returns a boolean if a field has been set.

### SetWithFirstEventNil

`func (o *Sourcequery) SetWithFirstEventNil(b bool)`

 SetWithFirstEventNil sets the value for WithFirstEvent to be an explicit nil

### UnsetWithFirstEvent
`func (o *Sourcequery) UnsetWithFirstEvent()`

UnsetWithFirstEvent ensures that no value is present for WithFirstEvent, not even an explicit nil
### GetWithLastEvent

`func (o *Sourcequery) GetWithLastEvent() bool`

GetWithLastEvent returns the WithLastEvent field if non-nil, zero value otherwise.

### GetWithLastEventOk

`func (o *Sourcequery) GetWithLastEventOk() (*bool, bool)`

GetWithLastEventOk returns a tuple with the WithLastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithLastEvent

`func (o *Sourcequery) SetWithLastEvent(v bool)`

SetWithLastEvent sets WithLastEvent field to given value.

### HasWithLastEvent

`func (o *Sourcequery) HasWithLastEvent() bool`

HasWithLastEvent returns a boolean if a field has been set.

### SetWithLastEventNil

`func (o *Sourcequery) SetWithLastEventNil(b bool)`

 SetWithLastEventNil sets the value for WithLastEvent to be an explicit nil

### UnsetWithLastEvent
`func (o *Sourcequery) UnsetWithLastEvent()`

UnsetWithLastEvent ensures that no value is present for WithLastEvent, not even an explicit nil
### GetMaxDistance

`func (o *Sourcequery) GetMaxDistance() float32`

GetMaxDistance returns the MaxDistance field if non-nil, zero value otherwise.

### GetMaxDistanceOk

`func (o *Sourcequery) GetMaxDistanceOk() (*float32, bool)`

GetMaxDistanceOk returns a tuple with the MaxDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistance

`func (o *Sourcequery) SetMaxDistance(v float32)`

SetMaxDistance sets MaxDistance field to given value.

### HasMaxDistance

`func (o *Sourcequery) HasMaxDistance() bool`

HasMaxDistance returns a boolean if a field has been set.

### SetMaxDistanceNil

`func (o *Sourcequery) SetMaxDistanceNil(b bool)`

 SetMaxDistanceNil sets the value for MaxDistance to be an explicit nil

### UnsetMaxDistance
`func (o *Sourcequery) UnsetMaxDistance()`

UnsetMaxDistance ensures that no value is present for MaxDistance, not even an explicit nil
### GetModelName

`func (o *Sourcequery) GetModelName() EmbeddingModelName`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Sourcequery) GetModelNameOk() (*EmbeddingModelName, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Sourcequery) SetModelName(v EmbeddingModelName)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *Sourcequery) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetRendering

`func (o *Sourcequery) GetRendering() string`

GetRendering returns the Rendering field if non-nil, zero value otherwise.

### GetRenderingOk

`func (o *Sourcequery) GetRenderingOk() (*string, bool)`

GetRenderingOk returns a tuple with the Rendering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendering

`func (o *Sourcequery) SetRendering(v string)`

SetRendering sets Rendering field to given value.

### HasRendering

`func (o *Sourcequery) HasRendering() bool`

HasRendering returns a boolean if a field has been set.

### SetRenderingNil

`func (o *Sourcequery) SetRenderingNil(b bool)`

 SetRenderingNil sets the value for Rendering to be an explicit nil

### UnsetRendering
`func (o *Sourcequery) UnsetRendering()`

UnsetRendering ensures that no value is present for Rendering, not even an explicit nil
### GetBreakdownProperties

`func (o *Sourcequery) GetBreakdownProperties() []string`

GetBreakdownProperties returns the BreakdownProperties field if non-nil, zero value otherwise.

### GetBreakdownPropertiesOk

`func (o *Sourcequery) GetBreakdownPropertiesOk() (*[]string, bool)`

GetBreakdownPropertiesOk returns a tuple with the BreakdownProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownProperties

`func (o *Sourcequery) SetBreakdownProperties(v []string)`

SetBreakdownProperties sets BreakdownProperties field to given value.


### GetMaxValuesPerProperty

`func (o *Sourcequery) GetMaxValuesPerProperty() int32`

GetMaxValuesPerProperty returns the MaxValuesPerProperty field if non-nil, zero value otherwise.

### GetMaxValuesPerPropertyOk

`func (o *Sourcequery) GetMaxValuesPerPropertyOk() (*int32, bool)`

GetMaxValuesPerPropertyOk returns a tuple with the MaxValuesPerProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValuesPerProperty

`func (o *Sourcequery) SetMaxValuesPerProperty(v int32)`

SetMaxValuesPerProperty sets MaxValuesPerProperty field to given value.

### HasMaxValuesPerProperty

`func (o *Sourcequery) HasMaxValuesPerProperty() bool`

HasMaxValuesPerProperty returns a boolean if a field has been set.

### SetMaxValuesPerPropertyNil

`func (o *Sourcequery) SetMaxValuesPerPropertyNil(b bool)`

 SetMaxValuesPerPropertyNil sets the value for MaxValuesPerProperty to be an explicit nil

### UnsetMaxValuesPerProperty
`func (o *Sourcequery) UnsetMaxValuesPerProperty()`

UnsetMaxValuesPerProperty ensures that no value is present for MaxValuesPerProperty, not even an explicit nil
### GetEvents

`func (o *Sourcequery) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Sourcequery) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Sourcequery) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.


### SetEventsNil

`func (o *Sourcequery) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *Sourcequery) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetServiceNames

`func (o *Sourcequery) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *Sourcequery) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *Sourcequery) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.


### GetSeverityLevels

`func (o *Sourcequery) GetSeverityLevels() []LogSeverityLevel`

GetSeverityLevels returns the SeverityLevels field if non-nil, zero value otherwise.

### GetSeverityLevelsOk

`func (o *Sourcequery) GetSeverityLevelsOk() (*[]LogSeverityLevel, bool)`

GetSeverityLevelsOk returns a tuple with the SeverityLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevels

`func (o *Sourcequery) SetSeverityLevels(v []LogSeverityLevel)`

SetSeverityLevels sets SeverityLevels field to given value.


### GetExperimentId

`func (o *Sourcequery) GetExperimentId() int32`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *Sourcequery) GetExperimentIdOk() (*int32, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *Sourcequery) SetExperimentId(v int32)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *Sourcequery) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### SetExperimentIdNil

`func (o *Sourcequery) SetExperimentIdNil(b bool)`

 SetExperimentIdNil sets the value for ExperimentId to be an explicit nil

### UnsetExperimentId
`func (o *Sourcequery) UnsetExperimentId()`

UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil
### GetFingerprint

`func (o *Sourcequery) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Sourcequery) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Sourcequery) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Sourcequery) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *Sourcequery) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Sourcequery) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetFunnelsQuery

`func (o *Sourcequery) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *Sourcequery) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *Sourcequery) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.


### GetUuid

`func (o *Sourcequery) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Sourcequery) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Sourcequery) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Sourcequery) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Sourcequery) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Sourcequery) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetCountQuery

`func (o *Sourcequery) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *Sourcequery) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *Sourcequery) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.


### GetExposureQuery

`func (o *Sourcequery) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *Sourcequery) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *Sourcequery) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *Sourcequery) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetAggregationGroupTypeIndex

`func (o *Sourcequery) GetAggregationGroupTypeIndex() int32`

GetAggregationGroupTypeIndex returns the AggregationGroupTypeIndex field if non-nil, zero value otherwise.

### GetAggregationGroupTypeIndexOk

`func (o *Sourcequery) GetAggregationGroupTypeIndexOk() (*int32, bool)`

GetAggregationGroupTypeIndexOk returns a tuple with the AggregationGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationGroupTypeIndex

`func (o *Sourcequery) SetAggregationGroupTypeIndex(v int32)`

SetAggregationGroupTypeIndex sets AggregationGroupTypeIndex field to given value.

### HasAggregationGroupTypeIndex

`func (o *Sourcequery) HasAggregationGroupTypeIndex() bool`

HasAggregationGroupTypeIndex returns a boolean if a field has been set.

### SetAggregationGroupTypeIndexNil

`func (o *Sourcequery) SetAggregationGroupTypeIndexNil(b bool)`

 SetAggregationGroupTypeIndexNil sets the value for AggregationGroupTypeIndex to be an explicit nil

### UnsetAggregationGroupTypeIndex
`func (o *Sourcequery) UnsetAggregationGroupTypeIndex()`

UnsetAggregationGroupTypeIndex ensures that no value is present for AggregationGroupTypeIndex, not even an explicit nil
### GetCalendarHeatmapFilter

`func (o *Sourcequery) GetCalendarHeatmapFilter() CalendarHeatmapFilter`

GetCalendarHeatmapFilter returns the CalendarHeatmapFilter field if non-nil, zero value otherwise.

### GetCalendarHeatmapFilterOk

`func (o *Sourcequery) GetCalendarHeatmapFilterOk() (*CalendarHeatmapFilter, bool)`

GetCalendarHeatmapFilterOk returns a tuple with the CalendarHeatmapFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarHeatmapFilter

`func (o *Sourcequery) SetCalendarHeatmapFilter(v CalendarHeatmapFilter)`

SetCalendarHeatmapFilter sets CalendarHeatmapFilter field to given value.

### HasCalendarHeatmapFilter

`func (o *Sourcequery) HasCalendarHeatmapFilter() bool`

HasCalendarHeatmapFilter returns a boolean if a field has been set.

### GetDataColorTheme

`func (o *Sourcequery) GetDataColorTheme() float32`

GetDataColorTheme returns the DataColorTheme field if non-nil, zero value otherwise.

### GetDataColorThemeOk

`func (o *Sourcequery) GetDataColorThemeOk() (*float32, bool)`

GetDataColorThemeOk returns a tuple with the DataColorTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorTheme

`func (o *Sourcequery) SetDataColorTheme(v float32)`

SetDataColorTheme sets DataColorTheme field to given value.

### HasDataColorTheme

`func (o *Sourcequery) HasDataColorTheme() bool`

HasDataColorTheme returns a boolean if a field has been set.

### SetDataColorThemeNil

`func (o *Sourcequery) SetDataColorThemeNil(b bool)`

 SetDataColorThemeNil sets the value for DataColorTheme to be an explicit nil

### UnsetDataColorTheme
`func (o *Sourcequery) UnsetDataColorTheme()`

UnsetDataColorTheme ensures that no value is present for DataColorTheme, not even an explicit nil
### GetSamplingFactor

`func (o *Sourcequery) GetSamplingFactor() float32`

GetSamplingFactor returns the SamplingFactor field if non-nil, zero value otherwise.

### GetSamplingFactorOk

`func (o *Sourcequery) GetSamplingFactorOk() (*float32, bool)`

GetSamplingFactorOk returns a tuple with the SamplingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingFactor

`func (o *Sourcequery) SetSamplingFactor(v float32)`

SetSamplingFactor sets SamplingFactor field to given value.

### HasSamplingFactor

`func (o *Sourcequery) HasSamplingFactor() bool`

HasSamplingFactor returns a boolean if a field has been set.

### SetSamplingFactorNil

`func (o *Sourcequery) SetSamplingFactorNil(b bool)`

 SetSamplingFactorNil sets the value for SamplingFactor to be an explicit nil

### UnsetSamplingFactor
`func (o *Sourcequery) UnsetSamplingFactor()`

UnsetSamplingFactor ensures that no value is present for SamplingFactor, not even an explicit nil
### GetActions

`func (o *Sourcequery) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Sourcequery) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Sourcequery) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *Sourcequery) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *Sourcequery) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *Sourcequery) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetCommentText

`func (o *Sourcequery) GetCommentText() RecordingPropertyFilter`

GetCommentText returns the CommentText field if non-nil, zero value otherwise.

### GetCommentTextOk

`func (o *Sourcequery) GetCommentTextOk() (*RecordingPropertyFilter, bool)`

GetCommentTextOk returns a tuple with the CommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentText

`func (o *Sourcequery) SetCommentText(v RecordingPropertyFilter)`

SetCommentText sets CommentText field to given value.

### HasCommentText

`func (o *Sourcequery) HasCommentText() bool`

HasCommentText returns a boolean if a field has been set.

### GetConsoleLogFilters

`func (o *Sourcequery) GetConsoleLogFilters() []LogEntryPropertyFilter`

GetConsoleLogFilters returns the ConsoleLogFilters field if non-nil, zero value otherwise.

### GetConsoleLogFiltersOk

`func (o *Sourcequery) GetConsoleLogFiltersOk() (*[]LogEntryPropertyFilter, bool)`

GetConsoleLogFiltersOk returns a tuple with the ConsoleLogFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleLogFilters

`func (o *Sourcequery) SetConsoleLogFilters(v []LogEntryPropertyFilter)`

SetConsoleLogFilters sets ConsoleLogFilters field to given value.

### HasConsoleLogFilters

`func (o *Sourcequery) HasConsoleLogFilters() bool`

HasConsoleLogFilters returns a boolean if a field has been set.

### SetConsoleLogFiltersNil

`func (o *Sourcequery) SetConsoleLogFiltersNil(b bool)`

 SetConsoleLogFiltersNil sets the value for ConsoleLogFilters to be an explicit nil

### UnsetConsoleLogFilters
`func (o *Sourcequery) UnsetConsoleLogFilters()`

UnsetConsoleLogFilters ensures that no value is present for ConsoleLogFilters, not even an explicit nil
### GetDateFrom

`func (o *Sourcequery) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *Sourcequery) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *Sourcequery) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *Sourcequery) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### SetDateFromNil

`func (o *Sourcequery) SetDateFromNil(b bool)`

 SetDateFromNil sets the value for DateFrom to be an explicit nil

### UnsetDateFrom
`func (o *Sourcequery) UnsetDateFrom()`

UnsetDateFrom ensures that no value is present for DateFrom, not even an explicit nil
### GetDateTo

`func (o *Sourcequery) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *Sourcequery) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *Sourcequery) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *Sourcequery) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### SetDateToNil

`func (o *Sourcequery) SetDateToNil(b bool)`

 SetDateToNil sets the value for DateTo to be an explicit nil

### UnsetDateTo
`func (o *Sourcequery) UnsetDateTo()`

UnsetDateTo ensures that no value is present for DateTo, not even an explicit nil
### GetDistinctIds

`func (o *Sourcequery) GetDistinctIds() []string`

GetDistinctIds returns the DistinctIds field if non-nil, zero value otherwise.

### GetDistinctIdsOk

`func (o *Sourcequery) GetDistinctIdsOk() (*[]string, bool)`

GetDistinctIdsOk returns a tuple with the DistinctIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctIds

`func (o *Sourcequery) SetDistinctIds(v []string)`

SetDistinctIds sets DistinctIds field to given value.

### HasDistinctIds

`func (o *Sourcequery) HasDistinctIds() bool`

HasDistinctIds returns a boolean if a field has been set.

### SetDistinctIdsNil

`func (o *Sourcequery) SetDistinctIdsNil(b bool)`

 SetDistinctIdsNil sets the value for DistinctIds to be an explicit nil

### UnsetDistinctIds
`func (o *Sourcequery) UnsetDistinctIds()`

UnsetDistinctIds ensures that no value is present for DistinctIds, not even an explicit nil
### GetFilterTestAccounts

`func (o *Sourcequery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *Sourcequery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *Sourcequery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *Sourcequery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *Sourcequery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *Sourcequery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetHavingPredicates

`func (o *Sourcequery) GetHavingPredicates() []FixedpropertiesInner`

GetHavingPredicates returns the HavingPredicates field if non-nil, zero value otherwise.

### GetHavingPredicatesOk

`func (o *Sourcequery) GetHavingPredicatesOk() (*[]FixedpropertiesInner, bool)`

GetHavingPredicatesOk returns a tuple with the HavingPredicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHavingPredicates

`func (o *Sourcequery) SetHavingPredicates(v []FixedpropertiesInner)`

SetHavingPredicates sets HavingPredicates field to given value.

### HasHavingPredicates

`func (o *Sourcequery) HasHavingPredicates() bool`

HasHavingPredicates returns a boolean if a field has been set.

### SetHavingPredicatesNil

`func (o *Sourcequery) SetHavingPredicatesNil(b bool)`

 SetHavingPredicatesNil sets the value for HavingPredicates to be an explicit nil

### UnsetHavingPredicates
`func (o *Sourcequery) UnsetHavingPredicates()`

UnsetHavingPredicates ensures that no value is present for HavingPredicates, not even an explicit nil
### GetOperand

`func (o *Sourcequery) GetOperand() FilterLogicalOperator`

GetOperand returns the Operand field if non-nil, zero value otherwise.

### GetOperandOk

`func (o *Sourcequery) GetOperandOk() (*FilterLogicalOperator, bool)`

GetOperandOk returns a tuple with the Operand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperand

`func (o *Sourcequery) SetOperand(v FilterLogicalOperator)`

SetOperand sets Operand field to given value.

### HasOperand

`func (o *Sourcequery) HasOperand() bool`

HasOperand returns a boolean if a field has been set.

### GetOrder

`func (o *Sourcequery) GetOrder() RecordingOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Sourcequery) GetOrderOk() (*RecordingOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Sourcequery) SetOrder(v RecordingOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Sourcequery) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderDirection

`func (o *Sourcequery) GetOrderDirection() RecordingOrderDirection`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *Sourcequery) GetOrderDirectionOk() (*RecordingOrderDirection, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *Sourcequery) SetOrderDirection(v RecordingOrderDirection)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *Sourcequery) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetPersonUuid

`func (o *Sourcequery) GetPersonUuid() string`

GetPersonUuid returns the PersonUuid field if non-nil, zero value otherwise.

### GetPersonUuidOk

`func (o *Sourcequery) GetPersonUuidOk() (*string, bool)`

GetPersonUuidOk returns a tuple with the PersonUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonUuid

`func (o *Sourcequery) SetPersonUuid(v string)`

SetPersonUuid sets PersonUuid field to given value.

### HasPersonUuid

`func (o *Sourcequery) HasPersonUuid() bool`

HasPersonUuid returns a boolean if a field has been set.

### SetPersonUuidNil

`func (o *Sourcequery) SetPersonUuidNil(b bool)`

 SetPersonUuidNil sets the value for PersonUuid to be an explicit nil

### UnsetPersonUuid
`func (o *Sourcequery) UnsetPersonUuid()`

UnsetPersonUuid ensures that no value is present for PersonUuid, not even an explicit nil
### GetSessionIds

`func (o *Sourcequery) GetSessionIds() []string`

GetSessionIds returns the SessionIds field if non-nil, zero value otherwise.

### GetSessionIdsOk

`func (o *Sourcequery) GetSessionIdsOk() (*[]string, bool)`

GetSessionIdsOk returns a tuple with the SessionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIds

`func (o *Sourcequery) SetSessionIds(v []string)`

SetSessionIds sets SessionIds field to given value.

### HasSessionIds

`func (o *Sourcequery) HasSessionIds() bool`

HasSessionIds returns a boolean if a field has been set.

### SetSessionIdsNil

`func (o *Sourcequery) SetSessionIdsNil(b bool)`

 SetSessionIdsNil sets the value for SessionIds to be an explicit nil

### UnsetSessionIds
`func (o *Sourcequery) UnsetSessionIds()`

UnsetSessionIds ensures that no value is present for SessionIds, not even an explicit nil
### GetSessionRecordingId

`func (o *Sourcequery) GetSessionRecordingId() string`

GetSessionRecordingId returns the SessionRecordingId field if non-nil, zero value otherwise.

### GetSessionRecordingIdOk

`func (o *Sourcequery) GetSessionRecordingIdOk() (*string, bool)`

GetSessionRecordingIdOk returns a tuple with the SessionRecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRecordingId

`func (o *Sourcequery) SetSessionRecordingId(v string)`

SetSessionRecordingId sets SessionRecordingId field to given value.

### HasSessionRecordingId

`func (o *Sourcequery) HasSessionRecordingId() bool`

HasSessionRecordingId returns a boolean if a field has been set.

### SetSessionRecordingIdNil

`func (o *Sourcequery) SetSessionRecordingIdNil(b bool)`

 SetSessionRecordingIdNil sets the value for SessionRecordingId to be an explicit nil

### UnsetSessionRecordingId
`func (o *Sourcequery) UnsetSessionRecordingId()`

UnsetSessionRecordingId ensures that no value is present for SessionRecordingId, not even an explicit nil
### GetUserModifiedFilters

`func (o *Sourcequery) GetUserModifiedFilters() map[string]interface{}`

GetUserModifiedFilters returns the UserModifiedFilters field if non-nil, zero value otherwise.

### GetUserModifiedFiltersOk

`func (o *Sourcequery) GetUserModifiedFiltersOk() (*map[string]interface{}, bool)`

GetUserModifiedFiltersOk returns a tuple with the UserModifiedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserModifiedFilters

`func (o *Sourcequery) SetUserModifiedFilters(v map[string]interface{})`

SetUserModifiedFilters sets UserModifiedFilters field to given value.

### HasUserModifiedFilters

`func (o *Sourcequery) HasUserModifiedFilters() bool`

HasUserModifiedFilters returns a boolean if a field has been set.

### SetUserModifiedFiltersNil

`func (o *Sourcequery) SetUserModifiedFiltersNil(b bool)`

 SetUserModifiedFiltersNil sets the value for UserModifiedFilters to be an explicit nil

### UnsetUserModifiedFilters
`func (o *Sourcequery) UnsetUserModifiedFilters()`

UnsetUserModifiedFilters ensures that no value is present for UserModifiedFilters, not even an explicit nil
### GetShowColumnConfigurator

`func (o *Sourcequery) GetShowColumnConfigurator() bool`

GetShowColumnConfigurator returns the ShowColumnConfigurator field if non-nil, zero value otherwise.

### GetShowColumnConfiguratorOk

`func (o *Sourcequery) GetShowColumnConfiguratorOk() (*bool, bool)`

GetShowColumnConfiguratorOk returns a tuple with the ShowColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowColumnConfigurator

`func (o *Sourcequery) SetShowColumnConfigurator(v bool)`

SetShowColumnConfigurator sets ShowColumnConfigurator field to given value.

### HasShowColumnConfigurator

`func (o *Sourcequery) HasShowColumnConfigurator() bool`

HasShowColumnConfigurator returns a boolean if a field has been set.

### SetShowColumnConfiguratorNil

`func (o *Sourcequery) SetShowColumnConfiguratorNil(b bool)`

 SetShowColumnConfiguratorNil sets the value for ShowColumnConfigurator to be an explicit nil

### UnsetShowColumnConfigurator
`func (o *Sourcequery) UnsetShowColumnConfigurator()`

UnsetShowColumnConfigurator ensures that no value is present for ShowColumnConfigurator, not even an explicit nil
### GetTraceId

`func (o *Sourcequery) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *Sourcequery) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *Sourcequery) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetEmbedding

`func (o *Sourcequery) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *Sourcequery) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *Sourcequery) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.


### GetEmbeddingVersion

`func (o *Sourcequery) GetEmbeddingVersion() float32`

GetEmbeddingVersion returns the EmbeddingVersion field if non-nil, zero value otherwise.

### GetEmbeddingVersionOk

`func (o *Sourcequery) GetEmbeddingVersionOk() (*float32, bool)`

GetEmbeddingVersionOk returns a tuple with the EmbeddingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingVersion

`func (o *Sourcequery) SetEmbeddingVersion(v float32)`

SetEmbeddingVersion sets EmbeddingVersion field to given value.

### HasEmbeddingVersion

`func (o *Sourcequery) HasEmbeddingVersion() bool`

HasEmbeddingVersion returns a boolean if a field has been set.

### SetEmbeddingVersionNil

`func (o *Sourcequery) SetEmbeddingVersionNil(b bool)`

 SetEmbeddingVersionNil sets the value for EmbeddingVersion to be an explicit nil

### UnsetEmbeddingVersion
`func (o *Sourcequery) UnsetEmbeddingVersion()`

UnsetEmbeddingVersion ensures that no value is present for EmbeddingVersion, not even an explicit nil
### GetGroupKey

`func (o *Sourcequery) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *Sourcequery) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *Sourcequery) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *Sourcequery) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### SetGroupKeyNil

`func (o *Sourcequery) SetGroupKeyNil(b bool)`

 SetGroupKeyNil sets the value for GroupKey to be an explicit nil

### UnsetGroupKey
`func (o *Sourcequery) UnsetGroupKey()`

UnsetGroupKey ensures that no value is present for GroupKey, not even an explicit nil
### GetPersonId

`func (o *Sourcequery) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Sourcequery) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Sourcequery) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Sourcequery) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *Sourcequery) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *Sourcequery) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


