# ExperimentVariantResultFrequentist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfidenceInterval** | Pointer to **[]float32** |  | [optional] 
**DenominatorSum** | Pointer to **NullableFloat32** |  | [optional] 
**DenominatorSumSquares** | Pointer to **NullableFloat32** |  | [optional] 
**Key** | **string** |  | 
**Method** | Pointer to **string** |  | [optional] [default to "frequentist"]
**NumberOfSamples** | **int32** |  | 
**NumeratorDenominatorSumProduct** | Pointer to **NullableFloat32** |  | [optional] 
**PValue** | Pointer to **NullableFloat32** |  | [optional] 
**Significant** | Pointer to **NullableBool** |  | [optional] 
**StepCounts** | Pointer to **[]int32** |  | [optional] 
**StepSessions** | Pointer to [**[][]SessionData**]([]SessionData.md) |  | [optional] 
**Sum** | **float32** |  | 
**SumSquares** | **float32** |  | 
**ValidationFailures** | Pointer to [**[]ExperimentStatsValidationFailure**](ExperimentStatsValidationFailure.md) |  | [optional] 

## Methods

### NewExperimentVariantResultFrequentist

`func NewExperimentVariantResultFrequentist(key string, numberOfSamples int32, sum float32, sumSquares float32, ) *ExperimentVariantResultFrequentist`

NewExperimentVariantResultFrequentist instantiates a new ExperimentVariantResultFrequentist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVariantResultFrequentistWithDefaults

`func NewExperimentVariantResultFrequentistWithDefaults() *ExperimentVariantResultFrequentist`

NewExperimentVariantResultFrequentistWithDefaults instantiates a new ExperimentVariantResultFrequentist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidenceInterval

`func (o *ExperimentVariantResultFrequentist) GetConfidenceInterval() []float32`

GetConfidenceInterval returns the ConfidenceInterval field if non-nil, zero value otherwise.

### GetConfidenceIntervalOk

`func (o *ExperimentVariantResultFrequentist) GetConfidenceIntervalOk() (*[]float32, bool)`

GetConfidenceIntervalOk returns a tuple with the ConfidenceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceInterval

`func (o *ExperimentVariantResultFrequentist) SetConfidenceInterval(v []float32)`

SetConfidenceInterval sets ConfidenceInterval field to given value.

### HasConfidenceInterval

`func (o *ExperimentVariantResultFrequentist) HasConfidenceInterval() bool`

HasConfidenceInterval returns a boolean if a field has been set.

### SetConfidenceIntervalNil

`func (o *ExperimentVariantResultFrequentist) SetConfidenceIntervalNil(b bool)`

 SetConfidenceIntervalNil sets the value for ConfidenceInterval to be an explicit nil

### UnsetConfidenceInterval
`func (o *ExperimentVariantResultFrequentist) UnsetConfidenceInterval()`

UnsetConfidenceInterval ensures that no value is present for ConfidenceInterval, not even an explicit nil
### GetDenominatorSum

`func (o *ExperimentVariantResultFrequentist) GetDenominatorSum() float32`

GetDenominatorSum returns the DenominatorSum field if non-nil, zero value otherwise.

### GetDenominatorSumOk

`func (o *ExperimentVariantResultFrequentist) GetDenominatorSumOk() (*float32, bool)`

GetDenominatorSumOk returns a tuple with the DenominatorSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorSum

`func (o *ExperimentVariantResultFrequentist) SetDenominatorSum(v float32)`

SetDenominatorSum sets DenominatorSum field to given value.

### HasDenominatorSum

`func (o *ExperimentVariantResultFrequentist) HasDenominatorSum() bool`

HasDenominatorSum returns a boolean if a field has been set.

### SetDenominatorSumNil

`func (o *ExperimentVariantResultFrequentist) SetDenominatorSumNil(b bool)`

 SetDenominatorSumNil sets the value for DenominatorSum to be an explicit nil

### UnsetDenominatorSum
`func (o *ExperimentVariantResultFrequentist) UnsetDenominatorSum()`

UnsetDenominatorSum ensures that no value is present for DenominatorSum, not even an explicit nil
### GetDenominatorSumSquares

`func (o *ExperimentVariantResultFrequentist) GetDenominatorSumSquares() float32`

GetDenominatorSumSquares returns the DenominatorSumSquares field if non-nil, zero value otherwise.

### GetDenominatorSumSquaresOk

`func (o *ExperimentVariantResultFrequentist) GetDenominatorSumSquaresOk() (*float32, bool)`

GetDenominatorSumSquaresOk returns a tuple with the DenominatorSumSquares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorSumSquares

