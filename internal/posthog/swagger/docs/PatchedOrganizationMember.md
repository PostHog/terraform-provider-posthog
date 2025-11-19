# PatchedOrganizationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**User** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**Level** | Pointer to [**OrganizationMembershipLevel**](OrganizationMembershipLevel.md) |  | [optional] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Is2faEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**HasSocialAuth** | Pointer to **bool** |  | [optional] [readonly] 
**LastLogin** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedOrganizationMember

`func NewPatchedOrganizationMember() *PatchedOrganizationMember`

NewPatchedOrganizationMember instantiates a new PatchedOrganizationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedOrganizationMemberWithDefaults

`func NewPatchedOrganizationMemberWithDefaults() *PatchedOrganizationMember`

NewPatchedOrganizationMemberWithDefaults instantiates a new PatchedOrganizationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedOrganizationMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedOrganizationMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedOrganizationMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedOrganizationMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUser

`func (o *PatchedOrganizationMember) GetUser() UserBasic`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedOrganizationMember) GetUserOk() (*UserBasic, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedOrganizationMember) SetUser(v UserBasic)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedOrganizationMember) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLevel

`func (o *PatchedOrganizationMember) GetLevel() OrganizationMembershipLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PatchedOrganizationMember) GetLevelOk() (*OrganizationMembershipLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PatchedOrganizationMember) SetLevel(v OrganizationMembershipLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *PatchedOrganizationMember) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetJoinedAt

`func (o *PatchedOrganizationMember) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *PatchedOrganizationMember) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *PatchedOrganizationMember) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *PatchedOrganizationMember) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedOrganizationMember) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedOrganizationMember) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedOrganizationMember) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedOrganizationMember) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIs2faEnabled

`func (o *PatchedOrganizationMember) GetIs2faEnabled() bool`

GetIs2faEnabled returns the Is2faEnabled field if non-nil, zero value otherwise.

### GetIs2faEnabledOk

`func (o *PatchedOrganizationMember) GetIs2faEnabledOk() (*bool, bool)`

GetIs2faEnabledOk returns a tuple with the Is2faEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2faEnabled

`func (o *PatchedOrganizationMember) SetIs2faEnabled(v bool)`

SetIs2faEnabled sets Is2faEnabled field to given value.

### HasIs2faEnabled

`func (o *PatchedOrganizationMember) HasIs2faEnabled() bool`

HasIs2faEnabled returns a boolean if a field has been set.

### GetHasSocialAuth

`func (o *PatchedOrganizationMember) GetHasSocialAuth() bool`

GetHasSocialAuth returns the HasSocialAuth field if non-nil, zero value otherwise.

### GetHasSocialAuthOk

`func (o *PatchedOrganizationMember) GetHasSocialAuthOk() (*bool, bool)`

GetHasSocialAuthOk returns a tuple with the HasSocialAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSocialAuth

`func (o *PatchedOrganizationMember) SetHasSocialAuth(v bool)`

SetHasSocialAuth sets HasSocialAuth field to given value.

### HasHasSocialAuth

`func (o *PatchedOrganizationMember) HasHasSocialAuth() bool`

HasHasSocialAuth returns a boolean if a field has been set.

### GetLastLogin

`func (o *PatchedOrganizationMember) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PatchedOrganizationMember) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PatchedOrganizationMember) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PatchedOrganizationMember) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


