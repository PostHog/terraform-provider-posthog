# ErrorTrackingQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**ErrorTrackingIssueAssignee**](ErrorTrackingIssueAssignee.md) |  | [optional] 
**DateRange** | [**DateRange**](DateRange.md) |  | 
**FilterGroup** | Pointer to [**PropertyGroupFilter**](PropertyGroupFilter.md) |  | [optional] 
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**GroupKey** | Pointer to **NullableString** |  | [optional] 
**GroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 
**IssueId** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ErrorTrackingQuery"]
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Offset** | Pointer to **NullableInt32** |  | [optional] 
**OrderBy** | [**OrderBy1**](OrderBy1.md) |  | 
**OrderDirection** | Pointer to [**OrderDirection1**](OrderDirection1.md) |  | [optional] 
**PersonId** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to [**ErrorTrackingQueryResponse**](ErrorTrackingQueryResponse.md) |  | [optional] 
**RevenueEntity** | Pointer to [**RevenueEntity**](RevenueEntity.md) |  | [optional] 
**RevenuePeriod** | Pointer to [**RevenuePeriod**](RevenuePeriod.md) |  | [optional] 
**SearchQuery** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**Status2**](Status2.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**VolumeResolution** | **int32** |  | 
**WithAggregations** | Pointer to **NullableBool** |  | [optional] 
**WithFirstEvent** | Pointer to **NullableBool** |  | [optional] 
**WithLastEvent** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewErrorTrackingQuery

`func NewErrorTrackingQuery(dateRange DateRange, orderBy OrderBy1, volumeResolution int32, ) *ErrorTrackingQuery`

NewErrorTrackingQuery instantiates a new ErrorTrackingQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorTrackingQueryWithDefaults

`func NewErrorTrackingQueryWithDefaults() *ErrorTrackingQuery`

NewErrorTrackingQueryWithDefaults instantiates a new ErrorTrackingQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *ErrorTrackingQuery) GetAssignee() ErrorTrackingIssueAssignee`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ErrorTrackingQuery) GetAssigneeOk() (*ErrorTrackingIssueAssignee, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ErrorTrackingQuery) SetAssignee(v ErrorTrackingIssueAssignee)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ErrorTrackingQuery) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetDateRange

`func (o *ErrorTrackingQuery) GetDateRange() DateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *ErrorTrackingQuery) GetDateRangeOk() (*DateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *ErrorTrackingQuery) SetDateRange(v DateRange)`

SetDateRange sets DateRange field to given value.


### GetFilterGroup

`func (o *ErrorTrackingQuery) GetFilterGroup() PropertyGroupFilter`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *ErrorTrackingQuery) GetFilterGroupOk() (*PropertyGroupFilter, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *ErrorTrackingQuery) SetFilterGroup(v PropertyGroupFilter)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *ErrorTrackingQuery) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### GetFilterTestAccounts

`func (o *ErrorTrackingQuery) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *ErrorTrackingQuery) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *ErrorTrackingQuery) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *ErrorTrackingQuery) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *ErrorTrackingQuery) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *ErrorTrackingQuery) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetGroupKey

`func (o *ErrorTrackingQuery) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *ErrorTrackingQuery) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *ErrorTrackingQuery) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *ErrorTrackingQuery) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### SetGroupKeyNil

`func (o *ErrorTrackingQuery) SetGroupKeyNil(b bool)`

 SetGroupKeyNil sets the value for GroupKey to be an explicit nil

### UnsetGroupKey
`func (o *ErrorTrackingQuery) UnsetGroupKey()`

UnsetGroupKey ensures that no value is present for GroupKey, not even an explicit nil
### GetGroupTypeIndex

`func (o *ErrorTrackingQuery) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *ErrorTrackingQuery) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *ErrorTrackingQuery) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *ErrorTrackingQuery) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *ErrorTrackingQuery) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *ErrorTrackingQuery) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetIssueId

