# FeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | Pointer to **string** | contains the description for the flag (field name &#x60;name&#x60; is kept for backwards-compatibility) | [optional] 
**Key** | **string** |  | 
**Filters** | Pointer to **map[string]interface{}** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | **NullableTime** |  | [readonly] 
**Version** | Pointer to **int32** |  | [optional] [default to 0]
**LastModifiedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**IsSimpleFlag** | **bool** |  | [readonly] 
**RolloutPercentage** | **NullableInt32** |  | [readonly] 
**EnsureExperienceContinuity** | Pointer to **NullableBool** |  | [optional] 
**ExperimentSet** | **string** |  | [readonly] 
**Surveys** | **map[string]interface{}** |  | [readonly] 
**Features** | **map[string]interface{}** |  | [readonly] 
**RollbackConditions** | Pointer to **interface{}** |  | [optional] 
**PerformedRollback** | Pointer to **NullableBool** |  | [optional] 
**CanEdit** | **bool** |  | [readonly] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 
**EvaluationTags** | Pointer to **[]interface{}** |  | [optional] 
**UsageDashboard** | **int32** |  | [readonly] 
**AnalyticsDashboards** | Pointer to **[]int32** |  | [optional] 
**HasEnrichedAnalytics** | Pointer to **NullableBool** |  | [optional] 
**UserAccessLevel** | **NullableString** | The effective access level the user has for this object | [readonly] 
**CreationContext** | Pointer to [**CreationContextEnum**](CreationContextEnum.md) | Indicates the origin product of the feature flag. Choices: &#39;feature_flags&#39;, &#39;experiments&#39;, &#39;surveys&#39;, &#39;early_access_features&#39;, &#39;web_experiments&#39;.  * &#x60;feature_flags&#x60; - feature_flags * &#x60;experiments&#x60; - experiments * &#x60;surveys&#x60; - surveys * &#x60;early_access_features&#x60; - early_access_features * &#x60;web_experiments&#x60; - web_experiments | [optional] 
**IsRemoteConfiguration** | Pointer to **NullableBool** |  | [optional] 
**HasEncryptedPayloads** | Pointer to **NullableBool** |  | [optional] 
**Status** | **string** |  | [readonly] 
**EvaluationRuntime** | Pointer to **NullableString** | Specifies where this feature flag should be evaluated  * &#x60;server&#x60; - Server * &#x60;client&#x60; - Client * &#x60;all&#x60; - All | [optional] 
**LastCalledAt** | Pointer to **NullableTime** | Last time this feature flag was called (from $feature_flag_called events) | [optional] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**ShouldCreateUsageDashboard** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewFeatureFlag

`func NewFeatureFlag(id int32, key string, createdBy UserBasic, updatedAt NullableTime, lastModifiedBy UserBasic, isSimpleFlag bool, rolloutPercentage NullableInt32, experimentSet string, surveys map[string]interface{}, features map[string]interface{}, canEdit bool, usageDashboard int32, userAccessLevel NullableString, status string, ) *FeatureFlag`

NewFeatureFlag instantiates a new FeatureFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagWithDefaults

`func NewFeatureFlagWithDefaults() *FeatureFlag`

NewFeatureFlagWithDefaults instantiates a new FeatureFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureFlag) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureFlag) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureFlag) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FeatureFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureFlag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FeatureFlag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *FeatureFlag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FeatureFlag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FeatureFlag) SetKey(v string)`

SetKey sets Key field to given value.


### GetFilters

`func (o *FeatureFlag) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *FeatureFlag) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *FeatureFlag) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *FeatureFlag) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetDeleted

`func (o *FeatureFlag) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *FeatureFlag) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *FeatureFlag) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *FeatureFlag) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetActive

`func (o *FeatureFlag) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FeatureFlag) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FeatureFlag) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FeatureFlag) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FeatureFlag) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FeatureFlag) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FeatureFlag) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *FeatureFlag) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeatureFlag) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeatureFlag) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FeatureFlag) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FeatureFlag) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FeatureFlag) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FeatureFlag) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *FeatureFlag) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FeatureFlag) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetVersion

`func (o *FeatureFlag) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeatureFlag) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeatureFlag) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FeatureFlag) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *FeatureFlag) GetLastModifiedBy() UserBasic`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *FeatureFlag) GetLastModifiedByOk() (*UserBasic, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *FeatureFlag) SetLastModifiedBy(v UserBasic)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetIsSimpleFlag

`func (o *FeatureFlag) GetIsSimpleFlag() bool`

GetIsSimpleFlag returns the IsSimpleFlag field if non-nil, zero value otherwise.

### GetIsSimpleFlagOk

`func (o *FeatureFlag) GetIsSimpleFlagOk() (*bool, bool)`

GetIsSimpleFlagOk returns a tuple with the IsSimpleFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSimpleFlag

`func (o *FeatureFlag) SetIsSimpleFlag(v bool)`

SetIsSimpleFlag sets IsSimpleFlag field to given value.


### GetRolloutPercentage

`func (o *FeatureFlag) GetRolloutPercentage() int32`

GetRolloutPercentage returns the RolloutPercentage field if non-nil, zero value otherwise.

### GetRolloutPercentageOk

`func (o *FeatureFlag) GetRolloutPercentageOk() (*int32, bool)`

GetRolloutPercentageOk returns a tuple with the RolloutPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentage

`func (o *FeatureFlag) SetRolloutPercentage(v int32)`

SetRolloutPercentage sets RolloutPercentage field to given value.


### SetRolloutPercentageNil

`func (o *FeatureFlag) SetRolloutPercentageNil(b bool)`

 SetRolloutPercentageNil sets the value for RolloutPercentage to be an explicit nil

### UnsetRolloutPercentage
`func (o *FeatureFlag) UnsetRolloutPercentage()`

UnsetRolloutPercentage ensures that no value is present for RolloutPercentage, not even an explicit nil
### GetEnsureExperienceContinuity

`func (o *FeatureFlag) GetEnsureExperienceContinuity() bool`

GetEnsureExperienceContinuity returns the EnsureExperienceContinuity field if non-nil, zero value otherwise.

### GetEnsureExperienceContinuityOk

`func (o *FeatureFlag) GetEnsureExperienceContinuityOk() (*bool, bool)`

GetEnsureExperienceContinuityOk returns a tuple with the EnsureExperienceContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureExperienceContinuity

`func (o *FeatureFlag) SetEnsureExperienceContinuity(v bool)`

SetEnsureExperienceContinuity sets EnsureExperienceContinuity field to given value.

### HasEnsureExperienceContinuity

`func (o *FeatureFlag) HasEnsureExperienceContinuity() bool`

HasEnsureExperienceContinuity returns a boolean if a field has been set.

### SetEnsureExperienceContinuityNil

`func (o *FeatureFlag) SetEnsureExperienceContinuityNil(b bool)`

 SetEnsureExperienceContinuityNil sets the value for EnsureExperienceContinuity to be an explicit nil

### UnsetEnsureExperienceContinuity
`func (o *FeatureFlag) UnsetEnsureExperienceContinuity()`

UnsetEnsureExperienceContinuity ensures that no value is present for EnsureExperienceContinuity, not even an explicit nil
### GetExperimentSet

`func (o *FeatureFlag) GetExperimentSet() string`

GetExperimentSet returns the ExperimentSet field if non-nil, zero value otherwise.

### GetExperimentSetOk

`func (o *FeatureFlag) GetExperimentSetOk() (*string, bool)`

GetExperimentSetOk returns a tuple with the ExperimentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentSet

`func (o *FeatureFlag) SetExperimentSet(v string)`

SetExperimentSet sets ExperimentSet field to given value.


### GetSurveys

`func (o *FeatureFlag) GetSurveys() map[string]interface{}`

GetSurveys returns the Surveys field if non-nil, zero value otherwise.

### GetSurveysOk

`func (o *FeatureFlag) GetSurveysOk() (*map[string]interface{}, bool)`

GetSurveysOk returns a tuple with the Surveys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveys

`func (o *FeatureFlag) SetSurveys(v map[string]interface{})`

SetSurveys sets Surveys field to given value.


### GetFeatures

