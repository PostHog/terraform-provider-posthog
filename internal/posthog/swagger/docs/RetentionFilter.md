# RetentionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cumulative** | Pointer to **NullableBool** |  | [optional] 
**DashboardDisplay** | Pointer to [**RetentionDashboardDisplayType**](RetentionDashboardDisplayType.md) |  | [optional] 
**Display** | Pointer to [**ChartDisplayType**](ChartDisplayType.md) |  | [optional] 
**MeanRetentionCalculation** | Pointer to [**MeanRetentionCalculation**](MeanRetentionCalculation.md) |  | [optional] 
**MinimumOccurrences** | Pointer to **NullableInt32** |  | [optional] 
**Period** | Pointer to [**RetentionPeriod**](RetentionPeriod.md) |  | [optional] 
**RetentionCustomBrackets** | Pointer to **[]float32** | Custom brackets for retention calculations | [optional] 
**RetentionReference** | Pointer to [**RetentionReference**](RetentionReference.md) |  | [optional] 
**RetentionType** | Pointer to [**RetentionType**](RetentionType.md) |  | [optional] 
**ReturningEntity** | Pointer to [**RetentionEntity**](RetentionEntity.md) |  | [optional] 
**SelectedInterval** | Pointer to **NullableInt32** | The selected interval to display across all cohorts (null &#x3D; show all intervals for each cohort) | [optional] 
**ShowTrendLines** | Pointer to **NullableBool** |  | [optional] 
**TargetEntity** | Pointer to [**RetentionEntity**](RetentionEntity.md) |  | [optional] 
**TimeWindowMode** | Pointer to [**TimeWindowMode**](TimeWindowMode.md) |  | [optional] 
**TotalIntervals** | Pointer to **NullableInt32** |  | [optional] [default to 8]

## Methods

### NewRetentionFilter

`func NewRetentionFilter() *RetentionFilter`

NewRetentionFilter instantiates a new RetentionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionFilterWithDefaults

`func NewRetentionFilterWithDefaults() *RetentionFilter`

NewRetentionFilterWithDefaults instantiates a new RetentionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCumulative

`func (o *RetentionFilter) GetCumulative() bool`

GetCumulative returns the Cumulative field if non-nil, zero value otherwise.

### GetCumulativeOk

`func (o *RetentionFilter) GetCumulativeOk() (*bool, bool)`

GetCumulativeOk returns a tuple with the Cumulative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulative

`func (o *RetentionFilter) SetCumulative(v bool)`

SetCumulative sets Cumulative field to given value.

### HasCumulative

`func (o *RetentionFilter) HasCumulative() bool`

HasCumulative returns a boolean if a field has been set.

### SetCumulativeNil

`func (o *RetentionFilter) SetCumulativeNil(b bool)`

 SetCumulativeNil sets the value for Cumulative to be an explicit nil

### UnsetCumulative
`func (o *RetentionFilter) UnsetCumulative()`

UnsetCumulative ensures that no value is present for Cumulative, not even an explicit nil
### GetDashboardDisplay

`func (o *RetentionFilter) GetDashboardDisplay() RetentionDashboardDisplayType`

GetDashboardDisplay returns the DashboardDisplay field if non-nil, zero value otherwise.

### GetDashboardDisplayOk

`func (o *RetentionFilter) GetDashboardDisplayOk() (*RetentionDashboardDisplayType, bool)`

GetDashboardDisplayOk returns a tuple with the DashboardDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardDisplay

`func (o *RetentionFilter) SetDashboardDisplay(v RetentionDashboardDisplayType)`

SetDashboardDisplay sets DashboardDisplay field to given value.

### HasDashboardDisplay

`func (o *RetentionFilter) HasDashboardDisplay() bool`

HasDashboardDisplay returns a boolean if a field has been set.

### GetDisplay

`func (o *RetentionFilter) GetDisplay() ChartDisplayType`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RetentionFilter) GetDisplayOk() (*ChartDisplayType, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RetentionFilter) SetDisplay(v ChartDisplayType)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *RetentionFilter) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetMeanRetentionCalculation

`func (o *RetentionFilter) GetMeanRetentionCalculation() MeanRetentionCalculation`

GetMeanRetentionCalculation returns the MeanRetentionCalculation field if non-nil, zero value otherwise.

### GetMeanRetentionCalculationOk

`func (o *RetentionFilter) GetMeanRetentionCalculationOk() (*MeanRetentionCalculation, bool)`

GetMeanRetentionCalculationOk returns a tuple with the MeanRetentionCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeanRetentionCalculation

`func (o *RetentionFilter) SetMeanRetentionCalculation(v MeanRetentionCalculation)`

