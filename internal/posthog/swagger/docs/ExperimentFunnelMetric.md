# ExperimentFunnelMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakdownFilter** | Pointer to [**BreakdownFilter**](BreakdownFilter.md) |  | [optional] 
**ConversionWindow** | Pointer to **NullableInt32** |  | [optional] 
**ConversionWindowUnit** | Pointer to [**FunnelConversionWindowTimeUnit**](FunnelConversionWindowTimeUnit.md) |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**FunnelOrderType** | Pointer to [**StepOrderValue**](StepOrderValue.md) |  | [optional] 
**Goal** | Pointer to [**ExperimentMetricGoal**](ExperimentMetricGoal.md) |  | [optional] 
**IsSharedMetric** | Pointer to **NullableBool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentMetric"]
**MetricType** | Pointer to **string** |  | [optional] [default to "funnel"]
**Name** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 
**Series** | [**[]Series1Inner1**](Series1Inner1.md) |  | 
**SharedMetricId** | Pointer to **NullableFloat32** |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewExperimentFunnelMetric

`func NewExperimentFunnelMetric(series []Series1Inner1, ) *ExperimentFunnelMetric`

NewExperimentFunnelMetric instantiates a new ExperimentFunnelMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentFunnelMetricWithDefaults

`func NewExperimentFunnelMetricWithDefaults() *ExperimentFunnelMetric`

NewExperimentFunnelMetricWithDefaults instantiates a new ExperimentFunnelMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdownFilter

`func (o *ExperimentFunnelMetric) GetBreakdownFilter() BreakdownFilter`

GetBreakdownFilter returns the BreakdownFilter field if non-nil, zero value otherwise.

### GetBreakdownFilterOk

`func (o *ExperimentFunnelMetric) GetBreakdownFilterOk() (*BreakdownFilter, bool)`

GetBreakdownFilterOk returns a tuple with the BreakdownFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownFilter

`func (o *ExperimentFunnelMetric) SetBreakdownFilter(v BreakdownFilter)`

SetBreakdownFilter sets BreakdownFilter field to given value.

### HasBreakdownFilter

`func (o *ExperimentFunnelMetric) HasBreakdownFilter() bool`

HasBreakdownFilter returns a boolean if a field has been set.

### GetConversionWindow

`func (o *ExperimentFunnelMetric) GetConversionWindow() int32`

GetConversionWindow returns the ConversionWindow field if non-nil, zero value otherwise.

### GetConversionWindowOk

`func (o *ExperimentFunnelMetric) GetConversionWindowOk() (*int32, bool)`

GetConversionWindowOk returns a tuple with the ConversionWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindow

`func (o *ExperimentFunnelMetric) SetConversionWindow(v int32)`

SetConversionWindow sets ConversionWindow field to given value.

### HasConversionWindow

`func (o *ExperimentFunnelMetric) HasConversionWindow() bool`

HasConversionWindow returns a boolean if a field has been set.

### SetConversionWindowNil

`func (o *ExperimentFunnelMetric) SetConversionWindowNil(b bool)`

 SetConversionWindowNil sets the value for ConversionWindow to be an explicit nil

### UnsetConversionWindow
`func (o *ExperimentFunnelMetric) UnsetConversionWindow()`

UnsetConversionWindow ensures that no value is present for ConversionWindow, not even an explicit nil
### GetConversionWindowUnit

`func (o *ExperimentFunnelMetric) GetConversionWindowUnit() FunnelConversionWindowTimeUnit`

GetConversionWindowUnit returns the ConversionWindowUnit field if non-nil, zero value otherwise.

### GetConversionWindowUnitOk

`func (o *ExperimentFunnelMetric) GetConversionWindowUnitOk() (*FunnelConversionWindowTimeUnit, bool)`

GetConversionWindowUnitOk returns a tuple with the ConversionWindowUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionWindowUnit

`func (o *ExperimentFunnelMetric) SetConversionWindowUnit(v FunnelConversionWindowTimeUnit)`

SetConversionWindowUnit sets ConversionWindowUnit field to given value.

### HasConversionWindowUnit

`func (o *ExperimentFunnelMetric) HasConversionWindowUnit() bool`

HasConversionWindowUnit returns a boolean if a field has been set.

### GetFingerprint

