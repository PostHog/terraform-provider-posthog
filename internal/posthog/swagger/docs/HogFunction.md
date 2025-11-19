# HogFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**CreatedBy** | [**UserBasic**](UserBasic.md) |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Hog** | Pointer to **string** |  | [optional] 
**Bytecode** | **interface{}** |  | [readonly] 
**Transpiled** | **NullableString** |  | [readonly] 
**InputsSchema** | Pointer to [**[]InputsSchemaItem**](InputsSchemaItem.md) |  | [optional] 
**Inputs** | Pointer to [**map[string]InputsItem**](InputsItem.md) |  | [optional] 
**Filters** | Pointer to [**HogFunctionFilters**](HogFunctionFilters.md) |  | [optional] 
**Masking** | Pointer to [**NullableHogFunctionMasking**](HogFunctionMasking.md) |  | [optional] 
**Mappings** | Pointer to [**[]Mappings**](Mappings.md) |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**Template** | [**HogFunctionTemplate**](HogFunctionTemplate.md) |  | [readonly] 
**TemplateId** | Pointer to **NullableString** |  | [optional] 
**Status** | [**NullableHogFunctionStatus**](HogFunctionStatus.md) |  | [readonly] 
**ExecutionOrder** | Pointer to **NullableInt32** |  | [optional] 
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewHogFunction

`func NewHogFunction(id string, createdAt time.Time, createdBy UserBasic, updatedAt time.Time, bytecode interface{}, transpiled NullableString, template HogFunctionTemplate, status NullableHogFunctionStatus, ) *HogFunction`

NewHogFunction instantiates a new HogFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogFunctionWithDefaults

`func NewHogFunctionWithDefaults() *HogFunction`

NewHogFunctionWithDefaults instantiates a new HogFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HogFunction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HogFunction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HogFunction) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *HogFunction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HogFunction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HogFunction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HogFunction) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HogFunction) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HogFunction) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *HogFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HogFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HogFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HogFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HogFunction) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HogFunction) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *HogFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HogFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HogFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HogFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HogFunction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HogFunction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HogFunction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *HogFunction) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *HogFunction) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *HogFunction) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedAt

`func (o *HogFunction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HogFunction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HogFunction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEnabled

`func (o *HogFunction) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HogFunction) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HogFunction) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *HogFunction) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDeleted

`func (o *HogFunction) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *HogFunction) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *HogFunction) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *HogFunction) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetHog

`func (o *HogFunction) GetHog() string`

GetHog returns the Hog field if non-nil, zero value otherwise.

### GetHogOk

`func (o *HogFunction) GetHogOk() (*string, bool)`

GetHogOk returns a tuple with the Hog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHog

`func (o *HogFunction) SetHog(v string)`

SetHog sets Hog field to given value.

### HasHog

`func (o *HogFunction) HasHog() bool`

HasHog returns a boolean if a field has been set.

### GetBytecode

`func (o *HogFunction) GetBytecode() interface{}`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *HogFunction) GetBytecodeOk() (*interface{}, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *HogFunction) SetBytecode(v interface{})`

SetBytecode sets Bytecode field to given value.


### SetBytecodeNil

`func (o *HogFunction) SetBytecodeNil(b bool)`

 SetBytecodeNil sets the value for Bytecode to be an explicit nil

### UnsetBytecode
`func (o *HogFunction) UnsetBytecode()`

UnsetBytecode ensures that no value is present for Bytecode, not even an explicit nil
### GetTranspiled

`func (o *HogFunction) GetTranspiled() string`

GetTranspiled returns the Transpiled field if non-nil, zero value otherwise.

### GetTranspiledOk

`func (o *HogFunction) GetTranspiledOk() (*string, bool)`

GetTranspiledOk returns a tuple with the Transpiled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranspiled

`func (o *HogFunction) SetTranspiled(v string)`

SetTranspiled sets Transpiled field to given value.


### SetTranspiledNil

`func (o *HogFunction) SetTranspiledNil(b bool)`

 SetTranspiledNil sets the value for Transpiled to be an explicit nil

### UnsetTranspiled
`func (o *HogFunction) UnsetTranspiled()`

UnsetTranspiled ensures that no value is present for Transpiled, not even an explicit nil
### GetInputsSchema

