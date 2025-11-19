# Source1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | Pointer to **NullableString** |  | [optional] 
**Event** | Pointer to **NullableString** | Limit to events matching this string | [optional] 
**FixedProperties** | Pointer to [**[]FixedpropertiesInner2**](FixedpropertiesInner2.md) | Fixed properties in the query, can&#39;t be edited in the interface (e.g. scoping down by person) | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "TraceQuery"]
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
**OrderBy** | [**OrderBy1**](OrderBy1.md) |  | 
**Properties** | [**[]FixedpropertiesInner**](FixedpropertiesInner.md) | Properties configurable in the interface | 
**Response** | Pointer to [**TraceQueryResponse**](TraceQueryResponse.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**ActionId** | Pointer to **NullableInt32** | Show events matching a given action | [optional] 
**After** | Pointer to **NullableString** | Only fetch sessions that started after this timestamp | [optional] 
**Before** | Pointer to **NullableString** | Only fetch sessions that started before this timestamp | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**PersonId** | Pointer to **NullableString** | Person who performed the event | [optional] 
**Select** | **[]string** | Return a limited set of data. Will use default columns if empty. | 
**Source** | [**Source3**](Source3.md) |  | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Where** | Pointer to **[]string** | HogQL filters to apply on returned data | [optional] 
**Cohort** | Pointer to **NullableInt32** |  | [optional] 
**DistinctId** | Pointer to **NullableString** |  | [optional] 
**Search** | Pointer to **NullableString** |  | [optional] 
**GroupTypeIndex** | **int32** |  | 
**Explain** | Pointer to **NullableBool** |  | [optional] 
**Filters** | Pointer to [**Filters**](Filters.md) |  | [optional] 
**Query** | **string** |  | 
**Values** | Pointer to **map[string]interface{}** | Constant values that can be referenced with the {placeholder} syntax in the query | [optional] 
**Variables** | Pointer to [**map[string]HogQLVariable**](HogQLVariable.md) | Variables to be substituted into the query | [optional] 
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal1**](Conversiongoal1.md) |  | [optional] [default to null]
**DateRange** | [**DateRange**](DateRange.md) |  | 
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**BreakdownBy** | [**WebStatsBreakdown**](WebStatsBreakdown.md) |  | 
**IncludeBounceRate** | Pointer to **NullableBool** |  | [optional] 
**IncludeScrollDepth** | Pointer to **NullableBool** |  | [optional] 
**StripQueryParams** | Pointer to **NullableBool** |  | [optional] 
**Metric** | [**WebVitalsMetric**](WebVitalsMetric.md) |  | 
**Percentile** | [**WebVitalsPercentile**](WebVitalsPercentile.md) |  | 
**Thresholds** | **[]float32** |  | 
**GroupBy** | [**RevenueAnalyticsTopCustomersGroupBy**](RevenueAnalyticsTopCustomersGroupBy.md) |  | 
**Breakdown** | [**[]RevenueAnalyticsBreakdown**](RevenueAnalyticsBreakdown.md) |  | 
**Interval** | [**SimpleIntervalType**](SimpleIntervalType.md) |  | 
**DraftConversionGoal** | Pointer to [**NullableDraftconversiongoal**](Draftconversiongoal.md) |  | [optional] [default to null]
**IncludeAllConversions** | Pointer to **NullableBool** | Include conversion goal rows even when they don&#39;t match campaign costs table | [optional] 
**IntegrationFilter** | Pointer to [**IntegrationFilter**](IntegrationFilter.md) |  | [optional] 
**Assignee** | Pointer to [**ErrorTrackingIssueAssignee**](ErrorTrackingIssueAssignee.md) |  | [optional] 
**FilterGroup** | Pointer to [**PropertyGroupFilter**](PropertyGroupFilter.md) |  | [optional] 
**GroupKey** | Pointer to **NullableString** |  | [optional] 
**GroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 
**IssueId** | Pointer to **NullableString** |  | [optional] 
**OrderDirection** | Pointer to [**OrderDirection1**](OrderDirection1.md) |  | [optional] 
**RevenueEntity** | Pointer to [**RevenueEntity**](RevenueEntity.md) |  | [optional] 
**RevenuePeriod** | Pointer to [**RevenuePeriod**](RevenuePeriod.md) |  | [optional] 
**SearchQuery** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**Status2**](Status2.md) |  | [optional] 
**VolumeResolution** | **int32** |  | 
**WithAggregations** | Pointer to **NullableBool** |  | [optional] 
**WithFirstEvent** | Pointer to **NullableBool** |  | [optional] 
**WithLastEvent** | Pointer to **NullableBool** |  | [optional] 
**Events** | **[]string** |  | 
**ExperimentId** | Pointer to **NullableInt32** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**FunnelsQuery** | [**FunnelsQuery**](FunnelsQuery.md) |  | 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**CountQuery** | [**TrendsQuery**](TrendsQuery.md) |  | 
**ExposureQuery** | Pointer to [**TrendsQuery**](TrendsQuery.md) |  | [optional] 
**ShowColumnConfigurator** | Pointer to **NullableBool** |  | [optional] 
**TraceId** | **string** |  | 

