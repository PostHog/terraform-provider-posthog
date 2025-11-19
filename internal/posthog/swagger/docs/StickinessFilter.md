# StickinessFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputedAs** | Pointer to [**StickinessComputationMode**](StickinessComputationMode.md) |  | [optional] 
**Display** | Pointer to [**ChartDisplayType**](ChartDisplayType.md) |  | [optional] 
**HiddenLegendIndexes** | Pointer to **[]int32** |  | [optional] 
**ResultCustomizationBy** | Pointer to [**ResultCustomizationBy**](ResultCustomizationBy.md) |  | [optional] 
**ResultCustomizations** | Pointer to [**NullableResultcustomizations**](Resultcustomizations.md) |  | [optional] [default to null]
**ShowLegend** | Pointer to **NullableBool** |  | [optional] 
**ShowMultipleYAxes** | Pointer to **NullableBool** |  | [optional] 
**ShowValuesOnSeries** | Pointer to **NullableBool** |  | [optional] 
**StickinessCriteria** | Pointer to [**StickinessCriteria**](StickinessCriteria.md) |  | [optional] 

## Methods

### NewStickinessFilter

`func NewStickinessFilter() *StickinessFilter`

NewStickinessFilter instantiates a new StickinessFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStickinessFilterWithDefaults

`func NewStickinessFilterWithDefaults() *StickinessFilter`

NewStickinessFilterWithDefaults instantiates a new StickinessFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputedAs

`func (o *StickinessFilter) GetComputedAs() StickinessComputationMode`

GetComputedAs returns the ComputedAs field if non-nil, zero value otherwise.

### GetComputedAsOk

`func (o *StickinessFilter) GetComputedAsOk() (*StickinessComputationMode, bool)`

GetComputedAsOk returns a tuple with the ComputedAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedAs

`func (o *StickinessFilter) SetComputedAs(v StickinessComputationMode)`

SetComputedAs sets ComputedAs field to given value.

### HasComputedAs

`func (o *StickinessFilter) HasComputedAs() bool`

HasComputedAs returns a boolean if a field has been set.

### GetDisplay

`func (o *StickinessFilter) GetDisplay() ChartDisplayType`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *StickinessFilter) GetDisplayOk() (*ChartDisplayType, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *StickinessFilter) SetDisplay(v ChartDisplayType)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *StickinessFilter) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetHiddenLegendIndexes

`func (o *StickinessFilter) GetHiddenLegendIndexes() []int32`

GetHiddenLegendIndexes returns the HiddenLegendIndexes field if non-nil, zero value otherwise.

### GetHiddenLegendIndexesOk

`func (o *StickinessFilter) GetHiddenLegendIndexesOk() (*[]int32, bool)`

GetHiddenLegendIndexesOk returns a tuple with the HiddenLegendIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenLegendIndexes

`func (o *StickinessFilter) SetHiddenLegendIndexes(v []int32)`

SetHiddenLegendIndexes sets HiddenLegendIndexes field to given value.

### HasHiddenLegendIndexes

`func (o *StickinessFilter) HasHiddenLegendIndexes() bool`

HasHiddenLegendIndexes returns a boolean if a field has been set.

### SetHiddenLegendIndexesNil

`func (o *StickinessFilter) SetHiddenLegendIndexesNil(b bool)`

 SetHiddenLegendIndexesNil sets the value for HiddenLegendIndexes to be an explicit nil

### UnsetHiddenLegendIndexes
`func (o *StickinessFilter) UnsetHiddenLegendIndexes()`

UnsetHiddenLegendIndexes ensures that no value is present for HiddenLegendIndexes, not even an explicit nil
### GetResultCustomizationBy

`func (o *StickinessFilter) GetResultCustomizationBy() ResultCustomizationBy`

GetResultCustomizationBy returns the ResultCustomizationBy field if non-nil, zero value otherwise.

### GetResultCustomizationByOk

`func (o *StickinessFilter) GetResultCustomizationByOk() (*ResultCustomizationBy, bool)`

GetResultCustomizationByOk returns a tuple with the ResultCustomizationBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCustomizationBy

`func (o *StickinessFilter) SetResultCustomizationBy(v ResultCustomizationBy)`

SetResultCustomizationBy sets ResultCustomizationBy field to given value.

### HasResultCustomizationBy

`func (o *StickinessFilter) HasResultCustomizationBy() bool`

HasResultCustomizationBy returns a boolean if a field has been set.

### GetResultCustomizations

`func (o *StickinessFilter) GetResultCustomizations() Resultcustomizations`

GetResultCustomizations returns the ResultCustomizations field if non-nil, zero value otherwise.

### GetResultCustomizationsOk

`func (o *StickinessFilter) GetResultCustomizationsOk() (*Resultcustomizations, bool)`

GetResultCustomizationsOk returns a tuple with the ResultCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCustomizations

