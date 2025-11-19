# \ActionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionsCreate**](ActionsAPI.md#ActionsCreate) | **Post** /api/projects/{project_id}/actions/ | 
[**ActionsDestroy**](ActionsAPI.md#ActionsDestroy) | **Delete** /api/projects/{project_id}/actions/{id}/ | 
[**ActionsList**](ActionsAPI.md#ActionsList) | **Get** /api/projects/{project_id}/actions/ | 
[**ActionsPartialUpdate**](ActionsAPI.md#ActionsPartialUpdate) | **Patch** /api/projects/{project_id}/actions/{id}/ | 
[**ActionsRetrieve**](ActionsAPI.md#ActionsRetrieve) | **Get** /api/projects/{project_id}/actions/{id}/ | 
[**ActionsUpdate**](ActionsAPI.md#ActionsUpdate) | **Put** /api/projects/{project_id}/actions/{id}/ | 



## ActionsCreate

> Action ActionsCreate(ctx, projectId).Format(format).Action(action).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/posthog/terraform-provider"
)

func main() {
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	action := *openapiclient.NewAction(int32(123), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), false, int32(123), false, "BytecodeError_example", "CreationContext_example", "UserAccessLevel_example") // Action |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionsCreate(context.Background(), projectId).Format(format).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsCreate`: Action
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **action** | [**Action**](Action.md) |  | 

### Return type

[**Action**](Action.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsDestroy

> ActionsDestroy(ctx, id, projectId).Format(format).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this action.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ActionsAPI.ActionsDestroy(context.Background(), id, projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this action. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsList

> PaginatedActionList ActionsList(ctx, projectId).Format(format).Limit(limit).Offset(offset).Execute()



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
	format := "format_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionsList(context.Background(), projectId).Format(format).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsList`: PaginatedActionList
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedActionList**](PaginatedActionList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsPartialUpdate

> Action ActionsPartialUpdate(ctx, id, projectId).Format(format).PatchedAction(patchedAction).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this action.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	patchedAction := *openapiclient.NewPatchedAction() // PatchedAction |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionsPartialUpdate(context.Background(), id, projectId).Format(format).PatchedAction(patchedAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsPartialUpdate`: Action
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this action. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 
 **patchedAction** | [**PatchedAction**](PatchedAction.md) |  | 

### Return type

[**Action**](Action.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRetrieve

> Action ActionsRetrieve(ctx, id, projectId).Format(format).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this action.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionsRetrieve(context.Background(), id, projectId).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsRetrieve`: Action
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this action. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 

### Return type

[**Action**](Action.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsUpdate

> Action ActionsUpdate(ctx, id, projectId).Format(format).Action(action).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/posthog/terraform-provider"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this action.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	format := "format_example" // string |  (optional)
	action := *openapiclient.NewAction(int32(123), time.Now(), *openapiclient.NewUserBasic(int32(123), "Uuid_example", "Email_example", map[string]interface{}{"key": interface{}(123)}), false, int32(123), false, "BytecodeError_example", "CreationContext_example", "UserAccessLevel_example") // Action |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionsUpdate(context.Background(), id, projectId).Format(format).Action(action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsUpdate`: Action
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this action. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** |  | 
 **action** | [**Action**](Action.md) |  | 

### Return type

[**Action**](Action.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

