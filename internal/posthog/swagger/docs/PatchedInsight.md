# PatchedInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**ShortId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DerivedName** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to **map[string]interface{}** |  | [optional] 
**Order** | Pointer to **NullableInt32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Dashboards** | Pointer to **[]int32** |          DEPRECATED. Will be removed in a future release. Use dashboard_tiles instead.         A dashboard ID for each of the dashboards that this insight is displayed on.          | [optional] 
**DashboardTiles** | Pointer to [**[]DashboardTileBasic**](DashboardTileBasic.md) |      A dashboard tile ID and dashboard_id for each of the dashboards that this insight is displayed on.      | [optional] [readonly] 
**LastRefresh** | Pointer to **string** |      The datetime this insight&#39;s results were generated.     If added to one or more dashboards the insight can be refreshed separately on each.     Returns the appropriate last_refresh datetime for the context the insight is viewed in     (see from_dashboard query parameter).      | [optional] [readonly] 
**CacheTargetAge** | Pointer to **string** | The target age of the cached results for this insight. | [optional] [readonly] 
**NextAllowedClientRefresh** | Pointer to **string** |      The earliest possible datetime at which we&#39;ll allow the cached results for this insight to be refreshed     by querying the database.      | [optional] [readonly] 
**Result** | Pointer to **string** |  | [optional] [readonly] 
**HasMore** | Pointer to **string** |  | [optional] [readonly] 
**Columns** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**Favorited** | Pointer to **bool** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastModifiedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**IsSample** | Pointer to **bool** |  | [optional] [readonly] 
**EffectiveRestrictionLevel** | Pointer to [**EffectiveRestrictionLevelEnum**](EffectiveRestrictionLevelEnum.md) |  | [optional] [readonly] 
**EffectivePrivilegeLevel** | Pointer to [**EffectivePrivilegeLevelEnum**](EffectivePrivilegeLevelEnum.md) |  | [optional] [readonly] 
**UserAccessLevel** | Pointer to **NullableString** | The effective access level the user has for this object | [optional] [readonly] 
**Timezone** | Pointer to **string** | The timezone this chart is displayed in. | [optional] [readonly] 
**IsCached** | Pointer to **string** |  | [optional] [readonly] 
**QueryStatus** | Pointer to **string** |  | [optional] [readonly] 
**Hogql** | Pointer to **string** |  | [optional] [readonly] 
**Types** | Pointer to **string** |  | [optional] [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**Alerts** | Pointer to **string** |  | [optional] [readonly] 
**LastViewedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedInsight

`func NewPatchedInsight() *PatchedInsight`

NewPatchedInsight instantiates a new PatchedInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedInsightWithDefaults

`func NewPatchedInsightWithDefaults() *PatchedInsight`

NewPatchedInsightWithDefaults instantiates a new PatchedInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedInsight) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedInsight) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedInsight) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedInsight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShortId

`func (o *PatchedInsight) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *PatchedInsight) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *PatchedInsight) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *PatchedInsight) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetName

`func (o *PatchedInsight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedInsight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedInsight) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedInsight) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedInsight) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedInsight) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDerivedName

`func (o *PatchedInsight) GetDerivedName() string`

GetDerivedName returns the DerivedName field if non-nil, zero value otherwise.

### GetDerivedNameOk

`func (o *PatchedInsight) GetDerivedNameOk() (*string, bool)`

GetDerivedNameOk returns a tuple with the DerivedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedName

`func (o *PatchedInsight) SetDerivedName(v string)`

SetDerivedName sets DerivedName field to given value.

### HasDerivedName

`func (o *PatchedInsight) HasDerivedName() bool`

HasDerivedName returns a boolean if a field has been set.

### SetDerivedNameNil

`func (o *PatchedInsight) SetDerivedNameNil(b bool)`

 SetDerivedNameNil sets the value for DerivedName to be an explicit nil

### UnsetDerivedName
`func (o *PatchedInsight) UnsetDerivedName()`

UnsetDerivedName ensures that no value is present for DerivedName, not even an explicit nil
### GetQuery

`func (o *PatchedInsight) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PatchedInsight) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PatchedInsight) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PatchedInsight) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *PatchedInsight) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *PatchedInsight) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetOrder

`func (o *PatchedInsight) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PatchedInsight) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PatchedInsight) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PatchedInsight) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *PatchedInsight) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *PatchedInsight) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetDeleted

`func (o *PatchedInsight) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedInsight) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedInsight) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedInsight) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDashboards

`func (o *PatchedInsight) GetDashboards() []int32`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *PatchedInsight) GetDashboardsOk() (*[]int32, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *PatchedInsight) SetDashboards(v []int32)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *PatchedInsight) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### GetDashboardTiles

`func (o *PatchedInsight) GetDashboardTiles() []DashboardTileBasic`

GetDashboardTiles returns the DashboardTiles field if non-nil, zero value otherwise.

### GetDashboardTilesOk

`func (o *PatchedInsight) GetDashboardTilesOk() (*[]DashboardTileBasic, bool)`

GetDashboardTilesOk returns a tuple with the DashboardTiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardTiles

`func (o *PatchedInsight) SetDashboardTiles(v []DashboardTileBasic)`

SetDashboardTiles sets DashboardTiles field to given value.

### HasDashboardTiles

`func (o *PatchedInsight) HasDashboardTiles() bool`

HasDashboardTiles returns a boolean if a field has been set.

### GetLastRefresh

`func (o *PatchedInsight) GetLastRefresh() string`

GetLastRefresh returns the LastRefresh field if non-nil, zero value otherwise.

### GetLastRefreshOk

`func (o *PatchedInsight) GetLastRefreshOk() (*string, bool)`

GetLastRefreshOk returns a tuple with the LastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefresh

`func (o *PatchedInsight) SetLastRefresh(v string)`

SetLastRefresh sets LastRefresh field to given value.

### HasLastRefresh

`func (o *PatchedInsight) HasLastRefresh() bool`

HasLastRefresh returns a boolean if a field has been set.

### GetCacheTargetAge

`func (o *PatchedInsight) GetCacheTargetAge() string`

GetCacheTargetAge returns the CacheTargetAge field if non-nil, zero value otherwise.

### GetCacheTargetAgeOk

`func (o *PatchedInsight) GetCacheTargetAgeOk() (*string, bool)`

GetCacheTargetAgeOk returns a tuple with the CacheTargetAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTargetAge

`func (o *PatchedInsight) SetCacheTargetAge(v string)`

SetCacheTargetAge sets CacheTargetAge field to given value.

### HasCacheTargetAge

`func (o *PatchedInsight) HasCacheTargetAge() bool`

HasCacheTargetAge returns a boolean if a field has been set.

### GetNextAllowedClientRefresh

`func (o *PatchedInsight) GetNextAllowedClientRefresh() string`

GetNextAllowedClientRefresh returns the NextAllowedClientRefresh field if non-nil, zero value otherwise.

### GetNextAllowedClientRefreshOk

`func (o *PatchedInsight) GetNextAllowedClientRefreshOk() (*string, bool)`

GetNextAllowedClientRefreshOk returns a tuple with the NextAllowedClientRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAllowedClientRefresh

`func (o *PatchedInsight) SetNextAllowedClientRefresh(v string)`

SetNextAllowedClientRefresh sets NextAllowedClientRefresh field to given value.

### HasNextAllowedClientRefresh

`func (o *PatchedInsight) HasNextAllowedClientRefresh() bool`

HasNextAllowedClientRefresh returns a boolean if a field has been set.

### GetResult

`func (o *PatchedInsight) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PatchedInsight) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PatchedInsight) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *PatchedInsight) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetHasMore

`func (o *PatchedInsight) GetHasMore() string`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PatchedInsight) GetHasMoreOk() (*string, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PatchedInsight) SetHasMore(v string)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *PatchedInsight) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetColumns

`func (o *PatchedInsight) GetColumns() string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *PatchedInsight) GetColumnsOk() (*string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *PatchedInsight) SetColumns(v string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *PatchedInsight) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedInsight) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedInsight) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedInsight) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedInsight) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedInsight) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedInsight) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *PatchedInsight) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedInsight) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedInsight) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedInsight) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedInsight) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedInsight) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedInsight) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedInsight) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedInsight) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedInsight) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUpdatedAt

`func (o *PatchedInsight) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedInsight) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedInsight) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedInsight) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *PatchedInsight) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedInsight) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedInsight) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedInsight) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFavorited

`func (o *PatchedInsight) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *PatchedInsight) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *PatchedInsight) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.

### HasFavorited

`func (o *PatchedInsight) HasFavorited() bool`

HasFavorited returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *PatchedInsight) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *PatchedInsight) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *PatchedInsight) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *PatchedInsight) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *PatchedInsight) GetLastModifiedBy() UserBasic`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *PatchedInsight) GetLastModifiedByOk() (*UserBasic, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *PatchedInsight) SetLastModifiedBy(v UserBasic)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *PatchedInsight) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetIsSample

`func (o *PatchedInsight) GetIsSample() bool`

GetIsSample returns the IsSample field if non-nil, zero value otherwise.

### GetIsSampleOk

`func (o *PatchedInsight) GetIsSampleOk() (*bool, bool)`

GetIsSampleOk returns a tuple with the IsSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSample

`func (o *PatchedInsight) SetIsSample(v bool)`

SetIsSample sets IsSample field to given value.

### HasIsSample

`func (o *PatchedInsight) HasIsSample() bool`

HasIsSample returns a boolean if a field has been set.

### GetEffectiveRestrictionLevel

