# ExperimentQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Baseline** | Pointer to [**ExperimentStatsBaseValidated**](ExperimentStatsBaseValidated.md) |  | [optional] 
**BreakdownResults** | Pointer to [**[]ExperimentBreakdownResult**](ExperimentBreakdownResult.md) | Results grouped by breakdown value. When present, baseline and variant_results contain aggregated data. | [optional] 
**CredibleIntervals** | Pointer to **map[string][]float32** |  | [optional] 
**Insight** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentQuery"]
**Metric** | Pointer to [**NullableMetric1**](Metric1.md) |  | [optional] [default to null]
**PValue** | Pointer to **NullableFloat32** |  | [optional] 
**Probability** | Pointer to **map[string]float32** |  | [optional] 
**SignificanceCode** | Pointer to [**ExperimentSignificanceCode**](ExperimentSignificanceCode.md) |  | [optional] 
**Significant** | Pointer to **NullableBool** |  | [optional] 
**StatsVersion** | Pointer to **NullableInt32** |  | [optional] 
**VariantResults** | Pointer to [**NullableVariantResults**](VariantResults.md) |  | [optional] [default to null]
**Variants** | Pointer to [**NullableVariants1**](Variants1.md) |  | [optional] [default to null]

## Methods

### NewExperimentQueryResponse

`func NewExperimentQueryResponse() *ExperimentQueryResponse`

NewExperimentQueryResponse instantiates a new ExperimentQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentQueryResponseWithDefaults

`func NewExperimentQueryResponseWithDefaults() *ExperimentQueryResponse`

NewExperimentQueryResponseWithDefaults instantiates a new ExperimentQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseline

`func (o *ExperimentQueryResponse) GetBaseline() ExperimentStatsBaseValidated`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *ExperimentQueryResponse) GetBaselineOk() (*ExperimentStatsBaseValidated, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *ExperimentQueryResponse) SetBaseline(v ExperimentStatsBaseValidated)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *ExperimentQueryResponse) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetBreakdownResults

`func (o *ExperimentQueryResponse) GetBreakdownResults() []ExperimentBreakdownResult`

GetBreakdownResults returns the BreakdownResults field if non-nil, zero value otherwise.

### GetBreakdownResultsOk

`func (o *ExperimentQueryResponse) GetBreakdownResultsOk() (*[]ExperimentBreakdownResult, bool)`

GetBreakdownResultsOk returns a tuple with the BreakdownResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownResults

`func (o *ExperimentQueryResponse) SetBreakdownResults(v []ExperimentBreakdownResult)`

SetBreakdownResults sets BreakdownResults field to given value.

### HasBreakdownResults

`func (o *ExperimentQueryResponse) HasBreakdownResults() bool`

HasBreakdownResults returns a boolean if a field has been set.

### SetBreakdownResultsNil

`func (o *ExperimentQueryResponse) SetBreakdownResultsNil(b bool)`

 SetBreakdownResultsNil sets the value for BreakdownResults to be an explicit nil

### UnsetBreakdownResults
`func (o *ExperimentQueryResponse) UnsetBreakdownResults()`

UnsetBreakdownResults ensures that no value is present for BreakdownResults, not even an explicit nil
### GetCredibleIntervals

`func (o *ExperimentQueryResponse) GetCredibleIntervals() map[string][]float32`

GetCredibleIntervals returns the CredibleIntervals field if non-nil, zero value otherwise.

### GetCredibleIntervalsOk

`func (o *ExperimentQueryResponse) GetCredibleIntervalsOk() (*map[string][]float32, bool)`

GetCredibleIntervalsOk returns a tuple with the CredibleIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleIntervals

`func (o *ExperimentQueryResponse) SetCredibleIntervals(v map[string][]float32)`

SetCredibleIntervals sets CredibleIntervals field to given value.

### HasCredibleIntervals

`func (o *ExperimentQueryResponse) HasCredibleIntervals() bool`

HasCredibleIntervals returns a boolean if a field has been set.

### SetCredibleIntervalsNil

`func (o *ExperimentQueryResponse) SetCredibleIntervalsNil(b bool)`

 SetCredibleIntervalsNil sets the value for CredibleIntervals to be an explicit nil

### UnsetCredibleIntervals
`func (o *ExperimentQueryResponse) UnsetCredibleIntervals()`

UnsetCredibleIntervals ensures that no value is present for CredibleIntervals, not even an explicit nil
### GetInsight

`func (o *ExperimentQueryResponse) GetInsight() []map[string]interface{}`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *ExperimentQueryResponse) GetInsightOk() (*[]map[string]interface{}, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *ExperimentQueryResponse) SetInsight(v []map[string]interface{})`

SetInsight sets Insight field to given value.

### HasInsight

`func (o *ExperimentQueryResponse) HasInsight() bool`

HasInsight returns a boolean if a field has been set.

### SetInsightNil

`func (o *ExperimentQueryResponse) SetInsightNil(b bool)`

 SetInsightNil sets the value for Insight to be an explicit nil

### UnsetInsight
`func (o *ExperimentQueryResponse) UnsetInsight()`

UnsetInsight ensures that no value is present for Insight, not even an explicit nil
### GetKind

`func (o *ExperimentQueryResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentQueryResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentQueryResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentQueryResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetric

`func (o *ExperimentQueryResponse) GetMetric() Metric1`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ExperimentQueryResponse) GetMetricOk() (*Metric1, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ExperimentQueryResponse) SetMetric(v Metric1)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ExperimentQueryResponse) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *ExperimentQueryResponse) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *ExperimentQueryResponse) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil
### GetPValue

