# \AdvancedActivityLogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvancedActivityLogsAvailableFiltersRetrieve**](AdvancedActivityLogsAPI.md#AdvancedActivityLogsAvailableFiltersRetrieve) | **Get** /api/projects/{project_id}/advanced_activity_logs/available_filters/ | 
[**AdvancedActivityLogsExportCreate**](AdvancedActivityLogsAPI.md#AdvancedActivityLogsExportCreate) | **Post** /api/projects/{project_id}/advanced_activity_logs/export/ | 
[**AdvancedActivityLogsList**](AdvancedActivityLogsAPI.md#AdvancedActivityLogsList) | **Get** /api/projects/{project_id}/advanced_activity_logs/ | 



## AdvancedActivityLogsAvailableFiltersRetrieve

> ActivityLog AdvancedActivityLogsAvailableFiltersRetrieve(ctx, projectId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedActivityLogsAPI.AdvancedActivityLogsAvailableFiltersRetrieve(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedActivityLogsAPI.AdvancedActivityLogsAvailableFiltersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdvancedActivityLogsAvailableFiltersRetrieve`: ActivityLog
	fmt.Fprintf(os.Stdout, "Response from `AdvancedActivityLogsAPI.AdvancedActivityLogsAvailableFiltersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdvancedActivityLogsAvailableFiltersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActivityLog**](ActivityLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdvancedActivityLogsExportCreate

> ActivityLog AdvancedActivityLogsExportCreate(ctx, projectId).ActivityLog(activityLog).Execute()



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
	activityLog := *openapiclient.NewActivityLog("Id_example", *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), false, "Activity_example", "Scope_example") // ActivityLog | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedActivityLogsAPI.AdvancedActivityLogsExportCreate(context.Background(), projectId).ActivityLog(activityLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedActivityLogsAPI.AdvancedActivityLogsExportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdvancedActivityLogsExportCreate`: ActivityLog
	fmt.Fprintf(os.Stdout, "Response from `AdvancedActivityLogsAPI.AdvancedActivityLogsExportCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdvancedActivityLogsExportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activityLog** | [**ActivityLog**](ActivityLog.md) |  | 

### Return type

[**ActivityLog**](ActivityLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdvancedActivityLogsList

> []ActivityLog AdvancedActivityLogsList(ctx, projectId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedActivityLogsAPI.AdvancedActivityLogsList(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedActivityLogsAPI.AdvancedActivityLogsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdvancedActivityLogsList`: []ActivityLog
	fmt.Fprintf(os.Stdout, "Response from `AdvancedActivityLogsAPI.AdvancedActivityLogsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdvancedActivityLogsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ActivityLog**](ActivityLog.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

