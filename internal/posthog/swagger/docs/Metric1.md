# Metric1

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

### NewMetric1

`func NewMetric1(source Source2, series []Series1Inner1, denominator Denominator, numerator Numerator, ) *Metric1`

NewMetric1 instantiates a new Metric1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetric1WithDefaults

`func NewMetric1WithDefaults() *Metric1`

NewMetric1WithDefaults instantiates a new Metric1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownFilter

`func (o *Metric1) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *Metric1) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *Metric1) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *Metric1) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetConversionWindow

`func (o *Metric1) GetConversionWindow() int32`

GetConversionWindow returns the ConversionWindow field if non-nil, zero value otherwise.

### GetConversionWindowOk

`func (o *Metric1) GetConversionWindowOk() (*int32, bool)`

GetConversionWindowOk returns a tuple with the ConversionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindow

`func (o *Metric1) SetConversionWindow(v int32)`

SetConversionWindow sets ConversionWindow field to given value.

### HasConversionWindow

`func (o *Metric1) HasConversionWindow() bool`

HasConversionWindow returns a boolean if a field has been set.

### SetConversionWindowNil

`func (o *Metric1) SetConversionWindowNil(b bool)`

 SetConversionWindowNil sets the value for ConversionWindow to be an explicit nil

### UnsetConversionWindow
`func (o *Metric1) UnsetConversionWindow()`

UnsetConversionWindow ensures that no value is present for ConversionWindow, not even an explicit nil
### GetConversionWindowUnit

`func (o *Metric1) GetConversionWindowUnit() FunnelConversionWindowTimeUnit`

GetConversionWindowUnit returns the ConversionWindowUnit field if non-nil, zero value otherwise.

### GetConversionWindowUnitOk

`func (o *Metric1) GetConversionWindowUnitOk() (*FunnelConversionWindowTimeUnit, bool)`

GetConversionWindowUnitOk returns a tuple with the ConversionWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindowUnit

`func (o *Metric1) SetConversionWindowUnit(v FunnelConversionWindowTimeUnit)`

SetConversionWindowUnit sets ConversionWindowUnit field to given value.

### HasConversionWindowUnit

`func (o *Metric1) HasConversionWindowUnit() bool`

HasConversionWindowUnit returns a boolean if a field has been set.

### GetFingerprint

`func (o *Metric1) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Metric1) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Metric1) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Metric1) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *Metric1) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *Metric1) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetGoal

`func (o *Metric1) GetGoal() ExperimentMetricGoal`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *Metric1) GetGoalOk() (*ExperimentMetricGoal, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *Metric1) SetGoal(v ExperimentMetricGoal)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *Metric1) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetIgnoreZeros

`func (o *Metric1) GetIgnoreZeros() bool`

GetIgnoreZeros returns the IgnoreZeros field if non-nil, zero value otherwise.

### GetIgnoreZerosOk

`func (o *Metric1) GetIgnoreZerosOk() (*bool, bool)`

GetIgnoreZerosOk returns a tuple with the IgnoreZeros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreZeros

`func (o *Metric1) SetIgnoreZeros(v bool)`

SetIgnoreZeros sets IgnoreZeros field to given value.

### HasIgnoreZeros

`func (o *Metric1) HasIgnoreZeros() bool`

HasIgnoreZeros returns a boolean if a field has been set.

### SetIgnoreZerosNil

`func (o *Metric1) SetIgnoreZerosNil(b bool)`

 SetIgnoreZerosNil sets the value for IgnoreZeros to be an explicit nil

### UnsetIgnoreZeros
`func (o *Metric1) UnsetIgnoreZeros()`

UnsetIgnoreZeros ensures that no value is present for IgnoreZeros, not even an explicit nil
### GetIsSharedMetric

`func (o *Metric1) GetIsSharedMetric() bool`

GetIsSharedMetric returns the IsSharedMetric field if non-nil, zero value otherwise.

### GetIsSharedMetricOk

`func (o *Metric1) GetIsSharedMetricOk() (*bool, bool)`

GetIsSharedMetricOk returns a tuple with the IsSharedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSharedMetric

`func (o *Metric1) SetIsSharedMetric(v bool)`

SetIsSharedMetric sets IsSharedMetric field to given value.

### HasIsSharedMetric

`func (o *Metric1) HasIsSharedMetric() bool`

HasIsSharedMetric returns a boolean if a field has been set.

### SetIsSharedMetricNil

`func (o *Metric1) SetIsSharedMetricNil(b bool)`

 SetIsSharedMetricNil sets the value for IsSharedMetric to be an explicit nil

### UnsetIsSharedMetric
`func (o *Metric1) UnsetIsSharedMetric()`

UnsetIsSharedMetric ensures that no value is present for IsSharedMetric, not even an explicit nil
### GetKind

