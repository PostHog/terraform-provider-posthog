# MinimalFeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**TeamId** | **int32** |  | [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Key** | **string** |  | 
**Filters** | Pointer to **map[string]interface{}** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**EnsureExperienceContinuity** | Pointer to **NullableBool** |  | [optional] 
**HasEncryptedPayloads** | Pointer to **NullableBool** |  | [optional] 
**Version** | Pointer to **NullableInt32** |  | [optional] 
**EvaluationRuntime** | Pointer to **NullableString** | Specifies where this feature flag should be evaluated  * &#x60;server&#x60; - Server * &#x60;client&#x60; - Client * &#x60;all&#x60; - All | [optional] 
**EvaluationTags** | **[]string** |  | [readonly] 

## Methods

### NewMinimalFeatureFlag

`func NewMinimalFeatureFlag(id int32, teamId int32, key string, evaluationTags []string, ) *MinimalFeatureFlag`

NewMinimalFeatureFlag instantiates a new MinimalFeatureFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalFeatureFlagWithDefaults

`func NewMinimalFeatureFlagWithDefaults() *MinimalFeatureFlag`

NewMinimalFeatureFlagWithDefaults instantiates a new MinimalFeatureFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimalFeatureFlag) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimalFeatureFlag) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimalFeatureFlag) SetId(v int32)`

SetId sets Id field to given value.


### GetTeamId

`func (o *MinimalFeatureFlag) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MinimalFeatureFlag) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MinimalFeatureFlag) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetName

`func (o *MinimalFeatureFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MinimalFeatureFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MinimalFeatureFlag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MinimalFeatureFlag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *MinimalFeatureFlag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MinimalFeatureFlag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MinimalFeatureFlag) SetKey(v string)`

SetKey sets Key field to given value.


### GetFilters

`func (o *MinimalFeatureFlag) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MinimalFeatureFlag) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MinimalFeatureFlag) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MinimalFeatureFlag) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetDeleted

`func (o *MinimalFeatureFlag) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MinimalFeatureFlag) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MinimalFeatureFlag) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MinimalFeatureFlag) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetActive

`func (o *MinimalFeatureFlag) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MinimalFeatureFlag) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MinimalFeatureFlag) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MinimalFeatureFlag) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnsureExperienceContinuity

`func (o *MinimalFeatureFlag) GetEnsureExperienceContinuity() bool`

GetEnsureExperienceContinuity returns the EnsureExperienceContinuity field if non-nil, zero value otherwise.

### GetEnsureExperienceContinuityOk

`func (o *MinimalFeatureFlag) GetEnsureExperienceContinuityOk() (*bool, bool)`

GetEnsureExperienceContinuityOk returns a tuple with the EnsureExperienceContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureExperienceContinuity

`func (o *MinimalFeatureFlag) SetEnsureExperienceContinuity(v bool)`

SetEnsureExperienceContinuity sets EnsureExperienceContinuity field to given value.

### HasEnsureExperienceContinuity

`func (o *MinimalFeatureFlag) HasEnsureExperienceContinuity() bool`

HasEnsureExperienceContinuity returns a boolean if a field has been set.

### SetEnsureExperienceContinuityNil

`func (o *MinimalFeatureFlag) SetEnsureExperienceContinuityNil(b bool)`

 SetEnsureExperienceContinuityNil sets the value for EnsureExperienceContinuity to be an explicit nil

### UnsetEnsureExperienceContinuity
`func (o *MinimalFeatureFlag) UnsetEnsureExperienceContinuity()`

UnsetEnsureExperienceContinuity ensures that no value is present for EnsureExperienceContinuity, not even an explicit nil
### GetHasEncryptedPayloads

`func (o *MinimalFeatureFlag) GetHasEncryptedPayloads() bool`

GetHasEncryptedPayloads returns the HasEncryptedPayloads field if non-nil, zero value otherwise.

### GetHasEncryptedPayloadsOk

`func (o *MinimalFeatureFlag) GetHasEncryptedPayloadsOk() (*bool, bool)`

GetHasEncryptedPayloadsOk returns a tuple with the HasEncryptedPayloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEncryptedPayloads

`func (o *MinimalFeatureFlag) SetHasEncryptedPayloads(v bool)`

SetHasEncryptedPayloads sets HasEncryptedPayloads field to given value.

### HasHasEncryptedPayloads

`func (o *MinimalFeatureFlag) HasHasEncryptedPayloads() bool`

HasHasEncryptedPayloads returns a boolean if a field has been set.

### SetHasEncryptedPayloadsNil

`func (o *MinimalFeatureFlag) SetHasEncryptedPayloadsNil(b bool)`

 SetHasEncryptedPayloadsNil sets the value for HasEncryptedPayloads to be an explicit nil

### UnsetHasEncryptedPayloads
`func (o *MinimalFeatureFlag) UnsetHasEncryptedPayloads()`

UnsetHasEncryptedPayloads ensures that no value is present for HasEncryptedPayloads, not even an explicit nil
### GetVersion

`func (o *MinimalFeatureFlag) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MinimalFeatureFlag) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MinimalFeatureFlag) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MinimalFeatureFlag) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MinimalFeatureFlag) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MinimalFeatureFlag) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetEvaluationRuntime

`func (o *MinimalFeatureFlag) GetEvaluationRuntime() string`

GetEvaluationRuntime returns the EvaluationRuntime field if non-nil, zero value otherwise.

### GetEvaluationRuntimeOk

`func (o *MinimalFeatureFlag) GetEvaluationRuntimeOk() (*string, bool)`

GetEvaluationRuntimeOk returns a tuple with the EvaluationRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationRuntime

`func (o *MinimalFeatureFlag) SetEvaluationRuntime(v string)`

SetEvaluationRuntime sets EvaluationRuntime field to given value.

### HasEvaluationRuntime

`func (o *MinimalFeatureFlag) HasEvaluationRuntime() bool`

HasEvaluationRuntime returns a boolean if a field has been set.

### SetEvaluationRuntimeNil

`func (o *MinimalFeatureFlag) SetEvaluationRuntimeNil(b bool)`

 SetEvaluationRuntimeNil sets the value for EvaluationRuntime to be an explicit nil

### UnsetEvaluationRuntime
`func (o *MinimalFeatureFlag) UnsetEvaluationRuntime()`

UnsetEvaluationRuntime ensures that no value is present for EvaluationRuntime, not even an explicit nil
### GetEvaluationTags

`func (o *MinimalFeatureFlag) GetEvaluationTags() []string`

GetEvaluationTags returns the EvaluationTags field if non-nil, zero value otherwise.

### GetEvaluationTagsOk

`func (o *MinimalFeatureFlag) GetEvaluationTagsOk() (*[]string, bool)`

GetEvaluationTagsOk returns a tuple with the EvaluationTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTags

`func (o *MinimalFeatureFlag) SetEvaluationTags(v []string)`

SetEvaluationTags sets EvaluationTags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


