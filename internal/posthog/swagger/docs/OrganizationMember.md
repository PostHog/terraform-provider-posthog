# OrganizationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**User** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**Level** | Pointer to [**OrganizationMembershipLevel**](OrganizationMembershipLevel.md) |  | [optional] 
**JoinedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Is2faEnabled** | **bool** |  | [readonly] 
**HasSocialAuth** | **bool** |  | [readonly] 
**LastLogin** | **time.Time** |  | [readonly] 

## Methods

### NewOrganizationMember

`func NewOrganizationMember(id string, user UserBasic, joinedAt time.Time, updatedAt time.Time, is2faEnabled bool, hasSocialAuth bool, lastLogin time.Time, ) *OrganizationMember`

NewOrganizationMember instantiates a new OrganizationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberWithDefaults

`func NewOrganizationMemberWithDefaults() *OrganizationMember`

NewOrganizationMemberWithDefaults instantiates a new OrganizationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationMember) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *OrganizationMember) GetUser() UserBasic`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationMember) GetUserOk() (*UserBasic, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationMember) SetUser(v UserBasic)`

SetUser sets User field to given value.


### GetLevel

`func (o *OrganizationMember) GetLevel() OrganizationMembershipLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *OrganizationMember) GetLevelOk() (*OrganizationMembershipLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *OrganizationMember) SetLevel(v OrganizationMembershipLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *OrganizationMember) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetJoinedAt

`func (o *OrganizationMember) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrganizationMember) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrganizationMember) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationMember) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationMember) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationMember) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIs2faEnabled

`func (o *OrganizationMember) GetIs2faEnabled() bool`

GetIs2faEnabled returns the Is2faEnabled field if non-nil, zero value otherwise.

### GetIs2faEnabledOk

`func (o *OrganizationMember) GetIs2faEnabledOk() (*bool, bool)`

GetIs2faEnabledOk returns a tuple with the Is2faEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs2faEnabled

`func (o *OrganizationMember) SetIs2faEnabled(v bool)`

SetIs2faEnabled sets Is2faEnabled field to given value.


### GetHasSocialAuth

`func (o *OrganizationMember) GetHasSocialAuth() bool`

GetHasSocialAuth returns the HasSocialAuth field if non-nil, zero value otherwise.

### GetHasSocialAuthOk

`func (o *OrganizationMember) GetHasSocialAuthOk() (*bool, bool)`

GetHasSocialAuthOk returns a tuple with the HasSocialAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSocialAuth

`func (o *OrganizationMember) SetHasSocialAuth(v bool)`

SetHasSocialAuth sets HasSocialAuth field to given value.


### GetLastLogin

`func (o *OrganizationMember) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *OrganizationMember) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *OrganizationMember) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


