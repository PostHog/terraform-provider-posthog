# PatchedRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**Members** | Pointer to **string** |  | [optional] [readonly] 
**IsDefault** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPatchedRole

`func NewPatchedRole() *PatchedRole`

NewPatchedRole instantiates a new PatchedRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRoleWithDefaults

`func NewPatchedRoleWithDefaults() *PatchedRole`

NewPatchedRoleWithDefaults instantiates a new PatchedRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedRole) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedRole) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedRole) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedRole) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedRole) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedRole) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedRole) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetMembers

`func (o *PatchedRole) GetMembers() string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PatchedRole) GetMembersOk() (*string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PatchedRole) SetMembers(v string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *PatchedRole) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetIsDefault

`func (o *PatchedRole) GetIsDefault() string`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PatchedRole) GetIsDefaultOk() (*string, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PatchedRole) SetIsDefault(v string)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PatchedRole) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


