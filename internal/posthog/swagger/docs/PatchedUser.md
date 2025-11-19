# PatchedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateJoined** | Pointer to **time.Time** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**DistinctId** | Pointer to **NullableString** |  | [optional] [readonly] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**PendingEmail** | Pointer to **NullableString** |  | [optional] [readonly] 
**IsEmailVerified** | Pointer to **NullableBool** |  | [optional] [readonly] 
**NotificationSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**AnonymizeData** | Pointer to **NullableBool** |  | [optional] 
**ToolbarMode** | Pointer to **NullableString** |  | [optional] 
**HasPassword** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**IsStaff** | Pointer to **bool** | Designates whether the user can log into this admin site. | [optional] 
**IsImpersonated** | Pointer to **NullableBool** |  | [optional] [readonly] 
**IsImpersonatedUntil** | Pointer to **NullableString** |  | [optional] [readonly] 
**SensitiveSessionExpiresAt** | Pointer to **NullableString** |  | [optional] [readonly] 
**Team** | Pointer to [**TeamBasic**](TeamBasic.md) |  | [optional] [readonly] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] [readonly] 
**Organizations** | Pointer to [**[]OrganizationBasic**](OrganizationBasic.md) |  | [optional] [readonly] 
**SetCurrentOrganization** | Pointer to **string** |  | [optional] 
**SetCurrentTeam** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**CurrentPassword** | Pointer to **string** |  | [optional] 
**EventsColumnConfig** | Pointer to **interface{}** |  | [optional] 
**Is2faEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**HasSocialAuth** | Pointer to **bool** |  | [optional] [readonly] 
**HasSsoEnforcement** | Pointer to **bool** |  | [optional] [readonly] 
**HasSeenProductIntroFor** | Pointer to **interface{}** |  | [optional] 
**ScenePersonalisation** | Pointer to [**[]ScenePersonalisationBasic**](ScenePersonalisationBasic.md) |  | [optional] [readonly] 
**ThemeMode** | Pointer to **NullableString** |  | [optional] 
**HedgehogConfig** | Pointer to **interface{}** |  | [optional] 
**RoleAtOrganization** | Pointer to [**RoleAtOrganizationEnum**](RoleAtOrganizationEnum.md) |  | [optional] 

## Methods

### NewPatchedUser

`func NewPatchedUser() *PatchedUser`

NewPatchedUser instantiates a new PatchedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUserWithDefaults

`func NewPatchedUserWithDefaults() *PatchedUser`

NewPatchedUserWithDefaults instantiates a new PatchedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateJoined

