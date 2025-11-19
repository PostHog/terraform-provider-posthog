# HogFunctionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**StateEnum**](StateEnum.md) |  | 
**Tokens** | **int32** |  | 

## Methods

### NewHogFunctionStatus

`func NewHogFunctionStatus(state StateEnum, tokens int32, ) *HogFunctionStatus`

NewHogFunctionStatus instantiates a new HogFunctionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogFunctionStatusWithDefaults

`func NewHogFunctionStatusWithDefaults() *HogFunctionStatus`

NewHogFunctionStatusWithDefaults instantiates a new HogFunctionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *HogFunctionStatus) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HogFunctionStatus) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HogFunctionStatus) SetState(v StateEnum)`

SetState sets State field to given value.


### GetTokens

`func (o *HogFunctionStatus) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *HogFunctionStatus) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *HogFunctionStatus) SetTokens(v int32)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


