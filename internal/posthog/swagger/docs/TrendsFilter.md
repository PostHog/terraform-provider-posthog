# TrendsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationAxisFormat** | Pointer to [**AggregationAxisFormat**](AggregationAxisFormat.md) |  | [optional] 
**AggregationAxisPostfix** | Pointer to **NullableString** |  | [optional] 
**AggregationAxisPrefix** | Pointer to **NullableString** |  | [optional] 
**BreakdownHistogramBinCount** | Pointer to **NullableFloat32** |  | [optional] 
**ConfidenceLevel** | Pointer to **NullableFloat32** |  | [optional] 
**DecimalPlaces** | Pointer to **NullableFloat32** |  | [optional] 
**DetailedResultsAggregationType** | Pointer to [**DetailedResultsAggregationType**](DetailedResultsAggregationType.md) |  | [optional] 
**Display** | Pointer to [**ChartDisplayType**](ChartDisplayType.md) |  | [optional] 
**Formula** | Pointer to **NullableString** |  | [optional] 
**FormulaNodes** | Pointer to [**[]TrendsFormulaNode**](TrendsFormulaNode.md) | List of formulas with optional custom names. Takes precedence over formula/formulas if set. | [optional] 
**Formulas** | Pointer to **[]string** |  | [optional] 
**GoalLines** | Pointer to [**[]GoalLine**](GoalLine.md) | Goal Lines | [optional] 
**HiddenLegendIndexes** | Pointer to **[]int32** |  | [optional] 
**MinDecimalPlaces** | Pointer to **NullableFloat32** |  | [optional] 
**MovingAverageIntervals** | Pointer to **NullableFloat32** |  | [optional] 
**ResultCustomizationBy** | Pointer to [**ResultCustomizationBy**](ResultCustomizationBy.md) |  | [optional] 
**ResultCustomizations** | Pointer to [**NullableResultcustomizations**](Resultcustomizations.md) |  | [optional] [default to null]
**ShowAlertThresholdLines** | Pointer to **NullableBool** |  | [optional] [default to false]
**ShowConfidenceIntervals** | Pointer to **NullableBool** |  | [optional] 
**ShowLabelsOnSeries** | Pointer to **NullableBool** |  | [optional] 
**ShowLegend** | Pointer to **NullableBool** |  | [optional] [default to false]
**ShowMovingAverage** | Pointer to **NullableBool** |  | [optional] 
**ShowMultipleYAxes** | Pointer to **NullableBool** |  | [optional] [default to false]
**ShowPercentStackView** | Pointer to **NullableBool** |  | [optional] [default to false]
**ShowTrendLines** | Pointer to **NullableBool** |  | [optional] 
**ShowValuesOnSeries** | Pointer to **NullableBool** |  | [optional] [default to false]
**SmoothingIntervals** | Pointer to **NullableInt32** |  | [optional] [default to 1]
**YAxisScaleType** | Pointer to [**YAxisScaleType**](YAxisScaleType.md) |  | [optional] 

## Methods

### NewTrendsFilter

`func NewTrendsFilter() *TrendsFilter`

NewTrendsFilter instantiates a new TrendsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrendsFilterWithDefaults

`func NewTrendsFilterWithDefaults() *TrendsFilter`

NewTrendsFilterWithDefaults instantiates a new TrendsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationAxisFormat

`func (o *TrendsFilter) GetAggregationAxisFormat() AggregationAxisFormat`

GetAggregationAxisFormat returns the AggregationAxisFormat field if non-nil, zero value otherwise.

### GetAggregationAxisFormatOk

`func (o *TrendsFilter) GetAggregationAxisFormatOk() (*AggregationAxisFormat, bool)`

GetAggregationAxisFormatOk returns a tuple with the AggregationAxisFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAxisFormat

`func (o *TrendsFilter) SetAggregationAxisFormat(v AggregationAxisFormat)`

SetAggregationAxisFormat sets AggregationAxisFormat field to given value.

### HasAggregationAxisFormat

`func (o *TrendsFilter) HasAggregationAxisFormat() bool`

