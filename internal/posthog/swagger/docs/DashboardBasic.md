# DashboardBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **NullableString** |  | [readonly] 
**Description** | **string** |  | [readonly] 
**Pinned** | **bool** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**LastAccessedAt** | **NullableTime** |  | [readonly] 
**LastViewedAt** | **NullableTime** |  | [readonly] 
**IsShared** | **bool** |  | [readonly] 
**Deleted** | **bool** |  | [readonly] 
**CreationMode** | [**CreationModeEnum**](CreationModeEnum.md) |  | [readonly] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**RestrictionLevel** | [**DashboardRestrictionLevel**](DashboardRestrictionLevel.md) |  | [readonly] 
**EffectiveRestrictionLevel** | [**EffectiveRestrictionLevelEnum**](EffectiveRestrictionLevelEnum.md) |  | [readonly] 
**EffectivePrivilegeLevel** | [**EffectivePrivilegeLevelEnum**](EffectivePrivilegeLevelEnum.md) |  | [readonly] 
**UserAccessLevel** | **NullableString** | The effective access level the user has for this object | [readonly] 
**AccessControlVersion** | **string** |  | [readonly] 
**LastRefresh** | **NullableTime** |  | [readonly] 
**TeamId** | **int32** |  | [readonly] 

## Methods

### NewDashboardBasic

`func NewDashboardBasic(id int32, name NullableString, description string, pinned bool, createdAt time.Time, createdBy UserBasic, lastAccessedAt NullableTime, lastViewedAt NullableTime, isShared bool, deleted bool, creationMode CreationModeEnum, restrictionLevel DashboardRestrictionLevel, effectiveRestrictionLevel EffectiveRestrictionLevelEnum, effectivePrivilegeLevel EffectivePrivilegeLevelEnum, userAccessLevel NullableString, accessControlVersion string, lastRefresh NullableTime, teamId int32, ) *DashboardBasic`

NewDashboardBasic instantiates a new DashboardBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardBasicWithDefaults

`func NewDashboardBasicWithDefaults() *DashboardBasic`

NewDashboardBasicWithDefaults instantiates a new DashboardBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardBasic) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardBasic) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardBasic) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DashboardBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardBasic) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DashboardBasic) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DashboardBasic) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *DashboardBasic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardBasic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardBasic) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPinned

`func (o *DashboardBasic) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *DashboardBasic) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *DashboardBasic) SetPinned(v bool)`

SetPinned sets Pinned field to given value.


### GetCreatedAt

`func (o *DashboardBasic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DashboardBasic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DashboardBasic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DashboardBasic) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DashboardBasic) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DashboardBasic) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastAccessedAt

`func (o *DashboardBasic) GetLastAccessedAt() time.Time`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *DashboardBasic) GetLastAccessedAtOk() (*time.Time, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *DashboardBasic) SetLastAccessedAt(v time.Time)`

SetLastAccessedAt sets LastAccessedAt field to given value.


### SetLastAccessedAtNil

`func (o *DashboardBasic) SetLastAccessedAtNil(b bool)`

 SetLastAccessedAtNil sets the value for LastAccessedAt to be an explicit nil

### UnsetLastAccessedAt
`func (o *DashboardBasic) UnsetLastAccessedAt()`

UnsetLastAccessedAt ensures that no value is present for LastAccessedAt, not even an explicit nil
### GetLastViewedAt

`func (o *DashboardBasic) GetLastViewedAt() time.Time`

GetLastViewedAt returns the LastViewedAt field if non-nil, zero value otherwise.

### GetLastViewedAtOk

`func (o *DashboardBasic) GetLastViewedAtOk() (*time.Time, bool)`

GetLastViewedAtOk returns a tuple with the LastViewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewedAt

`func (o *DashboardBasic) SetLastViewedAt(v time.Time)`

SetLastViewedAt sets LastViewedAt field to given value.


### SetLastViewedAtNil

`func (o *DashboardBasic) SetLastViewedAtNil(b bool)`

 SetLastViewedAtNil sets the value for LastViewedAt to be an explicit nil

### UnsetLastViewedAt
`func (o *DashboardBasic) UnsetLastViewedAt()`

UnsetLastViewedAt ensures that no value is present for LastViewedAt, not even an explicit nil
### GetIsShared

