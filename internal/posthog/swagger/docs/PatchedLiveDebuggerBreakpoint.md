# PatchedLiveDebuggerBreakpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Repository** | Pointer to **NullableString** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**LineNumber** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Condition** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPatchedLiveDebuggerBreakpoint

`func NewPatchedLiveDebuggerBreakpoint() *PatchedLiveDebuggerBreakpoint`

NewPatchedLiveDebuggerBreakpoint instantiates a new PatchedLiveDebuggerBreakpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLiveDebuggerBreakpointWithDefaults

`func NewPatchedLiveDebuggerBreakpointWithDefaults() *PatchedLiveDebuggerBreakpoint`

NewPatchedLiveDebuggerBreakpointWithDefaults instantiates a new PatchedLiveDebuggerBreakpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedLiveDebuggerBreakpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedLiveDebuggerBreakpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedLiveDebuggerBreakpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedLiveDebuggerBreakpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRepository

`func (o *PatchedLiveDebuggerBreakpoint) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PatchedLiveDebuggerBreakpoint) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PatchedLiveDebuggerBreakpoint) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PatchedLiveDebuggerBreakpoint) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *PatchedLiveDebuggerBreakpoint) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *PatchedLiveDebuggerBreakpoint) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetFilename

`func (o *PatchedLiveDebuggerBreakpoint) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PatchedLiveDebuggerBreakpoint) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PatchedLiveDebuggerBreakpoint) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PatchedLiveDebuggerBreakpoint) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetLineNumber

`func (o *PatchedLiveDebuggerBreakpoint) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *PatchedLiveDebuggerBreakpoint) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *PatchedLiveDebuggerBreakpoint) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *PatchedLiveDebuggerBreakpoint) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedLiveDebuggerBreakpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedLiveDebuggerBreakpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedLiveDebuggerBreakpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedLiveDebuggerBreakpoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCondition

`func (o *PatchedLiveDebuggerBreakpoint) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *PatchedLiveDebuggerBreakpoint) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *PatchedLiveDebuggerBreakpoint) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *PatchedLiveDebuggerBreakpoint) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *PatchedLiveDebuggerBreakpoint) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *PatchedLiveDebuggerBreakpoint) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetCreatedAt

`func (o *PatchedLiveDebuggerBreakpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedLiveDebuggerBreakpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedLiveDebuggerBreakpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedLiveDebuggerBreakpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedLiveDebuggerBreakpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedLiveDebuggerBreakpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedLiveDebuggerBreakpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedLiveDebuggerBreakpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


