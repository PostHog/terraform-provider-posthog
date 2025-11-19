# MinimalHedgehogConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessories** | **[]string** |  | 
**Color** | Pointer to [**HedgehogColorOptions**](HedgehogColorOptions.md) |  | [optional] 
**UseAsProfile** | **bool** |  | 

## Methods

### NewMinimalHedgehogConfig

`func NewMinimalHedgehogConfig(accessories []string, useAsProfile bool, ) *MinimalHedgehogConfig`

NewMinimalHedgehogConfig instantiates a new MinimalHedgehogConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalHedgehogConfigWithDefaults

`func NewMinimalHedgehogConfigWithDefaults() *MinimalHedgehogConfig`

NewMinimalHedgehogConfigWithDefaults instantiates a new MinimalHedgehogConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessories

`func (o *MinimalHedgehogConfig) GetAccessories() []string`

GetAccessories returns the Accessories field if non-nil, zero value otherwise.

### GetAccessoriesOk

`func (o *MinimalHedgehogConfig) GetAccessoriesOk() (*[]string, bool)`

GetAccessoriesOk returns a tuple with the Accessories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessories

`func (o *MinimalHedgehogConfig) SetAccessories(v []string)`

SetAccessories sets Accessories field to given value.


### GetColor

`func (o *MinimalHedgehogConfig) GetColor() HedgehogColorOptions`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MinimalHedgehogConfig) GetColorOk() (*HedgehogColorOptions, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MinimalHedgehogConfig) SetColor(v HedgehogColorOptions)`

SetColor sets Color field to given value.

### HasColor

`func (o *MinimalHedgehogConfig) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetUseAsProfile

`func (o *MinimalHedgehogConfig) GetUseAsProfile() bool`

GetUseAsProfile returns the UseAsProfile field if non-nil, zero value otherwise.

### GetUseAsProfileOk

`func (o *MinimalHedgehogConfig) GetUseAsProfileOk() (*bool, bool)`

GetUseAsProfileOk returns a tuple with the UseAsProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAsProfile

`func (o *MinimalHedgehogConfig) SetUseAsProfile(v bool)`

SetUseAsProfile sets UseAsProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


