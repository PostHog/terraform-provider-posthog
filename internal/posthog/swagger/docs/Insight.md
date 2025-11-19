# Insight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ShortId** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DerivedName** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to **map[string]interface{}** |  | [optional] 
**Order** | Pointer to **NullableInt32** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Dashboards** | Pointer to **[]int32** |          DEPRECATED. Will be removed in a future release. Use dashboard_tiles instead.         A dashboard ID for each of the dashboards that this insight is displayed on.          | [optional] 
**DashboardTiles** | [**[]DashboardTileBasic**](DashboardTileBasic.md) |      A dashboard tile ID and dashboard_id for each of the dashboards that this insight is displayed on.      | [readonly] 
**LastRefresh** | **string** |      The datetime this insight&#39;s results were generated.     If added to one or more dashboards the insight can be refreshed separately on each.     Returns the appropriate last_refresh datetime for the context the insight is viewed in     (see from_dashboard query parameter).      | [readonly] 
**CacheTargetAge** | **string** | The target age of the cached results for this insight. | [readonly] 
**NextAllowedClientRefresh** | **string** |      The earliest possible datetime at which we&#39;ll allow the cached results for this insight to be refreshed     by querying the database.      | [readonly] 
**Result** | **string** |  | [readonly] 
**HasMore** | **string** |  | [readonly] 
**Columns** | **string** |  | [readonly] 
**CreatedAt** | **NullableTime** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**Description** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**Favorited** | Pointer to **bool** |  | [optional] 
**LastModifiedAt** | **time.Time** |  | [readonly] 
**LastModifiedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**IsSample** | **bool** |  | [readonly] 
**EffectiveRestrictionLevel** | [**EffectiveRestrictionLevelEnum**](EffectiveRestrictionLevelEnum.md) |  | [readonly] 
**EffectivePrivilegeLevel** | [**EffectivePrivilegeLevelEnum**](EffectivePrivilegeLevelEnum.md) |  | [readonly] 
**UserAccessLevel** | **NullableString** | The effective access level the user has for this object | [readonly] 
**Timezone** | **string** | The timezone this chart is displayed in. | [readonly] 
**IsCached** | **string** |  | [readonly] 
**QueryStatus** | **string** |  | [readonly] 
**Hogql** | **string** |  | [readonly] 
**Types** | **string** |  | [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**Alerts** | **string** |  | [readonly] 
**LastViewedAt** | **string** |  | [readonly] 

## Methods

### NewInsight

`func NewInsight(id int32, shortId string, dashboardTiles []DashboardTileBasic, lastRefresh string, cacheTargetAge string, nextAllowedClientRefresh string, result string, hasMore string, columns string, createdAt NullableTime, createdBy UserBasic, updatedAt time.Time, lastModifiedAt time.Time, lastModifiedBy UserBasic, isSample bool, effectiveRestrictionLevel EffectiveRestrictionLevelEnum, effectivePrivilegeLevel EffectivePrivilegeLevelEnum, userAccessLevel NullableString, timezone string, isCached string, queryStatus string, hogql string, types string, alerts string, lastViewedAt string, ) *Insight`

NewInsight instantiates a new Insight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightWithDefaults

`func NewInsightWithDefaults() *Insight`

NewInsightWithDefaults instantiates a new Insight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Insight) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Insight) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Insight) SetId(v int32)`

SetId sets Id field to given value.


### GetShortId

`func (o *Insight) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *Insight) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *Insight) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetName

`func (o *Insight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Insight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Insight) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Insight) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Insight) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Insight) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDerivedName

`func (o *Insight) GetDerivedName() string`

GetDerivedName returns the DerivedName field if non-nil, zero value otherwise.

### GetDerivedNameOk

`func (o *Insight) GetDerivedNameOk() (*string, bool)`

GetDerivedNameOk returns a tuple with the DerivedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedName

`func (o *Insight) SetDerivedName(v string)`

SetDerivedName sets DerivedName field to given value.

### HasDerivedName

`func (o *Insight) HasDerivedName() bool`

HasDerivedName returns a boolean if a field has been set.

### SetDerivedNameNil

`func (o *Insight) SetDerivedNameNil(b bool)`

 SetDerivedNameNil sets the value for DerivedName to be an explicit nil

### UnsetDerivedName
`func (o *Insight) UnsetDerivedName()`

UnsetDerivedName ensures that no value is present for DerivedName, not even an explicit nil
### GetQuery

`func (o *Insight) GetQuery() map[string]interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Insight) GetQueryOk() (*map[string]interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Insight) SetQuery(v map[string]interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Insight) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *Insight) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *Insight) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetOrder

