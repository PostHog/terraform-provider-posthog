# PatchedTaskRunDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Task** | Pointer to **string** |  | [optional] [readonly] 
**Stage** | Pointer to **NullableString** | Current stage for this run (e.g., &#39;backlog&#39;, &#39;in_progress&#39;, &#39;done&#39;) | [optional] 
**Branch** | Pointer to **NullableString** | Branch name for the run | [optional] 
**Status** | Pointer to [**TaskRunDetailStatusEnum**](TaskRunDetailStatusEnum.md) |  | [optional] 
**LogUrl** | Pointer to **NullableString** | Presigned S3 URL for log access (valid for 1 hour). | [optional] [readonly] 
**ErrorMessage** | Pointer to **NullableString** | Error message if execution failed | [optional] 
**Output** | Pointer to **interface{}** | Run output data (e.g., PR URL, commit SHA, etc.) | [optional] 
**State** | Pointer to **interface{}** | Run state data for resuming or tracking execution state | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewPatchedTaskRunDetail

`func NewPatchedTaskRunDetail() *PatchedTaskRunDetail`

NewPatchedTaskRunDetail instantiates a new PatchedTaskRunDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTaskRunDetailWithDefaults

`func NewPatchedTaskRunDetailWithDefaults() *PatchedTaskRunDetail`

NewPatchedTaskRunDetailWithDefaults instantiates a new PatchedTaskRunDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedTaskRunDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedTaskRunDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedTaskRunDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedTaskRunDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTask

`func (o *PatchedTaskRunDetail) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *PatchedTaskRunDetail) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *PatchedTaskRunDetail) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *PatchedTaskRunDetail) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetStage

`func (o *PatchedTaskRunDetail) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PatchedTaskRunDetail) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PatchedTaskRunDetail) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PatchedTaskRunDetail) HasStage() bool`

HasStage returns a boolean if a field has been set.

### SetStageNil

`func (o *PatchedTaskRunDetail) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *PatchedTaskRunDetail) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil
### GetBranch

`func (o *PatchedTaskRunDetail) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PatchedTaskRunDetail) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PatchedTaskRunDetail) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PatchedTaskRunDetail) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *PatchedTaskRunDetail) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *PatchedTaskRunDetail) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetStatus

`func (o *PatchedTaskRunDetail) GetStatus() TaskRunDetailStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedTaskRunDetail) GetStatusOk() (*TaskRunDetailStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedTaskRunDetail) SetStatus(v TaskRunDetailStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedTaskRunDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogUrl

`func (o *PatchedTaskRunDetail) GetLogUrl() string`

GetLogUrl returns the LogUrl field if non-nil, zero value otherwise.

### GetLogUrlOk

`func (o *PatchedTaskRunDetail) GetLogUrlOk() (*string, bool)`

GetLogUrlOk returns a tuple with the LogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUrl

`func (o *PatchedTaskRunDetail) SetLogUrl(v string)`

SetLogUrl sets LogUrl field to given value.

### HasLogUrl

`func (o *PatchedTaskRunDetail) HasLogUrl() bool`

HasLogUrl returns a boolean if a field has been set.

### SetLogUrlNil

`func (o *PatchedTaskRunDetail) SetLogUrlNil(b bool)`

 SetLogUrlNil sets the value for LogUrl to be an explicit nil

### UnsetLogUrl
`func (o *PatchedTaskRunDetail) UnsetLogUrl()`

UnsetLogUrl ensures that no value is present for LogUrl, not even an explicit nil
### GetErrorMessage

`func (o *PatchedTaskRunDetail) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PatchedTaskRunDetail) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PatchedTaskRunDetail) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PatchedTaskRunDetail) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *PatchedTaskRunDetail) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *PatchedTaskRunDetail) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetOutput

`func (o *PatchedTaskRunDetail) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *PatchedTaskRunDetail) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *PatchedTaskRunDetail) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *PatchedTaskRunDetail) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *PatchedTaskRunDetail) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *PatchedTaskRunDetail) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetState

`func (o *PatchedTaskRunDetail) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PatchedTaskRunDetail) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PatchedTaskRunDetail) SetState(v interface{})`

SetState sets State field to given value.

### HasState

`func (o *PatchedTaskRunDetail) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *PatchedTaskRunDetail) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *PatchedTaskRunDetail) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCreatedAt

`func (o *PatchedTaskRunDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedTaskRunDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedTaskRunDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedTaskRunDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedTaskRunDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedTaskRunDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedTaskRunDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedTaskRunDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *PatchedTaskRunDetail) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *PatchedTaskRunDetail) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *PatchedTaskRunDetail) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *PatchedTaskRunDetail) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *PatchedTaskRunDetail) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *PatchedTaskRunDetail) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