## Methods

### NewSource1

`func NewSource1(orderBy OrderBy1, properties []FixedpropertiesInner, select_ []string, source Source3, groupTypeIndex int32, query string, dateRange DateRange, breakdownBy WebStatsBreakdown, metric WebVitalsMetric, percentile WebVitalsPercentile, thresholds []float32, groupBy RevenueAnalyticsTopCustomersGroupBy, breakdown []RevenueAnalyticsBreakdown, interval SimpleIntervalType, volumeResolution int32, events []string, funnelsQuery FunnelsQuery, countQuery TrendsQuery, traceId string, ) *Source1`

NewSource1 instantiates a new Source1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSource1WithDefaults

`func NewSource1WithDefaults() *Source1`

NewSource1WithDefaults instantiates a new Source1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *Source1) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *Source1) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *Source1) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *Source1) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *Source1) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *Source1) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetEvent

`func (o *Source1) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Source1) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Source1) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Source1) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *Source1) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *Source1) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetFixedProperties

`func (o *Source1) GetFixedProperties() []FixedpropertiesInner2`

GetFixedProperties returns the FixedProperties field if non-nil, zero value otherwise.

### GetFixedPropertiesOk

`func (o *Source1) GetFixedPropertiesOk() (*[]FixedpropertiesInner2, bool)`

GetFixedPropertiesOk returns a tuple with the FixedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProperties

`func (o *Source1) SetFixedProperties(v []FixedpropertiesInner2)`

SetFixedProperties sets FixedProperties field to given value.

### HasFixedProperties

`func (o *Source1) HasFixedProperties() bool`

HasFixedProperties returns a boolean if a field has been set.

### SetFixedPropertiesNil

`func (o *Source1) SetFixedPropertiesNil(b bool)`

 SetFixedPropertiesNil sets the value for FixedProperties to be an explicit nil

### UnsetFixedProperties
`func (o *Source1) UnsetFixedProperties()`

UnsetFixedProperties ensures that no value is present for FixedProperties, not even an explicit nil
### GetKind

`func (o *Source1) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Source1) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Source1) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Source1) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *Source1) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Source1) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Source1) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Source1) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *Source1) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *Source1) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMath

`func (o *Source1) GetMath() string`

GetMath returns the Math field if non-nil, zero value otherwise.

### GetMathOk

`func (o *Source1) GetMathOk() (*string, bool)`

GetMathOk returns a tuple with the Math field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMath

`func (o *Source1) SetMath(v string)`

SetMath sets Math field to given value.

### HasMath

`func (o *Source1) HasMath() bool`

HasMath returns a boolean if a field has been set.

### SetMathNil

`func (o *Source1) SetMathNil(b bool)`

 SetMathNil sets the value for Math to be an explicit nil

### UnsetMath
`func (o *Source1) UnsetMath()`

UnsetMath ensures that no value is present for Math, not even an explicit nil
### GetMathGroupTypeIndex

`func (o *Source1) GetMathGroupTypeIndex() MathGroupTypeIndex`

GetMathGroupTypeIndex returns the MathGroupTypeIndex field if non-nil, zero value otherwise.

### GetMathGroupTypeIndexOk

`func (o *Source1) GetMathGroupTypeIndexOk() (*MathGroupTypeIndex, bool)`

GetMathGroupTypeIndexOk returns a tuple with the MathGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathGroupTypeIndex

`func (o *Source1) SetMathGroupTypeIndex(v MathGroupTypeIndex)`

SetMathGroupTypeIndex sets MathGroupTypeIndex field to given value.

### HasMathGroupTypeIndex

`func (o *Source1) HasMathGroupTypeIndex() bool`

HasMathGroupTypeIndex returns a boolean if a field has been set.

### GetMathHogql

`func (o *Source1) GetMathHogql() string`

GetMathHogql returns the MathHogql field if non-nil, zero value otherwise.

### GetMathHogqlOk

`func (o *Source1) GetMathHogqlOk() (*string, bool)`

GetMathHogqlOk returns a tuple with the MathHogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathHogql

`func (o *Source1) SetMathHogql(v string)`

SetMathHogql sets MathHogql field to given value.

### HasMathHogql

`func (o *Source1) HasMathHogql() bool`

HasMathHogql returns a boolean if a field has been set.

### SetMathHogqlNil

