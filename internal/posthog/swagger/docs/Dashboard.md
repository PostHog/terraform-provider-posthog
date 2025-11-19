# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Pinned** | Pointer to **bool** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**LastAccessedAt** | Pointer to **NullableTime** |  | [optional] 
**LastViewedAt** | **NullableTime** |  | [readonly] 
**IsShared** | **bool** |  | [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**CreationMode** | [**CreationModeEnum**](CreationModeEnum.md) |  | [readonly] 
**Filters** | **map[string]interface{}** |  | [readonly] 
**Variables** | **map[string]interface{}** |  | [readonly] 
**BreakdownColors** | Pointer to **interface{}** |  | [optional] 
**DataColorThemeId** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**RestrictionLevel** | Pointer to [**DashboardRestrictionLevel**](DashboardRestrictionLevel.md) |  | [optional] 
**EffectiveRestrictionLevel** | [**EffectiveRestrictionLevelEnum**](EffectiveRestrictionLevelEnum.md) |  | [readonly] 
**EffectivePrivilegeLevel** | [**EffectivePrivilegeLevelEnum**](EffectivePrivilegeLevelEnum.md) |  | [readonly] 
**UserAccessLevel** | **NullableString** | The effective access level the user has for this object | [readonly] 
**AccessControlVersion** | **string** |  | [readonly] 
**LastRefresh** | Pointer to **NullableTime** |  | [optional] 
**PersistedFilters** | **map[string]interface{}** |  | [readonly] 
**PersistedVariables** | **map[string]interface{}** |  | [readonly] 
**TeamId** | **int32** |  | [readonly] 
**Tiles** | **[]map[string]interface{}** |  | [readonly] 
**UseTemplate** | Pointer to **string** |  | [optional] 
**UseDashboard** | Pointer to **NullableInt32** |  | [optional] 
**DeleteInsights** | Pointer to **bool** |  | [optional] [default to false]
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewDashboard

`func NewDashboard(id int32, createdAt time.Time, createdBy UserBasic, lastViewedAt NullableTime, isShared bool, creationMode CreationModeEnum, filters map[string]interface{}, variables map[string]interface{}, effectiveRestrictionLevel EffectiveRestrictionLevelEnum, effectivePrivilegeLevel EffectivePrivilegeLevelEnum, userAccessLevel NullableString, accessControlVersion string, persistedFilters map[string]interface{}, persistedVariables map[string]interface{}, teamId int32, tiles []map[string]interface{}, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dashboard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Dashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashboard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dashboard) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Dashboard) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Dashboard) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Dashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPinned

`func (o *Dashboard) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *Dashboard) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *Dashboard) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *Dashboard) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Dashboard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dashboard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Dashboard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Dashboard) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Dashboard) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Dashboard) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastAccessedAt

`func (o *Dashboard) GetLastAccessedAt() time.Time`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *Dashboard) GetLastAccessedAtOk() (*time.Time, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *Dashboard) SetLastAccessedAt(v time.Time)`

SetLastAccessedAt sets LastAccessedAt field to given value.

### HasLastAccessedAt

`func (o *Dashboard) HasLastAccessedAt() bool`

HasLastAccessedAt returns a boolean if a field has been set.

### SetLastAccessedAtNil

`func (o *Dashboard) SetLastAccessedAtNil(b bool)`

 SetLastAccessedAtNil sets the value for LastAccessedAt to be an explicit nil

### UnsetLastAccessedAt
`func (o *Dashboard) UnsetLastAccessedAt()`

UnsetLastAccessedAt ensures that no value is present for LastAccessedAt, not even an explicit nil
### GetLastViewedAt

`func (o *Dashboard) GetLastViewedAt() time.Time`

GetLastViewedAt returns the LastViewedAt field if non-nil, zero value otherwise.

### GetLastViewedAtOk

`func (o *Dashboard) GetLastViewedAtOk() (*time.Time, bool)`

GetLastViewedAtOk returns a tuple with the LastViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedAt

`func (o *Dashboard) SetLastViewedAt(v time.Time)`

SetLastViewedAt sets LastViewedAt field to given value.


### SetLastViewedAtNil

`func (o *Dashboard) SetLastViewedAtNil(b bool)`

 SetLastViewedAtNil sets the value for LastViewedAt to be an explicit nil

### UnsetLastViewedAt
`func (o *Dashboard) UnsetLastViewedAt()`

UnsetLastViewedAt ensures that no value is present for LastViewedAt, not even an explicit nil
### GetIsShared

`func (o *Dashboard) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *Dashboard) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *Dashboard) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.


