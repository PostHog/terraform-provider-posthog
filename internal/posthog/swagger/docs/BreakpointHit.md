# BreakpointHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the hit event | 
**LineNumber** | **int32** | Line number where the breakpoint was hit | 
**FunctionName** | **string** | Name of the function where breakpoint was hit | 
**Timestamp** | **time.Time** | When the breakpoint was hit | 
**Variables** | **map[string]interface{}** | Local variables at the time of the hit | 
**StackTrace** | **[]interface{}** | Stack trace at the time of the hit | 
**BreakpointId** | **string** | ID of the breakpoint that was hit | 
**Filename** | **string** | Filename where the breakpoint was hit | 

## Methods

### NewBreakpointHit

`func NewBreakpointHit(id string, lineNumber int32, functionName string, timestamp time.Time, variables map[string]interface{}, stackTrace []interface{}, breakpointId string, filename string, ) *BreakpointHit`

NewBreakpointHit instantiates a new BreakpointHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakpointHitWithDefaults

`func NewBreakpointHitWithDefaults() *BreakpointHit`

NewBreakpointHitWithDefaults instantiates a new BreakpointHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BreakpointHit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BreakpointHit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BreakpointHit) SetId(v string)`

SetId sets Id field to given value.


### GetLineNumber

`func (o *BreakpointHit) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *BreakpointHit) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *BreakpointHit) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.


### GetFunctionName

`func (o *BreakpointHit) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *BreakpointHit) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *BreakpointHit) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.


### GetTimestamp

`func (o *BreakpointHit) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BreakpointHit) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BreakpointHit) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetVariables

`func (o *BreakpointHit) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *BreakpointHit) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *BreakpointHit) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.


### GetStackTrace

`func (o *BreakpointHit) GetStackTrace() []interface{}`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *BreakpointHit) GetStackTraceOk() (*[]interface{}, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *BreakpointHit) SetStackTrace(v []interface{})`

SetStackTrace sets StackTrace field to given value.


### GetBreakpointId

`func (o *BreakpointHit) GetBreakpointId() string`

GetBreakpointId returns the BreakpointId field if non-nil, zero value otherwise.

### GetBreakpointIdOk

`func (o *BreakpointHit) GetBreakpointIdOk() (*string, bool)`

GetBreakpointIdOk returns a tuple with the BreakpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakpointId

`func (o *BreakpointHit) SetBreakpointId(v string)`

SetBreakpointId sets BreakpointId field to given value.


### GetFilename

`func (o *BreakpointHit) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *BreakpointHit) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *BreakpointHit) SetFilename(v string)`

SetFilename sets Filename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