`func (o *PatchedUser) GetDateJoined() time.Time`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *PatchedUser) GetDateJoinedOk() (*time.Time, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *PatchedUser) SetDateJoined(v time.Time)`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *PatchedUser) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### GetUuid

`func (o *PatchedUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PatchedUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PatchedUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PatchedUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDistinctId

`func (o *PatchedUser) GetDistinctId() string`

GetDistinctId returns the DistinctId field if non-nil, zero value otherwise.

### GetDistinctIdOk

`func (o *PatchedUser) GetDistinctIdOk() (*string, bool)`

GetDistinctIdOk returns a tuple with the DistinctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctId

`func (o *PatchedUser) SetDistinctId(v string)`

SetDistinctId sets DistinctId field to given value.

### HasDistinctId

`func (o *PatchedUser) HasDistinctId() bool`

HasDistinctId returns a boolean if a field has been set.

### SetDistinctIdNil

`func (o *PatchedUser) SetDistinctIdNil(b bool)`

 SetDistinctIdNil sets the value for DistinctId to be an explicit nil

### UnsetDistinctId
`func (o *PatchedUser) UnsetDistinctId()`

UnsetDistinctId ensures that no value is present for DistinctId, not even an explicit nil
### GetFirstName

`func (o *PatchedUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPendingEmail

`func (o *PatchedUser) GetPendingEmail() string`

GetPendingEmail returns the PendingEmail field if non-nil, zero value otherwise.

### GetPendingEmailOk

`func (o *PatchedUser) GetPendingEmailOk() (*string, bool)`

GetPendingEmailOk returns a tuple with the PendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingEmail

`func (o *PatchedUser) SetPendingEmail(v string)`

SetPendingEmail sets PendingEmail field to given value.

### HasPendingEmail

`func (o *PatchedUser) HasPendingEmail() bool`

HasPendingEmail returns a boolean if a field has been set.

### SetPendingEmailNil

`func (o *PatchedUser) SetPendingEmailNil(b bool)`

 SetPendingEmailNil sets the value for PendingEmail to be an explicit nil

### UnsetPendingEmail
`func (o *PatchedUser) UnsetPendingEmail()`

UnsetPendingEmail ensures that no value is present for PendingEmail, not even an explicit nil
### GetIsEmailVerified

`func (o *PatchedUser) GetIsEmailVerified() bool`

GetIsEmailVerified returns the IsEmailVerified field if non-nil, zero value otherwise.

### GetIsEmailVerifiedOk

`func (o *PatchedUser) GetIsEmailVerifiedOk() (*bool, bool)`

GetIsEmailVerifiedOk returns a tuple with the IsEmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailVerified

`func (o *PatchedUser) SetIsEmailVerified(v bool)`

SetIsEmailVerified sets IsEmailVerified field to given value.

### HasIsEmailVerified

`func (o *PatchedUser) HasIsEmailVerified() bool`

HasIsEmailVerified returns a boolean if a field has been set.

### SetIsEmailVerifiedNil

`func (o *PatchedUser) SetIsEmailVerifiedNil(b bool)`

 SetIsEmailVerifiedNil sets the value for IsEmailVerified to be an explicit nil

### UnsetIsEmailVerified
`func (o *PatchedUser) UnsetIsEmailVerified()`

UnsetIsEmailVerified ensures that no value is present for IsEmailVerified, not even an explicit nil
### GetNotificationSettings

`func (o *PatchedUser) GetNotificationSettings() map[string]interface{}`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *PatchedUser) GetNotificationSettingsOk() (*map[string]interface{}, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSettings

`func (o *PatchedUser) SetNotificationSettings(v map[string]interface{})`

SetNotificationSettings sets NotificationSettings field to given value.

### HasNotificationSettings

`func (o *PatchedUser) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### GetAnonymizeData

`func (o *PatchedUser) GetAnonymizeData() bool`

GetAnonymizeData returns the AnonymizeData field if non-nil, zero value otherwise.

### GetAnonymizeDataOk

`func (o *PatchedUser) GetAnonymizeDataOk() (*bool, bool)`

GetAnonymizeDataOk returns a tuple with the AnonymizeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymizeData

`func (o *PatchedUser) SetAnonymizeData(v bool)`

SetAnonymizeData sets AnonymizeData field to given value.

### HasAnonymizeData

`func (o *PatchedUser) HasAnonymizeData() bool`

HasAnonymizeData returns a boolean if a field has been set.

### SetAnonymizeDataNil

`func (o *PatchedUser) SetAnonymizeDataNil(b bool)`

 SetAnonymizeDataNil sets the value for AnonymizeData to be an explicit nil

### UnsetAnonymizeData
`func (o *PatchedUser) UnsetAnonymizeData()`

UnsetAnonymizeData ensures that no value is present for AnonymizeData, not even an explicit nil
### GetToolbarMode

`func (o *PatchedUser) GetToolbarMode() string`

GetToolbarMode returns the ToolbarMode field if non-nil, zero value otherwise.

### GetToolbarModeOk

`func (o *PatchedUser) GetToolbarModeOk() (*string, bool)`

GetToolbarModeOk returns a tuple with the ToolbarMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarMode

`func (o *PatchedUser) SetToolbarMode(v string)`

SetToolbarMode sets ToolbarMode field to given value.

### HasToolbarMode

`func (o *PatchedUser) HasToolbarMode() bool`

HasToolbarMode returns a boolean if a field has been set.

### SetToolbarModeNil

`func (o *PatchedUser) SetToolbarModeNil(b bool)`

 SetToolbarModeNil sets the value for ToolbarMode to be an explicit nil

### UnsetToolbarMode
`func (o *PatchedUser) UnsetToolbarMode()`

UnsetToolbarMode ensures that no value is present for ToolbarMode, not even an explicit nil
### GetHasPassword

`func (o *PatchedUser) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *PatchedUser) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *PatchedUser) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *PatchedUser) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetId

`func (o *PatchedUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsStaff

`func (o *PatchedUser) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *PatchedUser) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *PatchedUser) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *PatchedUser) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsImpersonated

`func (o *PatchedUser) GetIsImpersonated() bool`

GetIsImpersonated returns the IsImpersonated field if non-nil, zero value otherwise.

### GetIsImpersonatedOk

`func (o *PatchedUser) GetIsImpersonatedOk() (*bool, bool)`

GetIsImpersonatedOk returns a tuple with the IsImpersonated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImpersonated

`func (o *PatchedUser) SetIsImpersonated(v bool)`

SetIsImpersonated sets IsImpersonated field to given value.

### HasIsImpersonated

`func (o *PatchedUser) HasIsImpersonated() bool`

HasIsImpersonated returns a boolean if a field has been set.

### SetIsImpersonatedNil

`func (o *PatchedUser) SetIsImpersonatedNil(b bool)`

 SetIsImpersonatedNil sets the value for IsImpersonated to be an explicit nil

### UnsetIsImpersonated
`func (o *PatchedUser) UnsetIsImpersonated()`

UnsetIsImpersonated ensures that no value is present for IsImpersonated, not even an explicit nil
### GetIsImpersonatedUntil

`func (o *PatchedUser) GetIsImpersonatedUntil() string`

GetIsImpersonatedUntil returns the IsImpersonatedUntil field if non-nil, zero value otherwise.

### GetIsImpersonatedUntilOk

`func (o *PatchedUser) GetIsImpersonatedUntilOk() (*string, bool)`

GetIsImpersonatedUntilOk returns a tuple with the IsImpersonatedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImpersonatedUntil

`func (o *PatchedUser) SetIsImpersonatedUntil(v string)`

SetIsImpersonatedUntil sets IsImpersonatedUntil field to given value.

### HasIsImpersonatedUntil

`func (o *PatchedUser) HasIsImpersonatedUntil() bool`

HasIsImpersonatedUntil returns a boolean if a field has been set.

### SetIsImpersonatedUntilNil

`func (o *PatchedUser) SetIsImpersonatedUntilNil(b bool)`

 SetIsImpersonatedUntilNil sets the value for IsImpersonatedUntil to be an explicit nil

### UnsetIsImpersonatedUntil
`func (o *PatchedUser) UnsetIsImpersonatedUntil()`

UnsetIsImpersonatedUntil ensures that no value is present for IsImpersonatedUntil, not even an explicit nil
### GetSensitiveSessionExpiresAt

`func (o *PatchedUser) GetSensitiveSessionExpiresAt() string`

GetSensitiveSessionExpiresAt returns the SensitiveSessionExpiresAt field if non-nil, zero value otherwise.

### GetSensitiveSessionExpiresAtOk

`func (o *PatchedUser) GetSensitiveSessionExpiresAtOk() (*string, bool)`

GetSensitiveSessionExpiresAtOk returns a tuple with the SensitiveSessionExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveSessionExpiresAt

`func (o *PatchedUser) SetSensitiveSessionExpiresAt(v string)`

SetSensitiveSessionExpiresAt sets SensitiveSessionExpiresAt field to given value.

### HasSensitiveSessionExpiresAt

`func (o *PatchedUser) HasSensitiveSessionExpiresAt() bool`

HasSensitiveSessionExpiresAt returns a boolean if a field has been set.

### SetSensitiveSessionExpiresAtNil

`func (o *PatchedUser) SetSensitiveSessionExpiresAtNil(b bool)`

 SetSensitiveSessionExpiresAtNil sets the value for SensitiveSessionExpiresAt to be an explicit nil

### UnsetSensitiveSessionExpiresAt
`func (o *PatchedUser) UnsetSensitiveSessionExpiresAt()`

UnsetSensitiveSessionExpiresAt ensures that no value is present for SensitiveSessionExpiresAt, not even an explicit nil
### GetTeam

`func (o *PatchedUser) GetTeam() TeamBasic`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PatchedUser) GetTeamOk() (*TeamBasic, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PatchedUser) SetTeam(v TeamBasic)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PatchedUser) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetOrganization

`func (o *PatchedUser) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PatchedUser) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PatchedUser) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PatchedUser) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizations

