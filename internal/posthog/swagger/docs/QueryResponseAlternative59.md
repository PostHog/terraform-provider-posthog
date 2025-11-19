# QueryResponseAlternative59

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

### NewQueryResponseAlternative59

`func NewQueryResponseAlternative59(credibleIntervals map[string][]float32, expectedLoss float32, insight [][]map[string]interface{}, probability map[string]float32, significanceCode ExperimentSignificanceCode, significant bool, variants []ExperimentVariantFunnelsBaseStats, ) *QueryResponseAlternative59`

NewQueryResponseAlternative59 instantiates a new QueryResponseAlternative59 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative59WithDefaults

`func NewQueryResponseAlternative59WithDefaults() *QueryResponseAlternative59`

NewQueryResponseAlternative59WithDefaults instantiates a new QueryResponseAlternative59 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredibleIntervals

`func (o *QueryResponseAlternative59) GetCredibleIntervals() map[string][]float32`

GetCredibleIntervals returns the CredibleIntervals field if non-nil, zero value otherwise.

### GetCredibleIntervalsOk

`func (o *QueryResponseAlternative59) GetCredibleIntervalsOk() (*map[string][]float32, bool)`

GetCredibleIntervalsOk returns a tuple with the CredibleIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleIntervals

`func (o *QueryResponseAlternative59) SetCredibleIntervals(v map[string][]float32)`

SetCredibleIntervals sets CredibleIntervals field to given value.


### GetExpectedLoss

`func (o *QueryResponseAlternative59) GetExpectedLoss() float32`

GetExpectedLoss returns the ExpectedLoss field if non-nil, zero value otherwise.

### GetExpectedLossOk

`func (o *QueryResponseAlternative59) GetExpectedLossOk() (*float32, bool)`

GetExpectedLossOk returns a tuple with the ExpectedLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLoss

`func (o *QueryResponseAlternative59) SetExpectedLoss(v float32)`

SetExpectedLoss sets ExpectedLoss field to given value.


### GetFunnelsQuery

`func (o *QueryResponseAlternative59) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *QueryResponseAlternative59) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *QueryResponseAlternative59) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.

### HasFunnelsQuery

`func (o *QueryResponseAlternative59) HasFunnelsQuery() bool`

HasFunnelsQuery returns a boolean if a field has been set.

### GetInsight

`func (o *QueryResponseAlternative59) GetInsight() [][]map[string]interface{}`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *QueryResponseAlternative59) GetInsightOk() (*[][]map[string]interface{}, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *QueryResponseAlternative59) SetInsight(v [][]map[string]interface{})`

SetInsight sets Insight field to given value.


### GetKind

`func (o *QueryResponseAlternative59) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QueryResponseAlternative59) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QueryResponseAlternative59) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *QueryResponseAlternative59) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProbability

`func (o *QueryResponseAlternative59) GetProbability() map[string]float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *QueryResponseAlternative59) GetProbabilityOk() (*map[string]float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *QueryResponseAlternative59) SetProbability(v map[string]float32)`

SetProbability sets Probability field to given value.


### GetSignificanceCode

`func (o *QueryResponseAlternative59) GetSignificanceCode() ExperimentSignificanceCode`

GetSignificanceCode returns the SignificanceCode field if non-nil, zero value otherwise.

### GetSignificanceCodeOk

`func (o *QueryResponseAlternative59) GetSignificanceCodeOk() (*ExperimentSignificanceCode, bool)`

GetSignificanceCodeOk returns a tuple with the SignificanceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceCode

`func (o *QueryResponseAlternative59) SetSignificanceCode(v ExperimentSignificanceCode)`

SetSignificanceCode sets SignificanceCode field to given value.


### GetSignificant

`func (o *QueryResponseAlternative59) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *QueryResponseAlternative59) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *QueryResponseAlternative59) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.


### GetStatsVersion

`func (o *QueryResponseAlternative59) GetStatsVersion() int32`

GetStatsVersion returns the StatsVersion field if non-nil, zero value otherwise.

### GetStatsVersionOk

`func (o *QueryResponseAlternative59) GetStatsVersionOk() (*int32, bool)`

GetStatsVersionOk returns a tuple with the StatsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersion

`func (o *QueryResponseAlternative59) SetStatsVersion(v int32)`

SetStatsVersion sets StatsVersion field to given value.

### HasStatsVersion

`func (o *QueryResponseAlternative59) HasStatsVersion() bool`

HasStatsVersion returns a boolean if a field has been set.

### SetStatsVersionNil

`func (o *QueryResponseAlternative59) SetStatsVersionNil(b bool)`

 SetStatsVersionNil sets the value for StatsVersion to be an explicit nil

### UnsetStatsVersion
`func (o *QueryResponseAlternative59) UnsetStatsVersion()`

UnsetStatsVersion ensures that no value is present for StatsVersion, not even an explicit nil
### GetVariants

`func (o *QueryResponseAlternative59) GetVariants() []ExperimentVariantFunnelsBaseStats`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *QueryResponseAlternative59) GetVariantsOk() (*[]ExperimentVariantFunnelsBaseStats, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *QueryResponseAlternative59) SetVariants(v []ExperimentVariantFunnelsBaseStats)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


