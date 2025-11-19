# ExperimentBreakdownResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Baseline** | [**ExperimentStatsBaseValidated**](ExperimentStatsBaseValidated.md) |  | 
**BreakdownValue** | [**[]BreakdownValueInner**](BreakdownValueInner.md) | The breakdown values as an array (e.g., [\&quot;MacOS\&quot;, \&quot;Chrome\&quot;] for multi-breakdown, [\&quot;Chrome\&quot;] for single) | 
**Variants** | [**Variants**](Variants.md) |  | 

## Methods

### NewExperimentBreakdownResult

`func NewExperimentBreakdownResult(baseline ExperimentStatsBaseValidated, breakdownValue []BreakdownValueInner, variants Variants, ) *ExperimentBreakdownResult`

NewExperimentBreakdownResult instantiates a new ExperimentBreakdownResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentBreakdownResultWithDefaults

`func NewExperimentBreakdownResultWithDefaults() *ExperimentBreakdownResult`

NewExperimentBreakdownResultWithDefaults instantiates a new ExperimentBreakdownResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseline

`func (o *ExperimentBreakdownResult) GetBaseline() ExperimentStatsBaseValidated`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *ExperimentBreakdownResult) GetBaselineOk() (*ExperimentStatsBaseValidated, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *ExperimentBreakdownResult) SetBaseline(v ExperimentStatsBaseValidated)`

SetBaseline sets Baseline field to given value.


### GetBreakdownValue

`func (o *ExperimentBreakdownResult) GetBreakdownValue() []BreakdownValueInner`

GetBreakdownValue returns the BreakdownValue field if non-nil, zero value otherwise.

### GetBreakdownValueOk

`func (o *ExperimentBreakdownResult) GetBreakdownValueOk() (*[]BreakdownValueInner, bool)`

GetBreakdownValueOk returns a tuple with the BreakdownValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownValue

`func (o *ExperimentBreakdownResult) SetBreakdownValue(v []BreakdownValueInner)`

SetBreakdownValue sets BreakdownValue field to given value.


### GetVariants

`func (o *ExperimentBreakdownResult) GetVariants() Variants`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentBreakdownResult) GetVariantsOk() (*Variants, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentBreakdownResult) SetVariants(v Variants)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


