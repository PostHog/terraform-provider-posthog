# MarketingAnalyticsAggregatedQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareFilter** | Pointer to [**CompareFilter**](CompareFilter.md) |  | [optional] 
**ConversionGoal** | Pointer to [**NullableConversiongoal1**](Conversiongoal1.md) |  | [optional] [default to null]
**DateRange** | Pointer to [**DateRange**](DateRange.md) |  | [optional] 
**DoPathCleaning** | Pointer to **NullableBool** |  | [optional] 
**DraftConversionGoal** | Pointer to [**NullableDraftconversiongoal**](Draftconversiongoal.md) |  | [optional] [default to null]
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**IncludeRevenue** | Pointer to **NullableBool** |  | [optional] 
**IntegrationFilter** | Pointer to [**IntegrationFilter**](IntegrationFilter.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "MarketingAnalyticsAggregatedQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Properties** | [**[]Properties2Inner1**](Properties2Inner1.md) |  | 
**Response** | Pointer to [**MarketingAnalyticsAggregatedQueryResponse**](MarketingAnalyticsAggregatedQueryResponse.md) |  | [optional] 
**Sampling** | Pointer to [**WebAnalyticsSampling**](WebAnalyticsSampling.md) |  | [optional] 
**Select** | Pointer to **[]string** | Return a limited set of data. Will use default columns if empty. | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**UseSessionsTable** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewMarketingAnalyticsAggregatedQuery

`func NewMarketingAnalyticsAggregatedQuery(properties []Properties2Inner1, ) *MarketingAnalyticsAggregatedQuery`

NewMarketingAnalyticsAggregatedQuery instantiates a new MarketingAnalyticsAggregatedQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingAnalyticsAggregatedQueryWithDefaults

`func NewMarketingAnalyticsAggregatedQueryWithDefaults() *MarketingAnalyticsAggregatedQuery`

NewMarketingAnalyticsAggregatedQueryWithDefaults instantiates a new MarketingAnalyticsAggregatedQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareFilter

`func (o *MarketingAnalyticsAggregatedQuery) GetCompareFilter() CompareFilter`

GetCompareFilter returns the CompareFilter field if non-nil, zero value otherwise.

### GetCompareFilterOk

`func (o *MarketingAnalyticsAggregatedQuery) GetCompareFilterOk() (*CompareFilter, bool)`

GetCompareFilterOk returns a tuple with the CompareFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareFilter

`func (o *MarketingAnalyticsAggregatedQuery) SetCompareFilter(v CompareFilter)`

SetCompareFilter sets CompareFilter field to given value.

### HasCompareFilter

`func (o *MarketingAnalyticsAggregatedQuery) HasCompareFilter() bool`

HasCompareFilter returns a boolean if a field has been set.

### GetConversionGoal

`func (o *MarketingAnalyticsAggregatedQuery) GetConversionGoal() Conversiongoal1`

GetConversionGoal returns the ConversionGoal field if non-nil, zero value otherwise.

### GetConversionGoalOk

`func (o *MarketingAnalyticsAggregatedQuery) GetConversionGoalOk() (*Conversiongoal1, bool)`

GetConversionGoalOk returns a tuple with the ConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionGoal

`func (o *MarketingAnalyticsAggregatedQuery) SetConversionGoal(v Conversiongoal1)`

SetConversionGoal sets ConversionGoal field to given value.

### HasConversionGoal

`func (o *MarketingAnalyticsAggregatedQuery) HasConversionGoal() bool`

HasConversionGoal returns a boolean if a field has been set.

### SetConversionGoalNil

`func (o *MarketingAnalyticsAggregatedQuery) SetConversionGoalNil(b bool)`

 SetConversionGoalNil sets the value for ConversionGoal to be an explicit nil

### UnsetConversionGoal
`func (o *MarketingAnalyticsAggregatedQuery) UnsetConversionGoal()`

UnsetConversionGoal ensures that no value is present for ConversionGoal, not even an explicit nil
### GetDateRange

`func (o *MarketingAnalyticsAggregatedQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *MarketingAnalyticsAggregatedQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *MarketingAnalyticsAggregatedQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *MarketingAnalyticsAggregatedQuery) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetDoPathCleaning

`func (o *MarketingAnalyticsAggregatedQuery) GetDoPathCleaning() bool`

GetDoPathCleaning returns the DoPathCleaning field if non-nil, zero value otherwise.

### GetDoPathCleaningOk

`func (o *MarketingAnalyticsAggregatedQuery) GetDoPathCleaningOk() (*bool, bool)`

GetDoPathCleaningOk returns a tuple with the DoPathCleaning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoPathCleaning

`func (o *MarketingAnalyticsAggregatedQuery) SetDoPathCleaning(v bool)`

SetDoPathCleaning sets DoPathCleaning field to given value.

### HasDoPathCleaning

`func (o *MarketingAnalyticsAggregatedQuery) HasDoPathCleaning() bool`

HasDoPathCleaning returns a boolean if a field has been set.

### SetDoPathCleaningNil

`func (o *MarketingAnalyticsAggregatedQuery) SetDoPathCleaningNil(b bool)`

 SetDoPathCleaningNil sets the value for DoPathCleaning to be an explicit nil

### UnsetDoPathCleaning
`func (o *MarketingAnalyticsAggregatedQuery) UnsetDoPathCleaning()`

UnsetDoPathCleaning ensures that no value is present for DoPathCleaning, not even an explicit nil
### GetDraftConversionGoal

`func (o *MarketingAnalyticsAggregatedQuery) GetDraftConversionGoal() Draftconversiongoal`

GetDraftConversionGoal returns the DraftConversionGoal field if non-nil, zero value otherwise.

### GetDraftConversionGoalOk

`func (o *MarketingAnalyticsAggregatedQuery) GetDraftConversionGoalOk() (*Draftconversiongoal, bool)`

GetDraftConversionGoalOk returns a tuple with the DraftConversionGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftConversionGoal

`func (o *MarketingAnalyticsAggregatedQuery) SetDraftConversionGoal(v Draftconversiongoal)`

SetDraftConversionGoal sets DraftConversionGoal field to given value.

### HasDraftConversionGoal

`func (o *MarketingAnalyticsAggregatedQuery) HasDraftConversionGoal() bool`

HasDraftConversionGoal returns a boolean if a field has been set.

### SetDraftConversionGoalNil

`func (o *MarketingAnalyticsAggregatedQuery) SetDraftConversionGoalNil(b bool)`

 SetDraftConversionGoalNil sets the value for DraftConversionGoal to be an explicit nil

### UnsetDraftConversionGoal
`func (o *MarketingAnalyticsAggregatedQuery) UnsetDraftConversionGoal()`

UnsetDraftConversionGoal ensures that no value is present for DraftConversionGoal, not even an explicit nil
### GetFilterTestAccounts

`func (o *MarketingAnalyticsAggregatedQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *MarketingAnalyticsAggregatedQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *MarketingAnalyticsAggregatedQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *MarketingAnalyticsAggregatedQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *MarketingAnalyticsAggregatedQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *MarketingAnalyticsAggregatedQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetIncludeRevenue

`func (o *MarketingAnalyticsAggregatedQuery) GetIncludeRevenue() bool`

GetIncludeRevenue returns the IncludeRevenue field if non-nil, zero value otherwise.

### GetIncludeRevenueOk

`func (o *MarketingAnalyticsAggregatedQuery) GetIncludeRevenueOk() (*bool, bool)`

GetIncludeRevenueOk returns a tuple with the IncludeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevenue

`func (o *MarketingAnalyticsAggregatedQuery) SetIncludeRevenue(v bool)`

SetIncludeRevenue sets IncludeRevenue field to given value.

### HasIncludeRevenue

`func (o *MarketingAnalyticsAggregatedQuery) HasIncludeRevenue() bool`

HasIncludeRevenue returns a boolean if a field has been set.

### SetIncludeRevenueNil

`func (o *MarketingAnalyticsAggregatedQuery) SetIncludeRevenueNil(b bool)`

 SetIncludeRevenueNil sets the value for IncludeRevenue to be an explicit nil

### UnsetIncludeRevenue
`func (o *MarketingAnalyticsAggregatedQuery) UnsetIncludeRevenue()`

UnsetIncludeRevenue ensures that no value is present for IncludeRevenue, not even an explicit nil
### GetIntegrationFilter

`func (o *MarketingAnalyticsAggregatedQuery) GetIntegrationFilter() IntegrationFilter`

GetIntegrationFilter returns the IntegrationFilter field if non-nil, zero value otherwise.

### GetIntegrationFilterOk

`func (o *MarketingAnalyticsAggregatedQuery) GetIntegrationFilterOk() (*IntegrationFilter, bool)`

GetIntegrationFilterOk returns a tuple with the IntegrationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationFilter

`func (o *MarketingAnalyticsAggregatedQuery) SetIntegrationFilter(v IntegrationFilter)`

SetIntegrationFilter sets IntegrationFilter field to given value.

### HasIntegrationFilter

`func (o *MarketingAnalyticsAggregatedQuery) HasIntegrationFilter() bool`

HasIntegrationFilter returns a boolean if a field has been set.

### GetKind

`func (o *MarketingAnalyticsAggregatedQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketingAnalyticsAggregatedQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketingAnalyticsAggregatedQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MarketingAnalyticsAggregatedQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *MarketingAnalyticsAggregatedQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *MarketingAnalyticsAggregatedQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *MarketingAnalyticsAggregatedQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *MarketingAnalyticsAggregatedQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetProperties

`func (o *MarketingAnalyticsAggregatedQuery) GetProperties() []Properties2Inner1`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MarketingAnalyticsAggregatedQuery) GetPropertiesOk() (*[]Properties2Inner1, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MarketingAnalyticsAggregatedQuery) SetProperties(v []Properties2Inner1)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *MarketingAnalyticsAggregatedQuery) GetResponse() MarketingAnalyticsAggregatedQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MarketingAnalyticsAggregatedQuery) GetResponseOk() (*MarketingAnalyticsAggregatedQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *MarketingAnalyticsAggregatedQuery) SetResponse(v MarketingAnalyticsAggregatedQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *MarketingAnalyticsAggregatedQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSampling

`func (o *MarketingAnalyticsAggregatedQuery) GetSampling() WebAnalyticsSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *MarketingAnalyticsAggregatedQuery) GetSamplingOk() (*WebAnalyticsSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *MarketingAnalyticsAggregatedQuery) SetSampling(v WebAnalyticsSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *MarketingAnalyticsAggregatedQuery) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### GetSelect

`func (o *MarketingAnalyticsAggregatedQuery) GetSelect() []string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *MarketingAnalyticsAggregatedQuery) GetSelectOk() (*[]string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *MarketingAnalyticsAggregatedQuery) SetSelect(v []string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *MarketingAnalyticsAggregatedQuery) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### SetSelectNil

`func (o *MarketingAnalyticsAggregatedQuery) SetSelectNil(b bool)`

 SetSelectNil sets the value for Select to be an explicit nil

### UnsetSelect
`func (o *MarketingAnalyticsAggregatedQuery) UnsetSelect()`

UnsetSelect ensures that no value is present for Select, not even an explicit nil
### GetTags

`func (o *MarketingAnalyticsAggregatedQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MarketingAnalyticsAggregatedQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MarketingAnalyticsAggregatedQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MarketingAnalyticsAggregatedQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUseSessionsTable

`func (o *MarketingAnalyticsAggregatedQuery) GetUseSessionsTable() bool`

GetUseSessionsTable returns the UseSessionsTable field if non-nil, zero value otherwise.

### GetUseSessionsTableOk

`func (o *MarketingAnalyticsAggregatedQuery) GetUseSessionsTableOk() (*bool, bool)`

GetUseSessionsTableOk returns a tuple with the UseSessionsTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSessionsTable

`func (o *MarketingAnalyticsAggregatedQuery) SetUseSessionsTable(v bool)`

SetUseSessionsTable sets UseSessionsTable field to given value.

### HasUseSessionsTable

`func (o *MarketingAnalyticsAggregatedQuery) HasUseSessionsTable() bool`

HasUseSessionsTable returns a boolean if a field has been set.

### SetUseSessionsTableNil

`func (o *MarketingAnalyticsAggregatedQuery) SetUseSessionsTableNil(b bool)`

 SetUseSessionsTableNil sets the value for UseSessionsTable to be an explicit nil

### UnsetUseSessionsTable
`func (o *MarketingAnalyticsAggregatedQuery) UnsetUseSessionsTable()`

UnsetUseSessionsTable ensures that no value is present for UseSessionsTable, not even an explicit nil
### GetVersion

`func (o *MarketingAnalyticsAggregatedQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MarketingAnalyticsAggregatedQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MarketingAnalyticsAggregatedQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MarketingAnalyticsAggregatedQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MarketingAnalyticsAggregatedQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MarketingAnalyticsAggregatedQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


