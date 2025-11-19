# PersonPropertyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | [**PropertyOperator**](PropertyOperator.md) |  | 
**Type** | Pointer to **string** | Person properties | [optional] [default to "person"]
**Value** | Pointer to [**NullableValue2**](Value2.md) |  | [optional] [default to null]

## Methods

### NewPersonPropertyFilter

`func NewPersonPropertyFilter(key string, operator PropertyOperator, ) *PersonPropertyFilter`

NewPersonPropertyFilter instantiates a new PersonPropertyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonPropertyFilterWithDefaults

`func NewPersonPropertyFilterWithDefaults() *PersonPropertyFilter`

NewPersonPropertyFilterWithDefaults instantiates a new PersonPropertyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PersonPropertyFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PersonPropertyFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PersonPropertyFilter) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *PersonPropertyFilter) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PersonPropertyFilter) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PersonPropertyFilter) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PersonPropertyFilter) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *PersonPropertyFilter) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *PersonPropertyFilter) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *PersonPropertyFilter) GetOperator() PropertyOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PersonPropertyFilter) GetOperatorOk() (*PropertyOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PersonPropertyFilter) SetOperator(v PropertyOperator)`

SetOperator sets Operator field to given value.


### GetType

`func (o *PersonPropertyFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PersonPropertyFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PersonPropertyFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PersonPropertyFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *PersonPropertyFilter) GetValue() Value2`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PersonPropertyFilter) GetValueOk() (*Value2, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PersonPropertyFilter) SetValue(v Value2)`

SetValue sets Value field to given value.

### HasValue

`func (o *PersonPropertyFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PersonPropertyFilter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PersonPropertyFilter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


