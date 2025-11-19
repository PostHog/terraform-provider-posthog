# PropertyGroupFilterValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FilterLogicalOperator**](FilterLogicalOperator.md) |  | 
**Values** | [**[]ValuesInner**](ValuesInner.md) |  | 

## Methods

### NewPropertyGroupFilterValue

`func NewPropertyGroupFilterValue(type_ FilterLogicalOperator, values []ValuesInner, ) *PropertyGroupFilterValue`

NewPropertyGroupFilterValue instantiates a new PropertyGroupFilterValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyGroupFilterValueWithDefaults

`func NewPropertyGroupFilterValueWithDefaults() *PropertyGroupFilterValue`

NewPropertyGroupFilterValueWithDefaults instantiates a new PropertyGroupFilterValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PropertyGroupFilterValue) GetType() FilterLogicalOperator`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyGroupFilterValue) GetTypeOk() (*FilterLogicalOperator, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyGroupFilterValue) SetType(v FilterLogicalOperator)`

SetType sets Type field to given value.


### GetValues

`func (o *PropertyGroupFilterValue) GetValues() []ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PropertyGroupFilterValue) GetValuesOk() (*[]ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PropertyGroupFilterValue) SetValues(v []ValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


