# InputsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Templating** | Pointer to [**InputsItemTemplatingEnum**](InputsItemTemplatingEnum.md) |  | [optional] 
**Bytecode** | **[]interface{}** |  | [readonly] 
**Order** | **int32** |  | [readonly] 
**Transpiled** | **interface{}** |  | [readonly] 

## Methods

### NewInputsItem

`func NewInputsItem(bytecode []interface{}, order int32, transpiled interface{}, ) *InputsItem`

NewInputsItem instantiates a new InputsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputsItemWithDefaults

`func NewInputsItemWithDefaults() *InputsItem`

NewInputsItemWithDefaults instantiates a new InputsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *InputsItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InputsItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InputsItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InputsItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTemplating

`func (o *InputsItem) GetTemplating() InputsItemTemplatingEnum`

GetTemplating returns the Templating field if non-nil, zero value otherwise.

### GetTemplatingOk

`func (o *InputsItem) GetTemplatingOk() (*InputsItemTemplatingEnum, bool)`

GetTemplatingOk returns a tuple with the Templating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplating

`func (o *InputsItem) SetTemplating(v InputsItemTemplatingEnum)`

SetTemplating sets Templating field to given value.

### HasTemplating

`func (o *InputsItem) HasTemplating() bool`

HasTemplating returns a boolean if a field has been set.

### GetBytecode

`func (o *InputsItem) GetBytecode() []interface{}`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *InputsItem) GetBytecodeOk() (*[]interface{}, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *InputsItem) SetBytecode(v []interface{})`

SetBytecode sets Bytecode field to given value.


### GetOrder

`func (o *InputsItem) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InputsItem) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InputsItem) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetTranspiled

`func (o *InputsItem) GetTranspiled() interface{}`

GetTranspiled returns the Transpiled field if non-nil, zero value otherwise.

### GetTranspiledOk

`func (o *InputsItem) GetTranspiledOk() (*interface{}, bool)`

GetTranspiledOk returns a tuple with the Transpiled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranspiled

`func (o *InputsItem) SetTranspiled(v interface{})`

SetTranspiled sets Transpiled field to given value.


### SetTranspiledNil

`func (o *InputsItem) SetTranspiledNil(b bool)`

 SetTranspiledNil sets the value for Transpiled to be an explicit nil

### UnsetTranspiled
`func (o *InputsItem) UnsetTranspiled()`

UnsetTranspiled ensures that no value is present for Transpiled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