`func (o *ErrorTrackingQuery) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *ErrorTrackingQuery) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *ErrorTrackingQuery) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *ErrorTrackingQuery) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### SetIssueIdNil

`func (o *ErrorTrackingQuery) SetIssueIdNil(b bool)`

 SetIssueIdNil sets the value for IssueId to be an explicit nil

### UnsetIssueId
`func (o *ErrorTrackingQuery) UnsetIssueId()`

UnsetIssueId ensures that no value is present for IssueId, not even an explicit nil
### GetKind

`func (o *ErrorTrackingQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ErrorTrackingQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ErrorTrackingQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ErrorTrackingQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *ErrorTrackingQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ErrorTrackingQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ErrorTrackingQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ErrorTrackingQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ErrorTrackingQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ErrorTrackingQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModifiers

`func (o *ErrorTrackingQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ErrorTrackingQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ErrorTrackingQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ErrorTrackingQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetOffset

`func (o *ErrorTrackingQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ErrorTrackingQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ErrorTrackingQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ErrorTrackingQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *ErrorTrackingQuery) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *ErrorTrackingQuery) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOrderBy

`func (o *ErrorTrackingQuery) GetOrderBy() OrderBy1`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ErrorTrackingQuery) GetOrderByOk() (*OrderBy1, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ErrorTrackingQuery) SetOrderBy(v OrderBy1)`

SetOrderBy sets OrderBy field to given value.


### GetOrderDirection

`func (o *ErrorTrackingQuery) GetOrderDirection() OrderDirection1`

GetOrderDirection returns the OrderDirection field if non-nil, zero value otherwise.

### GetOrderDirectionOk

`func (o *ErrorTrackingQuery) GetOrderDirectionOk() (*OrderDirection1, bool)`

GetOrderDirectionOk returns a tuple with the OrderDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDirection

`func (o *ErrorTrackingQuery) SetOrderDirection(v OrderDirection1)`

SetOrderDirection sets OrderDirection field to given value.

### HasOrderDirection

`func (o *ErrorTrackingQuery) HasOrderDirection() bool`

HasOrderDirection returns a boolean if a field has been set.

### GetPersonId

`func (o *ErrorTrackingQuery) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *ErrorTrackingQuery) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *ErrorTrackingQuery) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *ErrorTrackingQuery) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### SetPersonIdNil

`func (o *ErrorTrackingQuery) SetPersonIdNil(b bool)`

 SetPersonIdNil sets the value for PersonId to be an explicit nil

### UnsetPersonId
`func (o *ErrorTrackingQuery) UnsetPersonId()`

UnsetPersonId ensures that no value is present for PersonId, not even an explicit nil
### GetResponse

`func (o *ErrorTrackingQuery) GetResponse() ErrorTrackingQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ErrorTrackingQuery) GetResponseOk() (*ErrorTrackingQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ErrorTrackingQuery) SetResponse(v ErrorTrackingQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ErrorTrackingQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetRevenueEntity

`func (o *ErrorTrackingQuery) GetRevenueEntity() RevenueEntity`

GetRevenueEntity returns the RevenueEntity field if non-nil, zero value otherwise.

### GetRevenueEntityOk

`func (o *ErrorTrackingQuery) GetRevenueEntityOk() (*RevenueEntity, bool)`

GetRevenueEntityOk returns a tuple with the RevenueEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueEntity

`func (o *ErrorTrackingQuery) SetRevenueEntity(v RevenueEntity)`

SetRevenueEntity sets RevenueEntity field to given value.

### HasRevenueEntity

`func (o *ErrorTrackingQuery) HasRevenueEntity() bool`

HasRevenueEntity returns a boolean if a field has been set.

### GetRevenuePeriod

`func (o *ErrorTrackingQuery) GetRevenuePeriod() RevenuePeriod`

GetRevenuePeriod returns the RevenuePeriod field if non-nil, zero value otherwise.

### GetRevenuePeriodOk

`func (o *ErrorTrackingQuery) GetRevenuePeriodOk() (*RevenuePeriod, bool)`

GetRevenuePeriodOk returns a tuple with the RevenuePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenuePeriod

`func (o *ErrorTrackingQuery) SetRevenuePeriod(v RevenuePeriod)`

SetRevenuePeriod sets RevenuePeriod field to given value.

### HasRevenuePeriod

`func (o *ErrorTrackingQuery) HasRevenuePeriod() bool`

HasRevenuePeriod returns a boolean if a field has been set.

### GetSearchQuery

`func (o *ErrorTrackingQuery) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *ErrorTrackingQuery) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *ErrorTrackingQuery) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *ErrorTrackingQuery) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### SetSearchQueryNil

`func (o *ErrorTrackingQuery) SetSearchQueryNil(b bool)`

 SetSearchQueryNil sets the value for SearchQuery to be an explicit nil

### UnsetSearchQuery
`func (o *ErrorTrackingQuery) UnsetSearchQuery()`

