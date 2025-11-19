# Evaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EvaluationType** | [**EvaluationTypeEnum**](EvaluationTypeEnum.md) |  | 
**EvaluationConfig** | Pointer to **interface{}** |  | [optional] 
**OutputType** | [**OutputTypeEnum**](OutputTypeEnum.md) |  | 
**OutputConfig** | Pointer to **interface{}** |  | [optional] 
**Conditions** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**Deleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewEvaluation

`func NewEvaluation(id string, name string, evaluationType EvaluationTypeEnum, outputType OutputTypeEnum, createdAt time.Time, updatedAt time.Time, createdBy UserBasic, ) *Evaluation`

NewEvaluation instantiates a new Evaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationWithDefaults

`func NewEvaluationWithDefaults() *Evaluation`

NewEvaluationWithDefaults instantiates a new Evaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Evaluation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Evaluation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Evaluation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Evaluation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Evaluation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Evaluation) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Evaluation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Evaluation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Evaluation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Evaluation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Evaluation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Evaluation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Evaluation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Evaluation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvaluationType

`func (o *Evaluation) GetEvaluationType() EvaluationTypeEnum`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *Evaluation) GetEvaluationTypeOk() (*EvaluationTypeEnum, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *Evaluation) SetEvaluationType(v EvaluationTypeEnum)`

SetEvaluationType sets EvaluationType field to given value.


### GetEvaluationConfig

`func (o *Evaluation) GetEvaluationConfig() interface{}`

GetEvaluationConfig returns the EvaluationConfig field if non-nil, zero value otherwise.

### GetEvaluationConfigOk

`func (o *Evaluation) GetEvaluationConfigOk() (*interface{}, bool)`

GetEvaluationConfigOk returns a tuple with the EvaluationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationConfig

`func (o *Evaluation) SetEvaluationConfig(v interface{})`

SetEvaluationConfig sets EvaluationConfig field to given value.

### HasEvaluationConfig

`func (o *Evaluation) HasEvaluationConfig() bool`

HasEvaluationConfig returns a boolean if a field has been set.

### SetEvaluationConfigNil

`func (o *Evaluation) SetEvaluationConfigNil(b bool)`

 SetEvaluationConfigNil sets the value for EvaluationConfig to be an explicit nil

### UnsetEvaluationConfig
`func (o *Evaluation) UnsetEvaluationConfig()`

UnsetEvaluationConfig ensures that no value is present for EvaluationConfig, not even an explicit nil
### GetOutputType

`func (o *Evaluation) GetOutputType() OutputTypeEnum`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *Evaluation) GetOutputTypeOk() (*OutputTypeEnum, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *Evaluation) SetOutputType(v OutputTypeEnum)`

SetOutputType sets OutputType field to given value.


### GetOutputConfig

`func (o *Evaluation) GetOutputConfig() interface{}`

GetOutputConfig returns the OutputConfig field if non-nil, zero value otherwise.

### GetOutputConfigOk

`func (o *Evaluation) GetOutputConfigOk() (*interface{}, bool)`

GetOutputConfigOk returns a tuple with the OutputConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputConfig

`func (o *Evaluation) SetOutputConfig(v interface{})`

SetOutputConfig sets OutputConfig field to given value.

### HasOutputConfig

`func (o *Evaluation) HasOutputConfig() bool`

HasOutputConfig returns a boolean if a field has been set.

### SetOutputConfigNil

`func (o *Evaluation) SetOutputConfigNil(b bool)`

 SetOutputConfigNil sets the value for OutputConfig to be an explicit nil

### UnsetOutputConfig
`func (o *Evaluation) UnsetOutputConfig()`

UnsetOutputConfig ensures that no value is present for OutputConfig, not even an explicit nil
### GetConditions

`func (o *Evaluation) GetConditions() interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Evaluation) GetConditionsOk() (*interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Evaluation) SetConditions(v interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Evaluation) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *Evaluation) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *Evaluation) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetCreatedAt

`func (o *Evaluation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Evaluation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Evaluation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Evaluation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Evaluation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Evaluation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedBy

`func (o *Evaluation) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Evaluation) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Evaluation) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetDeleted

`func (o *Evaluation) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Evaluation) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Evaluation) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Evaluation) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


