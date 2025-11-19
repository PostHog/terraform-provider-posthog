# TrendsQuery

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
**Kind** | Pointer to **string** |  | [optional] [default to "TrendsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | Pointer to [**NullableProperties1**](Properties1.md) |  | [optional] [default to []]
**Response** | Pointer to [**TrendsQueryResponse**](TrendsQueryResponse.md) |  | [optional] 
**SamplingFactor** | Pointer to **NullableFloat32** | Sampling rate | [optional] 
**Series** | [**[]Series1Inner**](Series1Inner.md) | Events and actions to include | 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**TrendsFilter** | Pointer to [**TrendsFilter**](TrendsFilter.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewTrendsQuery

`func NewTrendsQuery(series []Series1Inner, ) *TrendsQuery`

NewTrendsQuery instantiates a new TrendsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendsQueryWithDefaults

`func NewTrendsQueryWithDefaults() *TrendsQuery`

NewTrendsQueryWithDefaults instantiates a new TrendsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationGroupTypeIndex

`func (o *TrendsQuery) GetAggregationGroupTypeIndex() int32`

GetAggregationGroupTypeIndex returns the AggregationGroupTypeIndex field if non-nil, zero value otherwise.

### GetAggregationGroupTypeIndexOk

`func (o *TrendsQuery) GetAggregationGroupTypeIndexOk() (*int32, bool)`

GetAggregationGroupTypeIndexOk returns a tuple with the AggregationGroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationGroupTypeIndex

`func (o *TrendsQuery) SetAggregationGroupTypeIndex(v int32)`

SetAggregationGroupTypeIndex sets AggregationGroupTypeIndex field to given value.

### HasAggregationGroupTypeIndex

`func (o *TrendsQuery) HasAggregationGroupTypeIndex() bool`

HasAggregationGroupTypeIndex returns a boolean if a field has been set.

### SetAggregationGroupTypeIndexNil

`func (o *TrendsQuery) SetAggregationGroupTypeIndexNil(b bool)`

 SetAggregationGroupTypeIndexNil sets the value for AggregationGroupTypeIndex to be an explicit nil

### UnsetAggregationGroupTypeIndex
`func (o *TrendsQuery) UnsetAggregationGroupTypeIndex()`

UnsetAggregationGroupTypeIndex ensures that no value is present for AggregationGroupTypeIndex, not even an explicit nil
### GetBreakdownFilter

`func (o *TrendsQuery) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *TrendsQuery) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *TrendsQuery) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *TrendsQuery) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetCompareFilter

`func (o *TrendsQuery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *TrendsQuery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *TrendsQuery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *TrendsQuery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *TrendsQuery) GetConversionGoal() Conversiongoal`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *TrendsQuery) GetConversionGoalOk() (*Conversiongoal, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *TrendsQuery) SetConversionGoal(v Conversiongoal)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *TrendsQuery) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *TrendsQuery) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *TrendsQuery) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDataColorTheme

`func (o *TrendsQuery) GetDataColorTheme() float32`

GetDataColorTheme returns the DataColorTheme field if non-nil, zero value otherwise.

### GetDataColorThemeOk

`func (o *TrendsQuery) GetDataColorThemeOk() (*float32, bool)`

GetDataColorThemeOk returns a tuple with the DataColorTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorTheme

`func (o *TrendsQuery) SetDataColorTheme(v float32)`

SetDataColorTheme sets DataColorTheme field to given value.

### HasDataColorTheme

`func (o *TrendsQuery) HasDataColorTheme() bool`

HasDataColorTheme returns a boolean if a field has been set.

### SetDataColorThemeNil

`func (o *TrendsQuery) SetDataColorThemeNil(b bool)`

 SetDataColorThemeNil sets the value for DataColorTheme to be an explicit nil

### UnsetDataColorTheme
`func (o *TrendsQuery) UnsetDataColorTheme()`

UnsetDataColorTheme ensures that no value is present for DataColorTheme, not even an explicit nil
### GetDateRange

`func (o *TrendsQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *TrendsQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *TrendsQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *TrendsQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetFilterTestAccounts

`func (o *TrendsQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *TrendsQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *TrendsQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *TrendsQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *TrendsQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *TrendsQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetInterval

`func (o *TrendsQuery) GetInterval() IntervalType`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TrendsQuery) GetIntervalOk() (*IntervalType, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TrendsQuery) SetInterval(v IntervalType)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *TrendsQuery) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetKind

`func (o *TrendsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrendsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrendsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrendsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *TrendsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *TrendsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *TrendsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *TrendsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *TrendsQuery) GetProperties() Properties1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TrendsQuery) GetPropertiesOk() (*Properties1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TrendsQuery) SetProperties(v Properties1)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TrendsQuery) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *TrendsQuery) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *TrendsQuery) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetResponse

`func (o *TrendsQuery) GetResponse() TrendsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *TrendsQuery) GetResponseOk() (*TrendsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *TrendsQuery) SetResponse(v TrendsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *TrendsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSamplingFactor

`func (o *TrendsQuery) GetSamplingFactor() float32`

GetSamplingFactor returns the SamplingFactor field if non-nil, zero value otherwise.

### GetSamplingFactorOk

`func (o *TrendsQuery) GetSamplingFactorOk() (*float32, bool)`

GetSamplingFactorOk returns a tuple with the SamplingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingFactor

`func (o *TrendsQuery) SetSamplingFactor(v float32)`

SetSamplingFactor sets SamplingFactor field to given value.

### HasSamplingFactor

`func (o *TrendsQuery) HasSamplingFactor() bool`

HasSamplingFactor returns a boolean if a field has been set.

### SetSamplingFactorNil

`func (o *TrendsQuery) SetSamplingFactorNil(b bool)`

 SetSamplingFactorNil sets the value for SamplingFactor to be an explicit nil

### UnsetSamplingFactor
`func (o *TrendsQuery) UnsetSamplingFactor()`

UnsetSamplingFactor ensures that no value is present for SamplingFactor, not even an explicit nil
### GetSeries

`func (o *TrendsQuery) GetSeries() []Series1Inner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *TrendsQuery) GetSeriesOk() (*[]Series1Inner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *TrendsQuery) SetSeries(v []Series1Inner)`

SetSeries sets Series field to given value.


### GetTags

`func (o *TrendsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TrendsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TrendsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TrendsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTrendsFilter

`func (o *TrendsQuery) GetTrendsFilter() TrendsFilter`

GetTrendsFilter returns the TrendsFilter field if non-nil, zero value otherwise.

### GetTrendsFilterOk

`func (o *TrendsQuery) GetTrendsFilterOk() (*TrendsFilter, bool)`

GetTrendsFilterOk returns a tuple with the TrendsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrendsFilter

`func (o *TrendsQuery) SetTrendsFilter(v TrendsFilter)`

SetTrendsFilter sets TrendsFilter field to given value.

### HasTrendsFilter

`func (o *TrendsQuery) HasTrendsFilter() bool`

HasTrendsFilter returns a boolean if a field has been set.

### GetVersion

`func (o *TrendsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrendsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrendsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrendsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TrendsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TrendsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


