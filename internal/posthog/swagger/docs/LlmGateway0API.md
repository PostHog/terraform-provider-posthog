# \LlmGatewayAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LlmGatewayV1ChatCompletionsCreate**](LlmGatewayAPI.md#LlmGatewayV1ChatCompletionsCreate) | **Post** /api/projects/{project_id}/llm_gateway/v1/chat/completions/ | OpenAI Chat Completions API
[**LlmGatewayV1MessagesCreate**](LlmGatewayAPI.md#LlmGatewayV1MessagesCreate) | **Post** /api/projects/{project_id}/llm_gateway/v1/messages/ | Anthropic Messages API



## LlmGatewayV1ChatCompletionsCreate

> ChatCompletionResponse LlmGatewayV1ChatCompletionsCreate(ctx, projectId).ChatCompletionRequest(chatCompletionRequest).Format(format).Execute()

OpenAI Chat Completions API



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/posthog/terraform-provider"
)

func main() {
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	chatCompletionRequest := *openapiclient.NewChatCompletionRequest("Model_example", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // ChatCompletionRequest | 
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayAPI.LlmGatewayV1ChatCompletionsCreate(context.Background(), projectId).ChatCompletionRequest(chatCompletionRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayAPI.LlmGatewayV1ChatCompletionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LlmGatewayV1ChatCompletionsCreate`: ChatCompletionResponse
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayAPI.LlmGatewayV1ChatCompletionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLlmGatewayV1ChatCompletionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatCompletionRequest** | [**ChatCompletionRequest**](ChatCompletionRequest.md) |  | 
 **format** | **string** |  | 

### Return type

[**ChatCompletionResponse**](ChatCompletionResponse.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LlmGatewayV1MessagesCreate

> AnthropicMessagesResponse LlmGatewayV1MessagesCreate(ctx, projectId).AnthropicMessagesRequest(anthropicMessagesRequest).Format(format).Execute()

Anthropic Messages API



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/posthog/terraform-provider"
)

func main() {
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	anthropicMessagesRequest := *openapiclient.NewAnthropicMessagesRequest("Model_example", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // AnthropicMessagesRequest | 
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LlmGatewayAPI.LlmGatewayV1MessagesCreate(context.Background(), projectId).AnthropicMessagesRequest(anthropicMessagesRequest).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LlmGatewayAPI.LlmGatewayV1MessagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LlmGatewayV1MessagesCreate`: AnthropicMessagesResponse
	fmt.Fprintf(os.Stdout, "Response from `LlmGatewayAPI.LlmGatewayV1MessagesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLlmGatewayV1MessagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **anthropicMessagesRequest** | [**AnthropicMessagesRequest**](AnthropicMessagesRequest.md) |  | 
 **format** | **string** |  | 

### Return type

[**AnthropicMessagesResponse**](AnthropicMessagesResponse.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

