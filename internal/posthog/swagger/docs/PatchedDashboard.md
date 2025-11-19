# PatchedDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Pinned** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**LastAccessedAt** | Pointer to **NullableTime** |  | [optional] 
**LastViewedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**IsShared** | Pointer to **bool** |  | [optional] [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**CreationMode** | Pointer to [**CreationModeEnum**](CreationModeEnum.md) |  | [optional] [readonly] 
**Filters** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**BreakdownColors** | Pointer to **interface{}** |  | [optional] 
**DataColorThemeId** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**RestrictionLevel** | Pointer to [**DashboardRestrictionLevel**](DashboardRestrictionLevel.md) |  | [optional] 
**EffectiveRestrictionLevel** | Pointer to [**EffectiveRestrictionLevelEnum**](EffectiveRestrictionLevelEnum.md) |  | [optional] [readonly] 
**EffectivePrivilegeLevel** | Pointer to [**EffectivePrivilegeLevelEnum**](EffectivePrivilegeLevelEnum.md) |  | [optional] [readonly] 
**UserAccessLevel** | Pointer to **NullableString** | The effective access level the user has for this object | [optional] [readonly] 
**AccessControlVersion** | Pointer to **string** |  | [optional] [readonly] 
**LastRefresh** | Pointer to **NullableTime** |  | [optional] 
**PersistedFilters** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**PersistedVariables** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**TeamId** | Pointer to **int32** |  | [optional] [readonly] 
**Tiles** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**UseTemplate** | Pointer to **string** |  | [optional] 
**UseDashboard** | Pointer to **NullableInt32** |  | [optional] 
**DeleteInsights** | Pointer to **bool** |  | [optional] [default to false]
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedDashboard

`func NewPatchedDashboard() *PatchedDashboard`

NewPatchedDashboard instantiates a new PatchedDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDashboardWithDefaults

`func NewPatchedDashboardWithDefaults() *PatchedDashboard`

NewPatchedDashboardWithDefaults instantiates a new PatchedDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDashboard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDashboard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDashboard) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedDashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDashboard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDashboard) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedDashboard) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedDashboard) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PatchedDashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedDashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedDashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedDashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPinned

`func (o *PatchedDashboard) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *PatchedDashboard) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *PatchedDashboard) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *PatchedDashboard) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedDashboard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedDashboard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedDashboard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedDashboard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedDashboard) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedDashboard) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedDashboard) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedDashboard) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastAccessedAt

`func (o *PatchedDashboard) GetLastAccessedAt() time.Time`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *PatchedDashboard) GetLastAccessedAtOk() (*time.Time, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *PatchedDashboard) SetLastAccessedAt(v time.Time)`

SetLastAccessedAt sets LastAccessedAt field to given value.

### HasLastAccessedAt

`func (o *PatchedDashboard) HasLastAccessedAt() bool`

HasLastAccessedAt returns a boolean if a field has been set.

### SetLastAccessedAtNil

`func (o *PatchedDashboard) SetLastAccessedAtNil(b bool)`

 SetLastAccessedAtNil sets the value for LastAccessedAt to be an explicit nil

### UnsetLastAccessedAt
`func (o *PatchedDashboard) UnsetLastAccessedAt()`

UnsetLastAccessedAt ensures that no value is present for LastAccessedAt, not even an explicit nil
### GetLastViewedAt

`func (o *PatchedDashboard) GetLastViewedAt() time.Time`

GetLastViewedAt returns the LastViewedAt field if non-nil, zero value otherwise.

### GetLastViewedAtOk

`func (o *PatchedDashboard) GetLastViewedAtOk() (*time.Time, bool)`

GetLastViewedAtOk returns a tuple with the LastViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedAt

`func (o *PatchedDashboard) SetLastViewedAt(v time.Time)`

SetLastViewedAt sets LastViewedAt field to given value.

### HasLastViewedAt

`func (o *PatchedDashboard) HasLastViewedAt() bool`

HasLastViewedAt returns a boolean if a field has been set.

### SetLastViewedAtNil

`func (o *PatchedDashboard) SetLastViewedAtNil(b bool)`

 SetLastViewedAtNil sets the value for LastViewedAt to be an explicit nil

### UnsetLastViewedAt
`func (o *PatchedDashboard) UnsetLastViewedAt()`

UnsetLastViewedAt ensures that no value is present for LastViewedAt, not even an explicit nil
### GetIsShared

