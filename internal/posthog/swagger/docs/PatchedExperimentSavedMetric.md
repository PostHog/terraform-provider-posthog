# PatchedExperimentSavedMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to **interface{}** |  | [optional] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Tags** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewPatchedExperimentSavedMetric

`func NewPatchedExperimentSavedMetric() *PatchedExperimentSavedMetric`

NewPatchedExperimentSavedMetric instantiates a new PatchedExperimentSavedMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedExperimentSavedMetricWithDefaults

`func NewPatchedExperimentSavedMetricWithDefaults() *PatchedExperimentSavedMetric`

NewPatchedExperimentSavedMetricWithDefaults instantiates a new PatchedExperimentSavedMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedExperimentSavedMetric) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedExperimentSavedMetric) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedExperimentSavedMetric) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedExperimentSavedMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedExperimentSavedMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedExperimentSavedMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedExperimentSavedMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedExperimentSavedMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedExperimentSavedMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedExperimentSavedMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedExperimentSavedMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedExperimentSavedMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedExperimentSavedMetric) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedExperimentSavedMetric) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuery

`func (o *PatchedExperimentSavedMetric) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PatchedExperimentSavedMetric) GetQueryOk() (*interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PatchedExperimentSavedMetric) SetQuery(v interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *PatchedExperimentSavedMetric) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *PatchedExperimentSavedMetric) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *PatchedExperimentSavedMetric) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetCreatedBy

`func (o *PatchedExperimentSavedMetric) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedExperimentSavedMetric) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedExperimentSavedMetric) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedExperimentSavedMetric) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedExperimentSavedMetric) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedExperimentSavedMetric) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedExperimentSavedMetric) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedExperimentSavedMetric) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedExperimentSavedMetric) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedExperimentSavedMetric) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedExperimentSavedMetric) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedExperimentSavedMetric) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *PatchedExperimentSavedMetric) GetTags() []interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedExperimentSavedMetric) GetTagsOk() (*[]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedExperimentSavedMetric) SetTags(v []interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedExperimentSavedMetric) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


