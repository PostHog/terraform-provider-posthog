# TaskRunDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Task** | **string** |  | [readonly] 
**Stage** | Pointer to **NullableString** | Current stage for this run (e.g., &#39;backlog&#39;, &#39;in_progress&#39;, &#39;done&#39;) | [optional] 
**Branch** | Pointer to **NullableString** | Branch name for the run | [optional] 
**Status** | Pointer to [**TaskRunDetailStatusEnum**](TaskRunDetailStatusEnum.md) |  | [optional] 
**LogUrl** | **NullableString** | Presigned S3 URL for log access (valid for 1 hour). | [readonly] 
**ErrorMessage** | Pointer to **NullableString** | Error message if execution failed | [optional] 
**Output** | Pointer to **interface{}** | Run output data (e.g., PR URL, commit SHA, etc.) | [optional] 
**State** | Pointer to **interface{}** | Run state data for resuming or tracking execution state | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**CompletedAt** | **NullableTime** |  | [readonly] 

## Methods

### NewTaskRunDetail

`func NewTaskRunDetail(id string, task string, logUrl NullableString, createdAt time.Time, updatedAt time.Time, completedAt NullableTime, ) *TaskRunDetail`

NewTaskRunDetail instantiates a new TaskRunDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRunDetailWithDefaults

`func NewTaskRunDetailWithDefaults() *TaskRunDetail`

NewTaskRunDetailWithDefaults instantiates a new TaskRunDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskRunDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskRunDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskRunDetail) SetId(v string)`

SetId sets Id field to given value.


### GetTask

`func (o *TaskRunDetail) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TaskRunDetail) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TaskRunDetail) SetTask(v string)`

SetTask sets Task field to given value.


### GetStage

`func (o *TaskRunDetail) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *TaskRunDetail) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *TaskRunDetail) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *TaskRunDetail) HasStage() bool`

HasStage returns a boolean if a field has been set.

### SetStageNil

`func (o *TaskRunDetail) SetStageNil(b bool)`

 SetStageNil sets the value for Stage to be an explicit nil

### UnsetStage
`func (o *TaskRunDetail) UnsetStage()`

UnsetStage ensures that no value is present for Stage, not even an explicit nil
### GetBranch

`func (o *TaskRunDetail) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *TaskRunDetail) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *TaskRunDetail) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *TaskRunDetail) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *TaskRunDetail) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *TaskRunDetail) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetStatus

`func (o *TaskRunDetail) GetStatus() TaskRunDetailStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskRunDetail) GetStatusOk() (*TaskRunDetailStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskRunDetail) SetStatus(v TaskRunDetailStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskRunDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogUrl

`func (o *TaskRunDetail) GetLogUrl() string`

GetLogUrl returns the LogUrl field if non-nil, zero value otherwise.

### GetLogUrlOk

`func (o *TaskRunDetail) GetLogUrlOk() (*string, bool)`

GetLogUrlOk returns a tuple with the LogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogUrl

`func (o *TaskRunDetail) SetLogUrl(v string)`

SetLogUrl sets LogUrl field to given value.


### SetLogUrlNil

`func (o *TaskRunDetail) SetLogUrlNil(b bool)`

 SetLogUrlNil sets the value for LogUrl to be an explicit nil

### UnsetLogUrl
`func (o *TaskRunDetail) UnsetLogUrl()`

UnsetLogUrl ensures that no value is present for LogUrl, not even an explicit nil
### GetErrorMessage

`func (o *TaskRunDetail) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TaskRunDetail) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TaskRunDetail) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TaskRunDetail) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TaskRunDetail) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TaskRunDetail) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetOutput

`func (o *TaskRunDetail) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TaskRunDetail) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TaskRunDetail) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TaskRunDetail) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *TaskRunDetail) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *TaskRunDetail) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetState

`func (o *TaskRunDetail) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskRunDetail) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskRunDetail) SetState(v interface{})`

SetState sets State field to given value.

### HasState

`func (o *TaskRunDetail) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *TaskRunDetail) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *TaskRunDetail) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCreatedAt

`func (o *TaskRunDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskRunDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskRunDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TaskRunDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskRunDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskRunDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCompletedAt

`func (o *TaskRunDetail) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TaskRunDetail) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TaskRunDetail) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *TaskRunDetail) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *TaskRunDetail) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


