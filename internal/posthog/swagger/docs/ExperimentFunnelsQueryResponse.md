# ExperimentFunnelsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredibleIntervals** | **map[string][]float32** |  | 
**ExpectedLoss** | **float32** |  | 
**FunnelsQuery** | Pointer to [**FunnelsQuery**](FunnelsQuery.md) |  | [optional] 
**Insight** | **[][]map[string]interface{}** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentFunnelsQuery"]
**Probability** | **map[string]float32** |  | 
**SignificanceCode** | [**ExperimentSignificanceCode**](ExperimentSignificanceCode.md) |  | 
**Significant** | **bool** |  | 
**StatsVersion** | Pointer to **NullableInt32** |  | [optional] 
**Variants** | [**[]ExperimentVariantFunnelsBaseStats**](ExperimentVariantFunnelsBaseStats.md) |  | 

## Methods

### NewExperimentFunnelsQueryResponse

`func NewExperimentFunnelsQueryResponse(credibleIntervals map[string][]float32, expectedLoss float32, insight [][]map[string]interface{}, probability map[string]float32, significanceCode ExperimentSignificanceCode, significant bool, variants []ExperimentVariantFunnelsBaseStats, ) *ExperimentFunnelsQueryResponse`

NewExperimentFunnelsQueryResponse instantiates a new ExperimentFunnelsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentFunnelsQueryResponseWithDefaults

`func NewExperimentFunnelsQueryResponseWithDefaults() *ExperimentFunnelsQueryResponse`

NewExperimentFunnelsQueryResponseWithDefaults instantiates a new ExperimentFunnelsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredibleIntervals

`func (o *ExperimentFunnelsQueryResponse) GetCredibleIntervals() map[string][]float32`

GetCredibleIntervals returns the CredibleIntervals field if non-nil, zero value otherwise.

### GetCredibleIntervalsOk

`func (o *ExperimentFunnelsQueryResponse) GetCredibleIntervalsOk() (*map[string][]float32, bool)`

GetCredibleIntervalsOk returns a tuple with the CredibleIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleIntervals

`func (o *ExperimentFunnelsQueryResponse) SetCredibleIntervals(v map[string][]float32)`

SetCredibleIntervals sets CredibleIntervals field to given value.


### GetExpectedLoss

`func (o *ExperimentFunnelsQueryResponse) GetExpectedLoss() float32`

GetExpectedLoss returns the ExpectedLoss field if non-nil, zero value otherwise.

### GetExpectedLossOk

`func (o *ExperimentFunnelsQueryResponse) GetExpectedLossOk() (*float32, bool)`

GetExpectedLossOk returns a tuple with the ExpectedLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLoss

`func (o *ExperimentFunnelsQueryResponse) SetExpectedLoss(v float32)`

SetExpectedLoss sets ExpectedLoss field to given value.


### GetFunnelsQuery

`func (o *ExperimentFunnelsQueryResponse) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *ExperimentFunnelsQueryResponse) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *ExperimentFunnelsQueryResponse) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.

### HasFunnelsQuery

`func (o *ExperimentFunnelsQueryResponse) HasFunnelsQuery() bool`

HasFunnelsQuery returns a boolean if a field has been set.

### GetInsight

`func (o *ExperimentFunnelsQueryResponse) GetInsight() [][]map[string]interface{}`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *ExperimentFunnelsQueryResponse) GetInsightOk() (*[][]map[string]interface{}, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *ExperimentFunnelsQueryResponse) SetInsight(v [][]map[string]interface{})`

SetInsight sets Insight field to given value.


### GetKind

`func (o *ExperimentFunnelsQueryResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentFunnelsQueryResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentFunnelsQueryResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentFunnelsQueryResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProbability

`func (o *ExperimentFunnelsQueryResponse) GetProbability() map[string]float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *ExperimentFunnelsQueryResponse) GetProbabilityOk() (*map[string]float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *ExperimentFunnelsQueryResponse) SetProbability(v map[string]float32)`

SetProbability sets Probability field to given value.


### GetSignificanceCode

`func (o *ExperimentFunnelsQueryResponse) GetSignificanceCode() ExperimentSignificanceCode`

GetSignificanceCode returns the SignificanceCode field if non-nil, zero value otherwise.

### GetSignificanceCodeOk

`func (o *ExperimentFunnelsQueryResponse) GetSignificanceCodeOk() (*ExperimentSignificanceCode, bool)`

GetSignificanceCodeOk returns a tuple with the SignificanceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceCode

`func (o *ExperimentFunnelsQueryResponse) SetSignificanceCode(v ExperimentSignificanceCode)`

SetSignificanceCode sets SignificanceCode field to given value.


### GetSignificant

`func (o *ExperimentFunnelsQueryResponse) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *ExperimentFunnelsQueryResponse) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *ExperimentFunnelsQueryResponse) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.


### GetStatsVersion

`func (o *ExperimentFunnelsQueryResponse) GetStatsVersion() int32`

GetStatsVersion returns the StatsVersion field if non-nil, zero value otherwise.

### GetStatsVersionOk

`func (o *ExperimentFunnelsQueryResponse) GetStatsVersionOk() (*int32, bool)`

GetStatsVersionOk returns a tuple with the StatsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersion

`func (o *ExperimentFunnelsQueryResponse) SetStatsVersion(v int32)`

SetStatsVersion sets StatsVersion field to given value.

### HasStatsVersion

`func (o *ExperimentFunnelsQueryResponse) HasStatsVersion() bool`

HasStatsVersion returns a boolean if a field has been set.

### SetStatsVersionNil

`func (o *ExperimentFunnelsQueryResponse) SetStatsVersionNil(b bool)`

 SetStatsVersionNil sets the value for StatsVersion to be an explicit nil

### UnsetStatsVersion
`func (o *ExperimentFunnelsQueryResponse) UnsetStatsVersion()`

UnsetStatsVersion ensures that no value is present for StatsVersion, not even an explicit nil
### GetVariants

`func (o *ExperimentFunnelsQueryResponse) GetVariants() []ExperimentVariantFunnelsBaseStats`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentFunnelsQueryResponse) GetVariantsOk() (*[]ExperimentVariantFunnelsBaseStats, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentFunnelsQueryResponse) SetVariants(v []ExperimentVariantFunnelsBaseStats)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


