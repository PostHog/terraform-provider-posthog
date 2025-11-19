# PatchedAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**PostToSlack** | Pointer to **bool** |  | [optional] 
**SlackMessageFormat** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]ActionStepJSON**](ActionStepJSON.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 
**IsCalculating** | Pointer to **bool** |  | [optional] [readonly] 
**LastCalculatedAt** | Pointer to **time.Time** |  | [optional] 
**TeamId** | Pointer to **int32** |  | [optional] [readonly] 
**IsAction** | Pointer to **bool** |  | [optional] [readonly] [default to true]
**BytecodeError** | Pointer to **NullableString** |  | [optional] [readonly] 
**PinnedAt** | Pointer to **NullableTime** |  | [optional] 
**CreationContext** | Pointer to **string** |  | [optional] [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**UserAccessLevel** | Pointer to **NullableString** | The effective access level the user has for this object | [optional] [readonly] 

## Methods

### NewPatchedAction

`func NewPatchedAction() *PatchedAction`

NewPatchedAction instantiates a new PatchedAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedActionWithDefaults

`func NewPatchedActionWithDefaults() *PatchedAction`

NewPatchedActionWithDefaults instantiates a new PatchedAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedAction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedAction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedAction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAction) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedAction) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedAction) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PatchedAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedAction) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedAction) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedAction) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedAction) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPostToSlack

`func (o *PatchedAction) GetPostToSlack() bool`

GetPostToSlack returns the PostToSlack field if non-nil, zero value otherwise.

### GetPostToSlackOk

`func (o *PatchedAction) GetPostToSlackOk() (*bool, bool)`

GetPostToSlackOk returns a tuple with the PostToSlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostToSlack

`func (o *PatchedAction) SetPostToSlack(v bool)`

SetPostToSlack sets PostToSlack field to given value.

### HasPostToSlack

`func (o *PatchedAction) HasPostToSlack() bool`

HasPostToSlack returns a boolean if a field has been set.

### GetSlackMessageFormat

`func (o *PatchedAction) GetSlackMessageFormat() string`

GetSlackMessageFormat returns the SlackMessageFormat field if non-nil, zero value otherwise.

### GetSlackMessageFormatOk

`func (o *PatchedAction) GetSlackMessageFormatOk() (*string, bool)`

GetSlackMessageFormatOk returns a tuple with the SlackMessageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackMessageFormat

`func (o *PatchedAction) SetSlackMessageFormat(v string)`

SetSlackMessageFormat sets SlackMessageFormat field to given value.

### HasSlackMessageFormat

`func (o *PatchedAction) HasSlackMessageFormat() bool`

HasSlackMessageFormat returns a boolean if a field has been set.

### GetSteps

`func (o *PatchedAction) GetSteps() []ActionStepJSON`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *PatchedAction) GetStepsOk() (*[]ActionStepJSON, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *PatchedAction) SetSteps(v []ActionStepJSON)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *PatchedAction) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedAction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedAction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedAction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedAction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedAction) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedAction) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedAction) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedAction) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedAction) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedAction) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedAction) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedAction) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetIsCalculating

`func (o *PatchedAction) GetIsCalculating() bool`

GetIsCalculating returns the IsCalculating field if non-nil, zero value otherwise.

### GetIsCalculatingOk

`func (o *PatchedAction) GetIsCalculatingOk() (*bool, bool)`

GetIsCalculatingOk returns a tuple with the IsCalculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCalculating

`func (o *PatchedAction) SetIsCalculating(v bool)`

SetIsCalculating sets IsCalculating field to given value.

### HasIsCalculating

`func (o *PatchedAction) HasIsCalculating() bool`

HasIsCalculating returns a boolean if a field has been set.

### GetLastCalculatedAt

`func (o *PatchedAction) GetLastCalculatedAt() time.Time`

GetLastCalculatedAt returns the LastCalculatedAt field if non-nil, zero value otherwise.

### GetLastCalculatedAtOk

`func (o *PatchedAction) GetLastCalculatedAtOk() (*time.Time, bool)`