`func (o *Metric1) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Metric1) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Metric1) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Metric1) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLowerBoundPercentile

`func (o *Metric1) GetLowerBoundPercentile() float32`

GetLowerBoundPercentile returns the LowerBoundPercentile field if non-nil, zero value otherwise.

### GetLowerBoundPercentileOk

`func (o *Metric1) GetLowerBoundPercentileOk() (*float32, bool)`

GetLowerBoundPercentileOk returns a tuple with the LowerBoundPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundPercentile

`func (o *Metric1) SetLowerBoundPercentile(v float32)`

SetLowerBoundPercentile sets LowerBoundPercentile field to given value.

### HasLowerBoundPercentile

`func (o *Metric1) HasLowerBoundPercentile() bool`

HasLowerBoundPercentile returns a boolean if a field has been set.

### SetLowerBoundPercentileNil

`func (o *Metric1) SetLowerBoundPercentileNil(b bool)`

 SetLowerBoundPercentileNil sets the value for LowerBoundPercentile to be an explicit nil

### UnsetLowerBoundPercentile
`func (o *Metric1) UnsetLowerBoundPercentile()`

UnsetLowerBoundPercentile ensures that no value is present for LowerBoundPercentile, not even an explicit nil
### GetMetricType

`func (o *Metric1) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *Metric1) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *Metric1) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *Metric1) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetName

`func (o *Metric1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metric1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metric1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Metric1) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Metric1) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Metric1) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResponse

`func (o *Metric1) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Metric1) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Metric1) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Metric1) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *Metric1) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *Metric1) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSharedMetricId

`func (o *Metric1) GetSharedMetricId() float32`

GetSharedMetricId returns the SharedMetricId field if non-nil, zero value otherwise.

### GetSharedMetricIdOk

`func (o *Metric1) GetSharedMetricIdOk() (*float32, bool)`

GetSharedMetricIdOk returns a tuple with the SharedMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMetricId

`func (o *Metric1) SetSharedMetricId(v float32)`

SetSharedMetricId sets SharedMetricId field to given value.

### HasSharedMetricId

`func (o *Metric1) HasSharedMetricId() bool`

HasSharedMetricId returns a boolean if a field has been set.

### SetSharedMetricIdNil

`func (o *Metric1) SetSharedMetricIdNil(b bool)`

 SetSharedMetricIdNil sets the value for SharedMetricId to be an explicit nil

### UnsetSharedMetricId
`func (o *Metric1) UnsetSharedMetricId()`

UnsetSharedMetricId ensures that no value is present for SharedMetricId, not even an explicit nil
### GetSource

`func (o *Metric1) GetSource() Source2`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Metric1) GetSourceOk() (*Source2, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Metric1) SetSource(v Source2)`

SetSource sets Source field to given value.


### GetUpperBoundPercentile

`func (o *Metric1) GetUpperBoundPercentile() float32`

GetUpperBoundPercentile returns the UpperBoundPercentile field if non-nil, zero value otherwise.

### GetUpperBoundPercentileOk

`func (o *Metric1) GetUpperBoundPercentileOk() (*float32, bool)`

GetUpperBoundPercentileOk returns a tuple with the UpperBoundPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundPercentile

`func (o *Metric1) SetUpperBoundPercentile(v float32)`

SetUpperBoundPercentile sets UpperBoundPercentile field to given value.

### HasUpperBoundPercentile

`func (o *Metric1) HasUpperBoundPercentile() bool`

HasUpperBoundPercentile returns a boolean if a field has been set.

### SetUpperBoundPercentileNil

`func (o *Metric1) SetUpperBoundPercentileNil(b bool)`

 SetUpperBoundPercentileNil sets the value for UpperBoundPercentile to be an explicit nil

### UnsetUpperBoundPercentile
`func (o *Metric1) UnsetUpperBoundPercentile()`

UnsetUpperBoundPercentile ensures that no value is present for UpperBoundPercentile, not even an explicit nil
### GetUuid

`func (o *Metric1) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Metric1) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Metric1) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Metric1) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Metric1) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Metric1) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *Metric1) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Metric1) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Metric1) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Metric1) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Metric1) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Metric1) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFunnelOrderType

`func (o *Metric1) GetFunnelOrderType() StepOrderValue`

GetFunnelOrderType returns the FunnelOrderType field if non-nil, zero value otherwise.

### GetFunnelOrderTypeOk

`func (o *Metric1) GetFunnelOrderTypeOk() (*StepOrderValue, bool)`

GetFunnelOrderTypeOk returns a tuple with the FunnelOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelOrderType

`func (o *Metric1) SetFunnelOrderType(v StepOrderValue)`

SetFunnelOrderType sets FunnelOrderType field to given value.

### HasFunnelOrderType

`func (o *Metric1) HasFunnelOrderType() bool`

HasFunnelOrderType returns a boolean if a field has been set.

### GetSeries

`func (o *Metric1) GetSeries() []Series1Inner1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Metric1) GetSeriesOk() (*[]Series1Inner1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Metric1) SetSeries(v []Series1Inner1)`

SetSeries sets Series field to given value.


### GetDenominator

`func (o *Metric1) GetDenominator() Denominator`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *Metric1) GetDenominatorOk() (*Denominator, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *Metric1) SetDenominator(v Denominator)`

SetDenominator sets Denominator field to given value.


### GetNumerator

`func (o *Metric1) GetNumerator() Numerator`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *Metric1) GetNumeratorOk() (*Numerator, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *Metric1) SetNumerator(v Numerator)`

SetNumerator sets Numerator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


