# ChartSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoalLines** | Pointer to [**[]GoalLine**](GoalLine.md) |  | [optional] 
**LeftYAxisSettings** | Pointer to [**YAxisSettings**](YAxisSettings.md) |  | [optional] 
**RightYAxisSettings** | Pointer to [**YAxisSettings**](YAxisSettings.md) |  | [optional] 
**SeriesBreakdownColumn** | Pointer to **NullableString** |  | [optional] 
**ShowLegend** | Pointer to **NullableBool** |  | [optional] 
**ShowTotalRow** | Pointer to **NullableBool** |  | [optional] 
**ShowXAxisBorder** | Pointer to **NullableBool** |  | [optional] 
**ShowXAxisTicks** | Pointer to **NullableBool** |  | [optional] 
**ShowYAxisBorder** | Pointer to **NullableBool** |  | [optional] 
**StackBars100** | Pointer to **NullableBool** | Whether we fill the bars to 100% in stacked mode | [optional] 
**XAxis** | Pointer to [**ChartAxis**](ChartAxis.md) |  | [optional] 
**YAxis** | Pointer to [**[]ChartAxis**](ChartAxis.md) |  | [optional] 
**YAxisAtZero** | Pointer to **NullableBool** | Deprecated: use &#x60;[left|right]YAxisSettings&#x60;. Whether the Y axis should start at zero | [optional] 

## Methods

### NewChartSettings

`func NewChartSettings() *ChartSettings`

NewChartSettings instantiates a new ChartSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartSettingsWithDefaults

`func NewChartSettingsWithDefaults() *ChartSettings`

NewChartSettingsWithDefaults instantiates a new ChartSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoalLines

`func (o *ChartSettings) GetGoalLines() []GoalLine`

GetGoalLines returns the GoalLines field if non-nil, zero value otherwise.

### GetGoalLinesOk

`func (o *ChartSettings) GetGoalLinesOk() (*[]GoalLine, bool)`

GetGoalLinesOk returns a tuple with the GoalLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalLines

`func (o *ChartSettings) SetGoalLines(v []GoalLine)`

SetGoalLines sets GoalLines field to given value.

### HasGoalLines

`func (o *ChartSettings) HasGoalLines() bool`

HasGoalLines returns a boolean if a field has been set.

### SetGoalLinesNil

`func (o *ChartSettings) SetGoalLinesNil(b bool)`

 SetGoalLinesNil sets the value for GoalLines to be an explicit nil

### UnsetGoalLines
`func (o *ChartSettings) UnsetGoalLines()`

UnsetGoalLines ensures that no value is present for GoalLines, not even an explicit nil
### GetLeftYAxisSettings

`func (o *ChartSettings) GetLeftYAxisSettings() YAxisSettings`

GetLeftYAxisSettings returns the LeftYAxisSettings field if non-nil, zero value otherwise.

### GetLeftYAxisSettingsOk

`func (o *ChartSettings) GetLeftYAxisSettingsOk() (*YAxisSettings, bool)`

GetLeftYAxisSettingsOk returns a tuple with the LeftYAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftYAxisSettings

`func (o *ChartSettings) SetLeftYAxisSettings(v YAxisSettings)`

SetLeftYAxisSettings sets LeftYAxisSettings field to given value.

### HasLeftYAxisSettings

`func (o *ChartSettings) HasLeftYAxisSettings() bool`

HasLeftYAxisSettings returns a boolean if a field has been set.

### GetRightYAxisSettings

`func (o *ChartSettings) GetRightYAxisSettings() YAxisSettings`

GetRightYAxisSettings returns the RightYAxisSettings field if non-nil, zero value otherwise.

### GetRightYAxisSettingsOk

`func (o *ChartSettings) GetRightYAxisSettingsOk() (*YAxisSettings, bool)`

GetRightYAxisSettingsOk returns a tuple with the RightYAxisSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightYAxisSettings

`func (o *ChartSettings) SetRightYAxisSettings(v YAxisSettings)`

SetRightYAxisSettings sets RightYAxisSettings field to given value.

### HasRightYAxisSettings

`func (o *ChartSettings) HasRightYAxisSettings() bool`

HasRightYAxisSettings returns a boolean if a field has been set.

### GetSeriesBreakdownColumn

`func (o *ChartSettings) GetSeriesBreakdownColumn() string`

GetSeriesBreakdownColumn returns the SeriesBreakdownColumn field if non-nil, zero value otherwise.

### GetSeriesBreakdownColumnOk

`func (o *ChartSettings) GetSeriesBreakdownColumnOk() (*string, bool)`

GetSeriesBreakdownColumnOk returns a tuple with the SeriesBreakdownColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesBreakdownColumn

`func (o *ChartSettings) SetSeriesBreakdownColumn(v string)`