GetLastCalculatedAtOk returns a tuple with the LastCalculatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalculatedAt

`func (o *PatchedAction) SetLastCalculatedAt(v time.Time)`

SetLastCalculatedAt sets LastCalculatedAt field to given value.

### HasLastCalculatedAt

`func (o *PatchedAction) HasLastCalculatedAt() bool`

HasLastCalculatedAt returns a boolean if a field has been set.

### GetTeamId

`func (o *PatchedAction) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PatchedAction) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PatchedAction) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *PatchedAction) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetIsAction

`func (o *PatchedAction) GetIsAction() bool`

GetIsAction returns the IsAction field if non-nil, zero value otherwise.

### GetIsActionOk

`func (o *PatchedAction) GetIsActionOk() (*bool, bool)`

GetIsActionOk returns a tuple with the IsAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAction

`func (o *PatchedAction) SetIsAction(v bool)`

SetIsAction sets IsAction field to given value.

### HasIsAction

`func (o *PatchedAction) HasIsAction() bool`

HasIsAction returns a boolean if a field has been set.

### GetBytecodeError

`func (o *PatchedAction) GetBytecodeError() string`

GetBytecodeError returns the BytecodeError field if non-nil, zero value otherwise.

### GetBytecodeErrorOk

`func (o *PatchedAction) GetBytecodeErrorOk() (*string, bool)`

GetBytecodeErrorOk returns a tuple with the BytecodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecodeError

`func (o *PatchedAction) SetBytecodeError(v string)`

SetBytecodeError sets BytecodeError field to given value.

### HasBytecodeError

`func (o *PatchedAction) HasBytecodeError() bool`

HasBytecodeError returns a boolean if a field has been set.

### SetBytecodeErrorNil

`func (o *PatchedAction) SetBytecodeErrorNil(b bool)`

 SetBytecodeErrorNil sets the value for BytecodeError to be an explicit nil

### UnsetBytecodeError
`func (o *PatchedAction) UnsetBytecodeError()`

UnsetBytecodeError ensures that no value is present for BytecodeError, not even an explicit nil
### GetPinnedAt

`func (o *PatchedAction) GetPinnedAt() time.Time`

GetPinnedAt returns the PinnedAt field if non-nil, zero value otherwise.

### GetPinnedAtOk

`func (o *PatchedAction) GetPinnedAtOk() (*time.Time, bool)`

GetPinnedAtOk returns a tuple with the PinnedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedAt

`func (o *PatchedAction) SetPinnedAt(v time.Time)`

SetPinnedAt sets PinnedAt field to given value.

### HasPinnedAt

`func (o *PatchedAction) HasPinnedAt() bool`

HasPinnedAt returns a boolean if a field has been set.

### SetPinnedAtNil

`func (o *PatchedAction) SetPinnedAtNil(b bool)`

 SetPinnedAtNil sets the value for PinnedAt to be an explicit nil

### UnsetPinnedAt
`func (o *PatchedAction) UnsetPinnedAt()`

UnsetPinnedAt ensures that no value is present for PinnedAt, not even an explicit nil
### GetCreationContext

`func (o *PatchedAction) GetCreationContext() string`

GetCreationContext returns the CreationContext field if non-nil, zero value otherwise.

### GetCreationContextOk

`func (o *PatchedAction) GetCreationContextOk() (*string, bool)`

GetCreationContextOk returns a tuple with the CreationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationContext

`func (o *PatchedAction) SetCreationContext(v string)`

SetCreationContext sets CreationContext field to given value.

### HasCreationContext

`func (o *PatchedAction) HasCreationContext() bool`

HasCreationContext returns a boolean if a field has been set.

### GetCreateInFolder

`func (o *PatchedAction) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedAction) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedAction) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedAction) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetUserAccessLevel

`func (o *PatchedAction) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *PatchedAction) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *PatchedAction) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.

### HasUserAccessLevel

`func (o *PatchedAction) HasUserAccessLevel() bool`

HasUserAccessLevel returns a boolean if a field has been set.

### SetUserAccessLevelNil

`func (o *PatchedAction) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *PatchedAction) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


