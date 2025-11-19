# FunnelsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinCount** | Pointer to **NullableInt32** |  | [optional] 
**BreakdownAttributionType** | Pointer to [**BreakdownAttributionType**](BreakdownAttributionType.md) |  | [optional] 
**BreakdownAttributionValue** | Pointer to **NullableInt32** |  | [optional] 
**Exclusions** | Pointer to [**[]ExclusionsInner**](ExclusionsInner.md) |  | [optional] [default to {}]
**FunnelAggregateByHogQL** | Pointer to **NullableString** |  | [optional] 
**FunnelFromStep** | Pointer to **NullableInt32** |  | [optional] 
**FunnelOrderType** | Pointer to [**StepOrderValue**](StepOrderValue.md) |  | [optional] 
**FunnelStepReference** | Pointer to [**FunnelStepReference**](FunnelStepReference.md) |  | [optional] 
**FunnelToStep** | Pointer to **NullableInt32** | To select the range of steps for trends &amp; time to convert funnels, 0-indexed | [optional] 
**FunnelVizType** | Pointer to [**FunnelVizType**](FunnelVizType.md) |  | [optional] 
**FunnelWindowInterval** | Pointer to **NullableInt32** |  | [optional] [default to 14]
**FunnelWindowIntervalUnit** | Pointer to [**FunnelConversionWindowTimeUnit**](FunnelConversionWindowTimeUnit.md) |  | [optional] 
**GoalLines** | Pointer to [**[]GoalLine**](GoalLine.md) | Goal Lines | [optional] 
**HiddenLegendBreakdowns** | Pointer to **[]string** |  | [optional] 
**Layout** | Pointer to [**FunnelLayout**](FunnelLayout.md) |  | [optional] 
**ResultCustomizations** | Pointer to [**map[string]ResultCustomizationByValue**](ResultCustomizationByValue.md) | Customizations for the appearance of result datasets. | [optional] 
**ShowValuesOnSeries** | Pointer to **NullableBool** |  | [optional] [default to false]
**UseUdf** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewFunnelsFilter

`func NewFunnelsFilter() *FunnelsFilter`

NewFunnelsFilter instantiates a new FunnelsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunnelsFilterWithDefaults

`func NewFunnelsFilterWithDefaults() *FunnelsFilter`

NewFunnelsFilterWithDefaults instantiates a new FunnelsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinCount

`func (o *FunnelsFilter) GetBinCount() int32`

GetBinCount returns the BinCount field if non-nil, zero value otherwise.

### GetBinCountOk

`func (o *FunnelsFilter) GetBinCountOk() (*int32, bool)`

GetBinCountOk returns a tuple with the BinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCount

`func (o *FunnelsFilter) SetBinCount(v int32)`

SetBinCount sets BinCount field to given value.

### HasBinCount

`func (o *FunnelsFilter) HasBinCount() bool`

HasBinCount returns a boolean if a field has been set.

### SetBinCountNil

`func (o *FunnelsFilter) SetBinCountNil(b bool)`

 SetBinCountNil sets the value for BinCount to be an explicit nil

### UnsetBinCount
`func (o *FunnelsFilter) UnsetBinCount()`

UnsetBinCount ensures that no value is present for BinCount, not even an explicit nil
### GetBreakdownAttributionType

`func (o *FunnelsFilter) GetBreakdownAttributionType() BreakdownAttributionType`

GetBreakdownAttributionType returns the BreakdownAttributionType field if non-nil, zero value otherwise.

### GetBreakdownAttributionTypeOk

`func (o *FunnelsFilter) GetBreakdownAttributionTypeOk() (*BreakdownAttributionType, bool)`

GetBreakdownAttributionTypeOk returns a tuple with the BreakdownAttributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownAttributionType

`func (o *FunnelsFilter) SetBreakdownAttributionType(v BreakdownAttributionType)`

SetBreakdownAttributionType sets BreakdownAttributionType field to given value.

### HasBreakdownAttributionType

`func (o *FunnelsFilter) HasBreakdownAttributionType() bool`

HasBreakdownAttributionType returns a boolean if a field has been set.

### GetBreakdownAttributionValue

