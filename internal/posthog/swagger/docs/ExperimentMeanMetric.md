# ExperimentMeanMetric

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
**MetricType** | Pointer to **string** |  | [optional] [default to "mean"]
**Name** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 
**SharedMetricId** | Pointer to **NullableFloat32** |  | [optional] 
**Source** | [**Source2**](Source2.md) |  | 
**UpperBoundPercentile** | Pointer to **NullableFloat32** |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewExperimentMeanMetric

`func NewExperimentMeanMetric(source Source2, ) *ExperimentMeanMetric`

NewExperimentMeanMetric instantiates a new ExperimentMeanMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentMeanMetricWithDefaults

`func NewExperimentMeanMetricWithDefaults() *ExperimentMeanMetric`

NewExperimentMeanMetricWithDefaults instantiates a new ExperimentMeanMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownFilter

`func (o *ExperimentMeanMetric) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *ExperimentMeanMetric) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *ExperimentMeanMetric) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *ExperimentMeanMetric) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetConversionWindow

`func (o *ExperimentMeanMetric) GetConversionWindow() int32`

GetConversionWindow returns the ConversionWindow field if non-nil, zero value otherwise.

### GetConversionWindowOk

`func (o *ExperimentMeanMetric) GetConversionWindowOk() (*int32, bool)`

GetConversionWindowOk returns a tuple with the ConversionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindow

`func (o *ExperimentMeanMetric) SetConversionWindow(v int32)`

SetConversionWindow sets ConversionWindow field to given value.

### HasConversionWindow

`func (o *ExperimentMeanMetric) HasConversionWindow() bool`

HasConversionWindow returns a boolean if a field has been set.

### SetConversionWindowNil

`func (o *ExperimentMeanMetric) SetConversionWindowNil(b bool)`

 SetConversionWindowNil sets the value for ConversionWindow to be an explicit nil

### UnsetConversionWindow
`func (o *ExperimentMeanMetric) UnsetConversionWindow()`

UnsetConversionWindow ensures that no value is present for ConversionWindow, not even an explicit nil
### GetConversionWindowUnit

`func (o *ExperimentMeanMetric) GetConversionWindowUnit() FunnelConversionWindowTimeUnit`

GetConversionWindowUnit returns the ConversionWindowUnit field if non-nil, zero value otherwise.

### GetConversionWindowUnitOk

`func (o *ExperimentMeanMetric) GetConversionWindowUnitOk() (*FunnelConversionWindowTimeUnit, bool)`

GetConversionWindowUnitOk returns a tuple with the ConversionWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindowUnit

`func (o *ExperimentMeanMetric) SetConversionWindowUnit(v FunnelConversionWindowTimeUnit)`

SetConversionWindowUnit sets ConversionWindowUnit field to given value.

### HasConversionWindowUnit

`func (o *ExperimentMeanMetric) HasConversionWindowUnit() bool`

HasConversionWindowUnit returns a boolean if a field has been set.

### GetFingerprint

`func (o *ExperimentMeanMetric) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ExperimentMeanMetric) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ExperimentMeanMetric) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ExperimentMeanMetric) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *ExperimentMeanMetric) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *ExperimentMeanMetric) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetGoal

`func (o *ExperimentMeanMetric) GetGoal() ExperimentMetricGoal`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *ExperimentMeanMetric) GetGoalOk() (*ExperimentMetricGoal, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *ExperimentMeanMetric) SetGoal(v ExperimentMetricGoal)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *ExperimentMeanMetric) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetIgnoreZeros

`func (o *ExperimentMeanMetric) GetIgnoreZeros() bool`

GetIgnoreZeros returns the IgnoreZeros field if non-nil, zero value otherwise.

### GetIgnoreZerosOk

`func (o *ExperimentMeanMetric) GetIgnoreZerosOk() (*bool, bool)`

GetIgnoreZerosOk returns a tuple with the IgnoreZeros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreZeros

`func (o *ExperimentMeanMetric) SetIgnoreZeros(v bool)`

SetIgnoreZeros sets IgnoreZeros field to given value.

### HasIgnoreZeros

`func (o *ExperimentMeanMetric) HasIgnoreZeros() bool`

HasIgnoreZeros returns a boolean if a field has been set.

### SetIgnoreZerosNil

`func (o *ExperimentMeanMetric) SetIgnoreZerosNil(b bool)`

 SetIgnoreZerosNil sets the value for IgnoreZeros to be an explicit nil

### UnsetIgnoreZeros
`func (o *ExperimentMeanMetric) UnsetIgnoreZeros()`

UnsetIgnoreZeros ensures that no value is present for IgnoreZeros, not even an explicit nil
### GetIsSharedMetric

`func (o *ExperimentMeanMetric) GetIsSharedMetric() bool`

GetIsSharedMetric returns the IsSharedMetric field if non-nil, zero value otherwise.

### GetIsSharedMetricOk

`func (o *ExperimentMeanMetric) GetIsSharedMetricOk() (*bool, bool)`

GetIsSharedMetricOk returns a tuple with the IsSharedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSharedMetric

`func (o *ExperimentMeanMetric) SetIsSharedMetric(v bool)`

SetIsSharedMetric sets IsSharedMetric field to given value.

### HasIsSharedMetric

`func (o *ExperimentMeanMetric) HasIsSharedMetric() bool`

HasIsSharedMetric returns a boolean if a field has been set.

### SetIsSharedMetricNil