`func (o *ExperimentVariantResultFrequentist) SetDenominatorSumSquares(v float32)`

SetDenominatorSumSquares sets DenominatorSumSquares field to given value.

### HasDenominatorSumSquares

`func (o *ExperimentVariantResultFrequentist) HasDenominatorSumSquares() bool`

HasDenominatorSumSquares returns a boolean if a field has been set.

### SetDenominatorSumSquaresNil

`func (o *ExperimentVariantResultFrequentist) SetDenominatorSumSquaresNil(b bool)`

 SetDenominatorSumSquaresNil sets the value for DenominatorSumSquares to be an explicit nil

### UnsetDenominatorSumSquares
`func (o *ExperimentVariantResultFrequentist) UnsetDenominatorSumSquares()`

UnsetDenominatorSumSquares ensures that no value is present for DenominatorSumSquares, not even an explicit nil
### GetKey

`func (o *ExperimentVariantResultFrequentist) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExperimentVariantResultFrequentist) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExperimentVariantResultFrequentist) SetKey(v string)`

SetKey sets Key field to given value.


### GetMethod

`func (o *ExperimentVariantResultFrequentist) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ExperimentVariantResultFrequentist) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ExperimentVariantResultFrequentist) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ExperimentVariantResultFrequentist) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNumberOfSamples

`func (o *ExperimentVariantResultFrequentist) GetNumberOfSamples() int32`

GetNumberOfSamples returns the NumberOfSamples field if non-nil, zero value otherwise.

### GetNumberOfSamplesOk

`func (o *ExperimentVariantResultFrequentist) GetNumberOfSamplesOk() (*int32, bool)`

GetNumberOfSamplesOk returns a tuple with the NumberOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSamples

`func (o *ExperimentVariantResultFrequentist) SetNumberOfSamples(v int32)`

SetNumberOfSamples sets NumberOfSamples field to given value.


### GetNumeratorDenominatorSumProduct

`func (o *ExperimentVariantResultFrequentist) GetNumeratorDenominatorSumProduct() float32`

GetNumeratorDenominatorSumProduct returns the NumeratorDenominatorSumProduct field if non-nil, zero value otherwise.

### GetNumeratorDenominatorSumProductOk

`func (o *ExperimentVariantResultFrequentist) GetNumeratorDenominatorSumProductOk() (*float32, bool)`

GetNumeratorDenominatorSumProductOk returns a tuple with the NumeratorDenominatorSumProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeratorDenominatorSumProduct

`func (o *ExperimentVariantResultFrequentist) SetNumeratorDenominatorSumProduct(v float32)`

SetNumeratorDenominatorSumProduct sets NumeratorDenominatorSumProduct field to given value.

### HasNumeratorDenominatorSumProduct

`func (o *ExperimentVariantResultFrequentist) HasNumeratorDenominatorSumProduct() bool`

HasNumeratorDenominatorSumProduct returns a boolean if a field has been set.

### SetNumeratorDenominatorSumProductNil

`func (o *ExperimentVariantResultFrequentist) SetNumeratorDenominatorSumProductNil(b bool)`

 SetNumeratorDenominatorSumProductNil sets the value for NumeratorDenominatorSumProduct to be an explicit nil

### UnsetNumeratorDenominatorSumProduct
`func (o *ExperimentVariantResultFrequentist) UnsetNumeratorDenominatorSumProduct()`

UnsetNumeratorDenominatorSumProduct ensures that no value is present for NumeratorDenominatorSumProduct, not even an explicit nil
### GetPValue

`func (o *ExperimentVariantResultFrequentist) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *ExperimentVariantResultFrequentist) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *ExperimentVariantResultFrequentist) SetPValue(v float32)`

SetPValue sets PValue field to given value.

### HasPValue

`func (o *ExperimentVariantResultFrequentist) HasPValue() bool`

HasPValue returns a boolean if a field has been set.

### SetPValueNil

`func (o *ExperimentVariantResultFrequentist) SetPValueNil(b bool)`

 SetPValueNil sets the value for PValue to be an explicit nil

### UnsetPValue
`func (o *ExperimentVariantResultFrequentist) UnsetPValue()`

UnsetPValue ensures that no value is present for PValue, not even an explicit nil
### GetSignificant

`func (o *ExperimentVariantResultFrequentist) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *ExperimentVariantResultFrequentist) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *ExperimentVariantResultFrequentist) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.

### HasSignificant

`func (o *ExperimentVariantResultFrequentist) HasSignificant() bool`

HasSignificant returns a boolean if a field has been set.

