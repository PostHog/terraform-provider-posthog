# WebVitalsPathBreakdownQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal1**](Conversiongoal1.md) |  | [optional] [default to null]
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "WebVitalsPathBreakdownQuery"]
**Metric** | [**WebVitalsMetric**](WebVitalsMetric.md) |  | 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**OrderBy** | Pointer to **[]string** |  | [optional] 
**Percentile** | [**WebVitalsPercentile**](WebVitalsPercentile.md) |  | 
**Properties** | [**[]Properties2Inner1**](Properties2Inner1.md) |  | 
**Response** | Pointer to [**WebVitalsPathBreakdownQueryResponse**](WebVitalsPathBreakdownQueryResponse.md) |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Thresholds** | **[]float32** |  | 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewWebVitalsPathBreakdownQuery

`func NewWebVitalsPathBreakdownQuery(metric WebVitalsMetric, percentile WebVitalsPercentile, properties []Properties2Inner1, thresholds []float32, ) *WebVitalsPathBreakdownQuery`

NewWebVitalsPathBreakdownQuery instantiates a new WebVitalsPathBreakdownQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebVitalsPathBreakdownQueryWithDefaults

`func NewWebVitalsPathBreakdownQueryWithDefaults() *WebVitalsPathBreakdownQuery`

NewWebVitalsPathBreakdownQueryWithDefaults instantiates a new WebVitalsPathBreakdownQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareFilter

`func (o *WebVitalsPathBreakdownQuery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *WebVitalsPathBreakdownQuery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *WebVitalsPathBreakdownQuery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *WebVitalsPathBreakdownQuery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *WebVitalsPathBreakdownQuery) GetConversionGoal() Conversiongoal1`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *WebVitalsPathBreakdownQuery) GetConversionGoalOk() (*Conversiongoal1, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *WebVitalsPathBreakdownQuery) SetConversionGoal(v Conversiongoal1)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *WebVitalsPathBreakdownQuery) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *WebVitalsPathBreakdownQuery) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *WebVitalsPathBreakdownQuery) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDateRange

`func (o *WebVitalsPathBreakdownQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *WebVitalsPathBreakdownQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *WebVitalsPathBreakdownQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *WebVitalsPathBreakdownQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDoPathCleaning

`func (o *WebVitalsPathBreakdownQuery) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *WebVitalsPathBreakdownQuery) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *WebVitalsPathBreakdownQuery) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *WebVitalsPathBreakdownQuery) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *WebVitalsPathBreakdownQuery) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *WebVitalsPathBreakdownQuery) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetFilterTestAccounts

`func (o *WebVitalsPathBreakdownQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *WebVitalsPathBreakdownQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *WebVitalsPathBreakdownQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *WebVitalsPathBreakdownQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *WebVitalsPathBreakdownQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *WebVitalsPathBreakdownQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetIncludeRevenue

`func (o *WebVitalsPathBreakdownQuery) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *WebVitalsPathBreakdownQuery) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *WebVitalsPathBreakdownQuery) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *WebVitalsPathBreakdownQuery) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *WebVitalsPathBreakdownQuery) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *WebVitalsPathBreakdownQuery) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetKind

`func (o *WebVitalsPathBreakdownQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebVitalsPathBreakdownQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebVitalsPathBreakdownQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WebVitalsPathBreakdownQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetric

`func (o *WebVitalsPathBreakdownQuery) GetMetric() WebVitalsMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *WebVitalsPathBreakdownQuery) GetMetricOk() (*WebVitalsMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *WebVitalsPathBreakdownQuery) SetMetric(v WebVitalsMetric)`

SetMetric sets Metric field to given value.


### GetModifiers

`func (o *WebVitalsPathBreakdownQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *WebVitalsPathBreakdownQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *WebVitalsPathBreakdownQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *WebVitalsPathBreakdownQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOrderBy

`func (o *WebVitalsPathBreakdownQuery) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *WebVitalsPathBreakdownQuery) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *WebVitalsPathBreakdownQuery) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *WebVitalsPathBreakdownQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *WebVitalsPathBreakdownQuery) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *WebVitalsPathBreakdownQuery) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetPercentile

`func (o *WebVitalsPathBreakdownQuery) GetPercentile() WebVitalsPercentile`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *WebVitalsPathBreakdownQuery) GetPercentileOk() (*WebVitalsPercentile, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *WebVitalsPathBreakdownQuery) SetPercentile(v WebVitalsPercentile)`

SetPercentile sets Percentile field to given value.


### GetProperties

`func (o *WebVitalsPathBreakdownQuery) GetProperties() []Properties2Inner1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WebVitalsPathBreakdownQuery) GetPropertiesOk() (*[]Properties2Inner1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WebVitalsPathBreakdownQuery) SetProperties(v []Properties2Inner1)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *WebVitalsPathBreakdownQuery) GetResponse() WebVitalsPathBreakdownQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebVitalsPathBreakdownQuery) GetResponseOk() (*WebVitalsPathBreakdownQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebVitalsPathBreakdownQuery) SetResponse(v WebVitalsPathBreakdownQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebVitalsPathBreakdownQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSampling

`func (o *WebVitalsPathBreakdownQuery) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *WebVitalsPathBreakdownQuery) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *WebVitalsPathBreakdownQuery) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *WebVitalsPathBreakdownQuery) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetTags

`func (o *WebVitalsPathBreakdownQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WebVitalsPathBreakdownQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WebVitalsPathBreakdownQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WebVitalsPathBreakdownQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThresholds

`func (o *WebVitalsPathBreakdownQuery) GetThresholds() []float32`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *WebVitalsPathBreakdownQuery) GetThresholdsOk() (*[]float32, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *WebVitalsPathBreakdownQuery) SetThresholds(v []float32)`

SetThresholds sets Thresholds field to given value.


### GetUseSessionsTable

`func (o *WebVitalsPathBreakdownQuery) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *WebVitalsPathBreakdownQuery) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *WebVitalsPathBreakdownQuery) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *WebVitalsPathBreakdownQuery) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *WebVitalsPathBreakdownQuery) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *WebVitalsPathBreakdownQuery) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetVersion

`func (o *WebVitalsPathBreakdownQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebVitalsPathBreakdownQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebVitalsPathBreakdownQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WebVitalsPathBreakdownQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *WebVitalsPathBreakdownQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *WebVitalsPathBreakdownQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