`func (o *FunnelsFilter) GetBreakdownAttributionValue() int32`

GetBreakdownAttributionValue returns the BreakdownAttributionValue field if non-nil, zero value otherwise.

### GetBreakdownAttributionValueOk

`func (o *FunnelsFilter) GetBreakdownAttributionValueOk() (*int32, bool)`

GetBreakdownAttributionValueOk returns a tuple with the BreakdownAttributionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownAttributionValue

`func (o *FunnelsFilter) SetBreakdownAttributionValue(v int32)`

SetBreakdownAttributionValue sets BreakdownAttributionValue field to given value.

### HasBreakdownAttributionValue

`func (o *FunnelsFilter) HasBreakdownAttributionValue() bool`

HasBreakdownAttributionValue returns a boolean if a field has been set.

### SetBreakdownAttributionValueNil

`func (o *FunnelsFilter) SetBreakdownAttributionValueNil(b bool)`

 SetBreakdownAttributionValueNil sets the value for BreakdownAttributionValue to be an explicit nil

### UnsetBreakdownAttributionValue
`func (o *FunnelsFilter) UnsetBreakdownAttributionValue()`

UnsetBreakdownAttributionValue ensures that no value is present for BreakdownAttributionValue, not even an explicit nil
### GetExclusions

`func (o *FunnelsFilter) GetExclusions() []ExclusionsInner`

GetExclusions returns the Exclusions field if non-nil, zero value otherwise.

### GetExclusionsOk

`func (o *FunnelsFilter) GetExclusionsOk() (*[]ExclusionsInner, bool)`

GetExclusionsOk returns a tuple with the Exclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusions

`func (o *FunnelsFilter) SetExclusions(v []ExclusionsInner)`

SetExclusions sets Exclusions field to given value.

### HasExclusions

`func (o *FunnelsFilter) HasExclusions() bool`

HasExclusions returns a boolean if a field has been set.

### SetExclusionsNil

`func (o *FunnelsFilter) SetExclusionsNil(b bool)`

 SetExclusionsNil sets the value for Exclusions to be an explicit nil

### UnsetExclusions
`func (o *FunnelsFilter) UnsetExclusions()`

UnsetExclusions ensures that no value is present for Exclusions, not even an explicit nil
### GetFunnelAggregateByHogQL

`func (o *FunnelsFilter) GetFunnelAggregateByHogQL() string`

GetFunnelAggregateByHogQL returns the FunnelAggregateByHogQL field if non-nil, zero value otherwise.

### GetFunnelAggregateByHogQLOk

`func (o *FunnelsFilter) GetFunnelAggregateByHogQLOk() (*string, bool)`

GetFunnelAggregateByHogQLOk returns a tuple with the FunnelAggregateByHogQL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelAggregateByHogQL

`func (o *FunnelsFilter) SetFunnelAggregateByHogQL(v string)`

SetFunnelAggregateByHogQL sets FunnelAggregateByHogQL field to given value.

### HasFunnelAggregateByHogQL

`func (o *FunnelsFilter) HasFunnelAggregateByHogQL() bool`

HasFunnelAggregateByHogQL returns a boolean if a field has been set.

### SetFunnelAggregateByHogQLNil

`func (o *FunnelsFilter) SetFunnelAggregateByHogQLNil(b bool)`

 SetFunnelAggregateByHogQLNil sets the value for FunnelAggregateByHogQL to be an explicit nil

### UnsetFunnelAggregateByHogQL
`func (o *FunnelsFilter) UnsetFunnelAggregateByHogQL()`

UnsetFunnelAggregateByHogQL ensures that no value is present for FunnelAggregateByHogQL, not even an explicit nil
### GetFunnelFromStep

`func (o *FunnelsFilter) GetFunnelFromStep() int32`

GetFunnelFromStep returns the FunnelFromStep field if non-nil, zero value otherwise.

### GetFunnelFromStepOk

`func (o *FunnelsFilter) GetFunnelFromStepOk() (*int32, bool)`

GetFunnelFromStepOk returns a tuple with the FunnelFromStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelFromStep

`func (o *FunnelsFilter) SetFunnelFromStep(v int32)`

