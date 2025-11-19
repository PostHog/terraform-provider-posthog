# Source3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationGroupTypeIndex** | Pointer to **NullableInt32** | Groups aggregation | [optional] 
**BreakdownFilter** | Pointer to [**BreakdownFilter**](BreakdownFilter.md) |  | [optional] 
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal**](Conversiongoal.md) |  | [optional] [default to null]
**DataColorTheme** | Pointer to **NullableFloat32** | Colors used in the insight&#39;s visualization | [optional] 
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** | Exclude internal and test users by applying the respective filters | [optional] [default to false]
**Interval** | Pointer to [**IntervalType**](IntervalType.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "LifecycleQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | Pointer to [**NullableProperties1**](Properties1.md) |  | [optional] [default to []]
**Response** | Pointer to [**LifecycleQueryResponse**](LifecycleQueryResponse.md) |  | [optional] 
**SamplingFactor** | Pointer to **NullableFloat32** | Sampling rate | [optional] 
**Series** | [**[]Series1Inner**](Series1Inner.md) | Events and actions to include | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**TrendsFilter** | Pointer to [**TrendsFilter**](TrendsFilter.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**FunnelsFilter** | Pointer to [**FunnelsFilter**](FunnelsFilter.md) |  | [optional] 
**RetentionFilter** | [**RetentionFilter**](RetentionFilter.md) |  | 
**FunnelPathsFilter** | Pointer to [**FunnelPathsFilter**](FunnelPathsFilter.md) |  | [optional] 
**PathsFilter** | [**PathsFilter**](PathsFilter.md) |  | 
**IntervalCount** | Pointer to **NullableInt32** | How many intervals comprise a period. Only used for cohorts, otherwise default 1. | [optional] 
**StickinessFilter** | Pointer to [**StickinessFilter**](StickinessFilter.md) |  | [optional] 
**LifecycleFilter** | Pointer to [**LifecycleFilter**](LifecycleFilter.md) |  | [optional] 

## Methods

### NewSource3

`func NewSource3(series []Series1Inner, retentionFilter RetentionFilter, pathsFilter PathsFilter, ) *Source3`

NewSource3 instantiates a new Source3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSource3WithDefaults

`func NewSource3WithDefaults() *Source3`

NewSource3WithDefaults instantiates a new Source3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationGroupTypeIndex

`func (o *Source3) GetAggregationGroupTypeIndex() int32`

GetAggregationGroupTypeIndex returns the AggregationGroupTypeIndex field if non-nil, zero value otherwise.

### GetAggregationGroupTypeIndexOk

`func (o *Source3) GetAggregationGroupTypeIndexOk() (*int32, bool)`

GetAggregationGroupTypeIndexOk returns a tuple with the AggregationGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationGroupTypeIndex

`func (o *Source3) SetAggregationGroupTypeIndex(v int32)`

SetAggregationGroupTypeIndex sets AggregationGroupTypeIndex field to given value.

### HasAggregationGroupTypeIndex

`func (o *Source3) HasAggregationGroupTypeIndex() bool`

HasAggregationGroupTypeIndex returns a boolean if a field has been set.

### SetAggregationGroupTypeIndexNil

`func (o *Source3) SetAggregationGroupTypeIndexNil(b bool)`

 SetAggregationGroupTypeIndexNil sets the value for AggregationGroupTypeIndex to be an explicit nil

### UnsetAggregationGroupTypeIndex
`func (o *Source3) UnsetAggregationGroupTypeIndex()`

UnsetAggregationGroupTypeIndex ensures that no value is present for AggregationGroupTypeIndex, not even an explicit nil
### GetBreakdownFilter

`func (o *Source3) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *Source3) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *Source3) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *Source3) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetCompareFilter

`func (o *Source3) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *Source3) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *Source3) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *Source3) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *Source3) GetConversionGoal() Conversiongoal`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *Source3) GetConversionGoalOk() (*Conversiongoal, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *Source3) SetConversionGoal(v Conversiongoal)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *Source3) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *Source3) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *Source3) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDataColorTheme

`func (o *Source3) GetDataColorTheme() float32`

GetDataColorTheme returns the DataColorTheme field if non-nil, zero value otherwise.

### GetDataColorThemeOk

`func (o *Source3) GetDataColorThemeOk() (*float32, bool)`

GetDataColorThemeOk returns a tuple with the DataColorTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorTheme

`func (o *Source3) SetDataColorTheme(v float32)`

SetDataColorTheme sets DataColorTheme field to given value.

### HasDataColorTheme

`func (o *Source3) HasDataColorTheme() bool`

HasDataColorTheme returns a boolean if a field has been set.

### SetDataColorThemeNil

`func (o *Source3) SetDataColorThemeNil(b bool)`

 SetDataColorThemeNil sets the value for DataColorTheme to be an explicit nil

### UnsetDataColorTheme
`func (o *Source3) UnsetDataColorTheme()`

UnsetDataColorTheme ensures that no value is present for DataColorTheme, not even an explicit nil
### GetDateRange

`func (o *Source3) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *Source3) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *Source3) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *Source3) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetFilterTestAccounts

`func (o *Source3) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *Source3) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *Source3) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *Source3) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *Source3) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *Source3) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetInterval

`func (o *Source3) GetInterval() IntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Source3) GetIntervalOk() (*IntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Source3) SetInterval(v IntervalType)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Source3) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetKind

`func (o *Source3) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Source3) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Source3) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Source3) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *Source3) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *Source3) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *Source3) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *Source3) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *Source3) GetProperties() Properties1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Source3) GetPropertiesOk() (*Properties1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Source3) SetProperties(v Properties1)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Source3) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *Source3) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *Source3) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *Source3) GetResponse() LifecycleQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Source3) GetResponseOk() (*LifecycleQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Source3) SetResponse(v LifecycleQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Source3) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSamplingFactor

`func (o *Source3) GetSamplingFactor() float32`

GetSamplingFactor returns the SamplingFactor field if non-nil, zero value otherwise.

### GetSamplingFactorOk

`func (o *Source3) GetSamplingFactorOk() (*float32, bool)`

GetSamplingFactorOk returns a tuple with the SamplingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingFactor

`func (o *Source3) SetSamplingFactor(v float32)`

SetSamplingFactor sets SamplingFactor field to given value.

### HasSamplingFactor

`func (o *Source3) HasSamplingFactor() bool`

HasSamplingFactor returns a boolean if a field has been set.

### SetSamplingFactorNil

`func (o *Source3) SetSamplingFactorNil(b bool)`

 SetSamplingFactorNil sets the value for SamplingFactor to be an explicit nil

### UnsetSamplingFactor
`func (o *Source3) UnsetSamplingFactor()`

UnsetSamplingFactor ensures that no value is present for SamplingFactor, not even an explicit nil
### GetSeries

`func (o *Source3) GetSeries() []Series1Inner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Source3) GetSeriesOk() (*[]Series1Inner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Source3) SetSeries(v []Series1Inner)`

SetSeries sets Series field to given value.


### GetTags

`func (o *Source3) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Source3) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Source3) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Source3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTrendsFilter

`func (o *Source3) GetTrendsFilter() TrendsFilter`

GetTrendsFilter returns the TrendsFilter field if non-nil, zero value otherwise.

### GetTrendsFilterOk

`func (o *Source3) GetTrendsFilterOk() (*TrendsFilter, bool)`

GetTrendsFilterOk returns a tuple with the TrendsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendsFilter

`func (o *Source3) SetTrendsFilter(v TrendsFilter)`

SetTrendsFilter sets TrendsFilter field to given value.

### HasTrendsFilter

`func (o *Source3) HasTrendsFilter() bool`

HasTrendsFilter returns a boolean if a field has been set.

### GetVersion

