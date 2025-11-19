# WebExternalClicksTableQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal1**](Conversiongoal1.md) |  | [optional] [default to null]
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "WebExternalClicksTableQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**OrderBy** | Pointer to **[]string** |  | [optional] 
**Properties** | [**[]Properties2Inner1**](Properties2Inner1.md) |  | 
**Response** | Pointer to [**WebExternalClicksTableQueryResponse**](WebExternalClicksTableQueryResponse.md) |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**StripQueryParams** | Pointer to **NullableBool** |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewWebExternalClicksTableQuery

`func NewWebExternalClicksTableQuery(properties []Properties2Inner1, ) *WebExternalClicksTableQuery`

NewWebExternalClicksTableQuery instantiates a new WebExternalClicksTableQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebExternalClicksTableQueryWithDefaults

`func NewWebExternalClicksTableQueryWithDefaults() *WebExternalClicksTableQuery`

NewWebExternalClicksTableQueryWithDefaults instantiates a new WebExternalClicksTableQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareFilter

`func (o *WebExternalClicksTableQuery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *WebExternalClicksTableQuery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *WebExternalClicksTableQuery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *WebExternalClicksTableQuery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *WebExternalClicksTableQuery) GetConversionGoal() Conversiongoal1`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *WebExternalClicksTableQuery) GetConversionGoalOk() (*Conversiongoal1, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *WebExternalClicksTableQuery) SetConversionGoal(v Conversiongoal1)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *WebExternalClicksTableQuery) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *WebExternalClicksTableQuery) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *WebExternalClicksTableQuery) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDateRange

`func (o *WebExternalClicksTableQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *WebExternalClicksTableQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *WebExternalClicksTableQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *WebExternalClicksTableQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDoPathCleaning

`func (o *WebExternalClicksTableQuery) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *WebExternalClicksTableQuery) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *WebExternalClicksTableQuery) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *WebExternalClicksTableQuery) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *WebExternalClicksTableQuery) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *WebExternalClicksTableQuery) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetFilterTestAccounts

`func (o *WebExternalClicksTableQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *WebExternalClicksTableQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *WebExternalClicksTableQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *WebExternalClicksTableQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *WebExternalClicksTableQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *WebExternalClicksTableQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetIncludeRevenue

`func (o *WebExternalClicksTableQuery) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *WebExternalClicksTableQuery) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *WebExternalClicksTableQuery) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *WebExternalClicksTableQuery) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *WebExternalClicksTableQuery) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *WebExternalClicksTableQuery) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetKind

`func (o *WebExternalClicksTableQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebExternalClicksTableQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebExternalClicksTableQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WebExternalClicksTableQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *WebExternalClicksTableQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WebExternalClicksTableQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WebExternalClicksTableQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WebExternalClicksTableQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *WebExternalClicksTableQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *WebExternalClicksTableQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *WebExternalClicksTableQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *WebExternalClicksTableQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *WebExternalClicksTableQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *WebExternalClicksTableQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOrderBy

`func (o *WebExternalClicksTableQuery) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *WebExternalClicksTableQuery) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *WebExternalClicksTableQuery) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *WebExternalClicksTableQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *WebExternalClicksTableQuery) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *WebExternalClicksTableQuery) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetProperties

`func (o *WebExternalClicksTableQuery) GetProperties() []Properties2Inner1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WebExternalClicksTableQuery) GetPropertiesOk() (*[]Properties2Inner1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WebExternalClicksTableQuery) SetProperties(v []Properties2Inner1)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *WebExternalClicksTableQuery) GetResponse() WebExternalClicksTableQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebExternalClicksTableQuery) GetResponseOk() (*WebExternalClicksTableQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebExternalClicksTableQuery) SetResponse(v WebExternalClicksTableQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebExternalClicksTableQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSampling

`func (o *WebExternalClicksTableQuery) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *WebExternalClicksTableQuery) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *WebExternalClicksTableQuery) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *WebExternalClicksTableQuery) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetStripQueryParams

`func (o *WebExternalClicksTableQuery) GetStripQueryParams() bool`

GetStripQueryParams returns the StripQueryParams field if non-nil, zero value otherwise.

### GetStripQueryParamsOk

`func (o *WebExternalClicksTableQuery) GetStripQueryParamsOk() (*bool, bool)`

GetStripQueryParamsOk returns a tuple with the StripQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripQueryParams

`func (o *WebExternalClicksTableQuery) SetStripQueryParams(v bool)`

SetStripQueryParams sets StripQueryParams field to given value.

### HasStripQueryParams

`func (o *WebExternalClicksTableQuery) HasStripQueryParams() bool`

HasStripQueryParams returns a boolean if a field has been set.

### SetStripQueryParamsNil

`func (o *WebExternalClicksTableQuery) SetStripQueryParamsNil(b bool)`

 SetStripQueryParamsNil sets the value for StripQueryParams to be an explicit nil

### UnsetStripQueryParams
`func (o *WebExternalClicksTableQuery) UnsetStripQueryParams()`

UnsetStripQueryParams ensures that no value is present for StripQueryParams, not even an explicit nil
### GetTags

`func (o *WebExternalClicksTableQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WebExternalClicksTableQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WebExternalClicksTableQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WebExternalClicksTableQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *WebExternalClicksTableQuery) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *WebExternalClicksTableQuery) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *WebExternalClicksTableQuery) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *WebExternalClicksTableQuery) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *WebExternalClicksTableQuery) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *WebExternalClicksTableQuery) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetVersion

`func (o *WebExternalClicksTableQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebExternalClicksTableQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebExternalClicksTableQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WebExternalClicksTableQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *WebExternalClicksTableQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *WebExternalClicksTableQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


