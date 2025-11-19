# ExperimentExposureCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExposureConfig** | Pointer to [**NullableExposureConfig**](ExposureConfig.md) |  | [optional] [default to null]
**FilterTestAccounts** | Pointer to **NullableBool** |  | [optional] 
**MultipleVariantHandling** | Pointer to [**MultipleVariantHandling**](MultipleVariantHandling.md) |  | [optional] 

## Methods

### NewExperimentExposureCriteria

`func NewExperimentExposureCriteria() *ExperimentExposureCriteria`

NewExperimentExposureCriteria instantiates a new ExperimentExposureCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentExposureCriteriaWithDefaults

`func NewExperimentExposureCriteriaWithDefaults() *ExperimentExposureCriteria`

NewExperimentExposureCriteriaWithDefaults instantiates a new ExperimentExposureCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExposureConfig

`func (o *ExperimentExposureCriteria) GetExposureConfig() ExposureConfig`

GetExposureConfig returns the ExposureConfig field if non-nil, zero value otherwise.

### GetExposureConfigOk

`func (o *ExperimentExposureCriteria) GetExposureConfigOk() (*ExposureConfig, bool)`

GetExposureConfigOk returns a tuple with the ExposureConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureConfig

`func (o *ExperimentExposureCriteria) SetExposureConfig(v ExposureConfig)`

SetExposureConfig sets ExposureConfig field to given value.

### HasExposureConfig

`func (o *ExperimentExposureCriteria) HasExposureConfig() bool`

HasExposureConfig returns a boolean if a field has been set.

### SetExposureConfigNil

`func (o *ExperimentExposureCriteria) SetExposureConfigNil(b bool)`

 SetExposureConfigNil sets the value for ExposureConfig to be an explicit nil

### UnsetExposureConfig
`func (o *ExperimentExposureCriteria) UnsetExposureConfig()`

UnsetExposureConfig ensures that no value is present for ExposureConfig, not even an explicit nil
### GetFilterTestAccounts

`func (o *ExperimentExposureCriteria) GetFilterTestAccounts() bool`

GetFilterTestAccounts returns the FilterTestAccounts field if non-nil, zero value otherwise.

### GetFilterTestAccountsOk

`func (o *ExperimentExposureCriteria) GetFilterTestAccountsOk() (*bool, bool)`

GetFilterTestAccountsOk returns a tuple with the FilterTestAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTestAccounts

`func (o *ExperimentExposureCriteria) SetFilterTestAccounts(v bool)`

SetFilterTestAccounts sets FilterTestAccounts field to given value.

### HasFilterTestAccounts

`func (o *ExperimentExposureCriteria) HasFilterTestAccounts() bool`

HasFilterTestAccounts returns a boolean if a field has been set.

### SetFilterTestAccountsNil

`func (o *ExperimentExposureCriteria) SetFilterTestAccountsNil(b bool)`

 SetFilterTestAccountsNil sets the value for FilterTestAccounts to be an explicit nil

### UnsetFilterTestAccounts
`func (o *ExperimentExposureCriteria) UnsetFilterTestAccounts()`

UnsetFilterTestAccounts ensures that no value is present for FilterTestAccounts, not even an explicit nil
### GetMultipleVariantHandling

`func (o *ExperimentExposureCriteria) GetMultipleVariantHandling() MultipleVariantHandling`

GetMultipleVariantHandling returns the MultipleVariantHandling field if non-nil, zero value otherwise.

### GetMultipleVariantHandlingOk

`func (o *ExperimentExposureCriteria) GetMultipleVariantHandlingOk() (*MultipleVariantHandling, bool)`

GetMultipleVariantHandlingOk returns a tuple with the MultipleVariantHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVariantHandling

`func (o *ExperimentExposureCriteria) SetMultipleVariantHandling(v MultipleVariantHandling)`

SetMultipleVariantHandling sets MultipleVariantHandling field to given value.

### HasMultipleVariantHandling

`func (o *ExperimentExposureCriteria) HasMultipleVariantHandling() bool`

HasMultipleVariantHandling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


