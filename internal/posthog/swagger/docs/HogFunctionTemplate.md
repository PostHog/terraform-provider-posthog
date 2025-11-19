# HogFunctionTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Code** | **string** |  | 
**CodeLanguage** | Pointer to **string** |  | [optional] 
**InputsSchema** | **interface{}** |  | 
**Type** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **interface{}** |  | [optional] 
**Free** | Pointer to **bool** |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to **interface{}** |  | [optional] 
**Masking** | Pointer to **interface{}** |  | [optional] 
**MappingTemplates** | Pointer to [**[]HogFunctionMappingTemplate**](HogFunctionMappingTemplate.md) |  | [optional] 

## Methods

### NewHogFunctionTemplate

`func NewHogFunctionTemplate(id string, name string, code string, inputsSchema interface{}, type_ string, ) *HogFunctionTemplate`

NewHogFunctionTemplate instantiates a new HogFunctionTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHogFunctionTemplateWithDefaults

`func NewHogFunctionTemplateWithDefaults() *HogFunctionTemplate`

NewHogFunctionTemplateWithDefaults instantiates a new HogFunctionTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HogFunctionTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HogFunctionTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HogFunctionTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HogFunctionTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HogFunctionTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HogFunctionTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *HogFunctionTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HogFunctionTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HogFunctionTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HogFunctionTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HogFunctionTemplate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HogFunctionTemplate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *HogFunctionTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *HogFunctionTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *HogFunctionTemplate) SetCode(v string)`

SetCode sets Code field to given value.


### GetCodeLanguage

`func (o *HogFunctionTemplate) GetCodeLanguage() string`

GetCodeLanguage returns the CodeLanguage field if non-nil, zero value otherwise.

### GetCodeLanguageOk

`func (o *HogFunctionTemplate) GetCodeLanguageOk() (*string, bool)`

GetCodeLanguageOk returns a tuple with the CodeLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLanguage

`func (o *HogFunctionTemplate) SetCodeLanguage(v string)`

SetCodeLanguage sets CodeLanguage field to given value.

### HasCodeLanguage

`func (o *HogFunctionTemplate) HasCodeLanguage() bool`

HasCodeLanguage returns a boolean if a field has been set.

### GetInputsSchema

`func (o *HogFunctionTemplate) GetInputsSchema() interface{}`

GetInputsSchema returns the InputsSchema field if non-nil, zero value otherwise.

### GetInputsSchemaOk

`func (o *HogFunctionTemplate) GetInputsSchemaOk() (*interface{}, bool)`

GetInputsSchemaOk returns a tuple with the InputsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputsSchema

`func (o *HogFunctionTemplate) SetInputsSchema(v interface{})`

SetInputsSchema sets InputsSchema field to given value.


### SetInputsSchemaNil

`func (o *HogFunctionTemplate) SetInputsSchemaNil(b bool)`

 SetInputsSchemaNil sets the value for InputsSchema to be an explicit nil

### UnsetInputsSchema
`func (o *HogFunctionTemplate) UnsetInputsSchema()`

UnsetInputsSchema ensures that no value is present for InputsSchema, not even an explicit nil
### GetType

`func (o *HogFunctionTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HogFunctionTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HogFunctionTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *HogFunctionTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HogFunctionTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HogFunctionTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HogFunctionTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCategory

`func (o *HogFunctionTemplate) GetCategory() interface{}`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HogFunctionTemplate) GetCategoryOk() (*interface{}, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HogFunctionTemplate) SetCategory(v interface{})`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HogFunctionTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *HogFunctionTemplate) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *HogFunctionTemplate) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetFree

`func (o *HogFunctionTemplate) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *HogFunctionTemplate) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *HogFunctionTemplate) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *HogFunctionTemplate) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetIconUrl

`func (o *HogFunctionTemplate) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *HogFunctionTemplate) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *HogFunctionTemplate) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *HogFunctionTemplate) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *HogFunctionTemplate) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *HogFunctionTemplate) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetFilters

`func (o *HogFunctionTemplate) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *HogFunctionTemplate) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *HogFunctionTemplate) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *HogFunctionTemplate) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *HogFunctionTemplate) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *HogFunctionTemplate) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetMasking

`func (o *HogFunctionTemplate) GetMasking() interface{}`

GetMasking returns the Masking field if non-nil, zero value otherwise.

### GetMaskingOk

`func (o *HogFunctionTemplate) GetMaskingOk() (*interface{}, bool)`

GetMaskingOk returns a tuple with the Masking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasking

`func (o *HogFunctionTemplate) SetMasking(v interface{})`

SetMasking sets Masking field to given value.

### HasMasking

`func (o *HogFunctionTemplate) HasMasking() bool`

HasMasking returns a boolean if a field has been set.

### SetMaskingNil

`func (o *HogFunctionTemplate) SetMaskingNil(b bool)`

 SetMaskingNil sets the value for Masking to be an explicit nil

### UnsetMasking
`func (o *HogFunctionTemplate) UnsetMasking()`

UnsetMasking ensures that no value is present for Masking, not even an explicit nil
### GetMappingTemplates

`func (o *HogFunctionTemplate) GetMappingTemplates() []HogFunctionMappingTemplate`

GetMappingTemplates returns the MappingTemplates field if non-nil, zero value otherwise.

### GetMappingTemplatesOk

`func (o *HogFunctionTemplate) GetMappingTemplatesOk() (*[]HogFunctionMappingTemplate, bool)`

GetMappingTemplatesOk returns a tuple with the MappingTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingTemplates

`func (o *HogFunctionTemplate) SetMappingTemplates(v []HogFunctionMappingTemplate)`

SetMappingTemplates sets MappingTemplates field to given value.

### HasMappingTemplates

`func (o *HogFunctionTemplate) HasMappingTemplates() bool`

HasMappingTemplates returns a boolean if a field has been set.

### SetMappingTemplatesNil

`func (o *HogFunctionTemplate) SetMappingTemplatesNil(b bool)`

 SetMappingTemplatesNil sets the value for MappingTemplates to be an explicit nil

### UnsetMappingTemplates
`func (o *HogFunctionTemplate) UnsetMappingTemplates()`

UnsetMappingTemplates ensures that no value is present for MappingTemplates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