SetMeanRetentionCalculation sets MeanRetentionCalculation field to given value.

### HasMeanRetentionCalculation

`func (o *RetentionFilter) HasMeanRetentionCalculation() bool`

HasMeanRetentionCalculation returns a boolean if a field has been set.

### GetMinimumOccurrences

`func (o *RetentionFilter) GetMinimumOccurrences() int32`

GetMinimumOccurrences returns the MinimumOccurrences field if non-nil, zero value otherwise.

### GetMinimumOccurrencesOk

`func (o *RetentionFilter) GetMinimumOccurrencesOk() (*int32, bool)`

GetMinimumOccurrencesOk returns a tuple with the MinimumOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumOccurrences

`func (o *RetentionFilter) SetMinimumOccurrences(v int32)`

SetMinimumOccurrences sets MinimumOccurrences field to given value.

### HasMinimumOccurrences

`func (o *RetentionFilter) HasMinimumOccurrences() bool`

HasMinimumOccurrences returns a boolean if a field has been set.

### SetMinimumOccurrencesNil

`func (o *RetentionFilter) SetMinimumOccurrencesNil(b bool)`

 SetMinimumOccurrencesNil sets the value for MinimumOccurrences to be an explicit nil

### UnsetMinimumOccurrences
`func (o *RetentionFilter) UnsetMinimumOccurrences()`

UnsetMinimumOccurrences ensures that no value is present for MinimumOccurrences, not even an explicit nil
### GetPeriod