`func (o *PatchedDashboard) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *PatchedDashboard) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *PatchedDashboard) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *PatchedDashboard) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedDashboard) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedDashboard) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedDashboard) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedDashboard) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetCreationMode

`func (o *PatchedDashboard) GetCreationMode() CreationModeEnum`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *PatchedDashboard) GetCreationModeOk() (*CreationModeEnum, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *PatchedDashboard) SetCreationMode(v CreationModeEnum)`

SetCreationMode sets CreationMode field to given value.

### HasCreationMode

`func (o *PatchedDashboard) HasCreationMode() bool`

HasCreationMode returns a boolean if a field has been set.

### GetFilters

`func (o *PatchedDashboard) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedDashboard) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedDashboard) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedDashboard) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetVariables

`func (o *PatchedDashboard) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PatchedDashboard) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PatchedDashboard) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PatchedDashboard) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *PatchedDashboard) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *PatchedDashboard) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil
### GetBreakdownColors

`func (o *PatchedDashboard) GetBreakdownColors() interface{}`

GetBreakdownColors returns the BreakdownColors field if non-nil, zero value otherwise.

### GetBreakdownColorsOk

`func (o *PatchedDashboard) GetBreakdownColorsOk() (*interface{}, bool)`

GetBreakdownColorsOk returns a tuple with the BreakdownColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownColors

`func (o *PatchedDashboard) SetBreakdownColors(v interface{})`

SetBreakdownColors sets BreakdownColors field to given value.

### HasBreakdownColors

`func (o *PatchedDashboard) HasBreakdownColors() bool`

HasBreakdownColors returns a boolean if a field has been set.

### SetBreakdownColorsNil

`func (o *PatchedDashboard) SetBreakdownColorsNil(b bool)`

 SetBreakdownColorsNil sets the value for BreakdownColors to be an explicit nil

### UnsetBreakdownColors
`func (o *PatchedDashboard) UnsetBreakdownColors()`

UnsetBreakdownColors ensures that no value is present for BreakdownColors, not even an explicit nil
### GetDataColorThemeId

`func (o *PatchedDashboard) GetDataColorThemeId() int32`

GetDataColorThemeId returns the DataColorThemeId field if non-nil, zero value otherwise.

### GetDataColorThemeIdOk

`func (o *PatchedDashboard) GetDataColorThemeIdOk() (*int32, bool)`

GetDataColorThemeIdOk returns a tuple with the DataColorThemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataColorThemeId

`func (o *PatchedDashboard) SetDataColorThemeId(v int32)`

SetDataColorThemeId sets DataColorThemeId field to given value.

### HasDataColorThemeId

`func (o *PatchedDashboard) HasDataColorThemeId() bool`

HasDataColorThemeId returns a boolean if a field has been set.

### SetDataColorThemeIdNil

`func (o *PatchedDashboard) SetDataColorThemeIdNil(b bool)`

 SetDataColorThemeIdNil sets the value for DataColorThemeId to be an explicit nil

### UnsetDataColorThemeId
`func (o *PatchedDashboard) UnsetDataColorThemeId()`

UnsetDataColorThemeId ensures that no value is present for DataColorThemeId, not even an explicit nil
### GetTags

`func (o *PatchedDashboard) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedDashboard) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedDashboard) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedDashboard) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRestrictionLevel

`func (o *PatchedDashboard) GetRestrictionLevel() DashboardRestrictionLevel`

GetRestrictionLevel returns the RestrictionLevel field if non-nil, zero value otherwise.

### GetRestrictionLevelOk

`func (o *PatchedDashboard) GetRestrictionLevelOk() (*DashboardRestrictionLevel, bool)`

GetRestrictionLevelOk returns a tuple with the RestrictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionLevel

`func (o *PatchedDashboard) SetRestrictionLevel(v DashboardRestrictionLevel)`

SetRestrictionLevel sets RestrictionLevel field to given value.

### HasRestrictionLevel

`func (o *PatchedDashboard) HasRestrictionLevel() bool`

HasRestrictionLevel returns a boolean if a field has been set.

### GetEffectiveRestrictionLevel

`func (o *PatchedDashboard) GetEffectiveRestrictionLevel() EffectiveRestrictionLevelEnum`

GetEffectiveRestrictionLevel returns the EffectiveRestrictionLevel field if non-nil, zero value otherwise.

### GetEffectiveRestrictionLevelOk

`func (o *PatchedDashboard) GetEffectiveRestrictionLevelOk() (*EffectiveRestrictionLevelEnum, bool)`

GetEffectiveRestrictionLevelOk returns a tuple with the EffectiveRestrictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRestrictionLevel

`func (o *PatchedDashboard) SetEffectiveRestrictionLevel(v EffectiveRestrictionLevelEnum)`

SetEffectiveRestrictionLevel sets EffectiveRestrictionLevel field to given value.

### HasEffectiveRestrictionLevel

`func (o *PatchedDashboard) HasEffectiveRestrictionLevel() bool`

HasEffectiveRestrictionLevel returns a boolean if a field has been set.

### GetEffectivePrivilegeLevel

`func (o *PatchedDashboard) GetEffectivePrivilegeLevel() EffectivePrivilegeLevelEnum`

GetEffectivePrivilegeLevel returns the EffectivePrivilegeLevel field if non-nil, zero value otherwise.

### GetEffectivePrivilegeLevelOk

`func (o *PatchedDashboard) GetEffectivePrivilegeLevelOk() (*EffectivePrivilegeLevelEnum, bool)`

GetEffectivePrivilegeLevelOk returns a tuple with the EffectivePrivilegeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePrivilegeLevel

`func (o *PatchedDashboard) SetEffectivePrivilegeLevel(v EffectivePrivilegeLevelEnum)`

SetEffectivePrivilegeLevel sets EffectivePrivilegeLevel field to given value.

### HasEffectivePrivilegeLevel

`func (o *PatchedDashboard) HasEffectivePrivilegeLevel() bool`

HasEffectivePrivilegeLevel returns a boolean if a field has been set.

### GetUserAccessLevel

`func (o *PatchedDashboard) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *PatchedDashboard) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *PatchedDashboard) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.