SetFunnelFromStep sets FunnelFromStep field to given value.

### HasFunnelFromStep

`func (o *FunnelsFilter) HasFunnelFromStep() bool`

HasFunnelFromStep returns a boolean if a field has been set.

### SetFunnelFromStepNil

`func (o *FunnelsFilter) SetFunnelFromStepNil(b bool)`

 SetFunnelFromStepNil sets the value for FunnelFromStep to be an explicit nil

### UnsetFunnelFromStep
`func (o *FunnelsFilter) UnsetFunnelFromStep()`

UnsetFunnelFromStep ensures that no value is present for FunnelFromStep, not even an explicit nil
### GetFunnelOrderType

`func (o *FunnelsFilter) GetFunnelOrderType() StepOrderValue`

GetFunnelOrderType returns the FunnelOrderType field if non-nil, zero value otherwise.

### GetFunnelOrderTypeOk

`func (o *FunnelsFilter) GetFunnelOrderTypeOk() (*StepOrderValue, bool)`

GetFunnelOrderTypeOk returns a tuple with the FunnelOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelOrderType

`func (o *FunnelsFilter) SetFunnelOrderType(v StepOrderValue)`

SetFunnelOrderType sets FunnelOrderType field to given value.

### HasFunnelOrderType

`func (o *FunnelsFilter) HasFunnelOrderType() bool`

HasFunnelOrderType returns a boolean if a field has been set.

### GetFunnelStepReference

`func (o *FunnelsFilter) GetFunnelStepReference() FunnelStepReference`

GetFunnelStepReference returns the FunnelStepReference field if non-nil, zero value otherwise.

### GetFunnelStepReferenceOk

`func (o *FunnelsFilter) GetFunnelStepReferenceOk() (*FunnelStepReference, bool)`

GetFunnelStepReferenceOk returns a tuple with the FunnelStepReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStepReference

`func (o *FunnelsFilter) SetFunnelStepReference(v FunnelStepReference)`

SetFunnelStepReference sets FunnelStepReference field to given value.

### HasFunnelStepReference

`func (o *FunnelsFilter) HasFunnelStepReference() bool`

HasFunnelStepReference returns a boolean if a field has been set.

### GetFunnelToStep

`func (o *FunnelsFilter) GetFunnelToStep() int32`

GetFunnelToStep returns the FunnelToStep field if non-nil, zero value otherwise.

### GetFunnelToStepOk

`func (o *FunnelsFilter) GetFunnelToStepOk() (*int32, bool)`

GetFunnelToStepOk returns a tuple with the FunnelToStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelToStep

`func (o *FunnelsFilter) SetFunnelToStep(v int32)`

SetFunnelToStep sets FunnelToStep field to given value.

### HasFunnelToStep

`func (o *FunnelsFilter) HasFunnelToStep() bool`

HasFunnelToStep returns a boolean if a field has been set.

### SetFunnelToStepNil

`func (o *FunnelsFilter) SetFunnelToStepNil(b bool)`

 SetFunnelToStepNil sets the value for FunnelToStep to be an explicit nil

### UnsetFunnelToStep
`func (o *FunnelsFilter) UnsetFunnelToStep()`

UnsetFunnelToStep ensures that no value is present for FunnelToStep, not even an explicit nil
### GetFunnelVizType

`func (o *FunnelsFilter) GetFunnelVizType() FunnelVizType`

GetFunnelVizType returns the FunnelVizType field if non-nil, zero value otherwise.

### GetFunnelVizTypeOk

`func (o *FunnelsFilter) GetFunnelVizTypeOk() (*FunnelVizType, bool)`

GetFunnelVizTypeOk returns a tuple with the FunnelVizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelVizType

`func (o *FunnelsFilter) SetFunnelVizType(v FunnelVizType)`

SetFunnelVizType sets FunnelVizType field to given value.

### HasFunnelVizType

`func (o *FunnelsFilter) HasFunnelVizType() bool`

HasFunnelVizType returns a boolean if a field has been set.

### GetFunnelWindowInterval

`func (o *FunnelsFilter) GetFunnelWindowInterval() int32`

