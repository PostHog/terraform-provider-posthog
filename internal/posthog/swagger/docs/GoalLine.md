# GoalLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BorderColor** | Pointer to **NullableString** |  | [optional] 
**DisplayIfCrossed** | Pointer to **NullableBool** |  | [optional] 
**DisplayLabel** | Pointer to **NullableBool** |  | [optional] 
**Label** | **string** |  | 
**Position** | Pointer to [**Position**](Position.md) |  | [optional] 
**Value** | **float32** |  | 

## Methods

### NewGoalLine

`func NewGoalLine(label string, value float32, ) *GoalLine`

NewGoalLine instantiates a new GoalLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalLineWithDefaults

`func NewGoalLineWithDefaults() *GoalLine`

NewGoalLineWithDefaults instantiates a new GoalLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorderColor

`func (o *GoalLine) GetBorderColor() string`

GetBorderColor returns the BorderColor field if non-nil, zero value otherwise.

### GetBorderColorOk

`func (o *GoalLine) GetBorderColorOk() (*string, bool)`

GetBorderColorOk returns a tuple with the BorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColor

`func (o *GoalLine) SetBorderColor(v string)`

SetBorderColor sets BorderColor field to given value.

### HasBorderColor

`func (o *GoalLine) HasBorderColor() bool`

HasBorderColor returns a boolean if a field has been set.

### SetBorderColorNil

`func (o *GoalLine) SetBorderColorNil(b bool)`

 SetBorderColorNil sets the value for BorderColor to be an explicit nil

### UnsetBorderColor
`func (o *GoalLine) UnsetBorderColor()`

UnsetBorderColor ensures that no value is present for BorderColor, not even an explicit nil
### GetDisplayIfCrossed

`func (o *GoalLine) GetDisplayIfCrossed() bool`

GetDisplayIfCrossed returns the DisplayIfCrossed field if non-nil, zero value otherwise.

### GetDisplayIfCrossedOk

`func (o *GoalLine) GetDisplayIfCrossedOk() (*bool, bool)`

GetDisplayIfCrossedOk returns a tuple with the DisplayIfCrossed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIfCrossed

`func (o *GoalLine) SetDisplayIfCrossed(v bool)`

SetDisplayIfCrossed sets DisplayIfCrossed field to given value.

### HasDisplayIfCrossed

`func (o *GoalLine) HasDisplayIfCrossed() bool`

HasDisplayIfCrossed returns a boolean if a field has been set.

### SetDisplayIfCrossedNil

`func (o *GoalLine) SetDisplayIfCrossedNil(b bool)`

 SetDisplayIfCrossedNil sets the value for DisplayIfCrossed to be an explicit nil

### UnsetDisplayIfCrossed
`func (o *GoalLine) UnsetDisplayIfCrossed()`

UnsetDisplayIfCrossed ensures that no value is present for DisplayIfCrossed, not even an explicit nil
### GetDisplayLabel

`func (o *GoalLine) GetDisplayLabel() bool`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *GoalLine) GetDisplayLabelOk() (*bool, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *GoalLine) SetDisplayLabel(v bool)`

SetDisplayLabel sets DisplayLabel field to given value.

### HasDisplayLabel

`func (o *GoalLine) HasDisplayLabel() bool`

HasDisplayLabel returns a boolean if a field has been set.

### SetDisplayLabelNil

`func (o *GoalLine) SetDisplayLabelNil(b bool)`

 SetDisplayLabelNil sets the value for DisplayLabel to be an explicit nil

### UnsetDisplayLabel
`func (o *GoalLine) UnsetDisplayLabel()`

UnsetDisplayLabel ensures that no value is present for DisplayLabel, not even an explicit nil
### GetLabel

`func (o *GoalLine) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GoalLine) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GoalLine) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPosition

`func (o *GoalLine) GetPosition() Position`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GoalLine) GetPositionOk() (*Position, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GoalLine) SetPosition(v Position)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *GoalLine) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetValue

`func (o *GoalLine) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GoalLine) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GoalLine) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