`func (o *ExperimentQueryResponse) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *ExperimentQueryResponse) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *ExperimentQueryResponse) SetPValue(v float32)`

SetPValue sets PValue field to given value.

### HasPValue

`func (o *ExperimentQueryResponse) HasPValue() bool`

HasPValue returns a boolean if a field has been set.

### SetPValueNil

`func (o *ExperimentQueryResponse) SetPValueNil(b bool)`

 SetPValueNil sets the value for PValue to be an explicit nil

### UnsetPValue
`func (o *ExperimentQueryResponse) UnsetPValue()`

UnsetPValue ensures that no value is present for PValue, not even an explicit nil
### GetProbability

`func (o *ExperimentQueryResponse) GetProbability() map[string]float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *ExperimentQueryResponse) GetProbabilityOk() (*map[string]float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *ExperimentQueryResponse) SetProbability(v map[string]float32)`

SetProbability sets Probability field to given value.

### HasProbability

`func (o *ExperimentQueryResponse) HasProbability() bool`

HasProbability returns a boolean if a field has been set.

### SetProbabilityNil

`func (o *ExperimentQueryResponse) SetProbabilityNil(b bool)`

 SetProbabilityNil sets the value for Probability to be an explicit nil

### UnsetProbability
`func (o *ExperimentQueryResponse) UnsetProbability()`

UnsetProbability ensures that no value is present for Probability, not even an explicit nil
### GetSignificanceCode

`func (o *ExperimentQueryResponse) GetSignificanceCode() ExperimentSignificanceCode`

GetSignificanceCode returns the SignificanceCode field if non-nil, zero value otherwise.

### GetSignificanceCodeOk

`func (o *ExperimentQueryResponse) GetSignificanceCodeOk() (*ExperimentSignificanceCode, bool)`

GetSignificanceCodeOk returns a tuple with the SignificanceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceCode

`func (o *ExperimentQueryResponse) SetSignificanceCode(v ExperimentSignificanceCode)`

SetSignificanceCode sets SignificanceCode field to given value.

### HasSignificanceCode

`func (o *ExperimentQueryResponse) HasSignificanceCode() bool`

HasSignificanceCode returns a boolean if a field has been set.

### GetSignificant

`func (o *ExperimentQueryResponse) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *ExperimentQueryResponse) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *ExperimentQueryResponse) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.

### HasSignificant

`func (o *ExperimentQueryResponse) HasSignificant() bool`

HasSignificant returns a boolean if a field has been set.

### SetSignificantNil

`func (o *ExperimentQueryResponse) SetSignificantNil(b bool)`

 SetSignificantNil sets the value for Significant to be an explicit nil

### UnsetSignificant
`func (o *ExperimentQueryResponse) UnsetSignificant()`

UnsetSignificant ensures that no value is present for Significant, not even an explicit nil
### GetStatsVersion

`func (o *ExperimentQueryResponse) GetStatsVersion() int32`

GetStatsVersion returns the StatsVersion field if non-nil, zero value otherwise.

### GetStatsVersionOk

`func (o *ExperimentQueryResponse) GetStatsVersionOk() (*int32, bool)`

GetStatsVersionOk returns a tuple with the StatsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersion

`func (o *ExperimentQueryResponse) SetStatsVersion(v int32)`

SetStatsVersion sets StatsVersion field to given value.

### HasStatsVersion

`func (o *ExperimentQueryResponse) HasStatsVersion() bool`

HasStatsVersion returns a boolean if a field has been set.

### SetStatsVersionNil

`func (o *ExperimentQueryResponse) SetStatsVersionNil(b bool)`

 SetStatsVersionNil sets the value for StatsVersion to be an explicit nil

### UnsetStatsVersion
`func (o *ExperimentQueryResponse) UnsetStatsVersion()`

UnsetStatsVersion ensures that no value is present for StatsVersion, not even an explicit nil
### GetVariantResults

`func (o *ExperimentQueryResponse) GetVariantResults() VariantResults`

GetVariantResults returns the VariantResults field if non-nil, zero value otherwise.

### GetVariantResultsOk

`func (o *ExperimentQueryResponse) GetVariantResultsOk() (*VariantResults, bool)`

GetVariantResultsOk returns a tuple with the VariantResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantResults

`func (o *ExperimentQueryResponse) SetVariantResults(v VariantResults)`

SetVariantResults sets VariantResults field to given value.

### HasVariantResults

`func (o *ExperimentQueryResponse) HasVariantResults() bool`

HasVariantResults returns a boolean if a field has been set.

### SetVariantResultsNil

`func (o *ExperimentQueryResponse) SetVariantResultsNil(b bool)`

 SetVariantResultsNil sets the value for VariantResults to be an explicit nil

### UnsetVariantResults
`func (o *ExperimentQueryResponse) UnsetVariantResults()`

UnsetVariantResults ensures that no value is present for VariantResults, not even an explicit nil
### GetVariants

`func (o *ExperimentQueryResponse) GetVariants() Variants1`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentQueryResponse) GetVariantsOk() (*Variants1, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentQueryResponse) SetVariants(v Variants1)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ExperimentQueryResponse) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariantsNil

`func (o *ExperimentQueryResponse) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *ExperimentQueryResponse) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


