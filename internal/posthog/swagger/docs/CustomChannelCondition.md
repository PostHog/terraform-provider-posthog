# CustomChannelCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | [**CustomChannelField**](CustomChannelField.md) |  | 
**Op** | [**CustomChannelOperator**](CustomChannelOperator.md) |  | 
**Value** | Pointer to [**NullableValue1**](Value1.md) |  | [optional] [default to null]

## Methods

### NewCustomChannelCondition

`func NewCustomChannelCondition(id string, key CustomChannelField, op CustomChannelOperator, ) *CustomChannelCondition`

NewCustomChannelCondition instantiates a new CustomChannelCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomChannelConditionWithDefaults

`func NewCustomChannelConditionWithDefaults() *CustomChannelCondition`

NewCustomChannelConditionWithDefaults instantiates a new CustomChannelCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomChannelCondition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomChannelCondition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomChannelCondition) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *CustomChannelCondition) GetKey() CustomChannelField`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomChannelCondition) GetKeyOk() (*CustomChannelField, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomChannelCondition) SetKey(v CustomChannelField)`

SetKey sets Key field to given value.


### GetOp

`func (o *CustomChannelCondition) GetOp() CustomChannelOperator`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *CustomChannelCondition) GetOpOk() (*CustomChannelOperator, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *CustomChannelCondition) SetOp(v CustomChannelOperator)`

SetOp sets Op field to given value.


### GetValue

`func (o *CustomChannelCondition) GetValue() Value1`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomChannelCondition) GetValueOk() (*Value1, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomChannelCondition) SetValue(v Value1)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomChannelCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomChannelCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomChannelCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


