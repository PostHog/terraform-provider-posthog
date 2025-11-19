# ExperimentVariantResultBayesian

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChanceToWin** | Pointer to **NullableFloat32** |  | [optional] 
**CredibleInterval** | Pointer to **[]float32** |  | [optional] 
**DenominatorSum** | Pointer to **NullableFloat32** |  | [optional] 
**DenominatorSumSquares** | Pointer to **NullableFloat32** |  | [optional] 
**Key** | **string** |  | 
**Method** | Pointer to **string** |  | [optional] [default to "bayesian"]
**NumberOfSamples** | **int32** |  | 
**NumeratorDenominatorSumProduct** | Pointer to **NullableFloat32** |  | [optional] 
**Significant** | Pointer to **NullableBool** |  | [optional] 
**StepCounts** | Pointer to **[]int32** |  | [optional] 
**StepSessions** | Pointer to [**[][]SessionData**]([]SessionData.md) |  | [optional] 
**Sum** | **float32** |  | 
**SumSquares** | **float32** |  | 
**ValidationFailures** | Pointer to [**[]ExperimentStatsValidationFailure**](ExperimentStatsValidationFailure.md) |  | [optional] 

## Methods

### NewExperimentVariantResultBayesian

`func NewExperimentVariantResultBayesian(key string, numberOfSamples int32, sum float32, sumSquares float32, ) *ExperimentVariantResultBayesian`

NewExperimentVariantResultBayesian instantiates a new ExperimentVariantResultBayesian object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVariantResultBayesianWithDefaults

`func NewExperimentVariantResultBayesianWithDefaults() *ExperimentVariantResultBayesian`

NewExperimentVariantResultBayesianWithDefaults instantiates a new ExperimentVariantResultBayesian object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanceToWin

`func (o *ExperimentVariantResultBayesian) GetChanceToWin() float32`

GetChanceToWin returns the ChanceToWin field if non-nil, zero value otherwise.

### GetChanceToWinOk

`func (o *ExperimentVariantResultBayesian) GetChanceToWinOk() (*float32, bool)`

GetChanceToWinOk returns a tuple with the ChanceToWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanceToWin

`func (o *ExperimentVariantResultBayesian) SetChanceToWin(v float32)`

SetChanceToWin sets ChanceToWin field to given value.

### HasChanceToWin

`func (o *ExperimentVariantResultBayesian) HasChanceToWin() bool`

HasChanceToWin returns a boolean if a field has been set.

### SetChanceToWinNil

`func (o *ExperimentVariantResultBayesian) SetChanceToWinNil(b bool)`

 SetChanceToWinNil sets the value for ChanceToWin to be an explicit nil

### UnsetChanceToWin
`func (o *ExperimentVariantResultBayesian) UnsetChanceToWin()`

UnsetChanceToWin ensures that no value is present for ChanceToWin, not even an explicit nil
### GetCredibleInterval

`func (o *ExperimentVariantResultBayesian) GetCredibleInterval() []float32`

GetCredibleInterval returns the CredibleInterval field if non-nil, zero value otherwise.

### GetCredibleIntervalOk

`func (o *ExperimentVariantResultBayesian) GetCredibleIntervalOk() (*[]float32, bool)`

GetCredibleIntervalOk returns a tuple with the CredibleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleInterval

`func (o *ExperimentVariantResultBayesian) SetCredibleInterval(v []float32)`

SetCredibleInterval sets CredibleInterval field to given value.

### HasCredibleInterval

`func (o *ExperimentVariantResultBayesian) HasCredibleInterval() bool`

HasCredibleInterval returns a boolean if a field has been set.

### SetCredibleIntervalNil

`func (o *ExperimentVariantResultBayesian) SetCredibleIntervalNil(b bool)`

 SetCredibleIntervalNil sets the value for CredibleInterval to be an explicit nil

### UnsetCredibleInterval
`func (o *ExperimentVariantResultBayesian) UnsetCredibleInterval()`

UnsetCredibleInterval ensures that no value is present for CredibleInterval, not even an explicit nil
### GetDenominatorSum

`func (o *ExperimentVariantResultBayesian) GetDenominatorSum() float32`

GetDenominatorSum returns the DenominatorSum field if non-nil, zero value otherwise.

### GetDenominatorSumOk

`func (o *ExperimentVariantResultBayesian) GetDenominatorSumOk() (*float32, bool)`

GetDenominatorSumOk returns a tuple with the DenominatorSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorSum

`func (o *ExperimentVariantResultBayesian) SetDenominatorSum(v float32)`

SetDenominatorSum sets DenominatorSum field to given value.

### HasDenominatorSum