`func (o *ExperimentFunnelMetric) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ExperimentFunnelMetric) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ExperimentFunnelMetric) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ExperimentFunnelMetric) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *ExperimentFunnelMetric) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *ExperimentFunnelMetric) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetFunnelOrderType

`func (o *ExperimentFunnelMetric) GetFunnelOrderType() StepOrderValue`

GetFunnelOrderType returns the FunnelOrderType field if non-nil, zero value otherwise.

### GetFunnelOrderTypeOk

`func (o *ExperimentFunnelMetric) GetFunnelOrderTypeOk() (*StepOrderValue, bool)`

GetFunnelOrderTypeOk returns a tuple with the FunnelOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelOrderType

`func (o *ExperimentFunnelMetric) SetFunnelOrderType(v StepOrderValue)`

SetFunnelOrderType sets FunnelOrderType field to given value.

### HasFunnelOrderType

`func (o *ExperimentFunnelMetric) HasFunnelOrderType() bool`

HasFunnelOrderType returns a boolean if a field has been set.

### GetGoal

`func (o *ExperimentFunnelMetric) GetGoal() ExperimentMetricGoal`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *ExperimentFunnelMetric) GetGoalOk() (*ExperimentMetricGoal, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *ExperimentFunnelMetric) SetGoal(v ExperimentMetricGoal)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *ExperimentFunnelMetric) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetIsSharedMetric

`func (o *ExperimentFunnelMetric) GetIsSharedMetric() bool`

GetIsSharedMetric returns the IsSharedMetric field if non-nil, zero value otherwise.

### GetIsSharedMetricOk

`func (o *ExperimentFunnelMetric) GetIsSharedMetricOk() (*bool, bool)`

GetIsSharedMetricOk returns a tuple with the IsSharedMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSharedMetric

`func (o *ExperimentFunnelMetric) SetIsSharedMetric(v bool)`

SetIsSharedMetric sets IsSharedMetric field to given value.

### HasIsSharedMetric

`func (o *ExperimentFunnelMetric) HasIsSharedMetric() bool`

HasIsSharedMetric returns a boolean if a field has been set.

### SetIsSharedMetricNil

`func (o *ExperimentFunnelMetric) SetIsSharedMetricNil(b bool)`

 SetIsSharedMetricNil sets the value for IsSharedMetric to be an explicit nil

### UnsetIsSharedMetric
`func (o *ExperimentFunnelMetric) UnsetIsSharedMetric()`

UnsetIsSharedMetric ensures that no value is present for IsSharedMetric, not even an explicit nil
### GetKind

`func (o *ExperimentFunnelMetric) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentFunnelMetric) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentFunnelMetric) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentFunnelMetric) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetricType

`func (o *ExperimentFunnelMetric) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *ExperimentFunnelMetric) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *ExperimentFunnelMetric) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *ExperimentFunnelMetric) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetName

`func (o *ExperimentFunnelMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentFunnelMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentFunnelMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExperimentFunnelMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExperimentFunnelMetric) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExperimentFunnelMetric) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResponse

`func (o *ExperimentFunnelMetric) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExperimentFunnelMetric) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExperimentFunnelMetric) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExperimentFunnelMetric) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ExperimentFunnelMetric) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ExperimentFunnelMetric) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetSeries

`func (o *ExperimentFunnelMetric) GetSeries() []Series1Inner1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ExperimentFunnelMetric) GetSeriesOk() (*[]Series1Inner1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ExperimentFunnelMetric) SetSeries(v []Series1Inner1)`

SetSeries sets Series field to given value.


### GetSharedMetricId

`func (o *ExperimentFunnelMetric) GetSharedMetricId() float32`

GetSharedMetricId returns the SharedMetricId field if non-nil, zero value otherwise.

### GetSharedMetricIdOk

`func (o *ExperimentFunnelMetric) GetSharedMetricIdOk() (*float32, bool)`

GetSharedMetricIdOk returns a tuple with the SharedMetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedMetricId

`func (o *ExperimentFunnelMetric) SetSharedMetricId(v float32)`

SetSharedMetricId sets SharedMetricId field to given value.

### HasSharedMetricId

`func (o *ExperimentFunnelMetric) HasSharedMetricId() bool`

HasSharedMetricId returns a boolean if a field has been set.

### SetSharedMetricIdNil

`func (o *ExperimentFunnelMetric) SetSharedMetricIdNil(b bool)`

 SetSharedMetricIdNil sets the value for SharedMetricId to be an explicit nil

### UnsetSharedMetricId
`func (o *ExperimentFunnelMetric) UnsetSharedMetricId()`

UnsetSharedMetricId ensures that no value is present for SharedMetricId, not even an explicit nil
### GetUuid

`func (o *ExperimentFunnelMetric) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExperimentFunnelMetric) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExperimentFunnelMetric) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExperimentFunnelMetric) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ExperimentFunnelMetric) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ExperimentFunnelMetric) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ExperimentFunnelMetric) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExperimentFunnelMetric) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExperimentFunnelMetric) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExperimentFunnelMetric) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExperimentFunnelMetric) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExperimentFunnelMetric) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


