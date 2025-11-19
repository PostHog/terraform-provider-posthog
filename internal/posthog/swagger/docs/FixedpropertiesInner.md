# FixedpropertiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | [**PropertyOperator**](PropertyOperator.md) |  | 
**Type** | Pointer to **string** | Event properties | [optional] [default to "event"]
**Value** | [**NullableValue2**](Value2.md) |  | [default to null]
**CohortName** | Pointer to **NullableString** |  | [optional] 
**GroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFixedpropertiesInner

`func NewFixedpropertiesInner(key string, operator PropertyOperator, value NullableValue2, ) *FixedpropertiesInner`

NewFixedpropertiesInner instantiates a new FixedpropertiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedpropertiesInnerWithDefaults

`func NewFixedpropertiesInnerWithDefaults() *FixedpropertiesInner`

NewFixedpropertiesInnerWithDefaults instantiates a new FixedpropertiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FixedpropertiesInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FixedpropertiesInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FixedpropertiesInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FixedpropertiesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FixedpropertiesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FixedpropertiesInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FixedpropertiesInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *FixedpropertiesInner) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *FixedpropertiesInner) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *FixedpropertiesInner) GetOperator() PropertyOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FixedpropertiesInner) GetOperatorOk() (*PropertyOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FixedpropertiesInner) SetOperator(v PropertyOperator)`

SetOperator sets Operator field to given value.


### GetType

`func (o *FixedpropertiesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedpropertiesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedpropertiesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FixedpropertiesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *FixedpropertiesInner) GetValue() Value2`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FixedpropertiesInner) GetValueOk() (*Value2, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FixedpropertiesInner) SetValue(v Value2)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FixedpropertiesInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FixedpropertiesInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCohortName

`func (o *FixedpropertiesInner) GetCohortName() string`

GetCohortName returns the CohortName field if non-nil, zero value otherwise.

### GetCohortNameOk

`func (o *FixedpropertiesInner) GetCohortNameOk() (*string, bool)`

GetCohortNameOk returns a tuple with the CohortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortName

`func (o *FixedpropertiesInner) SetCohortName(v string)`

SetCohortName sets CohortName field to given value.

### HasCohortName

`func (o *FixedpropertiesInner) HasCohortName() bool`

HasCohortName returns a boolean if a field has been set.

### SetCohortNameNil

`func (o *FixedpropertiesInner) SetCohortNameNil(b bool)`

 SetCohortNameNil sets the value for CohortName to be an explicit nil

### UnsetCohortName
`func (o *FixedpropertiesInner) UnsetCohortName()`

UnsetCohortName ensures that no value is present for CohortName, not even an explicit nil
### GetGroupTypeIndex

`func (o *FixedpropertiesInner) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *FixedpropertiesInner) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *FixedpropertiesInner) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *FixedpropertiesInner) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *FixedpropertiesInner) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *FixedpropertiesInner) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


