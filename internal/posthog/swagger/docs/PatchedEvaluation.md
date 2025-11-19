# PatchedEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EvaluationType** | Pointer to [**EvaluationTypeEnum**](EvaluationTypeEnum.md) |  | [optional] 
**EvaluationConfig** | Pointer to **interface{}** |  | [optional] 
**OutputType** | Pointer to [**OutputTypeEnum**](OutputTypeEnum.md) |  | [optional] 
**OutputConfig** | Pointer to **interface{}** |  | [optional] 
**Conditions** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedEvaluation

`func NewPatchedEvaluation() *PatchedEvaluation`

NewPatchedEvaluation instantiates a new PatchedEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEvaluationWithDefaults

`func NewPatchedEvaluationWithDefaults() *PatchedEvaluation`

NewPatchedEvaluationWithDefaults instantiates a new PatchedEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedEvaluation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedEvaluation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedEvaluation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedEvaluation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedEvaluation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEvaluation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEvaluation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEvaluation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedEvaluation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedEvaluation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedEvaluation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedEvaluation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedEvaluation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedEvaluation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedEvaluation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedEvaluation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvaluationType

`func (o *PatchedEvaluation) GetEvaluationType() EvaluationTypeEnum`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *PatchedEvaluation) GetEvaluationTypeOk() (*EvaluationTypeEnum, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *PatchedEvaluation) SetEvaluationType(v EvaluationTypeEnum)`

SetEvaluationType sets EvaluationType field to given value.

### HasEvaluationType

`func (o *PatchedEvaluation) HasEvaluationType() bool`

HasEvaluationType returns a boolean if a field has been set.

### GetEvaluationConfig

`func (o *PatchedEvaluation) GetEvaluationConfig() interface{}`

GetEvaluationConfig returns the EvaluationConfig field if non-nil, zero value otherwise.

### GetEvaluationConfigOk

`func (o *PatchedEvaluation) GetEvaluationConfigOk() (*interface{}, bool)`

GetEvaluationConfigOk returns a tuple with the EvaluationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationConfig

`func (o *PatchedEvaluation) SetEvaluationConfig(v interface{})`

SetEvaluationConfig sets EvaluationConfig field to given value.

### HasEvaluationConfig

`func (o *PatchedEvaluation) HasEvaluationConfig() bool`

HasEvaluationConfig returns a boolean if a field has been set.

### SetEvaluationConfigNil

`func (o *PatchedEvaluation) SetEvaluationConfigNil(b bool)`

 SetEvaluationConfigNil sets the value for EvaluationConfig to be an explicit nil

### UnsetEvaluationConfig
`func (o *PatchedEvaluation) UnsetEvaluationConfig()`

UnsetEvaluationConfig ensures that no value is present for EvaluationConfig, not even an explicit nil
### GetOutputType

`func (o *PatchedEvaluation) GetOutputType() OutputTypeEnum`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *PatchedEvaluation) GetOutputTypeOk() (*OutputTypeEnum, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *PatchedEvaluation) SetOutputType(v OutputTypeEnum)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *PatchedEvaluation) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetOutputConfig

`func (o *PatchedEvaluation) GetOutputConfig() interface{}`

GetOutputConfig returns the OutputConfig field if non-nil, zero value otherwise.

### GetOutputConfigOk

`func (o *PatchedEvaluation) GetOutputConfigOk() (*interface{}, bool)`

GetOutputConfigOk returns a tuple with the OutputConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputConfig

`func (o *PatchedEvaluation) SetOutputConfig(v interface{})`

SetOutputConfig sets OutputConfig field to given value.

### HasOutputConfig

`func (o *PatchedEvaluation) HasOutputConfig() bool`

HasOutputConfig returns a boolean if a field has been set.

### SetOutputConfigNil

`func (o *PatchedEvaluation) SetOutputConfigNil(b bool)`

 SetOutputConfigNil sets the value for OutputConfig to be an explicit nil

### UnsetOutputConfig
`func (o *PatchedEvaluation) UnsetOutputConfig()`

UnsetOutputConfig ensures that no value is present for OutputConfig, not even an explicit nil
### GetConditions

`func (o *PatchedEvaluation) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PatchedEvaluation) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PatchedEvaluation) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PatchedEvaluation) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *PatchedEvaluation) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PatchedEvaluation) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetCreatedAt

`func (o *PatchedEvaluation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedEvaluation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedEvaluation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedEvaluation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedEvaluation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedEvaluation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedEvaluation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedEvaluation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedEvaluation) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedEvaluation) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedEvaluation) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedEvaluation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedEvaluation) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedEvaluation) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedEvaluation) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedEvaluation) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