`func (o *Source1) SetMathHogqlNil(b bool)`

 SetMathHogqlNil sets the value for MathHogql to be an explicit nil

### UnsetMathHogql
`func (o *Source1) UnsetMathHogql()`

UnsetMathHogql ensures that no value is present for MathHogql, not even an explicit nil
### GetMathMultiplier

`func (o *Source1) GetMathMultiplier() float32`

GetMathMultiplier returns the MathMultiplier field if non-nil, zero value otherwise.

### GetMathMultiplierOk

`func (o *Source1) GetMathMultiplierOk() (*float32, bool)`

GetMathMultiplierOk returns a tuple with the MathMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathMultiplier

`func (o *Source1) SetMathMultiplier(v float32)`

SetMathMultiplier sets MathMultiplier field to given value.

### HasMathMultiplier

`func (o *Source1) HasMathMultiplier() bool`

HasMathMultiplier returns a boolean if a field has been set.

### SetMathMultiplierNil

`func (o *Source1) SetMathMultiplierNil(b bool)`

 SetMathMultiplierNil sets the value for MathMultiplier to be an explicit nil

### UnsetMathMultiplier
`func (o *Source1) UnsetMathMultiplier()`

UnsetMathMultiplier ensures that no value is present for MathMultiplier, not even an explicit nil
### GetMathProperty

`func (o *Source1) GetMathProperty() string`

GetMathProperty returns the MathProperty field if non-nil, zero value otherwise.

### GetMathPropertyOk

`func (o *Source1) GetMathPropertyOk() (*string, bool)`

GetMathPropertyOk returns a tuple with the MathProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathProperty

`func (o *Source1) SetMathProperty(v string)`

SetMathProperty sets MathProperty field to given value.

### HasMathProperty

`func (o *Source1) HasMathProperty() bool`

HasMathProperty returns a boolean if a field has been set.

### SetMathPropertyNil

`func (o *Source1) SetMathPropertyNil(b bool)`

 SetMathPropertyNil sets the value for MathProperty to be an explicit nil

### UnsetMathProperty
`func (o *Source1) UnsetMathProperty()`

UnsetMathProperty ensures that no value is present for MathProperty, not even an explicit nil
### GetMathPropertyRevenueCurrency

`func (o *Source1) GetMathPropertyRevenueCurrency() RevenueCurrencyPropertyConfig`

GetMathPropertyRevenueCurrency returns the MathPropertyRevenueCurrency field if non-nil, zero value otherwise.

### GetMathPropertyRevenueCurrencyOk

`func (o *Source1) GetMathPropertyRevenueCurrencyOk() (*RevenueCurrencyPropertyConfig, bool)`

GetMathPropertyRevenueCurrencyOk returns a tuple with the MathPropertyRevenueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyRevenueCurrency

`func (o *Source1) SetMathPropertyRevenueCurrency(v RevenueCurrencyPropertyConfig)`

SetMathPropertyRevenueCurrency sets MathPropertyRevenueCurrency field to given value.

### HasMathPropertyRevenueCurrency

`func (o *Source1) HasMathPropertyRevenueCurrency() bool`

HasMathPropertyRevenueCurrency returns a boolean if a field has been set.

### GetMathPropertyType

`func (o *Source1) GetMathPropertyType() string`

GetMathPropertyType returns the MathPropertyType field if non-nil, zero value otherwise.

### GetMathPropertyTypeOk

`func (o *Source1) GetMathPropertyTypeOk() (*string, bool)`

GetMathPropertyTypeOk returns a tuple with the MathPropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathPropertyType

`func (o *Source1) SetMathPropertyType(v string)`

SetMathPropertyType sets MathPropertyType field to given value.

### HasMathPropertyType

`func (o *Source1) HasMathPropertyType() bool`

HasMathPropertyType returns a boolean if a field has been set.

### SetMathPropertyTypeNil

`func (o *Source1) SetMathPropertyTypeNil(b bool)`

 SetMathPropertyTypeNil sets the value for MathPropertyType to be an explicit nil

### UnsetMathPropertyType
`func (o *Source1) UnsetMathPropertyType()`

UnsetMathPropertyType ensures that no value is present for MathPropertyType, not even an explicit nil
### GetName