`func (o *FeatureFlag) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FeatureFlag) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FeatureFlag) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.


### GetRollbackConditions

`func (o *FeatureFlag) GetRollbackConditions() interface{}`

GetRollbackConditions returns the RollbackConditions field if non-nil, zero value otherwise.

### GetRollbackConditionsOk

`func (o *FeatureFlag) GetRollbackConditionsOk() (*interface{}, bool)`

GetRollbackConditionsOk returns a tuple with the RollbackConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackConditions

`func (o *FeatureFlag) SetRollbackConditions(v interface{})`

SetRollbackConditions sets RollbackConditions field to given value.

### HasRollbackConditions

`func (o *FeatureFlag) HasRollbackConditions() bool`

HasRollbackConditions returns a boolean if a field has been set.

### SetRollbackConditionsNil

`func (o *FeatureFlag) SetRollbackConditionsNil(b bool)`

 SetRollbackConditionsNil sets the value for RollbackConditions to be an explicit nil

### UnsetRollbackConditions
`func (o *FeatureFlag) UnsetRollbackConditions()`

UnsetRollbackConditions ensures that no value is present for RollbackConditions, not even an explicit nil
### GetPerformedRollback

`func (o *FeatureFlag) GetPerformedRollback() bool`

GetPerformedRollback returns the PerformedRollback field if non-nil, zero value otherwise.

### GetPerformedRollbackOk

`func (o *FeatureFlag) GetPerformedRollbackOk() (*bool, bool)`

GetPerformedRollbackOk returns a tuple with the PerformedRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedRollback

`func (o *FeatureFlag) SetPerformedRollback(v bool)`

SetPerformedRollback sets PerformedRollback field to given value.

### HasPerformedRollback

`func (o *FeatureFlag) HasPerformedRollback() bool`

HasPerformedRollback returns a boolean if a field has been set.

### SetPerformedRollbackNil

`func (o *FeatureFlag) SetPerformedRollbackNil(b bool)`

 SetPerformedRollbackNil sets the value for PerformedRollback to be an explicit nil

### UnsetPerformedRollback
`func (o *FeatureFlag) UnsetPerformedRollback()`

UnsetPerformedRollback ensures that no value is present for PerformedRollback, not even an explicit nil
### GetCanEdit

`func (o *FeatureFlag) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *FeatureFlag) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *FeatureFlag) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.


### GetTags