`func (o *Source3) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Source3) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Source3) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Source3) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Source3) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Source3) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFunnelsFilter

`func (o *Source3) GetFunnelsFilter() FunnelsFilter`

GetFunnelsFilter returns the FunnelsFilter field if non-nil, zero value otherwise.

### GetFunnelsFilterOk

`func (o *Source3) GetFunnelsFilterOk() (*FunnelsFilter, bool)`

GetFunnelsFilterOk returns a tuple with the FunnelsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsFilter

`func (o *Source3) SetFunnelsFilter(v FunnelsFilter)`

SetFunnelsFilter sets FunnelsFilter field to given value.

### HasFunnelsFilter

`func (o *Source3) HasFunnelsFilter() bool`

HasFunnelsFilter returns a boolean if a field has been set.

### GetRetentionFilter

`func (o *Source3) GetRetentionFilter() RetentionFilter`

GetRetentionFilter returns the RetentionFilter field if non-nil, zero value otherwise.

### GetRetentionFilterOk

`func (o *Source3) GetRetentionFilterOk() (*RetentionFilter, bool)`

GetRetentionFilterOk returns a tuple with the RetentionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionFilter

`func (o *Source3) SetRetentionFilter(v RetentionFilter)`

SetRetentionFilter sets RetentionFilter field to given value.


### GetFunnelPathsFilter

`func (o *Source3) GetFunnelPathsFilter() FunnelPathsFilter`

GetFunnelPathsFilter returns the FunnelPathsFilter field if non-nil, zero value otherwise.

### GetFunnelPathsFilterOk

`func (o *Source3) GetFunnelPathsFilterOk() (*FunnelPathsFilter, bool)`

GetFunnelPathsFilterOk returns a tuple with the FunnelPathsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelPathsFilter

`func (o *Source3) SetFunnelPathsFilter(v FunnelPathsFilter)`

SetFunnelPathsFilter sets FunnelPathsFilter field to given value.

### HasFunnelPathsFilter

`func (o *Source3) HasFunnelPathsFilter() bool`

HasFunnelPathsFilter returns a boolean if a field has been set.

### GetPathsFilter

`func (o *Source3) GetPathsFilter() PathsFilter`

GetPathsFilter returns the PathsFilter field if non-nil, zero value otherwise.

### GetPathsFilterOk

`func (o *Source3) GetPathsFilterOk() (*PathsFilter, bool)`

GetPathsFilterOk returns a tuple with the PathsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathsFilter

`func (o *Source3) SetPathsFilter(v PathsFilter)`

SetPathsFilter sets PathsFilter field to given value.


### GetIntervalCount

`func (o *Source3) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *Source3) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *Source3) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *Source3) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *Source3) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *Source3) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetStickinessFilter

`func (o *Source3) GetStickinessFilter() StickinessFilter`

GetStickinessFilter returns the StickinessFilter field if non-nil, zero value otherwise.

### GetStickinessFilterOk

`func (o *Source3) GetStickinessFilterOk() (*StickinessFilter, bool)`

GetStickinessFilterOk returns a tuple with the StickinessFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickinessFilter

`func (o *Source3) SetStickinessFilter(v StickinessFilter)`

SetStickinessFilter sets StickinessFilter field to given value.

### HasStickinessFilter

`func (o *Source3) HasStickinessFilter() bool`

HasStickinessFilter returns a boolean if a field has been set.

### GetLifecycleFilter

`func (o *Source3) GetLifecycleFilter() LifecycleFilter`

GetLifecycleFilter returns the LifecycleFilter field if non-nil, zero value otherwise.

### GetLifecycleFilterOk

`func (o *Source3) GetLifecycleFilterOk() (*LifecycleFilter, bool)`

GetLifecycleFilterOk returns a tuple with the LifecycleFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleFilter

`func (o *Source3) SetLifecycleFilter(v LifecycleFilter)`

SetLifecycleFilter sets LifecycleFilter field to given value.

### HasLifecycleFilter

`func (o *Source3) HasLifecycleFilter() bool`

HasLifecycleFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