HasAggregationAxisFormat returns a boolean if a field has been set.

### GetAggregationAxisPostfix

`func (o *TrendsFilter) GetAggregationAxisPostfix() string`

GetAggregationAxisPostfix returns the AggregationAxisPostfix field if non-nil, zero value otherwise.

### GetAggregationAxisPostfixOk

`func (o *TrendsFilter) GetAggregationAxisPostfixOk() (*string, bool)`

GetAggregationAxisPostfixOk returns a tuple with the AggregationAxisPostfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAxisPostfix

`func (o *TrendsFilter) SetAggregationAxisPostfix(v string)`

SetAggregationAxisPostfix sets AggregationAxisPostfix field to given value.

### HasAggregationAxisPostfix

`func (o *TrendsFilter) HasAggregationAxisPostfix() bool`

HasAggregationAxisPostfix returns a boolean if a field has been set.

### SetAggregationAxisPostfixNil

`func (o *TrendsFilter) SetAggregationAxisPostfixNil(b bool)`

 SetAggregationAxisPostfixNil sets the value for AggregationAxisPostfix to be an explicit nil

### UnsetAggregationAxisPostfix
`func (o *TrendsFilter) UnsetAggregationAxisPostfix()`

UnsetAggregationAxisPostfix ensures that no value is present for AggregationAxisPostfix, not even an explicit nil
### GetAggregationAxisPrefix

`func (o *TrendsFilter) GetAggregationAxisPrefix() string`

GetAggregationAxisPrefix returns the AggregationAxisPrefix field if non-nil, zero value otherwise.

### GetAggregationAxisPrefixOk

`func (o *TrendsFilter) GetAggregationAxisPrefixOk() (*string, bool)`

GetAggregationAxisPrefixOk returns a tuple with the AggregationAxisPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationAxisPrefix

`func (o *TrendsFilter) SetAggregationAxisPrefix(v string)`

SetAggregationAxisPrefix sets AggregationAxisPrefix field to given value.

### HasAggregationAxisPrefix

`func (o *TrendsFilter) HasAggregationAxisPrefix() bool`

HasAggregationAxisPrefix returns a boolean if a field has been set.

### SetAggregationAxisPrefixNil

`func (o *TrendsFilter) SetAggregationAxisPrefixNil(b bool)`

 SetAggregationAxisPrefixNil sets the value for AggregationAxisPrefix to be an explicit nil

### UnsetAggregationAxisPrefix
`func (o *TrendsFilter) UnsetAggregationAxisPrefix()`

UnsetAggregationAxisPrefix ensures that no value is present for AggregationAxisPrefix, not even an explicit nil
### GetBreakdownHistogramBinCount

`func (o *TrendsFilter) GetBreakdownHistogramBinCount() float32`

GetBreakdownHistogramBinCount returns the BreakdownHistogramBinCount field if non-nil, zero value otherwise.

### GetBreakdownHistogramBinCountOk

`func (o *TrendsFilter) GetBreakdownHistogramBinCountOk() (*float32, bool)`

GetBreakdownHistogramBinCountOk returns a tuple with the BreakdownHistogramBinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownHistogramBinCount

`func (o *TrendsFilter) SetBreakdownHistogramBinCount(v float32)`

SetBreakdownHistogramBinCount sets BreakdownHistogramBinCount field to given value.

### HasBreakdownHistogramBinCount

`func (o *TrendsFilter) HasBreakdownHistogramBinCount() bool`

HasBreakdownHistogramBinCount returns a boolean if a field has been set.

### SetBreakdownHistogramBinCountNil

`func (o *TrendsFilter) SetBreakdownHistogramBinCountNil(b bool)`

 SetBreakdownHistogramBinCountNil sets the value for BreakdownHistogramBinCount to be an explicit nil

### UnsetBreakdownHistogramBinCount
`func (o *TrendsFilter) UnsetBreakdownHistogramBinCount()`

UnsetBreakdownHistogramBinCount ensures that no value is present for BreakdownHistogramBinCount, not even an explicit nil
### GetConfidenceLevel

