# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**PostToSlack** | Pointer to **bool** |  | [optional] 
**SlackMessageFormat** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]ActionStepJSON**](ActionStepJSON.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**IsCalculating** | **bool** |  | [readonly] 
**LastCalculatedAt** | Pointer to **time.Time** |  | [optional] 
**TeamId** | **int32** |  | [readonly] 
**IsAction** | **bool** |  | [readonly] [default to true]
**BytecodeError** | **NullableString** |  | [readonly] 
**PinnedAt** | Pointer to **NullableTime** |  | [optional] 
**CreationContext** | **string** |  | [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**UserAccessLevel** | **NullableString** | The effective access level the user has for this object | [readonly] 

## Methods

### NewAction

`func NewAction(id int32, createdAt time.Time, createdBy UserBasic, isCalculating bool, teamId int32, isAction bool, bytecodeError NullableString, creationContext string, userAccessLevel NullableString, ) *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Action) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Action) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Action) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Action) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Action) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Action) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Action) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Action) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Action) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Action) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Action) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Action) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Action) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *Action) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Action) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Action) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Action) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPostToSlack

`func (o *Action) GetPostToSlack() bool`

GetPostToSlack returns the PostToSlack field if non-nil, zero value otherwise.

### GetPostToSlackOk

`func (o *Action) GetPostToSlackOk() (*bool, bool)`

GetPostToSlackOk returns a tuple with the PostToSlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostToSlack

`func (o *Action) SetPostToSlack(v bool)`

SetPostToSlack sets PostToSlack field to given value.

### HasPostToSlack

`func (o *Action) HasPostToSlack() bool`

HasPostToSlack returns a boolean if a field has been set.

### GetSlackMessageFormat

`func (o *Action) GetSlackMessageFormat() string`

GetSlackMessageFormat returns the SlackMessageFormat field if non-nil, zero value otherwise.

### GetSlackMessageFormatOk

`func (o *Action) GetSlackMessageFormatOk() (*string, bool)`

GetSlackMessageFormatOk returns a tuple with the SlackMessageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackMessageFormat

`func (o *Action) SetSlackMessageFormat(v string)`

SetSlackMessageFormat sets SlackMessageFormat field to given value.

### HasSlackMessageFormat

`func (o *Action) HasSlackMessageFormat() bool`

HasSlackMessageFormat returns a boolean if a field has been set.

### GetSteps

`func (o *Action) GetSteps() []ActionStepJSON`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Action) GetStepsOk() (*[]ActionStepJSON, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Action) SetSteps(v []ActionStepJSON)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Action) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Action) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Action) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Action) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Action) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Action) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Action) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetDeleted

`func (o *Action) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Action) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Action) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Action) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetIsCalculating

`func (o *Action) GetIsCalculating() bool`

GetIsCalculating returns the IsCalculating field if non-nil, zero value otherwise.

### GetIsCalculatingOk

`func (o *Action) GetIsCalculatingOk() (*bool, bool)`

GetIsCalculatingOk returns a tuple with the IsCalculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCalculating

`func (o *Action) SetIsCalculating(v bool)`

SetIsCalculating sets IsCalculating field to given value.


### GetLastCalculatedAt

`func (o *Action) GetLastCalculatedAt() time.Time`

GetLastCalculatedAt returns the LastCalculatedAt field if non-nil, zero value otherwise.

### GetLastCalculatedAtOk

`func (o *Action) GetLastCalculatedAtOk() (*time.Time, bool)`

GetLastCalculatedAtOk returns a tuple with the LastCalculatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalculatedAt

`func (o *Action) SetLastCalculatedAt(v time.Time)`

SetLastCalculatedAt sets LastCalculatedAt field to given value.

### HasLastCalculatedAt

`func (o *Action) HasLastCalculatedAt() bool`

HasLastCalculatedAt returns a boolean if a field has been set.

### GetTeamId

`func (o *Action) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Action) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Action) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetIsAction

`func (o *Action) GetIsAction() bool`

GetIsAction returns the IsAction field if non-nil, zero value otherwise.

### GetIsActionOk

`func (o *Action) GetIsActionOk() (*bool, bool)`

GetIsActionOk returns a tuple with the IsAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAction

`func (o *Action) SetIsAction(v bool)`

SetIsAction sets IsAction field to given value.


### GetBytecodeError

`func (o *Action) GetBytecodeError() string`

GetBytecodeError returns the BytecodeError field if non-nil, zero value otherwise.

### GetBytecodeErrorOk

`func (o *Action) GetBytecodeErrorOk() (*string, bool)`

GetBytecodeErrorOk returns a tuple with the BytecodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecodeError

`func (o *Action) SetBytecodeError(v string)`

SetBytecodeError sets BytecodeError field to given value.


### SetBytecodeErrorNil

`func (o *Action) SetBytecodeErrorNil(b bool)`

 SetBytecodeErrorNil sets the value for BytecodeError to be an explicit nil

### UnsetBytecodeError
`func (o *Action) UnsetBytecodeError()`

UnsetBytecodeError ensures that no value is present for BytecodeError, not even an explicit nil
### GetPinnedAt

`func (o *Action) GetPinnedAt() time.Time`

GetPinnedAt returns the PinnedAt field if non-nil, zero value otherwise.

### GetPinnedAtOk

`func (o *Action) GetPinnedAtOk() (*time.Time, bool)`

GetPinnedAtOk returns a tuple with the PinnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedAt

`func (o *Action) SetPinnedAt(v time.Time)`

SetPinnedAt sets PinnedAt field to given value.

### HasPinnedAt

`func (o *Action) HasPinnedAt() bool`

HasPinnedAt returns a boolean if a field has been set.

### SetPinnedAtNil

`func (o *Action) SetPinnedAtNil(b bool)`

 SetPinnedAtNil sets the value for PinnedAt to be an explicit nil

### UnsetPinnedAt
`func (o *Action) UnsetPinnedAt()`

UnsetPinnedAt ensures that no value is present for PinnedAt, not even an explicit nil
### GetCreationContext

`func (o *Action) GetCreationContext() string`

GetCreationContext returns the CreationContext field if non-nil, zero value otherwise.

### GetCreationContextOk

`func (o *Action) GetCreationContextOk() (*string, bool)`

GetCreationContextOk returns a tuple with the CreationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationContext

`func (o *Action) SetCreationContext(v string)`

SetCreationContext sets CreationContext field to given value.


### GetCreateInFolder

`func (o *Action) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *Action) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *Action) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *Action) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetUserAccessLevel

`func (o *Action) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *Action) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *Action) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.


### SetUserAccessLevelNil

`func (o *Action) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *Action) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


