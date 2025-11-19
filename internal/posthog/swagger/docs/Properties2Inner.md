# Properties2Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 
**Key** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | [**PropertyOperator**](PropertyOperator.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "group"]
**Value** | Pointer to [**NullableValue2**](Value2.md) |  | [optional] [default to null]

## Methods

### NewProperties2Inner

`func NewProperties2Inner(key string, operator PropertyOperator, ) *Properties2Inner`

NewProperties2Inner instantiates a new Properties2Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProperties2InnerWithDefaults

`func NewProperties2InnerWithDefaults() *Properties2Inner`

NewProperties2InnerWithDefaults instantiates a new Properties2Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupTypeIndex

`func (o *Properties2Inner) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *Properties2Inner) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *Properties2Inner) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *Properties2Inner) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *Properties2Inner) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *Properties2Inner) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil
### GetKey

`func (o *Properties2Inner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Properties2Inner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Properties2Inner) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *Properties2Inner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Properties2Inner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Properties2Inner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Properties2Inner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Properties2Inner) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Properties2Inner) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *Properties2Inner) GetOperator() PropertyOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Properties2Inner) GetOperatorOk() (*PropertyOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Properties2Inner) SetOperator(v PropertyOperator)`

SetOperator sets Operator field to given value.


### GetType

`func (o *Properties2Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Properties2Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Properties2Inner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Properties2Inner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *Properties2Inner) GetValue() Value2`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Properties2Inner) GetValueOk() (*Value2, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Properties2Inner) SetValue(v Value2)`

SetValue sets Value field to given value.

### HasValue

`func (o *Properties2Inner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Properties2Inner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Properties2Inner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


