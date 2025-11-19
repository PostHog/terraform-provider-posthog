# Experiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**FeatureFlagKey** | **string** |  | 
**FeatureFlag** | [**MinimalFeatureFlag**](MinimalFeatureFlag.md) |  | [readonly] 
**Holdout** | [**ExperimentHoldout**](ExperimentHoldout.md) |  | [readonly] 
**HoldoutId** | Pointer to **NullableInt32** |  | [optional] 
**ExposureCohort** | **NullableInt32** |  | [readonly] 
**Parameters** | Pointer to **interface{}** |  | [optional] 
**SecondaryMetrics** | Pointer to **interface{}** |  | [optional] 
**SavedMetrics** | [**[]ExperimentToSavedMetric**](ExperimentToSavedMetric.md) |  | [readonly] 
**SavedMetricsIds** | Pointer to **[]interface{}** |  | [optional] 
**Filters** | Pointer to **interface{}** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **NullableBool** |  | [optional] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
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
**UserAccessLevel** | **NullableString** | The effective access level the user has for this object | [readonly] 

## Methods

### NewExperiment

`func NewExperiment(id int32, name string, featureFlagKey string, featureFlag MinimalFeatureFlag, holdout ExperimentHoldout, exposureCohort NullableInt32, savedMetrics []ExperimentToSavedMetric, createdBy UserBasic, createdAt time.Time, updatedAt time.Time, userAccessLevel NullableString, ) *Experiment`

NewExperiment instantiates a new Experiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentWithDefaults

`func NewExperimentWithDefaults() *Experiment`

NewExperimentWithDefaults instantiates a new Experiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Experiment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Experiment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Experiment) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Experiment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Experiment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Experiment) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Experiment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Experiment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Experiment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Experiment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Experiment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Experiment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartDate

`func (o *Experiment) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Experiment) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Experiment) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Experiment) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Experiment) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Experiment) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *Experiment) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Experiment) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Experiment) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Experiment) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Experiment) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Experiment) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetFeatureFlagKey

`func (o *Experiment) GetFeatureFlagKey() string`

GetFeatureFlagKey returns the FeatureFlagKey field if non-nil, zero value otherwise.

### GetFeatureFlagKeyOk

`func (o *Experiment) GetFeatureFlagKeyOk() (*string, bool)`

GetFeatureFlagKeyOk returns a tuple with the FeatureFlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagKey

`func (o *Experiment) SetFeatureFlagKey(v string)`

SetFeatureFlagKey sets FeatureFlagKey field to given value.


### GetFeatureFlag

`func (o *Experiment) GetFeatureFlag() MinimalFeatureFlag`

GetFeatureFlag returns the FeatureFlag field if non-nil, zero value otherwise.

### GetFeatureFlagOk

`func (o *Experiment) GetFeatureFlagOk() (*MinimalFeatureFlag, bool)`

GetFeatureFlagOk returns a tuple with the FeatureFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlag

`func (o *Experiment) SetFeatureFlag(v MinimalFeatureFlag)`

SetFeatureFlag sets FeatureFlag field to given value.


### GetHoldout

`func (o *Experiment) GetHoldout() ExperimentHoldout`

GetHoldout returns the Holdout field if non-nil, zero value otherwise.

### GetHoldoutOk

`func (o *Experiment) GetHoldoutOk() (*ExperimentHoldout, bool)`

GetHoldoutOk returns a tuple with the Holdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldout

`func (o *Experiment) SetHoldout(v ExperimentHoldout)`

SetHoldout sets Holdout field to given value.


### GetHoldoutId

`func (o *Experiment) GetHoldoutId() int32`

GetHoldoutId returns the HoldoutId field if non-nil, zero value otherwise.

### GetHoldoutIdOk

`func (o *Experiment) GetHoldoutIdOk() (*int32, bool)`

GetHoldoutIdOk returns a tuple with the HoldoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldoutId

`func (o *Experiment) SetHoldoutId(v int32)`

SetHoldoutId sets HoldoutId field to given value.

### HasHoldoutId

`func (o *Experiment) HasHoldoutId() bool`

HasHoldoutId returns a boolean if a field has been set.

### SetHoldoutIdNil

`func (o *Experiment) SetHoldoutIdNil(b bool)`

 SetHoldoutIdNil sets the value for HoldoutId to be an explicit nil

### UnsetHoldoutId
`func (o *Experiment) UnsetHoldoutId()`

UnsetHoldoutId ensures that no value is present for HoldoutId, not even an explicit nil
### GetExposureCohort

`func (o *Experiment) GetExposureCohort() int32`

GetExposureCohort returns the ExposureCohort field if non-nil, zero value otherwise.

### GetExposureCohortOk

`func (o *Experiment) GetExposureCohortOk() (*int32, bool)`

GetExposureCohortOk returns a tuple with the ExposureCohort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureCohort

`func (o *Experiment) SetExposureCohort(v int32)`

SetExposureCohort sets ExposureCohort field to given value.


### SetExposureCohortNil

`func (o *Experiment) SetExposureCohortNil(b bool)`

 SetExposureCohortNil sets the value for ExposureCohort to be an explicit nil

### UnsetExposureCohort
`func (o *Experiment) UnsetExposureCohort()`

UnsetExposureCohort ensures that no value is present for ExposureCohort, not even an explicit nil
### GetParameters

`func (o *Experiment) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Experiment) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Experiment) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Experiment) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *Experiment) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *Experiment) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetSecondaryMetrics

`func (o *Experiment) GetSecondaryMetrics() interface{}`

GetSecondaryMetrics returns the SecondaryMetrics field if non-nil, zero value otherwise.

### GetSecondaryMetricsOk

`func (o *Experiment) GetSecondaryMetricsOk() (*interface{}, bool)`

GetSecondaryMetricsOk returns a tuple with the SecondaryMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetrics

`func (o *Experiment) SetSecondaryMetrics(v interface{})`

SetSecondaryMetrics sets SecondaryMetrics field to given value.

### HasSecondaryMetrics

`func (o *Experiment) HasSecondaryMetrics() bool`

HasSecondaryMetrics returns a boolean if a field has been set.

### SetSecondaryMetricsNil

`func (o *Experiment) SetSecondaryMetricsNil(b bool)`

 SetSecondaryMetricsNil sets the value for SecondaryMetrics to be an explicit nil

### UnsetSecondaryMetrics
`func (o *Experiment) UnsetSecondaryMetrics()`

UnsetSecondaryMetrics ensures that no value is present for SecondaryMetrics, not even an explicit nil
### GetSavedMetrics

`func (o *Experiment) GetSavedMetrics() []ExperimentToSavedMetric`

GetSavedMetrics returns the SavedMetrics field if non-nil, zero value otherwise.

### GetSavedMetricsOk

`func (o *Experiment) GetSavedMetricsOk() (*[]ExperimentToSavedMetric, bool)`

GetSavedMetricsOk returns a tuple with the SavedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedMetrics

`func (o *Experiment) SetSavedMetrics(v []ExperimentToSavedMetric)`

SetSavedMetrics sets SavedMetrics field to given value.


### GetSavedMetricsIds

`func (o *Experiment) GetSavedMetricsIds() []interface{}`

GetSavedMetricsIds returns the SavedMetricsIds field if non-nil, zero value otherwise.

### GetSavedMetricsIdsOk

`func (o *Experiment) GetSavedMetricsIdsOk() (*[]interface{}, bool)`

GetSavedMetricsIdsOk returns a tuple with the SavedMetricsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedMetricsIds

`func (o *Experiment) SetSavedMetricsIds(v []interface{})`

SetSavedMetricsIds sets SavedMetricsIds field to given value.

### HasSavedMetricsIds

`func (o *Experiment) HasSavedMetricsIds() bool`

HasSavedMetricsIds returns a boolean if a field has been set.

### SetSavedMetricsIdsNil

`func (o *Experiment) SetSavedMetricsIdsNil(b bool)`

 SetSavedMetricsIdsNil sets the value for SavedMetricsIds to be an explicit nil

### UnsetSavedMetricsIds
`func (o *Experiment) UnsetSavedMetricsIds()`

UnsetSavedMetricsIds ensures that no value is present for SavedMetricsIds, not even an explicit nil
### GetFilters

`func (o *Experiment) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Experiment) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Experiment) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Experiment) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *Experiment) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *Experiment) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetArchived

`func (o *Experiment) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Experiment) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Experiment) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Experiment) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDeleted

