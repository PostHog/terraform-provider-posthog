# Response23

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

### NewResponse23

`func NewResponse23(credibleIntervals map[string][]float32, insight []map[string]interface{}, pValue float32, probability map[string]float32, significanceCode ExperimentSignificanceCode, significant bool, variants []ExperimentVariantTrendsBaseStats, ) *Response23`

NewResponse23 instantiates a new Response23 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponse23WithDefaults

`func NewResponse23WithDefaults() *Response23`

NewResponse23WithDefaults instantiates a new Response23 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountQuery

`func (o *Response23) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *Response23) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *Response23) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.

### HasCountQuery

`func (o *Response23) HasCountQuery() bool`

HasCountQuery returns a boolean if a field has been set.

### GetCredibleIntervals

`func (o *Response23) GetCredibleIntervals() map[string][]float32`

GetCredibleIntervals returns the CredibleIntervals field if non-nil, zero value otherwise.

### GetCredibleIntervalsOk

`func (o *Response23) GetCredibleIntervalsOk() (*map[string][]float32, bool)`

GetCredibleIntervalsOk returns a tuple with the CredibleIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredibleIntervals

`func (o *Response23) SetCredibleIntervals(v map[string][]float32)`

SetCredibleIntervals sets CredibleIntervals field to given value.


### GetExposureQuery

`func (o *Response23) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *Response23) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *Response23) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *Response23) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetInsight

`func (o *Response23) GetInsight() []map[string]interface{}`

GetInsight returns the Insight field if non-nil, zero value otherwise.

### GetInsightOk

`func (o *Response23) GetInsightOk() (*[]map[string]interface{}, bool)`

GetInsightOk returns a tuple with the Insight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsight

`func (o *Response23) SetInsight(v []map[string]interface{})`

SetInsight sets Insight field to given value.


### GetKind

`func (o *Response23) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Response23) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Response23) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Response23) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPValue

`func (o *Response23) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *Response23) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *Response23) SetPValue(v float32)`

SetPValue sets PValue field to given value.


### GetProbability

`func (o *Response23) GetProbability() map[string]float32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *Response23) GetProbabilityOk() (*map[string]float32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *Response23) SetProbability(v map[string]float32)`

SetProbability sets Probability field to given value.


### GetSignificanceCode

`func (o *Response23) GetSignificanceCode() ExperimentSignificanceCode`

GetSignificanceCode returns the SignificanceCode field if non-nil, zero value otherwise.

### GetSignificanceCodeOk

`func (o *Response23) GetSignificanceCodeOk() (*ExperimentSignificanceCode, bool)`

GetSignificanceCodeOk returns a tuple with the SignificanceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificanceCode

`func (o *Response23) SetSignificanceCode(v ExperimentSignificanceCode)`

SetSignificanceCode sets SignificanceCode field to given value.


### GetSignificant

`func (o *Response23) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *Response23) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *Response23) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.


### GetStatsVersion

`func (o *Response23) GetStatsVersion() int32`

GetStatsVersion returns the StatsVersion field if non-nil, zero value otherwise.

### GetStatsVersionOk

`func (o *Response23) GetStatsVersionOk() (*int32, bool)`

GetStatsVersionOk returns a tuple with the StatsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersion

`func (o *Response23) SetStatsVersion(v int32)`

SetStatsVersion sets StatsVersion field to given value.

### HasStatsVersion

`func (o *Response23) HasStatsVersion() bool`

HasStatsVersion returns a boolean if a field has been set.

### SetStatsVersionNil

`func (o *Response23) SetStatsVersionNil(b bool)`

 SetStatsVersionNil sets the value for StatsVersion to be an explicit nil

### UnsetStatsVersion
`func (o *Response23) UnsetStatsVersion()`

UnsetStatsVersion ensures that no value is present for StatsVersion, not even an explicit nil
### GetVariants

`func (o *Response23) GetVariants() []ExperimentVariantTrendsBaseStats`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Response23) GetVariantsOk() (*[]ExperimentVariantTrendsBaseStats, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *Response23) SetVariants(v []ExperimentVariantTrendsBaseStats)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