`func (o *Source1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Source1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Source1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Source1) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Source1) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Source1) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptionalInFunnel

`func (o *Source1) GetOptionalInFunnel() bool`

GetOptionalInFunnel returns the OptionalInFunnel field if non-nil, zero value otherwise.

### GetOptionalInFunnelOk

`func (o *Source1) GetOptionalInFunnelOk() (*bool, bool)`

GetOptionalInFunnelOk returns a tuple with the OptionalInFunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalInFunnel

`func (o *Source1) SetOptionalInFunnel(v bool)`

SetOptionalInFunnel sets OptionalInFunnel field to given value.

### HasOptionalInFunnel

`func (o *Source1) HasOptionalInFunnel() bool`

HasOptionalInFunnel returns a boolean if a field has been set.

### SetOptionalInFunnelNil

`func (o *Source1) SetOptionalInFunnelNil(b bool)`

 SetOptionalInFunnelNil sets the value for OptionalInFunnel to be an explicit nil

### UnsetOptionalInFunnel
`func (o *Source1) UnsetOptionalInFunnel()`

UnsetOptionalInFunnel ensures that no value is present for OptionalInFunnel, not even an explicit nil
### GetOrderBy

`func (o *Source1) GetOrderBy() OrderBy1`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *Source1) GetOrderByOk() (*OrderBy1, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *Source1) SetOrderBy(v OrderBy1)`

SetOrderBy sets OrderBy field to given value.


### GetProperties

`func (o *Source1) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Source1) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Source1) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.


### SetPropertiesNil

`func (o *Source1) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Source1) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *Source1) GetResponse() TraceQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Source1) GetResponseOk() (*TraceQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Source1) SetResponse(v TraceQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Source1) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetVersion

`func (o *Source1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Source1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Source1) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Source1) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Source1) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Source1) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetActionId

`func (o *Source1) GetActionId() int32`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *Source1) GetActionIdOk() (*int32, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *Source1) SetActionId(v int32)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *Source1) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### SetActionIdNil

`func (o *Source1) SetActionIdNil(b bool)`

 SetActionIdNil sets the value for ActionId to be an explicit nil

### UnsetActionId
`func (o *Source1) UnsetActionId()`

UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
### GetAfter

`func (o *Source1) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *Source1) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *Source1) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *Source1) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *Source1) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *Source1) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetBefore

`func (o *Source1) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *Source1) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *Source1) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *Source1) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### SetBeforeNil

`func (o *Source1) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *Source1) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetFilterTestAccounts

`func (o *Source1) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *Source1) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *Source1) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *Source1) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *Source1) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *Source1) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetModifiers

`func (o *Source1) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Source1) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Source1) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Source1) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *Source1) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Source1) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Source1) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Source1) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *Source1) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *Source1) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetPersonId

`func (o *Source1) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Source1) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Source1) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Source1) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *Source1) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *Source1) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetSelect

`func (o *Source1) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *Source1) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *Source1) SetSelect(v []string)`

SetSelect sets Select field to given value.


### SetSelectNil

`func (o *Source1) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *Source1) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetSource

`func (o *Source1) GetSource() Source3`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Source1) GetSourceOk() (*Source3, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Source1) SetSource(v Source3)`

SetSource sets Source field to given value.


### GetTags

`func (o *Source1) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Source1) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Source1) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Source1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWhere

`func (o *Source1) GetWhere() []string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *Source1) GetWhereOk() (*[]string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *Source1) SetWhere(v []string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *Source1) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### SetWhereNil

`func (o *Source1) SetWhereNil(b bool)`

 SetWhereNil sets the value for Where to be an explicit nil

### UnsetWhere
`func (o *Source1) UnsetWhere()`

UnsetWhere ensures that no value is present for Where, not even an explicit nil
### GetCohort

`func (o *Source1) GetCohort() int32`

GetCohort returns the Cohort field if non-nil, zero value otherwise.

### GetCohortOk

`func (o *Source1) GetCohortOk() (*int32, bool)`

GetCohortOk returns a tuple with the Cohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohort

`func (o *Source1) SetCohort(v int32)`

SetCohort sets Cohort field to given value.

### HasCohort

`func (o *Source1) HasCohort() bool`

HasCohort returns a boolean if a field has been set.

### SetCohortNil

`func (o *Source1) SetCohortNil(b bool)`

 SetCohortNil sets the value for Cohort to be an explicit nil

### UnsetCohort
`func (o *Source1) UnsetCohort()`

UnsetCohort ensures that no value is present for Cohort, not even an explicit nil
### GetDistinctId

`func (o *Source1) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *Source1) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *Source1) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *Source1) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### SetDistinctIdNil

`func (o *Source1) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *Source1) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetSearch

`func (o *Source1) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *Source1) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *Source1) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *Source1) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### SetSearchNil

`func (o *Source1) SetSearchNil(b bool)`

 SetSearchNil sets the value for Search to be an explicit nil

### UnsetSearch
`func (o *Source1) UnsetSearch()`

UnsetSearch ensures that no value is present for Search, not even an explicit nil
### GetGroupTypeIndex

`func (o *Source1) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Source1) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Source1) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.


### GetExplain

`func (o *Source1) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *Source1) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *Source1) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *Source1) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *Source1) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *Source1) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetFilters

`func (o *Source1) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Source1) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Source1) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Source1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetQuery

`func (o *Source1) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Source1) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Source1) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetValues

`func (o *Source1) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Source1) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Source1) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *Source1) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *Source1) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *Source1) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetVariables

`func (o *Source1) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Source1) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Source1) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Source1) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *Source1) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Source1) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetCompareFilter

`func (o *Source1) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *Source1) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *Source1) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *Source1) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *Source1) GetConversionGoal() Conversiongoal1`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *Source1) GetConversionGoalOk() (*Conversiongoal1, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *Source1) SetConversionGoal(v Conversiongoal1)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *Source1) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *Source1) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *Source1) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDateRange

`func (o *Source1) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Source1) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Source1) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetDoPathCleaning

`func (o *Source1) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *Source1) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *Source1) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *Source1) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *Source1) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *Source1) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetIncludeRevenue

`func (o *Source1) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *Source1) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *Source1) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *Source1) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *Source1) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *Source1) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetSampling

`func (o *Source1) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *Source1) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *Source1) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *Source1) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *Source1) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *Source1) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *Source1) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *Source1) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *Source1) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *Source1) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetBreakdownBy

`func (o *Source1) GetBreakdownBy() WebStatsBreakdown`

GetBreakdownBy returns the BreakdownBy field if non-nil, zero value otherwise.

### GetBreakdownByOk

`func (o *Source1) GetBreakdownByOk() (*WebStatsBreakdown, bool)`

GetBreakdownByOk returns a tuple with the BreakdownBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownBy

`func (o *Source1) SetBreakdownBy(v WebStatsBreakdown)`

SetBreakdownBy sets BreakdownBy field to given value.


### GetIncludeBounceRate

`func (o *Source1) GetIncludeBounceRate() bool`

GetIncludeBounceRate returns the IncludeBounceRate field if non-nil, zero value otherwise.

### GetIncludeBounceRateOk

`func (o *Source1) GetIncludeBounceRateOk() (*bool, bool)`

GetIncludeBounceRateOk returns a tuple with the IncludeBounceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBounceRate

`func (o *Source1) SetIncludeBounceRate(v bool)`

SetIncludeBounceRate sets IncludeBounceRate field to given value.

### HasIncludeBounceRate

`func (o *Source1) HasIncludeBounceRate() bool`

HasIncludeBounceRate returns a boolean if a field has been set.

### SetIncludeBounceRateNil

`func (o *Source1) SetIncludeBounceRateNil(b bool)`

 SetIncludeBounceRateNil sets the value for IncludeBounceRate to be an explicit nil

### UnsetIncludeBounceRate
`func (o *Source1) UnsetIncludeBounceRate()`

UnsetIncludeBounceRate ensures that no value is present for IncludeBounceRate, not even an explicit nil
### GetIncludeScrollDepth

`func (o *Source1) GetIncludeScrollDepth() bool`

GetIncludeScrollDepth returns the IncludeScrollDepth field if non-nil, zero value otherwise.

### GetIncludeScrollDepthOk

`func (o *Source1) GetIncludeScrollDepthOk() (*bool, bool)`

GetIncludeScrollDepthOk returns a tuple with the IncludeScrollDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeScrollDepth

`func (o *Source1) SetIncludeScrollDepth(v bool)`

SetIncludeScrollDepth sets IncludeScrollDepth field to given value.

### HasIncludeScrollDepth

`func (o *Source1) HasIncludeScrollDepth() bool`

HasIncludeScrollDepth returns a boolean if a field has been set.

### SetIncludeScrollDepthNil

`func (o *Source1) SetIncludeScrollDepthNil(b bool)`

 SetIncludeScrollDepthNil sets the value for IncludeScrollDepth to be an explicit nil

### UnsetIncludeScrollDepth
`func (o *Source1) UnsetIncludeScrollDepth()`

UnsetIncludeScrollDepth ensures that no value is present for IncludeScrollDepth, not even an explicit nil
### GetStripQueryParams

`func (o *Source1) GetStripQueryParams() bool`

GetStripQueryParams returns the StripQueryParams field if non-nil, zero value otherwise.

### GetStripQueryParamsOk

`func (o *Source1) GetStripQueryParamsOk() (*bool, bool)`

GetStripQueryParamsOk returns a tuple with the StripQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripQueryParams

`func (o *Source1) SetStripQueryParams(v bool)`

SetStripQueryParams sets StripQueryParams field to given value.

### HasStripQueryParams

`func (o *Source1) HasStripQueryParams() bool`

HasStripQueryParams returns a boolean if a field has been set.

### SetStripQueryParamsNil

`func (o *Source1) SetStripQueryParamsNil(b bool)`

 SetStripQueryParamsNil sets the value for StripQueryParams to be an explicit nil