`func (o *TrendsFilter) GetConfidenceLevel() float32`

GetConfidenceLevel returns the ConfidenceLevel field if non-nil, zero value otherwise.

### GetConfidenceLevelOk

`func (o *TrendsFilter) GetConfidenceLevelOk() (*float32, bool)`

GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceLevel

`func (o *TrendsFilter) SetConfidenceLevel(v float32)`

SetConfidenceLevel sets ConfidenceLevel field to given value.

### HasConfidenceLevel

`func (o *TrendsFilter) HasConfidenceLevel() bool`

HasConfidenceLevel returns a boolean if a field has been set.

### SetConfidenceLevelNil

`func (o *TrendsFilter) SetConfidenceLevelNil(b bool)`

 SetConfidenceLevelNil sets the value for ConfidenceLevel to be an explicit nil

### UnsetConfidenceLevel
`func (o *TrendsFilter) UnsetConfidenceLevel()`

UnsetConfidenceLevel ensures that no value is present for ConfidenceLevel, not even an explicit nil
### GetDecimalPlaces

`func (o *TrendsFilter) GetDecimalPlaces() float32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *TrendsFilter) GetDecimalPlacesOk() (*float32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *TrendsFilter) SetDecimalPlaces(v float32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *TrendsFilter) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### SetDecimalPlacesNil

`func (o *TrendsFilter) SetDecimalPlacesNil(b bool)`

 SetDecimalPlacesNil sets the value for DecimalPlaces to be an explicit nil

### UnsetDecimalPlaces
`func (o *TrendsFilter) UnsetDecimalPlaces()`

UnsetDecimalPlaces ensures that no value is present for DecimalPlaces, not even an explicit nil
### GetDetailedResultsAggregationType

`func (o *TrendsFilter) GetDetailedResultsAggregationType() DetailedResultsAggregationType`

GetDetailedResultsAggregationType returns the DetailedResultsAggregationType field if non-nil, zero value otherwise.

### GetDetailedResultsAggregationTypeOk

`func (o *TrendsFilter) GetDetailedResultsAggregationTypeOk() (*DetailedResultsAggregationType, bool)`

GetDetailedResultsAggregationTypeOk returns a tuple with the DetailedResultsAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedResultsAggregationType

`func (o *TrendsFilter) SetDetailedResultsAggregationType(v DetailedResultsAggregationType)`

SetDetailedResultsAggregationType sets DetailedResultsAggregationType field to given value.

### HasDetailedResultsAggregationType

`func (o *TrendsFilter) HasDetailedResultsAggregationType() bool`

HasDetailedResultsAggregationType returns a boolean if a field has been set.

### GetDisplay

`func (o *TrendsFilter) GetDisplay() ChartDisplayType`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *TrendsFilter) GetDisplayOk() (*ChartDisplayType, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *TrendsFilter) SetDisplay(v ChartDisplayType)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *TrendsFilter) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetFormula

`func (o *TrendsFilter) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *TrendsFilter) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *TrendsFilter) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *TrendsFilter) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### SetFormulaNil

`func (o *TrendsFilter) SetFormulaNil(b bool)`

 SetFormulaNil sets the value for Formula to be an explicit nil

### UnsetFormula
`func (o *TrendsFilter) UnsetFormula()`

UnsetFormula ensures that no value is present for Formula, not even an explicit nil
### GetFormulaNodes

`func (o *TrendsFilter) GetFormulaNodes() []TrendsFormulaNode`

GetFormulaNodes returns the FormulaNodes field if non-nil, zero value otherwise.

### GetFormulaNodesOk

`func (o *TrendsFilter) GetFormulaNodesOk() (*[]TrendsFormulaNode, bool)`

GetFormulaNodesOk returns a tuple with the FormulaNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaNodes

`func (o *TrendsFilter) SetFormulaNodes(v []TrendsFormulaNode)`

SetFormulaNodes sets FormulaNodes field to given value.

### HasFormulaNodes

`func (o *TrendsFilter) HasFormulaNodes() bool`

