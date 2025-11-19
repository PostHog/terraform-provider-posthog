# Properties2Inner1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | [**PropertyOperator**](PropertyOperator.md) |  | 
**Type** | Pointer to **string** | Event properties | [optional] [default to "event"]
**Value** | Pointer to [**NullableValue2**](Value2.md) |  | [optional] [default to null]

## Methods

### NewProperties2Inner1

`func NewProperties2Inner1(key string, operator PropertyOperator, ) *Properties2Inner1`

NewProperties2Inner1 instantiates a new Properties2Inner1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProperties2Inner1WithDefaults

`func NewProperties2Inner1WithDefaults() *Properties2Inner1`

NewProperties2Inner1WithDefaults instantiates a new Properties2Inner1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Properties2Inner1) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Properties2Inner1) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Properties2Inner1) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *Properties2Inner1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Properties2Inner1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Properties2Inner1) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Properties2Inner1) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Properties2Inner1) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Properties2Inner1) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *Properties2Inner1) GetOperator() PropertyOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Properties2Inner1) GetOperatorOk() (*PropertyOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Properties2Inner1) SetOperator(v PropertyOperator)`

SetOperator sets Operator field to given value.


### GetType

`func (o *Properties2Inner1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Properties2Inner1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Properties2Inner1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Properties2Inner1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *Properties2Inner1) GetValue() Value2`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Properties2Inner1) GetValueOk() (*Value2, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Properties2Inner1) SetValue(v Value2)`

SetValue sets Value field to given value.

### HasValue

`func (o *Properties2Inner1) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Properties2Inner1) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Properties2Inner1) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


