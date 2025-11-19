# ExperimentStatsBaseValidated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DenominatorSum** | Pointer to **NullableFloat32** |  | [optional] 
**DenominatorSumSquares** | Pointer to **NullableFloat32** |  | [optional] 
**Key** | **string** |  | 
**NumberOfSamples** | **int32** |  | 
**NumeratorDenominatorSumProduct** | Pointer to **NullableFloat32** |  | [optional] 
**StepCounts** | Pointer to **[]int32** |  | [optional] 
**StepSessions** | Pointer to [**[][]SessionData**]([]SessionData.md) |  | [optional] 
**Sum** | **float32** |  | 
**SumSquares** | **float32** |  | 
**ValidationFailures** | Pointer to [**[]ExperimentStatsValidationFailure**](ExperimentStatsValidationFailure.md) |  | [optional] 

## Methods

### NewExperimentStatsBaseValidated

`func NewExperimentStatsBaseValidated(key string, numberOfSamples int32, sum float32, sumSquares float32, ) *ExperimentStatsBaseValidated`

NewExperimentStatsBaseValidated instantiates a new ExperimentStatsBaseValidated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentStatsBaseValidatedWithDefaults

`func NewExperimentStatsBaseValidatedWithDefaults() *ExperimentStatsBaseValidated`

NewExperimentStatsBaseValidatedWithDefaults instantiates a new ExperimentStatsBaseValidated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenominatorSum

`func (o *ExperimentStatsBaseValidated) GetDenominatorSum() float32`

GetDenominatorSum returns the DenominatorSum field if non-nil, zero value otherwise.

### GetDenominatorSumOk

`func (o *ExperimentStatsBaseValidated) GetDenominatorSumOk() (*float32, bool)`

GetDenominatorSumOk returns a tuple with the DenominatorSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorSum

`func (o *ExperimentStatsBaseValidated) SetDenominatorSum(v float32)`

SetDenominatorSum sets DenominatorSum field to given value.

### HasDenominatorSum

`func (o *ExperimentStatsBaseValidated) HasDenominatorSum() bool`

HasDenominatorSum returns a boolean if a field has been set.

### SetDenominatorSumNil

`func (o *ExperimentStatsBaseValidated) SetDenominatorSumNil(b bool)`

 SetDenominatorSumNil sets the value for DenominatorSum to be an explicit nil

### UnsetDenominatorSum
`func (o *ExperimentStatsBaseValidated) UnsetDenominatorSum()`

UnsetDenominatorSum ensures that no value is present for DenominatorSum, not even an explicit nil
### GetDenominatorSumSquares

`func (o *ExperimentStatsBaseValidated) GetDenominatorSumSquares() float32`

GetDenominatorSumSquares returns the DenominatorSumSquares field if non-nil, zero value otherwise.

### GetDenominatorSumSquaresOk

`func (o *ExperimentStatsBaseValidated) GetDenominatorSumSquaresOk() (*float32, bool)`

GetDenominatorSumSquaresOk returns a tuple with the DenominatorSumSquares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorSumSquares

`func (o *ExperimentStatsBaseValidated) SetDenominatorSumSquares(v float32)`

SetDenominatorSumSquares sets DenominatorSumSquares field to given value.

### HasDenominatorSumSquares

`func (o *ExperimentStatsBaseValidated) HasDenominatorSumSquares() bool`

HasDenominatorSumSquares returns a boolean if a field has been set.

### SetDenominatorSumSquaresNil

`func (o *ExperimentStatsBaseValidated) SetDenominatorSumSquaresNil(b bool)`

 SetDenominatorSumSquaresNil sets the value for DenominatorSumSquares to be an explicit nil

### UnsetDenominatorSumSquares
`func (o *ExperimentStatsBaseValidated) UnsetDenominatorSumSquares()`

UnsetDenominatorSumSquares ensures that no value is present for DenominatorSumSquares, not even an explicit nil
### GetKey

`func (o *ExperimentStatsBaseValidated) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExperimentStatsBaseValidated) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExperimentStatsBaseValidated) SetKey(v string)`

SetKey sets Key field to given value.


### GetNumberOfSamples

`func (o *ExperimentStatsBaseValidated) GetNumberOfSamples() int32`

GetNumberOfSamples returns the NumberOfSamples field if non-nil, zero value otherwise.

### GetNumberOfSamplesOk

`func (o *ExperimentStatsBaseValidated) GetNumberOfSamplesOk() (*int32, bool)`

GetNumberOfSamplesOk returns a tuple with the NumberOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSamples

`func (o *ExperimentStatsBaseValidated) SetNumberOfSamples(v int32)`

SetNumberOfSamples sets NumberOfSamples field to given value.


### GetNumeratorDenominatorSumProduct

`func (o *ExperimentStatsBaseValidated) GetNumeratorDenominatorSumProduct() float32`

GetNumeratorDenominatorSumProduct returns the NumeratorDenominatorSumProduct field if non-nil, zero value otherwise.

### GetNumeratorDenominatorSumProductOk

`func (o *ExperimentStatsBaseValidated) GetNumeratorDenominatorSumProductOk() (*float32, bool)`

GetNumeratorDenominatorSumProductOk returns a tuple with the NumeratorDenominatorSumProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeratorDenominatorSumProduct

`func (o *ExperimentStatsBaseValidated) SetNumeratorDenominatorSumProduct(v float32)`

SetNumeratorDenominatorSumProduct sets NumeratorDenominatorSumProduct field to given value.

### HasNumeratorDenominatorSumProduct

`func (o *ExperimentStatsBaseValidated) HasNumeratorDenominatorSumProduct() bool`

HasNumeratorDenominatorSumProduct returns a boolean if a field has been set.

### SetNumeratorDenominatorSumProductNil

`func (o *ExperimentStatsBaseValidated) SetNumeratorDenominatorSumProductNil(b bool)`

 SetNumeratorDenominatorSumProductNil sets the value for NumeratorDenominatorSumProduct to be an explicit nil

### UnsetNumeratorDenominatorSumProduct
`func (o *ExperimentStatsBaseValidated) UnsetNumeratorDenominatorSumProduct()`

UnsetNumeratorDenominatorSumProduct ensures that no value is present for NumeratorDenominatorSumProduct, not even an explicit nil
### GetStepCounts

`func (o *ExperimentStatsBaseValidated) GetStepCounts() []int32`

GetStepCounts returns the StepCounts field if non-nil, zero value otherwise.

### GetStepCountsOk

`func (o *ExperimentStatsBaseValidated) GetStepCountsOk() (*[]int32, bool)`

GetStepCountsOk returns a tuple with the StepCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepCounts

`func (o *ExperimentStatsBaseValidated) SetStepCounts(v []int32)`

SetStepCounts sets StepCounts field to given value.

### HasStepCounts

`func (o *ExperimentStatsBaseValidated) HasStepCounts() bool`

HasStepCounts returns a boolean if a field has been set.

### SetStepCountsNil

`func (o *ExperimentStatsBaseValidated) SetStepCountsNil(b bool)`

 SetStepCountsNil sets the value for StepCounts to be an explicit nil

### UnsetStepCounts
`func (o *ExperimentStatsBaseValidated) UnsetStepCounts()`

UnsetStepCounts ensures that no value is present for StepCounts, not even an explicit nil
### GetStepSessions

`func (o *ExperimentStatsBaseValidated) GetStepSessions() [][]SessionData`

GetStepSessions returns the StepSessions field if non-nil, zero value otherwise.

### GetStepSessionsOk

`func (o *ExperimentStatsBaseValidated) GetStepSessionsOk() (*[][]SessionData, bool)`

GetStepSessionsOk returns a tuple with the StepSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSessions

`func (o *ExperimentStatsBaseValidated) SetStepSessions(v [][]SessionData)`

SetStepSessions sets StepSessions field to given value.

### HasStepSessions

`func (o *ExperimentStatsBaseValidated) HasStepSessions() bool`

HasStepSessions returns a boolean if a field has been set.

### SetStepSessionsNil

`func (o *ExperimentStatsBaseValidated) SetStepSessionsNil(b bool)`

 SetStepSessionsNil sets the value for StepSessions to be an explicit nil

### UnsetStepSessions
`func (o *ExperimentStatsBaseValidated) UnsetStepSessions()`

UnsetStepSessions ensures that no value is present for StepSessions, not even an explicit nil
### GetSum

`func (o *ExperimentStatsBaseValidated) GetSum() float32`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *ExperimentStatsBaseValidated) GetSumOk() (*float32, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *ExperimentStatsBaseValidated) SetSum(v float32)`

SetSum sets Sum field to given value.


### GetSumSquares

`func (o *ExperimentStatsBaseValidated) GetSumSquares() float32`

GetSumSquares returns the SumSquares field if non-nil, zero value otherwise.

### GetSumSquaresOk

`func (o *ExperimentStatsBaseValidated) GetSumSquaresOk() (*float32, bool)`

GetSumSquaresOk returns a tuple with the SumSquares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumSquares

`func (o *ExperimentStatsBaseValidated) SetSumSquares(v float32)`

SetSumSquares sets SumSquares field to given value.


### GetValidationFailures

`func (o *ExperimentStatsBaseValidated) GetValidationFailures() []ExperimentStatsValidationFailure`

GetValidationFailures returns the ValidationFailures field if non-nil, zero value otherwise.

### GetValidationFailuresOk

`func (o *ExperimentStatsBaseValidated) GetValidationFailuresOk() (*[]ExperimentStatsValidationFailure, bool)`

GetValidationFailuresOk returns a tuple with the ValidationFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFailures

`func (o *ExperimentStatsBaseValidated) SetValidationFailures(v []ExperimentStatsValidationFailure)`

SetValidationFailures sets ValidationFailures field to given value.

### HasValidationFailures

`func (o *ExperimentStatsBaseValidated) HasValidationFailures() bool`

HasValidationFailures returns a boolean if a field has been set.

### SetValidationFailuresNil

`func (o *ExperimentStatsBaseValidated) SetValidationFailuresNil(b bool)`

 SetValidationFailuresNil sets the value for ValidationFailures to be an explicit nil

### UnsetValidationFailures
`func (o *ExperimentStatsBaseValidated) UnsetValidationFailures()`

UnsetValidationFailures ensures that no value is present for ValidationFailures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


