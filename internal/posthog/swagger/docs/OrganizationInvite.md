# OrganizationInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**TargetEmail** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**EmailingAttemptMade** | **bool** |  | [readonly] 
**Level** | Pointer to [**OrganizationMembershipLevel**](OrganizationMembershipLevel.md) |  | [optional] 
**IsExpired** | **bool** | Check if invite is older than INVITE_DAYS_VALIDITY days. | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Message** | Pointer to **NullableString** |  | [optional] 
**PrivateProjectAccess** | Pointer to **interface{}** | List of team IDs and corresponding access levels to private projects. | [optional] 
**SendEmail** | Pointer to **bool** |  | [optional] [default to true]
**CombinePendingInvites** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewOrganizationInvite

`func NewOrganizationInvite(id string, targetEmail string, emailingAttemptMade bool, isExpired bool, createdBy UserBasic, createdAt time.Time, updatedAt time.Time, ) *OrganizationInvite`

NewOrganizationInvite instantiates a new OrganizationInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInviteWithDefaults

`func NewOrganizationInviteWithDefaults() *OrganizationInvite`

NewOrganizationInviteWithDefaults instantiates a new OrganizationInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationInvite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationInvite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationInvite) SetId(v string)`

SetId sets Id field to given value.


### GetTargetEmail

`func (o *OrganizationInvite) GetTargetEmail() string`

GetTargetEmail returns the TargetEmail field if non-nil, zero value otherwise.

### GetTargetEmailOk

`func (o *OrganizationInvite) GetTargetEmailOk() (*string, bool)`

GetTargetEmailOk returns a tuple with the TargetEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEmail

`func (o *OrganizationInvite) SetTargetEmail(v string)`

SetTargetEmail sets TargetEmail field to given value.


### GetFirstName

`func (o *OrganizationInvite) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrganizationInvite) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrganizationInvite) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OrganizationInvite) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetEmailingAttemptMade

`func (o *OrganizationInvite) GetEmailingAttemptMade() bool`

GetEmailingAttemptMade returns the EmailingAttemptMade field if non-nil, zero value otherwise.

### GetEmailingAttemptMadeOk

`func (o *OrganizationInvite) GetEmailingAttemptMadeOk() (*bool, bool)`

GetEmailingAttemptMadeOk returns a tuple with the EmailingAttemptMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailingAttemptMade

`func (o *OrganizationInvite) SetEmailingAttemptMade(v bool)`

SetEmailingAttemptMade sets EmailingAttemptMade field to given value.


### GetLevel

`func (o *OrganizationInvite) GetLevel() OrganizationMembershipLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *OrganizationInvite) GetLevelOk() (*OrganizationMembershipLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *OrganizationInvite) SetLevel(v OrganizationMembershipLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *OrganizationInvite) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetIsExpired

`func (o *OrganizationInvite) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *OrganizationInvite) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *OrganizationInvite) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.


### GetCreatedBy

`func (o *OrganizationInvite) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganizationInvite) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganizationInvite) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *OrganizationInvite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationInvite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationInvite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationInvite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationInvite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationInvite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMessage

`func (o *OrganizationInvite) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrganizationInvite) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrganizationInvite) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrganizationInvite) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *OrganizationInvite) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *OrganizationInvite) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPrivateProjectAccess

`func (o *OrganizationInvite) GetPrivateProjectAccess() interface{}`

GetPrivateProjectAccess returns the PrivateProjectAccess field if non-nil, zero value otherwise.

### GetPrivateProjectAccessOk

`func (o *OrganizationInvite) GetPrivateProjectAccessOk() (*interface{}, bool)`

GetPrivateProjectAccessOk returns a tuple with the PrivateProjectAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateProjectAccess

`func (o *OrganizationInvite) SetPrivateProjectAccess(v interface{})`

SetPrivateProjectAccess sets PrivateProjectAccess field to given value.

### HasPrivateProjectAccess

`func (o *OrganizationInvite) HasPrivateProjectAccess() bool`

HasPrivateProjectAccess returns a boolean if a field has been set.

### SetPrivateProjectAccessNil

`func (o *OrganizationInvite) SetPrivateProjectAccessNil(b bool)`

 SetPrivateProjectAccessNil sets the value for PrivateProjectAccess to be an explicit nil

### UnsetPrivateProjectAccess
`func (o *OrganizationInvite) UnsetPrivateProjectAccess()`

UnsetPrivateProjectAccess ensures that no value is present for PrivateProjectAccess, not even an explicit nil
### GetSendEmail

`func (o *OrganizationInvite) GetSendEmail() bool`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *OrganizationInvite) GetSendEmailOk() (*bool, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *OrganizationInvite) SetSendEmail(v bool)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *OrganizationInvite) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetCombinePendingInvites

`func (o *OrganizationInvite) GetCombinePendingInvites() bool`

GetCombinePendingInvites returns the CombinePendingInvites field if non-nil, zero value otherwise.

### GetCombinePendingInvitesOk

`func (o *OrganizationInvite) GetCombinePendingInvitesOk() (*bool, bool)`

GetCombinePendingInvitesOk returns a tuple with the CombinePendingInvites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinePendingInvites

`func (o *OrganizationInvite) SetCombinePendingInvites(v bool)`

SetCombinePendingInvites sets CombinePendingInvites field to given value.

### HasCombinePendingInvites

`func (o *OrganizationInvite) HasCombinePendingInvites() bool`

HasCombinePendingInvites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


