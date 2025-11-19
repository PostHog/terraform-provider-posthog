# WebAnalyticsSampling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**ForceSamplingRate** | Pointer to [**SamplingRate**](SamplingRate.md) |  | [optional] 

## Methods

### NewWebAnalyticsSampling

`func NewWebAnalyticsSampling() *WebAnalyticsSampling`

NewWebAnalyticsSampling instantiates a new WebAnalyticsSampling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAnalyticsSamplingWithDefaults

`func NewWebAnalyticsSamplingWithDefaults() *WebAnalyticsSampling`

NewWebAnalyticsSamplingWithDefaults instantiates a new WebAnalyticsSampling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WebAnalyticsSampling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebAnalyticsSampling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebAnalyticsSampling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebAnalyticsSampling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *WebAnalyticsSampling) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *WebAnalyticsSampling) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetForceSamplingRate

`func (o *WebAnalyticsSampling) GetForceSamplingRate() SamplingRate`

GetForceSamplingRate returns the ForceSamplingRate field if non-nil, zero value otherwise.

### GetForceSamplingRateOk

`func (o *WebAnalyticsSampling) GetForceSamplingRateOk() (*SamplingRate, bool)`

GetForceSamplingRateOk returns a tuple with the ForceSamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSamplingRate

`func (o *WebAnalyticsSampling) SetForceSamplingRate(v SamplingRate)`

SetForceSamplingRate sets ForceSamplingRate field to given value.

### HasForceSamplingRate

`func (o *WebAnalyticsSampling) HasForceSamplingRate() bool`

HasForceSamplingRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