`func (o *PatchedUser) GetOrganizations() []OrganizationBasic`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *PatchedUser) GetOrganizationsOk() (*[]OrganizationBasic, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *PatchedUser) SetOrganizations(v []OrganizationBasic)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *PatchedUser) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetSetCurrentOrganization

`func (o *PatchedUser) GetSetCurrentOrganization() string`

GetSetCurrentOrganization returns the SetCurrentOrganization field if non-nil, zero value otherwise.

### GetSetCurrentOrganizationOk

`func (o *PatchedUser) GetSetCurrentOrganizationOk() (*string, bool)`

GetSetCurrentOrganizationOk returns a tuple with the SetCurrentOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCurrentOrganization

`func (o *PatchedUser) SetSetCurrentOrganization(v string)`

SetSetCurrentOrganization sets SetCurrentOrganization field to given value.

### HasSetCurrentOrganization

`func (o *PatchedUser) HasSetCurrentOrganization() bool`

HasSetCurrentOrganization returns a boolean if a field has been set.

### GetSetCurrentTeam

`func (o *PatchedUser) GetSetCurrentTeam() string`

GetSetCurrentTeam returns the SetCurrentTeam field if non-nil, zero value otherwise.

### GetSetCurrentTeamOk

`func (o *PatchedUser) GetSetCurrentTeamOk() (*string, bool)`

