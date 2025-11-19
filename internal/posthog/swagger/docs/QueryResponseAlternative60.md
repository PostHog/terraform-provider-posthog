# QueryResponseAlternative60

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountQuery** | Pointer to [**TrendsQuery**](TrendsQuery.md) |  | [optional] 
**CredibleIntervals** | **map[string][]float32** |  | 
**ExposureQuery** | Pointer to [**TrendsQuery**](TrendsQuery.md) |  | [optional] 
**Insight** | **[]map[string]interface{}** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentTrendsQuery"]
**PValue** | **float32** |  | 
**Probability** | **map[string]float32** |  | 
**SignificanceCode** | [**ExperimentSignificanceCode**](ExperimentSignificanceCode.md) |  | 
**Significant** | **bool** |  | 
**StatsVersion** | Pointer to **NullableInt32** |  | [optional] 
**Variants** | [**[]ExperimentVariantTrendsBaseStats**](ExperimentVariantTrendsBaseStats.md) |  | 

## Methods

### NewQueryResponseAlternative60

`func NewQueryResponseAlternative60(credibleIntervals map[string][]float32, insight []map[string]interface{}, pValue float32, probability map[string]float32, significanceCode ExperimentSignificanceCode, significant bool, variants []ExperimentVariantTrendsBaseStats, ) *QueryResponseAlternative60`

NewQueryResponseAlternative60 instantiates a new QueryResponseAlternative60 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseAlternative60WithDefaults

`func NewQueryResponseAlternative60WithDefaults() *QueryResponseAlternative60`

NewQueryResponseAlternative60WithDefaults instantiates a new QueryResponseAlternative60 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountQuery

`func (o *QueryResponseAlternative60) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *QueryResponseAlternative60) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *QueryResponseAlternative60) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.

### HasCountQuery

`func (o *QueryResponseAlternative60) HasCountQuery() bool`

HasCountQuery returns a boolean if a field has been set.

### GetCredibleIntervals

`func (o *QueryResponseAlternative60) GetCredibleIntervals() map[string][]float32`

GetCredibleIntervals returns the CredibleIntervals field if non-nil, zero value otherwise.

### GetCredibleIntervalsOk

`func (o *QueryResponseAlternative60) GetCredibleIntervalsOk() (*map[string][]float32, bool)`

GetCredibleIntervalsOk returns a tuple with the CredibleIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleIntervals

`func (o *QueryResponseAlternative60) SetCredibleIntervals(v map[string][]float32)`

SetCredibleIntervals sets CredibleIntervals field to given value.


### GetExposureQuery

`func (o *QueryResponseAlternative60) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *QueryResponseAlternative60) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *QueryResponseAlternative60) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *QueryResponseAlternative60) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetInsight

`func (o *QueryResponseAlternative60) GetInsight() []map[string]interface{}`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *QueryResponseAlternative60) GetInsightOk() (*[]map[string]interface{}, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *QueryResponseAlternative60) SetInsight(v []map[string]interface{})`

SetInsight sets Insight field to given value.


### GetKind

`func (o *QueryResponseAlternative60) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QueryResponseAlternative60) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QueryResponseAlternative60) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *QueryResponseAlternative60) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPValue

`func (o *QueryResponseAlternative60) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *QueryResponseAlternative60) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *QueryResponseAlternative60) SetPValue(v float32)`

SetPValue sets PValue field to given value.


### GetProbability

`func (o *QueryResponseAlternative60) GetProbability() map[string]float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *QueryResponseAlternative60) GetProbabilityOk() (*map[string]float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *QueryResponseAlternative60) SetProbability(v map[string]float32)`

SetProbability sets Probability field to given value.


### GetSignificanceCode

`func (o *QueryResponseAlternative60) GetSignificanceCode() ExperimentSignificanceCode`

GetSignificanceCode returns the SignificanceCode field if non-nil, zero value otherwise.

### GetSignificanceCodeOk

`func (o *QueryResponseAlternative60) GetSignificanceCodeOk() (*ExperimentSignificanceCode, bool)`

GetSignificanceCodeOk returns a tuple with the SignificanceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceCode

`func (o *QueryResponseAlternative60) SetSignificanceCode(v ExperimentSignificanceCode)`

SetSignificanceCode sets SignificanceCode field to given value.


### GetSignificant

`func (o *QueryResponseAlternative60) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *QueryResponseAlternative60) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *QueryResponseAlternative60) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.


### GetStatsVersion

`func (o *QueryResponseAlternative60) GetStatsVersion() int32`

GetStatsVersion returns the StatsVersion field if non-nil, zero value otherwise.

### GetStatsVersionOk

`func (o *QueryResponseAlternative60) GetStatsVersionOk() (*int32, bool)`

GetStatsVersionOk returns a tuple with the StatsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersion

`func (o *QueryResponseAlternative60) SetStatsVersion(v int32)`

SetStatsVersion sets StatsVersion field to given value.

### HasStatsVersion

`func (o *QueryResponseAlternative60) HasStatsVersion() bool`

HasStatsVersion returns a boolean if a field has been set.

### SetStatsVersionNil

`func (o *QueryResponseAlternative60) SetStatsVersionNil(b bool)`

 SetStatsVersionNil sets the value for StatsVersion to be an explicit nil

### UnsetStatsVersion
`func (o *QueryResponseAlternative60) UnsetStatsVersion()`

UnsetStatsVersion ensures that no value is present for StatsVersion, not even an explicit nil
### GetVariants

`func (o *QueryResponseAlternative60) GetVariants() []ExperimentVariantTrendsBaseStats`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *QueryResponseAlternative60) GetVariantsOk() (*[]ExperimentVariantTrendsBaseStats, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *QueryResponseAlternative60) SetVariants(v []ExperimentVariantTrendsBaseStats)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


