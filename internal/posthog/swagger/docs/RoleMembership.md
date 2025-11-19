# RoleMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**RoleId** | **string** |  | [readonly] 
**OrganizationMember** | [**OrganizationMember**](OrganizationMember.md) |  | [readonly] 
**User** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**JoinedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**UserUuid** | **string** |  | 

## Methods

### NewRoleMembership

`func NewRoleMembership(id string, roleId string, organizationMember OrganizationMember, user UserBasic, joinedAt time.Time, updatedAt time.Time, userUuid string, ) *RoleMembership`

NewRoleMembership instantiates a new RoleMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleMembershipWithDefaults

`func NewRoleMembershipWithDefaults() *RoleMembership`

NewRoleMembershipWithDefaults instantiates a new RoleMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleMembership) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleMembership) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleMembership) SetId(v string)`

SetId sets Id field to given value.


### GetRoleId

`func (o *RoleMembership) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RoleMembership) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RoleMembership) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetOrganizationMember

`func (o *RoleMembership) GetOrganizationMember() OrganizationMember`

GetOrganizationMember returns the OrganizationMember field if non-nil, zero value otherwise.

### GetOrganizationMemberOk

`func (o *RoleMembership) GetOrganizationMemberOk() (*OrganizationMember, bool)`

GetOrganizationMemberOk returns a tuple with the OrganizationMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationMember

`func (o *RoleMembership) SetOrganizationMember(v OrganizationMember)`

SetOrganizationMember sets OrganizationMember field to given value.


### GetUser

`func (o *RoleMembership) GetUser() UserBasic`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RoleMembership) GetUserOk() (*UserBasic, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RoleMembership) SetUser(v UserBasic)`

SetUser sets User field to given value.


### GetJoinedAt

`func (o *RoleMembership) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *RoleMembership) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *RoleMembership) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.


### GetUpdatedAt

`func (o *RoleMembership) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoleMembership) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoleMembership) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserUuid

`func (o *RoleMembership) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *RoleMembership) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *RoleMembership) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