`func (o *HogFunction) GetInputsSchema() []InputsSchemaItem`

GetInputsSchema returns the InputsSchema field if non-nil, zero value otherwise.

### GetInputsSchemaOk

`func (o *HogFunction) GetInputsSchemaOk() (*[]InputsSchemaItem, bool)`

GetInputsSchemaOk returns a tuple with the InputsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsSchema

`func (o *HogFunction) SetInputsSchema(v []InputsSchemaItem)`

SetInputsSchema sets InputsSchema field to given value.

### HasInputsSchema

`func (o *HogFunction) HasInputsSchema() bool`

HasInputsSchema returns a boolean if a field has been set.

### GetInputs

`func (o *HogFunction) GetInputs() map[string]InputsItem`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *HogFunction) GetInputsOk() (*map[string]InputsItem, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *HogFunction) SetInputs(v map[string]InputsItem)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *HogFunction) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetFilters

`func (o *HogFunction) GetFilters() HogFunctionFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HogFunction) GetFiltersOk() (*HogFunctionFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HogFunction) SetFilters(v HogFunctionFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *HogFunction) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMasking

`func (o *HogFunction) GetMasking() HogFunctionMasking`

GetMasking returns the Masking field if non-nil, zero value otherwise.

### GetMaskingOk

`func (o *HogFunction) GetMaskingOk() (*HogFunctionMasking, bool)`

GetMaskingOk returns a tuple with the Masking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasking

`func (o *HogFunction) SetMasking(v HogFunctionMasking)`

SetMasking sets Masking field to given value.

### HasMasking

`func (o *HogFunction) HasMasking() bool`

HasMasking returns a boolean if a field has been set.

### SetMaskingNil

`func (o *HogFunction) SetMaskingNil(b bool)`

 SetMaskingNil sets the value for Masking to be an explicit nil

### UnsetMasking
`func (o *HogFunction) UnsetMasking()`

UnsetMasking ensures that no value is present for Masking, not even an explicit nil
### GetMappings

`func (o *HogFunction) GetMappings() []Mappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *HogFunction) GetMappingsOk() (*[]Mappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *HogFunction) SetMappings(v []Mappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *HogFunction) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *HogFunction) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *HogFunction) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil
### GetIconUrl

`func (o *HogFunction) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *HogFunction) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *HogFunction) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *HogFunction) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *HogFunction) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *HogFunction) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetTemplate

`func (o *HogFunction) GetTemplate() HogFunctionTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *HogFunction) GetTemplateOk() (*HogFunctionTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *HogFunction) SetTemplate(v HogFunctionTemplate)`

SetTemplate sets Template field to given value.


### GetTemplateId

`func (o *HogFunction) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *HogFunction) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *HogFunction) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *HogFunction) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *HogFunction) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *HogFunction) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetStatus

`func (o *HogFunction) GetStatus() HogFunctionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HogFunction) GetStatusOk() (*HogFunctionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HogFunction) SetStatus(v HogFunctionStatus)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *HogFunction) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *HogFunction) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetExecutionOrder

`func (o *HogFunction) GetExecutionOrder() int32`

GetExecutionOrder returns the ExecutionOrder field if non-nil, zero value otherwise.

### GetExecutionOrderOk

`func (o *HogFunction) GetExecutionOrderOk() (*int32, bool)`

GetExecutionOrderOk returns a tuple with the ExecutionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrder

`func (o *HogFunction) SetExecutionOrder(v int32)`

SetExecutionOrder sets ExecutionOrder field to given value.

### HasExecutionOrder

`func (o *HogFunction) HasExecutionOrder() bool`

HasExecutionOrder returns a boolean if a field has been set.

### SetExecutionOrderNil

`func (o *HogFunction) SetExecutionOrderNil(b bool)`

 SetExecutionOrderNil sets the value for ExecutionOrder to be an explicit nil

### UnsetExecutionOrder
`func (o *HogFunction) UnsetExecutionOrder()`

UnsetExecutionOrder ensures that no value is present for ExecutionOrder, not even an explicit nil
### GetCreateInFolder

`func (o *HogFunction) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *HogFunction) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *HogFunction) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *HogFunction) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


