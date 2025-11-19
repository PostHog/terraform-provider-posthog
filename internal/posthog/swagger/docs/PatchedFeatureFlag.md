# PatchedFeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** | contains the description for the flag (field name &#x60;name&#x60; is kept for backwards-compatibility) | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **map[string]interface{}** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**Version** | Pointer to **int32** |  | [optional] [default to 0]
**LastModifiedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**IsSimpleFlag** | Pointer to **bool** |  | [optional] [readonly] 
**RolloutPercentage** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**EnsureExperienceContinuity** | Pointer to **NullableBool** |  | [optional] 
**ExperimentSet** | Pointer to **string** |  | [optional] [readonly] 
**Surveys** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Features** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**RollbackConditions** | Pointer to **interface{}** |  | [optional] 
**PerformedRollback** | Pointer to **NullableBool** |  | [optional] 
**CanEdit** | Pointer to **bool** |  | [optional] [readonly] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**EvaluationTags** | Pointer to **[]interface{}** |  | [optional] 
**UsageDashboard** | Pointer to **int32** |  | [optional] [readonly] 
**AnalyticsDashboards** | Pointer to **[]int32** |  | [optional] 
**HasEnrichedAnalytics** | Pointer to **NullableBool** |  | [optional] 
**UserAccessLevel** | Pointer to **NullableString** | The effective access level the user has for this object | [optional] [readonly] 
**CreationContext** | Pointer to [**CreationContextEnum**](CreationContextEnum.md) | Indicates the origin product of the feature flag. Choices: &#39;feature_flags&#39;, &#39;experiments&#39;, &#39;surveys&#39;, &#39;early_access_features&#39;, &#39;web_experiments&#39;.  * &#x60;feature_flags&#x60; - feature_flags * &#x60;experiments&#x60; - experiments * &#x60;surveys&#x60; - surveys * &#x60;early_access_features&#x60; - early_access_features * &#x60;web_experiments&#x60; - web_experiments | [optional] 
**IsRemoteConfiguration** | Pointer to **NullableBool** |  | [optional] 
**HasEncryptedPayloads** | Pointer to **NullableBool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**EvaluationRuntime** | Pointer to **NullableString** | Specifies where this feature flag should be evaluated  * &#x60;server&#x60; - Server * &#x60;client&#x60; - Client * &#x60;all&#x60; - All | [optional] 
**LastCalledAt** | Pointer to **NullableTime** | Last time this feature flag was called (from $feature_flag_called events) | [optional] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**ShouldCreateUsageDashboard** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPatchedFeatureFlag

`func NewPatchedFeatureFlag() *PatchedFeatureFlag`

NewPatchedFeatureFlag instantiates a new PatchedFeatureFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFeatureFlagWithDefaults

`func NewPatchedFeatureFlagWithDefaults() *PatchedFeatureFlag`

NewPatchedFeatureFlagWithDefaults instantiates a new PatchedFeatureFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedFeatureFlag) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedFeatureFlag) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedFeatureFlag) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedFeatureFlag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedFeatureFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedFeatureFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedFeatureFlag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedFeatureFlag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *PatchedFeatureFlag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedFeatureFlag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedFeatureFlag) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedFeatureFlag) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetFilters

`func (o *PatchedFeatureFlag) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedFeatureFlag) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedFeatureFlag) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedFeatureFlag) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedFeatureFlag) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedFeatureFlag) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedFeatureFlag) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedFeatureFlag) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetActive

`func (o *PatchedFeatureFlag) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedFeatureFlag) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedFeatureFlag) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedFeatureFlag) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedFeatureFlag) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedFeatureFlag) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedFeatureFlag) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedFeatureFlag) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedFeatureFlag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedFeatureFlag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedFeatureFlag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedFeatureFlag) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedFeatureFlag) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedFeatureFlag) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedFeatureFlag) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedFeatureFlag) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *PatchedFeatureFlag) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *PatchedFeatureFlag) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetVersion

`func (o *PatchedFeatureFlag) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedFeatureFlag) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedFeatureFlag) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedFeatureFlag) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *PatchedFeatureFlag) GetLastModifiedBy() UserBasic`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *PatchedFeatureFlag) GetLastModifiedByOk() (*UserBasic, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *PatchedFeatureFlag) SetLastModifiedBy(v UserBasic)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *PatchedFeatureFlag) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetIsSimpleFlag

`func (o *PatchedFeatureFlag) GetIsSimpleFlag() bool`

GetIsSimpleFlag returns the IsSimpleFlag field if non-nil, zero value otherwise.

