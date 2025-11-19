# Cohort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **interface{}** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Filters** | Pointer to **interface{}** | Filters for the cohort. Examples:          # Behavioral filter (performed event)         {             \&quot;properties\&quot;: {                 \&quot;type\&quot;: \&quot;OR\&quot;,                 \&quot;values\&quot;: [{                     \&quot;type\&quot;: \&quot;OR\&quot;,                     \&quot;values\&quot;: [{                         \&quot;key\&quot;: \&quot;address page viewed\&quot;,                         \&quot;type\&quot;: \&quot;behavioral\&quot;,                         \&quot;value\&quot;: \&quot;performed_event\&quot;,                         \&quot;negation\&quot;: false,                         \&quot;event_type\&quot;: \&quot;events\&quot;,                         \&quot;time_value\&quot;: \&quot;30\&quot;,                         \&quot;time_interval\&quot;: \&quot;day\&quot;                     }]                 }]             }         }          # Person property filter         {             \&quot;properties\&quot;: {                 \&quot;type\&quot;: \&quot;OR\&quot;,                 \&quot;values\&quot;: [{                     \&quot;type\&quot;: \&quot;AND\&quot;,                     \&quot;values\&quot;: [{                         \&quot;key\&quot;: \&quot;promoCodes\&quot;,                         \&quot;type\&quot;: \&quot;person\&quot;,                         \&quot;value\&quot;: [\&quot;1234567890\&quot;],                         \&quot;negation\&quot;: false,                         \&quot;operator\&quot;: \&quot;exact\&quot;                     }]                 }]             }         }          # Cohort filter         {             \&quot;properties\&quot;: {                 \&quot;type\&quot;: \&quot;OR\&quot;,                 \&quot;values\&quot;: [{                     \&quot;type\&quot;: \&quot;AND\&quot;,                     \&quot;values\&quot;: [{                         \&quot;key\&quot;: \&quot;id\&quot;,                         \&quot;type\&quot;: \&quot;cohort\&quot;,                         \&quot;value\&quot;: 8814,                         \&quot;negation\&quot;: false                     }]                 }]             }         } | [optional] 
**Query** | Pointer to **interface{}** |  | [optional] 
**IsCalculating** | **bool** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**CreatedAt** | **NullableTime** |  | [readonly] 
**LastCalculation** | **NullableTime** |  | [readonly] 
**ErrorsCalculating** | **int32** |  | [readonly] 
**Count** | **NullableInt32** |  | [readonly] 
**IsStatic** | Pointer to **bool** |  | [optional] 
**CohortType** | Pointer to **NullableString** | Type of cohort based on filter complexity  * &#x60;static&#x60; - static * &#x60;person_property&#x60; - person_property * &#x60;behavioral&#x60; - behavioral * &#x60;realtime&#x60; - realtime * &#x60;analytical&#x60; - analytical | [optional] 
**ExperimentSet** | **[]int32** |  | [readonly] 
**CreateInFolder** | Pointer to **string** |  | [optional] 
**CreateStaticPersonIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCohort

`func NewCohort(id int32, isCalculating bool, createdBy UserBasic, createdAt NullableTime, lastCalculation NullableTime, errorsCalculating int32, count NullableInt32, experimentSet []int32, ) *Cohort`

NewCohort instantiates a new Cohort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCohortWithDefaults

`func NewCohortWithDefaults() *Cohort`

NewCohortWithDefaults instantiates a new Cohort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cohort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cohort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cohort) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Cohort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cohort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cohort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cohort) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Cohort) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Cohort) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Cohort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cohort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cohort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cohort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroups

`func (o *Cohort) GetGroups() interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Cohort) GetGroupsOk() (*interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Cohort) SetGroups(v interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Cohort) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *Cohort) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *Cohort) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetDeleted

`func (o *Cohort) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Cohort) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Cohort) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Cohort) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFilters

`func (o *Cohort) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Cohort) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Cohort) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Cohort) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *Cohort) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *Cohort) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetQuery

`func (o *Cohort) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Cohort) GetQueryOk() (*interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Cohort) SetQuery(v interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Cohort) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *Cohort) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *Cohort) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetIsCalculating

`func (o *Cohort) GetIsCalculating() bool`

GetIsCalculating returns the IsCalculating field if non-nil, zero value otherwise.

### GetIsCalculatingOk

`func (o *Cohort) GetIsCalculatingOk() (*bool, bool)`

GetIsCalculatingOk returns a tuple with the IsCalculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCalculating

`func (o *Cohort) SetIsCalculating(v bool)`