GetSetCurrentTeamOk returns a tuple with the SetCurrentTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCurrentTeam

`func (o *PatchedUser) SetSetCurrentTeam(v string)`

SetSetCurrentTeam sets SetCurrentTeam field to given value.

### HasSetCurrentTeam

`func (o *PatchedUser) HasSetCurrentTeam() bool`

HasSetCurrentTeam returns a boolean if a field has been set.

### GetPassword

`func (o *PatchedUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCurrentPassword

`func (o *PatchedUser) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *PatchedUser) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *PatchedUser) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *PatchedUser) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetEventsColumnConfig

`func (o *PatchedUser) GetEventsColumnConfig() interface{}`

GetEventsColumnConfig returns the EventsColumnConfig field if non-nil, zero value otherwise.

### GetEventsColumnConfigOk

`func (o *PatchedUser) GetEventsColumnConfigOk() (*interface{}, bool)`

GetEventsColumnConfigOk returns a tuple with the EventsColumnConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsColumnConfig

`func (o *PatchedUser) SetEventsColumnConfig(v interface{})`

SetEventsColumnConfig sets EventsColumnConfig field to given value.

### HasEventsColumnConfig

`func (o *PatchedUser) HasEventsColumnConfig() bool`

HasEventsColumnConfig returns a boolean if a field has been set.

### SetEventsColumnConfigNil

`func (o *PatchedUser) SetEventsColumnConfigNil(b bool)`

 SetEventsColumnConfigNil sets the value for EventsColumnConfig to be an explicit nil

### UnsetEventsColumnConfig
`func (o *PatchedUser) UnsetEventsColumnConfig()`

UnsetEventsColumnConfig ensures that no value is present for EventsColumnConfig, not even an explicit nil
### GetIs2faEnabled

`func (o *PatchedUser) GetIs2faEnabled() bool`

GetIs2faEnabled returns the Is2faEnabled field if non-nil, zero value otherwise.

### GetIs2faEnabledOk

`func (o *PatchedUser) GetIs2faEnabledOk() (*bool, bool)`

GetIs2faEnabledOk returns a tuple with the Is2faEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2faEnabled

`func (o *PatchedUser) SetIs2faEnabled(v bool)`

SetIs2faEnabled sets Is2faEnabled field to given value.

### HasIs2faEnabled

`func (o *PatchedUser) HasIs2faEnabled() bool`

HasIs2faEnabled returns a boolean if a field has been set.

### GetHasSocialAuth

`func (o *PatchedUser) GetHasSocialAuth() bool`

GetHasSocialAuth returns the HasSocialAuth field if non-nil, zero value otherwise.

### GetHasSocialAuthOk

`func (o *PatchedUser) GetHasSocialAuthOk() (*bool, bool)`

GetHasSocialAuthOk returns a tuple with the HasSocialAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSocialAuth

`func (o *PatchedUser) SetHasSocialAuth(v bool)`

SetHasSocialAuth sets HasSocialAuth field to given value.

### HasHasSocialAuth

`func (o *PatchedUser) HasHasSocialAuth() bool`

HasHasSocialAuth returns a boolean if a field has been set.

### GetHasSsoEnforcement

`func (o *PatchedUser) GetHasSsoEnforcement() bool`

GetHasSsoEnforcement returns the HasSsoEnforcement field if non-nil, zero value otherwise.

### GetHasSsoEnforcementOk

`func (o *PatchedUser) GetHasSsoEnforcementOk() (*bool, bool)`

GetHasSsoEnforcementOk returns a tuple with the HasSsoEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSsoEnforcement

`func (o *PatchedUser) SetHasSsoEnforcement(v bool)`

SetHasSsoEnforcement sets HasSsoEnforcement field to given value.

### HasHasSsoEnforcement

`func (o *PatchedUser) HasHasSsoEnforcement() bool`

HasHasSsoEnforcement returns a boolean if a field has been set.

### GetHasSeenProductIntroFor

`func (o *PatchedUser) GetHasSeenProductIntroFor() interface{}`

GetHasSeenProductIntroFor returns the HasSeenProductIntroFor field if non-nil, zero value otherwise.

### GetHasSeenProductIntroForOk

`func (o *PatchedUser) GetHasSeenProductIntroForOk() (*interface{}, bool)`

