# LLMTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiSessionId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **string** |  | 
**ErrorCount** | Pointer to **NullableFloat32** |  | [optional] 
**Events** | [**[]LLMTraceEvent**](LLMTraceEvent.md) |  | 
**Id** | **string** |  | 
**InputCost** | Pointer to **NullableFloat32** |  | [optional] 
**InputState** | Pointer to **interface{}** |  | [optional] [default to null]
**InputTokens** | Pointer to **NullableFloat32** |  | [optional] 
**OutputCost** | Pointer to **NullableFloat32** |  | [optional] 
**OutputState** | Pointer to **interface{}** |  | [optional] [default to null]
**OutputTokens** | Pointer to **NullableFloat32** |  | [optional] 
**Person** | [**LLMTracePerson**](LLMTracePerson.md) |  | 
**TotalCost** | Pointer to **NullableFloat32** |  | [optional] 
**TotalLatency** | Pointer to **NullableFloat32** |  | [optional] 
**TraceName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLLMTrace

`func NewLLMTrace(createdAt string, events []LLMTraceEvent, id string, person LLMTracePerson, ) *LLMTrace`

NewLLMTrace instantiates a new LLMTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLLMTraceWithDefaults

`func NewLLMTraceWithDefaults() *LLMTrace`

NewLLMTraceWithDefaults instantiates a new LLMTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiSessionId

`func (o *LLMTrace) GetAiSessionId() string`

GetAiSessionId returns the AiSessionId field if non-nil, zero value otherwise.

### GetAiSessionIdOk

`func (o *LLMTrace) GetAiSessionIdOk() (*string, bool)`

GetAiSessionIdOk returns a tuple with the AiSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSessionId

`func (o *LLMTrace) SetAiSessionId(v string)`

SetAiSessionId sets AiSessionId field to given value.

### HasAiSessionId

`func (o *LLMTrace) HasAiSessionId() bool`

HasAiSessionId returns a boolean if a field has been set.

### SetAiSessionIdNil

`func (o *LLMTrace) SetAiSessionIdNil(b bool)`

 SetAiSessionIdNil sets the value for AiSessionId to be an explicit nil

### UnsetAiSessionId
`func (o *LLMTrace) UnsetAiSessionId()`

UnsetAiSessionId ensures that no value is present for AiSessionId, not even an explicit nil
### GetCreatedAt

`func (o *LLMTrace) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LLMTrace) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LLMTrace) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrorCount

`func (o *LLMTrace) GetErrorCount() float32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *LLMTrace) GetErrorCountOk() (*float32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *LLMTrace) SetErrorCount(v float32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *LLMTrace) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### SetErrorCountNil

`func (o *LLMTrace) SetErrorCountNil(b bool)`

 SetErrorCountNil sets the value for ErrorCount to be an explicit nil

### UnsetErrorCount
`func (o *LLMTrace) UnsetErrorCount()`

UnsetErrorCount ensures that no value is present for ErrorCount, not even an explicit nil
### GetEvents

`func (o *LLMTrace) GetEvents() []LLMTraceEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *LLMTrace) GetEventsOk() (*[]LLMTraceEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *LLMTrace) SetEvents(v []LLMTraceEvent)`

SetEvents sets Events field to given value.


### GetId

`func (o *LLMTrace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LLMTrace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LLMTrace) SetId(v string)`

SetId sets Id field to given value.


### GetInputCost

`func (o *LLMTrace) GetInputCost() float32`

GetInputCost returns the InputCost field if non-nil, zero value otherwise.

### GetInputCostOk

`func (o *LLMTrace) GetInputCostOk() (*float32, bool)`

GetInputCostOk returns a tuple with the InputCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCost

`func (o *LLMTrace) SetInputCost(v float32)`

SetInputCost sets InputCost field to given value.

### HasInputCost

`func (o *LLMTrace) HasInputCost() bool`

HasInputCost returns a boolean if a field has been set.

### SetInputCostNil

`func (o *LLMTrace) SetInputCostNil(b bool)`

 SetInputCostNil sets the value for InputCost to be an explicit nil

### UnsetInputCost
`func (o *LLMTrace) UnsetInputCost()`

UnsetInputCost ensures that no value is present for InputCost, not even an explicit nil
### GetInputState

`func (o *LLMTrace) GetInputState() interface{}`

GetInputState returns the InputState field if non-nil, zero value otherwise.

### GetInputStateOk

`func (o *LLMTrace) GetInputStateOk() (*interface{}, bool)`

GetInputStateOk returns a tuple with the InputState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputState

`func (o *LLMTrace) SetInputState(v interface{})`

SetInputState sets InputState field to given value.

### HasInputState

`func (o *LLMTrace) HasInputState() bool`

HasInputState returns a boolean if a field has been set.

### SetInputStateNil

`func (o *LLMTrace) SetInputStateNil(b bool)`

 SetInputStateNil sets the value for InputState to be an explicit nil

### UnsetInputState
`func (o *LLMTrace) UnsetInputState()`

UnsetInputState ensures that no value is present for InputState, not even an explicit nil
### GetInputTokens

`func (o *LLMTrace) GetInputTokens() float32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *LLMTrace) GetInputTokensOk() (*float32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *LLMTrace) SetInputTokens(v float32)`

SetInputTokens sets InputTokens field to given value.

### HasInputTokens

`func (o *LLMTrace) HasInputTokens() bool`

HasInputTokens returns a boolean if a field has been set.

### SetInputTokensNil

