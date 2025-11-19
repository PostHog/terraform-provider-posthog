# ExperimentTrendsQueryResponse

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

### NewExperimentTrendsQueryResponse

`func NewExperimentTrendsQueryResponse(credibleIntervals map[string][]float32, insight []map[string]interface{}, pValue float32, probability map[string]float32, significanceCode ExperimentSignificanceCode, significant bool, variants []ExperimentVariantTrendsBaseStats, ) *ExperimentTrendsQueryResponse`

NewExperimentTrendsQueryResponse instantiates a new ExperimentTrendsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentTrendsQueryResponseWithDefaults

`func NewExperimentTrendsQueryResponseWithDefaults() *ExperimentTrendsQueryResponse`

NewExperimentTrendsQueryResponseWithDefaults instantiates a new ExperimentTrendsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountQuery

`func (o *ExperimentTrendsQueryResponse) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *ExperimentTrendsQueryResponse) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *ExperimentTrendsQueryResponse) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.

### HasCountQuery

`func (o *ExperimentTrendsQueryResponse) HasCountQuery() bool`

HasCountQuery returns a boolean if a field has been set.

### GetCredibleIntervals

`func (o *ExperimentTrendsQueryResponse) GetCredibleIntervals() map[string][]float32`

GetCredibleIntervals returns the CredibleIntervals field if non-nil, zero value otherwise.

### GetCredibleIntervalsOk

`func (o *ExperimentTrendsQueryResponse) GetCredibleIntervalsOk() (*map[string][]float32, bool)`

GetCredibleIntervalsOk returns a tuple with the CredibleIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleIntervals

`func (o *ExperimentTrendsQueryResponse) SetCredibleIntervals(v map[string][]float32)`

SetCredibleIntervals sets CredibleIntervals field to given value.


### GetExposureQuery

`func (o *ExperimentTrendsQueryResponse) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *ExperimentTrendsQueryResponse) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *ExperimentTrendsQueryResponse) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *ExperimentTrendsQueryResponse) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetInsight

`func (o *ExperimentTrendsQueryResponse) GetInsight() []map[string]interface{}`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *ExperimentTrendsQueryResponse) GetInsightOk() (*[]map[string]interface{}, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *ExperimentTrendsQueryResponse) SetInsight(v []map[string]interface{})`

SetInsight sets Insight field to given value.


### GetKind

`func (o *ExperimentTrendsQueryResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentTrendsQueryResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentTrendsQueryResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentTrendsQueryResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPValue

`func (o *ExperimentTrendsQueryResponse) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *ExperimentTrendsQueryResponse) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *ExperimentTrendsQueryResponse) SetPValue(v float32)`

SetPValue sets PValue field to given value.


### GetProbability

`func (o *ExperimentTrendsQueryResponse) GetProbability() map[string]float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *ExperimentTrendsQueryResponse) GetProbabilityOk() (*map[string]float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *ExperimentTrendsQueryResponse) SetProbability(v map[string]float32)`

SetProbability sets Probability field to given value.


### GetSignificanceCode

`func (o *ExperimentTrendsQueryResponse) GetSignificanceCode() ExperimentSignificanceCode`

GetSignificanceCode returns the SignificanceCode field if non-nil, zero value otherwise.

### GetSignificanceCodeOk

`func (o *ExperimentTrendsQueryResponse) GetSignificanceCodeOk() (*ExperimentSignificanceCode, bool)`

GetSignificanceCodeOk returns a tuple with the SignificanceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceCode

`func (o *ExperimentTrendsQueryResponse) SetSignificanceCode(v ExperimentSignificanceCode)`

SetSignificanceCode sets SignificanceCode field to given value.


### GetSignificant

`func (o *ExperimentTrendsQueryResponse) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *ExperimentTrendsQueryResponse) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *ExperimentTrendsQueryResponse) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.


### GetStatsVersion

`func (o *ExperimentTrendsQueryResponse) GetStatsVersion() int32`

GetStatsVersion returns the StatsVersion field if non-nil, zero value otherwise.

### GetStatsVersionOk

`func (o *ExperimentTrendsQueryResponse) GetStatsVersionOk() (*int32, bool)`

GetStatsVersionOk returns a tuple with the StatsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersion

`func (o *ExperimentTrendsQueryResponse) SetStatsVersion(v int32)`

SetStatsVersion sets StatsVersion field to given value.

### HasStatsVersion

`func (o *ExperimentTrendsQueryResponse) HasStatsVersion() bool`

HasStatsVersion returns a boolean if a field has been set.

### SetStatsVersionNil

`func (o *ExperimentTrendsQueryResponse) SetStatsVersionNil(b bool)`

 SetStatsVersionNil sets the value for StatsVersion to be an explicit nil

### UnsetStatsVersion
`func (o *ExperimentTrendsQueryResponse) UnsetStatsVersion()`

UnsetStatsVersion ensures that no value is present for StatsVersion, not even an explicit nil
### GetVariants

`func (o *ExperimentTrendsQueryResponse) GetVariants() []ExperimentVariantTrendsBaseStats`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExperimentTrendsQueryResponse) GetVariantsOk() (*[]ExperimentVariantTrendsBaseStats, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExperimentTrendsQueryResponse) SetVariants(v []ExperimentVariantTrendsBaseStats)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


