# WebPageURLSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal1**](Conversiongoal1.md) |  | [optional] [default to null]
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "WebPageURLSearchQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**OrderBy** | Pointer to **[]string** |  | [optional] 
**Properties** | [**[]Properties2Inner1**](Properties2Inner1.md) |  | 
**Response** | Pointer to [**WebPageURLSearchQueryResponse**](WebPageURLSearchQueryResponse.md) |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**SearchTerm** | Pointer to **NullableString** |  | [optional] 
**StripQueryParams** | Pointer to **NullableBool** |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewWebPageURLSearchQuery

`func NewWebPageURLSearchQuery(properties []Properties2Inner1, ) *WebPageURLSearchQuery`

NewWebPageURLSearchQuery instantiates a new WebPageURLSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPageURLSearchQueryWithDefaults

`func NewWebPageURLSearchQueryWithDefaults() *WebPageURLSearchQuery`

NewWebPageURLSearchQueryWithDefaults instantiates a new WebPageURLSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareFilter

`func (o *WebPageURLSearchQuery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *WebPageURLSearchQuery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *WebPageURLSearchQuery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *WebPageURLSearchQuery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *WebPageURLSearchQuery) GetConversionGoal() Conversiongoal1`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *WebPageURLSearchQuery) GetConversionGoalOk() (*Conversiongoal1, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *WebPageURLSearchQuery) SetConversionGoal(v Conversiongoal1)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *WebPageURLSearchQuery) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *WebPageURLSearchQuery) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *WebPageURLSearchQuery) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDateRange

`func (o *WebPageURLSearchQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *WebPageURLSearchQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *WebPageURLSearchQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *WebPageURLSearchQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDoPathCleaning

`func (o *WebPageURLSearchQuery) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *WebPageURLSearchQuery) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *WebPageURLSearchQuery) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *WebPageURLSearchQuery) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *WebPageURLSearchQuery) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *WebPageURLSearchQuery) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetFilterTestAccounts

`func (o *WebPageURLSearchQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *WebPageURLSearchQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *WebPageURLSearchQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *WebPageURLSearchQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *WebPageURLSearchQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *WebPageURLSearchQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetIncludeRevenue

`func (o *WebPageURLSearchQuery) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *WebPageURLSearchQuery) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *WebPageURLSearchQuery) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *WebPageURLSearchQuery) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *WebPageURLSearchQuery) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *WebPageURLSearchQuery) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetKind

`func (o *WebPageURLSearchQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebPageURLSearchQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebPageURLSearchQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WebPageURLSearchQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *WebPageURLSearchQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *WebPageURLSearchQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *WebPageURLSearchQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *WebPageURLSearchQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *WebPageURLSearchQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *WebPageURLSearchQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *WebPageURLSearchQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *WebPageURLSearchQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *WebPageURLSearchQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *WebPageURLSearchQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOrderBy

`func (o *WebPageURLSearchQuery) GetOrderBy() []string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *WebPageURLSearchQuery) GetOrderByOk() (*[]string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *WebPageURLSearchQuery) SetOrderBy(v []string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *WebPageURLSearchQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *WebPageURLSearchQuery) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *WebPageURLSearchQuery) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetProperties

`func (o *WebPageURLSearchQuery) GetProperties() []Properties2Inner1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WebPageURLSearchQuery) GetPropertiesOk() (*[]Properties2Inner1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WebPageURLSearchQuery) SetProperties(v []Properties2Inner1)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *WebPageURLSearchQuery) GetResponse() WebPageURLSearchQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebPageURLSearchQuery) GetResponseOk() (*WebPageURLSearchQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebPageURLSearchQuery) SetResponse(v WebPageURLSearchQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebPageURLSearchQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSampling

`func (o *WebPageURLSearchQuery) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *WebPageURLSearchQuery) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *WebPageURLSearchQuery) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *WebPageURLSearchQuery) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetSearchTerm

`func (o *WebPageURLSearchQuery) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *WebPageURLSearchQuery) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *WebPageURLSearchQuery) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.

### HasSearchTerm

`func (o *WebPageURLSearchQuery) HasSearchTerm() bool`

HasSearchTerm returns a boolean if a field has been set.

### SetSearchTermNil

`func (o *WebPageURLSearchQuery) SetSearchTermNil(b bool)`

 SetSearchTermNil sets the value for SearchTerm to be an explicit nil

### UnsetSearchTerm
`func (o *WebPageURLSearchQuery) UnsetSearchTerm()`

UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil
### GetStripQueryParams

`func (o *WebPageURLSearchQuery) GetStripQueryParams() bool`

GetStripQueryParams returns the StripQueryParams field if non-nil, zero value otherwise.

### GetStripQueryParamsOk

`func (o *WebPageURLSearchQuery) GetStripQueryParamsOk() (*bool, bool)`

GetStripQueryParamsOk returns a tuple with the StripQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripQueryParams

`func (o *WebPageURLSearchQuery) SetStripQueryParams(v bool)`

SetStripQueryParams sets StripQueryParams field to given value.

### HasStripQueryParams

`func (o *WebPageURLSearchQuery) HasStripQueryParams() bool`

HasStripQueryParams returns a boolean if a field has been set.

### SetStripQueryParamsNil

`func (o *WebPageURLSearchQuery) SetStripQueryParamsNil(b bool)`

 SetStripQueryParamsNil sets the value for StripQueryParams to be an explicit nil

### UnsetStripQueryParams
`func (o *WebPageURLSearchQuery) UnsetStripQueryParams()`

UnsetStripQueryParams ensures that no value is present for StripQueryParams, not even an explicit nil
### GetTags

`func (o *WebPageURLSearchQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WebPageURLSearchQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WebPageURLSearchQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WebPageURLSearchQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *WebPageURLSearchQuery) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *WebPageURLSearchQuery) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *WebPageURLSearchQuery) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *WebPageURLSearchQuery) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *WebPageURLSearchQuery) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *WebPageURLSearchQuery) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetVersion

`func (o *WebPageURLSearchQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebPageURLSearchQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebPageURLSearchQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WebPageURLSearchQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *WebPageURLSearchQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *WebPageURLSearchQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