`func (o *ExperimentMeanMetric) SetIsSharedMetricNil(b bool)`

 SetIsSharedMetricNil sets the value for IsSharedMetric to be an explicit nil

### UnsetIsSharedMetric
`func (o *ExperimentMeanMetric) UnsetIsSharedMetric()`

UnsetIsSharedMetric ensures that no value is present for IsSharedMetric, not even an explicit nil
### GetKind

`func (o *ExperimentMeanMetric) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentMeanMetric) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentMeanMetric) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentMeanMetric) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLowerBoundPercentile

`func (o *ExperimentMeanMetric) GetLowerBoundPercentile() float32`

GetLowerBoundPercentile returns the LowerBoundPercentile field if non-nil, zero value otherwise.

### GetLowerBoundPercentileOk

`func (o *ExperimentMeanMetric) GetLowerBoundPercentileOk() (*float32, bool)`

GetLowerBoundPercentileOk returns a tuple with the LowerBoundPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundPercentile

`func (o *ExperimentMeanMetric) SetLowerBoundPercentile(v float32)`

SetLowerBoundPercentile sets LowerBoundPercentile field to given value.

### HasLowerBoundPercentile

`func (o *ExperimentMeanMetric) HasLowerBoundPercentile() bool`

HasLowerBoundPercentile returns a boolean if a field has been set.

### SetLowerBoundPercentileNil

`func (o *ExperimentMeanMetric) SetLowerBoundPercentileNil(b bool)`

 SetLowerBoundPercentileNil sets the value for LowerBoundPercentile to be an explicit nil

### UnsetLowerBoundPercentile
`func (o *ExperimentMeanMetric) UnsetLowerBoundPercentile()`

UnsetLowerBoundPercentile ensures that no value is present for LowerBoundPercentile, not even an explicit nil
### GetMetricType

`func (o *ExperimentMeanMetric) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *ExperimentMeanMetric) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *ExperimentMeanMetric) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *ExperimentMeanMetric) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetName

`func (o *ExperimentMeanMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentMeanMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentMeanMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExperimentMeanMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExperimentMeanMetric) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExperimentMeanMetric) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResponse

`func (o *ExperimentMeanMetric) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExperimentMeanMetric) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExperimentMeanMetric) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExperimentMeanMetric) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ExperimentMeanMetric) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ExperimentMeanMetric) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSharedMetricId

`func (o *ExperimentMeanMetric) GetSharedMetricId() float32`

GetSharedMetricId returns the SharedMetricId field if non-nil, zero value otherwise.

### GetSharedMetricIdOk

`func (o *ExperimentMeanMetric) GetSharedMetricIdOk() (*float32, bool)`

GetSharedMetricIdOk returns a tuple with the SharedMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMetricId

`func (o *ExperimentMeanMetric) SetSharedMetricId(v float32)`

SetSharedMetricId sets SharedMetricId field to given value.

### HasSharedMetricId

`func (o *ExperimentMeanMetric) HasSharedMetricId() bool`

HasSharedMetricId returns a boolean if a field has been set.

### SetSharedMetricIdNil

`func (o *ExperimentMeanMetric) SetSharedMetricIdNil(b bool)`

 SetSharedMetricIdNil sets the value for SharedMetricId to be an explicit nil

### UnsetSharedMetricId
`func (o *ExperimentMeanMetric) UnsetSharedMetricId()`

UnsetSharedMetricId ensures that no value is present for SharedMetricId, not even an explicit nil
### GetSource

`func (o *ExperimentMeanMetric) GetSource() Source2`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExperimentMeanMetric) GetSourceOk() (*Source2, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExperimentMeanMetric) SetSource(v Source2)`

SetSource sets Source field to given value.


### GetUpperBoundPercentile

`func (o *ExperimentMeanMetric) GetUpperBoundPercentile() float32`

GetUpperBoundPercentile returns the UpperBoundPercentile field if non-nil, zero value otherwise.

### GetUpperBoundPercentileOk

`func (o *ExperimentMeanMetric) GetUpperBoundPercentileOk() (*float32, bool)`

GetUpperBoundPercentileOk returns a tuple with the UpperBoundPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundPercentile

`func (o *ExperimentMeanMetric) SetUpperBoundPercentile(v float32)`

SetUpperBoundPercentile sets UpperBoundPercentile field to given value.

### HasUpperBoundPercentile

`func (o *ExperimentMeanMetric) HasUpperBoundPercentile() bool`

HasUpperBoundPercentile returns a boolean if a field has been set.

### SetUpperBoundPercentileNil

`func (o *ExperimentMeanMetric) SetUpperBoundPercentileNil(b bool)`

 SetUpperBoundPercentileNil sets the value for UpperBoundPercentile to be an explicit nil

### UnsetUpperBoundPercentile
`func (o *ExperimentMeanMetric) UnsetUpperBoundPercentile()`

UnsetUpperBoundPercentile ensures that no value is present for UpperBoundPercentile, not even an explicit nil
### GetUuid

`func (o *ExperimentMeanMetric) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExperimentMeanMetric) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExperimentMeanMetric) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExperimentMeanMetric) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ExperimentMeanMetric) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ExperimentMeanMetric) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ExperimentMeanMetric) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExperimentMeanMetric) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExperimentMeanMetric) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExperimentMeanMetric) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExperimentMeanMetric) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExperimentMeanMetric) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


