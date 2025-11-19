# ExperimentTrendsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountQuery** | [**TrendsQuery**](TrendsQuery.md) |  | 
**ExperimentId** | Pointer to **NullableInt32** |  | [optional] 
**ExposureQuery** | Pointer to [**TrendsQuery**](TrendsQuery.md) |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentTrendsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to [**ExperimentTrendsQueryResponse**](ExperimentTrendsQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewExperimentTrendsQuery

`func NewExperimentTrendsQuery(countQuery TrendsQuery, ) *ExperimentTrendsQuery`

NewExperimentTrendsQuery instantiates a new ExperimentTrendsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentTrendsQueryWithDefaults

`func NewExperimentTrendsQueryWithDefaults() *ExperimentTrendsQuery`

NewExperimentTrendsQueryWithDefaults instantiates a new ExperimentTrendsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountQuery

`func (o *ExperimentTrendsQuery) GetCountQuery() TrendsQuery`

GetCountQuery returns the CountQuery field if non-nil, zero value otherwise.

### GetCountQueryOk

`func (o *ExperimentTrendsQuery) GetCountQueryOk() (*TrendsQuery, bool)`

GetCountQueryOk returns a tuple with the CountQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountQuery

`func (o *ExperimentTrendsQuery) SetCountQuery(v TrendsQuery)`

SetCountQuery sets CountQuery field to given value.


### GetExperimentId

`func (o *ExperimentTrendsQuery) GetExperimentId() int32`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *ExperimentTrendsQuery) GetExperimentIdOk() (*int32, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *ExperimentTrendsQuery) SetExperimentId(v int32)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *ExperimentTrendsQuery) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### SetExperimentIdNil

`func (o *ExperimentTrendsQuery) SetExperimentIdNil(b bool)`

 SetExperimentIdNil sets the value for ExperimentId to be an explicit nil

### UnsetExperimentId
`func (o *ExperimentTrendsQuery) UnsetExperimentId()`

UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil
### GetExposureQuery

`func (o *ExperimentTrendsQuery) GetExposureQuery() TrendsQuery`

GetExposureQuery returns the ExposureQuery field if non-nil, zero value otherwise.

### GetExposureQueryOk

`func (o *ExperimentTrendsQuery) GetExposureQueryOk() (*TrendsQuery, bool)`

GetExposureQueryOk returns a tuple with the ExposureQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureQuery

`func (o *ExperimentTrendsQuery) SetExposureQuery(v TrendsQuery)`

SetExposureQuery sets ExposureQuery field to given value.

### HasExposureQuery

`func (o *ExperimentTrendsQuery) HasExposureQuery() bool`

HasExposureQuery returns a boolean if a field has been set.

### GetFingerprint

`func (o *ExperimentTrendsQuery) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ExperimentTrendsQuery) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ExperimentTrendsQuery) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ExperimentTrendsQuery) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *ExperimentTrendsQuery) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *ExperimentTrendsQuery) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetKind

`func (o *ExperimentTrendsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentTrendsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentTrendsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentTrendsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *ExperimentTrendsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ExperimentTrendsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ExperimentTrendsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ExperimentTrendsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetName

`func (o *ExperimentTrendsQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentTrendsQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentTrendsQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExperimentTrendsQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExperimentTrendsQuery) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExperimentTrendsQuery) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResponse

`func (o *ExperimentTrendsQuery) GetResponse() ExperimentTrendsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExperimentTrendsQuery) GetResponseOk() (*ExperimentTrendsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExperimentTrendsQuery) SetResponse(v ExperimentTrendsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExperimentTrendsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *ExperimentTrendsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExperimentTrendsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExperimentTrendsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExperimentTrendsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUuid

`func (o *ExperimentTrendsQuery) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExperimentTrendsQuery) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExperimentTrendsQuery) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExperimentTrendsQuery) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ExperimentTrendsQuery) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ExperimentTrendsQuery) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ExperimentTrendsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExperimentTrendsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExperimentTrendsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExperimentTrendsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExperimentTrendsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExperimentTrendsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