GetFunnelWindowInterval returns the FunnelWindowInterval field if non-nil, zero value otherwise.

### GetFunnelWindowIntervalOk

`func (o *FunnelsFilter) GetFunnelWindowIntervalOk() (*int32, bool)`

GetFunnelWindowIntervalOk returns a tuple with the FunnelWindowInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelWindowInterval

`func (o *FunnelsFilter) SetFunnelWindowInterval(v int32)`

SetFunnelWindowInterval sets FunnelWindowInterval field to given value.

### HasFunnelWindowInterval

`func (o *FunnelsFilter) HasFunnelWindowInterval() bool`

HasFunnelWindowInterval returns a boolean if a field has been set.

### SetFunnelWindowIntervalNil

`func (o *FunnelsFilter) SetFunnelWindowIntervalNil(b bool)`

 SetFunnelWindowIntervalNil sets the value for FunnelWindowInterval to be an explicit nil

### UnsetFunnelWindowInterval
`func (o *FunnelsFilter) UnsetFunnelWindowInterval()`

UnsetFunnelWindowInterval ensures that no value is present for FunnelWindowInterval, not even an explicit nil
### GetFunnelWindowIntervalUnit

`func (o *FunnelsFilter) GetFunnelWindowIntervalUnit() FunnelConversionWindowTimeUnit`

GetFunnelWindowIntervalUnit returns the FunnelWindowIntervalUnit field if non-nil, zero value otherwise.

### GetFunnelWindowIntervalUnitOk

`func (o *FunnelsFilter) GetFunnelWindowIntervalUnitOk() (*FunnelConversionWindowTimeUnit, bool)`

GetFunnelWindowIntervalUnitOk returns a tuple with the FunnelWindowIntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelWindowIntervalUnit

`func (o *FunnelsFilter) SetFunnelWindowIntervalUnit(v FunnelConversionWindowTimeUnit)`

SetFunnelWindowIntervalUnit sets FunnelWindowIntervalUnit field to given value.

### HasFunnelWindowIntervalUnit

`func (o *FunnelsFilter) HasFunnelWindowIntervalUnit() bool`

HasFunnelWindowIntervalUnit returns a boolean if a field has been set.

### GetGoalLines

`func (o *FunnelsFilter) GetGoalLines() []GoalLine`

GetGoalLines returns the GoalLines field if non-nil, zero value otherwise.

### GetGoalLinesOk

`func (o *FunnelsFilter) GetGoalLinesOk() (*[]GoalLine, bool)`

GetGoalLinesOk returns a tuple with the GoalLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalLines

`func (o *FunnelsFilter) SetGoalLines(v []GoalLine)`

SetGoalLines sets GoalLines field to given value.

### HasGoalLines

`func (o *FunnelsFilter) HasGoalLines() bool`

HasGoalLines returns a boolean if a field has been set.

### SetGoalLinesNil

`func (o *FunnelsFilter) SetGoalLinesNil(b bool)`

 SetGoalLinesNil sets the value for GoalLines to be an explicit nil

### UnsetGoalLines
`func (o *FunnelsFilter) UnsetGoalLines()`

UnsetGoalLines ensures that no value is present for GoalLines, not even an explicit nil
### GetHiddenLegendBreakdowns

`func (o *FunnelsFilter) GetHiddenLegendBreakdowns() []string`

GetHiddenLegendBreakdowns returns the HiddenLegendBreakdowns field if non-nil, zero value otherwise.

### GetHiddenLegendBreakdownsOk

`func (o *FunnelsFilter) GetHiddenLegendBreakdownsOk() (*[]string, bool)`

GetHiddenLegendBreakdownsOk returns a tuple with the HiddenLegendBreakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenLegendBreakdowns

`func (o *FunnelsFilter) SetHiddenLegendBreakdowns(v []string)`

SetHiddenLegendBreakdowns sets HiddenLegendBreakdowns field to given value.

### HasHiddenLegendBreakdowns

`func (o *FunnelsFilter) HasHiddenLegendBreakdowns() bool`

HasHiddenLegendBreakdowns returns a boolean if a field has been set.