### GetDeleted

`func (o *Dashboard) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Dashboard) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Dashboard) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Dashboard) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetCreationMode

`func (o *Dashboard) GetCreationMode() CreationModeEnum`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *Dashboard) GetCreationModeOk() (*CreationModeEnum, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *Dashboard) SetCreationMode(v CreationModeEnum)`

SetCreationMode sets CreationMode field to given value.


### GetFilters

`func (o *Dashboard) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Dashboard) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Dashboard) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.


### GetVariables

`func (o *Dashboard) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Dashboard) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Dashboard) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.


### SetVariablesNil

`func (o *Dashboard) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *Dashboard) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetBreakdownColors

`func (o *Dashboard) GetBreakdownColors() interface{}`

GetBreakdownColors returns the BreakdownColors field if non-nil, zero value otherwise.

### GetBreakdownColorsOk

`func (o *Dashboard) GetBreakdownColorsOk() (*interface{}, bool)`

GetBreakdownColorsOk returns a tuple with the BreakdownColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownColors

`func (o *Dashboard) SetBreakdownColors(v interface{})`

SetBreakdownColors sets BreakdownColors field to given value.

### HasBreakdownColors

`func (o *Dashboard) HasBreakdownColors() bool`

HasBreakdownColors returns a boolean if a field has been set.

### SetBreakdownColorsNil

`func (o *Dashboard) SetBreakdownColorsNil(b bool)`

 SetBreakdownColorsNil sets the value for BreakdownColors to be an explicit nil

### UnsetBreakdownColors
`func (o *Dashboard) UnsetBreakdownColors()`

UnsetBreakdownColors ensures that no value is present for BreakdownColors, not even an explicit nil
### GetDataColorThemeId

`func (o *Dashboard) GetDataColorThemeId() int32`

GetDataColorThemeId returns the DataColorThemeId field if non-nil, zero value otherwise.

### GetDataColorThemeIdOk

`func (o *Dashboard) GetDataColorThemeIdOk() (*int32, bool)`

GetDataColorThemeIdOk returns a tuple with the DataColorThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorThemeId

`func (o *Dashboard) SetDataColorThemeId(v int32)`

SetDataColorThemeId sets DataColorThemeId field to given value.

### HasDataColorThemeId

`func (o *Dashboard) HasDataColorThemeId() bool`

HasDataColorThemeId returns a boolean if a field has been set.

### SetDataColorThemeIdNil

`func (o *Dashboard) SetDataColorThemeIdNil(b bool)`

 SetDataColorThemeIdNil sets the value for DataColorThemeId to be an explicit nil

### UnsetDataColorThemeId
`func (o *Dashboard) UnsetDataColorThemeId()`

UnsetDataColorThemeId ensures that no value is present for DataColorThemeId, not even an explicit nil
### GetTags

`func (o *Dashboard) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Dashboard) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Dashboard) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Dashboard) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRestrictionLevel

`func (o *Dashboard) GetRestrictionLevel() DashboardRestrictionLevel`

GetRestrictionLevel returns the RestrictionLevel field if non-nil, zero value otherwise.

### GetRestrictionLevelOk

`func (o *Dashboard) GetRestrictionLevelOk() (*DashboardRestrictionLevel, bool)`

GetRestrictionLevelOk returns a tuple with the RestrictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionLevel

`func (o *Dashboard) SetRestrictionLevel(v DashboardRestrictionLevel)`

SetRestrictionLevel sets RestrictionLevel field to given value.

### HasRestrictionLevel

`func (o *Dashboard) HasRestrictionLevel() bool`

HasRestrictionLevel returns a boolean if a field has been set.

### GetEffectiveRestrictionLevel

