# FlagPropertyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key should be the flag ID | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | Pointer to **string** | Only flag_evaluates_to operator is allowed for flag dependencies | [optional] [default to "flag_evaluates_to"]
**Type** | Pointer to **string** | Feature flag dependency | [optional] [default to "flag"]
**Value** | [**Value4**](Value4.md) |  | 

## Methods

### NewFlagPropertyFilter

`func NewFlagPropertyFilter(key string, value Value4, ) *FlagPropertyFilter`

NewFlagPropertyFilter instantiates a new FlagPropertyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagPropertyFilterWithDefaults

`func NewFlagPropertyFilterWithDefaults() *FlagPropertyFilter`

NewFlagPropertyFilterWithDefaults instantiates a new FlagPropertyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FlagPropertyFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagPropertyFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagPropertyFilter) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FlagPropertyFilter) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FlagPropertyFilter) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FlagPropertyFilter) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FlagPropertyFilter) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *FlagPropertyFilter) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *FlagPropertyFilter) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *FlagPropertyFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FlagPropertyFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FlagPropertyFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *FlagPropertyFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetType

`func (o *FlagPropertyFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlagPropertyFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlagPropertyFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlagPropertyFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *FlagPropertyFilter) GetValue() Value4`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FlagPropertyFilter) GetValueOk() (*Value4, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FlagPropertyFilter) SetValue(v Value4)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