### UnsetStripQueryParams
`func (o *Source1) UnsetStripQueryParams()`

UnsetStripQueryParams ensures that no value is present for StripQueryParams, not even an explicit nil
### GetMetric

`func (o *Source1) GetMetric() WebVitalsMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *Source1) GetMetricOk() (*WebVitalsMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *Source1) SetMetric(v WebVitalsMetric)`

SetMetric sets Metric field to given value.


### GetPercentile

`func (o *Source1) GetPercentile() WebVitalsPercentile`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *Source1) GetPercentileOk() (*WebVitalsPercentile, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *Source1) SetPercentile(v WebVitalsPercentile)`

SetPercentile sets Percentile field to given value.


### GetThresholds

`func (o *Source1) GetThresholds() []float32`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *Source1) GetThresholdsOk() (*[]float32, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *Source1) SetThresholds(v []float32)`

SetThresholds sets Thresholds field to given value.


### GetGroupBy

`func (o *Source1) GetGroupBy() RevenueAnalyticsTopCustomersGroupBy`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Source1) GetGroupByOk() (*RevenueAnalyticsTopCustomersGroupBy, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Source1) SetGroupBy(v RevenueAnalyticsTopCustomersGroupBy)`

SetGroupBy sets GroupBy field to given value.


### GetBreakdown

`func (o *Source1) GetBreakdown() []RevenueAnalyticsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *Source1) GetBreakdownOk() (*[]RevenueAnalyticsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *Source1) SetBreakdown(v []RevenueAnalyticsBreakdown)`

SetBreakdown sets Breakdown field to given value.


### GetInterval

`func (o *Source1) GetInterval() SimpleIntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Source1) GetIntervalOk() (*SimpleIntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Source1) SetInterval(v SimpleIntervalType)`

SetInterval sets Interval field to given value.


### GetDraftConversionGoal

`func (o *Source1) GetDraftConversionGoal() Draftconversiongoal`

GetDraftConversionGoal returns the DraftConversionGoal field if non-nil, zero value otherwise.

### GetDraftConversionGoalOk

`func (o *Source1) GetDraftConversionGoalOk() (*Draftconversiongoal, bool)`

GetDraftConversionGoalOk returns a tuple with the DraftConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftConversionGoal

`func (o *Source1) SetDraftConversionGoal(v Draftconversiongoal)`

SetDraftConversionGoal sets DraftConversionGoal field to given value.

### HasDraftConversionGoal

`func (o *Source1) HasDraftConversionGoal() bool`

HasDraftConversionGoal returns a boolean if a field has been set.

### SetDraftConversionGoalNil

`func (o *Source1) SetDraftConversionGoalNil(b bool)`

 SetDraftConversionGoalNil sets the value for DraftConversionGoal to be an explicit nil

### UnsetDraftConversionGoal
`func (o *Source1) UnsetDraftConversionGoal()`

UnsetDraftConversionGoal ensures that no value is present for DraftConversionGoal, not even an explicit nil
### GetIncludeAllConversions

`func (o *Source1) GetIncludeAllConversions() bool`

GetIncludeAllConversions returns the IncludeAllConversions field if non-nil, zero value otherwise.

### GetIncludeAllConversionsOk

`func (o *Source1) GetIncludeAllConversionsOk() (*bool, bool)`

GetIncludeAllConversionsOk returns a tuple with the IncludeAllConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllConversions

`func (o *Source1) SetIncludeAllConversions(v bool)`

SetIncludeAllConversions sets IncludeAllConversions field to given value.

### HasIncludeAllConversions

`func (o *Source1) HasIncludeAllConversions() bool`

HasIncludeAllConversions returns a boolean if a field has been set.

### SetIncludeAllConversionsNil

`func (o *Source1) SetIncludeAllConversionsNil(b bool)`

 SetIncludeAllConversionsNil sets the value for IncludeAllConversions to be an explicit nil

### UnsetIncludeAllConversions
`func (o *Source1) UnsetIncludeAllConversions()`

UnsetIncludeAllConversions ensures that no value is present for IncludeAllConversions, not even an explicit nil
### GetIntegrationFilter

`func (o *Source1) GetIntegrationFilter() IntegrationFilter`

GetIntegrationFilter returns the IntegrationFilter field if non-nil, zero value otherwise.

### GetIntegrationFilterOk

`func (o *Source1) GetIntegrationFilterOk() (*IntegrationFilter, bool)`

GetIntegrationFilterOk returns a tuple with the IntegrationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationFilter

`func (o *Source1) SetIntegrationFilter(v IntegrationFilter)`

SetIntegrationFilter sets IntegrationFilter field to given value.

### HasIntegrationFilter

`func (o *Source1) HasIntegrationFilter() bool`

HasIntegrationFilter returns a boolean if a field has been set.

### GetAssignee

`func (o *Source1) GetAssignee() ErrorTrackingIssueAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Source1) GetAssigneeOk() (*ErrorTrackingIssueAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Source1) SetAssignee(v ErrorTrackingIssueAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *Source1) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetFilterGroup

`func (o *Source1) GetFilterGroup() PropertyGroupFilter`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *Source1) GetFilterGroupOk() (*PropertyGroupFilter, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *Source1) SetFilterGroup(v PropertyGroupFilter)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *Source1) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### GetGroupKey

`func (o *Source1) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *Source1) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *Source1) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *Source1) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### SetGroupKeyNil

`func (o *Source1) SetGroupKeyNil(b bool)`

 SetGroupKeyNil sets the value for GroupKey to be an explicit nil

### UnsetGroupKey
`func (o *Source1) UnsetGroupKey()`

UnsetGroupKey ensures that no value is present for GroupKey, not even an explicit nil
### GetGroupTypeIndex

`func (o *Source1) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Source1) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Source1) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *Source1) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *Source1) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Source1) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetIssueId

`func (o *Source1) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *Source1) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *Source1) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *Source1) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### SetIssueIdNil

`func (o *Source1) SetIssueIdNil(b bool)`

 SetIssueIdNil sets the value for IssueId to be an explicit nil

### UnsetIssueId
`func (o *Source1) UnsetIssueId()`

UnsetIssueId ensures that no value is present for IssueId, not even an explicit nil
### GetOrderDirection

`func (o *Source1) GetOrderDirection() OrderDirection1`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *Source1) GetOrderDirectionOk() (*OrderDirection1, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *Source1) SetOrderDirection(v OrderDirection1)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *Source1) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetRevenueEntity

`func (o *Source1) GetRevenueEntity() RevenueEntity`

GetRevenueEntity returns the RevenueEntity field if non-nil, zero value otherwise.

### GetRevenueEntityOk

`func (o *Source1) GetRevenueEntityOk() (*RevenueEntity, bool)`

GetRevenueEntityOk returns a tuple with the RevenueEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueEntity

`func (o *Source1) SetRevenueEntity(v RevenueEntity)`

SetRevenueEntity sets RevenueEntity field to given value.

### HasRevenueEntity

`func (o *Source1) HasRevenueEntity() bool`

HasRevenueEntity returns a boolean if a field has been set.

### GetRevenuePeriod

`func (o *Source1) GetRevenuePeriod() RevenuePeriod`

GetRevenuePeriod returns the RevenuePeriod field if non-nil, zero value otherwise.

### GetRevenuePeriodOk

`func (o *Source1) GetRevenuePeriodOk() (*RevenuePeriod, bool)`

GetRevenuePeriodOk returns a tuple with the RevenuePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePeriod

`func (o *Source1) SetRevenuePeriod(v RevenuePeriod)`

SetRevenuePeriod sets RevenuePeriod field to given value.

### HasRevenuePeriod

`func (o *Source1) HasRevenuePeriod() bool`

HasRevenuePeriod returns a boolean if a field has been set.

### GetSearchQuery

`func (o *Source1) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *Source1) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *Source1) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *Source1) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### SetSearchQueryNil

`func (o *Source1) SetSearchQueryNil(b bool)`

 SetSearchQueryNil sets the value for SearchQuery to be an explicit nil

### UnsetSearchQuery
`func (o *Source1) UnsetSearchQuery()`

UnsetSearchQuery ensures that no value is present for SearchQuery, not even an explicit nil
### GetStatus