`func (o *Dashboard) GetEffectiveRestrictionLevel() EffectiveRestrictionLevelEnum`

GetEffectiveRestrictionLevel returns the EffectiveRestrictionLevel field if non-nil, zero value otherwise.

### GetEffectiveRestrictionLevelOk

`func (o *Dashboard) GetEffectiveRestrictionLevelOk() (*EffectiveRestrictionLevelEnum, bool)`

GetEffectiveRestrictionLevelOk returns a tuple with the EffectiveRestrictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRestrictionLevel

`func (o *Dashboard) SetEffectiveRestrictionLevel(v EffectiveRestrictionLevelEnum)`

SetEffectiveRestrictionLevel sets EffectiveRestrictionLevel field to given value.


### GetEffectivePrivilegeLevel

`func (o *Dashboard) GetEffectivePrivilegeLevel() EffectivePrivilegeLevelEnum`

GetEffectivePrivilegeLevel returns the EffectivePrivilegeLevel field if non-nil, zero value otherwise.

### GetEffectivePrivilegeLevelOk

`func (o *Dashboard) GetEffectivePrivilegeLevelOk() (*EffectivePrivilegeLevelEnum, bool)`

GetEffectivePrivilegeLevelOk returns a tuple with the EffectivePrivilegeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePrivilegeLevel

`func (o *Dashboard) SetEffectivePrivilegeLevel(v EffectivePrivilegeLevelEnum)`

SetEffectivePrivilegeLevel sets EffectivePrivilegeLevel field to given value.


### GetUserAccessLevel

`func (o *Dashboard) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *Dashboard) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *Dashboard) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.


### SetUserAccessLevelNil

`func (o *Dashboard) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *Dashboard) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetAccessControlVersion

`func (o *Dashboard) GetAccessControlVersion() string`

GetAccessControlVersion returns the AccessControlVersion field if non-nil, zero value otherwise.

### GetAccessControlVersionOk

`func (o *Dashboard) GetAccessControlVersionOk() (*string, bool)`

GetAccessControlVersionOk returns a tuple with the AccessControlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlVersion

`func (o *Dashboard) SetAccessControlVersion(v string)`

SetAccessControlVersion sets AccessControlVersion field to given value.


### GetLastRefresh

`func (o *Dashboard) GetLastRefresh() time.Time`

GetLastRefresh returns the LastRefresh field if non-nil, zero value otherwise.

### GetLastRefreshOk

`func (o *Dashboard) GetLastRefreshOk() (*time.Time, bool)`

GetLastRefreshOk returns a tuple with the LastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefresh

`func (o *Dashboard) SetLastRefresh(v time.Time)`

SetLastRefresh sets LastRefresh field to given value.

### HasLastRefresh

`func (o *Dashboard) HasLastRefresh() bool`

HasLastRefresh returns a boolean if a field has been set.

### SetLastRefreshNil

`func (o *Dashboard) SetLastRefreshNil(b bool)`

 SetLastRefreshNil sets the value for LastRefresh to be an explicit nil

### UnsetLastRefresh
`func (o *Dashboard) UnsetLastRefresh()`

UnsetLastRefresh ensures that no value is present for LastRefresh, not even an explicit nil
### GetPersistedFilters

`func (o *Dashboard) GetPersistedFilters() map[string]interface{}`

GetPersistedFilters returns the PersistedFilters field if non-nil, zero value otherwise.

### GetPersistedFiltersOk

`func (o *Dashboard) GetPersistedFiltersOk() (*map[string]interface{}, bool)`

GetPersistedFiltersOk returns a tuple with the PersistedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistedFilters

`func (o *Dashboard) SetPersistedFilters(v map[string]interface{})`

SetPersistedFilters sets PersistedFilters field to given value.


### SetPersistedFiltersNil

`func (o *Dashboard) SetPersistedFiltersNil(b bool)`

 SetPersistedFiltersNil sets the value for PersistedFilters to be an explicit nil

### UnsetPersistedFilters
`func (o *Dashboard) UnsetPersistedFilters()`

UnsetPersistedFilters ensures that no value is present for PersistedFilters, not even an explicit nil
### GetPersistedVariables