`func (o *RetentionFilter) GetPeriod() RetentionPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RetentionFilter) GetPeriodOk() (*RetentionPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RetentionFilter) SetPeriod(v RetentionPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RetentionFilter) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRetentionCustomBrackets

`func (o *RetentionFilter) GetRetentionCustomBrackets() []float32`

GetRetentionCustomBrackets returns the RetentionCustomBrackets field if non-nil, zero value otherwise.

### GetRetentionCustomBracketsOk

`func (o *RetentionFilter) GetRetentionCustomBracketsOk() (*[]float32, bool)`

GetRetentionCustomBracketsOk returns a tuple with the RetentionCustomBrackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCustomBrackets

`func (o *RetentionFilter) SetRetentionCustomBrackets(v []float32)`

SetRetentionCustomBrackets sets RetentionCustomBrackets field to given value.

### HasRetentionCustomBrackets

`func (o *RetentionFilter) HasRetentionCustomBrackets() bool`

HasRetentionCustomBrackets returns a boolean if a field has been set.

### SetRetentionCustomBracketsNil

`func (o *RetentionFilter) SetRetentionCustomBracketsNil(b bool)`

 SetRetentionCustomBracketsNil sets the value for RetentionCustomBrackets to be an explicit nil

### UnsetRetentionCustomBrackets
`func (o *RetentionFilter) UnsetRetentionCustomBrackets()`

UnsetRetentionCustomBrackets ensures that no value is present for RetentionCustomBrackets, not even an explicit nil
### GetRetentionReference

`func (o *RetentionFilter) GetRetentionReference() RetentionReference`

GetRetentionReference returns the RetentionReference field if non-nil, zero value otherwise.

### GetRetentionReferenceOk

`func (o *RetentionFilter) GetRetentionReferenceOk() (*RetentionReference, bool)`

GetRetentionReferenceOk returns a tuple with the RetentionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionReference

`func (o *RetentionFilter) SetRetentionReference(v RetentionReference)`

SetRetentionReference sets RetentionReference field to given value.

### HasRetentionReference

`func (o *RetentionFilter) HasRetentionReference() bool`

HasRetentionReference returns a boolean if a field has been set.

### GetRetentionType

`func (o *RetentionFilter) GetRetentionType() RetentionType`

GetRetentionType returns the RetentionType field if non-nil, zero value otherwise.

### GetRetentionTypeOk

`func (o *RetentionFilter) GetRetentionTypeOk() (*RetentionType, bool)`

GetRetentionTypeOk returns a tuple with the RetentionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionType

`func (o *RetentionFilter) SetRetentionType(v RetentionType)`

SetRetentionType sets RetentionType field to given value.

### HasRetentionType

`func (o *RetentionFilter) HasRetentionType() bool`

HasRetentionType returns a boolean if a field has been set.

### GetReturningEntity

`func (o *RetentionFilter) GetReturningEntity() RetentionEntity`

GetReturningEntity returns the ReturningEntity field if non-nil, zero value otherwise.

### GetReturningEntityOk

`func (o *RetentionFilter) GetReturningEntityOk() (*RetentionEntity, bool)`

GetReturningEntityOk returns a tuple with the ReturningEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturningEntity

`func (o *RetentionFilter) SetReturningEntity(v RetentionEntity)`

SetReturningEntity sets ReturningEntity field to given value.

### HasReturningEntity

`func (o *RetentionFilter) HasReturningEntity() bool`

HasReturningEntity returns a boolean if a field has been set.

### GetSelectedInterval

`func (o *RetentionFilter) GetSelectedInterval() int32`

GetSelectedInterval returns the SelectedInterval field if non-nil, zero value otherwise.

### GetSelectedIntervalOk

`func (o *RetentionFilter) GetSelectedIntervalOk() (*int32, bool)`

GetSelectedIntervalOk returns a tuple with the SelectedInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedInterval

`func (o *RetentionFilter) SetSelectedInterval(v int32)`

SetSelectedInterval sets SelectedInterval field to given value.

### HasSelectedInterval

`func (o *RetentionFilter) HasSelectedInterval() bool`

HasSelectedInterval returns a boolean if a field has been set.

### SetSelectedIntervalNil

`func (o *RetentionFilter) SetSelectedIntervalNil(b bool)`

 SetSelectedIntervalNil sets the value for SelectedInterval to be an explicit nil

### UnsetSelectedInterval
`func (o *RetentionFilter) UnsetSelectedInterval()`

UnsetSelectedInterval ensures that no value is present for SelectedInterval, not even an explicit nil
### GetShowTrendLines

`func (o *RetentionFilter) GetShowTrendLines() bool`

GetShowTrendLines returns the ShowTrendLines field if non-nil, zero value otherwise.

### GetShowTrendLinesOk

`func (o *RetentionFilter) GetShowTrendLinesOk() (*bool, bool)`

GetShowTrendLinesOk returns a tuple with the ShowTrendLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTrendLines

`func (o *RetentionFilter) SetShowTrendLines(v bool)`

SetShowTrendLines sets ShowTrendLines field to given value.

### HasShowTrendLines

`func (o *RetentionFilter) HasShowTrendLines() bool`

HasShowTrendLines returns a boolean if a field has been set.

### SetShowTrendLinesNil

`func (o *RetentionFilter) SetShowTrendLinesNil(b bool)`

 SetShowTrendLinesNil sets the value for ShowTrendLines to be an explicit nil

### UnsetShowTrendLines
`func (o *RetentionFilter) UnsetShowTrendLines()`

UnsetShowTrendLines ensures that no value is present for ShowTrendLines, not even an explicit nil
### GetTargetEntity

`func (o *RetentionFilter) GetTargetEntity() RetentionEntity`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *RetentionFilter) GetTargetEntityOk() (*RetentionEntity, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *RetentionFilter) SetTargetEntity(v RetentionEntity)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *RetentionFilter) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### GetTimeWindowMode

`func (o *RetentionFilter) GetTimeWindowMode() TimeWindowMode`

GetTimeWindowMode returns the TimeWindowMode field if non-nil, zero value otherwise.

### GetTimeWindowModeOk

`func (o *RetentionFilter) GetTimeWindowModeOk() (*TimeWindowMode, bool)`

GetTimeWindowModeOk returns a tuple with the TimeWindowMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowMode

`func (o *RetentionFilter) SetTimeWindowMode(v TimeWindowMode)`

SetTimeWindowMode sets TimeWindowMode field to given value.

### HasTimeWindowMode

`func (o *RetentionFilter) HasTimeWindowMode() bool`

HasTimeWindowMode returns a boolean if a field has been set.

### GetTotalIntervals

`func (o *RetentionFilter) GetTotalIntervals() int32`

GetTotalIntervals returns the TotalIntervals field if non-nil, zero value otherwise.

### GetTotalIntervalsOk

`func (o *RetentionFilter) GetTotalIntervalsOk() (*int32, bool)`

GetTotalIntervalsOk returns a tuple with the TotalIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIntervals

`func (o *RetentionFilter) SetTotalIntervals(v int32)`

SetTotalIntervals sets TotalIntervals field to given value.

### HasTotalIntervals

`func (o *RetentionFilter) HasTotalIntervals() bool`

HasTotalIntervals returns a boolean if a field has been set.

### SetTotalIntervalsNil

`func (o *RetentionFilter) SetTotalIntervalsNil(b bool)`

 SetTotalIntervalsNil sets the value for TotalIntervals to be an explicit nil

### UnsetTotalIntervals
`func (o *RetentionFilter) UnsetTotalIntervals()`

UnsetTotalIntervals ensures that no value is present for TotalIntervals, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