### HasUserAccessLevel

`func (o *PatchedDashboard) HasUserAccessLevel() bool`

HasUserAccessLevel returns a boolean if a field has been set.

### SetUserAccessLevelNil

`func (o *PatchedDashboard) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *PatchedDashboard) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetAccessControlVersion

`func (o *PatchedDashboard) GetAccessControlVersion() string`

GetAccessControlVersion returns the AccessControlVersion field if non-nil, zero value otherwise.

### GetAccessControlVersionOk

`func (o *PatchedDashboard) GetAccessControlVersionOk() (*string, bool)`

GetAccessControlVersionOk returns a tuple with the AccessControlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlVersion

`func (o *PatchedDashboard) SetAccessControlVersion(v string)`

SetAccessControlVersion sets AccessControlVersion field to given value.

### HasAccessControlVersion

`func (o *PatchedDashboard) HasAccessControlVersion() bool`

HasAccessControlVersion returns a boolean if a field has been set.

### GetLastRefresh

`func (o *PatchedDashboard) GetLastRefresh() time.Time`

GetLastRefresh returns the LastRefresh field if non-nil, zero value otherwise.

### GetLastRefreshOk

`func (o *PatchedDashboard) GetLastRefreshOk() (*time.Time, bool)`

GetLastRefreshOk returns a tuple with the LastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefresh

`func (o *PatchedDashboard) SetLastRefresh(v time.Time)`

SetLastRefresh sets LastRefresh field to given value.

### HasLastRefresh

`func (o *PatchedDashboard) HasLastRefresh() bool`

HasLastRefresh returns a boolean if a field has been set.

### SetLastRefreshNil

`func (o *PatchedDashboard) SetLastRefreshNil(b bool)`

 SetLastRefreshNil sets the value for LastRefresh to be an explicit nil

### UnsetLastRefresh
`func (o *PatchedDashboard) UnsetLastRefresh()`

UnsetLastRefresh ensures that no value is present for LastRefresh, not even an explicit nil
### GetPersistedFilters

`func (o *PatchedDashboard) GetPersistedFilters() map[string]interface{}`

GetPersistedFilters returns the PersistedFilters field if non-nil, zero value otherwise.

### GetPersistedFiltersOk

`func (o *PatchedDashboard) GetPersistedFiltersOk() (*map[string]interface{}, bool)`

GetPersistedFiltersOk returns a tuple with the PersistedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistedFilters

`func (o *PatchedDashboard) SetPersistedFilters(v map[string]interface{})`

SetPersistedFilters sets PersistedFilters field to given value.

### HasPersistedFilters

`func (o *PatchedDashboard) HasPersistedFilters() bool`

HasPersistedFilters returns a boolean if a field has been set.

### SetPersistedFiltersNil

`func (o *PatchedDashboard) SetPersistedFiltersNil(b bool)`

 SetPersistedFiltersNil sets the value for PersistedFilters to be an explicit nil

### UnsetPersistedFilters
`func (o *PatchedDashboard) UnsetPersistedFilters()`

UnsetPersistedFilters ensures that no value is present for PersistedFilters, not even an explicit nil
### GetPersistedVariables

`func (o *PatchedDashboard) GetPersistedVariables() map[string]interface{}`

GetPersistedVariables returns the PersistedVariables field if non-nil, zero value otherwise.

### GetPersistedVariablesOk

`func (o *PatchedDashboard) GetPersistedVariablesOk() (*map[string]interface{}, bool)`

GetPersistedVariablesOk returns a tuple with the PersistedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistedVariables

`func (o *PatchedDashboard) SetPersistedVariables(v map[string]interface{})`

SetPersistedVariables sets PersistedVariables field to given value.

### HasPersistedVariables

`func (o *PatchedDashboard) HasPersistedVariables() bool`

HasPersistedVariables returns a boolean if a field has been set.

### SetPersistedVariablesNil

`func (o *PatchedDashboard) SetPersistedVariablesNil(b bool)`

 SetPersistedVariablesNil sets the value for PersistedVariables to be an explicit nil

### UnsetPersistedVariables
`func (o *PatchedDashboard) UnsetPersistedVariables()`

UnsetPersistedVariables ensures that no value is present for PersistedVariables, not even an explicit nil
### GetTeamId

`func (o *PatchedDashboard) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PatchedDashboard) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PatchedDashboard) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PatchedDashboard) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTiles

`func (o *PatchedDashboard) GetTiles() []map[string]interface{}`

GetTiles returns the Tiles field if non-nil, zero value otherwise.

### GetTilesOk

`func (o *PatchedDashboard) GetTilesOk() (*[]map[string]interface{}, bool)`

GetTilesOk returns a tuple with the Tiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiles

`func (o *PatchedDashboard) SetTiles(v []map[string]interface{})`

SetTiles sets Tiles field to given value.

### HasTiles

`func (o *PatchedDashboard) HasTiles() bool`

HasTiles returns a boolean if a field has been set.

### SetTilesNil

`func (o *PatchedDashboard) SetTilesNil(b bool)`

 SetTilesNil sets the value for Tiles to be an explicit nil

### UnsetTiles
`func (o *PatchedDashboard) UnsetTiles()`

UnsetTiles ensures that no value is present for Tiles, not even an explicit nil
### GetUseTemplate

`func (o *PatchedDashboard) GetUseTemplate() string`

GetUseTemplate returns the UseTemplate field if non-nil, zero value otherwise.

### GetUseTemplateOk

`func (o *PatchedDashboard) GetUseTemplateOk() (*string, bool)`

GetUseTemplateOk returns a tuple with the UseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTemplate

`func (o *PatchedDashboard) SetUseTemplate(v string)`

SetUseTemplate sets UseTemplate field to given value.

### HasUseTemplate

`func (o *PatchedDashboard) HasUseTemplate() bool`

HasUseTemplate returns a boolean if a field has been set.

### GetUseDashboard

`func (o *PatchedDashboard) GetUseDashboard() int32`

GetUseDashboard returns the UseDashboard field if non-nil, zero value otherwise.

### GetUseDashboardOk

`func (o *PatchedDashboard) GetUseDashboardOk() (*int32, bool)`

GetUseDashboardOk returns a tuple with the UseDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDashboard

`func (o *PatchedDashboard) SetUseDashboard(v int32)`

SetUseDashboard sets UseDashboard field to given value.

### HasUseDashboard

`func (o *PatchedDashboard) HasUseDashboard() bool`

HasUseDashboard returns a boolean if a field has been set.

### SetUseDashboardNil

`func (o *PatchedDashboard) SetUseDashboardNil(b bool)`

 SetUseDashboardNil sets the value for UseDashboard to be an explicit nil

### UnsetUseDashboard
`func (o *PatchedDashboard) UnsetUseDashboard()`

UnsetUseDashboard ensures that no value is present for UseDashboard, not even an explicit nil
### GetDeleteInsights

`func (o *PatchedDashboard) GetDeleteInsights() bool`

GetDeleteInsights returns the DeleteInsights field if non-nil, zero value otherwise.

### GetDeleteInsightsOk

`func (o *PatchedDashboard) GetDeleteInsightsOk() (*bool, bool)`

GetDeleteInsightsOk returns a tuple with the DeleteInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInsights

`func (o *PatchedDashboard) SetDeleteInsights(v bool)`

SetDeleteInsights sets DeleteInsights field to given value.

### HasDeleteInsights

`func (o *PatchedDashboard) HasDeleteInsights() bool`

HasDeleteInsights returns a boolean if a field has been set.

### GetCreateInFolder

`func (o *PatchedDashboard) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedDashboard) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedDashboard) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedDashboard) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


