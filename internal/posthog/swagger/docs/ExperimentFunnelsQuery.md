# ExperimentFunnelsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperimentId** | Pointer to **NullableInt32** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**FunnelsQuery** | [**FunnelsQuery**](FunnelsQuery.md) |  | 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentFunnelsQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to [**ExperimentFunnelsQueryResponse**](ExperimentFunnelsQueryResponse.md) |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewExperimentFunnelsQuery

`func NewExperimentFunnelsQuery(funnelsQuery FunnelsQuery, ) *ExperimentFunnelsQuery`

NewExperimentFunnelsQuery instantiates a new ExperimentFunnelsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentFunnelsQueryWithDefaults

`func NewExperimentFunnelsQueryWithDefaults() *ExperimentFunnelsQuery`

NewExperimentFunnelsQueryWithDefaults instantiates a new ExperimentFunnelsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimentId

`func (o *ExperimentFunnelsQuery) GetExperimentId() int32`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *ExperimentFunnelsQuery) GetExperimentIdOk() (*int32, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *ExperimentFunnelsQuery) SetExperimentId(v int32)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *ExperimentFunnelsQuery) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### SetExperimentIdNil

`func (o *ExperimentFunnelsQuery) SetExperimentIdNil(b bool)`

 SetExperimentIdNil sets the value for ExperimentId to be an explicit nil

### UnsetExperimentId
`func (o *ExperimentFunnelsQuery) UnsetExperimentId()`

UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil
### GetFingerprint

`func (o *ExperimentFunnelsQuery) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ExperimentFunnelsQuery) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ExperimentFunnelsQuery) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ExperimentFunnelsQuery) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *ExperimentFunnelsQuery) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *ExperimentFunnelsQuery) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetFunnelsQuery

`func (o *ExperimentFunnelsQuery) GetFunnelsQuery() FunnelsQuery`

GetFunnelsQuery returns the FunnelsQuery field if non-nil, zero value otherwise.

### GetFunnelsQueryOk

`func (o *ExperimentFunnelsQuery) GetFunnelsQueryOk() (*FunnelsQuery, bool)`

GetFunnelsQueryOk returns a tuple with the FunnelsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelsQuery

`func (o *ExperimentFunnelsQuery) SetFunnelsQuery(v FunnelsQuery)`

SetFunnelsQuery sets FunnelsQuery field to given value.


### GetKind

`func (o *ExperimentFunnelsQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentFunnelsQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentFunnelsQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentFunnelsQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *ExperimentFunnelsQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ExperimentFunnelsQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ExperimentFunnelsQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ExperimentFunnelsQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetName

`func (o *ExperimentFunnelsQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentFunnelsQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentFunnelsQuery) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExperimentFunnelsQuery) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExperimentFunnelsQuery) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExperimentFunnelsQuery) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetResponse

`func (o *ExperimentFunnelsQuery) GetResponse() ExperimentFunnelsQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExperimentFunnelsQuery) GetResponseOk() (*ExperimentFunnelsQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExperimentFunnelsQuery) SetResponse(v ExperimentFunnelsQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExperimentFunnelsQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetTags

`func (o *ExperimentFunnelsQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExperimentFunnelsQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExperimentFunnelsQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExperimentFunnelsQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUuid

`func (o *ExperimentFunnelsQuery) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExperimentFunnelsQuery) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExperimentFunnelsQuery) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExperimentFunnelsQuery) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ExperimentFunnelsQuery) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ExperimentFunnelsQuery) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ExperimentFunnelsQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExperimentFunnelsQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExperimentFunnelsQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExperimentFunnelsQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExperimentFunnelsQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExperimentFunnelsQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