SetSeriesBreakdownColumn sets SeriesBreakdownColumn field to given value.

### HasSeriesBreakdownColumn

`func (o *ChartSettings) HasSeriesBreakdownColumn() bool`

HasSeriesBreakdownColumn returns a boolean if a field has been set.

### SetSeriesBreakdownColumnNil

`func (o *ChartSettings) SetSeriesBreakdownColumnNil(b bool)`

 SetSeriesBreakdownColumnNil sets the value for SeriesBreakdownColumn to be an explicit nil

### UnsetSeriesBreakdownColumn
`func (o *ChartSettings) UnsetSeriesBreakdownColumn()`

UnsetSeriesBreakdownColumn ensures that no value is present for SeriesBreakdownColumn, not even an explicit nil
### GetShowLegend

`func (o *ChartSettings) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *ChartSettings) GetShowLegendOk() (*bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegend

`func (o *ChartSettings) SetShowLegend(v bool)`

SetShowLegend sets ShowLegend field to given value.

### HasShowLegend

`func (o *ChartSettings) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### SetShowLegendNil

`func (o *ChartSettings) SetShowLegendNil(b bool)`

 SetShowLegendNil sets the value for ShowLegend to be an explicit nil

### UnsetShowLegend
`func (o *ChartSettings) UnsetShowLegend()`

UnsetShowLegend ensures that no value is present for ShowLegend, not even an explicit nil
### GetShowTotalRow

`func (o *ChartSettings) GetShowTotalRow() bool`

GetShowTotalRow returns the ShowTotalRow field if non-nil, zero value otherwise.

### GetShowTotalRowOk

`func (o *ChartSettings) GetShowTotalRowOk() (*bool, bool)`

GetShowTotalRowOk returns a tuple with the ShowTotalRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTotalRow

`func (o *ChartSettings) SetShowTotalRow(v bool)`

SetShowTotalRow sets ShowTotalRow field to given value.

### HasShowTotalRow

`func (o *ChartSettings) HasShowTotalRow() bool`

HasShowTotalRow returns a boolean if a field has been set.

### SetShowTotalRowNil

`func (o *ChartSettings) SetShowTotalRowNil(b bool)`

 SetShowTotalRowNil sets the value for ShowTotalRow to be an explicit nil

### UnsetShowTotalRow
`func (o *ChartSettings) UnsetShowTotalRow()`

UnsetShowTotalRow ensures that no value is present for ShowTotalRow, not even an explicit nil
### GetShowXAxisBorder

`func (o *ChartSettings) GetShowXAxisBorder() bool`

GetShowXAxisBorder returns the ShowXAxisBorder field if non-nil, zero value otherwise.

### GetShowXAxisBorderOk

`func (o *ChartSettings) GetShowXAxisBorderOk() (*bool, bool)`

GetShowXAxisBorderOk returns a tuple with the ShowXAxisBorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowXAxisBorder

`func (o *ChartSettings) SetShowXAxisBorder(v bool)`

SetShowXAxisBorder sets ShowXAxisBorder field to given value.

### HasShowXAxisBorder

`func (o *ChartSettings) HasShowXAxisBorder() bool`

HasShowXAxisBorder returns a boolean if a field has been set.

### SetShowXAxisBorderNil

`func (o *ChartSettings) SetShowXAxisBorderNil(b bool)`

 SetShowXAxisBorderNil sets the value for ShowXAxisBorder to be an explicit nil

### UnsetShowXAxisBorder
`func (o *ChartSettings) UnsetShowXAxisBorder()`

UnsetShowXAxisBorder ensures that no value is present for ShowXAxisBorder, not even an explicit nil
### GetShowXAxisTicks

`func (o *ChartSettings) GetShowXAxisTicks() bool`

GetShowXAxisTicks returns the ShowXAxisTicks field if non-nil, zero value otherwise.

### GetShowXAxisTicksOk

`func (o *ChartSettings) GetShowXAxisTicksOk() (*bool, bool)`

GetShowXAxisTicksOk returns a tuple with the ShowXAxisTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowXAxisTicks

`func (o *ChartSettings) SetShowXAxisTicks(v bool)`

SetShowXAxisTicks sets ShowXAxisTicks field to given value.

### HasShowXAxisTicks

`func (o *ChartSettings) HasShowXAxisTicks() bool`

HasShowXAxisTicks returns a boolean if a field has been set.

### SetShowXAxisTicksNil

`func (o *ChartSettings) SetShowXAxisTicksNil(b bool)`

 SetShowXAxisTicksNil sets the value for ShowXAxisTicks to be an explicit nil

### UnsetShowXAxisTicks
`func (o *ChartSettings) UnsetShowXAxisTicks()`

