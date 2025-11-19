# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateJoined** | **time.Time** |  | [readonly] 
**Uuid** | **string** |  | [readonly] 
**DistinctId** | **NullableString** |  | [readonly] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**PendingEmail** | **NullableString** |  | [readonly] 
**IsEmailVerified** | **NullableBool** |  | [readonly] 
**NotificationSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**AnonymizeData** | Pointer to **NullableBool** |  | [optional] 
**ToolbarMode** | Pointer to **NullableString** |  | [optional] 
**HasPassword** | **bool** |  | [readonly] 
**Id** | **int32** |  | [readonly] 
**IsStaff** | Pointer to **bool** | Designates whether the user can log into this admin site. | [optional] 
**IsImpersonated** | **NullableBool** |  | [readonly] 
**IsImpersonatedUntil** | **NullableString** |  | [readonly] 
**SensitiveSessionExpiresAt** | **NullableString** |  | [readonly] 
**Team** | [**TeamBasic**](TeamBasic.md) |  | [readonly] 
**Organization** | [**Organization**](Organization.md) |  | [readonly] 
**Organizations** | [**[]OrganizationBasic**](OrganizationBasic.md) |  | [readonly] 
**SetCurrentOrganization** | Pointer to **string** |  | [optional] 
**SetCurrentTeam** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 
**CurrentPassword** | Pointer to **string** |  | [optional] 
**EventsColumnConfig** | Pointer to **interface{}** |  | [optional] 
**Is2faEnabled** | **bool** |  | [readonly] 
**HasSocialAuth** | **bool** |  | [readonly] 
**HasSsoEnforcement** | **bool** |  | [readonly] 
**HasSeenProductIntroFor** | Pointer to **interface{}** |  | [optional] 
**ScenePersonalisation** | [**[]ScenePersonalisationBasic**](ScenePersonalisationBasic.md) |  | [readonly] 
**ThemeMode** | Pointer to **NullableString** |  | [optional] 
**HedgehogConfig** | Pointer to **interface{}** |  | [optional] 
**RoleAtOrganization** | Pointer to [**RoleAtOrganizationEnum**](RoleAtOrganizationEnum.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(dateJoined time.Time, uuid string, distinctId NullableString, email string, pendingEmail NullableString, isEmailVerified NullableBool, hasPassword bool, id int32, isImpersonated NullableBool, isImpersonatedUntil NullableString, sensitiveSessionExpiresAt NullableString, team TeamBasic, organization Organization, organizations []OrganizationBasic, password string, is2faEnabled bool, hasSocialAuth bool, hasSsoEnforcement bool, scenePersonalisation []ScenePersonalisationBasic, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateJoined

`func (o *User) GetDateJoined() time.Time`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *User) GetDateJoinedOk() (*time.Time, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *User) SetDateJoined(v time.Time)`

SetDateJoined sets DateJoined field to given value.


### GetUuid

`func (o *User) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *User) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *User) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetDistinctId

`func (o *User) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *User) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *User) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.


### SetDistinctIdNil

`func (o *User) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *User) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPendingEmail

`func (o *User) GetPendingEmail() string`

GetPendingEmail returns the PendingEmail field if non-nil, zero value otherwise.

### GetPendingEmailOk

`func (o *User) GetPendingEmailOk() (*string, bool)`

GetPendingEmailOk returns a tuple with the PendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingEmail

`func (o *User) SetPendingEmail(v string)`

SetPendingEmail sets PendingEmail field to given value.


### SetPendingEmailNil

`func (o *User) SetPendingEmailNil(b bool)`

 SetPendingEmailNil sets the value for PendingEmail to be an explicit nil

### UnsetPendingEmail
`func (o *User) UnsetPendingEmail()`

UnsetPendingEmail ensures that no value is present for PendingEmail, not even an explicit nil
### GetIsEmailVerified

`func (o *User) GetIsEmailVerified() bool`

GetIsEmailVerified returns the IsEmailVerified field if non-nil, zero value otherwise.

### GetIsEmailVerifiedOk

`func (o *User) GetIsEmailVerifiedOk() (*bool, bool)`

GetIsEmailVerifiedOk returns a tuple with the IsEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailVerified

`func (o *User) SetIsEmailVerified(v bool)`

SetIsEmailVerified sets IsEmailVerified field to given value.


### SetIsEmailVerifiedNil

`func (o *User) SetIsEmailVerifiedNil(b bool)`

 SetIsEmailVerifiedNil sets the value for IsEmailVerified to be an explicit nil

### UnsetIsEmailVerified
`func (o *User) UnsetIsEmailVerified()`

UnsetIsEmailVerified ensures that no value is present for IsEmailVerified, not even an explicit nil
### GetNotificationSettings