`func (o *Experiment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Experiment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Experiment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Experiment) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### SetDeletedNil

`func (o *Experiment) SetDeletedNil(b bool)`

 SetDeletedNil sets the value for Deleted to be an explicit nil

### UnsetDeleted
`func (o *Experiment) UnsetDeleted()`

UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil
### GetCreatedBy

`func (o *Experiment) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Experiment) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Experiment) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Experiment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Experiment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Experiment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Experiment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Experiment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Experiment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetType

`func (o *Experiment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Experiment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Experiment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Experiment) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Experiment) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Experiment) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExposureCriteria

`func (o *Experiment) GetExposureCriteria() interface{}`

GetExposureCriteria returns the ExposureCriteria field if non-nil, zero value otherwise.

### GetExposureCriteriaOk

`func (o *Experiment) GetExposureCriteriaOk() (*interface{}, bool)`

GetExposureCriteriaOk returns a tuple with the ExposureCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureCriteria

`func (o *Experiment) SetExposureCriteria(v interface{})`

SetExposureCriteria sets ExposureCriteria field to given value.

### HasExposureCriteria

`func (o *Experiment) HasExposureCriteria() bool`

HasExposureCriteria returns a boolean if a field has been set.

### SetExposureCriteriaNil

`func (o *Experiment) SetExposureCriteriaNil(b bool)`

 SetExposureCriteriaNil sets the value for ExposureCriteria to be an explicit nil

### UnsetExposureCriteria
`func (o *Experiment) UnsetExposureCriteria()`

UnsetExposureCriteria ensures that no value is present for ExposureCriteria, not even an explicit nil
### GetMetrics

`func (o *Experiment) GetMetrics() interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Experiment) GetMetricsOk() (*interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Experiment) SetMetrics(v interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Experiment) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetricsNil

`func (o *Experiment) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *Experiment) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil
### GetMetricsSecondary

`func (o *Experiment) GetMetricsSecondary() interface{}`

GetMetricsSecondary returns the MetricsSecondary field if non-nil, zero value otherwise.

### GetMetricsSecondaryOk

`func (o *Experiment) GetMetricsSecondaryOk() (*interface{}, bool)`

GetMetricsSecondaryOk returns a tuple with the MetricsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsSecondary

`func (o *Experiment) SetMetricsSecondary(v interface{})`

SetMetricsSecondary sets MetricsSecondary field to given value.

### HasMetricsSecondary

`func (o *Experiment) HasMetricsSecondary() bool`

HasMetricsSecondary returns a boolean if a field has been set.

### SetMetricsSecondaryNil

`func (o *Experiment) SetMetricsSecondaryNil(b bool)`

 SetMetricsSecondaryNil sets the value for MetricsSecondary to be an explicit nil

### UnsetMetricsSecondary
`func (o *Experiment) UnsetMetricsSecondary()`

UnsetMetricsSecondary ensures that no value is present for MetricsSecondary, not even an explicit nil
### GetStatsConfig

`func (o *Experiment) GetStatsConfig() interface{}`

GetStatsConfig returns the StatsConfig field if non-nil, zero value otherwise.

### GetStatsConfigOk

`func (o *Experiment) GetStatsConfigOk() (*interface{}, bool)`

GetStatsConfigOk returns a tuple with the StatsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsConfig

`func (o *Experiment) SetStatsConfig(v interface{})`

SetStatsConfig sets StatsConfig field to given value.

### HasStatsConfig

`func (o *Experiment) HasStatsConfig() bool`

HasStatsConfig returns a boolean if a field has been set.

### SetStatsConfigNil

`func (o *Experiment) SetStatsConfigNil(b bool)`

 SetStatsConfigNil sets the value for StatsConfig to be an explicit nil

### UnsetStatsConfig
`func (o *Experiment) UnsetStatsConfig()`

UnsetStatsConfig ensures that no value is present for StatsConfig, not even an explicit nil
### GetCreateInFolder

`func (o *Experiment) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *Experiment) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *Experiment) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *Experiment) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetConclusion

`func (o *Experiment) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *Experiment) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *Experiment) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.

### HasConclusion

`func (o *Experiment) HasConclusion() bool`

HasConclusion returns a boolean if a field has been set.

### SetConclusionNil

`func (o *Experiment) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *Experiment) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetConclusionComment

`func (o *Experiment) GetConclusionComment() string`

GetConclusionComment returns the ConclusionComment field if non-nil, zero value otherwise.

### GetConclusionCommentOk

`func (o *Experiment) GetConclusionCommentOk() (*string, bool)`

GetConclusionCommentOk returns a tuple with the ConclusionComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusionComment

`func (o *Experiment) SetConclusionComment(v string)`

SetConclusionComment sets ConclusionComment field to given value.

### HasConclusionComment

`func (o *Experiment) HasConclusionComment() bool`

HasConclusionComment returns a boolean if a field has been set.

### SetConclusionCommentNil

`func (o *Experiment) SetConclusionCommentNil(b bool)`

 SetConclusionCommentNil sets the value for ConclusionComment to be an explicit nil

### UnsetConclusionComment
`func (o *Experiment) UnsetConclusionComment()`

UnsetConclusionComment ensures that no value is present for ConclusionComment, not even an explicit nil
### GetPrimaryMetricsOrderedUuids

`func (o *Experiment) GetPrimaryMetricsOrderedUuids() interface{}`

GetPrimaryMetricsOrderedUuids returns the PrimaryMetricsOrderedUuids field if non-nil, zero value otherwise.

### GetPrimaryMetricsOrderedUuidsOk

`func (o *Experiment) GetPrimaryMetricsOrderedUuidsOk() (*interface{}, bool)`

GetPrimaryMetricsOrderedUuidsOk returns a tuple with the PrimaryMetricsOrderedUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMetricsOrderedUuids

`func (o *Experiment) SetPrimaryMetricsOrderedUuids(v interface{})`

SetPrimaryMetricsOrderedUuids sets PrimaryMetricsOrderedUuids field to given value.

### HasPrimaryMetricsOrderedUuids

`func (o *Experiment) HasPrimaryMetricsOrderedUuids() bool`

HasPrimaryMetricsOrderedUuids returns a boolean if a field has been set.

### SetPrimaryMetricsOrderedUuidsNil

`func (o *Experiment) SetPrimaryMetricsOrderedUuidsNil(b bool)`

 SetPrimaryMetricsOrderedUuidsNil sets the value for PrimaryMetricsOrderedUuids to be an explicit nil

### UnsetPrimaryMetricsOrderedUuids
`func (o *Experiment) UnsetPrimaryMetricsOrderedUuids()`

UnsetPrimaryMetricsOrderedUuids ensures that no value is present for PrimaryMetricsOrderedUuids, not even an explicit nil
### GetSecondaryMetricsOrderedUuids

`func (o *Experiment) GetSecondaryMetricsOrderedUuids() interface{}`

GetSecondaryMetricsOrderedUuids returns the SecondaryMetricsOrderedUuids field if non-nil, zero value otherwise.

### GetSecondaryMetricsOrderedUuidsOk

`func (o *Experiment) GetSecondaryMetricsOrderedUuidsOk() (*interface{}, bool)`

GetSecondaryMetricsOrderedUuidsOk returns a tuple with the SecondaryMetricsOrderedUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryMetricsOrderedUuids

`func (o *Experiment) SetSecondaryMetricsOrderedUuids(v interface{})`

SetSecondaryMetricsOrderedUuids sets SecondaryMetricsOrderedUuids field to given value.

### HasSecondaryMetricsOrderedUuids

`func (o *Experiment) HasSecondaryMetricsOrderedUuids() bool`

HasSecondaryMetricsOrderedUuids returns a boolean if a field has been set.

### SetSecondaryMetricsOrderedUuidsNil

`func (o *Experiment) SetSecondaryMetricsOrderedUuidsNil(b bool)`

 SetSecondaryMetricsOrderedUuidsNil sets the value for SecondaryMetricsOrderedUuids to be an explicit nil

### UnsetSecondaryMetricsOrderedUuids
`func (o *Experiment) UnsetSecondaryMetricsOrderedUuids()`

UnsetSecondaryMetricsOrderedUuids ensures that no value is present for SecondaryMetricsOrderedUuids, not even an explicit nil
### GetUserAccessLevel

`func (o *Experiment) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *Experiment) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *Experiment) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.


### SetUserAccessLevelNil

`func (o *Experiment) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *Experiment) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