### GetIsSimpleFlagOk

`func (o *PatchedFeatureFlag) GetIsSimpleFlagOk() (*bool, bool)`

GetIsSimpleFlagOk returns a tuple with the IsSimpleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSimpleFlag

`func (o *PatchedFeatureFlag) SetIsSimpleFlag(v bool)`

SetIsSimpleFlag sets IsSimpleFlag field to given value.

### HasIsSimpleFlag

`func (o *PatchedFeatureFlag) HasIsSimpleFlag() bool`

HasIsSimpleFlag returns a boolean if a field has been set.

### GetRolloutPercentage

`func (o *PatchedFeatureFlag) GetRolloutPercentage() int32`

GetRolloutPercentage returns the RolloutPercentage field if non-nil, zero value otherwise.

### GetRolloutPercentageOk

`func (o *PatchedFeatureFlag) GetRolloutPercentageOk() (*int32, bool)`

GetRolloutPercentageOk returns a tuple with the RolloutPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentage

`func (o *PatchedFeatureFlag) SetRolloutPercentage(v int32)`

SetRolloutPercentage sets RolloutPercentage field to given value.

### HasRolloutPercentage

`func (o *PatchedFeatureFlag) HasRolloutPercentage() bool`

HasRolloutPercentage returns a boolean if a field has been set.

### SetRolloutPercentageNil

`func (o *PatchedFeatureFlag) SetRolloutPercentageNil(b bool)`

 SetRolloutPercentageNil sets the value for RolloutPercentage to be an explicit nil

### UnsetRolloutPercentage
`func (o *PatchedFeatureFlag) UnsetRolloutPercentage()`

UnsetRolloutPercentage ensures that no value is present for RolloutPercentage, not even an explicit nil
### GetEnsureExperienceContinuity

`func (o *PatchedFeatureFlag) GetEnsureExperienceContinuity() bool`

GetEnsureExperienceContinuity returns the EnsureExperienceContinuity field if non-nil, zero value otherwise.

### GetEnsureExperienceContinuityOk

`func (o *PatchedFeatureFlag) GetEnsureExperienceContinuityOk() (*bool, bool)`

GetEnsureExperienceContinuityOk returns a tuple with the EnsureExperienceContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureExperienceContinuity

`func (o *PatchedFeatureFlag) SetEnsureExperienceContinuity(v bool)`

SetEnsureExperienceContinuity sets EnsureExperienceContinuity field to given value.

### HasEnsureExperienceContinuity

`func (o *PatchedFeatureFlag) HasEnsureExperienceContinuity() bool`

HasEnsureExperienceContinuity returns a boolean if a field has been set.

### SetEnsureExperienceContinuityNil

`func (o *PatchedFeatureFlag) SetEnsureExperienceContinuityNil(b bool)`

 SetEnsureExperienceContinuityNil sets the value for EnsureExperienceContinuity to be an explicit nil

### UnsetEnsureExperienceContinuity
`func (o *PatchedFeatureFlag) UnsetEnsureExperienceContinuity()`

UnsetEnsureExperienceContinuity ensures that no value is present for EnsureExperienceContinuity, not even an explicit nil
### GetExperimentSet

`func (o *PatchedFeatureFlag) GetExperimentSet() string`

GetExperimentSet returns the ExperimentSet field if non-nil, zero value otherwise.

### GetExperimentSetOk

`func (o *PatchedFeatureFlag) GetExperimentSetOk() (*string, bool)`

GetExperimentSetOk returns a tuple with the ExperimentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentSet

`func (o *PatchedFeatureFlag) SetExperimentSet(v string)`

SetExperimentSet sets ExperimentSet field to given value.

### HasExperimentSet

`func (o *PatchedFeatureFlag) HasExperimentSet() bool`

HasExperimentSet returns a boolean if a field has been set.

### GetSurveys

`func (o *PatchedFeatureFlag) GetSurveys() map[string]interface{}`

GetSurveys returns the Surveys field if non-nil, zero value otherwise.

### GetSurveysOk

`func (o *PatchedFeatureFlag) GetSurveysOk() (*map[string]interface{}, bool)`

GetSurveysOk returns a tuple with the Surveys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveys

`func (o *PatchedFeatureFlag) SetSurveys(v map[string]interface{})`

SetSurveys sets Surveys field to given value.

### HasSurveys

`func (o *PatchedFeatureFlag) HasSurveys() bool`

HasSurveys returns a boolean if a field has been set.

### GetFeatures