`func (o *User) GetNotificationSettings() map[string]interface{}`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *User) GetNotificationSettingsOk() (*map[string]interface{}, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *User) SetNotificationSettings(v map[string]interface{})`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *User) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetAnonymizeData

`func (o *User) GetAnonymizeData() bool`

GetAnonymizeData returns the AnonymizeData field if non-nil, zero value otherwise.

### GetAnonymizeDataOk

`func (o *User) GetAnonymizeDataOk() (*bool, bool)`

GetAnonymizeDataOk returns a tuple with the AnonymizeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeData

`func (o *User) SetAnonymizeData(v bool)`

SetAnonymizeData sets AnonymizeData field to given value.

### HasAnonymizeData

`func (o *User) HasAnonymizeData() bool`

HasAnonymizeData returns a boolean if a field has been set.

### SetAnonymizeDataNil

`func (o *User) SetAnonymizeDataNil(b bool)`

 SetAnonymizeDataNil sets the value for AnonymizeData to be an explicit nil

### UnsetAnonymizeData
`func (o *User) UnsetAnonymizeData()`

UnsetAnonymizeData ensures that no value is present for AnonymizeData, not even an explicit nil
### GetToolbarMode

`func (o *User) GetToolbarMode() string`

GetToolbarMode returns the ToolbarMode field if non-nil, zero value otherwise.

### GetToolbarModeOk

`func (o *User) GetToolbarModeOk() (*string, bool)`

GetToolbarModeOk returns a tuple with the ToolbarMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarMode

`func (o *User) SetToolbarMode(v string)`

SetToolbarMode sets ToolbarMode field to given value.

### HasToolbarMode

`func (o *User) HasToolbarMode() bool`

HasToolbarMode returns a boolean if a field has been set.

### SetToolbarModeNil

`func (o *User) SetToolbarModeNil(b bool)`

 SetToolbarModeNil sets the value for ToolbarMode to be an explicit nil

### UnsetToolbarMode
`func (o *User) UnsetToolbarMode()`

UnsetToolbarMode ensures that no value is present for ToolbarMode, not even an explicit nil
### GetHasPassword

`func (o *User) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *User) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *User) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.


### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetIsStaff

`func (o *User) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *User) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *User) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *User) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsImpersonated

`func (o *User) GetIsImpersonated() bool`

GetIsImpersonated returns the IsImpersonated field if non-nil, zero value otherwise.

### GetIsImpersonatedOk

`func (o *User) GetIsImpersonatedOk() (*bool, bool)`

GetIsImpersonatedOk returns a tuple with the IsImpersonated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImpersonated

`func (o *User) SetIsImpersonated(v bool)`

SetIsImpersonated sets IsImpersonated field to given value.


### SetIsImpersonatedNil

`func (o *User) SetIsImpersonatedNil(b bool)`

 SetIsImpersonatedNil sets the value for IsImpersonated to be an explicit nil

### UnsetIsImpersonated
`func (o *User) UnsetIsImpersonated()`

UnsetIsImpersonated ensures that no value is present for IsImpersonated, not even an explicit nil
### GetIsImpersonatedUntil

`func (o *User) GetIsImpersonatedUntil() string`

GetIsImpersonatedUntil returns the IsImpersonatedUntil field if non-nil, zero value otherwise.

### GetIsImpersonatedUntilOk

`func (o *User) GetIsImpersonatedUntilOk() (*string, bool)`

GetIsImpersonatedUntilOk returns a tuple with the IsImpersonatedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImpersonatedUntil

`func (o *User) SetIsImpersonatedUntil(v string)`

SetIsImpersonatedUntil sets IsImpersonatedUntil field to given value.


### SetIsImpersonatedUntilNil

`func (o *User) SetIsImpersonatedUntilNil(b bool)`

 SetIsImpersonatedUntilNil sets the value for IsImpersonatedUntil to be an explicit nil

### UnsetIsImpersonatedUntil
`func (o *User) UnsetIsImpersonatedUntil()`

UnsetIsImpersonatedUntil ensures that no value is present for IsImpersonatedUntil, not even an explicit nil
### GetSensitiveSessionExpiresAt

`func (o *User) GetSensitiveSessionExpiresAt() string`

GetSensitiveSessionExpiresAt returns the SensitiveSessionExpiresAt field if non-nil, zero value otherwise.

### GetSensitiveSessionExpiresAtOk

`func (o *User) GetSensitiveSessionExpiresAtOk() (*string, bool)`

GetSensitiveSessionExpiresAtOk returns a tuple with the SensitiveSessionExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveSessionExpiresAt

`func (o *User) SetSensitiveSessionExpiresAt(v string)`

SetSensitiveSessionExpiresAt sets SensitiveSessionExpiresAt field to given value.


### SetSensitiveSessionExpiresAtNil

`func (o *User) SetSensitiveSessionExpiresAtNil(b bool)`

 SetSensitiveSessionExpiresAtNil sets the value for SensitiveSessionExpiresAt to be an explicit nil

### UnsetSensitiveSessionExpiresAt
`func (o *User) UnsetSensitiveSessionExpiresAt()`

UnsetSensitiveSessionExpiresAt ensures that no value is present for SensitiveSessionExpiresAt, not even an explicit nil
### GetTeam

`func (o *User) GetTeam() TeamBasic`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *User) GetTeamOk() (*TeamBasic, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *User) SetTeam(v TeamBasic)`

