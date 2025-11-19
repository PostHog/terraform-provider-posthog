# LiveDebuggerBreakpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Repository** | Pointer to **NullableString** |  | [optional] 
**Filename** | **string** |  | 
**LineNumber** | **int32** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Condition** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewLiveDebuggerBreakpoint

`func NewLiveDebuggerBreakpoint(id string, filename string, lineNumber int32, createdAt time.Time, updatedAt time.Time, ) *LiveDebuggerBreakpoint`

NewLiveDebuggerBreakpoint instantiates a new LiveDebuggerBreakpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveDebuggerBreakpointWithDefaults

`func NewLiveDebuggerBreakpointWithDefaults() *LiveDebuggerBreakpoint`

NewLiveDebuggerBreakpointWithDefaults instantiates a new LiveDebuggerBreakpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveDebuggerBreakpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveDebuggerBreakpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveDebuggerBreakpoint) SetId(v string)`

SetId sets Id field to given value.


### GetRepository

`func (o *LiveDebuggerBreakpoint) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *LiveDebuggerBreakpoint) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *LiveDebuggerBreakpoint) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *LiveDebuggerBreakpoint) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *LiveDebuggerBreakpoint) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *LiveDebuggerBreakpoint) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetFilename

`func (o *LiveDebuggerBreakpoint) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *LiveDebuggerBreakpoint) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *LiveDebuggerBreakpoint) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetLineNumber

`func (o *LiveDebuggerBreakpoint) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *LiveDebuggerBreakpoint) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *LiveDebuggerBreakpoint) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.


### GetEnabled

`func (o *LiveDebuggerBreakpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LiveDebuggerBreakpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LiveDebuggerBreakpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LiveDebuggerBreakpoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCondition

`func (o *LiveDebuggerBreakpoint) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *LiveDebuggerBreakpoint) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *LiveDebuggerBreakpoint) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *LiveDebuggerBreakpoint) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *LiveDebuggerBreakpoint) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *LiveDebuggerBreakpoint) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetCreatedAt

`func (o *LiveDebuggerBreakpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LiveDebuggerBreakpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LiveDebuggerBreakpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LiveDebuggerBreakpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LiveDebuggerBreakpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LiveDebuggerBreakpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