`func (o *FeatureFlag) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FeatureFlag) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FeatureFlag) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FeatureFlag) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEvaluationTags

`func (o *FeatureFlag) GetEvaluationTags() []interface{}`

GetEvaluationTags returns the EvaluationTags field if non-nil, zero value otherwise.

### GetEvaluationTagsOk

`func (o *FeatureFlag) GetEvaluationTagsOk() (*[]interface{}, bool)`

GetEvaluationTagsOk returns a tuple with the EvaluationTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTags

`func (o *FeatureFlag) SetEvaluationTags(v []interface{})`

SetEvaluationTags sets EvaluationTags field to given value.

### HasEvaluationTags

`func (o *FeatureFlag) HasEvaluationTags() bool`

HasEvaluationTags returns a boolean if a field has been set.

### GetUsageDashboard

`func (o *FeatureFlag) GetUsageDashboard() int32`

GetUsageDashboard returns the UsageDashboard field if non-nil, zero value otherwise.

### GetUsageDashboardOk

`func (o *FeatureFlag) GetUsageDashboardOk() (*int32, bool)`

GetUsageDashboardOk returns a tuple with the UsageDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDashboard

`func (o *FeatureFlag) SetUsageDashboard(v int32)`

SetUsageDashboard sets UsageDashboard field to given value.


### GetAnalyticsDashboards

`func (o *FeatureFlag) GetAnalyticsDashboards() []int32`

GetAnalyticsDashboards returns the AnalyticsDashboards field if non-nil, zero value otherwise.

### GetAnalyticsDashboardsOk

`func (o *FeatureFlag) GetAnalyticsDashboardsOk() (*[]int32, bool)`

GetAnalyticsDashboardsOk returns a tuple with the AnalyticsDashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsDashboards

`func (o *FeatureFlag) SetAnalyticsDashboards(v []int32)`

SetAnalyticsDashboards sets AnalyticsDashboards field to given value.

### HasAnalyticsDashboards

`func (o *FeatureFlag) HasAnalyticsDashboards() bool`

HasAnalyticsDashboards returns a boolean if a field has been set.

### GetHasEnrichedAnalytics

`func (o *FeatureFlag) GetHasEnrichedAnalytics() bool`

GetHasEnrichedAnalytics returns the HasEnrichedAnalytics field if non-nil, zero value otherwise.

### GetHasEnrichedAnalyticsOk

`func (o *FeatureFlag) GetHasEnrichedAnalyticsOk() (*bool, bool)`

GetHasEnrichedAnalyticsOk returns a tuple with the HasEnrichedAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEnrichedAnalytics

`func (o *FeatureFlag) SetHasEnrichedAnalytics(v bool)`

SetHasEnrichedAnalytics sets HasEnrichedAnalytics field to given value.

### HasHasEnrichedAnalytics

`func (o *FeatureFlag) HasHasEnrichedAnalytics() bool`

HasHasEnrichedAnalytics returns a boolean if a field has been set.

### SetHasEnrichedAnalyticsNil

`func (o *FeatureFlag) SetHasEnrichedAnalyticsNil(b bool)`

 SetHasEnrichedAnalyticsNil sets the value for HasEnrichedAnalytics to be an explicit nil

### UnsetHasEnrichedAnalytics
`func (o *FeatureFlag) UnsetHasEnrichedAnalytics()`

UnsetHasEnrichedAnalytics ensures that no value is present for HasEnrichedAnalytics, not even an explicit nil
### GetUserAccessLevel

`func (o *FeatureFlag) GetUserAccessLevel() string`

GetUserAccessLevel returns the UserAccessLevel field if non-nil, zero value otherwise.

### GetUserAccessLevelOk

`func (o *FeatureFlag) GetUserAccessLevelOk() (*string, bool)`

GetUserAccessLevelOk returns a tuple with the UserAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessLevel

`func (o *FeatureFlag) SetUserAccessLevel(v string)`

SetUserAccessLevel sets UserAccessLevel field to given value.


### SetUserAccessLevelNil

`func (o *FeatureFlag) SetUserAccessLevelNil(b bool)`

 SetUserAccessLevelNil sets the value for UserAccessLevel to be an explicit nil

### UnsetUserAccessLevel
`func (o *FeatureFlag) UnsetUserAccessLevel()`

UnsetUserAccessLevel ensures that no value is present for UserAccessLevel, not even an explicit nil
### GetCreationContext

`func (o *FeatureFlag) GetCreationContext() CreationContextEnum`

GetCreationContext returns the CreationContext field if non-nil, zero value otherwise.

### GetCreationContextOk

`func (o *FeatureFlag) GetCreationContextOk() (*CreationContextEnum, bool)`

GetCreationContextOk returns a tuple with the CreationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationContext

`func (o *FeatureFlag) SetCreationContext(v CreationContextEnum)`

SetCreationContext sets CreationContext field to given value.

### HasCreationContext

`func (o *FeatureFlag) HasCreationContext() bool`

HasCreationContext returns a boolean if a field has been set.

### GetIsRemoteConfiguration

`func (o *FeatureFlag) GetIsRemoteConfiguration() bool`

GetIsRemoteConfiguration returns the IsRemoteConfiguration field if non-nil, zero value otherwise.

### GetIsRemoteConfigurationOk

`func (o *FeatureFlag) GetIsRemoteConfigurationOk() (*bool, bool)`

GetIsRemoteConfigurationOk returns a tuple with the IsRemoteConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteConfiguration

`func (o *FeatureFlag) SetIsRemoteConfiguration(v bool)`

SetIsRemoteConfiguration sets IsRemoteConfiguration field to given value.

### HasIsRemoteConfiguration

`func (o *FeatureFlag) HasIsRemoteConfiguration() bool`

HasIsRemoteConfiguration returns a boolean if a field has been set.

### SetIsRemoteConfigurationNil

`func (o *FeatureFlag) SetIsRemoteConfigurationNil(b bool)`

 SetIsRemoteConfigurationNil sets the value for IsRemoteConfiguration to be an explicit nil

### UnsetIsRemoteConfiguration
`func (o *FeatureFlag) UnsetIsRemoteConfiguration()`

UnsetIsRemoteConfiguration ensures that no value is present for IsRemoteConfiguration, not even an explicit nil
### GetHasEncryptedPayloads

`func (o *FeatureFlag) GetHasEncryptedPayloads() bool`

GetHasEncryptedPayloads returns the HasEncryptedPayloads field if non-nil, zero value otherwise.

### GetHasEncryptedPayloadsOk

`func (o *FeatureFlag) GetHasEncryptedPayloadsOk() (*bool, bool)`

GetHasEncryptedPayloadsOk returns a tuple with the HasEncryptedPayloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEncryptedPayloads

`func (o *FeatureFlag) SetHasEncryptedPayloads(v bool)`

SetHasEncryptedPayloads sets HasEncryptedPayloads field to given value.

### HasHasEncryptedPayloads

`func (o *FeatureFlag) HasHasEncryptedPayloads() bool`

HasHasEncryptedPayloads returns a boolean if a field has been set.

### SetHasEncryptedPayloadsNil

`func (o *FeatureFlag) SetHasEncryptedPayloadsNil(b bool)`

 SetHasEncryptedPayloadsNil sets the value for HasEncryptedPayloads to be an explicit nil

### UnsetHasEncryptedPayloads
`func (o *FeatureFlag) UnsetHasEncryptedPayloads()`

UnsetHasEncryptedPayloads ensures that no value is present for HasEncryptedPayloads, not even an explicit nil
### GetStatus

`func (o *FeatureFlag) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeatureFlag) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeatureFlag) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEvaluationRuntime

`func (o *FeatureFlag) GetEvaluationRuntime() string`

GetEvaluationRuntime returns the EvaluationRuntime field if non-nil, zero value otherwise.

### GetEvaluationRuntimeOk

`func (o *FeatureFlag) GetEvaluationRuntimeOk() (*string, bool)`

GetEvaluationRuntimeOk returns a tuple with the EvaluationRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationRuntime

`func (o *FeatureFlag) SetEvaluationRuntime(v string)`

SetEvaluationRuntime sets EvaluationRuntime field to given value.

### HasEvaluationRuntime

`func (o *FeatureFlag) HasEvaluationRuntime() bool`

HasEvaluationRuntime returns a boolean if a field has been set.

### SetEvaluationRuntimeNil

`func (o *FeatureFlag) SetEvaluationRuntimeNil(b bool)`

 SetEvaluationRuntimeNil sets the value for EvaluationRuntime to be an explicit nil

### UnsetEvaluationRuntime
`func (o *FeatureFlag) UnsetEvaluationRuntime()`

UnsetEvaluationRuntime ensures that no value is present for EvaluationRuntime, not even an explicit nil
### GetLastCalledAt

`func (o *FeatureFlag) GetLastCalledAt() time.Time`

GetLastCalledAt returns the LastCalledAt field if non-nil, zero value otherwise.

### GetLastCalledAtOk

`func (o *FeatureFlag) GetLastCalledAtOk() (*time.Time, bool)`

GetLastCalledAtOk returns a tuple with the LastCalledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalledAt

`func (o *FeatureFlag) SetLastCalledAt(v time.Time)`

SetLastCalledAt sets LastCalledAt field to given value.

### HasLastCalledAt

`func (o *FeatureFlag) HasLastCalledAt() bool`

HasLastCalledAt returns a boolean if a field has been set.

### SetLastCalledAtNil

`func (o *FeatureFlag) SetLastCalledAtNil(b bool)`

 SetLastCalledAtNil sets the value for LastCalledAt to be an explicit nil

### UnsetLastCalledAt
`func (o *FeatureFlag) UnsetLastCalledAt()`

UnsetLastCalledAt ensures that no value is present for LastCalledAt, not even an explicit nil
### GetCreateInFolder

`func (o *FeatureFlag) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *FeatureFlag) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *FeatureFlag) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *FeatureFlag) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetShouldCreateUsageDashboard

`func (o *FeatureFlag) GetShouldCreateUsageDashboard() bool`

GetShouldCreateUsageDashboard returns the ShouldCreateUsageDashboard field if non-nil, zero value otherwise.

### GetShouldCreateUsageDashboardOk

`func (o *FeatureFlag) GetShouldCreateUsageDashboardOk() (*bool, bool)`

GetShouldCreateUsageDashboardOk returns a tuple with the ShouldCreateUsageDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldCreateUsageDashboard

`func (o *FeatureFlag) SetShouldCreateUsageDashboard(v bool)`

SetShouldCreateUsageDashboard sets ShouldCreateUsageDashboard field to given value.

### HasShouldCreateUsageDashboard

`func (o *FeatureFlag) HasShouldCreateUsageDashboard() bool`

HasShouldCreateUsageDashboard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


