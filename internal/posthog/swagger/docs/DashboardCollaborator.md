# DashboardCollaborator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**DashboardId** | **int32** |  | [readonly] 
**User** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**Level** | [**DashboardRestrictionLevel**](DashboardRestrictionLevel.md) |  | 
**AddedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**UserUuid** | **string** |  | 

## Methods

### NewDashboardCollaborator

`func NewDashboardCollaborator(id string, dashboardId int32, user UserBasic, level DashboardRestrictionLevel, addedAt time.Time, updatedAt time.Time, userUuid string, ) *DashboardCollaborator`

NewDashboardCollaborator instantiates a new DashboardCollaborator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardCollaboratorWithDefaults

`func NewDashboardCollaboratorWithDefaults() *DashboardCollaborator`

NewDashboardCollaboratorWithDefaults instantiates a new DashboardCollaborator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardCollaborator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardCollaborator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardCollaborator) SetId(v string)`

SetId sets Id field to given value.


### GetDashboardId

`func (o *DashboardCollaborator) GetDashboardId() int32`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *DashboardCollaborator) GetDashboardIdOk() (*int32, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *DashboardCollaborator) SetDashboardId(v int32)`

SetDashboardId sets DashboardId field to given value.


### GetUser

`func (o *DashboardCollaborator) GetUser() UserBasic`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DashboardCollaborator) GetUserOk() (*UserBasic, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DashboardCollaborator) SetUser(v UserBasic)`

SetUser sets User field to given value.


### GetLevel

`func (o *DashboardCollaborator) GetLevel() DashboardRestrictionLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *DashboardCollaborator) GetLevelOk() (*DashboardRestrictionLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *DashboardCollaborator) SetLevel(v DashboardRestrictionLevel)`

SetLevel sets Level field to given value.


### GetAddedAt

`func (o *DashboardCollaborator) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *DashboardCollaborator) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *DashboardCollaborator) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.


### GetUpdatedAt

`func (o *DashboardCollaborator) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DashboardCollaborator) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DashboardCollaborator) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserUuid

`func (o *DashboardCollaborator) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *DashboardCollaborator) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *DashboardCollaborator) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


