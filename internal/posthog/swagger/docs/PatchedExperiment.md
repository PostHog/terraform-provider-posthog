# PatchedExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**FeatureFlagKey** | Pointer to **string** |  | [optional] 
**FeatureFlag** | Pointer to [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [optional] [readonly] 
**Holdout** | Pointer to [**ExperimentHoldout**](ExperimentHoldout.md) |  | [optional] [readonly] 
**HoldoutId** | Pointer to **NullableInt32** |  | [optional] 
**ExposureCohort** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Parameters** | Pointer to **interface{}** |  | [optional] 
**SecondaryMetrics** | Pointer to **interface{}** |  | [optional] 
**SavedMetrics** | Pointer to [**[]ExperimentToSavedMetric**](ExperimentToSavedMetric.md) |  | [optional] [readonly] 
**SavedMetricsIds** | Pointer to **[]interface{}** |  | [optional] 
**Filters** | Pointer to **interface{}** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ExposureCriteria** | Pointer to **interface{}** |  | [optional] 
**Metrics** | Pointer to **interface{}** |  | [optional] 
**MetricsSecondary** | Pointer to **interface{}** |  | [optional] 
**StatsConfig** | Pointer to **interface{}** |  | [optional] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**Conclusion** | Pointer to **NullableString** |  | [optional] 
**ConclusionComment** | Pointer to **NullableString** |  | [optional] 
**PrimaryMetricsOrderedUuids** | Pointer to **interface{}** |  | [optional] 
**SecondaryMetricsOrderedUuids** | Pointer to **interface{}** |  | [optional] 
**UserAccessLevel** | Pointer to **NullableString** | The effective access level the user has for this object | [optional] [readonly] 

## Methods

### NewPatchedExperiment

`func NewPatchedExperiment() *PatchedExperiment`

NewPatchedExperiment instantiates a new PatchedExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedExperimentWithDefaults

`func NewPatchedExperimentWithDefaults() *PatchedExperiment`

NewPatchedExperimentWithDefaults instantiates a new PatchedExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedExperiment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedExperiment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedExperiment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedExperiment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedExperiment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedExperiment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedExperiment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedExperiment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedExperiment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedExperiment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedExperiment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedExperiment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedExperiment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedExperiment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartDate

`func (o *PatchedExperiment) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PatchedExperiment) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PatchedExperiment) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PatchedExperiment) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *PatchedExperiment) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *PatchedExperiment) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *PatchedExperiment) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PatchedExperiment) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PatchedExperiment) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PatchedExperiment) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *PatchedExperiment) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *PatchedExperiment) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetFeatureFlagKey

`func (o *PatchedExperiment) GetFeatureFlagKey() string`

GetFeatureFlagKey returns the FeatureFlagKey field if non-nil, zero value otherwise.

### GetFeatureFlagKeyOk

`func (o *PatchedExperiment) GetFeatureFlagKeyOk() (*string, bool)`

GetFeatureFlagKeyOk returns a tuple with the FeatureFlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagKey

`func (o *PatchedExperiment) SetFeatureFlagKey(v string)`

SetFeatureFlagKey sets FeatureFlagKey field to given value.

### HasFeatureFlagKey

`func (o *PatchedExperiment) HasFeatureFlagKey() bool`

HasFeatureFlagKey returns a boolean if a field has been set.

### GetFeatureFlag

`func (o *PatchedExperiment) GetFeatureFlag() MinimalFeatureFlag`

GetFeatureFlag returns the FeatureFlag field if non-nil, zero value otherwise.

### GetFeatureFlagOk

`func (o *PatchedExperiment) GetFeatureFlagOk() (*MinimalFeatureFlag, bool)`

GetFeatureFlagOk returns a tuple with the FeatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlag

`func (o *PatchedExperiment) SetFeatureFlag(v MinimalFeatureFlag)`

SetFeatureFlag sets FeatureFlag field to given value.

### HasFeatureFlag

`func (o *PatchedExperiment) HasFeatureFlag() bool`

HasFeatureFlag returns a boolean if a field has been set.

### GetHoldout

`func (o *PatchedExperiment) GetHoldout() ExperimentHoldout`

GetHoldout returns the Holdout field if non-nil, zero value otherwise.

### GetHoldoutOk

`func (o *PatchedExperiment) GetHoldoutOk() (*ExperimentHoldout, bool)`

GetHoldoutOk returns a tuple with the Holdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldout

`func (o *PatchedExperiment) SetHoldout(v ExperimentHoldout)`

SetHoldout sets Holdout field to given value.

### HasHoldout

`func (o *PatchedExperiment) HasHoldout() bool`

HasHoldout returns a boolean if a field has been set.

### GetHoldoutId

`func (o *PatchedExperiment) GetHoldoutId() int32`

GetHoldoutId returns the HoldoutId field if non-nil, zero value otherwise.

### GetHoldoutIdOk

`func (o *PatchedExperiment) GetHoldoutIdOk() (*int32, bool)`

GetHoldoutIdOk returns a tuple with the HoldoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutId

`func (o *PatchedExperiment) SetHoldoutId(v int32)`

SetHoldoutId sets HoldoutId field to given value.

### HasHoldoutId

`func (o *PatchedExperiment) HasHoldoutId() bool`

HasHoldoutId returns a boolean if a field has been set.

### SetHoldoutIdNil

`func (o *PatchedExperiment) SetHoldoutIdNil(b bool)`

 SetHoldoutIdNil sets the value for HoldoutId to be an explicit nil

### UnsetHoldoutId
`func (o *PatchedExperiment) UnsetHoldoutId()`

UnsetHoldoutId ensures that no value is present for HoldoutId, not even an explicit nil
### GetExposureCohort

`func (o *PatchedExperiment) GetExposureCohort() int32`

GetExposureCohort returns the ExposureCohort field if non-nil, zero value otherwise.

### GetExposureCohortOk

`func (o *PatchedExperiment) GetExposureCohortOk() (*int32, bool)`

GetExposureCohortOk returns a tuple with the ExposureCohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureCohort

`func (o *PatchedExperiment) SetExposureCohort(v int32)`

SetExposureCohort sets ExposureCohort field to given value.

### HasExposureCohort

`func (o *PatchedExperiment) HasExposureCohort() bool`

HasExposureCohort returns a boolean if a field has been set.

### SetExposureCohortNil

`func (o *PatchedExperiment) SetExposureCohortNil(b bool)`

 SetExposureCohortNil sets the value for ExposureCohort to be an explicit nil

### UnsetExposureCohort
`func (o *PatchedExperiment) UnsetExposureCohort()`

UnsetExposureCohort ensures that no value is present for ExposureCohort, not even an explicit nil
### GetParameters

`func (o *PatchedExperiment) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PatchedExperiment) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PatchedExperiment) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PatchedExperiment) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PatchedExperiment) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PatchedExperiment) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetSecondaryMetrics

`func (o *PatchedExperiment) GetSecondaryMetrics() interface{}`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *PatchedExperiment) GetSecondaryMetricsOk() (*interface{}, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *PatchedExperiment) SetSecondaryMetrics(v interface{})`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *PatchedExperiment) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.

### SetSecondaryMetricsNil

`func (o *PatchedExperiment) SetSecondaryMetricsNil(b bool)`

 SetSecondaryMetricsNil sets the value for SecondaryMetrics to be an explicit nil

### UnsetSecondaryMetrics
`func (o *PatchedExperiment) UnsetSecondaryMetrics()`

UnsetSecondaryMetrics ensures that no value is present for SecondaryMetrics, not even an explicit nil
### GetSavedMetrics

`func (o *PatchedExperiment) GetSavedMetrics() []ExperimentToSavedMetric`

GetSavedMetrics returns the SavedMetrics field if non-nil, zero value otherwise.

### GetSavedMetricsOk

`func (o *PatchedExperiment) GetSavedMetricsOk() (*[]ExperimentToSavedMetric, bool)`

GetSavedMetricsOk returns a tuple with the SavedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedMetrics

`func (o *PatchedExperiment) SetSavedMetrics(v []ExperimentToSavedMetric)`

SetSavedMetrics sets SavedMetrics field to given value.

### HasSavedMetrics

`func (o *PatchedExperiment) HasSavedMetrics() bool`

HasSavedMetrics returns a boolean if a field has been set.

### GetSavedMetricsIds

`func (o *PatchedExperiment) GetSavedMetricsIds() []interface{}`

GetSavedMetricsIds returns the SavedMetricsIds field if non-nil, zero value otherwise.

### GetSavedMetricsIdsOk

`func (o *PatchedExperiment) GetSavedMetricsIdsOk() (*[]interface{}, bool)`

GetSavedMetricsIdsOk returns a tuple with the SavedMetricsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedMetricsIds

`func (o *PatchedExperiment) SetSavedMetricsIds(v []interface{})`

SetSavedMetricsIds sets SavedMetricsIds field to given value.

### HasSavedMetricsIds

`func (o *PatchedExperiment) HasSavedMetricsIds() bool`

HasSavedMetricsIds returns a boolean if a field has been set.

### SetSavedMetricsIdsNil

`func (o *PatchedExperiment) SetSavedMetricsIdsNil(b bool)`

 SetSavedMetricsIdsNil sets the value for SavedMetricsIds to be an explicit nil

### UnsetSavedMetricsIds
`func (o *PatchedExperiment) UnsetSavedMetricsIds()`

UnsetSavedMetricsIds ensures that no value is present for SavedMetricsIds, not even an explicit nil
### GetFilters

`func (o *PatchedExperiment) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedExperiment) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedExperiment) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedExperiment) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PatchedExperiment) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PatchedExperiment) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetArchived

`func (o *PatchedExperiment) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *PatchedExperiment) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *PatchedExperiment) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *PatchedExperiment) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedExperiment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedExperiment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedExperiment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedExperiment) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *PatchedExperiment) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *PatchedExperiment) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetCreatedBy

`func (o *PatchedExperiment) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedExperiment) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedExperiment) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedExperiment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedExperiment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedExperiment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedExperiment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedExperiment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedExperiment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedExperiment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedExperiment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedExperiment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *PatchedExperiment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedExperiment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedExperiment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedExperiment) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchedExperiment) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchedExperiment) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExposureCriteria

`func (o *PatchedExperiment) GetExposureCriteria() interface{}`

GetExposureCriteria returns the ExposureCriteria field if non-nil, zero value otherwise.

### GetExposureCriteriaOk

`func (o *PatchedExperiment) GetExposureCriteriaOk() (*interface{}, bool)`

GetExposureCriteriaOk returns a tuple with the ExposureCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureCriteria

`func (o *PatchedExperiment) SetExposureCriteria(v interface{})`

SetExposureCriteria sets ExposureCriteria field to given value.

### HasExposureCriteria

`func (o *PatchedExperiment) HasExposureCriteria() bool`

HasExposureCriteria returns a boolean if a field has been set.

### SetExposureCriteriaNil

`func (o *PatchedExperiment) SetExposureCriteriaNil(b bool)`

 SetExposureCriteriaNil sets the value for ExposureCriteria to be an explicit nil

### UnsetExposureCriteria
`func (o *PatchedExperiment) UnsetExposureCriteria()`

UnsetExposureCriteria ensures that no value is present for ExposureCriteria, not even an explicit nil
### GetMetrics

`func (o *PatchedExperiment) GetMetrics() interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *PatchedExperiment) GetMetricsOk() (*interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *PatchedExperiment) SetMetrics(v interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *PatchedExperiment) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetricsNil

`func (o *PatchedExperiment) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *PatchedExperiment) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil
### GetMetricsSecondary

`func (o *PatchedExperiment) GetMetricsSecondary() interface{}`

GetMetricsSecondary returns the MetricsSecondary field if non-nil, zero value otherwise.

### GetMetricsSecondaryOk

`func (o *PatchedExperiment) GetMetricsSecondaryOk() (*interface{}, bool)`

GetMetricsSecondaryOk returns a tuple with the MetricsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsSecondary

`func (o *PatchedExperiment) SetMetricsSecondary(v interface{})`

SetMetricsSecondary sets MetricsSecondary field to given value.

### HasMetricsSecondary

`func (o *PatchedExperiment) HasMetricsSecondary() bool`

HasMetricsSecondary returns a boolean if a field has been set.

### SetMetricsSecondaryNil

`func (o *PatchedExperiment) SetMetricsSecondaryNil(b bool)`

 SetMetricsSecondaryNil sets the value for MetricsSecondary to be an explicit nil

### UnsetMetricsSecondary
`func (o *PatchedExperiment) UnsetMetricsSecondary()`

UnsetMetricsSecondary ensures that no value is present for MetricsSecondary, not even an explicit nil
### GetStatsConfig

`func (o *PatchedExperiment) GetStatsConfig() interface{}`

GetStatsConfig returns the StatsConfig field if non-nil, zero value otherwise.

### GetStatsConfigOk

`func (o *PatchedExperiment) GetStatsConfigOk() (*interface{}, bool)`

GetStatsConfigOk returns a tuple with the StatsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsConfig

`func (o *PatchedExperiment) SetStatsConfig(v interface{})`

SetStatsConfig sets StatsConfig field to given value.

### HasStatsConfig

`func (o *PatchedExperiment) HasStatsConfig() bool`

HasStatsConfig returns a boolean if a field has been set.

### SetStatsConfigNil

`func (o *PatchedExperiment) SetStatsConfigNil(b bool)`

 SetStatsConfigNil sets the value for StatsConfig to be an explicit nil

### UnsetStatsConfig
`func (o *PatchedExperiment) UnsetStatsConfig()`

UnsetStatsConfig ensures that no value is present for StatsConfig, not even an explicit nil
### GetCreateInFolder

`func (o *PatchedExperiment) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedExperiment) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedExperiment) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedExperiment) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetConclusion

`func (o *PatchedExperiment) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *PatchedExperiment) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *PatchedExperiment) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *PatchedExperiment) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### SetConclusionNil

