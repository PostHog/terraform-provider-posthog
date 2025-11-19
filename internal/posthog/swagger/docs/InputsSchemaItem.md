# InputsSchemaItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InputsSchemaItemTypeEnum**](InputsSchemaItemTypeEnum.md) |  | 
**Key** | **string** |  | 
**Label** | Pointer to **string** |  | [optional] 
**Choices** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] [default to false]
**Default** | Pointer to **interface{}** |  | [optional] 
**Secret** | Pointer to **bool** |  | [optional] [default to false]
**Hidden** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to **string** |  | [optional] 
**IntegrationKey** | Pointer to **string** |  | [optional] 
**RequiresField** | Pointer to **string** |  | [optional] 
**IntegrationField** | Pointer to **string** |  | [optional] 
**RequiredScopes** | Pointer to **string** |  | [optional] 
**Templating** | Pointer to [**InputsSchemaItemTemplatingEnum**](InputsSchemaItemTemplatingEnum.md) |  | [optional] 

## Methods

### NewInputsSchemaItem

`func NewInputsSchemaItem(type_ InputsSchemaItemTypeEnum, key string, ) *InputsSchemaItem`

NewInputsSchemaItem instantiates a new InputsSchemaItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputsSchemaItemWithDefaults

`func NewInputsSchemaItemWithDefaults() *InputsSchemaItem`

NewInputsSchemaItemWithDefaults instantiates a new InputsSchemaItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputsSchemaItem) GetType() InputsSchemaItemTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputsSchemaItem) GetTypeOk() (*InputsSchemaItemTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputsSchemaItem) SetType(v InputsSchemaItemTypeEnum)`

SetType sets Type field to given value.


### GetKey

`func (o *InputsSchemaItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *InputsSchemaItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *InputsSchemaItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *InputsSchemaItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InputsSchemaItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InputsSchemaItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InputsSchemaItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetChoices

`func (o *InputsSchemaItem) GetChoices() []map[string]interface{}`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *InputsSchemaItem) GetChoicesOk() (*[]map[string]interface{}, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *InputsSchemaItem) SetChoices(v []map[string]interface{})`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *InputsSchemaItem) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetRequired

`func (o *InputsSchemaItem) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *InputsSchemaItem) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *InputsSchemaItem) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *InputsSchemaItem) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDefault

`func (o *InputsSchemaItem) GetDefault() interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *InputsSchemaItem) GetDefaultOk() (*interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *InputsSchemaItem) SetDefault(v interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *InputsSchemaItem) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *InputsSchemaItem) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *InputsSchemaItem) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetSecret

`func (o *InputsSchemaItem) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InputsSchemaItem) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InputsSchemaItem) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InputsSchemaItem) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetHidden

`func (o *InputsSchemaItem) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *InputsSchemaItem) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *InputsSchemaItem) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *InputsSchemaItem) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDescription

`func (o *InputsSchemaItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InputsSchemaItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InputsSchemaItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InputsSchemaItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIntegration

`func (o *InputsSchemaItem) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *InputsSchemaItem) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *InputsSchemaItem) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *InputsSchemaItem) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetIntegrationKey

`func (o *InputsSchemaItem) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *InputsSchemaItem) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *InputsSchemaItem) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.

### HasIntegrationKey

`func (o *InputsSchemaItem) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.

### GetRequiresField

`func (o *InputsSchemaItem) GetRequiresField() string`

GetRequiresField returns the RequiresField field if non-nil, zero value otherwise.

### GetRequiresFieldOk

`func (o *InputsSchemaItem) GetRequiresFieldOk() (*string, bool)`

GetRequiresFieldOk returns a tuple with the RequiresField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresField

`func (o *InputsSchemaItem) SetRequiresField(v string)`

SetRequiresField sets RequiresField field to given value.

### HasRequiresField

`func (o *InputsSchemaItem) HasRequiresField() bool`

HasRequiresField returns a boolean if a field has been set.

### GetIntegrationField

`func (o *InputsSchemaItem) GetIntegrationField() string`

GetIntegrationField returns the IntegrationField field if non-nil, zero value otherwise.

### GetIntegrationFieldOk

`func (o *InputsSchemaItem) GetIntegrationFieldOk() (*string, bool)`

GetIntegrationFieldOk returns a tuple with the IntegrationField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationField

`func (o *InputsSchemaItem) SetIntegrationField(v string)`

SetIntegrationField sets IntegrationField field to given value.

### HasIntegrationField

`func (o *InputsSchemaItem) HasIntegrationField() bool`

HasIntegrationField returns a boolean if a field has been set.

### GetRequiredScopes

`func (o *InputsSchemaItem) GetRequiredScopes() string`

GetRequiredScopes returns the RequiredScopes field if non-nil, zero value otherwise.

### GetRequiredScopesOk

`func (o *InputsSchemaItem) GetRequiredScopesOk() (*string, bool)`

GetRequiredScopesOk returns a tuple with the RequiredScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredScopes

`func (o *InputsSchemaItem) SetRequiredScopes(v string)`

SetRequiredScopes sets RequiredScopes field to given value.

### HasRequiredScopes

`func (o *InputsSchemaItem) HasRequiredScopes() bool`

HasRequiredScopes returns a boolean if a field has been set.

### GetTemplating

`func (o *InputsSchemaItem) GetTemplating() InputsSchemaItemTemplatingEnum`

GetTemplating returns the Templating field if non-nil, zero value otherwise.

### GetTemplatingOk

`func (o *InputsSchemaItem) GetTemplatingOk() (*InputsSchemaItemTemplatingEnum, bool)`

GetTemplatingOk returns a tuple with the Templating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplating

`func (o *InputsSchemaItem) SetTemplating(v InputsSchemaItemTemplatingEnum)`

SetTemplating sets Templating field to given value.

### HasTemplating

`func (o *InputsSchemaItem) HasTemplating() bool`

HasTemplating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


