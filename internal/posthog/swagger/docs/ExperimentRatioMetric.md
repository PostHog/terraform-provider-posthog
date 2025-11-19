# ExperimentRatioMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakdownFilter** | Pointer to [**BreakdownFilter**](BreakdownFilter.md) |  | [optional] 
**ConversionWindow** | Pointer to **NullableInt32** |  | [optional] 
**ConversionWindowUnit** | Pointer to [**FunnelConversionWindowTimeUnit**](FunnelConversionWindowTimeUnit.md) |  | [optional] 
**Denominator** | [**Denominator**](Denominator.md) |  | 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**Goal** | Pointer to [**ExperimentMetricGoal**](ExperimentMetricGoal.md) |  | [optional] 
**IsSharedMetric** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentMetric"]
**MetricType** | Pointer to **string** |  | [optional] [default to "ratio"]
**Name** | Pointer to **NullableString** |  | [optional] 
**Numerator** | [**Numerator**](Numerator.md) |  | 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 
**SharedMetricId** | Pointer to **NullableFloat32** |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewExperimentRatioMetric

`func NewExperimentRatioMetric(denominator Denominator, numerator Numerator, ) *ExperimentRatioMetric`

NewExperimentRatioMetric instantiates a new ExperimentRatioMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentRatioMetricWithDefaults

`func NewExperimentRatioMetricWithDefaults() *ExperimentRatioMetric`

NewExperimentRatioMetricWithDefaults instantiates a new ExperimentRatioMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownFilter

`func (o *ExperimentRatioMetric) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *ExperimentRatioMetric) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *ExperimentRatioMetric) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *ExperimentRatioMetric) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetConversionWindow

`func (o *ExperimentRatioMetric) GetConversionWindow() int32`

GetConversionWindow returns the ConversionWindow field if non-nil, zero value otherwise.

### GetConversionWindowOk

`func (o *ExperimentRatioMetric) GetConversionWindowOk() (*int32, bool)`

GetConversionWindowOk returns a tuple with the ConversionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindow

`func (o *ExperimentRatioMetric) SetConversionWindow(v int32)`

SetConversionWindow sets ConversionWindow field to given value.

### HasConversionWindow

`func (o *ExperimentRatioMetric) HasConversionWindow() bool`

HasConversionWindow returns a boolean if a field has been set.

### SetConversionWindowNil

`func (o *ExperimentRatioMetric) SetConversionWindowNil(b bool)`

 SetConversionWindowNil sets the value for ConversionWindow to be an explicit nil

### UnsetConversionWindow
`func (o *ExperimentRatioMetric) UnsetConversionWindow()`

UnsetConversionWindow ensures that no value is present for ConversionWindow, not even an explicit nil
### GetConversionWindowUnit

`func (o *ExperimentRatioMetric) GetConversionWindowUnit() FunnelConversionWindowTimeUnit`

GetConversionWindowUnit returns the ConversionWindowUnit field if non-nil, zero value otherwise.

### GetConversionWindowUnitOk

`func (o *ExperimentRatioMetric) GetConversionWindowUnitOk() (*FunnelConversionWindowTimeUnit, bool)`

GetConversionWindowUnitOk returns a tuple with the ConversionWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindowUnit

`func (o *ExperimentRatioMetric) SetConversionWindowUnit(v FunnelConversionWindowTimeUnit)`

SetConversionWindowUnit sets ConversionWindowUnit field to given value.

### HasConversionWindowUnit

`func (o *ExperimentRatioMetric) HasConversionWindowUnit() bool`

HasConversionWindowUnit returns a boolean if a field has been set.

### GetDenominator

`func (o *ExperimentRatioMetric) GetDenominator() Denominator`

GetDenominator returns the Denominator field if non-nil, zero value otherwise.

### GetDenominatorOk

`func (o *ExperimentRatioMetric) GetDenominatorOk() (*Denominator, bool)`

GetDenominatorOk returns a tuple with the Denominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominator

`func (o *ExperimentRatioMetric) SetDenominator(v Denominator)`

SetDenominator sets Denominator field to given value.


### GetFingerprint

`func (o *ExperimentRatioMetric) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ExperimentRatioMetric) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ExperimentRatioMetric) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ExperimentRatioMetric) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *ExperimentRatioMetric) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *ExperimentRatioMetric) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetGoal

`func (o *ExperimentRatioMetric) GetGoal() ExperimentMetricGoal`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *ExperimentRatioMetric) GetGoalOk() (*ExperimentMetricGoal, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *ExperimentRatioMetric) SetGoal(v ExperimentMetricGoal)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *ExperimentRatioMetric) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetIsSharedMetric

`func (o *ExperimentRatioMetric) GetIsSharedMetric() bool`

GetIsSharedMetric returns the IsSharedMetric field if non-nil, zero value otherwise.

### GetIsSharedMetricOk

`func (o *ExperimentRatioMetric) GetIsSharedMetricOk() (*bool, bool)`

GetIsSharedMetricOk returns a tuple with the IsSharedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSharedMetric

`func (o *ExperimentRatioMetric) SetIsSharedMetric(v bool)`

SetIsSharedMetric sets IsSharedMetric field to given value.

### HasIsSharedMetric

`func (o *ExperimentRatioMetric) HasIsSharedMetric() bool`

HasIsSharedMetric returns a boolean if a field has been set.

### SetIsSharedMetricNil

`func (o *ExperimentRatioMetric) SetIsSharedMetricNil(b bool)`

 SetIsSharedMetricNil sets the value for IsSharedMetric to be an explicit nil

### UnsetIsSharedMetric
`func (o *ExperimentRatioMetric) UnsetIsSharedMetric()`

UnsetIsSharedMetric ensures that no value is present for IsSharedMetric, not even an explicit nil
### GetKind

`func (o *ExperimentRatioMetric) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentRatioMetric) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentRatioMetric) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentRatioMetric) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetricType

`func (o *ExperimentRatioMetric) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *ExperimentRatioMetric) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *ExperimentRatioMetric) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *ExperimentRatioMetric) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetName

`func (o *ExperimentRatioMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentRatioMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentRatioMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExperimentRatioMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExperimentRatioMetric) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExperimentRatioMetric) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumerator

`func (o *ExperimentRatioMetric) GetNumerator() Numerator`

GetNumerator returns the Numerator field if non-nil, zero value otherwise.

### GetNumeratorOk

`func (o *ExperimentRatioMetric) GetNumeratorOk() (*Numerator, bool)`

GetNumeratorOk returns a tuple with the Numerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumerator

`func (o *ExperimentRatioMetric) SetNumerator(v Numerator)`

SetNumerator sets Numerator field to given value.


### GetResponse

`func (o *ExperimentRatioMetric) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExperimentRatioMetric) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExperimentRatioMetric) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExperimentRatioMetric) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ExperimentRatioMetric) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ExperimentRatioMetric) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSharedMetricId

`func (o *ExperimentRatioMetric) GetSharedMetricId() float32`

GetSharedMetricId returns the SharedMetricId field if non-nil, zero value otherwise.

### GetSharedMetricIdOk

`func (o *ExperimentRatioMetric) GetSharedMetricIdOk() (*float32, bool)`

GetSharedMetricIdOk returns a tuple with the SharedMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMetricId

`func (o *ExperimentRatioMetric) SetSharedMetricId(v float32)`

SetSharedMetricId sets SharedMetricId field to given value.

### HasSharedMetricId

`func (o *ExperimentRatioMetric) HasSharedMetricId() bool`

HasSharedMetricId returns a boolean if a field has been set.

### SetSharedMetricIdNil

`func (o *ExperimentRatioMetric) SetSharedMetricIdNil(b bool)`

 SetSharedMetricIdNil sets the value for SharedMetricId to be an explicit nil

### UnsetSharedMetricId
`func (o *ExperimentRatioMetric) UnsetSharedMetricId()`

UnsetSharedMetricId ensures that no value is present for SharedMetricId, not even an explicit nil
### GetUuid

`func (o *ExperimentRatioMetric) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExperimentRatioMetric) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExperimentRatioMetric) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExperimentRatioMetric) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ExperimentRatioMetric) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ExperimentRatioMetric) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ExperimentRatioMetric) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExperimentRatioMetric) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExperimentRatioMetric) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExperimentRatioMetric) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExperimentRatioMetric) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExperimentRatioMetric) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