`func (o *ExperimentVariantResultBayesian) HasDenominatorSum() bool`

HasDenominatorSum returns a boolean if a field has been set.

### SetDenominatorSumNil

`func (o *ExperimentVariantResultBayesian) SetDenominatorSumNil(b bool)`

 SetDenominatorSumNil sets the value for DenominatorSum to be an explicit nil

### UnsetDenominatorSum
`func (o *ExperimentVariantResultBayesian) UnsetDenominatorSum()`

UnsetDenominatorSum ensures that no value is present for DenominatorSum, not even an explicit nil
### GetDenominatorSumSquares

`func (o *ExperimentVariantResultBayesian) GetDenominatorSumSquares() float32`

GetDenominatorSumSquares returns the DenominatorSumSquares field if non-nil, zero value otherwise.

### GetDenominatorSumSquaresOk

`func (o *ExperimentVariantResultBayesian) GetDenominatorSumSquaresOk() (*float32, bool)`

GetDenominatorSumSquaresOk returns a tuple with the DenominatorSumSquares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorSumSquares

`func (o *ExperimentVariantResultBayesian) SetDenominatorSumSquares(v float32)`

SetDenominatorSumSquares sets DenominatorSumSquares field to given value.

### HasDenominatorSumSquares

`func (o *ExperimentVariantResultBayesian) HasDenominatorSumSquares() bool`

HasDenominatorSumSquares returns a boolean if a field has been set.

### SetDenominatorSumSquaresNil

`func (o *ExperimentVariantResultBayesian) SetDenominatorSumSquaresNil(b bool)`

 SetDenominatorSumSquaresNil sets the value for DenominatorSumSquares to be an explicit nil

### UnsetDenominatorSumSquares
`func (o *ExperimentVariantResultBayesian) UnsetDenominatorSumSquares()`

UnsetDenominatorSumSquares ensures that no value is present for DenominatorSumSquares, not even an explicit nil
### GetKey

`func (o *ExperimentVariantResultBayesian) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExperimentVariantResultBayesian) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExperimentVariantResultBayesian) SetKey(v string)`

SetKey sets Key field to given value.


### GetMethod

`func (o *ExperimentVariantResultBayesian) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ExperimentVariantResultBayesian) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ExperimentVariantResultBayesian) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ExperimentVariantResultBayesian) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNumberOfSamples

`func (o *ExperimentVariantResultBayesian) GetNumberOfSamples() int32`

GetNumberOfSamples returns the NumberOfSamples field if non-nil, zero value otherwise.

### GetNumberOfSamplesOk

`func (o *ExperimentVariantResultBayesian) GetNumberOfSamplesOk() (*int32, bool)`

GetNumberOfSamplesOk returns a tuple with the NumberOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSamples

`func (o *ExperimentVariantResultBayesian) SetNumberOfSamples(v int32)`

SetNumberOfSamples sets NumberOfSamples field to given value.


### GetNumeratorDenominatorSumProduct

`func (o *ExperimentVariantResultBayesian) GetNumeratorDenominatorSumProduct() float32`

GetNumeratorDenominatorSumProduct returns the NumeratorDenominatorSumProduct field if non-nil, zero value otherwise.

### GetNumeratorDenominatorSumProductOk

`func (o *ExperimentVariantResultBayesian) GetNumeratorDenominatorSumProductOk() (*float32, bool)`

GetNumeratorDenominatorSumProductOk returns a tuple with the NumeratorDenominatorSumProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeratorDenominatorSumProduct

`func (o *ExperimentVariantResultBayesian) SetNumeratorDenominatorSumProduct(v float32)`

SetNumeratorDenominatorSumProduct sets NumeratorDenominatorSumProduct field to given value.

### HasNumeratorDenominatorSumProduct

`func (o *ExperimentVariantResultBayesian) HasNumeratorDenominatorSumProduct() bool`

HasNumeratorDenominatorSumProduct returns a boolean if a field has been set.

### SetNumeratorDenominatorSumProductNil

`func (o *ExperimentVariantResultBayesian) SetNumeratorDenominatorSumProductNil(b bool)`

 SetNumeratorDenominatorSumProductNil sets the value for NumeratorDenominatorSumProduct to be an explicit nil

### UnsetNumeratorDenominatorSumProduct
`func (o *ExperimentVariantResultBayesian) UnsetNumeratorDenominatorSumProduct()`

UnsetNumeratorDenominatorSumProduct ensures that no value is present for NumeratorDenominatorSumProduct, not even an explicit nil
### GetSignificant

`func (o *ExperimentVariantResultBayesian) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *ExperimentVariantResultBayesian) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *ExperimentVariantResultBayesian) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.