`func (o *PatchedInsight) GetEffectiveRestrictionLevel() EffectiveRestrictionLevelEnum`

GetEffectiveRestrictionLevel returns the EffectiveRestrictionLevel field if non-nil, zero value otherwise.

### GetEffectiveRestrictionLevelOk

`func (o *PatchedInsight) GetEffectiveRestrictionLevelOk() (*EffectiveRestrictionLevelEnum, bool)`

GetEffectiveRestrictionLevelOk returns a tuple with the EffectiveRestrictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRestrictionLevel

`func (o *PatchedInsight) SetEffectiveRestrictionLevel(v EffectiveRestrictionLevelEnum)`

SetEffectiveRestrictionLevel sets EffectiveRestrictionLevel field to given value.

### HasEffectiveRestrictionLevel

`func (o *PatchedInsight) HasEffectiveRestrictionLevel() bool`

HasEffectiveRestrictionLevel returns a boolean if a field has been set.

### GetEffectivePrivilegeLevel

`func (o *PatchedInsight) GetEffectivePrivilegeLevel() EffectivePrivilegeLevelEnum`

GetEffectivePrivilegeLevel returns the EffectivePrivilegeLevel field if non-nil, zero value otherwise.

### GetEffectivePrivilegeLevelOk

`func (o *PatchedInsight) GetEffectivePrivilegeLevelOk() (*EffectivePrivilegeLevelEnum, bool)`

GetEffectivePrivilegeLevelOk returns a tuple with the EffectivePrivilegeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePrivilegeLevel

`func (o *PatchedInsight) SetEffectivePrivilegeLevel(v EffectivePrivilegeLevelEnum)`

SetEffectivePrivilegeLevel sets EffectivePrivilegeLevel field to given value.

### HasEffectivePrivilegeLevel

`func (o *PatchedInsight) HasEffectivePrivilegeLevel() bool`

HasEffectivePrivilegeLevel returns a boolean if a field has been set.

### GetUserAccessLevel

`func (o *PatchedInsight) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *PatchedInsight) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *PatchedInsight) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.

### HasUserAccessLevel

`func (o *PatchedInsight) HasUserAccessLevel() bool`

HasUserAccessLevel returns a boolean if a field has been set.

### SetUserAccessLevelNil

`func (o *PatchedInsight) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *PatchedInsight) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetTimezone

`func (o *PatchedInsight) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *PatchedInsight) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *PatchedInsight) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *PatchedInsight) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetIsCached

`func (o *PatchedInsight) GetIsCached() string`

GetIsCached returns the IsCached field if non-nil, zero value otherwise.

### GetIsCachedOk

`func (o *PatchedInsight) GetIsCachedOk() (*string, bool)`

GetIsCachedOk returns a tuple with the IsCached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCached

`func (o *PatchedInsight) SetIsCached(v string)`

SetIsCached sets IsCached field to given value.

### HasIsCached

`func (o *PatchedInsight) HasIsCached() bool`

HasIsCached returns a boolean if a field has been set.

### GetQueryStatus

`func (o *PatchedInsight) GetQueryStatus() string`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *PatchedInsight) GetQueryStatusOk() (*string, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *PatchedInsight) SetQueryStatus(v string)`

SetQueryStatus sets QueryStatus field to given value.

### HasQueryStatus

`func (o *PatchedInsight) HasQueryStatus() bool`

HasQueryStatus returns a boolean if a field has been set.

### GetHogql

`func (o *PatchedInsight) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *PatchedInsight) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *PatchedInsight) SetHogql(v string)`

SetHogql sets Hogql field to given value.

### HasHogql

`func (o *PatchedInsight) HasHogql() bool`

HasHogql returns a boolean if a field has been set.

### GetTypes

`func (o *PatchedInsight) GetTypes() string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PatchedInsight) GetTypesOk() (*string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PatchedInsight) SetTypes(v string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *PatchedInsight) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetCreateInFolder

`func (o *PatchedInsight) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedInsight) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedInsight) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedInsight) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetAlerts

`func (o *PatchedInsight) GetAlerts() string`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *PatchedInsight) GetAlertsOk() (*string, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *PatchedInsight) SetAlerts(v string)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *PatchedInsight) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetLastViewedAt

`func (o *PatchedInsight) GetLastViewedAt() string`

GetLastViewedAt returns the LastViewedAt field if non-nil, zero value otherwise.

### GetLastViewedAtOk

`func (o *PatchedInsight) GetLastViewedAtOk() (*string, bool)`

GetLastViewedAtOk returns a tuple with the LastViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedAt

`func (o *PatchedInsight) SetLastViewedAt(v string)`

SetLastViewedAt sets LastViewedAt field to given value.

### HasLastViewedAt

`func (o *PatchedInsight) HasLastViewedAt() bool`

HasLastViewedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