GetHasSeenProductIntroForOk returns a tuple with the HasSeenProductIntroFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSeenProductIntroFor

`func (o *PatchedUser) SetHasSeenProductIntroFor(v interface{})`

SetHasSeenProductIntroFor sets HasSeenProductIntroFor field to given value.

### HasHasSeenProductIntroFor

`func (o *PatchedUser) HasHasSeenProductIntroFor() bool`

HasHasSeenProductIntroFor returns a boolean if a field has been set.

### SetHasSeenProductIntroForNil

`func (o *PatchedUser) SetHasSeenProductIntroForNil(b bool)`

 SetHasSeenProductIntroForNil sets the value for HasSeenProductIntroFor to be an explicit nil

### UnsetHasSeenProductIntroFor
`func (o *PatchedUser) UnsetHasSeenProductIntroFor()`

UnsetHasSeenProductIntroFor ensures that no value is present for HasSeenProductIntroFor, not even an explicit nil
### GetScenePersonalisation

`func (o *PatchedUser) GetScenePersonalisation() []ScenePersonalisationBasic`

GetScenePersonalisation returns the ScenePersonalisation field if non-nil, zero value otherwise.

### GetScenePersonalisationOk

`func (o *PatchedUser) GetScenePersonalisationOk() (*[]ScenePersonalisationBasic, bool)`

GetScenePersonalisationOk returns a tuple with the ScenePersonalisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenePersonalisation

`func (o *PatchedUser) SetScenePersonalisation(v []ScenePersonalisationBasic)`

SetScenePersonalisation sets ScenePersonalisation field to given value.

### HasScenePersonalisation

`func (o *PatchedUser) HasScenePersonalisation() bool`

HasScenePersonalisation returns a boolean if a field has been set.

### GetThemeMode

`func (o *PatchedUser) GetThemeMode() string`

GetThemeMode returns the ThemeMode field if non-nil, zero value otherwise.

### GetThemeModeOk

`func (o *PatchedUser) GetThemeModeOk() (*string, bool)`

GetThemeModeOk returns a tuple with the ThemeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeMode

`func (o *PatchedUser) SetThemeMode(v string)`

SetThemeMode sets ThemeMode field to given value.

### HasThemeMode

`func (o *PatchedUser) HasThemeMode() bool`

HasThemeMode returns a boolean if a field has been set.

### SetThemeModeNil

`func (o *PatchedUser) SetThemeModeNil(b bool)`

 SetThemeModeNil sets the value for ThemeMode to be an explicit nil

### UnsetThemeMode
`func (o *PatchedUser) UnsetThemeMode()`

UnsetThemeMode ensures that no value is present for ThemeMode, not even an explicit nil
### GetHedgehogConfig

`func (o *PatchedUser) GetHedgehogConfig() interface{}`

GetHedgehogConfig returns the HedgehogConfig field if non-nil, zero value otherwise.

### GetHedgehogConfigOk

`func (o *PatchedUser) GetHedgehogConfigOk() (*interface{}, bool)`

GetHedgehogConfigOk returns a tuple with the HedgehogConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHedgehogConfig

`func (o *PatchedUser) SetHedgehogConfig(v interface{})`

SetHedgehogConfig sets HedgehogConfig field to given value.

### HasHedgehogConfig

`func (o *PatchedUser) HasHedgehogConfig() bool`

HasHedgehogConfig returns a boolean if a field has been set.

### SetHedgehogConfigNil

`func (o *PatchedUser) SetHedgehogConfigNil(b bool)`

 SetHedgehogConfigNil sets the value for HedgehogConfig to be an explicit nil

### UnsetHedgehogConfig
`func (o *PatchedUser) UnsetHedgehogConfig()`

UnsetHedgehogConfig ensures that no value is present for HedgehogConfig, not even an explicit nil
### GetRoleAtOrganization

`func (o *PatchedUser) GetRoleAtOrganization() RoleAtOrganizationEnum`

GetRoleAtOrganization returns the RoleAtOrganization field if non-nil, zero value otherwise.

### GetRoleAtOrganizationOk

`func (o *PatchedUser) GetRoleAtOrganizationOk() (*RoleAtOrganizationEnum, bool)`

GetRoleAtOrganizationOk returns a tuple with the RoleAtOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAtOrganization

`func (o *PatchedUser) SetRoleAtOrganization(v RoleAtOrganizationEnum)`

SetRoleAtOrganization sets RoleAtOrganization field to given value.

### HasRoleAtOrganization

`func (o *PatchedUser) HasRoleAtOrganization() bool`

HasRoleAtOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