`func (o *StickinessFilter) SetResultCustomizations(v Resultcustomizations)`

SetResultCustomizations sets ResultCustomizations field to given value.

### HasResultCustomizations

`func (o *StickinessFilter) HasResultCustomizations() bool`

HasResultCustomizations returns a boolean if a field has been set.

### SetResultCustomizationsNil

`func (o *StickinessFilter) SetResultCustomizationsNil(b bool)`

 SetResultCustomizationsNil sets the value for ResultCustomizations to be an explicit nil

### UnsetResultCustomizations
`func (o *StickinessFilter) UnsetResultCustomizations()`

UnsetResultCustomizations ensures that no value is present for ResultCustomizations, not even an explicit nil
### GetShowLegend

`func (o *StickinessFilter) GetShowLegend() bool`

GetShowLegend returns the ShowLegend field if non-nil, zero value otherwise.

### GetShowLegendOk

`func (o *StickinessFilter) GetShowLegendOk() (*bool, bool)`

GetShowLegendOk returns a tuple with the ShowLegend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLegend

`func (o *StickinessFilter) SetShowLegend(v bool)`

SetShowLegend sets ShowLegend field to given value.

### HasShowLegend

`func (o *StickinessFilter) HasShowLegend() bool`

HasShowLegend returns a boolean if a field has been set.

### SetShowLegendNil

`func (o *StickinessFilter) SetShowLegendNil(b bool)`

 SetShowLegendNil sets the value for ShowLegend to be an explicit nil

### UnsetShowLegend
`func (o *StickinessFilter) UnsetShowLegend()`

UnsetShowLegend ensures that no value is present for ShowLegend, not even an explicit nil
### GetShowMultipleYAxes

`func (o *StickinessFilter) GetShowMultipleYAxes() bool`

GetShowMultipleYAxes returns the ShowMultipleYAxes field if non-nil, zero value otherwise.

### GetShowMultipleYAxesOk

`func (o *StickinessFilter) GetShowMultipleYAxesOk() (*bool, bool)`

GetShowMultipleYAxesOk returns a tuple with the ShowMultipleYAxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMultipleYAxes

`func (o *StickinessFilter) SetShowMultipleYAxes(v bool)`

SetShowMultipleYAxes sets ShowMultipleYAxes field to given value.

### HasShowMultipleYAxes

`func (o *StickinessFilter) HasShowMultipleYAxes() bool`

HasShowMultipleYAxes returns a boolean if a field has been set.

### SetShowMultipleYAxesNil

`func (o *StickinessFilter) SetShowMultipleYAxesNil(b bool)`

 SetShowMultipleYAxesNil sets the value for ShowMultipleYAxes to be an explicit nil

### UnsetShowMultipleYAxes
`func (o *StickinessFilter) UnsetShowMultipleYAxes()`

UnsetShowMultipleYAxes ensures that no value is present for ShowMultipleYAxes, not even an explicit nil
### GetShowValuesOnSeries

`func (o *StickinessFilter) GetShowValuesOnSeries() bool`

GetShowValuesOnSeries returns the ShowValuesOnSeries field if non-nil, zero value otherwise.

### GetShowValuesOnSeriesOk

`func (o *StickinessFilter) GetShowValuesOnSeriesOk() (*bool, bool)`

GetShowValuesOnSeriesOk returns a tuple with the ShowValuesOnSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowValuesOnSeries

`func (o *StickinessFilter) SetShowValuesOnSeries(v bool)`

SetShowValuesOnSeries sets ShowValuesOnSeries field to given value.

### HasShowValuesOnSeries

`func (o *StickinessFilter) HasShowValuesOnSeries() bool`

HasShowValuesOnSeries returns a boolean if a field has been set.

### SetShowValuesOnSeriesNil

`func (o *StickinessFilter) SetShowValuesOnSeriesNil(b bool)`

 SetShowValuesOnSeriesNil sets the value for ShowValuesOnSeries to be an explicit nil

### UnsetShowValuesOnSeries
`func (o *StickinessFilter) UnsetShowValuesOnSeries()`

UnsetShowValuesOnSeries ensures that no value is present for ShowValuesOnSeries, not even an explicit nil
### GetStickinessCriteria

`func (o *StickinessFilter) GetStickinessCriteria() StickinessCriteria`

GetStickinessCriteria returns the StickinessCriteria field if non-nil, zero value otherwise.

### GetStickinessCriteriaOk

`func (o *StickinessFilter) GetStickinessCriteriaOk() (*StickinessCriteria, bool)`

GetStickinessCriteriaOk returns a tuple with the StickinessCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickinessCriteria

`func (o *StickinessFilter) SetStickinessCriteria(v StickinessCriteria)`

SetStickinessCriteria sets StickinessCriteria field to given value.

### HasStickinessCriteria

`func (o *StickinessFilter) HasStickinessCriteria() bool`

HasStickinessCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


