# PatchedHogFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to [**UserBasic**](UserBasic.md) |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Hog** | Pointer to **string** |  | [optional] 
**Bytecode** | Pointer to **interface{}** |  | [optional] [readonly] 
**Transpiled** | Pointer to **NullableString** |  | [optional] [readonly] 
**InputsSchema** | Pointer to [**[]InputsSchemaItem**](InputsSchemaItem.md) |  | [optional] 
**Inputs** | Pointer to [**map[string]InputsItem**](InputsItem.md) |  | [optional] 
**Filters** | Pointer to [**HogFunctionFilters**](HogFunctionFilters.md) |  | [optional] 
**Masking** | Pointer to [**NullableHogFunctionMasking**](HogFunctionMasking.md) |  | [optional] 
**Mappings** | Pointer to [**[]Mappings**](Mappings.md) |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to [**HogFunctionTemplate**](HogFunctionTemplate.md) |  | [optional] [readonly] 
**TemplateId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableHogFunctionStatus**](HogFunctionStatus.md) |  | [optional] [readonly] 
**ExecutionOrder** | Pointer to **NullableInt32** |  | [optional] 
**CreateInFolder** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedHogFunction

`func NewPatchedHogFunction() *PatchedHogFunction`

NewPatchedHogFunction instantiates a new PatchedHogFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedHogFunctionWithDefaults

`func NewPatchedHogFunctionWithDefaults() *PatchedHogFunction`

NewPatchedHogFunctionWithDefaults instantiates a new PatchedHogFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedHogFunction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedHogFunction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedHogFunction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedHogFunction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PatchedHogFunction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedHogFunction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedHogFunction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedHogFunction) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PatchedHogFunction) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PatchedHogFunction) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetName

`func (o *PatchedHogFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedHogFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedHogFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedHogFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedHogFunction) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedHogFunction) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PatchedHogFunction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedHogFunction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedHogFunction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedHogFunction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchedHogFunction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchedHogFunction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchedHogFunction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchedHogFunction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PatchedHogFunction) GetCreatedBy() UserBasic`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PatchedHogFunction) GetCreatedByOk() (*UserBasic, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PatchedHogFunction) SetCreatedBy(v UserBasic)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PatchedHogFunction) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PatchedHogFunction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PatchedHogFunction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PatchedHogFunction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PatchedHogFunction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedHogFunction) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedHogFunction) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedHogFunction) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedHogFunction) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDeleted

`func (o *PatchedHogFunction) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *PatchedHogFunction) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *PatchedHogFunction) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *PatchedHogFunction) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetHog

`func (o *PatchedHogFunction) GetHog() string`

GetHog returns the Hog field if non-nil, zero value otherwise.

### GetHogOk

`func (o *PatchedHogFunction) GetHogOk() (*string, bool)`

GetHogOk returns a tuple with the Hog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHog

`func (o *PatchedHogFunction) SetHog(v string)`

SetHog sets Hog field to given value.

### HasHog

`func (o *PatchedHogFunction) HasHog() bool`

HasHog returns a boolean if a field has been set.

### GetBytecode

`func (o *PatchedHogFunction) GetBytecode() interface{}`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *PatchedHogFunction) GetBytecodeOk() (*interface{}, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *PatchedHogFunction) SetBytecode(v interface{})`

SetBytecode sets Bytecode field to given value.

### HasBytecode

`func (o *PatchedHogFunction) HasBytecode() bool`

HasBytecode returns a boolean if a field has been set.

### SetBytecodeNil

`func (o *PatchedHogFunction) SetBytecodeNil(b bool)`

 SetBytecodeNil sets the value for Bytecode to be an explicit nil

### UnsetBytecode
`func (o *PatchedHogFunction) UnsetBytecode()`

UnsetBytecode ensures that no value is present for Bytecode, not even an explicit nil
### GetTranspiled

`func (o *PatchedHogFunction) GetTranspiled() string`

GetTranspiled returns the Transpiled field if non-nil, zero value otherwise.

### GetTranspiledOk

`func (o *PatchedHogFunction) GetTranspiledOk() (*string, bool)`

GetTranspiledOk returns a tuple with the Transpiled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranspiled

`func (o *PatchedHogFunction) SetTranspiled(v string)`

SetTranspiled sets Transpiled field to given value.

### HasTranspiled

`func (o *PatchedHogFunction) HasTranspiled() bool`

HasTranspiled returns a boolean if a field has been set.

### SetTranspiledNil

`func (o *PatchedHogFunction) SetTranspiledNil(b bool)`

 SetTranspiledNil sets the value for Transpiled to be an explicit nil

### UnsetTranspiled
`func (o *PatchedHogFunction) UnsetTranspiled()`

UnsetTranspiled ensures that no value is present for Transpiled, not even an explicit nil
### GetInputsSchema

`func (o *PatchedHogFunction) GetInputsSchema() []InputsSchemaItem`

GetInputsSchema returns the InputsSchema field if non-nil, zero value otherwise.

### GetInputsSchemaOk

`func (o *PatchedHogFunction) GetInputsSchemaOk() (*[]InputsSchemaItem, bool)`

