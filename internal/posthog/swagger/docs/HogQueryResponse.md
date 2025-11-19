# HogQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytecode** | Pointer to **[]interface{}** |  | [optional] 
**ColoredBytecode** | Pointer to **[]interface{}** |  | [optional] 
**Results** | **interface{}** |  | 
**Stdout** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHogQueryResponse

`func NewHogQueryResponse(results interface{}, ) *HogQueryResponse`

NewHogQueryResponse instantiates a new HogQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQueryResponseWithDefaults

`func NewHogQueryResponseWithDefaults() *HogQueryResponse`

NewHogQueryResponseWithDefaults instantiates a new HogQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytecode

`func (o *HogQueryResponse) GetBytecode() []interface{}`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *HogQueryResponse) GetBytecodeOk() (*[]interface{}, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *HogQueryResponse) SetBytecode(v []interface{})`

SetBytecode sets Bytecode field to given value.

### HasBytecode

`func (o *HogQueryResponse) HasBytecode() bool`

HasBytecode returns a boolean if a field has been set.

### SetBytecodeNil

`func (o *HogQueryResponse) SetBytecodeNil(b bool)`

 SetBytecodeNil sets the value for Bytecode to be an explicit nil

### UnsetBytecode
`func (o *HogQueryResponse) UnsetBytecode()`

UnsetBytecode ensures that no value is present for Bytecode, not even an explicit nil
### GetColoredBytecode

`func (o *HogQueryResponse) GetColoredBytecode() []interface{}`

GetColoredBytecode returns the ColoredBytecode field if non-nil, zero value otherwise.

### GetColoredBytecodeOk

`func (o *HogQueryResponse) GetColoredBytecodeOk() (*[]interface{}, bool)`

GetColoredBytecodeOk returns a tuple with the ColoredBytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColoredBytecode

`func (o *HogQueryResponse) SetColoredBytecode(v []interface{})`

SetColoredBytecode sets ColoredBytecode field to given value.

### HasColoredBytecode

`func (o *HogQueryResponse) HasColoredBytecode() bool`

HasColoredBytecode returns a boolean if a field has been set.

### SetColoredBytecodeNil

`func (o *HogQueryResponse) SetColoredBytecodeNil(b bool)`

 SetColoredBytecodeNil sets the value for ColoredBytecode to be an explicit nil

### UnsetColoredBytecode
`func (o *HogQueryResponse) UnsetColoredBytecode()`

UnsetColoredBytecode ensures that no value is present for ColoredBytecode, not even an explicit nil
### GetResults

`func (o *HogQueryResponse) GetResults() interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *HogQueryResponse) GetResultsOk() (*interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *HogQueryResponse) SetResults(v interface{})`

SetResults sets Results field to given value.


### SetResultsNil

`func (o *HogQueryResponse) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *HogQueryResponse) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetStdout

`func (o *HogQueryResponse) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *HogQueryResponse) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *HogQueryResponse) SetStdout(v string)`

SetStdout sets Stdout field to given value.

### HasStdout

`func (o *HogQueryResponse) HasStdout() bool`

HasStdout returns a boolean if a field has been set.

### SetStdoutNil

`func (o *HogQueryResponse) SetStdoutNil(b bool)`

 SetStdoutNil sets the value for Stdout to be an explicit nil

### UnsetStdout
`func (o *HogQueryResponse) UnsetStdout()`

UnsetStdout ensures that no value is present for Stdout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


