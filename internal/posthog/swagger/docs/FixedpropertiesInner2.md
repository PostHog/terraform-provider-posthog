# FixedpropertiesInner2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FilterLogicalOperator**](FilterLogicalOperator.md) |  | 
**Values** | [**[]ValuesInner**](ValuesInner.md) |  | 
**Key** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**Operator** | [**PropertyOperator**](PropertyOperator.md) |  | 
**Value** | [**NullableValue2**](Value2.md) |  | [default to null]
**CohortName** | Pointer to **NullableString** |  | [optional] 
**GroupTypeIndex** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewFixedpropertiesInner2

`func NewFixedpropertiesInner2(type_ FilterLogicalOperator, values []ValuesInner, key string, operator PropertyOperator, value NullableValue2, ) *FixedpropertiesInner2`

NewFixedpropertiesInner2 instantiates a new FixedpropertiesInner2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedpropertiesInner2WithDefaults

`func NewFixedpropertiesInner2WithDefaults() *FixedpropertiesInner2`

NewFixedpropertiesInner2WithDefaults instantiates a new FixedpropertiesInner2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FixedpropertiesInner2) GetType() FilterLogicalOperator`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedpropertiesInner2) GetTypeOk() (*FilterLogicalOperator, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedpropertiesInner2) SetType(v FilterLogicalOperator)`

SetType sets Type field to given value.


### GetValues

`func (o *FixedpropertiesInner2) GetValues() []ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FixedpropertiesInner2) GetValuesOk() (*[]ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FixedpropertiesInner2) SetValues(v []ValuesInner)`

SetValues sets Values field to given value.


### GetKey

`func (o *FixedpropertiesInner2) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FixedpropertiesInner2) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FixedpropertiesInner2) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FixedpropertiesInner2) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FixedpropertiesInner2) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FixedpropertiesInner2) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FixedpropertiesInner2) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *FixedpropertiesInner2) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *FixedpropertiesInner2) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetOperator

`func (o *FixedpropertiesInner2) GetOperator() PropertyOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FixedpropertiesInner2) GetOperatorOk() (*PropertyOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FixedpropertiesInner2) SetOperator(v PropertyOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *FixedpropertiesInner2) GetValue() Value2`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FixedpropertiesInner2) GetValueOk() (*Value2, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FixedpropertiesInner2) SetValue(v Value2)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FixedpropertiesInner2) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FixedpropertiesInner2) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCohortName

`func (o *FixedpropertiesInner2) GetCohortName() string`

GetCohortName returns the CohortName field if non-nil, zero value otherwise.

### GetCohortNameOk

`func (o *FixedpropertiesInner2) GetCohortNameOk() (*string, bool)`

GetCohortNameOk returns a tuple with the CohortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortName

`func (o *FixedpropertiesInner2) SetCohortName(v string)`

SetCohortName sets CohortName field to given value.

### HasCohortName

`func (o *FixedpropertiesInner2) HasCohortName() bool`

HasCohortName returns a boolean if a field has been set.

### SetCohortNameNil

`func (o *FixedpropertiesInner2) SetCohortNameNil(b bool)`

 SetCohortNameNil sets the value for CohortName to be an explicit nil

### UnsetCohortName
`func (o *FixedpropertiesInner2) UnsetCohortName()`

UnsetCohortName ensures that no value is present for CohortName, not even an explicit nil
### GetGroupTypeIndex

`func (o *FixedpropertiesInner2) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *FixedpropertiesInner2) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *FixedpropertiesInner2) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *FixedpropertiesInner2) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### SetGroupTypeIndexNil

`func (o *FixedpropertiesInner2) SetGroupTypeIndexNil(b bool)`

 SetGroupTypeIndexNil sets the value for GroupTypeIndex to be an explicit nil

### UnsetGroupTypeIndex
`func (o *FixedpropertiesInner2) UnsetGroupTypeIndex()`

UnsetGroupTypeIndex ensures that no value is present for GroupTypeIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


