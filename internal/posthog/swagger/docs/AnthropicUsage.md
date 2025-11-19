# AnthropicUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputTokens** | **int32** |  | 
**OutputTokens** | **int32** |  | 
**CacheCreationInputTokens** | Pointer to **NullableInt32** |  | [optional] 
**CacheReadInputTokens** | Pointer to **NullableInt32** |  | [optional] 
**ServerToolUse** | Pointer to **interface{}** |  | [optional] 
**ServiceTier** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAnthropicUsage

`func NewAnthropicUsage(inputTokens int32, outputTokens int32, ) *AnthropicUsage`

NewAnthropicUsage instantiates a new AnthropicUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicUsageWithDefaults

`func NewAnthropicUsageWithDefaults() *AnthropicUsage`

NewAnthropicUsageWithDefaults instantiates a new AnthropicUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputTokens

`func (o *AnthropicUsage) GetInputTokens() int32`

GetInputTokens returns the InputTokens field if non-nil, zero value otherwise.

### GetInputTokensOk

`func (o *AnthropicUsage) GetInputTokensOk() (*int32, bool)`

GetInputTokensOk returns a tuple with the InputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTokens

`func (o *AnthropicUsage) SetInputTokens(v int32)`

SetInputTokens sets InputTokens field to given value.


### GetOutputTokens

`func (o *AnthropicUsage) GetOutputTokens() int32`

GetOutputTokens returns the OutputTokens field if non-nil, zero value otherwise.

### GetOutputTokensOk

`func (o *AnthropicUsage) GetOutputTokensOk() (*int32, bool)`

GetOutputTokensOk returns a tuple with the OutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputTokens

`func (o *AnthropicUsage) SetOutputTokens(v int32)`

SetOutputTokens sets OutputTokens field to given value.


### GetCacheCreationInputTokens

`func (o *AnthropicUsage) GetCacheCreationInputTokens() int32`

GetCacheCreationInputTokens returns the CacheCreationInputTokens field if non-nil, zero value otherwise.

### GetCacheCreationInputTokensOk

`func (o *AnthropicUsage) GetCacheCreationInputTokensOk() (*int32, bool)`

GetCacheCreationInputTokensOk returns a tuple with the CacheCreationInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationInputTokens

`func (o *AnthropicUsage) SetCacheCreationInputTokens(v int32)`

SetCacheCreationInputTokens sets CacheCreationInputTokens field to given value.

### HasCacheCreationInputTokens

`func (o *AnthropicUsage) HasCacheCreationInputTokens() bool`

HasCacheCreationInputTokens returns a boolean if a field has been set.

### SetCacheCreationInputTokensNil

`func (o *AnthropicUsage) SetCacheCreationInputTokensNil(b bool)`

 SetCacheCreationInputTokensNil sets the value for CacheCreationInputTokens to be an explicit nil

### UnsetCacheCreationInputTokens
`func (o *AnthropicUsage) UnsetCacheCreationInputTokens()`

UnsetCacheCreationInputTokens ensures that no value is present for CacheCreationInputTokens, not even an explicit nil
### GetCacheReadInputTokens

`func (o *AnthropicUsage) GetCacheReadInputTokens() int32`

GetCacheReadInputTokens returns the CacheReadInputTokens field if non-nil, zero value otherwise.

### GetCacheReadInputTokensOk

`func (o *AnthropicUsage) GetCacheReadInputTokensOk() (*int32, bool)`

GetCacheReadInputTokensOk returns a tuple with the CacheReadInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadInputTokens

`func (o *AnthropicUsage) SetCacheReadInputTokens(v int32)`

SetCacheReadInputTokens sets CacheReadInputTokens field to given value.

### HasCacheReadInputTokens

`func (o *AnthropicUsage) HasCacheReadInputTokens() bool`

HasCacheReadInputTokens returns a boolean if a field has been set.

### SetCacheReadInputTokensNil

`func (o *AnthropicUsage) SetCacheReadInputTokensNil(b bool)`

 SetCacheReadInputTokensNil sets the value for CacheReadInputTokens to be an explicit nil

### UnsetCacheReadInputTokens
`func (o *AnthropicUsage) UnsetCacheReadInputTokens()`

UnsetCacheReadInputTokens ensures that no value is present for CacheReadInputTokens, not even an explicit nil
### GetServerToolUse

`func (o *AnthropicUsage) GetServerToolUse() interface{}`

GetServerToolUse returns the ServerToolUse field if non-nil, zero value otherwise.

### GetServerToolUseOk

`func (o *AnthropicUsage) GetServerToolUseOk() (*interface{}, bool)`

GetServerToolUseOk returns a tuple with the ServerToolUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerToolUse

`func (o *AnthropicUsage) SetServerToolUse(v interface{})`

SetServerToolUse sets ServerToolUse field to given value.

### HasServerToolUse

`func (o *AnthropicUsage) HasServerToolUse() bool`

HasServerToolUse returns a boolean if a field has been set.

### SetServerToolUseNil

`func (o *AnthropicUsage) SetServerToolUseNil(b bool)`

 SetServerToolUseNil sets the value for ServerToolUse to be an explicit nil

### UnsetServerToolUse
`func (o *AnthropicUsage) UnsetServerToolUse()`

UnsetServerToolUse ensures that no value is present for ServerToolUse, not even an explicit nil
### GetServiceTier

`func (o *AnthropicUsage) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *AnthropicUsage) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *AnthropicUsage) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *AnthropicUsage) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *AnthropicUsage) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *AnthropicUsage) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