`func (o *PatchedFeatureFlag) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PatchedFeatureFlag) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PatchedFeatureFlag) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PatchedFeatureFlag) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetRollbackConditions

`func (o *PatchedFeatureFlag) GetRollbackConditions() interface{}`

GetRollbackConditions returns the RollbackConditions field if non-nil, zero value otherwise.

### GetRollbackConditionsOk

`func (o *PatchedFeatureFlag) GetRollbackConditionsOk() (*interface{}, bool)`

GetRollbackConditionsOk returns a tuple with the RollbackConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackConditions

`func (o *PatchedFeatureFlag) SetRollbackConditions(v interface{})`

SetRollbackConditions sets RollbackConditions field to given value.

### HasRollbackConditions

`func (o *PatchedFeatureFlag) HasRollbackConditions() bool`

HasRollbackConditions returns a boolean if a field has been set.

### SetRollbackConditionsNil

`func (o *PatchedFeatureFlag) SetRollbackConditionsNil(b bool)`

 SetRollbackConditionsNil sets the value for RollbackConditions to be an explicit nil

### UnsetRollbackConditions
`func (o *PatchedFeatureFlag) UnsetRollbackConditions()`

UnsetRollbackConditions ensures that no value is present for RollbackConditions, not even an explicit nil
### GetPerformedRollback

`func (o *PatchedFeatureFlag) GetPerformedRollback() bool`

GetPerformedRollback returns the PerformedRollback field if non-nil, zero value otherwise.

### GetPerformedRollbackOk

`func (o *PatchedFeatureFlag) GetPerformedRollbackOk() (*bool, bool)`

GetPerformedRollbackOk returns a tuple with the PerformedRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedRollback

`func (o *PatchedFeatureFlag) SetPerformedRollback(v bool)`

SetPerformedRollback sets PerformedRollback field to given value.

### HasPerformedRollback

`func (o *PatchedFeatureFlag) HasPerformedRollback() bool`

HasPerformedRollback returns a boolean if a field has been set.

### SetPerformedRollbackNil

`func (o *PatchedFeatureFlag) SetPerformedRollbackNil(b bool)`

 SetPerformedRollbackNil sets the value for PerformedRollback to be an explicit nil

### UnsetPerformedRollback
`func (o *PatchedFeatureFlag) UnsetPerformedRollback()`

UnsetPerformedRollback ensures that no value is present for PerformedRollback, not even an explicit nil
### GetCanEdit

`func (o *PatchedFeatureFlag) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *PatchedFeatureFlag) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *PatchedFeatureFlag) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *PatchedFeatureFlag) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetTags

`func (o *PatchedFeatureFlag) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedFeatureFlag) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedFeatureFlag) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedFeatureFlag) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvaluationTags

`func (o *PatchedFeatureFlag) GetEvaluationTags() []interface{}`

GetEvaluationTags returns the EvaluationTags field if non-nil, zero value otherwise.

### GetEvaluationTagsOk

`func (o *PatchedFeatureFlag) GetEvaluationTagsOk() (*[]interface{}, bool)`

GetEvaluationTagsOk returns a tuple with the EvaluationTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTags

`func (o *PatchedFeatureFlag) SetEvaluationTags(v []interface{})`

SetEvaluationTags sets EvaluationTags field to given value.

### HasEvaluationTags

`func (o *PatchedFeatureFlag) HasEvaluationTags() bool`

HasEvaluationTags returns a boolean if a field has been set.

### GetUsageDashboard

`func (o *PatchedFeatureFlag) GetUsageDashboard() int32`

GetUsageDashboard returns the UsageDashboard field if non-nil, zero value otherwise.

### GetUsageDashboardOk

`func (o *PatchedFeatureFlag) GetUsageDashboardOk() (*int32, bool)`

GetUsageDashboardOk returns a tuple with the UsageDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDashboard

`func (o *PatchedFeatureFlag) SetUsageDashboard(v int32)`

SetUsageDashboard sets UsageDashboard field to given value.

### HasUsageDashboard

`func (o *PatchedFeatureFlag) HasUsageDashboard() bool`

HasUsageDashboard returns a boolean if a field has been set.

### GetAnalyticsDashboards

`func (o *PatchedFeatureFlag) GetAnalyticsDashboards() []int32`

GetAnalyticsDashboards returns the AnalyticsDashboards field if non-nil, zero value otherwise.

### GetAnalyticsDashboardsOk

`func (o *PatchedFeatureFlag) GetAnalyticsDashboardsOk() (*[]int32, bool)`

GetAnalyticsDashboardsOk returns a tuple with the AnalyticsDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsDashboards

`func (o *PatchedFeatureFlag) SetAnalyticsDashboards(v []int32)`

SetAnalyticsDashboards sets AnalyticsDashboards field to given value.

### HasAnalyticsDashboards

`func (o *PatchedFeatureFlag) HasAnalyticsDashboards() bool`

HasAnalyticsDashboards returns a boolean if a field has been set.

### GetHasEnrichedAnalytics

`func (o *PatchedFeatureFlag) GetHasEnrichedAnalytics() bool`

GetHasEnrichedAnalytics returns the HasEnrichedAnalytics field if non-nil, zero value otherwise.

### GetHasEnrichedAnalyticsOk

`func (o *PatchedFeatureFlag) GetHasEnrichedAnalyticsOk() (*bool, bool)`

GetHasEnrichedAnalyticsOk returns a tuple with the HasEnrichedAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEnrichedAnalytics

`func (o *PatchedFeatureFlag) SetHasEnrichedAnalytics(v bool)`

SetHasEnrichedAnalytics sets HasEnrichedAnalytics field to given value.

### HasHasEnrichedAnalytics

`func (o *PatchedFeatureFlag) HasHasEnrichedAnalytics() bool`

HasHasEnrichedAnalytics returns a boolean if a field has been set.

### SetHasEnrichedAnalyticsNil

`func (o *PatchedFeatureFlag) SetHasEnrichedAnalyticsNil(b bool)`

 SetHasEnrichedAnalyticsNil sets the value for HasEnrichedAnalytics to be an explicit nil

### UnsetHasEnrichedAnalytics
`func (o *PatchedFeatureFlag) UnsetHasEnrichedAnalytics()`

UnsetHasEnrichedAnalytics ensures that no value is present for HasEnrichedAnalytics, not even an explicit nil
### GetUserAccessLevel

`func (o *PatchedFeatureFlag) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *PatchedFeatureFlag) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *PatchedFeatureFlag) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.

