# ExperimentHoldoutType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**UserBasicType**](UserBasicType.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Filters** | [**[]FeatureFlagGroupType**](FeatureFlagGroupType.md) |  | 
**Id** | Pointer to **NullableFloat32** |  | [optional] 
**Name** | **string** |  | 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExperimentHoldoutType

`func NewExperimentHoldoutType(filters []FeatureFlagGroupType, name string, ) *ExperimentHoldoutType`

NewExperimentHoldoutType instantiates a new ExperimentHoldoutType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentHoldoutTypeWithDefaults

`func NewExperimentHoldoutTypeWithDefaults() *ExperimentHoldoutType`

NewExperimentHoldoutTypeWithDefaults instantiates a new ExperimentHoldoutType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ExperimentHoldoutType) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExperimentHoldoutType) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExperimentHoldoutType) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExperimentHoldoutType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ExperimentHoldoutType) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ExperimentHoldoutType) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *ExperimentHoldoutType) GetCreatedBy() UserBasicType`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ExperimentHoldoutType) GetCreatedByOk() (*UserBasicType, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ExperimentHoldoutType) SetCreatedBy(v UserBasicType)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ExperimentHoldoutType) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *ExperimentHoldoutType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExperimentHoldoutType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExperimentHoldoutType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExperimentHoldoutType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExperimentHoldoutType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExperimentHoldoutType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFilters

`func (o *ExperimentHoldoutType) GetFilters() []FeatureFlagGroupType`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ExperimentHoldoutType) GetFiltersOk() (*[]FeatureFlagGroupType, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ExperimentHoldoutType) SetFilters(v []FeatureFlagGroupType)`

SetFilters sets Filters field to given value.


### GetId

`func (o *ExperimentHoldoutType) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentHoldoutType) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentHoldoutType) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *ExperimentHoldoutType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExperimentHoldoutType) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExperimentHoldoutType) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ExperimentHoldoutType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentHoldoutType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentHoldoutType) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *ExperimentHoldoutType) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExperimentHoldoutType) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExperimentHoldoutType) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExperimentHoldoutType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ExperimentHoldoutType) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ExperimentHoldoutType) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


