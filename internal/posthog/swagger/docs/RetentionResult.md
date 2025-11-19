# RetentionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakdownValue** | Pointer to [**NullableBreakdownValue**](BreakdownValue.md) |  | [optional] [default to null]
**Date** | **time.Time** |  | 
**Label** | **string** |  | 
**Values** | [**[]RetentionValue**](RetentionValue.md) |  | 

## Methods

### NewRetentionResult

`func NewRetentionResult(date time.Time, label string, values []RetentionValue, ) *RetentionResult`

NewRetentionResult instantiates a new RetentionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionResultWithDefaults

`func NewRetentionResultWithDefaults() *RetentionResult`

NewRetentionResultWithDefaults instantiates a new RetentionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownValue

`func (o *RetentionResult) GetBreakdownValue() BreakdownValue`

GetBreakdownValue returns the BreakdownValue field if non-nil, zero value otherwise.

### GetBreakdownValueOk

`func (o *RetentionResult) GetBreakdownValueOk() (*BreakdownValue, bool)`

GetBreakdownValueOk returns a tuple with the BreakdownValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownValue

`func (o *RetentionResult) SetBreakdownValue(v BreakdownValue)`

SetBreakdownValue sets BreakdownValue field to given value.

### HasBreakdownValue

`func (o *RetentionResult) HasBreakdownValue() bool`

HasBreakdownValue returns a boolean if a field has been set.

### SetBreakdownValueNil

`func (o *RetentionResult) SetBreakdownValueNil(b bool)`

 SetBreakdownValueNil sets the value for BreakdownValue to be an explicit nil

### UnsetBreakdownValue
`func (o *RetentionResult) UnsetBreakdownValue()`

UnsetBreakdownValue ensures that no value is present for BreakdownValue, not even an explicit nil
### GetDate

`func (o *RetentionResult) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RetentionResult) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RetentionResult) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetLabel

`func (o *RetentionResult) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RetentionResult) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RetentionResult) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValues

`func (o *RetentionResult) GetValues() []RetentionValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RetentionResult) GetValuesOk() (*[]RetentionValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RetentionResult) SetValues(v []RetentionValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