### HasUserAccessLevel

`func (o *PatchedFeatureFlag) HasUserAccessLevel() bool`

HasUserAccessLevel returns a boolean if a field has been set.

### SetUserAccessLevelNil

`func (o *PatchedFeatureFlag) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *PatchedFeatureFlag) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetCreationContext

`func (o *PatchedFeatureFlag) GetCreationContext() CreationContextEnum`

GetCreationContext returns the CreationContext field if non-nil, zero value otherwise.

### GetCreationContextOk

`func (o *PatchedFeatureFlag) GetCreationContextOk() (*CreationContextEnum, bool)`

GetCreationContextOk returns a tuple with the CreationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationContext

`func (o *PatchedFeatureFlag) SetCreationContext(v CreationContextEnum)`

SetCreationContext sets CreationContext field to given value.

### HasCreationContext

`func (o *PatchedFeatureFlag) HasCreationContext() bool`

HasCreationContext returns a boolean if a field has been set.

### GetIsRemoteConfiguration

`func (o *PatchedFeatureFlag) GetIsRemoteConfiguration() bool`

GetIsRemoteConfiguration returns the IsRemoteConfiguration field if non-nil, zero value otherwise.

### GetIsRemoteConfigurationOk

`func (o *PatchedFeatureFlag) GetIsRemoteConfigurationOk() (*bool, bool)`

GetIsRemoteConfigurationOk returns a tuple with the IsRemoteConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteConfiguration

`func (o *PatchedFeatureFlag) SetIsRemoteConfiguration(v bool)`

SetIsRemoteConfiguration sets IsRemoteConfiguration field to given value.

### HasIsRemoteConfiguration

`func (o *PatchedFeatureFlag) HasIsRemoteConfiguration() bool`

HasIsRemoteConfiguration returns a boolean if a field has been set.

### SetIsRemoteConfigurationNil

`func (o *PatchedFeatureFlag) SetIsRemoteConfigurationNil(b bool)`

 SetIsRemoteConfigurationNil sets the value for IsRemoteConfiguration to be an explicit nil

### UnsetIsRemoteConfiguration
`func (o *PatchedFeatureFlag) UnsetIsRemoteConfiguration()`

UnsetIsRemoteConfiguration ensures that no value is present for IsRemoteConfiguration, not even an explicit nil
### GetHasEncryptedPayloads

`func (o *PatchedFeatureFlag) GetHasEncryptedPayloads() bool`

GetHasEncryptedPayloads returns the HasEncryptedPayloads field if non-nil, zero value otherwise.

### GetHasEncryptedPayloadsOk

`func (o *PatchedFeatureFlag) GetHasEncryptedPayloadsOk() (*bool, bool)`

GetHasEncryptedPayloadsOk returns a tuple with the HasEncryptedPayloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEncryptedPayloads

`func (o *PatchedFeatureFlag) SetHasEncryptedPayloads(v bool)`

SetHasEncryptedPayloads sets HasEncryptedPayloads field to given value.

### HasHasEncryptedPayloads

`func (o *PatchedFeatureFlag) HasHasEncryptedPayloads() bool`

HasHasEncryptedPayloads returns a boolean if a field has been set.

### SetHasEncryptedPayloadsNil

`func (o *PatchedFeatureFlag) SetHasEncryptedPayloadsNil(b bool)`

 SetHasEncryptedPayloadsNil sets the value for HasEncryptedPayloads to be an explicit nil

### UnsetHasEncryptedPayloads
`func (o *PatchedFeatureFlag) UnsetHasEncryptedPayloads()`

UnsetHasEncryptedPayloads ensures that no value is present for HasEncryptedPayloads, not even an explicit nil
### GetStatus

`func (o *PatchedFeatureFlag) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedFeatureFlag) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedFeatureFlag) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedFeatureFlag) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEvaluationRuntime

`func (o *PatchedFeatureFlag) GetEvaluationRuntime() string`

GetEvaluationRuntime returns the EvaluationRuntime field if non-nil, zero value otherwise.

### GetEvaluationRuntimeOk

`func (o *PatchedFeatureFlag) GetEvaluationRuntimeOk() (*string, bool)`

GetEvaluationRuntimeOk returns a tuple with the EvaluationRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationRuntime

`func (o *PatchedFeatureFlag) SetEvaluationRuntime(v string)`

SetEvaluationRuntime sets EvaluationRuntime field to given value.

### HasEvaluationRuntime

`func (o *PatchedFeatureFlag) HasEvaluationRuntime() bool`

HasEvaluationRuntime returns a boolean if a field has been set.

### SetEvaluationRuntimeNil

`func (o *PatchedFeatureFlag) SetEvaluationRuntimeNil(b bool)`

 SetEvaluationRuntimeNil sets the value for EvaluationRuntime to be an explicit nil

### UnsetEvaluationRuntime
`func (o *PatchedFeatureFlag) UnsetEvaluationRuntime()`

UnsetEvaluationRuntime ensures that no value is present for EvaluationRuntime, not even an explicit nil
### GetLastCalledAt

`func (o *PatchedFeatureFlag) GetLastCalledAt() time.Time`

GetLastCalledAt returns the LastCalledAt field if non-nil, zero value otherwise.

### GetLastCalledAtOk

`func (o *PatchedFeatureFlag) GetLastCalledAtOk() (*time.Time, bool)`

GetLastCalledAtOk returns a tuple with the LastCalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalledAt

`func (o *PatchedFeatureFlag) SetLastCalledAt(v time.Time)`

SetLastCalledAt sets LastCalledAt field to given value.

### HasLastCalledAt

`func (o *PatchedFeatureFlag) HasLastCalledAt() bool`

HasLastCalledAt returns a boolean if a field has been set.

### SetLastCalledAtNil

`func (o *PatchedFeatureFlag) SetLastCalledAtNil(b bool)`

 SetLastCalledAtNil sets the value for LastCalledAt to be an explicit nil

### UnsetLastCalledAt
`func (o *PatchedFeatureFlag) UnsetLastCalledAt()`

UnsetLastCalledAt ensures that no value is present for LastCalledAt, not even an explicit nil
### GetCreateInFolder

`func (o *PatchedFeatureFlag) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedFeatureFlag) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedFeatureFlag) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedFeatureFlag) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetShouldCreateUsageDashboard

`func (o *PatchedFeatureFlag) GetShouldCreateUsageDashboard() bool`

GetShouldCreateUsageDashboard returns the ShouldCreateUsageDashboard field if non-nil, zero value otherwise.

### GetShouldCreateUsageDashboardOk

`func (o *PatchedFeatureFlag) GetShouldCreateUsageDashboardOk() (*bool, bool)`

GetShouldCreateUsageDashboardOk returns a tuple with the ShouldCreateUsageDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateUsageDashboard

`func (o *PatchedFeatureFlag) SetShouldCreateUsageDashboard(v bool)`

SetShouldCreateUsageDashboard sets ShouldCreateUsageDashboard field to given value.

### HasShouldCreateUsageDashboard

`func (o *PatchedFeatureFlag) HasShouldCreateUsageDashboard() bool`

HasShouldCreateUsageDashboard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