`func (o *PatchedExperiment) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *PatchedExperiment) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetConclusionComment

`func (o *PatchedExperiment) GetConclusionComment() string`

GetConclusionComment returns the ConclusionComment field if non-nil, zero value otherwise.

### GetConclusionCommentOk

`func (o *PatchedExperiment) GetConclusionCommentOk() (*string, bool)`

GetConclusionCommentOk returns a tuple with the ConclusionComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusionComment

`func (o *PatchedExperiment) SetConclusionComment(v string)`

SetConclusionComment sets ConclusionComment field to given value.

### HasConclusionComment

`func (o *PatchedExperiment) HasConclusionComment() bool`

HasConclusionComment returns a boolean if a field has been set.

### SetConclusionCommentNil

`func (o *PatchedExperiment) SetConclusionCommentNil(b bool)`

 SetConclusionCommentNil sets the value for ConclusionComment to be an explicit nil

### UnsetConclusionComment
`func (o *PatchedExperiment) UnsetConclusionComment()`

UnsetConclusionComment ensures that no value is present for ConclusionComment, not even an explicit nil
### GetPrimaryMetricsOrderedUuids

`func (o *PatchedExperiment) GetPrimaryMetricsOrderedUuids() interface{}`

GetPrimaryMetricsOrderedUuids returns the PrimaryMetricsOrderedUuids field if non-nil, zero value otherwise.

### GetPrimaryMetricsOrderedUuidsOk

`func (o *PatchedExperiment) GetPrimaryMetricsOrderedUuidsOk() (*interface{}, bool)`

GetPrimaryMetricsOrderedUuidsOk returns a tuple with the PrimaryMetricsOrderedUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetricsOrderedUuids

`func (o *PatchedExperiment) SetPrimaryMetricsOrderedUuids(v interface{})`

SetPrimaryMetricsOrderedUuids sets PrimaryMetricsOrderedUuids field to given value.

### HasPrimaryMetricsOrderedUuids

`func (o *PatchedExperiment) HasPrimaryMetricsOrderedUuids() bool`

HasPrimaryMetricsOrderedUuids returns a boolean if a field has been set.

### SetPrimaryMetricsOrderedUuidsNil

`func (o *PatchedExperiment) SetPrimaryMetricsOrderedUuidsNil(b bool)`

 SetPrimaryMetricsOrderedUuidsNil sets the value for PrimaryMetricsOrderedUuids to be an explicit nil

### UnsetPrimaryMetricsOrderedUuids
`func (o *PatchedExperiment) UnsetPrimaryMetricsOrderedUuids()`

UnsetPrimaryMetricsOrderedUuids ensures that no value is present for PrimaryMetricsOrderedUuids, not even an explicit nil
### GetSecondaryMetricsOrderedUuids

`func (o *PatchedExperiment) GetSecondaryMetricsOrderedUuids() interface{}`

GetSecondaryMetricsOrderedUuids returns the SecondaryMetricsOrderedUuids field if non-nil, zero value otherwise.

### GetSecondaryMetricsOrderedUuidsOk

`func (o *PatchedExperiment) GetSecondaryMetricsOrderedUuidsOk() (*interface{}, bool)`

GetSecondaryMetricsOrderedUuidsOk returns a tuple with the SecondaryMetricsOrderedUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetricsOrderedUuids

`func (o *PatchedExperiment) SetSecondaryMetricsOrderedUuids(v interface{})`

SetSecondaryMetricsOrderedUuids sets SecondaryMetricsOrderedUuids field to given value.

### HasSecondaryMetricsOrderedUuids

`func (o *PatchedExperiment) HasSecondaryMetricsOrderedUuids() bool`

HasSecondaryMetricsOrderedUuids returns a boolean if a field has been set.

### SetSecondaryMetricsOrderedUuidsNil

`func (o *PatchedExperiment) SetSecondaryMetricsOrderedUuidsNil(b bool)`

 SetSecondaryMetricsOrderedUuidsNil sets the value for SecondaryMetricsOrderedUuids to be an explicit nil

### UnsetSecondaryMetricsOrderedUuids
`func (o *PatchedExperiment) UnsetSecondaryMetricsOrderedUuids()`

UnsetSecondaryMetricsOrderedUuids ensures that no value is present for SecondaryMetricsOrderedUuids, not even an explicit nil
### GetUserAccessLevel

`func (o *PatchedExperiment) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *PatchedExperiment) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *PatchedExperiment) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.

### HasUserAccessLevel

`func (o *PatchedExperiment) HasUserAccessLevel() bool`

HasUserAccessLevel returns a boolean if a field has been set.

### SetUserAccessLevelNil

`func (o *PatchedExperiment) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *PatchedExperiment) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