`func (o *Insight) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Insight) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Insight) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Insight) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *Insight) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *Insight) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetDeleted

`func (o *Insight) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Insight) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Insight) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Insight) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDashboards

`func (o *Insight) GetDashboards() []int32`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *Insight) GetDashboardsOk() (*[]int32, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *Insight) SetDashboards(v []int32)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *Insight) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### GetDashboardTiles

`func (o *Insight) GetDashboardTiles() []DashboardTileBasic`

GetDashboardTiles returns the DashboardTiles field if non-nil, zero value otherwise.

### GetDashboardTilesOk

`func (o *Insight) GetDashboardTilesOk() (*[]DashboardTileBasic, bool)`

GetDashboardTilesOk returns a tuple with the DashboardTiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardTiles

`func (o *Insight) SetDashboardTiles(v []DashboardTileBasic)`

SetDashboardTiles sets DashboardTiles field to given value.


### GetLastRefresh

`func (o *Insight) GetLastRefresh() string`

GetLastRefresh returns the LastRefresh field if non-nil, zero value otherwise.

### GetLastRefreshOk

`func (o *Insight) GetLastRefreshOk() (*string, bool)`

GetLastRefreshOk returns a tuple with the LastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefresh

`func (o *Insight) SetLastRefresh(v string)`

SetLastRefresh sets LastRefresh field to given value.


### GetCacheTargetAge

`func (o *Insight) GetCacheTargetAge() string`

GetCacheTargetAge returns the CacheTargetAge field if non-nil, zero value otherwise.

### GetCacheTargetAgeOk

`func (o *Insight) GetCacheTargetAgeOk() (*string, bool)`

GetCacheTargetAgeOk returns a tuple with the CacheTargetAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTargetAge

`func (o *Insight) SetCacheTargetAge(v string)`

SetCacheTargetAge sets CacheTargetAge field to given value.


### GetNextAllowedClientRefresh

`func (o *Insight) GetNextAllowedClientRefresh() string`

GetNextAllowedClientRefresh returns the NextAllowedClientRefresh field if non-nil, zero value otherwise.

### GetNextAllowedClientRefreshOk

`func (o *Insight) GetNextAllowedClientRefreshOk() (*string, bool)`

GetNextAllowedClientRefreshOk returns a tuple with the NextAllowedClientRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAllowedClientRefresh

`func (o *Insight) SetNextAllowedClientRefresh(v string)`

SetNextAllowedClientRefresh sets NextAllowedClientRefresh field to given value.


### GetResult

`func (o *Insight) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Insight) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Insight) SetResult(v string)`

SetResult sets Result field to given value.


### GetHasMore

`func (o *Insight) GetHasMore() string`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *Insight) GetHasMoreOk() (*string, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *Insight) SetHasMore(v string)`

SetHasMore sets HasMore field to given value.


### GetColumns

`func (o *Insight) GetColumns() string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Insight) GetColumnsOk() (*string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Insight) SetColumns(v string)`

SetColumns sets Columns field to given value.


### GetCreatedAt

`func (o *Insight) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Insight) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Insight) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Insight) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Insight) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *Insight) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Insight) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Insight) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *Insight) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Insight) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Insight) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Insight) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Insight) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Insight) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUpdatedAt

`func (o *Insight) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Insight) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Insight) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTags

`func (o *Insight) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Insight) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Insight) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Insight) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFavorited

`func (o *Insight) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *Insight) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *Insight) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.

### HasFavorited

`func (o *Insight) HasFavorited() bool`

HasFavorited returns a boolean if a field has been set.

### GetLastModifiedAt

`func (o *Insight) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *Insight) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *Insight) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.


### GetLastModifiedBy

`func (o *Insight) GetLastModifiedBy() UserBasic`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *Insight) GetLastModifiedByOk() (*UserBasic, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *Insight) SetLastModifiedBy(v UserBasic)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetIsSample

`func (o *Insight) GetIsSample() bool`

GetIsSample returns the IsSample field if non-nil, zero value otherwise.

### GetIsSampleOk

`func (o *Insight) GetIsSampleOk() (*bool, bool)`

GetIsSampleOk returns a tuple with the IsSample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSample

`func (o *Insight) SetIsSample(v bool)`

SetIsSample sets IsSample field to given value.


### GetEffectiveRestrictionLevel

`func (o *Insight) GetEffectiveRestrictionLevel() EffectiveRestrictionLevelEnum`

GetEffectiveRestrictionLevel returns the EffectiveRestrictionLevel field if non-nil, zero value otherwise.

### GetEffectiveRestrictionLevelOk

`func (o *Insight) GetEffectiveRestrictionLevelOk() (*EffectiveRestrictionLevelEnum, bool)`

GetEffectiveRestrictionLevelOk returns a tuple with the EffectiveRestrictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRestrictionLevel

`func (o *Insight) SetEffectiveRestrictionLevel(v EffectiveRestrictionLevelEnum)`

SetEffectiveRestrictionLevel sets EffectiveRestrictionLevel field to given value.


### GetEffectivePrivilegeLevel

`func (o *Insight) GetEffectivePrivilegeLevel() EffectivePrivilegeLevelEnum`

GetEffectivePrivilegeLevel returns the EffectivePrivilegeLevel field if non-nil, zero value otherwise.

### GetEffectivePrivilegeLevelOk

`func (o *Insight) GetEffectivePrivilegeLevelOk() (*EffectivePrivilegeLevelEnum, bool)`

GetEffectivePrivilegeLevelOk returns a tuple with the EffectivePrivilegeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePrivilegeLevel

`func (o *Insight) SetEffectivePrivilegeLevel(v EffectivePrivilegeLevelEnum)`

SetEffectivePrivilegeLevel sets EffectivePrivilegeLevel field to given value.


### GetUserAccessLevel

`func (o *Insight) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *Insight) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *Insight) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.


### SetUserAccessLevelNil

`func (o *Insight) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *Insight) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetTimezone

`func (o *Insight) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Insight) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Insight) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetIsCached

`func (o *Insight) GetIsCached() string`

GetIsCached returns the IsCached field if non-nil, zero value otherwise.

### GetIsCachedOk

`func (o *Insight) GetIsCachedOk() (*string, bool)`

GetIsCachedOk returns a tuple with the IsCached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCached

`func (o *Insight) SetIsCached(v string)`

SetIsCached sets IsCached field to given value.


### GetQueryStatus

`func (o *Insight) GetQueryStatus() string`

GetQueryStatus returns the QueryStatus field if non-nil, zero value otherwise.

### GetQueryStatusOk

`func (o *Insight) GetQueryStatusOk() (*string, bool)`

GetQueryStatusOk returns a tuple with the QueryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatus

`func (o *Insight) SetQueryStatus(v string)`

SetQueryStatus sets QueryStatus field to given value.


### GetHogql

`func (o *Insight) GetHogql() string`

GetHogql returns the Hogql field if non-nil, zero value otherwise.

### GetHogqlOk

`func (o *Insight) GetHogqlOk() (*string, bool)`

GetHogqlOk returns a tuple with the Hogql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHogql

`func (o *Insight) SetHogql(v string)`

SetHogql sets Hogql field to given value.


### GetTypes

`func (o *Insight) GetTypes() string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Insight) GetTypesOk() (*string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Insight) SetTypes(v string)`

SetTypes sets Types field to given value.


### GetCreateInFolder

`func (o *Insight) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *Insight) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *Insight) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *Insight) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetAlerts

`func (o *Insight) GetAlerts() string`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *Insight) GetAlertsOk() (*string, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *Insight) SetAlerts(v string)`

SetAlerts sets Alerts field to given value.


### GetLastViewedAt

`func (o *Insight) GetLastViewedAt() string`

GetLastViewedAt returns the LastViewedAt field if non-nil, zero value otherwise.

### GetLastViewedAtOk

`func (o *Insight) GetLastViewedAtOk() (*string, bool)`

GetLastViewedAtOk returns a tuple with the LastViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedAt

`func (o *Insight) SetLastViewedAt(v string)`

SetLastViewedAt sets LastViewedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


