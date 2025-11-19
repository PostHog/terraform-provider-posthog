# HogFunctionMasking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | **int32** |  | 
**Threshold** | Pointer to **NullableInt32** |  | [optional] 
**Hash** | **string** |  | 
**Bytecode** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewHogFunctionMasking

`func NewHogFunctionMasking(ttl int32, hash string, ) *HogFunctionMasking`

NewHogFunctionMasking instantiates a new HogFunctionMasking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogFunctionMaskingWithDefaults

`func NewHogFunctionMaskingWithDefaults() *HogFunctionMasking`

NewHogFunctionMaskingWithDefaults instantiates a new HogFunctionMasking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *HogFunctionMasking) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *HogFunctionMasking) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *HogFunctionMasking) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetThreshold

`func (o *HogFunctionMasking) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *HogFunctionMasking) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *HogFunctionMasking) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *HogFunctionMasking) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### SetThresholdNil

`func (o *HogFunctionMasking) SetThresholdNil(b bool)`

 SetThresholdNil sets the value for Threshold to be an explicit nil

### UnsetThreshold
`func (o *HogFunctionMasking) UnsetThreshold()`

UnsetThreshold ensures that no value is present for Threshold, not even an explicit nil
### GetHash

`func (o *HogFunctionMasking) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *HogFunctionMasking) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *HogFunctionMasking) SetHash(v string)`

SetHash sets Hash field to given value.


### GetBytecode

`func (o *HogFunctionMasking) GetBytecode() interface{}`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *HogFunctionMasking) GetBytecodeOk() (*interface{}, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *HogFunctionMasking) SetBytecode(v interface{})`

SetBytecode sets Bytecode field to given value.

### HasBytecode

`func (o *HogFunctionMasking) HasBytecode() bool`

HasBytecode returns a boolean if a field has been set.

### SetBytecodeNil

`func (o *HogFunctionMasking) SetBytecodeNil(b bool)`

 SetBytecodeNil sets the value for Bytecode to be an explicit nil

### UnsetBytecode
`func (o *HogFunctionMasking) UnsetBytecode()`

UnsetBytecode ensures that no value is present for Bytecode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


