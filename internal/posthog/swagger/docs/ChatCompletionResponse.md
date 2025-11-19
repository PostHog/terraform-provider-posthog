# ChatCompletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | [**ObjectEnum**](ObjectEnum.md) |  | 
**Created** | **int32** |  | 
**Model** | **string** |  | 
**Choices** | [**[]ChatCompletionChoice**](ChatCompletionChoice.md) |  | 
**Usage** | Pointer to [**NullableChatCompletionUsage**](ChatCompletionUsage.md) |  | [optional] 
**SystemFingerprint** | Pointer to **NullableString** |  | [optional] 
**ServiceTier** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChatCompletionResponse

`func NewChatCompletionResponse(id string, object ObjectEnum, created int32, model string, choices []ChatCompletionChoice, ) *ChatCompletionResponse`

NewChatCompletionResponse instantiates a new ChatCompletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionResponseWithDefaults

`func NewChatCompletionResponseWithDefaults() *ChatCompletionResponse`

NewChatCompletionResponseWithDefaults instantiates a new ChatCompletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatCompletionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatCompletionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatCompletionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ChatCompletionResponse) GetObject() ObjectEnum`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChatCompletionResponse) GetObjectOk() (*ObjectEnum, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChatCompletionResponse) SetObject(v ObjectEnum)`

SetObject sets Object field to given value.


### GetCreated

`func (o *ChatCompletionResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChatCompletionResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChatCompletionResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetModel

`func (o *ChatCompletionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetChoices

`func (o *ChatCompletionResponse) GetChoices() []ChatCompletionChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ChatCompletionResponse) GetChoicesOk() (*[]ChatCompletionChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ChatCompletionResponse) SetChoices(v []ChatCompletionChoice)`

SetChoices sets Choices field to given value.


### GetUsage

`func (o *ChatCompletionResponse) GetUsage() ChatCompletionUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChatCompletionResponse) GetUsageOk() (*ChatCompletionUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChatCompletionResponse) SetUsage(v ChatCompletionUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChatCompletionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *ChatCompletionResponse) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *ChatCompletionResponse) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetSystemFingerprint

`func (o *ChatCompletionResponse) GetSystemFingerprint() string`

GetSystemFingerprint returns the SystemFingerprint field if non-nil, zero value otherwise.

### GetSystemFingerprintOk

`func (o *ChatCompletionResponse) GetSystemFingerprintOk() (*string, bool)`

GetSystemFingerprintOk returns a tuple with the SystemFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFingerprint

`func (o *ChatCompletionResponse) SetSystemFingerprint(v string)`

SetSystemFingerprint sets SystemFingerprint field to given value.

### HasSystemFingerprint

`func (o *ChatCompletionResponse) HasSystemFingerprint() bool`

HasSystemFingerprint returns a boolean if a field has been set.

### SetSystemFingerprintNil

`func (o *ChatCompletionResponse) SetSystemFingerprintNil(b bool)`

 SetSystemFingerprintNil sets the value for SystemFingerprint to be an explicit nil

### UnsetSystemFingerprint
`func (o *ChatCompletionResponse) UnsetSystemFingerprint()`

UnsetSystemFingerprint ensures that no value is present for SystemFingerprint, not even an explicit nil
### GetServiceTier

`func (o *ChatCompletionResponse) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *ChatCompletionResponse) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *ChatCompletionResponse) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *ChatCompletionResponse) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *ChatCompletionResponse) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *ChatCompletionResponse) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