UnsetShowXAxisTicks ensures that no value is present for ShowXAxisTicks, not even an explicit nil
### GetShowYAxisBorder

`func (o *ChartSettings) GetShowYAxisBorder() bool`

GetShowYAxisBorder returns the ShowYAxisBorder field if non-nil, zero value otherwise.

### GetShowYAxisBorderOk

`func (o *ChartSettings) GetShowYAxisBorderOk() (*bool, bool)`

GetShowYAxisBorderOk returns a tuple with the ShowYAxisBorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowYAxisBorder

`func (o *ChartSettings) SetShowYAxisBorder(v bool)`

SetShowYAxisBorder sets ShowYAxisBorder field to given value.

### HasShowYAxisBorder

`func (o *ChartSettings) HasShowYAxisBorder() bool`

HasShowYAxisBorder returns a boolean if a field has been set.

### SetShowYAxisBorderNil

`func (o *ChartSettings) SetShowYAxisBorderNil(b bool)`

 SetShowYAxisBorderNil sets the value for ShowYAxisBorder to be an explicit nil

### UnsetShowYAxisBorder
`func (o *ChartSettings) UnsetShowYAxisBorder()`

UnsetShowYAxisBorder ensures that no value is present for ShowYAxisBorder, not even an explicit nil
### GetStackBars100

`func (o *ChartSettings) GetStackBars100() bool`

GetStackBars100 returns the StackBars100 field if non-nil, zero value otherwise.

### GetStackBars100Ok

`func (o *ChartSettings) GetStackBars100Ok() (*bool, bool)`

GetStackBars100Ok returns a tuple with the StackBars100 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackBars100

`func (o *ChartSettings) SetStackBars100(v bool)`

SetStackBars100 sets StackBars100 field to given value.

### HasStackBars100

`func (o *ChartSettings) HasStackBars100() bool`

HasStackBars100 returns a boolean if a field has been set.

### SetStackBars100Nil

`func (o *ChartSettings) SetStackBars100Nil(b bool)`

 SetStackBars100Nil sets the value for StackBars100 to be an explicit nil

### UnsetStackBars100
`func (o *ChartSettings) UnsetStackBars100()`

UnsetStackBars100 ensures that no value is present for StackBars100, not even an explicit nil
### GetXAxis

`func (o *ChartSettings) GetXAxis() ChartAxis`

GetXAxis returns the XAxis field if non-nil, zero value otherwise.

### GetXAxisOk

`func (o *ChartSettings) GetXAxisOk() (*ChartAxis, bool)`

GetXAxisOk returns a tuple with the XAxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxis

`func (o *ChartSettings) SetXAxis(v ChartAxis)`

SetXAxis sets XAxis field to given value.

### HasXAxis

`func (o *ChartSettings) HasXAxis() bool`

HasXAxis returns a boolean if a field has been set.

### GetYAxis

`func (o *ChartSettings) GetYAxis() []ChartAxis`

GetYAxis returns the YAxis field if non-nil, zero value otherwise.

### GetYAxisOk

`func (o *ChartSettings) GetYAxisOk() (*[]ChartAxis, bool)`

GetYAxisOk returns a tuple with the YAxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxis

`func (o *ChartSettings) SetYAxis(v []ChartAxis)`

SetYAxis sets YAxis field to given value.

### HasYAxis

`func (o *ChartSettings) HasYAxis() bool`

HasYAxis returns a boolean if a field has been set.

### SetYAxisNil

`func (o *ChartSettings) SetYAxisNil(b bool)`

 SetYAxisNil sets the value for YAxis to be an explicit nil

### UnsetYAxis
`func (o *ChartSettings) UnsetYAxis()`

UnsetYAxis ensures that no value is present for YAxis, not even an explicit nil
### GetYAxisAtZero

`func (o *ChartSettings) GetYAxisAtZero() bool`

GetYAxisAtZero returns the YAxisAtZero field if non-nil, zero value otherwise.

### GetYAxisAtZeroOk

`func (o *ChartSettings) GetYAxisAtZeroOk() (*bool, bool)`

GetYAxisAtZeroOk returns a tuple with the YAxisAtZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisAtZero

`func (o *ChartSettings) SetYAxisAtZero(v bool)`

SetYAxisAtZero sets YAxisAtZero field to given value.

### HasYAxisAtZero

`func (o *ChartSettings) HasYAxisAtZero() bool`

HasYAxisAtZero returns a boolean if a field has been set.

### SetYAxisAtZeroNil

`func (o *ChartSettings) SetYAxisAtZeroNil(b bool)`

 SetYAxisAtZeroNil sets the value for YAxisAtZero to be an explicit nil

### UnsetYAxisAtZero
`func (o *ChartSettings) UnsetYAxisAtZero()`

UnsetYAxisAtZero ensures that no value is present for YAxisAtZero, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


