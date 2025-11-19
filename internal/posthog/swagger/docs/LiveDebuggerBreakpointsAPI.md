# \LiveDebuggerBreakpointsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LiveDebuggerBreakpointsActiveRetrieve**](LiveDebuggerBreakpointsAPI.md#LiveDebuggerBreakpointsActiveRetrieve) | **Get** /api/projects/{project_id}/live_debugger_breakpoints/active/ | Get active breakpoints (External API)
[**LiveDebuggerBreakpointsBreakpointHitsRetrieve**](LiveDebuggerBreakpointsAPI.md#LiveDebuggerBreakpointsBreakpointHitsRetrieve) | **Get** /api/projects/{project_id}/live_debugger_breakpoints/breakpoint_hits/ | Get breakpoint hits
[**LiveDebuggerBreakpointsCreate**](LiveDebuggerBreakpointsAPI.md#LiveDebuggerBreakpointsCreate) | **Post** /api/projects/{project_id}/live_debugger_breakpoints/ | 
[**LiveDebuggerBreakpointsDestroy**](LiveDebuggerBreakpointsAPI.md#LiveDebuggerBreakpointsDestroy) | **Delete** /api/projects/{project_id}/live_debugger_breakpoints/{id}/ | 
[**LiveDebuggerBreakpointsList**](LiveDebuggerBreakpointsAPI.md#LiveDebuggerBreakpointsList) | **Get** /api/projects/{project_id}/live_debugger_breakpoints/ | 
[**LiveDebuggerBreakpointsPartialUpdate**](LiveDebuggerBreakpointsAPI.md#LiveDebuggerBreakpointsPartialUpdate) | **Patch** /api/projects/{project_id}/live_debugger_breakpoints/{id}/ | 
[**LiveDebuggerBreakpointsRetrieve**](LiveDebuggerBreakpointsAPI.md#LiveDebuggerBreakpointsRetrieve) | **Get** /api/projects/{project_id}/live_debugger_breakpoints/{id}/ | 
[**LiveDebuggerBreakpointsUpdate**](LiveDebuggerBreakpointsAPI.md#LiveDebuggerBreakpointsUpdate) | **Put** /api/projects/{project_id}/live_debugger_breakpoints/{id}/ | 



## LiveDebuggerBreakpointsActiveRetrieve

> ActiveBreakpointsResponse LiveDebuggerBreakpointsActiveRetrieve(ctx, projectId).Enabled(enabled).Filename(filename).Repository(repository).Execute()

Get active breakpoints (External API)



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
	enabled := true // bool | Only return enabled breakpoints (optional)
	filename := "filename_example" // string | Filter breakpoints for a specific file (optional)
	repository := "repository_example" // string | Filter breakpoints for a specific repository (e.g., 'PostHog/posthog') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsActiveRetrieve(context.Background(), projectId).Enabled(enabled).Filename(filename).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsActiveRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveDebuggerBreakpointsActiveRetrieve`: ActiveBreakpointsResponse
	fmt.Fprintf(os.Stdout, "Response from `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsActiveRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveDebuggerBreakpointsActiveRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **bool** | Only return enabled breakpoints | 
 **filename** | **string** | Filter breakpoints for a specific file | 
 **repository** | **string** | Filter breakpoints for a specific repository (e.g., &#39;PostHog/posthog&#39;) | 

### Return type

[**ActiveBreakpointsResponse**](ActiveBreakpointsResponse.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveDebuggerBreakpointsBreakpointHitsRetrieve

> BreakpointHitsResponse LiveDebuggerBreakpointsBreakpointHitsRetrieve(ctx, projectId).BreakpointIds(breakpointIds).Limit(limit).Offset(offset).Execute()

Get breakpoint hits



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
	breakpointIds := "breakpointIds_example" // string | Filter hits for specific breakpoints (repeat parameter for multiple IDs, e.g., ?breakpoint_ids=uuid1&breakpoint_ids=uuid2) (optional)
	limit := int32(56) // int32 | Number of hits to return (default: 100, max: 1000) (optional)
	offset := int32(56) // int32 | Pagination offset for retrieving additional results (default: 0) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsBreakpointHitsRetrieve(context.Background(), projectId).BreakpointIds(breakpointIds).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsBreakpointHitsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveDebuggerBreakpointsBreakpointHitsRetrieve`: BreakpointHitsResponse
	fmt.Fprintf(os.Stdout, "Response from `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsBreakpointHitsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveDebuggerBreakpointsBreakpointHitsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **breakpointIds** | **string** | Filter hits for specific breakpoints (repeat parameter for multiple IDs, e.g., ?breakpoint_ids&#x3D;uuid1&amp;breakpoint_ids&#x3D;uuid2) | 
 **limit** | **int32** | Number of hits to return (default: 100, max: 1000) | 
 **offset** | **int32** | Pagination offset for retrieving additional results (default: 0) | 

### Return type

[**BreakpointHitsResponse**](BreakpointHitsResponse.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveDebuggerBreakpointsCreate

> LiveDebuggerBreakpoint LiveDebuggerBreakpointsCreate(ctx, projectId).LiveDebuggerBreakpoint(liveDebuggerBreakpoint).Execute()





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
	liveDebuggerBreakpoint := *openapiclient.NewLiveDebuggerBreakpoint("Id_example", "Filename_example", int32(123), time.Now(), time.Now()) // LiveDebuggerBreakpoint | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsCreate(context.Background(), projectId).LiveDebuggerBreakpoint(liveDebuggerBreakpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveDebuggerBreakpointsCreate`: LiveDebuggerBreakpoint
	fmt.Fprintf(os.Stdout, "Response from `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveDebuggerBreakpointsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveDebuggerBreakpoint** | [**LiveDebuggerBreakpoint**](LiveDebuggerBreakpoint.md) |  | 

### Return type

[**LiveDebuggerBreakpoint**](LiveDebuggerBreakpoint.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveDebuggerBreakpointsDestroy

> LiveDebuggerBreakpointsDestroy(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this live debugger breakpoint.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsDestroy(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this live debugger breakpoint. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveDebuggerBreakpointsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## LiveDebuggerBreakpointsList

> PaginatedLiveDebuggerBreakpointList LiveDebuggerBreakpointsList(ctx, projectId).Filename(filename).Limit(limit).Offset(offset).Repository(repository).Execute()





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
	filename := "filename_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	repository := "repository_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsList(context.Background(), projectId).Filename(filename).Limit(limit).Offset(offset).Repository(repository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveDebuggerBreakpointsList`: PaginatedLiveDebuggerBreakpointList
	fmt.Fprintf(os.Stdout, "Response from `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveDebuggerBreakpointsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **repository** | **string** |  | 

### Return type

[**PaginatedLiveDebuggerBreakpointList**](PaginatedLiveDebuggerBreakpointList.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveDebuggerBreakpointsPartialUpdate

> LiveDebuggerBreakpoint LiveDebuggerBreakpointsPartialUpdate(ctx, id, projectId).PatchedLiveDebuggerBreakpoint(patchedLiveDebuggerBreakpoint).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this live debugger breakpoint.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	patchedLiveDebuggerBreakpoint := *openapiclient.NewPatchedLiveDebuggerBreakpoint() // PatchedLiveDebuggerBreakpoint |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsPartialUpdate(context.Background(), id, projectId).PatchedLiveDebuggerBreakpoint(patchedLiveDebuggerBreakpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveDebuggerBreakpointsPartialUpdate`: LiveDebuggerBreakpoint
	fmt.Fprintf(os.Stdout, "Response from `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this live debugger breakpoint. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveDebuggerBreakpointsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchedLiveDebuggerBreakpoint** | [**PatchedLiveDebuggerBreakpoint**](PatchedLiveDebuggerBreakpoint.md) |  | 

### Return type

[**LiveDebuggerBreakpoint**](LiveDebuggerBreakpoint.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveDebuggerBreakpointsRetrieve

> LiveDebuggerBreakpoint LiveDebuggerBreakpointsRetrieve(ctx, id, projectId).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this live debugger breakpoint.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsRetrieve(context.Background(), id, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveDebuggerBreakpointsRetrieve`: LiveDebuggerBreakpoint
	fmt.Fprintf(os.Stdout, "Response from `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this live debugger breakpoint. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveDebuggerBreakpointsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LiveDebuggerBreakpoint**](LiveDebuggerBreakpoint.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveDebuggerBreakpointsUpdate

> LiveDebuggerBreakpoint LiveDebuggerBreakpointsUpdate(ctx, id, projectId).LiveDebuggerBreakpoint(liveDebuggerBreakpoint).Execute()





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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this live debugger breakpoint.
	projectId := "projectId_example" // string | Project ID of the project you're trying to access. To find the ID of the project, make a call to /api/projects/.
	liveDebuggerBreakpoint := *openapiclient.NewLiveDebuggerBreakpoint("Id_example", "Filename_example", int32(123), time.Now(), time.Now()) // LiveDebuggerBreakpoint | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsUpdate(context.Background(), id, projectId).LiveDebuggerBreakpoint(liveDebuggerBreakpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LiveDebuggerBreakpointsUpdate`: LiveDebuggerBreakpoint
	fmt.Fprintf(os.Stdout, "Response from `LiveDebuggerBreakpointsAPI.LiveDebuggerBreakpointsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | A UUID string identifying this live debugger breakpoint. | 
**projectId** | **string** | Project ID of the project you&#39;re trying to access. To find the ID of the project, make a call to /api/projects/. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLiveDebuggerBreakpointsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **liveDebuggerBreakpoint** | [**LiveDebuggerBreakpoint**](LiveDebuggerBreakpoint.md) |  | 

### Return type

[**LiveDebuggerBreakpoint**](LiveDebuggerBreakpoint.md)

### Authorization

[PersonalAPIKeyAuth](../README.md#PersonalAPIKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

