# GroupType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupType** | **string** |  | [readonly] 
**GroupTypeIndex** | **int32** |  | [readonly] 
**NameSingular** | Pointer to **NullableString** |  | [optional] 
**NamePlural** | Pointer to **NullableString** |  | [optional] 
**DetailDashboard** | Pointer to **NullableInt32** |  | [optional] 
**DefaultColumns** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGroupType

`func NewGroupType(groupType string, groupTypeIndex int32, ) *GroupType`

NewGroupType instantiates a new GroupType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupTypeWithDefaults

`func NewGroupTypeWithDefaults() *GroupType`

NewGroupTypeWithDefaults instantiates a new GroupType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupType

`func (o *GroupType) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *GroupType) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *GroupType) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.


### GetGroupTypeIndex

`func (o *GroupType) GetGroupTypeIndex() int32`

GetGroupTypeIndex returns the GroupTypeIndex field if non-nil, zero value otherwise.

### GetGroupTypeIndexOk

`func (o *GroupType) GetGroupTypeIndexOk() (*int32, bool)`

GetGroupTypeIndexOk returns a tuple with the GroupTypeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTypeIndex

`func (o *GroupType) SetGroupTypeIndex(v int32)`

SetGroupTypeIndex sets GroupTypeIndex field to given value.


### GetNameSingular

`func (o *GroupType) GetNameSingular() string`

GetNameSingular returns the NameSingular field if non-nil, zero value otherwise.

### GetNameSingularOk

`func (o *GroupType) GetNameSingularOk() (*string, bool)`

GetNameSingularOk returns a tuple with the NameSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSingular

`func (o *GroupType) SetNameSingular(v string)`

SetNameSingular sets NameSingular field to given value.

### HasNameSingular

`func (o *GroupType) HasNameSingular() bool`

HasNameSingular returns a boolean if a field has been set.

### SetNameSingularNil

`func (o *GroupType) SetNameSingularNil(b bool)`

 SetNameSingularNil sets the value for NameSingular to be an explicit nil

### UnsetNameSingular
`func (o *GroupType) UnsetNameSingular()`

UnsetNameSingular ensures that no value is present for NameSingular, not even an explicit nil
### GetNamePlural

`func (o *GroupType) GetNamePlural() string`

GetNamePlural returns the NamePlural field if non-nil, zero value otherwise.

### GetNamePluralOk

`func (o *GroupType) GetNamePluralOk() (*string, bool)`

GetNamePluralOk returns a tuple with the NamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePlural

`func (o *GroupType) SetNamePlural(v string)`

SetNamePlural sets NamePlural field to given value.

### HasNamePlural

`func (o *GroupType) HasNamePlural() bool`

HasNamePlural returns a boolean if a field has been set.

### SetNamePluralNil

`func (o *GroupType) SetNamePluralNil(b bool)`

 SetNamePluralNil sets the value for NamePlural to be an explicit nil

### UnsetNamePlural
`func (o *GroupType) UnsetNamePlural()`

UnsetNamePlural ensures that no value is present for NamePlural, not even an explicit nil
### GetDetailDashboard

`func (o *GroupType) GetDetailDashboard() int32`

GetDetailDashboard returns the DetailDashboard field if non-nil, zero value otherwise.

### GetDetailDashboardOk

`func (o *GroupType) GetDetailDashboardOk() (*int32, bool)`

GetDetailDashboardOk returns a tuple with the DetailDashboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailDashboard

`func (o *GroupType) SetDetailDashboard(v int32)`

SetDetailDashboard sets DetailDashboard field to given value.

### HasDetailDashboard

`func (o *GroupType) HasDetailDashboard() bool`

HasDetailDashboard returns a boolean if a field has been set.

### SetDetailDashboardNil

`func (o *GroupType) SetDetailDashboardNil(b bool)`

 SetDetailDashboardNil sets the value for DetailDashboard to be an explicit nil

### UnsetDetailDashboard
`func (o *GroupType) UnsetDetailDashboard()`

UnsetDetailDashboard ensures that no value is present for DetailDashboard, not even an explicit nil
### GetDefaultColumns

`func (o *GroupType) GetDefaultColumns() []string`

GetDefaultColumns returns the DefaultColumns field if non-nil, zero value otherwise.

### GetDefaultColumnsOk

`func (o *GroupType) GetDefaultColumnsOk() (*[]string, bool)`

GetDefaultColumnsOk returns a tuple with the DefaultColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColumns

`func (o *GroupType) SetDefaultColumns(v []string)`

SetDefaultColumns sets DefaultColumns field to given value.

### HasDefaultColumns

`func (o *GroupType) HasDefaultColumns() bool`

HasDefaultColumns returns a boolean if a field has been set.

### SetDefaultColumnsNil

`func (o *GroupType) SetDefaultColumnsNil(b bool)`

 SetDefaultColumnsNil sets the value for DefaultColumns to be an explicit nil

### UnsetDefaultColumns
`func (o *GroupType) UnsetDefaultColumns()`

UnsetDefaultColumns ensures that no value is present for DefaultColumns, not even an explicit nil
### GetCreatedAt

`func (o *GroupType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GroupType) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GroupType) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