`func (o *Source1) GetStatus() Status2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Source1) GetStatusOk() (*Status2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Source1) SetStatus(v Status2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Source1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumeResolution

`func (o *Source1) GetVolumeResolution() int32`

GetVolumeResolution returns the VolumeResolution field if non-nil, zero value otherwise.

### GetVolumeResolutionOk

`func (o *Source1) GetVolumeResolutionOk() (*int32, bool)`

GetVolumeResolutionOk returns a tuple with the VolumeResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeResolution

`func (o *Source1) SetVolumeResolution(v int32)`

SetVolumeResolution sets VolumeResolution field to given value.


### GetWithAggregations

`func (o *Source1) GetWithAggregations() bool`

GetWithAggregations returns the WithAggregations field if non-nil, zero value otherwise.

### GetWithAggregationsOk

`func (o *Source1) GetWithAggregationsOk() (*bool, bool)`

GetWithAggregationsOk returns a tuple with the WithAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAggregations

`func (o *Source1) SetWithAggregations(v bool)`

SetWithAggregations sets WithAggregations field to given value.

### HasWithAggregations

`func (o *Source1) HasWithAggregations() bool`

HasWithAggregations returns a boolean if a field has been set.

### SetWithAggregationsNil

`func (o *Source1) SetWithAggregationsNil(b bool)`

 SetWithAggregationsNil sets the value for WithAggregations to be an explicit nil

### UnsetWithAggregations
`func (o *Source1) UnsetWithAggregations()`

UnsetWithAggregations ensures that no value is present for WithAggregations, not even an explicit nil
### GetWithFirstEvent

`func (o *Source1) GetWithFirstEvent() bool`

GetWithFirstEvent returns the WithFirstEvent field if non-nil, zero value otherwise.

### GetWithFirstEventOk

`func (o *Source1) GetWithFirstEventOk() (*bool, bool)`

GetWithFirstEventOk returns a tuple with the WithFirstEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithFirstEvent

`func (o *Source1) SetWithFirstEvent(v bool)`

SetWithFirstEvent sets WithFirstEvent field to given value.

### HasWithFirstEvent

`func (o *Source1) HasWithFirstEvent() bool`

HasWithFirstEvent returns a boolean if a field has been set.

### SetWithFirstEventNil

`func (o *Source1) SetWithFirstEventNil(b bool)`

 SetWithFirstEventNil sets the value for WithFirstEvent to be an explicit nil

### UnsetWithFirstEvent
`func (o *Source1) UnsetWithFirstEvent()`

UnsetWithFirstEvent ensures that no value is present for WithFirstEvent, not even an explicit nil
### GetWithLastEvent

`func (o *Source1) GetWithLastEvent() bool`

GetWithLastEvent returns the WithLastEvent field if non-nil, zero value otherwise.

### GetWithLastEventOk

`func (o *Source1) GetWithLastEventOk() (*bool, bool)`

GetWithLastEventOk returns a tuple with the WithLastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithLastEvent

`func (o *Source1) SetWithLastEvent(v bool)`

SetWithLastEvent sets WithLastEvent field to given value.

### HasWithLastEvent

`func (o *Source1) HasWithLastEvent() bool`

HasWithLastEvent returns a boolean if a field has been set.

### SetWithLastEventNil

`func (o *Source1) SetWithLastEventNil(b bool)`

 SetWithLastEventNil sets the value for WithLastEvent to be an explicit nil

### UnsetWithLastEvent
`func (o *Source1) UnsetWithLastEvent()`

UnsetWithLastEvent ensures that no value is present for WithLastEvent, not even an explicit nil
### GetEvents

`func (o *Source1) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Source1) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Source1) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetExperimentId

`func (o *Source1) GetExperimentId() int32`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *Source1) GetExperimentIdOk() (*int32, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *Source1) SetExperimentId(v int32)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *Source1) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### SetExperimentIdNil

`func (o *Source1) SetExperimentIdNil(b bool)`

 SetExperimentIdNil sets the value for ExperimentId to be an explicit nil

### UnsetExperimentId
`func (o *Source1) UnsetExperimentId()`

UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil
### GetFingerprint

`func (o *Source1) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Source1) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Source1) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Source1) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *Source1) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Source1) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetFunnelsQuery

`func (o *Source1) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *Source1) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *Source1) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.


### GetUuid

`func (o *Source1) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Source1) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Source1) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Source1) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Source1) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Source1) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetCountQuery

`func (o *Source1) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *Source1) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *Source1) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.


### GetExposureQuery

`func (o *Source1) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *Source1) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *Source1) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *Source1) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetShowColumnConfigurator

`func (o *Source1) GetShowColumnConfigurator() bool`

GetShowColumnConfigurator returns the ShowColumnConfigurator field if non-nil, zero value otherwise.

### GetShowColumnConfiguratorOk

`func (o *Source1) GetShowColumnConfiguratorOk() (*bool, bool)`

GetShowColumnConfiguratorOk returns a tuple with the ShowColumnConfigurator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowColumnConfigurator

`func (o *Source1) SetShowColumnConfigurator(v bool)`

SetShowColumnConfigurator sets ShowColumnConfigurator field to given value.

### HasShowColumnConfigurator

`func (o *Source1) HasShowColumnConfigurator() bool`

HasShowColumnConfigurator returns a boolean if a field has been set.

### SetShowColumnConfiguratorNil

`func (o *Source1) SetShowColumnConfiguratorNil(b bool)`

 SetShowColumnConfiguratorNil sets the value for ShowColumnConfigurator to be an explicit nil

### UnsetShowColumnConfigurator
`func (o *Source1) UnsetShowColumnConfigurator()`

UnsetShowColumnConfigurator ensures that no value is present for ShowColumnConfigurator, not even an explicit nil
### GetTraceId

`func (o *Source1) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *Source1) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *Source1) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