GetInputsSchemaOk returns a tuple with the InputsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsSchema

`func (o *PatchedHogFunction) SetInputsSchema(v []InputsSchemaItem)`

SetInputsSchema sets InputsSchema field to given value.

### HasInputsSchema

`func (o *PatchedHogFunction) HasInputsSchema() bool`

HasInputsSchema returns a boolean if a field has been set.

### GetInputs

`func (o *PatchedHogFunction) GetInputs() map[string]InputsItem`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *PatchedHogFunction) GetInputsOk() (*map[string]InputsItem, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *PatchedHogFunction) SetInputs(v map[string]InputsItem)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *PatchedHogFunction) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetFilters

`func (o *PatchedHogFunction) GetFilters() HogFunctionFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PatchedHogFunction) GetFiltersOk() (*HogFunctionFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PatchedHogFunction) SetFilters(v HogFunctionFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PatchedHogFunction) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMasking

`func (o *PatchedHogFunction) GetMasking() HogFunctionMasking`

GetMasking returns the Masking field if non-nil, zero value otherwise.

### GetMaskingOk

`func (o *PatchedHogFunction) GetMaskingOk() (*HogFunctionMasking, bool)`

GetMaskingOk returns a tuple with the Masking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasking

`func (o *PatchedHogFunction) SetMasking(v HogFunctionMasking)`

SetMasking sets Masking field to given value.

### HasMasking

`func (o *PatchedHogFunction) HasMasking() bool`

HasMasking returns a boolean if a field has been set.

### SetMaskingNil

`func (o *PatchedHogFunction) SetMaskingNil(b bool)`

 SetMaskingNil sets the value for Masking to be an explicit nil

### UnsetMasking
`func (o *PatchedHogFunction) UnsetMasking()`

UnsetMasking ensures that no value is present for Masking, not even an explicit nil
### GetMappings

`func (o *PatchedHogFunction) GetMappings() []Mappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *PatchedHogFunction) GetMappingsOk() (*[]Mappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *PatchedHogFunction) SetMappings(v []Mappings)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *PatchedHogFunction) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### SetMappingsNil

`func (o *PatchedHogFunction) SetMappingsNil(b bool)`

 SetMappingsNil sets the value for Mappings to be an explicit nil

### UnsetMappings
`func (o *PatchedHogFunction) UnsetMappings()`

UnsetMappings ensures that no value is present for Mappings, not even an explicit nil
### GetIconUrl

`func (o *PatchedHogFunction) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *PatchedHogFunction) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *PatchedHogFunction) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *PatchedHogFunction) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *PatchedHogFunction) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *PatchedHogFunction) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetTemplate

`func (o *PatchedHogFunction) GetTemplate() HogFunctionTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PatchedHogFunction) GetTemplateOk() (*HogFunctionTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PatchedHogFunction) SetTemplate(v HogFunctionTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PatchedHogFunction) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTemplateId

`func (o *PatchedHogFunction) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PatchedHogFunction) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PatchedHogFunction) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *PatchedHogFunction) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *PatchedHogFunction) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *PatchedHogFunction) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetStatus

`func (o *PatchedHogFunction) GetStatus() HogFunctionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedHogFunction) GetStatusOk() (*HogFunctionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedHogFunction) SetStatus(v HogFunctionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedHogFunction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PatchedHogFunction) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PatchedHogFunction) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetExecutionOrder

`func (o *PatchedHogFunction) GetExecutionOrder() int32`

GetExecutionOrder returns the ExecutionOrder field if non-nil, zero value otherwise.

### GetExecutionOrderOk

`func (o *PatchedHogFunction) GetExecutionOrderOk() (*int32, bool)`

GetExecutionOrderOk returns a tuple with the ExecutionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionOrder

`func (o *PatchedHogFunction) SetExecutionOrder(v int32)`

SetExecutionOrder sets ExecutionOrder field to given value.

### HasExecutionOrder

`func (o *PatchedHogFunction) HasExecutionOrder() bool`

HasExecutionOrder returns a boolean if a field has been set.

### SetExecutionOrderNil

`func (o *PatchedHogFunction) SetExecutionOrderNil(b bool)`

 SetExecutionOrderNil sets the value for ExecutionOrder to be an explicit nil

### UnsetExecutionOrder
`func (o *PatchedHogFunction) UnsetExecutionOrder()`

UnsetExecutionOrder ensures that no value is present for ExecutionOrder, not even an explicit nil
### GetCreateInFolder

`func (o *PatchedHogFunction) GetCreateInFolder() string`

GetCreateInFolder returns the CreateInFolder field if non-nil, zero value otherwise.

### GetCreateInFolderOk

`func (o *PatchedHogFunction) GetCreateInFolderOk() (*string, bool)`

GetCreateInFolderOk returns a tuple with the CreateInFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateInFolder

`func (o *PatchedHogFunction) SetCreateInFolder(v string)`

SetCreateInFolder sets CreateInFolder field to given value.

### HasCreateInFolder

`func (o *PatchedHogFunction) HasCreateInFolder() bool`

HasCreateInFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


