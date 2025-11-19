# Mappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**InputsSchema** | Pointer to [**[]InputsSchemaItem**](InputsSchemaItem.md) |  | [optional] 
**Inputs** | Pointer to [**map[string]InputsItem**](InputsItem.md) |  | [optional] 
**Filters** | Pointer to [**HogFunctionFilters**](HogFunctionFilters.md) |  | [optional] 

## Methods

### NewMappings

`func NewMappings() *Mappings`

NewMappings instantiates a new Mappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingsWithDefaults

`func NewMappingsWithDefaults() *Mappings`

NewMappingsWithDefaults instantiates a new Mappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Mappings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mappings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mappings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Mappings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInputsSchema

`func (o *Mappings) GetInputsSchema() []InputsSchemaItem`

GetInputsSchema returns the InputsSchema field if non-nil, zero value otherwise.

### GetInputsSchemaOk

`func (o *Mappings) GetInputsSchemaOk() (*[]InputsSchemaItem, bool)`

GetInputsSchemaOk returns a tuple with the InputsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsSchema

`func (o *Mappings) SetInputsSchema(v []InputsSchemaItem)`

SetInputsSchema sets InputsSchema field to given value.

### HasInputsSchema

`func (o *Mappings) HasInputsSchema() bool`

HasInputsSchema returns a boolean if a field has been set.

### GetInputs

`func (o *Mappings) GetInputs() map[string]InputsItem`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *Mappings) GetInputsOk() (*map[string]InputsItem, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *Mappings) SetInputs(v map[string]InputsItem)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *Mappings) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetFilters

`func (o *Mappings) GetFilters() HogFunctionFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Mappings) GetFiltersOk() (*HogFunctionFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Mappings) SetFilters(v HogFunctionFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Mappings) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


