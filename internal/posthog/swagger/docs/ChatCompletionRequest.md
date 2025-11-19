# ChatCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | The model to use for completion (e.g., &#39;gpt-4&#39;, &#39;gpt-3.5-turbo&#39;) | 
**Messages** | **[]map[string]interface{}** | List of message objects with &#39;role&#39; and &#39;content&#39; | 
**Temperature** | Pointer to **float64** | Sampling temperature between 0 and 2 | [optional] 
**TopP** | Pointer to **float64** | Nucleus sampling parameter | [optional] 
**N** | Pointer to **int32** | Number of completions to generate | [optional] 
**Stream** | Pointer to **bool** | Whether to stream the response | [optional] [default to false]
**StreamOptions** | Pointer to **interface{}** | Additional options for streaming | [optional] 
**Stop** | Pointer to **[]string** | Stop sequences | [optional] 
**MaxTokens** | Pointer to **int32** | Maximum number of tokens to generate | [optional] 
**MaxCompletionTokens** | Pointer to **int32** | Maximum number of completion tokens (alternative to max_tokens) | [optional] 
**PresencePenalty** | Pointer to **float64** | Presence penalty between -2.0 and 2.0 | [optional] 
**FrequencyPenalty** | Pointer to **float64** | Frequency penalty between -2.0 and 2.0 | [optional] 
**LogitBias** | Pointer to **interface{}** | Logit bias mapping | [optional] 
**User** | Pointer to **string** | Unique user identifier | [optional] 
**Tools** | Pointer to **[]interface{}** | List of tools available to the model | [optional] 
**ToolChoice** | Pointer to **interface{}** | Controls which tool is called | [optional] 
**ParallelToolCalls** | Pointer to **bool** | Whether to allow parallel tool calls | [optional] 
**ResponseFormat** | Pointer to **interface{}** | Format for the model output | [optional] 
**Seed** | Pointer to **int32** | Random seed for deterministic sampling | [optional] 
**Logprobs** | Pointer to **bool** | Whether to return log probabilities | [optional] 
**TopLogprobs** | Pointer to **int32** | Number of most likely tokens to return at each position | [optional] 
**Modalities** | Pointer to [**[]ModalitiesEnum**](ModalitiesEnum.md) | Output modalities | [optional] 
**Prediction** | Pointer to **interface{}** | Prediction content for speculative decoding | [optional] 
**Audio** | Pointer to **interface{}** | Audio input parameters | [optional] 
**ReasoningEffort** | Pointer to [**ReasoningEffortEnum**](ReasoningEffortEnum.md) | Reasoning effort level for o-series models  * &#x60;none&#x60; - none * &#x60;minimal&#x60; - minimal * &#x60;low&#x60; - low * &#x60;medium&#x60; - medium * &#x60;high&#x60; - high * &#x60;default&#x60; - default | [optional] 
**Verbosity** | Pointer to [**VerbosityEnum**](VerbosityEnum.md) | Controls the verbosity level of the model&#39;s output  * &#x60;concise&#x60; - concise * &#x60;standard&#x60; - standard * &#x60;verbose&#x60; - verbose | [optional] 
**Store** | Pointer to **bool** | Whether to store the output for model distillation or evals | [optional] 
**WebSearchOptions** | Pointer to **interface{}** | Web search tool configuration | [optional] 
**Functions** | Pointer to **[]interface{}** | Deprecated in favor of tools. List of functions the model may call | [optional] 
**FunctionCall** | Pointer to **interface{}** | Deprecated in favor of tool_choice. Controls which function is called | [optional] 

## Methods

### NewChatCompletionRequest

`func NewChatCompletionRequest(model string, messages []map[string]interface{}, ) *ChatCompletionRequest`

NewChatCompletionRequest instantiates a new ChatCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestWithDefaults

`func NewChatCompletionRequestWithDefaults() *ChatCompletionRequest`

NewChatCompletionRequestWithDefaults instantiates a new ChatCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ChatCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMessages

`func (o *ChatCompletionRequest) GetMessages() []map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatCompletionRequest) GetMessagesOk() (*[]map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatCompletionRequest) SetMessages(v []map[string]interface{})`

SetMessages sets Messages field to given value.


### GetTemperature

`func (o *ChatCompletionRequest) GetTemperature() float64`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ChatCompletionRequest) GetTemperatureOk() (*float64, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ChatCompletionRequest) SetTemperature(v float64)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ChatCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *ChatCompletionRequest) GetTopP() float64`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ChatCompletionRequest) GetTopPOk() (*float64, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ChatCompletionRequest) SetTopP(v float64)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ChatCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetN

`func (o *ChatCompletionRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ChatCompletionRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ChatCompletionRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *ChatCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetStream

`func (o *ChatCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStreamOptions

`func (o *ChatCompletionRequest) GetStreamOptions() interface{}`

GetStreamOptions returns the StreamOptions field if non-nil, zero value otherwise.

### GetStreamOptionsOk

`func (o *ChatCompletionRequest) GetStreamOptionsOk() (*interface{}, bool)`

GetStreamOptionsOk returns a tuple with the StreamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamOptions

`func (o *ChatCompletionRequest) SetStreamOptions(v interface{})`

SetStreamOptions sets StreamOptions field to given value.

### HasStreamOptions

`func (o *ChatCompletionRequest) HasStreamOptions() bool`

HasStreamOptions returns a boolean if a field has been set.

### SetStreamOptionsNil

`func (o *ChatCompletionRequest) SetStreamOptionsNil(b bool)`

 SetStreamOptionsNil sets the value for StreamOptions to be an explicit nil

### UnsetStreamOptions
`func (o *ChatCompletionRequest) UnsetStreamOptions()`

UnsetStreamOptions ensures that no value is present for StreamOptions, not even an explicit nil
### GetStop

`func (o *ChatCompletionRequest) GetStop() []string`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *ChatCompletionRequest) GetStopOk() (*[]string, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *ChatCompletionRequest) SetStop(v []string)`

SetStop sets Stop field to given value.

### HasStop

`func (o *ChatCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetMaxTokens

`func (o *ChatCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ChatCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ChatCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ChatCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetMaxCompletionTokens

`func (o *ChatCompletionRequest) GetMaxCompletionTokens() int32`

GetMaxCompletionTokens returns the MaxCompletionTokens field if non-nil, zero value otherwise.

### GetMaxCompletionTokensOk

`func (o *ChatCompletionRequest) GetMaxCompletionTokensOk() (*int32, bool)`

GetMaxCompletionTokensOk returns a tuple with the MaxCompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCompletionTokens

`func (o *ChatCompletionRequest) SetMaxCompletionTokens(v int32)`

SetMaxCompletionTokens sets MaxCompletionTokens field to given value.

### HasMaxCompletionTokens

`func (o *ChatCompletionRequest) HasMaxCompletionTokens() bool`

HasMaxCompletionTokens returns a boolean if a field has been set.

### GetPresencePenalty

`func (o *ChatCompletionRequest) GetPresencePenalty() float64`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *ChatCompletionRequest) GetPresencePenaltyOk() (*float64, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *ChatCompletionRequest) SetPresencePenalty(v float64)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *ChatCompletionRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### GetFrequencyPenalty

`func (o *ChatCompletionRequest) GetFrequencyPenalty() float64`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *ChatCompletionRequest) GetFrequencyPenaltyOk() (*float64, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *ChatCompletionRequest) SetFrequencyPenalty(v float64)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *ChatCompletionRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### GetLogitBias

`func (o *ChatCompletionRequest) GetLogitBias() interface{}`

GetLogitBias returns the LogitBias field if non-nil, zero value otherwise.

### GetLogitBiasOk

`func (o *ChatCompletionRequest) GetLogitBiasOk() (*interface{}, bool)`

GetLogitBiasOk returns a tuple with the LogitBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogitBias

`func (o *ChatCompletionRequest) SetLogitBias(v interface{})`

SetLogitBias sets LogitBias field to given value.

### HasLogitBias