### SetHiddenLegendBreakdownsNil

`func (o *FunnelsFilter) SetHiddenLegendBreakdownsNil(b bool)`

 SetHiddenLegendBreakdownsNil sets the value for HiddenLegendBreakdowns to be an explicit nil

### UnsetHiddenLegendBreakdowns
`func (o *FunnelsFilter) UnsetHiddenLegendBreakdowns()`

UnsetHiddenLegendBreakdowns ensures that no value is present for HiddenLegendBreakdowns, not even an explicit nil
### GetLayout

`func (o *FunnelsFilter) GetLayout() FunnelLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FunnelsFilter) GetLayoutOk() (*FunnelLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FunnelsFilter) SetLayout(v FunnelLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *FunnelsFilter) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetResultCustomizations

`func (o *FunnelsFilter) GetResultCustomizations() map[string]ResultCustomizationByValue`

GetResultCustomizations returns the ResultCustomizations field if non-nil, zero value otherwise.

### GetResultCustomizationsOk

`func (o *FunnelsFilter) GetResultCustomizationsOk() (*map[string]ResultCustomizationByValue, bool)`

GetResultCustomizationsOk returns a tuple with the ResultCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCustomizations

`func (o *FunnelsFilter) SetResultCustomizations(v map[string]ResultCustomizationByValue)`

SetResultCustomizations sets ResultCustomizations field to given value.

### HasResultCustomizations

`func (o *FunnelsFilter) HasResultCustomizations() bool`

HasResultCustomizations returns a boolean if a field has been set.

### SetResultCustomizationsNil

`func (o *FunnelsFilter) SetResultCustomizationsNil(b bool)`

 SetResultCustomizationsNil sets the value for ResultCustomizations to be an explicit nil

### UnsetResultCustomizations
`func (o *FunnelsFilter) UnsetResultCustomizations()`

UnsetResultCustomizations ensures that no value is present for ResultCustomizations, not even an explicit nil
### GetShowValuesOnSeries

`func (o *FunnelsFilter) GetShowValuesOnSeries() bool`

GetShowValuesOnSeries returns the ShowValuesOnSeries field if non-nil, zero value otherwise.

### GetShowValuesOnSeriesOk

`func (o *FunnelsFilter) GetShowValuesOnSeriesOk() (*bool, bool)`

GetShowValuesOnSeriesOk returns a tuple with the ShowValuesOnSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowValuesOnSeries

`func (o *FunnelsFilter) SetShowValuesOnSeries(v bool)`

SetShowValuesOnSeries sets ShowValuesOnSeries field to given value.

### HasShowValuesOnSeries

`func (o *FunnelsFilter) HasShowValuesOnSeries() bool`

HasShowValuesOnSeries returns a boolean if a field has been set.

### SetShowValuesOnSeriesNil

`func (o *FunnelsFilter) SetShowValuesOnSeriesNil(b bool)`

 SetShowValuesOnSeriesNil sets the value for ShowValuesOnSeries to be an explicit nil

### UnsetShowValuesOnSeries
`func (o *FunnelsFilter) UnsetShowValuesOnSeries()`

UnsetShowValuesOnSeries ensures that no value is present for ShowValuesOnSeries, not even an explicit nil
### GetUseUdf

`func (o *FunnelsFilter) GetUseUdf() bool`

GetUseUdf returns the UseUdf field if non-nil, zero value otherwise.

### GetUseUdfOk

`func (o *FunnelsFilter) GetUseUdfOk() (*bool, bool)`

GetUseUdfOk returns a tuple with the UseUdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUdf

`func (o *FunnelsFilter) SetUseUdf(v bool)`

SetUseUdf sets UseUdf field to given value.

### HasUseUdf

`func (o *FunnelsFilter) HasUseUdf() bool`

HasUseUdf returns a boolean if a field has been set.

### SetUseUdfNil

`func (o *FunnelsFilter) SetUseUdfNil(b bool)`

 SetUseUdfNil sets the value for UseUdf to be an explicit nil

### UnsetUseUdf
`func (o *FunnelsFilter) UnsetUseUdf()`

UnsetUseUdf ensures that no value is present for UseUdf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


