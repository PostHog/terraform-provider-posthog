# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakdownFilter** | Pointer to [**BreakdownFilter**](BreakdownFilter.md) |  | [optional] 
**ConversionWindow** | Pointer to **NullableInt32** |  | [optional] 
**ConversionWindowUnit** | Pointer to [**FunnelConversionWindowTimeUnit**](FunnelConversionWindowTimeUnit.md) |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**Goal** | Pointer to [**ExperimentMetricGoal**](ExperimentMetricGoal.md) |  | [optional] 
**IgnoreZeros** | Pointer to **NullableBool** |  | [optional] 
**IsSharedMetric** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentMetric"]
**LowerBoundPercentile** | Pointer to **NullableFloat32** |  | [optional] 
**MetricType** | Pointer to **string** |  | [optional] [default to "ratio"]
**Name** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 
**SharedMetricId** | Pointer to **NullableFloat32** |  | [optional] 
**Source** | [**Source2**](Source2.md) |  | 
**UpperBoundPercentile** | Pointer to **NullableFloat32** |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 
**FunnelOrderType** | Pointer to [**StepOrderValue**](StepOrderValue.md) |  | [optional] 
**Series** | [**[]Series1Inner1**](Series1Inner1.md) |  | 
**Denominator** | [**Denominator**](Denominator.md) |  | 
**Numerator** | [**Numerator**](Numerator.md) |  | 

## Methods

### NewMetric

`func NewMetric(source Source2, series []Series1Inner1, denominator Denominator, numerator Numerator, ) *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownFilter

`func (o *Metric) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *Metric) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *Metric) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *Metric) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetConversionWindow

`func (o *Metric) GetConversionWindow() int32`

GetConversionWindow returns the ConversionWindow field if non-nil, zero value otherwise.

### GetConversionWindowOk

`func (o *Metric) GetConversionWindowOk() (*int32, bool)`

GetConversionWindowOk returns a tuple with the ConversionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindow

`func (o *Metric) SetConversionWindow(v int32)`

SetConversionWindow sets ConversionWindow field to given value.

### HasConversionWindow

`func (o *Metric) HasConversionWindow() bool`

HasConversionWindow returns a boolean if a field has been set.

### SetConversionWindowNil

`func (o *Metric) SetConversionWindowNil(b bool)`

 SetConversionWindowNil sets the value for ConversionWindow to be an explicit nil

### UnsetConversionWindow
`func (o *Metric) UnsetConversionWindow()`

UnsetConversionWindow ensures that no value is present for ConversionWindow, not even an explicit nil
### GetConversionWindowUnit

`func (o *Metric) GetConversionWindowUnit() FunnelConversionWindowTimeUnit`

GetConversionWindowUnit returns the ConversionWindowUnit field if non-nil, zero value otherwise.

### GetConversionWindowUnitOk

`func (o *Metric) GetConversionWindowUnitOk() (*FunnelConversionWindowTimeUnit, bool)`

GetConversionWindowUnitOk returns a tuple with the ConversionWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindowUnit

`func (o *Metric) SetConversionWindowUnit(v FunnelConversionWindowTimeUnit)`

SetConversionWindowUnit sets ConversionWindowUnit field to given value.

### HasConversionWindowUnit

`func (o *Metric) HasConversionWindowUnit() bool`

HasConversionWindowUnit returns a boolean if a field has been set.

### GetFingerprint

`func (o *Metric) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Metric) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Metric) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Metric) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *Metric) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Metric) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetGoal

`func (o *Metric) GetGoal() ExperimentMetricGoal`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *Metric) GetGoalOk() (*ExperimentMetricGoal, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *Metric) SetGoal(v ExperimentMetricGoal)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *Metric) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetIgnoreZeros

`func (o *Metric) GetIgnoreZeros() bool`

GetIgnoreZeros returns the IgnoreZeros field if non-nil, zero value otherwise.

### GetIgnoreZerosOk

`func (o *Metric) GetIgnoreZerosOk() (*bool, bool)`

GetIgnoreZerosOk returns a tuple with the IgnoreZeros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreZeros

`func (o *Metric) SetIgnoreZeros(v bool)`

SetIgnoreZeros sets IgnoreZeros field to given value.

### HasIgnoreZeros

`func (o *Metric) HasIgnoreZeros() bool`

HasIgnoreZeros returns a boolean if a field has been set.

### SetIgnoreZerosNil

`func (o *Metric) SetIgnoreZerosNil(b bool)`

 SetIgnoreZerosNil sets the value for IgnoreZeros to be an explicit nil

### UnsetIgnoreZeros
`func (o *Metric) UnsetIgnoreZeros()`

UnsetIgnoreZeros ensures that no value is present for IgnoreZeros, not even an explicit nil
### GetIsSharedMetric

`func (o *Metric) GetIsSharedMetric() bool`

GetIsSharedMetric returns the IsSharedMetric field if non-nil, zero value otherwise.

### GetIsSharedMetricOk

`func (o *Metric) GetIsSharedMetricOk() (*bool, bool)`

GetIsSharedMetricOk returns a tuple with the IsSharedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSharedMetric

`func (o *Metric) SetIsSharedMetric(v bool)`

SetIsSharedMetric sets IsSharedMetric field to given value.

### HasIsSharedMetric

`func (o *Metric) HasIsSharedMetric() bool`

HasIsSharedMetric returns a boolean if a field has been set.

### SetIsSharedMetricNil

`func (o *Metric) SetIsSharedMetricNil(b bool)`

 SetIsSharedMetricNil sets the value for IsSharedMetric to be an explicit nil

### UnsetIsSharedMetric
`func (o *Metric) UnsetIsSharedMetric()`

UnsetIsSharedMetric ensures that no value is present for IsSharedMetric, not even an explicit nil
### GetKind

`func (o *Metric) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Metric) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Metric) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Metric) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLowerBoundPercentile

`func (o *Metric) GetLowerBoundPercentile() float32`

GetLowerBoundPercentile returns the LowerBoundPercentile field if non-nil, zero value otherwise.

### GetLowerBoundPercentileOk

`func (o *Metric) GetLowerBoundPercentileOk() (*float32, bool)`

GetLowerBoundPercentileOk returns a tuple with the LowerBoundPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundPercentile

`func (o *Metric) SetLowerBoundPercentile(v float32)`

SetLowerBoundPercentile sets LowerBoundPercentile field to given value.

### HasLowerBoundPercentile

`func (o *Metric) HasLowerBoundPercentile() bool`

HasLowerBoundPercentile returns a boolean if a field has been set.

### SetLowerBoundPercentileNil

`func (o *Metric) SetLowerBoundPercentileNil(b bool)`

 SetLowerBoundPercentileNil sets the value for LowerBoundPercentile to be an explicit nil

### UnsetLowerBoundPercentile
`func (o *Metric) UnsetLowerBoundPercentile()`

UnsetLowerBoundPercentile ensures that no value is present for LowerBoundPercentile, not even an explicit nil
### GetMetricType

`func (o *Metric) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *Metric) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *Metric) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *Metric) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetName

`func (o *Metric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Metric) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Metric) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Metric) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResponse

`func (o *Metric) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Metric) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Metric) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Metric) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *Metric) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *Metric) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSharedMetricId

`func (o *Metric) GetSharedMetricId() float32`

GetSharedMetricId returns the SharedMetricId field if non-nil, zero value otherwise.

### GetSharedMetricIdOk

`func (o *Metric) GetSharedMetricIdOk() (*float32, bool)`

GetSharedMetricIdOk returns a tuple with the SharedMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMetricId

`func (o *Metric) SetSharedMetricId(v float32)`

SetSharedMetricId sets SharedMetricId field to given value.

### HasSharedMetricId

`func (o *Metric) HasSharedMetricId() bool`

HasSharedMetricId returns a boolean if a field has been set.

### SetSharedMetricIdNil

`func (o *Metric) SetSharedMetricIdNil(b bool)`

 SetSharedMetricIdNil sets the value for SharedMetricId to be an explicit nil

### UnsetSharedMetricId
`func (o *Metric) UnsetSharedMetricId()`

UnsetSharedMetricId ensures that no value is present for SharedMetricId, not even an explicit nil
### GetSource

`func (o *Metric) GetSource() Source2`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Metric) GetSourceOk() (*Source2, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Metric) SetSource(v Source2)`

SetSource sets Source field to given value.


### GetUpperBoundPercentile

`func (o *Metric) GetUpperBoundPercentile() float32`

GetUpperBoundPercentile returns the UpperBoundPercentile field if non-nil, zero value otherwise.

### GetUpperBoundPercentileOk

`func (o *Metric) GetUpperBoundPercentileOk() (*float32, bool)`

GetUpperBoundPercentileOk returns a tuple with the UpperBoundPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundPercentile

`func (o *Metric) SetUpperBoundPercentile(v float32)`

SetUpperBoundPercentile sets UpperBoundPercentile field to given value.

### HasUpperBoundPercentile

`func (o *Metric) HasUpperBoundPercentile() bool`

HasUpperBoundPercentile returns a boolean if a field has been set.

### SetUpperBoundPercentileNil

`func (o *Metric) SetUpperBoundPercentileNil(b bool)`

 SetUpperBoundPercentileNil sets the value for UpperBoundPercentile to be an explicit nil

### UnsetUpperBoundPercentile
`func (o *Metric) UnsetUpperBoundPercentile()`

UnsetUpperBoundPercentile ensures that no value is present for UpperBoundPercentile, not even an explicit nil
### GetUuid

`func (o *Metric) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Metric) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Metric) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Metric) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Metric) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Metric) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *Metric) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Metric) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Metric) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Metric) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Metric) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Metric) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFunnelOrderType

`func (o *Metric) GetFunnelOrderType() StepOrderValue`

GetFunnelOrderType returns the FunnelOrderType field if non-nil, zero value otherwise.

### GetFunnelOrderTypeOk

`func (o *Metric) GetFunnelOrderTypeOk() (*StepOrderValue, bool)`

GetFunnelOrderTypeOk returns a tuple with the FunnelOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelOrderType

`func (o *Metric) SetFunnelOrderType(v StepOrderValue)`

SetFunnelOrderType sets FunnelOrderType field to given value.

### HasFunnelOrderType

`func (o *Metric) HasFunnelOrderType() bool`

HasFunnelOrderType returns a boolean if a field has been set.

### GetSeries

`func (o *Metric) GetSeries() []Series1Inner1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Metric) GetSeriesOk() (*[]Series1Inner1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Metric) SetSeries(v []Series1Inner1)`

SetSeries sets Series field to given value.


### GetDenominator

`func (o *Metric) GetDenominator() Denominator`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *Metric) GetDenominatorOk() (*Denominator, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *Metric) SetDenominator(v Denominator)`

SetDenominator sets Denominator field to given value.


### GetNumerator

`func (o *Metric) GetNumerator() Numerator`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *Metric) GetNumeratorOk() (*Numerator, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *Metric) SetNumerator(v Numerator)`

SetNumerator sets Numerator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