`func (o *ChatCompletionRequest) HasLogitBias() bool`

HasLogitBias returns a boolean if a field has been set.

### SetLogitBiasNil

`func (o *ChatCompletionRequest) SetLogitBiasNil(b bool)`

 SetLogitBiasNil sets the value for LogitBias to be an explicit nil

### UnsetLogitBias
`func (o *ChatCompletionRequest) UnsetLogitBias()`

UnsetLogitBias ensures that no value is present for LogitBias, not even an explicit nil
### GetUser

`func (o *ChatCompletionRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatCompletionRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatCompletionRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ChatCompletionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTools

`func (o *ChatCompletionRequest) GetTools() []interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ChatCompletionRequest) GetToolsOk() (*[]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ChatCompletionRequest) SetTools(v []interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *ChatCompletionRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *ChatCompletionRequest) GetToolChoice() interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *ChatCompletionRequest) GetToolChoiceOk() (*interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *ChatCompletionRequest) SetToolChoice(v interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *ChatCompletionRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *ChatCompletionRequest) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *ChatCompletionRequest) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetParallelToolCalls

`func (o *ChatCompletionRequest) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *ChatCompletionRequest) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *ChatCompletionRequest) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *ChatCompletionRequest) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### GetResponseFormat

`func (o *ChatCompletionRequest) GetResponseFormat() interface{}`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ChatCompletionRequest) GetResponseFormatOk() (*interface{}, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ChatCompletionRequest) SetResponseFormat(v interface{})`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ChatCompletionRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### SetResponseFormatNil

`func (o *ChatCompletionRequest) SetResponseFormatNil(b bool)`

 SetResponseFormatNil sets the value for ResponseFormat to be an explicit nil

### UnsetResponseFormat
`func (o *ChatCompletionRequest) UnsetResponseFormat()`

UnsetResponseFormat ensures that no value is present for ResponseFormat, not even an explicit nil
### GetSeed

`func (o *ChatCompletionRequest) GetSeed() int32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *ChatCompletionRequest) GetSeedOk() (*int32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *ChatCompletionRequest) SetSeed(v int32)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *ChatCompletionRequest) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetLogprobs

`func (o *ChatCompletionRequest) GetLogprobs() bool`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *ChatCompletionRequest) GetLogprobsOk() (*bool, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *ChatCompletionRequest) SetLogprobs(v bool)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *ChatCompletionRequest) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### GetTopLogprobs

`func (o *ChatCompletionRequest) GetTopLogprobs() int32`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ChatCompletionRequest) GetTopLogprobsOk() (*int32, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ChatCompletionRequest) SetTopLogprobs(v int32)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *ChatCompletionRequest) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### GetModalities

`func (o *ChatCompletionRequest) GetModalities() []ModalitiesEnum`

GetModalities returns the Modalities field if non-nil, zero value otherwise.

### GetModalitiesOk

`func (o *ChatCompletionRequest) GetModalitiesOk() (*[]ModalitiesEnum, bool)`

GetModalitiesOk returns a tuple with the Modalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModalities

`func (o *ChatCompletionRequest) SetModalities(v []ModalitiesEnum)`

SetModalities sets Modalities field to given value.

### HasModalities

`func (o *ChatCompletionRequest) HasModalities() bool`

HasModalities returns a boolean if a field has been set.

### GetPrediction

`func (o *ChatCompletionRequest) GetPrediction() interface{}`

GetPrediction returns the Prediction field if non-nil, zero value otherwise.

### GetPredictionOk

`func (o *ChatCompletionRequest) GetPredictionOk() (*interface{}, bool)`

GetPredictionOk returns a tuple with the Prediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrediction

`func (o *ChatCompletionRequest) SetPrediction(v interface{})`

SetPrediction sets Prediction field to given value.

### HasPrediction

`func (o *ChatCompletionRequest) HasPrediction() bool`

HasPrediction returns a boolean if a field has been set.

### SetPredictionNil

`func (o *ChatCompletionRequest) SetPredictionNil(b bool)`

 SetPredictionNil sets the value for Prediction to be an explicit nil

### UnsetPrediction
`func (o *ChatCompletionRequest) UnsetPrediction()`

UnsetPrediction ensures that no value is present for Prediction, not even an explicit nil
### GetAudio

`func (o *ChatCompletionRequest) GetAudio() interface{}`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *ChatCompletionRequest) GetAudioOk() (*interface{}, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *ChatCompletionRequest) SetAudio(v interface{})`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *ChatCompletionRequest) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudioNil

`func (o *ChatCompletionRequest) SetAudioNil(b bool)`

 SetAudioNil sets the value for Audio to be an explicit nil

### UnsetAudio
`func (o *ChatCompletionRequest) UnsetAudio()`

UnsetAudio ensures that no value is present for Audio, not even an explicit nil
### GetReasoningEffort

`func (o *ChatCompletionRequest) GetReasoningEffort() ReasoningEffortEnum`

GetReasoningEffort returns the ReasoningEffort field if non-nil, zero value otherwise.

### GetReasoningEffortOk

`func (o *ChatCompletionRequest) GetReasoningEffortOk() (*ReasoningEffortEnum, bool)`

GetReasoningEffortOk returns a tuple with the ReasoningEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEffort

`func (o *ChatCompletionRequest) SetReasoningEffort(v ReasoningEffortEnum)`

SetReasoningEffort sets ReasoningEffort field to given value.

### HasReasoningEffort

`func (o *ChatCompletionRequest) HasReasoningEffort() bool`

HasReasoningEffort returns a boolean if a field has been set.

### GetVerbosity

`func (o *ChatCompletionRequest) GetVerbosity() VerbosityEnum`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *ChatCompletionRequest) GetVerbosityOk() (*VerbosityEnum, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *ChatCompletionRequest) SetVerbosity(v VerbosityEnum)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *ChatCompletionRequest) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.

### GetStore

`func (o *ChatCompletionRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ChatCompletionRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ChatCompletionRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ChatCompletionRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetWebSearchOptions

`func (o *ChatCompletionRequest) GetWebSearchOptions() interface{}`

GetWebSearchOptions returns the WebSearchOptions field if non-nil, zero value otherwise.

### GetWebSearchOptionsOk

`func (o *ChatCompletionRequest) GetWebSearchOptionsOk() (*interface{}, bool)`

GetWebSearchOptionsOk returns a tuple with the WebSearchOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchOptions

`func (o *ChatCompletionRequest) SetWebSearchOptions(v interface{})`

SetWebSearchOptions sets WebSearchOptions field to given value.

### HasWebSearchOptions

`func (o *ChatCompletionRequest) HasWebSearchOptions() bool`

HasWebSearchOptions returns a boolean if a field has been set.

### SetWebSearchOptionsNil

`func (o *ChatCompletionRequest) SetWebSearchOptionsNil(b bool)`

 SetWebSearchOptionsNil sets the value for WebSearchOptions to be an explicit nil

### UnsetWebSearchOptions
`func (o *ChatCompletionRequest) UnsetWebSearchOptions()`

UnsetWebSearchOptions ensures that no value is present for WebSearchOptions, not even an explicit nil
### GetFunctions

`func (o *ChatCompletionRequest) GetFunctions() []interface{}`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ChatCompletionRequest) GetFunctionsOk() (*[]interface{}, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ChatCompletionRequest) SetFunctions(v []interface{})`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *ChatCompletionRequest) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetFunctionCall

`func (o *ChatCompletionRequest) GetFunctionCall() interface{}`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *ChatCompletionRequest) GetFunctionCallOk() (*interface{}, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *ChatCompletionRequest) SetFunctionCall(v interface{})`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *ChatCompletionRequest) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### SetFunctionCallNil

`func (o *ChatCompletionRequest) SetFunctionCallNil(b bool)`

 SetFunctionCallNil sets the value for FunctionCall to be an explicit nil

### UnsetFunctionCall
`func (o *ChatCompletionRequest) UnsetFunctionCall()`

UnsetFunctionCall ensures that no value is present for FunctionCall, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