SetTeam sets Team field to given value.


### GetOrganization

`func (o *User) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *User) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *User) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.


### GetOrganizations

`func (o *User) GetOrganizations() []OrganizationBasic`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *User) GetOrganizationsOk() (*[]OrganizationBasic, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *User) SetOrganizations(v []OrganizationBasic)`

SetOrganizations sets Organizations field to given value.


### GetSetCurrentOrganization

`func (o *User) GetSetCurrentOrganization() string`

GetSetCurrentOrganization returns the SetCurrentOrganization field if non-nil, zero value otherwise.

### GetSetCurrentOrganizationOk

`func (o *User) GetSetCurrentOrganizationOk() (*string, bool)`

GetSetCurrentOrganizationOk returns a tuple with the SetCurrentOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCurrentOrganization

`func (o *User) SetSetCurrentOrganization(v string)`

SetSetCurrentOrganization sets SetCurrentOrganization field to given value.

### HasSetCurrentOrganization

`func (o *User) HasSetCurrentOrganization() bool`

HasSetCurrentOrganization returns a boolean if a field has been set.

### GetSetCurrentTeam

`func (o *User) GetSetCurrentTeam() string`

GetSetCurrentTeam returns the SetCurrentTeam field if non-nil, zero value otherwise.

### GetSetCurrentTeamOk

`func (o *User) GetSetCurrentTeamOk() (*string, bool)`

GetSetCurrentTeamOk returns a tuple with the SetCurrentTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCurrentTeam

`func (o *User) SetSetCurrentTeam(v string)`

SetSetCurrentTeam sets SetCurrentTeam field to given value.

### HasSetCurrentTeam

`func (o *User) HasSetCurrentTeam() bool`

HasSetCurrentTeam returns a boolean if a field has been set.

### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetCurrentPassword

`func (o *User) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *User) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *User) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *User) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetEventsColumnConfig

`func (o *User) GetEventsColumnConfig() interface{}`

GetEventsColumnConfig returns the EventsColumnConfig field if non-nil, zero value otherwise.

### GetEventsColumnConfigOk

`func (o *User) GetEventsColumnConfigOk() (*interface{}, bool)`

GetEventsColumnConfigOk returns a tuple with the EventsColumnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsColumnConfig

`func (o *User) SetEventsColumnConfig(v interface{})`

SetEventsColumnConfig sets EventsColumnConfig field to given value.

### HasEventsColumnConfig

`func (o *User) HasEventsColumnConfig() bool`

HasEventsColumnConfig returns a boolean if a field has been set.

### SetEventsColumnConfigNil

`func (o *User) SetEventsColumnConfigNil(b bool)`

 SetEventsColumnConfigNil sets the value for EventsColumnConfig to be an explicit nil

### UnsetEventsColumnConfig
`func (o *User) UnsetEventsColumnConfig()`

UnsetEventsColumnConfig ensures that no value is present for EventsColumnConfig, not even an explicit nil
### GetIs2faEnabled

`func (o *User) GetIs2faEnabled() bool`

GetIs2faEnabled returns the Is2faEnabled field if non-nil, zero value otherwise.

### GetIs2faEnabledOk

`func (o *User) GetIs2faEnabledOk() (*bool, bool)`

GetIs2faEnabledOk returns a tuple with the Is2faEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2faEnabled

`func (o *User) SetIs2faEnabled(v bool)`

SetIs2faEnabled sets Is2faEnabled field to given value.


### GetHasSocialAuth

`func (o *User) GetHasSocialAuth() bool`

GetHasSocialAuth returns the HasSocialAuth field if non-nil, zero value otherwise.

### GetHasSocialAuthOk

`func (o *User) GetHasSocialAuthOk() (*bool, bool)`

GetHasSocialAuthOk returns a tuple with the HasSocialAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSocialAuth

`func (o *User) SetHasSocialAuth(v bool)`

SetHasSocialAuth sets HasSocialAuth field to given value.


### GetHasSsoEnforcement

`func (o *User) GetHasSsoEnforcement() bool`

GetHasSsoEnforcement returns the HasSsoEnforcement field if non-nil, zero value otherwise.

### GetHasSsoEnforcementOk

`func (o *User) GetHasSsoEnforcementOk() (*bool, bool)`

GetHasSsoEnforcementOk returns a tuple with the HasSsoEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSsoEnforcement

`func (o *User) SetHasSsoEnforcement(v bool)`

SetHasSsoEnforcement sets HasSsoEnforcement field to given value.


### GetHasSeenProductIntroFor

`func (o *User) GetHasSeenProductIntroFor() interface{}`

GetHasSeenProductIntroFor returns the HasSeenProductIntroFor field if non-nil, zero value otherwise.

### GetHasSeenProductIntroForOk

`func (o *User) GetHasSeenProductIntroForOk() (*interface{}, bool)`

GetHasSeenProductIntroForOk returns a tuple with the HasSeenProductIntroFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSeenProductIntroFor

`func (o *User) SetHasSeenProductIntroFor(v interface{})`

SetHasSeenProductIntroFor sets HasSeenProductIntroFor field to given value.

### HasHasSeenProductIntroFor

`func (o *User) HasHasSeenProductIntroFor() bool`

HasHasSeenProductIntroFor returns a boolean if a field has been set.

### SetHasSeenProductIntroForNil

`func (o *User) SetHasSeenProductIntroForNil(b bool)`

 SetHasSeenProductIntroForNil sets the value for HasSeenProductIntroFor to be an explicit nil

### UnsetHasSeenProductIntroFor
`func (o *User) UnsetHasSeenProductIntroFor()`

UnsetHasSeenProductIntroFor ensures that no value is present for HasSeenProductIntroFor, not even an explicit nil
### GetScenePersonalisation

`func (o *User) GetScenePersonalisation() []ScenePersonalisationBasic`

GetScenePersonalisation returns the ScenePersonalisation field if non-nil, zero value otherwise.

### GetScenePersonalisationOk

`func (o *User) GetScenePersonalisationOk() (*[]ScenePersonalisationBasic, bool)`

GetScenePersonalisationOk returns a tuple with the ScenePersonalisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenePersonalisation

`func (o *User) SetScenePersonalisation(v []ScenePersonalisationBasic)`

SetScenePersonalisation sets ScenePersonalisation field to given value.


### GetThemeMode

`func (o *User) GetThemeMode() string`

GetThemeMode returns the ThemeMode field if non-nil, zero value otherwise.

### GetThemeModeOk

`func (o *User) GetThemeModeOk() (*string, bool)`

GetThemeModeOk returns a tuple with the ThemeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeMode

`func (o *User) SetThemeMode(v string)`

SetThemeMode sets ThemeMode field to given value.

### HasThemeMode

`func (o *User) HasThemeMode() bool`

HasThemeMode returns a boolean if a field has been set.

### SetThemeModeNil

`func (o *User) SetThemeModeNil(b bool)`

 SetThemeModeNil sets the value for ThemeMode to be an explicit nil

### UnsetThemeMode
`func (o *User) UnsetThemeMode()`

UnsetThemeMode ensures that no value is present for ThemeMode, not even an explicit nil
### GetHedgehogConfig

`func (o *User) GetHedgehogConfig() interface{}`

GetHedgehogConfig returns the HedgehogConfig field if non-nil, zero value otherwise.

### GetHedgehogConfigOk

`func (o *User) GetHedgehogConfigOk() (*interface{}, bool)`

GetHedgehogConfigOk returns a tuple with the HedgehogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHedgehogConfig

`func (o *User) SetHedgehogConfig(v interface{})`

SetHedgehogConfig sets HedgehogConfig field to given value.

### HasHedgehogConfig

`func (o *User) HasHedgehogConfig() bool`

HasHedgehogConfig returns a boolean if a field has been set.

### SetHedgehogConfigNil

`func (o *User) SetHedgehogConfigNil(b bool)`

 SetHedgehogConfigNil sets the value for HedgehogConfig to be an explicit nil

### UnsetHedgehogConfig
`func (o *User) UnsetHedgehogConfig()`

UnsetHedgehogConfig ensures that no value is present for HedgehogConfig, not even an explicit nil
### GetRoleAtOrganization

`func (o *User) GetRoleAtOrganization() RoleAtOrganizationEnum`

GetRoleAtOrganization returns the RoleAtOrganization field if non-nil, zero value otherwise.

### GetRoleAtOrganizationOk

`func (o *User) GetRoleAtOrganizationOk() (*RoleAtOrganizationEnum, bool)`

GetRoleAtOrganizationOk returns a tuple with the RoleAtOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAtOrganization

`func (o *User) SetRoleAtOrganization(v RoleAtOrganizationEnum)`

SetRoleAtOrganization sets RoleAtOrganization field to given value.

### HasRoleAtOrganization

`func (o *User) HasRoleAtOrganization() bool`

HasRoleAtOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


