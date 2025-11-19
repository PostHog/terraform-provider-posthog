# PatchedCohort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **interface{}** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Filters** | Pointer to **interface{}** | Filters for the cohort. Examples:          # Behavioral filter (performed event)         {             \&quot;properties\&quot;: {                 \&quot;type\&quot;: \&quot;OR\&quot;,                 \&quot;values\&quot;: [{                     \&quot;type\&quot;: \&quot;OR\&quot;,                     \&quot;values\&quot;: [{                         \&quot;key\&quot;: \&quot;address page viewed\&quot;,                         \&quot;type\&quot;: \&quot;behavioral\&quot;,                         \&quot;value\&quot;: \&quot;performed_event\&quot;,                         \&quot;negation\&quot;: false,                         \&quot;event_type\&quot;: \&quot;events\&quot;,                         \&quot;time_value\&quot;: \&quot;30\&quot;,                         \&quot;time_interval\&quot;: \&quot;day\&quot;                     }]                 }]             }         }          # Person property filter         {             \&quot;properties\&quot;: {                 \&quot;type\&quot;: \&quot;OR\&quot;,                 \&quot;values\&quot;: [{                     \&quot;type\&quot;: \&quot;AND\&quot;,                     \&quot;values\&quot;: [{                         \&quot;key\&quot;: \&quot;promoCodes\&quot;,                         \&quot;type\&quot;: \&quot;person\&quot;,                         \&quot;value\&quot;: [\&quot;1234567890\&quot;],                         \&quot;negation\&quot;: false,                         \&quot;operator\&quot;: \&quot;exact\&quot;                     }]                 }]             }         }          # Cohort filter         {             \&quot;properties\&quot;: {                 \&quot;type\&quot;: \&quot;OR\&quot;,                 \&quot;values\&quot;: [{                     \&quot;type\&quot;: \&quot;AND\&quot;,                     \&quot;values\&quot;: [{                         \&quot;key\&quot;: \&quot;id\&quot;,                         \&quot;type\&quot;: \&quot;cohort\&quot;,                         \&quot;value\&quot;: 8814,                         \&quot;negation\&quot;: false                     }]                 }]             }         } | [optional] 
**Query** | Pointer to **interface{}** |  | [optional] 
**IsCalculating** | Pointer to **bool** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastCalculation** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ErrorsCalculating** | Pointer to **int32** |  | [optional] [readonly] 
**Count** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**IsStatic** | Pointer to **bool** |  | [optional] 
**CohortType** | Pointer to **NullableString** | Type of cohort based on filter complexity  * &#x60;static&#x60; - static * &#x60;person_property&#x60; - person_property * &#x60;behavioral&#x60; - behavioral * &#x60;realtime&#x60; - realtime * &#x60;analytical&#x60; - analytical | [optional] 
**ExperimentSet** | Pointer to **[]int32** |  | [optional] [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**CreateStaticPersonIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPatchedCohort

`func NewPatchedCohort() *PatchedCohort`

NewPatchedCohort instantiates a new PatchedCohort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCohortWithDefaults

`func NewPatchedCohortWithDefaults() *PatchedCohort`

NewPatchedCohortWithDefaults instantiates a new PatchedCohort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedCohort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedCohort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedCohort) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedCohort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedCohort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCohort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCohort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCohort) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedCohort) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedCohort) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PatchedCohort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedCohort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedCohort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedCohort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroups

`func (o *PatchedCohort) GetGroups() interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PatchedCohort) GetGroupsOk() (*interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PatchedCohort) SetGroups(v interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PatchedCohort) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *PatchedCohort) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *PatchedCohort) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetDeleted

`func (o *PatchedCohort) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedCohort) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedCohort) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedCohort) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFilters

`func (o *PatchedCohort) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedCohort) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedCohort) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedCohort) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *PatchedCohort) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *PatchedCohort) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetQuery

`func (o *PatchedCohort) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PatchedCohort) GetQueryOk() (*interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PatchedCohort) SetQuery(v interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PatchedCohort) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *PatchedCohort) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *PatchedCohort) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetIsCalculating

`func (o *PatchedCohort) GetIsCalculating() bool`

GetIsCalculating returns the IsCalculating field if non-nil, zero value otherwise.

### GetIsCalculatingOk

`func (o *PatchedCohort) GetIsCalculatingOk() (*bool, bool)`

GetIsCalculatingOk returns a tuple with the IsCalculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCalculating

`func (o *PatchedCohort) SetIsCalculating(v bool)`

SetIsCalculating sets IsCalculating field to given value.

### HasIsCalculating

`func (o *PatchedCohort) HasIsCalculating() bool`

HasIsCalculating returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedCohort) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedCohort) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedCohort) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedCohort) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedCohort) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedCohort) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedCohort) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedCohort) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedCohort) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedCohort) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastCalculation

`func (o *PatchedCohort) GetLastCalculation() time.Time`

GetLastCalculation returns the LastCalculation field if non-nil, zero value otherwise.

### GetLastCalculationOk

`func (o *PatchedCohort) GetLastCalculationOk() (*time.Time, bool)`

GetLastCalculationOk returns a tuple with the LastCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalculation

`func (o *PatchedCohort) SetLastCalculation(v time.Time)`

SetLastCalculation sets LastCalculation field to given value.

### HasLastCalculation

`func (o *PatchedCohort) HasLastCalculation() bool`

HasLastCalculation returns a boolean if a field has been set.

### SetLastCalculationNil

`func (o *PatchedCohort) SetLastCalculationNil(b bool)`

 SetLastCalculationNil sets the value for LastCalculation to be an explicit nil

### UnsetLastCalculation
`func (o *PatchedCohort) UnsetLastCalculation()`

UnsetLastCalculation ensures that no value is present for LastCalculation, not even an explicit nil
### GetErrorsCalculating

`func (o *PatchedCohort) GetErrorsCalculating() int32`

GetErrorsCalculating returns the ErrorsCalculating field if non-nil, zero value otherwise.

### GetErrorsCalculatingOk

`func (o *PatchedCohort) GetErrorsCalculatingOk() (*int32, bool)`

GetErrorsCalculatingOk returns a tuple with the ErrorsCalculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCalculating

`func (o *PatchedCohort) SetErrorsCalculating(v int32)`

SetErrorsCalculating sets ErrorsCalculating field to given value.

### HasErrorsCalculating

`func (o *PatchedCohort) HasErrorsCalculating() bool`

HasErrorsCalculating returns a boolean if a field has been set.

### GetCount

`func (o *PatchedCohort) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PatchedCohort) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PatchedCohort) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PatchedCohort) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PatchedCohort) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PatchedCohort) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetIsStatic

`func (o *PatchedCohort) GetIsStatic() bool`

GetIsStatic returns the IsStatic field if non-nil, zero value otherwise.

### GetIsStaticOk

`func (o *PatchedCohort) GetIsStaticOk() (*bool, bool)`

GetIsStaticOk returns a tuple with the IsStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatic

`func (o *PatchedCohort) SetIsStatic(v bool)`

SetIsStatic sets IsStatic field to given value.

### HasIsStatic

`func (o *PatchedCohort) HasIsStatic() bool`

HasIsStatic returns a boolean if a field has been set.

### GetCohortType

`func (o *PatchedCohort) GetCohortType() string`

GetCohortType returns the CohortType field if non-nil, zero value otherwise.

### GetCohortTypeOk

`func (o *PatchedCohort) GetCohortTypeOk() (*string, bool)`

GetCohortTypeOk returns a tuple with the CohortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortType

`func (o *PatchedCohort) SetCohortType(v string)`

SetCohortType sets CohortType field to given value.

### HasCohortType

`func (o *PatchedCohort) HasCohortType() bool`

HasCohortType returns a boolean if a field has been set.

### SetCohortTypeNil

`func (o *PatchedCohort) SetCohortTypeNil(b bool)`

 SetCohortTypeNil sets the value for CohortType to be an explicit nil

### UnsetCohortType
`func (o *PatchedCohort) UnsetCohortType()`

UnsetCohortType ensures that no value is present for CohortType, not even an explicit nil
### GetExperimentSet

`func (o *PatchedCohort) GetExperimentSet() []int32`

GetExperimentSet returns the ExperimentSet field if non-nil, zero value otherwise.

### GetExperimentSetOk

`func (o *PatchedCohort) GetExperimentSetOk() (*[]int32, bool)`

GetExperimentSetOk returns a tuple with the ExperimentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentSet

`func (o *PatchedCohort) SetExperimentSet(v []int32)`

SetExperimentSet sets ExperimentSet field to given value.

### HasExperimentSet

`func (o *PatchedCohort) HasExperimentSet() bool`

HasExperimentSet returns a boolean if a field has been set.

### GetCreateInFolder

`func (o *PatchedCohort) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedCohort) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedCohort) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedCohort) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetCreateStaticPersonIds

`func (o *PatchedCohort) GetCreateStaticPersonIds() []string`

GetCreateStaticPersonIds returns the CreateStaticPersonIds field if non-nil, zero value otherwise.

### GetCreateStaticPersonIdsOk

`func (o *PatchedCohort) GetCreateStaticPersonIdsOk() (*[]string, bool)`

GetCreateStaticPersonIdsOk returns a tuple with the CreateStaticPersonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateStaticPersonIds

`func (o *PatchedCohort) SetCreateStaticPersonIds(v []string)`

SetCreateStaticPersonIds sets CreateStaticPersonIds field to given value.

### HasCreateStaticPersonIds

`func (o *PatchedCohort) HasCreateStaticPersonIds() bool`

HasCreateStaticPersonIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