`func (o *LLMTrace) SetInputTokensNil(b bool)`

 SetInputTokensNil sets the value for InputTokens to be an explicit nil

### UnsetInputTokens
`func (o *LLMTrace) UnsetInputTokens()`

UnsetInputTokens ensures that no value is present for InputTokens, not even an explicit nil
### GetOutputCost

`func (o *LLMTrace) GetOutputCost() float32`

GetOutputCost returns the OutputCost field if non-nil, zero value otherwise.

### GetOutputCostOk

`func (o *LLMTrace) GetOutputCostOk() (*float32, bool)`

GetOutputCostOk returns a tuple with the OutputCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputCost

`func (o *LLMTrace) SetOutputCost(v float32)`

SetOutputCost sets OutputCost field to given value.

### HasOutputCost

`func (o *LLMTrace) HasOutputCost() bool`

HasOutputCost returns a boolean if a field has been set.

### SetOutputCostNil

`func (o *LLMTrace) SetOutputCostNil(b bool)`

 SetOutputCostNil sets the value for OutputCost to be an explicit nil

### UnsetOutputCost
`func (o *LLMTrace) UnsetOutputCost()`

UnsetOutputCost ensures that no value is present for OutputCost, not even an explicit nil
### GetOutputState

`func (o *LLMTrace) GetOutputState() interface{}`

GetOutputState returns the OutputState field if non-nil, zero value otherwise.

### GetOutputStateOk

`func (o *LLMTrace) GetOutputStateOk() (*interface{}, bool)`

GetOutputStateOk returns a tuple with the OutputState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputState

`func (o *LLMTrace) SetOutputState(v interface{})`

SetOutputState sets OutputState field to given value.

### HasOutputState

`func (o *LLMTrace) HasOutputState() bool`

HasOutputState returns a boolean if a field has been set.

### SetOutputStateNil

`func (o *LLMTrace) SetOutputStateNil(b bool)`

 SetOutputStateNil sets the value for OutputState to be an explicit nil

### UnsetOutputState
`func (o *LLMTrace) UnsetOutputState()`

UnsetOutputState ensures that no value is present for OutputState, not even an explicit nil
### GetOutputTokens

`func (o *LLMTrace) GetOutputTokens() float32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *LLMTrace) GetOutputTokensOk() (*float32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *LLMTrace) SetOutputTokens(v float32)`

SetOutputTokens sets OutputTokens field to given value.

### HasOutputTokens

`func (o *LLMTrace) HasOutputTokens() bool`

HasOutputTokens returns a boolean if a field has been set.

### SetOutputTokensNil

`func (o *LLMTrace) SetOutputTokensNil(b bool)`

 SetOutputTokensNil sets the value for OutputTokens to be an explicit nil

### UnsetOutputTokens
`func (o *LLMTrace) UnsetOutputTokens()`

UnsetOutputTokens ensures that no value is present for OutputTokens, not even an explicit nil
### GetPerson

`func (o *LLMTrace) GetPerson() LLMTracePerson`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *LLMTrace) GetPersonOk() (*LLMTracePerson, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *LLMTrace) SetPerson(v LLMTracePerson)`

SetPerson sets Person field to given value.


### GetTotalCost

`func (o *LLMTrace) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *LLMTrace) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *LLMTrace) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *LLMTrace) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### SetTotalCostNil

`func (o *LLMTrace) SetTotalCostNil(b bool)`

 SetTotalCostNil sets the value for TotalCost to be an explicit nil

### UnsetTotalCost
`func (o *LLMTrace) UnsetTotalCost()`

UnsetTotalCost ensures that no value is present for TotalCost, not even an explicit nil
### GetTotalLatency

`func (o *LLMTrace) GetTotalLatency() float32`

GetTotalLatency returns the TotalLatency field if non-nil, zero value otherwise.

### GetTotalLatencyOk

`func (o *LLMTrace) GetTotalLatencyOk() (*float32, bool)`

GetTotalLatencyOk returns a tuple with the TotalLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLatency

`func (o *LLMTrace) SetTotalLatency(v float32)`

SetTotalLatency sets TotalLatency field to given value.

### HasTotalLatency

`func (o *LLMTrace) HasTotalLatency() bool`

HasTotalLatency returns a boolean if a field has been set.

### SetTotalLatencyNil

`func (o *LLMTrace) SetTotalLatencyNil(b bool)`

 SetTotalLatencyNil sets the value for TotalLatency to be an explicit nil

### UnsetTotalLatency
`func (o *LLMTrace) UnsetTotalLatency()`

UnsetTotalLatency ensures that no value is present for TotalLatency, not even an explicit nil
### GetTraceName

`func (o *LLMTrace) GetTraceName() string`

GetTraceName returns the TraceName field if non-nil, zero value otherwise.

### GetTraceNameOk

`func (o *LLMTrace) GetTraceNameOk() (*string, bool)`

GetTraceNameOk returns a tuple with the TraceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceName

`func (o *LLMTrace) SetTraceName(v string)`

SetTraceName sets TraceName field to given value.

### HasTraceName

`func (o *LLMTrace) HasTraceName() bool`

HasTraceName returns a boolean if a field has been set.

### SetTraceNameNil

`func (o *LLMTrace) SetTraceNameNil(b bool)`

 SetTraceNameNil sets the value for TraceName to be an explicit nil

### UnsetTraceName
`func (o *LLMTrace) UnsetTraceName()`

UnsetTraceName ensures that no value is present for TraceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


