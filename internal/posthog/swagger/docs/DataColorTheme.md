# DataColorTheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Colors** | Pointer to **interface{}** |  | [optional] 
**IsGlobal** | **string** |  | [readonly] 
**CreatedAt** | **NullableTime** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 

## Methods

### NewDataColorTheme

`func NewDataColorTheme(id int32, name string, isGlobal string, createdAt NullableTime, createdBy UserBasic, ) *DataColorTheme`

NewDataColorTheme instantiates a new DataColorTheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataColorThemeWithDefaults

`func NewDataColorThemeWithDefaults() *DataColorTheme`

NewDataColorThemeWithDefaults instantiates a new DataColorTheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataColorTheme) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataColorTheme) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataColorTheme) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *DataColorTheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataColorTheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataColorTheme) SetName(v string)`

SetName sets Name field to given value.


### GetColors

`func (o *DataColorTheme) GetColors() interface{}`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *DataColorTheme) GetColorsOk() (*interface{}, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *DataColorTheme) SetColors(v interface{})`

SetColors sets Colors field to given value.

### HasColors

`func (o *DataColorTheme) HasColors() bool`

HasColors returns a boolean if a field has been set.

### SetColorsNil

`func (o *DataColorTheme) SetColorsNil(b bool)`

 SetColorsNil sets the value for Colors to be an explicit nil

### UnsetColors
`func (o *DataColorTheme) UnsetColors()`

UnsetColors ensures that no value is present for Colors, not even an explicit nil
### GetIsGlobal

`func (o *DataColorTheme) GetIsGlobal() string`

GetIsGlobal returns the IsGlobal field if non-nil, zero value otherwise.

### GetIsGlobalOk

`func (o *DataColorTheme) GetIsGlobalOk() (*string, bool)`

GetIsGlobalOk returns a tuple with the IsGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobal

`func (o *DataColorTheme) SetIsGlobal(v string)`

SetIsGlobal sets IsGlobal field to given value.


### GetCreatedAt

`func (o *DataColorTheme) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataColorTheme) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataColorTheme) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DataColorTheme) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DataColorTheme) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *DataColorTheme) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DataColorTheme) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DataColorTheme) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


