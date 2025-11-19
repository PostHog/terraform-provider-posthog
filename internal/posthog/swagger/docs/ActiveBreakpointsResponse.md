# ActiveBreakpointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakpoints** | [**[]ActiveBreakpoint**](ActiveBreakpoint.md) | List of active breakpoints | 

## Methods

### NewActiveBreakpointsResponse

`func NewActiveBreakpointsResponse(breakpoints []ActiveBreakpoint, ) *ActiveBreakpointsResponse`

NewActiveBreakpointsResponse instantiates a new ActiveBreakpointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveBreakpointsResponseWithDefaults

`func NewActiveBreakpointsResponseWithDefaults() *ActiveBreakpointsResponse`

NewActiveBreakpointsResponseWithDefaults instantiates a new ActiveBreakpointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakpoints

`func (o *ActiveBreakpointsResponse) GetBreakpoints() []ActiveBreakpoint`

GetBreakpoints returns the Breakpoints field if non-nil, zero value otherwise.

### GetBreakpointsOk

`func (o *ActiveBreakpointsResponse) GetBreakpointsOk() (*[]ActiveBreakpoint, bool)`

GetBreakpointsOk returns a tuple with the Breakpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakpoints

`func (o *ActiveBreakpointsResponse) SetBreakpoints(v []ActiveBreakpoint)`

SetBreakpoints sets Breakpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


