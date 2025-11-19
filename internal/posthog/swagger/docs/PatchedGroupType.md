# PatchedGroupType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupType** | Pointer to **string** |  | [optional] [readonly] 
**GroupTypeIndex** | Pointer to **int32** |  | [optional] [readonly] 
**NameSingular** | Pointer to **NullableString** |  | [optional] 
**NamePlural** | Pointer to **NullableString** |  | [optional] 
**DetailDashboard** | Pointer to **NullableInt32** |  | [optional] 
**DefaultColumns** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewPatchedGroupType

`func NewPatchedGroupType() *PatchedGroupType`

NewPatchedGroupType instantiates a new PatchedGroupType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGroupTypeWithDefaults

`func NewPatchedGroupTypeWithDefaults() *PatchedGroupType`

NewPatchedGroupTypeWithDefaults instantiates a new PatchedGroupType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupType

`func (o *PatchedGroupType) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *PatchedGroupType) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *PatchedGroupType) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *PatchedGroupType) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetGroupTypeIndex

`func (o *PatchedGroupType) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *PatchedGroupType) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *PatchedGroupType) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.

### HasGroupTypeIndex

`func (o *PatchedGroupType) HasGroupTypeIndex() bool`

HasGroupTypeIndex returns a boolean if a field has been set.

### GetNameSingular

`func (o *PatchedGroupType) GetNameSingular() string`

GetNameSingular returns the NameSingular field if non-nil, zero value otherwise.

### GetNameSingularOk

`func (o *PatchedGroupType) GetNameSingularOk() (*string, bool)`

GetNameSingularOk returns a tuple with the NameSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSingular

`func (o *PatchedGroupType) SetNameSingular(v string)`

SetNameSingular sets NameSingular field to given value.

### HasNameSingular

`func (o *PatchedGroupType) HasNameSingular() bool`

HasNameSingular returns a boolean if a field has been set.

### SetNameSingularNil

`func (o *PatchedGroupType) SetNameSingularNil(b bool)`

 SetNameSingularNil sets the value for NameSingular to be an explicit nil

### UnsetNameSingular
`func (o *PatchedGroupType) UnsetNameSingular()`

UnsetNameSingular ensures that no value is present for NameSingular, not even an explicit nil
### GetNamePlural

`func (o *PatchedGroupType) GetNamePlural() string`

GetNamePlural returns the NamePlural field if non-nil, zero value otherwise.

### GetNamePluralOk

`func (o *PatchedGroupType) GetNamePluralOk() (*string, bool)`

GetNamePluralOk returns a tuple with the NamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePlural

`func (o *PatchedGroupType) SetNamePlural(v string)`

SetNamePlural sets NamePlural field to given value.

### HasNamePlural

`func (o *PatchedGroupType) HasNamePlural() bool`

HasNamePlural returns a boolean if a field has been set.

### SetNamePluralNil

`func (o *PatchedGroupType) SetNamePluralNil(b bool)`

 SetNamePluralNil sets the value for NamePlural to be an explicit nil

### UnsetNamePlural
`func (o *PatchedGroupType) UnsetNamePlural()`

UnsetNamePlural ensures that no value is present for NamePlural, not even an explicit nil
### GetDetailDashboard

`func (o *PatchedGroupType) GetDetailDashboard() int32`

GetDetailDashboard returns the DetailDashboard field if non-nil, zero value otherwise.

### GetDetailDashboardOk

`func (o *PatchedGroupType) GetDetailDashboardOk() (*int32, bool)`

GetDetailDashboardOk returns a tuple with the DetailDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailDashboard

`func (o *PatchedGroupType) SetDetailDashboard(v int32)`

SetDetailDashboard sets DetailDashboard field to given value.

### HasDetailDashboard

`func (o *PatchedGroupType) HasDetailDashboard() bool`

HasDetailDashboard returns a boolean if a field has been set.

### SetDetailDashboardNil

`func (o *PatchedGroupType) SetDetailDashboardNil(b bool)`

 SetDetailDashboardNil sets the value for DetailDashboard to be an explicit nil

### UnsetDetailDashboard
`func (o *PatchedGroupType) UnsetDetailDashboard()`

UnsetDetailDashboard ensures that no value is present for DetailDashboard, not even an explicit nil
### GetDefaultColumns

`func (o *PatchedGroupType) GetDefaultColumns() []string`

GetDefaultColumns returns the DefaultColumns field if non-nil, zero value otherwise.

### GetDefaultColumnsOk

`func (o *PatchedGroupType) GetDefaultColumnsOk() (*[]string, bool)`

GetDefaultColumnsOk returns a tuple with the DefaultColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColumns

`func (o *PatchedGroupType) SetDefaultColumns(v []string)`

SetDefaultColumns sets DefaultColumns field to given value.

### HasDefaultColumns

`func (o *PatchedGroupType) HasDefaultColumns() bool`

HasDefaultColumns returns a boolean if a field has been set.

### SetDefaultColumnsNil

`func (o *PatchedGroupType) SetDefaultColumnsNil(b bool)`

 SetDefaultColumnsNil sets the value for DefaultColumns to be an explicit nil

### UnsetDefaultColumns
`func (o *PatchedGroupType) UnsetDefaultColumns()`

UnsetDefaultColumns ensures that no value is present for DefaultColumns, not even an explicit nil
### GetCreatedAt

`func (o *PatchedGroupType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedGroupType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedGroupType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedGroupType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedGroupType) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedGroupType) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


