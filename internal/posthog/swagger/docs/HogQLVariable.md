# HogQLVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeName** | **string** |  | 
**IsNull** | Pointer to **NullableBool** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] [default to null]
**VariableId** | **string** |  | 

## Methods

### NewHogQLVariable

`func NewHogQLVariable(codeName string, variableId string, ) *HogQLVariable`

NewHogQLVariable instantiates a new HogQLVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogQLVariableWithDefaults

`func NewHogQLVariableWithDefaults() *HogQLVariable`

NewHogQLVariableWithDefaults instantiates a new HogQLVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeName

`func (o *HogQLVariable) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *HogQLVariable) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *HogQLVariable) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.


### GetIsNull

`func (o *HogQLVariable) GetIsNull() bool`

GetIsNull returns the IsNull field if non-nil, zero value otherwise.

### GetIsNullOk

`func (o *HogQLVariable) GetIsNullOk() (*bool, bool)`

GetIsNullOk returns a tuple with the IsNull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNull

`func (o *HogQLVariable) SetIsNull(v bool)`

SetIsNull sets IsNull field to given value.

### HasIsNull

`func (o *HogQLVariable) HasIsNull() bool`

HasIsNull returns a boolean if a field has been set.

### SetIsNullNil

`func (o *HogQLVariable) SetIsNullNil(b bool)`

 SetIsNullNil sets the value for IsNull to be an explicit nil

### UnsetIsNull
`func (o *HogQLVariable) UnsetIsNull()`

UnsetIsNull ensures that no value is present for IsNull, not even an explicit nil
### GetValue

`func (o *HogQLVariable) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HogQLVariable) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HogQLVariable) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *HogQLVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *HogQLVariable) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *HogQLVariable) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetVariableId

`func (o *HogQLVariable) GetVariableId() string`

GetVariableId returns the VariableId field if non-nil, zero value otherwise.

### GetVariableIdOk

`func (o *HogQLVariable) GetVariableIdOk() (*string, bool)`

GetVariableIdOk returns a tuple with the VariableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableId

`func (o *HogQLVariable) SetVariableId(v string)`

SetVariableId sets VariableId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


