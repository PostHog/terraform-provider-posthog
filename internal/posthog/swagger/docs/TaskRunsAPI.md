# \TaskRunsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksRunsAppendLogCreate**](TaskRunsAPI.md#TasksRunsAppendLogCreate) | **Post** /api/projects/{project_id}/tasks/{task_id}/runs/{id}/append_log/ | Append log entries
[**TasksRunsCreate**](TaskRunsAPI.md#TasksRunsCreate) | **Post** /api/projects/{project_id}/tasks/{task_id}/runs/ | Create task run
[**TasksRunsList**](TaskRunsAPI.md#TasksRunsList) | **Get** /api/projects/{project_id}/tasks/{task_id}/runs/ | List task runs
[**TasksRunsPartialUpdate**](TaskRunsAPI.md#TasksRunsPartialUpdate) | **Patch** /api/projects/{project_id}/tasks/{task_id}/runs/{id}/ | 
[**TasksRunsRetrieve**](TaskRunsAPI.md#TasksRunsRetrieve) | **Get** /api/projects/{project_id}/tasks/{task_id}/runs/{id}/ | 
[**TasksRunsSetOutputPartialUpdate**](TaskRunsAPI.md#TasksRunsSetOutputPartialUpdate) | **Patch** /api/projects/{project_id}/tasks/{task_id}/runs/{id}/set_output/ | Set run output



## TasksRunsAppendLogCreate

> TaskRunDetail TasksRunsAppendLogCreate(ctx, id, projectId, taskId).TaskRunAppendLogRequest(taskRunAppendLogRequest).Execute()

Append log entries



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task run.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	taskRunAppendLogRequest := *openapiclient.NewTaskRunAppendLogRequest([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // TaskRunAppendLogRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskRunsAPI.TasksRunsAppendLogCreate(context.Background(), id, projectId, taskId).TaskRunAppendLogRequest(taskRunAppendLogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskRunsAPI.TasksRunsAppendLogCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksRunsAppendLogCreate`: TaskRunDetail
	fmt.Fprintf(os.Stdout, "Response from `TaskRunsAPI.TasksRunsAppendLogCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this task run. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRunsAppendLogCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **taskRunAppendLogRequest** | [**TaskRunAppendLogRequest**](TaskRunAppendLogRequest.md) |  | 

### Return type

[**TaskRunDetail**](TaskRunDetail.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRunsCreate

> TaskRunDetail TasksRunsCreate(ctx, projectId, taskId).Execute()

Create task run



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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskRunsAPI.TasksRunsCreate(context.Background(), projectId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskRunsAPI.TasksRunsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksRunsCreate`: TaskRunDetail
	fmt.Fprintf(os.Stdout, "Response from `TaskRunsAPI.TasksRunsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRunsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaskRunDetail**](TaskRunDetail.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRunsList

> PaginatedTaskRunDetailList TasksRunsList(ctx, projectId, taskId).Limit(limit).Offset(offset).Execute()

List task runs



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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskRunsAPI.TasksRunsList(context.Background(), projectId, taskId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskRunsAPI.TasksRunsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksRunsList`: PaginatedTaskRunDetailList
	fmt.Fprintf(os.Stdout, "Response from `TaskRunsAPI.TasksRunsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRunsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedTaskRunDetailList**](PaginatedTaskRunDetailList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRunsPartialUpdate

> TaskRunDetail TasksRunsPartialUpdate(ctx, id, projectId, taskId).PatchedTaskRunDetail(patchedTaskRunDetail).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task run.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchedTaskRunDetail := *openapiclient.NewPatchedTaskRunDetail() // PatchedTaskRunDetail |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskRunsAPI.TasksRunsPartialUpdate(context.Background(), id, projectId, taskId).PatchedTaskRunDetail(patchedTaskRunDetail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskRunsAPI.TasksRunsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksRunsPartialUpdate`: TaskRunDetail
	fmt.Fprintf(os.Stdout, "Response from `TaskRunsAPI.TasksRunsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this task run. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRunsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchedTaskRunDetail** | [**PatchedTaskRunDetail**](PatchedTaskRunDetail.md) |  | 

### Return type

[**TaskRunDetail**](TaskRunDetail.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRunsRetrieve

> TaskRunDetail TasksRunsRetrieve(ctx, id, projectId, taskId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task run.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskRunsAPI.TasksRunsRetrieve(context.Background(), id, projectId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskRunsAPI.TasksRunsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksRunsRetrieve`: TaskRunDetail
	fmt.Fprintf(os.Stdout, "Response from `TaskRunsAPI.TasksRunsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this task run. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRunsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TaskRunDetail**](TaskRunDetail.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRunsSetOutputPartialUpdate

> TaskRunDetail TasksRunsSetOutputPartialUpdate(ctx, id, projectId, taskId).Execute()

Set run output



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this task run.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskRunsAPI.TasksRunsSetOutputPartialUpdate(context.Background(), id, projectId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskRunsAPI.TasksRunsSetOutputPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksRunsSetOutputPartialUpdate`: TaskRunDetail
	fmt.Fprintf(os.Stdout, "Response from `TaskRunsAPI.TasksRunsSetOutputPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this task run. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 
**taskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRunsSetOutputPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TaskRunDetail**](TaskRunDetail.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

