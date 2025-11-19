# ChatCompletionMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**ChatCompletionMessageRoleEnum**](ChatCompletionMessageRoleEnum.md) |  | 
**Content** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FunctionCall** | Pointer to **interface{}** |  | [optional] 
**ToolCalls** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewChatCompletionMessage

`func NewChatCompletionMessage(role ChatCompletionMessageRoleEnum, ) *ChatCompletionMessage`

NewChatCompletionMessage instantiates a new ChatCompletionMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionMessageWithDefaults

`func NewChatCompletionMessageWithDefaults() *ChatCompletionMessage`

NewChatCompletionMessageWithDefaults instantiates a new ChatCompletionMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ChatCompletionMessage) GetRole() ChatCompletionMessageRoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatCompletionMessage) GetRoleOk() (*ChatCompletionMessageRoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatCompletionMessage) SetRole(v ChatCompletionMessageRoleEnum)`

SetRole sets Role field to given value.


### GetContent

`func (o *ChatCompletionMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatCompletionMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatCompletionMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ChatCompletionMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *ChatCompletionMessage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ChatCompletionMessage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetName

`func (o *ChatCompletionMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatCompletionMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatCompletionMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChatCompletionMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFunctionCall

`func (o *ChatCompletionMessage) GetFunctionCall() interface{}`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *ChatCompletionMessage) GetFunctionCallOk() (*interface{}, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *ChatCompletionMessage) SetFunctionCall(v interface{})`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *ChatCompletionMessage) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### SetFunctionCallNil

`func (o *ChatCompletionMessage) SetFunctionCallNil(b bool)`

 SetFunctionCallNil sets the value for FunctionCall to be an explicit nil

### UnsetFunctionCall
`func (o *ChatCompletionMessage) UnsetFunctionCall()`

UnsetFunctionCall ensures that no value is present for FunctionCall, not even an explicit nil
### GetToolCalls

`func (o *ChatCompletionMessage) GetToolCalls() []interface{}`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *ChatCompletionMessage) GetToolCallsOk() (*[]interface{}, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *ChatCompletionMessage) SetToolCalls(v []interface{})`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *ChatCompletionMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


