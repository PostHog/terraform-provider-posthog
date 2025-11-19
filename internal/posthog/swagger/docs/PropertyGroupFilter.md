# PropertyGroupFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FilterLogicalOperator**](FilterLogicalOperator.md) |  | 
**Values** | [**[]PropertyGroupFilterValue**](PropertyGroupFilterValue.md) |  | 

## Methods

### NewPropertyGroupFilter

`func NewPropertyGroupFilter(type_ FilterLogicalOperator, values []PropertyGroupFilterValue, ) *PropertyGroupFilter`

NewPropertyGroupFilter instantiates a new PropertyGroupFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyGroupFilterWithDefaults

`func NewPropertyGroupFilterWithDefaults() *PropertyGroupFilter`

NewPropertyGroupFilterWithDefaults instantiates a new PropertyGroupFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PropertyGroupFilter) GetType() FilterLogicalOperator`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyGroupFilter) GetTypeOk() (*FilterLogicalOperator, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyGroupFilter) SetType(v FilterLogicalOperator)`

SetType sets Type field to given value.


### GetValues

`func (o *PropertyGroupFilter) GetValues() []PropertyGroupFilterValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PropertyGroupFilter) GetValuesOk() (*[]PropertyGroupFilterValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PropertyGroupFilter) SetValues(v []PropertyGroupFilterValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


