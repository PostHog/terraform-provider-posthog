# ActiveBreakpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the breakpoint | 
**Repository** | Pointer to **NullableString** | Repository identifier (e.g., &#39;PostHog/posthog&#39;) | [optional] 
**Filename** | **string** | File path where the breakpoint is set | 
**LineNumber** | **int32** | Line number of the breakpoint | 
**Enabled** | **bool** | Whether the breakpoint is enabled | 
**Condition** | Pointer to **NullableString** | Optional condition for the breakpoint | [optional] 

## Methods

### NewActiveBreakpoint

`func NewActiveBreakpoint(id string, filename string, lineNumber int32, enabled bool, ) *ActiveBreakpoint`

NewActiveBreakpoint instantiates a new ActiveBreakpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveBreakpointWithDefaults

`func NewActiveBreakpointWithDefaults() *ActiveBreakpoint`

NewActiveBreakpointWithDefaults instantiates a new ActiveBreakpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActiveBreakpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveBreakpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveBreakpoint) SetId(v string)`

SetId sets Id field to given value.


### GetRepository

`func (o *ActiveBreakpoint) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ActiveBreakpoint) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ActiveBreakpoint) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ActiveBreakpoint) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *ActiveBreakpoint) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *ActiveBreakpoint) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetFilename

`func (o *ActiveBreakpoint) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ActiveBreakpoint) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ActiveBreakpoint) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetLineNumber

`func (o *ActiveBreakpoint) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ActiveBreakpoint) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ActiveBreakpoint) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.


### GetEnabled

`func (o *ActiveBreakpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActiveBreakpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActiveBreakpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCondition

`func (o *ActiveBreakpoint) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ActiveBreakpoint) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ActiveBreakpoint) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ActiveBreakpoint) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ActiveBreakpoint) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ActiveBreakpoint) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


