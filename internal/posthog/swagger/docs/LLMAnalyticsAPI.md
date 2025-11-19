# \LLMAnalyticsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsLlmAnalyticsTextReprCreate**](LLMAnalyticsAPI.md#EnvironmentsLlmAnalyticsTextReprCreate) | **Post** /api/environments/{project_id}/llm_analytics/text_repr/ | 



## EnvironmentsLlmAnalyticsTextReprCreate

> TextReprResponse EnvironmentsLlmAnalyticsTextReprCreate(ctx, projectId).TextReprRequest(textReprRequest).Execute()





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
	textReprRequest := *openapiclient.NewTextReprRequest(openapiclient.EventTypeEnum("$ai_generation"), interface{}(123)) // TextReprRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LLMAnalyticsAPI.EnvironmentsLlmAnalyticsTextReprCreate(context.Background(), projectId).TextReprRequest(textReprRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LLMAnalyticsAPI.EnvironmentsLlmAnalyticsTextReprCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsLlmAnalyticsTextReprCreate`: TextReprResponse
	fmt.Fprintf(os.Stdout, "Response from `LLMAnalyticsAPI.EnvironmentsLlmAnalyticsTextReprCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsLlmAnalyticsTextReprCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **textReprRequest** | [**TextReprRequest**](TextReprRequest.md) |  | 

### Return type

[**TextReprResponse**](TextReprResponse.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

