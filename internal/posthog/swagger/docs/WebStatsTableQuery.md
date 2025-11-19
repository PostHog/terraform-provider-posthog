# WebStatsTableQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakdownBy** | [**WebStatsBreakdown**](WebStatsBreakdown.md) |  | 
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal1**](Conversiongoal1.md) |  | [optional] [default to null]
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**IncludeBounceRate** | Pointer to **NullableBool** |  | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**IncludeScrollDepth** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "WebStatsTableQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**OrderBy** | Pointer to **[]string** |  | [optional] 
**Properties** | [**[]Properties2Inner1**](Properties2Inner1.md) |  | 
**Response** | Pointer to [**WebStatsTableQueryResponse**](WebStatsTableQueryResponse.md) |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewWebStatsTableQuery

`func NewWebStatsTableQuery(breakdownBy WebStatsBreakdown, properties []Properties2Inner1, ) *WebStatsTableQuery`

NewWebStatsTableQuery instantiates a new WebStatsTableQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebStatsTableQueryWithDefaults

`func NewWebStatsTableQueryWithDefaults() *WebStatsTableQuery`

NewWebStatsTableQueryWithDefaults instantiates a new WebStatsTableQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownBy

`func (o *WebStatsTableQuery) GetBreakdownBy() WebStatsBreakdown`

GetBreakdownBy returns the BreakdownBy field if non-nil, zero value otherwise.

### GetBreakdownByOk

`func (o *WebStatsTableQuery) GetBreakdownByOk() (*WebStatsBreakdown, bool)`

GetBreakdownByOk returns a tuple with the BreakdownBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownBy

`func (o *WebStatsTableQuery) SetBreakdownBy(v WebStatsBreakdown)`

SetBreakdownBy sets BreakdownBy field to given value.


### GetCompareFilter

`func (o *WebStatsTableQuery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *WebStatsTableQuery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *WebStatsTableQuery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *WebStatsTableQuery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *WebStatsTableQuery) GetConversionGoal() Conversiongoal1`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *WebStatsTableQuery) GetConversionGoalOk() (*Conversiongoal1, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *WebStatsTableQuery) SetConversionGoal(v Conversiongoal1)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *WebStatsTableQuery) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *WebStatsTableQuery) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *WebStatsTableQuery) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDateRange

`func (o *WebStatsTableQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *WebStatsTableQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *WebStatsTableQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *WebStatsTableQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDoPathCleaning

`func (o *WebStatsTableQuery) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *WebStatsTableQuery) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *WebStatsTableQuery) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *WebStatsTableQuery) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *WebStatsTableQuery) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *WebStatsTableQuery) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetFilterTestAccounts

`func (o *WebStatsTableQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *WebStatsTableQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *WebStatsTableQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *WebStatsTableQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *WebStatsTableQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *WebStatsTableQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetIncludeBounceRate

`func (o *WebStatsTableQuery) GetIncludeBounceRate() bool`

GetIncludeBounceRate returns the IncludeBounceRate field if non-nil, zero value otherwise.

### GetIncludeBounceRateOk

`func (o *WebStatsTableQuery) GetIncludeBounceRateOk() (*bool, bool)`

GetIncludeBounceRateOk returns a tuple with the IncludeBounceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBounceRate

`func (o *WebStatsTableQuery) SetIncludeBounceRate(v bool)`

SetIncludeBounceRate sets IncludeBounceRate field to given value.

### HasIncludeBounceRate

`func (o *WebStatsTableQuery) HasIncludeBounceRate() bool`

HasIncludeBounceRate returns a boolean if a field has been set.

### SetIncludeBounceRateNil

`func (o *WebStatsTableQuery) SetIncludeBounceRateNil(b bool)`

 SetIncludeBounceRateNil sets the value for IncludeBounceRate to be an explicit nil

### UnsetIncludeBounceRate
`func (o *WebStatsTableQuery) UnsetIncludeBounceRate()`

UnsetIncludeBounceRate ensures that no value is present for IncludeBounceRate, not even an explicit nil
### GetIncludeRevenue

`func (o *WebStatsTableQuery) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *WebStatsTableQuery) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *WebStatsTableQuery) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *WebStatsTableQuery) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *WebStatsTableQuery) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *WebStatsTableQuery) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetIncludeScrollDepth

`func (o *WebStatsTableQuery) GetIncludeScrollDepth() bool`

GetIncludeScrollDepth returns the IncludeScrollDepth field if non-nil, zero value otherwise.

### GetIncludeScrollDepthOk

`func (o *WebStatsTableQuery) GetIncludeScrollDepthOk() (*bool, bool)`

GetIncludeScrollDepthOk returns a tuple with the IncludeScrollDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeScrollDepth

`func (o *WebStatsTableQuery) SetIncludeScrollDepth(v bool)`

SetIncludeScrollDepth sets IncludeScrollDepth field to given value.

### HasIncludeScrollDepth

`func (o *WebStatsTableQuery) HasIncludeScrollDepth() bool`

HasIncludeScrollDepth returns a boolean if a field has been set.

### SetIncludeScrollDepthNil

`func (o *WebStatsTableQuery) SetIncludeScrollDepthNil(b bool)`

 SetIncludeScrollDepthNil sets the value for IncludeScrollDepth to be an explicit nil

### UnsetIncludeScrollDepth
`func (o *WebStatsTableQuery) UnsetIncludeScrollDepth()`

UnsetIncludeScrollDepth ensures that no value is present for IncludeScrollDepth, not even an explicit nil
### GetKind

`func (o *WebStatsTableQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebStatsTableQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebStatsTableQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WebStatsTableQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *WebStatsTableQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WebStatsTableQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WebStatsTableQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WebStatsTableQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *WebStatsTableQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *WebStatsTableQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *WebStatsTableQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *WebStatsTableQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *WebStatsTableQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *WebStatsTableQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *WebStatsTableQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *WebStatsTableQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *WebStatsTableQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *WebStatsTableQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *WebStatsTableQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *WebStatsTableQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOrderBy

`func (o *WebStatsTableQuery) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *WebStatsTableQuery) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *WebStatsTableQuery) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *WebStatsTableQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *WebStatsTableQuery) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *WebStatsTableQuery) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetProperties

`func (o *WebStatsTableQuery) GetProperties() []Properties2Inner1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WebStatsTableQuery) GetPropertiesOk() (*[]Properties2Inner1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WebStatsTableQuery) SetProperties(v []Properties2Inner1)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *WebStatsTableQuery) GetResponse() WebStatsTableQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebStatsTableQuery) GetResponseOk() (*WebStatsTableQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebStatsTableQuery) SetResponse(v WebStatsTableQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebStatsTableQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSampling

`func (o *WebStatsTableQuery) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *WebStatsTableQuery) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *WebStatsTableQuery) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *WebStatsTableQuery) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetTags

`func (o *WebStatsTableQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WebStatsTableQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WebStatsTableQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WebStatsTableQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *WebStatsTableQuery) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *WebStatsTableQuery) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *WebStatsTableQuery) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *WebStatsTableQuery) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *WebStatsTableQuery) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *WebStatsTableQuery) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetVersion

`func (o *WebStatsTableQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebStatsTableQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebStatsTableQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WebStatsTableQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *WebStatsTableQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *WebStatsTableQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


