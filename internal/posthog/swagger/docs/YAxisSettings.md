# YAxisSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scale** | Pointer to [**Scale**](Scale.md) |  | [optional] 
**ShowGridLines** | Pointer to **NullableBool** |  | [optional] 
**ShowTicks** | Pointer to **NullableBool** |  | [optional] 
**StartAtZero** | Pointer to **NullableBool** | Whether the Y axis should start at zero | [optional] 

## Methods

### NewYAxisSettings

`func NewYAxisSettings() *YAxisSettings`

NewYAxisSettings instantiates a new YAxisSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYAxisSettingsWithDefaults

`func NewYAxisSettingsWithDefaults() *YAxisSettings`

NewYAxisSettingsWithDefaults instantiates a new YAxisSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScale

`func (o *YAxisSettings) GetScale() Scale`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *YAxisSettings) GetScaleOk() (*Scale, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *YAxisSettings) SetScale(v Scale)`

SetScale sets Scale field to given value.

### HasScale

`func (o *YAxisSettings) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetShowGridLines

`func (o *YAxisSettings) GetShowGridLines() bool`

GetShowGridLines returns the ShowGridLines field if non-nil, zero value otherwise.

### GetShowGridLinesOk

`func (o *YAxisSettings) GetShowGridLinesOk() (*bool, bool)`

GetShowGridLinesOk returns a tuple with the ShowGridLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowGridLines

`func (o *YAxisSettings) SetShowGridLines(v bool)`

SetShowGridLines sets ShowGridLines field to given value.

### HasShowGridLines

`func (o *YAxisSettings) HasShowGridLines() bool`

HasShowGridLines returns a boolean if a field has been set.

### SetShowGridLinesNil

`func (o *YAxisSettings) SetShowGridLinesNil(b bool)`

 SetShowGridLinesNil sets the value for ShowGridLines to be an explicit nil

### UnsetShowGridLines
`func (o *YAxisSettings) UnsetShowGridLines()`

UnsetShowGridLines ensures that no value is present for ShowGridLines, not even an explicit nil
### GetShowTicks

`func (o *YAxisSettings) GetShowTicks() bool`

GetShowTicks returns the ShowTicks field if non-nil, zero value otherwise.

### GetShowTicksOk

`func (o *YAxisSettings) GetShowTicksOk() (*bool, bool)`

GetShowTicksOk returns a tuple with the ShowTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTicks

`func (o *YAxisSettings) SetShowTicks(v bool)`

SetShowTicks sets ShowTicks field to given value.

### HasShowTicks

`func (o *YAxisSettings) HasShowTicks() bool`

HasShowTicks returns a boolean if a field has been set.

### SetShowTicksNil

`func (o *YAxisSettings) SetShowTicksNil(b bool)`

 SetShowTicksNil sets the value for ShowTicks to be an explicit nil

### UnsetShowTicks
`func (o *YAxisSettings) UnsetShowTicks()`

UnsetShowTicks ensures that no value is present for ShowTicks, not even an explicit nil
### GetStartAtZero

`func (o *YAxisSettings) GetStartAtZero() bool`

GetStartAtZero returns the StartAtZero field if non-nil, zero value otherwise.

### GetStartAtZeroOk

`func (o *YAxisSettings) GetStartAtZeroOk() (*bool, bool)`

GetStartAtZeroOk returns a tuple with the StartAtZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAtZero

`func (o *YAxisSettings) SetStartAtZero(v bool)`

SetStartAtZero sets StartAtZero field to given value.

### HasStartAtZero

`func (o *YAxisSettings) HasStartAtZero() bool`

HasStartAtZero returns a boolean if a field has been set.

### SetStartAtZeroNil

`func (o *YAxisSettings) SetStartAtZeroNil(b bool)`

 SetStartAtZeroNil sets the value for StartAtZero to be an explicit nil

### UnsetStartAtZero
`func (o *YAxisSettings) UnsetStartAtZero()`

UnsetStartAtZero ensures that no value is present for StartAtZero, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