HasFormulaNodes returns a boolean if a field has been set.

### SetFormulaNodesNil

`func (o *TrendsFilter) SetFormulaNodesNil(b bool)`

 SetFormulaNodesNil sets the value for FormulaNodes to be an explicit nil

### UnsetFormulaNodes
`func (o *TrendsFilter) UnsetFormulaNodes()`

UnsetFormulaNodes ensures that no value is present for FormulaNodes, not even an explicit nil
### GetFormulas

`func (o *TrendsFilter) GetFormulas() []string`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *TrendsFilter) GetFormulasOk() (*[]string, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulas

`func (o *TrendsFilter) SetFormulas(v []string)`

SetFormulas sets Formulas field to given value.

### HasFormulas

`func (o *TrendsFilter) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### SetFormulasNil

`func (o *TrendsFilter) SetFormulasNil(b bool)`

 SetFormulasNil sets the value for Formulas to be an explicit nil

### UnsetFormulas
`func (o *TrendsFilter) UnsetFormulas()`

UnsetFormulas ensures that no value is present for Formulas, not even an explicit nil
### GetGoalLines

`func (o *TrendsFilter) GetGoalLines() []GoalLine`

GetGoalLines returns the GoalLines field if non-nil, zero value otherwise.

### GetGoalLinesOk

`func (o *TrendsFilter) GetGoalLinesOk() (*[]GoalLine, bool)`

GetGoalLinesOk returns a tuple with the GoalLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalLines

`func (o *TrendsFilter) SetGoalLines(v []GoalLine)`

SetGoalLines sets GoalLines field to given value.

### HasGoalLines

`func (o *TrendsFilter) HasGoalLines() bool`

HasGoalLines returns a boolean if a field has been set.

### SetGoalLinesNil

`func (o *TrendsFilter) SetGoalLinesNil(b bool)`

 SetGoalLinesNil sets the value for GoalLines to be an explicit nil

### UnsetGoalLines
`func (o *TrendsFilter) UnsetGoalLines()`

UnsetGoalLines ensures that no value is present for GoalLines, not even an explicit nil
### GetHiddenLegendIndexes

`func (o *TrendsFilter) GetHiddenLegendIndexes() []int32`

GetHiddenLegendIndexes returns the HiddenLegendIndexes field if non-nil, zero value otherwise.

### GetHiddenLegendIndexesOk

`func (o *TrendsFilter) GetHiddenLegendIndexesOk() (*[]int32, bool)`

GetHiddenLegendIndexesOk returns a tuple with the HiddenLegendIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenLegendIndexes

`func (o *TrendsFilter) SetHiddenLegendIndexes(v []int32)`

SetHiddenLegendIndexes sets HiddenLegendIndexes field to given value.

### HasHiddenLegendIndexes

`func (o *TrendsFilter) HasHiddenLegendIndexes() bool`

HasHiddenLegendIndexes returns a boolean if a field has been set.

### SetHiddenLegendIndexesNil

`func (o *TrendsFilter) SetHiddenLegendIndexesNil(b bool)`

 SetHiddenLegendIndexesNil sets the value for HiddenLegendIndexes to be an explicit nil

### UnsetHiddenLegendIndexes
`func (o *TrendsFilter) UnsetHiddenLegendIndexes()`

UnsetHiddenLegendIndexes ensures that no value is present for HiddenLegendIndexes, not even an explicit nil
### GetMinDecimalPlaces

`func (o *TrendsFilter) GetMinDecimalPlaces() float32`

GetMinDecimalPlaces returns the MinDecimalPlaces field if non-nil, zero value otherwise.

### GetMinDecimalPlacesOk

`func (o *TrendsFilter) GetMinDecimalPlacesOk() (*float32, bool)`

GetMinDecimalPlacesOk returns a tuple with the MinDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDecimalPlaces

`func (o *TrendsFilter) SetMinDecimalPlaces(v float32)`

SetMinDecimalPlaces sets MinDecimalPlaces field to given value.

### HasMinDecimalPlaces

`func (o *TrendsFilter) HasMinDecimalPlaces() bool`

HasMinDecimalPlaces returns a boolean if a field has been set.

### SetMinDecimalPlacesNil

`func (o *TrendsFilter) SetMinDecimalPlacesNil(b bool)`

 SetMinDecimalPlacesNil sets the value for MinDecimalPlaces to be an explicit nil

### UnsetMinDecimalPlaces
`func (o *TrendsFilter) UnsetMinDecimalPlaces()`

UnsetMinDecimalPlaces ensures that no value is present for MinDecimalPlaces, not even an explicit nil
### GetMovingAverageIntervals

`func (o *TrendsFilter) GetMovingAverageIntervals() float32`

GetMovingAverageIntervals returns the MovingAverageIntervals field if non-nil, zero value otherwise.

### GetMovingAverageIntervalsOk

`func (o *TrendsFilter) GetMovingAverageIntervalsOk() (*float32, bool)`

GetMovingAverageIntervalsOk returns a tuple with the MovingAverageIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingAverageIntervals

`func (o *TrendsFilter) SetMovingAverageIntervals(v float32)`

SetMovingAverageIntervals sets MovingAverageIntervals field to given value.

### HasMovingAverageIntervals

`func (o *TrendsFilter) HasMovingAverageIntervals() bool`

HasMovingAverageIntervals returns a boolean if a field has been set.

### SetMovingAverageIntervalsNil

`func (o *TrendsFilter) SetMovingAverageIntervalsNil(b bool)`

 SetMovingAverageIntervalsNil sets the value for MovingAverageIntervals to be an explicit nil

### UnsetMovingAverageIntervals
`func (o *TrendsFilter) UnsetMovingAverageIntervals()`

UnsetMovingAverageIntervals ensures that no value is present for MovingAverageIntervals, not even an explicit nil
### GetResultCustomizationBy

`func (o *TrendsFilter) GetResultCustomizationBy() ResultCustomizationBy`

GetResultCustomizationBy returns the ResultCustomizationBy field if non-nil, zero value otherwise.

### GetResultCustomizationByOk

`func (o *TrendsFilter) GetResultCustomizationByOk() (*ResultCustomizationBy, bool)`

GetResultCustomizationByOk returns a tuple with the ResultCustomizationBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCustomizationBy

`func (o *TrendsFilter) SetResultCustomizationBy(v ResultCustomizationBy)`

SetResultCustomizationBy sets ResultCustomizationBy field to given value.

### HasResultCustomizationBy

`func (o *TrendsFilter) HasResultCustomizationBy() bool`

HasResultCustomizationBy returns a boolean if a field has been set.

### GetResultCustomizations

`func (o *TrendsFilter) GetResultCustomizations() Resultcustomizations`

GetResultCustomizations returns the ResultCustomizations field if non-nil, zero value otherwise.

### GetResultCustomizationsOk

`func (o *TrendsFilter) GetResultCustomizationsOk() (*Resultcustomizations, bool)`

GetResultCustomizationsOk returns a tuple with the ResultCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCustomizations

`func (o *TrendsFilter) SetResultCustomizations(v Resultcustomizations)`

SetResultCustomizations sets ResultCustomizations field to given value.

### HasResultCustomizations

`func (o *TrendsFilter) HasResultCustomizations() bool`

HasResultCustomizations returns a boolean if a field has been set.

### SetResultCustomizationsNil

`func (o *TrendsFilter) SetResultCustomizationsNil(b bool)`

 SetResultCustomizationsNil sets the value for ResultCustomizations to be an explicit nil

### UnsetResultCustomizations
`func (o *TrendsFilter) UnsetResultCustomizations()`

UnsetResultCustomizations ensures that no value is present for ResultCustomizations, not even an explicit nil
### GetShowAlertThresholdLines

`func (o *TrendsFilter) GetShowAlertThresholdLines() bool`

GetShowAlertThresholdLines returns the ShowAlertThresholdLines field if non-nil, zero value otherwise.

### GetShowAlertThresholdLinesOk

`func (o *TrendsFilter) GetShowAlertThresholdLinesOk() (*bool, bool)`

GetShowAlertThresholdLinesOk returns a tuple with the ShowAlertThresholdLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAlertThresholdLines

`func (o *TrendsFilter) SetShowAlertThresholdLines(v bool)`

SetShowAlertThresholdLines sets ShowAlertThresholdLines field to given value.

### HasShowAlertThresholdLines

`func (o *TrendsFilter) HasShowAlertThresholdLines() bool`

HasShowAlertThresholdLines returns a boolean if a field has been set.

### SetShowAlertThresholdLinesNil

`func (o *TrendsFilter) SetShowAlertThresholdLinesNil(b bool)`

 SetShowAlertThresholdLinesNil sets the value for ShowAlertThresholdLines to be an explicit nil

### UnsetShowAlertThresholdLines
`func (o *TrendsFilter) UnsetShowAlertThresholdLines()`

UnsetShowAlertThresholdLines ensures that no value is present for ShowAlertThresholdLines, not even an explicit nil
### GetShowConfidenceIntervals

`func (o *TrendsFilter) GetShowConfidenceIntervals() bool`

GetShowConfidenceIntervals returns the ShowConfidenceIntervals field if non-nil, zero value otherwise.

### GetShowConfidenceIntervalsOk

`func (o *TrendsFilter) GetShowConfidenceIntervalsOk() (*bool, bool)`

GetShowConfidenceIntervalsOk returns a tuple with the ShowConfidenceIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConfidenceIntervals

`func (o *TrendsFilter) SetShowConfidenceIntervals(v bool)`

SetShowConfidenceIntervals sets ShowConfidenceIntervals field to given value.

### HasShowConfidenceIntervals

`func (o *TrendsFilter) HasShowConfidenceIntervals() bool`

HasShowConfidenceIntervals returns a boolean if a field has been set.

### SetShowConfidenceIntervalsNil

`func (o *TrendsFilter) SetShowConfidenceIntervalsNil(b bool)`

 SetShowConfidenceIntervalsNil sets the value for ShowConfidenceIntervals to be an explicit nil

### UnsetShowConfidenceIntervals
`func (o *TrendsFilter) UnsetShowConfidenceIntervals()`

UnsetShowConfidenceIntervals ensures that no value is present for ShowConfidenceIntervals, not even an explicit nil
### GetShowLabelsOnSeries

`func (o *TrendsFilter) GetShowLabelsOnSeries() bool`

GetShowLabelsOnSeries returns the ShowLabelsOnSeries field if non-nil, zero value otherwise.

### GetShowLabelsOnSeriesOk

`func (o *TrendsFilter) GetShowLabelsOnSeriesOk() (*bool, bool)`

GetShowLabelsOnSeriesOk returns a tuple with the ShowLabelsOnSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLabelsOnSeries

`func (o *TrendsFilter) SetShowLabelsOnSeries(v bool)`

SetShowLabelsOnSeries sets ShowLabelsOnSeries field to given value.

### HasShowLabelsOnSeries

`func (o *TrendsFilter) HasShowLabelsOnSeries() bool`

HasShowLabelsOnSeries returns a boolean if a field has been set.

### SetShowLabelsOnSeriesNil

`func (o *TrendsFilter) SetShowLabelsOnSeriesNil(b bool)`

 SetShowLabelsOnSeriesNil sets the value for ShowLabelsOnSeries to be an explicit nil

### UnsetShowLabelsOnSeries
`func (o *TrendsFilter) UnsetShowLabelsOnSeries()`

UnsetShowLabelsOnSeries ensures that no value is present for ShowLabelsOnSeries, not even an explicit nil
### GetShowLegend

`func (o *TrendsFilter) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *TrendsFilter) GetShowLegendOk() (*bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegend

`func (o *TrendsFilter) SetShowLegend(v bool)`

SetShowLegend sets ShowLegend field to given value.

### HasShowLegend

`func (o *TrendsFilter) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### SetShowLegendNil

`func (o *TrendsFilter) SetShowLegendNil(b bool)`

 SetShowLegendNil sets the value for ShowLegend to be an explicit nil

### UnsetShowLegend
`func (o *TrendsFilter) UnsetShowLegend()`

UnsetShowLegend ensures that no value is present for ShowLegend, not even an explicit nil
### GetShowMovingAverage

`func (o *TrendsFilter) GetShowMovingAverage() bool`

GetShowMovingAverage returns the ShowMovingAverage field if non-nil, zero value otherwise.

### GetShowMovingAverageOk

`func (o *TrendsFilter) GetShowMovingAverageOk() (*bool, bool)`

GetShowMovingAverageOk returns a tuple with the ShowMovingAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMovingAverage

`func (o *TrendsFilter) SetShowMovingAverage(v bool)`

SetShowMovingAverage sets ShowMovingAverage field to given value.

### HasShowMovingAverage

`func (o *TrendsFilter) HasShowMovingAverage() bool`

HasShowMovingAverage returns a boolean if a field has been set.

### SetShowMovingAverageNil

`func (o *TrendsFilter) SetShowMovingAverageNil(b bool)`

 SetShowMovingAverageNil sets the value for ShowMovingAverage to be an explicit nil

### UnsetShowMovingAverage
`func (o *TrendsFilter) UnsetShowMovingAverage()`

UnsetShowMovingAverage ensures that no value is present for ShowMovingAverage, not even an explicit nil
### GetShowMultipleYAxes

`func (o *TrendsFilter) GetShowMultipleYAxes() bool`

GetShowMultipleYAxes returns the ShowMultipleYAxes field if non-nil, zero value otherwise.

### GetShowMultipleYAxesOk

`func (o *TrendsFilter) GetShowMultipleYAxesOk() (*bool, bool)`

GetShowMultipleYAxesOk returns a tuple with the ShowMultipleYAxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMultipleYAxes

`func (o *TrendsFilter) SetShowMultipleYAxes(v bool)`

SetShowMultipleYAxes sets ShowMultipleYAxes field to given value.

### HasShowMultipleYAxes

`func (o *TrendsFilter) HasShowMultipleYAxes() bool`

HasShowMultipleYAxes returns a boolean if a field has been set.

### SetShowMultipleYAxesNil

`func (o *TrendsFilter) SetShowMultipleYAxesNil(b bool)`

 SetShowMultipleYAxesNil sets the value for ShowMultipleYAxes to be an explicit nil

### UnsetShowMultipleYAxes
`func (o *TrendsFilter) UnsetShowMultipleYAxes()`

UnsetShowMultipleYAxes ensures that no value is present for ShowMultipleYAxes, not even an explicit nil
### GetShowPercentStackView

`func (o *TrendsFilter) GetShowPercentStackView() bool`

GetShowPercentStackView returns the ShowPercentStackView field if non-nil, zero value otherwise.

### GetShowPercentStackViewOk

`func (o *TrendsFilter) GetShowPercentStackViewOk() (*bool, bool)`

GetShowPercentStackViewOk returns a tuple with the ShowPercentStackView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPercentStackView

`func (o *TrendsFilter) SetShowPercentStackView(v bool)`

SetShowPercentStackView sets ShowPercentStackView field to given value.

### HasShowPercentStackView

`func (o *TrendsFilter) HasShowPercentStackView() bool`

HasShowPercentStackView returns a boolean if a field has been set.

### SetShowPercentStackViewNil

`func (o *TrendsFilter) SetShowPercentStackViewNil(b bool)`

 SetShowPercentStackViewNil sets the value for ShowPercentStackView to be an explicit nil

### UnsetShowPercentStackView
`func (o *TrendsFilter) UnsetShowPercentStackView()`

UnsetShowPercentStackView ensures that no value is present for ShowPercentStackView, not even an explicit nil
### GetShowTrendLines

`func (o *TrendsFilter) GetShowTrendLines() bool`

GetShowTrendLines returns the ShowTrendLines field if non-nil, zero value otherwise.

### GetShowTrendLinesOk

`func (o *TrendsFilter) GetShowTrendLinesOk() (*bool, bool)`

GetShowTrendLinesOk returns a tuple with the ShowTrendLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTrendLines

`func (o *TrendsFilter) SetShowTrendLines(v bool)`

SetShowTrendLines sets ShowTrendLines field to given value.

### HasShowTrendLines

`func (o *TrendsFilter) HasShowTrendLines() bool`

HasShowTrendLines returns a boolean if a field has been set.

### SetShowTrendLinesNil

`func (o *TrendsFilter) SetShowTrendLinesNil(b bool)`

 SetShowTrendLinesNil sets the value for ShowTrendLines to be an explicit nil

### UnsetShowTrendLines
`func (o *TrendsFilter) UnsetShowTrendLines()`

UnsetShowTrendLines ensures that no value is present for ShowTrendLines, not even an explicit nil
### GetShowValuesOnSeries

`func (o *TrendsFilter) GetShowValuesOnSeries() bool`

GetShowValuesOnSeries returns the ShowValuesOnSeries field if non-nil, zero value otherwise.

### GetShowValuesOnSeriesOk

`func (o *TrendsFilter) GetShowValuesOnSeriesOk() (*bool, bool)`

GetShowValuesOnSeriesOk returns a tuple with the ShowValuesOnSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowValuesOnSeries

`func (o *TrendsFilter) SetShowValuesOnSeries(v bool)`

SetShowValuesOnSeries sets ShowValuesOnSeries field to given value.

### HasShowValuesOnSeries

`func (o *TrendsFilter) HasShowValuesOnSeries() bool`

HasShowValuesOnSeries returns a boolean if a field has been set.

### SetShowValuesOnSeriesNil

`func (o *TrendsFilter) SetShowValuesOnSeriesNil(b bool)`

 SetShowValuesOnSeriesNil sets the value for ShowValuesOnSeries to be an explicit nil

### UnsetShowValuesOnSeries
`func (o *TrendsFilter) UnsetShowValuesOnSeries()`

UnsetShowValuesOnSeries ensures that no value is present for ShowValuesOnSeries, not even an explicit nil
### GetSmoothingIntervals

`func (o *TrendsFilter) GetSmoothingIntervals() int32`

GetSmoothingIntervals returns the SmoothingIntervals field if non-nil, zero value otherwise.

### GetSmoothingIntervalsOk

`func (o *TrendsFilter) GetSmoothingIntervalsOk() (*int32, bool)`

GetSmoothingIntervalsOk returns a tuple with the SmoothingIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothingIntervals

`func (o *TrendsFilter) SetSmoothingIntervals(v int32)`

SetSmoothingIntervals sets SmoothingIntervals field to given value.

### HasSmoothingIntervals

`func (o *TrendsFilter) HasSmoothingIntervals() bool`

HasSmoothingIntervals returns a boolean if a field has been set.

### SetSmoothingIntervalsNil

`func (o *TrendsFilter) SetSmoothingIntervalsNil(b bool)`

 SetSmoothingIntervalsNil sets the value for SmoothingIntervals to be an explicit nil

### UnsetSmoothingIntervals
`func (o *TrendsFilter) UnsetSmoothingIntervals()`

UnsetSmoothingIntervals ensures that no value is present for SmoothingIntervals, not even an explicit nil
### GetYAxisScaleType

`func (o *TrendsFilter) GetYAxisScaleType() YAxisScaleType`

GetYAxisScaleType returns the YAxisScaleType field if non-nil, zero value otherwise.

### GetYAxisScaleTypeOk

`func (o *TrendsFilter) GetYAxisScaleTypeOk() (*YAxisScaleType, bool)`

GetYAxisScaleTypeOk returns a tuple with the YAxisScaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisScaleType

`func (o *TrendsFilter) SetYAxisScaleType(v YAxisScaleType)`

SetYAxisScaleType sets YAxisScaleType field to given value.

### HasYAxisScaleType

`func (o *TrendsFilter) HasYAxisScaleType() bool`

HasYAxisScaleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