`func (o *DashboardBasic) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *DashboardBasic) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *DashboardBasic) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.


### GetDeleted

`func (o *DashboardBasic) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DashboardBasic) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DashboardBasic) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetCreationMode

`func (o *DashboardBasic) GetCreationMode() CreationModeEnum`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *DashboardBasic) GetCreationModeOk() (*CreationModeEnum, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *DashboardBasic) SetCreationMode(v CreationModeEnum)`

SetCreationMode sets CreationMode field to given value.


### GetTags

`func (o *DashboardBasic) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardBasic) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardBasic) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardBasic) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRestrictionLevel

`func (o *DashboardBasic) GetRestrictionLevel() DashboardRestrictionLevel`

GetRestrictionLevel returns the RestrictionLevel field if non-nil, zero value otherwise.

### GetRestrictionLevelOk

`func (o *DashboardBasic) GetRestrictionLevelOk() (*DashboardRestrictionLevel, bool)`

GetRestrictionLevelOk returns a tuple with the RestrictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionLevel

`func (o *DashboardBasic) SetRestrictionLevel(v DashboardRestrictionLevel)`

SetRestrictionLevel sets RestrictionLevel field to given value.


### GetEffectiveRestrictionLevel

`func (o *DashboardBasic) GetEffectiveRestrictionLevel() EffectiveRestrictionLevelEnum`

GetEffectiveRestrictionLevel returns the EffectiveRestrictionLevel field if non-nil, zero value otherwise.

### GetEffectiveRestrictionLevelOk

`func (o *DashboardBasic) GetEffectiveRestrictionLevelOk() (*EffectiveRestrictionLevelEnum, bool)`

GetEffectiveRestrictionLevelOk returns a tuple with the EffectiveRestrictionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRestrictionLevel

`func (o *DashboardBasic) SetEffectiveRestrictionLevel(v EffectiveRestrictionLevelEnum)`

SetEffectiveRestrictionLevel sets EffectiveRestrictionLevel field to given value.


### GetEffectivePrivilegeLevel

`func (o *DashboardBasic) GetEffectivePrivilegeLevel() EffectivePrivilegeLevelEnum`

GetEffectivePrivilegeLevel returns the EffectivePrivilegeLevel field if non-nil, zero value otherwise.

### GetEffectivePrivilegeLevelOk

`func (o *DashboardBasic) GetEffectivePrivilegeLevelOk() (*EffectivePrivilegeLevelEnum, bool)`

GetEffectivePrivilegeLevelOk returns a tuple with the EffectivePrivilegeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePrivilegeLevel

`func (o *DashboardBasic) SetEffectivePrivilegeLevel(v EffectivePrivilegeLevelEnum)`

SetEffectivePrivilegeLevel sets EffectivePrivilegeLevel field to given value.


### GetUserAccessLevel

`func (o *DashboardBasic) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *DashboardBasic) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *DashboardBasic) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.


### SetUserAccessLevelNil

`func (o *DashboardBasic) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *DashboardBasic) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetAccessControlVersion

`func (o *DashboardBasic) GetAccessControlVersion() string`

GetAccessControlVersion returns the AccessControlVersion field if non-nil, zero value otherwise.

### GetAccessControlVersionOk

`func (o *DashboardBasic) GetAccessControlVersionOk() (*string, bool)`

GetAccessControlVersionOk returns a tuple with the AccessControlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlVersion

`func (o *DashboardBasic) SetAccessControlVersion(v string)`

SetAccessControlVersion sets AccessControlVersion field to given value.


### GetLastRefresh

`func (o *DashboardBasic) GetLastRefresh() time.Time`

GetLastRefresh returns the LastRefresh field if non-nil, zero value otherwise.

### GetLastRefreshOk

`func (o *DashboardBasic) GetLastRefreshOk() (*time.Time, bool)`

GetLastRefreshOk returns a tuple with the LastRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefresh

`func (o *DashboardBasic) SetLastRefresh(v time.Time)`

SetLastRefresh sets LastRefresh field to given value.


### SetLastRefreshNil

`func (o *DashboardBasic) SetLastRefreshNil(b bool)`

 SetLastRefreshNil sets the value for LastRefresh to be an explicit nil

### UnsetLastRefresh
`func (o *DashboardBasic) UnsetLastRefresh()`

UnsetLastRefresh ensures that no value is present for LastRefresh, not even an explicit nil
### GetTeamId

`func (o *DashboardBasic) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *DashboardBasic) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *DashboardBasic) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


