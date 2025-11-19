# PatchedDataColorTheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Colors** | Pointer to **interface{}** |  | [optional] 
**IsGlobal** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 

## Methods

### NewPatchedDataColorTheme

`func NewPatchedDataColorTheme() *PatchedDataColorTheme`

NewPatchedDataColorTheme instantiates a new PatchedDataColorTheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDataColorThemeWithDefaults

`func NewPatchedDataColorThemeWithDefaults() *PatchedDataColorTheme`

NewPatchedDataColorThemeWithDefaults instantiates a new PatchedDataColorTheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedDataColorTheme) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedDataColorTheme) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedDataColorTheme) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedDataColorTheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedDataColorTheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDataColorTheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDataColorTheme) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDataColorTheme) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColors

`func (o *PatchedDataColorTheme) GetColors() interface{}`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *PatchedDataColorTheme) GetColorsOk() (*interface{}, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *PatchedDataColorTheme) SetColors(v interface{})`

SetColors sets Colors field to given value.

### HasColors

`func (o *PatchedDataColorTheme) HasColors() bool`

HasColors returns a boolean if a field has been set.

### SetColorsNil

`func (o *PatchedDataColorTheme) SetColorsNil(b bool)`

 SetColorsNil sets the value for Colors to be an explicit nil

### UnsetColors
`func (o *PatchedDataColorTheme) UnsetColors()`

UnsetColors ensures that no value is present for Colors, not even an explicit nil
### GetIsGlobal

`func (o *PatchedDataColorTheme) GetIsGlobal() string`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *PatchedDataColorTheme) GetIsGlobalOk() (*string, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *PatchedDataColorTheme) SetIsGlobal(v string)`

SetIsGlobal sets IsGlobal field to given value.

### HasIsGlobal

`func (o *PatchedDataColorTheme) HasIsGlobal() bool`

HasIsGlobal returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedDataColorTheme) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedDataColorTheme) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedDataColorTheme) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedDataColorTheme) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PatchedDataColorTheme) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PatchedDataColorTheme) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *PatchedDataColorTheme) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedDataColorTheme) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedDataColorTheme) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedDataColorTheme) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


