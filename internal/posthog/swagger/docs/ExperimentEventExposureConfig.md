# ExperimentEventExposureConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentEventExposureConfig"]
**Properties** | [**[]FixedpropertiesInner**](FixedpropertiesInner.md) |  | 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewExperimentEventExposureConfig

`func NewExperimentEventExposureConfig(event string, properties []FixedpropertiesInner, ) *ExperimentEventExposureConfig`

NewExperimentEventExposureConfig instantiates a new ExperimentEventExposureConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentEventExposureConfigWithDefaults

`func NewExperimentEventExposureConfigWithDefaults() *ExperimentEventExposureConfig`

NewExperimentEventExposureConfigWithDefaults instantiates a new ExperimentEventExposureConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *ExperimentEventExposureConfig) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ExperimentEventExposureConfig) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ExperimentEventExposureConfig) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetKind

`func (o *ExperimentEventExposureConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentEventExposureConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentEventExposureConfig) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentEventExposureConfig) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProperties

`func (o *ExperimentEventExposureConfig) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ExperimentEventExposureConfig) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ExperimentEventExposureConfig) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.


### GetResponse

`func (o *ExperimentEventExposureConfig) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExperimentEventExposureConfig) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExperimentEventExposureConfig) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExperimentEventExposureConfig) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ExperimentEventExposureConfig) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ExperimentEventExposureConfig) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetVersion

`func (o *ExperimentEventExposureConfig) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExperimentEventExposureConfig) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExperimentEventExposureConfig) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExperimentEventExposureConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExperimentEventExposureConfig) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExperimentEventExposureConfig) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