UnsetSearchQuery ensures that no value is present for SearchQuery, not even an explicit nil
### GetStatus

`func (o *ErrorTrackingQuery) GetStatus() Status2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorTrackingQuery) GetStatusOk() (*Status2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorTrackingQuery) SetStatus(v Status2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorTrackingQuery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *ErrorTrackingQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ErrorTrackingQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ErrorTrackingQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ErrorTrackingQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *ErrorTrackingQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ErrorTrackingQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ErrorTrackingQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ErrorTrackingQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ErrorTrackingQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ErrorTrackingQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetVolumeResolution

`func (o *ErrorTrackingQuery) GetVolumeResolution() int32`

GetVolumeResolution returns the VolumeResolution field if non-nil, zero value otherwise.

### GetVolumeResolutionOk

`func (o *ErrorTrackingQuery) GetVolumeResolutionOk() (*int32, bool)`

GetVolumeResolutionOk returns a tuple with the VolumeResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeResolution

`func (o *ErrorTrackingQuery) SetVolumeResolution(v int32)`

SetVolumeResolution sets VolumeResolution field to given value.


### GetWithAggregations

`func (o *ErrorTrackingQuery) GetWithAggregations() bool`

GetWithAggregations returns the WithAggregations field if non-nil, zero value otherwise.

### GetWithAggregationsOk

`func (o *ErrorTrackingQuery) GetWithAggregationsOk() (*bool, bool)`

GetWithAggregationsOk returns a tuple with the WithAggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithAggregations

`func (o *ErrorTrackingQuery) SetWithAggregations(v bool)`

SetWithAggregations sets WithAggregations field to given value.

### HasWithAggregations

`func (o *ErrorTrackingQuery) HasWithAggregations() bool`

HasWithAggregations returns a boolean if a field has been set.

### SetWithAggregationsNil

`func (o *ErrorTrackingQuery) SetWithAggregationsNil(b bool)`

 SetWithAggregationsNil sets the value for WithAggregations to be an explicit nil

### UnsetWithAggregations
`func (o *ErrorTrackingQuery) UnsetWithAggregations()`

UnsetWithAggregations ensures that no value is present for WithAggregations, not even an explicit nil
### GetWithFirstEvent

`func (o *ErrorTrackingQuery) GetWithFirstEvent() bool`

GetWithFirstEvent returns the WithFirstEvent field if non-nil, zero value otherwise.

### GetWithFirstEventOk

`func (o *ErrorTrackingQuery) GetWithFirstEventOk() (*bool, bool)`

GetWithFirstEventOk returns a tuple with the WithFirstEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithFirstEvent

`func (o *ErrorTrackingQuery) SetWithFirstEvent(v bool)`

SetWithFirstEvent sets WithFirstEvent field to given value.

### HasWithFirstEvent

`func (o *ErrorTrackingQuery) HasWithFirstEvent() bool`

HasWithFirstEvent returns a boolean if a field has been set.

### SetWithFirstEventNil

`func (o *ErrorTrackingQuery) SetWithFirstEventNil(b bool)`

 SetWithFirstEventNil sets the value for WithFirstEvent to be an explicit nil

### UnsetWithFirstEvent
`func (o *ErrorTrackingQuery) UnsetWithFirstEvent()`

UnsetWithFirstEvent ensures that no value is present for WithFirstEvent, not even an explicit nil
### GetWithLastEvent

`func (o *ErrorTrackingQuery) GetWithLastEvent() bool`

GetWithLastEvent returns the WithLastEvent field if non-nil, zero value otherwise.

### GetWithLastEventOk

`func (o *ErrorTrackingQuery) GetWithLastEventOk() (*bool, bool)`

GetWithLastEventOk returns a tuple with the WithLastEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithLastEvent

`func (o *ErrorTrackingQuery) SetWithLastEvent(v bool)`

SetWithLastEvent sets WithLastEvent field to given value.

### HasWithLastEvent

`func (o *ErrorTrackingQuery) HasWithLastEvent() bool`

HasWithLastEvent returns a boolean if a field has been set.

### SetWithLastEventNil

`func (o *ErrorTrackingQuery) SetWithLastEventNil(b bool)`

 SetWithLastEventNil sets the value for WithLastEvent to be an explicit nil

### UnsetWithLastEvent
`func (o *ErrorTrackingQuery) UnsetWithLastEvent()`

UnsetWithLastEvent ensures that no value is present for WithLastEvent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