`func (o *Dashboard) GetPersistedVariables() map[string]interface{}`

GetPersistedVariables returns the PersistedVariables field if non-nil, zero value otherwise.

### GetPersistedVariablesOk

`func (o *Dashboard) GetPersistedVariablesOk() (*map[string]interface{}, bool)`

GetPersistedVariablesOk returns a tuple with the PersistedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistedVariables

`func (o *Dashboard) SetPersistedVariables(v map[string]interface{})`

SetPersistedVariables sets PersistedVariables field to given value.


### SetPersistedVariablesNil

`func (o *Dashboard) SetPersistedVariablesNil(b bool)`

 SetPersistedVariablesNil sets the value for PersistedVariables to be an explicit nil

### UnsetPersistedVariables
`func (o *Dashboard) UnsetPersistedVariables()`

UnsetPersistedVariables ensures that no value is present for PersistedVariables, not even an explicit nil
### GetTeamId

`func (o *Dashboard) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Dashboard) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Dashboard) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetTiles

`func (o *Dashboard) GetTiles() []map[string]interface{}`

GetTiles returns the Tiles field if non-nil, zero value otherwise.

### GetTilesOk

`func (o *Dashboard) GetTilesOk() (*[]map[string]interface{}, bool)`

GetTilesOk returns a tuple with the Tiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiles

`func (o *Dashboard) SetTiles(v []map[string]interface{})`

SetTiles sets Tiles field to given value.


### SetTilesNil

`func (o *Dashboard) SetTilesNil(b bool)`

 SetTilesNil sets the value for Tiles to be an explicit nil

### UnsetTiles
`func (o *Dashboard) UnsetTiles()`

UnsetTiles ensures that no value is present for Tiles, not even an explicit nil
### GetUseTemplate

`func (o *Dashboard) GetUseTemplate() string`

GetUseTemplate returns the UseTemplate field if non-nil, zero value otherwise.

### GetUseTemplateOk

`func (o *Dashboard) GetUseTemplateOk() (*string, bool)`

GetUseTemplateOk returns a tuple with the UseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTemplate

`func (o *Dashboard) SetUseTemplate(v string)`

SetUseTemplate sets UseTemplate field to given value.

### HasUseTemplate

`func (o *Dashboard) HasUseTemplate() bool`

HasUseTemplate returns a boolean if a field has been set.

### GetUseDashboard

`func (o *Dashboard) GetUseDashboard() int32`

GetUseDashboard returns the UseDashboard field if non-nil, zero value otherwise.

### GetUseDashboardOk

`func (o *Dashboard) GetUseDashboardOk() (*int32, bool)`

GetUseDashboardOk returns a tuple with the UseDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDashboard

`func (o *Dashboard) SetUseDashboard(v int32)`

SetUseDashboard sets UseDashboard field to given value.

### HasUseDashboard

`func (o *Dashboard) HasUseDashboard() bool`

HasUseDashboard returns a boolean if a field has been set.

### SetUseDashboardNil

`func (o *Dashboard) SetUseDashboardNil(b bool)`

 SetUseDashboardNil sets the value for UseDashboard to be an explicit nil

### UnsetUseDashboard
`func (o *Dashboard) UnsetUseDashboard()`

UnsetUseDashboard ensures that no value is present for UseDashboard, not even an explicit nil
### GetDeleteInsights

`func (o *Dashboard) GetDeleteInsights() bool`

GetDeleteInsights returns the DeleteInsights field if non-nil, zero value otherwise.

### GetDeleteInsightsOk

`func (o *Dashboard) GetDeleteInsightsOk() (*bool, bool)`

GetDeleteInsightsOk returns a tuple with the DeleteInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInsights

`func (o *Dashboard) SetDeleteInsights(v bool)`

SetDeleteInsights sets DeleteInsights field to given value.

### HasDeleteInsights

`func (o *Dashboard) HasDeleteInsights() bool`

HasDeleteInsights returns a boolean if a field has been set.

### GetCreateInFolder

`func (o *Dashboard) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *Dashboard) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *Dashboard) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *Dashboard) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


