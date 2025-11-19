# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Explain** | Pointer to **NullableBool** |  | [optional] 
**Filters** | Pointer to [**HogQLFilters**](HogQLFilters.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "LifecycleQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Client provided name of the query | [optional] 
**Query** | **string** |  | 
**Response** | Pointer to [**LifecycleQueryResponse**](LifecycleQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Values** | Pointer to **map[string]interface{}** | Constant values that can be referenced with the {placeholder} syntax in the query | [optional] 
**Variables** | Pointer to [**map[string]HogQLVariable**](HogQLVariable.md) | Variables to be substituted into the query | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**AggregationGroupTypeIndex** | Pointer to **NullableInt32** | Groups aggregation | [optional] 
**BreakdownFilter** | Pointer to [**BreakdownFilter**](BreakdownFilter.md) |  | [optional] 
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal**](Conversiongoal.md) |  | [optional] [default to null]
**DataColorTheme** | Pointer to **NullableFloat32** | Colors used in the insight&#39;s visualization | [optional] 
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** | Exclude internal and test users by applying the respective filters | [optional] [default to false]
**Interval** | Pointer to [**IntervalType**](IntervalType.md) |  | [optional] 
**Properties** | Pointer to [**NullableProperties1**](Properties1.md) |  | [optional] [default to []]
**SamplingFactor** | Pointer to **NullableFloat32** | Sampling rate | [optional] 
**Series** | [**[]Series1Inner**](Series1Inner.md) | Events and actions to include | 
**TrendsFilter** | Pointer to [**TrendsFilter**](TrendsFilter.md) |  | [optional] 
**FunnelsFilter** | Pointer to [**FunnelsFilter**](FunnelsFilter.md) |  | [optional] 
**RetentionFilter** | [**RetentionFilter**](RetentionFilter.md) |  | 
**FunnelPathsFilter** | Pointer to [**FunnelPathsFilter**](FunnelPathsFilter.md) |  | [optional] 
**PathsFilter** | [**PathsFilter**](PathsFilter.md) |  | 
**IntervalCount** | Pointer to **NullableInt32** | How many intervals comprise a period. Only used for cohorts, otherwise default 1. | [optional] 
**StickinessFilter** | Pointer to [**StickinessFilter**](StickinessFilter.md) |  | [optional] 
**LifecycleFilter** | Pointer to [**LifecycleFilter**](LifecycleFilter.md) |  | [optional] 

## Methods

### NewQuery

`func NewQuery(query string, series []Series1Inner, retentionFilter RetentionFilter, pathsFilter PathsFilter, ) *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplain

`func (o *Query) GetExplain() bool`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *Query) GetExplainOk() (*bool, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *Query) SetExplain(v bool)`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *Query) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *Query) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *Query) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil
### GetFilters

`func (o *Query) GetFilters() HogQLFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Query) GetFiltersOk() (*HogQLFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Query) SetFilters(v HogQLFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Query) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetKind

`func (o *Query) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Query) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Query) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Query) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *Query) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Query) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Query) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Query) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetName

`func (o *Query) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Query) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Query) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Query) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Query) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Query) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetQuery

`func (o *Query) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Query) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Query) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetResponse

`func (o *Query) GetResponse() LifecycleQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Query) GetResponseOk() (*LifecycleQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Query) SetResponse(v LifecycleQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Query) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *Query) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Query) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Query) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Query) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetValues

`func (o *Query) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Query) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Query) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *Query) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *Query) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *Query) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetVariables

`func (o *Query) GetVariables() map[string]HogQLVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Query) GetVariablesOk() (*map[string]HogQLVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Query) SetVariables(v map[string]HogQLVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Query) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *Query) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Query) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetVersion

`func (o *Query) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Query) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Query) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Query) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Query) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Query) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAggregationGroupTypeIndex

`func (o *Query) GetAggregationGroupTypeIndex() int32`

GetAggregationGroupTypeIndex returns the AggregationGroupTypeIndex field if non-nil, zero value otherwise.

### GetAggregationGroupTypeIndexOk

`func (o *Query) GetAggregationGroupTypeIndexOk() (*int32, bool)`

GetAggregationGroupTypeIndexOk returns a tuple with the AggregationGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationGroupTypeIndex

`func (o *Query) SetAggregationGroupTypeIndex(v int32)`

SetAggregationGroupTypeIndex sets AggregationGroupTypeIndex field to given value.

### HasAggregationGroupTypeIndex

`func (o *Query) HasAggregationGroupTypeIndex() bool`

HasAggregationGroupTypeIndex returns a boolean if a field has been set.

### SetAggregationGroupTypeIndexNil

`func (o *Query) SetAggregationGroupTypeIndexNil(b bool)`

 SetAggregationGroupTypeIndexNil sets the value for AggregationGroupTypeIndex to be an explicit nil

### UnsetAggregationGroupTypeIndex
`func (o *Query) UnsetAggregationGroupTypeIndex()`

UnsetAggregationGroupTypeIndex ensures that no value is present for AggregationGroupTypeIndex, not even an explicit nil
### GetBreakdownFilter

`func (o *Query) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *Query) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *Query) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *Query) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetCompareFilter

`func (o *Query) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *Query) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *Query) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *Query) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *Query) GetConversionGoal() Conversiongoal`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *Query) GetConversionGoalOk() (*Conversiongoal, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *Query) SetConversionGoal(v Conversiongoal)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *Query) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *Query) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *Query) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDataColorTheme

`func (o *Query) GetDataColorTheme() float32`

GetDataColorTheme returns the DataColorTheme field if non-nil, zero value otherwise.

### GetDataColorThemeOk

`func (o *Query) GetDataColorThemeOk() (*float32, bool)`

GetDataColorThemeOk returns a tuple with the DataColorTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorTheme

`func (o *Query) SetDataColorTheme(v float32)`

SetDataColorTheme sets DataColorTheme field to given value.

### HasDataColorTheme

`func (o *Query) HasDataColorTheme() bool`

HasDataColorTheme returns a boolean if a field has been set.

### SetDataColorThemeNil

`func (o *Query) SetDataColorThemeNil(b bool)`

 SetDataColorThemeNil sets the value for DataColorTheme to be an explicit nil

### UnsetDataColorTheme
`func (o *Query) UnsetDataColorTheme()`

UnsetDataColorTheme ensures that no value is present for DataColorTheme, not even an explicit nil
### GetDateRange

`func (o *Query) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Query) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Query) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *Query) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetFilterTestAccounts

`func (o *Query) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *Query) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *Query) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *Query) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *Query) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *Query) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetInterval

`func (o *Query) GetInterval() IntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Query) GetIntervalOk() (*IntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Query) SetInterval(v IntervalType)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Query) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetProperties

`func (o *Query) GetProperties() Properties1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Query) GetPropertiesOk() (*Properties1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Query) SetProperties(v Properties1)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Query) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *Query) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Query) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSamplingFactor

`func (o *Query) GetSamplingFactor() float32`

GetSamplingFactor returns the SamplingFactor field if non-nil, zero value otherwise.

### GetSamplingFactorOk

`func (o *Query) GetSamplingFactorOk() (*float32, bool)`

GetSamplingFactorOk returns a tuple with the SamplingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingFactor

`func (o *Query) SetSamplingFactor(v float32)`

SetSamplingFactor sets SamplingFactor field to given value.

### HasSamplingFactor

`func (o *Query) HasSamplingFactor() bool`

HasSamplingFactor returns a boolean if a field has been set.

### SetSamplingFactorNil

`func (o *Query) SetSamplingFactorNil(b bool)`

 SetSamplingFactorNil sets the value for SamplingFactor to be an explicit nil

### UnsetSamplingFactor
`func (o *Query) UnsetSamplingFactor()`

UnsetSamplingFactor ensures that no value is present for SamplingFactor, not even an explicit nil
### GetSeries

`func (o *Query) GetSeries() []Series1Inner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Query) GetSeriesOk() (*[]Series1Inner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Query) SetSeries(v []Series1Inner)`

SetSeries sets Series field to given value.


### GetTrendsFilter

`func (o *Query) GetTrendsFilter() TrendsFilter`

GetTrendsFilter returns the TrendsFilter field if non-nil, zero value otherwise.

### GetTrendsFilterOk

`func (o *Query) GetTrendsFilterOk() (*TrendsFilter, bool)`

GetTrendsFilterOk returns a tuple with the TrendsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendsFilter

`func (o *Query) SetTrendsFilter(v TrendsFilter)`

SetTrendsFilter sets TrendsFilter field to given value.

### HasTrendsFilter

`func (o *Query) HasTrendsFilter() bool`

HasTrendsFilter returns a boolean if a field has been set.

### GetFunnelsFilter

`func (o *Query) GetFunnelsFilter() FunnelsFilter`

GetFunnelsFilter returns the FunnelsFilter field if non-nil, zero value otherwise.

### GetFunnelsFilterOk

`func (o *Query) GetFunnelsFilterOk() (*FunnelsFilter, bool)`

GetFunnelsFilterOk returns a tuple with the FunnelsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsFilter

`func (o *Query) SetFunnelsFilter(v FunnelsFilter)`

SetFunnelsFilter sets FunnelsFilter field to given value.

### HasFunnelsFilter

`func (o *Query) HasFunnelsFilter() bool`

HasFunnelsFilter returns a boolean if a field has been set.

### GetRetentionFilter

`func (o *Query) GetRetentionFilter() RetentionFilter`

GetRetentionFilter returns the RetentionFilter field if non-nil, zero value otherwise.

### GetRetentionFilterOk

`func (o *Query) GetRetentionFilterOk() (*RetentionFilter, bool)`

GetRetentionFilterOk returns a tuple with the RetentionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionFilter

`func (o *Query) SetRetentionFilter(v RetentionFilter)`

SetRetentionFilter sets RetentionFilter field to given value.


### GetFunnelPathsFilter

`func (o *Query) GetFunnelPathsFilter() FunnelPathsFilter`

GetFunnelPathsFilter returns the FunnelPathsFilter field if non-nil, zero value otherwise.

### GetFunnelPathsFilterOk

`func (o *Query) GetFunnelPathsFilterOk() (*FunnelPathsFilter, bool)`

GetFunnelPathsFilterOk returns a tuple with the FunnelPathsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelPathsFilter

`func (o *Query) SetFunnelPathsFilter(v FunnelPathsFilter)`

SetFunnelPathsFilter sets FunnelPathsFilter field to given value.

### HasFunnelPathsFilter

`func (o *Query) HasFunnelPathsFilter() bool`

HasFunnelPathsFilter returns a boolean if a field has been set.

### GetPathsFilter

`func (o *Query) GetPathsFilter() PathsFilter`

GetPathsFilter returns the PathsFilter field if non-nil, zero value otherwise.

### GetPathsFilterOk

`func (o *Query) GetPathsFilterOk() (*PathsFilter, bool)`

GetPathsFilterOk returns a tuple with the PathsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathsFilter

`func (o *Query) SetPathsFilter(v PathsFilter)`

SetPathsFilter sets PathsFilter field to given value.


### GetIntervalCount

`func (o *Query) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *Query) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *Query) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *Query) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *Query) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *Query) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetStickinessFilter

`func (o *Query) GetStickinessFilter() StickinessFilter`

GetStickinessFilter returns the StickinessFilter field if non-nil, zero value otherwise.

### GetStickinessFilterOk

`func (o *Query) GetStickinessFilterOk() (*StickinessFilter, bool)`

GetStickinessFilterOk returns a tuple with the StickinessFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickinessFilter

`func (o *Query) SetStickinessFilter(v StickinessFilter)`

SetStickinessFilter sets StickinessFilter field to given value.

### HasStickinessFilter

`func (o *Query) HasStickinessFilter() bool`

HasStickinessFilter returns a boolean if a field has been set.

### GetLifecycleFilter

`func (o *Query) GetLifecycleFilter() LifecycleFilter`

GetLifecycleFilter returns the LifecycleFilter field if non-nil, zero value otherwise.

### GetLifecycleFilterOk

`func (o *Query) GetLifecycleFilterOk() (*LifecycleFilter, bool)`

GetLifecycleFilterOk returns a tuple with the LifecycleFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleFilter

`func (o *Query) SetLifecycleFilter(v LifecycleFilter)`

SetLifecycleFilter sets LifecycleFilter field to given value.

### HasLifecycleFilter

`func (o *Query) HasLifecycleFilter() bool`

HasLifecycleFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


