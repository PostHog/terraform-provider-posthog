# HogFunctionFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**HogFunctionFiltersSourceEnum**](HogFunctionFiltersSourceEnum.md) |  | [optional] [default to EVENTS]
**Actions** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Events** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Properties** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Bytecode** | Pointer to **interface{}** |  | [optional] 
**Transpiled** | Pointer to **interface{}** |  | [optional] 
**FilterTestAccounts** | Pointer to **bool** |  | [optional] 
**BytecodeError** | Pointer to **string** |  | [optional] 

## Methods

### NewHogFunctionFilters

`func NewHogFunctionFilters() *HogFunctionFilters`

NewHogFunctionFilters instantiates a new HogFunctionFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogFunctionFiltersWithDefaults

`func NewHogFunctionFiltersWithDefaults() *HogFunctionFilters`

NewHogFunctionFiltersWithDefaults instantiates a new HogFunctionFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *HogFunctionFilters) GetSource() HogFunctionFiltersSourceEnum`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HogFunctionFilters) GetSourceOk() (*HogFunctionFiltersSourceEnum, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HogFunctionFilters) SetSource(v HogFunctionFiltersSourceEnum)`

SetSource sets Source field to given value.

### HasSource

`func (o *HogFunctionFilters) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetActions

`func (o *HogFunctionFilters) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *HogFunctionFilters) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *HogFunctionFilters) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *HogFunctionFilters) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetEvents

`func (o *HogFunctionFilters) GetEvents() []map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *HogFunctionFilters) GetEventsOk() (*[]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *HogFunctionFilters) SetEvents(v []map[string]interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *HogFunctionFilters) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetProperties

`func (o *HogFunctionFilters) GetProperties() []map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *HogFunctionFilters) GetPropertiesOk() (*[]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *HogFunctionFilters) SetProperties(v []map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *HogFunctionFilters) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetBytecode

`func (o *HogFunctionFilters) GetBytecode() interface{}`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *HogFunctionFilters) GetBytecodeOk() (*interface{}, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *HogFunctionFilters) SetBytecode(v interface{})`

SetBytecode sets Bytecode field to given value.

### HasBytecode

`func (o *HogFunctionFilters) HasBytecode() bool`

HasBytecode returns a boolean if a field has been set.

### SetBytecodeNil

`func (o *HogFunctionFilters) SetBytecodeNil(b bool)`

 SetBytecodeNil sets the value for Bytecode to be an explicit nil

### UnsetBytecode
`func (o *HogFunctionFilters) UnsetBytecode()`

UnsetBytecode ensures that no value is present for Bytecode, not even an explicit nil
### GetTranspiled

`func (o *HogFunctionFilters) GetTranspiled() interface{}`

GetTranspiled returns the Transpiled field if non-nil, zero value otherwise.

### GetTranspiledOk

`func (o *HogFunctionFilters) GetTranspiledOk() (*interface{}, bool)`

GetTranspiledOk returns a tuple with the Transpiled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranspiled

`func (o *HogFunctionFilters) SetTranspiled(v interface{})`

SetTranspiled sets Transpiled field to given value.

### HasTranspiled

`func (o *HogFunctionFilters) HasTranspiled() bool`

HasTranspiled returns a boolean if a field has been set.

### SetTranspiledNil

`func (o *HogFunctionFilters) SetTranspiledNil(b bool)`

 SetTranspiledNil sets the value for Transpiled to be an explicit nil

### UnsetTranspiled
`func (o *HogFunctionFilters) UnsetTranspiled()`

UnsetTranspiled ensures that no value is present for Transpiled, not even an explicit nil
### GetFilterTestAccounts

`func (o *HogFunctionFilters) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *HogFunctionFilters) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *HogFunctionFilters) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *HogFunctionFilters) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### GetBytecodeError

`func (o *HogFunctionFilters) GetBytecodeError() string`

GetBytecodeError returns the BytecodeError field if non-nil, zero value otherwise.

### GetBytecodeErrorOk

`func (o *HogFunctionFilters) GetBytecodeErrorOk() (*string, bool)`

GetBytecodeErrorOk returns a tuple with the BytecodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecodeError

`func (o *HogFunctionFilters) SetBytecodeError(v string)`

SetBytecodeError sets BytecodeError field to given value.

### HasBytecodeError

`func (o *HogFunctionFilters) HasBytecodeError() bool`

HasBytecodeError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