SetIsCalculating sets IsCalculating field to given value.


### GetCreatedBy

`func (o *Cohort) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Cohort) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Cohort) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Cohort) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cohort) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cohort) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Cohort) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Cohort) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastCalculation

`func (o *Cohort) GetLastCalculation() time.Time`

GetLastCalculation returns the LastCalculation field if non-nil, zero value otherwise.

### GetLastCalculationOk

`func (o *Cohort) GetLastCalculationOk() (*time.Time, bool)`

GetLastCalculationOk returns a tuple with the LastCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCalculation

`func (o *Cohort) SetLastCalculation(v time.Time)`

SetLastCalculation sets LastCalculation field to given value.


### SetLastCalculationNil

`func (o *Cohort) SetLastCalculationNil(b bool)`

 SetLastCalculationNil sets the value for LastCalculation to be an explicit nil

### UnsetLastCalculation
`func (o *Cohort) UnsetLastCalculation()`

UnsetLastCalculation ensures that no value is present for LastCalculation, not even an explicit nil
### GetErrorsCalculating

`func (o *Cohort) GetErrorsCalculating() int32`

GetErrorsCalculating returns the ErrorsCalculating field if non-nil, zero value otherwise.

### GetErrorsCalculatingOk

`func (o *Cohort) GetErrorsCalculatingOk() (*int32, bool)`

GetErrorsCalculatingOk returns a tuple with the ErrorsCalculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCalculating

`func (o *Cohort) SetErrorsCalculating(v int32)`

SetErrorsCalculating sets ErrorsCalculating field to given value.


### GetCount

`func (o *Cohort) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Cohort) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Cohort) SetCount(v int32)`

SetCount sets Count field to given value.


### SetCountNil

`func (o *Cohort) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *Cohort) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetIsStatic

`func (o *Cohort) GetIsStatic() bool`

GetIsStatic returns the IsStatic field if non-nil, zero value otherwise.

### GetIsStaticOk

`func (o *Cohort) GetIsStaticOk() (*bool, bool)`

GetIsStaticOk returns a tuple with the IsStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatic

`func (o *Cohort) SetIsStatic(v bool)`

SetIsStatic sets IsStatic field to given value.

### HasIsStatic

`func (o *Cohort) HasIsStatic() bool`

HasIsStatic returns a boolean if a field has been set.

### GetCohortType

`func (o *Cohort) GetCohortType() string`

GetCohortType returns the CohortType field if non-nil, zero value otherwise.

### GetCohortTypeOk

`func (o *Cohort) GetCohortTypeOk() (*string, bool)`

GetCohortTypeOk returns a tuple with the CohortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohortType

`func (o *Cohort) SetCohortType(v string)`

SetCohortType sets CohortType field to given value.

### HasCohortType

`func (o *Cohort) HasCohortType() bool`

HasCohortType returns a boolean if a field has been set.

### SetCohortTypeNil

`func (o *Cohort) SetCohortTypeNil(b bool)`

 SetCohortTypeNil sets the value for CohortType to be an explicit nil

### UnsetCohortType
`func (o *Cohort) UnsetCohortType()`

UnsetCohortType ensures that no value is present for CohortType, not even an explicit nil
### GetExperimentSet

`func (o *Cohort) GetExperimentSet() []int32`

GetExperimentSet returns the ExperimentSet field if non-nil, zero value otherwise.

### GetExperimentSetOk

`func (o *Cohort) GetExperimentSetOk() (*[]int32, bool)`

GetExperimentSetOk returns a tuple with the ExperimentSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentSet

`func (o *Cohort) SetExperimentSet(v []int32)`

SetExperimentSet sets ExperimentSet field to given value.


### GetCreateInFolder

`func (o *Cohort) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *Cohort) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *Cohort) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *Cohort) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.

### GetCreateStaticPersonIds

`func (o *Cohort) GetCreateStaticPersonIds() []string`

GetCreateStaticPersonIds returns the CreateStaticPersonIds field if non-nil, zero value otherwise.

### GetCreateStaticPersonIdsOk

`func (o *Cohort) GetCreateStaticPersonIdsOk() (*[]string, bool)`

GetCreateStaticPersonIdsOk returns a tuple with the CreateStaticPersonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateStaticPersonIds

`func (o *Cohort) SetCreateStaticPersonIds(v []string)`

SetCreateStaticPersonIds sets CreateStaticPersonIds field to given value.

### HasCreateStaticPersonIds

`func (o *Cohort) HasCreateStaticPersonIds() bool`

HasCreateStaticPersonIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


