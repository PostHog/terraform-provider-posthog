# ExperimentExposureQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **NullableString** |  | [optional] 
**ExperimentId** | Pointer to **NullableInt32** |  | [optional] 
**ExperimentName** | **string** |  | 
**ExposureCriteria** | Pointer to [**ExperimentExposureCriteria**](ExperimentExposureCriteria.md) |  | [optional] 
**FeatureFlag** | **map[string]interface{}** |  | 
**Holdout** | Pointer to [**ExperimentHoldoutType**](ExperimentHoldoutType.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] [default to "ExperimentExposureQuery"]
**Modifiers** | Pointer to [**HogQLQueryModifiers**](HogQLQueryModifiers.md) |  | [optional] 
**Response** | Pointer to [**ExperimentExposureQueryResponse**](ExperimentExposureQueryResponse.md) |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**QueryLogTags**](QueryLogTags.md) |  | [optional] 
**Version** | Pointer to **NullableFloat32** | version of the node, used for schema migrations | [optional] 

## Methods

### NewExperimentExposureQuery

`func NewExperimentExposureQuery(experimentName string, featureFlag map[string]interface{}, ) *ExperimentExposureQuery`

NewExperimentExposureQuery instantiates a new ExperimentExposureQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentExposureQueryWithDefaults

`func NewExperimentExposureQueryWithDefaults() *ExperimentExposureQuery`

NewExperimentExposureQueryWithDefaults instantiates a new ExperimentExposureQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *ExperimentExposureQuery) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExperimentExposureQuery) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExperimentExposureQuery) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ExperimentExposureQuery) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ExperimentExposureQuery) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ExperimentExposureQuery) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetExperimentId

`func (o *ExperimentExposureQuery) GetExperimentId() int32`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *ExperimentExposureQuery) GetExperimentIdOk() (*int32, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *ExperimentExposureQuery) SetExperimentId(v int32)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *ExperimentExposureQuery) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### SetExperimentIdNil

`func (o *ExperimentExposureQuery) SetExperimentIdNil(b bool)`

 SetExperimentIdNil sets the value for ExperimentId to be an explicit nil

### UnsetExperimentId
`func (o *ExperimentExposureQuery) UnsetExperimentId()`

UnsetExperimentId ensures that no value is present for ExperimentId, not even an explicit nil
### GetExperimentName

`func (o *ExperimentExposureQuery) GetExperimentName() string`

GetExperimentName returns the ExperimentName field if non-nil, zero value otherwise.

### GetExperimentNameOk

`func (o *ExperimentExposureQuery) GetExperimentNameOk() (*string, bool)`

GetExperimentNameOk returns a tuple with the ExperimentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentName

`func (o *ExperimentExposureQuery) SetExperimentName(v string)`

SetExperimentName sets ExperimentName field to given value.


### GetExposureCriteria

`func (o *ExperimentExposureQuery) GetExposureCriteria() ExperimentExposureCriteria`

GetExposureCriteria returns the ExposureCriteria field if non-nil, zero value otherwise.

### GetExposureCriteriaOk

`func (o *ExperimentExposureQuery) GetExposureCriteriaOk() (*ExperimentExposureCriteria, bool)`

GetExposureCriteriaOk returns a tuple with the ExposureCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureCriteria

`func (o *ExperimentExposureQuery) SetExposureCriteria(v ExperimentExposureCriteria)`

SetExposureCriteria sets ExposureCriteria field to given value.

### HasExposureCriteria

`func (o *ExperimentExposureQuery) HasExposureCriteria() bool`

HasExposureCriteria returns a boolean if a field has been set.

### GetFeatureFlag

`func (o *ExperimentExposureQuery) GetFeatureFlag() map[string]interface{}`

GetFeatureFlag returns the FeatureFlag field if non-nil, zero value otherwise.

### GetFeatureFlagOk

`func (o *ExperimentExposureQuery) GetFeatureFlagOk() (*map[string]interface{}, bool)`

GetFeatureFlagOk returns a tuple with the FeatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlag

`func (o *ExperimentExposureQuery) SetFeatureFlag(v map[string]interface{})`

SetFeatureFlag sets FeatureFlag field to given value.


### GetHoldout

`func (o *ExperimentExposureQuery) GetHoldout() ExperimentHoldoutType`

GetHoldout returns the Holdout field if non-nil, zero value otherwise.

### GetHoldoutOk

`func (o *ExperimentExposureQuery) GetHoldoutOk() (*ExperimentHoldoutType, bool)`

GetHoldoutOk returns a tuple with the Holdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldout

`func (o *ExperimentExposureQuery) SetHoldout(v ExperimentHoldoutType)`

SetHoldout sets Holdout field to given value.

### HasHoldout

`func (o *ExperimentExposureQuery) HasHoldout() bool`

HasHoldout returns a boolean if a field has been set.

### GetKind

`func (o *ExperimentExposureQuery) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExperimentExposureQuery) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExperimentExposureQuery) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExperimentExposureQuery) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetModifiers

`func (o *ExperimentExposureQuery) GetModifiers() HogQLQueryModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *ExperimentExposureQuery) GetModifiersOk() (*HogQLQueryModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *ExperimentExposureQuery) SetModifiers(v HogQLQueryModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *ExperimentExposureQuery) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetResponse

`func (o *ExperimentExposureQuery) GetResponse() ExperimentExposureQueryResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExperimentExposureQuery) GetResponseOk() (*ExperimentExposureQueryResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExperimentExposureQuery) SetResponse(v ExperimentExposureQueryResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExperimentExposureQuery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStartDate

`func (o *ExperimentExposureQuery) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExperimentExposureQuery) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExperimentExposureQuery) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ExperimentExposureQuery) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ExperimentExposureQuery) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ExperimentExposureQuery) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetTags

`func (o *ExperimentExposureQuery) GetTags() QueryLogTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExperimentExposureQuery) GetTagsOk() (*QueryLogTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExperimentExposureQuery) SetTags(v QueryLogTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExperimentExposureQuery) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *ExperimentExposureQuery) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExperimentExposureQuery) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExperimentExposureQuery) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExperimentExposureQuery) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ExperimentExposureQuery) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ExperimentExposureQuery) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


