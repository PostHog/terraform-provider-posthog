# ChatCompletionUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromptTokens** | **int32** |  | 
**CompletionTokens** | **int32** |  | 
**TotalTokens** | **int32** |  | 
**CompletionTokensDetails** | Pointer to **interface{}** |  | [optional] 
**PromptTokensDetails** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewChatCompletionUsage

`func NewChatCompletionUsage(promptTokens int32, completionTokens int32, totalTokens int32, ) *ChatCompletionUsage`

NewChatCompletionUsage instantiates a new ChatCompletionUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionUsageWithDefaults

`func NewChatCompletionUsageWithDefaults() *ChatCompletionUsage`

NewChatCompletionUsageWithDefaults instantiates a new ChatCompletionUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromptTokens

`func (o *ChatCompletionUsage) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *ChatCompletionUsage) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *ChatCompletionUsage) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### GetCompletionTokens

`func (o *ChatCompletionUsage) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *ChatCompletionUsage) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *ChatCompletionUsage) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.


### GetTotalTokens

`func (o *ChatCompletionUsage) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *ChatCompletionUsage) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *ChatCompletionUsage) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### GetCompletionTokensDetails

`func (o *ChatCompletionUsage) GetCompletionTokensDetails() interface{}`

GetCompletionTokensDetails returns the CompletionTokensDetails field if non-nil, zero value otherwise.

### GetCompletionTokensDetailsOk

`func (o *ChatCompletionUsage) GetCompletionTokensDetailsOk() (*interface{}, bool)`

GetCompletionTokensDetailsOk returns a tuple with the CompletionTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokensDetails

`func (o *ChatCompletionUsage) SetCompletionTokensDetails(v interface{})`

SetCompletionTokensDetails sets CompletionTokensDetails field to given value.

### HasCompletionTokensDetails

`func (o *ChatCompletionUsage) HasCompletionTokensDetails() bool`

HasCompletionTokensDetails returns a boolean if a field has been set.

### SetCompletionTokensDetailsNil

`func (o *ChatCompletionUsage) SetCompletionTokensDetailsNil(b bool)`

 SetCompletionTokensDetailsNil sets the value for CompletionTokensDetails to be an explicit nil

### UnsetCompletionTokensDetails
`func (o *ChatCompletionUsage) UnsetCompletionTokensDetails()`

UnsetCompletionTokensDetails ensures that no value is present for CompletionTokensDetails, not even an explicit nil
### GetPromptTokensDetails

`func (o *ChatCompletionUsage) GetPromptTokensDetails() interface{}`

GetPromptTokensDetails returns the PromptTokensDetails field if non-nil, zero value otherwise.

### GetPromptTokensDetailsOk

`func (o *ChatCompletionUsage) GetPromptTokensDetailsOk() (*interface{}, bool)`

GetPromptTokensDetailsOk returns a tuple with the PromptTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokensDetails

`func (o *ChatCompletionUsage) SetPromptTokensDetails(v interface{})`

SetPromptTokensDetails sets PromptTokensDetails field to given value.

### HasPromptTokensDetails

`func (o *ChatCompletionUsage) HasPromptTokensDetails() bool`

HasPromptTokensDetails returns a boolean if a field has been set.

### SetPromptTokensDetailsNil

`func (o *ChatCompletionUsage) SetPromptTokensDetailsNil(b bool)`

 SetPromptTokensDetailsNil sets the value for PromptTokensDetails to be an explicit nil

### UnsetPromptTokensDetails
`func (o *ChatCompletionUsage) UnsetPromptTokensDetails()`

UnsetPromptTokensDetails ensures that no value is present for PromptTokensDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


