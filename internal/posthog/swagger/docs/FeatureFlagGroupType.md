# FeatureFlagGroupType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Properties** | Pointer to [**[]FixedpropertiesInner**](FixedpropertiesInner.md) |  | [optional] 
**RolloutPercentage** | Pointer to **NullableFloat32** |  | [optional] 
**SortKey** | Pointer to **NullableString** |  | [optional] 
**UsersAffected** | Pointer to **NullableFloat32** |  | [optional] 
**Variant** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFeatureFlagGroupType

`func NewFeatureFlagGroupType() *FeatureFlagGroupType`

NewFeatureFlagGroupType instantiates a new FeatureFlagGroupType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagGroupTypeWithDefaults

`func NewFeatureFlagGroupTypeWithDefaults() *FeatureFlagGroupType`

NewFeatureFlagGroupTypeWithDefaults instantiates a new FeatureFlagGroupType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FeatureFlagGroupType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureFlagGroupType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureFlagGroupType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FeatureFlagGroupType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FeatureFlagGroupType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FeatureFlagGroupType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProperties

`func (o *FeatureFlagGroupType) GetProperties() []FixedpropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FeatureFlagGroupType) GetPropertiesOk() (*[]FixedpropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FeatureFlagGroupType) SetProperties(v []FixedpropertiesInner)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FeatureFlagGroupType) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *FeatureFlagGroupType) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *FeatureFlagGroupType) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetRolloutPercentage

`func (o *FeatureFlagGroupType) GetRolloutPercentage() float32`

GetRolloutPercentage returns the RolloutPercentage field if non-nil, zero value otherwise.

### GetRolloutPercentageOk

`func (o *FeatureFlagGroupType) GetRolloutPercentageOk() (*float32, bool)`

GetRolloutPercentageOk returns a tuple with the RolloutPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentage

`func (o *FeatureFlagGroupType) SetRolloutPercentage(v float32)`

SetRolloutPercentage sets RolloutPercentage field to given value.

### HasRolloutPercentage

`func (o *FeatureFlagGroupType) HasRolloutPercentage() bool`

HasRolloutPercentage returns a boolean if a field has been set.

### SetRolloutPercentageNil

`func (o *FeatureFlagGroupType) SetRolloutPercentageNil(b bool)`

 SetRolloutPercentageNil sets the value for RolloutPercentage to be an explicit nil

### UnsetRolloutPercentage
`func (o *FeatureFlagGroupType) UnsetRolloutPercentage()`

UnsetRolloutPercentage ensures that no value is present for RolloutPercentage, not even an explicit nil
### GetSortKey

`func (o *FeatureFlagGroupType) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *FeatureFlagGroupType) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *FeatureFlagGroupType) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *FeatureFlagGroupType) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### SetSortKeyNil

`func (o *FeatureFlagGroupType) SetSortKeyNil(b bool)`

 SetSortKeyNil sets the value for SortKey to be an explicit nil

### UnsetSortKey
`func (o *FeatureFlagGroupType) UnsetSortKey()`

UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
### GetUsersAffected

`func (o *FeatureFlagGroupType) GetUsersAffected() float32`

GetUsersAffected returns the UsersAffected field if non-nil, zero value otherwise.

### GetUsersAffectedOk

`func (o *FeatureFlagGroupType) GetUsersAffectedOk() (*float32, bool)`

GetUsersAffectedOk returns a tuple with the UsersAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersAffected

`func (o *FeatureFlagGroupType) SetUsersAffected(v float32)`

SetUsersAffected sets UsersAffected field to given value.

### HasUsersAffected

`func (o *FeatureFlagGroupType) HasUsersAffected() bool`

HasUsersAffected returns a boolean if a field has been set.

### SetUsersAffectedNil

`func (o *FeatureFlagGroupType) SetUsersAffectedNil(b bool)`

 SetUsersAffectedNil sets the value for UsersAffected to be an explicit nil

### UnsetUsersAffected
`func (o *FeatureFlagGroupType) UnsetUsersAffected()`

UnsetUsersAffected ensures that no value is present for UsersAffected, not even an explicit nil
### GetVariant

`func (o *FeatureFlagGroupType) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *FeatureFlagGroupType) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *FeatureFlagGroupType) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *FeatureFlagGroupType) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### SetVariantNil

`func (o *FeatureFlagGroupType) SetVariantNil(b bool)`

 SetVariantNil sets the value for Variant to be an explicit nil

### UnsetVariant
`func (o *FeatureFlagGroupType) UnsetVariant()`

UnsetVariant ensures that no value is present for Variant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


