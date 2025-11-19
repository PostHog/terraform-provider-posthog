# MarketingAnalyticsTableQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal1**](Conversiongoal1.md) |  | [optional] [default to null]
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**DraftConversionGoal** | Pointer to [**NullableDraftconversiongoal**](Draftconversiongoal.md) |  | [optional] [default to null]
**FilterTestAccounts** | Pointer to **NullableBool** | Filter test accounts | [optional] 
**IncludeAllConversions** | Pointer to **NullableBool** | Include conversion goal rows even when they don&#39;t match campaign costs table | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**IntegrationFilter** | Pointer to [**IntegrationFilter**](IntegrationFilter.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "MarketingAnalyticsTableQuery"]
**Limit** | Pointer to **NullableInt32** | Number of rows to return | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** | Number of rows to skip before returning rows | [optional] 
**OrderBy** | Pointer to [**[][]OrderbyInnerInner**]([]OrderbyInnerInner.md) | Columns to order by - similar to EventsQuery format | [optional] 
**Properties** | [**[]Properties2Inner1**](Properties2Inner1.md) |  | 
**Response** | Pointer to [**MarketingAnalyticsTableQueryResponse**](MarketingAnalyticsTableQueryResponse.md) |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**Select** | Pointer to **[]string** | Return a limited set of data. Will use default columns if empty. | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewMarketingAnalyticsTableQuery

`func NewMarketingAnalyticsTableQuery(properties []Properties2Inner1, ) *MarketingAnalyticsTableQuery`

NewMarketingAnalyticsTableQuery instantiates a new MarketingAnalyticsTableQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingAnalyticsTableQueryWithDefaults

`func NewMarketingAnalyticsTableQueryWithDefaults() *MarketingAnalyticsTableQuery`

NewMarketingAnalyticsTableQueryWithDefaults instantiates a new MarketingAnalyticsTableQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareFilter

`func (o *MarketingAnalyticsTableQuery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *MarketingAnalyticsTableQuery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *MarketingAnalyticsTableQuery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *MarketingAnalyticsTableQuery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *MarketingAnalyticsTableQuery) GetConversionGoal() Conversiongoal1`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *MarketingAnalyticsTableQuery) GetConversionGoalOk() (*Conversiongoal1, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *MarketingAnalyticsTableQuery) SetConversionGoal(v Conversiongoal1)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *MarketingAnalyticsTableQuery) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *MarketingAnalyticsTableQuery) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *MarketingAnalyticsTableQuery) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDateRange

`func (o *MarketingAnalyticsTableQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *MarketingAnalyticsTableQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *MarketingAnalyticsTableQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *MarketingAnalyticsTableQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDoPathCleaning

`func (o *MarketingAnalyticsTableQuery) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *MarketingAnalyticsTableQuery) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *MarketingAnalyticsTableQuery) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *MarketingAnalyticsTableQuery) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *MarketingAnalyticsTableQuery) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *MarketingAnalyticsTableQuery) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetDraftConversionGoal

`func (o *MarketingAnalyticsTableQuery) GetDraftConversionGoal() Draftconversiongoal`

GetDraftConversionGoal returns the DraftConversionGoal field if non-nil, zero value otherwise.

### GetDraftConversionGoalOk

`func (o *MarketingAnalyticsTableQuery) GetDraftConversionGoalOk() (*Draftconversiongoal, bool)`

GetDraftConversionGoalOk returns a tuple with the DraftConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftConversionGoal

`func (o *MarketingAnalyticsTableQuery) SetDraftConversionGoal(v Draftconversiongoal)`

SetDraftConversionGoal sets DraftConversionGoal field to given value.

### HasDraftConversionGoal

`func (o *MarketingAnalyticsTableQuery) HasDraftConversionGoal() bool`

HasDraftConversionGoal returns a boolean if a field has been set.

### SetDraftConversionGoalNil

`func (o *MarketingAnalyticsTableQuery) SetDraftConversionGoalNil(b bool)`

 SetDraftConversionGoalNil sets the value for DraftConversionGoal to be an explicit nil

### UnsetDraftConversionGoal
`func (o *MarketingAnalyticsTableQuery) UnsetDraftConversionGoal()`

UnsetDraftConversionGoal ensures that no value is present for DraftConversionGoal, not even an explicit nil
### GetFilterTestAccounts

`func (o *MarketingAnalyticsTableQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *MarketingAnalyticsTableQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *MarketingAnalyticsTableQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *MarketingAnalyticsTableQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *MarketingAnalyticsTableQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *MarketingAnalyticsTableQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetIncludeAllConversions

`func (o *MarketingAnalyticsTableQuery) GetIncludeAllConversions() bool`

GetIncludeAllConversions returns the IncludeAllConversions field if non-nil, zero value otherwise.

### GetIncludeAllConversionsOk

`func (o *MarketingAnalyticsTableQuery) GetIncludeAllConversionsOk() (*bool, bool)`

GetIncludeAllConversionsOk returns a tuple with the IncludeAllConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllConversions

`func (o *MarketingAnalyticsTableQuery) SetIncludeAllConversions(v bool)`

SetIncludeAllConversions sets IncludeAllConversions field to given value.

### HasIncludeAllConversions

`func (o *MarketingAnalyticsTableQuery) HasIncludeAllConversions() bool`

HasIncludeAllConversions returns a boolean if a field has been set.

### SetIncludeAllConversionsNil

`func (o *MarketingAnalyticsTableQuery) SetIncludeAllConversionsNil(b bool)`

 SetIncludeAllConversionsNil sets the value for IncludeAllConversions to be an explicit nil

### UnsetIncludeAllConversions
`func (o *MarketingAnalyticsTableQuery) UnsetIncludeAllConversions()`

UnsetIncludeAllConversions ensures that no value is present for IncludeAllConversions, not even an explicit nil
### GetIncludeRevenue

`func (o *MarketingAnalyticsTableQuery) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *MarketingAnalyticsTableQuery) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *MarketingAnalyticsTableQuery) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *MarketingAnalyticsTableQuery) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *MarketingAnalyticsTableQuery) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *MarketingAnalyticsTableQuery) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetIntegrationFilter

`func (o *MarketingAnalyticsTableQuery) GetIntegrationFilter() IntegrationFilter`

GetIntegrationFilter returns the IntegrationFilter field if non-nil, zero value otherwise.

### GetIntegrationFilterOk

`func (o *MarketingAnalyticsTableQuery) GetIntegrationFilterOk() (*IntegrationFilter, bool)`

GetIntegrationFilterOk returns a tuple with the IntegrationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationFilter

`func (o *MarketingAnalyticsTableQuery) SetIntegrationFilter(v IntegrationFilter)`

SetIntegrationFilter sets IntegrationFilter field to given value.

### HasIntegrationFilter

`func (o *MarketingAnalyticsTableQuery) HasIntegrationFilter() bool`

HasIntegrationFilter returns a boolean if a field has been set.

### GetKind

`func (o *MarketingAnalyticsTableQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketingAnalyticsTableQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketingAnalyticsTableQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MarketingAnalyticsTableQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *MarketingAnalyticsTableQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MarketingAnalyticsTableQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MarketingAnalyticsTableQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MarketingAnalyticsTableQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *MarketingAnalyticsTableQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *MarketingAnalyticsTableQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *MarketingAnalyticsTableQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *MarketingAnalyticsTableQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *MarketingAnalyticsTableQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *MarketingAnalyticsTableQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *MarketingAnalyticsTableQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *MarketingAnalyticsTableQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *MarketingAnalyticsTableQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *MarketingAnalyticsTableQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *MarketingAnalyticsTableQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *MarketingAnalyticsTableQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOrderBy

`func (o *MarketingAnalyticsTableQuery) GetOrderBy() [][]OrderbyInnerInner`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *MarketingAnalyticsTableQuery) GetOrderByOk() (*[][]OrderbyInnerInner, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *MarketingAnalyticsTableQuery) SetOrderBy(v [][]OrderbyInnerInner)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *MarketingAnalyticsTableQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *MarketingAnalyticsTableQuery) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *MarketingAnalyticsTableQuery) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetProperties

`func (o *MarketingAnalyticsTableQuery) GetProperties() []Properties2Inner1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MarketingAnalyticsTableQuery) GetPropertiesOk() (*[]Properties2Inner1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MarketingAnalyticsTableQuery) SetProperties(v []Properties2Inner1)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *MarketingAnalyticsTableQuery) GetResponse() MarketingAnalyticsTableQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MarketingAnalyticsTableQuery) GetResponseOk() (*MarketingAnalyticsTableQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *MarketingAnalyticsTableQuery) SetResponse(v MarketingAnalyticsTableQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *MarketingAnalyticsTableQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSampling

`func (o *MarketingAnalyticsTableQuery) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *MarketingAnalyticsTableQuery) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *MarketingAnalyticsTableQuery) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *MarketingAnalyticsTableQuery) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetSelect

`func (o *MarketingAnalyticsTableQuery) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *MarketingAnalyticsTableQuery) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *MarketingAnalyticsTableQuery) SetSelect(v []string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *MarketingAnalyticsTableQuery) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### SetSelectNil

`func (o *MarketingAnalyticsTableQuery) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *MarketingAnalyticsTableQuery) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetTags

`func (o *MarketingAnalyticsTableQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MarketingAnalyticsTableQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MarketingAnalyticsTableQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MarketingAnalyticsTableQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *MarketingAnalyticsTableQuery) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *MarketingAnalyticsTableQuery) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *MarketingAnalyticsTableQuery) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *MarketingAnalyticsTableQuery) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *MarketingAnalyticsTableQuery) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *MarketingAnalyticsTableQuery) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetVersion

`func (o *MarketingAnalyticsTableQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketingAnalyticsTableQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketingAnalyticsTableQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketingAnalyticsTableQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MarketingAnalyticsTableQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MarketingAnalyticsTableQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


