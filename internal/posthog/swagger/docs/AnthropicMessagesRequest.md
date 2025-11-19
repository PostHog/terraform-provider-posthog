# AnthropicMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | The model to use for completion (e.g., &#39;claude-3-5-sonnet-20241022&#39;) | 
**Messages** | **[]map[string]interface{}** | List of message objects with &#39;role&#39; and &#39;content&#39; | 
**MaxTokens** | Pointer to **int32** | Maximum number of tokens to generate | [optional] [default to 4096]
**Temperature** | Pointer to **float64** | Sampling temperature between 0 and 1 | [optional] 
**TopP** | Pointer to **float64** | Nucleus sampling parameter | [optional] 
**TopK** | Pointer to **int32** | Top-k sampling parameter | [optional] 
**Stream** | Pointer to **bool** | Whether to stream the response | [optional] [default to false]
**StopSequences** | Pointer to **[]string** | Custom stop sequences | [optional] 
**System** | Pointer to **interface{}** | System prompt (string or array of content blocks) | [optional] 
**Metadata** | Pointer to **interface{}** | Metadata to attach to the request | [optional] 
**Thinking** | Pointer to **interface{}** | Thinking configuration for extended thinking | [optional] 
**Tools** | Pointer to **[]interface{}** | List of tools available to the model | [optional] 
**ToolChoice** | Pointer to **interface{}** | Controls which tool is called | [optional] 
**ServiceTier** | Pointer to [**AnthropicMessagesRequestServiceTierEnum**](AnthropicMessagesRequestServiceTierEnum.md) | Service tier for the request  * &#x60;auto&#x60; - auto * &#x60;standard_only&#x60; - standard_only | [optional] 

## Methods

### NewAnthropicMessagesRequest

`func NewAnthropicMessagesRequest(model string, messages []map[string]interface{}, ) *AnthropicMessagesRequest`

NewAnthropicMessagesRequest instantiates a new AnthropicMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicMessagesRequestWithDefaults

`func NewAnthropicMessagesRequestWithDefaults() *AnthropicMessagesRequest`

NewAnthropicMessagesRequestWithDefaults instantiates a new AnthropicMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AnthropicMessagesRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AnthropicMessagesRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AnthropicMessagesRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMessages

`func (o *AnthropicMessagesRequest) GetMessages() []map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AnthropicMessagesRequest) GetMessagesOk() (*[]map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AnthropicMessagesRequest) SetMessages(v []map[string]interface{})`

SetMessages sets Messages field to given value.


### GetMaxTokens

`func (o *AnthropicMessagesRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AnthropicMessagesRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AnthropicMessagesRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AnthropicMessagesRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTemperature

`func (o *AnthropicMessagesRequest) GetTemperature() float64`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AnthropicMessagesRequest) GetTemperatureOk() (*float64, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AnthropicMessagesRequest) SetTemperature(v float64)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AnthropicMessagesRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *AnthropicMessagesRequest) GetTopP() float64`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *AnthropicMessagesRequest) GetTopPOk() (*float64, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *AnthropicMessagesRequest) SetTopP(v float64)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *AnthropicMessagesRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetTopK

`func (o *AnthropicMessagesRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *AnthropicMessagesRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *AnthropicMessagesRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *AnthropicMessagesRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetStream

`func (o *AnthropicMessagesRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AnthropicMessagesRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AnthropicMessagesRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AnthropicMessagesRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStopSequences

`func (o *AnthropicMessagesRequest) GetStopSequences() []string`

GetStopSequences returns the StopSequences field if non-nil, zero value otherwise.

### GetStopSequencesOk

`func (o *AnthropicMessagesRequest) GetStopSequencesOk() (*[]string, bool)`

GetStopSequencesOk returns a tuple with the StopSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSequences

`func (o *AnthropicMessagesRequest) SetStopSequences(v []string)`

SetStopSequences sets StopSequences field to given value.

### HasStopSequences

`func (o *AnthropicMessagesRequest) HasStopSequences() bool`

HasStopSequences returns a boolean if a field has been set.

### GetSystem

`func (o *AnthropicMessagesRequest) GetSystem() interface{}`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AnthropicMessagesRequest) GetSystemOk() (*interface{}, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AnthropicMessagesRequest) SetSystem(v interface{})`

SetSystem sets System field to given value.

### HasSystem

`func (o *AnthropicMessagesRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *AnthropicMessagesRequest) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *AnthropicMessagesRequest) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetMetadata

`func (o *AnthropicMessagesRequest) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AnthropicMessagesRequest) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AnthropicMessagesRequest) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AnthropicMessagesRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AnthropicMessagesRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AnthropicMessagesRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetThinking

`func (o *AnthropicMessagesRequest) GetThinking() interface{}`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *AnthropicMessagesRequest) GetThinkingOk() (*interface{}, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *AnthropicMessagesRequest) SetThinking(v interface{})`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *AnthropicMessagesRequest) HasThinking() bool`

HasThinking returns a boolean if a field has been set.

### SetThinkingNil

`func (o *AnthropicMessagesRequest) SetThinkingNil(b bool)`

 SetThinkingNil sets the value for Thinking to be an explicit nil

### UnsetThinking
`func (o *AnthropicMessagesRequest) UnsetThinking()`

UnsetThinking ensures that no value is present for Thinking, not even an explicit nil
### GetTools

`func (o *AnthropicMessagesRequest) GetTools() []interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AnthropicMessagesRequest) GetToolsOk() (*[]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AnthropicMessagesRequest) SetTools(v []interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *AnthropicMessagesRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *AnthropicMessagesRequest) GetToolChoice() interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *AnthropicMessagesRequest) GetToolChoiceOk() (*interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *AnthropicMessagesRequest) SetToolChoice(v interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *AnthropicMessagesRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *AnthropicMessagesRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *AnthropicMessagesRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetServiceTier

`func (o *AnthropicMessagesRequest) GetServiceTier() AnthropicMessagesRequestServiceTierEnum`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *AnthropicMessagesRequest) GetServiceTierOk() (*AnthropicMessagesRequestServiceTierEnum, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *AnthropicMessagesRequest) SetServiceTier(v AnthropicMessagesRequestServiceTierEnum)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *AnthropicMessagesRequest) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


