# StickinessCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**StickinessOperator**](StickinessOperator.md) |  | 
**Value** | **int32** |  | 

## Methods

### NewStickinessCriteria

`func NewStickinessCriteria(operator StickinessOperator, value int32, ) *StickinessCriteria`

NewStickinessCriteria instantiates a new StickinessCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStickinessCriteriaWithDefaults

`func NewStickinessCriteriaWithDefaults() *StickinessCriteria`

NewStickinessCriteriaWithDefaults instantiates a new StickinessCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *StickinessCriteria) GetOperator() StickinessOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StickinessCriteria) GetOperatorOk() (*StickinessOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StickinessCriteria) SetOperator(v StickinessOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *StickinessCriteria) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StickinessCriteria) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StickinessCriteria) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