### HasSignificant

`func (o *ExperimentVariantResultBayesian) HasSignificant() bool`

HasSignificant returns a boolean if a field has been set.

### SetSignificantNil

`func (o *ExperimentVariantResultBayesian) SetSignificantNil(b bool)`

 SetSignificantNil sets the value for Significant to be an explicit nil

### UnsetSignificant
`func (o *ExperimentVariantResultBayesian) UnsetSignificant()`

UnsetSignificant ensures that no value is present for Significant, not even an explicit nil
### GetStepCounts

`func (o *ExperimentVariantResultBayesian) GetStepCounts() []int32`

GetStepCounts returns the StepCounts field if non-nil, zero value otherwise.

### GetStepCountsOk

`func (o *ExperimentVariantResultBayesian) GetStepCountsOk() (*[]int32, bool)`

GetStepCountsOk returns a tuple with the StepCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepCounts

`func (o *ExperimentVariantResultBayesian) SetStepCounts(v []int32)`

SetStepCounts sets StepCounts field to given value.

### HasStepCounts

`func (o *ExperimentVariantResultBayesian) HasStepCounts() bool`

HasStepCounts returns a boolean if a field has been set.

### SetStepCountsNil

`func (o *ExperimentVariantResultBayesian) SetStepCountsNil(b bool)`

 SetStepCountsNil sets the value for StepCounts to be an explicit nil

### UnsetStepCounts
`func (o *ExperimentVariantResultBayesian) UnsetStepCounts()`

UnsetStepCounts ensures that no value is present for StepCounts, not even an explicit nil
### GetStepSessions

`func (o *ExperimentVariantResultBayesian) GetStepSessions() [][]SessionData`

GetStepSessions returns the StepSessions field if non-nil, zero value otherwise.

### GetStepSessionsOk

`func (o *ExperimentVariantResultBayesian) GetStepSessionsOk() (*[][]SessionData, bool)`

GetStepSessionsOk returns a tuple with the StepSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSessions

`func (o *ExperimentVariantResultBayesian) SetStepSessions(v [][]SessionData)`

SetStepSessions sets StepSessions field to given value.

### HasStepSessions

`func (o *ExperimentVariantResultBayesian) HasStepSessions() bool`

HasStepSessions returns a boolean if a field has been set.

### SetStepSessionsNil

`func (o *ExperimentVariantResultBayesian) SetStepSessionsNil(b bool)`

 SetStepSessionsNil sets the value for StepSessions to be an explicit nil

### UnsetStepSessions
`func (o *ExperimentVariantResultBayesian) UnsetStepSessions()`

UnsetStepSessions ensures that no value is present for StepSessions, not even an explicit nil
### GetSum

`func (o *ExperimentVariantResultBayesian) GetSum() float32`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *ExperimentVariantResultBayesian) GetSumOk() (*float32, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *ExperimentVariantResultBayesian) SetSum(v float32)`

SetSum sets Sum field to given value.


### GetSumSquares

`func (o *ExperimentVariantResultBayesian) GetSumSquares() float32`

GetSumSquares returns the SumSquares field if non-nil, zero value otherwise.

### GetSumSquaresOk

`func (o *ExperimentVariantResultBayesian) GetSumSquaresOk() (*float32, bool)`

GetSumSquaresOk returns a tuple with the SumSquares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumSquares

`func (o *ExperimentVariantResultBayesian) SetSumSquares(v float32)`

SetSumSquares sets SumSquares field to given value.


### GetValidationFailures

`func (o *ExperimentVariantResultBayesian) GetValidationFailures() []ExperimentStatsValidationFailure`

GetValidationFailures returns the ValidationFailures field if non-nil, zero value otherwise.

### GetValidationFailuresOk

`func (o *ExperimentVariantResultBayesian) GetValidationFailuresOk() (*[]ExperimentStatsValidationFailure, bool)`

GetValidationFailuresOk returns a tuple with the ValidationFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFailures

`func (o *ExperimentVariantResultBayesian) SetValidationFailures(v []ExperimentStatsValidationFailure)`

SetValidationFailures sets ValidationFailures field to given value.

### HasValidationFailures

`func (o *ExperimentVariantResultBayesian) HasValidationFailures() bool`

HasValidationFailures returns a boolean if a field has been set.

### SetValidationFailuresNil

`func (o *ExperimentVariantResultBayesian) SetValidationFailuresNil(b bool)`

 SetValidationFailuresNil sets the value for ValidationFailures to be an explicit nil

### UnsetValidationFailures
`func (o *ExperimentVariantResultBayesian) UnsetValidationFailures()`

UnsetValidationFailures ensures that no value is present for ValidationFailures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


