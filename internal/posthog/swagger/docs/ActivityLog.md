# ActivityLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**User** | [**UserBasic**](UserBasic.md) |  | 
**Unread** | **bool** | is the date of this log item newer than the user&#39;s bookmark | [readonly] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**WasImpersonated** | Pointer to **NullableBool** |  | [optional] 
**IsSystem** | Pointer to **NullableBool** |  | [optional] 
**Activity** | **string** |  | 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**Scope** | **string** |  | 
**Detail** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewActivityLog

`func NewActivityLog(id string, user UserBasic, unread bool, activity string, scope string, ) *ActivityLog`

NewActivityLog instantiates a new ActivityLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogWithDefaults

`func NewActivityLogWithDefaults() *ActivityLog`

NewActivityLogWithDefaults instantiates a new ActivityLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLog) SetId(v string)`

SetId sets Id field to given value.


### GetUser

`func (o *ActivityLog) GetUser() UserBasic`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ActivityLog) GetUserOk() (*UserBasic, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ActivityLog) SetUser(v UserBasic)`

SetUser sets User field to given value.


### GetUnread

`func (o *ActivityLog) GetUnread() bool`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *ActivityLog) GetUnreadOk() (*bool, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *ActivityLog) SetUnread(v bool)`

SetUnread sets Unread field to given value.


### GetOrganizationId

`func (o *ActivityLog) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ActivityLog) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ActivityLog) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ActivityLog) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *ActivityLog) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *ActivityLog) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetWasImpersonated

`func (o *ActivityLog) GetWasImpersonated() bool`

GetWasImpersonated returns the WasImpersonated field if non-nil, zero value otherwise.

### GetWasImpersonatedOk

`func (o *ActivityLog) GetWasImpersonatedOk() (*bool, bool)`

GetWasImpersonatedOk returns a tuple with the WasImpersonated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasImpersonated

`func (o *ActivityLog) SetWasImpersonated(v bool)`

SetWasImpersonated sets WasImpersonated field to given value.

### HasWasImpersonated

`func (o *ActivityLog) HasWasImpersonated() bool`

HasWasImpersonated returns a boolean if a field has been set.

### SetWasImpersonatedNil

`func (o *ActivityLog) SetWasImpersonatedNil(b bool)`

 SetWasImpersonatedNil sets the value for WasImpersonated to be an explicit nil

### UnsetWasImpersonated
`func (o *ActivityLog) UnsetWasImpersonated()`

UnsetWasImpersonated ensures that no value is present for WasImpersonated, not even an explicit nil
### GetIsSystem

`func (o *ActivityLog) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *ActivityLog) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *ActivityLog) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *ActivityLog) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### SetIsSystemNil

`func (o *ActivityLog) SetIsSystemNil(b bool)`

 SetIsSystemNil sets the value for IsSystem to be an explicit nil

### UnsetIsSystem
`func (o *ActivityLog) UnsetIsSystem()`

UnsetIsSystem ensures that no value is present for IsSystem, not even an explicit nil
### GetActivity

`func (o *ActivityLog) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ActivityLog) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ActivityLog) SetActivity(v string)`

SetActivity sets Activity field to given value.


### GetItemId

`func (o *ActivityLog) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ActivityLog) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ActivityLog) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ActivityLog) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ActivityLog) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ActivityLog) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetScope

`func (o *ActivityLog) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ActivityLog) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ActivityLog) SetScope(v string)`

SetScope sets Scope field to given value.


### GetDetail

`func (o *ActivityLog) GetDetail() interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ActivityLog) GetDetailOk() (*interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ActivityLog) SetDetail(v interface{})`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ActivityLog) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ActivityLog) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ActivityLog) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetCreatedAt

`func (o *ActivityLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActivityLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActivityLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ActivityLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