### SetSignificantNil

`func (o *ExperimentVariantResultFrequentist) SetSignificantNil(b bool)`

 SetSignificantNil sets the value for Significant to be an explicit nil

### UnsetSignificant
`func (o *ExperimentVariantResultFrequentist) UnsetSignificant()`

UnsetSignificant ensures that no value is present for Significant, not even an explicit nil
### GetStepCounts

`func (o *ExperimentVariantResultFrequentist) GetStepCounts() []int32`

GetStepCounts returns the StepCounts field if non-nil, zero value otherwise.

### GetStepCountsOk

`func (o *ExperimentVariantResultFrequentist) GetStepCountsOk() (*[]int32, bool)`

GetStepCountsOk returns a tuple with the StepCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepCounts

`func (o *ExperimentVariantResultFrequentist) SetStepCounts(v []int32)`

SetStepCounts sets StepCounts field to given value.

### HasStepCounts

`func (o *ExperimentVariantResultFrequentist) HasStepCounts() bool`

HasStepCounts returns a boolean if a field has been set.

### SetStepCountsNil

`func (o *ExperimentVariantResultFrequentist) SetStepCountsNil(b bool)`

 SetStepCountsNil sets the value for StepCounts to be an explicit nil

### UnsetStepCounts
`func (o *ExperimentVariantResultFrequentist) UnsetStepCounts()`

UnsetStepCounts ensures that no value is present for StepCounts, not even an explicit nil
### GetStepSessions

`func (o *ExperimentVariantResultFrequentist) GetStepSessions() [][]SessionData`

GetStepSessions returns the StepSessions field if non-nil, zero value otherwise.

### GetStepSessionsOk

`func (o *ExperimentVariantResultFrequentist) GetStepSessionsOk() (*[][]SessionData, bool)`

GetStepSessionsOk returns a tuple with the StepSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSessions

`func (o *ExperimentVariantResultFrequentist) SetStepSessions(v [][]SessionData)`

SetStepSessions sets StepSessions field to given value.

### HasStepSessions

`func (o *ExperimentVariantResultFrequentist) HasStepSessions() bool`

HasStepSessions returns a boolean if a field has been set.

### SetStepSessionsNil

`func (o *ExperimentVariantResultFrequentist) SetStepSessionsNil(b bool)`

 SetStepSessionsNil sets the value for StepSessions to be an explicit nil

### UnsetStepSessions
`func (o *ExperimentVariantResultFrequentist) UnsetStepSessions()`

UnsetStepSessions ensures that no value is present for StepSessions, not even an explicit nil
### GetSum

`func (o *ExperimentVariantResultFrequentist) GetSum() float32`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *ExperimentVariantResultFrequentist) GetSumOk() (*float32, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *ExperimentVariantResultFrequentist) SetSum(v float32)`

SetSum sets Sum field to given value.


### GetSumSquares

`func (o *ExperimentVariantResultFrequentist) GetSumSquares() float32`

GetSumSquares returns the SumSquares field if non-nil, zero value otherwise.

### GetSumSquaresOk

`func (o *ExperimentVariantResultFrequentist) GetSumSquaresOk() (*float32, bool)`

GetSumSquaresOk returns a tuple with the SumSquares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumSquares

`func (o *ExperimentVariantResultFrequentist) SetSumSquares(v float32)`

SetSumSquares sets SumSquares field to given value.


### GetValidationFailures

`func (o *ExperimentVariantResultFrequentist) GetValidationFailures() []ExperimentStatsValidationFailure`

GetValidationFailures returns the ValidationFailures field if non-nil, zero value otherwise.

### GetValidationFailuresOk

`func (o *ExperimentVariantResultFrequentist) GetValidationFailuresOk() (*[]ExperimentStatsValidationFailure, bool)`

GetValidationFailuresOk returns a tuple with the ValidationFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFailures

`func (o *ExperimentVariantResultFrequentist) SetValidationFailures(v []ExperimentStatsValidationFailure)`

SetValidationFailures sets ValidationFailures field to given value.

### HasValidationFailures

`func (o *ExperimentVariantResultFrequentist) HasValidationFailures() bool`

HasValidationFailures returns a boolean if a field has been set.

### SetValidationFailuresNil

`func (o *ExperimentVariantResultFrequentist) SetValidationFailuresNil(b bool)`

 SetValidationFailuresNil sets the value for ValidationFailures to be an explicit nil

### UnsetValidationFailures
`func (o *ExperimentVariantResultFrequentist) UnsetValidationFailures()`

UnsetValidationFailures ensures that no value is present for ValidationFailures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


